// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceManualMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceManualMemoryResponseBody
	GetCode() *string
	SetContent(v string) *ReplaceManualMemoryResponseBody
	GetContent() *string
	SetMessage(v string) *ReplaceManualMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReplaceManualMemoryResponseBody
	GetRequestId() *string
}

type ReplaceManualMemoryResponseBody struct {
	// 业务状态码；成功为 200，失败时为对应错误码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 保存后的手动记忆完整内容；清空成功时为空字符串
	//
	// example:
	//
	// 喜欢简洁的回答。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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

func (s ReplaceManualMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceManualMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceManualMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceManualMemoryResponseBody) GetContent() *string {
	return s.Content
}

func (s *ReplaceManualMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceManualMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceManualMemoryResponseBody) SetCode(v string) *ReplaceManualMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) SetContent(v string) *ReplaceManualMemoryResponseBody {
	s.Content = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) SetMessage(v string) *ReplaceManualMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) SetRequestId(v string) *ReplaceManualMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceManualMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
