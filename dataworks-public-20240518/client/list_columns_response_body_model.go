// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListColumnsResponseBodyPagingInfo) *ListColumnsResponseBody
	GetPagingInfo() *ListColumnsResponseBodyPagingInfo
	SetRequestId(v string) *ListColumnsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListColumnsResponseBody
	GetSuccess() *bool
}

type ListColumnsResponseBody struct {
	// The pagination result.
	PagingInfo *ListColumnsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D1E2E5BC-xxxx-xxxx-xxxx-xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBody) GetPagingInfo() *ListColumnsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListColumnsResponseBody) SetPagingInfo(v *ListColumnsResponseBodyPagingInfo) *ListColumnsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListColumnsResponseBody) SetRequestId(v string) *ListColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListColumnsResponseBody) SetSuccess(v bool) *ListColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *ListColumnsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListColumnsResponseBodyPagingInfo struct {
	// The columns in the table.
	Columns []*Column `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
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
	// The total number of records returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListColumnsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBodyPagingInfo) GetColumns() []*Column {
	return s.Columns
}

func (s *ListColumnsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListColumnsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListColumnsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListColumnsResponseBodyPagingInfo) SetColumns(v []*Column) *ListColumnsResponseBodyPagingInfo {
	s.Columns = v
	return s
}

func (s *ListColumnsResponseBodyPagingInfo) SetPageNumber(v int32) *ListColumnsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListColumnsResponseBodyPagingInfo) SetPageSize(v int32) *ListColumnsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListColumnsResponseBodyPagingInfo) SetTotalCount(v int64) *ListColumnsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListColumnsResponseBodyPagingInfo) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
