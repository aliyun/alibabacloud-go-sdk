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
	// The sender address configured in the Direct Mail console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The address type. Valid values:
	//
	// `0`: A random account.
	//
	// `1`: A sender address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// This feature is available only through the latest SDKs. It is not supported for OpenAPI calls or signature-based authentication. For more information, see [How do I send an email with an attachment by using an SDK?](https://help.aliyun.com/document_detail/2937843.html).
	Attachments []*SingleSendMailAdvanceRequestAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// - A comma-separated list of BCC recipients.
	//
	// - The system sends a copy of the email to each BCC recipient. The BCC information is hidden from all recipients, including those specified in `ToAddress` and `BccAddress`.
	//
	// - To protect privacy, email tracking features (such as open and click tracking) are disabled for emails sent to BCC recipients. However, billing and sending status are still tracked.
	//
	// - A maximum of two BCC recipients are allowed per request.
	//
	// Note: The `SingleSendMail` API operation does not support a CC field. To send carbon copies, use SMTP.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1@example.com,2@example.com
	BccAddress *string `json:"BccAddress,omitempty" xml:"BccAddress,omitempty"`
	// Specifies whether to enable click tracking. Valid values: `"1"` enables click tracking, and `"0"` disables it (default).
	//
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	// Specifies whether to enable domain-level authentication.
	//
	// - `true`
	//
	// - `false`
	//
	// This parameter is used only for domain-level authentication. Ignore it for sender address-level authentication.
	//
	// 1\\. Create the address `domain-auth-created-by-system@example.com` in the console. The prefix must be fixed, and the suffix must be your domain.
	//
	// 2\\.
	//
	// **API scenario**
	//
	// Set `AccountName` to your domain. Recipients will see the sender as `domain-auth-created-by-system@example.com`.
	//
	// **SMTP scenario**
	//
	// a. Call the `ModifyPWByDomain` API operation to set a password for the domain.
	//
	// b. Authenticate with the domain and the configured password. Pass a custom address, such as `user@example.com`, as the actual sender in the `MAIL FROM` command. Recipients will see `user@example.com` as the sender.
	//
	// example:
	//
	// true
	DomainAuth *bool `json:"DomainAuth,omitempty" xml:"DomainAuth,omitempty"`
	// The sender name. It must be 15 characters or shorter.
	//
	// For example, if you set the sender name to "Xiaohong" and the sender address is `test***@example.net`, the recipient sees the sender as "Xiaohong" \\<test\\*\\*\\*@example.net>.
	//
	// example:
	//
	// Jane
	FromAlias *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	// Custom email header settings.
	//
	// Both standard and non-standard fields must comply with standard header syntax. You can specify up to 10 headers for an API call. Excess headers are ignored. This limit does not apply to SMTP.
	//
	// 1\\. Standard fields
	//
	// `Message-ID`, `List-Unsubscribe`, `List-Unsubscribe-Post`
	//
	// Standard fields overwrite existing values in the email header.
	//
	// 2\\. Non-standard fields
	//
	// Case-insensitive.
	//
	// a. Fields starting with `X-User-`: These are not pushed to EventBridge or Message Service (MNS). This prefix is required only for API calls, not for SMTP.
	//
	// b. Fields starting with `X-User-Notify-`: These are pushed to EventBridge and MNS. This is supported for both API and SMTP calls.
	//
	// When pushed to EventBridge or MNS, the header object will contain these fields.
	//
	// example:
	//
	// {
	//
	//       "Message-ID": "<d52ce63e-a0d5-4f95-b6a9-e1256a44f5fb@example.net>",
	//
	//       "X-User-UID1": "UID-1-000001",
	//
	//       "X-User-UID2": "UID-2-000001",
	//
	//       "X-User-Notify-UID1": "UID-3-000001",
	//
	//       "X-User-Notify-UID2": "UID-4-000001"
	//
	//
	//
	// }
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The HTML body of the email.
	//
	// Note: You must specify either `HtmlBody` or `TextBody`.
	//
	// - The size of the body is limited to approximately 80 KB when passed as a URL parameter.
	//
	// - For recent SDKs (Java 1.4.0+, Python 3 1.4.0+, and PHP 1.4.0+), the request body is limited to approximately 8 MB.
	//
	// example:
	//
	// body
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	// The ID of the dedicated IP pool. If you have purchased dedicated IPs, you can use this parameter to select which dedicated IP pool to use for sending the email. For more information, see [Dedicated IP](https://help.aliyun.com/document_detail/2932088.html).
	//
	// example:
	//
	// e4xxxxxe-4xx0-4xx3-8xxa-74cxxxxx1cef
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reply-to address.
	//
	// example:
	//
	// test2***@example.net
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// The name displayed for the reply-to address.
	//
	// example:
	//
	// Jane
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// Specifies whether to use the default reply-to address configured in the console. This address must be verified. Valid values: true, false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The subject of the email, with a maximum length of 256 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// A tag for categorizing email batches, which you can create in the Direct Mail console. Tags allow you to query the sending status of each batch and are required if you enable email tracking. The tag must be 1 to 128 characters long and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The template information for sending a templated email.
	Template *SingleSendMailAdvanceRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// The text body of the email.
	//
	// Note: You must specify either `HtmlBody` or `TextBody`.
	//
	// - The size of the body is limited to approximately 80 KB when passed as a URL parameter.
	//
	// - For recent SDKs (Java 1.4.0+, Python 3 1.4.0+, and PHP 1.4.0+), the request body is limited to approximately 8 MB.
	//
	// example:
	//
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// The destination email address(es). To specify multiple addresses, separate them with commas (up to 100).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// The filtering level. For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// `disabled`: No filtering.
	//
	// `default`: Uses the default policy. For batch addresses, filtering is applied at the sender address level.
	//
	// `mailfrom`: Filters at the sender address level.
	//
	// `mailfrom_domain`: Filters at the sender domain level.
	//
	// `edm_id`: Filters at the account level.
	//
	// example:
	//
	// mailfrom_domain
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	// `disabled`: Does not generate an unsubscribe link.
	//
	// `default`: Uses the default policy. For batch sender addresses, an unsubscribe link is generated when sending to specific domains containing keywords such as "gmail", "yahoo",
	//
	// "google", "aol.com", "hotmail",
	//
	// "outlook", and "ymail.com". For more information, see [Unsubscribe link generation and filtering mechanism](https://help.aliyun.com/document_detail/2689048.html).
	//
	// The display language is automatically determined based on the recipient\\"s browser settings.
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
	// The filename of the attachment.
	//
	// example:
	//
	// test.txt
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The local file path of the attachment that the SDK will use.
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
	// The variables and their values for the template.
	TemplateData map[string]*string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	// The template ID.
	//
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
