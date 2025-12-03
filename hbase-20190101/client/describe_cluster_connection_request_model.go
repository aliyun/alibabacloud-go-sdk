// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterConnectionRequest
	GetClusterId() *string
	SetRegionId(v string) *DescribeClusterConnectionRequest
	GetRegionId() *string
}

type DescribeClusterConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeClusterConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterConnectionRequest) SetClusterId(v string) *DescribeClusterConnectionRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterConnectionRequest) SetRegionId(v string) *DescribeClusterConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterConnectionRequest) Validate() error {
	return dara.Validate(s)
}
