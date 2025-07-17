// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputFormat(v string) *DescribeScalingGroupDetailRequest
	GetOutputFormat() *string
	SetOwnerId(v int64) *DescribeScalingGroupDetailRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeScalingGroupDetailRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *DescribeScalingGroupDetailRequest
	GetScalingGroupId() *string
}

type DescribeScalingGroupDetailRequest struct {
	// The output format. Set the value to yaml.
	//
	// example:
	//
	// yaml
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group. For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s DescribeScalingGroupDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *DescribeScalingGroupDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingGroupDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingGroupDetailRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingGroupDetailRequest) SetOutputFormat(v string) *DescribeScalingGroupDetailRequest {
	s.OutputFormat = &v
	return s
}

func (s *DescribeScalingGroupDetailRequest) SetOwnerId(v int64) *DescribeScalingGroupDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingGroupDetailRequest) SetRegionId(v string) *DescribeScalingGroupDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupDetailRequest) SetScalingGroupId(v string) *DescribeScalingGroupDetailRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailRequest) Validate() error {
	return dara.Validate(s)
}
