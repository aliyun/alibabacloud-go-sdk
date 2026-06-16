// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
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
// # Document review results
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
// # Document review results
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
// Retrieves the results of an Image Moderation Pro task.
//
// Description:
//
// - Billing information: This operation is not billed.
//
// - QPS limit: This operation is limited to 100 queries per second (QPS) for each user. If you exceed this limit, your API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
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
// Retrieves the results of an Image Moderation Pro task.
//
// Description:
//
// - Billing information: This operation is not billed.
//
// - QPS limit: This operation is limited to 100 queries per second (QPS) for each user. If you exceed this limit, your API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
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
// The enhanced image moderation auxiliary information API operation retrieves additional auxiliary information detected by the enhanced image moderation API operation, including OCR results and custom image library hit information.
//
// Description:
//
// This API operation must be used with the enhanced image moderation API. After you call the enhanced image moderation API operation, you can call this API operation to obtain additional detection information. This API operation is free of charge.
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
// The enhanced image moderation auxiliary information API operation retrieves additional auxiliary information detected by the enhanced image moderation API operation, including OCR results and custom image library hit information.
//
// Description:
//
// This API operation must be used with the enhanced image moderation API. After you call the enhanced image moderation API operation, you can call this API operation to obtain additional detection information. This API operation is free of charge.
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
// Query the results of an asynchronous multimodal moderation task.
//
// Description:
//
// - Billing information: This API call is free.
//
// - Query timeout: Wait 30 seconds after you submit an asynchronous moderation task before querying the result. Do not wait longer than 24 hours, or the result will be automatically deleted.
//
// - This API has a per-user rate limiting limit of 10 requests per second. Exceeding this limit triggers rate limiting, which may affect your service. Call the API responsibly.
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
// Query the results of an asynchronous multimodal moderation task.
//
// Description:
//
// - Billing information: This API call is free.
//
// - Query timeout: Wait 30 seconds after you submit an asynchronous moderation task before querying the result. Do not wait longer than 24 hours, or the result will be automatically deleted.
//
// - This API has a per-user rate limiting limit of 10 requests per second. Exceeding this limit triggers rate limiting, which may affect your service. Call the API responsibly.
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
// Retrieves an upload token.
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
// Retrieves an upload token.
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
// Queries moderation results based on the ReqId returned by asynchronous URL moderation.
//
// Description:
//
// - Billing information: This operation is free of charge.
//
// - Query timeout: We recommend that you set the query interval to 480 seconds (query the results 480 seconds after you submit the asynchronous moderation task). The maximum timeout period is 3 days. After this period, the results are automatically deleted.
//
// - The QPS limit for this operation is 100 queries per second (QPS) per user. If the limit is exceeded, your API calls will be throttled, which may affect your business. Make sure you call the operation at a reasonable rate.
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
// Queries moderation results based on the ReqId returned by asynchronous URL moderation.
//
// Description:
//
// - Billing information: This operation is free of charge.
//
// - Query timeout: We recommend that you set the query interval to 480 seconds (query the results 480 seconds after you submit the asynchronous moderation task). The maximum timeout period is 3 days. After this period, the results are automatically deleted.
//
// - The QPS limit for this operation is 100 queries per second (QPS) per user. If the limit is exceeded, your API calls will be throttled, which may affect your business. Make sure you call the operation at a reasonable rate.
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
// Moderates document content.
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
// Moderates document content.
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
// This API is used for asynchronous image moderation. Asynchronous moderation tasks do not return detection results in real time. You can obtain the detection results using a callback or by polling. The detection results are retained for up to three days.
//
// Description:
//
// - The following image formats are supported: PNG, JPG, JPEG, BMP, WEBP, TIFF, ICO, HEIC, and SVG.
//
// - The image size cannot exceed 10 MB. The recommended image resolution is greater than 200 × 200 pixels. A low resolution may compromise the accuracy of the Content Moderation algorithm.
//
// - The timeout period for image downloads is 3 seconds. If an image download exceeds this duration, a download timeout error is returned.
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
// This API is used for asynchronous image moderation. Asynchronous moderation tasks do not return detection results in real time. You can obtain the detection results using a callback or by polling. The detection results are retained for up to three days.
//
// Description:
//
// - The following image formats are supported: PNG, JPG, JPEG, BMP, WEBP, TIFF, ICO, HEIC, and SVG.
//
// - The image size cannot exceed 10 MB. The recommended image resolution is greater than 200 × 200 pixels. A low resolution may compromise the accuracy of the Content Moderation algorithm.
//
// - The timeout period for image downloads is 3 seconds. If an image download exceeds this duration, a download timeout error is returned.
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
// # Batch Invocation of Images
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
// # Batch Invocation of Images
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
// # Image moderation
//
// Description:
//
// Before you call this operation, complete the following steps:
//
// 1. [Activate AI Guardrails-Enhanced Edition](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn).
//
// 2. Understand the [billing methods and pricing](https://help.aliyun.com/document_detail/467826.html?#section-h06-qz6-1pt) of the enhanced image moderation feature.
//
// 3. For more information about API usage and parameters, see the [API reference](https://help.aliyun.com/document_detail/467829.html).
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
// # Image moderation
//
// Description:
//
// Before you call this operation, complete the following steps:
//
// 1. [Activate AI Guardrails-Enhanced Edition](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn).
//
// 2. Understand the [billing methods and pricing](https://help.aliyun.com/document_detail/467826.html?#section-h06-qz6-1pt) of the enhanced image moderation feature.
//
// 3. For more information about API usage and parameters, see the [API reference](https://help.aliyun.com/document_detail/467829.html).
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
// The channel callback API for manual review results in Content Moderation.
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
// The channel callback API for manual review results in Content Moderation.
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
// # Interface for submitting Content Moderation manual review requests
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
// # Interface for submitting Content Moderation manual review requests
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
// Retrieves the manual moderation result.
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
// Retrieves the manual moderation result.
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
// This is the synchronous detection API for the multi-modal agent.
//
// Description:
//
// This is the AI Guardrails agent.
//
// @param request - MultiModalAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalAgentResponse
func (client *Client) MultiModalAgentWithOptions(request *MultiModalAgentRequest, runtime *dara.RuntimeOptions) (_result *MultiModalAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppID) {
		body["AppID"] = request.AppID
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MultiModalAgent"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MultiModalAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This is the synchronous detection API for the multi-modal agent.
//
// Description:
//
// This is the AI Guardrails agent.
//
// @param request - MultiModalAgentRequest
//
// @return MultiModalAgentResponse
func (client *Client) MultiModalAgent(request *MultiModalAgentRequest) (_result *MultiModalAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultiModalAgentResponse{}
	_body, _err := client.MultiModalAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # API for synchronous detection
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
// # API for synchronous detection
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
// An asynchronous multimodal AI safety guardrail API for audio and video. It provides comprehensive detection of non-compliant content, sensitive content, prompt injection attacks, malicious files, and malicious URLs.
//
// Description:
//
// If an API is subject to billing, add the following sentence in bold: "Before using this API, ensure that you fully understand the billing methods and pricing of the XXX product." The word "pricing" must be a hyperlink to https\\://www\\.aliyun.com/price/product#/ecs/detail.
//
// @param request - MultiModalGuardAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalGuardAsyncResponse
func (client *Client) MultiModalGuardAsyncWithOptions(request *MultiModalGuardAsyncRequest, runtime *dara.RuntimeOptions) (_result *MultiModalGuardAsyncResponse, _err error) {
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
		Action:      dara.String("MultiModalGuardAsync"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MultiModalGuardAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An asynchronous multimodal AI safety guardrail API for audio and video. It provides comprehensive detection of non-compliant content, sensitive content, prompt injection attacks, malicious files, and malicious URLs.
//
// Description:
//
// If an API is subject to billing, add the following sentence in bold: "Before using this API, ensure that you fully understand the billing methods and pricing of the XXX product." The word "pricing" must be a hyperlink to https\\://www\\.aliyun.com/price/product#/ecs/detail.
//
// @param request - MultiModalGuardAsyncRequest
//
// @return MultiModalGuardAsyncResponse
func (client *Client) MultiModalGuardAsync(request *MultiModalGuardAsyncRequest) (_result *MultiModalGuardAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultiModalGuardAsyncResponse{}
	_body, _err := client.MultiModalGuardAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This AI Security Guardrail API retrieves asynchronous multimodal results from both audio and video.
//
// Description:
//
// For APIs that incur charges, add the following sentence in bold at the beginning of the description: "Before you use this API, make sure that you fully understand the billing methods and pricing of the XXX product." Link the word \\"pricing\\" to https\\://www\\.aliyun.com/price/product#/ecs/detail.
//
// @param request - MultiModalGuardAsyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalGuardAsyncResultResponse
func (client *Client) MultiModalGuardAsyncResultWithOptions(request *MultiModalGuardAsyncResultRequest, runtime *dara.RuntimeOptions) (_result *MultiModalGuardAsyncResultResponse, _err error) {
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
		Action:      dara.String("MultiModalGuardAsyncResult"),
		Version:     dara.String("2022-03-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MultiModalGuardAsyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This AI Security Guardrail API retrieves asynchronous multimodal results from both audio and video.
//
// Description:
//
// For APIs that incur charges, add the following sentence in bold at the beginning of the description: "Before you use this API, make sure that you fully understand the billing methods and pricing of the XXX product." Link the word \\"pricing\\" to https\\://www\\.aliyun.com/price/product#/ecs/detail.
//
// @param request - MultiModalGuardAsyncResultRequest
//
// @return MultiModalGuardAsyncResultResponse
func (client *Client) MultiModalGuardAsyncResult(request *MultiModalGuardAsyncResultRequest) (_result *MultiModalGuardAsyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultiModalGuardAsyncResultResponse{}
	_body, _err := client.MultiModalGuardAsyncResultWithOptions(request, runtime)
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
// Provides a WebSocket-based multimodal detection API for AI safety guardrails. This API supports content compliance detection, sensitive content detection, prompt attack detection, malicious file detection, malicious URL detection, and other comprehensive detection capabilities.
//
// @param request - MultiModalGuardWsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalGuardWsResponse
func (client *Client) MultiModalGuardWsWithOptions(request *MultiModalGuardWsRequest, runtime *dara.RuntimeOptions) (_result *MultiModalGuardWsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:               dara.String("MultiModalGuardWs"),
		Version:              dara.String("2022-03-02"),
		Protocol:             dara.String("wss"),
		Pathname:             dara.String("/"),
		Method:               dara.String("GET"),
		AuthType:             dara.String("AK"),
		Style:                dara.String("RPC"),
		ReqBodyType:          dara.String("formData"),
		BodyType:             dara.String("json"),
		WebsocketSubProtocol: dara.String("awap"),
	}
	res := &MultiModalGuardWsResponse{}
	callApiTmp, err := client.CallApi(params, req, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}
	tmp := dara.ToMap(callApiTmp)
	if !dara.IsNil(tmp["WebSocketClient"]) {
		res.WebSocketClient = websocketutils.CreateWebSocketClient(tmp["WebSocketClient"])
	}

	_result = res
	return _result, _err
}

// Summary:
//
// Provides a WebSocket-based multimodal detection API for AI safety guardrails. This API supports content compliance detection, sensitive content detection, prompt attack detection, malicious file detection, malicious URL detection, and other comprehensive detection capabilities.
//
// @param request - MultiModalGuardWsRequest
//
// @return MultiModalGuardWsResponse
func (client *Client) MultiModalGuardWs(request *MultiModalGuardWsRequest) (_result *MultiModalGuardWsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MultiModalGuardWsResponse{}
	_body, _err := client.MultiModalGuardWsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Multimodal-Asynchronous Detection
//
// Description:
//
// The asynchronous URL moderation service supports two billing methods: pay-as-you-go and resource plan usage.
//
// - After you activate the enhanced text moderation service, the default billing method is pay-as-you-go. You are billed daily based on actual usage. No charges apply if you do not invoke the service.
//
// - If your moderation volume is large or your moderation needs are relatively stable, purchase a resource plan in advance. Larger resource plans offer greater discounts. You can stack multiple resource plans.
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
// # Multimodal-Asynchronous Detection
//
// Description:
//
// The asynchronous URL moderation service supports two billing methods: pay-as-you-go and resource plan usage.
//
// - After you activate the enhanced text moderation service, the default billing method is pay-as-you-go. You are billed daily based on actual usage. No charges apply if you do not invoke the service.
//
// - If your moderation volume is large or your moderation needs are relatively stable, purchase a resource plan in advance. Larger resource plans offer greater discounts. You can stack multiple resource plans.
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
// This service uses dynamic policies and models to defend against adversarial content. It provides moderation services for various business scenarios and detects different types of violations.
//
// Description:
//
// Before you use this operation, review the [billing methods and pricing](https://help.aliyun.com/document_detail/464388.html?#section-itm-m2s-ugq) for Text Moderation Plus.
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
// This service uses dynamic policies and models to defend against adversarial content. It provides moderation services for various business scenarios and detects different types of violations.
//
// Description:
//
// Before you use this operation, review the [billing methods and pricing](https://help.aliyun.com/document_detail/464388.html?#section-itm-m2s-ugq) for Text Moderation Plus.
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
// Text Moderation Plus is an upgraded service that moderates the input instructions and generated text of large language models (LLMs). This service can retrieve standard answers for specific input instructions and lets you enable or disable moderation labels.
//
// Description:
//
// Before you use this API, [activate AI Guardrails Pro](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) and make sure that you understand the [billing methods and pricing](https://help.aliyun.com/document_detail/2671445.html?#section-6od-32j-99n) for Text Moderation Plus.
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
// Text Moderation Plus is an upgraded service that moderates the input instructions and generated text of large language models (LLMs). This service can retrieve standard answers for specific input instructions and lets you enable or disable moderation labels.
//
// Description:
//
// Before you use this API, [activate AI Guardrails Pro](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) and make sure that you understand the [billing methods and pricing](https://help.aliyun.com/document_detail/2671445.html?#section-6od-32j-99n) for Text Moderation Plus.
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
// The URL asynchronous moderation service detects threats such as fraud, pornography, and gambling in URLs to protect the content ecosystem of your platform.
//
// Description:
//
// The URL asynchronous moderation service supports the pay-as-you-go and resource plan billing methods.
//
// - After you activate the enhanced edition of Text Moderation, the default billing method is pay-as-you-go. You are charged CNY 30 per 10,000 calls based on your daily usage. No fees are incurred if you do not call the service.
//
// - If you have many moderation requests or relatively fixed moderation requirements, we recommend that you purchase resource plans in advance. The larger the resource plan you purchase, the greater the discount you receive. You can purchase and use multiple resource plans.
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
// The URL asynchronous moderation service detects threats such as fraud, pornography, and gambling in URLs to protect the content ecosystem of your platform.
//
// Description:
//
// The URL asynchronous moderation service supports the pay-as-you-go and resource plan billing methods.
//
// - After you activate the enhanced edition of Text Moderation, the default billing method is pay-as-you-go. You are charged CNY 30 per 10,000 calls based on your daily usage. No fees are incurred if you do not call the service.
//
// - If you have many moderation requests or relatively fixed moderation requirements, we recommend that you purchase resource plans in advance. The larger the resource plan you purchase, the greater the discount you receive. You can purchase and use multiple resource plans.
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
// The enhanced video moderation feature of Content Moderation detects threats and non-compliant content in video files. Use this operation to submit a moderation task.
//
// Description:
//
// Before you call this operation, make sure that you have activated the [enhanced Content Moderation](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) service and understand the [billing methods](https://help.aliyun.com/document_detail/2505807.html) and [pricing](https://www.aliyun.com/price/product?#/lvwang/detail/cdibag) of the enhanced video moderation feature.
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
// The enhanced video moderation feature of Content Moderation detects threats and non-compliant content in video files. Use this operation to submit a moderation task.
//
// Description:
//
// Before you call this operation, make sure that you have activated the [enhanced Content Moderation](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) service and understand the [billing methods](https://help.aliyun.com/document_detail/2505807.html) and [pricing](https://www.aliyun.com/price/product?#/lvwang/detail/cdibag) of the enhanced video moderation feature.
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
// Cancels an ApsaraVideo Live moderation task.
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
// Cancels an ApsaraVideo Live moderation task.
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
// Retrieves the results of an enhanced video moderation task.
//
// Description:
//
// This operation is free. We recommend querying for the result 30 seconds after submitting the asynchronous detection task. You must retrieve the result within 24 hours, or it will be automatically deleted.
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
// Retrieves the results of an enhanced video moderation task.
//
// Description:
//
// This operation is free. We recommend querying for the result 30 seconds after submitting the asynchronous detection task. You must retrieve the result within 24 hours, or it will be automatically deleted.
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
// Submits a task for enhanced voice moderation.
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
// Submits a task for enhanced voice moderation.
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
// This operation cancels an enhanced voice moderation task.
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
// This operation cancels an enhanced voice moderation task.
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
// Retrieve the detection results for enhanced voice moderation.
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
// Retrieve the detection results for enhanced voice moderation.
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
