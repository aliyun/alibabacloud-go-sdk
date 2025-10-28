// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateK8sConfigMapRequest
	GetClusterId() *string
	SetData(v map[string]interface{}) *UpdateK8sConfigMapRequest
	GetData() map[string]interface{}
	SetName(v string) *UpdateK8sConfigMapRequest
	GetName() *string
	SetNamespace(v string) *UpdateK8sConfigMapRequest
	GetNamespace() *string
}

type UpdateK8sConfigMapRequest struct {
	// The ID of the Kubernetes cluster.
	//
	// example:
	//
	// 7246cxxx-53xx-xxxx-xxxx-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data of the ConfigMap. Set this parameter to a JSON string.
	//
	// example:
	//
	// [{"Key":"name","Value":"william"},{"Key":"age","Value":"12"}]
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The name of the ConfigMap. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 63 characters in length.
	//
	// example:
	//
	// my-configmap
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s UpdateK8sConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sConfigMapRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sConfigMapRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateK8sConfigMapRequest) GetData() map[string]interface{} {
	return s.Data
}

func (s *UpdateK8sConfigMapRequest) GetName() *string {
	return s.Name
}

func (s *UpdateK8sConfigMapRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateK8sConfigMapRequest) SetClusterId(v string) *UpdateK8sConfigMapRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sConfigMapRequest) SetData(v map[string]interface{}) *UpdateK8sConfigMapRequest {
	s.Data = v
	return s
}

func (s *UpdateK8sConfigMapRequest) SetName(v string) *UpdateK8sConfigMapRequest {
	s.Name = &v
	return s
}

func (s *UpdateK8sConfigMapRequest) SetNamespace(v string) *UpdateK8sConfigMapRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateK8sConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
