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
	// The abstract model name (model tier). If not specified, the current model of the session is not modified.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The new session title.
	//
	// example:
	//
	// Sample title
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
