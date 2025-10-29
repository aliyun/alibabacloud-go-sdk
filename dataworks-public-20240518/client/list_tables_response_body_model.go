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
	// The pagination result.
	PagingInfo *ListTablesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
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
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTablesResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of data tables.
	Tables []*Table `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The total number of records returned.
	//
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
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
