// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClusterNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeContainerServiceK8sClusterNamespacesRequest
	GetClusterId() *string
	SetResourceOwnerId(v int64) *DescribeContainerServiceK8sClusterNamespacesRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeContainerServiceK8sClusterNamespacesRequest
	GetSourceIp() *string
}

type DescribeContainerServiceK8sClusterNamespacesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cf4435fefd45d4b1b8643f3a0bea3****
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 140.205.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeContainerServiceK8sClusterNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) SetClusterId(v string) *DescribeContainerServiceK8sClusterNamespacesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) SetResourceOwnerId(v int64) *DescribeContainerServiceK8sClusterNamespacesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) SetSourceIp(v string) *DescribeContainerServiceK8sClusterNamespacesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
