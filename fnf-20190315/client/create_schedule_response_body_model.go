// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *CreateScheduleResponseBody
	GetCreatedTime() *string
	SetCronExpression(v string) *CreateScheduleResponseBody
	GetCronExpression() *string
	SetDescription(v string) *CreateScheduleResponseBody
	GetDescription() *string
	SetEnable(v bool) *CreateScheduleResponseBody
	GetEnable() *bool
	SetLastModifiedTime(v string) *CreateScheduleResponseBody
	GetLastModifiedTime() *string
	SetPayload(v string) *CreateScheduleResponseBody
	GetPayload() *string
	SetRequestId(v string) *CreateScheduleResponseBody
	GetRequestId() *string
	SetScheduleId(v string) *CreateScheduleResponseBody
	GetScheduleId() *string
	SetScheduleName(v string) *CreateScheduleResponseBody
	GetScheduleName() *string
}

type CreateScheduleResponseBody struct {
	// The time when the time-based schedule was created.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the time-based schedule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the time-based schedule was last modified.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The trigger message of the time-based schedule.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the time-based schedule.
	//
	// example:
	//
	// testScheduleId
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The name of the time-based schedule.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s CreateScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CreateScheduleResponseBody) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateScheduleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateScheduleResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *CreateScheduleResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *CreateScheduleResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *CreateScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduleResponseBody) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *CreateScheduleResponseBody) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *CreateScheduleResponseBody) SetCreatedTime(v string) *CreateScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateScheduleResponseBody) SetCronExpression(v string) *CreateScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduleResponseBody) SetDescription(v string) *CreateScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *CreateScheduleResponseBody) SetEnable(v bool) *CreateScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *CreateScheduleResponseBody) SetLastModifiedTime(v string) *CreateScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateScheduleResponseBody) SetPayload(v string) *CreateScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *CreateScheduleResponseBody) SetRequestId(v string) *CreateScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleResponseBody) SetScheduleId(v string) *CreateScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *CreateScheduleResponseBody) SetScheduleName(v string) *CreateScheduleResponseBody {
	s.ScheduleName = &v
	return s
}

func (s *CreateScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
