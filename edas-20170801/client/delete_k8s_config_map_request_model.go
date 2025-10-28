// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sConfigMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteK8sConfigMapRequest
	GetClusterId() *string
	SetName(v string) *DeleteK8sConfigMapRequest
	GetName() *string
	SetNamespace(v string) *DeleteK8sConfigMapRequest
	GetNamespace() *string
}

type DeleteK8sConfigMapRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51e37e91-6403-423a-9913-f236aa7a6a50
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the ConfigMap. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-configmap
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteK8sConfigMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sConfigMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sConfigMapRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteK8sConfigMapRequest) GetName() *string {
	return s.Name
}

func (s *DeleteK8sConfigMapRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteK8sConfigMapRequest) SetClusterId(v string) *DeleteK8sConfigMapRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteK8sConfigMapRequest) SetName(v string) *DeleteK8sConfigMapRequest {
	s.Name = &v
	return s
}

func (s *DeleteK8sConfigMapRequest) SetNamespace(v string) *DeleteK8sConfigMapRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteK8sConfigMapRequest) Validate() error {
	return dara.Validate(s)
}
