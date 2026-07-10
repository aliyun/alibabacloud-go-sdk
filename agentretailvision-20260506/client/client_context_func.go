// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Generates a composite image for single-item multi-image or multi-item scenarios.
//
// Description:
//
// ## Request description
//
// - When `groupType=1`, `platformItemIdList` must contain only one element.
//
// - When `groupType=2`, `platformItemIdList` can contain 1 to 10 elements.
//
// @param tmpReq - GenerateGroupImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateGroupImageResponse
func (client *Client) GenerateGroupImageWithContext(ctx context.Context, tmpReq *GenerateGroupImageRequest, runtime *dara.RuntimeOptions) (_result *GenerateGroupImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateGroupImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PlatformItemIdList) {
		request.PlatformItemIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlatformItemIdList, dara.String("PlatformItemIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackSecret) {
		query["CallbackSecret"] = request.CallbackSecret
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.PlatformItemIdListShrink) {
		query["PlatformItemIdList"] = request.PlatformItemIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateGroupImage"),
		Version:     dara.String("2026-05-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateGroupImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds product information. After a successful import, the platform returns a globally unique platform_item_id for subsequent updates and recognition result association.
//
// Description:
//
// ## Operation description
//
// - This operation is used to add product information.
//
// - After you import products to the product library, they are stored in Alibaba Cloud OSS for direct recall and retrieval by the product recognition API.
//
// - You must provide at least one main image URL, and the `item_unique_id` must be unique within the same business party.
//
// - You can optionally provide multi-angle views and extra images to improve recognition accuracy.
//
// - The `device_id` field can be used to establish an association between a device and product vectors, but it is not required.
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
// At least one result retrieval method must be integrated: webhook callback or task status query. Both methods can be used simultaneously.
//
//   - If the user chooses the webhook callback method, the receiving endpoint must be prepared in advance and implemented according to the following request and response parameters.
//
//   - After the recognition task is completed, the platform will push the results to the business party based on the callback URL bound to the task.
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
// Used for intelligent recognition scenarios. Requires uploading the OSS address of shopping videos. The platform creates an asynchronous recognition task and immediately returns a task_id. Notifications are sent via webhook, and the results need to be actively retrieved through the query API.
//
// Description:
//
// ## Request Description
//
// - The user must provide `caller_uid` and `order_unique_id` as required parameters.
//
// - The `video_urls` parameter supports video files in mp4, avi, mov, and mkv formats, with a size limit of 100 MB, a duration of no more than 3 minutes, a resolution between 480p and 1080p, and specific aspect ratio requirements.
//
// - At least one of `device_id` or `candidate_items` must be provided to specify the recognition scope. If both are provided, the system first filters by the device product library and then further filters based on the candidate items list.
//
// - Optionally, the user can specify a `callback_url` to receive notifications of the recognition results. If not provided, the pre-registered default webhook address is used.
//
// - If a request is submitted repeatedly with the same `order_unique_id`, the system directly returns the previously existing task status.
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
// Registers or updates the default webhook callback URL.
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
// Updates the information of an existing item on the platform.
//
// Description:
//
// ## Operation description
//
// - The platform_item_id parameter is used as the primary identifier for the update.
//
// - If both platform_item_id and item_unique_id are specified, they must point to the same item.
//
// - The item title (image_title) and the list of main image URLs (main_image) are required. The main_image parameter must contain at least one image.
//
// - Optional parameters include the multi-angle image list (multi_view_images), the list of additional image URLs (extra_images), and the device ID (device_id).
//
// - In multi_view_images, each object must contain the image OSS address (url) and the shooting angle (angle). Valid values of angle: top view (up), bottom view (down), left view (left), right view (right), front view (front), and back view (back).
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
