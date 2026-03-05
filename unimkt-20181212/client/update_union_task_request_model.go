// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *UpdateUnionTaskRequest
	GetChannelId() *string
	SetChargePloy(v int32) *UpdateUnionTaskRequest
	GetChargePloy() *int32
	SetEndTime(v string) *UpdateUnionTaskRequest
	GetEndTime() *string
	SetOptimizationSwitch(v int32) *UpdateUnionTaskRequest
	GetOptimizationSwitch() *int32
	SetProxyUserId(v int64) *UpdateUnionTaskRequest
	GetProxyUserId() *int64
	SetQuota(v int32) *UpdateUnionTaskRequest
	GetQuota() *int32
	SetQuotaDay(v int32) *UpdateUnionTaskRequest
	GetQuotaDay() *int32
	SetStartTime(v string) *UpdateUnionTaskRequest
	GetStartTime() *string
	SetTaskId(v int64) *UpdateUnionTaskRequest
	GetTaskId() *int64
}

type UpdateUnionTaskRequest struct {
	// example:
	//
	// test
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1
	ChargePloy *int32 `json:"ChargePloy,omitempty" xml:"ChargePloy,omitempty"`
	// example:
	//
	// 2024-02-28 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0
	OptimizationSwitch *int32 `json:"OptimizationSwitch,omitempty" xml:"OptimizationSwitch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// example:
	//
	// 20000
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// 7200
	QuotaDay *int32 `json:"QuotaDay,omitempty" xml:"QuotaDay,omitempty"`
	// example:
	//
	// 2024-01-30 12:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1687154040871094
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateUnionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnionTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateUnionTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateUnionTaskRequest) GetChargePloy() *int32 {
	return s.ChargePloy
}

func (s *UpdateUnionTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateUnionTaskRequest) GetOptimizationSwitch() *int32 {
	return s.OptimizationSwitch
}

func (s *UpdateUnionTaskRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *UpdateUnionTaskRequest) GetQuota() *int32 {
	return s.Quota
}

func (s *UpdateUnionTaskRequest) GetQuotaDay() *int32 {
	return s.QuotaDay
}

func (s *UpdateUnionTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateUnionTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateUnionTaskRequest) SetChannelId(v string) *UpdateUnionTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetChargePloy(v int32) *UpdateUnionTaskRequest {
	s.ChargePloy = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetEndTime(v string) *UpdateUnionTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetOptimizationSwitch(v int32) *UpdateUnionTaskRequest {
	s.OptimizationSwitch = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetProxyUserId(v int64) *UpdateUnionTaskRequest {
	s.ProxyUserId = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetQuota(v int32) *UpdateUnionTaskRequest {
	s.Quota = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetQuotaDay(v int32) *UpdateUnionTaskRequest {
	s.QuotaDay = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetStartTime(v string) *UpdateUnionTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateUnionTaskRequest) SetTaskId(v int64) *UpdateUnionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateUnionTaskRequest) Validate() error {
	return dara.Validate(s)
}
