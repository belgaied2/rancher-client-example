package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rancher/lasso/pkg/client"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/kubeconfig"
	"github.com/rancher/wrangler/pkg/schemes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	// define the path to the kubeconfig file
	kubeconfigPath := os.Getenv("HOME") + "/.kube/config"
	if os.Getenv("KUBECONFIG") != "" {
		kubeconfigPath = os.Getenv("KUBECONFIG")
	}

	// create a new clientConfig using Wrangler's kubeconfig package and GetNonInteractiveClientConfig function
	clientConfig, err := kubeconfig.GetNonInteractiveClientConfig(kubeconfigPath).ClientConfig()

	if err != nil {
		panic(err)
	}

	//create a new clientFactory, this is used to create clients for different resources
	// Since each Resource type needs its own client, the clientFactory is used to create the client for each resource.
	// However, the clientFactory is able to return the same http.Client for each client it creates, meaning that the same http.Client is used for all clients.
	clientFactory, err := client.NewSharedClientFactoryForConfig(clientConfig)
	if err != nil {
		panic(err)
	}

	schemes.Register(v3.AddToScheme)
	// settingGVK := schema.GroupVersionKind{
	// 	Group:   "management.cattle.io",
	// 	Version: "v3",
	// 	Kind:    "Setting",
	// }

	settingGVK, err := clientFactory.GVKForObject(&v3.Setting{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", settingGVK)

	// create a new client for the Setting resource
	settingsClient, err := clientFactory.ForKind(settingGVK)
	if err != nil {
		panic(err)
	}

	// instantiate a new empty SettingList
	settingsList := &v3.SettingList{}

	// list all settings in the default namespace using the settingsClient and store the results in the settingsList
	err = settingsClient.List(context.TODO(), "default", settingsList, v1.ListOptions{}) // this is a list of settings
	if err != nil {
		panic(err)
	}

	// print the name of each setting in the settingsList
	for _, setting := range settingsList.Items {
		println(setting.Name)
	}

}
