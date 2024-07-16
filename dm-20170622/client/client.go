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
	// F814E960-5AEE-5CB1-881B-6F1A3250B55A
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

type BatchSendMailRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AddressType *int32  `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	ClickTrace  *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ReceiversName        *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias    *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// This parameter is required.
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
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
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
	// This parameter is required.
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
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) SetDomainStatus(v string) *CheckDomainResponseBody {
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

type CheckDomainDnsRequest struct {
	// This parameter is required.
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDomainDnsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainDnsRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainDnsRequest) SetDomainId(v int32) *CheckDomainDnsRequest {
	s.DomainId = &v
	return s
}

func (s *CheckDomainDnsRequest) SetOwnerId(v int64) *CheckDomainDnsRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDomainDnsRequest) SetResourceOwnerAccount(v string) *CheckDomainDnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDomainDnsRequest) SetResourceOwnerId(v int64) *CheckDomainDnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckDomainDnsRequest) SetType(v string) *CheckDomainDnsRequest {
	s.Type = &v
	return s
}

type CheckDomainDnsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckDomainDnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainDnsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainDnsResponseBody) SetRequestId(v string) *CheckDomainDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainDnsResponseBody) SetStatus(v int32) *CheckDomainDnsResponseBody {
	s.Status = &v
	return s
}

type CheckDomainDnsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDomainDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDomainDnsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainDnsResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainDnsResponse) SetHeaders(v map[string]*string) *CheckDomainDnsResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainDnsResponse) SetStatusCode(v int32) *CheckDomainDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainDnsResponse) SetBody(v *CheckDomainDnsResponseBody) *CheckDomainDnsResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// This parameter is required.
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
	DomainId  *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
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
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
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
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ReceiversAlias *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	// This parameter is required.
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
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// test description
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagId     *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
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

type CreateTemplateRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SmsContent           *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsType              *int32  `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// This parameter is required.
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateSubject  *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	TemplateText     *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	TemplateType     *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetFromType(v int32) *CreateTemplateRequest {
	s.FromType = &v
	return s
}

func (s *CreateTemplateRequest) SetOwnerId(v int64) *CreateTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTemplateRequest) SetRemark(v string) *CreateTemplateRequest {
	s.Remark = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceOwnerAccount(v string) *CreateTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceOwnerId(v int64) *CreateTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTemplateRequest) SetSmsContent(v string) *CreateTemplateRequest {
	s.SmsContent = &v
	return s
}

func (s *CreateTemplateRequest) SetSmsType(v int32) *CreateTemplateRequest {
	s.SmsType = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateNickName(v string) *CreateTemplateRequest {
	s.TemplateNickName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateSubject(v string) *CreateTemplateRequest {
	s.TemplateSubject = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateText(v string) *CreateTemplateRequest {
	s.TemplateText = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateType(v int32) *CreateTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateId(v int32) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// This parameter is required.
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

type DeleteMailAddressRequest struct {
	// This parameter is required.
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
	// This parameter is required.
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
	// This parameter is required.
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

type DeleteTemplateRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetFromType(v int32) *DeleteTemplateRequest {
	s.FromType = &v
	return s
}

func (s *DeleteTemplateRequest) SetOwnerId(v int64) *DeleteTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) SetResourceOwnerAccount(v string) *DeleteTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTemplateRequest) SetResourceOwnerId(v int64) *DeleteTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateId(v int32) *DeleteTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
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
	DailyQuota      *int32  `json:"DailyQuota,omitempty" xml:"DailyQuota,omitempty"`
	DayuStatus      *int32  `json:"DayuStatus,omitempty" xml:"DayuStatus,omitempty"`
	Domains         *int32  `json:"Domains,omitempty" xml:"Domains,omitempty"`
	EnableTimes     *int32  `json:"EnableTimes,omitempty" xml:"EnableTimes,omitempty"`
	MailAddresses   *int32  `json:"MailAddresses,omitempty" xml:"MailAddresses,omitempty"`
	MaxQuotaLevel   *int32  `json:"MaxQuotaLevel,omitempty" xml:"MaxQuotaLevel,omitempty"`
	MonthQuota      *int32  `json:"MonthQuota,omitempty" xml:"MonthQuota,omitempty"`
	QuotaLevel      *int32  `json:"QuotaLevel,omitempty" xml:"QuotaLevel,omitempty"`
	Receivers       *int32  `json:"Receivers,omitempty" xml:"Receivers,omitempty"`
	RemainFreeQuota *int32  `json:"RemainFreeQuota,omitempty" xml:"RemainFreeQuota,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsRecord       *int32  `json:"SmsRecord,omitempty" xml:"SmsRecord,omitempty"`
	SmsSign         *int32  `json:"SmsSign,omitempty" xml:"SmsSign,omitempty"`
	SmsTemplates    *int32  `json:"SmsTemplates,omitempty" xml:"SmsTemplates,omitempty"`
	Tags            *int32  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Templates       *int32  `json:"Templates,omitempty" xml:"Templates,omitempty"`
	UserStatus      *int32  `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
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
	CnameAuthStatus    *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	CnameConfirmStatus *string `json:"CnameConfirmStatus,omitempty" xml:"CnameConfirmStatus,omitempty"`
	CnameRecord        *string `json:"CnameRecord,omitempty" xml:"CnameRecord,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultDomain      *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	DkimAuthStatus     *string `json:"DkimAuthStatus,omitempty" xml:"DkimAuthStatus,omitempty"`
	DkimPublicKey      *string `json:"DkimPublicKey,omitempty" xml:"DkimPublicKey,omitempty"`
	DkimRR             *string `json:"DkimRR,omitempty" xml:"DkimRR,omitempty"`
	DmarcAuthStatus    *int32  `json:"DmarcAuthStatus,omitempty" xml:"DmarcAuthStatus,omitempty"`
	DmarcHostRecord    *string `json:"DmarcHostRecord,omitempty" xml:"DmarcHostRecord,omitempty"`
	DmarcRecord        *string `json:"DmarcRecord,omitempty" xml:"DmarcRecord,omitempty"`
	DnsDmarc           *string `json:"DnsDmarc,omitempty" xml:"DnsDmarc,omitempty"`
	DnsMx              *string `json:"DnsMx,omitempty" xml:"DnsMx,omitempty"`
	DnsSpf             *string `json:"DnsSpf,omitempty" xml:"DnsSpf,omitempty"`
	DnsTxt             *string `json:"DnsTxt,omitempty" xml:"DnsTxt,omitempty"`
	DomainId           *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus       *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType         *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	HostRecord         *string `json:"HostRecord,omitempty" xml:"HostRecord,omitempty"`
	IcpStatus          *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	MxAuthStatus       *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	MxRecord           *string `json:"MxRecord,omitempty" xml:"MxRecord,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpfAuthStatus      *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	SpfRecord          *string `json:"SpfRecord,omitempty" xml:"SpfRecord,omitempty"`
	SpfRecordV2        *string `json:"SpfRecordV2,omitempty" xml:"SpfRecordV2,omitempty"`
	TlDomainName       *string `json:"TlDomainName,omitempty" xml:"TlDomainName,omitempty"`
	TracefRecord       *string `json:"TracefRecord,omitempty" xml:"TracefRecord,omitempty"`
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

type DescTemplateRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescTemplateRequest) SetFromType(v int32) *DescTemplateRequest {
	s.FromType = &v
	return s
}

func (s *DescTemplateRequest) SetOwnerId(v int64) *DescTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescTemplateRequest) SetResourceOwnerAccount(v string) *DescTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescTemplateRequest) SetResourceOwnerId(v int64) *DescTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescTemplateRequest) SetTemplateId(v int32) *DescTemplateRequest {
	s.TemplateId = &v
	return s
}

type DescTemplateResponseBody struct {
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Remark           *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsContent       *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsType          *string `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateStatus   *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateSubject  *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	TemplateText     *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	TemplateType     *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescTemplateResponseBody) SetCreateTime(v string) *DescTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescTemplateResponseBody) SetRemark(v string) *DescTemplateResponseBody {
	s.Remark = &v
	return s
}

func (s *DescTemplateResponseBody) SetRequestId(v string) *DescTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescTemplateResponseBody) SetSmsContent(v string) *DescTemplateResponseBody {
	s.SmsContent = &v
	return s
}

func (s *DescTemplateResponseBody) SetSmsType(v string) *DescTemplateResponseBody {
	s.SmsType = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateName(v string) *DescTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateNickName(v string) *DescTemplateResponseBody {
	s.TemplateNickName = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateStatus(v string) *DescTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateSubject(v string) *DescTemplateResponseBody {
	s.TemplateSubject = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateText(v string) *DescTemplateResponseBody {
	s.TemplateText = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateType(v string) *DescTemplateResponseBody {
	s.TemplateType = &v
	return s
}

type DescTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescTemplateResponse) SetHeaders(v map[string]*string) *DescTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescTemplateResponse) SetStatusCode(v int32) *DescTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescTemplateResponse) SetBody(v *DescTemplateResponseBody) *DescTemplateResponse {
	s.Body = v
	return s
}

type GetAccountListRequest struct {
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Total                *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAccountListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListRequest) GoString() string {
	return s.String()
}

func (s *GetAccountListRequest) SetOffset(v string) *GetAccountListRequest {
	s.Offset = &v
	return s
}

func (s *GetAccountListRequest) SetOffsetCreateTime(v string) *GetAccountListRequest {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetAccountListRequest) SetOffsetCreateTimeDesc(v string) *GetAccountListRequest {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetAccountListRequest) SetOwnerId(v int64) *GetAccountListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAccountListRequest) SetPageNumber(v string) *GetAccountListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAccountListRequest) SetPageSize(v string) *GetAccountListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountListRequest) SetResourceOwnerAccount(v string) *GetAccountListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAccountListRequest) SetResourceOwnerId(v int64) *GetAccountListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAccountListRequest) SetTotal(v string) *GetAccountListRequest {
	s.Total = &v
	return s
}

type GetAccountListResponseBody struct {
	PageNo    *int32                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Data      *GetAccountListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountListResponseBody) SetPageNo(v int32) *GetAccountListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetAccountListResponseBody) SetPageSize(v int32) *GetAccountListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAccountListResponseBody) SetRequestId(v string) *GetAccountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountListResponseBody) SetTotal(v int32) *GetAccountListResponseBody {
	s.Total = &v
	return s
}

func (s *GetAccountListResponseBody) SetData(v *GetAccountListResponseBodyData) *GetAccountListResponseBody {
	s.Data = v
	return s
}

type GetAccountListResponseBodyData struct {
	AccountNotificationInfo []*GetAccountListResponseBodyDataAccountNotificationInfo `json:"accountNotificationInfo,omitempty" xml:"accountNotificationInfo,omitempty" type:"Repeated"`
}

func (s GetAccountListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccountListResponseBodyData) SetAccountNotificationInfo(v []*GetAccountListResponseBodyDataAccountNotificationInfo) *GetAccountListResponseBodyData {
	s.AccountNotificationInfo = v
	return s
}

type GetAccountListResponseBodyDataAccountNotificationInfo struct {
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAccountListResponseBodyDataAccountNotificationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListResponseBodyDataAccountNotificationInfo) GoString() string {
	return s.String()
}

func (s *GetAccountListResponseBodyDataAccountNotificationInfo) SetRegion(v string) *GetAccountListResponseBodyDataAccountNotificationInfo {
	s.Region = &v
	return s
}

func (s *GetAccountListResponseBodyDataAccountNotificationInfo) SetStatus(v string) *GetAccountListResponseBodyDataAccountNotificationInfo {
	s.Status = &v
	return s
}

func (s *GetAccountListResponseBodyDataAccountNotificationInfo) SetUpdateTime(v string) *GetAccountListResponseBodyDataAccountNotificationInfo {
	s.UpdateTime = &v
	return s
}

type GetAccountListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListResponse) GoString() string {
	return s.String()
}

func (s *GetAccountListResponse) SetHeaders(v map[string]*string) *GetAccountListResponse {
	s.Headers = v
	return s
}

func (s *GetAccountListResponse) SetStatusCode(v int32) *GetAccountListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountListResponse) SetBody(v *GetAccountListResponseBody) *GetAccountListResponse {
	s.Body = v
	return s
}

type GetMailAddressMsgCallBackUrlRequest struct {
	// This parameter is required.
	MailFrom             *string `json:"MailFrom,omitempty" xml:"MailFrom,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMailAddressMsgCallBackUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMailAddressMsgCallBackUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMailAddressMsgCallBackUrlRequest) SetMailFrom(v string) *GetMailAddressMsgCallBackUrlRequest {
	s.MailFrom = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlRequest) SetOwnerId(v int64) *GetMailAddressMsgCallBackUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlRequest) SetResourceOwnerAccount(v string) *GetMailAddressMsgCallBackUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlRequest) SetResourceOwnerId(v int64) *GetMailAddressMsgCallBackUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetMailAddressMsgCallBackUrlResponseBody struct {
	NotifyUrl       *int32  `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	NotifyUrlStatus *int32  `json:"NotifyUrlStatus,omitempty" xml:"NotifyUrlStatus,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMailAddressMsgCallBackUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMailAddressMsgCallBackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMailAddressMsgCallBackUrlResponseBody) SetNotifyUrl(v int32) *GetMailAddressMsgCallBackUrlResponseBody {
	s.NotifyUrl = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlResponseBody) SetNotifyUrlStatus(v int32) *GetMailAddressMsgCallBackUrlResponseBody {
	s.NotifyUrlStatus = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlResponseBody) SetRequestId(v string) *GetMailAddressMsgCallBackUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetMailAddressMsgCallBackUrlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMailAddressMsgCallBackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMailAddressMsgCallBackUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMailAddressMsgCallBackUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMailAddressMsgCallBackUrlResponse) SetHeaders(v map[string]*string) *GetMailAddressMsgCallBackUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMailAddressMsgCallBackUrlResponse) SetStatusCode(v int32) *GetMailAddressMsgCallBackUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlResponse) SetBody(v *GetMailAddressMsgCallBackUrlResponseBody) *GetMailAddressMsgCallBackUrlResponse {
	s.Body = v
	return s
}

type GetTrackListRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-29
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
	// 2022-09-29
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
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	RcptClickCount *int64 `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	// example:
	//
	// 0
	RcptClickRate *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	// example:
	//
	// 0
	RcptOpenCount *int64 `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
	// example:
	//
	// 0
	RcptOpenRate *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	// example:
	//
	// 0
	RcptUniqueClickCount *int64 `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	// example:
	//
	// 0
	RcptUniqueClickRate *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	// example:
	//
	// 0
	RcptUniqueOpenCount *int64 `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	// example:
	//
	// 0
	RcptUniqueOpenRate *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	// example:
	//
	// 0
	TotalNumber *int64 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetTrackListResponseBodyDataStat) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBodyDataStat) SetCreateTime(v int64) *GetTrackListResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickCount(v int64) *GetTrackListResponseBodyDataStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenCount(v int64) *GetTrackListResponseBodyDataStat {
	s.RcptOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickCount(v int64) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenCount(v int64) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetTotalNumber(v int64) *GetTrackListResponseBodyDataStat {
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

type ModifyMailAddressRequest struct {
	// This parameter is required.
	MailAddressId        *int32  `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *ModifyPWByDomainRequest) SetPassword(v string) *ModifyPWByDomainRequest {
	s.Password = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetResourceOwnerId(v string) *ModifyPWByDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyPWByDomainResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// test description
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
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

type ModifyTemplateRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SmsContent           *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsType              *int32  `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// This parameter is required.
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateSubject  *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	TemplateText     *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
}

func (s ModifyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequest) SetFromType(v int32) *ModifyTemplateRequest {
	s.FromType = &v
	return s
}

func (s *ModifyTemplateRequest) SetOwnerId(v int64) *ModifyTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTemplateRequest) SetRemark(v string) *ModifyTemplateRequest {
	s.Remark = &v
	return s
}

func (s *ModifyTemplateRequest) SetResourceOwnerAccount(v string) *ModifyTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTemplateRequest) SetResourceOwnerId(v int64) *ModifyTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTemplateRequest) SetSmsContent(v string) *ModifyTemplateRequest {
	s.SmsContent = &v
	return s
}

func (s *ModifyTemplateRequest) SetSmsType(v int32) *ModifyTemplateRequest {
	s.SmsType = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateId(v int32) *ModifyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateName(v string) *ModifyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateNickName(v string) *ModifyTemplateRequest {
	s.TemplateNickName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateSubject(v string) *ModifyTemplateRequest {
	s.TemplateSubject = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateText(v string) *ModifyTemplateRequest {
	s.TemplateText = &v
	return s
}

type ModifyTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponseBody) SetRequestId(v string) *ModifyTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponse) SetHeaders(v map[string]*string) *ModifyTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateResponse) SetStatusCode(v int32) *ModifyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateResponse) SetBody(v *ModifyTemplateResponseBody) *ModifyTemplateResponse {
	s.Body = v
	return s
}

type QueryDomainByParamRequest struct {
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
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
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	ConfirmStatus   *string `json:"ConfirmStatus,omitempty" xml:"ConfirmStatus,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainId        *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainRecord    *string `json:"DomainRecord,omitempty" xml:"DomainRecord,omitempty"`
	DomainStatus    *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	IcpStatus       *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	MxAuthStatus    *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	SpfAuthStatus   *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	UtcCreateTime   *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
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
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Length               *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	NextStart  *int32                               `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryInvalidAddressResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryInvalidAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBody) SetNextStart(v int32) *QueryInvalidAddressResponseBody {
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
	LastUpdateTime    *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	ToAddress         *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UtcLastUpdateTime *int64  `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
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
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Sendtype             *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
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
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DailyCount    *string `json:"DailyCount,omitempty" xml:"DailyCount,omitempty"`
	DailyReqCount *string `json:"DailyReqCount,omitempty" xml:"DailyReqCount,omitempty"`
	DomainStatus  *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	MonthCount    *string `json:"MonthCount,omitempty" xml:"MonthCount,omitempty"`
	MonthReqCount *string `json:"MonthReqCount,omitempty" xml:"MonthReqCount,omitempty"`
	ReplyAddress  *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyStatus   *string `json:"ReplyStatus,omitempty" xml:"ReplyStatus,omitempty"`
	Sendtype      *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
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
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
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
	NextStart  *string                               `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Count           *string `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Desc            *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	ReceiverId      *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ReceiversAlias  *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	ReceiversName   *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ReceiversStatus *string `json:"ReceiversStatus,omitempty" xml:"ReceiversStatus,omitempty"`
	UtcCreateTime   *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
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
	KeyWord   *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	NextStart *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
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
	DataSchema *string                              `json:"DataSchema,omitempty" xml:"DataSchema,omitempty"`
	NextStart  *string                              `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
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
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
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
	// example:
	//
	// test description
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// example:
	//
	// 52366
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// test
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
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
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
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AddressType   *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	RequestCount  *string `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	TagName       *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus    *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
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

type QueryTemplateByParamRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTemplateByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamRequest) SetFromType(v int32) *QueryTemplateByParamRequest {
	s.FromType = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetKeyWord(v string) *QueryTemplateByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetOwnerId(v int64) *QueryTemplateByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetPageNo(v int32) *QueryTemplateByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetPageSize(v int32) *QueryTemplateByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetResourceOwnerAccount(v string) *QueryTemplateByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetResourceOwnerId(v int64) *QueryTemplateByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetStatus(v int32) *QueryTemplateByParamRequest {
	s.Status = &v
	return s
}

type QueryTemplateByParamResponseBody struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *QueryTemplateByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryTemplateByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBody) SetPageNumber(v int32) *QueryTemplateByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetPageSize(v int32) *QueryTemplateByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetRequestId(v string) *QueryTemplateByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetTotalCount(v int32) *QueryTemplateByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetData(v *QueryTemplateByParamResponseBodyData) *QueryTemplateByParamResponseBody {
	s.Data = v
	return s
}

type QueryTemplateByParamResponseBodyData struct {
	Template []*QueryTemplateByParamResponseBodyDataTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Repeated"`
}

func (s QueryTemplateByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBodyData) SetTemplate(v []*QueryTemplateByParamResponseBodyDataTemplate) *QueryTemplateByParamResponseBodyData {
	s.Template = v
	return s
}

type QueryTemplateByParamResponseBodyDataTemplate struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SmsStatus       *int32  `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	SmsTemplateCode *int32  `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	Smsrejectinfo   *int32  `json:"Smsrejectinfo,omitempty" xml:"Smsrejectinfo,omitempty"`
	TemplateComment *string `json:"TemplateComment,omitempty" xml:"TemplateComment,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateStatus  *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateType    *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	UtcCreatetime   *int64  `json:"UtcCreatetime,omitempty" xml:"UtcCreatetime,omitempty"`
}

func (s QueryTemplateByParamResponseBodyDataTemplate) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetCreateTime(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.CreateTime = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsStatus(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.SmsStatus = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsTemplateCode(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.SmsTemplateCode = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsrejectinfo(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.Smsrejectinfo = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateComment(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateComment = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateId(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateId = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateName(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateName = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateStatus(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateStatus = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateType(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateType = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetUtcCreatetime(v int64) *QueryTemplateByParamResponseBodyDataTemplate {
	s.UtcCreatetime = &v
	return s
}

type QueryTemplateByParamResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTemplateByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTemplateByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponse) SetHeaders(v map[string]*string) *QueryTemplateByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryTemplateByParamResponse) SetStatusCode(v int32) *QueryTemplateByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTemplateByParamResponse) SetBody(v *QueryTemplateByParamResponseBody) *QueryTemplateByParamResponse {
	s.Body = v
	return s
}

type SaveReceiverDetailRequest struct {
	// This parameter is required.
	Detail  *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
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
	Data         *SaveReceiverDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCount   *int32                              `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessCount *int32                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
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

type SenderStatisticsByTagNameAndBatchIDRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagName   *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaildCount         *string `json:"faildCount,omitempty" xml:"faildCount,omitempty"`
	RequestCount       *string `json:"requestCount,omitempty" xml:"requestCount,omitempty"`
	SucceededPercent   *string `json:"succeededPercent,omitempty" xml:"succeededPercent,omitempty"`
	SuccessCount       *string `json:"successCount,omitempty" xml:"successCount,omitempty"`
	UnavailableCount   *string `json:"unavailableCount,omitempty" xml:"unavailableCount,omitempty"`
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
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Length               *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	ToAddress            *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
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
	NextStart *int32                                         `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SenderStatisticsDetailByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s SenderStatisticsDetailByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBody) SetNextStart(v int32) *SenderStatisticsDetailByParamResponseBody {
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
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ErrorClassification *string `json:"ErrorClassification,omitempty" xml:"ErrorClassification,omitempty"`
	LastUpdateTime      *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Subject             *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	ToAddress           *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UtcLastUpdateTime   *string `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
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

type SingleSendMailRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AddressType       *int32  `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	ClickTrace        *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
	FromAlias         *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	HtmlBody          *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReplyAddress      *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
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
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
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

type SingleSendMailV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// test***@example.net
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
	HtmlBody             *string                                        `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	HtmlBodyPlaceHolders []*SingleSendMailV2RequestHtmlBodyPlaceHolders `json:"HtmlBodyPlaceHolders,omitempty" xml:"HtmlBodyPlaceHolders,omitempty" type:"Repeated"`
	OwnerId              *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// test2***@example.net
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
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s SingleSendMailV2Request) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailV2Request) GoString() string {
	return s.String()
}

func (s *SingleSendMailV2Request) SetAccountName(v string) *SingleSendMailV2Request {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailV2Request) SetAddressType(v int32) *SingleSendMailV2Request {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailV2Request) SetClickTrace(v string) *SingleSendMailV2Request {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailV2Request) SetFromAlias(v string) *SingleSendMailV2Request {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailV2Request) SetHtmlBody(v string) *SingleSendMailV2Request {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailV2Request) SetHtmlBodyPlaceHolders(v []*SingleSendMailV2RequestHtmlBodyPlaceHolders) *SingleSendMailV2Request {
	s.HtmlBodyPlaceHolders = v
	return s
}

func (s *SingleSendMailV2Request) SetOwnerId(v int64) *SingleSendMailV2Request {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailV2Request) SetReplyAddress(v string) *SingleSendMailV2Request {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailV2Request) SetReplyAddressAlias(v string) *SingleSendMailV2Request {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailV2Request) SetReplyToAddress(v bool) *SingleSendMailV2Request {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailV2Request) SetResourceOwnerAccount(v string) *SingleSendMailV2Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailV2Request) SetResourceOwnerId(v int64) *SingleSendMailV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailV2Request) SetSubject(v string) *SingleSendMailV2Request {
	s.Subject = &v
	return s
}

func (s *SingleSendMailV2Request) SetTagName(v string) *SingleSendMailV2Request {
	s.TagName = &v
	return s
}

func (s *SingleSendMailV2Request) SetTextBody(v string) *SingleSendMailV2Request {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailV2Request) SetToAddress(v string) *SingleSendMailV2Request {
	s.ToAddress = &v
	return s
}

type SingleSendMailV2RequestHtmlBodyPlaceHolders struct {
	PlaceHolders map[string]*string `json:"PlaceHolders,omitempty" xml:"PlaceHolders,omitempty"`
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s SingleSendMailV2RequestHtmlBodyPlaceHolders) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailV2RequestHtmlBodyPlaceHolders) GoString() string {
	return s.String()
}

func (s *SingleSendMailV2RequestHtmlBodyPlaceHolders) SetPlaceHolders(v map[string]*string) *SingleSendMailV2RequestHtmlBodyPlaceHolders {
	s.PlaceHolders = v
	return s
}

func (s *SingleSendMailV2RequestHtmlBodyPlaceHolders) SetToAddress(v string) *SingleSendMailV2RequestHtmlBodyPlaceHolders {
	s.ToAddress = &v
	return s
}

type SingleSendMailV2ShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test***@example.net
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
	HtmlBody                   *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	HtmlBodyPlaceHoldersShrink *string `json:"HtmlBodyPlaceHolders,omitempty" xml:"HtmlBodyPlaceHolders,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// test2***@example.net
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
	// body
	TextBody *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s SingleSendMailV2ShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailV2ShrinkRequest) SetAccountName(v string) *SingleSendMailV2ShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetAddressType(v int32) *SingleSendMailV2ShrinkRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetClickTrace(v string) *SingleSendMailV2ShrinkRequest {
	s.ClickTrace = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetFromAlias(v string) *SingleSendMailV2ShrinkRequest {
	s.FromAlias = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetHtmlBody(v string) *SingleSendMailV2ShrinkRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetHtmlBodyPlaceHoldersShrink(v string) *SingleSendMailV2ShrinkRequest {
	s.HtmlBodyPlaceHoldersShrink = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetOwnerId(v int64) *SingleSendMailV2ShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetReplyAddress(v string) *SingleSendMailV2ShrinkRequest {
	s.ReplyAddress = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetReplyAddressAlias(v string) *SingleSendMailV2ShrinkRequest {
	s.ReplyAddressAlias = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetReplyToAddress(v bool) *SingleSendMailV2ShrinkRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetResourceOwnerAccount(v string) *SingleSendMailV2ShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetResourceOwnerId(v int64) *SingleSendMailV2ShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetSubject(v string) *SingleSendMailV2ShrinkRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetTagName(v string) *SingleSendMailV2ShrinkRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetTextBody(v string) *SingleSendMailV2ShrinkRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailV2ShrinkRequest) SetToAddress(v string) *SingleSendMailV2ShrinkRequest {
	s.ToAddress = &v
	return s
}

type SingleSendMailV2ResponseBody struct {
	// example:
	//
	// xxxxxx
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// example:
	//
	// 2D086F6-8F31-4658-84C1-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleSendMailV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *SingleSendMailV2ResponseBody) SetEnvId(v string) *SingleSendMailV2ResponseBody {
	s.EnvId = &v
	return s
}

func (s *SingleSendMailV2ResponseBody) SetRequestId(v string) *SingleSendMailV2ResponseBody {
	s.RequestId = &v
	return s
}

type SingleSendMailV2Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleSendMailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleSendMailV2Response) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailV2Response) GoString() string {
	return s.String()
}

func (s *SingleSendMailV2Response) SetHeaders(v map[string]*string) *SingleSendMailV2Response {
	s.Headers = v
	return s
}

func (s *SingleSendMailV2Response) SetStatusCode(v int32) *SingleSendMailV2Response {
	s.StatusCode = &v
	return s
}

func (s *SingleSendMailV2Response) SetBody(v *SingleSendMailV2ResponseBody) *SingleSendMailV2Response {
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
// 添加IP白名单
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
		Version:     tea.String("2017-06-22"),
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
// 添加IP白名单
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
// 校验域名Dns状态
//
// @param request - CheckDomainDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainDnsResponse
func (client *Client) CheckDomainDnsWithOptions(request *CheckDomainDnsRequest, runtime *util.RuntimeOptions) (_result *CheckDomainDnsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDomainDns"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDomainDnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验域名Dns状态
//
// @param request - CheckDomainDnsRequest
//
// @return CheckDomainDnsResponse
func (client *Client) CheckDomainDns(request *CheckDomainDnsRequest) (_result *CheckDomainDnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDomainDnsResponse{}
	_body, _err := client.CheckDomainDnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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

// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromType)) {
		query["FromType"] = request.FromType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SmsContent)) {
		query["SmsContent"] = request.SmsContent
	}

	if !tea.BoolValue(util.IsUnset(request.SmsType)) {
		query["SmsType"] = request.SmsType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateNickName)) {
		query["TemplateNickName"] = request.TemplateNickName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateSubject)) {
		query["TemplateSubject"] = request.TemplateSubject
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateText)) {
		query["TemplateText"] = request.TemplateText
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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

// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromType)) {
		query["FromType"] = request.FromType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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

// @param request - DescTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescTemplateResponse
func (client *Client) DescTemplateWithOptions(request *DescTemplateRequest, runtime *util.RuntimeOptions) (_result *DescTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromType)) {
		query["FromType"] = request.FromType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescTemplate"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescTemplateRequest
//
// @return DescTemplateResponse
func (client *Client) DescTemplate(request *DescTemplateRequest) (_result *DescTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescTemplateResponse{}
	_body, _err := client.DescTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAccountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountListResponse
func (client *Client) GetAccountListWithOptions(request *GetAccountListRequest, runtime *util.RuntimeOptions) (_result *GetAccountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.Total)) {
		query["Total"] = request.Total
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountList"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAccountListRequest
//
// @return GetAccountListResponse
func (client *Client) GetAccountList(request *GetAccountListRequest) (_result *GetAccountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountListResponse{}
	_body, _err := client.GetAccountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetMailAddressMsgCallBackUrl is deprecated
//
// Summary:
//
// 查询发信地址消息回调地址信息
//
// @param request - GetMailAddressMsgCallBackUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMailAddressMsgCallBackUrlResponse
// Deprecated
func (client *Client) GetMailAddressMsgCallBackUrlWithOptions(request *GetMailAddressMsgCallBackUrlRequest, runtime *util.RuntimeOptions) (_result *GetMailAddressMsgCallBackUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MailFrom)) {
		query["MailFrom"] = request.MailFrom
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
		Action:      tea.String("GetMailAddressMsgCallBackUrl"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMailAddressMsgCallBackUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetMailAddressMsgCallBackUrl is deprecated
//
// Summary:
//
// 查询发信地址消息回调地址信息
//
// @param request - GetMailAddressMsgCallBackUrlRequest
//
// @return GetMailAddressMsgCallBackUrlResponse
// Deprecated
func (client *Client) GetMailAddressMsgCallBackUrl(request *GetMailAddressMsgCallBackUrlRequest) (_result *GetMailAddressMsgCallBackUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMailAddressMsgCallBackUrlResponse{}
	_body, _err := client.GetMailAddressMsgCallBackUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取打开、点击等跟踪行为的统计结果
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
		Version:     tea.String("2017-06-22"),
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
// 获取打开、点击等跟踪行为的统计结果
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
		Version:     tea.String("2017-06-22"),
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

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPWByDomain"),
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
// 修改模板信息
//
// @param request - ModifyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplateWithOptions(request *ModifyTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromType)) {
		query["FromType"] = request.FromType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SmsContent)) {
		query["SmsContent"] = request.SmsContent
	}

	if !tea.BoolValue(util.IsUnset(request.SmsType)) {
		query["SmsType"] = request.SmsType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateNickName)) {
		query["TemplateNickName"] = request.TemplateNickName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateSubject)) {
		query["TemplateSubject"] = request.TemplateSubject
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateText)) {
		query["TemplateText"] = request.TemplateText
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTemplate"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改模板信息
//
// @param request - ModifyTemplateRequest
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplate(request *ModifyTemplateRequest) (_result *ModifyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.ModifyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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

// @param request - QueryTemplateByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTemplateByParamResponse
func (client *Client) QueryTemplateByParamWithOptions(request *QueryTemplateByParamRequest, runtime *util.RuntimeOptions) (_result *QueryTemplateByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromType)) {
		query["FromType"] = request.FromType
	}

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
		Action:      tea.String("QueryTemplateByParam"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTemplateByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTemplateByParamRequest
//
// @return QueryTemplateByParamResponse
func (client *Client) QueryTemplateByParam(request *QueryTemplateByParamRequest) (_result *QueryTemplateByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTemplateByParamResponse{}
	_body, _err := client.QueryTemplateByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
		Version:     tea.String("2017-06-22"),
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
// 极高发信专用API
//
// @param tmpReq - SingleSendMailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SingleSendMailV2Response
func (client *Client) SingleSendMailV2WithOptions(tmpReq *SingleSendMailV2Request, runtime *util.RuntimeOptions) (_result *SingleSendMailV2Response, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SingleSendMailV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HtmlBodyPlaceHolders)) {
		request.HtmlBodyPlaceHoldersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HtmlBodyPlaceHolders, tea.String("HtmlBodyPlaceHolders"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.HtmlBodyPlaceHoldersShrink)) {
		query["HtmlBodyPlaceHolders"] = request.HtmlBodyPlaceHoldersShrink
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleSendMailV2"),
		Version:     tea.String("2017-06-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SingleSendMailV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 极高发信专用API
//
// @param request - SingleSendMailV2Request
//
// @return SingleSendMailV2Response
func (client *Client) SingleSendMailV2(request *SingleSendMailV2Request) (_result *SingleSendMailV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SingleSendMailV2Response{}
	_body, _err := client.SingleSendMailV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
