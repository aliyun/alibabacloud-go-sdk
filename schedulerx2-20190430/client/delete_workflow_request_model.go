// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteWorkflowRequest
	GetGroupId() *string
	SetNamespace(v string) *DeleteWorkflowRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *DeleteWorkflowRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *DeleteWorkflowRequest
	GetRegionId() *string
	SetWorkflowId(v int64) *DeleteWorkflowRequest
	GetWorkflowId() *int64
}

type DeleteWorkflowRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
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

func (s DeleteWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteWorkflowRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteWorkflowRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *DeleteWorkflowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *DeleteWorkflowRequest) SetGroupId(v string) *DeleteWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetNamespace(v string) *DeleteWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteWorkflowRequest) SetNamespaceSource(v string) *DeleteWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DeleteWorkflowRequest) SetRegionId(v string) *DeleteWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetWorkflowId(v int64) *DeleteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *DeleteWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
