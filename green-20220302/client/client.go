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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            dara.String("green.aliyuncs.com"),
		"cn-hongkong":           dara.String("green.aliyuncs.com"),
		"cn-huhehaote":          dara.String("green.aliyuncs.com"),
		"cn-qingdao":            dara.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("green.aliyuncs.com"),
		"eu-central-1":          dara.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             dara.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             dara.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 文档审核结果
//
// @param request - DescribeFileModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileModerationResultResponse
func (client *Client) DescribeFileModerationResultWithOptions(request *DescribeFileModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档审核结果
//
// @param request - DescribeFileModerationResultRequest
//
// @return DescribeFileModerationResultResponse
func (client *Client) DescribeFileModerationResult(request *DescribeFileModerationResultRequest) (_result *DescribeFileModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFileModerationResultResponse{}
	_body, _err := client.DescribeFileModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeImageModerationResultWithOptions(request *DescribeImageModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeImageModerationResultResponse
func (client *Client) DescribeImageModerationResult(request *DescribeImageModerationResultRequest) (_result *DescribeImageModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImageModerationResultResponse{}
	_body, _err := client.DescribeImageModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeImageResultExtWithOptions(request *DescribeImageResultExtRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageResultExtResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeImageResultExtResponse
func (client *Client) DescribeImageResultExt(request *DescribeImageResultExtRequest) (_result *DescribeImageResultExtResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImageResultExtResponse{}
	_body, _err := client.DescribeImageResultExtWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeMultimodalModerationResultWithOptions(request *DescribeMultimodalModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultimodalModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeMultimodalModerationResultResponse
func (client *Client) DescribeMultimodalModerationResult(request *DescribeMultimodalModerationResultRequest) (_result *DescribeMultimodalModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMultimodalModerationResultResponse{}
	_body, _err := client.DescribeMultimodalModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询上传token
//
// @param request - DescribeUploadTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUploadTokenResponse
func (client *Client) DescribeUploadTokenWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeUploadTokenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUploadToken"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上传token
//
// @return DescribeUploadTokenResponse
func (client *Client) DescribeUploadToken() (_result *DescribeUploadTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUploadTokenResponse{}
	_body, _err := client.DescribeUploadTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUrlModerationResultWithOptions(request *DescribeUrlModerationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeUrlModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUrlModerationResultResponse
func (client *Client) DescribeUrlModerationResult(request *DescribeUrlModerationResultRequest) (_result *DescribeUrlModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUrlModerationResultResponse{}
	_body, _err := client.DescribeUrlModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FileModerationWithOptions(request *FileModerationRequest, runtime *dara.RuntimeOptions) (_result *FileModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return FileModerationResponse
func (client *Client) FileModeration(request *FileModerationRequest) (_result *FileModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FileModerationResponse{}
	_body, _err := client.FileModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImageAsyncModerationWithOptions(request *ImageAsyncModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageAsyncModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImageAsyncModerationResponse
func (client *Client) ImageAsyncModeration(request *ImageAsyncModerationRequest) (_result *ImageAsyncModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImageAsyncModerationResponse{}
	_body, _err := client.ImageAsyncModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImageBatchModerationWithOptions(request *ImageBatchModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageBatchModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImageBatchModerationResponse
func (client *Client) ImageBatchModeration(request *ImageBatchModerationRequest) (_result *ImageBatchModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImageBatchModerationResponse{}
	_body, _err := client.ImageBatchModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImageModerationWithOptions(request *ImageModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImageModerationResponse
func (client *Client) ImageModeration(request *ImageModerationRequest) (_result *ImageModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImageModerationResponse{}
	_body, _err := client.ImageModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImageQueueModerationWithOptions(request *ImageQueueModerationRequest, runtime *dara.RuntimeOptions) (_result *ImageQueueModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImageQueueModerationResponse
func (client *Client) ImageQueueModeration(request *ImageQueueModerationRequest) (_result *ImageQueueModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImageQueueModerationResponse{}
	_body, _err := client.ImageQueueModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ManualCallbackWithOptions(request *ManualCallbackRequest, runtime *dara.RuntimeOptions) (_result *ManualCallbackResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ManualCallbackResponse
func (client *Client) ManualCallback(request *ManualCallbackRequest) (_result *ManualCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManualCallbackResponse{}
	_body, _err := client.ManualCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ManualModerationWithOptions(request *ManualModerationRequest, runtime *dara.RuntimeOptions) (_result *ManualModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ManualModerationResponse
func (client *Client) ManualModeration(request *ManualModerationRequest) (_result *ManualModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManualModerationResponse{}
	_body, _err := client.ManualModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ManualModerationResultWithOptions(request *ManualModerationResultRequest, runtime *dara.RuntimeOptions) (_result *ManualModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ManualModerationResultResponse
func (client *Client) ManualModerationResult(request *ManualModerationResultRequest) (_result *ManualModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManualModerationResultResponse{}
	_body, _err := client.ManualModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MultiModalGuardWithOptions(request *MultiModalGuardRequest, runtime *dara.RuntimeOptions) (_result *MultiModalGuardResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MultiModalGuardResponse
func (client *Client) MultiModalGuard(request *MultiModalGuardRequest) (_result *MultiModalGuardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultiModalGuardResponse{}
	_body, _err := client.MultiModalGuardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MultiModalGuardForBase64WithOptions(request *MultiModalGuardForBase64Request, runtime *dara.RuntimeOptions) (_result *MultiModalGuardForBase64Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MultiModalGuardForBase64Response
func (client *Client) MultiModalGuardForBase64(request *MultiModalGuardForBase64Request) (_result *MultiModalGuardForBase64Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultiModalGuardForBase64Response{}
	_body, _err := client.MultiModalGuardForBase64WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MultimodalAsyncModerationWithOptions(request *MultimodalAsyncModerationRequest, runtime *dara.RuntimeOptions) (_result *MultimodalAsyncModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MultimodalAsyncModerationResponse
func (client *Client) MultimodalAsyncModeration(request *MultimodalAsyncModerationRequest) (_result *MultimodalAsyncModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultimodalAsyncModerationResponse{}
	_body, _err := client.MultimodalAsyncModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TextModerationWithOptions(request *TextModerationRequest, runtime *dara.RuntimeOptions) (_result *TextModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TextModerationResponse
func (client *Client) TextModeration(request *TextModerationRequest) (_result *TextModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TextModerationResponse{}
	_body, _err := client.TextModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TextModerationPlusWithOptions(request *TextModerationPlusRequest, runtime *dara.RuntimeOptions) (_result *TextModerationPlusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TextModerationPlusResponse
func (client *Client) TextModerationPlus(request *TextModerationPlusRequest) (_result *TextModerationPlusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TextModerationPlusResponse{}
	_body, _err := client.TextModerationPlusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UrlAsyncModerationWithOptions(request *UrlAsyncModerationRequest, runtime *dara.RuntimeOptions) (_result *UrlAsyncModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UrlAsyncModerationResponse
func (client *Client) UrlAsyncModeration(request *UrlAsyncModerationRequest) (_result *UrlAsyncModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UrlAsyncModerationResponse{}
	_body, _err := client.UrlAsyncModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VideoModerationWithOptions(request *VideoModerationRequest, runtime *dara.RuntimeOptions) (_result *VideoModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VideoModerationResponse
func (client *Client) VideoModeration(request *VideoModerationRequest) (_result *VideoModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VideoModerationResponse{}
	_body, _err := client.VideoModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VideoModerationCancelWithOptions(request *VideoModerationCancelRequest, runtime *dara.RuntimeOptions) (_result *VideoModerationCancelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VideoModerationCancelResponse
func (client *Client) VideoModerationCancel(request *VideoModerationCancelRequest) (_result *VideoModerationCancelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VideoModerationCancelResponse{}
	_body, _err := client.VideoModerationCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VideoModerationResultWithOptions(request *VideoModerationResultRequest, runtime *dara.RuntimeOptions) (_result *VideoModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VideoModerationResultResponse
func (client *Client) VideoModerationResult(request *VideoModerationResultRequest) (_result *VideoModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VideoModerationResultResponse{}
	_body, _err := client.VideoModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VoiceModerationWithOptions(request *VoiceModerationRequest, runtime *dara.RuntimeOptions) (_result *VoiceModerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VoiceModerationResponse
func (client *Client) VoiceModeration(request *VoiceModerationRequest) (_result *VoiceModerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VoiceModerationResponse{}
	_body, _err := client.VoiceModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VoiceModerationCancelWithOptions(request *VoiceModerationCancelRequest, runtime *dara.RuntimeOptions) (_result *VoiceModerationCancelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VoiceModerationCancelResponse
func (client *Client) VoiceModerationCancel(request *VoiceModerationCancelRequest) (_result *VoiceModerationCancelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VoiceModerationCancelResponse{}
	_body, _err := client.VoiceModerationCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VoiceModerationResultWithOptions(request *VoiceModerationResultRequest, runtime *dara.RuntimeOptions) (_result *VoiceModerationResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VoiceModerationResultResponse
func (client *Client) VoiceModerationResult(request *VoiceModerationResultRequest) (_result *VoiceModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VoiceModerationResultResponse{}
	_body, _err := client.VoiceModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
