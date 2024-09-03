// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddIpfilterRequest struct {
	// This parameter is required.
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddIpfilterRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpfilterRequest) GoString() string {
	return s.String()
}

func (s *AddIpfilterRequest) SetIpAddress(v string) *AddIpfilterRequest {
	s.IpAddress = &v
	return s
}

func (s *AddIpfilterRequest) SetOwnerId(v int64) *AddIpfilterRequest {
	s.OwnerId = &v
	return s
}

func (s *AddIpfilterRequest) SetResourceOwnerAccount(v string) *AddIpfilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddIpfilterRequest) SetResourceOwnerId(v int64) *AddIpfilterRequest {
	s.ResourceOwnerId = &v
	return s
}

type AddIpfilterResponseBody struct {
	// example:
	//
	// 10795
	IpFilterId *string `json:"IpFilterId,omitempty" xml:"IpFilterId,omitempty"`
	// example:
	//
	// 0E9282E8-DC08-5445-8FB0-B9F0CA28B249
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpfilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpfilterResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpfilterResponseBody) SetIpFilterId(v string) *AddIpfilterResponseBody {
	s.IpFilterId = &v
	return s
}

func (s *AddIpfilterResponseBody) SetRequestId(v string) *AddIpfilterResponseBody {
	s.RequestId = &v
	return s
}

type AddIpfilterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpfilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpfilterResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpfilterResponse) GoString() string {
	return s.String()
}

func (s *AddIpfilterResponse) SetHeaders(v map[string]*string) *AddIpfilterResponse {
	s.Headers = v
	return s
}

func (s *AddIpfilterResponse) SetStatusCode(v int32) *AddIpfilterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpfilterResponse) SetBody(v *AddIpfilterResponseBody) *AddIpfilterResponse {
	s.Body = v
	return s
}

type ApproveReplyMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s ApproveReplyMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveReplyMailAddressRequest) GoString() string {
	return s.String()
}

func (s *ApproveReplyMailAddressRequest) SetOwnerId(v int64) *ApproveReplyMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) SetResourceOwnerAccount(v string) *ApproveReplyMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) SetResourceOwnerId(v int64) *ApproveReplyMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) SetTicket(v string) *ApproveReplyMailAddressRequest {
	s.Ticket = &v
	return s
}

type ApproveReplyMailAddressResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveReplyMailAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveReplyMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveReplyMailAddressResponseBody) SetRequestId(v string) *ApproveReplyMailAddressResponseBody {
	s.RequestId = &v
	return s
}

type ApproveReplyMailAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveReplyMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveReplyMailAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveReplyMailAddressResponse) GoString() string {
	return s.String()
}

func (s *ApproveReplyMailAddressResponse) SetHeaders(v map[string]*string) *ApproveReplyMailAddressResponse {
	s.Headers = v
	return s
}

func (s *ApproveReplyMailAddressResponse) SetStatusCode(v int32) *ApproveReplyMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveReplyMailAddressResponse) SetBody(v *ApproveReplyMailAddressResponseBody) *ApproveReplyMailAddressResponse {
	s.Body = v
	return s
}

type BatchSendMailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test2
	ReceiversName        *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias    *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test3
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1
	TemplateName           *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	UnSubscribeLinkType    *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s BatchSendMailRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMailRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMailRequest) SetAccountName(v string) *BatchSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *BatchSendMailRequest) SetAddressType(v int32) *BatchSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *BatchSendMailRequest) SetClickTrace(v string) *BatchSendMailRequest {
	s.ClickTrace = &v
	return s
}

func (s *BatchSendMailRequest) SetOwnerId(v int64) *BatchSendMailRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSendMailRequest) SetReceiversName(v string) *BatchSendMailRequest {
	s.ReceiversName = &v
	return s
}

func (s *BatchSendMailRequest) SetReplyAddress(v string) *BatchSendMailRequest {
	s.ReplyAddress = &v
	return s
}

func (s *BatchSendMailRequest) SetReplyAddressAlias(v string) *BatchSendMailRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *BatchSendMailRequest) SetResourceOwnerAccount(v string) *BatchSendMailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchSendMailRequest) SetResourceOwnerId(v int64) *BatchSendMailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BatchSendMailRequest) SetTagName(v string) *BatchSendMailRequest {
	s.TagName = &v
	return s
}

func (s *BatchSendMailRequest) SetTemplateName(v string) *BatchSendMailRequest {
	s.TemplateName = &v
	return s
}

func (s *BatchSendMailRequest) SetUnSubscribeFilterLevel(v string) *BatchSendMailRequest {
	s.UnSubscribeFilterLevel = &v
	return s
}

func (s *BatchSendMailRequest) SetUnSubscribeLinkType(v string) *BatchSendMailRequest {
	s.UnSubscribeLinkType = &v
	return s
}

type BatchSendMailResponseBody struct {
	// example:
	//
	// xxx
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// example:
	//
	// 12D086F6-8F31-4658-84C1-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSendMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMailResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMailResponseBody) SetEnvId(v string) *BatchSendMailResponseBody {
	s.EnvId = &v
	return s
}

func (s *BatchSendMailResponseBody) SetRequestId(v string) *BatchSendMailResponseBody {
	s.RequestId = &v
	return s
}

type BatchSendMailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendMailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendMailResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMailResponse) GoString() string {
	return s.String()
}

func (s *BatchSendMailResponse) SetHeaders(v map[string]*string) *BatchSendMailResponse {
	s.Headers = v
	return s
}

func (s *BatchSendMailResponse) SetStatusCode(v int32) *BatchSendMailResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendMailResponse) SetBody(v *BatchSendMailResponseBody) *BatchSendMailResponse {
	s.Body = v
	return s
}

type CheckDomainRequest struct {
	// The ID of the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 153345
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainRequest) SetDomainId(v int32) *CheckDomainRequest {
	s.DomainId = &v
	return s
}

func (s *CheckDomainRequest) SetOwnerId(v int64) *CheckDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDomainRequest) SetResourceOwnerAccount(v string) *CheckDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDomainRequest) SetResourceOwnerId(v int64) *CheckDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckDomainResponseBody struct {
	// The status of the domain name. Indicates whether the domain name is verified and available.
	//
	// 	- 0: indicates that the domain name is verified and available.
	//
	// 	- 1: indicates that the domain name fails to be verified and is unavailable.
	//
	// 	- 2: indicates that the domain name is available, but not filed or configured with a CNAME record.
	//
	// 	- 3: indicates that the domain name is available but not filed.
	//
	// 	- 4: indicates that the domain name is available but not configured with a CNAME record.
	//
	// example:
	//
	// 1
	DomainStatus *int32 `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F0B82E83-A1D9-4FE6-97D2-F4B231F80B02
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) SetDomainStatus(v int32) *CheckDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

type CheckDomainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainResponse) SetHeaders(v map[string]*string) *CheckDomainResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainResponse) SetStatusCode(v int32) *CheckDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainResponse) SetBody(v *CheckDomainResponseBody) *CheckDomainResponse {
	s.Body = v
	return s
}

type CheckReplyToMailAddressRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 126545
	MailAddressId *int32 `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckReplyToMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckReplyToMailAddressRequest) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressRequest) SetLang(v string) *CheckReplyToMailAddressRequest {
	s.Lang = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetMailAddressId(v int32) *CheckReplyToMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetOwnerId(v int64) *CheckReplyToMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetRegion(v string) *CheckReplyToMailAddressRequest {
	s.Region = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetResourceOwnerAccount(v string) *CheckReplyToMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetResourceOwnerId(v int64) *CheckReplyToMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckReplyToMailAddressResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckReplyToMailAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckReplyToMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressResponseBody) SetRequestId(v string) *CheckReplyToMailAddressResponseBody {
	s.RequestId = &v
	return s
}

type CheckReplyToMailAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckReplyToMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckReplyToMailAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckReplyToMailAddressResponse) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressResponse) SetHeaders(v map[string]*string) *CheckReplyToMailAddressResponse {
	s.Headers = v
	return s
}

func (s *CheckReplyToMailAddressResponse) SetStatusCode(v int32) *CheckReplyToMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckReplyToMailAddressResponse) SetBody(v *CheckReplyToMailAddressResponseBody) *CheckReplyToMailAddressResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDomainRequest) SetOwnerId(v int64) *CreateDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDomainRequest) SetResourceOwnerAccount(v string) *CreateDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDomainRequest) SetResourceOwnerId(v int64) *CreateDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateDomainResponseBody struct {
	// example:
	//
	// 158910
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// B49AD828-25D1-488C-90B7-8853C1944486
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetDomainId(v string) *CreateDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateMailAddressRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// reply@aliyun.com
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// batch
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s CreateMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMailAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateMailAddressRequest) SetAccountName(v string) *CreateMailAddressRequest {
	s.AccountName = &v
	return s
}

func (s *CreateMailAddressRequest) SetOwnerId(v int64) *CreateMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMailAddressRequest) SetReplyAddress(v string) *CreateMailAddressRequest {
	s.ReplyAddress = &v
	return s
}

func (s *CreateMailAddressRequest) SetResourceOwnerAccount(v string) *CreateMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMailAddressRequest) SetResourceOwnerId(v int64) *CreateMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMailAddressRequest) SetSendtype(v string) *CreateMailAddressRequest {
	s.Sendtype = &v
	return s
}

type CreateMailAddressResponseBody struct {
	// example:
	//
	// xxx
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	// example:
	//
	// 95A7D497-F8DD-4834-B81E-C1783236E55F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMailAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMailAddressResponseBody) SetMailAddressId(v string) *CreateMailAddressResponseBody {
	s.MailAddressId = &v
	return s
}

func (s *CreateMailAddressResponseBody) SetRequestId(v string) *CreateMailAddressResponseBody {
	s.RequestId = &v
	return s
}

type CreateMailAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMailAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMailAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateMailAddressResponse) SetHeaders(v map[string]*string) *CreateMailAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateMailAddressResponse) SetStatusCode(v int32) *CreateMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMailAddressResponse) SetBody(v *CreateMailAddressResponseBody) *CreateMailAddressResponse {
	s.Body = v
	return s
}

type CreateReceiverRequest struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vme
	ReceiversAlias *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xme
	ReceiversName        *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateReceiverRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiverRequest) GoString() string {
	return s.String()
}

func (s *CreateReceiverRequest) SetDesc(v string) *CreateReceiverRequest {
	s.Desc = &v
	return s
}

func (s *CreateReceiverRequest) SetOwnerId(v int64) *CreateReceiverRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiversAlias(v string) *CreateReceiverRequest {
	s.ReceiversAlias = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiversName(v string) *CreateReceiverRequest {
	s.ReceiversName = &v
	return s
}

func (s *CreateReceiverRequest) SetResourceOwnerAccount(v string) *CreateReceiverRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateReceiverRequest) SetResourceOwnerId(v int64) *CreateReceiverRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateReceiverResponseBody struct {
	// example:
	//
	// xxxx
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReceiverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiverResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReceiverResponseBody) SetReceiverId(v string) *CreateReceiverResponseBody {
	s.ReceiverId = &v
	return s
}

func (s *CreateReceiverResponseBody) SetRequestId(v string) *CreateReceiverResponseBody {
	s.RequestId = &v
	return s
}

type CreateReceiverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReceiverResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiverResponse) GoString() string {
	return s.String()
}

func (s *CreateReceiverResponse) SetHeaders(v map[string]*string) *CreateReceiverResponse {
	s.Headers = v
	return s
}

func (s *CreateReceiverResponse) SetStatusCode(v int32) *CreateReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReceiverResponse) SetBody(v *CreateReceiverResponseBody) *CreateReceiverResponse {
	s.Body = v
	return s
}

type CreateTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagDescription       *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vme
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s CreateTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) SetOwnerId(v int64) *CreateTagRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTagRequest) SetResourceOwnerAccount(v string) *CreateTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTagRequest) SetResourceOwnerId(v int64) *CreateTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTagRequest) SetTagDescription(v string) *CreateTagRequest {
	s.TagDescription = &v
	return s
}

func (s *CreateTagRequest) SetTagName(v string) *CreateTagRequest {
	s.TagName = &v
	return s
}

type CreateTagResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// xxxx
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetTagId(v string) *CreateTagResponseBody {
	s.TagId = &v
	return s
}

type CreateTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponse) GoString() string {
	return s.String()
}

func (s *CreateTagResponse) SetHeaders(v map[string]*string) *CreateTagResponse {
	s.Headers = v
	return s
}

func (s *CreateTagResponse) SetStatusCode(v int32) *CreateTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagResponse) SetBody(v *CreateTagResponseBody) *CreateTagResponse {
	s.Body = v
	return s
}

type CreateUserSuppressionRequest struct {
	Address              *string `json:"Address,omitempty" xml:"Address,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateUserSuppressionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSuppressionRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSuppressionRequest) SetAddress(v string) *CreateUserSuppressionRequest {
	s.Address = &v
	return s
}

func (s *CreateUserSuppressionRequest) SetOwnerId(v int64) *CreateUserSuppressionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUserSuppressionRequest) SetResourceOwnerAccount(v string) *CreateUserSuppressionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateUserSuppressionRequest) SetResourceOwnerId(v int64) *CreateUserSuppressionRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateUserSuppressionResponseBody struct {
	// example:
	//
	// 1A846D66-5EC7-551B-9687-5BF1963DCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 59511
	SuppressionId *string `json:"SuppressionId,omitempty" xml:"SuppressionId,omitempty"`
}

func (s CreateUserSuppressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSuppressionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserSuppressionResponseBody) SetRequestId(v string) *CreateUserSuppressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserSuppressionResponseBody) SetSuppressionId(v string) *CreateUserSuppressionResponseBody {
	s.SuppressionId = &v
	return s
}

type CreateUserSuppressionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserSuppressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserSuppressionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserSuppressionResponse) GoString() string {
	return s.String()
}

func (s *CreateUserSuppressionResponse) SetHeaders(v map[string]*string) *CreateUserSuppressionResponse {
	s.Headers = v
	return s
}

func (s *CreateUserSuppressionResponse) SetStatusCode(v int32) *CreateUserSuppressionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserSuppressionResponse) SetBody(v *CreateUserSuppressionResponseBody) *CreateUserSuppressionResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 326***
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetDomainId(v int32) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainRequest) SetOwnerId(v int64) *DeleteDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDomainRequest) SetResourceOwnerAccount(v string) *DeleteDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDomainRequest) SetResourceOwnerId(v int64) *DeleteDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDomainResponseBody struct {
	// example:
	//
	// E3DFF97B-00CF-5333-8125-3D6819471984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteInvalidAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s DeleteInvalidAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInvalidAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteInvalidAddressRequest) SetOwnerId(v int64) *DeleteInvalidAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInvalidAddressRequest) SetResourceOwnerAccount(v string) *DeleteInvalidAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInvalidAddressRequest) SetResourceOwnerId(v int64) *DeleteInvalidAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInvalidAddressRequest) SetToAddress(v string) *DeleteInvalidAddressRequest {
	s.ToAddress = &v
	return s
}

type DeleteInvalidAddressResponseBody struct {
	// example:
	//
	// 2D086F6-xxxx-xxxx-xxxx-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInvalidAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInvalidAddressResponseBody) SetRequestId(v string) *DeleteInvalidAddressResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInvalidAddressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInvalidAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInvalidAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteInvalidAddressResponse) SetHeaders(v map[string]*string) *DeleteInvalidAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteInvalidAddressResponse) SetStatusCode(v int32) *DeleteInvalidAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInvalidAddressResponse) SetBody(v *DeleteInvalidAddressResponseBody) *DeleteInvalidAddressResponse {
	s.Body = v
	return s
}

type DeleteIpfilterByEdmIdRequest struct {
	FromType *int32 `json:"FromType,omitempty" xml:"FromType,omitempty"`
	// example:
	//
	// 10120
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpfilterByEdmIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpfilterByEdmIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpfilterByEdmIdRequest) SetFromType(v int32) *DeleteIpfilterByEdmIdRequest {
	s.FromType = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetId(v string) *DeleteIpfilterByEdmIdRequest {
	s.Id = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetOwnerId(v int64) *DeleteIpfilterByEdmIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetResourceOwnerAccount(v string) *DeleteIpfilterByEdmIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetResourceOwnerId(v int64) *DeleteIpfilterByEdmIdRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteIpfilterByEdmIdResponseBody struct {
	// example:
	//
	// E3DFF97B-00CF-5333-8125-3D6819471984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpfilterByEdmIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpfilterByEdmIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpfilterByEdmIdResponseBody) SetRequestId(v string) *DeleteIpfilterByEdmIdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpfilterByEdmIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpfilterByEdmIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpfilterByEdmIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpfilterByEdmIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpfilterByEdmIdResponse) SetHeaders(v map[string]*string) *DeleteIpfilterByEdmIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpfilterByEdmIdResponse) SetStatusCode(v int32) *DeleteIpfilterByEdmIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpfilterByEdmIdResponse) SetBody(v *DeleteIpfilterByEdmIdResponseBody) *DeleteIpfilterByEdmIdResponse {
	s.Body = v
	return s
}

type DeleteMailAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23457
	MailAddressId        *int32  `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteMailAddressRequest) SetMailAddressId(v int32) *DeleteMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *DeleteMailAddressRequest) SetOwnerId(v int64) *DeleteMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMailAddressRequest) SetResourceOwnerAccount(v string) *DeleteMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMailAddressRequest) SetResourceOwnerId(v int64) *DeleteMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteMailAddressResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMailAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMailAddressResponseBody) SetRequestId(v string) *DeleteMailAddressResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMailAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMailAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteMailAddressResponse) SetHeaders(v map[string]*string) *DeleteMailAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteMailAddressResponse) SetStatusCode(v int32) *DeleteMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMailAddressResponse) SetBody(v *DeleteMailAddressResponseBody) *DeleteMailAddressResponse {
	s.Body = v
	return s
}

type DeleteReceiverRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the recipient list.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1453
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteReceiverRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverRequest) GoString() string {
	return s.String()
}

func (s *DeleteReceiverRequest) SetOwnerId(v int64) *DeleteReceiverRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReceiverRequest) SetReceiverId(v string) *DeleteReceiverRequest {
	s.ReceiverId = &v
	return s
}

func (s *DeleteReceiverRequest) SetResourceOwnerAccount(v string) *DeleteReceiverRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReceiverRequest) SetResourceOwnerId(v int64) *DeleteReceiverRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteReceiverResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReceiverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReceiverResponseBody) SetRequestId(v string) *DeleteReceiverResponseBody {
	s.RequestId = &v
	return s
}

type DeleteReceiverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReceiverResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverResponse) GoString() string {
	return s.String()
}

func (s *DeleteReceiverResponse) SetHeaders(v map[string]*string) *DeleteReceiverResponse {
	s.Headers = v
	return s
}

func (s *DeleteReceiverResponse) SetStatusCode(v int32) *DeleteReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReceiverResponse) SetBody(v *DeleteReceiverResponseBody) *DeleteReceiverResponse {
	s.Body = v
	return s
}

type DeleteReceiverDetailRequest struct {
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123534
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteReceiverDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailRequest) SetEmail(v string) *DeleteReceiverDetailRequest {
	s.Email = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetOwnerId(v int64) *DeleteReceiverDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetReceiverId(v string) *DeleteReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetResourceOwnerAccount(v string) *DeleteReceiverDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetResourceOwnerId(v int64) *DeleteReceiverDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteReceiverDetailResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReceiverDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailResponseBody) SetRequestId(v string) *DeleteReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

type DeleteReceiverDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReceiverDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverDetailResponse) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailResponse) SetHeaders(v map[string]*string) *DeleteReceiverDetailResponse {
	s.Headers = v
	return s
}

func (s *DeleteReceiverDetailResponse) SetStatusCode(v int32) *DeleteReceiverDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReceiverDetailResponse) SetBody(v *DeleteReceiverDetailResponseBody) *DeleteReceiverDetailResponse {
	s.Body = v
	return s
}

type DeleteTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) SetOwnerId(v int64) *DeleteTagRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTagRequest) SetResourceOwnerAccount(v string) *DeleteTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTagRequest) SetResourceOwnerId(v int64) *DeleteTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTagRequest) SetTagId(v int32) *DeleteTagRequest {
	s.TagId = &v
	return s
}

type DeleteTagResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagResponse) SetHeaders(v map[string]*string) *DeleteTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagResponse) SetStatusCode(v int32) *DeleteTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagResponse) SetBody(v *DeleteTagResponseBody) *DeleteTagResponse {
	s.Body = v
	return s
}

type DescAccountSummaryRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescAccountSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryRequest) SetOwnerId(v int64) *DescAccountSummaryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescAccountSummaryRequest) SetResourceOwnerAccount(v string) *DescAccountSummaryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescAccountSummaryRequest) SetResourceOwnerId(v int64) *DescAccountSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescAccountSummaryResponseBody struct {
	// example:
	//
	// 2000
	DailyQuota           *int32 `json:"DailyQuota,omitempty" xml:"DailyQuota,omitempty"`
	DailyRemainFreeQuota *int32 `json:"DailyRemainFreeQuota,omitempty" xml:"DailyRemainFreeQuota,omitempty"`
	// example:
	//
	// 0
	DayuStatus *int32 `json:"DayuStatus,omitempty" xml:"DayuStatus,omitempty"`
	// example:
	//
	// 1
	Domains *int32 `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// example:
	//
	// 0
	EnableTimes *int32 `json:"EnableTimes,omitempty" xml:"EnableTimes,omitempty"`
	// example:
	//
	// 0
	MailAddresses *int32 `json:"MailAddresses,omitempty" xml:"MailAddresses,omitempty"`
	// example:
	//
	// 10
	MaxQuotaLevel *int32 `json:"MaxQuotaLevel,omitempty" xml:"MaxQuotaLevel,omitempty"`
	// example:
	//
	// 60000
	MonthQuota *int32 `json:"MonthQuota,omitempty" xml:"MonthQuota,omitempty"`
	// example:
	//
	// 2
	QuotaLevel *int32 `json:"QuotaLevel,omitempty" xml:"QuotaLevel,omitempty"`
	// example:
	//
	// 0
	Receivers       *int32 `json:"Receivers,omitempty" xml:"Receivers,omitempty"`
	RemainFreeQuota *int32 `json:"RemainFreeQuota,omitempty" xml:"RemainFreeQuota,omitempty"`
	// example:
	//
	// 82B295BB-7E69-491F-9896-ECEAFF09E1A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	SmsRecord *int32 `json:"SmsRecord,omitempty" xml:"SmsRecord,omitempty"`
	// example:
	//
	// 0
	SmsSign *int32 `json:"SmsSign,omitempty" xml:"SmsSign,omitempty"`
	// example:
	//
	// 0
	SmsTemplates *int32 `json:"SmsTemplates,omitempty" xml:"SmsTemplates,omitempty"`
	// example:
	//
	// 0
	Tags *int32 `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 1
	Templates *int32 `json:"Templates,omitempty" xml:"Templates,omitempty"`
	// example:
	//
	// 0
	UserStatus *int32 `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescAccountSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryResponseBody) SetDailyQuota(v int32) *DescAccountSummaryResponseBody {
	s.DailyQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDailyRemainFreeQuota(v int32) *DescAccountSummaryResponseBody {
	s.DailyRemainFreeQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDayuStatus(v int32) *DescAccountSummaryResponseBody {
	s.DayuStatus = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDomains(v int32) *DescAccountSummaryResponseBody {
	s.Domains = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetEnableTimes(v int32) *DescAccountSummaryResponseBody {
	s.EnableTimes = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMailAddresses(v int32) *DescAccountSummaryResponseBody {
	s.MailAddresses = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMaxQuotaLevel(v int32) *DescAccountSummaryResponseBody {
	s.MaxQuotaLevel = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMonthQuota(v int32) *DescAccountSummaryResponseBody {
	s.MonthQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetQuotaLevel(v int32) *DescAccountSummaryResponseBody {
	s.QuotaLevel = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetReceivers(v int32) *DescAccountSummaryResponseBody {
	s.Receivers = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetRemainFreeQuota(v int32) *DescAccountSummaryResponseBody {
	s.RemainFreeQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetRequestId(v string) *DescAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsRecord(v int32) *DescAccountSummaryResponseBody {
	s.SmsRecord = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsSign(v int32) *DescAccountSummaryResponseBody {
	s.SmsSign = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsTemplates(v int32) *DescAccountSummaryResponseBody {
	s.SmsTemplates = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetTags(v int32) *DescAccountSummaryResponseBody {
	s.Tags = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetTemplates(v int32) *DescAccountSummaryResponseBody {
	s.Templates = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetUserStatus(v int32) *DescAccountSummaryResponseBody {
	s.UserStatus = &v
	return s
}

type DescAccountSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescAccountSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryResponse) SetHeaders(v map[string]*string) *DescAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescAccountSummaryResponse) SetStatusCode(v int32) *DescAccountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescAccountSummaryResponse) SetBody(v *DescAccountSummaryResponseBody) *DescAccountSummaryResponse {
	s.Body = v
	return s
}

type DescDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13464
	DomainId                  *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequireRealTimeDnsRecords *bool   `json:"RequireRealTimeDnsRecords,omitempty" xml:"RequireRealTimeDnsRecords,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescDomainRequest) GoString() string {
	return s.String()
}

func (s *DescDomainRequest) SetDomainId(v int32) *DescDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DescDomainRequest) SetOwnerId(v int64) *DescDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescDomainRequest) SetRequireRealTimeDnsRecords(v bool) *DescDomainRequest {
	s.RequireRealTimeDnsRecords = &v
	return s
}

func (s *DescDomainRequest) SetResourceOwnerAccount(v string) *DescDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescDomainRequest) SetResourceOwnerId(v int64) *DescDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescDomainResponseBody struct {
	// example:
	//
	// 1
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	// example:
	//
	// 0
	CnameConfirmStatus *string `json:"CnameConfirmStatus,omitempty" xml:"CnameConfirmStatus,omitempty"`
	// example:
	//
	// dmtrace
	CnameRecord *string `json:"CnameRecord,omitempty" xml:"CnameRecord,omitempty"`
	// example:
	//
	// 2019-09-29T12:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	DefaultDomain   *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	DkimAuthStatus  *string `json:"DkimAuthStatus,omitempty" xml:"DkimAuthStatus,omitempty"`
	DkimPublicKey   *string `json:"DkimPublicKey,omitempty" xml:"DkimPublicKey,omitempty"`
	DkimRR          *string `json:"DkimRR,omitempty" xml:"DkimRR,omitempty"`
	DmarcAuthStatus *int32  `json:"DmarcAuthStatus,omitempty" xml:"DmarcAuthStatus,omitempty"`
	DmarcHostRecord *string `json:"DmarcHostRecord,omitempty" xml:"DmarcHostRecord,omitempty"`
	DmarcRecord     *string `json:"DmarcRecord,omitempty" xml:"DmarcRecord,omitempty"`
	DnsDmarc        *string `json:"DnsDmarc,omitempty" xml:"DnsDmarc,omitempty"`
	// example:
	//
	// abc-com.xxxx.com
	DnsMx *string `json:"DnsMx,omitempty" xml:"DnsMx,omitempty"`
	// example:
	//
	// v=xxxx
	DnsSpf *string `json:"DnsSpf,omitempty" xml:"DnsSpf,omitempty"`
	// example:
	//
	// 121309ohdsa
	DnsTxt *string `json:"DnsTxt,omitempty" xml:"DnsTxt,omitempty"`
	// example:
	//
	// 158910
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// abc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// example:
	//
	// 0c40d5f125af4e42892a
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	HostRecord *string `json:"HostRecord,omitempty" xml:"HostRecord,omitempty"`
	// example:
	//
	// 1
	IcpStatus *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	// example:
	//
	// 1
	MxAuthStatus *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	// example:
	//
	// mx01.dm.aliyun.com
	MxRecord *string `json:"MxRecord,omitempty" xml:"MxRecord,omitempty"`
	// example:
	//
	// 51B74264-46B4-43C8-A9A0-6B8E8BC04F34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SpfAuthStatus *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	// example:
	//
	// include:spf1.dm.aliyun.com
	SpfRecord   *string `json:"SpfRecord,omitempty" xml:"SpfRecord,omitempty"`
	SpfRecordV2 *string `json:"SpfRecordV2,omitempty" xml:"SpfRecordV2,omitempty"`
	// example:
	//
	// abc.com
	TlDomainName *string `json:"TlDomainName,omitempty" xml:"TlDomainName,omitempty"`
	// example:
	//
	// tracedm.aliyuncs.com
	TracefRecord *string `json:"TracefRecord,omitempty" xml:"TracefRecord,omitempty"`
}

func (s DescDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescDomainResponseBody) SetCnameAuthStatus(v string) *DescDomainResponseBody {
	s.CnameAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameConfirmStatus(v string) *DescDomainResponseBody {
	s.CnameConfirmStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameRecord(v string) *DescDomainResponseBody {
	s.CnameRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetCreateTime(v string) *DescDomainResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescDomainResponseBody) SetDefaultDomain(v string) *DescDomainResponseBody {
	s.DefaultDomain = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimAuthStatus(v string) *DescDomainResponseBody {
	s.DkimAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimPublicKey(v string) *DescDomainResponseBody {
	s.DkimPublicKey = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimRR(v string) *DescDomainResponseBody {
	s.DkimRR = &v
	return s
}

func (s *DescDomainResponseBody) SetDmarcAuthStatus(v int32) *DescDomainResponseBody {
	s.DmarcAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDmarcHostRecord(v string) *DescDomainResponseBody {
	s.DmarcHostRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetDmarcRecord(v string) *DescDomainResponseBody {
	s.DmarcRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsDmarc(v string) *DescDomainResponseBody {
	s.DnsDmarc = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsMx(v string) *DescDomainResponseBody {
	s.DnsMx = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsSpf(v string) *DescDomainResponseBody {
	s.DnsSpf = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsTxt(v string) *DescDomainResponseBody {
	s.DnsTxt = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainId(v string) *DescDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainName(v string) *DescDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainStatus(v string) *DescDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainType(v string) *DescDomainResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescDomainResponseBody) SetHostRecord(v string) *DescDomainResponseBody {
	s.HostRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetIcpStatus(v string) *DescDomainResponseBody {
	s.IcpStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetMxAuthStatus(v string) *DescDomainResponseBody {
	s.MxAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetMxRecord(v string) *DescDomainResponseBody {
	s.MxRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetRequestId(v string) *DescDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfAuthStatus(v string) *DescDomainResponseBody {
	s.SpfAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfRecord(v string) *DescDomainResponseBody {
	s.SpfRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfRecordV2(v string) *DescDomainResponseBody {
	s.SpfRecordV2 = &v
	return s
}

func (s *DescDomainResponseBody) SetTlDomainName(v string) *DescDomainResponseBody {
	s.TlDomainName = &v
	return s
}

func (s *DescDomainResponseBody) SetTracefRecord(v string) *DescDomainResponseBody {
	s.TracefRecord = &v
	return s
}

type DescDomainResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescDomainResponse) GoString() string {
	return s.String()
}

func (s *DescDomainResponse) SetHeaders(v map[string]*string) *DescDomainResponse {
	s.Headers = v
	return s
}

func (s *DescDomainResponse) SetStatusCode(v int32) *DescDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescDomainResponse) SetBody(v *DescDomainResponseBody) *DescDomainResponse {
	s.Body = v
	return s
}

type GetIpProtectionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetIpProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpProtectionRequest) GoString() string {
	return s.String()
}

func (s *GetIpProtectionRequest) SetOwnerId(v int64) *GetIpProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *GetIpProtectionRequest) SetResourceOwnerAccount(v string) *GetIpProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetIpProtectionRequest) SetResourceOwnerId(v int64) *GetIpProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetIpProtectionResponseBody struct {
	// example:
	//
	// 0
	IpProtection *string `json:"IpProtection,omitempty" xml:"IpProtection,omitempty"`
	// example:
	//
	// B30E5A62-2E64-577D-A70E-8C6781D6C975
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpProtectionResponseBody) SetIpProtection(v string) *GetIpProtectionResponseBody {
	s.IpProtection = &v
	return s
}

func (s *GetIpProtectionResponseBody) SetRequestId(v string) *GetIpProtectionResponseBody {
	s.RequestId = &v
	return s
}

type GetIpProtectionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpProtectionResponse) GoString() string {
	return s.String()
}

func (s *GetIpProtectionResponse) SetHeaders(v map[string]*string) *GetIpProtectionResponse {
	s.Headers = v
	return s
}

func (s *GetIpProtectionResponse) SetStatusCode(v int32) *GetIpProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpProtectionResponse) SetBody(v *GetIpProtectionResponseBody) *GetIpProtectionResponse {
	s.Body = v
	return s
}

type GetIpfilterListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetIpfilterListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListRequest) GoString() string {
	return s.String()
}

func (s *GetIpfilterListRequest) SetOwnerId(v int64) *GetIpfilterListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetIpfilterListRequest) SetResourceOwnerAccount(v string) *GetIpfilterListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetIpfilterListRequest) SetResourceOwnerId(v int64) *GetIpfilterListRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetIpfilterListResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 84DD77C7-A091-5139-9530-2D1F7CCE59E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *GetIpfilterListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetIpfilterListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBody) SetPageNumber(v int32) *GetIpfilterListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetPageSize(v int32) *GetIpfilterListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetRequestId(v string) *GetIpfilterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetTotalCount(v int32) *GetIpfilterListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetIpfilterListResponseBody) SetData(v *GetIpfilterListResponseBodyData) *GetIpfilterListResponseBody {
	s.Data = v
	return s
}

type GetIpfilterListResponseBodyData struct {
	Ipfilters []*GetIpfilterListResponseBodyDataIpfilters `json:"ipfilters,omitempty" xml:"ipfilters,omitempty" type:"Repeated"`
}

func (s GetIpfilterListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBodyData) SetIpfilters(v []*GetIpfilterListResponseBodyDataIpfilters) *GetIpfilterListResponseBodyData {
	s.Ipfilters = v
	return s
}

type GetIpfilterListResponseBodyDataIpfilters struct {
	// example:
	//
	// 1653547140
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 10083
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s GetIpfilterListResponseBodyDataIpfilters) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListResponseBodyDataIpfilters) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetCreateTime(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.CreateTime = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetId(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.Id = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetIpAddress(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.IpAddress = &v
	return s
}

type GetIpfilterListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpfilterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpfilterListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListResponse) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponse) SetHeaders(v map[string]*string) *GetIpfilterListResponse {
	s.Headers = v
	return s
}

func (s *GetIpfilterListResponse) SetStatusCode(v int32) *GetIpfilterListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpfilterListResponse) SetBody(v *GetIpfilterListResponseBody) *GetIpfilterListResponse {
	s.Body = v
	return s
}

type GetSuppressionListLevelRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetSuppressionListLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuppressionListLevelRequest) GoString() string {
	return s.String()
}

func (s *GetSuppressionListLevelRequest) SetOwnerId(v int64) *GetSuppressionListLevelRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSuppressionListLevelRequest) SetResourceOwnerAccount(v string) *GetSuppressionListLevelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSuppressionListLevelRequest) SetResourceOwnerId(v int64) *GetSuppressionListLevelRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetSuppressionListLevelResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuppressionListLevel *string `json:"SuppressionListLevel,omitempty" xml:"SuppressionListLevel,omitempty"`
}

func (s GetSuppressionListLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuppressionListLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuppressionListLevelResponseBody) SetRequestId(v string) *GetSuppressionListLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuppressionListLevelResponseBody) SetSuppressionListLevel(v string) *GetSuppressionListLevelResponseBody {
	s.SuppressionListLevel = &v
	return s
}

type GetSuppressionListLevelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuppressionListLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuppressionListLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuppressionListLevelResponse) GoString() string {
	return s.String()
}

func (s *GetSuppressionListLevelResponse) SetHeaders(v map[string]*string) *GetSuppressionListLevelResponse {
	s.Headers = v
	return s
}

func (s *GetSuppressionListLevelResponse) SetStatusCode(v int32) *GetSuppressionListLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuppressionListLevelResponse) SetBody(v *GetSuppressionListLevelResponseBody) *GetSuppressionListLevelResponse {
	s.Body = v
	return s
}

type GetTrackListRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagName   *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Total     *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTrackListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListRequest) GoString() string {
	return s.String()
}

func (s *GetTrackListRequest) SetAccountName(v string) *GetTrackListRequest {
	s.AccountName = &v
	return s
}

func (s *GetTrackListRequest) SetEndTime(v string) *GetTrackListRequest {
	s.EndTime = &v
	return s
}

func (s *GetTrackListRequest) SetOffset(v string) *GetTrackListRequest {
	s.Offset = &v
	return s
}

func (s *GetTrackListRequest) SetOffsetCreateTime(v string) *GetTrackListRequest {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListRequest) SetOffsetCreateTimeDesc(v string) *GetTrackListRequest {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListRequest) SetOwnerId(v int64) *GetTrackListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTrackListRequest) SetPageNumber(v string) *GetTrackListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrackListRequest) SetPageSize(v string) *GetTrackListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTrackListRequest) SetResourceOwnerAccount(v string) *GetTrackListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTrackListRequest) SetResourceOwnerId(v int64) *GetTrackListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTrackListRequest) SetStartTime(v string) *GetTrackListRequest {
	s.StartTime = &v
	return s
}

func (s *GetTrackListRequest) SetTagName(v string) *GetTrackListRequest {
	s.TagName = &v
	return s
}

func (s *GetTrackListRequest) SetTotal(v string) *GetTrackListRequest {
	s.Total = &v
	return s
}

type GetTrackListResponseBody struct {
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	Total *int32                        `json:"Total,omitempty" xml:"Total,omitempty"`
	Data  *GetTrackListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetTrackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBody) SetOffsetCreateTime(v string) *GetTrackListResponseBody {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListResponseBody) SetOffsetCreateTimeDesc(v string) *GetTrackListResponseBody {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListResponseBody) SetPageNo(v int32) *GetTrackListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetTrackListResponseBody) SetPageSize(v int32) *GetTrackListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTrackListResponseBody) SetRequestId(v string) *GetTrackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrackListResponseBody) SetTotal(v int32) *GetTrackListResponseBody {
	s.Total = &v
	return s
}

func (s *GetTrackListResponseBody) SetData(v *GetTrackListResponseBodyData) *GetTrackListResponseBody {
	s.Data = v
	return s
}

type GetTrackListResponseBodyData struct {
	Stat []*GetTrackListResponseBodyDataStat `json:"stat,omitempty" xml:"stat,omitempty" type:"Repeated"`
}

func (s GetTrackListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBodyData) SetStat(v []*GetTrackListResponseBodyDataStat) *GetTrackListResponseBodyData {
	s.Stat = v
	return s
}

type GetTrackListResponseBodyDataStat struct {
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	RcptClickCount *string `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	// example:
	//
	// 0
	RcptClickRate *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	// example:
	//
	// 0
	RcptOpenCount *string `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
	// example:
	//
	// 0
	RcptOpenRate *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	// example:
	//
	// 0
	RcptUniqueClickCount *string `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	// example:
	//
	// 0
	RcptUniqueClickRate *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	// example:
	//
	// 0
	RcptUniqueOpenCount *string `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	// example:
	//
	// 0
	RcptUniqueOpenRate *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	// example:
	//
	// 0
	TotalNumber *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetTrackListResponseBodyDataStat) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBodyDataStat) SetCreateTime(v string) *GetTrackListResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetTotalNumber(v string) *GetTrackListResponseBodyDataStat {
	s.TotalNumber = &v
	return s
}

type GetTrackListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrackListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListResponse) GoString() string {
	return s.String()
}

func (s *GetTrackListResponse) SetHeaders(v map[string]*string) *GetTrackListResponse {
	s.Headers = v
	return s
}

func (s *GetTrackListResponse) SetStatusCode(v int32) *GetTrackListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrackListResponse) SetBody(v *GetTrackListResponseBody) *GetTrackListResponse {
	s.Body = v
	return s
}

type GetTrackListByMailFromAndTagNameRequest struct {
	// example:
	//
	// e-service@amegroups.cn
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Subscription
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Total   *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameRequest) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetAccountName(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.AccountName = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetEndTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.EndTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffset(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Offset = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffsetCreateTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetPageNumber(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetPageSize(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.PageSize = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetResourceOwnerAccount(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetResourceOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetStartTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.StartTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetTagName(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.TagName = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetTotal(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Total = &v
	return s
}

type GetTrackListByMailFromAndTagNameResponseBody struct {
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	Total     *int32                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	TrackList *GetTrackListByMailFromAndTagNameResponseBodyTrackList `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Struct"`
}

func (s GetTrackListByMailFromAndTagNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetOffsetCreateTime(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetPageNo(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetPageSize(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetRequestId(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetTotal(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.Total = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetTrackList(v *GetTrackListByMailFromAndTagNameResponseBodyTrackList) *GetTrackListByMailFromAndTagNameResponseBody {
	s.TrackList = v
	return s
}

type GetTrackListByMailFromAndTagNameResponseBodyTrackList struct {
	Stat []*GetTrackListByMailFromAndTagNameResponseBodyTrackListStat `json:"Stat,omitempty" xml:"Stat,omitempty" type:"Repeated"`
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackList) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackList) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackList) SetStat(v []*GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) *GetTrackListByMailFromAndTagNameResponseBodyTrackList {
	s.Stat = v
	return s
}

type GetTrackListByMailFromAndTagNameResponseBodyTrackListStat struct {
	// example:
	//
	// 2022-01-11T10:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	RcptClickCount *string `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	// example:
	//
	// 0
	RcptClickRate *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	// example:
	//
	// 0
	RcptOpenCount *string `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
	// example:
	//
	// 0
	RcptOpenRate *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	// example:
	//
	// 0
	RcptUniqueClickCount *string `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	// example:
	//
	// 0
	RcptUniqueClickRate *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	// example:
	//
	// 0
	RcptUniqueOpenCount *string `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	// example:
	//
	// 0
	RcptUniqueOpenRate *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	// example:
	//
	// 0
	TotalNumber *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetCreateTime(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptClickCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptClickRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptOpenCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptOpenCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptOpenRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueClickCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueClickRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueOpenCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueOpenRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetTotalNumber(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.TotalNumber = &v
	return s
}

type GetTrackListByMailFromAndTagNameResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrackListByMailFromAndTagNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponse) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponse) SetHeaders(v map[string]*string) *GetTrackListByMailFromAndTagNameResponse {
	s.Headers = v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponse) SetStatusCode(v int32) *GetTrackListByMailFromAndTagNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponse) SetBody(v *GetTrackListByMailFromAndTagNameResponseBody) *GetTrackListByMailFromAndTagNameResponse {
	s.Body = v
	return s
}

type GetUserResponseBody struct {
	Data *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResponseBodyData struct {
	// example:
	//
	// true
	EnableEventbridge *bool `json:"EnableEventbridge,omitempty" xml:"EnableEventbridge,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetEnableEventbridge(v bool) *GetUserResponseBodyData {
	s.EnableEventbridge = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListUserSuppressionRequest struct {
	// example:
	//
	// test@example.net
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 1715669077
	EndBounceTime *int32 `json:"EndBounceTime,omitempty" xml:"EndBounceTime,omitempty"`
	// example:
	//
	// 1715669077
	EndCreateTime *int32 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1715668852
	StartBounceTime *int32 `json:"StartBounceTime,omitempty" xml:"StartBounceTime,omitempty"`
	// example:
	//
	// 1715668852
	StartCreateTime *int32 `json:"StartCreateTime,omitempty" xml:"StartCreateTime,omitempty"`
}

func (s ListUserSuppressionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserSuppressionRequest) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionRequest) SetAddress(v string) *ListUserSuppressionRequest {
	s.Address = &v
	return s
}

func (s *ListUserSuppressionRequest) SetEndBounceTime(v int32) *ListUserSuppressionRequest {
	s.EndBounceTime = &v
	return s
}

func (s *ListUserSuppressionRequest) SetEndCreateTime(v int32) *ListUserSuppressionRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListUserSuppressionRequest) SetOwnerId(v int64) *ListUserSuppressionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListUserSuppressionRequest) SetPageNo(v int32) *ListUserSuppressionRequest {
	s.PageNo = &v
	return s
}

func (s *ListUserSuppressionRequest) SetPageSize(v int32) *ListUserSuppressionRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserSuppressionRequest) SetResourceOwnerAccount(v string) *ListUserSuppressionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListUserSuppressionRequest) SetResourceOwnerId(v int64) *ListUserSuppressionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListUserSuppressionRequest) SetStartBounceTime(v int32) *ListUserSuppressionRequest {
	s.StartBounceTime = &v
	return s
}

func (s *ListUserSuppressionRequest) SetStartCreateTime(v int32) *ListUserSuppressionRequest {
	s.StartCreateTime = &v
	return s
}

type ListUserSuppressionResponseBody struct {
	Data *ListUserSuppressionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1A846D66-5EC7-551B-9687-5BF1963DCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserSuppressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserSuppressionResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponseBody) SetData(v *ListUserSuppressionResponseBodyData) *ListUserSuppressionResponseBody {
	s.Data = v
	return s
}

func (s *ListUserSuppressionResponseBody) SetPageNumber(v int32) *ListUserSuppressionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserSuppressionResponseBody) SetPageSize(v int32) *ListUserSuppressionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSuppressionResponseBody) SetRequestId(v string) *ListUserSuppressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSuppressionResponseBody) SetTotalCount(v int32) *ListUserSuppressionResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserSuppressionResponseBodyData struct {
	UserSuppressions []*ListUserSuppressionResponseBodyDataUserSuppressions `json:"UserSuppressions,omitempty" xml:"UserSuppressions,omitempty" type:"Repeated"`
}

func (s ListUserSuppressionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserSuppressionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponseBodyData) SetUserSuppressions(v []*ListUserSuppressionResponseBodyDataUserSuppressions) *ListUserSuppressionResponseBodyData {
	s.UserSuppressions = v
	return s
}

type ListUserSuppressionResponseBodyDataUserSuppressions struct {
	// example:
	//
	// test@example.net
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 1715667435
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1715667451
	LastBounceTime *int32 `json:"LastBounceTime,omitempty" xml:"LastBounceTime,omitempty"`
	// example:
	//
	// 59511
	SuppressionId *int32 `json:"SuppressionId,omitempty" xml:"SuppressionId,omitempty"`
	// example:
	//
	// - system
	//
	// - user
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUserSuppressionResponseBodyDataUserSuppressions) String() string {
	return tea.Prettify(s)
}

func (s ListUserSuppressionResponseBodyDataUserSuppressions) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetAddress(v string) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.Address = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetCreateTime(v int32) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.CreateTime = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetLastBounceTime(v int32) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.LastBounceTime = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetSuppressionId(v int32) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.SuppressionId = &v
	return s
}

func (s *ListUserSuppressionResponseBodyDataUserSuppressions) SetType(v string) *ListUserSuppressionResponseBodyDataUserSuppressions {
	s.Type = &v
	return s
}

type ListUserSuppressionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserSuppressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserSuppressionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserSuppressionResponse) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionResponse) SetHeaders(v map[string]*string) *ListUserSuppressionResponse {
	s.Headers = v
	return s
}

func (s *ListUserSuppressionResponse) SetStatusCode(v int32) *ListUserSuppressionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSuppressionResponse) SetBody(v *ListUserSuppressionResponseBody) *ListUserSuppressionResponse {
	s.Body = v
	return s
}

type ModifyMailAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1344565
	MailAddressId *int32 `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pwd129WDCft9
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// a***@example.net
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMailAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressRequest) SetMailAddressId(v int32) *ModifyMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *ModifyMailAddressRequest) SetOwnerId(v int64) *ModifyMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMailAddressRequest) SetPassword(v string) *ModifyMailAddressRequest {
	s.Password = &v
	return s
}

func (s *ModifyMailAddressRequest) SetReplyAddress(v string) *ModifyMailAddressRequest {
	s.ReplyAddress = &v
	return s
}

func (s *ModifyMailAddressRequest) SetResourceOwnerAccount(v string) *ModifyMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMailAddressRequest) SetResourceOwnerId(v int64) *ModifyMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyMailAddressResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMailAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressResponseBody) SetRequestId(v string) *ModifyMailAddressResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMailAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMailAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMailAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressResponse) SetHeaders(v map[string]*string) *ModifyMailAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyMailAddressResponse) SetStatusCode(v int32) *ModifyMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMailAddressResponse) SetBody(v *ModifyMailAddressResponseBody) *ModifyMailAddressResponse {
	s.Body = v
	return s
}

type ModifyPWByDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pwd
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPWByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPWByDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainRequest) SetDomainName(v string) *ModifyPWByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetOwnerId(v int64) *ModifyPWByDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetPassword(v string) *ModifyPWByDomainRequest {
	s.Password = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetResourceOwnerAccount(v string) *ModifyPWByDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetResourceOwnerId(v int64) *ModifyPWByDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyPWByDomainResponseBody struct {
	// example:
	//
	// OK
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 02B2A890-CBD8-4806-9BCA-C93190CE7EF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPWByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPWByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainResponseBody) SetCode(v string) *ModifyPWByDomainResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetMessage(v string) *ModifyPWByDomainResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetRequestId(v string) *ModifyPWByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetSuccess(v bool) *ModifyPWByDomainResponseBody {
	s.Success = &v
	return s
}

type ModifyPWByDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPWByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPWByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPWByDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainResponse) SetHeaders(v map[string]*string) *ModifyPWByDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyPWByDomainResponse) SetStatusCode(v int32) *ModifyPWByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPWByDomainResponse) SetBody(v *ModifyPWByDomainResponseBody) *ModifyPWByDomainResponse {
	s.Body = v
	return s
}

type ModifyTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagDescription       *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// The ID of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// vme
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ModifyTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTagRequest) GoString() string {
	return s.String()
}

func (s *ModifyTagRequest) SetOwnerId(v int64) *ModifyTagRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTagRequest) SetResourceOwnerAccount(v string) *ModifyTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTagRequest) SetResourceOwnerId(v int64) *ModifyTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTagRequest) SetTagDescription(v string) *ModifyTagRequest {
	s.TagDescription = &v
	return s
}

func (s *ModifyTagRequest) SetTagId(v int32) *ModifyTagRequest {
	s.TagId = &v
	return s
}

func (s *ModifyTagRequest) SetTagName(v string) *ModifyTagRequest {
	s.TagName = &v
	return s
}

type ModifyTagResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTagResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTagResponseBody) SetRequestId(v string) *ModifyTagResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTagResponse) GoString() string {
	return s.String()
}

func (s *ModifyTagResponse) SetHeaders(v map[string]*string) *ModifyTagResponse {
	s.Headers = v
	return s
}

func (s *ModifyTagResponse) SetStatusCode(v int32) *ModifyTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTagResponse) SetBody(v *ModifyTagResponseBody) *ModifyTagResponse {
	s.Body = v
	return s
}

type QueryDomainByParamRequest struct {
	// example:
	//
	// abc.com
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDomainByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamRequest) SetKeyWord(v string) *QueryDomainByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryDomainByParamRequest) SetOwnerId(v int64) *QueryDomainByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDomainByParamRequest) SetPageNo(v int32) *QueryDomainByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryDomainByParamRequest) SetPageSize(v int32) *QueryDomainByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDomainByParamRequest) SetResourceOwnerAccount(v string) *QueryDomainByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDomainByParamRequest) SetResourceOwnerId(v int64) *QueryDomainByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryDomainByParamRequest) SetStatus(v int32) *QueryDomainByParamRequest {
	s.Status = &v
	return s
}

type QueryDomainByParamResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 8C90CCD3-627C-4F87-AD8C-2F03146071EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryDomainByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryDomainByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBody) SetPageNumber(v int32) *QueryDomainByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetPageSize(v int32) *QueryDomainByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetRequestId(v string) *QueryDomainByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetTotalCount(v int32) *QueryDomainByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryDomainByParamResponseBody) SetData(v *QueryDomainByParamResponseBodyData) *QueryDomainByParamResponseBody {
	s.Data = v
	return s
}

type QueryDomainByParamResponseBodyData struct {
	Domain []*QueryDomainByParamResponseBodyDataDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryDomainByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBodyData) SetDomain(v []*QueryDomainByParamResponseBodyDataDomain) *QueryDomainByParamResponseBodyData {
	s.Domain = v
	return s
}

type QueryDomainByParamResponseBodyDataDomain struct {
	// example:
	//
	// 0
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	// example:
	//
	// 0
	ConfirmStatus *string `json:"ConfirmStatus,omitempty" xml:"ConfirmStatus,omitempty"`
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 158923
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// vmeixme.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 6bd86901b9fe4618a046
	DomainRecord *string `json:"DomainRecord,omitempty" xml:"DomainRecord,omitempty"`
	// example:
	//
	// 0
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// example:
	//
	// 1
	IcpStatus *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	// example:
	//
	// 0
	MxAuthStatus *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	// example:
	//
	// 0
	SpfAuthStatus *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryDomainByParamResponseBodyDataDomain) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetCnameAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.CnameAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetConfirmStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.ConfirmStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetCreateTime(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.CreateTime = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainId(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainId = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainName(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainRecord(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainRecord = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetIcpStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.IcpStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetMxAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.MxAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetSpfAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.SpfAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetUtcCreateTime(v int64) *QueryDomainByParamResponseBodyDataDomain {
	s.UtcCreateTime = &v
	return s
}

type QueryDomainByParamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponse) SetHeaders(v map[string]*string) *QueryDomainByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainByParamResponse) SetStatusCode(v int32) *QueryDomainByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainByParamResponse) SetBody(v *QueryDomainByParamResponseBody) *QueryDomainByParamResponse {
	s.Body = v
	return s
}

type QueryInvalidAddressRequest struct {
	// example:
	//
	// 2019-09-29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// info
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// example:
	//
	// 100
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// ***
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryInvalidAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressRequest) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressRequest) SetEndTime(v string) *QueryInvalidAddressRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetKeyWord(v string) *QueryInvalidAddressRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetLength(v int32) *QueryInvalidAddressRequest {
	s.Length = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetNextStart(v string) *QueryInvalidAddressRequest {
	s.NextStart = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetOwnerId(v int64) *QueryInvalidAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetResourceOwnerAccount(v string) *QueryInvalidAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetResourceOwnerId(v int64) *QueryInvalidAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetStartTime(v string) *QueryInvalidAddressRequest {
	s.StartTime = &v
	return s
}

type QueryInvalidAddressResponseBody struct {
	// example:
	//
	// 2
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// example:
	//
	// 95A7D497-F8DD-4834-B81E-C1783236E55F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryInvalidAddressResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryInvalidAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBody) SetNextStart(v string) *QueryInvalidAddressResponseBody {
	s.NextStart = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetRequestId(v string) *QueryInvalidAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetTotalCount(v int32) *QueryInvalidAddressResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetData(v *QueryInvalidAddressResponseBodyData) *QueryInvalidAddressResponseBody {
	s.Data = v
	return s
}

type QueryInvalidAddressResponseBodyData struct {
	MailDetail []*QueryInvalidAddressResponseBodyDataMailDetail `json:"mailDetail,omitempty" xml:"mailDetail,omitempty" type:"Repeated"`
}

func (s QueryInvalidAddressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBodyData) SetMailDetail(v []*QueryInvalidAddressResponseBodyDataMailDetail) *QueryInvalidAddressResponseBodyData {
	s.MailDetail = v
	return s
}

type QueryInvalidAddressResponseBodyDataMailDetail struct {
	// example:
	//
	// 2021-04-28T17:11Z
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// example:
	//
	// toaddress@test.com
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// example:
	//
	// 1619601108
	UtcLastUpdateTime *int64 `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
}

func (s QueryInvalidAddressResponseBodyDataMailDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressResponseBodyDataMailDetail) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) SetLastUpdateTime(v string) *QueryInvalidAddressResponseBodyDataMailDetail {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) SetToAddress(v string) *QueryInvalidAddressResponseBodyDataMailDetail {
	s.ToAddress = &v
	return s
}

func (s *QueryInvalidAddressResponseBodyDataMailDetail) SetUtcLastUpdateTime(v int64) *QueryInvalidAddressResponseBodyDataMailDetail {
	s.UtcLastUpdateTime = &v
	return s
}

type QueryInvalidAddressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInvalidAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressResponse) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponse) SetHeaders(v map[string]*string) *QueryInvalidAddressResponse {
	s.Headers = v
	return s
}

func (s *QueryInvalidAddressResponse) SetStatusCode(v int32) *QueryInvalidAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInvalidAddressResponse) SetBody(v *QueryInvalidAddressResponseBody) *QueryInvalidAddressResponse {
	s.Body = v
	return s
}

type QueryMailAddressByParamRequest struct {
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// batch
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s QueryMailAddressByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMailAddressByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamRequest) SetKeyWord(v string) *QueryMailAddressByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetOwnerId(v int64) *QueryMailAddressByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetPageNo(v int32) *QueryMailAddressByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetPageSize(v int32) *QueryMailAddressByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetResourceOwnerAccount(v string) *QueryMailAddressByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetResourceOwnerId(v int64) *QueryMailAddressByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMailAddressByParamRequest) SetSendtype(v string) *QueryMailAddressByParamRequest {
	s.Sendtype = &v
	return s
}

type QueryMailAddressByParamResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 95A7D497-F8DD-4834-B81E-C1783236E55F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryMailAddressByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryMailAddressByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMailAddressByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponseBody) SetPageNumber(v int32) *QueryMailAddressByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetPageSize(v int32) *QueryMailAddressByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetRequestId(v string) *QueryMailAddressByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetTotalCount(v int32) *QueryMailAddressByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBody) SetData(v *QueryMailAddressByParamResponseBodyData) *QueryMailAddressByParamResponseBody {
	s.Data = v
	return s
}

type QueryMailAddressByParamResponseBodyData struct {
	MailAddress []*QueryMailAddressByParamResponseBodyDataMailAddress `json:"mailAddress,omitempty" xml:"mailAddress,omitempty" type:"Repeated"`
}

func (s QueryMailAddressByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMailAddressByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponseBodyData) SetMailAddress(v []*QueryMailAddressByParamResponseBodyDataMailAddress) *QueryMailAddressByParamResponseBodyData {
	s.MailAddress = v
	return s
}

type QueryMailAddressByParamResponseBodyDataMailAddress struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 0
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 10000
	DailyCount *string `json:"DailyCount,omitempty" xml:"DailyCount,omitempty"`
	// example:
	//
	// 100
	DailyReqCount *string `json:"DailyReqCount,omitempty" xml:"DailyReqCount,omitempty"`
	// example:
	//
	// 0
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// example:
	//
	// 12122
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	// example:
	//
	// 300000
	MonthCount *string `json:"MonthCount,omitempty" xml:"MonthCount,omitempty"`
	// example:
	//
	// 20000
	MonthReqCount *string `json:"MonthReqCount,omitempty" xml:"MonthReqCount,omitempty"`
	// example:
	//
	// 112@aliyun.com
	ReplyAddress *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	// example:
	//
	// 0
	ReplyStatus *string `json:"ReplyStatus,omitempty" xml:"ReplyStatus,omitempty"`
	// example:
	//
	// batch
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s QueryMailAddressByParamResponseBodyDataMailAddress) String() string {
	return tea.Prettify(s)
}

func (s QueryMailAddressByParamResponseBodyDataMailAddress) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetAccountName(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.AccountName = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetAccountStatus(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.AccountStatus = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetCreateTime(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.CreateTime = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetDailyCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.DailyCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetDailyReqCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.DailyReqCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetDomainStatus(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.DomainStatus = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetMailAddressId(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.MailAddressId = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetMonthCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.MonthCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetMonthReqCount(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.MonthReqCount = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetReplyAddress(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.ReplyAddress = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetReplyStatus(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.ReplyStatus = &v
	return s
}

func (s *QueryMailAddressByParamResponseBodyDataMailAddress) SetSendtype(v string) *QueryMailAddressByParamResponseBodyDataMailAddress {
	s.Sendtype = &v
	return s
}

type QueryMailAddressByParamResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMailAddressByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMailAddressByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMailAddressByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponse) SetHeaders(v map[string]*string) *QueryMailAddressByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryMailAddressByParamResponse) SetStatusCode(v int32) *QueryMailAddressByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMailAddressByParamResponse) SetBody(v *QueryMailAddressByParamResponseBody) *QueryMailAddressByParamResponse {
	s.Body = v
	return s
}

type QueryReceiverByParamRequest struct {
	// example:
	//
	// mesh-notification
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryReceiverByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamRequest) SetKeyWord(v string) *QueryReceiverByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetOwnerId(v int64) *QueryReceiverByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetPageNo(v int32) *QueryReceiverByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetPageSize(v int32) *QueryReceiverByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetResourceOwnerAccount(v string) *QueryReceiverByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetResourceOwnerId(v int64) *QueryReceiverByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetStatus(v int32) *QueryReceiverByParamRequest {
	s.Status = &v
	return s
}

type QueryReceiverByParamResponseBody struct {
	// example:
	//
	// 6aec200853#102#1638894326#rdwilbur@verizon.net
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryReceiverByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryReceiverByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBody) SetNextStart(v string) *QueryReceiverByParamResponseBody {
	s.NextStart = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetPageSize(v int32) *QueryReceiverByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetRequestId(v string) *QueryReceiverByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetTotalCount(v int32) *QueryReceiverByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetData(v *QueryReceiverByParamResponseBodyData) *QueryReceiverByParamResponseBody {
	s.Data = v
	return s
}

type QueryReceiverByParamResponseBodyData struct {
	Receiver []*QueryReceiverByParamResponseBodyDataReceiver `json:"receiver,omitempty" xml:"receiver,omitempty" type:"Repeated"`
}

func (s QueryReceiverByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBodyData) SetReceiver(v []*QueryReceiverByParamResponseBodyDataReceiver) *QueryReceiverByParamResponseBodyData {
	s.Receiver = v
	return s
}

type QueryReceiverByParamResponseBodyDataReceiver struct {
	// example:
	//
	// 3
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// 0c910a7143044b1e116719eb678907b3
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// example:
	//
	// 10***@163.com
	ReceiversAlias *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	// example:
	//
	// TKP000442-333
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	// example:
	//
	// 0
	ReceiversStatus *string `json:"ReceiversStatus,omitempty" xml:"ReceiversStatus,omitempty"`
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryReceiverByParamResponseBodyDataReceiver) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamResponseBodyDataReceiver) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetCount(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.Count = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetCreateTime(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetDesc(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.Desc = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiverId(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiverId = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversAlias(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversAlias = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversName(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversName = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversStatus(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversStatus = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetUtcCreateTime(v int64) *QueryReceiverByParamResponseBodyDataReceiver {
	s.UtcCreateTime = &v
	return s
}

type QueryReceiverByParamResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReceiverByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReceiverByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponse) SetHeaders(v map[string]*string) *QueryReceiverByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiverByParamResponse) SetStatusCode(v int32) *QueryReceiverByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReceiverByParamResponse) SetBody(v *QueryReceiverByParamResponseBody) *QueryReceiverByParamResponse {
	s.Body = v
	return s
}

type QueryReceiverDetailRequest struct {
	// example:
	//
	// b***@example.net
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// example:
	//
	// 0
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1235
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryReceiverDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailRequest) SetKeyWord(v string) *QueryReceiverDetailRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetNextStart(v string) *QueryReceiverDetailRequest {
	s.NextStart = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetOwnerId(v int64) *QueryReceiverDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetPageSize(v int32) *QueryReceiverDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetReceiverId(v string) *QueryReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetResourceOwnerAccount(v string) *QueryReceiverDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetResourceOwnerId(v int64) *QueryReceiverDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryReceiverDetailResponseBody struct {
	// example:
	//
	// UserName,NickName,Gender,Birthday,Mobile
	DataSchema *string `json:"DataSchema,omitempty" xml:"DataSchema,omitempty"`
	// example:
	//
	// 90f0243616#401#b68c2af70b9664b2843f90fd3630b1a3-1650779410#xxx@demo.aliyun.com
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 361
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryReceiverDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryReceiverDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBody) SetDataSchema(v string) *QueryReceiverDetailResponseBody {
	s.DataSchema = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetNextStart(v string) *QueryReceiverDetailResponseBody {
	s.NextStart = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetRequestId(v string) *QueryReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetTotalCount(v int32) *QueryReceiverDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetData(v *QueryReceiverDetailResponseBodyData) *QueryReceiverDetailResponseBody {
	s.Data = v
	return s
}

type QueryReceiverDetailResponseBodyData struct {
	Detail []*QueryReceiverDetailResponseBodyDataDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Repeated"`
}

func (s QueryReceiverDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBodyData) SetDetail(v []*QueryReceiverDetailResponseBodyDataDetail) *QueryReceiverDetailResponseBodyData {
	s.Detail = v
	return s
}

type QueryReceiverDetailResponseBodyDataDetail struct {
	// example:
	//
	// 2019-09-29T13:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {\\"Domains\\": [\\"a.example.net\\", \\"b.example.net\\", \\"c.example.net\\", \\"d.example.net\\"]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// a***@example.net
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryReceiverDetailResponseBodyDataDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetCreateTime(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetData(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.Data = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetEmail(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.Email = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetUtcCreateTime(v int64) *QueryReceiverDetailResponseBodyDataDetail {
	s.UtcCreateTime = &v
	return s
}

type QueryReceiverDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReceiverDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponse) SetHeaders(v map[string]*string) *QueryReceiverDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiverDetailResponse) SetStatusCode(v int32) *QueryReceiverDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReceiverDetailResponse) SetBody(v *QueryReceiverDetailResponseBody) *QueryReceiverDetailResponse {
	s.Body = v
	return s
}

type QueryTagByParamRequest struct {
	// example:
	//
	// 1aTag
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTagByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTagByParamRequest) SetKeyWord(v string) *QueryTagByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTagByParamRequest) SetOwnerId(v int64) *QueryTagByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagByParamRequest) SetPageNo(v int32) *QueryTagByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTagByParamRequest) SetPageSize(v int32) *QueryTagByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTagByParamRequest) SetResourceOwnerAccount(v string) *QueryTagByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagByParamRequest) SetResourceOwnerId(v int64) *QueryTagByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryTagByParamResponseBody struct {
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryTagByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryTagByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBody) SetPageNumber(v int32) *QueryTagByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetPageSize(v int32) *QueryTagByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetRequestId(v string) *QueryTagByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetTotalCount(v int32) *QueryTagByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetData(v *QueryTagByParamResponseBodyData) *QueryTagByParamResponseBody {
	s.Data = v
	return s
}

type QueryTagByParamResponseBodyData struct {
	Tag []*QueryTagByParamResponseBodyDataTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s QueryTagByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBodyData) SetTag(v []*QueryTagByParamResponseBodyDataTag) *QueryTagByParamResponseBodyData {
	s.Tag = v
	return s
}

type QueryTagByParamResponseBodyDataTag struct {
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// example:
	//
	// 52366
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// hellopal
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryTagByParamResponseBodyDataTag) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamResponseBodyDataTag) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagDescription(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagDescription = &v
	return s
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagId(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagId = &v
	return s
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagName(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagName = &v
	return s
}

type QueryTagByParamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponse) SetHeaders(v map[string]*string) *QueryTagByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryTagByParamResponse) SetStatusCode(v int32) *QueryTagByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagByParamResponse) SetBody(v *QueryTagByParamResponseBody) *QueryTagByParamResponse {
	s.Body = v
	return s
}

type QueryTaskByParamRequest struct {
	// example:
	//
	// mesh-notification-788717
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTaskByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamRequest) SetKeyWord(v string) *QueryTaskByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTaskByParamRequest) SetOwnerId(v int64) *QueryTaskByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTaskByParamRequest) SetPageNo(v int32) *QueryTaskByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTaskByParamRequest) SetPageSize(v int32) *QueryTaskByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskByParamRequest) SetResourceOwnerAccount(v string) *QueryTaskByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTaskByParamRequest) SetResourceOwnerId(v int64) *QueryTaskByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTaskByParamRequest) SetStatus(v int32) *QueryTaskByParamRequest {
	s.Status = &v
	return s
}

type QueryTaskByParamResponseBody struct {
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryTaskByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryTaskByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBody) SetPageNumber(v int32) *QueryTaskByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetPageSize(v int32) *QueryTaskByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetRequestId(v string) *QueryTaskByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetTotalCount(v int32) *QueryTaskByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryTaskByParamResponseBody) SetData(v *QueryTaskByParamResponseBodyData) *QueryTaskByParamResponseBody {
	s.Data = v
	return s
}

type QueryTaskByParamResponseBodyData struct {
	Task []*QueryTaskByParamResponseBodyDataTask `json:"task,omitempty" xml:"task,omitempty" type:"Repeated"`
}

func (s QueryTaskByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBodyData) SetTask(v []*QueryTaskByParamResponseBodyDataTask) *QueryTaskByParamResponseBodyData {
	s.Task = v
	return s
}

type QueryTaskByParamResponseBodyDataTask struct {
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// TKP000442-333
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	// example:
	//
	// 1
	RequestCount *string `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// example:
	//
	// 202201
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// 1054296
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// Short Simple
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 1569734892
	UtcCreateTime *int64 `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryTaskByParamResponseBodyDataTask) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamResponseBodyDataTask) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBodyDataTask) SetAddressType(v string) *QueryTaskByParamResponseBodyDataTask {
	s.AddressType = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetCreateTime(v string) *QueryTaskByParamResponseBodyDataTask {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetReceiversName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.ReceiversName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetRequestCount(v string) *QueryTaskByParamResponseBodyDataTask {
	s.RequestCount = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTagName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TagName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTaskId(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TaskId = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTaskStatus(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTemplateName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TemplateName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetUtcCreateTime(v int64) *QueryTaskByParamResponseBodyDataTask {
	s.UtcCreateTime = &v
	return s
}

type QueryTaskByParamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponse) SetHeaders(v map[string]*string) *QueryTaskByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskByParamResponse) SetStatusCode(v int32) *QueryTaskByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskByParamResponse) SetBody(v *QueryTaskByParamResponseBody) *QueryTaskByParamResponse {
	s.Body = v
	return s
}

type RemoveUserSuppressionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SuppressionIds       *string `json:"SuppressionIds,omitempty" xml:"SuppressionIds,omitempty"`
}

func (s RemoveUserSuppressionRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserSuppressionRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserSuppressionRequest) SetOwnerId(v int64) *RemoveUserSuppressionRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveUserSuppressionRequest) SetResourceOwnerAccount(v string) *RemoveUserSuppressionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveUserSuppressionRequest) SetResourceOwnerId(v int64) *RemoveUserSuppressionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveUserSuppressionRequest) SetSuppressionIds(v string) *RemoveUserSuppressionRequest {
	s.SuppressionIds = &v
	return s
}

type RemoveUserSuppressionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserSuppressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserSuppressionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserSuppressionResponseBody) SetRequestId(v string) *RemoveUserSuppressionResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserSuppressionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserSuppressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserSuppressionResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserSuppressionResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserSuppressionResponse) SetHeaders(v map[string]*string) *RemoveUserSuppressionResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserSuppressionResponse) SetStatusCode(v int32) *RemoveUserSuppressionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserSuppressionResponse) SetBody(v *RemoveUserSuppressionResponseBody) *RemoveUserSuppressionResponse {
	s.Body = v
	return s
}

type SaveReceiverDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{“b”:” birthday”,”e”:”xxx@alibaba-inc.com”,”g”:” gender”,”m”:” mobile “,”n”:” nickname “,”u”:”name “}]
	Detail  *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 34642
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SaveReceiverDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailRequest) SetDetail(v string) *SaveReceiverDetailRequest {
	s.Detail = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetOwnerId(v int64) *SaveReceiverDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetReceiverId(v string) *SaveReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetResourceOwnerAccount(v string) *SaveReceiverDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetResourceOwnerId(v int64) *SaveReceiverDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type SaveReceiverDetailResponseBody struct {
	Data *SaveReceiverDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 638
	ErrorCount *int32 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 274
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s SaveReceiverDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBody) SetData(v *SaveReceiverDetailResponseBodyData) *SaveReceiverDetailResponseBody {
	s.Data = v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetErrorCount(v int32) *SaveReceiverDetailResponseBody {
	s.ErrorCount = &v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetRequestId(v string) *SaveReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetSuccessCount(v int32) *SaveReceiverDetailResponseBody {
	s.SuccessCount = &v
	return s
}

type SaveReceiverDetailResponseBodyData struct {
	Detail []*SaveReceiverDetailResponseBodyDataDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s SaveReceiverDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBodyData) SetDetail(v []*SaveReceiverDetailResponseBodyDataDetail) *SaveReceiverDetailResponseBodyData {
	s.Detail = v
	return s
}

type SaveReceiverDetailResponseBodyDataDetail struct {
	// example:
	//
	// test@mail.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s SaveReceiverDetailResponseBodyDataDetail) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBodyDataDetail) SetEmail(v string) *SaveReceiverDetailResponseBodyDataDetail {
	s.Email = &v
	return s
}

type SaveReceiverDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveReceiverDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailResponse) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponse) SetHeaders(v map[string]*string) *SaveReceiverDetailResponse {
	s.Headers = v
	return s
}

func (s *SaveReceiverDetailResponse) SetStatusCode(v int32) *SaveReceiverDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveReceiverDetailResponse) SetBody(v *SaveReceiverDetailResponseBody) *SaveReceiverDetailResponse {
	s.Body = v
	return s
}

type SendTestByTemplateRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 2000/01/01
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// This parameter is required.
	Email  *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 1380000****
	Mobile               *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	NickName             *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TemplateId *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SendTestByTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTestByTemplateRequest) GoString() string {
	return s.String()
}

func (s *SendTestByTemplateRequest) SetAccountName(v string) *SendTestByTemplateRequest {
	s.AccountName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetBirthday(v string) *SendTestByTemplateRequest {
	s.Birthday = &v
	return s
}

func (s *SendTestByTemplateRequest) SetEmail(v string) *SendTestByTemplateRequest {
	s.Email = &v
	return s
}

func (s *SendTestByTemplateRequest) SetGender(v string) *SendTestByTemplateRequest {
	s.Gender = &v
	return s
}

func (s *SendTestByTemplateRequest) SetMobile(v string) *SendTestByTemplateRequest {
	s.Mobile = &v
	return s
}

func (s *SendTestByTemplateRequest) SetNickName(v string) *SendTestByTemplateRequest {
	s.NickName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetOwnerId(v int64) *SendTestByTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *SendTestByTemplateRequest) SetResourceOwnerAccount(v string) *SendTestByTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendTestByTemplateRequest) SetResourceOwnerId(v int64) *SendTestByTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendTestByTemplateRequest) SetTemplateId(v int32) *SendTestByTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *SendTestByTemplateRequest) SetUserName(v string) *SendTestByTemplateRequest {
	s.UserName = &v
	return s
}

type SendTestByTemplateResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTestByTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTestByTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SendTestByTemplateResponseBody) SetRequestId(v string) *SendTestByTemplateResponseBody {
	s.RequestId = &v
	return s
}

type SendTestByTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTestByTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTestByTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTestByTemplateResponse) GoString() string {
	return s.String()
}

func (s *SendTestByTemplateResponse) SetHeaders(v map[string]*string) *SendTestByTemplateResponse {
	s.Headers = v
	return s
}

func (s *SendTestByTemplateResponse) SetStatusCode(v int32) *SendTestByTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTestByTemplateResponse) SetBody(v *SendTestByTemplateResponseBody) *SendTestByTemplateResponse {
	s.Body = v
	return s
}

type SenderStatisticsByTagNameAndBatchIDRequest struct {
	// example:
	//
	// xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xxx
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDRequest) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDRequest) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetAccountName(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetEndTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.EndTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetOwnerId(v int64) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.OwnerId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetResourceOwnerAccount(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetResourceOwnerId(v int64) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetStartTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.StartTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetTagName(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.TagName = &v
	return s
}

type SenderStatisticsByTagNameAndBatchIDResponseBody struct {
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *SenderStatisticsByTagNameAndBatchIDResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetRequestId(v string) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetTotalCount(v int32) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetData(v *SenderStatisticsByTagNameAndBatchIDResponseBodyData) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.Data = v
	return s
}

type SenderStatisticsByTagNameAndBatchIDResponseBodyData struct {
	Stat []*SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat `json:"stat,omitempty" xml:"stat,omitempty" type:"Repeated"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyData) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyData) SetStat(v []*SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) *SenderStatisticsByTagNameAndBatchIDResponseBodyData {
	s.Stat = v
	return s
}

type SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat struct {
	// example:
	//
	// 2021-07-02
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	FaildCount *string `json:"faildCount,omitempty" xml:"faildCount,omitempty"`
	// example:
	//
	// 4
	RequestCount *string `json:"requestCount,omitempty" xml:"requestCount,omitempty"`
	// example:
	//
	// 100.00%
	SucceededPercent *string `json:"succeededPercent,omitempty" xml:"succeededPercent,omitempty"`
	// example:
	//
	// 4
	SuccessCount *string `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// example:
	//
	// 0
	UnavailableCount *string `json:"unavailableCount,omitempty" xml:"unavailableCount,omitempty"`
	// example:
	//
	// 0%
	UnavailablePercent *string `json:"unavailablePercent,omitempty" xml:"unavailablePercent,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetCreateTime(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetFaildCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.FaildCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetRequestCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.RequestCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetSucceededPercent(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.SucceededPercent = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetSuccessCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.SuccessCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetUnavailableCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.UnavailableCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetUnavailablePercent(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.UnavailablePercent = &v
	return s
}

type SenderStatisticsByTagNameAndBatchIDResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SenderStatisticsByTagNameAndBatchIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponse) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponse) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetHeaders(v map[string]*string) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.Headers = v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetStatusCode(v int32) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.StatusCode = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetBody(v *SenderStatisticsByTagNameAndBatchIDResponseBody) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.Body = v
	return s
}

type SenderStatisticsDetailByParamRequest struct {
	// example:
	//
	// s***@example.net
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 2021-04-29 00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 5
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// 90f0243616#203#a***@example.net-1658817837#a***@example.net.247475288187
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2021-04-28 00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// EmailQuestionnaireHelioscam
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// b***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s SenderStatisticsDetailByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamRequest) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamRequest) SetAccountName(v string) *SenderStatisticsDetailByParamRequest {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetEndTime(v string) *SenderStatisticsDetailByParamRequest {
	s.EndTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetLength(v int32) *SenderStatisticsDetailByParamRequest {
	s.Length = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetNextStart(v string) *SenderStatisticsDetailByParamRequest {
	s.NextStart = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetOwnerId(v int64) *SenderStatisticsDetailByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetResourceOwnerAccount(v string) *SenderStatisticsDetailByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetResourceOwnerId(v int64) *SenderStatisticsDetailByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetStartTime(v string) *SenderStatisticsDetailByParamRequest {
	s.StartTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetStatus(v int32) *SenderStatisticsDetailByParamRequest {
	s.Status = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetTagName(v string) *SenderStatisticsDetailByParamRequest {
	s.TagName = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetToAddress(v string) *SenderStatisticsDetailByParamRequest {
	s.ToAddress = &v
	return s
}

type SenderStatisticsDetailByParamResponseBody struct {
	// example:
	//
	// 90f0243616#203#a***@example.net-1658817689#a***@example.net.247141122178
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	// example:
	//
	// B5AB8EBB-EE64-4BB2-B085-B92CC5DEDC41
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SenderStatisticsDetailByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s SenderStatisticsDetailByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBody) SetNextStart(v string) *SenderStatisticsDetailByParamResponseBody {
	s.NextStart = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) SetRequestId(v string) *SenderStatisticsDetailByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) SetData(v *SenderStatisticsDetailByParamResponseBodyData) *SenderStatisticsDetailByParamResponseBody {
	s.Data = v
	return s
}

type SenderStatisticsDetailByParamResponseBodyData struct {
	MailDetail []*SenderStatisticsDetailByParamResponseBodyDataMailDetail `json:"mailDetail,omitempty" xml:"mailDetail,omitempty" type:"Repeated"`
}

func (s SenderStatisticsDetailByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBodyData) SetMailDetail(v []*SenderStatisticsDetailByParamResponseBodyDataMailDetail) *SenderStatisticsDetailByParamResponseBodyData {
	s.MailDetail = v
	return s
}

type SenderStatisticsDetailByParamResponseBodyDataMailDetail struct {
	// example:
	//
	// s***@example.net
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ErrorClassification *string `json:"ErrorClassification,omitempty" xml:"ErrorClassification,omitempty"`
	// example:
	//
	// 2021-04-28T17:11Z
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// example:
	//
	// 250 Send Mail OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// b***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	// example:
	//
	// 1619601108
	UtcLastUpdateTime *string `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
}

func (s SenderStatisticsDetailByParamResponseBodyDataMailDetail) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBodyDataMailDetail) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetAccountName(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetErrorClassification(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ErrorClassification = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetLastUpdateTime(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.LastUpdateTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetMessage(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Message = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetStatus(v int32) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Status = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetSubject(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Subject = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetToAddress(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ToAddress = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetUtcLastUpdateTime(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.UtcLastUpdateTime = &v
	return s
}

type SenderStatisticsDetailByParamResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SenderStatisticsDetailByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SenderStatisticsDetailByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponse) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponse) SetHeaders(v map[string]*string) *SenderStatisticsDetailByParamResponse {
	s.Headers = v
	return s
}

func (s *SenderStatisticsDetailByParamResponse) SetStatusCode(v int32) *SenderStatisticsDetailByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponse) SetBody(v *SenderStatisticsDetailByParamResponseBody) *SenderStatisticsDetailByParamResponse {
	s.Body = v
	return s
}

type SetSuppressionListLevelRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SuppressionListLevel *string `json:"SuppressionListLevel,omitempty" xml:"SuppressionListLevel,omitempty"`
}

func (s SetSuppressionListLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSuppressionListLevelRequest) GoString() string {
	return s.String()
}

func (s *SetSuppressionListLevelRequest) SetOwnerId(v int64) *SetSuppressionListLevelRequest {
	s.OwnerId = &v
	return s
}

func (s *SetSuppressionListLevelRequest) SetResourceOwnerAccount(v string) *SetSuppressionListLevelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetSuppressionListLevelRequest) SetResourceOwnerId(v int64) *SetSuppressionListLevelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetSuppressionListLevelRequest) SetSuppressionListLevel(v string) *SetSuppressionListLevelRequest {
	s.SuppressionListLevel = &v
	return s
}

type SetSuppressionListLevelResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuppressionListLevel *string `json:"SuppressionListLevel,omitempty" xml:"SuppressionListLevel,omitempty"`
}

func (s SetSuppressionListLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSuppressionListLevelResponseBody) GoString() string {
	return s.String()
}

func (s *SetSuppressionListLevelResponseBody) SetRequestId(v string) *SetSuppressionListLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSuppressionListLevelResponseBody) SetSuppressionListLevel(v string) *SetSuppressionListLevelResponseBody {
	s.SuppressionListLevel = &v
	return s
}

type SetSuppressionListLevelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSuppressionListLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSuppressionListLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSuppressionListLevelResponse) GoString() string {
	return s.String()
}

func (s *SetSuppressionListLevelResponse) SetHeaders(v map[string]*string) *SetSuppressionListLevelResponse {
	s.Headers = v
	return s
}

func (s *SetSuppressionListLevelResponse) SetStatusCode(v int32) *SetSuppressionListLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSuppressionListLevelResponse) SetBody(v *SetSuppressionListLevelResponseBody) *SetSuppressionListLevelResponse {
	s.Body = v
	return s
}

type SingleSendMailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressType *int32 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 0
	ClickTrace *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	FromAlias  *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	// example:
	//
	// body
	HtmlBody *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 111@aliyun.com
	ReplyAddress      *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// 12121
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1@example.com
	ToAddress              *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UnSubscribeFilterLevel *string `json:"UnSubscribeFilterLevel,omitempty" xml:"UnSubscribeFilterLevel,omitempty"`
	UnSubscribeLinkType    *string `json:"UnSubscribeLinkType,omitempty" xml:"UnSubscribeLinkType,omitempty"`
}

func (s SingleSendMailRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequest) SetAccountName(v string) *SingleSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailRequest) SetAddressType(v int32) *SingleSendMailRequest {
	s.AddressType = &v
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

func (s *SingleSendMailRequest) SetHtmlBody(v string) *SingleSendMailRequest {
	s.HtmlBody = &v
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

type SingleSendMailResponseBody struct {
	// example:
	//
	// xxxxxx
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// example:
	//
	// 2D086F6-8F31-4658-84C1-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleSendMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailResponseBody) GoString() string {
	return s.String()
}

func (s *SingleSendMailResponseBody) SetEnvId(v string) *SingleSendMailResponseBody {
	s.EnvId = &v
	return s
}

func (s *SingleSendMailResponseBody) SetRequestId(v string) *SingleSendMailResponseBody {
	s.RequestId = &v
	return s
}

type SingleSendMailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleSendMailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleSendMailResponse) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailResponse) GoString() string {
	return s.String()
}

func (s *SingleSendMailResponse) SetHeaders(v map[string]*string) *SingleSendMailResponse {
	s.Headers = v
	return s
}

func (s *SingleSendMailResponse) SetStatusCode(v int32) *SingleSendMailResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleSendMailResponse) SetBody(v *SingleSendMailResponseBody) *SingleSendMailResponse {
	s.Body = v
	return s
}

type UpdateIpProtectionRequest struct {
	// example:
	//
	// 0
	IpProtection         *string `json:"IpProtection,omitempty" xml:"IpProtection,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpProtectionRequest) SetIpProtection(v string) *UpdateIpProtectionRequest {
	s.IpProtection = &v
	return s
}

func (s *UpdateIpProtectionRequest) SetOwnerId(v int64) *UpdateIpProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpProtectionRequest) SetResourceOwnerAccount(v string) *UpdateIpProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpProtectionRequest) SetResourceOwnerId(v int64) *UpdateIpProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateIpProtectionResponseBody struct {
	// example:
	//
	// B653A6FC-D1AD-5936-A262-F50994ED2574
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpProtectionResponseBody) SetRequestId(v string) *UpdateIpProtectionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpProtectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpProtectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpProtectionResponse) SetHeaders(v map[string]*string) *UpdateIpProtectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpProtectionResponse) SetStatusCode(v int32) *UpdateIpProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpProtectionResponse) SetBody(v *UpdateIpProtectionResponseBody) *UpdateIpProtectionResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	User *UpdateUserRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetUser(v *UpdateUserRequestUser) *UpdateUserRequest {
	s.User = v
	return s
}

type UpdateUserRequestUser struct {
	// example:
	//
	// true
	EnableEventbridge *bool `json:"EnableEventbridge,omitempty" xml:"EnableEventbridge,omitempty"`
}

func (s UpdateUserRequestUser) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequestUser) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestUser) SetEnableEventbridge(v bool) *UpdateUserRequestUser {
	s.EnableEventbridge = &v
	return s
}

type UpdateUserShrinkRequest struct {
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s UpdateUserShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserShrinkRequest) SetUserShrink(v string) *UpdateUserShrinkRequest {
	s.UserShrink = &v
	return s
}

type UpdateUserResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7BC346F6-1092-5852-B6E2-CCE2E5AAE51F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加IP防护信息
//
// @param request - AddIpfilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpfilterResponse
func (client *Client) AddIpfilterWithOptions(request *AddIpfilterRequest, runtime *util.RuntimeOptions) (_result *AddIpfilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpAddress)) {
		query["IpAddress"] = request.IpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIpfilter"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpfilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加IP防护信息
//
// @param request - AddIpfilterRequest
//
// @return AddIpfilterResponse
func (client *Client) AddIpfilter(request *AddIpfilterRequest) (_result *AddIpfilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpfilterResponse{}
	_body, _err := client.AddIpfilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证回信地址
//
// @param request - ApproveReplyMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveReplyMailAddressResponse
func (client *Client) ApproveReplyMailAddressWithOptions(request *ApproveReplyMailAddressRequest, runtime *util.RuntimeOptions) (_result *ApproveReplyMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		query["Ticket"] = request.Ticket
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveReplyMailAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveReplyMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证回信地址
//
// @param request - ApproveReplyMailAddressRequest
//
// @return ApproveReplyMailAddressResponse
func (client *Client) ApproveReplyMailAddress(request *ApproveReplyMailAddressRequest) (_result *ApproveReplyMailAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveReplyMailAddressResponse{}
	_body, _err := client.ApproveReplyMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量发送邮件
//
// @param request - BatchSendMailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendMailResponse
func (client *Client) BatchSendMailWithOptions(request *BatchSendMailRequest, runtime *util.RuntimeOptions) (_result *BatchSendMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ClickTrace)) {
		query["ClickTrace"] = request.ClickTrace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiversName)) {
		query["ReceiversName"] = request.ReceiversName
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyAddress)) {
		query["ReplyAddress"] = request.ReplyAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyAddressAlias)) {
		query["ReplyAddressAlias"] = request.ReplyAddressAlias
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.UnSubscribeFilterLevel)) {
		query["UnSubscribeFilterLevel"] = request.UnSubscribeFilterLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UnSubscribeLinkType)) {
		query["UnSubscribeLinkType"] = request.UnSubscribeLinkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSendMail"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSendMailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量发送邮件
//
// @param request - BatchSendMailRequest
//
// @return BatchSendMailResponse
func (client *Client) BatchSendMail(request *BatchSendMailRequest) (_result *BatchSendMailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSendMailResponse{}
	_body, _err := client.BatchSendMailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验域名状态
//
// @param request - CheckDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainResponse
func (client *Client) CheckDomainWithOptions(request *CheckDomainRequest, runtime *util.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDomain"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验域名状态
//
// @param request - CheckDomainRequest
//
// @return CheckDomainResponse
func (client *Client) CheckDomain(request *CheckDomainRequest) (_result *CheckDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDomainResponse{}
	_body, _err := client.CheckDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验回信地址
//
// @param request - CheckReplyToMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckReplyToMailAddressResponse
func (client *Client) CheckReplyToMailAddressWithOptions(request *CheckReplyToMailAddressRequest, runtime *util.RuntimeOptions) (_result *CheckReplyToMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MailAddressId)) {
		query["MailAddressId"] = request.MailAddressId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckReplyToMailAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckReplyToMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验回信地址
//
// @param request - CheckReplyToMailAddressRequest
//
// @return CheckReplyToMailAddressResponse
func (client *Client) CheckReplyToMailAddress(request *CheckReplyToMailAddressRequest) (_result *CheckReplyToMailAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckReplyToMailAddressResponse{}
	_body, _err := client.CheckReplyToMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建域名
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建域名
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMailAddressResponse
func (client *Client) CreateMailAddressWithOptions(request *CreateMailAddressRequest, runtime *util.RuntimeOptions) (_result *CreateMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyAddress)) {
		query["ReplyAddress"] = request.ReplyAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Sendtype)) {
		query["Sendtype"] = request.Sendtype
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMailAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMailAddressRequest
//
// @return CreateMailAddressResponse
func (client *Client) CreateMailAddress(request *CreateMailAddressRequest) (_result *CreateMailAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMailAddressResponse{}
	_body, _err := client.CreateMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建收件人列表
//
// @param request - CreateReceiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReceiverResponse
func (client *Client) CreateReceiverWithOptions(request *CreateReceiverRequest, runtime *util.RuntimeOptions) (_result *CreateReceiverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		query["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiversAlias)) {
		query["ReceiversAlias"] = request.ReceiversAlias
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiversName)) {
		query["ReceiversName"] = request.ReceiversName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReceiver"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReceiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建收件人列表
//
// @param request - CreateReceiverRequest
//
// @return CreateReceiverResponse
func (client *Client) CreateReceiver(request *CreateReceiverRequest) (_result *CreateReceiverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReceiverResponse{}
	_body, _err := client.CreateReceiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建标签
//
// @param request - CreateTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTagResponse
func (client *Client) CreateTagWithOptions(request *CreateTagRequest, runtime *util.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagDescription)) {
		query["TagDescription"] = request.TagDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTag"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标签
//
// @param request - CreateTagRequest
//
// @return CreateTagResponse
func (client *Client) CreateTag(request *CreateTagRequest) (_result *CreateTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTagResponse{}
	_body, _err := client.CreateTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用户无效地址
//
// @param request - CreateUserSuppressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserSuppressionResponse
func (client *Client) CreateUserSuppressionWithOptions(request *CreateUserSuppressionRequest, runtime *util.RuntimeOptions) (_result *CreateUserSuppressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserSuppression"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserSuppressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户无效地址
//
// @param request - CreateUserSuppressionRequest
//
// @return CreateUserSuppressionResponse
func (client *Client) CreateUserSuppression(request *CreateUserSuppressionRequest) (_result *CreateUserSuppressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserSuppressionResponse{}
	_body, _err := client.CreateUserSuppressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除域名
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从无效地址库删除无效地址
//
// @param request - DeleteInvalidAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInvalidAddressResponse
func (client *Client) DeleteInvalidAddressWithOptions(request *DeleteInvalidAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteInvalidAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ToAddress)) {
		query["ToAddress"] = request.ToAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInvalidAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInvalidAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从无效地址库删除无效地址
//
// @param request - DeleteInvalidAddressRequest
//
// @return DeleteInvalidAddressResponse
func (client *Client) DeleteInvalidAddress(request *DeleteInvalidAddressRequest) (_result *DeleteInvalidAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInvalidAddressResponse{}
	_body, _err := client.DeleteInvalidAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除IP保护信息
//
// @param request - DeleteIpfilterByEdmIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpfilterByEdmIdResponse
func (client *Client) DeleteIpfilterByEdmIdWithOptions(request *DeleteIpfilterByEdmIdRequest, runtime *util.RuntimeOptions) (_result *DeleteIpfilterByEdmIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromType)) {
		query["FromType"] = request.FromType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpfilterByEdmId"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpfilterByEdmIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除IP保护信息
//
// @param request - DeleteIpfilterByEdmIdRequest
//
// @return DeleteIpfilterByEdmIdResponse
func (client *Client) DeleteIpfilterByEdmId(request *DeleteIpfilterByEdmIdRequest) (_result *DeleteIpfilterByEdmIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpfilterByEdmIdResponse{}
	_body, _err := client.DeleteIpfilterByEdmIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除发信地址
//
// @param request - DeleteMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMailAddressResponse
func (client *Client) DeleteMailAddressWithOptions(request *DeleteMailAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MailAddressId)) {
		query["MailAddressId"] = request.MailAddressId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMailAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除发信地址
//
// @param request - DeleteMailAddressRequest
//
// @return DeleteMailAddressResponse
func (client *Client) DeleteMailAddress(request *DeleteMailAddressRequest) (_result *DeleteMailAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMailAddressResponse{}
	_body, _err := client.DeleteMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除收件人列表
//
// @param request - DeleteReceiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReceiverResponse
func (client *Client) DeleteReceiverWithOptions(request *DeleteReceiverRequest, runtime *util.RuntimeOptions) (_result *DeleteReceiverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteReceiver"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteReceiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除收件人列表
//
// @param request - DeleteReceiverRequest
//
// @return DeleteReceiverResponse
func (client *Client) DeleteReceiver(request *DeleteReceiverRequest) (_result *DeleteReceiverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteReceiverResponse{}
	_body, _err := client.DeleteReceiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除单个收件人
//
// @param request - DeleteReceiverDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReceiverDetailResponse
func (client *Client) DeleteReceiverDetailWithOptions(request *DeleteReceiverDetailRequest, runtime *util.RuntimeOptions) (_result *DeleteReceiverDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteReceiverDetail"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteReceiverDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除单个收件人
//
// @param request - DeleteReceiverDetailRequest
//
// @return DeleteReceiverDetailResponse
func (client *Client) DeleteReceiverDetail(request *DeleteReceiverDetailRequest) (_result *DeleteReceiverDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteReceiverDetailResponse{}
	_body, _err := client.DeleteReceiverDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTagResponse
func (client *Client) DeleteTagWithOptions(request *DeleteTagRequest, runtime *util.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTag"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteTagRequest
//
// @return DeleteTagResponse
func (client *Client) DeleteTag(request *DeleteTagRequest) (_result *DeleteTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagResponse{}
	_body, _err := client.DeleteTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescAccountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescAccountSummaryResponse
func (client *Client) DescAccountSummaryWithOptions(request *DescAccountSummaryRequest, runtime *util.RuntimeOptions) (_result *DescAccountSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescAccountSummary"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescAccountSummaryRequest
//
// @return DescAccountSummaryResponse
func (client *Client) DescAccountSummary(request *DescAccountSummaryRequest) (_result *DescAccountSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescAccountSummaryResponse{}
	_body, _err := client.DescAccountSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取域名详情
//
// @param request - DescDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescDomainResponse
func (client *Client) DescDomainWithOptions(request *DescDomainRequest, runtime *util.RuntimeOptions) (_result *DescDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RequireRealTimeDnsRecords)) {
		query["RequireRealTimeDnsRecords"] = request.RequireRealTimeDnsRecords
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescDomain"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取域名详情
//
// @param request - DescDomainRequest
//
// @return DescDomainResponse
func (client *Client) DescDomain(request *DescDomainRequest) (_result *DescDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescDomainResponse{}
	_body, _err := client.DescDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取IP保护信息
//
// @param request - GetIpProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpProtectionResponse
func (client *Client) GetIpProtectionWithOptions(request *GetIpProtectionRequest, runtime *util.RuntimeOptions) (_result *GetIpProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIpProtection"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIpProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取IP保护信息
//
// @param request - GetIpProtectionRequest
//
// @return GetIpProtectionResponse
func (client *Client) GetIpProtection(request *GetIpProtectionRequest) (_result *GetIpProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpProtectionResponse{}
	_body, _err := client.GetIpProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取IP防护信息
//
// @param request - GetIpfilterListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpfilterListResponse
func (client *Client) GetIpfilterListWithOptions(request *GetIpfilterListRequest, runtime *util.RuntimeOptions) (_result *GetIpfilterListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIpfilterList"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIpfilterListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取IP防护信息
//
// @param request - GetIpfilterListRequest
//
// @return GetIpfilterListResponse
func (client *Client) GetIpfilterList(request *GetIpfilterListRequest) (_result *GetIpfilterListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpfilterListResponse{}
	_body, _err := client.GetIpfilterListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户无效地址级别配置
//
// @param request - GetSuppressionListLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuppressionListLevelResponse
func (client *Client) GetSuppressionListLevelWithOptions(request *GetSuppressionListLevelRequest, runtime *util.RuntimeOptions) (_result *GetSuppressionListLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSuppressionListLevel"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSuppressionListLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户无效地址级别配置
//
// @param request - GetSuppressionListLevelRequest
//
// @return GetSuppressionListLevelResponse
func (client *Client) GetSuppressionListLevel(request *GetSuppressionListLevelRequest) (_result *GetSuppressionListLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuppressionListLevelResponse{}
	_body, _err := client.GetSuppressionListLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取跟踪信息
//
// @param request - GetTrackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrackListResponse
func (client *Client) GetTrackListWithOptions(request *GetTrackListRequest, runtime *util.RuntimeOptions) (_result *GetTrackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.OffsetCreateTime)) {
		query["OffsetCreateTime"] = request.OffsetCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.OffsetCreateTimeDesc)) {
		query["OffsetCreateTimeDesc"] = request.OffsetCreateTimeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.Total)) {
		query["Total"] = request.Total
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrackList"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取跟踪信息
//
// @param request - GetTrackListRequest
//
// @return GetTrackListResponse
func (client *Client) GetTrackList(request *GetTrackListRequest) (_result *GetTrackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrackListResponse{}
	_body, _err := client.GetTrackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据发信地址和Tag名称获取跟踪信息
//
// @param request - GetTrackListByMailFromAndTagNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrackListByMailFromAndTagNameResponse
func (client *Client) GetTrackListByMailFromAndTagNameWithOptions(request *GetTrackListByMailFromAndTagNameRequest, runtime *util.RuntimeOptions) (_result *GetTrackListByMailFromAndTagNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.OffsetCreateTime)) {
		query["OffsetCreateTime"] = request.OffsetCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.OffsetCreateTimeDesc)) {
		query["OffsetCreateTimeDesc"] = request.OffsetCreateTimeDesc
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.Total)) {
		query["Total"] = request.Total
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrackListByMailFromAndTagName"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrackListByMailFromAndTagNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据发信地址和Tag名称获取跟踪信息
//
// @param request - GetTrackListByMailFromAndTagNameRequest
//
// @return GetTrackListByMailFromAndTagNameResponse
func (client *Client) GetTrackListByMailFromAndTagName(request *GetTrackListByMailFromAndTagNameRequest) (_result *GetTrackListByMailFromAndTagNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrackListByMailFromAndTagNameResponse{}
	_body, _err := client.GetTrackListByMailFromAndTagNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取账号详情
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取账号详情
//
// @return GetUserResponse
func (client *Client) GetUser() (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出用户无效地址
//
// @param request - ListUserSuppressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserSuppressionResponse
func (client *Client) ListUserSuppressionWithOptions(request *ListUserSuppressionRequest, runtime *util.RuntimeOptions) (_result *ListUserSuppressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.EndBounceTime)) {
		query["EndBounceTime"] = request.EndBounceTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndCreateTime)) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartBounceTime)) {
		query["StartBounceTime"] = request.StartBounceTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartCreateTime)) {
		query["StartCreateTime"] = request.StartCreateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserSuppression"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserSuppressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出用户无效地址
//
// @param request - ListUserSuppressionRequest
//
// @return ListUserSuppressionResponse
func (client *Client) ListUserSuppression(request *ListUserSuppressionRequest) (_result *ListUserSuppressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserSuppressionResponse{}
	_body, _err := client.ListUserSuppressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改发信地址
//
// @param request - ModifyMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMailAddressResponse
func (client *Client) ModifyMailAddressWithOptions(request *ModifyMailAddressRequest, runtime *util.RuntimeOptions) (_result *ModifyMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MailAddressId)) {
		query["MailAddressId"] = request.MailAddressId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyAddress)) {
		query["ReplyAddress"] = request.ReplyAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMailAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改发信地址
//
// @param request - ModifyMailAddressRequest
//
// @return ModifyMailAddressResponse
func (client *Client) ModifyMailAddress(request *ModifyMailAddressRequest) (_result *ModifyMailAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMailAddressResponse{}
	_body, _err := client.ModifyMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改域级别密码
//
// @param request - ModifyPWByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPWByDomainResponse
func (client *Client) ModifyPWByDomainWithOptions(request *ModifyPWByDomainRequest, runtime *util.RuntimeOptions) (_result *ModifyPWByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPWByDomain"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPWByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改域级别密码
//
// @param request - ModifyPWByDomainRequest
//
// @return ModifyPWByDomainResponse
func (client *Client) ModifyPWByDomain(request *ModifyPWByDomainRequest) (_result *ModifyPWByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPWByDomainResponse{}
	_body, _err := client.ModifyPWByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改标签
//
// @param request - ModifyTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTagResponse
func (client *Client) ModifyTagWithOptions(request *ModifyTagRequest, runtime *util.RuntimeOptions) (_result *ModifyTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagDescription)) {
		query["TagDescription"] = request.TagDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTag"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改标签
//
// @param request - ModifyTagRequest
//
// @return ModifyTagResponse
func (client *Client) ModifyTag(request *ModifyTagRequest) (_result *ModifyTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTagResponse{}
	_body, _err := client.ModifyTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询域名信息
//
// @param request - QueryDomainByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainByParamResponse
func (client *Client) QueryDomainByParamWithOptions(request *QueryDomainByParamRequest, runtime *util.RuntimeOptions) (_result *QueryDomainByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainByParam"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名信息
//
// @param request - QueryDomainByParamRequest
//
// @return QueryDomainByParamResponse
func (client *Client) QueryDomainByParam(request *QueryDomainByParamRequest) (_result *QueryDomainByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainByParamResponse{}
	_body, _err := client.QueryDomainByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// NextStart修改为string
//
// @param request - QueryInvalidAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInvalidAddressResponse
func (client *Client) QueryInvalidAddressWithOptions(request *QueryInvalidAddressRequest, runtime *util.RuntimeOptions) (_result *QueryInvalidAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.NextStart)) {
		query["NextStart"] = request.NextStart
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInvalidAddress"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInvalidAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// NextStart修改为string
//
// @param request - QueryInvalidAddressRequest
//
// @return QueryInvalidAddressResponse
func (client *Client) QueryInvalidAddress(request *QueryInvalidAddressRequest) (_result *QueryInvalidAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInvalidAddressResponse{}
	_body, _err := client.QueryInvalidAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryMailAddressByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMailAddressByParamResponse
func (client *Client) QueryMailAddressByParamWithOptions(request *QueryMailAddressByParamRequest, runtime *util.RuntimeOptions) (_result *QueryMailAddressByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Sendtype)) {
		query["Sendtype"] = request.Sendtype
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMailAddressByParam"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMailAddressByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMailAddressByParamRequest
//
// @return QueryMailAddressByParamResponse
func (client *Client) QueryMailAddressByParam(request *QueryMailAddressByParamRequest) (_result *QueryMailAddressByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMailAddressByParamResponse{}
	_body, _err := client.QueryMailAddressByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询收信人列表详情
//
// @param request - QueryReceiverByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReceiverByParamResponse
func (client *Client) QueryReceiverByParamWithOptions(request *QueryReceiverByParamRequest, runtime *util.RuntimeOptions) (_result *QueryReceiverByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryReceiverByParam"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReceiverByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询收信人列表详情
//
// @param request - QueryReceiverByParamRequest
//
// @return QueryReceiverByParamResponse
func (client *Client) QueryReceiverByParam(request *QueryReceiverByParamRequest) (_result *QueryReceiverByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryReceiverByParamResponse{}
	_body, _err := client.QueryReceiverByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询收信人列表详情信息
//
// @param request - QueryReceiverDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReceiverDetailResponse
func (client *Client) QueryReceiverDetailWithOptions(request *QueryReceiverDetailRequest, runtime *util.RuntimeOptions) (_result *QueryReceiverDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.NextStart)) {
		query["NextStart"] = request.NextStart
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryReceiverDetail"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReceiverDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询收信人列表详情信息
//
// @param request - QueryReceiverDetailRequest
//
// @return QueryReceiverDetailResponse
func (client *Client) QueryReceiverDetail(request *QueryReceiverDetailRequest) (_result *QueryReceiverDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryReceiverDetailResponse{}
	_body, _err := client.QueryReceiverDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryTagByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagByParamResponse
func (client *Client) QueryTagByParamWithOptions(request *QueryTagByParamRequest, runtime *util.RuntimeOptions) (_result *QueryTagByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTagByParam"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTagByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTagByParamRequest
//
// @return QueryTagByParamResponse
func (client *Client) QueryTagByParam(request *QueryTagByParamRequest) (_result *QueryTagByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTagByParamResponse{}
	_body, _err := client.QueryTagByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryTaskByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskByParamResponse
func (client *Client) QueryTaskByParamWithOptions(request *QueryTaskByParamRequest, runtime *util.RuntimeOptions) (_result *QueryTaskByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskByParam"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskByParamRequest
//
// @return QueryTaskByParamResponse
func (client *Client) QueryTaskByParam(request *QueryTaskByParamRequest) (_result *QueryTaskByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskByParamResponse{}
	_body, _err := client.QueryTaskByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户无效地址
//
// @param request - RemoveUserSuppressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserSuppressionResponse
func (client *Client) RemoveUserSuppressionWithOptions(request *RemoveUserSuppressionRequest, runtime *util.RuntimeOptions) (_result *RemoveUserSuppressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SuppressionIds)) {
		query["SuppressionIds"] = request.SuppressionIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserSuppression"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserSuppressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户无效地址
//
// @param request - RemoveUserSuppressionRequest
//
// @return RemoveUserSuppressionResponse
func (client *Client) RemoveUserSuppression(request *RemoveUserSuppressionRequest) (_result *RemoveUserSuppressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserSuppressionResponse{}
	_body, _err := client.RemoveUserSuppressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 建单个收件人
//
// @param request - SaveReceiverDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveReceiverDetailResponse
func (client *Client) SaveReceiverDetailWithOptions(request *SaveReceiverDetailRequest, runtime *util.RuntimeOptions) (_result *SaveReceiverDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		query["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveReceiverDetail"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveReceiverDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 建单个收件人
//
// @param request - SaveReceiverDetailRequest
//
// @return SaveReceiverDetailResponse
func (client *Client) SaveReceiverDetail(request *SaveReceiverDetailRequest) (_result *SaveReceiverDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveReceiverDetailResponse{}
	_body, _err := client.SaveReceiverDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送模板测试邮件
//
// @param request - SendTestByTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTestByTemplateResponse
func (client *Client) SendTestByTemplateWithOptions(request *SendTestByTemplateRequest, runtime *util.RuntimeOptions) (_result *SendTestByTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.Birthday)) {
		query["Birthday"] = request.Birthday
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		query["Gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		query["NickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendTestByTemplate"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTestByTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送模板测试邮件
//
// @param request - SendTestByTemplateRequest
//
// @return SendTestByTemplateResponse
func (client *Client) SendTestByTemplate(request *SendTestByTemplateRequest) (_result *SendTestByTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendTestByTemplateResponse{}
	_body, _err := client.SendTestByTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定条件下的发送数据
//
// @param request - SenderStatisticsByTagNameAndBatchIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SenderStatisticsByTagNameAndBatchIDResponse
func (client *Client) SenderStatisticsByTagNameAndBatchIDWithOptions(request *SenderStatisticsByTagNameAndBatchIDRequest, runtime *util.RuntimeOptions) (_result *SenderStatisticsByTagNameAndBatchIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SenderStatisticsByTagNameAndBatchID"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SenderStatisticsByTagNameAndBatchIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定条件下的发送数据
//
// @param request - SenderStatisticsByTagNameAndBatchIDRequest
//
// @return SenderStatisticsByTagNameAndBatchIDResponse
func (client *Client) SenderStatisticsByTagNameAndBatchID(request *SenderStatisticsByTagNameAndBatchIDRequest) (_result *SenderStatisticsByTagNameAndBatchIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SenderStatisticsByTagNameAndBatchIDResponse{}
	_body, _err := client.SenderStatisticsByTagNameAndBatchIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询投递结果详情
//
// @param request - SenderStatisticsDetailByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SenderStatisticsDetailByParamResponse
func (client *Client) SenderStatisticsDetailByParamWithOptions(request *SenderStatisticsDetailByParamRequest, runtime *util.RuntimeOptions) (_result *SenderStatisticsDetailByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.NextStart)) {
		query["NextStart"] = request.NextStart
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.ToAddress)) {
		query["ToAddress"] = request.ToAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SenderStatisticsDetailByParam"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SenderStatisticsDetailByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询投递结果详情
//
// @param request - SenderStatisticsDetailByParamRequest
//
// @return SenderStatisticsDetailByParamResponse
func (client *Client) SenderStatisticsDetailByParam(request *SenderStatisticsDetailByParamRequest) (_result *SenderStatisticsDetailByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SenderStatisticsDetailByParamResponse{}
	_body, _err := client.SenderStatisticsDetailByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置用户无效地址级别配置
//
// @param request - SetSuppressionListLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSuppressionListLevelResponse
func (client *Client) SetSuppressionListLevelWithOptions(request *SetSuppressionListLevelRequest, runtime *util.RuntimeOptions) (_result *SetSuppressionListLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SuppressionListLevel)) {
		query["SuppressionListLevel"] = request.SuppressionListLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSuppressionListLevel"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSuppressionListLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置用户无效地址级别配置
//
// @param request - SetSuppressionListLevelRequest
//
// @return SetSuppressionListLevelResponse
func (client *Client) SetSuppressionListLevel(request *SetSuppressionListLevelRequest) (_result *SetSuppressionListLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSuppressionListLevelResponse{}
	_body, _err := client.SetSuppressionListLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// API发信
//
// @param request - SingleSendMailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SingleSendMailResponse
func (client *Client) SingleSendMailWithOptions(request *SingleSendMailRequest, runtime *util.RuntimeOptions) (_result *SingleSendMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ClickTrace)) {
		query["ClickTrace"] = request.ClickTrace
	}

	if !tea.BoolValue(util.IsUnset(request.FromAlias)) {
		query["FromAlias"] = request.FromAlias
	}

	if !tea.BoolValue(util.IsUnset(request.HtmlBody)) {
		query["HtmlBody"] = request.HtmlBody
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyAddress)) {
		query["ReplyAddress"] = request.ReplyAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyAddressAlias)) {
		query["ReplyAddressAlias"] = request.ReplyAddressAlias
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyToAddress)) {
		query["ReplyToAddress"] = request.ReplyToAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.TextBody)) {
		query["TextBody"] = request.TextBody
	}

	if !tea.BoolValue(util.IsUnset(request.ToAddress)) {
		query["ToAddress"] = request.ToAddress
	}

	if !tea.BoolValue(util.IsUnset(request.UnSubscribeFilterLevel)) {
		query["UnSubscribeFilterLevel"] = request.UnSubscribeFilterLevel
	}

	if !tea.BoolValue(util.IsUnset(request.UnSubscribeLinkType)) {
		query["UnSubscribeLinkType"] = request.UnSubscribeLinkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleSendMail"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SingleSendMailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// API发信
//
// @param request - SingleSendMailRequest
//
// @return SingleSendMailResponse
func (client *Client) SingleSendMail(request *SingleSendMailRequest) (_result *SingleSendMailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SingleSendMailResponse{}
	_body, _err := client.SingleSendMailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新IP防护API
//
// @param request - UpdateIpProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpProtectionResponse
func (client *Client) UpdateIpProtectionWithOptions(request *UpdateIpProtectionRequest, runtime *util.RuntimeOptions) (_result *UpdateIpProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpProtection)) {
		query["IpProtection"] = request.IpProtection
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpProtection"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新IP防护API
//
// @param request - UpdateIpProtectionRequest
//
// @return UpdateIpProtectionResponse
func (client *Client) UpdateIpProtection(request *UpdateIpProtectionRequest) (_result *UpdateIpProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpProtectionResponse{}
	_body, _err := client.UpdateIpProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新帐号信息
//
// @param tmpReq - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(tmpReq *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.User)) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, tea.String("User"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserShrink)) {
		body["User"] = request.UserShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2015-11-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新帐号信息
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
