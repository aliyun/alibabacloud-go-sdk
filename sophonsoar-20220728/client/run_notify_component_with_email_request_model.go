// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *RunNotifyComponentWithEmailRequest
	GetActionName() *string
	SetAssetId(v string) *RunNotifyComponentWithEmailRequest
	GetAssetId() *string
	SetComponentName(v string) *RunNotifyComponentWithEmailRequest
	GetComponentName() *string
	SetContent(v string) *RunNotifyComponentWithEmailRequest
	GetContent() *string
	SetLang(v string) *RunNotifyComponentWithEmailRequest
	GetLang() *string
	SetNodeName(v string) *RunNotifyComponentWithEmailRequest
	GetNodeName() *string
	SetPlaybookUuid(v string) *RunNotifyComponentWithEmailRequest
	GetPlaybookUuid() *string
	SetReceivers(v []*string) *RunNotifyComponentWithEmailRequest
	GetReceivers() []*string
	SetRoleFor(v int64) *RunNotifyComponentWithEmailRequest
	GetRoleFor() *int64
	SetRoleType(v string) *RunNotifyComponentWithEmailRequest
	GetRoleType() *string
	SetSubject(v string) *RunNotifyComponentWithEmailRequest
	GetSubject() *string
}

type RunNotifyComponentWithEmailRequest struct {
	// The action name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// notifyByCustom
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The resource instance ID of the email sender.
	//
	// >  You can call the [DescribeComponentAssets](~~DescribeComponentAssets~~) operation to query the ID.
	//
	// example:
	//
	// 10
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The name of component in the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// NotifyMessage
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The body of the email.
	//
	// This parameter is required.
	//
	// example:
	//
	// email content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The language of the content within the request and the response. Valid value:
	//
	// 	- **zh*	- (default): Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the node in the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// notify_message_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The email addresses.
	//
	// This parameter is required.
	Receivers []*string `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 137602xxx718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- 0: the view of the current Alibaba Cloud account.
	//
	// 	- 1: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The subject of the email.
	//
	// This parameter is required.
	//
	// example:
	//
	// title
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s RunNotifyComponentWithEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithEmailRequest) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailRequest) GetActionName() *string {
	return s.ActionName
}

func (s *RunNotifyComponentWithEmailRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *RunNotifyComponentWithEmailRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *RunNotifyComponentWithEmailRequest) GetContent() *string {
	return s.Content
}

func (s *RunNotifyComponentWithEmailRequest) GetLang() *string {
	return s.Lang
}

func (s *RunNotifyComponentWithEmailRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *RunNotifyComponentWithEmailRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *RunNotifyComponentWithEmailRequest) GetReceivers() []*string {
	return s.Receivers
}

func (s *RunNotifyComponentWithEmailRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *RunNotifyComponentWithEmailRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *RunNotifyComponentWithEmailRequest) GetSubject() *string {
	return s.Subject
}

func (s *RunNotifyComponentWithEmailRequest) SetActionName(v string) *RunNotifyComponentWithEmailRequest {
	s.ActionName = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetAssetId(v string) *RunNotifyComponentWithEmailRequest {
	s.AssetId = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetComponentName(v string) *RunNotifyComponentWithEmailRequest {
	s.ComponentName = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetContent(v string) *RunNotifyComponentWithEmailRequest {
	s.Content = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetLang(v string) *RunNotifyComponentWithEmailRequest {
	s.Lang = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetNodeName(v string) *RunNotifyComponentWithEmailRequest {
	s.NodeName = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetPlaybookUuid(v string) *RunNotifyComponentWithEmailRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetReceivers(v []*string) *RunNotifyComponentWithEmailRequest {
	s.Receivers = v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetRoleFor(v int64) *RunNotifyComponentWithEmailRequest {
	s.RoleFor = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetRoleType(v string) *RunNotifyComponentWithEmailRequest {
	s.RoleType = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetSubject(v string) *RunNotifyComponentWithEmailRequest {
	s.Subject = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) Validate() error {
	return dara.Validate(s)
}
