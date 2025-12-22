// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 文档审核结果
//
// @param request - DescribeFileModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileModerationResultResponse
func (client *Client) DescribeFileModerationResultWithContext(ctx context.Context, request *DescribeFileModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFileModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the moderation results of an Image Moderation 2.0 task.
//
// Description:
//
//	  Billing: This operation is free of charge.
//
//		- QPS limit: You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeImageModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageModerationResultResponse
func (client *Client) DescribeImageModerationResultWithContext(ctx context.Context, request *DescribeImageModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检测结果辅助信息
//
// @param request - DescribeImageResultExtRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageResultExtResponse
func (client *Client) DescribeImageResultExtWithContext(ctx context.Context, request *DescribeImageResultExtRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageResultExtResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InfoType) {
		body["InfoType"] = request.InfoType
	}

	if !dara.IsNil(request.ReqId) {
		body["ReqId"] = request.ReqId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageResultExt"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageResultExtResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询异步多模态检测结果
//
// @param request - DescribeMultimodalModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultimodalModerationResultResponse
func (client *Client) DescribeMultimodalModerationResultWithContext(ctx context.Context, request *DescribeMultimodalModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultimodalModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReqId) {
		query["ReqId"] = request.ReqId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMultimodalModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMultimodalModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the moderation results based on the ReqId returned by asynchronous URL moderation.
//
// Description:
//
//	  Billing: This operation is free of charge.
//
//		- Query timeout: We recommend that you query moderation results at least 480 seconds after you send an asynchronous moderation request. Content Moderation retains moderation results for up to 3 days. After 3 days, the results are deleted.
//
//		- You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUrlModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUrlModerationResultResponse
func (client *Client) DescribeUrlModerationResultWithContext(ctx context.Context, request *DescribeUrlModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeUrlModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReqId) {
		body["ReqId"] = request.ReqId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUrlModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUrlModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档审核
//
// @param request - FileModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileModerationResponse
func (client *Client) FileModerationWithContext(ctx context.Context, request *FileModerationRequest, runtime *dara.RuntimeOptions) (_result *FileModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FileModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片异步检测
//
// @param request - ImageAsyncModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageAsyncModerationResponse
func (client *Client) ImageAsyncModerationWithContext(ctx context.Context, request *ImageAsyncModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageAsyncModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImageAsyncModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImageAsyncModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片批量调用
//
// @param request - ImageBatchModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageBatchModerationResponse
func (client *Client) ImageBatchModerationWithContext(ctx context.Context, request *ImageBatchModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageBatchModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImageBatchModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImageBatchModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Identifies whether an image contains content or elements that violate relevant regulations on network content dissemination, affect the content order of a specific platform, or affect user experience. Image Moderation 2.0 supports over 90 content risk labels and over 100 risk control items. Image Moderation 2.0 of Content Moderation allows you to develop further moderation or governance measures for specific image content based on business scenarios, platform-specific content governance rules, or rich risk labels and scores of confidence levels returned by API calls.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/467826.html)[](https://www.aliyun.com/price/product?#/lvwang/detail/cdibag) of Image Moderation 2.0.
//
// @param request - ImageModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageModerationResponse
func (client *Client) ImageModerationWithContext(ctx context.Context, request *ImageModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImageModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImageModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片审核
//
// @param request - ImageQueueModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageQueueModerationResponse
func (client *Client) ImageQueueModerationWithContext(ctx context.Context, request *ImageQueueModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageQueueModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImageQueueModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImageQueueModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Content Security Manual Review Result Callback Interface
//
// @param request - ManualCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualCallbackResponse
func (client *Client) ManualCallbackWithContext(ctx context.Context, request *ManualCallbackRequest, runtime *dara.RuntimeOptions) (_result *ManualCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.Checksum) {
		body["Checksum"] = request.Checksum
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Msg) {
		body["Msg"] = request.Msg
	}

	if !dara.IsNil(request.ReqId) {
		body["ReqId"] = request.ReqId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManualCallback"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManualCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Content Security Manual Review Request Interface
//
// @param request - ManualModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualModerationResponse
func (client *Client) ManualModerationWithContext(ctx context.Context, request *ManualModerationRequest, runtime *dara.RuntimeOptions) (_result *ManualModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManualModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManualModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve manual review results
//
// @param request - ManualModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualModerationResultResponse
func (client *Client) ManualModerationResultWithContext(ctx context.Context, request *ManualModerationResultRequest, runtime *dara.RuntimeOptions) (_result *ManualModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManualModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManualModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步检测接口
//
// @param request - MultiModalGuardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalGuardResponse
func (client *Client) MultiModalGuardWithContext(ctx context.Context, request *MultiModalGuardRequest, runtime *dara.RuntimeOptions) (_result *MultiModalGuardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MultiModalGuard"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MultiModalGuardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态同步检测接口，支持图片base64字符串
//
// @param request - MultiModalGuardForBase64Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalGuardForBase64Response
func (client *Client) MultiModalGuardForBase64WithContext(ctx context.Context, request *MultiModalGuardForBase64Request, runtime *dara.RuntimeOptions) (_result *MultiModalGuardForBase64Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageBase64Str) {
		body["ImageBase64Str"] = request.ImageBase64Str
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MultiModalGuardForBase64"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MultiModalGuardForBase64Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态-异步检测
//
// @param request - MultimodalAsyncModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultimodalAsyncModerationResponse
func (client *Client) MultimodalAsyncModerationWithContext(ctx context.Context, request *MultimodalAsyncModerationRequest, runtime *dara.RuntimeOptions) (_result *MultimodalAsyncModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MultimodalAsyncModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MultimodalAsyncModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides moderation services for multiple business scenarios and identifies various violation risks.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/464388.html?#section-itm-m2s-ugq) of Text Moderation 2.0.
//
// @param request - TextModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextModerationResponse
func (client *Client) TextModerationWithContext(ctx context.Context, request *TextModerationRequest, runtime *dara.RuntimeOptions) (_result *TextModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TextModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TextModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moderates the input command and generated text of large language models (LLMs). Specific model input commands can be used to retrieve standard answers. The feature of enabling and disabling the moderation labels is also available.
//
// Description:
//
// Before you call this operation, make sure that you have [activated the Content Moderation 2.0 service](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) and are familiar with the [billing](https://help.aliyun.com/document_detail/2671445.html?#section-6od-32j-99n) of the Text Moderation 2.0 Plus service.
//
// @param request - TextModerationPlusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextModerationPlusResponse
func (client *Client) TextModerationPlusWithContext(ctx context.Context, request *TextModerationPlusRequest, runtime *dara.RuntimeOptions) (_result *TextModerationPlusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TextModerationPlus"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TextModerationPlusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// url异步检测
//
// @param request - UrlAsyncModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UrlAsyncModerationResponse
func (client *Client) UrlAsyncModerationWithContext(ctx context.Context, request *UrlAsyncModerationRequest, runtime *dara.RuntimeOptions) (_result *UrlAsyncModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UrlAsyncModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UrlAsyncModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频检测任务提交
//
// @param request - VideoModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoModerationResponse
func (client *Client) VideoModerationWithContext(ctx context.Context, request *VideoModerationRequest, runtime *dara.RuntimeOptions) (_result *VideoModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VideoModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VideoModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消视频直播流检测
//
// @param request - VideoModerationCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoModerationCancelResponse
func (client *Client) VideoModerationCancelWithContext(ctx context.Context, request *VideoModerationCancelRequest, runtime *dara.RuntimeOptions) (_result *VideoModerationCancelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VideoModerationCancel"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VideoModerationCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Video Detection Results
//
// Description:
//
// This operation is free of charge. We recommend that you query moderation results at least 30 seconds after you send an asynchronous moderation request. Content Moderation retains moderation results for at most 24 hours. After 24 hours, the results are deleted.
//
// @param request - VideoModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoModerationResultResponse
func (client *Client) VideoModerationResultWithContext(ctx context.Context, request *VideoModerationResultRequest, runtime *dara.RuntimeOptions) (_result *VideoModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VideoModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VideoModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音审核
//
// @param request - VoiceModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceModerationResponse
func (client *Client) VoiceModerationWithContext(ctx context.Context, request *VoiceModerationRequest, runtime *dara.RuntimeOptions) (_result *VoiceModerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VoiceModeration"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VoiceModerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消检测
//
// @param request - VoiceModerationCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceModerationCancelResponse
func (client *Client) VoiceModerationCancelWithContext(ctx context.Context, request *VoiceModerationCancelRequest, runtime *dara.RuntimeOptions) (_result *VoiceModerationCancelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VoiceModerationCancel"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VoiceModerationCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the moderation results of a Voice Moderation 2.0 task.
//
// @param request - VoiceModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceModerationResultResponse
func (client *Client) VoiceModerationResultWithContext(ctx context.Context, request *VoiceModerationResultRequest, runtime *dara.RuntimeOptions) (_result *VoiceModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		body["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VoiceModerationResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VoiceModerationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
