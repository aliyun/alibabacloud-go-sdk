// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *ListFlowsRequest
	GetFlowName() *string
	SetPageNumber(v int32) *ListFlowsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlowsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListFlowsRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListFlowsRequest
	GetWorkspaceIds() *string
}

type ListFlowsRequest struct {
	// Filter by flow name
	//
	// example:
	//
	// my-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// List of workspace IDs
	//
	// example:
	//
	// ws-1234567890abcdef,ws-1234567890bcdefg
	WorkspaceIds *string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlowsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListFlowsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListFlowsRequest) SetFlowName(v string) *ListFlowsRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowsRequest) SetPageNumber(v int32) *ListFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsRequest) SetPageSize(v int32) *ListFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowsRequest) SetWorkspaceId(v string) *ListFlowsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListFlowsRequest) SetWorkspaceIds(v string) *ListFlowsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListFlowsRequest) Validate() error {
	return dara.Validate(s)
}
