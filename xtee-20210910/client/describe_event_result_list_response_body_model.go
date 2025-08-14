// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventResultListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventResultListResponseBody
	GetCode() *string
	SetCurrentPage(v int64) *DescribeEventResultListResponseBody
	GetCurrentPage() *int64
	SetHttpStatusCode(v string) *DescribeEventResultListResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeEventResultListResponseBody
	GetMessage() *string
	SetPageSize(v int64) *DescribeEventResultListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeEventResultListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeEventResultListResponseBodyResultObject) *DescribeEventResultListResponseBody
	GetResultObject() []*DescribeEventResultListResponseBodyResultObject
	SetSuccess(v bool) *DescribeEventResultListResponseBody
	GetSuccess() *bool
	SetTotalItem(v int64) *DescribeEventResultListResponseBody
	GetTotalItem() *int64
	SetTotalPage(v int64) *DescribeEventResultListResponseBody
	GetTotalPage() *int64
}

type DescribeEventResultListResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error details
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeEventResultListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Whether the query was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 31
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 9
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeEventResultListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventResultListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventResultListResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeEventResultListResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeEventResultListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventResultListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventResultListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventResultListResponseBody) GetResultObject() []*DescribeEventResultListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventResultListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventResultListResponseBody) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeEventResultListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeEventResultListResponseBody) SetCode(v string) *DescribeEventResultListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetCurrentPage(v int64) *DescribeEventResultListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetHttpStatusCode(v string) *DescribeEventResultListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetMessage(v string) *DescribeEventResultListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetPageSize(v int64) *DescribeEventResultListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetRequestId(v string) *DescribeEventResultListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetResultObject(v []*DescribeEventResultListResponseBodyResultObject) *DescribeEventResultListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventResultListResponseBody) SetSuccess(v bool) *DescribeEventResultListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetTotalItem(v int64) *DescribeEventResultListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeEventResultListResponseBody) SetTotalPage(v int64) *DescribeEventResultListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeEventResultListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventResultListResponseBodyResultObject struct {
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Number of passed checks.
	//
	// example:
	//
	// 90
	PassNum *int64 `json:"passNum,omitempty" xml:"passNum,omitempty"`
	// Number of pending items.
	//
	// example:
	//
	// 5
	PendingNum *int64 `json:"pendingNum,omitempty" xml:"pendingNum,omitempty"`
	// Number of rejected approvals.
	//
	// example:
	//
	// 5
	RejectNum *int64 `json:"rejectNum,omitempty" xml:"rejectNum,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 100
	TotalNum *int64 `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s DescribeEventResultListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventResultListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventResultListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventResultListResponseBodyResultObject) GetPassNum() *int64 {
	return s.PassNum
}

func (s *DescribeEventResultListResponseBodyResultObject) GetPendingNum() *int64 {
	return s.PendingNum
}

func (s *DescribeEventResultListResponseBodyResultObject) GetRejectNum() *int64 {
	return s.RejectNum
}

func (s *DescribeEventResultListResponseBodyResultObject) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeEventResultListResponseBodyResultObject) SetEventCode(v string) *DescribeEventResultListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeEventResultListResponseBodyResultObject) SetEventName(v string) *DescribeEventResultListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeEventResultListResponseBodyResultObject) SetPassNum(v int64) *DescribeEventResultListResponseBodyResultObject {
	s.PassNum = &v
	return s
}

func (s *DescribeEventResultListResponseBodyResultObject) SetPendingNum(v int64) *DescribeEventResultListResponseBodyResultObject {
	s.PendingNum = &v
	return s
}

func (s *DescribeEventResultListResponseBodyResultObject) SetRejectNum(v int64) *DescribeEventResultListResponseBodyResultObject {
	s.RejectNum = &v
	return s
}

func (s *DescribeEventResultListResponseBodyResultObject) SetTotalNum(v int64) *DescribeEventResultListResponseBodyResultObject {
	s.TotalNum = &v
	return s
}

func (s *DescribeEventResultListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
