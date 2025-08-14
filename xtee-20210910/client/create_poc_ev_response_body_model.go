// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocEvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePocEvResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreatePocEvResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreatePocEvResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePocEvResponseBody
	GetRequestId() *string
	SetResultObject(v string) *CreatePocEvResponseBody
	GetResultObject() *string
}

type CreatePocEvResponseBody struct {
	// Response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result
	//
	// example:
	//
	// True
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s CreatePocEvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePocEvResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePocEvResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePocEvResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreatePocEvResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePocEvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePocEvResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *CreatePocEvResponseBody) SetCode(v string) *CreatePocEvResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePocEvResponseBody) SetHttpStatusCode(v string) *CreatePocEvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePocEvResponseBody) SetMessage(v string) *CreatePocEvResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePocEvResponseBody) SetRequestId(v string) *CreatePocEvResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePocEvResponseBody) SetResultObject(v string) *CreatePocEvResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreatePocEvResponseBody) Validate() error {
	return dara.Validate(s)
}
