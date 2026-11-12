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
	// The custom calendar. This parameter is optional when the time type is cron.
	//
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can generate the value from your client and ensure that the value is unique across different requests. The client token can contain only ASCII characters. Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId may be different for each API request.
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
	// The maximum concurrency of the workflow.
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
	// The workflow status. Default value: 0. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time expression. Set the time expression based on the selected time type.
	//
	// - **none**: Not required.
	//
	// - **cron**: Enter a standard cron expression. Online validation is supported.
	//
	// - **api**: Not required.
	//
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// - -1: none<br/>
	//
	// - 1: cron<br/>
	//
	// - 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone.
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
