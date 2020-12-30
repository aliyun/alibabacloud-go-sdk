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

type QueryCdpOfferRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Vendor               *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	ChannelType          *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Province             *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s QueryCdpOfferRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferRequest) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferRequest) SetOwnerId(v int64) *QueryCdpOfferRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCdpOfferRequest) SetResourceOwnerAccount(v string) *QueryCdpOfferRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCdpOfferRequest) SetResourceOwnerId(v int64) *QueryCdpOfferRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCdpOfferRequest) SetVendor(v string) *QueryCdpOfferRequest {
	s.Vendor = &v
	return s
}

func (s *QueryCdpOfferRequest) SetChannelType(v string) *QueryCdpOfferRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryCdpOfferRequest) SetProvince(v string) *QueryCdpOfferRequest {
	s.Province = &v
	return s
}

type QueryCdpOfferResponseBody struct {
	FlowOffers *QueryCdpOfferResponseBodyFlowOffers `json:"FlowOffers,omitempty" xml:"FlowOffers,omitempty" type:"Struct"`
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryCdpOfferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferResponseBody) SetFlowOffers(v *QueryCdpOfferResponseBodyFlowOffers) *QueryCdpOfferResponseBody {
	s.FlowOffers = v
	return s
}

func (s *QueryCdpOfferResponseBody) SetMessage(v string) *QueryCdpOfferResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCdpOfferResponseBody) SetRequestId(v string) *QueryCdpOfferResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCdpOfferResponseBody) SetCode(v string) *QueryCdpOfferResponseBody {
	s.Code = &v
	return s
}

type QueryCdpOfferResponseBodyFlowOffers struct {
	FlowOffer []*QueryCdpOfferResponseBodyFlowOffersFlowOffer `json:"FlowOffer,omitempty" xml:"FlowOffer,omitempty" type:"Repeated"`
}

func (s QueryCdpOfferResponseBodyFlowOffers) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferResponseBodyFlowOffers) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferResponseBodyFlowOffers) SetFlowOffer(v []*QueryCdpOfferResponseBodyFlowOffersFlowOffer) *QueryCdpOfferResponseBodyFlowOffers {
	s.FlowOffer = v
	return s
}

type QueryCdpOfferResponseBodyFlowOffersFlowOffer struct {
	Right       *string `json:"Right,omitempty" xml:"Right,omitempty"`
	UseEff      *string `json:"UseEff,omitempty" xml:"UseEff,omitempty"`
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	UseLimit    *string `json:"UseLimit,omitempty" xml:"UseLimit,omitempty"`
	UseScene    *string `json:"UseScene,omitempty" xml:"UseScene,omitempty"`
	Vendor      *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	Grade       *string `json:"Grade,omitempty" xml:"Grade,omitempty"`
	OfferId     *int64  `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
	Price       *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	FlowType    *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	Discount    *string `json:"Discount,omitempty" xml:"Discount,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s QueryCdpOfferResponseBodyFlowOffersFlowOffer) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferResponseBodyFlowOffersFlowOffer) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetRight(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.Right = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetUseEff(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.UseEff = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetChannelType(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.ChannelType = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetUseLimit(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.UseLimit = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetUseScene(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.UseScene = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetVendor(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.Vendor = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetGrade(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.Grade = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetOfferId(v int64) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.OfferId = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetPrice(v int64) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.Price = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetFlowType(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.FlowType = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetDiscount(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.Discount = &v
	return s
}

func (s *QueryCdpOfferResponseBodyFlowOffersFlowOffer) SetProvince(v string) *QueryCdpOfferResponseBodyFlowOffersFlowOffer {
	s.Province = &v
	return s
}

type QueryCdpOfferResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCdpOfferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCdpOfferResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferResponse) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferResponse) SetHeaders(v map[string]*string) *QueryCdpOfferResponse {
	s.Headers = v
	return s
}

func (s *QueryCdpOfferResponse) SetBody(v *QueryCdpOfferResponseBody) *QueryCdpOfferResponse {
	s.Body = v
	return s
}

type QueryCdpOfferByIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OfferId              *int64  `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
}

func (s QueryCdpOfferByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferByIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferByIdRequest) SetOwnerId(v int64) *QueryCdpOfferByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCdpOfferByIdRequest) SetResourceOwnerAccount(v string) *QueryCdpOfferByIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCdpOfferByIdRequest) SetResourceOwnerId(v int64) *QueryCdpOfferByIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCdpOfferByIdRequest) SetOfferId(v int64) *QueryCdpOfferByIdRequest {
	s.OfferId = &v
	return s
}

type QueryCdpOfferByIdResponseBody struct {
	FlowOffers *QueryCdpOfferByIdResponseBodyFlowOffers `json:"FlowOffers,omitempty" xml:"FlowOffers,omitempty" type:"Struct"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryCdpOfferByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferByIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferByIdResponseBody) SetFlowOffers(v *QueryCdpOfferByIdResponseBodyFlowOffers) *QueryCdpOfferByIdResponseBody {
	s.FlowOffers = v
	return s
}

func (s *QueryCdpOfferByIdResponseBody) SetMessage(v string) *QueryCdpOfferByIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBody) SetRequestId(v string) *QueryCdpOfferByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBody) SetCode(v string) *QueryCdpOfferByIdResponseBody {
	s.Code = &v
	return s
}

type QueryCdpOfferByIdResponseBodyFlowOffers struct {
	FlowOffer []*QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer `json:"FlowOffer,omitempty" xml:"FlowOffer,omitempty" type:"Repeated"`
}

func (s QueryCdpOfferByIdResponseBodyFlowOffers) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferByIdResponseBodyFlowOffers) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffers) SetFlowOffer(v []*QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) *QueryCdpOfferByIdResponseBodyFlowOffers {
	s.FlowOffer = v
	return s
}

type QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer struct {
	Right       *string `json:"Right,omitempty" xml:"Right,omitempty"`
	UseEff      *string `json:"UseEff,omitempty" xml:"UseEff,omitempty"`
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	UseLimit    *string `json:"UseLimit,omitempty" xml:"UseLimit,omitempty"`
	UseScene    *string `json:"UseScene,omitempty" xml:"UseScene,omitempty"`
	Vendor      *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	Grade       *string `json:"Grade,omitempty" xml:"Grade,omitempty"`
	OfferId     *int64  `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
	Price       *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	FlowType    *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	Discount    *string `json:"Discount,omitempty" xml:"Discount,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetRight(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.Right = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetUseEff(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.UseEff = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetChannelType(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.ChannelType = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetUseLimit(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.UseLimit = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetUseScene(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.UseScene = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetVendor(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.Vendor = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetGrade(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.Grade = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetOfferId(v int64) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.OfferId = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetPrice(v int64) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.Price = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetFlowType(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.FlowType = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetDiscount(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.Discount = &v
	return s
}

func (s *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer) SetProvince(v string) *QueryCdpOfferByIdResponseBodyFlowOffersFlowOffer {
	s.Province = &v
	return s
}

type QueryCdpOfferByIdResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCdpOfferByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCdpOfferByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOfferByIdResponse) GoString() string {
	return s.String()
}

func (s *QueryCdpOfferByIdResponse) SetHeaders(v map[string]*string) *QueryCdpOfferByIdResponse {
	s.Headers = v
	return s
}

func (s *QueryCdpOfferByIdResponse) SetBody(v *QueryCdpOfferByIdResponseBody) *QueryCdpOfferByIdResponse {
	s.Body = v
	return s
}

type QueryCdpOrderRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OutOrderId           *string `json:"OutOrderId,omitempty" xml:"OutOrderId,omitempty"`
}

func (s QueryCdpOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryCdpOrderRequest) SetOwnerId(v int64) *QueryCdpOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCdpOrderRequest) SetResourceOwnerAccount(v string) *QueryCdpOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCdpOrderRequest) SetResourceOwnerId(v int64) *QueryCdpOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCdpOrderRequest) SetOutOrderId(v string) *QueryCdpOrderRequest {
	s.OutOrderId = &v
	return s
}

type QueryCdpOrderResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryCdpOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryCdpOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCdpOrderResponseBody) SetMessage(v string) *QueryCdpOrderResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCdpOrderResponseBody) SetRequestId(v string) *QueryCdpOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCdpOrderResponseBody) SetData(v *QueryCdpOrderResponseBodyData) *QueryCdpOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryCdpOrderResponseBody) SetCode(v string) *QueryCdpOrderResponseBody {
	s.Code = &v
	return s
}

type QueryCdpOrderResponseBodyData struct {
	ExtendParam *string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	ResultCode  *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMsg   *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
}

func (s QueryCdpOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCdpOrderResponseBodyData) SetExtendParam(v string) *QueryCdpOrderResponseBodyData {
	s.ExtendParam = &v
	return s
}

func (s *QueryCdpOrderResponseBodyData) SetResultCode(v string) *QueryCdpOrderResponseBodyData {
	s.ResultCode = &v
	return s
}

func (s *QueryCdpOrderResponseBodyData) SetResultMsg(v string) *QueryCdpOrderResponseBodyData {
	s.ResultMsg = &v
	return s
}

type QueryCdpOrderResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCdpOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCdpOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCdpOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryCdpOrderResponse) SetHeaders(v map[string]*string) *QueryCdpOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryCdpOrderResponse) SetBody(v *QueryCdpOrderResponseBody) *QueryCdpOrderResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dycdpapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) QueryCdpOfferWithOptions(request *QueryCdpOfferRequest, runtime *util.RuntimeOptions) (_result *QueryCdpOfferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCdpOfferResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCdpOffer"), tea.String("2018-06-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCdpOffer(request *QueryCdpOfferRequest) (_result *QueryCdpOfferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCdpOfferResponse{}
	_body, _err := client.QueryCdpOfferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCdpOfferByIdWithOptions(request *QueryCdpOfferByIdRequest, runtime *util.RuntimeOptions) (_result *QueryCdpOfferByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCdpOfferByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCdpOfferById"), tea.String("2018-06-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCdpOfferById(request *QueryCdpOfferByIdRequest) (_result *QueryCdpOfferByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCdpOfferByIdResponse{}
	_body, _err := client.QueryCdpOfferByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCdpOrderWithOptions(request *QueryCdpOrderRequest, runtime *util.RuntimeOptions) (_result *QueryCdpOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCdpOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCdpOrder"), tea.String("2018-06-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCdpOrder(request *QueryCdpOrderRequest) (_result *QueryCdpOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCdpOrderResponse{}
	_body, _err := client.QueryCdpOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
