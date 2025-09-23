// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # IPC目标检测
//
// @param request - DetectIPCObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectIPCObjectResponse
func (client *Client) DetectIPCObjectWithContext(ctx context.Context, request *DetectIPCObjectRequest, runtime *dara.RuntimeOptions) (_result *DetectIPCObjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectIPCObject"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectIPCObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 猫鼠识别
//
// @param request - DetectKitchenAnimalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectKitchenAnimalsResponse
func (client *Client) DetectKitchenAnimalsWithContext(ctx context.Context, request *DetectKitchenAnimalsRequest, runtime *dara.RuntimeOptions) (_result *DetectKitchenAnimalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURLA) {
		body["ImageURLA"] = request.ImageURLA
	}

	if !dara.IsNil(request.ImageURLB) {
		body["ImageURLB"] = request.ImageURLB
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectKitchenAnimals"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectKitchenAnimalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectMainBodyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectMainBodyResponse
func (client *Client) DetectMainBodyWithContext(ctx context.Context, request *DetectMainBodyRequest, runtime *dara.RuntimeOptions) (_result *DetectMainBodyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectMainBody"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectMainBodyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectObjectResponse
func (client *Client) DetectObjectWithContext(ctx context.Context, request *DetectObjectRequest, runtime *dara.RuntimeOptions) (_result *DetectObjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectObject"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆拥堵检测
//
// @param tmpReq - DetectVehicleICongestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectVehicleICongestionResponse
func (client *Client) DetectVehicleICongestionWithContext(ctx context.Context, tmpReq *DetectVehicleICongestionRequest, runtime *dara.RuntimeOptions) (_result *DetectVehicleICongestionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DetectVehicleICongestionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PreRegionIntersectFeatures) {
		request.PreRegionIntersectFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreRegionIntersectFeatures, dara.String("PreRegionIntersectFeatures"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoadRegions) {
		request.RoadRegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoadRegions, dara.String("RoadRegions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.PreRegionIntersectFeaturesShrink) {
		body["PreRegionIntersectFeatures"] = request.PreRegionIntersectFeaturesShrink
	}

	if !dara.IsNil(request.RoadRegionsShrink) {
		body["RoadRegions"] = request.RoadRegionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectVehicleICongestion"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectVehicleICongestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆违停检测
//
// @param tmpReq - DetectVehicleIllegalParkingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectVehicleIllegalParkingResponse
func (client *Client) DetectVehicleIllegalParkingWithContext(ctx context.Context, tmpReq *DetectVehicleIllegalParkingRequest, runtime *dara.RuntimeOptions) (_result *DetectVehicleIllegalParkingResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DetectVehicleIllegalParkingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoadRegions) {
		request.RoadRegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoadRegions, dara.String("RoadRegions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	if !dara.IsNil(request.RoadRegionsShrink) {
		body["RoadRegions"] = request.RoadRegionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectVehicleIllegalParking"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectVehicleIllegalParkingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # IPC视频文件目标检测
//
// @param request - DetectVideoIPCObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectVideoIPCObjectResponse
func (client *Client) DetectVideoIPCObjectWithContext(ctx context.Context, request *DetectVideoIPCObjectRequest, runtime *dara.RuntimeOptions) (_result *DetectVideoIPCObjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackOnlyHasObject) {
		body["CallbackOnlyHasObject"] = request.CallbackOnlyHasObject
	}

	if !dara.IsNil(request.StartTimestamp) {
		body["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.VideoURL) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectVideoIPCObject"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectVideoIPCObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetectWhiteBaseImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectWhiteBaseImageResponse
func (client *Client) DetectWhiteBaseImageWithContext(ctx context.Context, request *DetectWhiteBaseImageRequest, runtime *dara.RuntimeOptions) (_result *DetectWhiteBaseImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageURL) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectWhiteBaseImage"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectWhiteBaseImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 着装检测
//
// @param tmpReq - DetectWorkwearRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectWorkwearResponse
func (client *Client) DetectWorkwearWithContext(ctx context.Context, tmpReq *DetectWorkwearRequest, runtime *dara.RuntimeOptions) (_result *DetectWorkwearResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DetectWorkwearShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Clothes) {
		request.ClothesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Clothes, dara.String("Clothes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClothesShrink) {
		body["Clothes"] = request.ClothesShrink
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectWorkwear"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectWorkwearResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询异步任务结果
//
// @param request - GetAsyncJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncJobResultResponse
func (client *Client) GetAsyncJobResultWithContext(ctx context.Context, request *GetAsyncJobResultRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncJobResult"),
		Version:     dara.String("2019-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
