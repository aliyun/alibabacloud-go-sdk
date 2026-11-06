// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNumberBindingsOfInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNumberBindingsOfInstanceResponseBody
	GetCode() *string
	SetData(v string) *UpdateNumberBindingsOfInstanceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateNumberBindingsOfInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateNumberBindingsOfInstanceResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateNumberBindingsOfInstanceResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateNumberBindingsOfInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNumberBindingsOfInstanceResponseBody
	GetSuccess() *bool
}

type UpdateNumberBindingsOfInstanceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNumberBindingsOfInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNumberBindingsOfInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetCode(v string) *UpdateNumberBindingsOfInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetData(v string) *UpdateNumberBindingsOfInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateNumberBindingsOfInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetMessage(v string) *UpdateNumberBindingsOfInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetParams(v []*string) *UpdateNumberBindingsOfInstanceResponseBody {
	s.Params = v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetRequestId(v string) *UpdateNumberBindingsOfInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) SetSuccess(v bool) *UpdateNumberBindingsOfInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
