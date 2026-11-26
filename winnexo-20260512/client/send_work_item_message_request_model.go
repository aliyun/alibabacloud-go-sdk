// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendWorkItemMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SendWorkItemMessageRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *SendWorkItemMessageRequest
	GetCollaborationGroupId() *string
	SetContent(v string) *SendWorkItemMessageRequest
	GetContent() *string
	SetConversationId(v string) *SendWorkItemMessageRequest
	GetConversationId() *string
	SetOpsUserCode(v string) *SendWorkItemMessageRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SendWorkItemMessageRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SendWorkItemMessageRequest
	GetTenantId() *string
	SetWorkItemId(v string) *SendWorkItemMessageRequest
	GetWorkItemId() *string
}

type SendWorkItemMessageRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration space.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The body of the user message. The value cannot be empty and can be up to 32,768 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Compile a list of next steps based on the previous discussion
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The conversation ID returned by the API operation for creating a free exploration. You must specify at least one of this parameter and workItemId. If both are specified, they must match.
	//
	// example:
	//
	// 01234567-89ab-4def-8123-456789abcdef
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the free exploration work item. This parameter is compatible with the legacy-topic:<conversationId> format. You must specify at least one of workItemId and conversationId. If both are specified, they must point to the same session.
	//
	// example:
	//
	// wi_0123456789abcdef0123456789abcdef
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s SendWorkItemMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendWorkItemMessageRequest) GoString() string {
	return s.String()
}

func (s *SendWorkItemMessageRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SendWorkItemMessageRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SendWorkItemMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendWorkItemMessageRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *SendWorkItemMessageRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SendWorkItemMessageRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SendWorkItemMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendWorkItemMessageRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *SendWorkItemMessageRequest) SetAppToken(v string) *SendWorkItemMessageRequest {
	s.AppToken = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetCollaborationGroupId(v string) *SendWorkItemMessageRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetContent(v string) *SendWorkItemMessageRequest {
	s.Content = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetConversationId(v string) *SendWorkItemMessageRequest {
	s.ConversationId = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetOpsUserCode(v string) *SendWorkItemMessageRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetOpsUserId(v string) *SendWorkItemMessageRequest {
	s.OpsUserId = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetTenantId(v string) *SendWorkItemMessageRequest {
	s.TenantId = &v
	return s
}

func (s *SendWorkItemMessageRequest) SetWorkItemId(v string) *SendWorkItemMessageRequest {
	s.WorkItemId = &v
	return s
}

func (s *SendWorkItemMessageRequest) Validate() error {
	return dara.Validate(s)
}
