// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertDirectConversationToWorkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *ConvertDirectConversationToWorkRequest
	GetSessionId() *string
	SetTenantId(v string) *ConvertDirectConversationToWorkRequest
	GetTenantId() *string
}

type ConvertDirectConversationToWorkRequest struct {
	// The source direct session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ConvertDirectConversationToWorkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertDirectConversationToWorkRequest) GoString() string {
	return s.String()
}

func (s *ConvertDirectConversationToWorkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ConvertDirectConversationToWorkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ConvertDirectConversationToWorkRequest) SetSessionId(v string) *ConvertDirectConversationToWorkRequest {
	s.SessionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) SetTenantId(v string) *ConvertDirectConversationToWorkRequest {
	s.TenantId = &v
	return s
}

func (s *ConvertDirectConversationToWorkRequest) Validate() error {
	return dara.Validate(s)
}
