/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	// "os/exec"
	// "os/exec"

	"github.com/spf13/cobra"
	// configv1 "github.com/openshift/api/config/v1"
	// routev1 "github.com/openshift/api/route/v1"
	// appsv1 "k8s.io/api/apps/v1"
	// batchv1 "k8s.io/api/batch/v1"
	// corev1 "k8s.io/api/core/v1"
	// velerov1 "github.com/heptio/velero/pkg/apis/velero/v1"
	// oappsv1 "github.com/openshift/api/apps/v1"
	// orbacv1 "github.com/openshift/api/authorization/v1"
	// oconfigv1 "github.com/openshift/api/config/v1"
	// machineapi "github.com/openshift/api/machine/v1beta1"
	// ingresscontroller "github.com/openshift/api/operator/v1"
	// autoscalingv1 "github.com/openshift/cluster-autoscaler-operator/pkg/apis/autoscaling/v1"
	// autoscalingv1beta1 "github.com/openshift/cluster-autoscaler-operator/pkg/apis/autoscaling/v1beta1"
	// "github.com/openshift/hive/apis"
	// monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	// admissionv1beta1 "k8s.io/api/admission/v1beta1"
	// rbacv1 "k8s.io/api/rbac/v1"
	// apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	// crv1alpha1 "k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
	// apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	// openstackprovider "sigs.k8s.io/cluster-api-provider-openstack/pkg/apis"

	// ovirtprovider "github.com/openshift/cluster-api-provider-ovirt/pkg/apis"
	// hivev1 "github.com/openshift/hive/apis/hive/v1"
	// hivecontractsv1alpha1 "github.com/openshift/hive/apis/hivecontracts/v1alpha1"
	// hiveintv1alpha1 "github.com/openshift/hive/apis/hiveinternal/v1alpha1"
	// admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	// "k8s.io/apimachinery/pkg/runtime"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oc-snapshot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}


type ApiResourceKnownNamespaces struct {
	apiResource     SwaggerDocable
	allNamespaces   bool
	knownNamespaces []string
}

// var ApiResourcesKnownNamespaces = map[string]string{
// }

var (
	namespace string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	// fetch namespace
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oc-diffpp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&namespace, "--namespace", "-n", "", "Diff a namespace.")
}
