// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBasicSearchPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBasicSearchPageListResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeBasicSearchPageListResponseBodyResultObject) *DescribeBasicSearchPageListResponseBody
	GetResultObject() *DescribeBasicSearchPageListResponseBodyResultObject
}

type DescribeBasicSearchPageListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned object.
	ResultObject *DescribeBasicSearchPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeBasicSearchPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicSearchPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBasicSearchPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBasicSearchPageListResponseBody) GetResultObject() *DescribeBasicSearchPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeBasicSearchPageListResponseBody) SetRequestId(v string) *DescribeBasicSearchPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBody) SetResultObject(v *DescribeBasicSearchPageListResponseBodyResultObject) *DescribeBasicSearchPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeBasicSearchPageListResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBasicSearchPageListResponseBodyResultObject struct {
	// The current page number in the paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The returned data object.
	Data []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The table header.
	Header []*DescribeBasicSearchPageListResponseBodyResultObjectHeader `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 6
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeBasicSearchPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicSearchPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) GetData() []map[string]interface{} {
	return s.Data
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) GetHeader() []*DescribeBasicSearchPageListResponseBodyResultObjectHeader {
	return s.Header
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) SetCurrentPage(v int64) *DescribeBasicSearchPageListResponseBodyResultObject {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) SetData(v []map[string]interface{}) *DescribeBasicSearchPageListResponseBodyResultObject {
	s.Data = v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) SetHeader(v []*DescribeBasicSearchPageListResponseBodyResultObjectHeader) *DescribeBasicSearchPageListResponseBodyResultObject {
	s.Header = v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) SetPageSize(v int64) *DescribeBasicSearchPageListResponseBodyResultObject {
	s.PageSize = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) SetTotalItem(v int64) *DescribeBasicSearchPageListResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) SetTotalPage(v int64) *DescribeBasicSearchPageListResponseBodyResultObject {
	s.TotalPage = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObject) Validate() error {
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

type DescribeBasicSearchPageListResponseBodyResultObjectHeader struct {
	// The field name.
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// The field name.
	//
	// example:
	//
	// 年龄
	FieldTitle *string `json:"fieldTitle,omitempty" xml:"fieldTitle,omitempty"`
	// Indicates whether the field is displayed by default. This field is included in the response only and is not used as a request parameter. Valid values:
	//
	// - true: The field is displayed by default.
	//
	// - false: The field is not displayed by default.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
}

func (s DescribeBasicSearchPageListResponseBodyResultObjectHeader) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicSearchPageListResponseBodyResultObjectHeader) GoString() string {
	return s.String()
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) GetFieldTitle() *string {
	return s.FieldTitle
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) SetFieldName(v string) *DescribeBasicSearchPageListResponseBodyResultObjectHeader {
	s.FieldName = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) SetFieldTitle(v string) *DescribeBasicSearchPageListResponseBodyResultObjectHeader {
	s.FieldTitle = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) SetIsDefault(v bool) *DescribeBasicSearchPageListResponseBodyResultObjectHeader {
	s.IsDefault = &v
	return s
}

func (s *DescribeBasicSearchPageListResponseBodyResultObjectHeader) Validate() error {
	return dara.Validate(s)
}
