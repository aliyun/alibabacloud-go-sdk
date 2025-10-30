// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvanceSearchPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAdvanceSearchPageListResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeAdvanceSearchPageListResponseBodyResultObject) *DescribeAdvanceSearchPageListResponseBody
	GetResultObject() *DescribeAdvanceSearchPageListResponseBodyResultObject
}

type DescribeAdvanceSearchPageListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeAdvanceSearchPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeAdvanceSearchPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvanceSearchPageListResponseBody) GetResultObject() *DescribeAdvanceSearchPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAdvanceSearchPageListResponseBody) SetRequestId(v string) *DescribeAdvanceSearchPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBody) SetResultObject(v *DescribeAdvanceSearchPageListResponseBodyResultObject) *DescribeAdvanceSearchPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvanceSearchPageListResponseBodyResultObject struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Returned data object
	Data []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Table header
	Header []*DescribeAdvanceSearchPageListResponseBodyResultObjectHeader `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// Number of items per page. Default value: 20, minimum value: 1, maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 37
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 4
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeAdvanceSearchPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) GetData() []map[string]interface{} {
	return s.Data
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) GetHeader() []*DescribeAdvanceSearchPageListResponseBodyResultObjectHeader {
	return s.Header
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) SetCurrentPage(v int64) *DescribeAdvanceSearchPageListResponseBodyResultObject {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) SetData(v []map[string]interface{}) *DescribeAdvanceSearchPageListResponseBodyResultObject {
	s.Data = v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) SetHeader(v []*DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) *DescribeAdvanceSearchPageListResponseBodyResultObject {
	s.Header = v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) SetPageSize(v int64) *DescribeAdvanceSearchPageListResponseBodyResultObject {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) SetTotalItem(v int64) *DescribeAdvanceSearchPageListResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) SetTotalPage(v int64) *DescribeAdvanceSearchPageListResponseBodyResultObject {
	s.TotalPage = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObject) Validate() error {
	if s.Header != nil {
		for _, item := range s.Header {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvanceSearchPageListResponseBodyResultObjectHeader struct {
	// Field name
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field title.
	//
	// example:
	//
	// 年龄
	FieldTitle *string `json:"fieldTitle,omitempty" xml:"fieldTitle,omitempty"`
	// Whether it is a default display field (displayed in the response, not used as a parameter)
	//
	// - true: Yes
	//
	// - false: No
	//
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
}

func (s DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) GetFieldTitle() *string {
	return s.FieldTitle
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) SetFieldName(v string) *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader {
	s.FieldName = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) SetFieldTitle(v string) *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader {
	s.FieldTitle = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) SetIsDefault(v bool) *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader {
	s.IsDefault = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponseBodyResultObjectHeader) Validate() error {
	return dara.Validate(s)
}
