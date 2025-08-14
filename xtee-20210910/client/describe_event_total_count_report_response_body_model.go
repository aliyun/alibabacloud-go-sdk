// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTotalCountReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventTotalCountReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeEventTotalCountReportResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeEventTotalCountReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventTotalCountReportResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeEventTotalCountReportResponseBodyResultObject) *DescribeEventTotalCountReportResponseBody
	GetResultObject() *DescribeEventTotalCountReportResponseBodyResultObject
	SetSuccess(v bool) *DescribeEventTotalCountReportResponseBody
	GetSuccess() *bool
}

type DescribeEventTotalCountReportResponseBody struct {
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
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeEventTotalCountReportResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Indicates whether this operation was successful, `true` means success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeEventTotalCountReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTotalCountReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTotalCountReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventTotalCountReportResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeEventTotalCountReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventTotalCountReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventTotalCountReportResponseBody) GetResultObject() *DescribeEventTotalCountReportResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventTotalCountReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventTotalCountReportResponseBody) SetCode(v string) *DescribeEventTotalCountReportResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBody) SetHttpStatusCode(v string) *DescribeEventTotalCountReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBody) SetMessage(v string) *DescribeEventTotalCountReportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBody) SetRequestId(v string) *DescribeEventTotalCountReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBody) SetResultObject(v *DescribeEventTotalCountReportResponseBodyResultObject) *DescribeEventTotalCountReportResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventTotalCountReportResponseBody) SetSuccess(v bool) *DescribeEventTotalCountReportResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventTotalCountReportResponseBodyResultObject struct {
	// Comparison with yesterday\\"s event invocation count
	//
	// example:
	//
	// 101
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Today\\"s event invocation count
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeEventTotalCountReportResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTotalCountReportResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventTotalCountReportResponseBodyResultObject) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeEventTotalCountReportResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeEventTotalCountReportResponseBodyResultObject) SetRatio(v string) *DescribeEventTotalCountReportResponseBodyResultObject {
	s.Ratio = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBodyResultObject) SetValue(v string) *DescribeEventTotalCountReportResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeEventTotalCountReportResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
