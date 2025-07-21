// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSendMailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SingleSendMailRequest
	GetAccountName() *string
	SetAddressType(v int32) *SingleSendMailRequest
	GetAddressType() *int32
	SetAttachments(v []*SingleSendMailRequestAttachments) *SingleSendMailRequest
	GetAttachments() []*SingleSendMailRequestAttachments
	SetClickTrace(v string) *SingleSendMailRequest
	GetClickTrace() *string
	SetFromAlias(v string) *SingleSendMailRequest
	GetFromAlias() *string
	SetHeaders(v string) *SingleSendMailRequest
	GetHeaders() *string
	SetHtmlBody(v string) *SingleSendMailRequest
	GetHtmlBody() *string
	SetIpPoolId(v string) *SingleSendMailRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *SingleSendMailRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *SingleSendMailRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *SingleSendMailRequest
	GetReplyAddressAlias() *string
	SetReplyToAddress(v bool) *SingleSendMailRequest
	GetReplyToAddress() *bool
	SetResourceOwnerAccount(v string) *SingleSendMailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleSendMailRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *SingleSendMailRequest
	GetSubject() *string
	SetTagName(v string) *SingleSendMailRequest
	GetTagName() *string
	SetTextBody(v string) *SingleSendMailRequest
	GetTextBody() *string
	SetToAddress(v string) *SingleSendMailRequest
	GetToAddress() *string
	SetUnSubscribeFilterLevel(v string) *SingleSendMailRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *SingleSendMailRequest
	GetUnSubscribeLinkType() *string
}

type SingleSendMailRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AddressType       *int32                              `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Attachments       []*SingleSendMailRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	ClickTrace        *string                             `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	FromAlias         *string                             `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	Headers           *string                             `json:"Headers,omitempty" xml:"Headers,omitempty"`
	HtmlBody          *string                             `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	IpPoolId          *string                             `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId           *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReplyAddress      *string                             `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias *string                             `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// This parameter is required.
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Subject  *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// This parameter is required.
	ToAddress              *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	UnSubscribeLinkType    *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s SingleSendMailRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SingleSendMailRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *SingleSendMailRequest) GetAttachments() []*SingleSendMailRequestAttachments {
	return s.Attachments
}

func (s *SingleSendMailRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailRequest) GetFromAlias() *string {
	return s.FromAlias
}

func (s *SingleSendMailRequest) GetHeaders() *string {
	return s.Headers
}

func (s *SingleSendMailRequest) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *SingleSendMailRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SingleSendMailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleSendMailRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *SingleSendMailRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *SingleSendMailRequest) GetReplyToAddress() *bool {
	return s.ReplyToAddress
}

func (s *SingleSendMailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleSendMailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleSendMailRequest) GetSubject() *string {
	return s.Subject
}

func (s *SingleSendMailRequest) GetTagName() *string {
	return s.TagName
}

func (s *SingleSendMailRequest) GetTextBody() *string {
	return s.TextBody
}

func (s *SingleSendMailRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SingleSendMailRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *SingleSendMailRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *SingleSendMailRequest) SetAccountName(v string) *SingleSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailRequest) SetAddressType(v int32) *SingleSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailRequest) SetAttachments(v []*SingleSendMailRequestAttachments) *SingleSendMailRequest {
	s.Attachments = v
	return s
}

func (s *SingleSendMailRequest) SetClickTrace(v string) *SingleSendMailRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailRequest) SetFromAlias(v string) *SingleSendMailRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailRequest) SetHeaders(v string) *SingleSendMailRequest {
	s.Headers = &v
	return s
}

func (s *SingleSendMailRequest) SetHtmlBody(v string) *SingleSendMailRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailRequest) SetIpPoolId(v string) *SingleSendMailRequest {
	s.IpPoolId = &v
	return s
}

func (s *SingleSendMailRequest) SetOwnerId(v int64) *SingleSendMailRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyAddress(v string) *SingleSendMailRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyAddressAlias(v string) *SingleSendMailRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyToAddress(v bool) *SingleSendMailRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetResourceOwnerAccount(v string) *SingleSendMailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailRequest) SetResourceOwnerId(v int64) *SingleSendMailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailRequest) SetSubject(v string) *SingleSendMailRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailRequest) SetTagName(v string) *SingleSendMailRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailRequest) SetTextBody(v string) *SingleSendMailRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailRequest) SetToAddress(v string) *SingleSendMailRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetUnSubscribeFilterLevel(v string) *SingleSendMailRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *SingleSendMailRequest) SetUnSubscribeLinkType(v string) *SingleSendMailRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *SingleSendMailRequest) Validate() error {
	return dara.Validate(s)
}

type SingleSendMailRequestAttachments struct {
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	AttachmentUrl  *string `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
}

func (s SingleSendMailRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailRequestAttachments) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequestAttachments) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *SingleSendMailRequestAttachments) GetAttachmentUrl() *string {
	return s.AttachmentUrl
}

func (s *SingleSendMailRequestAttachments) SetAttachmentName(v string) *SingleSendMailRequestAttachments {
	s.AttachmentName = &v
	return s
}

func (s *SingleSendMailRequestAttachments) SetAttachmentUrl(v string) *SingleSendMailRequestAttachments {
	s.AttachmentUrl = &v
	return s
}

func (s *SingleSendMailRequestAttachments) Validate() error {
	return dara.Validate(s)
}
