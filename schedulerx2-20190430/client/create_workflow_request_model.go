// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkflowRequest
	GetDescription() *string
	SetGroupId(v string) *CreateWorkflowRequest
	GetGroupId() *string
	SetMaxConcurrency(v int32) *CreateWorkflowRequest
	GetMaxConcurrency() *int32
	SetName(v string) *CreateWorkflowRequest
	GetName() *string
	SetNamespace(v string) *CreateWorkflowRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *CreateWorkflowRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *CreateWorkflowRequest
	GetRegionId() *string
	SetTimeExpression(v string) *CreateWorkflowRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *CreateWorkflowRequest
	GetTimeType() *int32
	SetTimezone(v string) *CreateWorkflowRequest
	GetTimezone() *string
}

type CreateWorkflowRequest struct {
	// The description of the workflow.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group ID. You can obtain the ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of workflow instances that can be run at the same time. Default value: 1. The value 1 indicates that only one workflow instance is allowed. In this case, if the triggered workflow instance is still ongoing, no more workflow instances can be triggered even the time to schedule the next workflow arrives.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time expression. You can set the time expression based on the selected method that is used to specify time.
	//
	// - If you set the TimeType parameter to cron, you need to enter a standard cron expression. Online verification is supported.
	//
	// - If you set the TimeType parameter to api, no time expression is required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The method that is used to specify the time. Valid values:
	//
	// - 1: cron
	//
	// - 100: api
	//
	// This parameter is required.
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

func (s *CreateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkflowRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateWorkflowRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *CreateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateWorkflowRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *CreateWorkflowRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *CreateWorkflowRequest) SetDescription(v string) *CreateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkflowRequest) SetGroupId(v string) *CreateWorkflowRequest {
	s.GroupId = &v
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

func (s *CreateWorkflowRequest) SetNamespace(v string) *CreateWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *CreateWorkflowRequest) SetNamespaceSource(v string) *CreateWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *CreateWorkflowRequest) SetRegionId(v string) *CreateWorkflowRequest {
	s.RegionId = &v
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
