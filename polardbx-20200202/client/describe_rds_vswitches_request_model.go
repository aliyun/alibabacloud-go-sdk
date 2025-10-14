// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVswitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRdsVswitchesRequest
	GetRegionId() *string
	SetVpcId(v string) *DescribeRdsVswitchesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeRdsVswitchesRequest
	GetZoneId() *string
}

type DescribeRdsVswitchesRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vpc-bp1ndou****twoedlmru0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVswitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVswitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVswitchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRdsVswitchesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRdsVswitchesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRdsVswitchesRequest) SetRegionId(v string) *DescribeRdsVswitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVswitchesRequest) SetVpcId(v string) *DescribeRdsVswitchesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVswitchesRequest) SetZoneId(v string) *DescribeRdsVswitchesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRdsVswitchesRequest) Validate() error {
	return dara.Validate(s)
}
