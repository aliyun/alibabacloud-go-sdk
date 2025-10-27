// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *UpdateScheduleResponseBody
	GetCreatedTime() *string
	SetCronExpression(v string) *UpdateScheduleResponseBody
	GetCronExpression() *string
	SetDescription(v string) *UpdateScheduleResponseBody
	GetDescription() *string
	SetEnable(v bool) *UpdateScheduleResponseBody
	GetEnable() *bool
	SetLastModifiedTime(v string) *UpdateScheduleResponseBody
	GetLastModifiedTime() *string
	SetPayload(v string) *UpdateScheduleResponseBody
	GetPayload() *string
	SetRequestId(v string) *UpdateScheduleResponseBody
	GetRequestId() *string
	SetScheduleId(v string) *UpdateScheduleResponseBody
	GetScheduleId() *string
	SetScheduleName(v string) *UpdateScheduleResponseBody
	GetScheduleName() *string
}

type UpdateScheduleResponseBody struct {
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
	// Indicates whether the time-based schedule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the time-based schedule was last updated.
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

func (s UpdateScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *UpdateScheduleResponseBody) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateScheduleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateScheduleResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateScheduleResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *UpdateScheduleResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *UpdateScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduleResponseBody) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *UpdateScheduleResponseBody) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *UpdateScheduleResponseBody) SetCreatedTime(v string) *UpdateScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetCronExpression(v string) *UpdateScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetDescription(v string) *UpdateScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetEnable(v bool) *UpdateScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetLastModifiedTime(v string) *UpdateScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetPayload(v string) *UpdateScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetRequestId(v string) *UpdateScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetScheduleId(v string) *UpdateScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetScheduleName(v string) *UpdateScheduleResponseBody {
	s.ScheduleName = &v
	return s
}

func (s *UpdateScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
