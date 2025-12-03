// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("bss"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建订单
//
// @param request - CreateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamStr) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrder"),
		Version:     dara.String("2014-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建订单
//
// @param request - CreateOrderRequest
//
// @return CreateOrderResponse
func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeCashDetailWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeCashDetailResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCashDetail"),
		Version:     dara.String("2014-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCashDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取现金明细
//
// @return DescribeCashDetailResponse
func (client *Client) DescribeCashDetail() (_result *DescribeCashDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCashDetailResponse{}
	_body, _err := client.DescribeCashDetailWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询卡券列表
//
// @param request - DescribeCouponListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponListResponse
func (client *Client) DescribeCouponListWithOptions(request *DescribeCouponListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDeliveryTime) {
		query["EndDeliveryTime"] = request.EndDeliveryTime
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDeliveryTime) {
		query["StartDeliveryTime"] = request.StartDeliveryTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCouponList"),
		Version:     dara.String("2014-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCouponListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询卡券列表
//
// @param request - DescribeCouponListRequest
//
// @return DescribeCouponListResponse
func (client *Client) DescribeCouponList(request *DescribeCouponListRequest) (_result *DescribeCouponListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
// 生产开通回调接口
//
// @param request - OpenCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenCallbackResponse
func (client *Client) OpenCallbackWithOptions(request *OpenCallbackRequest, runtime *dara.RuntimeOptions) (_result *OpenCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamStr) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenCallback"),
		Version:     dara.String("2014-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生产开通回调接口
//
// @param request - OpenCallbackRequest
//
// @return OpenCallbackResponse
func (client *Client) OpenCallback(request *OpenCallbackRequest) (_result *OpenCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenCallbackResponse{}
	_body, _err := client.OpenCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订单询价
//
// @param request - QueryForCssOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryForCssOrderResponse
func (client *Client) QueryForCssOrderWithOptions(request *QueryForCssOrderRequest, runtime *dara.RuntimeOptions) (_result *QueryForCssOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamStr) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryForCssOrder"),
		Version:     dara.String("2014-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryForCssOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订单询价
//
// @param request - QueryForCssOrderRequest
//
// @return QueryForCssOrderResponse
func (client *Client) QueryForCssOrder(request *QueryForCssOrderRequest) (_result *QueryForCssOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryForCssOrderResponse{}
	_body, _err := client.QueryForCssOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// vnoBatchRefundOrder
//
// @param request - VnoBatchRefundOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VnoBatchRefundOrderResponse
func (client *Client) VnoBatchRefundOrderWithOptions(request *VnoBatchRefundOrderRequest, runtime *dara.RuntimeOptions) (_result *VnoBatchRefundOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamStr) {
		query["paramStr"] = request.ParamStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("vnoBatchRefundOrder"),
		Version:     dara.String("2014-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VnoBatchRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// vnoBatchRefundOrder
//
// @param request - VnoBatchRefundOrderRequest
//
// @return VnoBatchRefundOrderResponse
func (client *Client) VnoBatchRefundOrder(request *VnoBatchRefundOrderRequest) (_result *VnoBatchRefundOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VnoBatchRefundOrderResponse{}
	_body, _err := client.VnoBatchRefundOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
