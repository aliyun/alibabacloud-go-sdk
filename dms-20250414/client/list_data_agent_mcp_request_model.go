// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDataAgentMcpRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentMcpRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataAgentMcpRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentMcpRequest
	GetPageSize() *int32
	SetReadyOnly(v bool) *ListDataAgentMcpRequest
	GetReadyOnly() *bool
	SetSearchKey(v string) *ListDataAgentMcpRequest
	GetSearchKey() *string
	SetType(v string) *ListDataAgentMcpRequest
	GetType() *string
	SetWorkspaceId(v string) *ListDataAgentMcpRequest
	GetWorkspaceId() *string
}

type ListDataAgentMcpRequest struct {
	// A compatible pagination parameter. The actual number of records per page is controlled by PageSize.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A compatible pagination token. The actual page sequence is controlled by PageNumber.
	//
	// example:
	//
	// page-2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values: 1 to 500. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return only MCP Servers that are enabled and in the ready state. Default value: false.
	ReadyOnly *bool `json:"ReadyOnly,omitempty" xml:"ReadyOnly,omitempty"`
	// The keyword for name search. The server performs a fuzzy match against MCP Server names.
	//
	// example:
	//
	// analytics
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The MCP Server type. Valid values:
	//
	// - system: system MCP.
	//
	// - customer: custom MCP.
	//
	// example:
	//
	// customer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The Data Agent workspace ID. The caller must have at least MEMBER permissions on this workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// atvx***xmz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMcpRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentMcpRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentMcpRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentMcpRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentMcpRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentMcpRequest) GetReadyOnly() *bool {
	return s.ReadyOnly
}

func (s *ListDataAgentMcpRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataAgentMcpRequest) GetType() *string {
	return s.Type
}

func (s *ListDataAgentMcpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentMcpRequest) SetMaxResults(v int32) *ListDataAgentMcpRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetNextToken(v string) *ListDataAgentMcpRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetPageNumber(v int32) *ListDataAgentMcpRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetPageSize(v int32) *ListDataAgentMcpRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetReadyOnly(v bool) *ListDataAgentMcpRequest {
	s.ReadyOnly = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetSearchKey(v string) *ListDataAgentMcpRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetType(v string) *ListDataAgentMcpRequest {
	s.Type = &v
	return s
}

func (s *ListDataAgentMcpRequest) SetWorkspaceId(v string) *ListDataAgentMcpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentMcpRequest) Validate() error {
	return dara.Validate(s)
}
