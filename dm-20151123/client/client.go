// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddIpfilterRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s AddIpfilterRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpfilterRequest) GoString() string {
	return s.String()
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

func (s *AddIpfilterRequest) SetIpAddress(v string) *AddIpfilterRequest {
	s.IpAddress = &v
	return s
}

type AddIpfilterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpfilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpfilterResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpfilterResponseBody) SetRequestId(v string) *AddIpfilterResponseBody {
	s.RequestId = &v
	return s
}

type AddIpfilterResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddIpfilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddIpfilterResponse) SetBody(v *AddIpfilterResponseBody) *AddIpfilterResponse {
	s.Body = v
	return s
}

type ApproveMailTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s ApproveMailTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveMailTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApproveMailTemplateRequest) SetOwnerId(v int64) *ApproveMailTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ApproveMailTemplateRequest) SetResourceOwnerAccount(v string) *ApproveMailTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApproveMailTemplateRequest) SetResourceOwnerId(v int64) *ApproveMailTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApproveMailTemplateRequest) SetTemplateId(v int32) *ApproveMailTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ApproveMailTemplateRequest) SetFromType(v int32) *ApproveMailTemplateRequest {
	s.FromType = &v
	return s
}

type ApproveMailTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveMailTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveMailTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveMailTemplateResponseBody) SetRequestId(v string) *ApproveMailTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ApproveMailTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveMailTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveMailTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveMailTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApproveMailTemplateResponse) SetHeaders(v map[string]*string) *ApproveMailTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApproveMailTemplateResponse) SetBody(v *ApproveMailTemplateResponseBody) *ApproveMailTemplateResponse {
	s.Body = v
	return s
}

type ApproveReplyMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Ticket               *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveReplyMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ApproveReplyMailAddressResponse) SetBody(v *ApproveReplyMailAddressResponseBody) *ApproveReplyMailAddressResponse {
	s.Body = v
	return s
}

type ApproveSmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s ApproveSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApproveSmsTemplateRequest) SetOwnerId(v int64) *ApproveSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ApproveSmsTemplateRequest) SetResourceOwnerAccount(v string) *ApproveSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApproveSmsTemplateRequest) SetResourceOwnerId(v int64) *ApproveSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApproveSmsTemplateRequest) SetTemplateId(v int32) *ApproveSmsTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ApproveSmsTemplateRequest) SetFromType(v int32) *ApproveSmsTemplateRequest {
	s.FromType = &v
	return s
}

type ApproveSmsTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveSmsTemplateResponseBody) SetRequestId(v string) *ApproveSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ApproveSmsTemplateResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApproveSmsTemplateResponse) SetHeaders(v map[string]*string) *ApproveSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApproveSmsTemplateResponse) SetBody(v *ApproveSmsTemplateResponseBody) *ApproveSmsTemplateResponse {
	s.Body = v
	return s
}

type ApproveTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s ApproveTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApproveTemplateRequest) SetOwnerId(v int64) *ApproveTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ApproveTemplateRequest) SetResourceOwnerAccount(v string) *ApproveTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApproveTemplateRequest) SetResourceOwnerId(v int64) *ApproveTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApproveTemplateRequest) SetTemplateId(v int32) *ApproveTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ApproveTemplateRequest) SetFromType(v int32) *ApproveTemplateRequest {
	s.FromType = &v
	return s
}

type ApproveTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveTemplateResponseBody) SetRequestId(v string) *ApproveTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ApproveTemplateResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApproveTemplateResponse) SetHeaders(v map[string]*string) *ApproveTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApproveTemplateResponse) SetBody(v *ApproveTemplateResponseBody) *ApproveTemplateResponse {
	s.Body = v
	return s
}

type BatchSendMailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ReceiversName        *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	AddressType          *int32  `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias    *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	ClickTrace           *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
}

func (s BatchSendMailRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMailRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMailRequest) SetOwnerId(v int64) *BatchSendMailRequest {
	s.OwnerId = &v
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

func (s *BatchSendMailRequest) SetTemplateName(v string) *BatchSendMailRequest {
	s.TemplateName = &v
	return s
}

func (s *BatchSendMailRequest) SetAccountName(v string) *BatchSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *BatchSendMailRequest) SetReceiversName(v string) *BatchSendMailRequest {
	s.ReceiversName = &v
	return s
}

func (s *BatchSendMailRequest) SetAddressType(v int32) *BatchSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *BatchSendMailRequest) SetTagName(v string) *BatchSendMailRequest {
	s.TagName = &v
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

func (s *BatchSendMailRequest) SetClickTrace(v string) *BatchSendMailRequest {
	s.ClickTrace = &v
	return s
}

type BatchSendMailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s BatchSendMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMailResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMailResponseBody) SetRequestId(v string) *BatchSendMailResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendMailResponseBody) SetEnvId(v string) *BatchSendMailResponseBody {
	s.EnvId = &v
	return s
}

type BatchSendMailResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSendMailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchSendMailResponse) SetBody(v *BatchSendMailResponseBody) *BatchSendMailResponse {
	s.Body = v
	return s
}

type CheckDomainRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
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

func (s *CheckDomainRequest) SetDomainId(v int32) *CheckDomainRequest {
	s.DomainId = &v
	return s
}

type CheckDomainResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainStatus *int32  `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainResponseBody) SetDomainStatus(v int32) *CheckDomainResponseBody {
	s.DomainStatus = &v
	return s
}

type CheckDomainResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckDomainResponse) SetBody(v *CheckDomainResponseBody) *CheckDomainResponse {
	s.Body = v
	return s
}

type CheckInvalidAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ToAddress            *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s CheckInvalidAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInvalidAddressRequest) GoString() string {
	return s.String()
}

func (s *CheckInvalidAddressRequest) SetOwnerId(v int64) *CheckInvalidAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckInvalidAddressRequest) SetResourceOwnerAccount(v string) *CheckInvalidAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckInvalidAddressRequest) SetResourceOwnerId(v int64) *CheckInvalidAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckInvalidAddressRequest) SetToAddress(v string) *CheckInvalidAddressRequest {
	s.ToAddress = &v
	return s
}

type CheckInvalidAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInvalidAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInvalidAddressResponseBody) SetRequestId(v string) *CheckInvalidAddressResponseBody {
	s.RequestId = &v
	return s
}

type CheckInvalidAddressResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckInvalidAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInvalidAddressResponse) GoString() string {
	return s.String()
}

func (s *CheckInvalidAddressResponse) SetHeaders(v map[string]*string) *CheckInvalidAddressResponse {
	s.Headers = v
	return s
}

func (s *CheckInvalidAddressResponse) SetBody(v *CheckInvalidAddressResponseBody) *CheckInvalidAddressResponse {
	s.Body = v
	return s
}

type CheckReplyToMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	MailAddressId        *int32  `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
}

func (s CheckReplyToMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckReplyToMailAddressRequest) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressRequest) SetOwnerId(v int64) *CheckReplyToMailAddressRequest {
	s.OwnerId = &v
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

func (s *CheckReplyToMailAddressRequest) SetLang(v string) *CheckReplyToMailAddressRequest {
	s.Lang = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetRegion(v string) *CheckReplyToMailAddressRequest {
	s.Region = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetMailAddressId(v int32) *CheckReplyToMailAddressRequest {
	s.MailAddressId = &v
	return s
}

type CheckReplyToMailAddressResponseBody struct {
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckReplyToMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckReplyToMailAddressResponse) SetBody(v *CheckReplyToMailAddressResponseBody) *CheckReplyToMailAddressResponse {
	s.Body = v
	return s
}

type CreateDayuRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccountType          *int32  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s CreateDayuRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDayuRequest) GoString() string {
	return s.String()
}

func (s *CreateDayuRequest) SetOwnerId(v int64) *CreateDayuRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDayuRequest) SetResourceOwnerAccount(v string) *CreateDayuRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDayuRequest) SetResourceOwnerId(v int64) *CreateDayuRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDayuRequest) SetAccountType(v int32) *CreateDayuRequest {
	s.AccountType = &v
	return s
}

type CreateDayuResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDayuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDayuResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDayuResponseBody) SetRequestId(v string) *CreateDayuResponseBody {
	s.RequestId = &v
	return s
}

type CreateDayuResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDayuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDayuResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDayuResponse) GoString() string {
	return s.String()
}

func (s *CreateDayuResponse) SetHeaders(v map[string]*string) *CreateDayuResponse {
	s.Headers = v
	return s
}

func (s *CreateDayuResponse) SetBody(v *CreateDayuResponseBody) *CreateDayuResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
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

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	Sendtype             *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s CreateMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMailAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateMailAddressRequest) SetOwnerId(v int64) *CreateMailAddressRequest {
	s.OwnerId = &v
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

func (s *CreateMailAddressRequest) SetAccountName(v string) *CreateMailAddressRequest {
	s.AccountName = &v
	return s
}

func (s *CreateMailAddressRequest) SetReplyAddress(v string) *CreateMailAddressRequest {
	s.ReplyAddress = &v
	return s
}

func (s *CreateMailAddressRequest) SetSendtype(v string) *CreateMailAddressRequest {
	s.Sendtype = &v
	return s
}

type CreateMailAddressResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
}

func (s CreateMailAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMailAddressResponseBody) SetRequestId(v string) *CreateMailAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMailAddressResponseBody) SetMailAddressId(v string) *CreateMailAddressResponseBody {
	s.MailAddressId = &v
	return s
}

type CreateMailAddressResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateMailAddressResponse) SetBody(v *CreateMailAddressResponseBody) *CreateMailAddressResponse {
	s.Body = v
	return s
}

type CreateReceiverRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ReceiversName        *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ReceiversAlias       *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	Desc                 *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s CreateReceiverRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReceiverRequest) GoString() string {
	return s.String()
}

func (s *CreateReceiverRequest) SetOwnerId(v int64) *CreateReceiverRequest {
	s.OwnerId = &v
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

func (s *CreateReceiverRequest) SetReceiversName(v string) *CreateReceiverRequest {
	s.ReceiversName = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiversAlias(v string) *CreateReceiverRequest {
	s.ReceiversAlias = &v
	return s
}

func (s *CreateReceiverRequest) SetDesc(v string) *CreateReceiverRequest {
	s.Desc = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateReceiverResponse) SetBody(v *CreateReceiverResponseBody) *CreateReceiverResponse {
	s.Body = v
	return s
}

type CreateSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SignType             *int32  `json:"SignType,omitempty" xml:"SignType,omitempty"`
	FileNames            *string `json:"FileNames,omitempty" xml:"FileNames,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s CreateSignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSignRequest) SetOwnerId(v int64) *CreateSignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSignRequest) SetResourceOwnerAccount(v string) *CreateSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSignRequest) SetResourceOwnerId(v int64) *CreateSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSignRequest) SetSignName(v string) *CreateSignRequest {
	s.SignName = &v
	return s
}

func (s *CreateSignRequest) SetRemark(v string) *CreateSignRequest {
	s.Remark = &v
	return s
}

func (s *CreateSignRequest) SetSignType(v int32) *CreateSignRequest {
	s.SignType = &v
	return s
}

func (s *CreateSignRequest) SetFileNames(v string) *CreateSignRequest {
	s.FileNames = &v
	return s
}

func (s *CreateSignRequest) SetFromType(v int32) *CreateSignRequest {
	s.FromType = &v
	return s
}

type CreateSignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignResponseBody) SetRequestId(v string) *CreateSignResponseBody {
	s.RequestId = &v
	return s
}

type CreateSignResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSignResponse) GoString() string {
	return s.String()
}

func (s *CreateSignResponse) SetHeaders(v map[string]*string) *CreateSignResponse {
	s.Headers = v
	return s
}

func (s *CreateSignResponse) SetBody(v *CreateSignResponseBody) *CreateSignResponse {
	s.Body = v
	return s
}

type CreateTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTagResponse) SetBody(v *CreateTagResponseBody) *CreateTagResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateType         *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateSubject      *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	TemplateNickName     *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateText         *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	SmsType              *int32  `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	SmsContent           *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetOwnerId(v int64) *CreateTemplateRequest {
	s.OwnerId = &v
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

func (s *CreateTemplateRequest) SetTemplateType(v int32) *CreateTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateSubject(v string) *CreateTemplateRequest {
	s.TemplateSubject = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateNickName(v string) *CreateTemplateRequest {
	s.TemplateNickName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateText(v string) *CreateTemplateRequest {
	s.TemplateText = &v
	return s
}

func (s *CreateTemplateRequest) SetSmsType(v int32) *CreateTemplateRequest {
	s.SmsType = &v
	return s
}

func (s *CreateTemplateRequest) SetSmsContent(v string) *CreateTemplateRequest {
	s.SmsContent = &v
	return s
}

func (s *CreateTemplateRequest) SetRemark(v string) *CreateTemplateRequest {
	s.Remark = &v
	return s
}

func (s *CreateTemplateRequest) SetFromType(v int32) *CreateTemplateRequest {
	s.FromType = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
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

func (s *DeleteDomainRequest) SetDomainId(v int32) *DeleteDomainRequest {
	s.DomainId = &v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteInvalidAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ToAddress            *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInvalidAddressResponse) SetBody(v *DeleteInvalidAddressResponseBody) *DeleteInvalidAddressResponse {
	s.Body = v
	return s
}

type DeleteIpfilterByEdmIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteIpfilterByEdmIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpfilterByEdmIdRequest) GoString() string {
	return s.String()
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

func (s *DeleteIpfilterByEdmIdRequest) SetFromType(v int32) *DeleteIpfilterByEdmIdRequest {
	s.FromType = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetId(v string) *DeleteIpfilterByEdmIdRequest {
	s.Id = &v
	return s
}

type DeleteIpfilterByEdmIdResponseBody struct {
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIpfilterByEdmIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteIpfilterByEdmIdResponse) SetBody(v *DeleteIpfilterByEdmIdResponseBody) *DeleteIpfilterByEdmIdResponse {
	s.Body = v
	return s
}

type DeleteMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	MailAddressId        *int32  `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
}

func (s DeleteMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailAddressRequest) GoString() string {
	return s.String()
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

func (s *DeleteMailAddressRequest) SetMailAddressId(v int32) *DeleteMailAddressRequest {
	s.MailAddressId = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteMailAddressResponse) SetBody(v *DeleteMailAddressResponseBody) *DeleteMailAddressResponse {
	s.Body = v
	return s
}

type DeleteReceiverRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
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

func (s *DeleteReceiverRequest) SetResourceOwnerAccount(v string) *DeleteReceiverRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReceiverRequest) SetResourceOwnerId(v int64) *DeleteReceiverRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteReceiverRequest) SetReceiverId(v string) *DeleteReceiverRequest {
	s.ReceiverId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteReceiverResponse) SetBody(v *DeleteReceiverResponseBody) *DeleteReceiverResponse {
	s.Body = v
	return s
}

type DeleteReceiverDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	Email                *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s DeleteReceiverDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailRequest) SetOwnerId(v int64) *DeleteReceiverDetailRequest {
	s.OwnerId = &v
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

func (s *DeleteReceiverDetailRequest) SetReceiverId(v string) *DeleteReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetEmail(v string) *DeleteReceiverDetailRequest {
	s.Email = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteReceiverDetailResponse) SetBody(v *DeleteReceiverDetailResponseBody) *DeleteReceiverDetailResponse {
	s.Body = v
	return s
}

type DeleteSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignId               *int64  `json:"SignId,omitempty" xml:"SignId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s DeleteSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignRequest) GoString() string {
	return s.String()
}

func (s *DeleteSignRequest) SetOwnerId(v int64) *DeleteSignRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSignRequest) SetResourceOwnerAccount(v string) *DeleteSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSignRequest) SetResourceOwnerId(v int64) *DeleteSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSignRequest) SetSignId(v int64) *DeleteSignRequest {
	s.SignId = &v
	return s
}

func (s *DeleteSignRequest) SetFromType(v int32) *DeleteSignRequest {
	s.FromType = &v
	return s
}

type DeleteSignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSignResponseBody) SetRequestId(v string) *DeleteSignResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSignResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignResponse) GoString() string {
	return s.String()
}

func (s *DeleteSignResponse) SetHeaders(v map[string]*string) *DeleteSignResponse {
	s.Headers = v
	return s
}

func (s *DeleteSignResponse) SetBody(v *DeleteSignResponseBody) *DeleteSignResponse {
	s.Body = v
	return s
}

type DeleteTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagId                *int32  `json:"TagId,omitempty" xml:"TagId,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTagResponse) SetBody(v *DeleteTagResponseBody) *DeleteTagResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
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

func (s *DeleteTemplateRequest) SetFromType(v int32) *DeleteTemplateRequest {
	s.FromType = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DayuStatus    *int32  `json:"DayuStatus,omitempty" xml:"DayuStatus,omitempty"`
	SmsRecord     *int32  `json:"SmsRecord,omitempty" xml:"SmsRecord,omitempty"`
	MonthQuota    *int32  `json:"MonthQuota,omitempty" xml:"MonthQuota,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Receivers     *int32  `json:"Receivers,omitempty" xml:"Receivers,omitempty"`
	SmsTemplates  *int32  `json:"SmsTemplates,omitempty" xml:"SmsTemplates,omitempty"`
	Templates     *int32  `json:"Templates,omitempty" xml:"Templates,omitempty"`
	DailyQuota    *int32  `json:"DailyQuota,omitempty" xml:"DailyQuota,omitempty"`
	UserStatus    *int32  `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	Domains       *int32  `json:"Domains,omitempty" xml:"Domains,omitempty"`
	QuotaLevel    *int32  `json:"QuotaLevel,omitempty" xml:"QuotaLevel,omitempty"`
	SmsSign       *int32  `json:"SmsSign,omitempty" xml:"SmsSign,omitempty"`
	MaxQuotaLevel *int32  `json:"MaxQuotaLevel,omitempty" xml:"MaxQuotaLevel,omitempty"`
	EnableTimes   *int32  `json:"EnableTimes,omitempty" xml:"EnableTimes,omitempty"`
	Tags          *int32  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	MailAddresses *int32  `json:"MailAddresses,omitempty" xml:"MailAddresses,omitempty"`
}

func (s DescAccountSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryResponseBody) SetDayuStatus(v int32) *DescAccountSummaryResponseBody {
	s.DayuStatus = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsRecord(v int32) *DescAccountSummaryResponseBody {
	s.SmsRecord = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMonthQuota(v int32) *DescAccountSummaryResponseBody {
	s.MonthQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetRequestId(v string) *DescAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetReceivers(v int32) *DescAccountSummaryResponseBody {
	s.Receivers = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsTemplates(v int32) *DescAccountSummaryResponseBody {
	s.SmsTemplates = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetTemplates(v int32) *DescAccountSummaryResponseBody {
	s.Templates = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDailyQuota(v int32) *DescAccountSummaryResponseBody {
	s.DailyQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetUserStatus(v int32) *DescAccountSummaryResponseBody {
	s.UserStatus = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDomains(v int32) *DescAccountSummaryResponseBody {
	s.Domains = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetQuotaLevel(v int32) *DescAccountSummaryResponseBody {
	s.QuotaLevel = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsSign(v int32) *DescAccountSummaryResponseBody {
	s.SmsSign = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMaxQuotaLevel(v int32) *DescAccountSummaryResponseBody {
	s.MaxQuotaLevel = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetEnableTimes(v int32) *DescAccountSummaryResponseBody {
	s.EnableTimes = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetTags(v int32) *DescAccountSummaryResponseBody {
	s.Tags = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMailAddresses(v int32) *DescAccountSummaryResponseBody {
	s.MailAddresses = &v
	return s
}

type DescAccountSummaryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescAccountSummaryResponse) SetBody(v *DescAccountSummaryResponseBody) *DescAccountSummaryResponse {
	s.Body = v
	return s
}

type DescAccountSummary2Request struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s DescAccountSummary2Request) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummary2Request) GoString() string {
	return s.String()
}

func (s *DescAccountSummary2Request) SetOwnerId(v int64) *DescAccountSummary2Request {
	s.OwnerId = &v
	return s
}

func (s *DescAccountSummary2Request) SetResourceOwnerAccount(v string) *DescAccountSummary2Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescAccountSummary2Request) SetResourceOwnerId(v int64) *DescAccountSummary2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescAccountSummary2Request) SetFromType(v int32) *DescAccountSummary2Request {
	s.FromType = &v
	return s
}

type DescAccountSummary2ResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MnsForceMigrating *int32  `json:"MnsForceMigrating,omitempty" xml:"MnsForceMigrating,omitempty"`
	MnsBag            *int32  `json:"MnsBag,omitempty" xml:"MnsBag,omitempty"`
	MnsMigrating      *int32  `json:"MnsMigrating,omitempty" xml:"MnsMigrating,omitempty"`
}

func (s DescAccountSummary2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummary2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescAccountSummary2ResponseBody) SetRequestId(v string) *DescAccountSummary2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescAccountSummary2ResponseBody) SetMnsForceMigrating(v int32) *DescAccountSummary2ResponseBody {
	s.MnsForceMigrating = &v
	return s
}

func (s *DescAccountSummary2ResponseBody) SetMnsBag(v int32) *DescAccountSummary2ResponseBody {
	s.MnsBag = &v
	return s
}

func (s *DescAccountSummary2ResponseBody) SetMnsMigrating(v int32) *DescAccountSummary2ResponseBody {
	s.MnsMigrating = &v
	return s
}

type DescAccountSummary2Response struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescAccountSummary2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescAccountSummary2Response) String() string {
	return tea.Prettify(s)
}

func (s DescAccountSummary2Response) GoString() string {
	return s.String()
}

func (s *DescAccountSummary2Response) SetHeaders(v map[string]*string) *DescAccountSummary2Response {
	s.Headers = v
	return s
}

func (s *DescAccountSummary2Response) SetBody(v *DescAccountSummary2ResponseBody) *DescAccountSummary2Response {
	s.Body = v
	return s
}

type DescDomainRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s DescDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescDomainRequest) GoString() string {
	return s.String()
}

func (s *DescDomainRequest) SetOwnerId(v int64) *DescDomainRequest {
	s.OwnerId = &v
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

func (s *DescDomainRequest) SetDomainId(v int32) *DescDomainRequest {
	s.DomainId = &v
	return s
}

type DescDomainResponseBody struct {
	SpfRecord          *string `json:"SpfRecord,omitempty" xml:"SpfRecord,omitempty"`
	SpfAuthStatus      *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	CnameAuthStatus    *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DnsMx              *string `json:"DnsMx,omitempty" xml:"DnsMx,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CnameRecord        *string `json:"CnameRecord,omitempty" xml:"CnameRecord,omitempty"`
	DnsTxt             *string `json:"DnsTxt,omitempty" xml:"DnsTxt,omitempty"`
	CnameConfirmStatus *string `json:"CnameConfirmStatus,omitempty" xml:"CnameConfirmStatus,omitempty"`
	IcpStatus          *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	DefaultDomain      *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	DnsSpf             *string `json:"DnsSpf,omitempty" xml:"DnsSpf,omitempty"`
	MxRecord           *string `json:"MxRecord,omitempty" xml:"MxRecord,omitempty"`
	DomainId           *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	DomainType         *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	MxAuthStatus       *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	TlDomainName       *string `json:"TlDomainName,omitempty" xml:"TlDomainName,omitempty"`
	TracefRecord       *string `json:"TracefRecord,omitempty" xml:"TracefRecord,omitempty"`
	DomainStatus       *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
}

func (s DescDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescDomainResponseBody) SetSpfRecord(v string) *DescDomainResponseBody {
	s.SpfRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfAuthStatus(v string) *DescDomainResponseBody {
	s.SpfAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameAuthStatus(v string) *DescDomainResponseBody {
	s.CnameAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetRequestId(v string) *DescDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainName(v string) *DescDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsMx(v string) *DescDomainResponseBody {
	s.DnsMx = &v
	return s
}

func (s *DescDomainResponseBody) SetCreateTime(v string) *DescDomainResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameRecord(v string) *DescDomainResponseBody {
	s.CnameRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsTxt(v string) *DescDomainResponseBody {
	s.DnsTxt = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameConfirmStatus(v string) *DescDomainResponseBody {
	s.CnameConfirmStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetIcpStatus(v string) *DescDomainResponseBody {
	s.IcpStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDefaultDomain(v string) *DescDomainResponseBody {
	s.DefaultDomain = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsSpf(v string) *DescDomainResponseBody {
	s.DnsSpf = &v
	return s
}

func (s *DescDomainResponseBody) SetMxRecord(v string) *DescDomainResponseBody {
	s.MxRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainId(v string) *DescDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainType(v string) *DescDomainResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescDomainResponseBody) SetMxAuthStatus(v string) *DescDomainResponseBody {
	s.MxAuthStatus = &v
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

func (s *DescDomainResponseBody) SetDomainStatus(v string) *DescDomainResponseBody {
	s.DomainStatus = &v
	return s
}

type DescDomainResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescDomainResponse) SetBody(v *DescDomainResponseBody) *DescDomainResponse {
	s.Body = v
	return s
}

type DescTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s DescTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescTemplateRequest) GoString() string {
	return s.String()
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

func (s *DescTemplateRequest) SetFromType(v int32) *DescTemplateRequest {
	s.FromType = &v
	return s
}

type DescTemplateResponseBody struct {
	SmsType          *string `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateText     *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	SmsContent       *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateType     *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateSubject  *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	Remark           *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TemplateStatus   *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
}

func (s DescTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescTemplateResponseBody) SetSmsType(v string) *DescTemplateResponseBody {
	s.SmsType = &v
	return s
}

func (s *DescTemplateResponseBody) SetRequestId(v string) *DescTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescTemplateResponseBody) SetCreateTime(v string) *DescTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateText(v string) *DescTemplateResponseBody {
	s.TemplateText = &v
	return s
}

func (s *DescTemplateResponseBody) SetSmsContent(v string) *DescTemplateResponseBody {
	s.SmsContent = &v
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

func (s *DescTemplateResponseBody) SetTemplateType(v string) *DescTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateSubject(v string) *DescTemplateResponseBody {
	s.TemplateSubject = &v
	return s
}

func (s *DescTemplateResponseBody) SetRemark(v string) *DescTemplateResponseBody {
	s.Remark = &v
	return s
}

func (s *DescTemplateResponseBody) SetTemplateStatus(v string) *DescTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

type DescTemplateResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescTemplateResponse) SetBody(v *DescTemplateResponseBody) *DescTemplateResponse {
	s.Body = v
	return s
}

type EnableAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAccountRequest) GoString() string {
	return s.String()
}

func (s *EnableAccountRequest) SetOwnerId(v int64) *EnableAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableAccountRequest) SetResourceOwnerAccount(v string) *EnableAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableAccountRequest) SetResourceOwnerId(v int64) *EnableAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

type EnableAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAccountResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAccountResponseBody) SetRequestId(v string) *EnableAccountResponseBody {
	s.RequestId = &v
	return s
}

type EnableAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAccountResponse) GoString() string {
	return s.String()
}

func (s *EnableAccountResponse) SetHeaders(v map[string]*string) *EnableAccountResponse {
	s.Headers = v
	return s
}

func (s *EnableAccountResponse) SetBody(v *EnableAccountResponseBody) *EnableAccountResponse {
	s.Body = v
	return s
}

type GetAccountListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Total                *string `json:"Total,omitempty" xml:"Total,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetAccountListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListRequest) GoString() string {
	return s.String()
}

func (s *GetAccountListRequest) SetOwnerId(v int64) *GetAccountListRequest {
	s.OwnerId = &v
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

func (s *GetAccountListRequest) SetOffset(v string) *GetAccountListRequest {
	s.Offset = &v
	return s
}

func (s *GetAccountListRequest) SetPageSize(v string) *GetAccountListRequest {
	s.PageSize = &v
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

func (s *GetAccountListRequest) SetPageNumber(v string) *GetAccountListRequest {
	s.PageNumber = &v
	return s
}

type GetAccountListResponseBody struct {
	PageSize  *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAccountListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Total     *int32                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PageNo    *int32                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s GetAccountListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountListResponseBody) SetPageSize(v int32) *GetAccountListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAccountListResponseBody) SetRequestId(v string) *GetAccountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountListResponseBody) SetData(v *GetAccountListResponseBodyData) *GetAccountListResponseBody {
	s.Data = v
	return s
}

func (s *GetAccountListResponseBody) SetTotal(v int32) *GetAccountListResponseBody {
	s.Total = &v
	return s
}

func (s *GetAccountListResponseBody) SetPageNo(v int32) *GetAccountListResponseBody {
	s.PageNo = &v
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
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetAccountListResponseBodyDataAccountNotificationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountListResponseBodyDataAccountNotificationInfo) GoString() string {
	return s.String()
}

func (s *GetAccountListResponseBodyDataAccountNotificationInfo) SetStatus(v string) *GetAccountListResponseBodyDataAccountNotificationInfo {
	s.Status = &v
	return s
}

func (s *GetAccountListResponseBodyDataAccountNotificationInfo) SetUpdateTime(v string) *GetAccountListResponseBodyDataAccountNotificationInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetAccountListResponseBodyDataAccountNotificationInfo) SetRegion(v string) *GetAccountListResponseBodyDataAccountNotificationInfo {
	s.Region = &v
	return s
}

type GetAccountListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccountListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAccountListResponse) SetBody(v *GetAccountListResponseBody) *GetAccountListResponse {
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
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *GetIpfilterListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetIpfilterListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBody) SetTotalCount(v int32) *GetIpfilterListResponseBody {
	s.TotalCount = &v
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

func (s *GetIpfilterListResponseBody) SetData(v *GetIpfilterListResponseBodyData) *GetIpfilterListResponseBody {
	s.Data = v
	return s
}

func (s *GetIpfilterListResponseBody) SetPageNumber(v int32) *GetIpfilterListResponseBody {
	s.PageNumber = &v
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
	IpAddress  *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetIpfilterListResponseBodyDataIpfilters) String() string {
	return tea.Prettify(s)
}

func (s GetIpfilterListResponseBodyDataIpfilters) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetIpAddress(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.IpAddress = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetCreateTime(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.CreateTime = &v
	return s
}

func (s *GetIpfilterListResponseBodyDataIpfilters) SetId(v string) *GetIpfilterListResponseBodyDataIpfilters {
	s.Id = &v
	return s
}

type GetIpfilterListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIpfilterListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetIpfilterListResponse) SetBody(v *GetIpfilterListResponseBody) *GetIpfilterListResponse {
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpProtection *string `json:"IpProtection,omitempty" xml:"IpProtection,omitempty"`
}

func (s GetIpProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpProtectionResponseBody) SetRequestId(v string) *GetIpProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpProtectionResponseBody) SetIpProtection(v string) *GetIpProtectionResponseBody {
	s.IpProtection = &v
	return s
}

type GetIpProtectionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIpProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetIpProtectionResponse) SetBody(v *GetIpProtectionResponseBody) *GetIpProtectionResponse {
	s.Body = v
	return s
}

type GetMailAddressMsgCallBackUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	MailFrom             *string `json:"MailFrom,omitempty" xml:"MailFrom,omitempty"`
}

func (s GetMailAddressMsgCallBackUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMailAddressMsgCallBackUrlRequest) GoString() string {
	return s.String()
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

func (s *GetMailAddressMsgCallBackUrlRequest) SetMailFrom(v string) *GetMailAddressMsgCallBackUrlRequest {
	s.MailFrom = &v
	return s
}

type GetMailAddressMsgCallBackUrlResponseBody struct {
	NotifyUrlStatus *int32  `json:"NotifyUrlStatus,omitempty" xml:"NotifyUrlStatus,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NotifyUrl       *int32  `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s GetMailAddressMsgCallBackUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMailAddressMsgCallBackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMailAddressMsgCallBackUrlResponseBody) SetNotifyUrlStatus(v int32) *GetMailAddressMsgCallBackUrlResponseBody {
	s.NotifyUrlStatus = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlResponseBody) SetRequestId(v string) *GetMailAddressMsgCallBackUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMailAddressMsgCallBackUrlResponseBody) SetNotifyUrl(v int32) *GetMailAddressMsgCallBackUrlResponseBody {
	s.NotifyUrl = &v
	return s
}

type GetMailAddressMsgCallBackUrlResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMailAddressMsgCallBackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMailAddressMsgCallBackUrlResponse) SetBody(v *GetMailAddressMsgCallBackUrlResponseBody) *GetMailAddressMsgCallBackUrlResponse {
	s.Body = v
	return s
}

type GetRegionListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Total                *string `json:"Total,omitempty" xml:"Total,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetRegionListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListRequest) GoString() string {
	return s.String()
}

func (s *GetRegionListRequest) SetOwnerId(v int64) *GetRegionListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetRegionListRequest) SetResourceOwnerAccount(v string) *GetRegionListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetRegionListRequest) SetResourceOwnerId(v int64) *GetRegionListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetRegionListRequest) SetTotal(v string) *GetRegionListRequest {
	s.Total = &v
	return s
}

func (s *GetRegionListRequest) SetOffset(v string) *GetRegionListRequest {
	s.Offset = &v
	return s
}

func (s *GetRegionListRequest) SetPageSize(v string) *GetRegionListRequest {
	s.PageSize = &v
	return s
}

func (s *GetRegionListRequest) SetOffsetCreateTime(v string) *GetRegionListRequest {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetRegionListRequest) SetOffsetCreateTimeDesc(v string) *GetRegionListRequest {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetRegionListRequest) SetPageNumber(v string) *GetRegionListRequest {
	s.PageNumber = &v
	return s
}

type GetRegionListResponseBody struct {
	PageSize  *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetRegionListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Total     *int32                         `json:"Total,omitempty" xml:"Total,omitempty"`
	PageNo    *int32                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s GetRegionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionListResponseBody) SetPageSize(v int32) *GetRegionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetRegionListResponseBody) SetRequestId(v string) *GetRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionListResponseBody) SetData(v *GetRegionListResponseBodyData) *GetRegionListResponseBody {
	s.Data = v
	return s
}

func (s *GetRegionListResponseBody) SetTotal(v int32) *GetRegionListResponseBody {
	s.Total = &v
	return s
}

func (s *GetRegionListResponseBody) SetPageNo(v int32) *GetRegionListResponseBody {
	s.PageNo = &v
	return s
}

type GetRegionListResponseBodyData struct {
	RegionList []*GetRegionListResponseBodyDataRegionList `json:"regionList,omitempty" xml:"regionList,omitempty" type:"Repeated"`
}

func (s GetRegionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRegionListResponseBodyData) SetRegionList(v []*GetRegionListResponseBodyDataRegionList) *GetRegionListResponseBodyData {
	s.RegionList = v
	return s
}

type GetRegionListResponseBodyDataRegionList struct {
	RegionDesc *string `json:"RegionDesc,omitempty" xml:"RegionDesc,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetRegionListResponseBodyDataRegionList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponseBodyDataRegionList) GoString() string {
	return s.String()
}

func (s *GetRegionListResponseBodyDataRegionList) SetRegionDesc(v string) *GetRegionListResponseBodyDataRegionList {
	s.RegionDesc = &v
	return s
}

func (s *GetRegionListResponseBodyDataRegionList) SetRegion(v string) *GetRegionListResponseBodyDataRegionList {
	s.Region = &v
	return s
}

type GetRegionListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegionListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponse) GoString() string {
	return s.String()
}

func (s *GetRegionListResponse) SetHeaders(v map[string]*string) *GetRegionListResponse {
	s.Headers = v
	return s
}

func (s *GetRegionListResponse) SetBody(v *GetRegionListResponseBody) *GetRegionListResponse {
	s.Body = v
	return s
}

type GetSenderAddressListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Total                *string `json:"Total,omitempty" xml:"Total,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo               *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	NotifyUrl            *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s GetSenderAddressListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSenderAddressListRequest) GoString() string {
	return s.String()
}

func (s *GetSenderAddressListRequest) SetOwnerId(v int64) *GetSenderAddressListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSenderAddressListRequest) SetResourceOwnerAccount(v string) *GetSenderAddressListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSenderAddressListRequest) SetResourceOwnerId(v int64) *GetSenderAddressListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSenderAddressListRequest) SetTotal(v string) *GetSenderAddressListRequest {
	s.Total = &v
	return s
}

func (s *GetSenderAddressListRequest) SetOffset(v string) *GetSenderAddressListRequest {
	s.Offset = &v
	return s
}

func (s *GetSenderAddressListRequest) SetPageSize(v string) *GetSenderAddressListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSenderAddressListRequest) SetPageNo(v string) *GetSenderAddressListRequest {
	s.PageNo = &v
	return s
}

func (s *GetSenderAddressListRequest) SetKeyword(v string) *GetSenderAddressListRequest {
	s.Keyword = &v
	return s
}

func (s *GetSenderAddressListRequest) SetNotifyUrl(v string) *GetSenderAddressListRequest {
	s.NotifyUrl = &v
	return s
}

type GetSenderAddressListResponseBody struct {
	PageSize  *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetSenderAddressListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Total     *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PageNo    *int32                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s GetSenderAddressListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSenderAddressListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSenderAddressListResponseBody) SetPageSize(v int32) *GetSenderAddressListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSenderAddressListResponseBody) SetRequestId(v string) *GetSenderAddressListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSenderAddressListResponseBody) SetData(v *GetSenderAddressListResponseBodyData) *GetSenderAddressListResponseBody {
	s.Data = v
	return s
}

func (s *GetSenderAddressListResponseBody) SetTotal(v int32) *GetSenderAddressListResponseBody {
	s.Total = &v
	return s
}

func (s *GetSenderAddressListResponseBody) SetPageNo(v int32) *GetSenderAddressListResponseBody {
	s.PageNo = &v
	return s
}

type GetSenderAddressListResponseBodyData struct {
	SenderAddressNotificationInfo []*GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo `json:"senderAddressNotificationInfo,omitempty" xml:"senderAddressNotificationInfo,omitempty" type:"Repeated"`
}

func (s GetSenderAddressListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSenderAddressListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSenderAddressListResponseBodyData) SetSenderAddressNotificationInfo(v []*GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) *GetSenderAddressListResponseBodyData {
	s.SenderAddressNotificationInfo = v
	return s
}

type GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	SenderAddress   *string `json:"SenderAddress,omitempty" xml:"SenderAddress,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SenderAddressId *string `json:"SenderAddressId,omitempty" xml:"SenderAddressId,omitempty"`
}

func (s GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) GoString() string {
	return s.String()
}

func (s *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) SetStatus(v string) *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo {
	s.Status = &v
	return s
}

func (s *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) SetUpdateTime(v string) *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) SetSenderAddress(v string) *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo {
	s.SenderAddress = &v
	return s
}

func (s *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) SetRegion(v string) *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo {
	s.Region = &v
	return s
}

func (s *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo) SetSenderAddressId(v string) *GetSenderAddressListResponseBodyDataSenderAddressNotificationInfo {
	s.SenderAddressId = &v
	return s
}

type GetSenderAddressListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSenderAddressListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSenderAddressListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSenderAddressListResponse) GoString() string {
	return s.String()
}

func (s *GetSenderAddressListResponse) SetHeaders(v map[string]*string) *GetSenderAddressListResponse {
	s.Headers = v
	return s
}

func (s *GetSenderAddressListResponse) SetBody(v *GetSenderAddressListResponseBody) *GetSenderAddressListResponse {
	s.Body = v
	return s
}

type GetTrackListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Total                *string `json:"Total,omitempty" xml:"Total,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetTrackListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListRequest) GoString() string {
	return s.String()
}

func (s *GetTrackListRequest) SetOwnerId(v int64) *GetTrackListRequest {
	s.OwnerId = &v
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

func (s *GetTrackListRequest) SetEndTime(v string) *GetTrackListRequest {
	s.EndTime = &v
	return s
}

func (s *GetTrackListRequest) SetTotal(v string) *GetTrackListRequest {
	s.Total = &v
	return s
}

func (s *GetTrackListRequest) SetOffset(v string) *GetTrackListRequest {
	s.Offset = &v
	return s
}

func (s *GetTrackListRequest) SetPageSize(v string) *GetTrackListRequest {
	s.PageSize = &v
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

func (s *GetTrackListRequest) SetPageNumber(v string) *GetTrackListRequest {
	s.PageNumber = &v
	return s
}

type GetTrackListResponseBody struct {
	OffsetCreateTime     *string                       `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	RequestId            *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize             *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data                 *GetTrackListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Total                *int32                        `json:"Total,omitempty" xml:"Total,omitempty"`
	PageNo               *int32                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	OffsetCreateTimeDesc *string                       `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
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

func (s *GetTrackListResponseBody) SetRequestId(v string) *GetTrackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrackListResponseBody) SetPageSize(v int32) *GetTrackListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTrackListResponseBody) SetData(v *GetTrackListResponseBodyData) *GetTrackListResponseBody {
	s.Data = v
	return s
}

func (s *GetTrackListResponseBody) SetTotal(v int32) *GetTrackListResponseBody {
	s.Total = &v
	return s
}

func (s *GetTrackListResponseBody) SetPageNo(v int32) *GetTrackListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetTrackListResponseBody) SetOffsetCreateTimeDesc(v string) *GetTrackListResponseBody {
	s.OffsetCreateTimeDesc = &v
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
	RcptClickRate        *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	RcptUniqueOpenCount  *string `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	RcptClickCount       *string `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	RcptUniqueClickCount *string `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RcptUniqueOpenRate   *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	RcptUniqueClickRate  *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	TotalNumber          *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	RcptOpenRate         *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	RcptOpenCount        *string `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
}

func (s GetTrackListResponseBodyDataStat) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetCreateTime(v string) *GetTrackListResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetTotalNumber(v string) *GetTrackListResponseBodyDataStat {
	s.TotalNumber = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenCount = &v
	return s
}

type GetTrackListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTrackListResponse) SetBody(v *GetTrackListResponseBody) *GetTrackListResponse {
	s.Body = v
	return s
}

type GetTrackListByMailFromAndTagNameRequest struct {
	Total                *string `json:"Total,omitempty" xml:"Total,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Offset               *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OffsetCreateTime     *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	PageNumber           *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameRequest) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetTotal(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Total = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest {
	s.OwnerId = &v
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

func (s *GetTrackListByMailFromAndTagNameRequest) SetEndTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.EndTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffset(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Offset = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetPageSize(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.PageSize = &v
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

func (s *GetTrackListByMailFromAndTagNameRequest) SetPageNumber(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetAccountName(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.AccountName = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetTagName(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.TagName = &v
	return s
}

type GetTrackListByMailFromAndTagNameResponseBody struct {
	OffsetCreateTime     *string                                                `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	RequestId            *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize             *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total                *int32                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	TrackList            *GetTrackListByMailFromAndTagNameResponseBodyTrackList `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Struct"`
	PageNo               *int32                                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	OffsetCreateTimeDesc *string                                                `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
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

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetRequestId(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetPageSize(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.PageSize = &v
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

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetPageNo(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.OffsetCreateTimeDesc = &v
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
	RcptClickRate        *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	RcptUniqueOpenCount  *string `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	RcptClickCount       *string `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	RcptUniqueClickCount *string `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RcptUniqueOpenRate   *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	RcptUniqueClickRate  *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	TotalNumber          *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	RcptOpenRate         *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	RcptOpenCount        *string `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) String() string {
	return tea.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptClickRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueOpenCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptClickCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueClickCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetCreateTime(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueOpenRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueClickRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetTotalNumber(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.TotalNumber = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptOpenRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptOpenCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptOpenCount = &v
	return s
}

type GetTrackListByMailFromAndTagNameResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrackListByMailFromAndTagNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTrackListByMailFromAndTagNameResponse) SetBody(v *GetTrackListByMailFromAndTagNameResponseBody) *GetTrackListByMailFromAndTagNameResponse {
	s.Body = v
	return s
}

type MigrateMarketRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s MigrateMarketRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateMarketRequest) GoString() string {
	return s.String()
}

func (s *MigrateMarketRequest) SetOwnerId(v int64) *MigrateMarketRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateMarketRequest) SetResourceOwnerAccount(v string) *MigrateMarketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateMarketRequest) SetResourceOwnerId(v int64) *MigrateMarketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateMarketRequest) SetVersion(v string) *MigrateMarketRequest {
	s.Version = &v
	return s
}

func (s *MigrateMarketRequest) SetFromType(v int32) *MigrateMarketRequest {
	s.FromType = &v
	return s
}

type MigrateMarketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateMarketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateMarketResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateMarketResponseBody) SetRequestId(v string) *MigrateMarketResponseBody {
	s.RequestId = &v
	return s
}

type MigrateMarketResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateMarketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateMarketResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateMarketResponse) GoString() string {
	return s.String()
}

func (s *MigrateMarketResponse) SetHeaders(v map[string]*string) *MigrateMarketResponse {
	s.Headers = v
	return s
}

func (s *MigrateMarketResponse) SetBody(v *MigrateMarketResponseBody) *MigrateMarketResponse {
	s.Body = v
	return s
}

type ModifyAccountNotificationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyAccountNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountNotificationRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountNotificationRequest) SetOwnerId(v int64) *ModifyAccountNotificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountNotificationRequest) SetResourceOwnerAccount(v string) *ModifyAccountNotificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountNotificationRequest) SetResourceOwnerId(v int64) *ModifyAccountNotificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountNotificationRequest) SetRegion(v string) *ModifyAccountNotificationRequest {
	s.Region = &v
	return s
}

func (s *ModifyAccountNotificationRequest) SetStatus(v string) *ModifyAccountNotificationRequest {
	s.Status = &v
	return s
}

type ModifyAccountNotificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountNotificationResponseBody) SetRequestId(v string) *ModifyAccountNotificationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountNotificationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountNotificationResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountNotificationResponse) SetHeaders(v map[string]*string) *ModifyAccountNotificationResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountNotificationResponse) SetBody(v *ModifyAccountNotificationResponseBody) *ModifyAccountNotificationResponse {
	s.Body = v
	return s
}

type ModifyMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	MailAddressId        *int32  `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ModifyMailAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMailAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressRequest) SetOwnerId(v int64) *ModifyMailAddressRequest {
	s.OwnerId = &v
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

func (s *ModifyMailAddressRequest) SetMailAddressId(v int32) *ModifyMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *ModifyMailAddressRequest) SetReplyAddress(v string) *ModifyMailAddressRequest {
	s.ReplyAddress = &v
	return s
}

func (s *ModifyMailAddressRequest) SetPassword(v string) *ModifyMailAddressRequest {
	s.Password = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyMailAddressResponse) SetBody(v *ModifyMailAddressResponseBody) *ModifyMailAddressResponse {
	s.Body = v
	return s
}

type ModifyPWByDomainRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
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
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPWByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPWByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainResponseBody) SetMessage(v string) *ModifyPWByDomainResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetRequestId(v string) *ModifyPWByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetCode(v string) *ModifyPWByDomainResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPWByDomainResponseBody) SetSuccess(v bool) *ModifyPWByDomainResponseBody {
	s.Success = &v
	return s
}

type ModifyPWByDomainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPWByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyPWByDomainResponse) SetBody(v *ModifyPWByDomainResponseBody) *ModifyPWByDomainResponse {
	s.Body = v
	return s
}

type ModifySenderAddressNotificationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SenderAddressId      *string `json:"SenderAddressId,omitempty" xml:"SenderAddressId,omitempty"`
	SenderAddress        *string `json:"SenderAddress,omitempty" xml:"SenderAddress,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifySenderAddressNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySenderAddressNotificationRequest) GoString() string {
	return s.String()
}

func (s *ModifySenderAddressNotificationRequest) SetOwnerId(v int64) *ModifySenderAddressNotificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySenderAddressNotificationRequest) SetResourceOwnerAccount(v string) *ModifySenderAddressNotificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySenderAddressNotificationRequest) SetResourceOwnerId(v int64) *ModifySenderAddressNotificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySenderAddressNotificationRequest) SetSenderAddressId(v string) *ModifySenderAddressNotificationRequest {
	s.SenderAddressId = &v
	return s
}

func (s *ModifySenderAddressNotificationRequest) SetSenderAddress(v string) *ModifySenderAddressNotificationRequest {
	s.SenderAddress = &v
	return s
}

func (s *ModifySenderAddressNotificationRequest) SetRegion(v string) *ModifySenderAddressNotificationRequest {
	s.Region = &v
	return s
}

func (s *ModifySenderAddressNotificationRequest) SetStatus(v string) *ModifySenderAddressNotificationRequest {
	s.Status = &v
	return s
}

type ModifySenderAddressNotificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySenderAddressNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySenderAddressNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySenderAddressNotificationResponseBody) SetRequestId(v string) *ModifySenderAddressNotificationResponseBody {
	s.RequestId = &v
	return s
}

type ModifySenderAddressNotificationResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySenderAddressNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySenderAddressNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySenderAddressNotificationResponse) GoString() string {
	return s.String()
}

func (s *ModifySenderAddressNotificationResponse) SetHeaders(v map[string]*string) *ModifySenderAddressNotificationResponse {
	s.Headers = v
	return s
}

func (s *ModifySenderAddressNotificationResponse) SetBody(v *ModifySenderAddressNotificationResponseBody) *ModifySenderAddressNotificationResponse {
	s.Body = v
	return s
}

type ModifyTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagId                *int32  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyTagResponse) SetBody(v *ModifyTagResponseBody) *ModifyTagResponse {
	s.Body = v
	return s
}

type ModifyTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateSubject      *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	TemplateNickName     *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateText         *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	SmsType              *int32  `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	SmsContent           *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s ModifyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequest) SetOwnerId(v int64) *ModifyTemplateRequest {
	s.OwnerId = &v
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

func (s *ModifyTemplateRequest) SetTemplateId(v int32) *ModifyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateName(v string) *ModifyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateSubject(v string) *ModifyTemplateRequest {
	s.TemplateSubject = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateNickName(v string) *ModifyTemplateRequest {
	s.TemplateNickName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateText(v string) *ModifyTemplateRequest {
	s.TemplateText = &v
	return s
}

func (s *ModifyTemplateRequest) SetSmsType(v int32) *ModifyTemplateRequest {
	s.SmsType = &v
	return s
}

func (s *ModifyTemplateRequest) SetSmsContent(v string) *ModifyTemplateRequest {
	s.SmsContent = &v
	return s
}

func (s *ModifyTemplateRequest) SetRemark(v string) *ModifyTemplateRequest {
	s.Remark = &v
	return s
}

func (s *ModifyTemplateRequest) SetFromType(v int32) *ModifyTemplateRequest {
	s.FromType = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyTemplateResponse) SetBody(v *ModifyTemplateResponseBody) *ModifyTemplateResponse {
	s.Body = v
	return s
}

type QueryDomainByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDomainByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamRequest) SetOwnerId(v int64) *QueryDomainByParamRequest {
	s.OwnerId = &v
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

func (s *QueryDomainByParamRequest) SetPageNo(v int32) *QueryDomainByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryDomainByParamRequest) SetPageSize(v int32) *QueryDomainByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDomainByParamRequest) SetKeyWord(v string) *QueryDomainByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryDomainByParamRequest) SetStatus(v int32) *QueryDomainByParamRequest {
	s.Status = &v
	return s
}

type QueryDomainByParamResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryDomainByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryDomainByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBody) SetTotalCount(v int32) *QueryDomainByParamResponseBody {
	s.TotalCount = &v
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

func (s *QueryDomainByParamResponseBody) SetData(v *QueryDomainByParamResponseBodyData) *QueryDomainByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryDomainByParamResponseBody) SetPageNumber(v int32) *QueryDomainByParamResponseBody {
	s.PageNumber = &v
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
	DomainRecord    *string `json:"DomainRecord,omitempty" xml:"DomainRecord,omitempty"`
	SpfAuthStatus   *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	MxAuthStatus    *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	ConfirmStatus   *string `json:"ConfirmStatus,omitempty" xml:"ConfirmStatus,omitempty"`
	IcpStatus       *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	UtcCreateTime   *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	DomainStatus    *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainId        *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s QueryDomainByParamResponseBodyDataDomain) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByParamResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainRecord(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainRecord = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetSpfAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.SpfAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetMxAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.MxAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetCreateTime(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.CreateTime = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetCnameAuthStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.CnameAuthStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetConfirmStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.ConfirmStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetIcpStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.IcpStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetUtcCreateTime(v int64) *QueryDomainByParamResponseBodyDataDomain {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainStatus(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainName(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByParamResponseBodyDataDomain) SetDomainId(v string) *QueryDomainByParamResponseBodyDataDomain {
	s.DomainId = &v
	return s
}

type QueryDomainByParamResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDomainByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDomainByParamResponse) SetBody(v *QueryDomainByParamResponseBody) *QueryDomainByParamResponse {
	s.Body = v
	return s
}

type QueryInvalidAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Length               *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
}

func (s QueryInvalidAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressRequest) GoString() string {
	return s.String()
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

type QueryInvalidAddressResponseBody struct {
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryInvalidAddressResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	NextStart  *int32                               `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
}

func (s QueryInvalidAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponseBody) SetTotalCount(v int32) *QueryInvalidAddressResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetRequestId(v string) *QueryInvalidAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetData(v *QueryInvalidAddressResponseBodyData) *QueryInvalidAddressResponseBody {
	s.Data = v
	return s
}

func (s *QueryInvalidAddressResponseBody) SetNextStart(v int32) *QueryInvalidAddressResponseBody {
	s.NextStart = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryInvalidAddressResponse) SetBody(v *QueryInvalidAddressResponseBody) *QueryInvalidAddressResponse {
	s.Body = v
	return s
}

type QueryReceiverByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryReceiverByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamRequest) SetOwnerId(v int64) *QueryReceiverByParamRequest {
	s.OwnerId = &v
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

func (s *QueryReceiverByParamRequest) SetPageNo(v int32) *QueryReceiverByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetPageSize(v int32) *QueryReceiverByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetKeyWord(v string) *QueryReceiverByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetStatus(v int32) *QueryReceiverByParamRequest {
	s.Status = &v
	return s
}

type QueryReceiverByParamResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryReceiverByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	NextStart  *string                               `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
}

func (s QueryReceiverByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBody) SetTotalCount(v int32) *QueryReceiverByParamResponseBody {
	s.TotalCount = &v
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

func (s *QueryReceiverByParamResponseBody) SetData(v *QueryReceiverByParamResponseBodyData) *QueryReceiverByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryReceiverByParamResponseBody) SetNextStart(v string) *QueryReceiverByParamResponseBody {
	s.NextStart = &v
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
	ReceiversAlias  *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	ReceiversName   *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ReceiverId      *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	UtcCreateTime   *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	ReceiversStatus *string `json:"ReceiversStatus,omitempty" xml:"ReceiversStatus,omitempty"`
	Count           *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Desc            *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s QueryReceiverByParamResponseBodyDataReceiver) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverByParamResponseBodyDataReceiver) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversAlias(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversAlias = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversName(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversName = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetCreateTime(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiverId(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiverId = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetUtcCreateTime(v int64) *QueryReceiverByParamResponseBodyDataReceiver {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetReceiversStatus(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.ReceiversStatus = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetCount(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.Count = &v
	return s
}

func (s *QueryReceiverByParamResponseBodyDataReceiver) SetDesc(v string) *QueryReceiverByParamResponseBodyDataReceiver {
	s.Desc = &v
	return s
}

type QueryReceiverByParamResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReceiverByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryReceiverByParamResponse) SetBody(v *QueryReceiverByParamResponseBody) *QueryReceiverByParamResponse {
	s.Body = v
	return s
}

type QueryReceiverDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
}

func (s QueryReceiverDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailRequest) SetOwnerId(v int64) *QueryReceiverDetailRequest {
	s.OwnerId = &v
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

func (s *QueryReceiverDetailRequest) SetReceiverId(v string) *QueryReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetPageSize(v int32) *QueryReceiverDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetKeyWord(v string) *QueryReceiverDetailRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryReceiverDetailRequest) SetNextStart(v string) *QueryReceiverDetailRequest {
	s.NextStart = &v
	return s
}

type QueryReceiverDetailResponseBody struct {
	DataSchema *string                              `json:"DataSchema,omitempty" xml:"DataSchema,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryReceiverDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	NextStart  *string                              `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
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

func (s *QueryReceiverDetailResponseBody) SetTotalCount(v int32) *QueryReceiverDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetRequestId(v string) *QueryReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetData(v *QueryReceiverDetailResponseBodyData) *QueryReceiverDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryReceiverDetailResponseBody) SetNextStart(v string) *QueryReceiverDetailResponseBody {
	s.NextStart = &v
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
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
}

func (s QueryReceiverDetailResponseBodyDataDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryReceiverDetailResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetData(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.Data = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetEmail(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.Email = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetCreateTime(v string) *QueryReceiverDetailResponseBodyDataDetail {
	s.CreateTime = &v
	return s
}

func (s *QueryReceiverDetailResponseBodyDataDetail) SetUtcCreateTime(v int64) *QueryReceiverDetailResponseBodyDataDetail {
	s.UtcCreateTime = &v
	return s
}

type QueryReceiverDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryReceiverDetailResponse) SetBody(v *QueryReceiverDetailResponseBody) *QueryReceiverDetailResponse {
	s.Body = v
	return s
}

type QuerySignByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s QuerySignByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySignByParamRequest) GoString() string {
	return s.String()
}

func (s *QuerySignByParamRequest) SetOwnerId(v int64) *QuerySignByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySignByParamRequest) SetResourceOwnerAccount(v string) *QuerySignByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySignByParamRequest) SetResourceOwnerId(v int64) *QuerySignByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySignByParamRequest) SetPageNo(v int32) *QuerySignByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySignByParamRequest) SetPageSize(v int32) *QuerySignByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySignByParamRequest) SetKeyWord(v string) *QuerySignByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QuerySignByParamRequest) SetFromType(v int32) *QuerySignByParamRequest {
	s.FromType = &v
	return s
}

type QuerySignByParamResponseBody struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data       *QuerySignByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QuerySignByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySignByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySignByParamResponseBody) SetRequestId(v string) *QuerySignByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySignByParamResponseBody) SetPageSize(v int32) *QuerySignByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuerySignByParamResponseBody) SetData(v *QuerySignByParamResponseBodyData) *QuerySignByParamResponseBody {
	s.Data = v
	return s
}

func (s *QuerySignByParamResponseBody) SetPageNumber(v int32) *QuerySignByParamResponseBody {
	s.PageNumber = &v
	return s
}

type QuerySignByParamResponseBodyData struct {
	Sign []*QuerySignByParamResponseBodyDataSign `json:"sign,omitempty" xml:"sign,omitempty" type:"Repeated"`
}

func (s QuerySignByParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySignByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySignByParamResponseBodyData) SetSign(v []*QuerySignByParamResponseBodyDataSign) *QuerySignByParamResponseBodyData {
	s.Sign = v
	return s
}

type QuerySignByParamResponseBodyDataSign struct {
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	AuditState *string `json:"AuditState,omitempty" xml:"AuditState,omitempty"`
	GmtCreate  *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SignId     *int64  `json:"SignId,omitempty" xml:"SignId,omitempty"`
	SignName   *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	SignType   *string `json:"SignType,omitempty" xml:"SignType,omitempty"`
}

func (s QuerySignByParamResponseBodyDataSign) String() string {
	return tea.Prettify(s)
}

func (s QuerySignByParamResponseBodyDataSign) GoString() string {
	return s.String()
}

func (s *QuerySignByParamResponseBodyDataSign) SetRemark(v string) *QuerySignByParamResponseBodyDataSign {
	s.Remark = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetAuditState(v string) *QuerySignByParamResponseBodyDataSign {
	s.AuditState = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetGmtCreate(v string) *QuerySignByParamResponseBodyDataSign {
	s.GmtCreate = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetSignId(v int64) *QuerySignByParamResponseBodyDataSign {
	s.SignId = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetSignName(v string) *QuerySignByParamResponseBodyDataSign {
	s.SignName = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetOrderId(v string) *QuerySignByParamResponseBodyDataSign {
	s.OrderId = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetRejectInfo(v string) *QuerySignByParamResponseBodyDataSign {
	s.RejectInfo = &v
	return s
}

func (s *QuerySignByParamResponseBodyDataSign) SetSignType(v string) *QuerySignByParamResponseBodyDataSign {
	s.SignType = &v
	return s
}

type QuerySignByParamResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySignByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySignByParamResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySignByParamResponse) GoString() string {
	return s.String()
}

func (s *QuerySignByParamResponse) SetHeaders(v map[string]*string) *QuerySignByParamResponse {
	s.Headers = v
	return s
}

func (s *QuerySignByParamResponse) SetBody(v *QuerySignByParamResponseBody) *QuerySignByParamResponse {
	s.Body = v
	return s
}

type QuerySmsStatisticsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s QuerySmsStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsStatisticsRequest) SetOwnerId(v int64) *QuerySmsStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsStatisticsRequest) SetResourceOwnerAccount(v string) *QuerySmsStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsStatisticsRequest) SetResourceOwnerId(v int64) *QuerySmsStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsStatisticsRequest) SetStartTime(v string) *QuerySmsStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySmsStatisticsRequest) SetEndTime(v string) *QuerySmsStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySmsStatisticsRequest) SetFromType(v int32) *QuerySmsStatisticsRequest {
	s.FromType = &v
	return s
}

type QuerySmsStatisticsResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QuerySmsStatisticsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QuerySmsStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsStatisticsResponseBody) SetTotalCount(v int32) *QuerySmsStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySmsStatisticsResponseBody) SetRequestId(v string) *QuerySmsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsStatisticsResponseBody) SetData(v *QuerySmsStatisticsResponseBodyData) *QuerySmsStatisticsResponseBody {
	s.Data = v
	return s
}

type QuerySmsStatisticsResponseBodyData struct {
	Stat []*QuerySmsStatisticsResponseBodyDataStat `json:"stat,omitempty" xml:"stat,omitempty" type:"Repeated"`
}

func (s QuerySmsStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySmsStatisticsResponseBodyData) SetStat(v []*QuerySmsStatisticsResponseBodyDataStat) *QuerySmsStatisticsResponseBodyData {
	s.Stat = v
	return s
}

type QuerySmsStatisticsResponseBodyDataStat struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaildCount   *string `json:"faildCount,omitempty" xml:"faildCount,omitempty"`
	SuccessCount *string `json:"successCount,omitempty" xml:"successCount,omitempty"`
	RequestCount *string `json:"requestCount,omitempty" xml:"requestCount,omitempty"`
}

func (s QuerySmsStatisticsResponseBodyDataStat) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsStatisticsResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *QuerySmsStatisticsResponseBodyDataStat) SetCreateTime(v string) *QuerySmsStatisticsResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *QuerySmsStatisticsResponseBodyDataStat) SetFaildCount(v string) *QuerySmsStatisticsResponseBodyDataStat {
	s.FaildCount = &v
	return s
}

func (s *QuerySmsStatisticsResponseBodyDataStat) SetSuccessCount(v string) *QuerySmsStatisticsResponseBodyDataStat {
	s.SuccessCount = &v
	return s
}

func (s *QuerySmsStatisticsResponseBodyDataStat) SetRequestCount(v string) *QuerySmsStatisticsResponseBodyDataStat {
	s.RequestCount = &v
	return s
}

type QuerySmsStatisticsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySmsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySmsStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsStatisticsResponse) SetHeaders(v map[string]*string) *QuerySmsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsStatisticsResponse) SetBody(v *QuerySmsStatisticsResponseBody) *QuerySmsStatisticsResponse {
	s.Body = v
	return s
}

type QueryTagByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
}

func (s QueryTagByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTagByParamRequest) SetOwnerId(v int64) *QueryTagByParamRequest {
	s.OwnerId = &v
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

func (s *QueryTagByParamRequest) SetPageNo(v int32) *QueryTagByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTagByParamRequest) SetPageSize(v int32) *QueryTagByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTagByParamRequest) SetKeyWord(v string) *QueryTagByParamRequest {
	s.KeyWord = &v
	return s
}

type QueryTagByParamResponseBody struct {
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryTagByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryTagByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBody) SetTotalCount(v int32) *QueryTagByParamResponseBody {
	s.TotalCount = &v
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

func (s *QueryTagByParamResponseBody) SetData(v *QueryTagByParamResponseBodyData) *QueryTagByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagByParamResponseBody) SetPageNumber(v int32) *QueryTagByParamResponseBody {
	s.PageNumber = &v
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
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryTagByParamResponseBodyDataTag) String() string {
	return tea.Prettify(s)
}

func (s QueryTagByParamResponseBodyDataTag) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagName(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagName = &v
	return s
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagId(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagId = &v
	return s
}

type QueryTagByParamResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTagByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTagByParamResponse) SetBody(v *QueryTagByParamResponseBody) *QueryTagByParamResponse {
	s.Body = v
	return s
}

type QueryTaskByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTaskByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamRequest) SetOwnerId(v int64) *QueryTaskByParamRequest {
	s.OwnerId = &v
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

func (s *QueryTaskByParamRequest) SetPageNo(v int32) *QueryTaskByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTaskByParamRequest) SetPageSize(v int32) *QueryTaskByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskByParamRequest) SetKeyWord(v string) *QueryTaskByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTaskByParamRequest) SetStatus(v int32) *QueryTaskByParamRequest {
	s.Status = &v
	return s
}

type QueryTaskByParamResponseBody struct {
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryTaskByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryTaskByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBody) SetTotalCount(v int32) *QueryTaskByParamResponseBody {
	s.TotalCount = &v
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

func (s *QueryTaskByParamResponseBody) SetData(v *QueryTaskByParamResponseBodyData) *QueryTaskByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskByParamResponseBody) SetPageNumber(v int32) *QueryTaskByParamResponseBody {
	s.PageNumber = &v
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
	ReceiversName *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	TagName       *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TaskStatus    *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RequestCount  *string `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	AddressType   *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	UtcCreateTime *int64  `json:"UtcCreateTime,omitempty" xml:"UtcCreateTime,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryTaskByParamResponseBodyDataTask) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskByParamResponseBodyDataTask) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponseBodyDataTask) SetReceiversName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.ReceiversName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTagName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TagName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTaskStatus(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetCreateTime(v string) *QueryTaskByParamResponseBodyDataTask {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetRequestCount(v string) *QueryTaskByParamResponseBodyDataTask {
	s.RequestCount = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetAddressType(v string) *QueryTaskByParamResponseBodyDataTask {
	s.AddressType = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetUtcCreateTime(v int64) *QueryTaskByParamResponseBodyDataTask {
	s.UtcCreateTime = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTemplateName(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TemplateName = &v
	return s
}

func (s *QueryTaskByParamResponseBodyDataTask) SetTaskId(v string) *QueryTaskByParamResponseBodyDataTask {
	s.TaskId = &v
	return s
}

type QueryTaskByParamResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTaskByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTaskByParamResponse) SetBody(v *QueryTaskByParamResponseBody) *QueryTaskByParamResponse {
	s.Body = v
	return s
}

type QueryTemplateByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	KeyWord              *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s QueryTemplateByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamRequest) SetOwnerId(v int64) *QueryTemplateByParamRequest {
	s.OwnerId = &v
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

func (s *QueryTemplateByParamRequest) SetPageNo(v int32) *QueryTemplateByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetPageSize(v int32) *QueryTemplateByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetKeyWord(v string) *QueryTemplateByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetStatus(v int32) *QueryTemplateByParamRequest {
	s.Status = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetFromType(v int32) *QueryTemplateByParamRequest {
	s.FromType = &v
	return s
}

type QueryTemplateByParamResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *QueryTemplateByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s QueryTemplateByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBody) SetTotalCount(v int32) *QueryTemplateByParamResponseBody {
	s.TotalCount = &v
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

func (s *QueryTemplateByParamResponseBody) SetData(v *QueryTemplateByParamResponseBodyData) *QueryTemplateByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryTemplateByParamResponseBody) SetPageNumber(v int32) *QueryTemplateByParamResponseBody {
	s.PageNumber = &v
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
	TemplateComment *string `json:"TemplateComment,omitempty" xml:"TemplateComment,omitempty"`
	UtcCreatetime   *int64  `json:"UtcCreatetime,omitempty" xml:"UtcCreatetime,omitempty"`
	Smsrejectinfo   *int32  `json:"Smsrejectinfo,omitempty" xml:"Smsrejectinfo,omitempty"`
	SmsTemplateCode *int32  `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateStatus  *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateType    *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	SmsStatus       *int32  `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryTemplateByParamResponseBodyDataTemplate) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateByParamResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateComment(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateComment = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetUtcCreatetime(v int64) *QueryTemplateByParamResponseBodyDataTemplate {
	s.UtcCreatetime = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsrejectinfo(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.Smsrejectinfo = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsTemplateCode(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.SmsTemplateCode = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetCreateTime(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.CreateTime = &v
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

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateName(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateName = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetSmsStatus(v int32) *QueryTemplateByParamResponseBodyDataTemplate {
	s.SmsStatus = &v
	return s
}

func (s *QueryTemplateByParamResponseBodyDataTemplate) SetTemplateId(v string) *QueryTemplateByParamResponseBodyDataTemplate {
	s.TemplateId = &v
	return s
}

type QueryTemplateByParamResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTemplateByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTemplateByParamResponse) SetBody(v *QueryTemplateByParamResponseBody) *QueryTemplateByParamResponse {
	s.Body = v
	return s
}

type SaveReceiverDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s SaveReceiverDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailRequest) SetOwnerId(v int64) *SaveReceiverDetailRequest {
	s.OwnerId = &v
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

func (s *SaveReceiverDetailRequest) SetReceiverId(v string) *SaveReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *SaveReceiverDetailRequest) SetDetail(v string) *SaveReceiverDetailRequest {
	s.Detail = &v
	return s
}

type SaveReceiverDetailResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *SaveReceiverDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCount   *int32                              `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	SuccessCount *int32                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s SaveReceiverDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveReceiverDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponseBody) SetRequestId(v string) *SaveReceiverDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetData(v *SaveReceiverDetailResponseBodyData) *SaveReceiverDetailResponseBody {
	s.Data = v
	return s
}

func (s *SaveReceiverDetailResponseBody) SetErrorCount(v int32) *SaveReceiverDetailResponseBody {
	s.ErrorCount = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SaveReceiverDetailResponse) SetBody(v *SaveReceiverDetailResponseBody) *SaveReceiverDetailResponse {
	s.Body = v
	return s
}

type SenderStatisticsByTagNameAndBatchIDRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDRequest) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDRequest) GoString() string {
	return s.String()
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

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetAccountName(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetStartTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.StartTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetEndTime(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.EndTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDRequest) SetTagName(v string) *SenderStatisticsByTagNameAndBatchIDRequest {
	s.TagName = &v
	return s
}

type SenderStatisticsByTagNameAndBatchIDResponseBody struct {
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       *SenderStatisticsByTagNameAndBatchIDResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetTotalCount(v int32) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBody) SetRequestId(v string) *SenderStatisticsByTagNameAndBatchIDResponseBody {
	s.RequestId = &v
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
	UnavailablePercent *string `json:"unavailablePercent,omitempty" xml:"unavailablePercent,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SucceededPercent   *string `json:"succeededPercent,omitempty" xml:"succeededPercent,omitempty"`
	FaildCount         *string `json:"faildCount,omitempty" xml:"faildCount,omitempty"`
	UnavailableCount   *string `json:"unavailableCount,omitempty" xml:"unavailableCount,omitempty"`
	SuccessCount       *string `json:"successCount,omitempty" xml:"successCount,omitempty"`
	RequestCount       *string `json:"requestCount,omitempty" xml:"requestCount,omitempty"`
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetUnavailablePercent(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.UnavailablePercent = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetCreateTime(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetSucceededPercent(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.SucceededPercent = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetFaildCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.FaildCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetUnavailableCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.UnavailableCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetSuccessCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.SuccessCount = &v
	return s
}

func (s *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat) SetRequestCount(v string) *SenderStatisticsByTagNameAndBatchIDResponseBodyDataStat {
	s.RequestCount = &v
	return s
}

type SenderStatisticsByTagNameAndBatchIDResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SenderStatisticsByTagNameAndBatchIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SenderStatisticsByTagNameAndBatchIDResponse) SetBody(v *SenderStatisticsByTagNameAndBatchIDResponseBody) *SenderStatisticsByTagNameAndBatchIDResponse {
	s.Body = v
	return s
}

type SenderStatisticsDetailByParamRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ToAddress            *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Length               *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
}

func (s SenderStatisticsDetailByParamRequest) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamRequest) GoString() string {
	return s.String()
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

func (s *SenderStatisticsDetailByParamRequest) SetAccountName(v string) *SenderStatisticsDetailByParamRequest {
	s.AccountName = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetToAddress(v string) *SenderStatisticsDetailByParamRequest {
	s.ToAddress = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetStatus(v int32) *SenderStatisticsDetailByParamRequest {
	s.Status = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetStartTime(v string) *SenderStatisticsDetailByParamRequest {
	s.StartTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetEndTime(v string) *SenderStatisticsDetailByParamRequest {
	s.EndTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamRequest) SetTagName(v string) *SenderStatisticsDetailByParamRequest {
	s.TagName = &v
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

type SenderStatisticsDetailByParamResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SenderStatisticsDetailByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	NextStart *int32                                         `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
}

func (s SenderStatisticsDetailByParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBody) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBody) SetRequestId(v string) *SenderStatisticsDetailByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) SetData(v *SenderStatisticsDetailByParamResponseBodyData) *SenderStatisticsDetailByParamResponseBody {
	s.Data = v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBody) SetNextStart(v int32) *SenderStatisticsDetailByParamResponseBody {
	s.NextStart = &v
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
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	LastUpdateTime    *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ToAddress         *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UtcLastUpdateTime *string `json:"UtcLastUpdateTime,omitempty" xml:"UtcLastUpdateTime,omitempty"`
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s SenderStatisticsDetailByParamResponseBodyDataMailDetail) String() string {
	return tea.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponseBodyDataMailDetail) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetStatus(v int32) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.Status = &v
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

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetToAddress(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.ToAddress = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetUtcLastUpdateTime(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.UtcLastUpdateTime = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponseBodyDataMailDetail) SetAccountName(v string) *SenderStatisticsDetailByParamResponseBodyDataMailDetail {
	s.AccountName = &v
	return s
}

type SenderStatisticsDetailByParamResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SenderStatisticsDetailByParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SenderStatisticsDetailByParamResponse) SetBody(v *SenderStatisticsDetailByParamResponseBody) *SenderStatisticsDetailByParamResponse {
	s.Body = v
	return s
}

type SendTestByTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	NickName             *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Birthday             *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	Gender               *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Mobile               *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Email                *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s SendTestByTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTestByTemplateRequest) GoString() string {
	return s.String()
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

func (s *SendTestByTemplateRequest) SetAccountName(v string) *SendTestByTemplateRequest {
	s.AccountName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetUserName(v string) *SendTestByTemplateRequest {
	s.UserName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetNickName(v string) *SendTestByTemplateRequest {
	s.NickName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetBirthday(v string) *SendTestByTemplateRequest {
	s.Birthday = &v
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

func (s *SendTestByTemplateRequest) SetEmail(v string) *SendTestByTemplateRequest {
	s.Email = &v
	return s
}

type SendTestByTemplateResponseBody struct {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendTestByTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendTestByTemplateResponse) SetBody(v *SendTestByTemplateResponseBody) *SendTestByTemplateResponse {
	s.Body = v
	return s
}

type SingleSendMailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AddressType          *int32  `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	ReplyToAddress       *bool   `json:"ReplyToAddress,omitempty" xml:"ReplyToAddress,omitempty"`
	ToAddress            *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	Subject              *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	HtmlBody             *string `json:"HtmlBody,omitempty" xml:"HtmlBody,omitempty"`
	TextBody             *string `json:"TextBody,omitempty" xml:"TextBody,omitempty"`
	FromAlias            *string `json:"FromAlias,omitempty" xml:"FromAlias,omitempty"`
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ReplyAddressAlias    *string `json:"ReplyAddressAlias,omitempty" xml:"ReplyAddressAlias,omitempty"`
	ClickTrace           *string `json:"ClickTrace,omitempty" xml:"ClickTrace,omitempty"`
}

func (s SingleSendMailRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailRequest) GoString() string {
	return s.String()
}

func (s *SingleSendMailRequest) SetOwnerId(v int64) *SingleSendMailRequest {
	s.OwnerId = &v
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

func (s *SingleSendMailRequest) SetAccountName(v string) *SingleSendMailRequest {
	s.AccountName = &v
	return s
}

func (s *SingleSendMailRequest) SetAddressType(v int32) *SingleSendMailRequest {
	s.AddressType = &v
	return s
}

func (s *SingleSendMailRequest) SetTagName(v string) *SingleSendMailRequest {
	s.TagName = &v
	return s
}

func (s *SingleSendMailRequest) SetReplyToAddress(v bool) *SingleSendMailRequest {
	s.ReplyToAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetToAddress(v string) *SingleSendMailRequest {
	s.ToAddress = &v
	return s
}

func (s *SingleSendMailRequest) SetSubject(v string) *SingleSendMailRequest {
	s.Subject = &v
	return s
}

func (s *SingleSendMailRequest) SetHtmlBody(v string) *SingleSendMailRequest {
	s.HtmlBody = &v
	return s
}

func (s *SingleSendMailRequest) SetTextBody(v string) *SingleSendMailRequest {
	s.TextBody = &v
	return s
}

func (s *SingleSendMailRequest) SetFromAlias(v string) *SingleSendMailRequest {
	s.FromAlias = &v
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

func (s *SingleSendMailRequest) SetClickTrace(v string) *SingleSendMailRequest {
	s.ClickTrace = &v
	return s
}

type SingleSendMailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnvId     *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s SingleSendMailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleSendMailResponseBody) GoString() string {
	return s.String()
}

func (s *SingleSendMailResponseBody) SetRequestId(v string) *SingleSendMailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleSendMailResponseBody) SetEnvId(v string) *SingleSendMailResponseBody {
	s.EnvId = &v
	return s
}

type SingleSendMailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SingleSendMailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SingleSendMailResponse) SetBody(v *SingleSendMailResponseBody) *SingleSendMailResponse {
	s.Body = v
	return s
}

type SingleSendSmsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	RecNum               *string `json:"RecNum,omitempty" xml:"RecNum,omitempty"`
	ParamString          *string `json:"ParamString,omitempty" xml:"ParamString,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SingleSendSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleSendSmsRequest) GoString() string {
	return s.String()
}

func (s *SingleSendSmsRequest) SetOwnerId(v int64) *SingleSendSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleSendSmsRequest) SetResourceOwnerAccount(v string) *SingleSendSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleSendSmsRequest) SetResourceOwnerId(v int64) *SingleSendSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleSendSmsRequest) SetSignName(v string) *SingleSendSmsRequest {
	s.SignName = &v
	return s
}

func (s *SingleSendSmsRequest) SetTemplateCode(v string) *SingleSendSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SingleSendSmsRequest) SetRecNum(v string) *SingleSendSmsRequest {
	s.RecNum = &v
	return s
}

func (s *SingleSendSmsRequest) SetParamString(v string) *SingleSendSmsRequest {
	s.ParamString = &v
	return s
}

func (s *SingleSendSmsRequest) SetVersion(v string) *SingleSendSmsRequest {
	s.Version = &v
	return s
}

type SingleSendSmsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleSendSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleSendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SingleSendSmsResponseBody) SetRequestId(v string) *SingleSendSmsResponseBody {
	s.RequestId = &v
	return s
}

type SingleSendSmsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SingleSendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SingleSendSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SingleSendSmsResponse) GoString() string {
	return s.String()
}

func (s *SingleSendSmsResponse) SetHeaders(v map[string]*string) *SingleSendSmsResponse {
	s.Headers = v
	return s
}

func (s *SingleSendSmsResponse) SetBody(v *SingleSendSmsResponseBody) *SingleSendSmsResponse {
	s.Body = v
	return s
}

type UpdateDomainTrackNameRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	CnameTrackRecord     *string `json:"CnameTrackRecord,omitempty" xml:"CnameTrackRecord,omitempty"`
}

func (s UpdateDomainTrackNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainTrackNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainTrackNameRequest) SetOwnerId(v int64) *UpdateDomainTrackNameRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDomainTrackNameRequest) SetResourceOwnerAccount(v string) *UpdateDomainTrackNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateDomainTrackNameRequest) SetResourceOwnerId(v int64) *UpdateDomainTrackNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateDomainTrackNameRequest) SetDomainId(v int32) *UpdateDomainTrackNameRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateDomainTrackNameRequest) SetCnameTrackRecord(v string) *UpdateDomainTrackNameRequest {
	s.CnameTrackRecord = &v
	return s
}

type UpdateDomainTrackNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainTrackNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainTrackNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainTrackNameResponseBody) SetRequestId(v string) *UpdateDomainTrackNameResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDomainTrackNameResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDomainTrackNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDomainTrackNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainTrackNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainTrackNameResponse) SetHeaders(v map[string]*string) *UpdateDomainTrackNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainTrackNameResponse) SetBody(v *UpdateDomainTrackNameResponseBody) *UpdateDomainTrackNameResponse {
	s.Body = v
	return s
}

type UpdateIpProtectionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	IpProtection         *string `json:"IpProtection,omitempty" xml:"IpProtection,omitempty"`
}

func (s UpdateIpProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpProtectionRequest) GoString() string {
	return s.String()
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

func (s *UpdateIpProtectionRequest) SetIpProtection(v string) *UpdateIpProtectionRequest {
	s.IpProtection = &v
	return s
}

type UpdateIpProtectionResponseBody struct {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIpProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateIpProtectionResponse) SetBody(v *UpdateIpProtectionResponseBody) *UpdateIpProtectionResponse {
	s.Body = v
	return s
}

type UpdateMailAddressMsgCallBackUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	MailFrom             *string `json:"MailFrom,omitempty" xml:"MailFrom,omitempty"`
	NotifyUrl            *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s UpdateMailAddressMsgCallBackUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailAddressMsgCallBackUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateMailAddressMsgCallBackUrlRequest) SetOwnerId(v int64) *UpdateMailAddressMsgCallBackUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMailAddressMsgCallBackUrlRequest) SetResourceOwnerAccount(v string) *UpdateMailAddressMsgCallBackUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMailAddressMsgCallBackUrlRequest) SetResourceOwnerId(v int64) *UpdateMailAddressMsgCallBackUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMailAddressMsgCallBackUrlRequest) SetMailFrom(v string) *UpdateMailAddressMsgCallBackUrlRequest {
	s.MailFrom = &v
	return s
}

func (s *UpdateMailAddressMsgCallBackUrlRequest) SetNotifyUrl(v string) *UpdateMailAddressMsgCallBackUrlRequest {
	s.NotifyUrl = &v
	return s
}

type UpdateMailAddressMsgCallBackUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMailAddressMsgCallBackUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailAddressMsgCallBackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMailAddressMsgCallBackUrlResponseBody) SetRequestId(v string) *UpdateMailAddressMsgCallBackUrlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMailAddressMsgCallBackUrlResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMailAddressMsgCallBackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMailAddressMsgCallBackUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailAddressMsgCallBackUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateMailAddressMsgCallBackUrlResponse) SetHeaders(v map[string]*string) *UpdateMailAddressMsgCallBackUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateMailAddressMsgCallBackUrlResponse) SetBody(v *UpdateMailAddressMsgCallBackUrlResponseBody) *UpdateMailAddressMsgCallBackUrlResponse {
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

func (client *Client) AddIpfilterWithOptions(request *AddIpfilterRequest, runtime *util.RuntimeOptions) (_result *AddIpfilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddIpfilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddIpfilter"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ApproveMailTemplateWithOptions(request *ApproveMailTemplateRequest, runtime *util.RuntimeOptions) (_result *ApproveMailTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApproveMailTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApproveMailTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveMailTemplate(request *ApproveMailTemplateRequest) (_result *ApproveMailTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveMailTemplateResponse{}
	_body, _err := client.ApproveMailTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveReplyMailAddressWithOptions(request *ApproveReplyMailAddressRequest, runtime *util.RuntimeOptions) (_result *ApproveReplyMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApproveReplyMailAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApproveReplyMailAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ApproveSmsTemplateWithOptions(request *ApproveSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *ApproveSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApproveSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApproveSmsTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveSmsTemplate(request *ApproveSmsTemplateRequest) (_result *ApproveSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveSmsTemplateResponse{}
	_body, _err := client.ApproveSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveTemplateWithOptions(request *ApproveTemplateRequest, runtime *util.RuntimeOptions) (_result *ApproveTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApproveTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApproveTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveTemplate(request *ApproveTemplateRequest) (_result *ApproveTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveTemplateResponse{}
	_body, _err := client.ApproveTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSendMailWithOptions(request *BatchSendMailRequest, runtime *util.RuntimeOptions) (_result *BatchSendMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchSendMailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchSendMail"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CheckDomainWithOptions(request *CheckDomainRequest, runtime *util.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckDomain"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CheckInvalidAddressWithOptions(request *CheckInvalidAddressRequest, runtime *util.RuntimeOptions) (_result *CheckInvalidAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckInvalidAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckInvalidAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckInvalidAddress(request *CheckInvalidAddressRequest) (_result *CheckInvalidAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckInvalidAddressResponse{}
	_body, _err := client.CheckInvalidAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckReplyToMailAddressWithOptions(request *CheckReplyToMailAddressRequest, runtime *util.RuntimeOptions) (_result *CheckReplyToMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckReplyToMailAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckReplyToMailAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateDayuWithOptions(request *CreateDayuRequest, runtime *util.RuntimeOptions) (_result *CreateDayuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDayuResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDayu"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDayu(request *CreateDayuRequest) (_result *CreateDayuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDayuResponse{}
	_body, _err := client.CreateDayuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDomain"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateMailAddressWithOptions(request *CreateMailAddressRequest, runtime *util.RuntimeOptions) (_result *CreateMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMailAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMailAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateReceiverWithOptions(request *CreateReceiverRequest, runtime *util.RuntimeOptions) (_result *CreateReceiverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateReceiverResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateReceiver"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateSignWithOptions(request *CreateSignRequest, runtime *util.RuntimeOptions) (_result *CreateSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSign"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSign(request *CreateSignRequest) (_result *CreateSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSignResponse{}
	_body, _err := client.CreateSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTagWithOptions(request *CreateTagRequest, runtime *util.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTag"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomain"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteInvalidAddressWithOptions(request *DeleteInvalidAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteInvalidAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInvalidAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInvalidAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteIpfilterByEdmIdWithOptions(request *DeleteIpfilterByEdmIdRequest, runtime *util.RuntimeOptions) (_result *DeleteIpfilterByEdmIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIpfilterByEdmIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIpfilterByEdmId"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteMailAddressWithOptions(request *DeleteMailAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMailAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMailAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteReceiverWithOptions(request *DeleteReceiverRequest, runtime *util.RuntimeOptions) (_result *DeleteReceiverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteReceiverResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteReceiver"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteReceiverDetailWithOptions(request *DeleteReceiverDetailRequest, runtime *util.RuntimeOptions) (_result *DeleteReceiverDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteReceiverDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteReceiverDetail"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteSignWithOptions(request *DeleteSignRequest, runtime *util.RuntimeOptions) (_result *DeleteSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSign"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSign(request *DeleteSignRequest) (_result *DeleteSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSignResponse{}
	_body, _err := client.DeleteSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagWithOptions(request *DeleteTagRequest, runtime *util.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTag"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescAccountSummaryWithOptions(request *DescAccountSummaryRequest, runtime *util.RuntimeOptions) (_result *DescAccountSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescAccountSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescAccountSummary"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescAccountSummary2WithOptions(request *DescAccountSummary2Request, runtime *util.RuntimeOptions) (_result *DescAccountSummary2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescAccountSummary2Response{}
	_body, _err := client.DoRPCRequest(tea.String("DescAccountSummary2"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescAccountSummary2(request *DescAccountSummary2Request) (_result *DescAccountSummary2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescAccountSummary2Response{}
	_body, _err := client.DescAccountSummary2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescDomainWithOptions(request *DescDomainRequest, runtime *util.RuntimeOptions) (_result *DescDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescDomain"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) DescTemplateWithOptions(request *DescTemplateRequest, runtime *util.RuntimeOptions) (_result *DescTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) EnableAccountWithOptions(request *EnableAccountRequest, runtime *util.RuntimeOptions) (_result *EnableAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableAccount"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAccount(request *EnableAccountRequest) (_result *EnableAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAccountResponse{}
	_body, _err := client.EnableAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountListWithOptions(request *GetAccountListRequest, runtime *util.RuntimeOptions) (_result *GetAccountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAccountListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAccountList"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetIpfilterListWithOptions(request *GetIpfilterListRequest, runtime *util.RuntimeOptions) (_result *GetIpfilterListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIpfilterListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIpfilterList"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetIpProtectionWithOptions(request *GetIpProtectionRequest, runtime *util.RuntimeOptions) (_result *GetIpProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIpProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIpProtection"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetMailAddressMsgCallBackUrlWithOptions(request *GetMailAddressMsgCallBackUrlRequest, runtime *util.RuntimeOptions) (_result *GetMailAddressMsgCallBackUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMailAddressMsgCallBackUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMailAddressMsgCallBackUrl"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetRegionListWithOptions(request *GetRegionListRequest, runtime *util.RuntimeOptions) (_result *GetRegionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRegionListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRegionList"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegionList(request *GetRegionListRequest) (_result *GetRegionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegionListResponse{}
	_body, _err := client.GetRegionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSenderAddressListWithOptions(request *GetSenderAddressListRequest, runtime *util.RuntimeOptions) (_result *GetSenderAddressListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSenderAddressListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSenderAddressList"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSenderAddressList(request *GetSenderAddressListRequest) (_result *GetSenderAddressListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSenderAddressListResponse{}
	_body, _err := client.GetSenderAddressListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrackListWithOptions(request *GetTrackListRequest, runtime *util.RuntimeOptions) (_result *GetTrackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTrackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrackList"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) GetTrackListByMailFromAndTagNameWithOptions(request *GetTrackListByMailFromAndTagNameRequest, runtime *util.RuntimeOptions) (_result *GetTrackListByMailFromAndTagNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTrackListByMailFromAndTagNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrackListByMailFromAndTagName"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) MigrateMarketWithOptions(request *MigrateMarketRequest, runtime *util.RuntimeOptions) (_result *MigrateMarketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MigrateMarketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MigrateMarket"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateMarket(request *MigrateMarketRequest) (_result *MigrateMarketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateMarketResponse{}
	_body, _err := client.MigrateMarketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountNotificationWithOptions(request *ModifyAccountNotificationRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountNotificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountNotification"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountNotification(request *ModifyAccountNotificationRequest) (_result *ModifyAccountNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountNotificationResponse{}
	_body, _err := client.ModifyAccountNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMailAddressWithOptions(request *ModifyMailAddressRequest, runtime *util.RuntimeOptions) (_result *ModifyMailAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMailAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMailAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyPWByDomainWithOptions(request *ModifyPWByDomainRequest, runtime *util.RuntimeOptions) (_result *ModifyPWByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPWByDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPWByDomain"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifySenderAddressNotificationWithOptions(request *ModifySenderAddressNotificationRequest, runtime *util.RuntimeOptions) (_result *ModifySenderAddressNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySenderAddressNotificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySenderAddressNotification"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySenderAddressNotification(request *ModifySenderAddressNotificationRequest) (_result *ModifySenderAddressNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySenderAddressNotificationResponse{}
	_body, _err := client.ModifySenderAddressNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTagWithOptions(request *ModifyTagRequest, runtime *util.RuntimeOptions) (_result *ModifyTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyTag"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyTemplateWithOptions(request *ModifyTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryDomainByParamWithOptions(request *QueryDomainByParamRequest, runtime *util.RuntimeOptions) (_result *QueryDomainByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDomainByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDomainByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryInvalidAddressWithOptions(request *QueryInvalidAddressRequest, runtime *util.RuntimeOptions) (_result *QueryInvalidAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryInvalidAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInvalidAddress"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryReceiverByParamWithOptions(request *QueryReceiverByParamRequest, runtime *util.RuntimeOptions) (_result *QueryReceiverByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryReceiverByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryReceiverByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryReceiverDetailWithOptions(request *QueryReceiverDetailRequest, runtime *util.RuntimeOptions) (_result *QueryReceiverDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryReceiverDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryReceiverDetail"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QuerySignByParamWithOptions(request *QuerySignByParamRequest, runtime *util.RuntimeOptions) (_result *QuerySignByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySignByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySignByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySignByParam(request *QuerySignByParamRequest) (_result *QuerySignByParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySignByParamResponse{}
	_body, _err := client.QuerySignByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySmsStatisticsWithOptions(request *QuerySmsStatisticsRequest, runtime *util.RuntimeOptions) (_result *QuerySmsStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySmsStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySmsStatistics"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySmsStatistics(request *QuerySmsStatisticsRequest) (_result *QuerySmsStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsStatisticsResponse{}
	_body, _err := client.QuerySmsStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTagByParamWithOptions(request *QueryTagByParamRequest, runtime *util.RuntimeOptions) (_result *QueryTagByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTagByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTagByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryTaskByParamWithOptions(request *QueryTaskByParamRequest, runtime *util.RuntimeOptions) (_result *QueryTaskByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTaskByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTaskByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) QueryTemplateByParamWithOptions(request *QueryTemplateByParamRequest, runtime *util.RuntimeOptions) (_result *QueryTemplateByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTemplateByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTemplateByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SaveReceiverDetailWithOptions(request *SaveReceiverDetailRequest, runtime *util.RuntimeOptions) (_result *SaveReceiverDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveReceiverDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveReceiverDetail"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SenderStatisticsByTagNameAndBatchIDWithOptions(request *SenderStatisticsByTagNameAndBatchIDRequest, runtime *util.RuntimeOptions) (_result *SenderStatisticsByTagNameAndBatchIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SenderStatisticsByTagNameAndBatchIDResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SenderStatisticsByTagNameAndBatchID"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SenderStatisticsDetailByParamWithOptions(request *SenderStatisticsDetailByParamRequest, runtime *util.RuntimeOptions) (_result *SenderStatisticsDetailByParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SenderStatisticsDetailByParamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SenderStatisticsDetailByParam"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SendTestByTemplateWithOptions(request *SendTestByTemplateRequest, runtime *util.RuntimeOptions) (_result *SendTestByTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendTestByTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendTestByTemplate"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SingleSendMailWithOptions(request *SingleSendMailRequest, runtime *util.RuntimeOptions) (_result *SingleSendMailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SingleSendMailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SingleSendMail"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SingleSendSmsWithOptions(request *SingleSendSmsRequest, runtime *util.RuntimeOptions) (_result *SingleSendSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SingleSendSmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SingleSendSms"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SingleSendSms(request *SingleSendSmsRequest) (_result *SingleSendSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SingleSendSmsResponse{}
	_body, _err := client.SingleSendSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDomainTrackNameWithOptions(request *UpdateDomainTrackNameRequest, runtime *util.RuntimeOptions) (_result *UpdateDomainTrackNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDomainTrackNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDomainTrackName"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDomainTrackName(request *UpdateDomainTrackNameRequest) (_result *UpdateDomainTrackNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDomainTrackNameResponse{}
	_body, _err := client.UpdateDomainTrackNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpProtectionWithOptions(request *UpdateIpProtectionRequest, runtime *util.RuntimeOptions) (_result *UpdateIpProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIpProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIpProtection"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateMailAddressMsgCallBackUrlWithOptions(request *UpdateMailAddressMsgCallBackUrlRequest, runtime *util.RuntimeOptions) (_result *UpdateMailAddressMsgCallBackUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMailAddressMsgCallBackUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMailAddressMsgCallBackUrl"), tea.String("2015-11-23"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMailAddressMsgCallBackUrl(request *UpdateMailAddressMsgCallBackUrlRequest) (_result *UpdateMailAddressMsgCallBackUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMailAddressMsgCallBackUrlResponse{}
	_body, _err := client.UpdateMailAddressMsgCallBackUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
