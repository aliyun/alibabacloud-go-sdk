// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceOwnerId(v int64) *DescribeContainerServiceK8sClustersRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeContainerServiceK8sClustersRequest
	GetSourceIp() *string
}

type DescribeContainerServiceK8sClustersRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 42.120.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeContainerServiceK8sClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeContainerServiceK8sClustersRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeContainerServiceK8sClustersRequest) SetResourceOwnerId(v int64) *DescribeContainerServiceK8sClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerServiceK8sClustersRequest) SetSourceIp(v string) *DescribeContainerServiceK8sClustersRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeContainerServiceK8sClustersRequest) Validate() error {
	return dara.Validate(s)
}
