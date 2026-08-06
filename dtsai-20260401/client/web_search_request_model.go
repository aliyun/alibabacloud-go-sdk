// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *WebSearchRequest
	GetMaxResults() *int32
	SetQuery(v string) *WebSearchRequest
	GetQuery() *string
	SetRegionId(v string) *WebSearchRequest
	GetRegionId() *string
	SetUrlScopeDomains(v string) *WebSearchRequest
	GetUrlScopeDomains() *string
	SetUrlScopeMode(v string) *WebSearchRequest
	GetUrlScopeMode() *string
}

type WebSearchRequest struct {
	// The maximum number of results to return. Default value: 10. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The search query statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// Spring Boot
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UrlScopeDomains *string `json:"UrlScopeDomains,omitempty" xml:"UrlScopeDomains,omitempty"`
	UrlScopeMode    *string `json:"UrlScopeMode,omitempty" xml:"UrlScopeMode,omitempty"`
}

func (s WebSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s WebSearchRequest) GoString() string {
	return s.String()
}

func (s *WebSearchRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *WebSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *WebSearchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WebSearchRequest) GetUrlScopeDomains() *string {
	return s.UrlScopeDomains
}

func (s *WebSearchRequest) GetUrlScopeMode() *string {
	return s.UrlScopeMode
}

func (s *WebSearchRequest) SetMaxResults(v int32) *WebSearchRequest {
	s.MaxResults = &v
	return s
}

func (s *WebSearchRequest) SetQuery(v string) *WebSearchRequest {
	s.Query = &v
	return s
}

func (s *WebSearchRequest) SetRegionId(v string) *WebSearchRequest {
	s.RegionId = &v
	return s
}

func (s *WebSearchRequest) SetUrlScopeDomains(v string) *WebSearchRequest {
	s.UrlScopeDomains = &v
	return s
}

func (s *WebSearchRequest) SetUrlScopeMode(v string) *WebSearchRequest {
	s.UrlScopeMode = &v
	return s
}

func (s *WebSearchRequest) Validate() error {
	return dara.Validate(s)
}
