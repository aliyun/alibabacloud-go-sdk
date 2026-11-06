// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLlmAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLlmAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *CreateLlmAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateLlmAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateLlmAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateLlmAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateLlmAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLlmAccessProfileResponseBody
	GetSuccess() *bool
}

type CreateLlmAccessProfileResponseBody struct {
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

func (s CreateLlmAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLlmAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLlmAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLlmAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateLlmAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateLlmAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLlmAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateLlmAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLlmAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLlmAccessProfileResponseBody) SetCode(v string) *CreateLlmAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetData(v string) *CreateLlmAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetHttpStatusCode(v int32) *CreateLlmAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetMessage(v string) *CreateLlmAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetParams(v []*string) *CreateLlmAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetRequestId(v string) *CreateLlmAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetSuccess(v bool) *CreateLlmAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
