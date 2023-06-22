package main

import (
	"context"

	"github.com/rancher/lasso/pkg/client"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/kubeconfig"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	kubeconfigPath := "/Users/mbelgaied/.kube/config"
	// create a new clientConfig using Wrangler's kubeconfig package and GetNonInteractiveClientConfig function
	clientConfig, err := kubeconfig.GetNonInteractiveClientConfig(kubeconfigPath).ClientConfig()

	if err != nil {
		panic(err)
	}

	//create a new clientFactory
	clientFactory, err := client.NewSharedClientFactoryForConfig(clientConfig)
	if err != nil {
		panic(err)
	}

	settingGVK := schema.GroupVersionKind{
		Group:   "management.cattle.io",
		Version: "v3",
		Kind:    "Setting",
	}

	settingsClient, err := clientFactory.ForKind(settingGVK)
	if err != nil {
		panic(err)
	}

	settingsList := &v3.SettingList{}

	err = settingsClient.List(context.TODO(), "default", settingsList, v1.ListOptions{}) // this is a list of settings
	if err != nil {
		panic(err)
	}

	for _, setting := range settingsList.Items {
		println(setting.Name)
	}

	//create a new managementClient
	// managementClient, err := v3.NewFactoryFromConfigWithOptions(clientConfig, clientFactory, &v3.FactoryOptions{})
	// if err != nil {
	// 	panic(err)
	// }

}
