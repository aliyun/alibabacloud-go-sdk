// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *CreateScheduleRequest
	GetCronExpression() *string
	SetDescription(v string) *CreateScheduleRequest
	GetDescription() *string
	SetEnable(v bool) *CreateScheduleRequest
	GetEnable() *bool
	SetFlowName(v string) *CreateScheduleRequest
	GetFlowName() *string
	SetPayload(v string) *CreateScheduleRequest
	GetPayload() *string
	SetScheduleName(v string) *CreateScheduleRequest
	GetScheduleName() *string
	SetSignatureVersion(v string) *CreateScheduleRequest
	GetSignatureVersion() *string
}

type CreateScheduleRequest struct {
	// The CRON expression.
	//
	// This parameter is required.
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
	// The name of the workflow that is associated with the time-based schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The trigger message of the time-based schedule. Specify the value in the JSON format.
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
	// 	- It is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testScheduleName
	ScheduleName     *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	SignatureVersion *string `json:"SignatureVersion,omitempty" xml:"SignatureVersion,omitempty"`
}

func (s CreateScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateScheduleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScheduleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateScheduleRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateScheduleRequest) GetPayload() *string {
	return s.Payload
}

func (s *CreateScheduleRequest) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *CreateScheduleRequest) GetSignatureVersion() *string {
	return s.SignatureVersion
}

func (s *CreateScheduleRequest) SetCronExpression(v string) *CreateScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduleRequest) SetDescription(v string) *CreateScheduleRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduleRequest) SetEnable(v bool) *CreateScheduleRequest {
	s.Enable = &v
	return s
}

func (s *CreateScheduleRequest) SetFlowName(v string) *CreateScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *CreateScheduleRequest) SetPayload(v string) *CreateScheduleRequest {
	s.Payload = &v
	return s
}

func (s *CreateScheduleRequest) SetScheduleName(v string) *CreateScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *CreateScheduleRequest) SetSignatureVersion(v string) *CreateScheduleRequest {
	s.SignatureVersion = &v
	return s
}

func (s *CreateScheduleRequest) Validate() error {
	return dara.Validate(s)
}
