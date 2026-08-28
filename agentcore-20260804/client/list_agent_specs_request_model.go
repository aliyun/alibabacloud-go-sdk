// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpecName(v string) *ListAgentSpecsRequest
	GetAgentSpecName() *string
	SetBizTag(v string) *ListAgentSpecsRequest
	GetBizTag() *string
	SetOrderBy(v string) *ListAgentSpecsRequest
	GetOrderBy() *string
	SetOwner(v string) *ListAgentSpecsRequest
	GetOwner() *string
	SetPageNo(v int32) *ListAgentSpecsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAgentSpecsRequest
	GetPageSize() *int32
	SetScope(v string) *ListAgentSpecsRequest
	GetScope() *string
	SetSearch(v string) *ListAgentSpecsRequest
	GetSearch() *string
	SetWithCapabilities(v bool) *ListAgentSpecsRequest
	GetWithCapabilities() *bool
}

type ListAgentSpecsRequest struct {
	// The AgentSpec name used as a search keyword. Use this parameter together with the search parameter.
	//
	// example:
	//
	// my-worker
	AgentSpecName *string `json:"agentSpecName,omitempty" xml:"agentSpecName,omitempty"`
	// The business tag used for fuzzy filtering.
	//
	// example:
	//
	// ai
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// The field by which to sort results. Set this parameter to download_count to sort by download count. By default, results are sorted by update time.
	//
	// example:
	//
	// download_count
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The owner used to filter results.
	//
	// example:
	//
	// user1
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The visibility scope used to filter results. Valid values:
	//
	// - PUBLIC
	//
	// - PRIVATE
	//
	// example:
	//
	// PUBLIC
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The search mode. Valid values:
	//
	// - accurate: exact match.
	//
	// - blur: fuzzy match.
	//
	// Default value: blur.
	//
	// example:
	//
	// blur
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// Specifies whether to return the Skills and McpServers lists. Default value: false.
	//
	// example:
	//
	// true
	WithCapabilities *bool `json:"withCapabilities,omitempty" xml:"withCapabilities,omitempty"`
}

func (s ListAgentSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsRequest) GetAgentSpecName() *string {
	return s.AgentSpecName
}

func (s *ListAgentSpecsRequest) GetBizTag() *string {
	return s.BizTag
}

func (s *ListAgentSpecsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListAgentSpecsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListAgentSpecsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAgentSpecsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentSpecsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListAgentSpecsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListAgentSpecsRequest) GetWithCapabilities() *bool {
	return s.WithCapabilities
}

func (s *ListAgentSpecsRequest) SetAgentSpecName(v string) *ListAgentSpecsRequest {
	s.AgentSpecName = &v
	return s
}

func (s *ListAgentSpecsRequest) SetBizTag(v string) *ListAgentSpecsRequest {
	s.BizTag = &v
	return s
}

func (s *ListAgentSpecsRequest) SetOrderBy(v string) *ListAgentSpecsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAgentSpecsRequest) SetOwner(v string) *ListAgentSpecsRequest {
	s.Owner = &v
	return s
}

func (s *ListAgentSpecsRequest) SetPageNo(v int32) *ListAgentSpecsRequest {
	s.PageNo = &v
	return s
}

func (s *ListAgentSpecsRequest) SetPageSize(v int32) *ListAgentSpecsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentSpecsRequest) SetScope(v string) *ListAgentSpecsRequest {
	s.Scope = &v
	return s
}

func (s *ListAgentSpecsRequest) SetSearch(v string) *ListAgentSpecsRequest {
	s.Search = &v
	return s
}

func (s *ListAgentSpecsRequest) SetWithCapabilities(v bool) *ListAgentSpecsRequest {
	s.WithCapabilities = &v
	return s
}

func (s *ListAgentSpecsRequest) Validate() error {
	return dara.Validate(s)
}
