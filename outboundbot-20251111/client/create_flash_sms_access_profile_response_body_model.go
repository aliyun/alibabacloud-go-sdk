// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlashSmsAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFlashSmsAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *CreateFlashSmsAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateFlashSmsAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateFlashSmsAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateFlashSmsAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateFlashSmsAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFlashSmsAccessProfileResponseBody
	GetSuccess() *bool
}

type CreateFlashSmsAccessProfileResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// xxxxx
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
	// Instance does not exist. Instance=df408e55-63dc-4c52-9161-74265381b6a4
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

func (s CreateFlashSmsAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlashSmsAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlashSmsAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetCode(v string) *CreateFlashSmsAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetData(v string) *CreateFlashSmsAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetHttpStatusCode(v int32) *CreateFlashSmsAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetMessage(v string) *CreateFlashSmsAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetParams(v []*string) *CreateFlashSmsAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetRequestId(v string) *CreateFlashSmsAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) SetSuccess(v bool) *CreateFlashSmsAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
