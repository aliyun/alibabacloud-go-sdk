// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeServiceAppKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeServiceAppKeyResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeServiceAppKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeServiceAppKeyResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeServiceAppKeyResponseBody
	GetResultObject() *bool
	SetUccess(v bool) *DescribeServiceAppKeyResponseBody
	GetUccess() *bool
}

type DescribeServiceAppKeyResponseBody struct {
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
	// Return object
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
	Uccess *bool `json:"uccess,omitempty" xml:"uccess,omitempty"`
}

func (s DescribeServiceAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceAppKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeServiceAppKeyResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeServiceAppKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceAppKeyResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeServiceAppKeyResponseBody) GetUccess() *bool {
	return s.Uccess
}

func (s *DescribeServiceAppKeyResponseBody) SetCode(v string) *DescribeServiceAppKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceAppKeyResponseBody) SetHttpStatusCode(v string) *DescribeServiceAppKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeServiceAppKeyResponseBody) SetMessage(v string) *DescribeServiceAppKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceAppKeyResponseBody) SetRequestId(v string) *DescribeServiceAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceAppKeyResponseBody) SetResultObject(v bool) *DescribeServiceAppKeyResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeServiceAppKeyResponseBody) SetUccess(v bool) *DescribeServiceAppKeyResponseBody {
	s.Uccess = &v
	return s
}

func (s *DescribeServiceAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
