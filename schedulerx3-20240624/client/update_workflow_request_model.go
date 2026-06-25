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
	// The name of a custom calendar to exclude specific dates from the schedule.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// A unique token that you provide to ensure the request is idempotent. The token can be up to 64 ASCII characters long.
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
	// The maximum concurrency for the workflow.
	//
	// > The maximum number of concurrent instances that can run for the same workflow. A value of `1` prevents overlapping runs. If the number of running instances reaches this limit, subsequent scheduled runs are skipped.
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
	// The cron expression for the schedule. This parameter is required only when `TimeType` is set to `1` (cron).
	//
	// - If `TimeType` is `-1` (none), this parameter is not required.
	//
	// - If `TimeType` is `1` (cron), specify a standard cron expression.
	//
	// - If `TimeType` is `100` (api), this parameter is not required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The scheduling type. Valid values:
	//
	// - `-1` (none): The workflow is not scheduled and must be triggered on demand.
	//
	// - `1` (cron): The workflow runs based on the cron expression in the `TimeExpression` parameter.
	//
	// - `100` (api): The workflow is triggered by an API call.
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone for the schedule.
	//
	// > Defaults to the time zone of the SchedulerX server.
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
