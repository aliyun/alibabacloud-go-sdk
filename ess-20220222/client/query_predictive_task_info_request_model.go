// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryPredictiveTaskInfoRequest
	GetEndTime() *string
	SetOwnerId(v int64) *QueryPredictiveTaskInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryPredictiveTaskInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *QueryPredictiveTaskInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPredictiveTaskInfoRequest
	GetResourceOwnerId() *int64
	SetScalingRuleId(v string) *QueryPredictiveTaskInfoRequest
	GetScalingRuleId() *string
	SetStartTime(v string) *QueryPredictiveTaskInfoRequest
	GetStartTime() *string
}

type QueryPredictiveTaskInfoRequest struct {
	// The end time of the prediction task. The time follows the ISO8601 standard and uses UTC time.
	//
	// Format: yyyy-MM-ddTHH:mmZ.
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
	// The ID of the prediction scaling rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// asr-bp1dp8yd9f8xxxxx
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	// The start time of the prediction task. The time follows the ISO8601 standard and uses UTC time.
	//
	// Format: yyyy-MM-ddTHH:mmZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-17T08:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryPredictiveTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPredictiveTaskInfoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPredictiveTaskInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPredictiveTaskInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryPredictiveTaskInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPredictiveTaskInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPredictiveTaskInfoRequest) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *QueryPredictiveTaskInfoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPredictiveTaskInfoRequest) SetEndTime(v string) *QueryPredictiveTaskInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) SetOwnerId(v int64) *QueryPredictiveTaskInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) SetRegionId(v string) *QueryPredictiveTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) SetResourceOwnerAccount(v string) *QueryPredictiveTaskInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) SetResourceOwnerId(v int64) *QueryPredictiveTaskInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) SetScalingRuleId(v string) *QueryPredictiveTaskInfoRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) SetStartTime(v string) *QueryPredictiveTaskInfoRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPredictiveTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
