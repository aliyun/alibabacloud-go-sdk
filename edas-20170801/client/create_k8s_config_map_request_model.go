// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateK8sConfigMapRequest
	GetClusterId() *string
	SetData(v map[string]interface{}) *CreateK8sConfigMapRequest
	GetData() map[string]interface{}
	SetName(v string) *CreateK8sConfigMapRequest
	GetName() *string
	SetNamespace(v string) *CreateK8sConfigMapRequest
	GetNamespace() *string
}

type CreateK8sConfigMapRequest struct {
	// The ID of the Kubernetes cluster.
	//
	// example:
	//
	// b07c8192-****-adf4f7447720
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data of the ConfigMap. The value must be a JSON array string.
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

func (s CreateK8sConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sConfigMapRequest) GoString() string {
	return s.String()
}

func (s *CreateK8sConfigMapRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateK8sConfigMapRequest) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateK8sConfigMapRequest) GetName() *string {
	return s.Name
}

func (s *CreateK8sConfigMapRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateK8sConfigMapRequest) SetClusterId(v string) *CreateK8sConfigMapRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateK8sConfigMapRequest) SetData(v map[string]interface{}) *CreateK8sConfigMapRequest {
	s.Data = v
	return s
}

func (s *CreateK8sConfigMapRequest) SetName(v string) *CreateK8sConfigMapRequest {
	s.Name = &v
	return s
}

func (s *CreateK8sConfigMapRequest) SetNamespace(v string) *CreateK8sConfigMapRequest {
	s.Namespace = &v
	return s
}

func (s *CreateK8sConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
