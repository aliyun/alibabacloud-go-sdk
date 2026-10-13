// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncJobScheduleResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncJobScheduleResponseBodyData) *GetContextDatabaseSyncJobScheduleResponseBody
	GetData() *GetContextDatabaseSyncJobScheduleResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncJobScheduleResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncJobScheduleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncJobScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncJobScheduleResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncJobScheduleResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncJobScheduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Example value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContextDatabaseSyncJobScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) GetData() *GetContextDatabaseSyncJobScheduleResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) SetCode(v string) *GetContextDatabaseSyncJobScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) SetData(v *GetContextDatabaseSyncJobScheduleResponseBodyData) *GetContextDatabaseSyncJobScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncJobScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) SetMessage(v string) *GetContextDatabaseSyncJobScheduleResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) SetRequestId(v string) *GetContextDatabaseSyncJobScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncJobScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncJobScheduleResponseBodyData struct {
	// The reason for the automatic pause.
	//
	// example:
	//
	// Three consecutive execution failures
	AutoPauseReason *string `json:"AutoPauseReason,omitempty" xml:"AutoPauseReason,omitempty"`
	// Indicates whether the task is automatically paused.
	//
	// example:
	//
	// true
	AutoPaused *bool `json:"AutoPaused,omitempty" xml:"AutoPaused,omitempty"`
	// The list of daily trigger times for daily-times scheduling in the HH:mm 24-hour format. The number of times ranges from 1 to 6.
	//
	// example:
	//
	// ["09:00", "18:00"]
	DailyTimes []*string `json:"DailyTimes,omitempty" xml:"DailyTimes,omitempty" type:"Repeated"`
	// The interval in minutes for interval scheduling. The minimum value is 10.
	//
	// example:
	//
	// 12
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
	// The processing policy for missed schedules. Valid values: skip and catch-up.
	//
	// example:
	//
	// skip
	MissedPolicy *string `json:"MissedPolicy,omitempty" xml:"MissedPolicy,omitempty"`
	// The next run time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	NextRunAt *string `json:"NextRunAt,omitempty" xml:"NextRunAt,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The scheduling type. Valid values: interval and daily-times.
	//
	// example:
	//
	// daily-times
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetContextDatabaseSyncJobScheduleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetAutoPauseReason() *string {
	return s.AutoPauseReason
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetAutoPaused() *bool {
	return s.AutoPaused
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetDailyTimes() []*string {
	return s.DailyTimes
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetMissedPolicy() *string {
	return s.MissedPolicy
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetAutoPauseReason(v string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.AutoPauseReason = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetAutoPaused(v bool) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.AutoPaused = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetDailyTimes(v []*string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.DailyTimes = v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetIntervalMinutes(v int32) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.IntervalMinutes = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetMissedPolicy(v string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.MissedPolicy = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetNextRunAt(v string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.NextRunAt = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetTimezone(v string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetType(v string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) SetUpdatedAt(v string) *GetContextDatabaseSyncJobScheduleResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
