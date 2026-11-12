// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateWorkflowRequest
	GetAppName() *string
	SetCalendar(v string) *UpdateWorkflowRequest
	GetCalendar() *string
	SetClientToken(v string) *UpdateWorkflowRequest
	GetClientToken() *string
	SetClusterId(v string) *UpdateWorkflowRequest
	GetClusterId() *string
	SetDescription(v string) *UpdateWorkflowRequest
	GetDescription() *string
	SetMaxConcurrency(v int32) *UpdateWorkflowRequest
	GetMaxConcurrency() *int32
	SetName(v string) *UpdateWorkflowRequest
	GetName() *string
	SetTimeExpression(v string) *UpdateWorkflowRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *UpdateWorkflowRequest
	GetTimeType() *int32
	SetTimezone(v string) *UpdateWorkflowRequest
	GetTimezone() *string
	SetWorkflowId(v int64) *UpdateWorkflowRequest
	GetWorkflowId() *int64
}

type UpdateWorkflowRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The custom calendar.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can generate the value on the client side. Ensure that the value is unique across different requests and contains a maximum of 64 ASCII characters.
	//
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The workflow description.
	//
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum concurrency.
	//
	// > The maximum number of instances that can run concurrently for the same workflow. A value of 1 indicates that concurrent execution is not allowed. If the concurrency limit is exceeded, the current scheduling task is skipped.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// myWorkflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time expression. Configure the time expression based on the selected time type.
	//
	// - none: Leave this parameter empty.
	//
	// - cron: Enter a standard cron expression. Online verification is supported.
	//
	// - api: Leave this parameter empty.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The scheduling type. Valid values:
	//
	// - -1: none
	//
	// - 1: cron
	//
	// - 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone.
	//
	// > By default, the time zone of the SchedulerX server is used.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWorkflowRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *UpdateWorkflowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkflowRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *UpdateWorkflowRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *UpdateWorkflowRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *UpdateWorkflowRequest) SetAppName(v string) *UpdateWorkflowRequest {
	s.AppName = &v
	return s
}

func (s *UpdateWorkflowRequest) SetCalendar(v string) *UpdateWorkflowRequest {
	s.Calendar = &v
	return s
}

func (s *UpdateWorkflowRequest) SetClientToken(v string) *UpdateWorkflowRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkflowRequest) SetClusterId(v string) *UpdateWorkflowRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetDescription(v string) *UpdateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequest) SetMaxConcurrency(v int32) *UpdateWorkflowRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateWorkflowRequest) SetName(v string) *UpdateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimeExpression(v string) *UpdateWorkflowRequest {
	s.TimeExpression = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimeType(v int32) *UpdateWorkflowRequest {
	s.TimeType = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimezone(v string) *UpdateWorkflowRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateWorkflowRequest) SetWorkflowId(v int64) *UpdateWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
