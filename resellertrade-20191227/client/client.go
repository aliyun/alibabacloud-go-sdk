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

type CreateCouponTemplateRequest struct {
	ActivitySite    *int32                               `json:"ActivitySite,omitempty" xml:"ActivitySite,omitempty"`
	Bid             *int64                               `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CertainMoney    *float64                             `json:"CertainMoney,omitempty" xml:"CertainMoney,omitempty"`
	ClientType      *string                              `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CommodityType   *string                              `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	ControlType     *string                              `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	CouponAmount    *float64                             `json:"CouponAmount,omitempty" xml:"CouponAmount,omitempty"`
	CouponEndTime   *string                              `json:"CouponEndTime,omitempty" xml:"CouponEndTime,omitempty"`
	CouponFixedType *string                              `json:"CouponFixedType,omitempty" xml:"CouponFixedType,omitempty"`
	CouponStartTime *string                              `json:"CouponStartTime,omitempty" xml:"CouponStartTime,omitempty"`
	CouponType      *string                              `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	Currency        *CreateCouponTemplateRequestCurrency `json:"Currency,omitempty" xml:"Currency,omitempty" type:"Struct"`
	Description     *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DiscountRate    *float64                             `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`
	EndTime         *string                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtendsMap      map[string]*string                   `json:"ExtendsMap,omitempty" xml:"ExtendsMap,omitempty"`
	FromApp         *string                              `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	ItemCodeSet     []*string                            `json:"ItemCodeSet,omitempty" xml:"ItemCodeSet,omitempty" type:"Repeated"`
	Market          *string                              `json:"Market,omitempty" xml:"Market,omitempty"`
	MarketType      *int64                               `json:"MarketType,omitempty" xml:"MarketType,omitempty"`
	MaxRelease      *float64                             `json:"MaxRelease,omitempty" xml:"MaxRelease,omitempty"`
	MessageId       *string                              `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Name            *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator        *string                              `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OrderTypeSet    []*string                            `json:"OrderTypeSet,omitempty" xml:"OrderTypeSet,omitempty" type:"Repeated"`
	PerLimitNum     *int32                               `json:"PerLimitNum,omitempty" xml:"PerLimitNum,omitempty"`
	PromotionId     *int64                               `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	Reason          *string                              `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RelativeSecond  *int32                               `json:"RelativeSecond,omitempty" xml:"RelativeSecond,omitempty"`
	RequestId       *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelectionIdSet  []*string                            `json:"SelectionIdSet,omitempty" xml:"SelectionIdSet,omitempty" type:"Repeated"`
	SellerId        *int64                               `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	Site            *string                              `json:"Site,omitempty" xml:"Site,omitempty"`
	SpId            *int64                               `json:"SpId,omitempty" xml:"SpId,omitempty"`
	StartTime       *string                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type            *string                              `json:"Type,omitempty" xml:"Type,omitempty"`
	UniversalType   *string                              `json:"UniversalType,omitempty" xml:"UniversalType,omitempty"`
	UpperLimit      *float64                             `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
	UsageCount      *int32                               `json:"UsageCount,omitempty" xml:"UsageCount,omitempty"`
	UseScene        *string                              `json:"UseScene,omitempty" xml:"UseScene,omitempty"`
	UserPkAmount    *string                              `json:"UserPkAmount,omitempty" xml:"UserPkAmount,omitempty"`
	ValidityType    *string                              `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s CreateCouponTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateRequest) SetActivitySite(v int32) *CreateCouponTemplateRequest {
	s.ActivitySite = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetBid(v int64) *CreateCouponTemplateRequest {
	s.Bid = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCertainMoney(v float64) *CreateCouponTemplateRequest {
	s.CertainMoney = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetClientType(v string) *CreateCouponTemplateRequest {
	s.ClientType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCommodityType(v string) *CreateCouponTemplateRequest {
	s.CommodityType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetControlType(v string) *CreateCouponTemplateRequest {
	s.ControlType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponAmount(v float64) *CreateCouponTemplateRequest {
	s.CouponAmount = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponEndTime(v string) *CreateCouponTemplateRequest {
	s.CouponEndTime = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponFixedType(v string) *CreateCouponTemplateRequest {
	s.CouponFixedType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponStartTime(v string) *CreateCouponTemplateRequest {
	s.CouponStartTime = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCouponType(v string) *CreateCouponTemplateRequest {
	s.CouponType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetCurrency(v *CreateCouponTemplateRequestCurrency) *CreateCouponTemplateRequest {
	s.Currency = v
	return s
}

func (s *CreateCouponTemplateRequest) SetDescription(v string) *CreateCouponTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetDiscountRate(v float64) *CreateCouponTemplateRequest {
	s.DiscountRate = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetEndTime(v string) *CreateCouponTemplateRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetExtendsMap(v map[string]*string) *CreateCouponTemplateRequest {
	s.ExtendsMap = v
	return s
}

func (s *CreateCouponTemplateRequest) SetFromApp(v string) *CreateCouponTemplateRequest {
	s.FromApp = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetItemCodeSet(v []*string) *CreateCouponTemplateRequest {
	s.ItemCodeSet = v
	return s
}

func (s *CreateCouponTemplateRequest) SetMarket(v string) *CreateCouponTemplateRequest {
	s.Market = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetMarketType(v int64) *CreateCouponTemplateRequest {
	s.MarketType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetMaxRelease(v float64) *CreateCouponTemplateRequest {
	s.MaxRelease = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetMessageId(v string) *CreateCouponTemplateRequest {
	s.MessageId = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetName(v string) *CreateCouponTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetOperator(v string) *CreateCouponTemplateRequest {
	s.Operator = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetOrderTypeSet(v []*string) *CreateCouponTemplateRequest {
	s.OrderTypeSet = v
	return s
}

func (s *CreateCouponTemplateRequest) SetPerLimitNum(v int32) *CreateCouponTemplateRequest {
	s.PerLimitNum = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetPromotionId(v int64) *CreateCouponTemplateRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetReason(v string) *CreateCouponTemplateRequest {
	s.Reason = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetRelativeSecond(v int32) *CreateCouponTemplateRequest {
	s.RelativeSecond = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetRequestId(v string) *CreateCouponTemplateRequest {
	s.RequestId = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetSelectionIdSet(v []*string) *CreateCouponTemplateRequest {
	s.SelectionIdSet = v
	return s
}

func (s *CreateCouponTemplateRequest) SetSellerId(v int64) *CreateCouponTemplateRequest {
	s.SellerId = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetSite(v string) *CreateCouponTemplateRequest {
	s.Site = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetSpId(v int64) *CreateCouponTemplateRequest {
	s.SpId = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetStartTime(v string) *CreateCouponTemplateRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetType(v string) *CreateCouponTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetUniversalType(v string) *CreateCouponTemplateRequest {
	s.UniversalType = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetUpperLimit(v float64) *CreateCouponTemplateRequest {
	s.UpperLimit = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetUsageCount(v int32) *CreateCouponTemplateRequest {
	s.UsageCount = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetUseScene(v string) *CreateCouponTemplateRequest {
	s.UseScene = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetUserPkAmount(v string) *CreateCouponTemplateRequest {
	s.UserPkAmount = &v
	return s
}

func (s *CreateCouponTemplateRequest) SetValidityType(v string) *CreateCouponTemplateRequest {
	s.ValidityType = &v
	return s
}

type CreateCouponTemplateRequestCurrency struct {
	CurrencyCode          *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	DefaultFractionDigits *int32  `json:"DefaultFractionDigits,omitempty" xml:"DefaultFractionDigits,omitempty"`
	NumericCode           *int32  `json:"NumericCode,omitempty" xml:"NumericCode,omitempty"`
}

func (s CreateCouponTemplateRequestCurrency) String() string {
	return tea.Prettify(s)
}

func (s CreateCouponTemplateRequestCurrency) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateRequestCurrency) SetCurrencyCode(v string) *CreateCouponTemplateRequestCurrency {
	s.CurrencyCode = &v
	return s
}

func (s *CreateCouponTemplateRequestCurrency) SetDefaultFractionDigits(v int32) *CreateCouponTemplateRequestCurrency {
	s.DefaultFractionDigits = &v
	return s
}

func (s *CreateCouponTemplateRequestCurrency) SetNumericCode(v int32) *CreateCouponTemplateRequestCurrency {
	s.NumericCode = &v
	return s
}

type CreateCouponTemplateResponseBody struct {
	Code       *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{} `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	// result data
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CreateCouponTemplateResponseBody) SetContextMap(v map[string]interface{}) *CreateCouponTemplateResponseBody {
	s.ContextMap = v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetData(v int64) *CreateCouponTemplateResponseBody {
	s.Data = &v
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

type DiscardCouponListRequest struct {
	CouponList []*int64 `json:"CouponList,omitempty" xml:"CouponList,omitempty" type:"Repeated"`
}

func (s DiscardCouponListRequest) String() string {
	return tea.Prettify(s)
}

func (s DiscardCouponListRequest) GoString() string {
	return s.String()
}

func (s *DiscardCouponListRequest) SetCouponList(v []*int64) *DiscardCouponListRequest {
	s.CouponList = v
	return s
}

type DiscardCouponListResponseBody struct {
	Code       *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{} `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	// result data
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DiscardCouponListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DiscardCouponListResponseBody) GoString() string {
	return s.String()
}

func (s *DiscardCouponListResponseBody) SetCode(v string) *DiscardCouponListResponseBody {
	s.Code = &v
	return s
}

func (s *DiscardCouponListResponseBody) SetContextMap(v map[string]interface{}) *DiscardCouponListResponseBody {
	s.ContextMap = v
	return s
}

func (s *DiscardCouponListResponseBody) SetData(v bool) *DiscardCouponListResponseBody {
	s.Data = &v
	return s
}

func (s *DiscardCouponListResponseBody) SetMessage(v string) *DiscardCouponListResponseBody {
	s.Message = &v
	return s
}

func (s *DiscardCouponListResponseBody) SetRequestId(v string) *DiscardCouponListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiscardCouponListResponseBody) SetSuccess(v bool) *DiscardCouponListResponseBody {
	s.Success = &v
	return s
}

type DiscardCouponListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiscardCouponListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiscardCouponListResponse) String() string {
	return tea.Prettify(s)
}

func (s DiscardCouponListResponse) GoString() string {
	return s.String()
}

func (s *DiscardCouponListResponse) SetHeaders(v map[string]*string) *DiscardCouponListResponse {
	s.Headers = v
	return s
}

func (s *DiscardCouponListResponse) SetStatusCode(v int32) *DiscardCouponListResponse {
	s.StatusCode = &v
	return s
}

func (s *DiscardCouponListResponse) SetBody(v *DiscardCouponListResponseBody) *DiscardCouponListResponse {
	s.Body = v
	return s
}

type GetCouponPageRequest struct {
	FromApp    *string  `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	PageNo     *int32   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SellerId   *int64   `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	TemplateId *int64   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Uids       []*int64 `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
}

func (s GetCouponPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCouponPageRequest) GoString() string {
	return s.String()
}

func (s *GetCouponPageRequest) SetFromApp(v string) *GetCouponPageRequest {
	s.FromApp = &v
	return s
}

func (s *GetCouponPageRequest) SetPageNo(v int32) *GetCouponPageRequest {
	s.PageNo = &v
	return s
}

func (s *GetCouponPageRequest) SetPageSize(v int32) *GetCouponPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetCouponPageRequest) SetRequestId(v string) *GetCouponPageRequest {
	s.RequestId = &v
	return s
}

func (s *GetCouponPageRequest) SetSellerId(v int64) *GetCouponPageRequest {
	s.SellerId = &v
	return s
}

func (s *GetCouponPageRequest) SetTemplateId(v int64) *GetCouponPageRequest {
	s.TemplateId = &v
	return s
}

func (s *GetCouponPageRequest) SetUids(v []*int64) *GetCouponPageRequest {
	s.Uids = v
	return s
}

type GetCouponPageResponseBody struct {
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{}           `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Count      *int32                           `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       []*GetCouponPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCouponPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCouponPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetCouponPageResponseBody) SetCode(v string) *GetCouponPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetCouponPageResponseBody) SetContextMap(v map[string]interface{}) *GetCouponPageResponseBody {
	s.ContextMap = v
	return s
}

func (s *GetCouponPageResponseBody) SetCount(v int32) *GetCouponPageResponseBody {
	s.Count = &v
	return s
}

func (s *GetCouponPageResponseBody) SetData(v []*GetCouponPageResponseBodyData) *GetCouponPageResponseBody {
	s.Data = v
	return s
}

func (s *GetCouponPageResponseBody) SetMessage(v string) *GetCouponPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetCouponPageResponseBody) SetRequestId(v string) *GetCouponPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCouponPageResponseBody) SetSuccess(v bool) *GetCouponPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetCouponPageResponseBody) SetTotalCount(v int32) *GetCouponPageResponseBody {
	s.TotalCount = &v
	return s
}

type GetCouponPageResponseBodyData struct {
	Amount        *float64                               `json:"Amount,omitempty" xml:"Amount,omitempty"`
	CertainMoney  *float64                               `json:"CertainMoney,omitempty" xml:"CertainMoney,omitempty"`
	CouponId      *int64                                 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	CouponType    *string                                `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	Currency      *GetCouponPageResponseBodyDataCurrency `json:"Currency,omitempty" xml:"Currency,omitempty" type:"Struct"`
	Description   *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DiscountRate  *float64                               `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`
	EndTime       *string                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FrozenAmount  *float64                               `json:"FrozenAmount,omitempty" xml:"FrozenAmount,omitempty"`
	FrozenCount   *int32                                 `json:"FrozenCount,omitempty" xml:"FrozenCount,omitempty"`
	PromotionId   *int64                                 `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	SellerId      *int64                                 `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	StartTime     *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId    *int64                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UniversalType *string                                `json:"UniversalType,omitempty" xml:"UniversalType,omitempty"`
	UpperLimit    *float64                               `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
	UsageCount    *int32                                 `json:"UsageCount,omitempty" xml:"UsageCount,omitempty"`
	UsedAmount    *float64                               `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
	UsedCount     *int32                                 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	UserId        *int64                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCouponPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCouponPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCouponPageResponseBodyData) SetAmount(v float64) *GetCouponPageResponseBodyData {
	s.Amount = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetCertainMoney(v float64) *GetCouponPageResponseBodyData {
	s.CertainMoney = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetCouponId(v int64) *GetCouponPageResponseBodyData {
	s.CouponId = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetCouponType(v string) *GetCouponPageResponseBodyData {
	s.CouponType = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetCurrency(v *GetCouponPageResponseBodyDataCurrency) *GetCouponPageResponseBodyData {
	s.Currency = v
	return s
}

func (s *GetCouponPageResponseBodyData) SetDescription(v string) *GetCouponPageResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetDiscountRate(v float64) *GetCouponPageResponseBodyData {
	s.DiscountRate = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetEndTime(v string) *GetCouponPageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetFrozenAmount(v float64) *GetCouponPageResponseBodyData {
	s.FrozenAmount = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetFrozenCount(v int32) *GetCouponPageResponseBodyData {
	s.FrozenCount = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetPromotionId(v int64) *GetCouponPageResponseBodyData {
	s.PromotionId = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetSellerId(v int64) *GetCouponPageResponseBodyData {
	s.SellerId = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetStartTime(v string) *GetCouponPageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetStatus(v string) *GetCouponPageResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetTemplateId(v int64) *GetCouponPageResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetUniversalType(v string) *GetCouponPageResponseBodyData {
	s.UniversalType = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetUpperLimit(v float64) *GetCouponPageResponseBodyData {
	s.UpperLimit = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetUsageCount(v int32) *GetCouponPageResponseBodyData {
	s.UsageCount = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetUsedAmount(v float64) *GetCouponPageResponseBodyData {
	s.UsedAmount = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetUsedCount(v int32) *GetCouponPageResponseBodyData {
	s.UsedCount = &v
	return s
}

func (s *GetCouponPageResponseBodyData) SetUserId(v int64) *GetCouponPageResponseBodyData {
	s.UserId = &v
	return s
}

type GetCouponPageResponseBodyDataCurrency struct {
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s GetCouponPageResponseBodyDataCurrency) String() string {
	return tea.Prettify(s)
}

func (s GetCouponPageResponseBodyDataCurrency) GoString() string {
	return s.String()
}

func (s *GetCouponPageResponseBodyDataCurrency) SetCurrencyCode(v string) *GetCouponPageResponseBodyDataCurrency {
	s.CurrencyCode = &v
	return s
}

type GetCouponPageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCouponPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCouponPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCouponPageResponse) GoString() string {
	return s.String()
}

func (s *GetCouponPageResponse) SetHeaders(v map[string]*string) *GetCouponPageResponse {
	s.Headers = v
	return s
}

func (s *GetCouponPageResponse) SetStatusCode(v int32) *GetCouponPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCouponPageResponse) SetBody(v *GetCouponPageResponseBody) *GetCouponPageResponse {
	s.Body = v
	return s
}

type GetCustomerListRequest struct {
	PageNo   *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetCustomerListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerListRequest) SetPageNo(v int32) *GetCustomerListRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerListRequest) SetPageSize(v int32) *GetCustomerListRequest {
	s.PageSize = &v
	return s
}

type GetCustomerListResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCustomerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponseBody) SetCode(v string) *GetCustomerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerListResponseBody) SetData(v *GetCustomerListResponseBodyData) *GetCustomerListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerListResponseBody) SetMessage(v string) *GetCustomerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerListResponseBody) SetRequestId(v string) *GetCustomerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerListResponseBody) SetSuccess(v bool) *GetCustomerListResponseBody {
	s.Success = &v
	return s
}

type GetCustomerListResponseBodyData struct {
	TotalSize *int32    `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	UidList   []*string `json:"UidList,omitempty" xml:"UidList,omitempty" type:"Repeated"`
}

func (s GetCustomerListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponseBodyData) SetTotalSize(v int32) *GetCustomerListResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GetCustomerListResponseBodyData) SetUidList(v []*string) *GetCustomerListResponseBodyData {
	s.UidList = v
	return s
}

type GetCustomerListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponse) SetHeaders(v map[string]*string) *GetCustomerListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerListResponse) SetStatusCode(v int32) *GetCustomerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerListResponse) SetBody(v *GetCustomerListResponseBody) *GetCustomerListResponse {
	s.Body = v
	return s
}

type GetRelationRequest struct {
	RelationTime *int64 `json:"RelationTime,omitempty" xml:"RelationTime,omitempty"`
	Uid          *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelationRequest) GoString() string {
	return s.String()
}

func (s *GetRelationRequest) SetRelationTime(v int64) *GetRelationRequest {
	s.RelationTime = &v
	return s
}

func (s *GetRelationRequest) SetUid(v int64) *GetRelationRequest {
	s.Uid = &v
	return s
}

type GetRelationResponseBody struct {
	Class     *string                      `json:"Class,omitempty" xml:"Class,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelationResponseBody) SetClass(v string) *GetRelationResponseBody {
	s.Class = &v
	return s
}

func (s *GetRelationResponseBody) SetCode(v string) *GetRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetRelationResponseBody) SetData(v *GetRelationResponseBodyData) *GetRelationResponseBody {
	s.Data = v
	return s
}

func (s *GetRelationResponseBody) SetMessage(v string) *GetRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetRelationResponseBody) SetRequestId(v string) *GetRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRelationResponseBody) SetSuccess(v bool) *GetRelationResponseBody {
	s.Success = &v
	return s
}

type GetRelationResponseBodyData struct {
	CanLoginOfficial      *bool                                             `json:"CanLoginOfficial,omitempty" xml:"CanLoginOfficial,omitempty"`
	Class                 *string                                           `json:"Class,omitempty" xml:"Class,omitempty"`
	EndTime               *int64                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ResellerProductModeDO *GetRelationResponseBodyDataResellerProductModeDO `json:"ResellerProductModeDO,omitempty" xml:"ResellerProductModeDO,omitempty" type:"Struct"`
	ResellerProductRuleDO *GetRelationResponseBodyDataResellerProductRuleDO `json:"ResellerProductRuleDO,omitempty" xml:"ResellerProductRuleDO,omitempty" type:"Struct"`
	ResellerUid           *int64                                            `json:"ResellerUid,omitempty" xml:"ResellerUid,omitempty"`
	StartTime             *int64                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	TopReseller           *bool                                             `json:"TopReseller,omitempty" xml:"TopReseller,omitempty"`
	Uid                   *int64                                            `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserType              *string                                           `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRelationResponseBodyData) SetCanLoginOfficial(v bool) *GetRelationResponseBodyData {
	s.CanLoginOfficial = &v
	return s
}

func (s *GetRelationResponseBodyData) SetClass(v string) *GetRelationResponseBodyData {
	s.Class = &v
	return s
}

func (s *GetRelationResponseBodyData) SetEndTime(v int64) *GetRelationResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetRelationResponseBodyData) SetResellerProductModeDO(v *GetRelationResponseBodyDataResellerProductModeDO) *GetRelationResponseBodyData {
	s.ResellerProductModeDO = v
	return s
}

func (s *GetRelationResponseBodyData) SetResellerProductRuleDO(v *GetRelationResponseBodyDataResellerProductRuleDO) *GetRelationResponseBodyData {
	s.ResellerProductRuleDO = v
	return s
}

func (s *GetRelationResponseBodyData) SetResellerUid(v int64) *GetRelationResponseBodyData {
	s.ResellerUid = &v
	return s
}

func (s *GetRelationResponseBodyData) SetStartTime(v int64) *GetRelationResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetRelationResponseBodyData) SetStatus(v string) *GetRelationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRelationResponseBodyData) SetTopReseller(v bool) *GetRelationResponseBodyData {
	s.TopReseller = &v
	return s
}

func (s *GetRelationResponseBodyData) SetUid(v int64) *GetRelationResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetRelationResponseBodyData) SetUserType(v string) *GetRelationResponseBodyData {
	s.UserType = &v
	return s
}

type GetRelationResponseBodyDataResellerProductModeDO struct {
	Class       *string `json:"Class,omitempty" xml:"Class,omitempty"`
	IsService   *int64  `json:"IsService,omitempty" xml:"IsService,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s GetRelationResponseBodyDataResellerProductModeDO) String() string {
	return tea.Prettify(s)
}

func (s GetRelationResponseBodyDataResellerProductModeDO) GoString() string {
	return s.String()
}

func (s *GetRelationResponseBodyDataResellerProductModeDO) SetClass(v string) *GetRelationResponseBodyDataResellerProductModeDO {
	s.Class = &v
	return s
}

func (s *GetRelationResponseBodyDataResellerProductModeDO) SetIsService(v int64) *GetRelationResponseBodyDataResellerProductModeDO {
	s.IsService = &v
	return s
}

func (s *GetRelationResponseBodyDataResellerProductModeDO) SetProductCode(v string) *GetRelationResponseBodyDataResellerProductModeDO {
	s.ProductCode = &v
	return s
}

func (s *GetRelationResponseBodyDataResellerProductModeDO) SetProductName(v string) *GetRelationResponseBodyDataResellerProductModeDO {
	s.ProductName = &v
	return s
}

type GetRelationResponseBodyDataResellerProductRuleDO struct {
	Class          *string `json:"Class,omitempty" xml:"Class,omitempty"`
	CommodityRoute *bool   `json:"CommodityRoute,omitempty" xml:"CommodityRoute,omitempty"`
	MultiOrder     *bool   `json:"MultiOrder,omitempty" xml:"MultiOrder,omitempty"`
	PayMode        *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
}

func (s GetRelationResponseBodyDataResellerProductRuleDO) String() string {
	return tea.Prettify(s)
}

func (s GetRelationResponseBodyDataResellerProductRuleDO) GoString() string {
	return s.String()
}

func (s *GetRelationResponseBodyDataResellerProductRuleDO) SetClass(v string) *GetRelationResponseBodyDataResellerProductRuleDO {
	s.Class = &v
	return s
}

func (s *GetRelationResponseBodyDataResellerProductRuleDO) SetCommodityRoute(v bool) *GetRelationResponseBodyDataResellerProductRuleDO {
	s.CommodityRoute = &v
	return s
}

func (s *GetRelationResponseBodyDataResellerProductRuleDO) SetMultiOrder(v bool) *GetRelationResponseBodyDataResellerProductRuleDO {
	s.MultiOrder = &v
	return s
}

func (s *GetRelationResponseBodyDataResellerProductRuleDO) SetPayMode(v string) *GetRelationResponseBodyDataResellerProductRuleDO {
	s.PayMode = &v
	return s
}

type GetRelationResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelationResponse) GoString() string {
	return s.String()
}

func (s *GetRelationResponse) SetHeaders(v map[string]*string) *GetRelationResponse {
	s.Headers = v
	return s
}

func (s *GetRelationResponse) SetStatusCode(v int32) *GetRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelationResponse) SetBody(v *GetRelationResponseBody) *GetRelationResponse {
	s.Body = v
	return s
}

type GetResellerPayOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Uid     *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetResellerPayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResellerPayOrderRequest) GoString() string {
	return s.String()
}

func (s *GetResellerPayOrderRequest) SetOrderId(v string) *GetResellerPayOrderRequest {
	s.OrderId = &v
	return s
}

func (s *GetResellerPayOrderRequest) SetUid(v int64) *GetResellerPayOrderRequest {
	s.Uid = &v
	return s
}

type GetResellerPayOrderResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetResellerPayOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResellerPayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResellerPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetResellerPayOrderResponseBody) SetCode(v string) *GetResellerPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *GetResellerPayOrderResponseBody) SetData(v *GetResellerPayOrderResponseBodyData) *GetResellerPayOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetResellerPayOrderResponseBody) SetMessage(v string) *GetResellerPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *GetResellerPayOrderResponseBody) SetRequestId(v string) *GetResellerPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResellerPayOrderResponseBody) SetSuccess(v bool) *GetResellerPayOrderResponseBody {
	s.Success = &v
	return s
}

type GetResellerPayOrderResponseBodyData struct {
	BuyerId     *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	PayAmount   *string `json:"PayAmount,omitempty" xml:"PayAmount,omitempty"`
}

func (s GetResellerPayOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResellerPayOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResellerPayOrderResponseBodyData) SetBuyerId(v string) *GetResellerPayOrderResponseBodyData {
	s.BuyerId = &v
	return s
}

func (s *GetResellerPayOrderResponseBodyData) SetOrderId(v string) *GetResellerPayOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetResellerPayOrderResponseBodyData) SetOrderStatus(v string) *GetResellerPayOrderResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetResellerPayOrderResponseBodyData) SetPayAmount(v string) *GetResellerPayOrderResponseBodyData {
	s.PayAmount = &v
	return s
}

type GetResellerPayOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResellerPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResellerPayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResellerPayOrderResponse) GoString() string {
	return s.String()
}

func (s *GetResellerPayOrderResponse) SetHeaders(v map[string]*string) *GetResellerPayOrderResponse {
	s.Headers = v
	return s
}

func (s *GetResellerPayOrderResponse) SetStatusCode(v int32) *GetResellerPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResellerPayOrderResponse) SetBody(v *GetResellerPayOrderResponseBody) *GetResellerPayOrderResponse {
	s.Body = v
	return s
}

type LabelPartnerUserRequest struct {
	UserPK  *int64  `json:"UserPK,omitempty" xml:"UserPK,omitempty"`
	UserTag *string `json:"UserTag,omitempty" xml:"UserTag,omitempty"`
}

func (s LabelPartnerUserRequest) String() string {
	return tea.Prettify(s)
}

func (s LabelPartnerUserRequest) GoString() string {
	return s.String()
}

func (s *LabelPartnerUserRequest) SetUserPK(v int64) *LabelPartnerUserRequest {
	s.UserPK = &v
	return s
}

func (s *LabelPartnerUserRequest) SetUserTag(v string) *LabelPartnerUserRequest {
	s.UserTag = &v
	return s
}

type LabelPartnerUserResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LabelPartnerUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LabelPartnerUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LabelPartnerUserResponseBody) GoString() string {
	return s.String()
}

func (s *LabelPartnerUserResponseBody) SetCode(v string) *LabelPartnerUserResponseBody {
	s.Code = &v
	return s
}

func (s *LabelPartnerUserResponseBody) SetData(v *LabelPartnerUserResponseBodyData) *LabelPartnerUserResponseBody {
	s.Data = v
	return s
}

func (s *LabelPartnerUserResponseBody) SetMessage(v string) *LabelPartnerUserResponseBody {
	s.Message = &v
	return s
}

func (s *LabelPartnerUserResponseBody) SetRequestId(v string) *LabelPartnerUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *LabelPartnerUserResponseBody) SetSuccess(v bool) *LabelPartnerUserResponseBody {
	s.Success = &v
	return s
}

type LabelPartnerUserResponseBodyData struct {
	UserPK  *int64  `json:"UserPK,omitempty" xml:"UserPK,omitempty"`
	UserTag *string `json:"UserTag,omitempty" xml:"UserTag,omitempty"`
}

func (s LabelPartnerUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LabelPartnerUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *LabelPartnerUserResponseBodyData) SetUserPK(v int64) *LabelPartnerUserResponseBodyData {
	s.UserPK = &v
	return s
}

func (s *LabelPartnerUserResponseBodyData) SetUserTag(v string) *LabelPartnerUserResponseBodyData {
	s.UserTag = &v
	return s
}

type LabelPartnerUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LabelPartnerUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LabelPartnerUserResponse) String() string {
	return tea.Prettify(s)
}

func (s LabelPartnerUserResponse) GoString() string {
	return s.String()
}

func (s *LabelPartnerUserResponse) SetHeaders(v map[string]*string) *LabelPartnerUserResponse {
	s.Headers = v
	return s
}

func (s *LabelPartnerUserResponse) SetStatusCode(v int32) *LabelPartnerUserResponse {
	s.StatusCode = &v
	return s
}

func (s *LabelPartnerUserResponse) SetBody(v *LabelPartnerUserResponseBody) *LabelPartnerUserResponse {
	s.Body = v
	return s
}

type MigrateResourceRequest struct {
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s MigrateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateResourceRequest) GoString() string {
	return s.String()
}

func (s *MigrateResourceRequest) SetActionCode(v string) *MigrateResourceRequest {
	s.ActionCode = &v
	return s
}

func (s *MigrateResourceRequest) SetContent(v string) *MigrateResourceRequest {
	s.Content = &v
	return s
}

type MigrateResourceResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MigrateResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MigrateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateResourceResponseBody) SetCode(v string) *MigrateResourceResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateResourceResponseBody) SetData(v *MigrateResourceResponseBodyData) *MigrateResourceResponseBody {
	s.Data = v
	return s
}

func (s *MigrateResourceResponseBody) SetMessage(v string) *MigrateResourceResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateResourceResponseBody) SetRequestId(v string) *MigrateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateResourceResponseBody) SetSuccess(v bool) *MigrateResourceResponseBody {
	s.Success = &v
	return s
}

type MigrateResourceResponseBodyData struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProcEnvir *string `json:"ProcEnvir,omitempty" xml:"ProcEnvir,omitempty"`
}

func (s MigrateResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MigrateResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MigrateResourceResponseBodyData) SetContent(v string) *MigrateResourceResponseBodyData {
	s.Content = &v
	return s
}

func (s *MigrateResourceResponseBodyData) SetProcEnvir(v string) *MigrateResourceResponseBodyData {
	s.ProcEnvir = &v
	return s
}

type MigrateResourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateResourceResponse) GoString() string {
	return s.String()
}

func (s *MigrateResourceResponse) SetHeaders(v map[string]*string) *MigrateResourceResponse {
	s.Headers = v
	return s
}

func (s *MigrateResourceResponse) SetStatusCode(v int32) *MigrateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateResourceResponse) SetBody(v *MigrateResourceResponseBody) *MigrateResourceResponse {
	s.Body = v
	return s
}

type OfflineActivityRequest struct {
	ActivityList []*int64 `json:"ActivityList,omitempty" xml:"ActivityList,omitempty" type:"Repeated"`
	BizId        *string  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Token        *string  `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s OfflineActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineActivityRequest) GoString() string {
	return s.String()
}

func (s *OfflineActivityRequest) SetActivityList(v []*int64) *OfflineActivityRequest {
	s.ActivityList = v
	return s
}

func (s *OfflineActivityRequest) SetBizId(v string) *OfflineActivityRequest {
	s.BizId = &v
	return s
}

func (s *OfflineActivityRequest) SetToken(v string) *OfflineActivityRequest {
	s.Token = &v
	return s
}

type OfflineActivityResponseBody struct {
	Code       *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{} `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Data       []*int64               `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineActivityResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineActivityResponseBody) SetCode(v string) *OfflineActivityResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineActivityResponseBody) SetContextMap(v map[string]interface{}) *OfflineActivityResponseBody {
	s.ContextMap = v
	return s
}

func (s *OfflineActivityResponseBody) SetData(v []*int64) *OfflineActivityResponseBody {
	s.Data = v
	return s
}

func (s *OfflineActivityResponseBody) SetMessage(v string) *OfflineActivityResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineActivityResponseBody) SetRequestId(v string) *OfflineActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineActivityResponseBody) SetSuccess(v bool) *OfflineActivityResponseBody {
	s.Success = &v
	return s
}

type OfflineActivityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineActivityResponse) GoString() string {
	return s.String()
}

func (s *OfflineActivityResponse) SetHeaders(v map[string]*string) *OfflineActivityResponse {
	s.Headers = v
	return s
}

func (s *OfflineActivityResponse) SetStatusCode(v int32) *OfflineActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineActivityResponse) SetBody(v *OfflineActivityResponseBody) *OfflineActivityResponse {
	s.Body = v
	return s
}

type PayResultCallbackRequest struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PayStatus *string `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	PayTime   *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	Uid       *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s PayResultCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s PayResultCallbackRequest) GoString() string {
	return s.String()
}

func (s *PayResultCallbackRequest) SetOrderId(v string) *PayResultCallbackRequest {
	s.OrderId = &v
	return s
}

func (s *PayResultCallbackRequest) SetPayStatus(v string) *PayResultCallbackRequest {
	s.PayStatus = &v
	return s
}

func (s *PayResultCallbackRequest) SetPayTime(v string) *PayResultCallbackRequest {
	s.PayTime = &v
	return s
}

func (s *PayResultCallbackRequest) SetUid(v int64) *PayResultCallbackRequest {
	s.Uid = &v
	return s
}

type PayResultCallbackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PayResultCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PayResultCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *PayResultCallbackResponseBody) SetCode(v string) *PayResultCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *PayResultCallbackResponseBody) SetData(v bool) *PayResultCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *PayResultCallbackResponseBody) SetMessage(v string) *PayResultCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *PayResultCallbackResponseBody) SetRequestId(v string) *PayResultCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *PayResultCallbackResponseBody) SetSuccess(v bool) *PayResultCallbackResponseBody {
	s.Success = &v
	return s
}

type PayResultCallbackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PayResultCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PayResultCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s PayResultCallbackResponse) GoString() string {
	return s.String()
}

func (s *PayResultCallbackResponse) SetHeaders(v map[string]*string) *PayResultCallbackResponse {
	s.Headers = v
	return s
}

func (s *PayResultCallbackResponse) SetStatusCode(v int32) *PayResultCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *PayResultCallbackResponse) SetBody(v *PayResultCallbackResponseBody) *PayResultCallbackResponse {
	s.Body = v
	return s
}

type PublicActivityRequest struct {
	ActivityList []*int64 `json:"ActivityList,omitempty" xml:"ActivityList,omitempty" type:"Repeated"`
	BizId        *string  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SnapType     *string  `json:"SnapType,omitempty" xml:"SnapType,omitempty"`
	Token        *string  `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s PublicActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s PublicActivityRequest) GoString() string {
	return s.String()
}

func (s *PublicActivityRequest) SetActivityList(v []*int64) *PublicActivityRequest {
	s.ActivityList = v
	return s
}

func (s *PublicActivityRequest) SetBizId(v string) *PublicActivityRequest {
	s.BizId = &v
	return s
}

func (s *PublicActivityRequest) SetSnapType(v string) *PublicActivityRequest {
	s.SnapType = &v
	return s
}

func (s *PublicActivityRequest) SetToken(v string) *PublicActivityRequest {
	s.Token = &v
	return s
}

type PublicActivityResponseBody struct {
	Code       *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{} `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Data       []*int64               `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublicActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublicActivityResponseBody) GoString() string {
	return s.String()
}

func (s *PublicActivityResponseBody) SetCode(v string) *PublicActivityResponseBody {
	s.Code = &v
	return s
}

func (s *PublicActivityResponseBody) SetContextMap(v map[string]interface{}) *PublicActivityResponseBody {
	s.ContextMap = v
	return s
}

func (s *PublicActivityResponseBody) SetData(v []*int64) *PublicActivityResponseBody {
	s.Data = v
	return s
}

func (s *PublicActivityResponseBody) SetMessage(v string) *PublicActivityResponseBody {
	s.Message = &v
	return s
}

func (s *PublicActivityResponseBody) SetRequestId(v string) *PublicActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublicActivityResponseBody) SetSuccess(v bool) *PublicActivityResponseBody {
	s.Success = &v
	return s
}

type PublicActivityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublicActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublicActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s PublicActivityResponse) GoString() string {
	return s.String()
}

func (s *PublicActivityResponse) SetHeaders(v map[string]*string) *PublicActivityResponse {
	s.Headers = v
	return s
}

func (s *PublicActivityResponse) SetStatusCode(v int32) *PublicActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *PublicActivityResponse) SetBody(v *PublicActivityResponseBody) *PublicActivityResponse {
	s.Body = v
	return s
}

type QueryActivityRequest struct {
	ActivityId *int64  `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SnapType   *string `json:"SnapType,omitempty" xml:"SnapType,omitempty"`
}

func (s QueryActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityRequest) GoString() string {
	return s.String()
}

func (s *QueryActivityRequest) SetActivityId(v int64) *QueryActivityRequest {
	s.ActivityId = &v
	return s
}

func (s *QueryActivityRequest) SetBizId(v string) *QueryActivityRequest {
	s.BizId = &v
	return s
}

func (s *QueryActivityRequest) SetSnapType(v string) *QueryActivityRequest {
	s.SnapType = &v
	return s
}

type QueryActivityResponseBody struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{}         `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Data       *QueryActivityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActivityResponseBody) SetCode(v string) *QueryActivityResponseBody {
	s.Code = &v
	return s
}

func (s *QueryActivityResponseBody) SetContextMap(v map[string]interface{}) *QueryActivityResponseBody {
	s.ContextMap = v
	return s
}

func (s *QueryActivityResponseBody) SetData(v *QueryActivityResponseBodyData) *QueryActivityResponseBody {
	s.Data = v
	return s
}

func (s *QueryActivityResponseBody) SetMessage(v string) *QueryActivityResponseBody {
	s.Message = &v
	return s
}

func (s *QueryActivityResponseBody) SetRequestId(v string) *QueryActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryActivityResponseBody) SetSuccess(v bool) *QueryActivityResponseBody {
	s.Success = &v
	return s
}

type QueryActivityResponseBodyData struct {
	ActivityId           *int64             `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ActivityName         *string            `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	BlackList            []*int64           `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	CommodityCodeList    []*string          `json:"CommodityCodeList,omitempty" xml:"CommodityCodeList,omitempty" type:"Repeated"`
	Description          *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime              *string            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtMap               map[string]*string `json:"ExtMap,omitempty" xml:"ExtMap,omitempty"`
	PromotionDescription *string            `json:"PromotionDescription,omitempty" xml:"PromotionDescription,omitempty"`
	PromotionValue       *float64           `json:"PromotionValue,omitempty" xml:"PromotionValue,omitempty"`
	RegionList           []*string          `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	SellerIdList         []*int64           `json:"SellerIdList,omitempty" xml:"SellerIdList,omitempty" type:"Repeated"`
	StartTime            *string            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	TestAccountList      []*int64           `json:"TestAccountList,omitempty" xml:"TestAccountList,omitempty" type:"Repeated"`
	WhiteList            []*int64           `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s QueryActivityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryActivityResponseBodyData) SetActivityId(v int64) *QueryActivityResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetActivityName(v string) *QueryActivityResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetBlackList(v []*int64) *QueryActivityResponseBodyData {
	s.BlackList = v
	return s
}

func (s *QueryActivityResponseBodyData) SetCommodityCodeList(v []*string) *QueryActivityResponseBodyData {
	s.CommodityCodeList = v
	return s
}

func (s *QueryActivityResponseBodyData) SetDescription(v string) *QueryActivityResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetEndTime(v string) *QueryActivityResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetExtMap(v map[string]*string) *QueryActivityResponseBodyData {
	s.ExtMap = v
	return s
}

func (s *QueryActivityResponseBodyData) SetPromotionDescription(v string) *QueryActivityResponseBodyData {
	s.PromotionDescription = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetPromotionValue(v float64) *QueryActivityResponseBodyData {
	s.PromotionValue = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetRegionList(v []*string) *QueryActivityResponseBodyData {
	s.RegionList = v
	return s
}

func (s *QueryActivityResponseBodyData) SetSellerIdList(v []*int64) *QueryActivityResponseBodyData {
	s.SellerIdList = v
	return s
}

func (s *QueryActivityResponseBodyData) SetStartTime(v string) *QueryActivityResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetStatus(v string) *QueryActivityResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryActivityResponseBodyData) SetTestAccountList(v []*int64) *QueryActivityResponseBodyData {
	s.TestAccountList = v
	return s
}

func (s *QueryActivityResponseBodyData) SetWhiteList(v []*int64) *QueryActivityResponseBodyData {
	s.WhiteList = v
	return s
}

type QueryActivityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityResponse) GoString() string {
	return s.String()
}

func (s *QueryActivityResponse) SetHeaders(v map[string]*string) *QueryActivityResponse {
	s.Headers = v
	return s
}

func (s *QueryActivityResponse) SetStatusCode(v int32) *QueryActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryActivityResponse) SetBody(v *QueryActivityResponseBody) *QueryActivityResponse {
	s.Body = v
	return s
}

type QueryEcoRelationRequest struct {
	RelationTime *string `json:"RelationTime,omitempty" xml:"RelationTime,omitempty"`
	Uid          *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryEcoRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEcoRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryEcoRelationRequest) SetRelationTime(v string) *QueryEcoRelationRequest {
	s.RelationTime = &v
	return s
}

func (s *QueryEcoRelationRequest) SetUid(v int64) *QueryEcoRelationRequest {
	s.Uid = &v
	return s
}

type QueryEcoRelationResponseBody struct {
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{}            `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Data       *QueryEcoRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEcoRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEcoRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEcoRelationResponseBody) SetCode(v string) *QueryEcoRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEcoRelationResponseBody) SetContextMap(v map[string]interface{}) *QueryEcoRelationResponseBody {
	s.ContextMap = v
	return s
}

func (s *QueryEcoRelationResponseBody) SetData(v *QueryEcoRelationResponseBodyData) *QueryEcoRelationResponseBody {
	s.Data = v
	return s
}

func (s *QueryEcoRelationResponseBody) SetMessage(v string) *QueryEcoRelationResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEcoRelationResponseBody) SetRequestId(v string) *QueryEcoRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEcoRelationResponseBody) SetSuccess(v bool) *QueryEcoRelationResponseBody {
	s.Success = &v
	return s
}

type QueryEcoRelationResponseBodyData struct {
	CanLoginOfficial      *bool                                                  `json:"CanLoginOfficial,omitempty" xml:"CanLoginOfficial,omitempty"`
	EndTime               *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsTopReseller         *bool                                                  `json:"IsTopReseller,omitempty" xml:"IsTopReseller,omitempty"`
	ResellerProductModeDO *QueryEcoRelationResponseBodyDataResellerProductModeDO `json:"ResellerProductModeDO,omitempty" xml:"ResellerProductModeDO,omitempty" type:"Struct"`
	ResellerProductRuleDO *QueryEcoRelationResponseBodyDataResellerProductRuleDO `json:"ResellerProductRuleDO,omitempty" xml:"ResellerProductRuleDO,omitempty" type:"Struct"`
	ResellerUid           *int64                                                 `json:"ResellerUid,omitempty" xml:"ResellerUid,omitempty"`
	StartTime             *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Uid                   *int64                                                 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserType              *string                                                `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s QueryEcoRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEcoRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEcoRelationResponseBodyData) SetCanLoginOfficial(v bool) *QueryEcoRelationResponseBodyData {
	s.CanLoginOfficial = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetEndTime(v string) *QueryEcoRelationResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetIsTopReseller(v bool) *QueryEcoRelationResponseBodyData {
	s.IsTopReseller = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetResellerProductModeDO(v *QueryEcoRelationResponseBodyDataResellerProductModeDO) *QueryEcoRelationResponseBodyData {
	s.ResellerProductModeDO = v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetResellerProductRuleDO(v *QueryEcoRelationResponseBodyDataResellerProductRuleDO) *QueryEcoRelationResponseBodyData {
	s.ResellerProductRuleDO = v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetResellerUid(v int64) *QueryEcoRelationResponseBodyData {
	s.ResellerUid = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetStartTime(v string) *QueryEcoRelationResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetStatus(v string) *QueryEcoRelationResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetUid(v int64) *QueryEcoRelationResponseBodyData {
	s.Uid = &v
	return s
}

func (s *QueryEcoRelationResponseBodyData) SetUserType(v string) *QueryEcoRelationResponseBodyData {
	s.UserType = &v
	return s
}

type QueryEcoRelationResponseBodyDataResellerProductModeDO struct {
	IsService   *int32  `json:"IsService,omitempty" xml:"IsService,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s QueryEcoRelationResponseBodyDataResellerProductModeDO) String() string {
	return tea.Prettify(s)
}

func (s QueryEcoRelationResponseBodyDataResellerProductModeDO) GoString() string {
	return s.String()
}

func (s *QueryEcoRelationResponseBodyDataResellerProductModeDO) SetIsService(v int32) *QueryEcoRelationResponseBodyDataResellerProductModeDO {
	s.IsService = &v
	return s
}

func (s *QueryEcoRelationResponseBodyDataResellerProductModeDO) SetProductCode(v string) *QueryEcoRelationResponseBodyDataResellerProductModeDO {
	s.ProductCode = &v
	return s
}

func (s *QueryEcoRelationResponseBodyDataResellerProductModeDO) SetProductName(v string) *QueryEcoRelationResponseBodyDataResellerProductModeDO {
	s.ProductName = &v
	return s
}

type QueryEcoRelationResponseBodyDataResellerProductRuleDO struct {
	CommodityRoute *bool   `json:"CommodityRoute,omitempty" xml:"CommodityRoute,omitempty"`
	MultiOrder     *bool   `json:"MultiOrder,omitempty" xml:"MultiOrder,omitempty"`
	PayMode        *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
}

func (s QueryEcoRelationResponseBodyDataResellerProductRuleDO) String() string {
	return tea.Prettify(s)
}

func (s QueryEcoRelationResponseBodyDataResellerProductRuleDO) GoString() string {
	return s.String()
}

func (s *QueryEcoRelationResponseBodyDataResellerProductRuleDO) SetCommodityRoute(v bool) *QueryEcoRelationResponseBodyDataResellerProductRuleDO {
	s.CommodityRoute = &v
	return s
}

func (s *QueryEcoRelationResponseBodyDataResellerProductRuleDO) SetMultiOrder(v bool) *QueryEcoRelationResponseBodyDataResellerProductRuleDO {
	s.MultiOrder = &v
	return s
}

func (s *QueryEcoRelationResponseBodyDataResellerProductRuleDO) SetPayMode(v string) *QueryEcoRelationResponseBodyDataResellerProductRuleDO {
	s.PayMode = &v
	return s
}

type QueryEcoRelationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEcoRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEcoRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEcoRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryEcoRelationResponse) SetHeaders(v map[string]*string) *QueryEcoRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryEcoRelationResponse) SetStatusCode(v int32) *QueryEcoRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEcoRelationResponse) SetBody(v *QueryEcoRelationResponseBody) *QueryEcoRelationResponse {
	s.Body = v
	return s
}

type SaveActivityRequest struct {
	ActivityName             *string                                        `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	BizId                    *string                                        `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BlackUidList             []*int64                                       `json:"BlackUidList,omitempty" xml:"BlackUidList,omitempty" type:"Repeated"`
	Description              *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                  *string                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtendMap                map[string]interface{}                         `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	FusionPromotionParamList []*SaveActivityRequestFusionPromotionParamList `json:"FusionPromotionParamList,omitempty" xml:"FusionPromotionParamList,omitempty" type:"Repeated"`
	PublishTag               *string                                        `json:"PublishTag,omitempty" xml:"PublishTag,omitempty"`
	StartTime                *string                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TestAccountUidList       []*int64                                       `json:"TestAccountUidList,omitempty" xml:"TestAccountUidList,omitempty" type:"Repeated"`
	Token                    *string                                        `json:"Token,omitempty" xml:"Token,omitempty"`
	WhiteUidList             []*int64                                       `json:"WhiteUidList,omitempty" xml:"WhiteUidList,omitempty" type:"Repeated"`
}

func (s SaveActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityRequest) GoString() string {
	return s.String()
}

func (s *SaveActivityRequest) SetActivityName(v string) *SaveActivityRequest {
	s.ActivityName = &v
	return s
}

func (s *SaveActivityRequest) SetBizId(v string) *SaveActivityRequest {
	s.BizId = &v
	return s
}

func (s *SaveActivityRequest) SetBlackUidList(v []*int64) *SaveActivityRequest {
	s.BlackUidList = v
	return s
}

func (s *SaveActivityRequest) SetDescription(v string) *SaveActivityRequest {
	s.Description = &v
	return s
}

func (s *SaveActivityRequest) SetEndTime(v string) *SaveActivityRequest {
	s.EndTime = &v
	return s
}

func (s *SaveActivityRequest) SetExtendMap(v map[string]interface{}) *SaveActivityRequest {
	s.ExtendMap = v
	return s
}

func (s *SaveActivityRequest) SetFusionPromotionParamList(v []*SaveActivityRequestFusionPromotionParamList) *SaveActivityRequest {
	s.FusionPromotionParamList = v
	return s
}

func (s *SaveActivityRequest) SetPublishTag(v string) *SaveActivityRequest {
	s.PublishTag = &v
	return s
}

func (s *SaveActivityRequest) SetStartTime(v string) *SaveActivityRequest {
	s.StartTime = &v
	return s
}

func (s *SaveActivityRequest) SetTestAccountUidList(v []*int64) *SaveActivityRequest {
	s.TestAccountUidList = v
	return s
}

func (s *SaveActivityRequest) SetToken(v string) *SaveActivityRequest {
	s.Token = &v
	return s
}

func (s *SaveActivityRequest) SetWhiteUidList(v []*int64) *SaveActivityRequest {
	s.WhiteUidList = v
	return s
}

type SaveActivityRequestFusionPromotionParamList struct {
	CommodityCodeList    []*string                                                         `json:"CommodityCodeList,omitempty" xml:"CommodityCodeList,omitempty" type:"Repeated"`
	PromotionValue       *float64                                                          `json:"PromotionValue,omitempty" xml:"PromotionValue,omitempty"`
	RestrictedRegionList []*string                                                         `json:"RestrictedRegionList,omitempty" xml:"RestrictedRegionList,omitempty" type:"Repeated"`
	ModuleInfoParamList  []*SaveActivityRequestFusionPromotionParamListModuleInfoParamList `json:"moduleInfoParamList,omitempty" xml:"moduleInfoParamList,omitempty" type:"Repeated"`
}

func (s SaveActivityRequestFusionPromotionParamList) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityRequestFusionPromotionParamList) GoString() string {
	return s.String()
}

func (s *SaveActivityRequestFusionPromotionParamList) SetCommodityCodeList(v []*string) *SaveActivityRequestFusionPromotionParamList {
	s.CommodityCodeList = v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamList) SetPromotionValue(v float64) *SaveActivityRequestFusionPromotionParamList {
	s.PromotionValue = &v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamList) SetRestrictedRegionList(v []*string) *SaveActivityRequestFusionPromotionParamList {
	s.RestrictedRegionList = v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamList) SetModuleInfoParamList(v []*SaveActivityRequestFusionPromotionParamListModuleInfoParamList) *SaveActivityRequestFusionPromotionParamList {
	s.ModuleInfoParamList = v
	return s
}

type SaveActivityRequestFusionPromotionParamListModuleInfoParamList struct {
	ComponentCode   *string   `json:"componentCode,omitempty" xml:"componentCode,omitempty"`
	ItemCode        *string   `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	ModuleId        *int64    `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleKey       *string   `json:"moduleKey,omitempty" xml:"moduleKey,omitempty"`
	ModuleName      *string   `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	ModuleValueList []*string `json:"moduleValueList,omitempty" xml:"moduleValueList,omitempty" type:"Repeated"`
	PricePlanId     *int64    `json:"pricePlanId,omitempty" xml:"pricePlanId,omitempty"`
}

func (s SaveActivityRequestFusionPromotionParamListModuleInfoParamList) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityRequestFusionPromotionParamListModuleInfoParamList) GoString() string {
	return s.String()
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetComponentCode(v string) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.ComponentCode = &v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetItemCode(v string) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.ItemCode = &v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetModuleId(v int64) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleId = &v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetModuleKey(v string) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleKey = &v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetModuleName(v string) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleName = &v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetModuleValueList(v []*string) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleValueList = v
	return s
}

func (s *SaveActivityRequestFusionPromotionParamListModuleInfoParamList) SetPricePlanId(v int64) *SaveActivityRequestFusionPromotionParamListModuleInfoParamList {
	s.PricePlanId = &v
	return s
}

type SaveActivityShrinkRequest struct {
	ActivityName             *string                                              `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	BizId                    *string                                              `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BlackUidList             []*int64                                             `json:"BlackUidList,omitempty" xml:"BlackUidList,omitempty" type:"Repeated"`
	Description              *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                  *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtendMapShrink          *string                                              `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	FusionPromotionParamList []*SaveActivityShrinkRequestFusionPromotionParamList `json:"FusionPromotionParamList,omitempty" xml:"FusionPromotionParamList,omitempty" type:"Repeated"`
	PublishTag               *string                                              `json:"PublishTag,omitempty" xml:"PublishTag,omitempty"`
	StartTime                *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TestAccountUidList       []*int64                                             `json:"TestAccountUidList,omitempty" xml:"TestAccountUidList,omitempty" type:"Repeated"`
	Token                    *string                                              `json:"Token,omitempty" xml:"Token,omitempty"`
	WhiteUidList             []*int64                                             `json:"WhiteUidList,omitempty" xml:"WhiteUidList,omitempty" type:"Repeated"`
}

func (s SaveActivityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveActivityShrinkRequest) SetActivityName(v string) *SaveActivityShrinkRequest {
	s.ActivityName = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetBizId(v string) *SaveActivityShrinkRequest {
	s.BizId = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetBlackUidList(v []*int64) *SaveActivityShrinkRequest {
	s.BlackUidList = v
	return s
}

func (s *SaveActivityShrinkRequest) SetDescription(v string) *SaveActivityShrinkRequest {
	s.Description = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetEndTime(v string) *SaveActivityShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetExtendMapShrink(v string) *SaveActivityShrinkRequest {
	s.ExtendMapShrink = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetFusionPromotionParamList(v []*SaveActivityShrinkRequestFusionPromotionParamList) *SaveActivityShrinkRequest {
	s.FusionPromotionParamList = v
	return s
}

func (s *SaveActivityShrinkRequest) SetPublishTag(v string) *SaveActivityShrinkRequest {
	s.PublishTag = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetStartTime(v string) *SaveActivityShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetTestAccountUidList(v []*int64) *SaveActivityShrinkRequest {
	s.TestAccountUidList = v
	return s
}

func (s *SaveActivityShrinkRequest) SetToken(v string) *SaveActivityShrinkRequest {
	s.Token = &v
	return s
}

func (s *SaveActivityShrinkRequest) SetWhiteUidList(v []*int64) *SaveActivityShrinkRequest {
	s.WhiteUidList = v
	return s
}

type SaveActivityShrinkRequestFusionPromotionParamList struct {
	CommodityCodeList    []*string                                                               `json:"CommodityCodeList,omitempty" xml:"CommodityCodeList,omitempty" type:"Repeated"`
	PromotionValue       *float64                                                                `json:"PromotionValue,omitempty" xml:"PromotionValue,omitempty"`
	RestrictedRegionList []*string                                                               `json:"RestrictedRegionList,omitempty" xml:"RestrictedRegionList,omitempty" type:"Repeated"`
	ModuleInfoParamList  []*SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList `json:"moduleInfoParamList,omitempty" xml:"moduleInfoParamList,omitempty" type:"Repeated"`
}

func (s SaveActivityShrinkRequestFusionPromotionParamList) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityShrinkRequestFusionPromotionParamList) GoString() string {
	return s.String()
}

func (s *SaveActivityShrinkRequestFusionPromotionParamList) SetCommodityCodeList(v []*string) *SaveActivityShrinkRequestFusionPromotionParamList {
	s.CommodityCodeList = v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamList) SetPromotionValue(v float64) *SaveActivityShrinkRequestFusionPromotionParamList {
	s.PromotionValue = &v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamList) SetRestrictedRegionList(v []*string) *SaveActivityShrinkRequestFusionPromotionParamList {
	s.RestrictedRegionList = v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamList) SetModuleInfoParamList(v []*SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) *SaveActivityShrinkRequestFusionPromotionParamList {
	s.ModuleInfoParamList = v
	return s
}

type SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList struct {
	ComponentCode   *string   `json:"componentCode,omitempty" xml:"componentCode,omitempty"`
	ItemCode        *string   `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	ModuleId        *int64    `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleKey       *string   `json:"moduleKey,omitempty" xml:"moduleKey,omitempty"`
	ModuleName      *string   `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	ModuleValueList []*string `json:"moduleValueList,omitempty" xml:"moduleValueList,omitempty" type:"Repeated"`
	PricePlanId     *int64    `json:"pricePlanId,omitempty" xml:"pricePlanId,omitempty"`
}

func (s SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) GoString() string {
	return s.String()
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetComponentCode(v string) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.ComponentCode = &v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetItemCode(v string) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.ItemCode = &v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetModuleId(v int64) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleId = &v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetModuleKey(v string) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleKey = &v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetModuleName(v string) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleName = &v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetModuleValueList(v []*string) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.ModuleValueList = v
	return s
}

func (s *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList) SetPricePlanId(v int64) *SaveActivityShrinkRequestFusionPromotionParamListModuleInfoParamList {
	s.PricePlanId = &v
	return s
}

type SaveActivityResponseBody struct {
	Code       *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{} `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Data       []*int64               `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityResponseBody) GoString() string {
	return s.String()
}

func (s *SaveActivityResponseBody) SetCode(v string) *SaveActivityResponseBody {
	s.Code = &v
	return s
}

func (s *SaveActivityResponseBody) SetContextMap(v map[string]interface{}) *SaveActivityResponseBody {
	s.ContextMap = v
	return s
}

func (s *SaveActivityResponseBody) SetData(v []*int64) *SaveActivityResponseBody {
	s.Data = v
	return s
}

func (s *SaveActivityResponseBody) SetMessage(v string) *SaveActivityResponseBody {
	s.Message = &v
	return s
}

func (s *SaveActivityResponseBody) SetRequestId(v string) *SaveActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveActivityResponseBody) SetSuccess(v bool) *SaveActivityResponseBody {
	s.Success = &v
	return s
}

type SaveActivityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveActivityResponse) GoString() string {
	return s.String()
}

func (s *SaveActivityResponse) SetHeaders(v map[string]*string) *SaveActivityResponse {
	s.Headers = v
	return s
}

func (s *SaveActivityResponse) SetStatusCode(v int32) *SaveActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveActivityResponse) SetBody(v *SaveActivityResponseBody) *SaveActivityResponse {
	s.Body = v
	return s
}

type SendCouponRequest struct {
	Bid                 *int64                                  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	FromApp             *string                                 `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	Operator            *string                                 `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RequestId           *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SellerId            *int64                                  `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	TemplateId          *int64                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UserAmountModelList []*SendCouponRequestUserAmountModelList `json:"UserAmountModelList,omitempty" xml:"UserAmountModelList,omitempty" type:"Repeated"`
}

func (s SendCouponRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCouponRequest) GoString() string {
	return s.String()
}

func (s *SendCouponRequest) SetBid(v int64) *SendCouponRequest {
	s.Bid = &v
	return s
}

func (s *SendCouponRequest) SetFromApp(v string) *SendCouponRequest {
	s.FromApp = &v
	return s
}

func (s *SendCouponRequest) SetOperator(v string) *SendCouponRequest {
	s.Operator = &v
	return s
}

func (s *SendCouponRequest) SetRequestId(v string) *SendCouponRequest {
	s.RequestId = &v
	return s
}

func (s *SendCouponRequest) SetSellerId(v int64) *SendCouponRequest {
	s.SellerId = &v
	return s
}

func (s *SendCouponRequest) SetTemplateId(v int64) *SendCouponRequest {
	s.TemplateId = &v
	return s
}

func (s *SendCouponRequest) SetUserAmountModelList(v []*SendCouponRequestUserAmountModelList) *SendCouponRequest {
	s.UserAmountModelList = v
	return s
}

type SendCouponRequestUserAmountModelList struct {
	Amount   *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Uid      *int64   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId   *int64   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	YouhuiId *int64   `json:"YouhuiId,omitempty" xml:"YouhuiId,omitempty"`
}

func (s SendCouponRequestUserAmountModelList) String() string {
	return tea.Prettify(s)
}

func (s SendCouponRequestUserAmountModelList) GoString() string {
	return s.String()
}

func (s *SendCouponRequestUserAmountModelList) SetAmount(v float64) *SendCouponRequestUserAmountModelList {
	s.Amount = &v
	return s
}

func (s *SendCouponRequestUserAmountModelList) SetUid(v int64) *SendCouponRequestUserAmountModelList {
	s.Uid = &v
	return s
}

func (s *SendCouponRequestUserAmountModelList) SetUserId(v int64) *SendCouponRequestUserAmountModelList {
	s.UserId = &v
	return s
}

func (s *SendCouponRequestUserAmountModelList) SetYouhuiId(v int64) *SendCouponRequestUserAmountModelList {
	s.YouhuiId = &v
	return s
}

type SendCouponResponseBody struct {
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	ContextMap map[string]interface{}        `json:"ContextMap,omitempty" xml:"ContextMap,omitempty"`
	Count      *int32                        `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       []*SendCouponResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SendCouponResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCouponResponseBody) GoString() string {
	return s.String()
}

func (s *SendCouponResponseBody) SetCode(v string) *SendCouponResponseBody {
	s.Code = &v
	return s
}

func (s *SendCouponResponseBody) SetContextMap(v map[string]interface{}) *SendCouponResponseBody {
	s.ContextMap = v
	return s
}

func (s *SendCouponResponseBody) SetCount(v int32) *SendCouponResponseBody {
	s.Count = &v
	return s
}

func (s *SendCouponResponseBody) SetData(v []*SendCouponResponseBodyData) *SendCouponResponseBody {
	s.Data = v
	return s
}

func (s *SendCouponResponseBody) SetMessage(v string) *SendCouponResponseBody {
	s.Message = &v
	return s
}

func (s *SendCouponResponseBody) SetRequestId(v string) *SendCouponResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCouponResponseBody) SetSuccess(v bool) *SendCouponResponseBody {
	s.Success = &v
	return s
}

func (s *SendCouponResponseBody) SetTotalCount(v int32) *SendCouponResponseBody {
	s.TotalCount = &v
	return s
}

type SendCouponResponseBodyData struct {
	Amount   *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Uid      *int64   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId   *int64   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	YouhuiId *int64   `json:"YouhuiId,omitempty" xml:"YouhuiId,omitempty"`
}

func (s SendCouponResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendCouponResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCouponResponseBodyData) SetAmount(v float64) *SendCouponResponseBodyData {
	s.Amount = &v
	return s
}

func (s *SendCouponResponseBodyData) SetUid(v int64) *SendCouponResponseBodyData {
	s.Uid = &v
	return s
}

func (s *SendCouponResponseBodyData) SetUserId(v int64) *SendCouponResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SendCouponResponseBodyData) SetYouhuiId(v int64) *SendCouponResponseBodyData {
	s.YouhuiId = &v
	return s
}

type SendCouponResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCouponResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCouponResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCouponResponse) GoString() string {
	return s.String()
}

func (s *SendCouponResponse) SetHeaders(v map[string]*string) *SendCouponResponse {
	s.Headers = v
	return s
}

func (s *SendCouponResponse) SetStatusCode(v int32) *SendCouponResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCouponResponse) SetBody(v *SendCouponResponseBody) *SendCouponResponse {
	s.Body = v
	return s
}

type TransferResourceRequest struct {
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s TransferResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferResourceRequest) GoString() string {
	return s.String()
}

func (s *TransferResourceRequest) SetActionCode(v string) *TransferResourceRequest {
	s.ActionCode = &v
	return s
}

func (s *TransferResourceRequest) SetContent(v string) *TransferResourceRequest {
	s.Content = &v
	return s
}

type TransferResourceResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TransferResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferResourceResponseBody) GoString() string {
	return s.String()
}

func (s *TransferResourceResponseBody) SetCode(v string) *TransferResourceResponseBody {
	s.Code = &v
	return s
}

func (s *TransferResourceResponseBody) SetData(v *TransferResourceResponseBodyData) *TransferResourceResponseBody {
	s.Data = v
	return s
}

func (s *TransferResourceResponseBody) SetMessage(v string) *TransferResourceResponseBody {
	s.Message = &v
	return s
}

func (s *TransferResourceResponseBody) SetRequestId(v string) *TransferResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferResourceResponseBody) SetSuccess(v bool) *TransferResourceResponseBody {
	s.Success = &v
	return s
}

type TransferResourceResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProcEnv *string `json:"ProcEnv,omitempty" xml:"ProcEnv,omitempty"`
}

func (s TransferResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TransferResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransferResourceResponseBodyData) SetContent(v string) *TransferResourceResponseBodyData {
	s.Content = &v
	return s
}

func (s *TransferResourceResponseBodyData) SetProcEnv(v string) *TransferResourceResponseBodyData {
	s.ProcEnv = &v
	return s
}

type TransferResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferResourceResponse) GoString() string {
	return s.String()
}

func (s *TransferResourceResponse) SetHeaders(v map[string]*string) *TransferResourceResponse {
	s.Headers = v
	return s
}

func (s *TransferResourceResponse) SetStatusCode(v int32) *TransferResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferResourceResponse) SetBody(v *TransferResourceResponseBody) *TransferResourceResponse {
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
		"ap-northeast-1":              tea.String("resellertrade.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("resellertrade.aliyuncs.com"),
		"ap-south-1":                  tea.String("resellertrade.aliyuncs.com"),
		"ap-southeast-1":              tea.String("resellertrade.aliyuncs.com"),
		"ap-southeast-2":              tea.String("resellertrade.aliyuncs.com"),
		"ap-southeast-3":              tea.String("resellertrade.aliyuncs.com"),
		"ap-southeast-5":              tea.String("resellertrade.aliyuncs.com"),
		"cn-beijing":                  tea.String("resellertrade.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("resellertrade.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("resellertrade.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("resellertrade.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("resellertrade.aliyuncs.com"),
		"cn-chengdu":                  tea.String("resellertrade.aliyuncs.com"),
		"cn-edge-1":                   tea.String("resellertrade.aliyuncs.com"),
		"cn-fujian":                   tea.String("resellertrade.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("resellertrade.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("resellertrade.aliyuncs.com"),
		"cn-hongkong":                 tea.String("resellertrade.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("resellertrade.aliyuncs.com"),
		"cn-huhehaote":                tea.String("resellertrade.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("resellertrade.aliyuncs.com"),
		"cn-qingdao":                  tea.String("resellertrade.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("resellertrade.aliyuncs.com"),
		"cn-shanghai":                 tea.String("resellertrade.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("resellertrade.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("resellertrade.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("resellertrade.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("resellertrade.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("resellertrade.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("resellertrade.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("resellertrade.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("resellertrade.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("resellertrade.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("resellertrade.aliyuncs.com"),
		"cn-wuhan":                    tea.String("resellertrade.aliyuncs.com"),
		"cn-yushanfang":               tea.String("resellertrade.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("resellertrade.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("resellertrade.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("resellertrade.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("resellertrade.aliyuncs.com"),
		"eu-central-1":                tea.String("resellertrade.aliyuncs.com"),
		"eu-west-1":                   tea.String("resellertrade.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("resellertrade.aliyuncs.com"),
		"me-east-1":                   tea.String("resellertrade.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("resellertrade.aliyuncs.com"),
		"us-east-1":                   tea.String("resellertrade.aliyuncs.com"),
		"us-west-1":                   tea.String("resellertrade.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("resellertrade"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateCouponTemplateWithOptions(request *CreateCouponTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateCouponTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivitySite)) {
		body["ActivitySite"] = request.ActivitySite
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		body["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.CertainMoney)) {
		body["CertainMoney"] = request.CertainMoney
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		body["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityType)) {
		body["CommodityType"] = request.CommodityType
	}

	if !tea.BoolValue(util.IsUnset(request.ControlType)) {
		body["ControlType"] = request.ControlType
	}

	if !tea.BoolValue(util.IsUnset(request.CouponAmount)) {
		body["CouponAmount"] = request.CouponAmount
	}

	if !tea.BoolValue(util.IsUnset(request.CouponEndTime)) {
		body["CouponEndTime"] = request.CouponEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CouponFixedType)) {
		body["CouponFixedType"] = request.CouponFixedType
	}

	if !tea.BoolValue(util.IsUnset(request.CouponStartTime)) {
		body["CouponStartTime"] = request.CouponStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CouponType)) {
		body["CouponType"] = request.CouponType
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		bodyFlat["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DiscountRate)) {
		body["DiscountRate"] = request.DiscountRate
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendsMap)) {
		bodyFlat["ExtendsMap"] = request.ExtendsMap
	}

	if !tea.BoolValue(util.IsUnset(request.FromApp)) {
		body["FromApp"] = request.FromApp
	}

	if !tea.BoolValue(util.IsUnset(request.ItemCodeSet)) {
		body["ItemCodeSet"] = request.ItemCodeSet
	}

	if !tea.BoolValue(util.IsUnset(request.Market)) {
		body["Market"] = request.Market
	}

	if !tea.BoolValue(util.IsUnset(request.MarketType)) {
		body["MarketType"] = request.MarketType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRelease)) {
		body["MaxRelease"] = request.MaxRelease
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OrderTypeSet)) {
		body["OrderTypeSet"] = request.OrderTypeSet
	}

	if !tea.BoolValue(util.IsUnset(request.PerLimitNum)) {
		body["PerLimitNum"] = request.PerLimitNum
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		body["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.RelativeSecond)) {
		body["RelativeSecond"] = request.RelativeSecond
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SelectionIdSet)) {
		body["SelectionIdSet"] = request.SelectionIdSet
	}

	if !tea.BoolValue(util.IsUnset(request.SellerId)) {
		body["SellerId"] = request.SellerId
	}

	if !tea.BoolValue(util.IsUnset(request.Site)) {
		body["Site"] = request.Site
	}

	if !tea.BoolValue(util.IsUnset(request.SpId)) {
		body["SpId"] = request.SpId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UniversalType)) {
		body["UniversalType"] = request.UniversalType
	}

	if !tea.BoolValue(util.IsUnset(request.UpperLimit)) {
		body["UpperLimit"] = request.UpperLimit
	}

	if !tea.BoolValue(util.IsUnset(request.UsageCount)) {
		body["UsageCount"] = request.UsageCount
	}

	if !tea.BoolValue(util.IsUnset(request.UseScene)) {
		body["UseScene"] = request.UseScene
	}

	if !tea.BoolValue(util.IsUnset(request.UserPkAmount)) {
		body["UserPkAmount"] = request.UserPkAmount
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityType)) {
		body["ValidityType"] = request.ValidityType
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCouponTemplate"),
		Version:     tea.String("2019-12-27"),
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

func (client *Client) DiscardCouponListWithOptions(request *DiscardCouponListRequest, runtime *util.RuntimeOptions) (_result *DiscardCouponListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponList)) {
		body["CouponList"] = request.CouponList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DiscardCouponList"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DiscardCouponListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DiscardCouponList(request *DiscardCouponListRequest) (_result *DiscardCouponListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DiscardCouponListResponse{}
	_body, _err := client.DiscardCouponListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCouponPageWithOptions(request *GetCouponPageRequest, runtime *util.RuntimeOptions) (_result *GetCouponPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCouponPage"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCouponPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCouponPage(request *GetCouponPageRequest) (_result *GetCouponPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCouponPageResponse{}
	_body, _err := client.GetCouponPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomerListWithOptions(request *GetCustomerListRequest, runtime *util.RuntimeOptions) (_result *GetCustomerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomerList"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomerList(request *GetCustomerListRequest) (_result *GetCustomerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomerListResponse{}
	_body, _err := client.GetCustomerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRelationWithOptions(request *GetRelationRequest, runtime *util.RuntimeOptions) (_result *GetRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRelation"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRelation(request *GetRelationRequest) (_result *GetRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRelationResponse{}
	_body, _err := client.GetRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResellerPayOrderWithOptions(request *GetResellerPayOrderRequest, runtime *util.RuntimeOptions) (_result *GetResellerPayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResellerPayOrder"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResellerPayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResellerPayOrder(request *GetResellerPayOrderRequest) (_result *GetResellerPayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResellerPayOrderResponse{}
	_body, _err := client.GetResellerPayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LabelPartnerUserWithOptions(request *LabelPartnerUserRequest, runtime *util.RuntimeOptions) (_result *LabelPartnerUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPK)) {
		body["UserPK"] = request.UserPK
	}

	if !tea.BoolValue(util.IsUnset(request.UserTag)) {
		body["UserTag"] = request.UserTag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LabelPartnerUser"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LabelPartnerUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LabelPartnerUser(request *LabelPartnerUserRequest) (_result *LabelPartnerUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LabelPartnerUserResponse{}
	_body, _err := client.LabelPartnerUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateResourceWithOptions(request *MigrateResourceRequest, runtime *util.RuntimeOptions) (_result *MigrateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCode)) {
		body["ActionCode"] = request.ActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MigrateResource"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MigrateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateResource(request *MigrateResourceRequest) (_result *MigrateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateResourceResponse{}
	_body, _err := client.MigrateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OfflineActivityWithOptions(request *OfflineActivityRequest, runtime *util.RuntimeOptions) (_result *OfflineActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityList)) {
		body["ActivityList"] = request.ActivityList
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OfflineActivity"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OfflineActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OfflineActivity(request *OfflineActivityRequest) (_result *OfflineActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OfflineActivityResponse{}
	_body, _err := client.OfflineActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PayResultCallbackWithOptions(request *PayResultCallbackRequest, runtime *util.RuntimeOptions) (_result *PayResultCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayStatus)) {
		query["PayStatus"] = request.PayStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PayTime)) {
		query["PayTime"] = request.PayTime
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PayResultCallback"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PayResultCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PayResultCallback(request *PayResultCallbackRequest) (_result *PayResultCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PayResultCallbackResponse{}
	_body, _err := client.PayResultCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublicActivityWithOptions(request *PublicActivityRequest, runtime *util.RuntimeOptions) (_result *PublicActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityList)) {
		body["ActivityList"] = request.ActivityList
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapType)) {
		body["SnapType"] = request.SnapType
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublicActivity"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublicActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublicActivity(request *PublicActivityRequest) (_result *PublicActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublicActivityResponse{}
	_body, _err := client.PublicActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryActivityWithOptions(request *QueryActivityRequest, runtime *util.RuntimeOptions) (_result *QueryActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		body["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapType)) {
		body["SnapType"] = request.SnapType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryActivity"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryActivity(request *QueryActivityRequest) (_result *QueryActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryActivityResponse{}
	_body, _err := client.QueryActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEcoRelationWithOptions(request *QueryEcoRelationRequest, runtime *util.RuntimeOptions) (_result *QueryEcoRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationTime)) {
		body["RelationTime"] = request.RelationTime
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEcoRelation"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEcoRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEcoRelation(request *QueryEcoRelationRequest) (_result *QueryEcoRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEcoRelationResponse{}
	_body, _err := client.QueryEcoRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveActivityWithOptions(tmpReq *SaveActivityRequest, runtime *util.RuntimeOptions) (_result *SaveActivityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveActivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendMap)) {
		request.ExtendMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendMap, tea.String("ExtendMap"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityName)) {
		body["ActivityName"] = request.ActivityName
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BlackUidList)) {
		body["BlackUidList"] = request.BlackUidList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendMapShrink)) {
		body["ExtendMap"] = request.ExtendMapShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FusionPromotionParamList)) {
		body["FusionPromotionParamList"] = request.FusionPromotionParamList
	}

	if !tea.BoolValue(util.IsUnset(request.PublishTag)) {
		body["PublishTag"] = request.PublishTag
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TestAccountUidList)) {
		body["TestAccountUidList"] = request.TestAccountUidList
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteUidList)) {
		body["WhiteUidList"] = request.WhiteUidList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveActivity"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveActivity(request *SaveActivityRequest) (_result *SaveActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveActivityResponse{}
	_body, _err := client.SaveActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCouponWithOptions(request *SendCouponRequest, runtime *util.RuntimeOptions) (_result *SendCouponResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		body["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.FromApp)) {
		body["FromApp"] = request.FromApp
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SellerId)) {
		body["SellerId"] = request.SellerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAmountModelList)) {
		body["UserAmountModelList"] = request.UserAmountModelList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCoupon"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCouponResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCoupon(request *SendCouponRequest) (_result *SendCouponResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCouponResponse{}
	_body, _err := client.SendCouponWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferResourceWithOptions(request *TransferResourceRequest, runtime *util.RuntimeOptions) (_result *TransferResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCode)) {
		body["ActionCode"] = request.ActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferResource"),
		Version:     tea.String("2019-12-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferResource(request *TransferResourceRequest) (_result *TransferResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferResourceResponse{}
	_body, _err := client.TransferResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
