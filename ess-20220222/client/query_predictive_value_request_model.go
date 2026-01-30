// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryPredictiveValueRequest
	GetEndTime() *string
	SetMetricName(v string) *QueryPredictiveValueRequest
	GetMetricName() *string
	SetOwnerId(v int64) *QueryPredictiveValueRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryPredictiveValueRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *QueryPredictiveValueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPredictiveValueRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *QueryPredictiveValueRequest
	GetScalingGroupId() *string
	SetStartTime(v string) *QueryPredictiveValueRequest
	GetStartTime() *string
	SetTargetValue(v float32) *QueryPredictiveValueRequest
	GetTargetValue() *float32
}

type QueryPredictiveValueRequest struct {
	// The end time of the predicted number of instances in the scaling group. The time follows the ISO8601 standard and uses UTC time.
	//
	// Format: yyyy-MM-ddTHH:mmZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-18T08:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the monitoring metric of the prediction rule. Valid values:
	//
	// 	- CpuUtilization: the average CPU utilization.
	//
	// 	- IntranetRx: the inbound traffic over an internal network.
	//
	// 	- IntranetTx: the outbound traffic over an internal network.
	//
	// This parameter is required.
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// asg-2zee0wsy0zkdxxxxx
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The start time of the predicted number of instances in the scaling group. The time follows the ISO8601 standard and uses UTC time.
	//
	// Format: yyyy-MM-ddTHH:mmZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-17T08:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The current value of the monitoring metric of the prediction rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	TargetValue *float32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s QueryPredictiveValueRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveValueRequest) GoString() string {
	return s.String()
}

func (s *QueryPredictiveValueRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPredictiveValueRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *QueryPredictiveValueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPredictiveValueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryPredictiveValueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPredictiveValueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPredictiveValueRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *QueryPredictiveValueRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPredictiveValueRequest) GetTargetValue() *float32 {
	return s.TargetValue
}

func (s *QueryPredictiveValueRequest) SetEndTime(v string) *QueryPredictiveValueRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetMetricName(v string) *QueryPredictiveValueRequest {
	s.MetricName = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetOwnerId(v int64) *QueryPredictiveValueRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetRegionId(v string) *QueryPredictiveValueRequest {
	s.RegionId = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetResourceOwnerAccount(v string) *QueryPredictiveValueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetResourceOwnerId(v int64) *QueryPredictiveValueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetScalingGroupId(v string) *QueryPredictiveValueRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetStartTime(v string) *QueryPredictiveValueRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPredictiveValueRequest) SetTargetValue(v float32) *QueryPredictiveValueRequest {
	s.TargetValue = &v
	return s
}

func (s *QueryPredictiveValueRequest) Validate() error {
	return dara.Validate(s)
}
