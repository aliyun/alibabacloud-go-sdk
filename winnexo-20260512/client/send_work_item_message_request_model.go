// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendWorkItemMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *SendWorkItemMessageRequest
	GetCollaborationGroupId() *string
	SetContent(v string) *SendWorkItemMessageRequest
	GetContent() *string
	SetConversationId(v string) *SendWorkItemMessageRequest
	GetConversationId() *string
	SetTenantId(v string) *SendWorkItemMessageRequest
	GetTenantId() *string
	SetWorkItemId(v string) *SendWorkItemMessageRequest
	GetWorkItemId() *string
}

type SendWorkItemMessageRequest struct {
	// 目标会话所属的协作空间ID，必填，长度1至64，不能为空白。调用者须具有目标话题的查看和发送权限。
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 用户发送的文本正文，必填，长度1至32768，不能全为空白；保留原有文本格式。
	//
	// This parameter is required.
	//
	// example:
	//
	// 请根据刚才的讨论整理下一步行动清单。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 已有自由探索话题的会话ID，长度1至128，可使用创建自由探索接口返回的conversation_id。与workItemId至少传一个，同时传入时须指向同一会话。
	//
	// example:
	//
	// 01234567-89ab-4def-8123-456789abcdef
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 业务租户ID，可选。必须属于当前调用身份的可访问租户；无权访问时拒绝请求。未传时按现有鉴权规则解析租户，多租户用户无法唯一确定时需显式指定。
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 自由探索工作项ID，长度1至256；兼容legacy-topic:<conversationId>形式。与conversationId至少传一个，同时传入时须指向同一会话。普通工作项必须为FREE_EXPLORATION类型。
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

func (s *SendWorkItemMessageRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SendWorkItemMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendWorkItemMessageRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *SendWorkItemMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendWorkItemMessageRequest) GetWorkItemId() *string {
	return s.WorkItemId
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
