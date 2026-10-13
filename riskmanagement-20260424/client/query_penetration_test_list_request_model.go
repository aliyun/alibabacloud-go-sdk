// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPenetrationTestListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryPenetrationTestListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryPenetrationTestListRequest
	GetPageSize() *int32
}

type QueryPenetrationTestListRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPenetrationTestListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestListRequest) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPenetrationTestListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPenetrationTestListRequest) SetCurrentPage(v int32) *QueryPenetrationTestListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPenetrationTestListRequest) SetPageSize(v int32) *QueryPenetrationTestListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPenetrationTestListRequest) Validate() error {
	return dara.Validate(s)
}
