// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetWorkFlowRequest
	GetGroupId() *string
	SetNamespace(v string) *GetWorkFlowRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetWorkFlowRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *GetWorkFlowRequest
	GetRegionId() *string
	SetWorkflowId(v int64) *GetWorkFlowRequest
	GetWorkflowId() *int64
}

type GetWorkFlowRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// hxm.test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a06d5ea-f576-4326-842c-fb14ea043d8d
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace.
	//
	// example:
	//
	// source
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region information.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkFlowRequest) GoString() string {
	return s.String()
}

func (s *GetWorkFlowRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetWorkFlowRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetWorkFlowRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetWorkFlowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkFlowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkFlowRequest) SetGroupId(v string) *GetWorkFlowRequest {
	s.GroupId = &v
	return s
}

func (s *GetWorkFlowRequest) SetNamespace(v string) *GetWorkFlowRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkFlowRequest) SetNamespaceSource(v string) *GetWorkFlowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetWorkFlowRequest) SetRegionId(v string) *GetWorkFlowRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkFlowRequest) SetWorkflowId(v int64) *GetWorkFlowRequest {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkFlowRequest) Validate() error {
	return dara.Validate(s)
}
