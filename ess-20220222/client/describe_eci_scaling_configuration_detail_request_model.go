// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEciScalingConfigurationDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputFormat(v string) *DescribeEciScalingConfigurationDetailRequest
	GetOutputFormat() *string
	SetRegionId(v string) *DescribeEciScalingConfigurationDetailRequest
	GetRegionId() *string
	SetScalingConfigurationId(v string) *DescribeEciScalingConfigurationDetailRequest
	GetScalingConfigurationId() *string
	SetScalingGroupId(v string) *DescribeEciScalingConfigurationDetailRequest
	GetScalingGroupId() *string
}

type DescribeEciScalingConfigurationDetailRequest struct {
	// The output format. Set the value to YAML.
	//
	// example:
	//
	// yaml
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The region ID of the scaling group to which the scaling configuration belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling configuration based on which elastic container instances are created.
	//
	// This parameter is required.
	//
	// example:
	//
	// asc-bp1ffogfdauy0nu5****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The ID of the scaling group to which the scaling configuration belongs.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *DescribeEciScalingConfigurationDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEciScalingConfigurationDetailRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeEciScalingConfigurationDetailRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeEciScalingConfigurationDetailRequest) SetOutputFormat(v string) *DescribeEciScalingConfigurationDetailRequest {
	s.OutputFormat = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailRequest) SetRegionId(v string) *DescribeEciScalingConfigurationDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailRequest) SetScalingConfigurationId(v string) *DescribeEciScalingConfigurationDetailRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailRequest) SetScalingGroupId(v string) *DescribeEciScalingConfigurationDetailRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailRequest) Validate() error {
	return dara.Validate(s)
}
