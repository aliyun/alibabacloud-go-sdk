// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListToolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListToolsRequest
	GetPageSize() *int32
	SetToolType(v string) *ListToolsRequest
	GetToolType() *string
	SetWorkspaceId(v string) *ListToolsRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListToolsRequest
	GetWorkspaceIds() *string
}

type ListToolsRequest struct {
	// 当前页码，从 1 开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的工具数量，用于分页查询
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按工具类型过滤，可选值：MCP、FUNCTIONCALL、SKILL
	//
	// example:
	//
	// MCP
	ToolType *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
	// 按工作空间ID过滤，查询指定工作空间下的工具
	//
	// example:
	//
	// workspace-xyz789
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// ws1,ws2
	WorkspaceIds *string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty"`
}

func (s ListToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListToolsRequest) GoString() string {
	return s.String()
}

func (s *ListToolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListToolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListToolsRequest) GetToolType() *string {
	return s.ToolType
}

func (s *ListToolsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListToolsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListToolsRequest) SetPageNumber(v int32) *ListToolsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListToolsRequest) SetPageSize(v int32) *ListToolsRequest {
	s.PageSize = &v
	return s
}

func (s *ListToolsRequest) SetToolType(v string) *ListToolsRequest {
	s.ToolType = &v
	return s
}

func (s *ListToolsRequest) SetWorkspaceId(v string) *ListToolsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListToolsRequest) SetWorkspaceIds(v string) *ListToolsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListToolsRequest) Validate() error {
	return dara.Validate(s)
}
