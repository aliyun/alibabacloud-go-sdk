// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAiAppByPageRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListAiAppByPageRequest
	GetPageSize() *int32
	SetQuery(v string) *ListAiAppByPageRequest
	GetQuery() *string
	SetRegionId(v string) *ListAiAppByPageRequest
	GetRegionId() *string
}

type ListAiAppByPageRequest struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition.
	//
	// example:
	//
	// {}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAiAppByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppByPageRequest) GoString() string {
	return s.String()
}

func (s *ListAiAppByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAiAppByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiAppByPageRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAiAppByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAiAppByPageRequest) SetCurrentPage(v int32) *ListAiAppByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAiAppByPageRequest) SetPageSize(v int32) *ListAiAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAiAppByPageRequest) SetQuery(v string) *ListAiAppByPageRequest {
	s.Query = &v
	return s
}

func (s *ListAiAppByPageRequest) SetRegionId(v string) *ListAiAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *ListAiAppByPageRequest) Validate() error {
	return dara.Validate(s)
}
