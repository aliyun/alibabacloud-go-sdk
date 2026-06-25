// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateWorkflowRequest
	GetAppName() *string
	SetCalendar(v string) *CreateWorkflowRequest
	GetCalendar() *string
	SetClientToken(v string) *CreateWorkflowRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateWorkflowRequest
	GetClusterId() *string
	SetDescription(v string) *CreateWorkflowRequest
	GetDescription() *string
	SetMaxConcurrency(v int32) *CreateWorkflowRequest
	GetMaxConcurrency() *int32
	SetName(v string) *CreateWorkflowRequest
	GetName() *string
	SetStatus(v int32) *CreateWorkflowRequest
	GetStatus() *int32
	SetTimeExpression(v string) *CreateWorkflowRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *CreateWorkflowRequest
	GetTimeType() *int32
	SetTimezone(v string) *CreateWorkflowRequest
	GetTimezone() *string
}

type CreateWorkflowRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxl-job-executor-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The custom calendar. This parameter applies only when `TimeType` is `cron`.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// A unique client token to ensure request idempotence. The token must contain only ASCII characters. If you omit this parameter, the system uses the RequestId as the ClientToken. The RequestId is unique to each request.
	//
	// example:
	//
	// 123456789
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The workflow description.
	//
	// example:
	//
	// my first workflow for data analyse
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum concurrency for the workflow.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The workflow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-workflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the workflow. By default, the workflow is disabled. Valid values:
	//
	// - 0: Disabled
	//
	// - 1: Enabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time expression, which depends on the `TimeType` parameter.
	//
	// - **none**: This parameter is not required.
	//
	// - **cron**: Enter a standard cron expression. Online validation is supported.
	//
	// - **api**: This parameter is not required.
	//
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The schedule type. Valid values:
	//
	// - -1: none<br>
	//
	// - 1: cron<br>
	//
	// - 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone for the schedule.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateWorkflowRequest) GetCalendar() *string {
	return s.Calendar
}

func (s *CreateWorkflowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkflowRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkflowRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *CreateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateWorkflowRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *CreateWorkflowRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *CreateWorkflowRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateWorkflowRequest) SetAppName(v string) *CreateWorkflowRequest {
	s.AppName = &v
	return s
}

func (s *CreateWorkflowRequest) SetCalendar(v string) *CreateWorkflowRequest {
	s.Calendar = &v
	return s
}

func (s *CreateWorkflowRequest) SetClientToken(v string) *CreateWorkflowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkflowRequest) SetClusterId(v string) *CreateWorkflowRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateWorkflowRequest) SetDescription(v string) *CreateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkflowRequest) SetMaxConcurrency(v int32) *CreateWorkflowRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateWorkflowRequest) SetName(v string) *CreateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkflowRequest) SetStatus(v int32) *CreateWorkflowRequest {
	s.Status = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimeExpression(v string) *CreateWorkflowRequest {
	s.TimeExpression = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimeType(v int32) *CreateWorkflowRequest {
	s.TimeType = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimezone(v string) *CreateWorkflowRequest {
	s.Timezone = &v
	return s
}

func (s *CreateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
