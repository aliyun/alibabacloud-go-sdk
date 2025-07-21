// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSingleSendMailAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SingleSendMailAdvanceRequest
	GetAccountName() *string
	SetAddressType(v int32) *SingleSendMailAdvanceRequest
	GetAddressType() *int32
	SetAttachments(v []*SingleSendMailAdvanceRequestAttachments) *SingleSendMailAdvanceRequest
	GetAttachments() []*SingleSendMailAdvanceRequestAttachments
	SetClickTrace(v string) *SingleSendMailAdvanceRequest
	GetClickTrace() *string
	SetFromAlias(v string) *SingleSendMailAdvanceRequest
	GetFromAlias() *string
	SetHeaders(v string) *SingleSendMailAdvanceRequest
	GetHeaders() *string
	SetHtmlBody(v string) *SingleSendMailAdvanceRequest
	GetHtmlBody() *string
	SetIpPoolId(v string) *SingleSendMailAdvanceRequest
	GetIpPoolId() *string
	SetOwnerId(v int64) *SingleSendMailAdvanceRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *SingleSendMailAdvanceRequest
	GetReplyAddress() *string
	SetReplyAddressAlias(v string) *SingleSendMailAdvanceRequest
	GetReplyAddressAlias() *string
	SetReplyToAddress(v bool) *SingleSendMailAdvanceRequest
	GetReplyToAddress() *bool
	SetResourceOwnerAccount(v string) *SingleSendMailAdvanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SingleSendMailAdvanceRequest
	GetResourceOwnerId() *int64
	SetSubject(v string) *SingleSendMailAdvanceRequest
	GetSubject() *string
	SetTagName(v string) *SingleSendMailAdvanceRequest
	GetTagName() *string
	SetTextBody(v string) *SingleSendMailAdvanceRequest
	GetTextBody() *string
	SetToAddress(v string) *SingleSendMailAdvanceRequest
	GetToAddress() *string
	SetUnSubscribeFilterLevel(v string) *SingleSendMailAdvanceRequest
	GetUnSubscribeFilterLevel() *string
	SetUnSubscribeLinkType(v string) *SingleSendMailAdvanceRequest
	GetUnSubscribeLinkType() *string
}

type SingleSendMailAdvanceRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AddressType       *int32                                     `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Attachments       []*SingleSendMailAdvanceRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	ClickTrace        *string                                    `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	FromAlias         *string                                    `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	Headers           *string                                    `json:"Headers,omitempty" xml:"Headers,omitempty"`
	HtmlBody          *string                                    `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	IpPoolId          *string                                    `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId           *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReplyAddress      *string                                    `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias *string                                    `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
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

func (s SingleSendMailAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailAdvanceRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SingleSendMailAdvanceRequest) GetAddressType() *int32 {
	return s.AddressType
}

func (s *SingleSendMailAdvanceRequest) GetAttachments() []*SingleSendMailAdvanceRequestAttachments {
	return s.Attachments
}

func (s *SingleSendMailAdvanceRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailAdvanceRequest) GetFromAlias() *string {
	return s.FromAlias
}

func (s *SingleSendMailAdvanceRequest) GetHeaders() *string {
	return s.Headers
}

func (s *SingleSendMailAdvanceRequest) GetHtmlBody() *string {
	return s.HtmlBody
}

func (s *SingleSendMailAdvanceRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *SingleSendMailAdvanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SingleSendMailAdvanceRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *SingleSendMailAdvanceRequest) GetReplyAddressAlias() *string {
	return s.ReplyAddressAlias
}

func (s *SingleSendMailAdvanceRequest) GetReplyToAddress() *bool {
	return s.ReplyToAddress
}

func (s *SingleSendMailAdvanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SingleSendMailAdvanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SingleSendMailAdvanceRequest) GetSubject() *string {
	return s.Subject
}

func (s *SingleSendMailAdvanceRequest) GetTagName() *string {
	return s.TagName
}

func (s *SingleSendMailAdvanceRequest) GetTextBody() *string {
	return s.TextBody
}

func (s *SingleSendMailAdvanceRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *SingleSendMailAdvanceRequest) GetUnSubscribeFilterLevel() *string {
	return s.UnSubscribeFilterLevel
}

func (s *SingleSendMailAdvanceRequest) GetUnSubscribeLinkType() *string {
	return s.UnSubscribeLinkType
}

func (s *SingleSendMailAdvanceRequest) SetAccountName(v string) *SingleSendMailAdvanceRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetAddressType(v int32) *SingleSendMailAdvanceRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetAttachments(v []*SingleSendMailAdvanceRequestAttachments) *SingleSendMailAdvanceRequest {
	s.Attachments = v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetClickTrace(v string) *SingleSendMailAdvanceRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetFromAlias(v string) *SingleSendMailAdvanceRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetHeaders(v string) *SingleSendMailAdvanceRequest {
	s.Headers = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetHtmlBody(v string) *SingleSendMailAdvanceRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetIpPoolId(v string) *SingleSendMailAdvanceRequest {
	s.IpPoolId = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetOwnerId(v int64) *SingleSendMailAdvanceRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetReplyAddress(v string) *SingleSendMailAdvanceRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetReplyAddressAlias(v string) *SingleSendMailAdvanceRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetReplyToAddress(v bool) *SingleSendMailAdvanceRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetResourceOwnerAccount(v string) *SingleSendMailAdvanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetResourceOwnerId(v int64) *SingleSendMailAdvanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetSubject(v string) *SingleSendMailAdvanceRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetTagName(v string) *SingleSendMailAdvanceRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetTextBody(v string) *SingleSendMailAdvanceRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetToAddress(v string) *SingleSendMailAdvanceRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetUnSubscribeFilterLevel(v string) *SingleSendMailAdvanceRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetUnSubscribeLinkType(v string) *SingleSendMailAdvanceRequest {
	s.UnSubscribeLinkType = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type SingleSendMailAdvanceRequestAttachments struct {
	AttachmentName      *string   `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	AttachmentUrlObject io.Reader `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
}

func (s SingleSendMailAdvanceRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailAdvanceRequestAttachments) GoString() string {
	return s.String()
}

func (s *SingleSendMailAdvanceRequestAttachments) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *SingleSendMailAdvanceRequestAttachments) GetAttachmentUrlObject() io.Reader {
	return s.AttachmentUrlObject
}

func (s *SingleSendMailAdvanceRequestAttachments) SetAttachmentName(v string) *SingleSendMailAdvanceRequestAttachments {
	s.AttachmentName = &v
	return s
}

func (s *SingleSendMailAdvanceRequestAttachments) SetAttachmentUrlObject(v io.Reader) *SingleSendMailAdvanceRequestAttachments {
	s.AttachmentUrlObject = v
	return s
}

func (s *SingleSendMailAdvanceRequestAttachments) Validate() error {
	return dara.Validate(s)
}
