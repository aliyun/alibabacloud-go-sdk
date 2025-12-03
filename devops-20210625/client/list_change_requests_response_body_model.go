// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListChangeRequestsResponseBody
	GetCurrent() *int64
	SetData(v []interface{}) *ListChangeRequestsResponseBody
	GetData() []interface{}
	SetNextToken(v string) *ListChangeRequestsResponseBody
	GetNextToken() *string
	SetPages(v int64) *ListChangeRequestsResponseBody
	GetPages() *int64
	SetPerPage(v int64) *ListChangeRequestsResponseBody
	GetPerPage() *int64
	SetTotal(v int64) *ListChangeRequestsResponseBody
	GetTotal() *int64
}

type ListChangeRequestsResponseBody struct {
	// example:
	//
	// 1
	Current *int64        `json:"current,omitempty" xml:"current,omitempty"`
	Data    []interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// eb13ac6049d3d78159d60f84af
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 5
	Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// example:
	//
	// 20
	PerPage *int64 `json:"perPage,omitempty" xml:"perPage,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListChangeRequestsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeRequestsResponseBody) GetCurrent() *int64 {
	return s.Current
}

func (s *ListChangeRequestsResponseBody) GetData() []interface{} {
	return s.Data
}

func (s *ListChangeRequestsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChangeRequestsResponseBody) GetPages() *int64 {
	return s.Pages
}

func (s *ListChangeRequestsResponseBody) GetPerPage() *int64 {
	return s.PerPage
}

func (s *ListChangeRequestsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListChangeRequestsResponseBody) SetCurrent(v int64) *ListChangeRequestsResponseBody {
	s.Current = &v
	return s
}

func (s *ListChangeRequestsResponseBody) SetData(v []interface{}) *ListChangeRequestsResponseBody {
	s.Data = v
	return s
}

func (s *ListChangeRequestsResponseBody) SetNextToken(v string) *ListChangeRequestsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListChangeRequestsResponseBody) SetPages(v int64) *ListChangeRequestsResponseBody {
	s.Pages = &v
	return s
}

func (s *ListChangeRequestsResponseBody) SetPerPage(v int64) *ListChangeRequestsResponseBody {
	s.PerPage = &v
	return s
}

func (s *ListChangeRequestsResponseBody) SetTotal(v int64) *ListChangeRequestsResponseBody {
	s.Total = &v
	return s
}

func (s *ListChangeRequestsResponseBody) Validate() error {
	return dara.Validate(s)
}
