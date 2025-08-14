// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAuditConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeAuditConfigResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeAuditConfigResponseBody
	GetMessage() *string
	SetResultObject(v bool) *DescribeAuditConfigResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DescribeAuditConfigResponseBody
	GetSuccess() *bool
}

type DescribeAuditConfigResponseBody struct {
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
	// Error message
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeAuditConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAuditConfigResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeAuditConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAuditConfigResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeAuditConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAuditConfigResponseBody) SetCode(v string) *DescribeAuditConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAuditConfigResponseBody) SetHttpStatusCode(v string) *DescribeAuditConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAuditConfigResponseBody) SetMessage(v string) *DescribeAuditConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAuditConfigResponseBody) SetResultObject(v bool) *DescribeAuditConfigResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeAuditConfigResponseBody) SetSuccess(v bool) *DescribeAuditConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAuditConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
