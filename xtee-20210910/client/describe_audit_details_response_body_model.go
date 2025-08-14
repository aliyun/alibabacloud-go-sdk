// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAuditDetailsResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeAuditDetailsResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeAuditDetailsResponseBody
	GetMessage() *string
	SetResultObject(v bool) *DescribeAuditDetailsResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DescribeAuditDetailsResponseBody
	GetSuccess() *bool
}

type DescribeAuditDetailsResponseBody struct {
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
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Whether the call was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeAuditDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAuditDetailsResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeAuditDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAuditDetailsResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeAuditDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAuditDetailsResponseBody) SetCode(v string) *DescribeAuditDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAuditDetailsResponseBody) SetHttpStatusCode(v string) *DescribeAuditDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAuditDetailsResponseBody) SetMessage(v string) *DescribeAuditDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAuditDetailsResponseBody) SetResultObject(v bool) *DescribeAuditDetailsResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeAuditDetailsResponseBody) SetSuccess(v bool) *DescribeAuditDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAuditDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
