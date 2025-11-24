// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceScopeSidecarConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DescribeNamespaceScopeSidecarConfigRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *DescribeNamespaceScopeSidecarConfigRequest
	GetServiceMeshId() *string
}

type DescribeNamespaceScopeSidecarConfigRequest struct {
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c7120e75a202d4fd8acb028a86b6a****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) SetNamespace(v string) *DescribeNamespaceScopeSidecarConfigRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) SetServiceMeshId(v string) *DescribeNamespaceScopeSidecarConfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) Validate() error {
	return dara.Validate(s)
}
