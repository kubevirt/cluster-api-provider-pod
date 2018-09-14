/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package google

import (
	"fmt"
	"github.com/golang/glog"

	providerconfigv1 "github.com/kubevirt/cluster-api-provider-external/cloud/external/providerconfig/v1alpha1"

	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	client "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset/typed/cluster/v1alpha1"
)

type ExtClusterClient struct {
	clusterClient       client.ClusterInterface
	providerConfigCodec *providerconfigv1.ExtProviderConfigCodec
}

type ClusterActuatorParams struct {
	ClusterClient client.ClusterInterface
}

func NewClusterActuator(params ClusterActuatorParams) (*ExtClusterClient, error) {
	codec, err := provicderconfigv1.NewCodec()
	if err != nil {
		return nil, err
	}

	return &ExtClusterClient{
		clusterClient:       params.ClusterClient,
		providerConfigCodec: codec,
	}, nil
}

func (gce *ExtClusterClient) Reconcile(cluster *clusterv1.Cluster) error {
	glog.Infof("Reconciling cluster %v.", cluster.Name)
	return nil
}

func (gce *ExtClusterClient) Delete(cluster *clusterv1.Cluster) error {
	glog.Infof("Deleting cluster %v.", cluster.Name)
	return nil
}
