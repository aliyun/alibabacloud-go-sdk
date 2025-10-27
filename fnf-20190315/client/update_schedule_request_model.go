// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *UpdateScheduleRequest
	GetCronExpression() *string
	SetDescription(v string) *UpdateScheduleRequest
	GetDescription() *string
	SetEnable(v bool) *UpdateScheduleRequest
	GetEnable() *bool
	SetFlowName(v string) *UpdateScheduleRequest
	GetFlowName() *string
	SetPayload(v string) *UpdateScheduleRequest
	GetPayload() *string
	SetScheduleName(v string) *UpdateScheduleRequest
	GetScheduleName() *string
}

type UpdateScheduleRequest struct {
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
	// Specifies whether to enable the time-based schedule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The name of the flow that is associated with the time-based schedule. The name must be unique within the region and cannot be modified after the time-based schedule is created. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testFlowName
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The trigger message of the time-based schedule. It must be in the JSON format.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The name of the time-based schedule. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s UpdateScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateScheduleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateScheduleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateScheduleRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateScheduleRequest) GetPayload() *string {
	return s.Payload
}

func (s *UpdateScheduleRequest) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *UpdateScheduleRequest) SetCronExpression(v string) *UpdateScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateScheduleRequest) SetDescription(v string) *UpdateScheduleRequest {
	s.Description = &v
	return s
}

func (s *UpdateScheduleRequest) SetEnable(v bool) *UpdateScheduleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateScheduleRequest) SetFlowName(v string) *UpdateScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateScheduleRequest) SetPayload(v string) *UpdateScheduleRequest {
	s.Payload = &v
	return s
}

func (s *UpdateScheduleRequest) SetScheduleName(v string) *UpdateScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *UpdateScheduleRequest) Validate() error {
	return dara.Validate(s)
}
