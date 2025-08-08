// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Connection) *ListConnectionsResponseBody
	GetData() []*Connection
	SetPageNumber(v int64) *ListConnectionsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListConnectionsResponseBody
	GetPageSize() *int64
	SetTotalCount(v int64) *ListConnectionsResponseBody
	GetTotalCount() *int64
}

type ListConnectionsResponseBody struct {
	Data []*Connection `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBody) GetData() []*Connection {
	return s.Data
}

func (s *ListConnectionsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListConnectionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConnectionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConnectionsResponseBody) SetData(v []*Connection) *ListConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConnectionsResponseBody) SetPageNumber(v int64) *ListConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsResponseBody) SetPageSize(v int64) *ListConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConnectionsResponseBody) SetTotalCount(v int64) *ListConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}
