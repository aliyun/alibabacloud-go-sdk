// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivityStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeScalingActivityStatisticsRequest
	GetEndTime() *string
	SetMetricType(v string) *DescribeScalingActivityStatisticsRequest
	GetMetricType() *string
	SetOwnerId(v int64) *DescribeScalingActivityStatisticsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeScalingActivityStatisticsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScalingActivityStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingActivityStatisticsRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DescribeScalingActivityStatisticsRequest
	GetScalingGroupId() *string
	SetStartTime(v string) *DescribeScalingActivityStatisticsRequest
	GetStartTime() *string
}

type DescribeScalingActivityStatisticsRequest struct {
	// The end time of the statistical interval. The time follows the ISO 8601 standard and uses UTC time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-18T08:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the metric on which the scaling activity is counted. Valid values:
	//
	// 	- ScalingActivityStatus: collects statistics on the distribution of scaling activities in different states within a time range.
	//
	// 	- ScalingActivityErrorCodes: the distribution of error codes in failed scaling activities within a time range.
	//
	// Default value: ScalingActivityStatus.
	//
	// example:
	//
	// ScalingActivityStatus
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
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
	// asg-8vbje5pofxxxxxxxx
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

func (s DescribeScalingActivityStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeScalingActivityStatisticsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeScalingActivityStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingActivityStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingActivityStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingActivityStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingActivityStatisticsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingActivityStatisticsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeScalingActivityStatisticsRequest) SetEndTime(v string) *DescribeScalingActivityStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetMetricType(v string) *DescribeScalingActivityStatisticsRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetOwnerId(v int64) *DescribeScalingActivityStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetRegionId(v string) *DescribeScalingActivityStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivityStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetResourceOwnerId(v int64) *DescribeScalingActivityStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetScalingGroupId(v string) *DescribeScalingActivityStatisticsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) SetStartTime(v string) *DescribeScalingActivityStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeScalingActivityStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
