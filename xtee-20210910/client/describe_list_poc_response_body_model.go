// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListPocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeListPocResponseBody
	GetCode() *string
	SetCurrentPage(v string) *DescribeListPocResponseBody
	GetCurrentPage() *string
	SetHttpStatusCode(v string) *DescribeListPocResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeListPocResponseBody
	GetMessage() *string
	SetPageSize(v string) *DescribeListPocResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeListPocResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeListPocResponseBody
	GetResultObject() *bool
	SetTotalItem(v string) *DescribeListPocResponseBody
	GetTotalItem() *string
	SetTotalPage(v string) *DescribeListPocResponseBody
	GetTotalPage() *string
}

type DescribeListPocResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Total number of items returned.
	//
	// example:
	//
	// 0
	TotalItem *string `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 0
	TotalPage *string `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeListPocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListPocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListPocResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeListPocResponseBody) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeListPocResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeListPocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeListPocResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeListPocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListPocResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeListPocResponseBody) GetTotalItem() *string {
	return s.TotalItem
}

func (s *DescribeListPocResponseBody) GetTotalPage() *string {
	return s.TotalPage
}

func (s *DescribeListPocResponseBody) SetCode(v string) *DescribeListPocResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeListPocResponseBody) SetCurrentPage(v string) *DescribeListPocResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeListPocResponseBody) SetHttpStatusCode(v string) *DescribeListPocResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeListPocResponseBody) SetMessage(v string) *DescribeListPocResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeListPocResponseBody) SetPageSize(v string) *DescribeListPocResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeListPocResponseBody) SetRequestId(v string) *DescribeListPocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListPocResponseBody) SetResultObject(v bool) *DescribeListPocResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeListPocResponseBody) SetTotalItem(v string) *DescribeListPocResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeListPocResponseBody) SetTotalPage(v string) *DescribeListPocResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeListPocResponseBody) Validate() error {
	return dara.Validate(s)
}
