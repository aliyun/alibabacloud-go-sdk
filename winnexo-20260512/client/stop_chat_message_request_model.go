// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *StopChatMessageRequest
	GetSessionId() *string
	SetTenantId(v string) *StopChatMessageRequest
	GetTenantId() *string
}

type StopChatMessageRequest struct {
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StopChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s StopChatMessageRequest) GoString() string {
	return s.String()
}

func (s *StopChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *StopChatMessageRequest) SetSessionId(v string) *StopChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *StopChatMessageRequest) SetTenantId(v string) *StopChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *StopChatMessageRequest) Validate() error {
	return dara.Validate(s)
}
