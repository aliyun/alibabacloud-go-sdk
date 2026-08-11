// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScriptResponseBody
	GetCode() *string
	SetData(v string) *CreateScriptResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateScriptResponseBody
	GetSuccess() *bool
}

type CreateScriptResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scenario ID.
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
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=outb003.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScriptResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScriptResponseBody) SetCode(v string) *CreateScriptResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScriptResponseBody) SetData(v string) *CreateScriptResponseBody {
	s.Data = &v
	return s
}

func (s *CreateScriptResponseBody) SetHttpStatusCode(v int32) *CreateScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScriptResponseBody) SetMessage(v string) *CreateScriptResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScriptResponseBody) SetParams(v []*string) *CreateScriptResponseBody {
	s.Params = v
	return s
}

func (s *CreateScriptResponseBody) SetRequestId(v string) *CreateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptResponseBody) SetSuccess(v bool) *CreateScriptResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
