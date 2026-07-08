// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SearchNewsRequest
	GetAgentKey() *string
	SetFilterNotNull(v bool) *SearchNewsRequest
	GetFilterNotNull() *bool
	SetIncludeContent(v bool) *SearchNewsRequest
	GetIncludeContent() *bool
	SetPage(v int32) *SearchNewsRequest
	GetPage() *int32
	SetPageSize(v int32) *SearchNewsRequest
	GetPageSize() *int32
	SetQuery(v string) *SearchNewsRequest
	GetQuery() *string
	SetSearchSources(v []*string) *SearchNewsRequest
	GetSearchSources() []*string
}

type SearchNewsRequest struct {
	// The unique ID of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/3027170.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Specifies whether to filter out results with empty content.
	//
	// example:
	//
	// false
	FilterNotNull *bool `json:"FilterNotNull,omitempty" xml:"FilterNotNull,omitempty"`
	// Specifies whether to include the full text of the article.
	//
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records to return on each page. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search query.
	//
	// example:
	//
	// 检索Query
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// A list of search sources.
	SearchSources []*string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s SearchNewsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchNewsRequest) GoString() string {
	return s.String()
}

func (s *SearchNewsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SearchNewsRequest) GetFilterNotNull() *bool {
	return s.FilterNotNull
}

func (s *SearchNewsRequest) GetIncludeContent() *bool {
	return s.IncludeContent
}

func (s *SearchNewsRequest) GetPage() *int32 {
	return s.Page
}

func (s *SearchNewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchNewsRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchNewsRequest) GetSearchSources() []*string {
	return s.SearchSources
}

func (s *SearchNewsRequest) SetAgentKey(v string) *SearchNewsRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchNewsRequest) SetFilterNotNull(v bool) *SearchNewsRequest {
	s.FilterNotNull = &v
	return s
}

func (s *SearchNewsRequest) SetIncludeContent(v bool) *SearchNewsRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchNewsRequest) SetPage(v int32) *SearchNewsRequest {
	s.Page = &v
	return s
}

func (s *SearchNewsRequest) SetPageSize(v int32) *SearchNewsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchNewsRequest) SetQuery(v string) *SearchNewsRequest {
	s.Query = &v
	return s
}

func (s *SearchNewsRequest) SetSearchSources(v []*string) *SearchNewsRequest {
	s.SearchSources = v
	return s
}

func (s *SearchNewsRequest) Validate() error {
	return dara.Validate(s)
}
