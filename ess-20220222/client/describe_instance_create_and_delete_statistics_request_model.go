// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceCreateAndDeleteStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetScalingGroupId() *string
	SetStartTime(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest
	GetStartTime() *string
}

type DescribeInstanceCreateAndDeleteStatisticsRequest struct {
	// The end time of the statistical interval. The time follows the ISO 8601 standard and uses UTC time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-18T08:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-2ze4057nqfbxxxxxxxx
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The start time of the statistical interval. The time follows the ISO 8601 standard and uses UTC time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-15T08:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceCreateAndDeleteStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCreateAndDeleteStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetEndTime(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetOwnerId(v int64) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetRegionId(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetResourceOwnerId(v int64) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetScalingGroupId(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) SetStartTime(v string) *DescribeInstanceCreateAndDeleteStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceCreateAndDeleteStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
