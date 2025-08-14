// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBasicStartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeBasicStartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeBasicStartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeBasicStartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBasicStartResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeBasicStartResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DescribeBasicStartResponseBody
	GetSuccess() *bool
}

type DescribeBasicStartResponseBody struct {
	// Status code.
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
	// Return message
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
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeBasicStartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicStartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBasicStartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBasicStartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeBasicStartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBasicStartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBasicStartResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeBasicStartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBasicStartResponseBody) SetCode(v string) *DescribeBasicStartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBasicStartResponseBody) SetHttpStatusCode(v string) *DescribeBasicStartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBasicStartResponseBody) SetMessage(v string) *DescribeBasicStartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBasicStartResponseBody) SetRequestId(v string) *DescribeBasicStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBasicStartResponseBody) SetResultObject(v bool) *DescribeBasicStartResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeBasicStartResponseBody) SetSuccess(v bool) *DescribeBasicStartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBasicStartResponseBody) Validate() error {
	return dara.Validate(s)
}
