// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlashSmsAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFlashSmsAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *UpdateFlashSmsAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateFlashSmsAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateFlashSmsAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateFlashSmsAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateFlashSmsAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFlashSmsAccessProfileResponseBody
	GetSuccess() *bool
}

type UpdateFlashSmsAccessProfileResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The flash message configuration ID.
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
	// Instance does not exist. Instance=outb003
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFlashSmsAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlashSmsAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlashSmsAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetCode(v string) *UpdateFlashSmsAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetData(v string) *UpdateFlashSmsAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetHttpStatusCode(v int32) *UpdateFlashSmsAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetMessage(v string) *UpdateFlashSmsAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetParams(v []*string) *UpdateFlashSmsAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetRequestId(v string) *UpdateFlashSmsAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) SetSuccess(v bool) *UpdateFlashSmsAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
