// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListTablesResponseBodyPagingInfo) *ListTablesResponseBody
	GetPagingInfo() *ListTablesResponseBodyPagingInfo
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTablesResponseBody
	GetSuccess() *bool
}

type ListTablesResponseBody struct {
	PagingInfo *ListTablesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetPagingInfo() *ListTablesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTablesResponseBody) SetPagingInfo(v *ListTablesResponseBodyPagingInfo) *ListTablesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetSuccess(v bool) *ListTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tables   []*Table `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTablesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTablesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesResponseBodyPagingInfo) GetTables() []*Table {
	return s.Tables
}

func (s *ListTablesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTablesResponseBodyPagingInfo) SetPageNumber(v int32) *ListTablesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListTablesResponseBodyPagingInfo) SetPageSize(v int32) *ListTablesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListTablesResponseBodyPagingInfo) SetTables(v []*Table) *ListTablesResponseBodyPagingInfo {
	s.Tables = v
	return s
}

func (s *ListTablesResponseBodyPagingInfo) SetTotalCount(v int64) *ListTablesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListTablesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}
