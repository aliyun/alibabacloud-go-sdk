// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListWorkflowInstanceRequest
	GetGroupId() *string
	SetNamespace(v string) *ListWorkflowInstanceRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ListWorkflowInstanceRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *ListWorkflowInstanceRequest
	GetRegionId() *string
	SetWorkflowId(v string) *ListWorkflowInstanceRequest
	GetWorkflowId() *string
}

type ListWorkflowInstanceRequest struct {
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
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListWorkflowInstanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWorkflowInstanceRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ListWorkflowInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkflowInstanceRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowInstanceRequest) SetGroupId(v string) *ListWorkflowInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetNamespace(v string) *ListWorkflowInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetNamespaceSource(v string) *ListWorkflowInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetRegionId(v string) *ListWorkflowInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetWorkflowId(v string) *ListWorkflowInstanceRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
