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
}

type WebSearchRequest struct {
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *WebSearchRequest) Validate() error {
	return dara.Validate(s)
}
