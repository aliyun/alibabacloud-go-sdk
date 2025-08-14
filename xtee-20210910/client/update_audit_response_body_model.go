// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAuditResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *UpdateAuditResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UpdateAuditResponseBody
	GetMessage() *string
	SetResultObject(v bool) *UpdateAuditResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *UpdateAuditResponseBody
	GetSuccess() *bool
}

type UpdateAuditResponseBody struct {
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

func (s UpdateAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuditResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAuditResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateAuditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAuditResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateAuditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAuditResponseBody) SetCode(v string) *UpdateAuditResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAuditResponseBody) SetHttpStatusCode(v string) *UpdateAuditResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAuditResponseBody) SetMessage(v string) *UpdateAuditResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAuditResponseBody) SetResultObject(v bool) *UpdateAuditResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateAuditResponseBody) SetSuccess(v bool) *UpdateAuditResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
