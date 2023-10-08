package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfigPath() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".kube", "config")
}

func getKubeClients() (*kubernetes.Clientset, dynamic.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", getKubeConfigPath())
	if err != nil {
		return nil, nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	return clientset, dynamicClient, nil
}

func DiscoverAPIGroups(client *kubernetes.Clientset) ([]string, error) {
	groups, err := client.Discovery().ServerGroups()
	if err != nil {
		return nil, err
	}

	var apiGroups []string
	for _, group := range groups.Groups {
		for _, version := range group.Versions {
			apiGroups = append(apiGroups, version.GroupVersion)
		}
	}
	return apiGroups, nil
}

func FetchObjectsForAPIGroup(client *kubernetes.Clientset, dynamicClient dynamic.Interface, namespace string, apiGroup string) ([]string, error) {
	fmt.Println(apiGroup)
	resourceList, err := client.Discovery().ServerResourcesForGroupVersion(apiGroup)
	fmt.Println(resourceList)
	if err != nil {
		return nil, err
	}

	var objects []string
	for _, apiResource := range resourceList.APIResources {
		gvr := schema.GroupVersionResource{
			Group:    apiResource.Group,
			Version:  apiResource.Version,
			Resource: apiResource.Name,
		}

		fmt.Println(gvr)
		objs, err := dynamicClient.Resource(gvr).Namespace(namespace).List(context.TODO(), v1.ListOptions{})
		fmt.Println(objs)
		if err != nil {
			return nil, err
		}
		for _, obj := range objs.Items {
			objects = append(objects, obj.GetName())
		}
	}
	return objects, nil
}

func GetResourcesMap(namespace string, apiGroups []string, client *kubernetes.Clientset, dynamicClient dynamic.Interface) (map[string][]string, error) {
	fmt.Println(apiGroups)
	result := make(map[string][]string)
	for _, apiGroup := range apiGroups {
		objects, err := FetchObjectsForAPIGroup(client, dynamicClient, namespace, apiGroup)
		if err != nil {
			return nil, err
		}
		result[apiGroup] = objects
	}
	return result, nil
}

func main() {
	client, dynamicClient, err := getKubeClients()
	if err != nil {
		fmt.Println("Error fetching client:", err)
		return
	}

	apiGroups, err := DiscoverAPIGroups(client)
	if err != nil {
		fmt.Println("Error discovering API groups:", err)
		return
	}

	namespace := "default"
	resourcesMap, err := GetResourcesMap(namespace, apiGroups, client, dynamicClient)
	if err != nil {
		fmt.Println("Error fetching resources:", err)
		return
	}

	for group, objects := range resourcesMap {
		fmt.Println(group)
		for _, object := range objects {
			fmt.Printf("  - %s\n", object)
		}
	}
}
