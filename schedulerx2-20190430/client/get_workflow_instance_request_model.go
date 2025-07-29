// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetWorkflowInstanceRequest
	GetGroupId() *string
	SetNamespace(v string) *GetWorkflowInstanceRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetWorkflowInstanceRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *GetWorkflowInstanceRequest
	GetRegionId() *string
	SetWfInstanceId(v int64) *GetWorkflowInstanceRequest
	GetWfInstanceId() *int64
	SetWorkflowId(v int64) *GetWorkflowInstanceRequest
	GetWorkflowId() *int64
}

type GetWorkflowInstanceRequest struct {
	// The application group ID. You can obtain the ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	// The workflow instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetWorkflowInstanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetWorkflowInstanceRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetWorkflowInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkflowInstanceRequest) GetWfInstanceId() *int64 {
	return s.WfInstanceId
}

func (s *GetWorkflowInstanceRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowInstanceRequest) SetGroupId(v string) *GetWorkflowInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetNamespace(v string) *GetWorkflowInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetNamespaceSource(v string) *GetWorkflowInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetRegionId(v string) *GetWorkflowInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetWfInstanceId(v int64) *GetWorkflowInstanceRequest {
	s.WfInstanceId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetWorkflowId(v int64) *GetWorkflowInstanceRequest {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
