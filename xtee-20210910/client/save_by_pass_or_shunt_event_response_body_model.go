// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveByPassOrShuntEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveByPassOrShuntEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *SaveByPassOrShuntEventResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SaveByPassOrShuntEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveByPassOrShuntEventResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SaveByPassOrShuntEventResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *SaveByPassOrShuntEventResponseBody
	GetSuccess() *bool
}

type SaveByPassOrShuntEventResponseBody struct {
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveByPassOrShuntEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveByPassOrShuntEventResponseBody) GoString() string {
	return s.String()
}

func (s *SaveByPassOrShuntEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveByPassOrShuntEventResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SaveByPassOrShuntEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveByPassOrShuntEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveByPassOrShuntEventResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SaveByPassOrShuntEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveByPassOrShuntEventResponseBody) SetCode(v string) *SaveByPassOrShuntEventResponseBody {
	s.Code = &v
	return s
}

func (s *SaveByPassOrShuntEventResponseBody) SetHttpStatusCode(v string) *SaveByPassOrShuntEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveByPassOrShuntEventResponseBody) SetMessage(v string) *SaveByPassOrShuntEventResponseBody {
	s.Message = &v
	return s
}

func (s *SaveByPassOrShuntEventResponseBody) SetRequestId(v string) *SaveByPassOrShuntEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveByPassOrShuntEventResponseBody) SetResultObject(v bool) *SaveByPassOrShuntEventResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SaveByPassOrShuntEventResponseBody) SetSuccess(v bool) *SaveByPassOrShuntEventResponseBody {
	s.Success = &v
	return s
}

func (s *SaveByPassOrShuntEventResponseBody) Validate() error {
	return dara.Validate(s)
}
