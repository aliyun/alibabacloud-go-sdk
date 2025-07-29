// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateWorkflowRequest
	GetDescription() *string
	SetGroupId(v string) *UpdateWorkflowRequest
	GetGroupId() *string
	SetName(v string) *UpdateWorkflowRequest
	GetName() *string
	SetNamespace(v string) *UpdateWorkflowRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *UpdateWorkflowRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *UpdateWorkflowRequest
	GetRegionId() *string
	SetTimeExpression(v string) *UpdateWorkflowRequest
	GetTimeExpression() *string
	SetTimeType(v int32) *UpdateWorkflowRequest
	GetTimeType() *int32
	SetWorkflowId(v string) *UpdateWorkflowRequest
	GetWorkflowId() *string
}

type UpdateWorkflowRequest struct {
	// The description of the workflow.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the workflow.
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
	// 	- If you set TimeType to cron, you need to enter a standard cron expression. Online verification is supported.
	//
	// 	- If you set TimeType to api, no time expression is required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The method that is used to specify the time. Valid values:
	//
	// 	- 1: cron
	//
	// 	- 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateWorkflowRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *UpdateWorkflowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWorkflowRequest) GetTimeExpression() *string {
	return s.TimeExpression
}

func (s *UpdateWorkflowRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *UpdateWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *UpdateWorkflowRequest) SetDescription(v string) *UpdateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequest) SetGroupId(v string) *UpdateWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetName(v string) *UpdateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequest) SetNamespace(v string) *UpdateWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateWorkflowRequest) SetNamespaceSource(v string) *UpdateWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateWorkflowRequest) SetRegionId(v string) *UpdateWorkflowRequest {
	s.RegionId = &v
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

func (s *UpdateWorkflowRequest) SetWorkflowId(v string) *UpdateWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
