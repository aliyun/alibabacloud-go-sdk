// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlashSmsAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFlashSmsAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *DeleteFlashSmsAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteFlashSmsAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteFlashSmsAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteFlashSmsAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteFlashSmsAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFlashSmsAccessProfileResponseBody
	GetSuccess() *bool
}

type DeleteFlashSmsAccessProfileResponseBody struct {
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
	// 闪信配置 ID
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
	// Instance does not exist. Instance=outb003.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlashSmsAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlashSmsAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlashSmsAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetCode(v string) *DeleteFlashSmsAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetData(v string) *DeleteFlashSmsAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetHttpStatusCode(v int32) *DeleteFlashSmsAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetMessage(v string) *DeleteFlashSmsAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetParams(v []*string) *DeleteFlashSmsAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetRequestId(v string) *DeleteFlashSmsAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) SetSuccess(v bool) *DeleteFlashSmsAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
