// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceBindingsOfNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceBindingsOfNumberResponseBody
	GetCode() *string
	SetData(v string) *UpdateInstanceBindingsOfNumberResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateInstanceBindingsOfNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstanceBindingsOfNumberResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateInstanceBindingsOfNumberResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateInstanceBindingsOfNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceBindingsOfNumberResponseBody
	GetSuccess() *bool
}

type UpdateInstanceBindingsOfNumberResponseBody struct {
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
	// Data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceBindingsOfNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceBindingsOfNumberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetCode(v string) *UpdateInstanceBindingsOfNumberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetData(v string) *UpdateInstanceBindingsOfNumberResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceBindingsOfNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetMessage(v string) *UpdateInstanceBindingsOfNumberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetParams(v []*string) *UpdateInstanceBindingsOfNumberResponseBody {
	s.Params = v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetRequestId(v string) *UpdateInstanceBindingsOfNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) SetSuccess(v bool) *UpdateInstanceBindingsOfNumberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
