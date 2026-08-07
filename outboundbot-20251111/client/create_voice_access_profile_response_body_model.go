// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVoiceAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *CreateVoiceAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateVoiceAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVoiceAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateVoiceAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateVoiceAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVoiceAccessProfileResponseBody
	GetSuccess() *bool
}

type CreateVoiceAccessProfileResponseBody struct {
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
	// Instance does not exist. Instance=af81a389-91f0-4157-8d82-720edd02b66a.
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

func (s CreateVoiceAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVoiceAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVoiceAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateVoiceAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVoiceAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVoiceAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateVoiceAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVoiceAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVoiceAccessProfileResponseBody) SetCode(v string) *CreateVoiceAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetData(v string) *CreateVoiceAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetHttpStatusCode(v int32) *CreateVoiceAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetMessage(v string) *CreateVoiceAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetParams(v []*string) *CreateVoiceAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetRequestId(v string) *CreateVoiceAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetSuccess(v bool) *CreateVoiceAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
