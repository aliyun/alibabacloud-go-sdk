// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSendMailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SingleSendMailShrinkRequest
	GetAccountName() *string
	SetAddressType(v int32) *SingleSendMailShrinkRequest
	GetAddressType() *int32
	SetAttachments(v []*SingleSendMailShrinkRequestAttachments) *SingleSendMailShrinkRequest
	GetAttachments() []*SingleSendMailShrinkRequestAttachments
	SetClickTrace(v string) *SingleSendMailShrinkRequest
	GetClickTrace() *string
	SetFromAlias(v string) *SingleSendMailShrinkRequest
	GetFromAlias() *string
	SetHeaders(v string) *SingleSendMailShrinkRequest
	GetHeaders() *string
	SetHtmlBody(v string) *SingleSendMailShrinkRequest
	GetHtmlBody() *string
	SetIpPoolId(v string) *SingleSendMailShrinkRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *SingleSendMailShrinkRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *SingleSendMailShrinkRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *SingleSendMailShrinkRequest
	GetReplyAddressAlias() *string
	SetReplyToAddress(v bool) *SingleSendMailShrinkRequest
	GetReplyToAddress() *bool
	SetResourceOwnerAccount(v string) *SingleSendMailShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleSendMailShrinkRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *SingleSendMailShrinkRequest
	GetSubject() *string
	SetTagName(v string) *SingleSendMailShrinkRequest
	GetTagName() *string
	SetTemplateShrink(v string) *SingleSendMailShrinkRequest
	GetTemplateShrink() *string
	SetTextBody(v string) *SingleSendMailShrinkRequest
	GetTextBody() *string
	SetToAddress(v string) *SingleSendMailShrinkRequest
	GetToAddress() *string
	SetUnSubscribeFilterLevel(v string) *SingleSendMailShrinkRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *SingleSendMailShrinkRequest
	GetUnSubscribeLinkType() *string
}

type SingleSendMailShrinkRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AddressType       *int32                                    `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Attachments       []*SingleSendMailShrinkRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	ClickTrace        *string                                   `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	FromAlias         *string                                   `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	Headers           *string                                   `json:"Headers,omitempty" xml:"Headers,omitempty"`
	HtmlBody          *string                                   `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	IpPoolId          *string                                   `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId           *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReplyAddress      *string                                   `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias *string                                   `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// This parameter is required.
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Subject        *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	TextBody       *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// This parameter is required.
	ToAddress              *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	UnSubscribeLinkType    *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s SingleSendMailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailShrinkRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SingleSendMailShrinkRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *SingleSendMailShrinkRequest) GetAttachments() []*SingleSendMailShrinkRequestAttachments {
	return s.Attachments
}

func (s *SingleSendMailShrinkRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailShrinkRequest) GetFromAlias() *string {
	return s.FromAlias
}

func (s *SingleSendMailShrinkRequest) GetHeaders() *string {
	return s.Headers
}

func (s *SingleSendMailShrinkRequest) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *SingleSendMailShrinkRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SingleSendMailShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleSendMailShrinkRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *SingleSendMailShrinkRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *SingleSendMailShrinkRequest) GetReplyToAddress() *bool {
	return s.ReplyToAddress
}

func (s *SingleSendMailShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleSendMailShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleSendMailShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *SingleSendMailShrinkRequest) GetTagName() *string {
	return s.TagName
}

func (s *SingleSendMailShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *SingleSendMailShrinkRequest) GetTextBody() *string {
	return s.TextBody
}

func (s *SingleSendMailShrinkRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SingleSendMailShrinkRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *SingleSendMailShrinkRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *SingleSendMailShrinkRequest) SetAccountName(v string) *SingleSendMailShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetAddressType(v int32) *SingleSendMailShrinkRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetAttachments(v []*SingleSendMailShrinkRequestAttachments) *SingleSendMailShrinkRequest {
	s.Attachments = v
	return s
}

func (s *SingleSendMailShrinkRequest) SetClickTrace(v string) *SingleSendMailShrinkRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetFromAlias(v string) *SingleSendMailShrinkRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetHeaders(v string) *SingleSendMailShrinkRequest {
	s.Headers = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetHtmlBody(v string) *SingleSendMailShrinkRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetIpPoolId(v string) *SingleSendMailShrinkRequest {
	s.IpPoolId = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetOwnerId(v int64) *SingleSendMailShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetReplyAddress(v string) *SingleSendMailShrinkRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetReplyAddressAlias(v string) *SingleSendMailShrinkRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetReplyToAddress(v bool) *SingleSendMailShrinkRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetResourceOwnerAccount(v string) *SingleSendMailShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetResourceOwnerId(v int64) *SingleSendMailShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetSubject(v string) *SingleSendMailShrinkRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetTagName(v string) *SingleSendMailShrinkRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetTemplateShrink(v string) *SingleSendMailShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetTextBody(v string) *SingleSendMailShrinkRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetToAddress(v string) *SingleSendMailShrinkRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetUnSubscribeFilterLevel(v string) *SingleSendMailShrinkRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetUnSubscribeLinkType(v string) *SingleSendMailShrinkRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *SingleSendMailShrinkRequest) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SingleSendMailShrinkRequestAttachments struct {
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	AttachmentUrl  *string `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
}

func (s SingleSendMailShrinkRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailShrinkRequestAttachments) GoString() string {
	return s.String()
}

func (s *SingleSendMailShrinkRequestAttachments) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *SingleSendMailShrinkRequestAttachments) GetAttachmentUrl() *string {
	return s.AttachmentUrl
}

func (s *SingleSendMailShrinkRequestAttachments) SetAttachmentName(v string) *SingleSendMailShrinkRequestAttachments {
	s.AttachmentName = &v
	return s
}

func (s *SingleSendMailShrinkRequestAttachments) SetAttachmentUrl(v string) *SingleSendMailShrinkRequestAttachments {
	s.AttachmentUrl = &v
	return s
}

func (s *SingleSendMailShrinkRequestAttachments) Validate() error {
	return dara.Validate(s)
}
