// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatappBindWabaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryChatappBindWabaResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryChatappBindWabaResponseBody
	GetCode() *string
	SetData(v *QueryChatappBindWabaResponseBodyData) *QueryChatappBindWabaResponseBody
	GetData() *QueryChatappBindWabaResponseBodyData
	SetMessage(v string) *QueryChatappBindWabaResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryChatappBindWabaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryChatappBindWabaResponseBody
	GetSuccess() *bool
}

type QueryChatappBindWabaResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryChatappBindWabaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryChatappBindWabaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappBindWabaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryChatappBindWabaResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryChatappBindWabaResponseBody) GetData() *QueryChatappBindWabaResponseBodyData {
	return s.Data
}

func (s *QueryChatappBindWabaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryChatappBindWabaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryChatappBindWabaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryChatappBindWabaResponseBody) SetAccessDeniedDetail(v string) *QueryChatappBindWabaResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetCode(v string) *QueryChatappBindWabaResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetData(v *QueryChatappBindWabaResponseBodyData) *QueryChatappBindWabaResponseBody {
	s.Data = v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetMessage(v string) *QueryChatappBindWabaResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetRequestId(v string) *QueryChatappBindWabaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) SetSuccess(v bool) *QueryChatappBindWabaResponseBody {
	s.Success = &v
	return s
}

func (s *QueryChatappBindWabaResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryChatappBindWabaResponseBodyData struct {
	// The review state of the WhatsApp Business account (WABA).
	//
	// >  Valid values:
	//
	// 	- PENDING: The WABA is to be reviewed.
	//
	// 	- APPROVED: The WABA was approved.
	//
	// 	- REJECTED: The WABA was rejected.
	//
	// 	- DISABLED: The WABA was forbidden.
	//
	// example:
	//
	// APPROVED
	AccountReviewStatus *string `json:"AccountReviewStatus,omitempty" xml:"AccountReviewStatus,omitempty"`
	// WABA related information.
	AuthInternationalRateEligibility map[string]interface{} `json:"AuthInternationalRateEligibility,omitempty" xml:"AuthInternationalRateEligibility,omitempty"`
	// The business ID.
	//
	// example:
	//
	// 19293988***
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The business name.
	//
	// example:
	//
	// Alibaba
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The currency.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the WhatsApp Business account.
	//
	// example:
	//
	// 20393988393993***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Marketing Messaging Lite status.
	//
	// example:
	//
	// Y
	MarketingMessageLiteStatus *string `json:"MarketingMessageLiteStatus,omitempty" xml:"MarketingMessageLiteStatus,omitempty"`
	// The namespace of the message template.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	MessageTemplateNamespace *string `json:"MessageTemplateNamespace,omitempty" xml:"MessageTemplateNamespace,omitempty"`
	// The name of the WhatsApp Business account.
	//
	// example:
	//
	// Alibaba
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The start time when the authentication-international rate applies.
	//
	// example:
	//
	// "start_time":1721952000
	PrimaryBusinessLocation *string `json:"PrimaryBusinessLocation,omitempty" xml:"PrimaryBusinessLocation,omitempty"`
}

func (s QueryChatappBindWabaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappBindWabaResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponseBodyData) GetAccountReviewStatus() *string {
	return s.AccountReviewStatus
}

func (s *QueryChatappBindWabaResponseBodyData) GetAuthInternationalRateEligibility() map[string]interface{} {
	return s.AuthInternationalRateEligibility
}

func (s *QueryChatappBindWabaResponseBodyData) GetBusinessId() *string {
	return s.BusinessId
}

func (s *QueryChatappBindWabaResponseBodyData) GetBusinessName() *string {
	return s.BusinessName
}

func (s *QueryChatappBindWabaResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *QueryChatappBindWabaResponseBodyData) GetId() *string {
	return s.Id
}

func (s *QueryChatappBindWabaResponseBodyData) GetMarketingMessageLiteStatus() *string {
	return s.MarketingMessageLiteStatus
}

func (s *QueryChatappBindWabaResponseBodyData) GetMessageTemplateNamespace() *string {
	return s.MessageTemplateNamespace
}

func (s *QueryChatappBindWabaResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryChatappBindWabaResponseBodyData) GetPrimaryBusinessLocation() *string {
	return s.PrimaryBusinessLocation
}

func (s *QueryChatappBindWabaResponseBodyData) SetAccountReviewStatus(v string) *QueryChatappBindWabaResponseBodyData {
	s.AccountReviewStatus = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetAuthInternationalRateEligibility(v map[string]interface{}) *QueryChatappBindWabaResponseBodyData {
	s.AuthInternationalRateEligibility = v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetBusinessId(v string) *QueryChatappBindWabaResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetBusinessName(v string) *QueryChatappBindWabaResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetCurrency(v string) *QueryChatappBindWabaResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetId(v string) *QueryChatappBindWabaResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetMarketingMessageLiteStatus(v string) *QueryChatappBindWabaResponseBodyData {
	s.MarketingMessageLiteStatus = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetMessageTemplateNamespace(v string) *QueryChatappBindWabaResponseBodyData {
	s.MessageTemplateNamespace = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetName(v string) *QueryChatappBindWabaResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) SetPrimaryBusinessLocation(v string) *QueryChatappBindWabaResponseBodyData {
	s.PrimaryBusinessLocation = &v
	return s
}

func (s *QueryChatappBindWabaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
