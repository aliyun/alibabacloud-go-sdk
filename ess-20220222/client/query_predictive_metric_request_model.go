// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryPredictiveMetricRequest
	GetEndTime() *string
	SetMetricName(v string) *QueryPredictiveMetricRequest
	GetMetricName() *string
	SetOwnerId(v int64) *QueryPredictiveMetricRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryPredictiveMetricRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *QueryPredictiveMetricRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPredictiveMetricRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *QueryPredictiveMetricRequest
	GetScalingGroupId() *string
	SetStartTime(v string) *QueryPredictiveMetricRequest
	GetStartTime() *string
}

type QueryPredictiveMetricRequest struct {
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
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-17T08:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPredictiveMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryPredictiveMetricRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPredictiveMetricRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *QueryPredictiveMetricRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPredictiveMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryPredictiveMetricRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPredictiveMetricRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPredictiveMetricRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *QueryPredictiveMetricRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPredictiveMetricRequest) SetEndTime(v string) *QueryPredictiveMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetMetricName(v string) *QueryPredictiveMetricRequest {
	s.MetricName = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetOwnerId(v int64) *QueryPredictiveMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetRegionId(v string) *QueryPredictiveMetricRequest {
	s.RegionId = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetResourceOwnerAccount(v string) *QueryPredictiveMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetResourceOwnerId(v int64) *QueryPredictiveMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetScalingGroupId(v string) *QueryPredictiveMetricRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *QueryPredictiveMetricRequest) SetStartTime(v string) *QueryPredictiveMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPredictiveMetricRequest) Validate() error {
	return dara.Validate(s)
}
