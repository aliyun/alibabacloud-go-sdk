// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryAutoExtractionSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryAutoExtractionSwitchResponseBody
	GetCode() *string
	SetEnabled(v bool) *GetMemoryAutoExtractionSwitchResponseBody
	GetEnabled() *bool
	SetMessage(v string) *GetMemoryAutoExtractionSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMemoryAutoExtractionSwitchResponseBody
	GetRequestId() *string
}

type GetMemoryAutoExtractionSwitchResponseBody struct {
	// 业务状态码；成功为 200，失败时为对应错误码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前记忆范围的自动记忆拾取开关状态；true 表示已开启，false 表示已关闭
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 错误描述；成功时为空字符串，失败时返回具体原因
	//
	// example:
	//
	// 当前用户不是群组 cg_12345 的活跃成员
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID，用于定位本次调用
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemoryAutoExtractionSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryAutoExtractionSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetCode(v string) *GetMemoryAutoExtractionSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetEnabled(v bool) *GetMemoryAutoExtractionSwitchResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetMessage(v string) *GetMemoryAutoExtractionSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetRequestId(v string) *GetMemoryAutoExtractionSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
