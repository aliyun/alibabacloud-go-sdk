// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInterpreterSessionOut interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterId(v string) *CodeInterpreterSessionOut
	GetCodeInterpreterId() *string
	SetCodeInterpreterName(v string) *CodeInterpreterSessionOut
	GetCodeInterpreterName() *string
	SetCreatedAt(v string) *CodeInterpreterSessionOut
	GetCreatedAt() *string
	SetLastUpdatedAt(v string) *CodeInterpreterSessionOut
	GetLastUpdatedAt() *string
	SetSessionId(v string) *CodeInterpreterSessionOut
	GetSessionId() *string
	SetSessionIdleTimeoutSeconds(v int32) *CodeInterpreterSessionOut
	GetSessionIdleTimeoutSeconds() *int32
	SetStatus(v string) *CodeInterpreterSessionOut
	GetStatus() *string
}

type CodeInterpreterSessionOut struct {
	// 关联的代码解释器的唯一标识符
	//
	// This parameter is required.
	CodeInterpreterId *string `json:"codeInterpreterId,omitempty" xml:"codeInterpreterId,omitempty"`
	// 代码解释器会话的名称
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	// 代码解释器会话的创建时间，采用ISO 8601格式
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 代码解释器会话的最后更新时间，采用ISO 8601格式
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// 代码解释器会话的唯一标识符
	//
	// This parameter is required.
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 代码解释器会话的空闲超时时间，单位为秒。实例没有会话请求后处于空闲状态，空闲态为闲置计费模式，超过此超时时间后会话自动过期，不可继续使用
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// 代码解释器会话的当前状态
	//
	// This parameter is required.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CodeInterpreterSessionOut) String() string {
	return dara.Prettify(s)
}

func (s CodeInterpreterSessionOut) GoString() string {
	return s.String()
}

func (s *CodeInterpreterSessionOut) GetCodeInterpreterId() *string {
	return s.CodeInterpreterId
}

func (s *CodeInterpreterSessionOut) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *CodeInterpreterSessionOut) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CodeInterpreterSessionOut) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *CodeInterpreterSessionOut) GetSessionId() *string {
	return s.SessionId
}

func (s *CodeInterpreterSessionOut) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CodeInterpreterSessionOut) GetStatus() *string {
	return s.Status
}

func (s *CodeInterpreterSessionOut) SetCodeInterpreterId(v string) *CodeInterpreterSessionOut {
	s.CodeInterpreterId = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetCodeInterpreterName(v string) *CodeInterpreterSessionOut {
	s.CodeInterpreterName = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetCreatedAt(v string) *CodeInterpreterSessionOut {
	s.CreatedAt = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetLastUpdatedAt(v string) *CodeInterpreterSessionOut {
	s.LastUpdatedAt = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetSessionId(v string) *CodeInterpreterSessionOut {
	s.SessionId = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetSessionIdleTimeoutSeconds(v int32) *CodeInterpreterSessionOut {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CodeInterpreterSessionOut) SetStatus(v string) *CodeInterpreterSessionOut {
	s.Status = &v
	return s
}

func (s *CodeInterpreterSessionOut) Validate() error {
	return dara.Validate(s)
}
