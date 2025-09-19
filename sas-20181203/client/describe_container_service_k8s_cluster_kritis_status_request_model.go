// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClusterKritisStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeContainerServiceK8sClusterKritisStatusRequest
	GetClusterId() *string
	SetResourceOwnerId(v int64) *DescribeContainerServiceK8sClusterKritisStatusRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeContainerServiceK8sClusterKritisStatusRequest
	GetSourceIp() *string
}

type DescribeContainerServiceK8sClusterKritisStatusRequest struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c0e9efc6dea5f41db93b7e977123c****
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 119.145.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeContainerServiceK8sClusterKritisStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterKritisStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) SetClusterId(v string) *DescribeContainerServiceK8sClusterKritisStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) SetResourceOwnerId(v int64) *DescribeContainerServiceK8sClusterKritisStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) SetSourceIp(v string) *DescribeContainerServiceK8sClusterKritisStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusRequest) Validate() error {
	return dara.Validate(s)
}
