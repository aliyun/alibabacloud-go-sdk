// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *ListCustomAgentRequest
	GetAgentType() *string
	SetPageNumber(v int32) *ListCustomAgentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomAgentRequest
	GetPageSize() *int32
	SetQueryAllReleased(v bool) *ListCustomAgentRequest
	GetQueryAllReleased() *bool
	SetSearchKey(v string) *ListCustomAgentRequest
	GetSearchKey() *string
	SetStatus(v string) *ListCustomAgentRequest
	GetStatus() *string
	SetWorkspaceId(v string) *ListCustomAgentRequest
	GetWorkspaceId() *string
}

type ListCustomAgentRequest struct {
	// The custom agent type. Valid values:
	//
	// - Default: system built-in custom agent
	//
	// - Common: regular custom agent
	//
	// example:
	//
	// Common
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query all published custom agents. Default value: false.
	//
	// example:
	//
	// true
	QueryAllReleased *bool `json:"QueryAllReleased,omitempty" xml:"QueryAllReleased,omitempty"`
	// The search keyword. Supports fuzzy search for custom agent names and descriptions.
	//
	// example:
	//
	// Test agent
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The status of the custom agent.
	//
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *ListCustomAgentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomAgentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomAgentRequest) GetQueryAllReleased() *bool {
	return s.QueryAllReleased
}

func (s *ListCustomAgentRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListCustomAgentRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCustomAgentRequest) SetAgentType(v string) *ListCustomAgentRequest {
	s.AgentType = &v
	return s
}

func (s *ListCustomAgentRequest) SetPageNumber(v int32) *ListCustomAgentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAgentRequest) SetPageSize(v int32) *ListCustomAgentRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomAgentRequest) SetQueryAllReleased(v bool) *ListCustomAgentRequest {
	s.QueryAllReleased = &v
	return s
}

func (s *ListCustomAgentRequest) SetSearchKey(v string) *ListCustomAgentRequest {
	s.SearchKey = &v
	return s
}

func (s *ListCustomAgentRequest) SetStatus(v string) *ListCustomAgentRequest {
	s.Status = &v
	return s
}

func (s *ListCustomAgentRequest) SetWorkspaceId(v string) *ListCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
