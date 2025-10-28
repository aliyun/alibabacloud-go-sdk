// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteK8sSecretRequest
	GetClusterId() *string
	SetName(v string) *DeleteK8sSecretRequest
	GetName() *string
	SetNamespace(v string) *DeleteK8sSecretRequest
	GetNamespace() *string
}

type DeleteK8sSecretRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2b3a1013-31c7-433b-8fe7-3895d838b824
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the Secret. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-secret
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

func (s DeleteK8sSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sSecretRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteK8sSecretRequest) GetName() *string {
	return s.Name
}

func (s *DeleteK8sSecretRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteK8sSecretRequest) SetClusterId(v string) *DeleteK8sSecretRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteK8sSecretRequest) SetName(v string) *DeleteK8sSecretRequest {
	s.Name = &v
	return s
}

func (s *DeleteK8sSecretRequest) SetNamespace(v string) *DeleteK8sSecretRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteK8sSecretRequest) Validate() error {
	return dara.Validate(s)
}
