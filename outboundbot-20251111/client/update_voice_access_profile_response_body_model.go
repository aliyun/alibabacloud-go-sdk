// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVoiceAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVoiceAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *UpdateVoiceAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateVoiceAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVoiceAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateVoiceAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateVoiceAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVoiceAccessProfileResponseBody
	GetSuccess() *bool
}

type UpdateVoiceAccessProfileResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Instance does not exist. Instance=52VzexWY
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateVoiceAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVoiceAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVoiceAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVoiceAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateVoiceAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVoiceAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVoiceAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateVoiceAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVoiceAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVoiceAccessProfileResponseBody) SetCode(v string) *UpdateVoiceAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetData(v string) *UpdateVoiceAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetHttpStatusCode(v int32) *UpdateVoiceAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetMessage(v string) *UpdateVoiceAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetParams(v []*string) *UpdateVoiceAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetRequestId(v string) *UpdateVoiceAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetSuccess(v bool) *UpdateVoiceAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
