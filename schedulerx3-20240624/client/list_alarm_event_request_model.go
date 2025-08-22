// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmChannel(v string) *ListAlarmEventRequest
	GetAlarmChannel() *string
	SetAlarmStatus(v string) *ListAlarmEventRequest
	GetAlarmStatus() *string
	SetAlarmType(v string) *ListAlarmEventRequest
	GetAlarmType() *string
	SetAppName(v string) *ListAlarmEventRequest
	GetAppName() *string
	SetClusterId(v string) *ListAlarmEventRequest
	GetClusterId() *string
	SetEndTime(v int64) *ListAlarmEventRequest
	GetEndTime() *int64
	SetJobName(v string) *ListAlarmEventRequest
	GetJobName() *string
	SetPageNum(v string) *ListAlarmEventRequest
	GetPageNum() *string
	SetPageSize(v string) *ListAlarmEventRequest
	GetPageSize() *string
	SetReverse(v bool) *ListAlarmEventRequest
	GetReverse() *bool
	SetStartTime(v int64) *ListAlarmEventRequest
	GetStartTime() *int64
}

type ListAlarmEventRequest struct {
	// example:
	//
	// webhook
	AlarmChannel *string `json:"AlarmChannel,omitempty" xml:"AlarmChannel,omitempty"`
	// example:
	//
	// true
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// example:
	//
	// schedulerx3_fail_alarm
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1731636011558
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// example:
	//
	// 1690419316000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmEventRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmEventRequest) GetAlarmChannel() *string {
	return s.AlarmChannel
}

func (s *ListAlarmEventRequest) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *ListAlarmEventRequest) GetAlarmType() *string {
	return s.AlarmType
}

func (s *ListAlarmEventRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAlarmEventRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAlarmEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAlarmEventRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListAlarmEventRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *ListAlarmEventRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAlarmEventRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListAlarmEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAlarmEventRequest) SetAlarmChannel(v string) *ListAlarmEventRequest {
	s.AlarmChannel = &v
	return s
}

func (s *ListAlarmEventRequest) SetAlarmStatus(v string) *ListAlarmEventRequest {
	s.AlarmStatus = &v
	return s
}

func (s *ListAlarmEventRequest) SetAlarmType(v string) *ListAlarmEventRequest {
	s.AlarmType = &v
	return s
}

func (s *ListAlarmEventRequest) SetAppName(v string) *ListAlarmEventRequest {
	s.AppName = &v
	return s
}

func (s *ListAlarmEventRequest) SetClusterId(v string) *ListAlarmEventRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAlarmEventRequest) SetEndTime(v int64) *ListAlarmEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmEventRequest) SetJobName(v string) *ListAlarmEventRequest {
	s.JobName = &v
	return s
}

func (s *ListAlarmEventRequest) SetPageNum(v string) *ListAlarmEventRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlarmEventRequest) SetPageSize(v string) *ListAlarmEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlarmEventRequest) SetReverse(v bool) *ListAlarmEventRequest {
	s.Reverse = &v
	return s
}

func (s *ListAlarmEventRequest) SetStartTime(v int64) *ListAlarmEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlarmEventRequest) Validate() error {
	return dara.Validate(s)
}
