// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProfile(v string) *DescribeRCClustersRequest
	GetProfile() *string
	SetRegionId(v string) *DescribeRCClustersRequest
	GetRegionId() *string
	SetVpcId(v string) *DescribeRCClustersRequest
	GetVpcId() *string
}

type DescribeRCClustersRequest struct {
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCClustersRequest) GetProfile() *string {
	return s.Profile
}

func (s *DescribeRCClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCClustersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCClustersRequest) SetProfile(v string) *DescribeRCClustersRequest {
	s.Profile = &v
	return s
}

func (s *DescribeRCClustersRequest) SetRegionId(v string) *DescribeRCClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCClustersRequest) SetVpcId(v string) *DescribeRCClustersRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRCClustersRequest) Validate() error {
	return dara.Validate(s)
}
