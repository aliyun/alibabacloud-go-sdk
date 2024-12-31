// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateOrderRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetParamStr(v string) *CreateOrderRequest {
	s.ParamStr = &v
	return s
}

type CreateOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetCode(v string) *CreateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrderResponseBody) SetData(v string) *CreateOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOrderResponseBody) SetMessage(v string) *CreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

type CreateOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetStatusCode(v int32) *CreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type DescribeCashDetailResponseBody struct {
	AmountOwed           *string `json:"AmountOwed,omitempty" xml:"AmountOwed,omitempty"`
	AvailableCredit      *string `json:"AvailableCredit,omitempty" xml:"AvailableCredit,omitempty"`
	BalanceAmount        *string `json:"BalanceAmount,omitempty" xml:"BalanceAmount,omitempty"`
	CreditCardAmount     *string `json:"CreditCardAmount,omitempty" xml:"CreditCardAmount,omitempty"`
	CreditLimit          *string `json:"CreditLimit,omitempty" xml:"CreditLimit,omitempty"`
	EnableThresholdAlert *string `json:"EnableThresholdAlert,omitempty" xml:"EnableThresholdAlert,omitempty"`
	FrozenAmount         *string `json:"FrozenAmount,omitempty" xml:"FrozenAmount,omitempty"`
	MiniAlertThreshold   *int64  `json:"MiniAlertThreshold,omitempty" xml:"MiniAlertThreshold,omitempty"`
	RemmitanceAmount     *string `json:"RemmitanceAmount,omitempty" xml:"RemmitanceAmount,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCashDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCashDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCashDetailResponseBody) SetAmountOwed(v string) *DescribeCashDetailResponseBody {
	s.AmountOwed = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetAvailableCredit(v string) *DescribeCashDetailResponseBody {
	s.AvailableCredit = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetBalanceAmount(v string) *DescribeCashDetailResponseBody {
	s.BalanceAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetCreditCardAmount(v string) *DescribeCashDetailResponseBody {
	s.CreditCardAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetCreditLimit(v string) *DescribeCashDetailResponseBody {
	s.CreditLimit = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetEnableThresholdAlert(v string) *DescribeCashDetailResponseBody {
	s.EnableThresholdAlert = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetFrozenAmount(v string) *DescribeCashDetailResponseBody {
	s.FrozenAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetMiniAlertThreshold(v int64) *DescribeCashDetailResponseBody {
	s.MiniAlertThreshold = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetRemmitanceAmount(v string) *DescribeCashDetailResponseBody {
	s.RemmitanceAmount = &v
	return s
}

func (s *DescribeCashDetailResponseBody) SetRequestId(v string) *DescribeCashDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCashDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCashDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCashDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCashDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCashDetailResponse) SetHeaders(v map[string]*string) *DescribeCashDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCashDetailResponse) SetStatusCode(v int32) *DescribeCashDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCashDetailResponse) SetBody(v *DescribeCashDetailResponseBody) *DescribeCashDetailResponse {
	s.Body = v
	return s
}

type DescribeCouponListRequest struct {
	EndDeliveryTime   *string `json:"EndDeliveryTime,omitempty" xml:"EndDeliveryTime,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDeliveryTime *string `json:"StartDeliveryTime,omitempty" xml:"StartDeliveryTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCouponListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponListRequest) SetEndDeliveryTime(v string) *DescribeCouponListRequest {
	s.EndDeliveryTime = &v
	return s
}

func (s *DescribeCouponListRequest) SetPageNum(v int32) *DescribeCouponListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCouponListRequest) SetPageSize(v int32) *DescribeCouponListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponListRequest) SetStartDeliveryTime(v string) *DescribeCouponListRequest {
	s.StartDeliveryTime = &v
	return s
}

func (s *DescribeCouponListRequest) SetStatus(v string) *DescribeCouponListRequest {
	s.Status = &v
	return s
}

type DescribeCouponListResponseBody struct {
	Coupons   *DescribeCouponListResponseBodyCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCouponListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBody) SetCoupons(v *DescribeCouponListResponseBodyCoupons) *DescribeCouponListResponseBody {
	s.Coupons = v
	return s
}

func (s *DescribeCouponListResponseBody) SetRequestId(v string) *DescribeCouponListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCouponListResponseBodyCoupons struct {
	Coupon []*DescribeCouponListResponseBodyCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribeCouponListResponseBodyCoupons) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListResponseBodyCoupons) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCoupons) SetCoupon(v []*DescribeCouponListResponseBodyCouponsCoupon) *DescribeCouponListResponseBodyCoupons {
	s.Coupon = v
	return s
}

type DescribeCouponListResponseBodyCouponsCoupon struct {
	Application      *string                                                  `json:"Application,omitempty" xml:"Application,omitempty"`
	BalanceAmount    *string                                                  `json:"BalanceAmount,omitempty" xml:"BalanceAmount,omitempty"`
	CouponNumber     *string                                                  `json:"CouponNumber,omitempty" xml:"CouponNumber,omitempty"`
	CouponTemplateId *int64                                                   `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	CreationTime     *string                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeliveryTime     *string                                                  `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	Description      *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredAmount    *string                                                  `json:"ExpiredAmount,omitempty" xml:"ExpiredAmount,omitempty"`
	ExpiredTime      *string                                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FrozenAmount     *string                                                  `json:"FrozenAmount,omitempty" xml:"FrozenAmount,omitempty"`
	ModificationTime *string                                                  `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PriceLimit       *string                                                  `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	ProductCodes     *DescribeCouponListResponseBodyCouponsCouponProductCodes `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty" type:"Struct"`
	Status           *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalAmount      *string                                                  `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	TradeTypes       *DescribeCouponListResponseBodyCouponsCouponTradeTypes   `json:"TradeTypes,omitempty" xml:"TradeTypes,omitempty" type:"Struct"`
}

func (s DescribeCouponListResponseBodyCouponsCoupon) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListResponseBodyCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetApplication(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.Application = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetBalanceAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.BalanceAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetCouponNumber(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.CouponNumber = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetCouponTemplateId(v int64) *DescribeCouponListResponseBodyCouponsCoupon {
	s.CouponTemplateId = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetCreationTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.CreationTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetDeliveryTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.DeliveryTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetDescription(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetExpiredAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ExpiredAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetExpiredTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetFrozenAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.FrozenAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetModificationTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ModificationTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetPriceLimit(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.PriceLimit = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetProductCodes(v *DescribeCouponListResponseBodyCouponsCouponProductCodes) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ProductCodes = v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetStatus(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.Status = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetTotalAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.TotalAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetTradeTypes(v *DescribeCouponListResponseBodyCouponsCouponTradeTypes) *DescribeCouponListResponseBodyCouponsCoupon {
	s.TradeTypes = v
	return s
}

type DescribeCouponListResponseBodyCouponsCouponProductCodes struct {
	ProductCode []*string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty" type:"Repeated"`
}

func (s DescribeCouponListResponseBodyCouponsCouponProductCodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListResponseBodyCouponsCouponProductCodes) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCouponsCouponProductCodes) SetProductCode(v []*string) *DescribeCouponListResponseBodyCouponsCouponProductCodes {
	s.ProductCode = v
	return s
}

type DescribeCouponListResponseBodyCouponsCouponTradeTypes struct {
	TradeType []*string `json:"TradeType,omitempty" xml:"TradeType,omitempty" type:"Repeated"`
}

func (s DescribeCouponListResponseBodyCouponsCouponTradeTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListResponseBodyCouponsCouponTradeTypes) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCouponsCouponTradeTypes) SetTradeType(v []*string) *DescribeCouponListResponseBodyCouponsCouponTradeTypes {
	s.TradeType = v
	return s
}

type DescribeCouponListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCouponListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCouponListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponse) SetHeaders(v map[string]*string) *DescribeCouponListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCouponListResponse) SetStatusCode(v int32) *DescribeCouponListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCouponListResponse) SetBody(v *DescribeCouponListResponseBody) *DescribeCouponListResponse {
	s.Body = v
	return s
}

type OpenCallbackRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s OpenCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCallbackRequest) GoString() string {
	return s.String()
}

func (s *OpenCallbackRequest) SetParamStr(v string) *OpenCallbackRequest {
	s.ParamStr = &v
	return s
}

type OpenCallbackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCallbackResponseBody) SetCode(v string) *OpenCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *OpenCallbackResponseBody) SetData(v string) *OpenCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *OpenCallbackResponseBody) SetMessage(v string) *OpenCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *OpenCallbackResponseBody) SetRequestId(v string) *OpenCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCallbackResponseBody) SetSuccess(v bool) *OpenCallbackResponseBody {
	s.Success = &v
	return s
}

type OpenCallbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCallbackResponse) GoString() string {
	return s.String()
}

func (s *OpenCallbackResponse) SetHeaders(v map[string]*string) *OpenCallbackResponse {
	s.Headers = v
	return s
}

func (s *OpenCallbackResponse) SetStatusCode(v int32) *OpenCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCallbackResponse) SetBody(v *OpenCallbackResponseBody) *OpenCallbackResponse {
	s.Body = v
	return s
}

type QueryForCssOrderRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s QueryForCssOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryForCssOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryForCssOrderRequest) SetParamStr(v string) *QueryForCssOrderRequest {
	s.ParamStr = &v
	return s
}

type QueryForCssOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryForCssOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryForCssOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryForCssOrderResponseBody) SetCode(v string) *QueryForCssOrderResponseBody {
	s.Code = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetData(v string) *QueryForCssOrderResponseBody {
	s.Data = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetMessage(v string) *QueryForCssOrderResponseBody {
	s.Message = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetRequestId(v string) *QueryForCssOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetSuccess(v bool) *QueryForCssOrderResponseBody {
	s.Success = &v
	return s
}

type QueryForCssOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryForCssOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryForCssOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryForCssOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryForCssOrderResponse) SetHeaders(v map[string]*string) *QueryForCssOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryForCssOrderResponse) SetStatusCode(v int32) *QueryForCssOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryForCssOrderResponse) SetBody(v *QueryForCssOrderResponseBody) *QueryForCssOrderResponse {
	s.Body = v
	return s
}

type VnoBatchRefundOrderRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s VnoBatchRefundOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s VnoBatchRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *VnoBatchRefundOrderRequest) SetParamStr(v string) *VnoBatchRefundOrderRequest {
	s.ParamStr = &v
	return s
}

type VnoBatchRefundOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VnoBatchRefundOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VnoBatchRefundOrderResponseBody) GoString() string {
	return s.String()
}

func (s *VnoBatchRefundOrderResponseBody) SetCode(v string) *VnoBatchRefundOrderResponseBody {
	s.Code = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetData(v string) *VnoBatchRefundOrderResponseBody {
	s.Data = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetMessage(v string) *VnoBatchRefundOrderResponseBody {
	s.Message = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetRequestId(v string) *VnoBatchRefundOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetSuccess(v bool) *VnoBatchRefundOrderResponseBody {
	s.Success = &v
	return s
}

type VnoBatchRefundOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VnoBatchRefundOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VnoBatchRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s VnoBatchRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *VnoBatchRefundOrderResponse) SetHeaders(v map[string]*string) *VnoBatchRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *VnoBatchRefundOrderResponse) SetStatusCode(v int32) *VnoBatchRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *VnoBatchRefundOrderResponse) SetBody(v *VnoBatchRefundOrderResponseBody) *VnoBatchRefundOrderResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("bss"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CreateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamStr)) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrder"),
		Version:     tea.String("2014-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateOrderRequest
//
// @return CreateOrderResponse
func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取现金明细
//
// @param request - DescribeCashDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCashDetailResponse
func (client *Client) DescribeCashDetailWithOptions(runtime *util.RuntimeOptions) (_result *DescribeCashDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeCashDetail"),
		Version:     tea.String("2014-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCashDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取现金明细
//
// @return DescribeCashDetailResponse
func (client *Client) DescribeCashDetail() (_result *DescribeCashDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCashDetailResponse{}
	_body, _err := client.DescribeCashDetailWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCouponListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponListResponse
func (client *Client) DescribeCouponListWithOptions(request *DescribeCouponListRequest, runtime *util.RuntimeOptions) (_result *DescribeCouponListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDeliveryTime)) {
		query["EndDeliveryTime"] = request.EndDeliveryTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDeliveryTime)) {
		query["StartDeliveryTime"] = request.StartDeliveryTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCouponList"),
		Version:     tea.String("2014-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCouponListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCouponListRequest
//
// @return DescribeCouponListResponse
func (client *Client) DescribeCouponList(request *DescribeCouponListRequest) (_result *DescribeCouponListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCouponListResponse{}
	_body, _err := client.DescribeCouponListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// OpenCallback
//
// @param request - OpenCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenCallbackResponse
func (client *Client) OpenCallbackWithOptions(request *OpenCallbackRequest, runtime *util.RuntimeOptions) (_result *OpenCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamStr)) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenCallback"),
		Version:     tea.String("2014-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// OpenCallback
//
// @param request - OpenCallbackRequest
//
// @return OpenCallbackResponse
func (client *Client) OpenCallback(request *OpenCallbackRequest) (_result *OpenCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCallbackResponse{}
	_body, _err := client.OpenCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryForCssOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryForCssOrderResponse
func (client *Client) QueryForCssOrderWithOptions(request *QueryForCssOrderRequest, runtime *util.RuntimeOptions) (_result *QueryForCssOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamStr)) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryForCssOrder"),
		Version:     tea.String("2014-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryForCssOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryForCssOrderRequest
//
// @return QueryForCssOrderResponse
func (client *Client) QueryForCssOrder(request *QueryForCssOrderRequest) (_result *QueryForCssOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryForCssOrderResponse{}
	_body, _err := client.QueryForCssOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VnoBatchRefundOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VnoBatchRefundOrderResponse
func (client *Client) VnoBatchRefundOrderWithOptions(request *VnoBatchRefundOrderRequest, runtime *util.RuntimeOptions) (_result *VnoBatchRefundOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamStr)) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("vnoBatchRefundOrder"),
		Version:     tea.String("2014-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VnoBatchRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - VnoBatchRefundOrderRequest
//
// @return VnoBatchRefundOrderResponse
func (client *Client) VnoBatchRefundOrder(request *VnoBatchRefundOrderRequest) (_result *VnoBatchRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VnoBatchRefundOrderResponse{}
	_body, _err := client.VnoBatchRefundOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
