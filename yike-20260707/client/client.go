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
// 批量获取媒资信息
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// 批量获取媒资信息
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// 创建分类
//
// Description:
//
// 分类最多支持三级分类，每级分类最多支持创建 100 个子分类。
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
// 创建分类
//
// Description:
//
// 分类最多支持三级分类，每级分类最多支持创建 100 个子分类。
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
// 删除分类
//
// Description:
//
// 此接口会同时删除其子分类（包括二级分类和三级分类），请慎重操作。
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
// 删除分类
//
// Description:
//
// 此接口会同时删除其子分类（包括二级分类和三级分类），请慎重操作。
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
// 删除媒资信息
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
// 删除媒资信息
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
// 查询分类
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
// 查询分类
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
// 查询图片生成任务
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
// 查询图片生成任务
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
// 查询媒资内容理解作业
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// 查询媒资内容理解作业
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// 查询视频生成任务
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
// 查询视频生成任务
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
// 获取一刻主账户会员计划及积分情况
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
// 获取一刻主账户会员计划及积分情况
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
// 查询一刻任务实际消耗积分
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
// 查询一刻任务实际消耗积分
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
// 导入媒资
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// 导入媒资
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// 列出分类
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
// 列出分类
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
// 搜索媒资
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
// 搜索媒资
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
// 提交图像生成接口
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
// 提交图像生成接口
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
	if !dara.IsNil(request.JobParams) {
		query["JobParams"] = request.JobParams
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
// 提交视频生成接口
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
// 提交视频生成接口
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
// 提交视频翻译任务
//
// Description:
//
// ## 请求说明
//
// - 该 API 支持多种视频翻译功能，包括字幕翻译和声音翻译。
//
// - `JobType` 参数定义了任务类型，如 `SubtitleTranslate`和`VoiceTranslate` 。
//
// - `Input` 和 `Output` 参数分别指定了输入资源和输出路径。
//
// - `JobParameters` 包含了语言配置和其他能力开关，如 `SourceLanguage`、`TargetLanguage`、`NeedDetext` 和 `NeedVisualTranslate` 等。
//
// - `EditingConfig` 可以用来指定最终剪辑合成的样式配置。
//
// - `ClientToken` 是一个可选参数，用于保证请求的幂等性。
//
// - 请确保所有必填字段都已正确填写，否则可能会导致请求失败。
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
// 提交视频翻译任务
//
// Description:
//
// ## 请求说明
//
// - 该 API 支持多种视频翻译功能，包括字幕翻译和声音翻译。
//
// - `JobType` 参数定义了任务类型，如 `SubtitleTranslate`和`VoiceTranslate` 。
//
// - `Input` 和 `Output` 参数分别指定了输入资源和输出路径。
//
// - `JobParameters` 包含了语言配置和其他能力开关，如 `SourceLanguage`、`TargetLanguage`、`NeedDetext` 和 `NeedVisualTranslate` 等。
//
// - `EditingConfig` 可以用来指定最终剪辑合成的样式配置。
//
// - `ClientToken` 是一个可选参数，用于保证请求的幂等性。
//
// - 请确保所有必填字段都已正确填写，否则可能会导致请求失败。
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
// 更新媒资分类
//
// Description:
//
// 创建媒资分类后，可调用本接口通过分类 ID 来定位并更新媒资分类的名称。
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
// 更新媒资分类
//
// Description:
//
// 创建媒资分类后，可调用本接口通过分类 ID 来定位并更新媒资分类的名称。
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
// # UpdateMedia
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
// # UpdateMedia
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
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
