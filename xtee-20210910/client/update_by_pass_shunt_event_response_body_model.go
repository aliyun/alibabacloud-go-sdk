// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateByPassShuntEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateByPassShuntEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *UpdateByPassShuntEventResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UpdateByPassShuntEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateByPassShuntEventResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateByPassShuntEventResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *UpdateByPassShuntEventResponseBody
	GetSuccess() *bool
}

type UpdateByPassShuntEventResponseBody struct {
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
	// Whether it was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateByPassShuntEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateByPassShuntEventResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateByPassShuntEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateByPassShuntEventResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateByPassShuntEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateByPassShuntEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateByPassShuntEventResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateByPassShuntEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateByPassShuntEventResponseBody) SetCode(v string) *UpdateByPassShuntEventResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateByPassShuntEventResponseBody) SetHttpStatusCode(v string) *UpdateByPassShuntEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateByPassShuntEventResponseBody) SetMessage(v string) *UpdateByPassShuntEventResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateByPassShuntEventResponseBody) SetRequestId(v string) *UpdateByPassShuntEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateByPassShuntEventResponseBody) SetResultObject(v bool) *UpdateByPassShuntEventResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateByPassShuntEventResponseBody) SetSuccess(v bool) *UpdateByPassShuntEventResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateByPassShuntEventResponseBody) Validate() error {
	return dara.Validate(s)
}
