// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMassMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *SendChatappMassMessageShrinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *SendChatappMassMessageShrinkRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *SendChatappMassMessageShrinkRequest
	GetCustWabaId() *string
	SetFallBackContent(v string) *SendChatappMassMessageShrinkRequest
	GetFallBackContent() *string
	SetFallBackDuration(v int32) *SendChatappMassMessageShrinkRequest
	GetFallBackDuration() *int32
	SetFallBackId(v string) *SendChatappMassMessageShrinkRequest
	GetFallBackId() *string
	SetFallBackRule(v string) *SendChatappMassMessageShrinkRequest
	GetFallBackRule() *string
	SetFrom(v string) *SendChatappMassMessageShrinkRequest
	GetFrom() *string
	SetIsvCode(v string) *SendChatappMassMessageShrinkRequest
	GetIsvCode() *string
	SetLabel(v string) *SendChatappMassMessageShrinkRequest
	GetLabel() *string
	SetLanguage(v string) *SendChatappMassMessageShrinkRequest
	GetLanguage() *string
	SetOwnerId(v int64) *SendChatappMassMessageShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SendChatappMassMessageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendChatappMassMessageShrinkRequest
	GetResourceOwnerId() *int64
	SetSenderListShrink(v string) *SendChatappMassMessageShrinkRequest
	GetSenderListShrink() *string
	SetTag(v string) *SendChatappMassMessageShrinkRequest
	GetTag() *string
	SetTaskId(v string) *SendChatappMassMessageShrinkRequest
	GetTaskId() *string
	SetTemplateCode(v string) *SendChatappMassMessageShrinkRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *SendChatappMassMessageShrinkRequest
	GetTemplateName() *string
	SetTtl(v int64) *SendChatappMassMessageShrinkRequest
	GetTtl() *int64
}

type SendChatappMassMessageShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackContent  *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	FallBackDuration *int32  `json:"FallBackDuration,omitempty" xml:"FallBackDuration,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FallBackRule *string `json:"FallBackRule,omitempty" xml:"FallBackRule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 示例值示例值
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// example:
	//
	// 示例值示例值
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SenderListShrink     *string `json:"SenderList,omitempty" xml:"SenderList,omitempty"`
	// example:
	//
	// 示例值示例值
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 示例值示例值
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 46
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SendChatappMassMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendChatappMassMessageShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SendChatappMassMessageShrinkRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackContent() *string {
	return s.FallBackContent
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackDuration() *int32 {
	return s.FallBackDuration
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackId() *string {
	return s.FallBackId
}

func (s *SendChatappMassMessageShrinkRequest) GetFallBackRule() *string {
	return s.FallBackRule
}

func (s *SendChatappMassMessageShrinkRequest) GetFrom() *string {
	return s.From
}

func (s *SendChatappMassMessageShrinkRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *SendChatappMassMessageShrinkRequest) GetLabel() *string {
	return s.Label
}

func (s *SendChatappMassMessageShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *SendChatappMassMessageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendChatappMassMessageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendChatappMassMessageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendChatappMassMessageShrinkRequest) GetSenderListShrink() *string {
	return s.SenderListShrink
}

func (s *SendChatappMassMessageShrinkRequest) GetTag() *string {
	return s.Tag
}

func (s *SendChatappMassMessageShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SendChatappMassMessageShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendChatappMassMessageShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SendChatappMassMessageShrinkRequest) GetTtl() *int64 {
	return s.Ttl
}

func (s *SendChatappMassMessageShrinkRequest) SetChannelType(v string) *SendChatappMassMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetCustSpaceId(v string) *SendChatappMassMessageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetCustWabaId(v string) *SendChatappMassMessageShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackContent(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackDuration(v int32) *SendChatappMassMessageShrinkRequest {
	s.FallBackDuration = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackId(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFallBackRule(v string) *SendChatappMassMessageShrinkRequest {
	s.FallBackRule = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetFrom(v string) *SendChatappMassMessageShrinkRequest {
	s.From = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetIsvCode(v string) *SendChatappMassMessageShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetLabel(v string) *SendChatappMassMessageShrinkRequest {
	s.Label = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetLanguage(v string) *SendChatappMassMessageShrinkRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetOwnerId(v int64) *SendChatappMassMessageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetResourceOwnerAccount(v string) *SendChatappMassMessageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetResourceOwnerId(v int64) *SendChatappMassMessageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetSenderListShrink(v string) *SendChatappMassMessageShrinkRequest {
	s.SenderListShrink = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTag(v string) *SendChatappMassMessageShrinkRequest {
	s.Tag = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTaskId(v string) *SendChatappMassMessageShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMassMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTemplateName(v string) *SendChatappMassMessageShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) SetTtl(v int64) *SendChatappMassMessageShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *SendChatappMassMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
