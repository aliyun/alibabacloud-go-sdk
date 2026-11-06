// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLlmAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLlmAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *UpdateLlmAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateLlmAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateLlmAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateLlmAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateLlmAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLlmAccessProfileResponseBody
	GetSuccess() *bool
}

type UpdateLlmAccessProfileResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This field is not used.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLlmAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLlmAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLlmAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLlmAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateLlmAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateLlmAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLlmAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateLlmAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLlmAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLlmAccessProfileResponseBody) SetCode(v string) *UpdateLlmAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetData(v string) *UpdateLlmAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetHttpStatusCode(v int32) *UpdateLlmAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetMessage(v string) *UpdateLlmAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetParams(v []*string) *UpdateLlmAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetRequestId(v string) *UpdateLlmAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetSuccess(v bool) *UpdateLlmAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
