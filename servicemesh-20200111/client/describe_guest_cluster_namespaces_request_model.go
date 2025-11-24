// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGuestClusterID(v string) *DescribeGuestClusterNamespacesRequest
	GetGuestClusterID() *string
	SetServiceMeshId(v string) *DescribeGuestClusterNamespacesRequest
	GetServiceMeshId() *string
	SetShowNsLabels(v bool) *DescribeGuestClusterNamespacesRequest
	GetShowNsLabels() *bool
}

type DescribeGuestClusterNamespacesRequest struct {
	// The ID of the Kubernetes cluster that is added to the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c584d805c7bd442b8bac421f9849f****
	GuestClusterID *string `json:"GuestClusterID,omitempty" xml:"GuestClusterID,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce134b0727aa2492db69f6c3880e****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// Specifies whether to return the labels of the namespaces.
	//
	// example:
	//
	// true
	ShowNsLabels *bool `json:"ShowNsLabels,omitempty" xml:"ShowNsLabels,omitempty"`
}

func (s DescribeGuestClusterNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterNamespacesRequest) GetGuestClusterID() *string {
	return s.GuestClusterID
}

func (s *DescribeGuestClusterNamespacesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeGuestClusterNamespacesRequest) GetShowNsLabels() *bool {
	return s.ShowNsLabels
}

func (s *DescribeGuestClusterNamespacesRequest) SetGuestClusterID(v string) *DescribeGuestClusterNamespacesRequest {
	s.GuestClusterID = &v
	return s
}

func (s *DescribeGuestClusterNamespacesRequest) SetServiceMeshId(v string) *DescribeGuestClusterNamespacesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeGuestClusterNamespacesRequest) SetShowNsLabels(v bool) *DescribeGuestClusterNamespacesRequest {
	s.ShowNsLabels = &v
	return s
}

func (s *DescribeGuestClusterNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
