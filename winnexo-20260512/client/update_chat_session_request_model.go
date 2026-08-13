// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModel(v string) *UpdateChatSessionRequest
	GetModel() *string
	SetSessionId(v string) *UpdateChatSessionRequest
	GetSessionId() *string
	SetTenantId(v string) *UpdateChatSessionRequest
	GetTenantId() *string
	SetTitle(v string) *UpdateChatSessionRequest
	GetTitle() *string
}

type UpdateChatSessionRequest struct {
	// 抽象模型名（模型档位）；不传则不修改会话当前模型
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 会话 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 新的会话标题
	//
	// example:
	//
	// 示例标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatSessionRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateChatSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateChatSessionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateChatSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateChatSessionRequest) SetModel(v string) *UpdateChatSessionRequest {
	s.Model = &v
	return s
}

func (s *UpdateChatSessionRequest) SetSessionId(v string) *UpdateChatSessionRequest {
	s.SessionId = &v
	return s
}

func (s *UpdateChatSessionRequest) SetTenantId(v string) *UpdateChatSessionRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateChatSessionRequest) SetTitle(v string) *UpdateChatSessionRequest {
	s.Title = &v
	return s
}

func (s *UpdateChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
