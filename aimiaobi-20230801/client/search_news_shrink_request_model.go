// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNewsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SearchNewsShrinkRequest
	GetAgentKey() *string
	SetFilterNotNull(v bool) *SearchNewsShrinkRequest
	GetFilterNotNull() *bool
	SetIncludeContent(v bool) *SearchNewsShrinkRequest
	GetIncludeContent() *bool
	SetPage(v int32) *SearchNewsShrinkRequest
	GetPage() *int32
	SetPageSize(v int32) *SearchNewsShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *SearchNewsShrinkRequest
	GetQuery() *string
	SetSearchSourcesShrink(v string) *SearchNewsShrinkRequest
	GetSearchSourcesShrink() *string
}

type SearchNewsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// false
	FilterNotNull *bool `json:"FilterNotNull,omitempty" xml:"FilterNotNull,omitempty"`
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// example:
	//
	// 81
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 35
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 检索Query
	Query               *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchSourcesShrink *string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty"`
}

func (s SearchNewsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchNewsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchNewsShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SearchNewsShrinkRequest) GetFilterNotNull() *bool {
	return s.FilterNotNull
}

func (s *SearchNewsShrinkRequest) GetIncludeContent() *bool {
	return s.IncludeContent
}

func (s *SearchNewsShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *SearchNewsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchNewsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchNewsShrinkRequest) GetSearchSourcesShrink() *string {
	return s.SearchSourcesShrink
}

func (s *SearchNewsShrinkRequest) SetAgentKey(v string) *SearchNewsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetFilterNotNull(v bool) *SearchNewsShrinkRequest {
	s.FilterNotNull = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetIncludeContent(v bool) *SearchNewsShrinkRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetPage(v int32) *SearchNewsShrinkRequest {
	s.Page = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetPageSize(v int32) *SearchNewsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetQuery(v string) *SearchNewsShrinkRequest {
	s.Query = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetSearchSourcesShrink(v string) *SearchNewsShrinkRequest {
	s.SearchSourcesShrink = &v
	return s
}

func (s *SearchNewsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
