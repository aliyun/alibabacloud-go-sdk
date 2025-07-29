// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DisableWorkflowRequest
	GetGroupId() *string
	SetNamespace(v string) *DisableWorkflowRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *DisableWorkflowRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *DisableWorkflowRequest
	GetRegionId() *string
	SetWorkflowId(v int64) *DisableWorkflowRequest
	GetWorkflowId() *int64
}

type DisableWorkflowRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
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
	// 111
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DisableWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DisableWorkflowRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DisableWorkflowRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DisableWorkflowRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *DisableWorkflowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *DisableWorkflowRequest) SetGroupId(v string) *DisableWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *DisableWorkflowRequest) SetNamespace(v string) *DisableWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *DisableWorkflowRequest) SetNamespaceSource(v string) *DisableWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DisableWorkflowRequest) SetRegionId(v string) *DisableWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *DisableWorkflowRequest) SetWorkflowId(v int64) *DisableWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *DisableWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
