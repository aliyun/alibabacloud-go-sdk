// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableListByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeTableListByTypeResponseBodyList) *DescribeTableListByTypeResponseBody
	GetList() []*DescribeTableListByTypeResponseBodyList
	SetPageNumber(v int32) *DescribeTableListByTypeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTableListByTypeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTableListByTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTableListByTypeResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeTableListByTypeResponseBody
	GetTotal() *int32
}

type DescribeTableListByTypeResponseBody struct {
	// Indicates the information about tables.
	List []*DescribeTableListByTypeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	//
	// example:
	//
	// B360F47B-59E3-4D1C-BA03-6BFB1C993F88
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of returned tables.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTableListByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableListByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBody) GetList() []*DescribeTableListByTypeResponseBodyList {
	return s.List
}

func (s *DescribeTableListByTypeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTableListByTypeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTableListByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableListByTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTableListByTypeResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeTableListByTypeResponseBody) SetList(v []*DescribeTableListByTypeResponseBodyList) *DescribeTableListByTypeResponseBody {
	s.List = v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageNumber(v int32) *DescribeTableListByTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetPageSize(v int32) *DescribeTableListByTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetRequestId(v string) *DescribeTableListByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetSuccess(v bool) *DescribeTableListByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) SetTotal(v int32) *DescribeTableListByTypeResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTableListByTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTableListByTypeResponseBodyList struct {
	// Indicates the property of a table.
	//
	// example:
	//
	// single
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// Indicates the name of the table.
	//
	// example:
	//
	// employee_split
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableListByTypeResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableListByTypeResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponseBodyList) GetProperty() *string {
	return s.Property
}

func (s *DescribeTableListByTypeResponseBodyList) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableListByTypeResponseBodyList) SetProperty(v string) *DescribeTableListByTypeResponseBodyList {
	s.Property = &v
	return s
}

func (s *DescribeTableListByTypeResponseBodyList) SetTableName(v string) *DescribeTableListByTypeResponseBodyList {
	s.TableName = &v
	return s
}

func (s *DescribeTableListByTypeResponseBodyList) Validate() error {
	return dara.Validate(s)
}
