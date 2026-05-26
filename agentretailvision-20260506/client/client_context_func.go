// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 商品导入
//
// @param tmpReq - ImportProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportProductsResponse
func (client *Client) ImportProductsWithContext(ctx context.Context, tmpReq *ImportProductsRequest, runtime *dara.RuntimeOptions) (_result *ImportProductsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecognitionResultResponse
func (client *Client) QueryRecognitionResultWithContext(ctx context.Context, request *QueryRecognitionResultRequest, runtime *dara.RuntimeOptions) (_result *QueryRecognitionResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - RecognizeOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeOrderResponse
func (client *Client) RecognizeOrderWithContext(ctx context.Context, tmpReq *RecognizeOrderRequest, runtime *dara.RuntimeOptions) (_result *RecognizeOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterWebhookResponse
func (client *Client) RegisterWebhookWithContext(ctx context.Context, request *RegisterWebhookRequest, runtime *dara.RuntimeOptions) (_result *RegisterWebhookResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProductResponse
func (client *Client) UpdateProductWithContext(ctx context.Context, tmpReq *UpdateProductRequest, runtime *dara.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
