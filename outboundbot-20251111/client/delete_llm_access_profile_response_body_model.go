// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLlmAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *DeleteLlmAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteLlmAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteLlmAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteLlmAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteLlmAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLlmAccessProfileResponseBody
	GetSuccess() *bool
}

type DeleteLlmAccessProfileResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The access configuration ID.
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
	// 0956D5DA-0978-5DC9-94B0-C68527DA7475
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLlmAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLlmAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLlmAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteLlmAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteLlmAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLlmAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteLlmAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLlmAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLlmAccessProfileResponseBody) SetCode(v string) *DeleteLlmAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetData(v string) *DeleteLlmAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetHttpStatusCode(v int32) *DeleteLlmAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetMessage(v string) *DeleteLlmAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetParams(v []*string) *DeleteLlmAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetRequestId(v string) *DeleteLlmAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetSuccess(v bool) *DeleteLlmAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
