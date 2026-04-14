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
	SetBccAddress(v string) *SingleSendMailAdvanceRequest
	GetBccAddress() *string
	SetClickTrace(v string) *SingleSendMailAdvanceRequest
	GetClickTrace() *string
	SetDomainAuth(v bool) *SingleSendMailAdvanceRequest
	GetDomainAuth() *bool
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
	SetTemplate(v *SingleSendMailAdvanceRequestTemplate) *SingleSendMailAdvanceRequest
	GetTemplate() *SingleSendMailAdvanceRequestTemplate
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
	Attachments []*SingleSendMailAdvanceRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
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
	TagName  *string                               `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Template *SingleSendMailAdvanceRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
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

func (s *SingleSendMailAdvanceRequest) GetBccAddress() *string {
	return s.BccAddress
}

func (s *SingleSendMailAdvanceRequest) GetClickTrace() *string {
	return s.ClickTrace
}

func (s *SingleSendMailAdvanceRequest) GetDomainAuth() *bool {
	return s.DomainAuth
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

func (s *SingleSendMailAdvanceRequest) GetTemplate() *SingleSendMailAdvanceRequestTemplate {
	return s.Template
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

func (s *SingleSendMailAdvanceRequest) SetBccAddress(v string) *SingleSendMailAdvanceRequest {
	s.BccAddress = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetClickTrace(v string) *SingleSendMailAdvanceRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailAdvanceRequest) SetDomainAuth(v bool) *SingleSendMailAdvanceRequest {
	s.DomainAuth = &v
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

func (s *SingleSendMailAdvanceRequest) SetTemplate(v *SingleSendMailAdvanceRequestTemplate) *SingleSendMailAdvanceRequest {
	s.Template = v
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

type SingleSendMailAdvanceRequestAttachments struct {
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

type SingleSendMailAdvanceRequestTemplate struct {
	TemplateData map[string]*string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	// example:
	//
	// xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SingleSendMailAdvanceRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailAdvanceRequestTemplate) GoString() string {
	return s.String()
}

func (s *SingleSendMailAdvanceRequestTemplate) GetTemplateData() map[string]*string {
	return s.TemplateData
}

func (s *SingleSendMailAdvanceRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SingleSendMailAdvanceRequestTemplate) SetTemplateData(v map[string]*string) *SingleSendMailAdvanceRequestTemplate {
	s.TemplateData = v
	return s
}

func (s *SingleSendMailAdvanceRequestTemplate) SetTemplateId(v string) *SingleSendMailAdvanceRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *SingleSendMailAdvanceRequestTemplate) Validate() error {
	return dara.Validate(s)
}
