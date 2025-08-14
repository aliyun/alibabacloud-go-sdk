// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventCountResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeEventCountResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeEventCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventCountResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeEventCountResponseBodyResultObject) *DescribeEventCountResponseBody
	GetResultObject() *DescribeEventCountResponseBodyResultObject
	SetSuccess(v bool) *DescribeEventCountResponseBody
	GetSuccess() *bool
}

type DescribeEventCountResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeEventCountResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeEventCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventCountResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeEventCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventCountResponseBody) GetResultObject() *DescribeEventCountResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventCountResponseBody) SetCode(v string) *DescribeEventCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventCountResponseBody) SetHttpStatusCode(v string) *DescribeEventCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEventCountResponseBody) SetMessage(v string) *DescribeEventCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventCountResponseBody) SetRequestId(v string) *DescribeEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventCountResponseBody) SetResultObject(v *DescribeEventCountResponseBodyResultObject) *DescribeEventCountResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventCountResponseBody) SetSuccess(v bool) *DescribeEventCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventCountResponseBodyResultObject struct {
	// Whether it exceeds the maximum number
	//
	// example:
	//
	// true
	Limit *bool `json:"limit,omitempty" xml:"limit,omitempty"`
	// Maximum creation count
	//
	// example:
	//
	// 100
	MaxTotalItem *int32 `json:"maxTotalItem,omitempty" xml:"maxTotalItem,omitempty"`
	// Total count
	//
	// example:
	//
	// 101
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
}

func (s DescribeEventCountResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventCountResponseBodyResultObject) GetLimit() *bool {
	return s.Limit
}

func (s *DescribeEventCountResponseBodyResultObject) GetMaxTotalItem() *int32 {
	return s.MaxTotalItem
}

func (s *DescribeEventCountResponseBodyResultObject) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeEventCountResponseBodyResultObject) SetLimit(v bool) *DescribeEventCountResponseBodyResultObject {
	s.Limit = &v
	return s
}

func (s *DescribeEventCountResponseBodyResultObject) SetMaxTotalItem(v int32) *DescribeEventCountResponseBodyResultObject {
	s.MaxTotalItem = &v
	return s
}

func (s *DescribeEventCountResponseBodyResultObject) SetTotalItem(v int32) *DescribeEventCountResponseBodyResultObject {
	s.TotalItem = &v
	return s
}

func (s *DescribeEventCountResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
