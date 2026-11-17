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
	// 源会话 ID；首次调用建立转换，后续使用同一 ID 查询或继续原转换
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 租户 ID，公共参数，缺省时使用调用方默认租户；仍需通过源会话归属及权限校验
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
