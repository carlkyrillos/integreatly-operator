/*
Copyright 2019 The OpenShift Authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO: these types should eventually be broken out, along with the actuator,
// to a separate repo.

// KubevirtProviderSpec the specification of the credentials request in Kubevirt.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KubevirtProviderSpec struct {
	metav1.TypeMeta `json:",inline"`
}

// KubevirtProviderSpec contains the status of the credentials request in Kubevirt.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KubevirtProviderStatus struct {
	metav1.TypeMeta `json:",inline"`
}