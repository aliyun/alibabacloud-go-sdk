// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *RunNotifyComponentWithWebhookRequest
	GetActionName() *string
	SetAssetId(v string) *RunNotifyComponentWithWebhookRequest
	GetAssetId() *string
	SetComponentName(v string) *RunNotifyComponentWithWebhookRequest
	GetComponentName() *string
	SetContent(v string) *RunNotifyComponentWithWebhookRequest
	GetContent() *string
	SetLang(v string) *RunNotifyComponentWithWebhookRequest
	GetLang() *string
	SetMsgType(v string) *RunNotifyComponentWithWebhookRequest
	GetMsgType() *string
	SetNodeName(v string) *RunNotifyComponentWithWebhookRequest
	GetNodeName() *string
	SetPlaybookUuid(v string) *RunNotifyComponentWithWebhookRequest
	GetPlaybookUuid() *string
	SetRoleFor(v int64) *RunNotifyComponentWithWebhookRequest
	GetRoleFor() *int64
	SetRoleType(v string) *RunNotifyComponentWithWebhookRequest
	GetRoleType() *string
	SetSecret(v string) *RunNotifyComponentWithWebhookRequest
	GetSecret() *string
	SetWebhook(v string) *RunNotifyComponentWithWebhookRequest
	GetWebhook() *string
}

type RunNotifyComponentWithWebhookRequest struct {
	// The name of the action in the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// notifyByCustom
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The ID of the resource. This parameter is deprecated.
	//
	// example:
	//
	// 1
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The name of the component in the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// NotifyMessage
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The message body sent by the DingTalk group chatbot webhook.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "at": {
	//
	//         "atMobiles":[
	//
	//             "180xxxxxx"
	//
	//         ],
	//
	//         "atUserIds":[
	//
	//             "user123"
	//
	//         ],
	//
	//         "isAtAll": false
	//
	//     },
	//
	//     "text": {
	//
	//         "content":"1234"
	//
	//     },
	//
	//     "msgtype":"text"
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the webhook message. Valid values:
	//
	// 	- text.
	//
	// 	- markdown.
	//
	// 	- actionCard.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	// The name of the node in the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// notify_message_node
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 94bc318c-****-4cba-****-801ccb0d739f
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 126339xxxx805497
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- 0 (default): the view of the current Alibaba Cloud account.
	//
	// 	- 1: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The message key of the DingTalk chatbot webhook. This parameter is deprecated.
	//
	// example:
	//
	// SECc1*****e157b32b380f********bb8c70e1a67a22072
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The IDs of chatbots that are configured in the message center. Only DingTalk chatbots are supported.
	//
	// >  You can call the [ListEncryptWebhooks](~~ListEncryptWebhooks~~) operation to query the chatbot IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"10651\\"]
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s RunNotifyComponentWithWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithWebhookRequest) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookRequest) GetActionName() *string {
	return s.ActionName
}

func (s *RunNotifyComponentWithWebhookRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *RunNotifyComponentWithWebhookRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *RunNotifyComponentWithWebhookRequest) GetContent() *string {
	return s.Content
}

func (s *RunNotifyComponentWithWebhookRequest) GetLang() *string {
	return s.Lang
}

func (s *RunNotifyComponentWithWebhookRequest) GetMsgType() *string {
	return s.MsgType
}

func (s *RunNotifyComponentWithWebhookRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *RunNotifyComponentWithWebhookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *RunNotifyComponentWithWebhookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *RunNotifyComponentWithWebhookRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *RunNotifyComponentWithWebhookRequest) GetSecret() *string {
	return s.Secret
}

func (s *RunNotifyComponentWithWebhookRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *RunNotifyComponentWithWebhookRequest) SetActionName(v string) *RunNotifyComponentWithWebhookRequest {
	s.ActionName = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetAssetId(v string) *RunNotifyComponentWithWebhookRequest {
	s.AssetId = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetComponentName(v string) *RunNotifyComponentWithWebhookRequest {
	s.ComponentName = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetContent(v string) *RunNotifyComponentWithWebhookRequest {
	s.Content = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetLang(v string) *RunNotifyComponentWithWebhookRequest {
	s.Lang = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetMsgType(v string) *RunNotifyComponentWithWebhookRequest {
	s.MsgType = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetNodeName(v string) *RunNotifyComponentWithWebhookRequest {
	s.NodeName = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetPlaybookUuid(v string) *RunNotifyComponentWithWebhookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetRoleFor(v int64) *RunNotifyComponentWithWebhookRequest {
	s.RoleFor = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetRoleType(v string) *RunNotifyComponentWithWebhookRequest {
	s.RoleType = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetSecret(v string) *RunNotifyComponentWithWebhookRequest {
	s.Secret = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetWebhook(v string) *RunNotifyComponentWithWebhookRequest {
	s.Webhook = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) Validate() error {
	return dara.Validate(s)
}
