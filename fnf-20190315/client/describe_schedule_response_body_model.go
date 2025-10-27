// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *DescribeScheduleResponseBody
	GetCreatedTime() *string
	SetCronExpression(v string) *DescribeScheduleResponseBody
	GetCronExpression() *string
	SetDescription(v string) *DescribeScheduleResponseBody
	GetDescription() *string
	SetEnable(v bool) *DescribeScheduleResponseBody
	GetEnable() *bool
	SetLastModifiedTime(v string) *DescribeScheduleResponseBody
	GetLastModifiedTime() *string
	SetPayload(v string) *DescribeScheduleResponseBody
	GetPayload() *string
	SetRequestId(v string) *DescribeScheduleResponseBody
	GetRequestId() *string
	SetScheduleId(v string) *DescribeScheduleResponseBody
	GetScheduleId() *string
	SetScheduleName(v string) *DescribeScheduleResponseBody
	GetScheduleName() *string
}

type DescribeScheduleResponseBody struct {
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

func (s DescribeScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeScheduleResponseBody) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeScheduleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeScheduleResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeScheduleResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *DescribeScheduleResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *DescribeScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScheduleResponseBody) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *DescribeScheduleResponseBody) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *DescribeScheduleResponseBody) SetCreatedTime(v string) *DescribeScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetCronExpression(v string) *DescribeScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetDescription(v string) *DescribeScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetEnable(v bool) *DescribeScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetLastModifiedTime(v string) *DescribeScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetPayload(v string) *DescribeScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetRequestId(v string) *DescribeScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetScheduleId(v string) *DescribeScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetScheduleName(v string) *DescribeScheduleResponseBody {
	s.ScheduleName = &v
	return s
}

func (s *DescribeScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
