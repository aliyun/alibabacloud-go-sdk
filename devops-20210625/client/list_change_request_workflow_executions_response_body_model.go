// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestWorkflowExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListChangeRequestWorkflowExecutionsResponseBody
	GetCurrent() *int64
	SetPageSize(v int64) *ListChangeRequestWorkflowExecutionsResponseBody
	GetPageSize() *int64
	SetPages(v int64) *ListChangeRequestWorkflowExecutionsResponseBody
	GetPages() *int64
	SetRecords(v []interface{}) *ListChangeRequestWorkflowExecutionsResponseBody
	GetRecords() []interface{}
	SetTotal(v int64) *ListChangeRequestWorkflowExecutionsResponseBody
	GetTotal() *int64
}

type ListChangeRequestWorkflowExecutionsResponseBody struct {
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	Pages   *int64        `json:"pages,omitempty" xml:"pages,omitempty"`
	Records []interface{} `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListChangeRequestWorkflowExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestWorkflowExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) GetCurrent() *int64 {
	return s.Current
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) GetPages() *int64 {
	return s.Pages
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) GetRecords() []interface{} {
	return s.Records
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) SetCurrent(v int64) *ListChangeRequestWorkflowExecutionsResponseBody {
	s.Current = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) SetPageSize(v int64) *ListChangeRequestWorkflowExecutionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) SetPages(v int64) *ListChangeRequestWorkflowExecutionsResponseBody {
	s.Pages = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) SetRecords(v []interface{}) *ListChangeRequestWorkflowExecutionsResponseBody {
	s.Records = v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) SetTotal(v int64) *ListChangeRequestWorkflowExecutionsResponseBody {
	s.Total = &v
	return s
}

func (s *ListChangeRequestWorkflowExecutionsResponseBody) Validate() error {
	return dara.Validate(s)
}
