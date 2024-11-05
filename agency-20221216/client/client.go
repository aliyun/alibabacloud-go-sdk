// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelSubscriptionBillRequest struct {
	// The type of the bill to which you want to cancel the subscription. Valid values: PartnerBillingItemDetailForBillingPeriod, PartnerBillingItemDetailMonthly, PartnerInstanceDetailForBillingPeriod, and PartnerInstanceDetailMonthly.
	//
	// This parameter is required.
	//
	// example:
	//
	// PartnerBillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s CancelSubscriptionBillRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelSubscriptionBillRequest) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillRequest) SetSubscribeType(v string) *CancelSubscriptionBillRequest {
	s.SubscribeType = &v
	return s
}

type CancelSubscriptionBillResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelSubscriptionBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelSubscriptionBillResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillResponseBody) SetCode(v string) *CancelSubscriptionBillResponseBody {
	s.Code = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetData(v bool) *CancelSubscriptionBillResponseBody {
	s.Data = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetMessage(v string) *CancelSubscriptionBillResponseBody {
	s.Message = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetRequestId(v string) *CancelSubscriptionBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetSuccess(v bool) *CancelSubscriptionBillResponseBody {
	s.Success = &v
	return s
}

type CancelSubscriptionBillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSubscriptionBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSubscriptionBillResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelSubscriptionBillResponse) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillResponse) SetHeaders(v map[string]*string) *CancelSubscriptionBillResponse {
	s.Headers = v
	return s
}

func (s *CancelSubscriptionBillResponse) SetStatusCode(v int32) *CancelSubscriptionBillResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSubscriptionBillResponse) SetBody(v *CancelSubscriptionBillResponseBody) *CancelSubscriptionBillResponse {
	s.Body = v
	return s
}

type CreateCouponTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// All Products
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Partner
	CostBearer        *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	CouponDescription *string `json:"CouponDescription,omitempty" xml:"CouponDescription,omitempty"`
	// example:
	//
	// 2024-08-26
	Expireddate *string `json:"Expireddate,omitempty" xml:"Expireddate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Unlimited
	LimitPerPerson *string   `json:"LimitPerPerson,omitempty" xml:"LimitPerPerson,omitempty"`
	ProductType    []*string `json:"ProductType,omitempty" xml:"ProductType,omitempty" type:"Repeated"`
	// example:
	//
	// ALL
	PurchaseType *string `json:"PurchaseType,omitempty" xml:"PurchaseType,omitempty"`
	// This parameter is required.
	ReasonForApplication *string `json:"ReasonForApplication,omitempty" xml:"ReasonForApplication,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2024-08-26
	Vailddate *string `json:"Vailddate,omitempty" xml:"Vailddate,omitempty"`
	// example:
	//
	// 1
	Vaildperioddays *string `json:"Vaildperioddays,omitempty" xml:"Vaildperioddays,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Validity Duration
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCouponTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateRequest) SetAcceptLanguage(v string) *CreateCouponTemplateRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetApplicableProducts(v string) *CreateCouponTemplateRequest {
	s.ApplicableProducts = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCostBearer(v string) *CreateCouponTemplateRequest {
	s.CostBearer = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponDescription(v string) *CreateCouponTemplateRequest {
	s.CouponDescription = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetExpireddate(v string) *CreateCouponTemplateRequest {
	s.Expireddate = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetLimitPerPerson(v string) *CreateCouponTemplateRequest {
	s.LimitPerPerson = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetProductType(v []*string) *CreateCouponTemplateRequest {
	s.ProductType = v
	return s
}

func (s *CreateCouponTemplateRequest) SetPurchaseType(v string) *CreateCouponTemplateRequest {
	s.PurchaseType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetReasonForApplication(v string) *CreateCouponTemplateRequest {
	s.ReasonForApplication = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetTemplateName(v string) *CreateCouponTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetVailddate(v string) *CreateCouponTemplateRequest {
	s.Vailddate = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetVaildperioddays(v string) *CreateCouponTemplateRequest {
	s.Vaildperioddays = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetValidUntil(v string) *CreateCouponTemplateRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetValue(v string) *CreateCouponTemplateRequest {
	s.Value = &v
	return s
}

type CreateCouponTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// All Products
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Partner
	CostBearer        *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	CouponDescription *string `json:"CouponDescription,omitempty" xml:"CouponDescription,omitempty"`
	// example:
	//
	// 2024-08-26
	Expireddate *string `json:"Expireddate,omitempty" xml:"Expireddate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Unlimited
	LimitPerPerson    *string `json:"LimitPerPerson,omitempty" xml:"LimitPerPerson,omitempty"`
	ProductTypeShrink *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ALL
	PurchaseType *string `json:"PurchaseType,omitempty" xml:"PurchaseType,omitempty"`
	// This parameter is required.
	ReasonForApplication *string `json:"ReasonForApplication,omitempty" xml:"ReasonForApplication,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2024-08-26
	Vailddate *string `json:"Vailddate,omitempty" xml:"Vailddate,omitempty"`
	// example:
	//
	// 1
	Vaildperioddays *string `json:"Vaildperioddays,omitempty" xml:"Vaildperioddays,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Validity Duration
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCouponTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateShrinkRequest) SetAcceptLanguage(v string) *CreateCouponTemplateShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetApplicableProducts(v string) *CreateCouponTemplateShrinkRequest {
	s.ApplicableProducts = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetCostBearer(v string) *CreateCouponTemplateShrinkRequest {
	s.CostBearer = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetCouponDescription(v string) *CreateCouponTemplateShrinkRequest {
	s.CouponDescription = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetExpireddate(v string) *CreateCouponTemplateShrinkRequest {
	s.Expireddate = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetLimitPerPerson(v string) *CreateCouponTemplateShrinkRequest {
	s.LimitPerPerson = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetProductTypeShrink(v string) *CreateCouponTemplateShrinkRequest {
	s.ProductTypeShrink = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetPurchaseType(v string) *CreateCouponTemplateShrinkRequest {
	s.PurchaseType = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetReasonForApplication(v string) *CreateCouponTemplateShrinkRequest {
	s.ReasonForApplication = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetTemplateName(v string) *CreateCouponTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetVailddate(v string) *CreateCouponTemplateShrinkRequest {
	s.Vailddate = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetVaildperioddays(v string) *CreateCouponTemplateShrinkRequest {
	s.Vaildperioddays = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetValidUntil(v string) *CreateCouponTemplateShrinkRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreateCouponTemplateShrinkRequest) SetValue(v string) *CreateCouponTemplateShrinkRequest {
	s.Value = &v
	return s
}

type CreateCouponTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateCouponTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2103a30617045934095083027d88c5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCouponTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateResponseBody) SetCode(v string) *CreateCouponTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetData(v *CreateCouponTemplateResponseBodyData) *CreateCouponTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetMessage(v string) *CreateCouponTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetRequestId(v string) *CreateCouponTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetSuccess(v bool) *CreateCouponTemplateResponseBody {
	s.Success = &v
	return s
}

type CreateCouponTemplateResponseBodyData struct {
	// example:
	//
	// Custom
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// example:
	//
	// Partner
	CostBearer *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	// example:
	//
	// 111111
	CouponTemplateID *int64 `json:"CouponTemplateID,omitempty" xml:"CouponTemplateID,omitempty"`
	// example:
	//
	// 2024-04-02 16:15:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-01-01
	Expireddate *string   `json:"Expireddate,omitempty" xml:"Expireddate,omitempty"`
	ProductType []*string `json:"ProductType,omitempty" xml:"ProductType,omitempty" type:"Repeated"`
	// example:
	//
	// APPROVED
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2024-01-01
	Vailddate *string `json:"Vailddate,omitempty" xml:"Vailddate,omitempty"`
	// example:
	//
	// 1
	Vaildperioddays *string `json:"Vaildperioddays,omitempty" xml:"Vaildperioddays,omitempty"`
	// example:
	//
	// Validity Duration
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCouponTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateResponseBodyData) SetApplicableProducts(v string) *CreateCouponTemplateResponseBodyData {
	s.ApplicableProducts = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetCostBearer(v string) *CreateCouponTemplateResponseBodyData {
	s.CostBearer = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetCouponTemplateID(v int64) *CreateCouponTemplateResponseBodyData {
	s.CouponTemplateID = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetCreateTime(v string) *CreateCouponTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetExpireddate(v string) *CreateCouponTemplateResponseBodyData {
	s.Expireddate = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetProductType(v []*string) *CreateCouponTemplateResponseBodyData {
	s.ProductType = v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetStatus(v string) *CreateCouponTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetTemplateName(v string) *CreateCouponTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetVailddate(v string) *CreateCouponTemplateResponseBodyData {
	s.Vailddate = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetVaildperioddays(v string) *CreateCouponTemplateResponseBodyData {
	s.Vaildperioddays = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetValidUntil(v string) *CreateCouponTemplateResponseBodyData {
	s.ValidUntil = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetValue(v string) *CreateCouponTemplateResponseBodyData {
	s.Value = &v
	return s
}

type CreateCouponTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCouponTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCouponTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateResponse) SetHeaders(v map[string]*string) *CreateCouponTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCouponTemplateResponse) SetStatusCode(v int32) *CreateCouponTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCouponTemplateResponse) SetBody(v *CreateCouponTemplateResponseBody) *CreateCouponTemplateResponse {
	s.Body = v
	return s
}

type CreateCustomerRequest struct {
	// Customer\\"s name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DoorBell Marketing
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// The source/channel that allow client to connected with us. Please enumerate with Customer Source.
	//
	// This parameter is required.
	//
	// example:
	//
	// website
	CustomerSource *string `json:"CustomerSource,omitempty" xml:"CustomerSource,omitempty"`
	// The sub-industry that Customer\\"s business belongs to. Please enumerate with Customer Trade.
	//
	// example:
	//
	// 0101
	CustomerSubTrade *string `json:"CustomerSubTrade,omitempty" xml:"CustomerSubTrade,omitempty"`
	// The industry that Customer\\"s business belongs to. Please enumerate with Customer Trade.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	CustomerTrade *string `json:"CustomerTrade,omitempty" xml:"CustomerTrade,omitempty"`
	// The region that Customer choose to launch the Cloud Service. Please use ListCountries to confirm the valid region list for current UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AR
	Nation *string `json:"Nation,omitempty" xml:"Nation,omitempty"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetCustomerName(v string) *CreateCustomerRequest {
	s.CustomerName = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerSource(v string) *CreateCustomerRequest {
	s.CustomerSource = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerSubTrade(v string) *CreateCustomerRequest {
	s.CustomerSubTrade = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerTrade(v string) *CreateCustomerRequest {
	s.CustomerTrade = &v
	return s
}

func (s *CreateCustomerRequest) SetNation(v string) *CreateCustomerRequest {
	s.Nation = &v
	return s
}

type CreateCustomerResponseBody struct {
	// Code indicating whether the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data indicating whether a customer was successfully created. If it\\"s "true", the Message contains CID.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Massage indicating whether the call was successful.
	//
	// example:
	//
	// 12345
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// A9B725C7-3DBD-576B-AC91-F6F22AB99A77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call it self was successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetCode(v string) *CreateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerResponseBody) SetData(v bool) *CreateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCustomerResponseBody) SetMessage(v string) *CreateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomerResponseBody) SetRequestId(v string) *CreateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetSuccess(v bool) *CreateCustomerResponseBody {
	s.Success = &v
	return s
}

type CreateCustomerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetStatusCode(v int32) *CreateCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

type CustomerQuotaRecordListRequest struct {
	// End Date Format: yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-09-24
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Customer UID
	//
	// This parameter is required.
	//
	// example:
	//
	// 5113766248601929
	EndUserPk *int64 `json:"EndUserPk,omitempty" xml:"EndUserPk,omitempty"`
	// Multilingual Parameters, the default language is English.</br>
	//
	// en: English</br>
	//
	// zh: Chinese</br>
	//
	// ja: Japanese </br>
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Operation Type Enum</br>
	//
	// all All types</br>
	//
	// quota_create Create quota</br>
	//
	// quota_amount_adjust Adjust the amount of quota</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// Pagination, current page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Pagination, record number on each page. Maximum 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start Date Format: yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-02
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CustomerQuotaRecordListRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomerQuotaRecordListRequest) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListRequest) SetEndDate(v string) *CustomerQuotaRecordListRequest {
	s.EndDate = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetEndUserPk(v int64) *CustomerQuotaRecordListRequest {
	s.EndUserPk = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetLanguage(v string) *CustomerQuotaRecordListRequest {
	s.Language = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetOperationType(v string) *CustomerQuotaRecordListRequest {
	s.OperationType = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetPageNo(v int32) *CustomerQuotaRecordListRequest {
	s.PageNo = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetPageSize(v int32) *CustomerQuotaRecordListRequest {
	s.PageSize = &v
	return s
}

func (s *CustomerQuotaRecordListRequest) SetStartDate(v string) *CustomerQuotaRecordListRequest {
	s.StartDate = &v
	return s
}

type CustomerQuotaRecordListResponseBody struct {
	// Status code of returning result, 200 means success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Listed data of returning result
	Data []*CustomerQuotaRecordListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of returning data
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Record number on each page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of request
	//
	// example:
	//
	// 2103a0ae16849855284594613d874e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total volume
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CustomerQuotaRecordListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomerQuotaRecordListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListResponseBody) SetCode(v string) *CustomerQuotaRecordListResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetData(v []*CustomerQuotaRecordListResponseBodyData) *CustomerQuotaRecordListResponseBody {
	s.Data = v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetMsg(v string) *CustomerQuotaRecordListResponseBody {
	s.Msg = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetPageNo(v int32) *CustomerQuotaRecordListResponseBody {
	s.PageNo = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetPageSize(v int32) *CustomerQuotaRecordListResponseBody {
	s.PageSize = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetRequestId(v string) *CustomerQuotaRecordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetTotal(v int32) *CustomerQuotaRecordListResponseBody {
	s.Total = &v
	return s
}

type CustomerQuotaRecordListResponseBodyData struct {
	// The way to submit the quota adjustment operation. API/ACPN
	//
	// example:
	//
	// ACPN
	OperationSubmitType *string `json:"OperationSubmitType,omitempty" xml:"OperationSubmitType,omitempty"`
	// The time of submit the quota adjustment operation.
	//
	// example:
	//
	// 2023-12-15 10:34:36 UTC+8
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// Operation Type Enum</br>
	//
	// all All types</br>
	//
	// quota_create Create quota</br>
	//
	// quota_amount_adjust Adjust the amount of quota</br>
	//
	// example:
	//
	// quota_amount_adjust
	OperationTypeCode *string `json:"OperationTypeCode,omitempty" xml:"OperationTypeCode,omitempty"`
	// The description of submitted quota adjustment operation.
	//
	// example:
	//
	// Quota Adjustment
	OperationTypeDesc *string `json:"OperationTypeDesc,omitempty" xml:"OperationTypeDesc,omitempty"`
	// The UID of operator(Partner\\"s UID).
	//
	// example:
	//
	// 5113766248601929
	OperationUid *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	// Updated quota amount
	//
	// example:
	//
	// 121.00
	UpdateAfterAmount *string `json:"UpdateAfterAmount,omitempty" xml:"UpdateAfterAmount,omitempty"`
	// The difference amount between updated quota and original quota.
	//
	// example:
	//
	// -100.00
	UpdateAmount *string `json:"UpdateAmount,omitempty" xml:"UpdateAmount,omitempty"`
	// Original quota amount
	//
	// example:
	//
	// 221.00
	UpdateBeforeAmount *string `json:"UpdateBeforeAmount,omitempty" xml:"UpdateBeforeAmount,omitempty"`
}

func (s CustomerQuotaRecordListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CustomerQuotaRecordListResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationSubmitType(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationSubmitType = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationTime(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationTypeCode(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationTypeCode = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationTypeDesc(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationTypeDesc = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationUid(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationUid = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetUpdateAfterAmount(v string) *CustomerQuotaRecordListResponseBodyData {
	s.UpdateAfterAmount = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetUpdateAmount(v string) *CustomerQuotaRecordListResponseBodyData {
	s.UpdateAmount = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetUpdateBeforeAmount(v string) *CustomerQuotaRecordListResponseBodyData {
	s.UpdateBeforeAmount = &v
	return s
}

type CustomerQuotaRecordListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerQuotaRecordListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerQuotaRecordListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomerQuotaRecordListResponse) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListResponse) SetHeaders(v map[string]*string) *CustomerQuotaRecordListResponse {
	s.Headers = v
	return s
}

func (s *CustomerQuotaRecordListResponse) SetStatusCode(v int32) *CustomerQuotaRecordListResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerQuotaRecordListResponse) SetBody(v *CustomerQuotaRecordListResponseBody) *CustomerQuotaRecordListResponse {
	s.Body = v
	return s
}

type DeductOutstandingBalanceRequest struct {
	// The Deducted Credit to be offset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	DeductAmount *string `json:"DeductAmount,omitempty" xml:"DeductAmount,omitempty"`
	// Account UID of Distribution Customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeductOutstandingBalanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeductOutstandingBalanceRequest) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceRequest) SetDeductAmount(v string) *DeductOutstandingBalanceRequest {
	s.DeductAmount = &v
	return s
}

func (s *DeductOutstandingBalanceRequest) SetUid(v int64) *DeductOutstandingBalanceRequest {
	s.Uid = &v
	return s
}

type DeductOutstandingBalanceResponseBody struct {
	// Result Code. Value Range:
	//
	// - 200 OK
	//
	// - 1109 System Error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Same as Code Parameter Value.
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeductOutstandingBalanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeductOutstandingBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceResponseBody) SetCode(v string) *DeductOutstandingBalanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetMessage(v string) *DeductOutstandingBalanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetRequestId(v string) *DeductOutstandingBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetSuccess(v bool) *DeductOutstandingBalanceResponseBody {
	s.Success = &v
	return s
}

type DeductOutstandingBalanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeductOutstandingBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeductOutstandingBalanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeductOutstandingBalanceResponse) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceResponse) SetHeaders(v map[string]*string) *DeductOutstandingBalanceResponse {
	s.Headers = v
	return s
}

func (s *DeductOutstandingBalanceResponse) SetStatusCode(v int32) *DeductOutstandingBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeductOutstandingBalanceResponse) SetBody(v *DeductOutstandingBalanceResponseBody) *DeductOutstandingBalanceResponse {
	s.Body = v
	return s
}

type EditEndUserStatusRequest struct {
	// Shutdown Status</br>
	//
	// - postPayFreeze, the account have been blocked</br>
	//
	// - postPayThaw, the account have been unlocked</br>
	//
	// example:
	//
	// postPayFreeze
	CreditStatus *string `json:"CreditStatus,omitempty" xml:"CreditStatus,omitempty"`
	// UID
	//
	// example:
	//
	// 1792155717328010
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditEndUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s EditEndUserStatusRequest) GoString() string {
	return s.String()
}

func (s *EditEndUserStatusRequest) SetCreditStatus(v string) *EditEndUserStatusRequest {
	s.CreditStatus = &v
	return s
}

func (s *EditEndUserStatusRequest) SetUid(v int64) *EditEndUserStatusRequest {
	s.Uid = &v
	return s
}

type EditEndUserStatusResponseBody struct {
	// Status Code</br>
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Success or not</br>
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Message</br>
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Message</br>
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Request ID</br>
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditEndUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditEndUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *EditEndUserStatusResponseBody) SetCode(v string) *EditEndUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetData(v string) *EditEndUserStatusResponseBody {
	s.Data = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetMessage(v string) *EditEndUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetMsg(v string) *EditEndUserStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetRequestId(v string) *EditEndUserStatusResponseBody {
	s.RequestId = &v
	return s
}

type EditEndUserStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditEndUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditEndUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s EditEndUserStatusResponse) GoString() string {
	return s.String()
}

func (s *EditEndUserStatusResponse) SetHeaders(v map[string]*string) *EditEndUserStatusResponse {
	s.Headers = v
	return s
}

func (s *EditEndUserStatusResponse) SetStatusCode(v int32) *EditEndUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *EditEndUserStatusResponse) SetBody(v *EditEndUserStatusResponseBody) *EditEndUserStatusResponse {
	s.Body = v
	return s
}

type EditNewBuyStatusRequest struct {
	// New Purchase Status</br>
	//
	// - cancelBan: Cancel the restriction for New Purchase request</br>
	//
	// - ban: ban the New Purchase request</br>
	//
	// example:
	//
	// cancelBan
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	// Customer UID
	//
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditNewBuyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s EditNewBuyStatusRequest) GoString() string {
	return s.String()
}

func (s *EditNewBuyStatusRequest) SetNewBuyStatus(v string) *EditNewBuyStatusRequest {
	s.NewBuyStatus = &v
	return s
}

func (s *EditNewBuyStatusRequest) SetUid(v int64) *EditNewBuyStatusRequest {
	s.Uid = &v
	return s
}

type EditNewBuyStatusResponseBody struct {
	// Status Code</br>
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Success or not</br>
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Message</br>
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Message</br>
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Request ID</br>
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditNewBuyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditNewBuyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *EditNewBuyStatusResponseBody) SetCode(v string) *EditNewBuyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetData(v string) *EditNewBuyStatusResponseBody {
	s.Data = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetMessage(v string) *EditNewBuyStatusResponseBody {
	s.Message = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetMsg(v string) *EditNewBuyStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetRequestId(v string) *EditNewBuyStatusResponseBody {
	s.RequestId = &v
	return s
}

type EditNewBuyStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditNewBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditNewBuyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s EditNewBuyStatusResponse) GoString() string {
	return s.String()
}

func (s *EditNewBuyStatusResponse) SetHeaders(v map[string]*string) *EditNewBuyStatusResponse {
	s.Headers = v
	return s
}

func (s *EditNewBuyStatusResponse) SetStatusCode(v int32) *EditNewBuyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *EditNewBuyStatusResponse) SetBody(v *EditNewBuyStatusResponseBody) *EditNewBuyStatusResponse {
	s.Body = v
	return s
}

type EditZeroCreditShutdownRequest struct {
	// UID
	//
	// example:
	//
	// Shutdown Policy</br>
	//
	// - immediatelyStop, The instances of the specified End User\\"s account will be shutdown immediately once EU triggered the Shutdown Policy.</br>
	//
	// - delayStop, The instances of the specified End User\\"s account will be shutdown later, even EU have triggered the Shutdown Policy.</br>
	//
	// - noStop, The instances of the specified End User\\"s account will not be shutdown, after EU have triggered the Shutdown Policy.</br>
	ShutdownPolicy *string `json:"ShutdownPolicy,omitempty" xml:"ShutdownPolicy,omitempty"`
	// No Change History
	//
	// example:
	//
	// 1263644979775567
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditZeroCreditShutdownRequest) String() string {
	return tea.Prettify(s)
}

func (s EditZeroCreditShutdownRequest) GoString() string {
	return s.String()
}

func (s *EditZeroCreditShutdownRequest) SetShutdownPolicy(v string) *EditZeroCreditShutdownRequest {
	s.ShutdownPolicy = &v
	return s
}

func (s *EditZeroCreditShutdownRequest) SetUid(v int64) *EditZeroCreditShutdownRequest {
	s.Uid = &v
	return s
}

type EditZeroCreditShutdownResponseBody struct {
	// Success or not</br>
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID</br>
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Message</br>
	//
	// example:
	//
	// Message</br>
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// NO_STOP
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// success
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditZeroCreditShutdownResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditZeroCreditShutdownResponseBody) GoString() string {
	return s.String()
}

func (s *EditZeroCreditShutdownResponseBody) SetCode(v string) *EditZeroCreditShutdownResponseBody {
	s.Code = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetData(v string) *EditZeroCreditShutdownResponseBody {
	s.Data = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetMessage(v string) *EditZeroCreditShutdownResponseBody {
	s.Message = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetMsg(v string) *EditZeroCreditShutdownResponseBody {
	s.Msg = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetRequestId(v string) *EditZeroCreditShutdownResponseBody {
	s.RequestId = &v
	return s
}

type EditZeroCreditShutdownResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditZeroCreditShutdownResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditZeroCreditShutdownResponse) String() string {
	return tea.Prettify(s)
}

func (s EditZeroCreditShutdownResponse) GoString() string {
	return s.String()
}

func (s *EditZeroCreditShutdownResponse) SetHeaders(v map[string]*string) *EditZeroCreditShutdownResponse {
	s.Headers = v
	return s
}

func (s *EditZeroCreditShutdownResponse) SetStatusCode(v int32) *EditZeroCreditShutdownResponse {
	s.StatusCode = &v
	return s
}

func (s *EditZeroCreditShutdownResponse) SetBody(v *EditZeroCreditShutdownResponseBody) *EditZeroCreditShutdownResponse {
	s.Body = v
	return s
}

type ExportCustomerQuotaRecordRequest struct {
	// End Date Format:  yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-24
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Customer UID
	//
	// This parameter is required.
	//
	// example:
	//
	// 5113766248601929
	EndUserPk *int64 `json:"EndUserPk,omitempty" xml:"EndUserPk,omitempty"`
	// Multilingual Parameters, the default language is English.</br>
	//
	// en: English</br>
	//
	// zh: Chinese</br>
	//
	// ja: Japanese </br>
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Operation Type Enum</br>
	//
	// all All types</br>
	//
	// quota_create Create quota</br>
	//
	// quota_amount_adjust Adjust the amount of quota</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// Start Date Format:  yyyy-MM-dd
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-11-10
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportCustomerQuotaRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportCustomerQuotaRecordRequest) GoString() string {
	return s.String()
}

func (s *ExportCustomerQuotaRecordRequest) SetEndDate(v string) *ExportCustomerQuotaRecordRequest {
	s.EndDate = &v
	return s
}

func (s *ExportCustomerQuotaRecordRequest) SetEndUserPk(v int64) *ExportCustomerQuotaRecordRequest {
	s.EndUserPk = &v
	return s
}

func (s *ExportCustomerQuotaRecordRequest) SetLanguage(v string) *ExportCustomerQuotaRecordRequest {
	s.Language = &v
	return s
}

func (s *ExportCustomerQuotaRecordRequest) SetOperationType(v string) *ExportCustomerQuotaRecordRequest {
	s.OperationType = &v
	return s
}

func (s *ExportCustomerQuotaRecordRequest) SetStartDate(v string) *ExportCustomerQuotaRecordRequest {
	s.StartDate = &v
	return s
}

type ExportCustomerQuotaRecordResponseBody struct {
	// Code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	Data *ExportCustomerQuotaRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID of the Request
	//
	// example:
	//
	// 210bc4b416874189683843905d9f9a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportCustomerQuotaRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportCustomerQuotaRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ExportCustomerQuotaRecordResponseBody) SetCode(v string) *ExportCustomerQuotaRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ExportCustomerQuotaRecordResponseBody) SetData(v *ExportCustomerQuotaRecordResponseBodyData) *ExportCustomerQuotaRecordResponseBody {
	s.Data = v
	return s
}

func (s *ExportCustomerQuotaRecordResponseBody) SetMsg(v string) *ExportCustomerQuotaRecordResponseBody {
	s.Msg = &v
	return s
}

func (s *ExportCustomerQuotaRecordResponseBody) SetRequestId(v string) *ExportCustomerQuotaRecordResponseBody {
	s.RequestId = &v
	return s
}

type ExportCustomerQuotaRecordResponseBodyData struct {
	// Estimated duration, in minutes.
	//
	// example:
	//
	// 1
	Cost *int32 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// ID of Export task
	//
	// example:
	//
	// 1231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportCustomerQuotaRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExportCustomerQuotaRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExportCustomerQuotaRecordResponseBodyData) SetCost(v int32) *ExportCustomerQuotaRecordResponseBodyData {
	s.Cost = &v
	return s
}

func (s *ExportCustomerQuotaRecordResponseBodyData) SetId(v int64) *ExportCustomerQuotaRecordResponseBodyData {
	s.Id = &v
	return s
}

type ExportCustomerQuotaRecordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportCustomerQuotaRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCustomerQuotaRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportCustomerQuotaRecordResponse) GoString() string {
	return s.String()
}

func (s *ExportCustomerQuotaRecordResponse) SetHeaders(v map[string]*string) *ExportCustomerQuotaRecordResponse {
	s.Headers = v
	return s
}

func (s *ExportCustomerQuotaRecordResponse) SetStatusCode(v int32) *ExportCustomerQuotaRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportCustomerQuotaRecordResponse) SetBody(v *ExportCustomerQuotaRecordResponseBody) *ExportCustomerQuotaRecordResponse {
	s.Body = v
	return s
}

type GetAccountInfoRequest struct {
	// Message
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Success
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 10 (Value <= 20)
	//
	// example:
	//
	// 1215848086704806
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// Result Code - Error Code. Value Range:
	//
	// - 200 OK
	//
	// - 1109 System Error
	//
	// - 3029: Invalid UID
	//
	// - 3062: UID and UserType are both empty.
	//
	// - 3063: UserType value out of range.
	//
	// example:
	//
	// 1
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAccountInfoRequest) SetCurrentPage(v int32) *GetAccountInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAccountInfoRequest) SetPageSize(v int32) *GetAccountInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountInfoRequest) SetUid(v int64) *GetAccountInfoRequest {
	s.Uid = &v
	return s
}

func (s *GetAccountInfoRequest) SetUserType(v string) *GetAccountInfoRequest {
	s.UserType = &v
	return s
}

type GetAccountInfoResponseBody struct {
	// Account Type:
	//
	// - 1 Agency\\"s End User
	//
	// - 2 Reseller\\"s End User
	//
	// - 3 Enterprise
	//
	// - 4 T2 Agency Partner
	//
	// - 5 T2 Reseller Partner
	//
	// - 6 T2 Agency+Reseller Partner
	AccountInfoList *GetAccountInfoResponseBodyAccountInfoList `json:"AccountInfoList,omitempty" xml:"AccountInfoList,omitempty" type:"Struct"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	//
	// example:
	//
	// Pagination Information
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination, page volume in total.
	PageInfo *GetAccountInfoResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request id, a unique identifier generated by Alibaba cloud for the request.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Pagination, record number on each page.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBody) SetAccountInfoList(v *GetAccountInfoResponseBodyAccountInfoList) *GetAccountInfoResponseBody {
	s.AccountInfoList = v
	return s
}

func (s *GetAccountInfoResponseBody) SetCode(v string) *GetAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetMessage(v string) *GetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetPageInfo(v *GetAccountInfoResponseBodyPageInfo) *GetAccountInfoResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetAccountInfoResponseBody) SetRequestId(v string) *GetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetSuccess(v bool) *GetAccountInfoResponseBody {
	s.Success = &v
	return s
}

type GetAccountInfoResponseBodyAccountInfoList struct {
	AccountInfo []*GetAccountInfoResponseBodyAccountInfoListAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Repeated"`
}

func (s GetAccountInfoResponseBodyAccountInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfoList) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfoList) SetAccountInfo(v []*GetAccountInfoResponseBodyAccountInfoListAccountInfo) *GetAccountInfoResponseBodyAccountInfoList {
	s.AccountInfo = v
	return s
}

type GetAccountInfoResponseBodyAccountInfoListAccountInfo struct {
	// The E-mail of Distribution Customer.
	//
	// example:
	//
	// Description of Distribution Customer.
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// Account CID of Distribution Customer.
	//
	// example:
	//
	// 1234@qq.com
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// XXX Technology LTD.
	//
	// example:
	//
	// 2021-01-01
	AssociationSuccessTime *string `json:"AssociationSuccessTime,omitempty" xml:"AssociationSuccessTime,omitempty"`
	// customer\\"s CID
	//
	// example:
	//
	// 61479572
	Cid                 *int64 `json:"Cid,omitempty" xml:"Cid,omitempty"`
	CustomerAccountType *int32 `json:"CustomerAccountType,omitempty" xml:"CustomerAccountType,omitempty"`
	// customer manager
	//
	// example:
	//
	// abc
	CustomerBd                  *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	CustomerEnterpriseCertified *int32  `json:"CustomerEnterpriseCertified,omitempty" xml:"CustomerEnterpriseCertified,omitempty"`
	// The account have Shutdown-delay Privilege, After Shutdown-delay Credit is ran out, Alibaba Cloud will take over resources and keep the instance for 15 days. In addition, the instance will be released if Sub Account failed to pay the bill within these 15 days.
	//
	// example:
	//
	// 600
	DelayAmount *string `json:"DelayAmount,omitempty" xml:"DelayAmount,omitempty"`
	// Partner\\"s Shutdown Policy Management for Sub Account.
	//
	// 1: delayStop. The account have Shutdown-delay Privilege, After Shutdown-delay Credit is ran out, Alibaba Cloud will take over resources and keep the instance for 15 days. In addition, the instance will be released if Sub Account failed to pay the bill within these 15 days.
	//
	// 2: noStop. Partner will manually manage Shutdown Status for Sub Account. Meanwhile, System would not manage the resource\\"s life-circle of Sub Account.
	//
	// 3: immediatelyStop. Once valid quota of Sub Account falls below 0 and be identified as defaulting account, it will trigger the instance shutdown immediately.
	//
	// example:
	//
	// noStop
	DelayStatus *string `json:"DelayStatus,omitempty" xml:"DelayStatus,omitempty"`
	// Sub Account
	//
	// example:
	//
	// 1234@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Account UID of Distribution Customer.
	//
	// example:
	//
	// 13641588680
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Purchase Forbidden：Ban the new purchase action
	//
	// normal：Normal--End Use can issue Cloud Resource order immediately.
	//
	// example:
	//
	// Normal
	NewBuyStatus        *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	RegisterCountryCode *string `json:"RegisterCountryCode,omitempty" xml:"RegisterCountryCode,omitempty"`
	// Valid mobile number of Distribution Customer.
	//
	// example:
	//
	// Alibaba Cloud Login name of Distribution Customer.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of Sub Account:
	//
	// 1.	Use the official name of Company, if Sub Account is an enterprise.
	//
	// 2.	Use the official name of Partner, if Sub Account is a T2 reseller.
	//
	// example:
	//
	// 1
	SubAccountType *int32 `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 1415740779475837
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetAccountInfoResponseBodyAccountInfoListAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfoListAccountInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetAccountNickname(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.AccountNickname = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetAliyunId(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.AliyunId = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetAssociationSuccessTime(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.AssociationSuccessTime = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetCid(v int64) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Cid = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetCustomerAccountType(v int32) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.CustomerAccountType = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetCustomerBd(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.CustomerBd = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetCustomerEnterpriseCertified(v int32) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.CustomerEnterpriseCertified = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetDelayAmount(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.DelayAmount = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetDelayStatus(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.DelayStatus = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetEmail(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Email = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetMobile(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Mobile = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetNewBuyStatus(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.NewBuyStatus = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetRegisterCountryCode(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.RegisterCountryCode = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetRemark(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Remark = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetSubAccountType(v int32) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.SubAccountType = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetUid(v int64) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Uid = &v
	return s
}

type GetAccountInfoResponseBodyPageInfo struct {
	// Account Information
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Pagination, current page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of Account Information
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAccountInfoResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyPageInfo) SetPage(v int32) *GetAccountInfoResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *GetAccountInfoResponseBodyPageInfo) SetPageSize(v int32) *GetAccountInfoResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetAccountInfoResponseBodyPageInfo) SetTotal(v int32) *GetAccountInfoResponseBodyPageInfo {
	s.Total = &v
	return s
}

type GetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponse) SetHeaders(v map[string]*string) *GetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountInfoResponse) SetStatusCode(v int32) *GetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountInfoResponse) SetBody(v *GetAccountInfoResponseBody) *GetAccountInfoResponse {
	s.Body = v
	return s
}

type GetCoupondeductProductCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s GetCoupondeductProductCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCoupondeductProductCodeRequest) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeRequest) SetAcceptLanguage(v string) *GetCoupondeductProductCodeRequest {
	s.AcceptLanguage = &v
	return s
}

type GetCoupondeductProductCodeResponseBody struct {
	// example:
	//
	// code
	Code    *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetCoupondeductProductCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCoupondeductProductCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCoupondeductProductCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeResponseBody) SetCode(v string) *GetCoupondeductProductCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetData(v []*GetCoupondeductProductCodeResponseBodyData) *GetCoupondeductProductCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetMessage(v string) *GetCoupondeductProductCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetRequestId(v string) *GetCoupondeductProductCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetSuccess(v bool) *GetCoupondeductProductCodeResponseBody {
	s.Success = &v
	return s
}

type GetCoupondeductProductCodeResponseBodyData struct {
	// example:
	//
	// code1
	ProductType interface{} `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetCoupondeductProductCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCoupondeductProductCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeResponseBodyData) SetProductType(v interface{}) *GetCoupondeductProductCodeResponseBodyData {
	s.ProductType = v
	return s
}

type GetCoupondeductProductCodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCoupondeductProductCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCoupondeductProductCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCoupondeductProductCodeResponse) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeResponse) SetHeaders(v map[string]*string) *GetCoupondeductProductCodeResponse {
	s.Headers = v
	return s
}

func (s *GetCoupondeductProductCodeResponse) SetStatusCode(v int32) *GetCoupondeductProductCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCoupondeductProductCodeResponse) SetBody(v *GetCoupondeductProductCodeResponseBody) *GetCoupondeductProductCodeResponse {
	s.Body = v
	return s
}

type GetCreditInfoRequest struct {
	// Sub Account UID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792155717328010
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetCreditInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCreditInfoRequest) SetUid(v int64) *GetCreditInfoRequest {
	s.Uid = &v
	return s
}

type GetCreditInfoResponseBody struct {
	// Result Code:
	//
	// - 200 OK
	//
	// - 1109 System Error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCreditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Message Information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCreditInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponseBody) SetCode(v string) *GetCreditInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetData(v *GetCreditInfoResponseBodyData) *GetCreditInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCreditInfoResponseBody) SetMessage(v string) *GetCreditInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetRequestId(v string) *GetCreditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetSuccess(v bool) *GetCreditInfoResponseBody {
	s.Success = &v
	return s
}

type GetCreditInfoResponseBodyData struct {
	// The Credit Control status, Value Range:</br>
	//
	// 1. normal - Sub Account status is running as usual.
	//
	// 2. arrearsNotShutdown - Sub Account status is running as usual, but have outstanding bill(s).
	//
	// 3. shutdown -  Sub Account status is down.
	//
	// example:
	//
	// normal
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// Percentage value, when the available credit limit is lower than this credit limit percentage, a notification E-mail will be sent to the main account.
	//
	// example:
	//
	// 20
	AlarmThreshold *string `json:"AlarmThreshold,omitempty" xml:"AlarmThreshold,omitempty"`
	// The Credit available to consume.
	//
	// example:
	//
	// 800
	AvailableCredit *string `json:"AvailableCredit,omitempty" xml:"AvailableCredit,omitempty"`
	// Obtain total unpaid amount on demo bill before simulated deduction.
	//
	// example:
	//
	// 0.000000
	ConsumedUndeductedValue *string `json:"ConsumedUndeductedValue,omitempty" xml:"ConsumedUndeductedValue,omitempty"`
	// The Credit Line of Sub Account
	//
	// example:
	//
	// 1000
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The Credit have been consumed by Sub Account, and haven\\"t be paid.
	//
	// example:
	//
	// 200
	OutstandingBalance *string `json:"OutstandingBalance,omitempty" xml:"OutstandingBalance,omitempty"`
	// The systematic controlling policy for resource management, specifically when the available Credit of Sub Account falls to 0 or less.</br>
	//
	// - 1: delayStop. The account have Shutdown-delay Privilege,  After Shutdown-delay Credit is ran out, Alibaba Cloud will take over resources and keep the instance for 15 days. In addition, the instance will be released if Sub Account failed to pay the bill within these 15 days.</br>
	//
	// - 2: noStop. Partner will manually manage Shutdown Status for Sub Account. Meanwhile, System would not manage the resource\\"s life-circle of Sub Account.</br>
	//
	// - 3: immediatelyStop. Once valid quota of Sub Account falls below 0 and be identified as defaulting account, it will trigger the instance shutdown immediately.</br>
	//
	// example:
	//
	// delayStop
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
	// Manage order operation.
	//
	// - ban：Ban the new purchase action.
	//
	// - normal：The account could raise new purchase order as usual.
	//
	// example:
	//
	// ban
	NewBuyStatus *string `json:"newBuyStatus,omitempty" xml:"newBuyStatus,omitempty"`
}

func (s GetCreditInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponseBodyData) SetAccountStatus(v string) *GetCreditInfoResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetAlarmThreshold(v string) *GetCreditInfoResponseBodyData {
	s.AlarmThreshold = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetAvailableCredit(v string) *GetCreditInfoResponseBodyData {
	s.AvailableCredit = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetConsumedUndeductedValue(v string) *GetCreditInfoResponseBodyData {
	s.ConsumedUndeductedValue = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetCreditLine(v string) *GetCreditInfoResponseBodyData {
	s.CreditLine = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetOutstandingBalance(v string) *GetCreditInfoResponseBodyData {
	s.OutstandingBalance = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetZeroCreditShutdownPolicy(v string) *GetCreditInfoResponseBodyData {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetNewBuyStatus(v string) *GetCreditInfoResponseBodyData {
	s.NewBuyStatus = &v
	return s
}

type GetCreditInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreditInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponse) SetHeaders(v map[string]*string) *GetCreditInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCreditInfoResponse) SetStatusCode(v int32) *GetCreditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreditInfoResponse) SetBody(v *GetCreditInfoResponseBody) *GetCreditInfoResponse {
	s.Body = v
	return s
}

type GetCustomerOrdersRequest struct {
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// example:
	//
	// myBd
	CustomerManager *string `json:"CustomerManager,omitempty" xml:"CustomerManager,omitempty"`
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-23 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 209335720330622
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 0
	OrderSource *int32 `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// RENEW
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// vm_intl
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-13 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
}

func (s GetCustomerOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrdersRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersRequest) SetCustomerAccount(v string) *GetCustomerOrdersRequest {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetCustomerManager(v string) *GetCustomerOrdersRequest {
	s.CustomerManager = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetCustomerUid(v int64) *GetCustomerOrdersRequest {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetEndDate(v string) *GetCustomerOrdersRequest {
	s.EndDate = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderId(v int64) *GetCustomerOrdersRequest {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderSource(v int32) *GetCustomerOrdersRequest {
	s.OrderSource = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderStatus(v int32) *GetCustomerOrdersRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderType(v string) *GetCustomerOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetPageNo(v int32) *GetCustomerOrdersRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetPageSize(v int32) *GetCustomerOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetProductType(v string) *GetCustomerOrdersRequest {
	s.ProductType = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetStartDate(v string) *GetCustomerOrdersRequest {
	s.StartDate = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetTimeType(v int32) *GetCustomerOrdersRequest {
	s.TimeType = &v
	return s
}

type GetCustomerOrdersResponseBody struct {
	// example:
	//
	// 200
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetCustomerOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg     *string                              `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 23309219-4A34-589D-A3E0-9B2A3BFFD24F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetCustomerOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersResponseBody) SetCode(v string) *GetCustomerOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetData(v []*GetCustomerOrdersResponseBodyData) *GetCustomerOrdersResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetMessage(v string) *GetCustomerOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetMsg(v string) *GetCustomerOrdersResponseBody {
	s.Msg = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetPageNo(v int32) *GetCustomerOrdersResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetPageSize(v int32) *GetCustomerOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetRequestId(v string) *GetCustomerOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetSuccess(v bool) *GetCustomerOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetTotal(v int32) *GetCustomerOrdersResponseBody {
	s.Total = &v
	return s
}

type GetCustomerOrdersResponseBodyData struct {
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// example:
	//
	// myBd
	CustomerManager *string `json:"CustomerManager,omitempty" xml:"CustomerManager,omitempty"`
	// example:
	//
	// 123456
	CustomerNo *int64 `json:"CustomerNo,omitempty" xml:"CustomerNo,omitempty"`
	// example:
	//
	// 236414227150922
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 0
	OrderSource *string `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 3.92
	OriginalCost *float64 `json:"OriginalCost,omitempty" xml:"OriginalCost,omitempty"`
	// example:
	//
	// 3:32
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// example:
	//
	// 2024-08-13 13:02:02
	PaymentTime *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	// example:
	//
	// 3.92
	PretaxCost *float64 `json:"PretaxCost,omitempty" xml:"PretaxCost,omitempty"`
	// example:
	//
	// oss
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// example:
	//
	// snapshot
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 2024-08-13 13:02:02
	TimeToOrder *string `json:"TimeToOrder,omitempty" xml:"TimeToOrder,omitempty"`
}

func (s GetCustomerOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersResponseBodyData) SetCustomerAccount(v string) *GetCustomerOrdersResponseBodyData {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetCustomerManager(v string) *GetCustomerOrdersResponseBodyData {
	s.CustomerManager = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetCustomerNo(v int64) *GetCustomerOrdersResponseBodyData {
	s.CustomerNo = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderId(v int64) *GetCustomerOrdersResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderSource(v string) *GetCustomerOrdersResponseBodyData {
	s.OrderSource = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderStatus(v int32) *GetCustomerOrdersResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderType(v string) *GetCustomerOrdersResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOriginalCost(v float64) *GetCustomerOrdersResponseBodyData {
	s.OriginalCost = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetPaymentMethod(v string) *GetCustomerOrdersResponseBodyData {
	s.PaymentMethod = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetPaymentTime(v string) *GetCustomerOrdersResponseBodyData {
	s.PaymentTime = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetPretaxCost(v float64) *GetCustomerOrdersResponseBodyData {
	s.PretaxCost = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetProductDetail(v string) *GetCustomerOrdersResponseBodyData {
	s.ProductDetail = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetProductType(v string) *GetCustomerOrdersResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetTimeToOrder(v string) *GetCustomerOrdersResponseBodyData {
	s.TimeToOrder = &v
	return s
}

type GetCustomerOrdersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerOrdersResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersResponse) SetHeaders(v map[string]*string) *GetCustomerOrdersResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerOrdersResponse) SetStatusCode(v int32) *GetCustomerOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerOrdersResponse) SetBody(v *GetCustomerOrdersResponseBody) *GetCustomerOrdersResponse {
	s.Body = v
	return s
}

type GetDailyBillRequest struct {
	// Bill Owner type. Value Range:</br>
	//
	// 1: Master account</br>
	//
	// 2: Sub account</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// BillType. Value Range:</br>
	//
	// - DailyOrder(Deprecated)
	//
	// - DailyBill (Deprecated)
	//
	// - DailyInstanceBill (Deprecated)
	//
	// - DailyInstanceBillV2
	//
	// This parameter is required.
	//
	// example:
	//
	// DailyInstanceBillV2
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Billing date. Format YYYY-MM-DD
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-24
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetDailyBillRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillRequest) GoString() string {
	return s.String()
}

func (s *GetDailyBillRequest) SetBillOwner(v string) *GetDailyBillRequest {
	s.BillOwner = &v
	return s
}

func (s *GetDailyBillRequest) SetBillType(v string) *GetDailyBillRequest {
	s.BillType = &v
	return s
}

func (s *GetDailyBillRequest) SetDate(v string) *GetDailyBillRequest {
	s.Date = &v
	return s
}

type GetDailyBillResponseBody struct {
	// Result Code:
	//
	// 	- 200 OK
	//
	// 	- 1109 System error
	//
	// 	- 3050 Bill Type can only be DailyOrder, DailyBill or DailyInstanceBill.
	//
	// 	- 3049 Incorrect format of Spending Time.
	//
	// 	- 3048 Bill Owner can only be 1 or 2.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetDailyBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Same as Code parameters.
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDailyBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillResponseBody) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponseBody) SetCode(v string) *GetDailyBillResponseBody {
	s.Code = &v
	return s
}

func (s *GetDailyBillResponseBody) SetData(v *GetDailyBillResponseBodyData) *GetDailyBillResponseBody {
	s.Data = v
	return s
}

func (s *GetDailyBillResponseBody) SetMessage(v string) *GetDailyBillResponseBody {
	s.Message = &v
	return s
}

func (s *GetDailyBillResponseBody) SetRequestId(v string) *GetDailyBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDailyBillResponseBody) SetSuccess(v bool) *GetDailyBillResponseBody {
	s.Success = &v
	return s
}

type GetDailyBillResponseBodyData struct {
	// The link to download CSV file, please use HTTP Protocol.
	//
	// example:
	//
	// intl-reseller-month-bill.oss-ap-southeast-1.aliyuncs.com/statements/month/dts/1064252248461886/202104/Month%20Billing%20Invoice%20202104.pdf?Expires=1671160973&OSSAccessKeyId=TMP.3KhvoD9pW264cPv8sYe8E2zJ5HLWmrijNHgKiKpXcy8yS472BcrvemgTfNrrXKu5fCCdbLr2XhmYAyYPmbCe8zJyEkmYzL&Signature=hns1PgiiYl1WdI%2FdiOqbEdsgmfI%3D
	BillLinkCSV *string `json:"BillLinkCSV,omitempty" xml:"BillLinkCSV,omitempty"`
	// The link to download XLSX file, please use HTTP Protocol.
	//
	// example:
	//
	// intl-reseller-month-bill.oss-ap-southeast-1.aliyuncs.com/statements/month/dts/1064252248461886/202104/Month%20Billing%20Invoice%20202104.pdf?Expires=1671160973&OSSAccessKeyId=TMP.3KhvoD9pW264cPv8sYe8E2zJ5HLWmrijNHgKiKpXcy8yS472BcrvemgTfNrrXKu5fCCdbLr2XhmYAyYPmbCe8zJyEkmYzL&Signature=hns1PgiiYl1WdI%2FdiOqbEdsgmfI%3D
	BillLinkXLSX *string `json:"BillLinkXLSX,omitempty" xml:"BillLinkXLSX,omitempty"`
	// Same as inserted parameter BillOwner.
	//
	// example:
	//
	// 1
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// Same as inserted parameter BillType.
	//
	// example:
	//
	// DailyInstanceBillV2
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Spending Time, refer to the exact time of costuming.
	//
	// example:
	//
	// 20221201
	SpendingTime *string `json:"SpendingTime,omitempty" xml:"SpendingTime,omitempty"`
}

func (s GetDailyBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponseBodyData) SetBillLinkCSV(v string) *GetDailyBillResponseBodyData {
	s.BillLinkCSV = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetBillLinkXLSX(v string) *GetDailyBillResponseBodyData {
	s.BillLinkXLSX = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetBillOwner(v string) *GetDailyBillResponseBodyData {
	s.BillOwner = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetBillType(v string) *GetDailyBillResponseBodyData {
	s.BillType = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetSpendingTime(v string) *GetDailyBillResponseBodyData {
	s.SpendingTime = &v
	return s
}

type GetDailyBillResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDailyBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDailyBillResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillResponse) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponse) SetHeaders(v map[string]*string) *GetDailyBillResponse {
	s.Headers = v
	return s
}

func (s *GetDailyBillResponse) SetStatusCode(v int32) *GetDailyBillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDailyBillResponse) SetBody(v *GetDailyBillResponseBody) *GetDailyBillResponse {
	s.Body = v
	return s
}

type GetInviteStatusRequest struct {
	// inviteId list</br>
	//
	// `Sub-levels <= 5`
	//
	// This parameter is required.
	InviteStatusList []*GetInviteStatusRequestInviteStatusList `json:"InviteStatusList,omitempty" xml:"InviteStatusList,omitempty" type:"Repeated"`
}

func (s GetInviteStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusRequest) GoString() string {
	return s.String()
}

func (s *GetInviteStatusRequest) SetInviteStatusList(v []*GetInviteStatusRequestInviteStatusList) *GetInviteStatusRequest {
	s.InviteStatusList = v
	return s
}

type GetInviteStatusRequestInviteStatusList struct {
	// Invitation ID, From interface InviteSubAccount
	//
	// example:
	//
	// 123
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
}

func (s GetInviteStatusRequestInviteStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusRequestInviteStatusList) GoString() string {
	return s.String()
}

func (s *GetInviteStatusRequestInviteStatusList) SetInviteId(v int64) *GetInviteStatusRequestInviteStatusList {
	s.InviteId = &v
	return s
}

type GetInviteStatusResponseBody struct {
	// Status Code. Error Code:
	//
	// - 3057 InviteId is empty
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetInviteStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInviteStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBody) SetCode(v string) *GetInviteStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetData(v *GetInviteStatusResponseBodyData) *GetInviteStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetInviteStatusResponseBody) SetMessage(v string) *GetInviteStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetRequestId(v string) *GetInviteStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetSuccess(v bool) *GetInviteStatusResponseBody {
	s.Success = &v
	return s
}

type GetInviteStatusResponseBodyData struct {
	InviteStatus []*GetInviteStatusResponseBodyDataInviteStatus `json:"InviteStatus,omitempty" xml:"InviteStatus,omitempty" type:"Repeated"`
}

func (s GetInviteStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyData) SetInviteStatus(v []*GetInviteStatusResponseBodyDataInviteStatus) *GetInviteStatusResponseBodyData {
	s.InviteStatus = v
	return s
}

type GetInviteStatusResponseBodyDataInviteStatus struct {
	// Result Code. Value Range:
	//
	// 	- 200 OK
	//
	// 	- 1109 system error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of Invitation Status result
	InviteStatusList *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList `json:"InviteStatusList,omitempty" xml:"InviteStatusList,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInviteStatusResponseBodyDataInviteStatus) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBodyDataInviteStatus) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetCode(v string) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Code = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetInviteStatusList(v *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) *GetInviteStatusResponseBodyDataInviteStatus {
	s.InviteStatusList = v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetMessage(v string) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Message = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetSuccess(v bool) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Success = &v
	return s
}

type GetInviteStatusResponseBodyDataInviteStatusInviteStatusList struct {
	// The time that Distribution Customer successfully associated with Distributor.</br>
	//
	// This value will be empty if there is no existing association.
	//
	// example:
	//
	// 2018-02-13
	AssociationSuccessTime *string `json:"AssociationSuccessTime,omitempty" xml:"AssociationSuccessTime,omitempty"`
	// Distribution Customer\\"s CID
	//
	// example:
	//
	// 1234567890123
	Cid *int64 `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// The time of email been sent out.
	//
	// example:
	//
	// 2018-02-12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// 1093238769140523
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// Invitation Status:
	//
	// 	- 0 No visit on registration URL
	//
	// 	- 1 Successful Registration
	//
	// 	- 2 Unsuccessful Registration
	//
	// 	- 3 Registration URL have been visited, but no submitted sheet/ticket.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Account Type:
	//
	// - 1 Agency\\"s End User
	//
	// - 2 Reseller\\"s End User
	//
	// - 5 T2 Reseller Partner
	//
	// example:
	//
	// 1
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// Distribution Customer\\"s UID
	//
	// example:
	//
	// 1234567890123
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetAssociationSuccessTime(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.AssociationSuccessTime = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetCid(v int64) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Cid = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetGmtCreate(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.GmtCreate = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetParentId(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.ParentId = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetStatus(v int32) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Status = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetSubAccountType(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.SubAccountType = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetUid(v int64) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Uid = &v
	return s
}

type GetInviteStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInviteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInviteStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponse) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponse) SetHeaders(v map[string]*string) *GetInviteStatusResponse {
	s.Headers = v
	return s
}

func (s *GetInviteStatusResponse) SetStatusCode(v int32) *GetInviteStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInviteStatusResponse) SetBody(v *GetInviteStatusResponseBody) *GetInviteStatusResponse {
	s.Body = v
	return s
}

type GetMonthlyBillRequest struct {
	// Bill Owner type. Value Range:</br>
	//
	// 1: Master account</br>
	//
	// 2: Sub account</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// Value Range:
	//
	// - MonthlyInvoice
	//
	// - MonthRefundInvoice
	//
	// - MonthlySummary
	//
	// - MonthlyInstanceAddAdjustBill
	//
	// - MonthlyInstanceRefundBill
	//
	// - MonthlyAddAdjustInvoce
	//
	// - MonthlyRefundAdjustInvoce
	//
	// - MonthlyInstanceConsumeV2
	//
	// - MarginReportV2
	//
	// This parameter is required.
	//
	// example:
	//
	// MonthlyInvoice
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Billing Month, Format is YYYY-MM
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s GetMonthlyBillRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillRequest) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillRequest) SetBillOwner(v string) *GetMonthlyBillRequest {
	s.BillOwner = &v
	return s
}

func (s *GetMonthlyBillRequest) SetBillType(v string) *GetMonthlyBillRequest {
	s.BillType = &v
	return s
}

func (s *GetMonthlyBillRequest) SetMonth(v string) *GetMonthlyBillRequest {
	s.Month = &v
	return s
}

type GetMonthlyBillResponseBody struct {
	// Result Code:
	//
	// 	- 200 OK
	//
	// 	- 1109 System error
	//
	// 	- 3030 Sub Account Nickname exceeds maximum length, maximum length 150 bytes.
	//
	// 	- 3031 Remark exceeds maximum length, maximum length 3000 bytes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetMonthlyBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Same as Code parameters.
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMonthlyBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponseBody) SetCode(v string) *GetMonthlyBillResponseBody {
	s.Code = &v
	return s
}

func (s *GetMonthlyBillResponseBody) SetData(v *GetMonthlyBillResponseBodyData) *GetMonthlyBillResponseBody {
	s.Data = v
	return s
}

func (s *GetMonthlyBillResponseBody) SetMessage(v string) *GetMonthlyBillResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonthlyBillResponseBody) SetRequestId(v string) *GetMonthlyBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonthlyBillResponseBody) SetSuccess(v bool) *GetMonthlyBillResponseBody {
	s.Success = &v
	return s
}

type GetMonthlyBillResponseBodyData struct {
	// The link to download CSV file, please use HTTP Protocol.
	//
	// example:
	//
	// intl-reseller-month-bill.oss-ap-southeast-1.aliyuncs.com/statements/month/dts/1064252248461886/202104/Month%20Billing%20Invoice%20202104.pdf?Expires=1671160973&OSSAccessKeyId=TMP.3KhvoD9pW264cPv8sYe8E2zJ5HLWmrijNHgKiKpXcy8yS472BcrvemgTfNrrXKu5fCCdbLr2XhmYAyYPmbCe8zJyEkmYzL&Signature=hns1PgiiYl1WdI%2FdiOqbEdsgmfI%3D
	BillLinkCSV *string `json:"BillLinkCSV,omitempty" xml:"BillLinkCSV,omitempty"`
	// The link to download XLSX file, please use HTTP Protocol.
	//
	// example:
	//
	// intl-reseller-month-bill.oss-ap-southeast-1.aliyuncs.com/statements/month/dts/1064252248461886/202104/Month%20Billing%20Invoice%20202104.pdf?Expires=1671160973&OSSAccessKeyId=TMP.3KhvoD9pW264cPv8sYe8E2zJ5HLWmrijNHgKiKpXcy8yS472BcrvemgTfNrrXKu5fCCdbLr2XhmYAyYPmbCe8zJyEkmYzL&Signature=hns1PgiiYl1WdI%2FdiOqbEdsgmfI%3D
	BillLinkXLSX *string `json:"BillLinkXLSX,omitempty" xml:"BillLinkXLSX,omitempty"`
	// Same as inserted parameter BillOwner.
	//
	// example:
	//
	// 1
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// Same as inserted parameter BillType.
	//
	// example:
	//
	// MonthlyInvoice
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The URL to download invoice.
	//
	// example:
	//
	// intl-reseller-month-bill.oss-ap-southeast-1.aliyuncs.com/statements/month/dts/1064252248461886/202104/Month%20Billing%20Invoice%20202104.pdf?Expires=1671160973&OSSAccessKeyId=TMP.3KhvoD9pW264cPv8sYe8E2zJ5HLWmrijNHgKiKpXcy8yS472BcrvemgTfNrrXKu5fCCdbLr2XhmYAyYPmbCe8zJyEkmYzL&Signature=hns1PgiiYl1WdI%2FdiOqbEdsgmfI%3D
	InvoiceLink *string `json:"InvoiceLink,omitempty" xml:"InvoiceLink,omitempty"`
	// It states the existence of refund invoice. </br>
	//
	// Candidate Values: True/False
	//
	// example:
	//
	// True
	RefundInvoiceFlag *bool `json:"RefundInvoiceFlag,omitempty" xml:"RefundInvoiceFlag,omitempty"`
	// The URL to download refund invoice.
	//
	// example:
	//
	// intl-reseller-month-bill.oss-ap-southeast-1.aliyuncs.com/statements/month/dts/1064252248461886/202104/Month%20Billing%20Invoice%20202104.pdf?Expires=1671160973&OSSAccessKeyId=TMP.3KhvoD9pW264cPv8sYe8E2zJ5HLWmrijNHgKiKpXcy8yS472BcrvemgTfNrrXKu5fCCdbLr2XhmYAyYPmbCe8zJyEkmYzL&Signature=hns1PgiiYl1WdI%2FdiOqbEdsgmfI%3D
	RefundInvoiceLink *string `json:"RefundInvoiceLink,omitempty" xml:"RefundInvoiceLink,omitempty"`
	// Spending Time, refer to the exact time of costuming.
	//
	// example:
	//
	// 20221201
	SpendingTime *string `json:"SpendingTime,omitempty" xml:"SpendingTime,omitempty"`
}

func (s GetMonthlyBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponseBodyData) SetBillLinkCSV(v string) *GetMonthlyBillResponseBodyData {
	s.BillLinkCSV = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetBillLinkXLSX(v string) *GetMonthlyBillResponseBodyData {
	s.BillLinkXLSX = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetBillOwner(v string) *GetMonthlyBillResponseBodyData {
	s.BillOwner = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetBillType(v string) *GetMonthlyBillResponseBodyData {
	s.BillType = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetInvoiceLink(v string) *GetMonthlyBillResponseBodyData {
	s.InvoiceLink = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetRefundInvoiceFlag(v bool) *GetMonthlyBillResponseBodyData {
	s.RefundInvoiceFlag = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetRefundInvoiceLink(v string) *GetMonthlyBillResponseBodyData {
	s.RefundInvoiceLink = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetSpendingTime(v string) *GetMonthlyBillResponseBodyData {
	s.SpendingTime = &v
	return s
}

type GetMonthlyBillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonthlyBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonthlyBillResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillResponse) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponse) SetHeaders(v map[string]*string) *GetMonthlyBillResponse {
	s.Headers = v
	return s
}

func (s *GetMonthlyBillResponse) SetStatusCode(v int32) *GetMonthlyBillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonthlyBillResponse) SetBody(v *GetMonthlyBillResponseBody) *GetMonthlyBillResponse {
	s.Body = v
	return s
}

type GetUnassociatedCustomerRequest struct {
	// Pagination, current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Pagination, record number on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetUnassociatedCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerRequest) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerRequest) SetCurrentPage(v int32) *GetUnassociatedCustomerRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetUnassociatedCustomerRequest) SetPageSize(v int32) *GetUnassociatedCustomerRequest {
	s.PageSize = &v
	return s
}

type GetUnassociatedCustomerResponseBody struct {
	// Error Code, Candidate Value：
	//
	// 	- 200: OK
	//
	// 	- 1109: System error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of Invitation Information
	InviteInfoList *GetUnassociatedCustomerResponseBodyInviteInfoList `json:"InviteInfoList,omitempty" xml:"InviteInfoList,omitempty" type:"Struct"`
	// Message information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination Information
	PageInfo *GetUnassociatedCustomerResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// 23309219-4A34-589D-A3E0-9B2A3BFFD24F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUnassociatedCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBody) SetCode(v string) *GetUnassociatedCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetInviteInfoList(v *GetUnassociatedCustomerResponseBodyInviteInfoList) *GetUnassociatedCustomerResponseBody {
	s.InviteInfoList = v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetMessage(v string) *GetUnassociatedCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetPageInfo(v *GetUnassociatedCustomerResponseBodyPageInfo) *GetUnassociatedCustomerResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetRequestId(v string) *GetUnassociatedCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetSuccess(v bool) *GetUnassociatedCustomerResponseBody {
	s.Success = &v
	return s
}

type GetUnassociatedCustomerResponseBodyInviteInfoList struct {
	InviteInfo []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo `json:"InviteInfo,omitempty" xml:"InviteInfo,omitempty" type:"Repeated"`
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoList) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoList) SetInviteInfo(v []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) *GetUnassociatedCustomerResponseBodyInviteInfoList {
	s.InviteInfo = v
	return s
}

type GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo struct {
	// The name of Customer who are to be invited.
	//
	// example:
	//
	// My Client
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The Email of Customer who are to be invited.
	//
	// example:
	//
	// 12345@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The time of email been sent out.
	//
	// example:
	//
	// 2023-05-10
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Invitation ID
	//
	// example:
	//
	// 190
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	// Invitation Status:
	//
	// 	- 0 No visit on registration URL
	//
	// 	- 1 Successful Registration
	//
	// 	- 2 Unsuccessful Registration
	//
	// 	- 3 Registration URL have been visited, but no submitted sheet/ticket.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetAccountNickname(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.AccountNickname = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetEmail(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.Email = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetGmtCreate(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetInviteId(v int64) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.InviteId = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetStatus(v int32) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.Status = &v
	return s
}

type GetUnassociatedCustomerResponseBodyPageInfo struct {
	// Pagination, current page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Pagination, record number on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Pagination, page volume in total.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetUnassociatedCustomerResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetPage(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetPageSize(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetTotal(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.Total = &v
	return s
}

type GetUnassociatedCustomerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnassociatedCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnassociatedCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponse) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponse) SetHeaders(v map[string]*string) *GetUnassociatedCustomerResponse {
	s.Headers = v
	return s
}

func (s *GetUnassociatedCustomerResponse) SetStatusCode(v int32) *GetUnassociatedCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnassociatedCustomerResponse) SetBody(v *GetUnassociatedCustomerResponseBody) *GetUnassociatedCustomerResponse {
	s.Body = v
	return s
}

type InviteSubAccountRequest struct {
	// List of invited account information,  less than 5 accounts at a time.</br>
	//
	// `Sub-levels <= 5`
	//
	// This parameter is required.
	AccountInfoList []*InviteSubAccountRequestAccountInfoList `json:"AccountInfoList,omitempty" xml:"AccountInfoList,omitempty" type:"Repeated"`
}

func (s InviteSubAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountRequest) GoString() string {
	return s.String()
}

func (s *InviteSubAccountRequest) SetAccountInfoList(v []*InviteSubAccountRequestAccountInfoList) *InviteSubAccountRequest {
	s.AccountInfoList = v
	return s
}

type InviteSubAccountRequestAccountInfoList struct {
	// The name of Sub Account:</br>
	//
	// 1. Use the official name of Company, if Sub Account is an enterprise.</br>
	//
	// 2. Use the official name of Partner, if Sub Account is a T2 reseller.</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// XXX Technology LTD.
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The total budget Credit of Sub Account that distributed by Partner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// Customer ID, Returning ID from CreateCustomer API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The email address of End User,  which will receive the invitation email.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345@163.com
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	// Initial Order Status</br>
	//
	// 1. ban：Ban the new purchase action--After End User finish registration and authorization, they can\\"t issue Cloud Resource order immediately. Partner should manually update the "Order Control" settings as "Normal" to enable new order.</br>
	//
	// 2. normal：Normal--After End User finished registration and authorization, they can issue Cloud Resource order immediately.</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// ban
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	// Description of Sub Account.
	//
	// example:
	//
	// The invitation to develop XX as a Sub Account
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of Sub Account</br>
	//
	// 1 Agency\\"s End User</br>
	//
	// 2 Reseller\\"s End user</br>
	//
	// 5 Reseller\\"s T2 Partner</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// Partner\\"s Shutdown Policy Management for Sub Account.</br>
	//
	// 1: delayStop. The account have Shutdown-delay Privilege,  After Shutdown-delay Credit is ran out, Alibaba Cloud will take over resources and keep the instance for 15 days. In addition, the instance will be released if Sub Account failed to pay the bill within these 15 days.</br>
	//
	// 2: noStop. Partner will manually manage Shutdown Status for Sub Account. Meanwhile, System would not manage the resource\\"s life-circle of Sub Account.</br>
	//
	// 3: immediatelyStop. Once valid quota of Sub Account falls below 0 and be identified as defaulting account, it will trigger the instance shutdown immediately.</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
}

func (s InviteSubAccountRequestAccountInfoList) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountRequestAccountInfoList) GoString() string {
	return s.String()
}

func (s *InviteSubAccountRequestAccountInfoList) SetAccountNickname(v string) *InviteSubAccountRequestAccountInfoList {
	s.AccountNickname = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCreditLine(v string) *InviteSubAccountRequestAccountInfoList {
	s.CreditLine = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCustomerBd(v string) *InviteSubAccountRequestAccountInfoList {
	s.CustomerBd = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCustomerId(v string) *InviteSubAccountRequestAccountInfoList {
	s.CustomerId = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetEmailAddress(v string) *InviteSubAccountRequestAccountInfoList {
	s.EmailAddress = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetNewBuyStatus(v string) *InviteSubAccountRequestAccountInfoList {
	s.NewBuyStatus = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetRemark(v string) *InviteSubAccountRequestAccountInfoList {
	s.Remark = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetSubAccountType(v string) *InviteSubAccountRequestAccountInfoList {
	s.SubAccountType = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetZeroCreditShutdownPolicy(v string) *InviteSubAccountRequestAccountInfoList {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

type InviteSubAccountResponseBody struct {
	// Error Code: </br>
	//
	// • 200 OK</br>
	//
	// • 1109 System Error</br>
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message</br>
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of invitation sending results
	Results *InviteSubAccountResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// Candidate Values: True/False, this value states if the current API calling action is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBody) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBody) SetCode(v string) *InviteSubAccountResponseBody {
	s.Code = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetMessage(v string) *InviteSubAccountResponseBody {
	s.Message = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetRequestId(v string) *InviteSubAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetResults(v *InviteSubAccountResponseBodyResults) *InviteSubAccountResponseBody {
	s.Results = v
	return s
}

func (s *InviteSubAccountResponseBody) SetSuccess(v bool) *InviteSubAccountResponseBody {
	s.Success = &v
	return s
}

type InviteSubAccountResponseBodyResults struct {
	Result []*InviteSubAccountResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InviteSubAccountResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBodyResults) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResults) SetResult(v []*InviteSubAccountResponseBodyResultsResult) *InviteSubAccountResponseBodyResults {
	s.Result = v
	return s
}

type InviteSubAccountResponseBodyResultsResult struct {
	// Error Code, 200 OK
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message, Notes of Code
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Returning Message of Invitation Results
	Result *InviteSubAccountResponseBodyResultsResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Always true.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubAccountResponseBodyResultsResult) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResultsResult) SetCode(v string) *InviteSubAccountResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetMessage(v string) *InviteSubAccountResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetResult(v *InviteSubAccountResponseBodyResultsResultResult) *InviteSubAccountResponseBodyResultsResult {
	s.Result = v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetSuccess(v bool) *InviteSubAccountResponseBodyResultsResult {
	s.Success = &v
	return s
}

type InviteSubAccountResponseBodyResultsResultResult struct {
	// Valid days of registration URL, count on daily basis.
	//
	// example:
	//
	// 15
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Invitation ID, The invitation status tracking code.
	//
	// example:
	//
	// 12345
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	// URL for Partner Customer Registration.
	//
	// example:
	//
	// http://agency-intl.console.aliyun.com/customer/register?intl=true&fxinfo=-4uT%2FMWHnnUdvr5GXVd1AYK8luTnGgH3M7Y3lSCd5M1fxRwAkViTWtDJDpckh0HL
	RegUrl *string `json:"RegUrl,omitempty" xml:"RegUrl,omitempty"`
}

func (s InviteSubAccountResponseBodyResultsResultResult) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBodyResultsResultResult) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetDays(v int32) *InviteSubAccountResponseBodyResultsResultResult {
	s.Days = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetInviteId(v int64) *InviteSubAccountResponseBodyResultsResultResult {
	s.InviteId = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetRegUrl(v string) *InviteSubAccountResponseBodyResultsResultResult {
	s.RegUrl = &v
	return s
}

type InviteSubAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteSubAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteSubAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponse) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponse) SetHeaders(v map[string]*string) *InviteSubAccountResponse {
	s.Headers = v
	return s
}

func (s *InviteSubAccountResponse) SetStatusCode(v int32) *InviteSubAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteSubAccountResponse) SetBody(v *InviteSubAccountResponseBody) *InviteSubAccountResponse {
	s.Body = v
	return s
}

type IssueCouponForCustomerRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111,2222
	Uidlist *string `json:"Uidlist,omitempty" xml:"Uidlist,omitempty"`
}

func (s IssueCouponForCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s IssueCouponForCustomerRequest) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerRequest) SetAcceptLanguage(v string) *IssueCouponForCustomerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *IssueCouponForCustomerRequest) SetCouponTemplateId(v int64) *IssueCouponForCustomerRequest {
	s.CouponTemplateId = &v
	return s
}

func (s *IssueCouponForCustomerRequest) SetUidlist(v string) *IssueCouponForCustomerRequest {
	s.Uidlist = &v
	return s
}

type IssueCouponForCustomerResponseBody struct {
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data    *IssueCouponForCustomerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s IssueCouponForCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IssueCouponForCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerResponseBody) SetCode(v string) *IssueCouponForCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetMessage(v string) *IssueCouponForCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetRequestId(v string) *IssueCouponForCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetSuccess(v bool) *IssueCouponForCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetData(v *IssueCouponForCustomerResponseBodyData) *IssueCouponForCustomerResponseBody {
	s.Data = v
	return s
}

type IssueCouponForCustomerResponseBodyData struct {
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// example:
	//
	// 2024-03-05 18:24:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 111,2222
	Uidlist *string `json:"Uidlist,omitempty" xml:"Uidlist,omitempty"`
}

func (s IssueCouponForCustomerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IssueCouponForCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerResponseBodyData) SetCouponTemplateId(v int64) *IssueCouponForCustomerResponseBodyData {
	s.CouponTemplateId = &v
	return s
}

func (s *IssueCouponForCustomerResponseBodyData) SetCreateTime(v string) *IssueCouponForCustomerResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *IssueCouponForCustomerResponseBodyData) SetUidlist(v string) *IssueCouponForCustomerResponseBodyData {
	s.Uidlist = &v
	return s
}

type IssueCouponForCustomerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IssueCouponForCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IssueCouponForCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s IssueCouponForCustomerResponse) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerResponse) SetHeaders(v map[string]*string) *IssueCouponForCustomerResponse {
	s.Headers = v
	return s
}

func (s *IssueCouponForCustomerResponse) SetStatusCode(v int32) *IssueCouponForCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *IssueCouponForCustomerResponse) SetBody(v *IssueCouponForCustomerResponseBody) *IssueCouponForCustomerResponse {
	s.Body = v
	return s
}

type ListCountriesResponseBody struct {
	// Error Code
	//
	// 	- 200: OK
	//
	// 	- 1109: System error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of Region Code
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Message information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// A747A00F-E096-5244-88B3-3E474BAE3AE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCountriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCountriesResponseBody) SetCode(v string) *ListCountriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCountriesResponseBody) SetData(v []*string) *ListCountriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCountriesResponseBody) SetMessage(v string) *ListCountriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCountriesResponseBody) SetRequestId(v string) *ListCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCountriesResponseBody) SetSuccess(v bool) *ListCountriesResponseBody {
	s.Success = &v
	return s
}

type ListCountriesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCountriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCountriesResponse) GoString() string {
	return s.String()
}

func (s *ListCountriesResponse) SetHeaders(v map[string]*string) *ListCountriesResponse {
	s.Headers = v
	return s
}

func (s *ListCountriesResponse) SetStatusCode(v int32) *ListCountriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCountriesResponse) SetBody(v *ListCountriesResponseBody) *ListCountriesResponse {
	s.Body = v
	return s
}

type ListCouponUsageRequest struct {
	// example:
	//
	// oqevfbveuadcrduzmf@ttirv.net
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCouponUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCouponUsageRequest) GoString() string {
	return s.String()
}

func (s *ListCouponUsageRequest) SetAccount(v string) *ListCouponUsageRequest {
	s.Account = &v
	return s
}

func (s *ListCouponUsageRequest) SetCouponTemplateId(v int64) *ListCouponUsageRequest {
	s.CouponTemplateId = &v
	return s
}

func (s *ListCouponUsageRequest) SetPage(v int32) *ListCouponUsageRequest {
	s.Page = &v
	return s
}

func (s *ListCouponUsageRequest) SetPageSize(v int32) *ListCouponUsageRequest {
	s.PageSize = &v
	return s
}

func (s *ListCouponUsageRequest) SetStatus(v string) *ListCouponUsageRequest {
	s.Status = &v
	return s
}

func (s *ListCouponUsageRequest) SetUid(v int64) *ListCouponUsageRequest {
	s.Uid = &v
	return s
}

type ListCouponUsageResponseBody struct {
	// example:
	//
	// 200
	Code     *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data     []*ListCouponUsageResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message  *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo *ListCouponUsageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCouponUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCouponUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponseBody) SetCode(v string) *ListCouponUsageResponseBody {
	s.Code = &v
	return s
}

func (s *ListCouponUsageResponseBody) SetData(v []*ListCouponUsageResponseBodyData) *ListCouponUsageResponseBody {
	s.Data = v
	return s
}

func (s *ListCouponUsageResponseBody) SetMessage(v string) *ListCouponUsageResponseBody {
	s.Message = &v
	return s
}

func (s *ListCouponUsageResponseBody) SetPageInfo(v *ListCouponUsageResponseBodyPageInfo) *ListCouponUsageResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCouponUsageResponseBody) SetRequestId(v string) *ListCouponUsageResponseBody {
	s.RequestId = &v
	return s
}

type ListCouponUsageResponseBodyData struct {
	// example:
	//
	// oqevfbveuadcrduzmf@ttirv.net
	Account *string  `json:"Account,omitempty" xml:"Account,omitempty"`
	Amount  *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 0.01
	Balance *float64 `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// example:
	//
	// 59226280
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 503802
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// example:
	//
	// 2023-04-06 00:00:00 ~ 2023-04-07 00:00:00
	EffDate *string `json:"EffDate,omitempty" xml:"EffDate,omitempty"`
	// example:
	//
	// 2023-04-06 19:32:10
	PublishDate *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1647668856741998
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCouponUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCouponUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponseBodyData) SetAccount(v string) *ListCouponUsageResponseBodyData {
	s.Account = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetAmount(v float64) *ListCouponUsageResponseBodyData {
	s.Amount = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetBalance(v float64) *ListCouponUsageResponseBodyData {
	s.Balance = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetCouponId(v string) *ListCouponUsageResponseBodyData {
	s.CouponId = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetCouponTemplateId(v int64) *ListCouponUsageResponseBodyData {
	s.CouponTemplateId = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetEffDate(v string) *ListCouponUsageResponseBodyData {
	s.EffDate = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetPublishDate(v string) *ListCouponUsageResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetStatus(v string) *ListCouponUsageResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetUid(v int64) *ListCouponUsageResponseBodyData {
	s.Uid = &v
	return s
}

type ListCouponUsageResponseBodyPageInfo struct {
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 300
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCouponUsageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCouponUsageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponseBodyPageInfo) SetPage(v int32) *ListCouponUsageResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *ListCouponUsageResponseBodyPageInfo) SetPageSize(v int32) *ListCouponUsageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCouponUsageResponseBodyPageInfo) SetTotal(v int32) *ListCouponUsageResponseBodyPageInfo {
	s.Total = &v
	return s
}

type ListCouponUsageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCouponUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCouponUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCouponUsageResponse) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponse) SetHeaders(v map[string]*string) *ListCouponUsageResponse {
	s.Headers = v
	return s
}

func (s *ListCouponUsageResponse) SetStatusCode(v int32) *ListCouponUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCouponUsageResponse) SetBody(v *ListCouponUsageResponseBody) *ListCouponUsageResponse {
	s.Body = v
	return s
}

type QuotaListExportPagedRequest struct {
	// Pagination, current page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Multilingual Parameters, the default language is English.</br>
	//
	// en: English</br>
	//
	// zh: Chinese</br>
	//
	// ja: Japanese </br>
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Pagination, record number on each page, maximum 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QuotaListExportPagedRequest) String() string {
	return tea.Prettify(s)
}

func (s QuotaListExportPagedRequest) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedRequest) SetCurrentPage(v int32) *QuotaListExportPagedRequest {
	s.CurrentPage = &v
	return s
}

func (s *QuotaListExportPagedRequest) SetLanguage(v string) *QuotaListExportPagedRequest {
	s.Language = &v
	return s
}

func (s *QuotaListExportPagedRequest) SetPageSize(v int32) *QuotaListExportPagedRequest {
	s.PageSize = &v
	return s
}

type QuotaListExportPagedResponseBody struct {
	// Status code of returning result, 200 means success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Listed data of returning result
	Data []*QuotaListExportPagedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of returning result
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Record number on each page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of the Request
	//
	// example:
	//
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total volume
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuotaListExportPagedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuotaListExportPagedResponseBody) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedResponseBody) SetCode(v string) *QuotaListExportPagedResponseBody {
	s.Code = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetData(v []*QuotaListExportPagedResponseBodyData) *QuotaListExportPagedResponseBody {
	s.Data = v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetMsg(v string) *QuotaListExportPagedResponseBody {
	s.Msg = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetPageNo(v int32) *QuotaListExportPagedResponseBody {
	s.PageNo = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetPageSize(v int32) *QuotaListExportPagedResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetRequestId(v string) *QuotaListExportPagedResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuotaListExportPagedResponseBody) SetTotal(v int32) *QuotaListExportPagedResponseBody {
	s.Total = &v
	return s
}

type QuotaListExportPagedResponseBodyData struct {
	// Create Time
	//
	// example:
	//
	// 2023-12-21 21:31:57 UTC+8
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// File Name
	//
	// example:
	//
	// 5113766248601929_quota_2023-06-22_2023-12-21_all_2023122121310057
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Notification Message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Display of Task Status
	//
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task Status Enum</br>
	//
	// 2: Exporting</br>
	//
	// 3: Export Success</br>
	//
	// -1: Export Fail</br>
	//
	// example:
	//
	// Export Success
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The link to download exported file.
	//
	// example:
	//
	// //aliyun-eco-market-servic-singapore.oss-ap-southeast-1.aliyuncs.com/5113766248601929_quota_2023-06-22_2023-12-21_all_2023122121310057
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QuotaListExportPagedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuotaListExportPagedResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedResponseBodyData) SetCreateTime(v string) *QuotaListExportPagedResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetFileName(v string) *QuotaListExportPagedResponseBodyData {
	s.FileName = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetMessage(v string) *QuotaListExportPagedResponseBodyData {
	s.Message = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetStatus(v string) *QuotaListExportPagedResponseBodyData {
	s.Status = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetStatusCode(v string) *QuotaListExportPagedResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *QuotaListExportPagedResponseBodyData) SetUrl(v string) *QuotaListExportPagedResponseBodyData {
	s.Url = &v
	return s
}

type QuotaListExportPagedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuotaListExportPagedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuotaListExportPagedResponse) String() string {
	return tea.Prettify(s)
}

func (s QuotaListExportPagedResponse) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedResponse) SetHeaders(v map[string]*string) *QuotaListExportPagedResponse {
	s.Headers = v
	return s
}

func (s *QuotaListExportPagedResponse) SetStatusCode(v int32) *QuotaListExportPagedResponse {
	s.StatusCode = &v
	return s
}

func (s *QuotaListExportPagedResponse) SetBody(v *QuotaListExportPagedResponseBody) *QuotaListExportPagedResponse {
	s.Body = v
	return s
}

type ResendEmailRequest struct {
	// Invitation ID, from interface InviteSubAccount </br>
	//
	// Note: This field type is Long, which may result in precision loss in serialization/deserialization process. Please ensure the value does not exceed 9007199254740991.
	//
	// This parameter is required.
	//
	// example:
	//
	// 176
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
}

func (s ResendEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendEmailRequest) SetInviteId(v int64) *ResendEmailRequest {
	s.InviteId = &v
	return s
}

type ResendEmailResponseBody struct {
	// Result Code, Error code.</br>
	//
	// Candidate Value: </br>
	//
	// 	- 200: OK
	//
	// 	- 1109: System error
	//
	// 	- 3058: Frequent sending, the limit is 10 emails in every 5 minutes.
	//
	// 	- 3057: InviteId is empty.
	//
	// 	- 3060: Can\\"t find sending record of given InviteId.
	//
	// 	- 3061: Registration URL is expired, unable to resend.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Result message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResendEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendEmailResponseBody) SetCode(v string) *ResendEmailResponseBody {
	s.Code = &v
	return s
}

func (s *ResendEmailResponseBody) SetMessage(v string) *ResendEmailResponseBody {
	s.Message = &v
	return s
}

func (s *ResendEmailResponseBody) SetRequestId(v string) *ResendEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendEmailResponseBody) SetSuccess(v bool) *ResendEmailResponseBody {
	s.Success = &v
	return s
}

type ResendEmailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendEmailResponse) SetHeaders(v map[string]*string) *ResendEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendEmailResponse) SetStatusCode(v int32) *ResendEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendEmailResponse) SetBody(v *ResendEmailResponseBody) *ResendEmailResponse {
	s.Body = v
	return s
}

type SetAccountInfoRequest struct {
	// Result Code:
	//
	// 	- 200 OK
	//
	// 	- 1109 System error
	//
	// 	- 3030 Sub Account Nickname exceeds maximum length,  maximum length 150 bytes.
	//
	// 	- 3031 Remark exceeds maximum length,  maximum length 3000 bytes.
	//
	// example:
	//
	// Message information
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// Customer manager（limited 50 character）
	//
	// example:
	//
	// abc
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// success
	//
	// example:
	//
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *SetAccountInfoRequest) SetAccountNickname(v string) *SetAccountInfoRequest {
	s.AccountNickname = &v
	return s
}

func (s *SetAccountInfoRequest) SetCustomerBd(v string) *SetAccountInfoRequest {
	s.CustomerBd = &v
	return s
}

func (s *SetAccountInfoRequest) SetRemark(v string) *SetAccountInfoRequest {
	s.Remark = &v
	return s
}

func (s *SetAccountInfoRequest) SetUid(v int64) *SetAccountInfoRequest {
	s.Uid = &v
	return s
}

type SetAccountInfoResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountInfoResponseBody) SetCode(v string) *SetAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetMessage(v string) *SetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetRequestId(v string) *SetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetSuccess(v bool) *SetAccountInfoResponseBody {
	s.Success = &v
	return s
}

type SetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *SetAccountInfoResponse) SetHeaders(v map[string]*string) *SetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *SetAccountInfoResponse) SetStatusCode(v int32) *SetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccountInfoResponse) SetBody(v *SetAccountInfoResponseBody) *SetAccountInfoResponse {
	s.Body = v
	return s
}

type SetCreditLineRequest struct {
	// New Credit Line
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The UID of Sub Account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1263644979775567
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetCreditLineRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLineRequest) GoString() string {
	return s.String()
}

func (s *SetCreditLineRequest) SetCreditLine(v string) *SetCreditLineRequest {
	s.CreditLine = &v
	return s
}

func (s *SetCreditLineRequest) SetUid(v int64) *SetCreditLineRequest {
	s.Uid = &v
	return s
}

type SetCreditLineResponseBody struct {
	// Result Code:
	//
	// 	- 200 OK
	//
	// 	- 1109 system error
	//
	// 	- 3040 Sub Account is in a frozen state and cannot be operated.
	//
	// 	- 3041 Credit is not a proper number
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Same as Code parameter value
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetCreditLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLineResponseBody) GoString() string {
	return s.String()
}

func (s *SetCreditLineResponseBody) SetCode(v string) *SetCreditLineResponseBody {
	s.Code = &v
	return s
}

func (s *SetCreditLineResponseBody) SetMessage(v string) *SetCreditLineResponseBody {
	s.Message = &v
	return s
}

func (s *SetCreditLineResponseBody) SetRequestId(v string) *SetCreditLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCreditLineResponseBody) SetSuccess(v bool) *SetCreditLineResponseBody {
	s.Success = &v
	return s
}

type SetCreditLineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCreditLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCreditLineResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLineResponse) GoString() string {
	return s.String()
}

func (s *SetCreditLineResponse) SetHeaders(v map[string]*string) *SetCreditLineResponse {
	s.Headers = v
	return s
}

func (s *SetCreditLineResponse) SetStatusCode(v int32) *SetCreditLineResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCreditLineResponse) SetBody(v *SetCreditLineResponseBody) *SetCreditLineResponse {
	s.Body = v
	return s
}

type SetWarningThresholdRequest struct {
	// The UID of the partner‘s customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792155717328010
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// Percentage, 1 to 100. When the available credit limit is lower than the credit limit percentage, an email is sent to the main account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	WarningValue *string `json:"WarningValue,omitempty" xml:"WarningValue,omitempty"`
}

func (s SetWarningThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s SetWarningThresholdRequest) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdRequest) SetUid(v int64) *SetWarningThresholdRequest {
	s.Uid = &v
	return s
}

func (s *SetWarningThresholdRequest) SetWarningValue(v string) *SetWarningThresholdRequest {
	s.WarningValue = &v
	return s
}

type SetWarningThresholdResponseBody struct {
	// Result Code:
	//
	// 	- 200 OK
	//
	// 	- 1109 System Error
	//
	// 	- 3040 The Sub Account is frozen, the operation cannot be completed.
	//
	// 	- 3044 Alert proportion value is not a number.
	//
	// 	- 3045 Alert proportion value should between 1 to 100.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Same as Code parameter value
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, the unique request identifier generated by Alibaba Cloud.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True or False, which indicates whether the current API call itself is successful. does not represent the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetWarningThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetWarningThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdResponseBody) SetCode(v string) *SetWarningThresholdResponseBody {
	s.Code = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetMessage(v string) *SetWarningThresholdResponseBody {
	s.Message = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetRequestId(v string) *SetWarningThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetSuccess(v bool) *SetWarningThresholdResponseBody {
	s.Success = &v
	return s
}

type SetWarningThresholdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWarningThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWarningThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s SetWarningThresholdResponse) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdResponse) SetHeaders(v map[string]*string) *SetWarningThresholdResponse {
	s.Headers = v
	return s
}

func (s *SetWarningThresholdResponse) SetStatusCode(v int32) *SetWarningThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWarningThresholdResponse) SetBody(v *SetWarningThresholdResponseBody) *SetWarningThresholdResponse {
	s.Body = v
	return s
}

type SubscriptionBillRequest struct {
	// The start month from which the bills are pushed. Specify the value in the yyyy-MM format.
	//
	// After the subscription is generated, the system automatically pushes the bill data that is generated from the month that you specified to the current point in time. Data of up to six months can be pushed. The current month is included. If you subscribe to the bills for more than six months, the subscription is invalid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The file format of the bill. Valid values: csv and parquet.
	//
	// If you subscribe to the bills of multiple file formats, we recommend that you store the bills in different OSS buckets to prevent file overwriting.
	//
	// This parameter is required.
	//
	// example:
	//
	// csv
	BillFormat *string `json:"BillFormat,omitempty" xml:"BillFormat,omitempty"`
	// The ID of the user to which the OSS bucket belongs.
	//
	// If you are an eco-partner of Alibaba Cloud and you need to push the bills to the OSS bucket of a subordinate partner account, you must set this parameter to the ID of the subordinate partner account and grant the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission to the subordinate partner account.
	//
	// If you are an eco-partner of Alibaba Cloud and you need to push the bills to the OSS bucket of your own account, your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5569414254138836
	BucketOwnerId *int64 `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	// The name of the Object Storage Service (OSS) bucket in which you want to store the bills.
	//
	// This parameter is required.
	//
	// example:
	//
	// bill-bucket
	SubscribeBucket *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	// The maximum rows in a single bill file. If the number of bill rows exceed the upper limit, the bill is automatically split into multiple files. The name of each split file is in the `uid_billType_billCycle_SquenceNo_fileNo` format.
	//
	// Files whose names are the same except for the fileNo field are of the same type and belong to the same billing cycle.
	//
	// example:
	//
	// 100000
	SubscribeSegmentSize *int32 `json:"SubscribeSegmentSize,omitempty" xml:"SubscribeSegmentSize,omitempty"`
	// The type of the bill to which you want to subscribe. Valid values: PartnerBillingItemDetailForBillingPeriod, PartnerBillingItemDetailMonthly, PartnerInstanceDetailForBillingPeriod, and PartnerInstanceDetailMonthly.
	//
	// This parameter is required.
	//
	// example:
	//
	// PartnerBillingItemDetailForBillingPeriod
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s SubscriptionBillRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionBillRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionBillRequest) SetBeginBillingCycle(v string) *SubscriptionBillRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *SubscriptionBillRequest) SetBillFormat(v string) *SubscriptionBillRequest {
	s.BillFormat = &v
	return s
}

func (s *SubscriptionBillRequest) SetBucketOwnerId(v int64) *SubscriptionBillRequest {
	s.BucketOwnerId = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeBucket(v string) *SubscriptionBillRequest {
	s.SubscribeBucket = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeSegmentSize(v int32) *SubscriptionBillRequest {
	s.SubscribeSegmentSize = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeType(v string) *SubscriptionBillRequest {
	s.SubscribeType = &v
	return s
}

type SubscriptionBillResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubscriptionBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionBillResponseBody) GoString() string {
	return s.String()
}

func (s *SubscriptionBillResponseBody) SetCode(v string) *SubscriptionBillResponseBody {
	s.Code = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetData(v bool) *SubscriptionBillResponseBody {
	s.Data = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetMessage(v string) *SubscriptionBillResponseBody {
	s.Message = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetRequestId(v string) *SubscriptionBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetSuccess(v bool) *SubscriptionBillResponseBody {
	s.Success = &v
	return s
}

type SubscriptionBillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscriptionBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscriptionBillResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionBillResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionBillResponse) SetHeaders(v map[string]*string) *SubscriptionBillResponse {
	s.Headers = v
	return s
}

func (s *SubscriptionBillResponse) SetStatusCode(v int32) *SubscriptionBillResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscriptionBillResponse) SetBody(v *SubscriptionBillResponseBody) *SubscriptionBillResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("agency.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("agency.aliyuncs.com"),
		"ap-south-1":                  tea.String("agency.aliyuncs.com"),
		"ap-southeast-2":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-3":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-5":              tea.String("agency.aliyuncs.com"),
		"cn-beijing":                  tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("agency.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("agency.aliyuncs.com"),
		"cn-chengdu":                  tea.String("agency.aliyuncs.com"),
		"cn-edge-1":                   tea.String("agency.aliyuncs.com"),
		"cn-fujian":                   tea.String("agency.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("agency.aliyuncs.com"),
		"cn-hongkong":                 tea.String("agency.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("agency.aliyuncs.com"),
		"cn-huhehaote":                tea.String("agency.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("agency.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-qingdao":                  tea.String("agency.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai":                 tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("agency.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("agency.aliyuncs.com"),
		"cn-wuhan":                    tea.String("agency.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("agency.aliyuncs.com"),
		"cn-yushanfang":               tea.String("agency.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("agency.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("agency.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("agency.aliyuncs.com"),
		"eu-central-1":                tea.String("agency.aliyuncs.com"),
		"eu-west-1":                   tea.String("agency.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("agency.aliyuncs.com"),
		"me-east-1":                   tea.String("agency.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("agency.aliyuncs.com"),
		"us-east-1":                   tea.String("agency.aliyuncs.com"),
		"us-west-1":                   tea.String("agency.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("agency"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
// Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
// You can call this operation to cancel the subscription to only one type of bill at a time.
//
// After the subscription to a type of bill is canceled, bills of this type are no longer pushed to the specified Object Storage Service (OSS) bucket.
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - CancelSubscriptionBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSubscriptionBillResponse
func (client *Client) CancelSubscriptionBillWithOptions(request *CancelSubscriptionBillRequest, runtime *util.RuntimeOptions) (_result *CancelSubscriptionBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscribeType)) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelSubscriptionBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelSubscriptionBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
// Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
// You can call this operation to cancel the subscription to only one type of bill at a time.
//
// After the subscription to a type of bill is canceled, bills of this type are no longer pushed to the specified Object Storage Service (OSS) bucket.
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - CancelSubscriptionBillRequest
//
// @return CancelSubscriptionBillResponse
func (client *Client) CancelSubscriptionBill(request *CancelSubscriptionBillRequest) (_result *CancelSubscriptionBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelSubscriptionBillResponse{}
	_body, _err := client.CancelSubscriptionBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建优惠券模板
//
// @param tmpReq - CreateCouponTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCouponTemplateResponse
func (client *Client) CreateCouponTemplateWithOptions(tmpReq *CreateCouponTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateCouponTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCouponTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ProductType)) {
		request.ProductTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductType, tea.String("ProductType"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicableProducts)) {
		query["ApplicableProducts"] = request.ApplicableProducts
	}

	if !tea.BoolValue(util.IsUnset(request.CostBearer)) {
		query["CostBearer"] = request.CostBearer
	}

	if !tea.BoolValue(util.IsUnset(request.CouponDescription)) {
		query["CouponDescription"] = request.CouponDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Expireddate)) {
		query["Expireddate"] = request.Expireddate
	}

	if !tea.BoolValue(util.IsUnset(request.LimitPerPerson)) {
		query["LimitPerPerson"] = request.LimitPerPerson
	}

	if !tea.BoolValue(util.IsUnset(request.ProductTypeShrink)) {
		query["ProductType"] = request.ProductTypeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseType)) {
		query["PurchaseType"] = request.PurchaseType
	}

	if !tea.BoolValue(util.IsUnset(request.ReasonForApplication)) {
		query["ReasonForApplication"] = request.ReasonForApplication
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Vailddate)) {
		query["Vailddate"] = request.Vailddate
	}

	if !tea.BoolValue(util.IsUnset(request.Vaildperioddays)) {
		query["Vaildperioddays"] = request.Vaildperioddays
	}

	if !tea.BoolValue(util.IsUnset(request.ValidUntil)) {
		query["ValidUntil"] = request.ValidUntil
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCouponTemplate"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCouponTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建优惠券模板
//
// @param request - CreateCouponTemplateRequest
//
// @return CreateCouponTemplateResponse
func (client *Client) CreateCouponTemplate(request *CreateCouponTemplateRequest) (_result *CreateCouponTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCouponTemplateResponse{}
	_body, _err := client.CreateCouponTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This function is designed for create a customer who is to be invited.
//
// @param request - CreateCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerName)) {
		query["CustomerName"] = request.CustomerName
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerSource)) {
		query["CustomerSource"] = request.CustomerSource
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerSubTrade)) {
		query["CustomerSubTrade"] = request.CustomerSubTrade
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerTrade)) {
		query["CustomerTrade"] = request.CustomerTrade
	}

	if !tea.BoolValue(util.IsUnset(request.Nation)) {
		query["Nation"] = request.Nation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomer"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function is designed for create a customer who is to be invited.
//
// @param request - CreateCustomerRequest
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query quota adjustment list of Distribution Customer from International Site. Not available on Domestic Site.
//
// @param request - CustomerQuotaRecordListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerQuotaRecordListResponse
func (client *Client) CustomerQuotaRecordListWithOptions(request *CustomerQuotaRecordListRequest, runtime *util.RuntimeOptions) (_result *CustomerQuotaRecordListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomerQuotaRecordList"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomerQuotaRecordListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query quota adjustment list of Distribution Customer from International Site. Not available on Domestic Site.
//
// @param request - CustomerQuotaRecordListRequest
//
// @return CustomerQuotaRecordListResponse
func (client *Client) CustomerQuotaRecordList(request *CustomerQuotaRecordListRequest) (_result *CustomerQuotaRecordListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CustomerQuotaRecordListResponse{}
	_body, _err := client.CustomerQuotaRecordListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API is used to offset the Deducted Credit of a Distribution Customer. For example, if the current Deducted Credit is 500 and the Available Credit is 1000, by offsetting 300, the Deducted Credit will then become 200, and the Available Credit becomes 1300.
//
// Description:
//
// Note that sometimes you may find that the customer\\"s Used Credit is negative. This indicates that there is no need to restore the Used Credit, and its ready for customer\\"s usage. This phenomenon occurs because a refund is generated while the customer\\"s credit is full, thereby triggered additional increasing on the customer\\"s credit.
//
// For example, if the customer\\"s maximum Available Credit is 1000 with no usage, and a refund of 300 occurs, the Used Credit will become -300.
//
// @param request - DeductOutstandingBalanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeductOutstandingBalanceResponse
func (client *Client) DeductOutstandingBalanceWithOptions(request *DeductOutstandingBalanceRequest, runtime *util.RuntimeOptions) (_result *DeductOutstandingBalanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeductAmount)) {
		query["DeductAmount"] = request.DeductAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeductOutstandingBalance"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeductOutstandingBalanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to offset the Deducted Credit of a Distribution Customer. For example, if the current Deducted Credit is 500 and the Available Credit is 1000, by offsetting 300, the Deducted Credit will then become 200, and the Available Credit becomes 1300.
//
// Description:
//
// Note that sometimes you may find that the customer\\"s Used Credit is negative. This indicates that there is no need to restore the Used Credit, and its ready for customer\\"s usage. This phenomenon occurs because a refund is generated while the customer\\"s credit is full, thereby triggered additional increasing on the customer\\"s credit.
//
// For example, if the customer\\"s maximum Available Credit is 1000 with no usage, and a refund of 300 occurs, the Used Credit will become -300.
//
// @param request - DeductOutstandingBalanceRequest
//
// @return DeductOutstandingBalanceResponse
func (client *Client) DeductOutstandingBalance(request *DeductOutstandingBalanceRequest) (_result *DeductOutstandingBalanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeductOutstandingBalanceResponse{}
	_body, _err := client.DeductOutstandingBalanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set the after-shutdown instance status for post-pay End Users as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditEndUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditEndUserStatusResponse
func (client *Client) EditEndUserStatusWithOptions(request *EditEndUserStatusRequest, runtime *util.RuntimeOptions) (_result *EditEndUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditEndUserStatus"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditEndUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set the after-shutdown instance status for post-pay End Users as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditEndUserStatusRequest
//
// @return EditEndUserStatusResponse
func (client *Client) EditEndUserStatus(request *EditEndUserStatusRequest) (_result *EditEndUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditEndUserStatusResponse{}
	_body, _err := client.EditEndUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set the New Buy status for Sub-Customer as a Partner.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditNewBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditNewBuyStatusResponse
func (client *Client) EditNewBuyStatusWithOptions(request *EditNewBuyStatusRequest, runtime *util.RuntimeOptions) (_result *EditNewBuyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewBuyStatus)) {
		query["NewBuyStatus"] = request.NewBuyStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditNewBuyStatus"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditNewBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set the New Buy status for Sub-Customer as a Partner.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditNewBuyStatusRequest
//
// @return EditNewBuyStatusResponse
func (client *Client) EditNewBuyStatus(request *EditNewBuyStatusRequest) (_result *EditNewBuyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditNewBuyStatusResponse{}
	_body, _err := client.EditNewBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the End User\\"s Shutdown Policy as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditZeroCreditShutdownRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditZeroCreditShutdownResponse
func (client *Client) EditZeroCreditShutdownWithOptions(request *EditZeroCreditShutdownRequest, runtime *util.RuntimeOptions) (_result *EditZeroCreditShutdownResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShutdownPolicy)) {
		query["ShutdownPolicy"] = request.ShutdownPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditZeroCreditShutdown"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditZeroCreditShutdownResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the End User\\"s Shutdown Policy as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditZeroCreditShutdownRequest
//
// @return EditZeroCreditShutdownResponse
func (client *Client) EditZeroCreditShutdown(request *EditZeroCreditShutdownRequest) (_result *EditZeroCreditShutdownResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditZeroCreditShutdownResponse{}
	_body, _err := client.EditZeroCreditShutdownWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Export quota amount adjustment history as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - ExportCustomerQuotaRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCustomerQuotaRecordResponse
func (client *Client) ExportCustomerQuotaRecordWithOptions(request *ExportCustomerQuotaRecordRequest, runtime *util.RuntimeOptions) (_result *ExportCustomerQuotaRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserPk)) {
		query["EndUserPk"] = request.EndUserPk
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportCustomerQuotaRecord"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportCustomerQuotaRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Export quota amount adjustment history as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - ExportCustomerQuotaRecordRequest
//
// @return ExportCustomerQuotaRecordResponse
func (client *Client) ExportCustomerQuotaRecord(request *ExportCustomerQuotaRecordRequest) (_result *ExportCustomerQuotaRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportCustomerQuotaRecordResponse{}
	_body, _err := client.ExportCustomerQuotaRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Return Distribution Customer\\"s account information.
//
// @param request - GetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfoWithOptions(request *GetAccountInfoRequest, runtime *util.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountInfo"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Return Distribution Customer\\"s account information.
//
// @param request - GetAccountInfoRequest
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfo(request *GetAccountInfoRequest) (_result *GetAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.GetAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际渠道分销优惠券可抵扣产品
//
// @param request - GetCoupondeductProductCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCoupondeductProductCodeResponse
func (client *Client) GetCoupondeductProductCodeWithOptions(request *GetCoupondeductProductCodeRequest, runtime *util.RuntimeOptions) (_result *GetCoupondeductProductCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCoupondeductProductCode"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCoupondeductProductCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际渠道分销优惠券可抵扣产品
//
// @param request - GetCoupondeductProductCodeRequest
//
// @return GetCoupondeductProductCodeResponse
func (client *Client) GetCoupondeductProductCode(request *GetCoupondeductProductCodeRequest) (_result *GetCoupondeductProductCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCoupondeductProductCodeResponse{}
	_body, _err := client.GetCoupondeductProductCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Credit Control information of Distribution Customers. The PopCreditInfoJson in the Return Parameter will be empty if the Distribution Customer is an Agency. This function is only available for Resellers and Distributors.
//
// @param request - GetCreditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreditInfoResponse
func (client *Client) GetCreditInfoWithOptions(request *GetCreditInfoRequest, runtime *util.RuntimeOptions) (_result *GetCreditInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreditInfo"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Credit Control information of Distribution Customers. The PopCreditInfoJson in the Return Parameter will be empty if the Distribution Customer is an Agency. This function is only available for Resellers and Distributors.
//
// @param request - GetCreditInfoRequest
//
// @return GetCreditInfoResponse
func (client *Client) GetCreditInfo(request *GetCreditInfoRequest) (_result *GetCreditInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreditInfoResponse{}
	_body, _err := client.GetCreditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 客户订单查询
//
// @param request - GetCustomerOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerOrdersResponse
func (client *Client) GetCustomerOrdersWithOptions(request *GetCustomerOrdersRequest, runtime *util.RuntimeOptions) (_result *GetCustomerOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomerOrders"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomerOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户订单查询
//
// @param request - GetCustomerOrdersRequest
//
// @return GetCustomerOrdersResponse
func (client *Client) GetCustomerOrders(request *GetCustomerOrdersRequest) (_result *GetCustomerOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomerOrdersResponse{}
	_body, _err := client.GetCustomerOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s daily Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetDailyBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDailyBillResponse
func (client *Client) GetDailyBillWithOptions(request *GetDailyBillRequest, runtime *util.RuntimeOptions) (_result *GetDailyBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillOwner)) {
		query["BillOwner"] = request.BillOwner
	}

	if !tea.BoolValue(util.IsUnset(request.BillType)) {
		query["BillType"] = request.BillType
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["Date"] = request.Date
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDailyBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDailyBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s daily Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetDailyBillRequest
//
// @return GetDailyBillResponse
func (client *Client) GetDailyBill(request *GetDailyBillRequest) (_result *GetDailyBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDailyBillResponse{}
	_body, _err := client.GetDailyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query invitation status of customer who have been created and invited.
//
// @param request - GetInviteStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInviteStatusResponse
func (client *Client) GetInviteStatusWithOptions(request *GetInviteStatusRequest, runtime *util.RuntimeOptions) (_result *GetInviteStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteStatusList)) {
		query["InviteStatusList"] = request.InviteStatusList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInviteStatus"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInviteStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query invitation status of customer who have been created and invited.
//
// @param request - GetInviteStatusRequest
//
// @return GetInviteStatusResponse
func (client *Client) GetInviteStatus(request *GetInviteStatusRequest) (_result *GetInviteStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInviteStatusResponse{}
	_body, _err := client.GetInviteStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s Monthly Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetMonthlyBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonthlyBillResponse
func (client *Client) GetMonthlyBillWithOptions(request *GetMonthlyBillRequest, runtime *util.RuntimeOptions) (_result *GetMonthlyBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillOwner)) {
		query["BillOwner"] = request.BillOwner
	}

	if !tea.BoolValue(util.IsUnset(request.BillType)) {
		query["BillType"] = request.BillType
	}

	if !tea.BoolValue(util.IsUnset(request.Month)) {
		query["Month"] = request.Month
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMonthlyBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMonthlyBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s Monthly Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetMonthlyBillRequest
//
// @return GetMonthlyBillResponse
func (client *Client) GetMonthlyBill(request *GetMonthlyBillRequest) (_result *GetMonthlyBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMonthlyBillResponse{}
	_body, _err := client.GetMonthlyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query all the Unassociated Customer.
//
// @param request - GetUnassociatedCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUnassociatedCustomerResponse
func (client *Client) GetUnassociatedCustomerWithOptions(request *GetUnassociatedCustomerRequest, runtime *util.RuntimeOptions) (_result *GetUnassociatedCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUnassociatedCustomer"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUnassociatedCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query all the Unassociated Customer.
//
// @param request - GetUnassociatedCustomerRequest
//
// @return GetUnassociatedCustomerResponse
func (client *Client) GetUnassociatedCustomer(request *GetUnassociatedCustomerRequest) (_result *GetUnassociatedCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUnassociatedCustomerResponse{}
	_body, _err := client.GetUnassociatedCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiate the Partner registration invitation.
//
// Description:
//
// The current API request rate for the Cloud Product has not been disclosed.
//
// @param request - InviteSubAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteSubAccountResponse
func (client *Client) InviteSubAccountWithOptions(request *InviteSubAccountRequest, runtime *util.RuntimeOptions) (_result *InviteSubAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountInfoList)) {
		query["AccountInfoList"] = request.AccountInfoList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InviteSubAccount"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InviteSubAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate the Partner registration invitation.
//
// Description:
//
// The current API request rate for the Cloud Product has not been disclosed.
//
// @param request - InviteSubAccountRequest
//
// @return InviteSubAccountResponse
func (client *Client) InviteSubAccount(request *InviteSubAccountRequest) (_result *InviteSubAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InviteSubAccountResponse{}
	_body, _err := client.InviteSubAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发放优惠券
//
// @param request - IssueCouponForCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IssueCouponForCustomerResponse
func (client *Client) IssueCouponForCustomerWithOptions(request *IssueCouponForCustomerRequest, runtime *util.RuntimeOptions) (_result *IssueCouponForCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.CouponTemplateId)) {
		query["CouponTemplateId"] = request.CouponTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Uidlist)) {
		query["Uidlist"] = request.Uidlist
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IssueCouponForCustomer"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IssueCouponForCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发放优惠券
//
// @param request - IssueCouponForCustomerRequest
//
// @return IssueCouponForCustomerResponse
func (client *Client) IssueCouponForCustomer(request *IssueCouponForCustomerRequest) (_result *IssueCouponForCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IssueCouponForCustomerResponse{}
	_body, _err := client.IssueCouponForCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This function is available for all Distributors. It displays the corresponding region code information based on the operable countries as agreed in the Distributor\\"s contract.
//
// Description:
//
// The current API request rate for cloud products has not been disclosed.
//
// @param request - ListCountriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCountriesResponse
func (client *Client) ListCountriesWithOptions(runtime *util.RuntimeOptions) (_result *ListCountriesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListCountries"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCountriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function is available for all Distributors. It displays the corresponding region code information based on the operable countries as agreed in the Distributor\\"s contract.
//
// Description:
//
// The current API request rate for cloud products has not been disclosed.
//
// @return ListCountriesResponse
func (client *Client) ListCountries() (_result *ListCountriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCountriesResponse{}
	_body, _err := client.ListCountriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 优惠券使用量列表查询
//
// @param request - ListCouponUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCouponUsageResponse
func (client *Client) ListCouponUsageWithOptions(request *ListCouponUsageRequest, runtime *util.RuntimeOptions) (_result *ListCouponUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.CouponTemplateId)) {
		query["CouponTemplateId"] = request.CouponTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCouponUsage"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCouponUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 优惠券使用量列表查询
//
// @param request - ListCouponUsageRequest
//
// @return ListCouponUsageResponse
func (client *Client) ListCouponUsage(request *ListCouponUsageRequest) (_result *ListCouponUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCouponUsageResponse{}
	_body, _err := client.ListCouponUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check the result of export quota list as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - QuotaListExportPagedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuotaListExportPagedResponse
func (client *Client) QuotaListExportPagedWithOptions(request *QuotaListExportPagedRequest, runtime *util.RuntimeOptions) (_result *QuotaListExportPagedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuotaListExportPaged"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuotaListExportPagedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check the result of export quota list as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - QuotaListExportPagedRequest
//
// @return QuotaListExportPagedResponse
func (client *Client) QuotaListExportPaged(request *QuotaListExportPagedRequest) (_result *QuotaListExportPagedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuotaListExportPagedResponse{}
	_body, _err := client.QuotaListExportPagedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resend invitation email.
//
// @param request - ResendEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendEmailResponse
func (client *Client) ResendEmailWithOptions(request *ResendEmailRequest, runtime *util.RuntimeOptions) (_result *ResendEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteId)) {
		query["InviteId"] = request.InviteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendEmail"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resend invitation email.
//
// @param request - ResendEmailRequest
//
// @return ResendEmailResponse
func (client *Client) ResendEmail(request *ResendEmailRequest) (_result *ResendEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendEmailResponse{}
	_body, _err := client.ResendEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This function is designed for Sub Account information maintenance, including Nickname and Remark.
//
// @param request - SetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccountInfoResponse
func (client *Client) SetAccountInfoWithOptions(request *SetAccountInfoRequest, runtime *util.RuntimeOptions) (_result *SetAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNickname)) {
		query["AccountNickname"] = request.AccountNickname
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerBd)) {
		query["CustomerBd"] = request.CustomerBd
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccountInfo"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function is designed for Sub Account information maintenance, including Nickname and Remark.
//
// @param request - SetAccountInfoRequest
//
// @return SetAccountInfoResponse
func (client *Client) SetAccountInfo(request *SetAccountInfoRequest) (_result *SetAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccountInfoResponse{}
	_body, _err := client.SetAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set Credit Line for Distribution Customers. This function is only available for Resellers and Distributors.
//
// @param request - SetCreditLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCreditLineResponse
func (client *Client) SetCreditLineWithOptions(request *SetCreditLineRequest, runtime *util.RuntimeOptions) (_result *SetCreditLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreditLine)) {
		query["CreditLine"] = request.CreditLine
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetCreditLine"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetCreditLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set Credit Line for Distribution Customers. This function is only available for Resellers and Distributors.
//
// @param request - SetCreditLineRequest
//
// @return SetCreditLineResponse
func (client *Client) SetCreditLine(request *SetCreditLineRequest) (_result *SetCreditLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCreditLineResponse{}
	_body, _err := client.SetCreditLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can use this API to set the threshold for the use of credit control. When the customer credit control reaches below the threshold, it will pass through the notification email distributor. This feature is for Reseller and Distributor only.
//
// @param request - SetWarningThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWarningThresholdResponse
func (client *Client) SetWarningThresholdWithOptions(request *SetWarningThresholdRequest, runtime *util.RuntimeOptions) (_result *SetWarningThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.WarningValue)) {
		query["WarningValue"] = request.WarningValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetWarningThreshold"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetWarningThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can use this API to set the threshold for the use of credit control. When the customer credit control reaches below the threshold, it will pass through the notification email distributor. This feature is for Reseller and Distributor only.
//
// @param request - SetWarningThresholdRequest
//
// @return SetWarningThresholdResponse
func (client *Client) SetWarningThreshold(request *SetWarningThresholdRequest) (_result *SetWarningThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetWarningThresholdResponse{}
	_body, _err := client.SetWarningThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
//   Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
// 	- You can call this operation to subscribe to only one type of bill at a time.
//
// 	- After the subscription to a type of bill is generated, the bill for the previous day is pushed on a daily basis from the next day. On the fifth day of each month, the full-data bill for the previous month is pushed.
//
// 	- A daily bill may be delayed. The delayed bill is pushed the next day after it is generated. The delayed bill may contain the bill data that is delayed until the previous day. We recommend that you query the full-data bill for the previous month at the beginning of each month.
//
// 	- Your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
//
// 	- The following file name formats are supported for bills:
//
// ```
//
// BillingItemDetailForBillingPeriod
//
//
//
// File name format of a daily bill: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
//
//
//
// File name format of a monthly full-data bill: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
//
// InstanceDetailForBillingPeriod
//
//
//
//  File name format of a daily bill: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
//
//
//
// File name format of a monthly full-data bill: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
//
// BillingItemDetailMonthly
//
//
//
// File name format of a daily bill: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// InstanceDetailMonthly
//
//
//
// File name format of a daily bill: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// The fileNo field exists only when the number of bill rows reaches the maximum rows in a single bill file and the bill is split into multiple files.
//
// ```
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - SubscriptionBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscriptionBillResponse
func (client *Client) SubscriptionBillWithOptions(request *SubscriptionBillRequest, runtime *util.RuntimeOptions) (_result *SubscriptionBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginBillingCycle)) {
		query["BeginBillingCycle"] = request.BeginBillingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.BillFormat)) {
		query["BillFormat"] = request.BillFormat
	}

	if !tea.BoolValue(util.IsUnset(request.BucketOwnerId)) {
		query["BucketOwnerId"] = request.BucketOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeBucket)) {
		query["SubscribeBucket"] = request.SubscribeBucket
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeSegmentSize)) {
		query["SubscribeSegmentSize"] = request.SubscribeSegmentSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeType)) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubscriptionBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscriptionBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
//   Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
// 	- You can call this operation to subscribe to only one type of bill at a time.
//
// 	- After the subscription to a type of bill is generated, the bill for the previous day is pushed on a daily basis from the next day. On the fifth day of each month, the full-data bill for the previous month is pushed.
//
// 	- A daily bill may be delayed. The delayed bill is pushed the next day after it is generated. The delayed bill may contain the bill data that is delayed until the previous day. We recommend that you query the full-data bill for the previous month at the beginning of each month.
//
// 	- Your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
//
// 	- The following file name formats are supported for bills:
//
// ```
//
// BillingItemDetailForBillingPeriod
//
//
//
// File name format of a daily bill: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
//
//
//
// File name format of a monthly full-data bill: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
//
// InstanceDetailForBillingPeriod
//
//
//
//  File name format of a daily bill: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
//
//
//
// File name format of a monthly full-data bill: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
//
// BillingItemDetailMonthly
//
//
//
// File name format of a daily bill: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// InstanceDetailMonthly
//
//
//
// File name format of a daily bill: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// The fileNo field exists only when the number of bill rows reaches the maximum rows in a single bill file and the bill is split into multiple files.
//
// ```
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - SubscriptionBillRequest
//
// @return SubscriptionBillResponse
func (client *Client) SubscriptionBill(request *SubscriptionBillRequest) (_result *SubscriptionBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscriptionBillResponse{}
	_body, _err := client.SubscriptionBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
