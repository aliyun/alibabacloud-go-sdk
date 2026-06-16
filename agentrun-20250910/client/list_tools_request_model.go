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
	SetToolName(v string) *ListToolsRequest
	GetToolName() *string
	SetToolType(v string) *ListToolsRequest
	GetToolType() *string
	SetWorkspaceId(v string) *ListToolsRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListToolsRequest
	GetWorkspaceIds() *string
}

type ListToolsRequest struct {
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
	// Tool Name, supports fuzzy search
	//
	// example:
	//
	// tool-1
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
	// Tool type
	//
	// example:
	//
	// MCP
	ToolType *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// workspace-xyz789
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// List of workspace IDs, separated by commas
	//
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

func (s *ListToolsRequest) GetToolName() *string {
	return s.ToolName
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

func (s *ListToolsRequest) SetToolName(v string) *ListToolsRequest {
	s.ToolName = &v
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
