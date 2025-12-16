// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListWorkFlowsRequest
	GetGroupId() *string
	SetNamespace(v string) *ListWorkFlowsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ListWorkFlowsRequest
	GetNamespaceSource() *string
	SetPageNum(v int32) *ListWorkFlowsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListWorkFlowsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListWorkFlowsRequest
	GetRegionId() *string
	SetStatus(v int32) *ListWorkFlowsRequest
	GetStatus() *int32
	SetWorkflowName(v string) *ListWorkFlowsRequest
	GetWorkflowName() *string
}

type ListWorkFlowsRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// hxm.test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a06d5ea-f576-4326-842c-fb14ea043d8d
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for special sources.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job status.
	//
	// 	- **0**: disables the job.
	//
	// 	- **1**: enables the routing policy.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// test3
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ListWorkFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkFlowsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListWorkFlowsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWorkFlowsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ListWorkFlowsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListWorkFlowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkFlowsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkFlowsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListWorkFlowsRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListWorkFlowsRequest) SetGroupId(v string) *ListWorkFlowsRequest {
	s.GroupId = &v
	return s
}

func (s *ListWorkFlowsRequest) SetNamespace(v string) *ListWorkFlowsRequest {
	s.Namespace = &v
	return s
}

func (s *ListWorkFlowsRequest) SetNamespaceSource(v string) *ListWorkFlowsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListWorkFlowsRequest) SetPageNum(v int32) *ListWorkFlowsRequest {
	s.PageNum = &v
	return s
}

func (s *ListWorkFlowsRequest) SetPageSize(v int32) *ListWorkFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkFlowsRequest) SetRegionId(v string) *ListWorkFlowsRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkFlowsRequest) SetStatus(v int32) *ListWorkFlowsRequest {
	s.Status = &v
	return s
}

func (s *ListWorkFlowsRequest) SetWorkflowName(v string) *ListWorkFlowsRequest {
	s.WorkflowName = &v
	return s
}

func (s *ListWorkFlowsRequest) Validate() error {
	return dara.Validate(s)
}
