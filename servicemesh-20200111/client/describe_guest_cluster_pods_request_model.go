// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterPodsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGuestClusterID(v string) *DescribeGuestClusterPodsRequest
	GetGuestClusterID() *string
	SetNamespace(v string) *DescribeGuestClusterPodsRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *DescribeGuestClusterPodsRequest
	GetServiceMeshId() *string
}

type DescribeGuestClusterPodsRequest struct {
	// The ID of the Kubernetes cluster that is added to the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c49ad2169d5a04f2d89dfc4b6bcu****
	GuestClusterID *string `json:"GuestClusterID,omitempty" xml:"GuestClusterID,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c78d60f98fa43403ab6e0701b2678****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeGuestClusterPodsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterPodsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterPodsRequest) GetGuestClusterID() *string {
	return s.GuestClusterID
}

func (s *DescribeGuestClusterPodsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeGuestClusterPodsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeGuestClusterPodsRequest) SetGuestClusterID(v string) *DescribeGuestClusterPodsRequest {
	s.GuestClusterID = &v
	return s
}

func (s *DescribeGuestClusterPodsRequest) SetNamespace(v string) *DescribeGuestClusterPodsRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeGuestClusterPodsRequest) SetServiceMeshId(v string) *DescribeGuestClusterPodsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeGuestClusterPodsRequest) Validate() error {
	return dara.Validate(s)
}
