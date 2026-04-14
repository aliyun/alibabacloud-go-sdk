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
	SetBccAddress(v string) *SingleSendMailShrinkRequest
	GetBccAddress() *string
	SetClickTrace(v string) *SingleSendMailShrinkRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *SingleSendMailShrinkRequest
	GetDomainAuth() *bool
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
	// The sending address configured in the management console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Address type. Values:
	//
	// 0: Random account
	//
	// 1: Sending address
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Only supported for use with the new version of the SDK; not currently supported by openapi and signature mechanisms.
	Attachments []*SingleSendMailShrinkRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1@example.com,2@example.com
	BccAddress *string `json:"BccAddress,omitempty" xml:"BccAddress,omitempty"`
	// 1: Enable data tracking function
	//
	// 0 (default): Disable data tracking function.
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	DomainAuth *bool   `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// Sender alias, with a maximum length of 15 characters.
	//
	// For example, if the sender alias is set to "Xiaohong" and the sending address is test***@example.net, the recipient will see the sending address as "Xiaohong" <test***@example.net>.
	//
	// example:
	//
	// Xiaohong
	FromAlias *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	// Currently, the standard fields that can be added to the email header are Message-ID, List-Unsubscribe, and List-Unsubscribe-Post. Standard fields will overwrite the existing values in the email header, while non-standard fields need to start with X-User- and will be appended to the email header.
	//
	// Currently, up to 10 headers can be passed in JSON format, and both standard and non-standard fields must comply with the syntax requirements for headers.
	//
	// example:
	//
	// {
	//
	//   "Message-ID": "<msg0001@example.com>",
	//
	//   "X-User-UID1": "UID-1-000001",
	//
	//   "X-User-UID2": "UID-2-000001"
	//
	// }
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// Email HTML body, limited to 80K by the SDK. Note: HtmlBody and TextBody are for different types of email content, and one of them must be provided.
	//
	// example:
	//
	// body
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	// dedicated IP pool ID. Users who have purchased an dedicated IP can use this parameter to specify the outgoing IP for this email.
	//
	// example:
	//
	// xxx
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Reply-to address
	//
	// example:
	//
	// test2***@example.net
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// Reply-to address alias
	//
	// example:
	//
	// Xiaohong
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// Whether to enable the reply-to address configured in the management console (the status must be verified). The value range is the string `true` or `false` (not a boolean value).
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Email subject, with a maximum length of 100 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// A tag created in the email push console, used to categorize batches of sent emails. You can use tags to query the sending status of each batch. Additionally, if the email tracking feature is enabled, you must use an email tag when sending emails.
	//
	// example:
	//
	// test
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Email text body, limited to 80K by the SDK. Note: HtmlBody and TextBody are for different types of email content, and one of them must be provided.
	//
	// example:
	//
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// Recipient addresses. Multiple email addresses can be separated by commas, with a maximum of 100 addresses (supports mailing lists).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// Filtering level. Refer to the [Unsubscribe Function Link Generation and Filtering Mechanism](https://help.aliyun.com/document_detail/2689048.html) document.
	//
	// disabled: Do not filter
	//
	// default: Use the default strategy, bulk addresses use the sending address level filtering
	//
	// mailfrom: Sending address level filtering
	//
	// mailfrom_domain: Sending domain level filtering
	//
	// edm_id: Account level filtering
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// Type of generated unsubscribe link. Refer to the [Unsubscribe Function Link Generation and Filtering Mechanism](https://help.aliyun.com/document_detail/2689048.html) document.
	//
	// disabled: Do not generate
	//
	// default: Use the default strategy: Generate unsubscribe links for bulk-type sending addresses to specific domains, such as those containing the keywords "gmail", "yahoo",
	//
	// "google", "aol.com", "hotmail",
	//
	// "outlook", "ymail.com", etc.
	//
	// zh-cn: Generate, for future content preparation
	//
	// en-us: Generate, for future content preparation
	//
	// example:
	//
	// default
	UnSubscribeLinkType *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
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

func (s *SingleSendMailShrinkRequest) GetBccAddress() *string {
	return s.BccAddress
}

func (s *SingleSendMailShrinkRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailShrinkRequest) GetDomainAuth() *bool {
	return s.DomainAuth
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

func (s *SingleSendMailShrinkRequest) SetBccAddress(v string) *SingleSendMailShrinkRequest {
	s.BccAddress = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetClickTrace(v string) *SingleSendMailShrinkRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailShrinkRequest) SetDomainAuth(v bool) *SingleSendMailShrinkRequest {
	s.DomainAuth = &v
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
	// Only supported for use with the new version of the SDK; not currently supported by openapi and signature mechanisms.
	//
	// example:
	//
	// test.txt
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// Only supported for use with the new version of the SDK; not currently supported by openapi and signature mechanisms.
	//
	// example:
	//
	// C:\\Users\\Downloads\\test.txt
	AttachmentUrl *string `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
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
