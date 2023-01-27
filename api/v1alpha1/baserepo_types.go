/*
Copyright 2023 microsoft.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BaseRepoSpec defines the desired state of BaseRepo
type BaseRepoSpec struct {
	ManifestsSpec `json:",inline"`

	//+optional
	Commit string `json:"commit,omitempty"`
}

// BaseRepoStatus defines the observed state of BaseRepo
type BaseRepoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// BaseRepo is the Schema for the baserepoes API
type BaseRepo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseRepoSpec   `json:"spec,omitempty"`
	Status BaseRepoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BaseRepoList contains a list of BaseRepo
type BaseRepoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseRepo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseRepo{}, &BaseRepoList{})
}
