// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoricalMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryHistoricalMetricRequest
	GetEndTime() *string
	SetMetricName(v string) *QueryHistoricalMetricRequest
	GetMetricName() *string
	SetOwnerId(v int64) *QueryHistoricalMetricRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryHistoricalMetricRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *QueryHistoricalMetricRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryHistoricalMetricRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *QueryHistoricalMetricRequest
	GetScalingGroupId() *string
	SetStartTime(v string) *QueryHistoricalMetricRequest
	GetStartTime() *string
}

type QueryHistoricalMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-18T08:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asg-bp15oubotmrq11xe****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-17T08:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryHistoricalMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoricalMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoricalMetricRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryHistoricalMetricRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *QueryHistoricalMetricRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryHistoricalMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryHistoricalMetricRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryHistoricalMetricRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryHistoricalMetricRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *QueryHistoricalMetricRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryHistoricalMetricRequest) SetEndTime(v string) *QueryHistoricalMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetMetricName(v string) *QueryHistoricalMetricRequest {
	s.MetricName = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetOwnerId(v int64) *QueryHistoricalMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetRegionId(v string) *QueryHistoricalMetricRequest {
	s.RegionId = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetResourceOwnerAccount(v string) *QueryHistoricalMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetResourceOwnerId(v int64) *QueryHistoricalMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetScalingGroupId(v string) *QueryHistoricalMetricRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *QueryHistoricalMetricRequest) SetStartTime(v string) *QueryHistoricalMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryHistoricalMetricRequest) Validate() error {
	return dara.Validate(s)
}
