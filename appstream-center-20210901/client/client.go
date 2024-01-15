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

type AccessPageGetAclRequest struct {
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
}

func (s AccessPageGetAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessPageGetAclRequest) GoString() string {
	return s.String()
}

func (s *AccessPageGetAclRequest) SetAccessPageId(v string) *AccessPageGetAclRequest {
	s.AccessPageId = &v
	return s
}

type AccessPageGetAclResponseBody struct {
	Code    *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*AccessPageGetAclResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccessPageGetAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccessPageGetAclResponseBody) GoString() string {
	return s.String()
}

func (s *AccessPageGetAclResponseBody) SetCode(v string) *AccessPageGetAclResponseBody {
	s.Code = &v
	return s
}

func (s *AccessPageGetAclResponseBody) SetData(v []*AccessPageGetAclResponseBodyData) *AccessPageGetAclResponseBody {
	s.Data = v
	return s
}

func (s *AccessPageGetAclResponseBody) SetMessage(v string) *AccessPageGetAclResponseBody {
	s.Message = &v
	return s
}

func (s *AccessPageGetAclResponseBody) SetRequestId(v string) *AccessPageGetAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessPageGetAclResponseBody) SetSuccess(v string) *AccessPageGetAclResponseBody {
	s.Success = &v
	return s
}

type AccessPageGetAclResponseBodyData struct {
	AccessMode    *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	AccessUrl     *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
	EffectTime    *int32  `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	Unit          *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	UrlExpireTime *string `json:"UrlExpireTime,omitempty" xml:"UrlExpireTime,omitempty"`
}

func (s AccessPageGetAclResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AccessPageGetAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccessPageGetAclResponseBodyData) SetAccessMode(v string) *AccessPageGetAclResponseBodyData {
	s.AccessMode = &v
	return s
}

func (s *AccessPageGetAclResponseBodyData) SetAccessUrl(v string) *AccessPageGetAclResponseBodyData {
	s.AccessUrl = &v
	return s
}

func (s *AccessPageGetAclResponseBodyData) SetEffectTime(v int32) *AccessPageGetAclResponseBodyData {
	s.EffectTime = &v
	return s
}

func (s *AccessPageGetAclResponseBodyData) SetUnit(v string) *AccessPageGetAclResponseBodyData {
	s.Unit = &v
	return s
}

func (s *AccessPageGetAclResponseBodyData) SetUrlExpireTime(v string) *AccessPageGetAclResponseBodyData {
	s.UrlExpireTime = &v
	return s
}

type AccessPageGetAclResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AccessPageGetAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AccessPageGetAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessPageGetAclResponse) GoString() string {
	return s.String()
}

func (s *AccessPageGetAclResponse) SetHeaders(v map[string]*string) *AccessPageGetAclResponse {
	s.Headers = v
	return s
}

func (s *AccessPageGetAclResponse) SetStatusCode(v int32) *AccessPageGetAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessPageGetAclResponse) SetBody(v *AccessPageGetAclResponseBody) *AccessPageGetAclResponse {
	s.Body = v
	return s
}

type AccessPageSetAclRequest struct {
	AccessMode     *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	AccessPageId   *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	EffectTime     *int32  `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	Unit           *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s AccessPageSetAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessPageSetAclRequest) GoString() string {
	return s.String()
}

func (s *AccessPageSetAclRequest) SetAccessMode(v string) *AccessPageSetAclRequest {
	s.AccessMode = &v
	return s
}

func (s *AccessPageSetAclRequest) SetAccessPageId(v string) *AccessPageSetAclRequest {
	s.AccessPageId = &v
	return s
}

func (s *AccessPageSetAclRequest) SetAccessPageName(v string) *AccessPageSetAclRequest {
	s.AccessPageName = &v
	return s
}

func (s *AccessPageSetAclRequest) SetEffectTime(v int32) *AccessPageSetAclRequest {
	s.EffectTime = &v
	return s
}

func (s *AccessPageSetAclRequest) SetUnit(v string) *AccessPageSetAclRequest {
	s.Unit = &v
	return s
}

type AccessPageSetAclResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccessPageSetAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccessPageSetAclResponseBody) GoString() string {
	return s.String()
}

func (s *AccessPageSetAclResponseBody) SetCode(v string) *AccessPageSetAclResponseBody {
	s.Code = &v
	return s
}

func (s *AccessPageSetAclResponseBody) SetMessage(v string) *AccessPageSetAclResponseBody {
	s.Message = &v
	return s
}

func (s *AccessPageSetAclResponseBody) SetRequestId(v string) *AccessPageSetAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessPageSetAclResponseBody) SetSuccess(v string) *AccessPageSetAclResponseBody {
	s.Success = &v
	return s
}

type AccessPageSetAclResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AccessPageSetAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AccessPageSetAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessPageSetAclResponse) GoString() string {
	return s.String()
}

func (s *AccessPageSetAclResponse) SetHeaders(v map[string]*string) *AccessPageSetAclResponse {
	s.Headers = v
	return s
}

func (s *AccessPageSetAclResponse) SetStatusCode(v int32) *AccessPageSetAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessPageSetAclResponse) SetBody(v *AccessPageSetAclResponseBody) *AccessPageSetAclResponse {
	s.Body = v
	return s
}

type ApproveOtaTaskRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	OtaType            *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApproveOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskRequest) SetAppInstanceGroupId(v string) *ApproveOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetBizRegionId(v string) *ApproveOtaTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetOtaType(v string) *ApproveOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetStartTime(v string) *ApproveOtaTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetTaskId(v string) *ApproveOtaTaskRequest {
	s.TaskId = &v
	return s
}

type ApproveOtaTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponseBody) SetCode(v string) *ApproveOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetMessage(v string) *ApproveOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetRequestId(v string) *ApproveOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

type ApproveOtaTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApproveOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponse) SetHeaders(v map[string]*string) *ApproveOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ApproveOtaTaskResponse) SetStatusCode(v int32) *ApproveOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOtaTaskResponse) SetBody(v *ApproveOtaTaskResponseBody) *ApproveOtaTaskResponse {
	s.Body = v
	return s
}

type AskSessionPackagePriceRequest struct {
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	MaxSessions        *int32  `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SessionPackageType *string `json:"SessionPackageType,omitempty" xml:"SessionPackageType,omitempty"`
	SessionSpec        *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	SessionType        *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
}

func (s AskSessionPackagePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceRequest) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceRequest) SetChargeType(v string) *AskSessionPackagePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetMaxSessions(v int32) *AskSessionPackagePriceRequest {
	s.MaxSessions = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetPeriod(v int32) *AskSessionPackagePriceRequest {
	s.Period = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetPeriodUnit(v string) *AskSessionPackagePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetRegion(v string) *AskSessionPackagePriceRequest {
	s.Region = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetSessionPackageType(v string) *AskSessionPackagePriceRequest {
	s.SessionPackageType = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetSessionSpec(v string) *AskSessionPackagePriceRequest {
	s.SessionSpec = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetSessionType(v string) *AskSessionPackagePriceRequest {
	s.SessionType = &v
	return s
}

type AskSessionPackagePriceResponseBody struct {
	Data      []*AskSessionPackagePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AskSessionPackagePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponseBody) SetData(v []*AskSessionPackagePriceResponseBodyData) *AskSessionPackagePriceResponseBody {
	s.Data = v
	return s
}

func (s *AskSessionPackagePriceResponseBody) SetRequestId(v string) *AskSessionPackagePriceResponseBody {
	s.RequestId = &v
	return s
}

type AskSessionPackagePriceResponseBodyData struct {
	Price *AskSessionPackagePriceResponseBodyDataPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s AskSessionPackagePriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponseBodyData) SetPrice(v *AskSessionPackagePriceResponseBodyDataPrice) *AskSessionPackagePriceResponseBodyData {
	s.Price = v
	return s
}

type AskSessionPackagePriceResponseBodyDataPrice struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s AskSessionPackagePriceResponseBodyDataPrice) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetCurrency(v string) *AskSessionPackagePriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetDiscountPrice(v float32) *AskSessionPackagePriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetOriginalPrice(v float32) *AskSessionPackagePriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetTradePrice(v float32) *AskSessionPackagePriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

type AskSessionPackagePriceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AskSessionPackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AskSessionPackagePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponse) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponse) SetHeaders(v map[string]*string) *AskSessionPackagePriceResponse {
	s.Headers = v
	return s
}

func (s *AskSessionPackagePriceResponse) SetStatusCode(v int32) *AskSessionPackagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *AskSessionPackagePriceResponse) SetBody(v *AskSessionPackagePriceResponseBody) *AskSessionPackagePriceResponse {
	s.Body = v
	return s
}

type AskSessionPackageRenewPriceRequest struct {
	Period           *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SessionPackageId *string `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
}

func (s AskSessionPackageRenewPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackageRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *AskSessionPackageRenewPriceRequest) SetPeriod(v int32) *AskSessionPackageRenewPriceRequest {
	s.Period = &v
	return s
}

func (s *AskSessionPackageRenewPriceRequest) SetPeriodUnit(v string) *AskSessionPackageRenewPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AskSessionPackageRenewPriceRequest) SetSessionPackageId(v string) *AskSessionPackageRenewPriceRequest {
	s.SessionPackageId = &v
	return s
}

type AskSessionPackageRenewPriceResponseBody struct {
	Data      []*AskSessionPackageRenewPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AskSessionPackageRenewPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackageRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *AskSessionPackageRenewPriceResponseBody) SetData(v []*AskSessionPackageRenewPriceResponseBodyData) *AskSessionPackageRenewPriceResponseBody {
	s.Data = v
	return s
}

func (s *AskSessionPackageRenewPriceResponseBody) SetRequestId(v string) *AskSessionPackageRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

type AskSessionPackageRenewPriceResponseBodyData struct {
	Price *AskSessionPackageRenewPriceResponseBodyDataPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s AskSessionPackageRenewPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackageRenewPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AskSessionPackageRenewPriceResponseBodyData) SetPrice(v *AskSessionPackageRenewPriceResponseBodyDataPrice) *AskSessionPackageRenewPriceResponseBodyData {
	s.Price = v
	return s
}

type AskSessionPackageRenewPriceResponseBodyDataPrice struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s AskSessionPackageRenewPriceResponseBodyDataPrice) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackageRenewPriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *AskSessionPackageRenewPriceResponseBodyDataPrice) SetCurrency(v string) *AskSessionPackageRenewPriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *AskSessionPackageRenewPriceResponseBodyDataPrice) SetDiscountPrice(v float32) *AskSessionPackageRenewPriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *AskSessionPackageRenewPriceResponseBodyDataPrice) SetOriginalPrice(v float32) *AskSessionPackageRenewPriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *AskSessionPackageRenewPriceResponseBodyDataPrice) SetTradePrice(v float32) *AskSessionPackageRenewPriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

type AskSessionPackageRenewPriceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AskSessionPackageRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AskSessionPackageRenewPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackageRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *AskSessionPackageRenewPriceResponse) SetHeaders(v map[string]*string) *AskSessionPackageRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *AskSessionPackageRenewPriceResponse) SetStatusCode(v int32) *AskSessionPackageRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *AskSessionPackageRenewPriceResponse) SetBody(v *AskSessionPackageRenewPriceResponseBody) *AskSessionPackageRenewPriceResponse {
	s.Body = v
	return s
}

type AuthorizeInstanceGroupRequest struct {
	AppInstanceGroupId *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AuthorizeUserIds   []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	ProductType        *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
}

func (s AuthorizeInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupRequest) SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetProductType(v string) *AuthorizeInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.UnAuthorizeUserIds = v
	return s
}

type AuthorizeInstanceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponseBody) SetRequestId(v string) *AuthorizeInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthorizeInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetStatusCode(v int32) *AuthorizeInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetBody(v *AuthorizeInstanceGroupResponseBody) *AuthorizeInstanceGroupResponse {
	s.Body = v
	return s
}

type BuySessionPackageRequest struct {
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	MaxSessions        *int32  `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProjectId          *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SessionPackageName *string `json:"SessionPackageName,omitempty" xml:"SessionPackageName,omitempty"`
	SessionPackageType *string `json:"SessionPackageType,omitempty" xml:"SessionPackageType,omitempty"`
	SessionSpec        *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	SessionType        *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
}

func (s BuySessionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s BuySessionPackageRequest) GoString() string {
	return s.String()
}

func (s *BuySessionPackageRequest) SetChargeType(v string) *BuySessionPackageRequest {
	s.ChargeType = &v
	return s
}

func (s *BuySessionPackageRequest) SetMaxSessions(v int32) *BuySessionPackageRequest {
	s.MaxSessions = &v
	return s
}

func (s *BuySessionPackageRequest) SetPeriod(v int32) *BuySessionPackageRequest {
	s.Period = &v
	return s
}

func (s *BuySessionPackageRequest) SetPeriodUnit(v string) *BuySessionPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *BuySessionPackageRequest) SetProjectId(v string) *BuySessionPackageRequest {
	s.ProjectId = &v
	return s
}

func (s *BuySessionPackageRequest) SetRegion(v string) *BuySessionPackageRequest {
	s.Region = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionPackageName(v string) *BuySessionPackageRequest {
	s.SessionPackageName = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionPackageType(v string) *BuySessionPackageRequest {
	s.SessionPackageType = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionSpec(v string) *BuySessionPackageRequest {
	s.SessionSpec = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionType(v string) *BuySessionPackageRequest {
	s.SessionType = &v
	return s
}

type BuySessionPackageResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionPackageId *int64  `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	Success          *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuySessionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuySessionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *BuySessionPackageResponseBody) SetCode(v string) *BuySessionPackageResponseBody {
	s.Code = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetMessage(v string) *BuySessionPackageResponseBody {
	s.Message = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetRequestId(v string) *BuySessionPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetSessionPackageId(v int64) *BuySessionPackageResponseBody {
	s.SessionPackageId = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetSuccess(v string) *BuySessionPackageResponseBody {
	s.Success = &v
	return s
}

type BuySessionPackageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BuySessionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BuySessionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s BuySessionPackageResponse) GoString() string {
	return s.String()
}

func (s *BuySessionPackageResponse) SetHeaders(v map[string]*string) *BuySessionPackageResponse {
	s.Headers = v
	return s
}

func (s *BuySessionPackageResponse) SetStatusCode(v int32) *BuySessionPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *BuySessionPackageResponse) SetBody(v *BuySessionPackageResponseBody) *BuySessionPackageResponse {
	s.Body = v
	return s
}

type CancelOtaTaskRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelOtaTaskRequest) SetAppInstanceGroupId(v string) *CancelOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CancelOtaTaskRequest) SetTaskId(v string) *CancelOtaTaskRequest {
	s.TaskId = &v
	return s
}

type CancelOtaTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOtaTaskResponseBody) SetCode(v string) *CancelOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOtaTaskResponseBody) SetMessage(v string) *CancelOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOtaTaskResponseBody) SetRequestId(v string) *CancelOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelOtaTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelOtaTaskResponse) SetHeaders(v map[string]*string) *CancelOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelOtaTaskResponse) SetStatusCode(v int32) *CancelOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOtaTaskResponse) SetBody(v *CancelOtaTaskResponseBody) *CancelOtaTaskResponse {
	s.Body = v
	return s
}

type CreateAccessPageRequest struct {
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	CloudEnvId     *string `json:"CloudEnvId,omitempty" xml:"CloudEnvId,omitempty"`
	EffectTime     *int32  `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Unit           *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s CreateAccessPageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPageRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessPageRequest) SetAccessPageName(v string) *CreateAccessPageRequest {
	s.AccessPageName = &v
	return s
}

func (s *CreateAccessPageRequest) SetCloudEnvId(v string) *CreateAccessPageRequest {
	s.CloudEnvId = &v
	return s
}

func (s *CreateAccessPageRequest) SetEffectTime(v int32) *CreateAccessPageRequest {
	s.EffectTime = &v
	return s
}

func (s *CreateAccessPageRequest) SetProjectId(v string) *CreateAccessPageRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAccessPageRequest) SetProjectName(v string) *CreateAccessPageRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateAccessPageRequest) SetUnit(v string) *CreateAccessPageRequest {
	s.Unit = &v
	return s
}

type CreateAccessPageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccessPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessPageResponseBody) SetCode(v string) *CreateAccessPageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetData(v string) *CreateAccessPageResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetMessage(v string) *CreateAccessPageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetRequestId(v string) *CreateAccessPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetSuccess(v string) *CreateAccessPageResponseBody {
	s.Success = &v
	return s
}

type CreateAccessPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessPageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPageResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessPageResponse) SetHeaders(v map[string]*string) *CreateAccessPageResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessPageResponse) SetStatusCode(v int32) *CreateAccessPageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessPageResponse) SetBody(v *CreateAccessPageResponseBody) *CreateAccessPageResponse {
	s.Body = v
	return s
}

type CreateAppInstanceGroupRequest struct {
	AppCenterImageId     *string                                      `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string                                      `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AutoPay              *bool                                        `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool                                        `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BizRegionId          *string                                      `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChargeResourceMode   *string                                      `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	ChargeType           *string                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Network              *CreateAppInstanceGroupRequestNetwork        `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	NodePool             *CreateAppInstanceGroupRequestNodePool       `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	Period               *int32                                       `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string                                      `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PreOpenAppId         *string                                      `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	ProductType          *string                                      `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	PromotionId          *string                                      `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RuntimePolicy        *CreateAppInstanceGroupRequestRuntimePolicy  `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty" type:"Struct"`
	SecurityPolicy       *CreateAppInstanceGroupRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	SessionTimeout       *int32                                       `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicy        *CreateAppInstanceGroupRequestStoragePolicy  `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
	UserInfo             *CreateAppInstanceGroupRequestUserInfo       `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	Users                []*string                                    `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoPay(v bool) *CreateAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetBizRegionId(v string) *CreateAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeType(v string) *CreateAppInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNetwork(v *CreateAppInstanceGroupRequestNetwork) *CreateAppInstanceGroupRequest {
	s.Network = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNodePool(v *CreateAppInstanceGroupRequestNodePool) *CreateAppInstanceGroupRequest {
	s.NodePool = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriod(v int32) *CreateAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPreOpenAppId(v string) *CreateAppInstanceGroupRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetProductType(v string) *CreateAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPromotionId(v string) *CreateAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetRuntimePolicy(v *CreateAppInstanceGroupRequestRuntimePolicy) *CreateAppInstanceGroupRequest {
	s.RuntimePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSecurityPolicy(v *CreateAppInstanceGroupRequestSecurityPolicy) *CreateAppInstanceGroupRequest {
	s.SecurityPolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetStoragePolicy(v *CreateAppInstanceGroupRequestStoragePolicy) *CreateAppInstanceGroupRequest {
	s.StoragePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserInfo(v *CreateAppInstanceGroupRequestUserInfo) *CreateAppInstanceGroupRequest {
	s.UserInfo = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUsers(v []*string) *CreateAppInstanceGroupRequest {
	s.Users = v
	return s
}

type CreateAppInstanceGroupRequestNetwork struct {
	DomainRules     []*CreateAppInstanceGroupRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
	IpExpireMinutes *int32                                             `json:"IpExpireMinutes,omitempty" xml:"IpExpireMinutes,omitempty"`
	Routes          []*CreateAppInstanceGroupRequestNetworkRoutes      `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	StrategyType    *string                                            `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetwork) SetDomainRules(v []*CreateAppInstanceGroupRequestNetworkDomainRules) *CreateAppInstanceGroupRequestNetwork {
	s.DomainRules = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetIpExpireMinutes(v int32) *CreateAppInstanceGroupRequestNetwork {
	s.IpExpireMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetRoutes(v []*CreateAppInstanceGroupRequestNetworkRoutes) *CreateAppInstanceGroupRequestNetwork {
	s.Routes = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetStrategyType(v string) *CreateAppInstanceGroupRequestNetwork {
	s.StrategyType = &v
	return s
}

type CreateAppInstanceGroupRequestNetworkDomainRules struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetworkDomainRules) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetworkDomainRules) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) SetDomain(v string) *CreateAppInstanceGroupRequestNetworkDomainRules {
	s.Domain = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) SetPolicy(v string) *CreateAppInstanceGroupRequestNetworkDomainRules {
	s.Policy = &v
	return s
}

type CreateAppInstanceGroupRequestNetworkRoutes struct {
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetworkRoutes) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetworkRoutes) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) SetDestination(v string) *CreateAppInstanceGroupRequestNetworkRoutes {
	s.Destination = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) SetMode(v string) *CreateAppInstanceGroupRequestNetworkRoutes {
	s.Mode = &v
	return s
}

type CreateAppInstanceGroupRequestNodePool struct {
	MaxScalingAmount            *int32                                                      `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	NodeAmount                  *int32                                                      `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	NodeCapacity                *int32                                                      `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodeInstanceType            *string                                                     `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	RecurrenceSchedules         []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	ScalingDownAfterIdleMinutes *int32                                                      `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingStep                 *int32                                                      `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                     `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	StrategyDisableDate         *string                                                     `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	StrategyEnableDate          *string                                                     `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	StrategyType                *string                                                     `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	WarmUp                      *bool                                                       `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePool) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePool) SetMaxScalingAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeCapacity(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeInstanceType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetRecurrenceSchedules(v []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) *CreateAppInstanceGroupRequestNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingDownAfterIdleMinutes(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingStep(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingStep = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingUsageThreshold(v string) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyDisableDate(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyEnableDate(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetWarmUp(v bool) *CreateAppInstanceGroupRequestNodePool {
	s.WarmUp = &v
	return s
}

type CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules struct {
	RecurrenceType   *string                                                                 `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetTimerPeriods(v []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods struct {
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type CreateAppInstanceGroupRequestRuntimePolicy struct {
	DebugMode *string `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	// 会话类型。
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
}

func (s CreateAppInstanceGroupRequestRuntimePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestRuntimePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetDebugMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.DebugMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionType(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionType = &v
	return s
}

type CreateAppInstanceGroupRequestSecurityPolicy struct {
	ResetAfterUnbind  *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s CreateAppInstanceGroupRequestSecurityPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) SetResetAfterUnbind(v bool) *CreateAppInstanceGroupRequestSecurityPolicy {
	s.ResetAfterUnbind = &v
	return s
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *CreateAppInstanceGroupRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

type CreateAppInstanceGroupRequestStoragePolicy struct {
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequestStoragePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) SetStorageTypeList(v []*string) *CreateAppInstanceGroupRequestStoragePolicy {
	s.StorageTypeList = v
	return s
}

type CreateAppInstanceGroupRequestUserInfo struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppInstanceGroupRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestUserInfo) SetType(v string) *CreateAppInstanceGroupRequestUserInfo {
	s.Type = &v
	return s
}

type CreateAppInstanceGroupShrinkRequest struct {
	AppCenterImageId     *string   `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string   `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AutoPay              *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BizRegionId          *string   `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChargeResourceMode   *string   `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	ChargeType           *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NetworkShrink        *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NodePoolShrink       *string   `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	Period               *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PreOpenAppId         *string   `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	ProductType          *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	PromotionId          *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RuntimePolicyShrink  *string   `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty"`
	SecurityPolicyShrink *string   `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	SessionTimeout       *int32    `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicyShrink  *string   `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
	UserInfoShrink       *string   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	Users                []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoPay(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetBizRegionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNetworkShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNodePoolShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriod(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPreOpenAppId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetProductType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPromotionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetRuntimePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.RuntimePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSecurityPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetStoragePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserInfoShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUsers(v []*string) *CreateAppInstanceGroupShrinkRequest {
	s.Users = v
	return s
}

type CreateAppInstanceGroupResponseBody struct {
	AppInstanceGroupModel *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel `json:"AppInstanceGroupModel,omitempty" xml:"AppInstanceGroupModel,omitempty" type:"Struct"`
	RequestId             *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponseBody) SetAppInstanceGroupModel(v *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) *CreateAppInstanceGroupResponseBody {
	s.AppInstanceGroupModel = v
	return s
}

func (s *CreateAppInstanceGroupResponseBody) SetRequestId(v string) *CreateAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppInstanceGroupResponseBodyAppInstanceGroupModel struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	NodePoolId         *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetAppInstanceGroupId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetNodePoolId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.NodePoolId = &v
	return s
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetOrderId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.OrderId = &v
	return s
}

type CreateAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetStatusCode(v int32) *CreateAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetBody(v *CreateAppInstanceGroupResponseBody) *CreateAppInstanceGroupResponse {
	s.Body = v
	return s
}

type CreateImageFromAppInstanceGroupRequest struct {
	AppCenterImageName *string `json:"AppCenterImageName,omitempty" xml:"AppCenterImageName,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s CreateImageFromAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupRequest) SetAppCenterImageName(v string) *CreateImageFromAppInstanceGroupRequest {
	s.AppCenterImageName = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *CreateImageFromAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) SetProductType(v string) *CreateImageFromAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type CreateImageFromAppInstanceGroupResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageFromAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupResponseBody) SetImageId(v string) *CreateImageFromAppInstanceGroupResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponseBody) SetRequestId(v string) *CreateImageFromAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageFromAppInstanceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageFromAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageFromAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateImageFromAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) SetStatusCode(v int32) *CreateImageFromAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) SetBody(v *CreateImageFromAppInstanceGroupResponseBody) *CreateImageFromAppInstanceGroupResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	Clipboard                    *int32  `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	CloudEnvId                   *string `json:"CloudEnvId,omitempty" xml:"CloudEnvId,omitempty"`
	ContentId                    *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileTransfer                 *int32  `json:"FileTransfer,omitempty" xml:"FileTransfer,omitempty"`
	FrameRate                    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	KeepAliveDuration            *int32  `json:"KeepAliveDuration,omitempty" xml:"KeepAliveDuration,omitempty"`
	ProjectName                  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SessionResolutionHeight      *int32  `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	SessionResolutionWidth       *int32  `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	SessionSpec                  *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	StreamingMode                *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	TerminalResolutionAdaptation *bool   `json:"TerminalResolutionAdaptation,omitempty" xml:"TerminalResolutionAdaptation,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetClipboard(v int32) *CreateProjectRequest {
	s.Clipboard = &v
	return s
}

func (s *CreateProjectRequest) SetCloudEnvId(v string) *CreateProjectRequest {
	s.CloudEnvId = &v
	return s
}

func (s *CreateProjectRequest) SetContentId(v string) *CreateProjectRequest {
	s.ContentId = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetFileTransfer(v int32) *CreateProjectRequest {
	s.FileTransfer = &v
	return s
}

func (s *CreateProjectRequest) SetFrameRate(v int32) *CreateProjectRequest {
	s.FrameRate = &v
	return s
}

func (s *CreateProjectRequest) SetKeepAliveDuration(v int32) *CreateProjectRequest {
	s.KeepAliveDuration = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetSessionResolutionHeight(v int32) *CreateProjectRequest {
	s.SessionResolutionHeight = &v
	return s
}

func (s *CreateProjectRequest) SetSessionResolutionWidth(v int32) *CreateProjectRequest {
	s.SessionResolutionWidth = &v
	return s
}

func (s *CreateProjectRequest) SetSessionSpec(v string) *CreateProjectRequest {
	s.SessionSpec = &v
	return s
}

func (s *CreateProjectRequest) SetStreamingMode(v string) *CreateProjectRequest {
	s.StreamingMode = &v
	return s
}

func (s *CreateProjectRequest) SetTerminalResolutionAdaptation(v bool) *CreateProjectRequest {
	s.TerminalResolutionAdaptation = &v
	return s
}

type CreateProjectResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetCode(v string) *CreateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProjectResponseBody) SetData(v string) *CreateProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProjectResponseBody) SetMessage(v string) *CreateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBody) SetPageNumber(v int32) *CreateProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CreateProjectResponseBody) SetPageSize(v int32) *CreateProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v string) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type DeleteAccessPageRequest struct {
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
}

func (s DeleteAccessPageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPageRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessPageRequest) SetAccessPageId(v string) *DeleteAccessPageRequest {
	s.AccessPageId = &v
	return s
}

type DeleteAccessPageResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccessPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessPageResponseBody) SetCode(v string) *DeleteAccessPageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAccessPageResponseBody) SetMessage(v string) *DeleteAccessPageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccessPageResponseBody) SetRequestId(v string) *DeleteAccessPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessPageResponseBody) SetSuccess(v string) *DeleteAccessPageResponseBody {
	s.Success = &v
	return s
}

type DeleteAccessPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessPageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPageResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPageResponse) SetHeaders(v map[string]*string) *DeleteAccessPageResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPageResponse) SetStatusCode(v int32) *DeleteAccessPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPageResponse) SetBody(v *DeleteAccessPageResponseBody) *DeleteAccessPageResponse {
	s.Body = v
	return s
}

type DeleteAppInstanceGroupRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DeleteAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *DeleteAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DeleteAppInstanceGroupRequest) SetProductType(v string) *DeleteAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type DeleteAppInstanceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupResponseBody) SetRequestId(v string) *DeleteAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupResponse) SetHeaders(v map[string]*string) *DeleteAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstanceGroupResponse) SetStatusCode(v int32) *DeleteAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstanceGroupResponse) SetBody(v *DeleteAppInstanceGroupResponseBody) *DeleteAppInstanceGroupResponse {
	s.Body = v
	return s
}

type DeleteAppInstancesRequest struct {
	AppInstanceGroupId *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceIds     []*string `json:"AppInstanceIds,omitempty" xml:"AppInstanceIds,omitempty" type:"Repeated"`
	ProductType        *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DeleteAppInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesRequest) SetAppInstanceGroupId(v string) *DeleteAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DeleteAppInstancesRequest) SetAppInstanceIds(v []*string) *DeleteAppInstancesRequest {
	s.AppInstanceIds = v
	return s
}

func (s *DeleteAppInstancesRequest) SetProductType(v string) *DeleteAppInstancesRequest {
	s.ProductType = &v
	return s
}

type DeleteAppInstancesResponseBody struct {
	DeleteAppInstanceModels []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels `json:"DeleteAppInstanceModels,omitempty" xml:"DeleteAppInstanceModels,omitempty" type:"Repeated"`
	RequestId               *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponseBody) SetDeleteAppInstanceModels(v []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels) *DeleteAppInstancesResponseBody {
	s.DeleteAppInstanceModels = v
	return s
}

func (s *DeleteAppInstancesResponseBody) SetRequestId(v string) *DeleteAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppInstancesResponseBodyDeleteAppInstanceModels struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppInstancesResponseBodyDeleteAppInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetAppInstanceId(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetCode(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Code = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetMessage(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Message = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetSuccess(v bool) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Success = &v
	return s
}

type DeleteAppInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponse) SetHeaders(v map[string]*string) *DeleteAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstancesResponse) SetStatusCode(v int32) *DeleteAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstancesResponse) SetBody(v *DeleteAppInstancesResponseBody) *DeleteAppInstancesResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetCode(v string) *DeleteProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProjectResponseBody) SetData(v bool) *DeleteProjectResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteProjectResponseBody) SetMessage(v string) *DeleteProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v string) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type GetAccessPageSessionRequest struct {
	AccessPageId    *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	AccessPageToken *string `json:"AccessPageToken,omitempty" xml:"AccessPageToken,omitempty"`
	ExternalUserId  *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
}

func (s GetAccessPageSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionRequest) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionRequest) SetAccessPageId(v string) *GetAccessPageSessionRequest {
	s.AccessPageId = &v
	return s
}

func (s *GetAccessPageSessionRequest) SetAccessPageToken(v string) *GetAccessPageSessionRequest {
	s.AccessPageToken = &v
	return s
}

func (s *GetAccessPageSessionRequest) SetExternalUserId(v string) *GetAccessPageSessionRequest {
	s.ExternalUserId = &v
	return s
}

type GetAccessPageSessionResponseBody struct {
	Code    *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetAccessPageSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccessPageSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionResponseBody) SetCode(v string) *GetAccessPageSessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetData(v *GetAccessPageSessionResponseBodyData) *GetAccessPageSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetMessage(v string) *GetAccessPageSessionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetRequestId(v string) *GetAccessPageSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetSuccess(v string) *GetAccessPageSessionResponseBody {
	s.Success = &v
	return s
}

type GetAccessPageSessionResponseBodyData struct {
	ConnectTicket *string `json:"ConnectTicket,omitempty" xml:"ConnectTicket,omitempty"`
	// flow ID
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetAccessPageSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionResponseBodyData) SetConnectTicket(v string) *GetAccessPageSessionResponseBodyData {
	s.ConnectTicket = &v
	return s
}

func (s *GetAccessPageSessionResponseBodyData) SetFlowId(v string) *GetAccessPageSessionResponseBodyData {
	s.FlowId = &v
	return s
}

type GetAccessPageSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessPageSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessPageSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionResponse) SetHeaders(v map[string]*string) *GetAccessPageSessionResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPageSessionResponse) SetStatusCode(v int32) *GetAccessPageSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPageSessionResponse) SetBody(v *GetAccessPageSessionResponseBody) *GetAccessPageSessionResponse {
	s.Body = v
	return s
}

type GetAppInstanceGroupRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *GetAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAppInstanceGroupRequest) SetProductType(v string) *GetAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type GetAppInstanceGroupResponseBody struct {
	// AppInstanceGroupModels
	AppInstanceGroupModels *GetAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Struct"`
	RequestId              *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) *GetAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *GetAppInstanceGroupResponseBody) SetRequestId(v string) *GetAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	Amount                      *int32                                                           `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCenterImageId            *string                                                          `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppCenterImageName          *string                                                          `json:"AppCenterImageName,omitempty" xml:"AppCenterImageName,omitempty"`
	AppInstanceGroupId          *string                                                          `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName        *string                                                          `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AppInstanceType             *string                                                          `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	AppInstanceTypeName         *string                                                          `json:"AppInstanceTypeName,omitempty" xml:"AppInstanceTypeName,omitempty"`
	AppPolicyId                 *string                                                          `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	Apps                        []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps     `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	ChargeResourceMode          *string                                                          `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	ChargeType                  *string                                                          `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime                 *string                                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate                   *string                                                          `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MaxAmount                   *int32                                                           `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	MinAmount                   *int32                                                           `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	NodePool                    []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	OsType                      *string                                                          `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaInfo                     *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo    `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	ProductType                 *string                                                          `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId                    *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReserveAmountRatio          *string                                                          `json:"ReserveAmountRatio,omitempty" xml:"ReserveAmountRatio,omitempty"`
	ReserveMaxAmount            *int32                                                           `json:"ReserveMaxAmount,omitempty" xml:"ReserveMaxAmount,omitempty"`
	ReserveMinAmount            *int32                                                           `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	ResourceStatus              *string                                                          `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ScalingDownAfterIdleMinutes *int32                                                           `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingStep                 *int32                                                           `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                          `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	SessionTimeout              *string                                                          `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	SkipUserAuthCheck           *bool                                                            `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
	SpecId                      *string                                                          `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	Status                      *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceTypeName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceTypeName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeResourceMode(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeResourceMode = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMaxAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MaxAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMinAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MinAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveAmountRatio(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveAmountRatio = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMaxAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMaxAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMinAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMinAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceStatus(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceStatus = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingDownAfterIdleMinutes(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingStep(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingStep = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingUsageThreshold(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSkipUserAuthCheck(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	AppIcon        *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion     *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	Amount                      *int32                                                                              `json:"Amount,omitempty" xml:"Amount,omitempty"`
	MaxScalingAmount            *int32                                                                              `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	NodeAmount                  *int32                                                                              `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	NodeCapacity                *int32                                                                              `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodeInstanceType            *string                                                                             `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodePoolId                  *string                                                                             `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	NodeTypeName                *string                                                                             `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
	NodeUsed                    *int32                                                                              `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	RecurrenceSchedules         []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	ScalingDownAfterIdleMinutes *int32                                                                              `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingNodeAmount           *int32                                                                              `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	ScalingNodeUsed             *int32                                                                              `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	ScalingStep                 *int32                                                                              `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                                             `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	StrategyDisableDate         *string                                                                             `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	StrategyEnableDate          *string                                                                             `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	StrategyType                *string                                                                             `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	WarmUp                      *bool                                                                               `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeTypeName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeTypeName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	RecurrenceType   *string                                                                                         `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                                        `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
	NewOtaVersion *string `json:"NewOtaVersion,omitempty" xml:"NewOtaVersion,omitempty"`
	OtaVersion    *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

type GetAppInstanceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponse) SetHeaders(v map[string]*string) *GetAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceGroupResponse) SetStatusCode(v int32) *GetAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceGroupResponse) SetBody(v *GetAppInstanceGroupResponseBody) *GetAppInstanceGroupResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	AppId                   *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceGroupIdList  []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	AppInstanceId           *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppInstancePersistentId *string   `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	AppStartParam           *string   `json:"AppStartParam,omitempty" xml:"AppStartParam,omitempty"`
	AppVersion              *string   `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizRegionId             *string   `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	EndUserId               *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ProductType             *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	TaskId                  *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetAppId(v string) *GetConnectionTicketRequest {
	s.AppId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupIdList(v []*string) *GetConnectionTicketRequest {
	s.AppInstanceGroupIdList = v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceId(v string) *GetConnectionTicketRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstancePersistentId(v string) *GetConnectionTicketRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppStartParam(v string) *GetConnectionTicketRequest {
	s.AppStartParam = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppVersion(v string) *GetConnectionTicketRequest {
	s.AppVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetBizRegionId(v string) *GetConnectionTicketRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetProductType(v string) *GetConnectionTicketRequest {
	s.ProductType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	AppInstanceGroupId      *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId           *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	BizRegionId             *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	OsType                  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus              *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TenantId                *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Ticket                  *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetBizRegionId(v string) *GetConnectionTicketResponseBody {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetOsType(v string) *GetConnectionTicketResponseBody {
	s.OsType = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTenantId(v int64) *GetConnectionTicketResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetStatusCode(v int32) *GetConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetDebugAppInstanceRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetDebugAppInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDebugAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceRequest) SetAppInstanceGroupId(v string) *GetDebugAppInstanceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetDebugAppInstanceRequest) SetProductType(v string) *GetDebugAppInstanceRequest {
	s.ProductType = &v
	return s
}

type GetDebugAppInstanceResponseBody struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppVersion         *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AuthCode           *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDebugAppInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDebugAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceResponseBody) SetAppId(v string) *GetDebugAppInstanceResponseBody {
	s.AppId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppInstanceGroupId(v string) *GetDebugAppInstanceResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppInstanceId(v string) *GetDebugAppInstanceResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppVersion(v string) *GetDebugAppInstanceResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAuthCode(v string) *GetDebugAppInstanceResponseBody {
	s.AuthCode = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetRequestId(v string) *GetDebugAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetUserId(v string) *GetDebugAppInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetDebugAppInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDebugAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDebugAppInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDebugAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceResponse) SetHeaders(v map[string]*string) *GetDebugAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetDebugAppInstanceResponse) SetStatusCode(v int32) *GetDebugAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDebugAppInstanceResponse) SetBody(v *GetDebugAppInstanceResponseBody) *GetDebugAppInstanceResponse {
	s.Body = v
	return s
}

type GetOtaTaskByTaskIdRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetOtaTaskByTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdRequest) SetTaskId(v string) *GetOtaTaskByTaskIdRequest {
	s.TaskId = &v
	return s
}

type GetOtaTaskByTaskIdResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OtaVersion    *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	ReleaseNote   *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s GetOtaTaskByTaskIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponseBody) SetCode(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetMessage(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetOtaVersion(v string) *GetOtaTaskByTaskIdResponseBody {
	s.OtaVersion = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetReleaseNote(v string) *GetOtaTaskByTaskIdResponseBody {
	s.ReleaseNote = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetRequestId(v string) *GetOtaTaskByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetTaskStartTime(v string) *GetOtaTaskByTaskIdResponseBody {
	s.TaskStartTime = &v
	return s
}

type GetOtaTaskByTaskIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOtaTaskByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOtaTaskByTaskIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponse) SetHeaders(v map[string]*string) *GetOtaTaskByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetStatusCode(v int32) *GetOtaTaskByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetBody(v *GetOtaTaskByTaskIdResponseBody) *GetOtaTaskByTaskIdResponse {
	s.Body = v
	return s
}

type GetProjectPoliciesRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectPoliciesRequest) GoString() string {
	return s.String()
}

func (s *GetProjectPoliciesRequest) SetProjectId(v string) *GetProjectPoliciesRequest {
	s.ProjectId = &v
	return s
}

type GetProjectPoliciesResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetProjectPoliciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectPoliciesResponseBody) SetCode(v string) *GetProjectPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectPoliciesResponseBody) SetData(v *GetProjectPoliciesResponseBodyData) *GetProjectPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectPoliciesResponseBody) SetMessage(v string) *GetProjectPoliciesResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectPoliciesResponseBody) SetRequestId(v string) *GetProjectPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectPoliciesResponseBody) SetSuccess(v string) *GetProjectPoliciesResponseBody {
	s.Success = &v
	return s
}

type GetProjectPoliciesResponseBodyData struct {
	Clipboard                    *int32  `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	FileTransfer                 *int32  `json:"FileTransfer,omitempty" xml:"FileTransfer,omitempty"`
	FrameRate                    *string `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	KeepAliveDuration            *int32  `json:"KeepAliveDuration,omitempty" xml:"KeepAliveDuration,omitempty"`
	MaxHours                     *int32  `json:"MaxHours,omitempty" xml:"MaxHours,omitempty"`
	MaxSessions                  *int32  `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	ProjectId                    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SessionResolutionHeight      *int32  `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	SessionResolutionWidth       *int32  `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	SessionSpec                  *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	StreamingMode                *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	TerminalResolutionAdaptation *bool   `json:"TerminalResolutionAdaptation,omitempty" xml:"TerminalResolutionAdaptation,omitempty"`
}

func (s GetProjectPoliciesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProjectPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectPoliciesResponseBodyData) SetClipboard(v int32) *GetProjectPoliciesResponseBodyData {
	s.Clipboard = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetFileTransfer(v int32) *GetProjectPoliciesResponseBodyData {
	s.FileTransfer = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetFrameRate(v string) *GetProjectPoliciesResponseBodyData {
	s.FrameRate = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetKeepAliveDuration(v int32) *GetProjectPoliciesResponseBodyData {
	s.KeepAliveDuration = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetMaxHours(v int32) *GetProjectPoliciesResponseBodyData {
	s.MaxHours = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetMaxSessions(v int32) *GetProjectPoliciesResponseBodyData {
	s.MaxSessions = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetProjectId(v string) *GetProjectPoliciesResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetSessionResolutionHeight(v int32) *GetProjectPoliciesResponseBodyData {
	s.SessionResolutionHeight = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetSessionResolutionWidth(v int32) *GetProjectPoliciesResponseBodyData {
	s.SessionResolutionWidth = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetSessionSpec(v string) *GetProjectPoliciesResponseBodyData {
	s.SessionSpec = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetStreamingMode(v string) *GetProjectPoliciesResponseBodyData {
	s.StreamingMode = &v
	return s
}

func (s *GetProjectPoliciesResponseBodyData) SetTerminalResolutionAdaptation(v bool) *GetProjectPoliciesResponseBodyData {
	s.TerminalResolutionAdaptation = &v
	return s
}

type GetProjectPoliciesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectPoliciesResponse) GoString() string {
	return s.String()
}

func (s *GetProjectPoliciesResponse) SetHeaders(v map[string]*string) *GetProjectPoliciesResponse {
	s.Headers = v
	return s
}

func (s *GetProjectPoliciesResponse) SetStatusCode(v int32) *GetProjectPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectPoliciesResponse) SetBody(v *GetProjectPoliciesResponseBody) *GetProjectPoliciesResponse {
	s.Body = v
	return s
}

type GetResourcePriceRequest struct {
	Amount           *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppInstanceType  *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	BizRegionId      *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	Period           *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourcePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePriceRequest) SetAmount(v int64) *GetResourcePriceRequest {
	s.Amount = &v
	return s
}

func (s *GetResourcePriceRequest) SetAppInstanceType(v string) *GetResourcePriceRequest {
	s.AppInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetBizRegionId(v string) *GetResourcePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetResourcePriceRequest) SetChargeType(v string) *GetResourcePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *GetResourcePriceRequest) SetNodeInstanceType(v string) *GetResourcePriceRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriod(v int64) *GetResourcePriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriodUnit(v string) *GetResourcePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourcePriceRequest) SetProductType(v string) *GetResourcePriceRequest {
	s.ProductType = &v
	return s
}

type GetResourcePriceResponseBody struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PriceList  []*GetResourcePriceResponseBodyPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	PriceModel *GetResourcePriceResponseBodyPriceModel  `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourcePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBody) SetCode(v string) *GetResourcePriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetMessage(v string) *GetResourcePriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceList(v []*GetResourcePriceResponseBodyPriceList) *GetResourcePriceResponseBody {
	s.PriceList = v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceModel(v *GetResourcePriceResponseBodyPriceModel) *GetResourcePriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *GetResourcePriceResponseBody) SetRequestId(v string) *GetResourcePriceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourcePriceResponseBodyPriceList struct {
	Price     *GetResourcePriceResponseBodyPriceListPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	PriceType *string                                       `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	Rules     []*GetResourcePriceResponseBodyPriceListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceList) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceList) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceList) SetPrice(v *GetResourcePriceResponseBodyPriceListPrice) *GetResourcePriceResponseBodyPriceList {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) SetPriceType(v string) *GetResourcePriceResponseBodyPriceList {
	s.PriceType = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) SetRules(v []*GetResourcePriceResponseBodyPriceListRules) *GetResourcePriceResponseBodyPriceList {
	s.Rules = v
	return s
}

type GetResourcePriceResponseBodyPriceListPrice struct {
	Currency      *string                                                 `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *string                                                 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *string                                                 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*GetResourcePriceResponseBodyPriceListPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *string                                                 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceListPricePromotions) *GetResourcePriceResponseBodyPriceListPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.TradePrice = &v
	return s
}

type GetResourcePriceResponseBodyPriceListPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.Selected = &v
	return s
}

type GetResourcePriceResponseBodyPriceListRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceListRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceListRules {
	s.RuleId = &v
	return s
}

type GetResourcePriceResponseBodyPriceModel struct {
	Price *GetResourcePriceResponseBodyPriceModelPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*GetResourcePriceResponseBodyPriceModelRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceModel) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModel) SetPrice(v *GetResourcePriceResponseBodyPriceModelPrice) *GetResourcePriceResponseBodyPriceModel {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModel) SetRules(v []*GetResourcePriceResponseBodyPriceModelRules) *GetResourcePriceResponseBodyPriceModel {
	s.Rules = v
	return s
}

type GetResourcePriceResponseBodyPriceModelPrice struct {
	Currency      *string                                                  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *string                                                  `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *string                                                  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*GetResourcePriceResponseBodyPriceModelPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *string                                                  `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceModelPricePromotions) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.TradePrice = &v
	return s
}

type GetResourcePriceResponseBodyPriceModelPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.Selected = &v
	return s
}

type GetResourcePriceResponseBodyPriceModelRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceModelRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceModelRules {
	s.RuleId = &v
	return s
}

type GetResourcePriceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourcePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourcePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponse) SetHeaders(v map[string]*string) *GetResourcePriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePriceResponse) SetStatusCode(v int32) *GetResourcePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePriceResponse) SetBody(v *GetResourcePriceResponseBody) *GetResourcePriceResponse {
	s.Body = v
	return s
}

type GetResourceRenewPriceRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	Period             *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourceRenewPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceRequest) SetAppInstanceGroupId(v string) *GetResourceRenewPriceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetPeriod(v int64) *GetResourceRenewPriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetPeriodUnit(v string) *GetResourceRenewPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetProductType(v string) *GetResourceRenewPriceRequest {
	s.ProductType = &v
	return s
}

type GetResourceRenewPriceResponseBody struct {
	Data      *GetResourceRenewPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceRenewPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBody) SetData(v *GetResourceRenewPriceResponseBodyData) *GetResourceRenewPriceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceRenewPriceResponseBody) SetRequestId(v string) *GetResourceRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceRenewPriceResponseBodyData struct {
	Price *GetResourceRenewPriceResponseBodyDataPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*GetResourceRenewPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourceRenewPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyData) SetPrice(v *GetResourceRenewPriceResponseBodyDataPrice) *GetResourceRenewPriceResponseBodyData {
	s.Price = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyData) SetRules(v []*GetResourceRenewPriceResponseBodyDataRules) *GetResourceRenewPriceResponseBodyData {
	s.Rules = v
	return s
}

type GetResourceRenewPriceResponseBodyDataPrice struct {
	Currency      *string                                                 `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *string                                                 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *string                                                 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*GetResourceRenewPriceResponseBodyDataPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *string                                                 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetCurrency(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetDiscountPrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetOriginalPrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetPromotions(v []*GetResourceRenewPriceResponseBodyDataPricePromotions) *GetResourceRenewPriceResponseBodyDataPrice {
	s.Promotions = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetTradePrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

type GetResourceRenewPriceResponseBodyDataPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetOptionCode(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionDesc(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionId(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionName(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetSelected(v bool) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.Selected = &v
	return s
}

type GetResourceRenewPriceResponseBodyDataRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataRules) SetDescription(v string) *GetResourceRenewPriceResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataRules) SetRuleId(v int64) *GetResourceRenewPriceResponseBodyDataRules {
	s.RuleId = &v
	return s
}

type GetResourceRenewPriceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceRenewPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponse) SetHeaders(v map[string]*string) *GetResourceRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceRenewPriceResponse) SetStatusCode(v int32) *GetResourceRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceRenewPriceResponse) SetBody(v *GetResourceRenewPriceResponseBody) *GetResourceRenewPriceResponse {
	s.Body = v
	return s
}

type ListAccessPagesRequest struct {
	AccessPageId   *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SortType       *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListAccessPagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPagesRequest) SetAccessPageId(v string) *ListAccessPagesRequest {
	s.AccessPageId = &v
	return s
}

func (s *ListAccessPagesRequest) SetAccessPageName(v string) *ListAccessPagesRequest {
	s.AccessPageName = &v
	return s
}

func (s *ListAccessPagesRequest) SetPageNumber(v int32) *ListAccessPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccessPagesRequest) SetPageSize(v int32) *ListAccessPagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccessPagesRequest) SetProjectId(v string) *ListAccessPagesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListAccessPagesRequest) SetSortType(v string) *ListAccessPagesRequest {
	s.SortType = &v
	return s
}

type ListAccessPagesResponseBody struct {
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string                            `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       []*ListAccessPagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *string                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccessPagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPagesResponseBody) SetCode(v string) *ListAccessPagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetCount(v string) *ListAccessPagesResponseBody {
	s.Count = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetData(v []*ListAccessPagesResponseBodyData) *ListAccessPagesResponseBody {
	s.Data = v
	return s
}

func (s *ListAccessPagesResponseBody) SetMessage(v string) *ListAccessPagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetPageNumber(v string) *ListAccessPagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetPageSize(v string) *ListAccessPagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetRequestId(v string) *ListAccessPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetSuccess(v string) *ListAccessPagesResponseBody {
	s.Success = &v
	return s
}

type ListAccessPagesResponseBodyData struct {
	AccessMode      *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	AccessPageId    *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	AccessPageName  *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	AccessPageState *string `json:"AccessPageState,omitempty" xml:"AccessPageState,omitempty"`
	AccessUrl       *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
	ContentId       *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ContentName     *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	EffectTime      *int32  `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Unit            *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	UrlExpireTime   *string `json:"UrlExpireTime,omitempty" xml:"UrlExpireTime,omitempty"`
}

func (s ListAccessPagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccessPagesResponseBodyData) SetAccessMode(v string) *ListAccessPagesResponseBodyData {
	s.AccessMode = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessPageId(v string) *ListAccessPagesResponseBodyData {
	s.AccessPageId = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessPageName(v string) *ListAccessPagesResponseBodyData {
	s.AccessPageName = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessPageState(v string) *ListAccessPagesResponseBodyData {
	s.AccessPageState = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessUrl(v string) *ListAccessPagesResponseBodyData {
	s.AccessUrl = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetContentId(v string) *ListAccessPagesResponseBodyData {
	s.ContentId = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetContentName(v string) *ListAccessPagesResponseBodyData {
	s.ContentName = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetEffectTime(v int32) *ListAccessPagesResponseBodyData {
	s.EffectTime = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetGmtCreate(v string) *ListAccessPagesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetProjectId(v string) *ListAccessPagesResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetProjectName(v string) *ListAccessPagesResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetUnit(v string) *ListAccessPagesResponseBodyData {
	s.Unit = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetUrlExpireTime(v string) *ListAccessPagesResponseBodyData {
	s.UrlExpireTime = &v
	return s
}

type ListAccessPagesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccessPagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccessPagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesResponse) GoString() string {
	return s.String()
}

func (s *ListAccessPagesResponse) SetHeaders(v map[string]*string) *ListAccessPagesResponse {
	s.Headers = v
	return s
}

func (s *ListAccessPagesResponse) SetStatusCode(v int32) *ListAccessPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessPagesResponse) SetBody(v *ListAccessPagesResponseBody) *ListAccessPagesResponse {
	s.Body = v
	return s
}

type ListAppInstanceGroupRequest struct {
	AppCenterImageId     *string   `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupId   *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string   `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	BizRegionId          *string   `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NodeInstanceType     *string   `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType          *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Status               []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupRequest) SetAppCenterImageId(v string) *ListAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetBizRegionId(v string) *ListAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetNodeInstanceType(v string) *ListAppInstanceGroupRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageNumber(v int32) *ListAppInstanceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageSize(v int32) *ListAppInstanceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetProductType(v string) *ListAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetStatus(v []*string) *ListAppInstanceGroupRequest {
	s.Status = v
	return s
}

type ListAppInstanceGroupResponseBody struct {
	AppInstanceGroupModels []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	PageNumber             *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageNumber(v int32) *ListAppInstanceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageSize(v int32) *ListAppInstanceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetRequestId(v string) *ListAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetTotalCount(v int32) *ListAppInstanceGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	Amount               *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCenterImageId     *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupId   *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AppInstanceType      *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// 策略ID。
	AppPolicyId *string                                                       `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	Apps        []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// 售卖模式。
	ChargeResourceMode          *string                                                           `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	ChargeType                  *string                                                           `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime                 *string                                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate                   *string                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MaxAmount                   *int32                                                            `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	MinAmount                   *int32                                                            `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	NodePool                    []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	OsType                      *string                                                           `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaInfo                     *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo    `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	ProductType                 *string                                                           `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId                    *string                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReserveAmountRatio          *string                                                           `json:"ReserveAmountRatio,omitempty" xml:"ReserveAmountRatio,omitempty"`
	ReserveMaxAmount            *int32                                                            `json:"ReserveMaxAmount,omitempty" xml:"ReserveMaxAmount,omitempty"`
	ReserveMinAmount            *int32                                                            `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	ResourceStatus              *string                                                           `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ScalingDownAfterIdleMinutes *int32                                                            `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingStep                 *int32                                                            `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                           `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	SessionTimeout              *string                                                           `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	SkipUserAuthCheck           *bool                                                             `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
	SpecId                      *string                                                           `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	Status                      *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeResourceMode(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeResourceMode = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMaxAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MaxAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMinAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MinAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveAmountRatio(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveAmountRatio = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMaxAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMaxAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMinAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMinAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceStatus = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSkipUserAuthCheck(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	// 应用图标。
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本。
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 应用版本名称。
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	Amount                      *int32                                                                               `json:"Amount,omitempty" xml:"Amount,omitempty"`
	MaxScalingAmount            *int32                                                                               `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	NodeAmount                  *int32                                                                               `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	NodeCapacity                *int32                                                                               `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodeInstanceType            *string                                                                              `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodePoolId                  *string                                                                              `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	NodeTypeName                *string                                                                              `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
	NodeUsed                    *int32                                                                               `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	RecurrenceSchedules         []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	ScalingDownAfterIdleMinutes *int32                                                                               `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingNodeAmount           *int32                                                                               `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	ScalingNodeUsed             *int32                                                                               `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	ScalingStep                 *int32                                                                               `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                                              `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	StrategyDisableDate         *string                                                                              `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	StrategyEnableDate          *string                                                                              `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	StrategyType                *string                                                                              `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	WarmUp                      *bool                                                                                `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeTypeName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeTypeName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	RecurrenceType   *string                                                                                          `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                                         `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
	NewOtaVersion *string `json:"NewOtaVersion,omitempty" xml:"NewOtaVersion,omitempty"`
	OtaVersion    *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

type ListAppInstanceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponse) SetHeaders(v map[string]*string) *ListAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceGroupResponse) SetStatusCode(v int32) *ListAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceGroupResponse) SetBody(v *ListAppInstanceGroupResponseBody) *ListAppInstanceGroupResponse {
	s.Body = v
	return s
}

type ListAppInstancesRequest struct {
	AppInstanceGroupId *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppInstanceIdList  []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	IncludeDeleted     *bool     `json:"IncludeDeleted,omitempty" xml:"IncludeDeleted,omitempty"`
	PageNumber         *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status             []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListAppInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstancesRequest) SetAppInstanceGroupId(v string) *ListAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstancesRequest) SetAppInstanceId(v string) *ListAppInstancesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstancesRequest) SetAppInstanceIdList(v []*string) *ListAppInstancesRequest {
	s.AppInstanceIdList = v
	return s
}

func (s *ListAppInstancesRequest) SetIncludeDeleted(v bool) *ListAppInstancesRequest {
	s.IncludeDeleted = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageNumber(v int32) *ListAppInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageSize(v int32) *ListAppInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesRequest) SetStatus(v []*string) *ListAppInstancesRequest {
	s.Status = v
	return s
}

type ListAppInstancesResponseBody struct {
	AppInstanceModels []*ListAppInstancesResponseBodyAppInstanceModels `json:"AppInstanceModels,omitempty" xml:"AppInstanceModels,omitempty" type:"Repeated"`
	PageNumber        *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBody) SetAppInstanceModels(v []*ListAppInstancesResponseBodyAppInstanceModels) *ListAppInstancesResponseBody {
	s.AppInstanceModels = v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageNumber(v int32) *ListAppInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageSize(v int32) *ListAppInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetRequestId(v string) *ListAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetTotalCount(v int32) *ListAppInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppInstancesResponseBodyAppInstanceModels struct {
	AppInstanceGroupId *string                                                `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string                                                `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	BindInfo           *ListAppInstancesResponseBodyAppInstanceModelsBindInfo `json:"BindInfo,omitempty" xml:"BindInfo,omitempty" type:"Struct"`
	GmtCreate          *string                                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MainEthPublicIp    *string                                                `json:"MainEthPublicIp,omitempty" xml:"MainEthPublicIp,omitempty"`
	SessionStatus      *string                                                `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	Status             *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInstancesResponseBodyAppInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponseBodyAppInstanceModels) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetAppInstanceGroupId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetAppInstanceId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetBindInfo(v *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) *ListAppInstancesResponseBodyAppInstanceModels {
	s.BindInfo = v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetGmtCreate(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetGmtModified(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.GmtModified = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetMainEthPublicIp(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.MainEthPublicIp = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetSessionStatus(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.SessionStatus = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetStatus(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.Status = &v
	return s
}

type ListAppInstancesResponseBodyAppInstanceModelsBindInfo struct {
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	UsageDuration *int64  `json:"UsageDuration,omitempty" xml:"UsageDuration,omitempty"`
}

func (s ListAppInstancesResponseBodyAppInstanceModelsBindInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponseBodyAppInstanceModelsBindInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) SetEndUserId(v string) *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	s.EndUserId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) SetUsageDuration(v int64) *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	s.UsageDuration = &v
	return s
}

type ListAppInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponse) SetHeaders(v map[string]*string) *ListAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstancesResponse) SetStatusCode(v int32) *ListAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstancesResponse) SetBody(v *ListAppInstancesResponseBody) *ListAppInstancesResponse {
	s.Body = v
	return s
}

type ListNodeInstanceTypeRequest struct {
	// 资源所属的地域ID。关于支持的地域详情，请参见[使用限制](~~426036~~)。
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// 语言类型。
	Language         *string `json:"Language,omitempty" xml:"Language,omitempty"`
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// 支持的操作系统类型。
	OsType      *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListNodeInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeRequest) SetBizRegionId(v string) *ListNodeInstanceTypeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetLanguage(v string) *ListNodeInstanceTypeRequest {
	s.Language = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetNodeInstanceType(v string) *ListNodeInstanceTypeRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOsType(v string) *ListNodeInstanceTypeRequest {
	s.OsType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageNumber(v int32) *ListNodeInstanceTypeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageSize(v int32) *ListNodeInstanceTypeRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetProductType(v string) *ListNodeInstanceTypeRequest {
	s.ProductType = &v
	return s
}

type ListNodeInstanceTypeResponseBody struct {
	NodeInstanceTypeModels []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels `json:"NodeInstanceTypeModels,omitempty" xml:"NodeInstanceTypeModels,omitempty" type:"Repeated"`
	PageNumber             *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBody) SetNodeInstanceTypeModels(v []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) *ListNodeInstanceTypeResponseBody {
	s.NodeInstanceTypeModels = v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageNumber(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageSize(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetRequestId(v string) *ListNodeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetTotalCount(v int32) *ListNodeInstanceTypeResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels struct {
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// 显卡内存大小，单位为MB。
	GpuMemory *int64 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// 最大并发会话数，即单个资源可同时连接的会话数。如果同时连接的会话数过多，可能导致应用的使用体验下降。取值范围因资源规格不同而不同。各资源规格对应的取值范围分别是：
	//
	// - appstreaming.general.4c8g：1\~2；
	// - appstreaming.general.8c16g：1\~4；
	// - appstreaming.vgpu.8c16g.4g：1\~4；
	// - appstreaming.vgpu.8c31g.16g：1\~4；
	// - appstreaming.vgpu.14c93g.12g：1\~6；
	MaxCapacity            *int32  `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	Memory                 *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeInstanceType       *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	// 资源规格名称。
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetCpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Cpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Gpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpuMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.GpuMemory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMaxCapacity(v int32) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.MaxCapacity = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Memory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceType(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceTypeFamily = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeTypeName(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeTypeName = &v
	return s
}

type ListNodeInstanceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponse) SetHeaders(v map[string]*string) *ListNodeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetStatusCode(v int32) *ListNodeInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetBody(v *ListNodeInstanceTypeResponseBody) *ListNodeInstanceTypeResponse {
	s.Body = v
	return s
}

type ListOtaTaskRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	OtaType            *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOtaTaskRequest) SetAppInstanceGroupId(v string) *ListOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListOtaTaskRequest) SetOtaType(v string) *ListOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageNumber(v int32) *ListOtaTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageSize(v int32) *ListOtaTaskRequest {
	s.PageSize = &v
	return s
}

type ListOtaTaskResponseBody struct {
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList   []*ListOtaTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBody) SetPageNumber(v int32) *ListOtaTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetPageSize(v int32) *ListOtaTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetRequestId(v string) *ListOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetTaskList(v []*ListOtaTaskResponseBodyTaskList) *ListOtaTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *ListOtaTaskResponseBody) SetTotalCount(v int32) *ListOtaTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListOtaTaskResponseBodyTaskList struct {
	OtaVersion        *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	TaskDisplayStatus *string `json:"TaskDisplayStatus,omitempty" xml:"TaskDisplayStatus,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStartTime     *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s ListOtaTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBodyTaskList) SetOtaVersion(v string) *ListOtaTaskResponseBodyTaskList {
	s.OtaVersion = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskDisplayStatus(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskDisplayStatus = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskId(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskStartTime(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskStartTime = &v
	return s
}

type ListOtaTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponse) SetHeaders(v map[string]*string) *ListOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOtaTaskResponse) SetStatusCode(v int32) *ListOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOtaTaskResponse) SetBody(v *ListOtaTaskResponseBody) *ListOtaTaskResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	PageNumber  *int32   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId   *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SortType    *string  `json:"SortType,omitempty" xml:"SortType,omitempty"`
	StateList   []*int32 `json:"StateList,omitempty" xml:"StateList,omitempty" type:"Repeated"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetProjectId(v string) *ListProjectsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsRequest) SetProjectName(v string) *ListProjectsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListProjectsRequest) SetSortType(v string) *ListProjectsRequest {
	s.SortType = &v
	return s
}

func (s *ListProjectsRequest) SetStateList(v []*int32) *ListProjectsRequest {
	s.StateList = v
	return s
}

type ListProjectsResponseBody struct {
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListProjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetCode(v string) *ListProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectsResponseBody) SetData(v []*ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetMessage(v string) *ListProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageNumber(v int32) *ListProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageSize(v int32) *ListProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetSuccess(v string) *ListProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int32) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyData struct {
	AccessPageId   []*int64 `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty" type:"Repeated"`
	AvailableHours *int32   `json:"AvailableHours,omitempty" xml:"AvailableHours,omitempty"`
	ContentId      *string  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ContentName    *string  `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	CreateTime     *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	InUseSessions  *int64   `json:"InUseSessions,omitempty" xml:"InUseSessions,omitempty"`
	MaxHours       *int64   `json:"MaxHours,omitempty" xml:"MaxHours,omitempty"`
	MaxSessions    *int64   `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	ProjectId      *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName    *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectState   *string  `json:"ProjectState,omitempty" xml:"ProjectState,omitempty"`
	SessionSpec    *string  `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
}

func (s ListProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) SetAccessPageId(v []*int64) *ListProjectsResponseBodyData {
	s.AccessPageId = v
	return s
}

func (s *ListProjectsResponseBodyData) SetAvailableHours(v int32) *ListProjectsResponseBodyData {
	s.AvailableHours = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetContentId(v string) *ListProjectsResponseBodyData {
	s.ContentId = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetContentName(v string) *ListProjectsResponseBodyData {
	s.ContentName = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetCreateTime(v string) *ListProjectsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetDescription(v string) *ListProjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetInUseSessions(v int64) *ListProjectsResponseBodyData {
	s.InUseSessions = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMaxHours(v int64) *ListProjectsResponseBodyData {
	s.MaxHours = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMaxSessions(v int64) *ListProjectsResponseBodyData {
	s.MaxSessions = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectId(v string) *ListProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectName(v string) *ListProjectsResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectState(v string) *ListProjectsResponseBodyData {
	s.ProjectState = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetSessionSpec(v string) *ListProjectsResponseBodyData {
	s.SessionSpec = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	RegionModels []*ListRegionsResponseBodyRegionModels `json:"RegionModels,omitempty" xml:"RegionModels,omitempty" type:"Repeated"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegionModels(v []*ListRegionsResponseBodyRegionModels) *ListRegionsResponseBody {
	s.RegionModels = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegionModels struct {
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionModels) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionModels) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionModels) SetRegionId(v string) *ListRegionsResponseBodyRegionModels {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSessionPackagesRequest struct {
	PageNumber         *int32   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId          *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SessionPackageId   *string  `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	SessionPackageName *string  `json:"SessionPackageName,omitempty" xml:"SessionPackageName,omitempty"`
	SortType           *string  `json:"SortType,omitempty" xml:"SortType,omitempty"`
	StateList          []*int32 `json:"StateList,omitempty" xml:"StateList,omitempty" type:"Repeated"`
}

func (s ListSessionPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesRequest) SetPageNumber(v int32) *ListSessionPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSessionPackagesRequest) SetPageSize(v int32) *ListSessionPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSessionPackagesRequest) SetProjectId(v string) *ListSessionPackagesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSessionPackagesRequest) SetSessionPackageId(v string) *ListSessionPackagesRequest {
	s.SessionPackageId = &v
	return s
}

func (s *ListSessionPackagesRequest) SetSessionPackageName(v string) *ListSessionPackagesRequest {
	s.SessionPackageName = &v
	return s
}

func (s *ListSessionPackagesRequest) SetSortType(v string) *ListSessionPackagesRequest {
	s.SortType = &v
	return s
}

func (s *ListSessionPackagesRequest) SetStateList(v []*int32) *ListSessionPackagesRequest {
	s.StateList = v
	return s
}

type ListSessionPackagesResponseBody struct {
	Data       []*ListSessionPackagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSessionPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponseBody) SetData(v []*ListSessionPackagesResponseBodyData) *ListSessionPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionPackagesResponseBody) SetPageNumber(v int32) *ListSessionPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSessionPackagesResponseBody) SetPageSize(v int32) *ListSessionPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSessionPackagesResponseBody) SetRequestId(v string) *ListSessionPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionPackagesResponseBody) SetTotalCount(v int64) *ListSessionPackagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListSessionPackagesResponseBodyData struct {
	AvailableHours       *int64                                             `json:"AvailableHours,omitempty" xml:"AvailableHours,omitempty"`
	ChargeType           *string                                            `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DeleteStatus         *int32                                             `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	GmtCreate            *string                                            `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModifiedTime      *string                                            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceObject       *ListSessionPackagesResponseBodyDataInstanceObject `json:"InstanceObject,omitempty" xml:"InstanceObject,omitempty" type:"Struct"`
	MaxHours             *int64                                             `json:"MaxHours,omitempty" xml:"MaxHours,omitempty"`
	MaxSessions          *int64                                             `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	ProjectId            *string                                            `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName          *string                                            `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Region               *string                                            `json:"Region,omitempty" xml:"Region,omitempty"`
	SessionPackageId     *string                                            `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	SessionPackageName   *string                                            `json:"SessionPackageName,omitempty" xml:"SessionPackageName,omitempty"`
	SessionPackageTypeId *string                                            `json:"SessionPackageTypeId,omitempty" xml:"SessionPackageTypeId,omitempty"`
	SessionSpec          *string                                            `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	SessionUsageRate     *int64                                             `json:"SessionUsageRate,omitempty" xml:"SessionUsageRate,omitempty"`
	State                *int32                                             `json:"State,omitempty" xml:"State,omitempty"`
	UserIdentification   *int64                                             `json:"UserIdentification,omitempty" xml:"UserIdentification,omitempty"`
}

func (s ListSessionPackagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponseBodyData) SetAvailableHours(v int64) *ListSessionPackagesResponseBodyData {
	s.AvailableHours = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetChargeType(v string) *ListSessionPackagesResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetDeleteStatus(v int32) *ListSessionPackagesResponseBodyData {
	s.DeleteStatus = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetGmtCreate(v string) *ListSessionPackagesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetGmtModifiedTime(v string) *ListSessionPackagesResponseBodyData {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetInstanceObject(v *ListSessionPackagesResponseBodyDataInstanceObject) *ListSessionPackagesResponseBodyData {
	s.InstanceObject = v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetMaxHours(v int64) *ListSessionPackagesResponseBodyData {
	s.MaxHours = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetMaxSessions(v int64) *ListSessionPackagesResponseBodyData {
	s.MaxSessions = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetProjectId(v string) *ListSessionPackagesResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetProjectName(v string) *ListSessionPackagesResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetRegion(v string) *ListSessionPackagesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionPackageId(v string) *ListSessionPackagesResponseBodyData {
	s.SessionPackageId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionPackageName(v string) *ListSessionPackagesResponseBodyData {
	s.SessionPackageName = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionPackageTypeId(v string) *ListSessionPackagesResponseBodyData {
	s.SessionPackageTypeId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionSpec(v string) *ListSessionPackagesResponseBodyData {
	s.SessionSpec = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionUsageRate(v int64) *ListSessionPackagesResponseBodyData {
	s.SessionUsageRate = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetState(v int32) *ListSessionPackagesResponseBodyData {
	s.State = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetUserIdentification(v int64) *ListSessionPackagesResponseBodyData {
	s.UserIdentification = &v
	return s
}

type ListSessionPackagesResponseBodyDataInstanceObject struct {
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalTime    *int64  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	UsedTime     *int64  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s ListSessionPackagesResponseBodyDataInstanceObject) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponseBodyDataInstanceObject) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetExpiredTime(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.ExpiredTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetInstanceId(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.InstanceId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetInstanceType(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.InstanceType = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetResourceId(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.ResourceId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetResourceType(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.ResourceType = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetStartTime(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.StartTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetTotalTime(v int64) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.TotalTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetUsedTime(v int64) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.UsedTime = &v
	return s
}

type ListSessionPackagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSessionPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSessionPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponse) SetHeaders(v map[string]*string) *ListSessionPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListSessionPackagesResponse) SetStatusCode(v int32) *ListSessionPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionPackagesResponse) SetBody(v *ListSessionPackagesResponseBody) *ListSessionPackagesResponse {
	s.Body = v
	return s
}

type ListTenantConfigResponseBody struct {
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantConfigModel *ListTenantConfigResponseBodyTenantConfigModel `json:"TenantConfigModel,omitempty" xml:"TenantConfigModel,omitempty" type:"Struct"`
}

func (s ListTenantConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTenantConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponseBody) SetRequestId(v string) *ListTenantConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantConfigResponseBody) SetTenantConfigModel(v *ListTenantConfigResponseBodyTenantConfigModel) *ListTenantConfigResponseBody {
	s.TenantConfigModel = v
	return s
}

type ListTenantConfigResponseBodyTenantConfigModel struct {
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
}

func (s ListTenantConfigResponseBodyTenantConfigModel) String() string {
	return tea.Prettify(s)
}

func (s ListTenantConfigResponseBodyTenantConfigModel) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) SetAppInstanceGroupExpireRemind(v bool) *ListTenantConfigResponseBodyTenantConfigModel {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

type ListTenantConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTenantConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTenantConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTenantConfigResponse) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponse) SetHeaders(v map[string]*string) *ListTenantConfigResponse {
	s.Headers = v
	return s
}

func (s *ListTenantConfigResponse) SetStatusCode(v int32) *ListTenantConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantConfigResponse) SetBody(v *ListTenantConfigResponseBody) *ListTenantConfigResponse {
	s.Body = v
	return s
}

type LogOffAllSessionsInAppInstanceGroupRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *LogOffAllSessionsInAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) SetProductType(v string) *LogOffAllSessionsInAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type LogOffAllSessionsInAppInstanceGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetCode(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetMessage(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetRequestId(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type LogOffAllSessionsInAppInstanceGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogOffAllSessionsInAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LogOffAllSessionsInAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetHeaders(v map[string]*string) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetStatusCode(v int32) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetBody(v *LogOffAllSessionsInAppInstanceGroupResponseBody) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.Body = v
	return s
}

type MigrateSessionPackageRequest struct {
	DestProjectId    *string `json:"DestProjectId,omitempty" xml:"DestProjectId,omitempty"`
	SessionPackageId *string `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	SourceProjectId  *string `json:"SourceProjectId,omitempty" xml:"SourceProjectId,omitempty"`
}

func (s MigrateSessionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateSessionPackageRequest) GoString() string {
	return s.String()
}

func (s *MigrateSessionPackageRequest) SetDestProjectId(v string) *MigrateSessionPackageRequest {
	s.DestProjectId = &v
	return s
}

func (s *MigrateSessionPackageRequest) SetSessionPackageId(v string) *MigrateSessionPackageRequest {
	s.SessionPackageId = &v
	return s
}

func (s *MigrateSessionPackageRequest) SetSourceProjectId(v string) *MigrateSessionPackageRequest {
	s.SourceProjectId = &v
	return s
}

type MigrateSessionPackageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MigrateSessionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateSessionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateSessionPackageResponseBody) SetCode(v string) *MigrateSessionPackageResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateSessionPackageResponseBody) SetMessage(v string) *MigrateSessionPackageResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateSessionPackageResponseBody) SetRequestId(v string) *MigrateSessionPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateSessionPackageResponseBody) SetSuccess(v bool) *MigrateSessionPackageResponseBody {
	s.Success = &v
	return s
}

type MigrateSessionPackageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MigrateSessionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateSessionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateSessionPackageResponse) GoString() string {
	return s.String()
}

func (s *MigrateSessionPackageResponse) SetHeaders(v map[string]*string) *MigrateSessionPackageResponse {
	s.Headers = v
	return s
}

func (s *MigrateSessionPackageResponse) SetStatusCode(v int32) *MigrateSessionPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateSessionPackageResponse) SetBody(v *MigrateSessionPackageResponseBody) *MigrateSessionPackageResponse {
	s.Body = v
	return s
}

type ModifyAppInstanceGroupAttributeRequest struct {
	AppInstanceGroupId   *string                                               `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string                                               `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	Network              *ModifyAppInstanceGroupAttributeRequestNetwork        `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	NodePool             *ModifyAppInstanceGroupAttributeRequestNodePool       `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	PreOpenAppId         *string                                               `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	PreOpenMode          *string                                               `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	ProductType          *string                                               `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SecurityPolicy       *ModifyAppInstanceGroupAttributeRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	SessionTimeout       *int32                                                `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicy        *ModifyAppInstanceGroupAttributeRequestStoragePolicy  `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
}

func (s ModifyAppInstanceGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNetwork(v *ModifyAppInstanceGroupAttributeRequestNetwork) *ModifyAppInstanceGroupAttributeRequest {
	s.Network = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNodePool(v *ModifyAppInstanceGroupAttributeRequestNodePool) *ModifyAppInstanceGroupAttributeRequest {
	s.NodePool = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.PreOpenMode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSecurityPolicy(v *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) *ModifyAppInstanceGroupAttributeRequest {
	s.SecurityPolicy = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetStoragePolicy(v *ModifyAppInstanceGroupAttributeRequestStoragePolicy) *ModifyAppInstanceGroupAttributeRequest {
	s.StoragePolicy = v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNetwork struct {
	DomainRules []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
}

func (s ModifyAppInstanceGroupAttributeRequestNetwork) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNetwork) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNetwork) SetDomainRules(v []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) *ModifyAppInstanceGroupAttributeRequestNetwork {
	s.DomainRules = v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNetworkDomainRules struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) SetDomain(v string) *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	s.Domain = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) SetPolicy(v string) *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	s.Policy = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNodePool struct {
	NodeCapacity *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolId   *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodeCapacity(v int32) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodePoolId(v string) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodePoolId = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestSecurityPolicy struct {
	ResetAfterUnbind  *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestSecurityPolicy) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) SetResetAfterUnbind(v bool) *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	s.ResetAfterUnbind = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicy struct {
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicy) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetStorageTypeList(v []*string) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.StorageTypeList = v
	return s
}

type ModifyAppInstanceGroupAttributeShrinkRequest struct {
	AppInstanceGroupId   *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	NetworkShrink        *string `json:"Network,omitempty" xml:"Network,omitempty"`
	NodePoolShrink       *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	PreOpenAppId         *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	PreOpenMode          *string `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicyShrink  *string `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNetworkShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNodePoolShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PreOpenMode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSecurityPolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetStoragePolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

type ModifyAppInstanceGroupAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetCode(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetMessage(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetRequestId(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppInstanceGroupAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppInstanceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppInstanceGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetStatusCode(v int32) *ModifyAppInstanceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetBody(v *ModifyAppInstanceGroupAttributeResponseBody) *ModifyAppInstanceGroupAttributeResponse {
	s.Body = v
	return s
}

type ModifyNodePoolAttributeRequest struct {
	BizRegionId      *string                                         `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NodeCapacity     *int32                                          `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategy *ModifyNodePoolAttributeRequestNodePoolStrategy `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty" type:"Struct"`
	PoolId           *string                                         `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// 产品类型。
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodePoolStrategy(v *ModifyNodePoolAttributeRequestNodePoolStrategy) *ModifyNodePoolAttributeRequest {
	s.NodePoolStrategy = v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetPoolId(v string) *ModifyNodePoolAttributeRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetProductType(v string) *ModifyNodePoolAttributeRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategy struct {
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// 购买资源的数量。取值范围：1~100。
	//
	// >
	// - 若为包年包月资源，则该参数不可修改。
	// - 若为按量付费资源，则当弹性模式（`StrategyType`）为固定数量（`NODE_FIXED`）或自动扩缩容（`NODE_SCALING_BY_USAGE`）时该参数可修改。
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// 策略执行周期列表。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	RecurrenceSchedules         []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	ScalingDownAfterIdleMinutes *int32                                                               `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingStep                 *int32                                                               `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                              `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// 策略失效日期。格式为：yyyy-MM-dd。失效日期与生效日期的间隔必须介于7天到1年之间（含7天和1年）。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// 策略生效日期。格式为：yyyy-MM-dd。该日期必须大于或等于当前日期。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	StrategyType       *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// 是否开启资源预热策略。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetMaxScalingAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.MaxScalingAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetNodeAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.NodeAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetRecurrenceSchedules(v []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.RecurrenceSchedules = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingDownAfterIdleMinutes(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingStep(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingStep = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingUsageThreshold(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyDisableDate(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyDisableDate = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyEnableDate(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyEnableDate = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetWarmUp(v bool) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.WarmUp = &v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules struct {
	// 策略执行周期的类型。必须同时指定`RecurrenceType`和`RecurrenceValues`。
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// 策略执行周期的数值列表。
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// 策略执行周期的时间段列表。时间段设置要求：
	//
	// - 最多可添加3个时间段。
	// - 时间段之间不重叠。
	// - 时间段之间的间隔大于或等于5分钟。
	// - 单个时间段的时长大于或等于15分钟。
	// - 所有时间段累计不跨天。
	TimerPeriods []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetRecurrenceType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetTimerPeriods(v []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods struct {
	// 资源数量。
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// 结束时间。格式为HH:mm。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 开始时间。格式为HH:mm。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type ModifyNodePoolAttributeShrinkRequest struct {
	BizRegionId            *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NodeCapacity           *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategyShrink *string `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty"`
	PoolId                 *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// 产品类型。
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeShrinkRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodePoolStrategyShrink(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.NodePoolStrategyShrink = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetPoolId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetProductType(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodePoolAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponseBody) SetCode(v string) *ModifyNodePoolAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetMessage(v string) *ModifyNodePoolAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetRequestId(v string) *ModifyNodePoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNodePoolAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyNodePoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNodePoolAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponse) SetHeaders(v map[string]*string) *ModifyNodePoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetStatusCode(v int32) *ModifyNodePoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetBody(v *ModifyNodePoolAttributeResponseBody) *ModifyNodePoolAttributeResponse {
	s.Body = v
	return s
}

type ModifyProjectPolicyRequest struct {
	Clipboard                    *int32  `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	FileTransfer                 *int32  `json:"FileTransfer,omitempty" xml:"FileTransfer,omitempty"`
	FrameRate                    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	KeepAliveDuration            *int32  `json:"KeepAliveDuration,omitempty" xml:"KeepAliveDuration,omitempty"`
	ProjectId                    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SessionResolutionHeight      *int32  `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	SessionResolutionWidth       *int32  `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	StreamingMode                *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	TerminalResolutionAdaptation *bool   `json:"TerminalResolutionAdaptation,omitempty" xml:"TerminalResolutionAdaptation,omitempty"`
}

func (s ModifyProjectPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyProjectPolicyRequest) SetClipboard(v int32) *ModifyProjectPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetFileTransfer(v int32) *ModifyProjectPolicyRequest {
	s.FileTransfer = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetFrameRate(v int32) *ModifyProjectPolicyRequest {
	s.FrameRate = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetKeepAliveDuration(v int32) *ModifyProjectPolicyRequest {
	s.KeepAliveDuration = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetProjectId(v string) *ModifyProjectPolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetSessionResolutionHeight(v int32) *ModifyProjectPolicyRequest {
	s.SessionResolutionHeight = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetSessionResolutionWidth(v int32) *ModifyProjectPolicyRequest {
	s.SessionResolutionWidth = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetStreamingMode(v string) *ModifyProjectPolicyRequest {
	s.StreamingMode = &v
	return s
}

func (s *ModifyProjectPolicyRequest) SetTerminalResolutionAdaptation(v bool) *ModifyProjectPolicyRequest {
	s.TerminalResolutionAdaptation = &v
	return s
}

type ModifyProjectPolicyResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyProjectPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProjectPolicyResponseBody) SetCode(v string) *ModifyProjectPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyProjectPolicyResponseBody) SetData(v string) *ModifyProjectPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyProjectPolicyResponseBody) SetMessage(v string) *ModifyProjectPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyProjectPolicyResponseBody) SetRequestId(v string) *ModifyProjectPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyProjectPolicyResponseBody) SetSuccess(v string) *ModifyProjectPolicyResponseBody {
	s.Success = &v
	return s
}

type ModifyProjectPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyProjectPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProjectPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyProjectPolicyResponse) SetHeaders(v map[string]*string) *ModifyProjectPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyProjectPolicyResponse) SetStatusCode(v int32) *ModifyProjectPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProjectPolicyResponse) SetBody(v *ModifyProjectPolicyResponseBody) *ModifyProjectPolicyResponse {
	s.Body = v
	return s
}

type ModifyTenantConfigRequest struct {
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
}

func (s ModifyTenantConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigRequest) SetAppInstanceGroupExpireRemind(v bool) *ModifyTenantConfigRequest {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

type ModifyTenantConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigResponseBody) SetRequestId(v string) *ModifyTenantConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigResponse) SetHeaders(v map[string]*string) *ModifyTenantConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantConfigResponse) SetStatusCode(v int32) *ModifyTenantConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantConfigResponse) SetBody(v *ModifyTenantConfigResponseBody) *ModifyTenantConfigResponse {
	s.Body = v
	return s
}

type PageListAppInstanceGroupUserRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s PageListAppInstanceGroupUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserRequest) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserRequest) SetAppInstanceGroupId(v string) *PageListAppInstanceGroupUserRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageNumber(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageNumber = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageSize(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageSize = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetProductType(v string) *PageListAppInstanceGroupUserRequest {
	s.ProductType = &v
	return s
}

type PageListAppInstanceGroupUserResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s PageListAppInstanceGroupUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponseBody) SetRequestId(v string) *PageListAppInstanceGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponseBody) SetUsers(v []*string) *PageListAppInstanceGroupUserResponseBody {
	s.Users = v
	return s
}

type PageListAppInstanceGroupUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageListAppInstanceGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageListAppInstanceGroupUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponse) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponse) SetHeaders(v map[string]*string) *PageListAppInstanceGroupUserResponse {
	s.Headers = v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetStatusCode(v int32) *PageListAppInstanceGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetBody(v *PageListAppInstanceGroupUserResponseBody) *PageListAppInstanceGroupUserResponse {
	s.Body = v
	return s
}

type RefreshAccessUrlRequest struct {
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
}

func (s RefreshAccessUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshAccessUrlRequest) GoString() string {
	return s.String()
}

func (s *RefreshAccessUrlRequest) SetAccessPageId(v string) *RefreshAccessUrlRequest {
	s.AccessPageId = &v
	return s
}

type RefreshAccessUrlResponseBody struct {
	AccessUrl *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshAccessUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshAccessUrlResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAccessUrlResponseBody) SetAccessUrl(v string) *RefreshAccessUrlResponseBody {
	s.AccessUrl = &v
	return s
}

func (s *RefreshAccessUrlResponseBody) SetCode(v string) *RefreshAccessUrlResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshAccessUrlResponseBody) SetMessage(v string) *RefreshAccessUrlResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshAccessUrlResponseBody) SetRequestId(v string) *RefreshAccessUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAccessUrlResponseBody) SetSuccess(v string) *RefreshAccessUrlResponseBody {
	s.Success = &v
	return s
}

type RefreshAccessUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshAccessUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshAccessUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshAccessUrlResponse) GoString() string {
	return s.String()
}

func (s *RefreshAccessUrlResponse) SetHeaders(v map[string]*string) *RefreshAccessUrlResponse {
	s.Headers = v
	return s
}

func (s *RefreshAccessUrlResponse) SetStatusCode(v int32) *RefreshAccessUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAccessUrlResponse) SetBody(v *RefreshAccessUrlResponseBody) *RefreshAccessUrlResponse {
	s.Body = v
	return s
}

type RenewAppInstanceGroupRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	PromotionId        *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s RenewAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *RenewAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetAutoPay(v bool) *RenewAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPeriod(v int32) *RenewAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPeriodUnit(v string) *RenewAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetProductType(v string) *RenewAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPromotionId(v string) *RenewAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

type RenewAppInstanceGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupResponseBody) SetCode(v string) *RenewAppInstanceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetMessage(v string) *RenewAppInstanceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetOrderId(v string) *RenewAppInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetRequestId(v string) *RenewAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type RenewAppInstanceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupResponse) SetHeaders(v map[string]*string) *RenewAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewAppInstanceGroupResponse) SetStatusCode(v int32) *RenewAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppInstanceGroupResponse) SetBody(v *RenewAppInstanceGroupResponseBody) *RenewAppInstanceGroupResponse {
	s.Body = v
	return s
}

type RenewSessionPackageRequest struct {
	Period           *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SessionPackageId *string `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
}

func (s RenewSessionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewSessionPackageRequest) GoString() string {
	return s.String()
}

func (s *RenewSessionPackageRequest) SetPeriod(v int32) *RenewSessionPackageRequest {
	s.Period = &v
	return s
}

func (s *RenewSessionPackageRequest) SetPeriodUnit(v string) *RenewSessionPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewSessionPackageRequest) SetSessionPackageId(v string) *RenewSessionPackageRequest {
	s.SessionPackageId = &v
	return s
}

type RenewSessionPackageResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionPackageId *int64  `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	Success          *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewSessionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewSessionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *RenewSessionPackageResponseBody) SetCode(v string) *RenewSessionPackageResponseBody {
	s.Code = &v
	return s
}

func (s *RenewSessionPackageResponseBody) SetMessage(v string) *RenewSessionPackageResponseBody {
	s.Message = &v
	return s
}

func (s *RenewSessionPackageResponseBody) SetRequestId(v string) *RenewSessionPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewSessionPackageResponseBody) SetSessionPackageId(v int64) *RenewSessionPackageResponseBody {
	s.SessionPackageId = &v
	return s
}

func (s *RenewSessionPackageResponseBody) SetSuccess(v string) *RenewSessionPackageResponseBody {
	s.Success = &v
	return s
}

type RenewSessionPackageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewSessionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewSessionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewSessionPackageResponse) GoString() string {
	return s.String()
}

func (s *RenewSessionPackageResponse) SetHeaders(v map[string]*string) *RenewSessionPackageResponse {
	s.Headers = v
	return s
}

func (s *RenewSessionPackageResponse) SetStatusCode(v int32) *RenewSessionPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewSessionPackageResponse) SetBody(v *RenewSessionPackageResponseBody) *RenewSessionPackageResponse {
	s.Body = v
	return s
}

type UnbindRequest struct {
	AppInstanceGroupId      *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId           *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	EndUserId               *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ProductType             *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s UnbindRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) SetAppInstanceGroupId(v string) *UnbindRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceId(v string) *UnbindRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindRequest) SetAppInstancePersistentId(v string) *UnbindRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *UnbindRequest) SetEndUserId(v string) *UnbindRequest {
	s.EndUserId = &v
	return s
}

func (s *UnbindRequest) SetProductType(v string) *UnbindRequest {
	s.ProductType = &v
	return s
}

type UnbindResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResponseBody) SetRequestId(v string) *UnbindResponseBody {
	s.RequestId = &v
	return s
}

type UnbindResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponse) GoString() string {
	return s.String()
}

func (s *UnbindResponse) SetHeaders(v map[string]*string) *UnbindResponse {
	s.Headers = v
	return s
}

func (s *UnbindResponse) SetStatusCode(v int32) *UnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResponse) SetBody(v *UnbindResponseBody) *UnbindResponse {
	s.Body = v
	return s
}

type UpdateAccessPageStateRequest struct {
	AccessPageId    *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	AccessPageState *int32  `json:"AccessPageState,omitempty" xml:"AccessPageState,omitempty"`
}

func (s UpdateAccessPageStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessPageStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessPageStateRequest) SetAccessPageId(v string) *UpdateAccessPageStateRequest {
	s.AccessPageId = &v
	return s
}

func (s *UpdateAccessPageStateRequest) SetAccessPageState(v int32) *UpdateAccessPageStateRequest {
	s.AccessPageState = &v
	return s
}

type UpdateAccessPageStateResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAccessPageStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessPageStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessPageStateResponseBody) SetCode(v string) *UpdateAccessPageStateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccessPageStateResponseBody) SetMessage(v string) *UpdateAccessPageStateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccessPageStateResponseBody) SetRequestId(v string) *UpdateAccessPageStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccessPageStateResponseBody) SetSuccess(v string) *UpdateAccessPageStateResponseBody {
	s.Success = &v
	return s
}

type UpdateAccessPageStateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAccessPageStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAccessPageStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessPageStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessPageStateResponse) SetHeaders(v map[string]*string) *UpdateAccessPageStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessPageStateResponse) SetStatusCode(v int32) *UpdateAccessPageStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessPageStateResponse) SetBody(v *UpdateAccessPageStateResponseBody) *UpdateAccessPageStateResponse {
	s.Body = v
	return s
}

type UpdateAppInstanceGroupImageRequest struct {
	AppCenterImageId   *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s UpdateAppInstanceGroupImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppCenterImageId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppInstanceGroupId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetBizRegionId(v string) *UpdateAppInstanceGroupImageRequest {
	s.BizRegionId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetProductType(v string) *UpdateAppInstanceGroupImageRequest {
	s.ProductType = &v
	return s
}

type UpdateAppInstanceGroupImageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppInstanceGroupImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetCode(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetMessage(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetRequestId(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppInstanceGroupImageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppInstanceGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppInstanceGroupImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponse) SetHeaders(v map[string]*string) *UpdateAppInstanceGroupImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetStatusCode(v int32) *UpdateAppInstanceGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetBody(v *UpdateAppInstanceGroupImageResponseBody) *UpdateAppInstanceGroupImageResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AccessPageGetAclWithOptions(request *AccessPageGetAclRequest, runtime *util.RuntimeOptions) (_result *AccessPageGetAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AccessPageGetAcl"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AccessPageGetAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AccessPageGetAcl(request *AccessPageGetAclRequest) (_result *AccessPageGetAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AccessPageGetAclResponse{}
	_body, _err := client.AccessPageGetAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AccessPageSetAclWithOptions(request *AccessPageSetAclRequest, runtime *util.RuntimeOptions) (_result *AccessPageSetAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessMode)) {
		query["AccessMode"] = request.AccessMode
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageName)) {
		query["AccessPageName"] = request.AccessPageName
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTime)) {
		query["EffectTime"] = request.EffectTime
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		query["Unit"] = request.Unit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AccessPageSetAcl"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AccessPageSetAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AccessPageSetAcl(request *AccessPageSetAclRequest) (_result *AccessPageSetAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AccessPageSetAclResponse{}
	_body, _err := client.AccessPageSetAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveOtaTaskWithOptions(request *ApproveOtaTaskRequest, runtime *util.RuntimeOptions) (_result *ApproveOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		body["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveOtaTask(request *ApproveOtaTaskRequest) (_result *ApproveOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveOtaTaskResponse{}
	_body, _err := client.ApproveOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AskSessionPackagePriceWithOptions(request *AskSessionPackagePriceRequest, runtime *util.RuntimeOptions) (_result *AskSessionPackagePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSessions)) {
		query["MaxSessions"] = request.MaxSessions
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageType)) {
		query["SessionPackageType"] = request.SessionPackageType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionSpec)) {
		query["SessionSpec"] = request.SessionSpec
	}

	if !tea.BoolValue(util.IsUnset(request.SessionType)) {
		query["SessionType"] = request.SessionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AskSessionPackagePrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AskSessionPackagePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AskSessionPackagePrice(request *AskSessionPackagePriceRequest) (_result *AskSessionPackagePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AskSessionPackagePriceResponse{}
	_body, _err := client.AskSessionPackagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AskSessionPackageRenewPriceWithOptions(request *AskSessionPackageRenewPriceRequest, runtime *util.RuntimeOptions) (_result *AskSessionPackageRenewPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageId)) {
		query["SessionPackageId"] = request.SessionPackageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AskSessionPackageRenewPrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AskSessionPackageRenewPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AskSessionPackageRenewPrice(request *AskSessionPackageRenewPriceRequest) (_result *AskSessionPackageRenewPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AskSessionPackageRenewPriceResponse{}
	_body, _err := client.AskSessionPackageRenewPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeInstanceGroupWithOptions(request *AuthorizeInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeUserIds)) {
		body["AuthorizeUserIds"] = request.AuthorizeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.UnAuthorizeUserIds)) {
		body["UnAuthorizeUserIds"] = request.UnAuthorizeUserIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeInstanceGroup(request *AuthorizeInstanceGroupRequest) (_result *AuthorizeInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeInstanceGroupResponse{}
	_body, _err := client.AuthorizeInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BuySessionPackageWithOptions(request *BuySessionPackageRequest, runtime *util.RuntimeOptions) (_result *BuySessionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSessions)) {
		query["MaxSessions"] = request.MaxSessions
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageName)) {
		query["SessionPackageName"] = request.SessionPackageName
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageType)) {
		query["SessionPackageType"] = request.SessionPackageType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionSpec)) {
		query["SessionSpec"] = request.SessionSpec
	}

	if !tea.BoolValue(util.IsUnset(request.SessionType)) {
		query["SessionType"] = request.SessionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BuySessionPackage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BuySessionPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BuySessionPackage(request *BuySessionPackageRequest) (_result *BuySessionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BuySessionPackageResponse{}
	_body, _err := client.BuySessionPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOtaTaskWithOptions(request *CancelOtaTaskRequest, runtime *util.RuntimeOptions) (_result *CancelOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOtaTask(request *CancelOtaTaskRequest) (_result *CancelOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOtaTaskResponse{}
	_body, _err := client.CancelOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessPageWithOptions(request *CreateAccessPageRequest, runtime *util.RuntimeOptions) (_result *CreateAccessPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageName)) {
		query["AccessPageName"] = request.AccessPageName
	}

	if !tea.BoolValue(util.IsUnset(request.CloudEnvId)) {
		query["CloudEnvId"] = request.CloudEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTime)) {
		query["EffectTime"] = request.EffectTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		query["Unit"] = request.Unit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessPage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessPage(request *CreateAccessPageRequest) (_result *CreateAccessPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessPageResponse{}
	_body, _err := client.CreateAccessPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppInstanceGroupWithOptions(tmpReq *CreateAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Network)) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, tea.String("Network"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodePool)) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, tea.String("NodePool"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RuntimePolicy)) {
		request.RuntimePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimePolicy, tea.String("RuntimePolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SecurityPolicy)) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, tea.String("SecurityPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StoragePolicy)) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, tea.String("StoragePolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		body["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		body["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeResourceMode)) {
		body["ChargeResourceMode"] = request.ChargeResourceMode
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkShrink)) {
		body["Network"] = request.NetworkShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PreOpenAppId)) {
		body["PreOpenAppId"] = request.PreOpenAppId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		body["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuntimePolicyShrink)) {
		body["RuntimePolicy"] = request.RuntimePolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyShrink)) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		body["SessionTimeout"] = request.SessionTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.StoragePolicyShrink)) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppInstanceGroup(request *CreateAppInstanceGroupRequest) (_result *CreateAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppInstanceGroupResponse{}
	_body, _err := client.CreateAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageFromAppInstanceGroupWithOptions(request *CreateImageFromAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateImageFromAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageName)) {
		body["AppCenterImageName"] = request.AppCenterImageName
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageFromAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageFromAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageFromAppInstanceGroup(request *CreateImageFromAppInstanceGroupRequest) (_result *CreateImageFromAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageFromAppInstanceGroupResponse{}
	_body, _err := client.CreateImageFromAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		query["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.CloudEnvId)) {
		query["CloudEnvId"] = request.CloudEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentId)) {
		query["ContentId"] = request.ContentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileTransfer)) {
		query["FileTransfer"] = request.FileTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.FrameRate)) {
		query["FrameRate"] = request.FrameRate
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAliveDuration)) {
		query["KeepAliveDuration"] = request.KeepAliveDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SessionResolutionHeight)) {
		query["SessionResolutionHeight"] = request.SessionResolutionHeight
	}

	if !tea.BoolValue(util.IsUnset(request.SessionResolutionWidth)) {
		query["SessionResolutionWidth"] = request.SessionResolutionWidth
	}

	if !tea.BoolValue(util.IsUnset(request.SessionSpec)) {
		query["SessionSpec"] = request.SessionSpec
	}

	if !tea.BoolValue(util.IsUnset(request.StreamingMode)) {
		query["StreamingMode"] = request.StreamingMode
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalResolutionAdaptation)) {
		query["TerminalResolutionAdaptation"] = request.TerminalResolutionAdaptation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessPageWithOptions(request *DeleteAccessPageRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessPage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessPage(request *DeleteAccessPageRequest) (_result *DeleteAccessPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessPageResponse{}
	_body, _err := client.DeleteAccessPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppInstanceGroupWithOptions(request *DeleteAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppInstanceGroup(request *DeleteAppInstanceGroupRequest) (_result *DeleteAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppInstanceGroupResponse{}
	_body, _err := client.DeleteAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppInstancesWithOptions(request *DeleteAppInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteAppInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceIds)) {
		body["AppInstanceIds"] = request.AppInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppInstances"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppInstances(request *DeleteAppInstancesRequest) (_result *DeleteAppInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppInstancesResponse{}
	_body, _err := client.DeleteAppInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessPageSessionWithOptions(request *GetAccessPageSessionRequest, runtime *util.RuntimeOptions) (_result *GetAccessPageSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageToken)) {
		query["AccessPageToken"] = request.AccessPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUserId)) {
		query["ExternalUserId"] = request.ExternalUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessPageSession"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessPageSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessPageSession(request *GetAccessPageSessionRequest) (_result *GetAccessPageSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessPageSessionResponse{}
	_body, _err := client.GetAccessPageSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppInstanceGroupWithOptions(request *GetAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *GetAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppInstanceGroup(request *GetAppInstanceGroupRequest) (_result *GetAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppInstanceGroupResponse{}
	_body, _err := client.GetAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupIdList)) {
		body["AppInstanceGroupIdList"] = request.AppInstanceGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstancePersistentId)) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !tea.BoolValue(util.IsUnset(request.AppStartParam)) {
		body["AppStartParam"] = request.AppStartParam
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDebugAppInstanceWithOptions(request *GetDebugAppInstanceRequest, runtime *util.RuntimeOptions) (_result *GetDebugAppInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDebugAppInstance"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDebugAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDebugAppInstance(request *GetDebugAppInstanceRequest) (_result *GetDebugAppInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDebugAppInstanceResponse{}
	_body, _err := client.GetDebugAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOtaTaskByTaskIdWithOptions(request *GetOtaTaskByTaskIdRequest, runtime *util.RuntimeOptions) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOtaTaskByTaskId"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOtaTaskByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOtaTaskByTaskId(request *GetOtaTaskByTaskIdRequest) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOtaTaskByTaskIdResponse{}
	_body, _err := client.GetOtaTaskByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectPoliciesWithOptions(request *GetProjectPoliciesRequest, runtime *util.RuntimeOptions) (_result *GetProjectPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectPolicies"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectPolicies(request *GetProjectPoliciesRequest) (_result *GetProjectPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectPoliciesResponse{}
	_body, _err := client.GetProjectPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourcePriceWithOptions(request *GetResourcePriceRequest, runtime *util.RuntimeOptions) (_result *GetResourcePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceType)) {
		query["AppInstanceType"] = request.AppInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourcePrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourcePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourcePrice(request *GetResourcePriceRequest) (_result *GetResourcePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourcePriceResponse{}
	_body, _err := client.GetResourcePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceRenewPriceWithOptions(request *GetResourceRenewPriceRequest, runtime *util.RuntimeOptions) (_result *GetResourceRenewPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceRenewPrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceRenewPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceRenewPrice(request *GetResourceRenewPriceRequest) (_result *GetResourceRenewPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceRenewPriceResponse{}
	_body, _err := client.GetResourceRenewPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessPagesWithOptions(request *ListAccessPagesRequest, runtime *util.RuntimeOptions) (_result *ListAccessPagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageName)) {
		query["AccessPageName"] = request.AccessPageName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessPages"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessPagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessPages(request *ListAccessPagesRequest) (_result *ListAccessPagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessPagesResponse{}
	_body, _err := client.ListAccessPagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppInstanceGroupWithOptions(request *ListAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *ListAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppInstanceGroup(request *ListAppInstanceGroupRequest) (_result *ListAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInstanceGroupResponse{}
	_body, _err := client.ListAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppInstancesWithOptions(request *ListAppInstancesRequest, runtime *util.RuntimeOptions) (_result *ListAppInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDeleted)) {
		query["IncludeDeleted"] = request.IncludeDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceIdList)) {
		body["AppInstanceIdList"] = request.AppInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppInstances"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppInstances(request *ListAppInstancesRequest) (_result *ListAppInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInstancesResponse{}
	_body, _err := client.ListAppInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeInstanceTypeWithOptions(request *ListNodeInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *ListNodeInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeInstanceType"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeInstanceType(request *ListNodeInstanceTypeRequest) (_result *ListNodeInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeInstanceTypeResponse{}
	_body, _err := client.ListNodeInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOtaTaskWithOptions(request *ListOtaTaskRequest, runtime *util.RuntimeOptions) (_result *ListOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		body["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOtaTask(request *ListOtaTaskRequest) (_result *ListOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOtaTaskResponse{}
	_body, _err := client.ListOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.StateList)) {
		query["StateList"] = request.StateList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSessionPackagesWithOptions(request *ListSessionPackagesRequest, runtime *util.RuntimeOptions) (_result *ListSessionPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageId)) {
		query["SessionPackageId"] = request.SessionPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageName)) {
		query["SessionPackageName"] = request.SessionPackageName
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.StateList)) {
		query["StateList"] = request.StateList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSessionPackages"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSessionPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSessionPackages(request *ListSessionPackagesRequest) (_result *ListSessionPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSessionPackagesResponse{}
	_body, _err := client.ListSessionPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTenantConfigWithOptions(runtime *util.RuntimeOptions) (_result *ListTenantConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListTenantConfig"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTenantConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTenantConfig() (_result *ListTenantConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTenantConfigResponse{}
	_body, _err := client.ListTenantConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LogOffAllSessionsInAppInstanceGroupWithOptions(request *LogOffAllSessionsInAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *LogOffAllSessionsInAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LogOffAllSessionsInAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LogOffAllSessionsInAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LogOffAllSessionsInAppInstanceGroup(request *LogOffAllSessionsInAppInstanceGroupRequest) (_result *LogOffAllSessionsInAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LogOffAllSessionsInAppInstanceGroupResponse{}
	_body, _err := client.LogOffAllSessionsInAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateSessionPackageWithOptions(request *MigrateSessionPackageRequest, runtime *util.RuntimeOptions) (_result *MigrateSessionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestProjectId)) {
		body["DestProjectId"] = request.DestProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageId)) {
		body["SessionPackageId"] = request.SessionPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceProjectId)) {
		body["SourceProjectId"] = request.SourceProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MigrateSessionPackage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MigrateSessionPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateSessionPackage(request *MigrateSessionPackageRequest) (_result *MigrateSessionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateSessionPackageResponse{}
	_body, _err := client.MigrateSessionPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppInstanceGroupAttributeWithOptions(tmpReq *ModifyAppInstanceGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppInstanceGroupAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Network)) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, tea.String("Network"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodePool)) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, tea.String("NodePool"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SecurityPolicy)) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, tea.String("SecurityPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StoragePolicy)) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, tea.String("StoragePolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		query["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkShrink)) {
		body["Network"] = request.NetworkShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PreOpenAppId)) {
		body["PreOpenAppId"] = request.PreOpenAppId
	}

	if !tea.BoolValue(util.IsUnset(request.PreOpenMode)) {
		body["PreOpenMode"] = request.PreOpenMode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyShrink)) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StoragePolicyShrink)) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppInstanceGroupAttribute"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppInstanceGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppInstanceGroupAttribute(request *ModifyAppInstanceGroupAttributeRequest) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppInstanceGroupAttributeResponse{}
	_body, _err := client.ModifyAppInstanceGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNodePoolAttributeWithOptions(tmpReq *ModifyNodePoolAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyNodePoolAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyNodePoolAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodePoolStrategy)) {
		request.NodePoolStrategyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePoolStrategy, tea.String("NodePoolStrategy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCapacity)) {
		body["NodeCapacity"] = request.NodeCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolStrategyShrink)) {
		body["NodePoolStrategy"] = request.NodePoolStrategyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PoolId)) {
		body["PoolId"] = request.PoolId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodePoolAttribute"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNodePoolAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNodePoolAttribute(request *ModifyNodePoolAttributeRequest) (_result *ModifyNodePoolAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNodePoolAttributeResponse{}
	_body, _err := client.ModifyNodePoolAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProjectPolicyWithOptions(request *ModifyProjectPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyProjectPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		query["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.FileTransfer)) {
		query["FileTransfer"] = request.FileTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.FrameRate)) {
		query["FrameRate"] = request.FrameRate
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAliveDuration)) {
		query["KeepAliveDuration"] = request.KeepAliveDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionResolutionHeight)) {
		query["SessionResolutionHeight"] = request.SessionResolutionHeight
	}

	if !tea.BoolValue(util.IsUnset(request.SessionResolutionWidth)) {
		query["SessionResolutionWidth"] = request.SessionResolutionWidth
	}

	if !tea.BoolValue(util.IsUnset(request.StreamingMode)) {
		query["StreamingMode"] = request.StreamingMode
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalResolutionAdaptation)) {
		query["TerminalResolutionAdaptation"] = request.TerminalResolutionAdaptation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProjectPolicy"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProjectPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProjectPolicy(request *ModifyProjectPolicyRequest) (_result *ModifyProjectPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProjectPolicyResponse{}
	_body, _err := client.ModifyProjectPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantConfigWithOptions(request *ModifyTenantConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupExpireRemind)) {
		body["AppInstanceGroupExpireRemind"] = request.AppInstanceGroupExpireRemind
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantConfig"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantConfig(request *ModifyTenantConfigRequest) (_result *ModifyTenantConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantConfigResponse{}
	_body, _err := client.ModifyTenantConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListAppInstanceGroupUserWithOptions(request *PageListAppInstanceGroupUserRequest, runtime *util.RuntimeOptions) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageListAppInstanceGroupUser"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListAppInstanceGroupUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListAppInstanceGroupUser(request *PageListAppInstanceGroupUserRequest) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageListAppInstanceGroupUserResponse{}
	_body, _err := client.PageListAppInstanceGroupUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshAccessUrlWithOptions(request *RefreshAccessUrlRequest, runtime *util.RuntimeOptions) (_result *RefreshAccessUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshAccessUrl"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshAccessUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshAccessUrl(request *RefreshAccessUrlRequest) (_result *RefreshAccessUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshAccessUrlResponse{}
	_body, _err := client.RefreshAccessUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewAppInstanceGroupWithOptions(request *RenewAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *RenewAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewAppInstanceGroup(request *RenewAppInstanceGroupRequest) (_result *RenewAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewAppInstanceGroupResponse{}
	_body, _err := client.RenewAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewSessionPackageWithOptions(request *RenewSessionPackageRequest, runtime *util.RuntimeOptions) (_result *RenewSessionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageId)) {
		query["SessionPackageId"] = request.SessionPackageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewSessionPackage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewSessionPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewSessionPackage(request *RenewSessionPackageRequest) (_result *RenewSessionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewSessionPackageResponse{}
	_body, _err := client.RenewSessionPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindWithOptions(request *UnbindRequest, runtime *util.RuntimeOptions) (_result *UnbindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstancePersistentId)) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Unbind"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Unbind(request *UnbindRequest) (_result *UnbindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindResponse{}
	_body, _err := client.UnbindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAccessPageStateWithOptions(request *UpdateAccessPageStateRequest, runtime *util.RuntimeOptions) (_result *UpdateAccessPageStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageState)) {
		query["AccessPageState"] = request.AccessPageState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccessPageState"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccessPageStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAccessPageState(request *UpdateAccessPageStateRequest) (_result *UpdateAccessPageStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccessPageStateResponse{}
	_body, _err := client.UpdateAccessPageStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppInstanceGroupImageWithOptions(request *UpdateAppInstanceGroupImageRequest, runtime *util.RuntimeOptions) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppInstanceGroupImage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppInstanceGroupImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppInstanceGroupImage(request *UpdateAppInstanceGroupImageRequest) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppInstanceGroupImageResponse{}
	_body, _err := client.UpdateAppInstanceGroupImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
