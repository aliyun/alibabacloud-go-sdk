// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateContactRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Locale      *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Position    *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s CreateContactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateContactRequest) SetClientToken(v string) *CreateContactRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContactRequest) SetContactName(v string) *CreateContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateContactRequest) SetEmail(v string) *CreateContactRequest {
	s.Email = &v
	return s
}

func (s *CreateContactRequest) SetLocale(v string) *CreateContactRequest {
	s.Locale = &v
	return s
}

func (s *CreateContactRequest) SetMobile(v string) *CreateContactRequest {
	s.Mobile = &v
	return s
}

func (s *CreateContactRequest) SetPosition(v string) *CreateContactRequest {
	s.Position = &v
	return s
}

type CreateContactResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactResponseBody) SetCode(v string) *CreateContactResponseBody {
	s.Code = &v
	return s
}

func (s *CreateContactResponseBody) SetContactId(v int64) *CreateContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateContactResponseBody) SetMessage(v string) *CreateContactResponseBody {
	s.Message = &v
	return s
}

func (s *CreateContactResponseBody) SetRequestId(v string) *CreateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContactResponseBody) SetSuccess(v bool) *CreateContactResponseBody {
	s.Success = &v
	return s
}

type CreateContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateContactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContactResponse) GoString() string {
	return s.String()
}

func (s *CreateContactResponse) SetHeaders(v map[string]*string) *CreateContactResponse {
	s.Headers = v
	return s
}

func (s *CreateContactResponse) SetStatusCode(v int32) *CreateContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContactResponse) SetBody(v *CreateContactResponseBody) *CreateContactResponse {
	s.Body = v
	return s
}

type CreateSubscriptionItemRequest struct {
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	Locale   *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s CreateSubscriptionItemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionItemRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionItemRequest) SetItemName(v string) *CreateSubscriptionItemRequest {
	s.ItemName = &v
	return s
}

func (s *CreateSubscriptionItemRequest) SetLocale(v string) *CreateSubscriptionItemRequest {
	s.Locale = &v
	return s
}

type CreateSubscriptionItemResponseBody struct {
	Code             *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionItem *CreateSubscriptionItemResponseBodySubscriptionItem `json:"SubscriptionItem,omitempty" xml:"SubscriptionItem,omitempty" type:"Struct"`
	Success          *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSubscriptionItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionItemResponseBody) SetCode(v string) *CreateSubscriptionItemResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubscriptionItemResponseBody) SetMessage(v string) *CreateSubscriptionItemResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSubscriptionItemResponseBody) SetRequestId(v string) *CreateSubscriptionItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionItemResponseBody) SetSubscriptionItem(v *CreateSubscriptionItemResponseBodySubscriptionItem) *CreateSubscriptionItemResponseBody {
	s.SubscriptionItem = v
	return s
}

func (s *CreateSubscriptionItemResponseBody) SetSuccess(v bool) *CreateSubscriptionItemResponseBody {
	s.Success = &v
	return s
}

type CreateSubscriptionItemResponseBodySubscriptionItem struct {
	Channel       *string  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ContactIds    []*int64 `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	Description   *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EmailStatus   *int32   `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId        *int32   `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName      *string  `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	PmsgStatus    *int32   `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	SmsStatus     *int32   `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus     *int32   `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookIds    []*int64 `json:"WebhookIds,omitempty" xml:"WebhookIds,omitempty" type:"Repeated"`
	WebhookStatus *int32   `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
}

func (s CreateSubscriptionItemResponseBodySubscriptionItem) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionItemResponseBodySubscriptionItem) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetChannel(v string) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.Channel = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetContactIds(v []*int64) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.ContactIds = v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetDescription(v string) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.Description = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetEmailStatus(v int32) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.EmailStatus = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetItemId(v int32) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.ItemId = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetItemName(v string) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.ItemName = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetPmsgStatus(v int32) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.PmsgStatus = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetSmsStatus(v int32) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.SmsStatus = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetTtsStatus(v int32) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.TtsStatus = &v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetWebhookIds(v []*int64) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.WebhookIds = v
	return s
}

func (s *CreateSubscriptionItemResponseBodySubscriptionItem) SetWebhookStatus(v int32) *CreateSubscriptionItemResponseBodySubscriptionItem {
	s.WebhookStatus = &v
	return s
}

type CreateSubscriptionItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubscriptionItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubscriptionItemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionItemResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionItemResponse) SetHeaders(v map[string]*string) *CreateSubscriptionItemResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionItemResponse) SetStatusCode(v int32) *CreateSubscriptionItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscriptionItemResponse) SetBody(v *CreateSubscriptionItemResponseBody) *CreateSubscriptionItemResponse {
	s.Body = v
	return s
}

type CreateWebhookRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Locale      *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	ServerUrl   *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s CreateWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWebhookRequest) GoString() string {
	return s.String()
}

func (s *CreateWebhookRequest) SetClientToken(v string) *CreateWebhookRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWebhookRequest) SetLocale(v string) *CreateWebhookRequest {
	s.Locale = &v
	return s
}

func (s *CreateWebhookRequest) SetServerUrl(v string) *CreateWebhookRequest {
	s.ServerUrl = &v
	return s
}

func (s *CreateWebhookRequest) SetWebhookName(v string) *CreateWebhookRequest {
	s.WebhookName = &v
	return s
}

type CreateWebhookResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	WebhookId *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s CreateWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebhookResponseBody) SetCode(v string) *CreateWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWebhookResponseBody) SetMessage(v string) *CreateWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWebhookResponseBody) SetRequestId(v string) *CreateWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebhookResponseBody) SetSuccess(v bool) *CreateWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWebhookResponseBody) SetWebhookId(v int64) *CreateWebhookResponseBody {
	s.WebhookId = &v
	return s
}

type CreateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWebhookResponse) GoString() string {
	return s.String()
}

func (s *CreateWebhookResponse) SetHeaders(v map[string]*string) *CreateWebhookResponse {
	s.Headers = v
	return s
}

func (s *CreateWebhookResponse) SetStatusCode(v int32) *CreateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebhookResponse) SetBody(v *CreateWebhookResponseBody) *CreateWebhookResponse {
	s.Body = v
	return s
}

type DeleteContactRequest struct {
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Locale    *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s DeleteContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactRequest) SetContactId(v int64) *DeleteContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteContactRequest) SetLocale(v string) *DeleteContactRequest {
	s.Locale = &v
	return s
}

type DeleteContactResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactResponseBody) SetCode(v string) *DeleteContactResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactResponseBody) SetMessage(v string) *DeleteContactResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactResponseBody) SetRequestId(v string) *DeleteContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactResponseBody) SetResult(v bool) *DeleteContactResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteContactResponseBody) SetSuccess(v bool) *DeleteContactResponseBody {
	s.Success = &v
	return s
}

type DeleteContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactResponse) SetHeaders(v map[string]*string) *DeleteContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactResponse) SetStatusCode(v int32) *DeleteContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactResponse) SetBody(v *DeleteContactResponseBody) *DeleteContactResponse {
	s.Body = v
	return s
}

type DeleteWebhookRequest struct {
	Locale    *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	WebhookId *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s DeleteWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebhookRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebhookRequest) SetLocale(v string) *DeleteWebhookRequest {
	s.Locale = &v
	return s
}

func (s *DeleteWebhookRequest) SetWebhookId(v int64) *DeleteWebhookRequest {
	s.WebhookId = &v
	return s
}

type DeleteWebhookResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebhookResponseBody) SetCode(v string) *DeleteWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWebhookResponseBody) SetMessage(v string) *DeleteWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWebhookResponseBody) SetRequestId(v string) *DeleteWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebhookResponseBody) SetResult(v bool) *DeleteWebhookResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteWebhookResponseBody) SetSuccess(v bool) *DeleteWebhookResponseBody {
	s.Success = &v
	return s
}

type DeleteWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebhookResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebhookResponse) SetHeaders(v map[string]*string) *DeleteWebhookResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebhookResponse) SetStatusCode(v int32) *DeleteWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebhookResponse) SetBody(v *DeleteWebhookResponseBody) *DeleteWebhookResponse {
	s.Body = v
	return s
}

type GetContactRequest struct {
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Locale    *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s GetContactRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContactRequest) GoString() string {
	return s.String()
}

func (s *GetContactRequest) SetContactId(v int64) *GetContactRequest {
	s.ContactId = &v
	return s
}

func (s *GetContactRequest) SetLocale(v string) *GetContactRequest {
	s.Locale = &v
	return s
}

type GetContactResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Contact   *GetContactResponseBodyContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContactResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactResponseBody) SetCode(v string) *GetContactResponseBody {
	s.Code = &v
	return s
}

func (s *GetContactResponseBody) SetContact(v *GetContactResponseBodyContact) *GetContactResponseBody {
	s.Contact = v
	return s
}

func (s *GetContactResponseBody) SetMessage(v string) *GetContactResponseBody {
	s.Message = &v
	return s
}

func (s *GetContactResponseBody) SetRequestId(v string) *GetContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContactResponseBody) SetSuccess(v bool) *GetContactResponseBody {
	s.Success = &v
	return s
}

type GetContactResponseBodyContact struct {
	AccountUid                      *int64  `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	ContactId                       *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName                     *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email                           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IsAccount                       *bool   `json:"IsAccount,omitempty" xml:"IsAccount,omitempty"`
	IsObsolete                      *bool   `json:"IsObsolete,omitempty" xml:"IsObsolete,omitempty"`
	IsVerifiedEmail                 *bool   `json:"IsVerifiedEmail,omitempty" xml:"IsVerifiedEmail,omitempty"`
	IsVerifiedMobile                *bool   `json:"IsVerifiedMobile,omitempty" xml:"IsVerifiedMobile,omitempty"`
	LastEmailVerificationTimeStamp  *int64  `json:"LastEmailVerificationTimeStamp,omitempty" xml:"LastEmailVerificationTimeStamp,omitempty"`
	LastMobileVerificationTimeStamp *int64  `json:"LastMobileVerificationTimeStamp,omitempty" xml:"LastMobileVerificationTimeStamp,omitempty"`
	Mobile                          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Position                        *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetContactResponseBodyContact) String() string {
	return tea.Prettify(s)
}

func (s GetContactResponseBodyContact) GoString() string {
	return s.String()
}

func (s *GetContactResponseBodyContact) SetAccountUid(v int64) *GetContactResponseBodyContact {
	s.AccountUid = &v
	return s
}

func (s *GetContactResponseBodyContact) SetContactId(v int64) *GetContactResponseBodyContact {
	s.ContactId = &v
	return s
}

func (s *GetContactResponseBodyContact) SetContactName(v string) *GetContactResponseBodyContact {
	s.ContactName = &v
	return s
}

func (s *GetContactResponseBodyContact) SetEmail(v string) *GetContactResponseBodyContact {
	s.Email = &v
	return s
}

func (s *GetContactResponseBodyContact) SetIsAccount(v bool) *GetContactResponseBodyContact {
	s.IsAccount = &v
	return s
}

func (s *GetContactResponseBodyContact) SetIsObsolete(v bool) *GetContactResponseBodyContact {
	s.IsObsolete = &v
	return s
}

func (s *GetContactResponseBodyContact) SetIsVerifiedEmail(v bool) *GetContactResponseBodyContact {
	s.IsVerifiedEmail = &v
	return s
}

func (s *GetContactResponseBodyContact) SetIsVerifiedMobile(v bool) *GetContactResponseBodyContact {
	s.IsVerifiedMobile = &v
	return s
}

func (s *GetContactResponseBodyContact) SetLastEmailVerificationTimeStamp(v int64) *GetContactResponseBodyContact {
	s.LastEmailVerificationTimeStamp = &v
	return s
}

func (s *GetContactResponseBodyContact) SetLastMobileVerificationTimeStamp(v int64) *GetContactResponseBodyContact {
	s.LastMobileVerificationTimeStamp = &v
	return s
}

func (s *GetContactResponseBodyContact) SetMobile(v string) *GetContactResponseBodyContact {
	s.Mobile = &v
	return s
}

func (s *GetContactResponseBodyContact) SetPosition(v string) *GetContactResponseBodyContact {
	s.Position = &v
	return s
}

type GetContactResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContactResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContactResponse) GoString() string {
	return s.String()
}

func (s *GetContactResponse) SetHeaders(v map[string]*string) *GetContactResponse {
	s.Headers = v
	return s
}

func (s *GetContactResponse) SetStatusCode(v int32) *GetContactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactResponse) SetBody(v *GetContactResponseBody) *GetContactResponse {
	s.Body = v
	return s
}

type GetSubscriptionItemRequest struct {
	ItemId *int32  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s GetSubscriptionItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemRequest) SetItemId(v int32) *GetSubscriptionItemRequest {
	s.ItemId = &v
	return s
}

func (s *GetSubscriptionItemRequest) SetLocale(v string) *GetSubscriptionItemRequest {
	s.Locale = &v
	return s
}

type GetSubscriptionItemResponseBody struct {
	Code             *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionItem *GetSubscriptionItemResponseBodySubscriptionItem `json:"SubscriptionItem,omitempty" xml:"SubscriptionItem,omitempty" type:"Struct"`
	Success          *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubscriptionItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemResponseBody) SetCode(v string) *GetSubscriptionItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionItemResponseBody) SetMessage(v string) *GetSubscriptionItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionItemResponseBody) SetRequestId(v string) *GetSubscriptionItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionItemResponseBody) SetSubscriptionItem(v *GetSubscriptionItemResponseBodySubscriptionItem) *GetSubscriptionItemResponseBody {
	s.SubscriptionItem = v
	return s
}

func (s *GetSubscriptionItemResponseBody) SetSuccess(v bool) *GetSubscriptionItemResponseBody {
	s.Success = &v
	return s
}

type GetSubscriptionItemResponseBodySubscriptionItem struct {
	Channel       *string  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ContactIds    []*int64 `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	Description   *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EmailStatus   *int32   `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId        *int32   `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName      *string  `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	PmsgStatus    *int32   `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	SmsStatus     *int32   `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus     *int32   `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookIds    []*int64 `json:"WebhookIds,omitempty" xml:"WebhookIds,omitempty" type:"Repeated"`
	WebhookStatus *int32   `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
}

func (s GetSubscriptionItemResponseBodySubscriptionItem) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemResponseBodySubscriptionItem) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetChannel(v string) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.Channel = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetContactIds(v []*int64) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.ContactIds = v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetDescription(v string) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.Description = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetEmailStatus(v int32) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.EmailStatus = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetItemId(v int32) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.ItemId = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetItemName(v string) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.ItemName = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetPmsgStatus(v int32) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.PmsgStatus = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetSmsStatus(v int32) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.SmsStatus = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetTtsStatus(v int32) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.TtsStatus = &v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetWebhookIds(v []*int64) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.WebhookIds = v
	return s
}

func (s *GetSubscriptionItemResponseBodySubscriptionItem) SetWebhookStatus(v int32) *GetSubscriptionItemResponseBodySubscriptionItem {
	s.WebhookStatus = &v
	return s
}

type GetSubscriptionItemResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSubscriptionItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubscriptionItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemResponse) SetHeaders(v map[string]*string) *GetSubscriptionItemResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionItemResponse) SetStatusCode(v int32) *GetSubscriptionItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionItemResponse) SetBody(v *GetSubscriptionItemResponseBody) *GetSubscriptionItemResponse {
	s.Body = v
	return s
}

type GetSubscriptionItemDetailRequest struct {
	ItemId *int32  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s GetSubscriptionItemDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemDetailRequest) SetItemId(v int32) *GetSubscriptionItemDetailRequest {
	s.ItemId = &v
	return s
}

func (s *GetSubscriptionItemDetailRequest) SetLocale(v string) *GetSubscriptionItemDetailRequest {
	s.Locale = &v
	return s
}

type GetSubscriptionItemDetailResponseBody struct {
	Code                   *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionItemDetail *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail `json:"SubscriptionItemDetail,omitempty" xml:"SubscriptionItemDetail,omitempty" type:"Struct"`
	Success                *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubscriptionItemDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemDetailResponseBody) SetCode(v string) *GetSubscriptionItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBody) SetMessage(v string) *GetSubscriptionItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBody) SetRequestId(v string) *GetSubscriptionItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBody) SetSubscriptionItemDetail(v *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) *GetSubscriptionItemDetailResponseBody {
	s.SubscriptionItemDetail = v
	return s
}

func (s *GetSubscriptionItemDetailResponseBody) SetSuccess(v bool) *GetSubscriptionItemDetailResponseBody {
	s.Success = &v
	return s
}

type GetSubscriptionItemDetailResponseBodySubscriptionItemDetail struct {
	Channel       *string                                                                `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Contacts      []*GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Description   *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EmailStatus   *int32                                                                 `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId        *int32                                                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName      *string                                                                `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	PmsgStatus    *int32                                                                 `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	RegionId      *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SmsStatus     *int32                                                                 `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus     *int32                                                                 `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookStatus *int32                                                                 `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
	Webhooks      []*GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks `json:"Webhooks,omitempty" xml:"Webhooks,omitempty" type:"Repeated"`
}

func (s GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetChannel(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.Channel = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetContacts(v []*GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.Contacts = v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetDescription(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.Description = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetEmailStatus(v int32) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.EmailStatus = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetItemId(v int32) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.ItemId = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetItemName(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.ItemName = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetPmsgStatus(v int32) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.PmsgStatus = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetRegionId(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.RegionId = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetSmsStatus(v int32) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.SmsStatus = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetTtsStatus(v int32) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.TtsStatus = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetWebhookStatus(v int32) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.WebhookStatus = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail) SetWebhooks(v []*GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetail {
	s.Webhooks = v
	return s
}

type GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts struct {
	AccountUID                      *int64  `json:"AccountUID,omitempty" xml:"AccountUID,omitempty"`
	ContactId                       *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Email                           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IsAccount                       *bool   `json:"IsAccount,omitempty" xml:"IsAccount,omitempty"`
	IsObsolete                      *bool   `json:"IsObsolete,omitempty" xml:"IsObsolete,omitempty"`
	IsVerifiedEmail                 *bool   `json:"IsVerifiedEmail,omitempty" xml:"IsVerifiedEmail,omitempty"`
	IsVerifiedMobile                *bool   `json:"IsVerifiedMobile,omitempty" xml:"IsVerifiedMobile,omitempty"`
	LastEmailVerificationTimeStamp  *int64  `json:"LastEmailVerificationTimeStamp,omitempty" xml:"LastEmailVerificationTimeStamp,omitempty"`
	LastMobileVerificationTimeStamp *int64  `json:"LastMobileVerificationTimeStamp,omitempty" xml:"LastMobileVerificationTimeStamp,omitempty"`
	Mobile                          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Name                            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Position                        *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetAccountUID(v int64) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.AccountUID = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetContactId(v int64) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.ContactId = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetEmail(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.Email = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetIsAccount(v bool) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.IsAccount = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetIsObsolete(v bool) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.IsObsolete = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetIsVerifiedEmail(v bool) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.IsVerifiedEmail = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetIsVerifiedMobile(v bool) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.IsVerifiedMobile = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetLastEmailVerificationTimeStamp(v int64) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.LastEmailVerificationTimeStamp = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetLastMobileVerificationTimeStamp(v int64) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.LastMobileVerificationTimeStamp = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetMobile(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.Mobile = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetName(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.Name = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts) SetPosition(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailContacts {
	s.Position = &v
	return s
}

type GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	WebhookId *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks) SetName(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks {
	s.Name = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks) SetServerUrl(v string) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks {
	s.ServerUrl = &v
	return s
}

func (s *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks) SetWebhookId(v int64) *GetSubscriptionItemDetailResponseBodySubscriptionItemDetailWebhooks {
	s.WebhookId = &v
	return s
}

type GetSubscriptionItemDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSubscriptionItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubscriptionItemDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionItemDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionItemDetailResponse) SetHeaders(v map[string]*string) *GetSubscriptionItemDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionItemDetailResponse) SetStatusCode(v int32) *GetSubscriptionItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionItemDetailResponse) SetBody(v *GetSubscriptionItemDetailResponseBody) *GetSubscriptionItemDetailResponse {
	s.Body = v
	return s
}

type GetWebhookRequest struct {
	Locale    *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	WebhookId *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s GetWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebhookRequest) GoString() string {
	return s.String()
}

func (s *GetWebhookRequest) SetLocale(v string) *GetWebhookRequest {
	s.Locale = &v
	return s
}

func (s *GetWebhookRequest) SetWebhookId(v int64) *GetWebhookRequest {
	s.WebhookId = &v
	return s
}

type GetWebhookResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Webhook   *GetWebhookResponseBodyWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Struct"`
}

func (s GetWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebhookResponseBody) SetCode(v string) *GetWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *GetWebhookResponseBody) SetMessage(v string) *GetWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *GetWebhookResponseBody) SetRequestId(v string) *GetWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebhookResponseBody) SetSuccess(v bool) *GetWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *GetWebhookResponseBody) SetWebhook(v *GetWebhookResponseBodyWebhook) *GetWebhookResponseBody {
	s.Webhook = v
	return s
}

type GetWebhookResponseBodyWebhook struct {
	ServerUrl   *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	WebhookId   *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s GetWebhookResponseBodyWebhook) String() string {
	return tea.Prettify(s)
}

func (s GetWebhookResponseBodyWebhook) GoString() string {
	return s.String()
}

func (s *GetWebhookResponseBodyWebhook) SetServerUrl(v string) *GetWebhookResponseBodyWebhook {
	s.ServerUrl = &v
	return s
}

func (s *GetWebhookResponseBodyWebhook) SetWebhookId(v int64) *GetWebhookResponseBodyWebhook {
	s.WebhookId = &v
	return s
}

func (s *GetWebhookResponseBodyWebhook) SetWebhookName(v string) *GetWebhookResponseBodyWebhook {
	s.WebhookName = &v
	return s
}

type GetWebhookResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebhookResponse) GoString() string {
	return s.String()
}

func (s *GetWebhookResponse) SetHeaders(v map[string]*string) *GetWebhookResponse {
	s.Headers = v
	return s
}

func (s *GetWebhookResponse) SetStatusCode(v int32) *GetWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebhookResponse) SetBody(v *GetWebhookResponseBody) *GetWebhookResponse {
	s.Body = v
	return s
}

type ListContactsRequest struct {
	ContactId  *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Locale     *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactsRequest) GoString() string {
	return s.String()
}

func (s *ListContactsRequest) SetContactId(v int64) *ListContactsRequest {
	s.ContactId = &v
	return s
}

func (s *ListContactsRequest) SetFilter(v string) *ListContactsRequest {
	s.Filter = &v
	return s
}

func (s *ListContactsRequest) SetLocale(v string) *ListContactsRequest {
	s.Locale = &v
	return s
}

func (s *ListContactsRequest) SetMaxResults(v int32) *ListContactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContactsRequest) SetNextToken(v string) *ListContactsRequest {
	s.NextToken = &v
	return s
}

type ListContactsResponseBody struct {
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Contacts   []*ListContactsResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *int32                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBody) SetCode(v string) *ListContactsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContactsResponseBody) SetContacts(v []*ListContactsResponseBodyContacts) *ListContactsResponseBody {
	s.Contacts = v
	return s
}

func (s *ListContactsResponseBody) SetMessage(v string) *ListContactsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContactsResponseBody) SetNextToken(v int32) *ListContactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContactsResponseBody) SetRequestId(v string) *ListContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactsResponseBody) SetSuccess(v bool) *ListContactsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContactsResponseBody) SetTotalCount(v int32) *ListContactsResponseBody {
	s.TotalCount = &v
	return s
}

type ListContactsResponseBodyContacts struct {
	AccountUid                      *int64  `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	ContactId                       *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName                     *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email                           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IsAccount                       *bool   `json:"IsAccount,omitempty" xml:"IsAccount,omitempty"`
	IsObsolete                      *bool   `json:"IsObsolete,omitempty" xml:"IsObsolete,omitempty"`
	IsVerifiedEmail                 *bool   `json:"IsVerifiedEmail,omitempty" xml:"IsVerifiedEmail,omitempty"`
	IsVerifiedMobile                *bool   `json:"IsVerifiedMobile,omitempty" xml:"IsVerifiedMobile,omitempty"`
	LastEmailVerificationTimeStamp  *int64  `json:"LastEmailVerificationTimeStamp,omitempty" xml:"LastEmailVerificationTimeStamp,omitempty"`
	LastMobileVerificationTimeStamp *int64  `json:"LastMobileVerificationTimeStamp,omitempty" xml:"LastMobileVerificationTimeStamp,omitempty"`
	Mobile                          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Position                        *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ListContactsResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s ListContactsResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBodyContacts) SetAccountUid(v int64) *ListContactsResponseBodyContacts {
	s.AccountUid = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetContactId(v int64) *ListContactsResponseBodyContacts {
	s.ContactId = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetContactName(v string) *ListContactsResponseBodyContacts {
	s.ContactName = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetEmail(v string) *ListContactsResponseBodyContacts {
	s.Email = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetIsAccount(v bool) *ListContactsResponseBodyContacts {
	s.IsAccount = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetIsObsolete(v bool) *ListContactsResponseBodyContacts {
	s.IsObsolete = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetIsVerifiedEmail(v bool) *ListContactsResponseBodyContacts {
	s.IsVerifiedEmail = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetIsVerifiedMobile(v bool) *ListContactsResponseBodyContacts {
	s.IsVerifiedMobile = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetLastEmailVerificationTimeStamp(v int64) *ListContactsResponseBodyContacts {
	s.LastEmailVerificationTimeStamp = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetLastMobileVerificationTimeStamp(v int64) *ListContactsResponseBodyContacts {
	s.LastMobileVerificationTimeStamp = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetMobile(v string) *ListContactsResponseBodyContacts {
	s.Mobile = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetPosition(v string) *ListContactsResponseBodyContacts {
	s.Position = &v
	return s
}

type ListContactsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactsResponse) GoString() string {
	return s.String()
}

func (s *ListContactsResponse) SetHeaders(v map[string]*string) *ListContactsResponse {
	s.Headers = v
	return s
}

func (s *ListContactsResponse) SetStatusCode(v int32) *ListContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactsResponse) SetBody(v *ListContactsResponseBody) *ListContactsResponse {
	s.Body = v
	return s
}

type ListSubscriptionItemGroupDetailsRequest struct {
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s ListSubscriptionItemGroupDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsRequest) SetLocale(v string) *ListSubscriptionItemGroupDetailsRequest {
	s.Locale = &v
	return s
}

type ListSubscriptionItemGroupDetailsResponseBody struct {
	Code                         *string                                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                      *string                                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                    *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionItemGroupDetails []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails `json:"SubscriptionItemGroupDetails,omitempty" xml:"SubscriptionItemGroupDetails,omitempty" type:"Repeated"`
	Success                      *bool                                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSubscriptionItemGroupDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsResponseBody) SetCode(v string) *ListSubscriptionItemGroupDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBody) SetMessage(v string) *ListSubscriptionItemGroupDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBody) SetRequestId(v string) *ListSubscriptionItemGroupDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBody) SetSubscriptionItemGroupDetails(v []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) *ListSubscriptionItemGroupDetailsResponseBody {
	s.SubscriptionItemGroupDetails = v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBody) SetSuccess(v bool) *ListSubscriptionItemGroupDetailsResponseBody {
	s.Success = &v
	return s
}

type ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails struct {
	Description   *string                                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ItemDetails   []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails `json:"ItemDetails,omitempty" xml:"ItemDetails,omitempty" type:"Repeated"`
	ItemGroupId   *int32                                                                                 `json:"ItemGroupId,omitempty" xml:"ItemGroupId,omitempty"`
	ItemGroupName *string                                                                                `json:"ItemGroupName,omitempty" xml:"ItemGroupName,omitempty"`
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) SetDescription(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails {
	s.Description = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) SetItemDetails(v []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails {
	s.ItemDetails = v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) SetItemGroupId(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails {
	s.ItemGroupId = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails) SetItemGroupName(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetails {
	s.ItemGroupName = &v
	return s
}

type ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails struct {
	Channel       *string                                                                                        `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Contacts      []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Description   *string                                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EmailStatus   *int32                                                                                         `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId        *int32                                                                                         `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName      *string                                                                                        `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	PmsgStatus    *int32                                                                                         `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	RegionId      *string                                                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SmsStatus     *int32                                                                                         `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus     *int32                                                                                         `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookStatus *int32                                                                                         `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
	Webhooks      []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks `json:"Webhooks,omitempty" xml:"Webhooks,omitempty" type:"Repeated"`
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetChannel(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.Channel = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetContacts(v []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.Contacts = v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetDescription(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.Description = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetEmailStatus(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.EmailStatus = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetItemId(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.ItemId = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetItemName(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.ItemName = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetPmsgStatus(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.PmsgStatus = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetRegionId(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.RegionId = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetSmsStatus(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.SmsStatus = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetTtsStatus(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.TtsStatus = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetWebhookStatus(v int32) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.WebhookStatus = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails) SetWebhooks(v []*ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetails {
	s.Webhooks = v
	return s
}

type ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts struct {
	AccountUID                      *int64  `json:"AccountUID,omitempty" xml:"AccountUID,omitempty"`
	ContactId                       *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Email                           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IsAccount                       *bool   `json:"IsAccount,omitempty" xml:"IsAccount,omitempty"`
	IsObsolete                      *bool   `json:"IsObsolete,omitempty" xml:"IsObsolete,omitempty"`
	IsVerifiedEmail                 *bool   `json:"IsVerifiedEmail,omitempty" xml:"IsVerifiedEmail,omitempty"`
	IsVerifiedMobile                *bool   `json:"IsVerifiedMobile,omitempty" xml:"IsVerifiedMobile,omitempty"`
	LastEmailVerificationTimeStamp  *int64  `json:"LastEmailVerificationTimeStamp,omitempty" xml:"LastEmailVerificationTimeStamp,omitempty"`
	LastMobileVerificationTimeStamp *int64  `json:"LastMobileVerificationTimeStamp,omitempty" xml:"LastMobileVerificationTimeStamp,omitempty"`
	Mobile                          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Name                            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Position                        *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetAccountUID(v int64) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.AccountUID = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetContactId(v int64) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.ContactId = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetEmail(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.Email = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetIsAccount(v bool) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.IsAccount = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetIsObsolete(v bool) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.IsObsolete = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetIsVerifiedEmail(v bool) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.IsVerifiedEmail = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetIsVerifiedMobile(v bool) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.IsVerifiedMobile = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetLastEmailVerificationTimeStamp(v int64) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.LastEmailVerificationTimeStamp = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetLastMobileVerificationTimeStamp(v int64) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.LastMobileVerificationTimeStamp = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetMobile(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.Mobile = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetName(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.Name = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts) SetPosition(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsContacts {
	s.Position = &v
	return s
}

type ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	WebhookId *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks) SetName(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks {
	s.Name = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks) SetServerUrl(v string) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks {
	s.ServerUrl = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks) SetWebhookId(v int64) *ListSubscriptionItemGroupDetailsResponseBodySubscriptionItemGroupDetailsItemDetailsWebhooks {
	s.WebhookId = &v
	return s
}

type ListSubscriptionItemGroupDetailsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubscriptionItemGroupDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubscriptionItemGroupDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemGroupDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemGroupDetailsResponse) SetHeaders(v map[string]*string) *ListSubscriptionItemGroupDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponse) SetStatusCode(v int32) *ListSubscriptionItemGroupDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionItemGroupDetailsResponse) SetBody(v *ListSubscriptionItemGroupDetailsResponseBody) *ListSubscriptionItemGroupDetailsResponse {
	s.Body = v
	return s
}

type ListSubscriptionItemsRequest struct {
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	ItemId     *int32  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Locale     *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSubscriptionItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemsRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemsRequest) SetFilter(v string) *ListSubscriptionItemsRequest {
	s.Filter = &v
	return s
}

func (s *ListSubscriptionItemsRequest) SetItemId(v int32) *ListSubscriptionItemsRequest {
	s.ItemId = &v
	return s
}

func (s *ListSubscriptionItemsRequest) SetLocale(v string) *ListSubscriptionItemsRequest {
	s.Locale = &v
	return s
}

func (s *ListSubscriptionItemsRequest) SetMaxResults(v int32) *ListSubscriptionItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSubscriptionItemsRequest) SetNextToken(v string) *ListSubscriptionItemsRequest {
	s.NextToken = &v
	return s
}

type ListSubscriptionItemsResponseBody struct {
	Code              *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message           *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken         *int32                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionItems []*ListSubscriptionItemsResponseBodySubscriptionItems `json:"SubscriptionItems,omitempty" xml:"SubscriptionItems,omitempty" type:"Repeated"`
	Success           *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubscriptionItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemsResponseBody) SetCode(v string) *ListSubscriptionItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionItemsResponseBody) SetMessage(v string) *ListSubscriptionItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionItemsResponseBody) SetNextToken(v int32) *ListSubscriptionItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSubscriptionItemsResponseBody) SetRequestId(v string) *ListSubscriptionItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionItemsResponseBody) SetSubscriptionItems(v []*ListSubscriptionItemsResponseBodySubscriptionItems) *ListSubscriptionItemsResponseBody {
	s.SubscriptionItems = v
	return s
}

func (s *ListSubscriptionItemsResponseBody) SetSuccess(v bool) *ListSubscriptionItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubscriptionItemsResponseBody) SetTotalCount(v int32) *ListSubscriptionItemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSubscriptionItemsResponseBodySubscriptionItems struct {
	Channel       *string  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ContactIds    []*int64 `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	Description   *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EmailStatus   *int32   `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId        *int32   `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName      *string  `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	PmsgStatus    *int32   `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	SmsStatus     *int32   `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus     *int32   `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookIds    []*int64 `json:"WebhookIds,omitempty" xml:"WebhookIds,omitempty" type:"Repeated"`
	WebhookStatus *int32   `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
}

func (s ListSubscriptionItemsResponseBodySubscriptionItems) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemsResponseBodySubscriptionItems) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetChannel(v string) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.Channel = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetContactIds(v []*int64) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.ContactIds = v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetDescription(v string) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.Description = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetEmailStatus(v int32) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.EmailStatus = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetItemId(v int32) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.ItemId = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetItemName(v string) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.ItemName = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetPmsgStatus(v int32) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.PmsgStatus = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetSmsStatus(v int32) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.SmsStatus = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetTtsStatus(v int32) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.TtsStatus = &v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetWebhookIds(v []*int64) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.WebhookIds = v
	return s
}

func (s *ListSubscriptionItemsResponseBodySubscriptionItems) SetWebhookStatus(v int32) *ListSubscriptionItemsResponseBodySubscriptionItems {
	s.WebhookStatus = &v
	return s
}

type ListSubscriptionItemsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubscriptionItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubscriptionItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionItemsResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionItemsResponse) SetHeaders(v map[string]*string) *ListSubscriptionItemsResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionItemsResponse) SetStatusCode(v int32) *ListSubscriptionItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionItemsResponse) SetBody(v *ListSubscriptionItemsResponseBody) *ListSubscriptionItemsResponse {
	s.Body = v
	return s
}

type ListWebhooksRequest struct {
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Locale     *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	WebhookId  *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s ListWebhooksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWebhooksRequest) GoString() string {
	return s.String()
}

func (s *ListWebhooksRequest) SetFilter(v string) *ListWebhooksRequest {
	s.Filter = &v
	return s
}

func (s *ListWebhooksRequest) SetLocale(v string) *ListWebhooksRequest {
	s.Locale = &v
	return s
}

func (s *ListWebhooksRequest) SetMaxResults(v int32) *ListWebhooksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWebhooksRequest) SetNextToken(v string) *ListWebhooksRequest {
	s.NextToken = &v
	return s
}

func (s *ListWebhooksRequest) SetWebhookId(v int64) *ListWebhooksRequest {
	s.WebhookId = &v
	return s
}

type ListWebhooksResponseBody struct {
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *int32                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Webhooks   []*ListWebhooksResponseBodyWebhooks `json:"Webhooks,omitempty" xml:"Webhooks,omitempty" type:"Repeated"`
}

func (s ListWebhooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWebhooksResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebhooksResponseBody) SetCode(v string) *ListWebhooksResponseBody {
	s.Code = &v
	return s
}

func (s *ListWebhooksResponseBody) SetMessage(v string) *ListWebhooksResponseBody {
	s.Message = &v
	return s
}

func (s *ListWebhooksResponseBody) SetNextToken(v int32) *ListWebhooksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWebhooksResponseBody) SetRequestId(v string) *ListWebhooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebhooksResponseBody) SetSuccess(v bool) *ListWebhooksResponseBody {
	s.Success = &v
	return s
}

func (s *ListWebhooksResponseBody) SetTotalCount(v int32) *ListWebhooksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWebhooksResponseBody) SetWebhooks(v []*ListWebhooksResponseBodyWebhooks) *ListWebhooksResponseBody {
	s.Webhooks = v
	return s
}

type ListWebhooksResponseBodyWebhooks struct {
	ServerUrl   *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	WebhookId   *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s ListWebhooksResponseBodyWebhooks) String() string {
	return tea.Prettify(s)
}

func (s ListWebhooksResponseBodyWebhooks) GoString() string {
	return s.String()
}

func (s *ListWebhooksResponseBodyWebhooks) SetServerUrl(v string) *ListWebhooksResponseBodyWebhooks {
	s.ServerUrl = &v
	return s
}

func (s *ListWebhooksResponseBodyWebhooks) SetWebhookId(v int64) *ListWebhooksResponseBodyWebhooks {
	s.WebhookId = &v
	return s
}

func (s *ListWebhooksResponseBodyWebhooks) SetWebhookName(v string) *ListWebhooksResponseBodyWebhooks {
	s.WebhookName = &v
	return s
}

type ListWebhooksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWebhooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWebhooksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWebhooksResponse) GoString() string {
	return s.String()
}

func (s *ListWebhooksResponse) SetHeaders(v map[string]*string) *ListWebhooksResponse {
	s.Headers = v
	return s
}

func (s *ListWebhooksResponse) SetStatusCode(v int32) *ListWebhooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebhooksResponse) SetBody(v *ListWebhooksResponseBody) *ListWebhooksResponse {
	s.Body = v
	return s
}

type SendVerificationMessageRequest struct {
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Locale    *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendVerificationMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationMessageRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationMessageRequest) SetContactId(v int64) *SendVerificationMessageRequest {
	s.ContactId = &v
	return s
}

func (s *SendVerificationMessageRequest) SetLocale(v string) *SendVerificationMessageRequest {
	s.Locale = &v
	return s
}

func (s *SendVerificationMessageRequest) SetType(v int32) *SendVerificationMessageRequest {
	s.Type = &v
	return s
}

type SendVerificationMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendVerificationMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationMessageResponseBody) SetCode(v string) *SendVerificationMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendVerificationMessageResponseBody) SetMessage(v string) *SendVerificationMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendVerificationMessageResponseBody) SetRequestId(v string) *SendVerificationMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerificationMessageResponseBody) SetResult(v int32) *SendVerificationMessageResponseBody {
	s.Result = &v
	return s
}

func (s *SendVerificationMessageResponseBody) SetSuccess(v bool) *SendVerificationMessageResponseBody {
	s.Success = &v
	return s
}

type SendVerificationMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendVerificationMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendVerificationMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationMessageResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationMessageResponse) SetHeaders(v map[string]*string) *SendVerificationMessageResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationMessageResponse) SetStatusCode(v int32) *SendVerificationMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationMessageResponse) SetBody(v *SendVerificationMessageResponseBody) *SendVerificationMessageResponse {
	s.Body = v
	return s
}

type UpdateContactRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Locale      *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s UpdateContactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactRequest) SetClientToken(v string) *UpdateContactRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContactRequest) SetContactId(v int64) *UpdateContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateContactRequest) SetContactName(v string) *UpdateContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateContactRequest) SetEmail(v string) *UpdateContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateContactRequest) SetLocale(v string) *UpdateContactRequest {
	s.Locale = &v
	return s
}

func (s *UpdateContactRequest) SetMobile(v string) *UpdateContactRequest {
	s.Mobile = &v
	return s
}

type UpdateContactResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactResponseBody) SetCode(v string) *UpdateContactResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContactResponseBody) SetMessage(v string) *UpdateContactResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContactResponseBody) SetRequestId(v string) *UpdateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactResponseBody) SetResult(v bool) *UpdateContactResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateContactResponseBody) SetSuccess(v bool) *UpdateContactResponseBody {
	s.Success = &v
	return s
}

type UpdateContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactResponse) SetHeaders(v map[string]*string) *UpdateContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactResponse) SetStatusCode(v int32) *UpdateContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactResponse) SetBody(v *UpdateContactResponseBody) *UpdateContactResponse {
	s.Body = v
	return s
}

type UpdateSubscriptionItemRequest struct {
	ClientToken   *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContactIds    []*int64 `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	EmailStatus   *int32   `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId        *int32   `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Locale        *string  `json:"Locale,omitempty" xml:"Locale,omitempty"`
	PmsgStatus    *int32   `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	RegionId      *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SmsStatus     *int32   `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus     *int32   `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookIds    []*int64 `json:"WebhookIds,omitempty" xml:"WebhookIds,omitempty" type:"Repeated"`
	WebhookStatus *int32   `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
}

func (s UpdateSubscriptionItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionItemRequest) SetClientToken(v string) *UpdateSubscriptionItemRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetContactIds(v []*int64) *UpdateSubscriptionItemRequest {
	s.ContactIds = v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetEmailStatus(v int32) *UpdateSubscriptionItemRequest {
	s.EmailStatus = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetItemId(v int32) *UpdateSubscriptionItemRequest {
	s.ItemId = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetLocale(v string) *UpdateSubscriptionItemRequest {
	s.Locale = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetPmsgStatus(v int32) *UpdateSubscriptionItemRequest {
	s.PmsgStatus = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetRegionId(v string) *UpdateSubscriptionItemRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetSmsStatus(v int32) *UpdateSubscriptionItemRequest {
	s.SmsStatus = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetTtsStatus(v int32) *UpdateSubscriptionItemRequest {
	s.TtsStatus = &v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetWebhookIds(v []*int64) *UpdateSubscriptionItemRequest {
	s.WebhookIds = v
	return s
}

func (s *UpdateSubscriptionItemRequest) SetWebhookStatus(v int32) *UpdateSubscriptionItemRequest {
	s.WebhookStatus = &v
	return s
}

type UpdateSubscriptionItemShrinkRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContactIdsShrink *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	EmailStatus      *int32  `json:"EmailStatus,omitempty" xml:"EmailStatus,omitempty"`
	ItemId           *int32  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	PmsgStatus       *int32  `json:"PmsgStatus,omitempty" xml:"PmsgStatus,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SmsStatus        *int32  `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	TtsStatus        *int32  `json:"TtsStatus,omitempty" xml:"TtsStatus,omitempty"`
	WebhookIdsShrink *string `json:"WebhookIds,omitempty" xml:"WebhookIds,omitempty"`
	WebhookStatus    *int32  `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
}

func (s UpdateSubscriptionItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionItemShrinkRequest) SetClientToken(v string) *UpdateSubscriptionItemShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetContactIdsShrink(v string) *UpdateSubscriptionItemShrinkRequest {
	s.ContactIdsShrink = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetEmailStatus(v int32) *UpdateSubscriptionItemShrinkRequest {
	s.EmailStatus = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetItemId(v int32) *UpdateSubscriptionItemShrinkRequest {
	s.ItemId = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetLocale(v string) *UpdateSubscriptionItemShrinkRequest {
	s.Locale = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetPmsgStatus(v int32) *UpdateSubscriptionItemShrinkRequest {
	s.PmsgStatus = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetRegionId(v string) *UpdateSubscriptionItemShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetSmsStatus(v int32) *UpdateSubscriptionItemShrinkRequest {
	s.SmsStatus = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetTtsStatus(v int32) *UpdateSubscriptionItemShrinkRequest {
	s.TtsStatus = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetWebhookIdsShrink(v string) *UpdateSubscriptionItemShrinkRequest {
	s.WebhookIdsShrink = &v
	return s
}

func (s *UpdateSubscriptionItemShrinkRequest) SetWebhookStatus(v int32) *UpdateSubscriptionItemShrinkRequest {
	s.WebhookStatus = &v
	return s
}

type UpdateSubscriptionItemResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSubscriptionItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionItemResponseBody) SetCode(v string) *UpdateSubscriptionItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubscriptionItemResponseBody) SetMessage(v string) *UpdateSubscriptionItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubscriptionItemResponseBody) SetRequestId(v string) *UpdateSubscriptionItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubscriptionItemResponseBody) SetResult(v bool) *UpdateSubscriptionItemResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateSubscriptionItemResponseBody) SetSuccess(v bool) *UpdateSubscriptionItemResponseBody {
	s.Success = &v
	return s
}

type UpdateSubscriptionItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSubscriptionItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubscriptionItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubscriptionItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionItemResponse) SetHeaders(v map[string]*string) *UpdateSubscriptionItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubscriptionItemResponse) SetStatusCode(v int32) *UpdateSubscriptionItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubscriptionItemResponse) SetBody(v *UpdateSubscriptionItemResponseBody) *UpdateSubscriptionItemResponse {
	s.Body = v
	return s
}

type UpdateWebhookRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Locale      *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	ServerUrl   *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	WebhookId   *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s UpdateWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebhookRequest) SetClientToken(v string) *UpdateWebhookRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWebhookRequest) SetLocale(v string) *UpdateWebhookRequest {
	s.Locale = &v
	return s
}

func (s *UpdateWebhookRequest) SetServerUrl(v string) *UpdateWebhookRequest {
	s.ServerUrl = &v
	return s
}

func (s *UpdateWebhookRequest) SetWebhookId(v int64) *UpdateWebhookRequest {
	s.WebhookId = &v
	return s
}

func (s *UpdateWebhookRequest) SetWebhookName(v string) *UpdateWebhookRequest {
	s.WebhookName = &v
	return s
}

type UpdateWebhookResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponseBody) SetCode(v string) *UpdateWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetMessage(v string) *UpdateWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetRequestId(v string) *UpdateWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetResult(v bool) *UpdateWebhookResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetSuccess(v bool) *UpdateWebhookResponseBody {
	s.Success = &v
	return s
}

type UpdateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponse) SetHeaders(v map[string]*string) *UpdateWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebhookResponse) SetStatusCode(v int32) *UpdateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebhookResponse) SetBody(v *UpdateWebhookResponseBody) *UpdateWebhookResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mscopensubscription"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateContactWithOptions(request *CreateContactRequest, runtime *util.RuntimeOptions) (_result *CreateContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		body["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Position)) {
		body["Position"] = request.Position
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateContact"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateContact(request *CreateContactRequest) (_result *CreateContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateContactResponse{}
	_body, _err := client.CreateContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscriptionItemWithOptions(request *CreateSubscriptionItemRequest, runtime *util.RuntimeOptions) (_result *CreateSubscriptionItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemName)) {
		body["ItemName"] = request.ItemName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubscriptionItem"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubscriptionItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscriptionItem(request *CreateSubscriptionItemRequest) (_result *CreateSubscriptionItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubscriptionItemResponse{}
	_body, _err := client.CreateSubscriptionItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWebhookWithOptions(request *CreateWebhookRequest, runtime *util.RuntimeOptions) (_result *CreateWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServerUrl)) {
		body["ServerUrl"] = request.ServerUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookName)) {
		body["WebhookName"] = request.WebhookName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWebhook"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWebhook(request *CreateWebhookRequest) (_result *CreateWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWebhookResponse{}
	_body, _err := client.CreateWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactWithOptions(request *DeleteContactRequest, runtime *util.RuntimeOptions) (_result *DeleteContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		body["ContactId"] = request.ContactId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContact"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContact(request *DeleteContactRequest) (_result *DeleteContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactResponse{}
	_body, _err := client.DeleteContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebhookWithOptions(request *DeleteWebhookRequest, runtime *util.RuntimeOptions) (_result *DeleteWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WebhookId)) {
		body["WebhookId"] = request.WebhookId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWebhook"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebhook(request *DeleteWebhookRequest) (_result *DeleteWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebhookResponse{}
	_body, _err := client.DeleteWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContactWithOptions(request *GetContactRequest, runtime *util.RuntimeOptions) (_result *GetContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetContact"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContact(request *GetContactRequest) (_result *GetContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetContactResponse{}
	_body, _err := client.GetContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubscriptionItemWithOptions(request *GetSubscriptionItemRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubscriptionItem"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubscriptionItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubscriptionItem(request *GetSubscriptionItemRequest) (_result *GetSubscriptionItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubscriptionItemResponse{}
	_body, _err := client.GetSubscriptionItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubscriptionItemDetailWithOptions(request *GetSubscriptionItemDetailRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionItemDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubscriptionItemDetail"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubscriptionItemDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubscriptionItemDetail(request *GetSubscriptionItemDetailRequest) (_result *GetSubscriptionItemDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubscriptionItemDetailResponse{}
	_body, _err := client.GetSubscriptionItemDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebhookWithOptions(request *GetWebhookRequest, runtime *util.RuntimeOptions) (_result *GetWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebhook"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebhook(request *GetWebhookRequest) (_result *GetWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebhookResponse{}
	_body, _err := client.GetWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContactsWithOptions(request *ListContactsRequest, runtime *util.RuntimeOptions) (_result *ListContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContacts"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContacts(request *ListContactsRequest) (_result *ListContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContactsResponse{}
	_body, _err := client.ListContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubscriptionItemGroupDetailsWithOptions(request *ListSubscriptionItemGroupDetailsRequest, runtime *util.RuntimeOptions) (_result *ListSubscriptionItemGroupDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptionItemGroupDetails"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionItemGroupDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubscriptionItemGroupDetails(request *ListSubscriptionItemGroupDetailsRequest) (_result *ListSubscriptionItemGroupDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubscriptionItemGroupDetailsResponse{}
	_body, _err := client.ListSubscriptionItemGroupDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubscriptionItemsWithOptions(request *ListSubscriptionItemsRequest, runtime *util.RuntimeOptions) (_result *ListSubscriptionItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptionItems"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubscriptionItems(request *ListSubscriptionItemsRequest) (_result *ListSubscriptionItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubscriptionItemsResponse{}
	_body, _err := client.ListSubscriptionItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWebhooksWithOptions(request *ListWebhooksRequest, runtime *util.RuntimeOptions) (_result *ListWebhooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWebhooks"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWebhooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWebhooks(request *ListWebhooksRequest) (_result *ListWebhooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWebhooksResponse{}
	_body, _err := client.ListWebhooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerificationMessageWithOptions(request *SendVerificationMessageRequest, runtime *util.RuntimeOptions) (_result *SendVerificationMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerificationMessage"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerificationMessage(request *SendVerificationMessageRequest) (_result *SendVerificationMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationMessageResponse{}
	_body, _err := client.SendVerificationMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContactWithOptions(request *UpdateContactRequest, runtime *util.RuntimeOptions) (_result *UpdateContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		body["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		body["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["Mobile"] = request.Mobile
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateContact"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContact(request *UpdateContactRequest) (_result *UpdateContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContactResponse{}
	_body, _err := client.UpdateContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubscriptionItemWithOptions(tmpReq *UpdateSubscriptionItemRequest, runtime *util.RuntimeOptions) (_result *UpdateSubscriptionItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSubscriptionItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ContactIds)) {
		request.ContactIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactIds, tea.String("ContactIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.WebhookIds)) {
		request.WebhookIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WebhookIds, tea.String("WebhookIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactIdsShrink)) {
		body["ContactIds"] = request.ContactIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EmailStatus)) {
		body["EmailStatus"] = request.EmailStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		body["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.PmsgStatus)) {
		body["PmsgStatus"] = request.PmsgStatus
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SmsStatus)) {
		body["SmsStatus"] = request.SmsStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TtsStatus)) {
		body["TtsStatus"] = request.TtsStatus
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookIdsShrink)) {
		body["WebhookIds"] = request.WebhookIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookStatus)) {
		body["WebhookStatus"] = request.WebhookStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubscriptionItem"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubscriptionItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubscriptionItem(request *UpdateSubscriptionItemRequest) (_result *UpdateSubscriptionItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubscriptionItemResponse{}
	_body, _err := client.UpdateSubscriptionItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWebhookWithOptions(request *UpdateWebhookRequest, runtime *util.RuntimeOptions) (_result *UpdateWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServerUrl)) {
		body["ServerUrl"] = request.ServerUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookId)) {
		body["WebhookId"] = request.WebhookId
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookName)) {
		body["WebhookName"] = request.WebhookName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWebhook"),
		Version:     tea.String("2021-07-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWebhook(request *UpdateWebhookRequest) (_result *UpdateWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.UpdateWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
