// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupDiagnoseDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeScalingGroupDiagnoseDetailsRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *DescribeScalingGroupDiagnoseDetailsRequest
	GetScalingGroupId() *string
}

type DescribeScalingGroupDiagnoseDetailsRequest struct {
	// The ID of the region to which the scaling group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingGroupDiagnoseDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDiagnoseDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDiagnoseDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingGroupDiagnoseDetailsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingGroupDiagnoseDetailsRequest) SetRegionId(v string) *DescribeScalingGroupDiagnoseDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsRequest) SetScalingGroupId(v string) *DescribeScalingGroupDiagnoseDetailsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsRequest) Validate() error {
	return dara.Validate(s)
}
