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
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentretailvision"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 商品导入
//
// @param tmpReq - ImportProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportProductsResponse
func (client *Client) ImportProductsWithOptions(tmpReq *ImportProductsRequest, runtime *dara.RuntimeOptions) (_result *ImportProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportProductsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtraImages) {
		request.ExtraImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraImages, dara.String("ExtraImages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MainImage) {
		request.MainImageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MainImage, dara.String("MainImage"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MultiViewImages) {
		request.MultiViewImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiViewImages, dara.String("MultiViewImages"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.ExtraImagesShrink) {
		query["ExtraImages"] = request.ExtraImagesShrink
	}

	if !dara.IsNil(request.ImageTitle) {
		query["ImageTitle"] = request.ImageTitle
	}

	if !dara.IsNil(request.ItemUniqueId) {
		query["ItemUniqueId"] = request.ItemUniqueId
	}

	if !dara.IsNil(request.MainImageShrink) {
		query["MainImage"] = request.MainImageShrink
	}

	if !dara.IsNil(request.MultiViewImagesShrink) {
		query["MultiViewImages"] = request.MultiViewImagesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportProducts"),
		Version:     dara.String("2026-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商品导入
//
// @param request - ImportProductsRequest
//
// @return ImportProductsResponse
func (client *Client) ImportProducts(request *ImportProductsRequest) (_result *ImportProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportProductsResponse{}
	_body, _err := client.ImportProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务状态
//
// @param request - QueryRecognitionResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecognitionResultResponse
func (client *Client) QueryRecognitionResultWithOptions(request *QueryRecognitionResultRequest, runtime *dara.RuntimeOptions) (_result *QueryRecognitionResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderUniqueId) {
		query["OrderUniqueId"] = request.OrderUniqueId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecognitionResult"),
		Version:     dara.String("2026-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecognitionResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务状态
//
// @param request - QueryRecognitionResultRequest
//
// @return QueryRecognitionResultResponse
func (client *Client) QueryRecognitionResult(request *QueryRecognitionResultRequest) (_result *QueryRecognitionResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecognitionResultResponse{}
	_body, _err := client.QueryRecognitionResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 购物识别
//
// @param tmpReq - RecognizeOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeOrderResponse
func (client *Client) RecognizeOrderWithOptions(tmpReq *RecognizeOrderRequest, runtime *dara.RuntimeOptions) (_result *RecognizeOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RecognizeOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CandidateItems) {
		request.CandidateItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CandidateItems, dara.String("CandidateItems"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoUrls) {
		request.VideoUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoUrls, dara.String("VideoUrls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.CandidateItemsShrink) {
		query["CandidateItems"] = request.CandidateItemsShrink
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OrderUniqueId) {
		query["OrderUniqueId"] = request.OrderUniqueId
	}

	if !dara.IsNil(request.VideoUrlsShrink) {
		query["VideoUrls"] = request.VideoUrlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecognizeOrder"),
		Version:     dara.String("2026-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 购物识别
//
// @param request - RecognizeOrderRequest
//
// @return RecognizeOrderResponse
func (client *Client) RecognizeOrder(request *RecognizeOrderRequest) (_result *RecognizeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecognizeOrderResponse{}
	_body, _err := client.RecognizeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Webhook注册
//
// @param request - RegisterWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterWebhookResponse
func (client *Client) RegisterWebhookWithOptions(request *RegisterWebhookRequest, runtime *dara.RuntimeOptions) (_result *RegisterWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackSecret) {
		query["CallbackSecret"] = request.CallbackSecret
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterWebhook"),
		Version:     dara.String("2026-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Webhook注册
//
// @param request - RegisterWebhookRequest
//
// @return RegisterWebhookResponse
func (client *Client) RegisterWebhook(request *RegisterWebhookRequest) (_result *RegisterWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterWebhookResponse{}
	_body, _err := client.RegisterWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 商品更新
//
// @param tmpReq - UpdateProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProductResponse
func (client *Client) UpdateProductWithOptions(tmpReq *UpdateProductRequest, runtime *dara.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProductShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtraImages) {
		request.ExtraImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraImages, dara.String("ExtraImages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MainImage) {
		request.MainImageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MainImage, dara.String("MainImage"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MultiViewImages) {
		request.MultiViewImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiViewImages, dara.String("MultiViewImages"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.ExtraImagesShrink) {
		query["ExtraImages"] = request.ExtraImagesShrink
	}

	if !dara.IsNil(request.ImageTitle) {
		query["ImageTitle"] = request.ImageTitle
	}

	if !dara.IsNil(request.ItemUniqueId) {
		query["ItemUniqueId"] = request.ItemUniqueId
	}

	if !dara.IsNil(request.MainImageShrink) {
		query["MainImage"] = request.MainImageShrink
	}

	if !dara.IsNil(request.MultiViewImagesShrink) {
		query["MultiViewImages"] = request.MultiViewImagesShrink
	}

	if !dara.IsNil(request.PlatformItemId) {
		query["PlatformItemId"] = request.PlatformItemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProduct"),
		Version:     dara.String("2026-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商品更新
//
// @param request - UpdateProductRequest
//
// @return UpdateProductResponse
func (client *Client) UpdateProduct(request *UpdateProductRequest) (_result *UpdateProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProductResponse{}
	_body, _err := client.UpdateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
