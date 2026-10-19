// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryAutoExtractionSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClearedCount(v int64) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetClearedCount() *int64
	SetCode(v string) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetCode() *string
	SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetEnabled() *bool
	SetMessage(v string) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetRequestId() *string
}

type UpdateMemoryAutoExtractionSwitchResponseBody struct {
	// 本次同步清空的自动记忆文件数量；未清空或无可清空记忆时为 0
	//
	// example:
	//
	// 3
	ClearedCount *int64 `json:"clearedCount,omitempty" xml:"clearedCount,omitempty"`
	// 业务状态码；成功为 200，失败时为对应错误码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 更新后的自动记忆拾取开关状态；true 为开启，false 为关闭
	//
	// example:
	//
	// false
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

func (s UpdateMemoryAutoExtractionSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryAutoExtractionSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetClearedCount() *int64 {
	return s.ClearedCount
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetClearedCount(v int64) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.ClearedCount = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetCode(v string) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetMessage(v string) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetRequestId(v string) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
