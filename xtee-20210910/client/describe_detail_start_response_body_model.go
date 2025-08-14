// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetailStartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDetailStartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeDetailStartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeDetailStartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDetailStartResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeDetailStartResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DescribeDetailStartResponseBody
	GetSuccess() *bool
}

type DescribeDetailStartResponseBody struct {
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
	// Error details
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
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Indicates whether the operation was successful, with true representing success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeDetailStartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetailStartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDetailStartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDetailStartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeDetailStartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDetailStartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDetailStartResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeDetailStartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDetailStartResponseBody) SetCode(v string) *DescribeDetailStartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDetailStartResponseBody) SetHttpStatusCode(v string) *DescribeDetailStartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDetailStartResponseBody) SetMessage(v string) *DescribeDetailStartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDetailStartResponseBody) SetRequestId(v string) *DescribeDetailStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDetailStartResponseBody) SetResultObject(v bool) *DescribeDetailStartResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeDetailStartResponseBody) SetSuccess(v bool) *DescribeDetailStartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDetailStartResponseBody) Validate() error {
	return dara.Validate(s)
}
