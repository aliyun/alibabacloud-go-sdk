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
	// Filters results by the business platform to which the agent belongs. If this parameter is omitted, no restriction is applied to the business platform.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// Filters results by agent provider. If this parameter is omitted, the system queries providers that support memory management. `JVSCopilot` is not supported for memory management in this API operation.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// Filters results by business type code. If this parameter is omitted, no restriction is applied to the business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of end user IDs to be included in the query results. Exact matching is applied. Maximum value: 200.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The list of end user IDs to be excluded from the query results. Maximum value: 200. When used together with EndUserIds, these users are removed from the inclusion scope.
	ExcludedEndUserIds []*string `json:"ExcludedEndUserIds,omitempty" xml:"ExcludedEndUserIds,omitempty" type:"Repeated"`
	// Filters results by the memory switch configured for the instance.
	//
	// - `ENABLED`: The instance memory switch is turned on.
	//
	// - `DISABLED`: The instance memory switch is turned off.
	//
	// If both values are passed, instances matching either status are returned.
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
	// The keyword used for inclusion matching. Use SearchType to specify the search field. If this parameter is left empty, no keyword filtering is applied.
	//
	// example:
	//
	// ecd-d8ut
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	// The scope of the keyword search. Default value: `ALL`.
	//
	// - `ALL`: Searches resource IDs, resource names, runtime IDs, and end user IDs.
	//
	// - `RESOURCE_ID`: Searches resource IDs.
	//
	// - `RESOURCE_NAME`: Searches resource names.
	//
	// - `RUNTIME_ID`: Searches runtime IDs.
	//
	// - `END_USER_ID`: Searches end user IDs.
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
