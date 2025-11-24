// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListServiceAccountsRequest
	GetClusterId() *string
	SetNamespace(v string) *ListServiceAccountsRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *ListServiceAccountsRequest
	GetServiceMeshId() *string
}

type ListServiceAccountsRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce134b0727aa2492db69f6c3880e1****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ListServiceAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceAccountsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListServiceAccountsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListServiceAccountsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *ListServiceAccountsRequest) SetClusterId(v string) *ListServiceAccountsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListServiceAccountsRequest) SetNamespace(v string) *ListServiceAccountsRequest {
	s.Namespace = &v
	return s
}

func (s *ListServiceAccountsRequest) SetServiceMeshId(v string) *ListServiceAccountsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *ListServiceAccountsRequest) Validate() error {
	return dara.Validate(s)
}
