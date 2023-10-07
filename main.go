package main

import (
	// "context"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getKubeConfigPath retrieves the path to the kubeconfig file.
func getKubeConfigPath() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".kube", "config")
}

// getKubeClient gets a Kubernetes client using the local kubeconfig.
func getKubeClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", getKubeConfigPath())
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// FetchResourcesForAPIGroup retrieves resources for a specific API group and namespace.
func FetchResourcesForAPIGroup(client *kubernetes.Clientset, namespace string, apiGroup string) ([]string, error) {
	resourceList, err := client.Discovery().ServerResourcesForGroupVersion(apiGroup)
	if err != nil {
		return nil, err
	}

	var resources []string
	for _, apiResource := range resourceList.APIResources {
		resources = append(resources, apiResource.Name)
	}
	return resources, nil
}

// GetResourcesMap retrieves a hierarchical map of API groups to resources.
func GetResourcesMap(namespace string, apiGroups []string) (map[string][]string, error) {
	client, err := getKubeClient()
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, apiGroup := range apiGroups {
		resources, err := FetchResourcesForAPIGroup(client, namespace, apiGroup)
		if err != nil {
			return nil, err
		}
		result[apiGroup] = resources
	}
	return result, nil
}

func main() {
	apiGroups := []string{"apps/v1", "v1"}
	namespace := "default"

	resourcesMap, err := GetResourcesMap(namespace, apiGroups)
	if err != nil {
		fmt.Println("Error fetching resources:", err)
		return
	}

	for group, resources := range resourcesMap {
		fmt.Println(group)
		for _, resource := range resources {
			fmt.Printf("  - %s\n", resource)
		}
	}
}
