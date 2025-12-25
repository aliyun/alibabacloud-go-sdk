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
	// This parameter is required.
	//
	// example:
	//
	// xxl-job-executor-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// workday
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// example:
	//
	// 123456789
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// my first workflow for data analyse
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-workflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0 0 4 ? 	- Mon/1
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
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
