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
		"cn-shanghai":    dara.String("yike.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1": dara.String("yike.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("yike"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Retrieves information about multiple media assets in a batch.
//
// Description:
//
// ## Request description.
//
// @param request - BatchGetMediasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetMediasResponse
func (client *Client) BatchGetMediasWithOptions(request *BatchGetMediasRequest, runtime *dara.RuntimeOptions) (_result *BatchGetMediasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetMedias"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetMediasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about multiple media assets in a batch.
//
// Description:
//
// ## Request description.
//
// @param request - BatchGetMediasRequest
//
// @return BatchGetMediasResponse
func (client *Client) BatchGetMedias(request *BatchGetMediasRequest) (_result *BatchGetMediasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetMediasResponse{}
	_body, _err := client.BatchGetMediasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a media asset category.
//
// Description:
//
// Categories support up to three levels, and each level supports up to 100 subcategories.
//
// @param request - CreateAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssetCategoryResponse
func (client *Client) CreateAssetCategoryWithOptions(request *CreateAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *CreateAssetCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAssetCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a media asset category.
//
// Description:
//
// Categories support up to three levels, and each level supports up to 100 subcategories.
//
// @param request - CreateAssetCategoryRequest
//
// @return CreateAssetCategoryResponse
func (client *Client) CreateAssetCategory(request *CreateAssetCategoryRequest) (_result *CreateAssetCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAssetCategoryResponse{}
	_body, _err := client.CreateAssetCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a media asset category.
//
// Description:
//
// This operation also deletes all subcategories (including second-level and third-level categories). Proceed with caution.
//
// @param request - DeleteAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAssetCategoryResponse
func (client *Client) DeleteAssetCategoryWithOptions(request *DeleteAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteAssetCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAssetCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a media asset category.
//
// Description:
//
// This operation also deletes all subcategories (including second-level and third-level categories). Proceed with caution.
//
// @param request - DeleteAssetCategoryRequest
//
// @return DeleteAssetCategoryResponse
func (client *Client) DeleteAssetCategory(request *DeleteAssetCategoryRequest) (_result *DeleteAssetCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAssetCategoryResponse{}
	_body, _err := client.DeleteAssetCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes media asset information.
//
// @param request - DeleteMediasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediasResponse
func (client *Client) DeleteMediasWithOptions(request *DeleteMediasRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletePhysicalFiles) {
		query["DeletePhysicalFiles"] = request.DeletePhysicalFiles
	}

	if !dara.IsNil(request.InputURLs) {
		query["InputURLs"] = request.InputURLs
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMedias"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes media asset information.
//
// @param request - DeleteMediasRequest
//
// @return DeleteMediasResponse
func (client *Client) DeleteMedias(request *DeleteMediasRequest) (_result *DeleteMediasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediasResponse{}
	_body, _err := client.DeleteMediasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified category and the list of its subcategories (immediate child categories).
//
// @param request - GetAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssetCategoryResponse
func (client *Client) GetAssetCategoryWithOptions(request *GetAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetAssetCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssetCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified category and the list of its subcategories (immediate child categories).
//
// @param request - GetAssetCategoryRequest
//
// @return GetAssetCategoryResponse
func (client *Client) GetAssetCategory(request *GetAssetCategoryRequest) (_result *GetAssetCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAssetCategoryResponse{}
	_body, _err := client.GetAssetCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an image generation task.
//
// @param request - GetImageGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageGenerationJobResponse
func (client *Client) GetImageGenerationJobWithOptions(request *GetImageGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *GetImageGenerationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageGenerationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an image generation task.
//
// @param request - GetImageGenerationJobRequest
//
// @return GetImageGenerationJobResponse
func (client *Client) GetImageGenerationJob(request *GetImageGenerationJobRequest) (_result *GetImageGenerationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageGenerationJobResponse{}
	_body, _err := client.GetImageGenerationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询媒资
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
//
// @param request - GetMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaResponse
func (client *Client) GetMediaWithOptions(request *GetMediaRequest, runtime *dara.RuntimeOptions) (_result *GetMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMedia"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询媒资
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
//
// @param request - GetMediaRequest
//
// @return GetMediaResponse
func (client *Client) GetMedia(request *GetMediaRequest) (_result *GetMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaResponse{}
	_body, _err := client.GetMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a media asset content understanding job.
//
// Description:
//
// ## Description
//
// This API is used to query a media asset content understanding job.
//
// @param request - GetMediaComprehensionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaComprehensionJobResponse
func (client *Client) GetMediaComprehensionJobWithOptions(request *GetMediaComprehensionJobRequest, runtime *dara.RuntimeOptions) (_result *GetMediaComprehensionJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaComprehensionJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaComprehensionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a media asset content understanding job.
//
// Description:
//
// ## Description
//
// This API is used to query a media asset content understanding job.
//
// @param request - GetMediaComprehensionJobRequest
//
// @return GetMediaComprehensionJobResponse
func (client *Client) GetMediaComprehensionJob(request *GetMediaComprehensionJobRequest) (_result *GetMediaComprehensionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaComprehensionJobResponse{}
	_body, _err := client.GetMediaComprehensionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a creative script generation task.
//
// @param request - GetRemakeScriptJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRemakeScriptJobResponse
func (client *Client) GetRemakeScriptJobWithOptions(request *GetRemakeScriptJobRequest, runtime *dara.RuntimeOptions) (_result *GetRemakeScriptJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRemakeScriptJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRemakeScriptJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a creative script generation task.
//
// @param request - GetRemakeScriptJobRequest
//
// @return GetRemakeScriptJobResponse
func (client *Client) GetRemakeScriptJob(request *GetRemakeScriptJobRequest) (_result *GetRemakeScriptJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRemakeScriptJobResponse{}
	_body, _err := client.GetRemakeScriptJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a video generation task.
//
// @param request - GetVideoGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoGenerationJobResponse
func (client *Client) GetVideoGenerationJobWithOptions(request *GetVideoGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *GetVideoGenerationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoGenerationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a video generation task.
//
// @param request - GetVideoGenerationJobRequest
//
// @return GetVideoGenerationJobResponse
func (client *Client) GetVideoGenerationJob(request *GetVideoGenerationJobRequest) (_result *GetVideoGenerationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoGenerationJobResponse{}
	_body, _err := client.GetVideoGenerationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询爆款新视频渲染任务
//
// @param request - GetVideoRenderJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoRenderJobResponse
func (client *Client) GetVideoRenderJobWithOptions(request *GetVideoRenderJobRequest, runtime *dara.RuntimeOptions) (_result *GetVideoRenderJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoRenderJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoRenderJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询爆款新视频渲染任务
//
// @param request - GetVideoRenderJobRequest
//
// @return GetVideoRenderJobResponse
func (client *Client) GetVideoRenderJob(request *GetVideoRenderJobRequest) (_result *GetVideoRenderJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoRenderJobResponse{}
	_body, _err := client.GetVideoRenderJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the membership plan and credit information for a Yike primary account.
//
// @param request - GetYikeAccountCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAccountCreditResponse
func (client *Client) GetYikeAccountCreditWithOptions(request *GetYikeAccountCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAccountCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeAccountCredit"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeAccountCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the membership plan and credit information for a Yike primary account.
//
// @param request - GetYikeAccountCreditRequest
//
// @return GetYikeAccountCreditResponse
func (client *Client) GetYikeAccountCredit(request *GetYikeAccountCreditRequest) (_result *GetYikeAccountCreditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeAccountCreditResponse{}
	_body, _err := client.GetYikeAccountCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the actual credit consumption of a task.
//
// @param request - GetYikeJobCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeJobCreditResponse
func (client *Client) GetYikeJobCreditWithOptions(request *GetYikeJobCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeJobCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeJobCredit"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeJobCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the actual credit consumption of a task.
//
// @param request - GetYikeJobCreditRequest
//
// @return GetYikeJobCreditResponse
func (client *Client) GetYikeJobCredit(request *GetYikeJobCreditRequest) (_result *GetYikeJobCreditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeJobCreditResponse{}
	_body, _err := client.GetYikeJobCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a media asset.
//
// Description:
//
// ## Operation description
//
// This API is used to query media content understanding jobs.
//
// @param request - ImportMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportMediaResponse
func (client *Client) ImportMediaWithOptions(request *ImportMediaRequest, runtime *dara.RuntimeOptions) (_result *ImportMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicMetaData) {
		query["DynamicMetaData"] = request.DynamicMetaData
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ImportSource) {
		query["ImportSource"] = request.ImportSource
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaTags) {
		query["MediaTags"] = request.MediaTags
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.RegisterConfig) {
		query["RegisterConfig"] = request.RegisterConfig
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportMedia"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a media asset.
//
// Description:
//
// ## Operation description
//
// This API is used to query media content understanding jobs.
//
// @param request - ImportMediaRequest
//
// @return ImportMediaResponse
func (client *Client) ImportMedia(request *ImportMediaRequest) (_result *ImportMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportMediaResponse{}
	_body, _err := client.ImportMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of categories.
//
// @param request - ListAssetCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssetCategoriesResponse
func (client *Client) ListAssetCategoriesWithOptions(request *ListAssetCategoriesRequest, runtime *dara.RuntimeOptions) (_result *ListAssetCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAssetCategories"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAssetCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of categories.
//
// @param request - ListAssetCategoriesRequest
//
// @return ListAssetCategoriesResponse
func (client *Client) ListAssetCategories(request *ListAssetCategoriesRequest) (_result *ListAssetCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAssetCategoriesResponse{}
	_body, _err := client.ListAssetCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns media asset information that matches the specified filter conditions.
//
// @param request - SearchMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaResponse
func (client *Client) SearchMediaWithOptions(request *SearchMediaRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Match) {
		query["Match"] = request.Match
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScrollToken) {
		query["ScrollToken"] = request.ScrollToken
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMedia"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns media asset information that matches the specified filter conditions.
//
// @param request - SearchMediaRequest
//
// @return SearchMediaResponse
func (client *Client) SearchMedia(request *SearchMediaRequest) (_result *SearchMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaResponse{}
	_body, _err := client.SearchMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an image generation task.
//
// @param request - SubmitImageGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageGenerationJobResponse
func (client *Client) SubmitImageGenerationJobWithOptions(request *SubmitImageGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitImageGenerationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AspectRatio) {
		query["AspectRatio"] = request.AspectRatio
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParameters) {
		query["JobParameters"] = request.JobParameters
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		query["N"] = request.N
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImageGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageGenerationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an image generation task.
//
// @param request - SubmitImageGenerationJobRequest
//
// @return SubmitImageGenerationJobResponse
func (client *Client) SubmitImageGenerationJob(request *SubmitImageGenerationJobRequest) (_result *SubmitImageGenerationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitImageGenerationJobResponse{}
	_body, _err := client.SubmitImageGenerationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交媒资内容理解作业
//
// Description:
//
// ## 请求说明
//
// 该API用于根据提供的媒资文件（比如视频链接）进行内容理解。此外，支持通过`UserData`字段传递自定义参数，在回调时原样返回。
//
// @param request - SubmitMediaComprehensionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaComprehensionJobResponse
func (client *Client) SubmitMediaComprehensionJobWithOptions(request *SubmitMediaComprehensionJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaComprehensionJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParams) {
		query["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaComprehensionJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaComprehensionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交媒资内容理解作业
//
// Description:
//
// ## 请求说明
//
// 该API用于根据提供的媒资文件（比如视频链接）进行内容理解。此外，支持通过`UserData`字段传递自定义参数，在回调时原样返回。
//
// @param request - SubmitMediaComprehensionJobRequest
//
// @return SubmitMediaComprehensionJobResponse
func (client *Client) SubmitMediaComprehensionJob(request *SubmitMediaComprehensionJobRequest) (_result *SubmitMediaComprehensionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMediaComprehensionJobResponse{}
	_body, _err := client.SubmitMediaComprehensionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交爆款复刻新脚本生成任务
//
// Description:
//
// 该 API 用于根据内容理解的结果与新商品/模特信息，仿写生成新的口播脚本。此外，支持通过UserData字段传递自定义参数，在回调时原样返回。
//
// @param request - SubmitRemakeScriptJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitRemakeScriptJobResponse
func (client *Client) SubmitRemakeScriptJobWithOptions(request *SubmitRemakeScriptJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitRemakeScriptJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RemakeParams) {
		query["RemakeParams"] = request.RemakeParams
	}

	if !dara.IsNil(request.RemakeType) {
		query["RemakeType"] = request.RemakeType
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitRemakeScriptJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitRemakeScriptJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交爆款复刻新脚本生成任务
//
// Description:
//
// 该 API 用于根据内容理解的结果与新商品/模特信息，仿写生成新的口播脚本。此外，支持通过UserData字段传递自定义参数，在回调时原样返回。
//
// @param request - SubmitRemakeScriptJobRequest
//
// @return SubmitRemakeScriptJobResponse
func (client *Client) SubmitRemakeScriptJob(request *SubmitRemakeScriptJobRequest) (_result *SubmitRemakeScriptJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitRemakeScriptJobResponse{}
	_body, _err := client.SubmitRemakeScriptJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a video generation task.
//
// @param request - SubmitVideoGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoGenerationJobResponse
func (client *Client) SubmitVideoGenerationJobWithOptions(request *SubmitVideoGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoGenerationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AspectRatio) {
		query["AspectRatio"] = request.AspectRatio
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParameters) {
		query["JobParameters"] = request.JobParameters
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		query["N"] = request.N
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoGenerationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video generation task.
//
// @param request - SubmitVideoGenerationJobRequest
//
// @return SubmitVideoGenerationJobResponse
func (client *Client) SubmitVideoGenerationJob(request *SubmitVideoGenerationJobRequest) (_result *SubmitVideoGenerationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitVideoGenerationJobResponse{}
	_body, _err := client.SubmitVideoGenerationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交爆款新视频渲染任务
//
// @param request - SubmitVideoRenderJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoRenderJobResponse
func (client *Client) SubmitVideoRenderJobWithOptions(request *SubmitVideoRenderJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoRenderJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Script) {
		query["Script"] = request.Script
	}

	if !dara.IsNil(request.Settings) {
		query["Settings"] = request.Settings
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoRenderJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoRenderJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交爆款新视频渲染任务
//
// @param request - SubmitVideoRenderJobRequest
//
// @return SubmitVideoRenderJobResponse
func (client *Client) SubmitVideoRenderJob(request *SubmitVideoRenderJobRequest) (_result *SubmitVideoRenderJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitVideoRenderJobResponse{}
	_body, _err := client.SubmitVideoRenderJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a video translation task that supports subtitle translation, voice translation, and on-screen text translation.
//
// Description:
//
// ## Request description
//
// - This API supports multiple video translation features, including subtitle translation and voice translation.
//
// - The `JobType` parameter defines the task type, such as `SubtitleTranslate` and `VoiceTranslate`.
//
// - The `Input` and `Output` parameters specify the input resource and output path, respectively.
//
// - `JobParameters` contains language configuration and other feature switches, such as `SourceLanguage`, `TargetLanguage`, `NeedDetext`, and `NeedVisualTranslate`.
//
// - `EditingConfig` can be used to specify the style configuration for the final editing and compositing.
//
// - `ClientToken` is an optional parameter used to ensure the idempotence of the request.
//
// - Ensure that all required fields are correctly filled in. Otherwise, the request may fail.
//
// @param request - SubmitVideoTranslationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoTranslationJobResponse
func (client *Client) SubmitVideoTranslationJobWithOptions(request *SubmitVideoTranslationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoTranslationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Input) {
		body["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParameters) {
		body["JobParameters"] = request.JobParameters
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Output) {
		body["Output"] = request.Output
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoTranslationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoTranslationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video translation task that supports subtitle translation, voice translation, and on-screen text translation.
//
// Description:
//
// ## Request description
//
// - This API supports multiple video translation features, including subtitle translation and voice translation.
//
// - The `JobType` parameter defines the task type, such as `SubtitleTranslate` and `VoiceTranslate`.
//
// - The `Input` and `Output` parameters specify the input resource and output path, respectively.
//
// - `JobParameters` contains language configuration and other feature switches, such as `SourceLanguage`, `TargetLanguage`, `NeedDetext`, and `NeedVisualTranslate`.
//
// - `EditingConfig` can be used to specify the style configuration for the final editing and compositing.
//
// - `ClientToken` is an optional parameter used to ensure the idempotence of the request.
//
// - Ensure that all required fields are correctly filled in. Otherwise, the request may fail.
//
// @param request - SubmitVideoTranslationJobRequest
//
// @return SubmitVideoTranslationJobResponse
func (client *Client) SubmitVideoTranslationJob(request *SubmitVideoTranslationJobRequest) (_result *SubmitVideoTranslationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitVideoTranslationJobResponse{}
	_body, _err := client.SubmitVideoTranslationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a media asset category.
//
// Description:
//
// After you create a media asset category, you can call this operation to locate and update the name of the media asset category by category ID.
//
// @param request - UpdateAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssetCategoryResponse
func (client *Client) UpdateAssetCategoryWithOptions(request *UpdateAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateAssetCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAssetCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a media asset category.
//
// Description:
//
// After you create a media asset category, you can call this operation to locate and update the name of the media asset category by category ID.
//
// @param request - UpdateAssetCategoryRequest
//
// @return UpdateAssetCategoryResponse
func (client *Client) UpdateAssetCategory(request *UpdateAssetCategoryRequest) (_result *UpdateAssetCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAssetCategoryResponse{}
	_body, _err := client.UpdateAssetCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates media asset information.
//
// Description:
//
// ## Request description
//
// This API is used to query media content understanding jobs.
//
// @param request - UpdateMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaResponse
func (client *Client) UpdateMediaWithOptions(request *UpdateMediaRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppendTags) {
		query["AppendTags"] = request.AppendTags
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicMetaData) {
		query["DynamicMetaData"] = request.DynamicMetaData
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaTags) {
		query["MediaTags"] = request.MediaTags
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMedia"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates media asset information.
//
// Description:
//
// ## Request description
//
// This API is used to query media content understanding jobs.
//
// @param request - UpdateMediaRequest
//
// @return UpdateMediaResponse
func (client *Client) UpdateMedia(request *UpdateMediaRequest) (_result *UpdateMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaResponse{}
	_body, _err := client.UpdateMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
