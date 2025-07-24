// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *GlobalSearchRequest
	GetPage() *int32
	SetPageSize(v int32) *GlobalSearchRequest
	GetPageSize() *int32
	SetQuery(v string) *GlobalSearchRequest
	GetQuery() *string
	SetTimeRange(v string) *GlobalSearchRequest
	GetTimeRange() *string
}

type GlobalSearchRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GlobalSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalSearchRequest) GoString() string {
	return s.String()
}

func (s *GlobalSearchRequest) GetPage() *int32 {
	return s.Page
}

func (s *GlobalSearchRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GlobalSearchRequest) GetQuery() *string {
	return s.Query
}

func (s *GlobalSearchRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *GlobalSearchRequest) SetPage(v int32) *GlobalSearchRequest {
	s.Page = &v
	return s
}

func (s *GlobalSearchRequest) SetPageSize(v int32) *GlobalSearchRequest {
	s.PageSize = &v
	return s
}

func (s *GlobalSearchRequest) SetQuery(v string) *GlobalSearchRequest {
	s.Query = &v
	return s
}

func (s *GlobalSearchRequest) SetTimeRange(v string) *GlobalSearchRequest {
	s.TimeRange = &v
	return s
}

func (s *GlobalSearchRequest) Validate() error {
	return dara.Validate(s)
}
