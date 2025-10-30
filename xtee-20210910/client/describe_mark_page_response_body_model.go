// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarkPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeMarkPageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeMarkPageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeMarkPageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeMarkPageResponseBodyResultObject) *DescribeMarkPageResponseBody
	GetResultObject() []*DescribeMarkPageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeMarkPageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeMarkPageResponseBody
	GetTotalPage() *int32
}

type DescribeMarkPageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeMarkPageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeMarkPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarkPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMarkPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMarkPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeMarkPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMarkPageResponseBody) GetResultObject() []*DescribeMarkPageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeMarkPageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeMarkPageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeMarkPageResponseBody) SetRequestId(v string) *DescribeMarkPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMarkPageResponseBody) SetCurrentPage(v int32) *DescribeMarkPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeMarkPageResponseBody) SetPageSize(v int32) *DescribeMarkPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMarkPageResponseBody) SetResultObject(v []*DescribeMarkPageResponseBodyResultObject) *DescribeMarkPageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeMarkPageResponseBody) SetTotalItem(v int32) *DescribeMarkPageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeMarkPageResponseBody) SetTotalPage(v int32) *DescribeMarkPageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeMarkPageResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMarkPageResponseBodyResultObject struct {
	// Field name.
	//
	// example:
	//
	// mobile
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field value.
	//
	// example:
	//
	// 18000000000
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// Primary key ID.
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Mark (0 No / 1 Yes).
	//
	// example:
	//
	// 1
	MarkType *string `json:"markType,omitempty" xml:"markType,omitempty"`
}

func (s DescribeMarkPageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarkPageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeMarkPageResponseBodyResultObject) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeMarkPageResponseBodyResultObject) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeMarkPageResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeMarkPageResponseBodyResultObject) GetMarkType() *string {
	return s.MarkType
}

func (s *DescribeMarkPageResponseBodyResultObject) SetFieldName(v string) *DescribeMarkPageResponseBodyResultObject {
	s.FieldName = &v
	return s
}

func (s *DescribeMarkPageResponseBodyResultObject) SetFieldValue(v string) *DescribeMarkPageResponseBodyResultObject {
	s.FieldValue = &v
	return s
}

func (s *DescribeMarkPageResponseBodyResultObject) SetId(v int64) *DescribeMarkPageResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeMarkPageResponseBodyResultObject) SetMarkType(v string) *DescribeMarkPageResponseBodyResultObject {
	s.MarkType = &v
	return s
}

func (s *DescribeMarkPageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
