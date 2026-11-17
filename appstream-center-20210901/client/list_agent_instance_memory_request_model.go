// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstanceMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListAgentInstanceMemoryRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ListAgentInstanceMemoryRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ListAgentInstanceMemoryRequest
	GetBizType() *int32
	SetEndUserIds(v []*string) *ListAgentInstanceMemoryRequest
	GetEndUserIds() []*string
	SetExcludedEndUserIds(v []*string) *ListAgentInstanceMemoryRequest
	GetExcludedEndUserIds() []*string
	SetMemoryStatuses(v []*string) *ListAgentInstanceMemoryRequest
	GetMemoryStatuses() []*string
	SetPageNumber(v int32) *ListAgentInstanceMemoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentInstanceMemoryRequest
	GetPageSize() *int32
	SetSearchKeyword(v string) *ListAgentInstanceMemoryRequest
	GetSearchKeyword() *string
	SetSearchType(v string) *ListAgentInstanceMemoryRequest
	GetSearchType() *string
}

type ListAgentInstanceMemoryRequest struct {
	// Filters by the business platform to which the Agent belongs. If omitted, no restriction is applied to the business platform.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// Filters by Agent provider. If omitted, queries providers that support memory management. `JVSCopilot` is not supported for memory management in this operation.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// Filters by business type code. If omitted, no restriction is applied to the business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of end user IDs to include in the query results. Exact match is applied. You can specify up to 200 IDs.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The list of end user IDs to exclude from the query results. You can specify up to 200 IDs. If used together with EndUserIds, these users are excluded from the inclusion scope.
	ExcludedEndUserIds []*string `json:"ExcludedEndUserIds,omitempty" xml:"ExcludedEndUserIds,omitempty" type:"Repeated"`
	// Filters by the memory switch configured for the instance.
	//
	// - `ENABLED`: The instance memory switch is turned on.
	//
	// - `DISABLED`: The instance memory switch is turned off.
	//
	// If both values are specified, instances matching either status are returned.
	MemoryStatuses []*string `json:"MemoryStatuses,omitempty" xml:"MemoryStatuses,omitempty" type:"Repeated"`
	// The page number. Pages start from 1. If this parameter is not specified or is set to 0, the first page is queried.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used for inclusion matching. Specify the search field by using SearchType. If left empty, no keyword filtering is applied.
	//
	// example:
	//
	// ecd-d8ut
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	// The search scope for the keyword. Default value: `ALL`.
	//
	// - `ALL`: Searches the resource ID, resource name, runtime ID, and end user ID.
	//
	// - `RESOURCE_ID`: Searches the resource ID.
	//
	// - `RESOURCE_NAME`: Searches the resource name.
	//
	// - `RUNTIME_ID`: Searches the runtime ID.
	//
	// - `END_USER_ID`: Searches the end user ID.
	//
	// example:
	//
	// RESOURCE_ID
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
}

func (s ListAgentInstanceMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceMemoryRequest) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceMemoryRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListAgentInstanceMemoryRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListAgentInstanceMemoryRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListAgentInstanceMemoryRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *ListAgentInstanceMemoryRequest) GetExcludedEndUserIds() []*string {
	return s.ExcludedEndUserIds
}

func (s *ListAgentInstanceMemoryRequest) GetMemoryStatuses() []*string {
	return s.MemoryStatuses
}

func (s *ListAgentInstanceMemoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentInstanceMemoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentInstanceMemoryRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *ListAgentInstanceMemoryRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListAgentInstanceMemoryRequest) SetAgentPlatform(v string) *ListAgentInstanceMemoryRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetAgentProvider(v string) *ListAgentInstanceMemoryRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetBizType(v int32) *ListAgentInstanceMemoryRequest {
	s.BizType = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetEndUserIds(v []*string) *ListAgentInstanceMemoryRequest {
	s.EndUserIds = v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetExcludedEndUserIds(v []*string) *ListAgentInstanceMemoryRequest {
	s.ExcludedEndUserIds = v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetMemoryStatuses(v []*string) *ListAgentInstanceMemoryRequest {
	s.MemoryStatuses = v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetPageNumber(v int32) *ListAgentInstanceMemoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetPageSize(v int32) *ListAgentInstanceMemoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetSearchKeyword(v string) *ListAgentInstanceMemoryRequest {
	s.SearchKeyword = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) SetSearchType(v string) *ListAgentInstanceMemoryRequest {
	s.SearchType = &v
	return s
}

func (s *ListAgentInstanceMemoryRequest) Validate() error {
	return dara.Validate(s)
}
