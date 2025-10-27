// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSchedulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSchedulesResponseBody
	GetRequestId() *string
	SetSchedules(v []*ListSchedulesResponseBodySchedules) *ListSchedulesResponseBody
	GetSchedules() []*ListSchedulesResponseBodySchedules
}

type ListSchedulesResponseBody struct {
	// The token for the next query.
	//
	// example:
	//
	// testNextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time-based schedules that are queried.
	Schedules []*ListSchedulesResponseBodySchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s ListSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchedulesResponseBody) GetSchedules() []*ListSchedulesResponseBodySchedules {
	return s.Schedules
}

func (s *ListSchedulesResponseBody) SetNextToken(v string) *ListSchedulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSchedulesResponseBody) SetRequestId(v string) *ListSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchedulesResponseBody) SetSchedules(v []*ListSchedulesResponseBodySchedules) *ListSchedulesResponseBody {
	s.Schedules = v
	return s
}

func (s *ListSchedulesResponseBody) Validate() error {
	if s.Schedules != nil {
		for _, item := range s.Schedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSchedulesResponseBodySchedules struct {
	// The time when the time-based schedule was created.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cron expression of the scheduled task.
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

func (s ListSchedulesResponseBodySchedules) String() string {
	return dara.Prettify(s)
}

func (s ListSchedulesResponseBodySchedules) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodySchedules) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListSchedulesResponseBodySchedules) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListSchedulesResponseBodySchedules) GetDescription() *string {
	return s.Description
}

func (s *ListSchedulesResponseBodySchedules) GetEnable() *bool {
	return s.Enable
}

func (s *ListSchedulesResponseBodySchedules) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *ListSchedulesResponseBodySchedules) GetPayload() *string {
	return s.Payload
}

func (s *ListSchedulesResponseBodySchedules) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *ListSchedulesResponseBodySchedules) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *ListSchedulesResponseBodySchedules) SetCreatedTime(v string) *ListSchedulesResponseBodySchedules {
	s.CreatedTime = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetCronExpression(v string) *ListSchedulesResponseBodySchedules {
	s.CronExpression = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetDescription(v string) *ListSchedulesResponseBodySchedules {
	s.Description = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetEnable(v bool) *ListSchedulesResponseBodySchedules {
	s.Enable = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetLastModifiedTime(v string) *ListSchedulesResponseBodySchedules {
	s.LastModifiedTime = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetPayload(v string) *ListSchedulesResponseBodySchedules {
	s.Payload = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetScheduleId(v string) *ListSchedulesResponseBodySchedules {
	s.ScheduleId = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetScheduleName(v string) *ListSchedulesResponseBodySchedules {
	s.ScheduleName = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) Validate() error {
	return dara.Validate(s)
}
