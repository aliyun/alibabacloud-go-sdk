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
	SetBccAddress(v string) *SingleSendMailRequest
	GetBccAddress() *string
	SetClickTrace(v string) *SingleSendMailRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *SingleSendMailRequest
	GetDomainAuth() *bool
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
	SetTemplate(v *SingleSendMailRequestTemplate) *SingleSendMailRequest
	GetTemplate() *SingleSendMailRequestTemplate
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
	Attachments []*SingleSendMailRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
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
	TagName  *string                        `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Template *SingleSendMailRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
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

func (s *SingleSendMailRequest) GetBccAddress() *string {
	return s.BccAddress
}

func (s *SingleSendMailRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailRequest) GetDomainAuth() *bool {
	return s.DomainAuth
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

func (s *SingleSendMailRequest) GetTemplate() *SingleSendMailRequestTemplate {
	return s.Template
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

func (s *SingleSendMailRequest) SetBccAddress(v string) *SingleSendMailRequest {
	s.BccAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetClickTrace(v string) *SingleSendMailRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailRequest) SetDomainAuth(v bool) *SingleSendMailRequest {
	s.DomainAuth = &v
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

func (s *SingleSendMailRequest) SetTemplate(v *SingleSendMailRequestTemplate) *SingleSendMailRequest {
	s.Template = v
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
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SingleSendMailRequestAttachments struct {
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

type SingleSendMailRequestTemplate struct {
	TemplateData map[string]*string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	// example:
	//
	// xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SingleSendMailRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailRequestTemplate) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequestTemplate) GetTemplateData() map[string]*string {
	return s.TemplateData
}

func (s *SingleSendMailRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SingleSendMailRequestTemplate) SetTemplateData(v map[string]*string) *SingleSendMailRequestTemplate {
	s.TemplateData = v
	return s
}

func (s *SingleSendMailRequestTemplate) SetTemplateId(v string) *SingleSendMailRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *SingleSendMailRequestTemplate) Validate() error {
	return dara.Validate(s)
}
