// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateScriptVersionResponseBody
	GetCode() *string
	SetData(v string) *UpdateScriptVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateScriptVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateScriptVersionResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateScriptVersionResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateScriptVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScriptVersionResponseBody
	GetSuccess() *bool
}

type UpdateScriptVersionResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scenario version ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b26
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
	// Instance does not exist. Instance=outb003
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

func (s UpdateScriptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScriptVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateScriptVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateScriptVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateScriptVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateScriptVersionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateScriptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScriptVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScriptVersionResponseBody) SetCode(v string) *UpdateScriptVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScriptVersionResponseBody) SetData(v string) *UpdateScriptVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateScriptVersionResponseBody) SetHttpStatusCode(v int32) *UpdateScriptVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateScriptVersionResponseBody) SetMessage(v string) *UpdateScriptVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateScriptVersionResponseBody) SetParams(v []*string) *UpdateScriptVersionResponseBody {
	s.Params = v
	return s
}

func (s *UpdateScriptVersionResponseBody) SetRequestId(v string) *UpdateScriptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScriptVersionResponseBody) SetSuccess(v bool) *UpdateScriptVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScriptVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
