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
// 批量获取一刻AI应用生成任务
//
// @param request - BatchGetYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetYikeAIAppJobResponse
func (client *Client) BatchGetYikeAIAppJobWithOptions(request *BatchGetYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *BatchGetYikeAIAppJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetYikeAIAppJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetYikeAIAppJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取一刻AI应用生成任务
//
// @param request - BatchGetYikeAIAppJobRequest
//
// @return BatchGetYikeAIAppJobResponse
func (client *Client) BatchGetYikeAIAppJob(request *BatchGetYikeAIAppJobRequest) (_result *BatchGetYikeAIAppJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetYikeAIAppJobResponse{}
	_body, _err := client.BatchGetYikeAIAppJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一刻媒资上传凭证
//
// @param request - CreateYikeAssetUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeAssetUploadResponse
func (client *Client) CreateYikeAssetUploadWithOptions(request *CreateYikeAssetUploadRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeAssetUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileExt) {
		query["FileExt"] = request.FileExt
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateYikeAssetUpload"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateYikeAssetUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻媒资上传凭证
//
// @param request - CreateYikeAssetUploadRequest
//
// @return CreateYikeAssetUploadResponse
func (client *Client) CreateYikeAssetUpload(request *CreateYikeAssetUploadRequest) (_result *CreateYikeAssetUploadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateYikeAssetUploadResponse{}
	_body, _err := client.CreateYikeAssetUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一刻AI应用任务
//
// @param request - GetYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAIAppJobResponse
func (client *Client) GetYikeAIAppJobWithOptions(request *GetYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAIAppJobResponse, _err error) {
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
		Action:      dara.String("GetYikeAIAppJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeAIAppJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻AI应用任务
//
// @param request - GetYikeAIAppJobRequest
//
// @return GetYikeAIAppJobResponse
func (client *Client) GetYikeAIAppJob(request *GetYikeAIAppJobRequest) (_result *GetYikeAIAppJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeAIAppJobResponse{}
	_body, _err := client.GetYikeAIAppJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一刻故事板任务
//
// @param request - GetYikeStoryboardJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeStoryboardJobResponse
func (client *Client) GetYikeStoryboardJobWithOptions(request *GetYikeStoryboardJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeStoryboardJobResponse, _err error) {
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
		Action:      dara.String("GetYikeStoryboardJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeStoryboardJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻故事板任务
//
// @param request - GetYikeStoryboardJobRequest
//
// @return GetYikeStoryboardJobResponse
func (client *Client) GetYikeStoryboardJob(request *GetYikeStoryboardJobRequest) (_result *GetYikeStoryboardJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeStoryboardJobResponse{}
	_body, _err := client.GetYikeStoryboardJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交一刻AI应用任务
//
// @param request - SubmitYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeAIAppJobResponse
func (client *Client) SubmitYikeAIAppJobWithOptions(request *SubmitYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeAIAppJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppParams) {
		body["AppParams"] = request.AppParams
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProductionId) {
		body["ProductionId"] = request.ProductionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikeAIAppJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikeAIAppJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交一刻AI应用任务
//
// @param request - SubmitYikeAIAppJobRequest
//
// @return SubmitYikeAIAppJobResponse
func (client *Client) SubmitYikeAIAppJob(request *SubmitYikeAIAppJobRequest) (_result *SubmitYikeAIAppJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitYikeAIAppJobResponse{}
	_body, _err := client.SubmitYikeAIAppJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交一刻故事板任务
//
// @param request - SubmitYikeStoryboardJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeStoryboardJobResponse
func (client *Client) SubmitYikeStoryboardJobWithOptions(request *SubmitYikeStoryboardJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeStoryboardJobResponse, _err error) {
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

	if !dara.IsNil(request.ExecMode) {
		query["ExecMode"] = request.ExecMode
	}

	if !dara.IsNil(request.ModelParams) {
		query["ModelParams"] = request.ModelParams
	}

	if !dara.IsNil(request.NarrationVoiceId) {
		query["NarrationVoiceId"] = request.NarrationVoiceId
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.ShotPromptMode) {
		query["ShotPromptMode"] = request.ShotPromptMode
	}

	if !dara.IsNil(request.SkipFailureShot) {
		query["SkipFailureShot"] = request.SkipFailureShot
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoModel) {
		query["VideoModel"] = request.VideoModel
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileURL) {
		body["FileURL"] = request.FileURL
	}

	if !dara.IsNil(request.ShotSplitMode) {
		body["ShotSplitMode"] = request.ShotSplitMode
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StyleId) {
		body["StyleId"] = request.StyleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikeStoryboardJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikeStoryboardJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交一刻故事板任务
//
// @param request - SubmitYikeStoryboardJobRequest
//
// @return SubmitYikeStoryboardJobResponse
func (client *Client) SubmitYikeStoryboardJob(request *SubmitYikeStoryboardJobRequest) (_result *SubmitYikeStoryboardJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitYikeStoryboardJobResponse{}
	_body, _err := client.SubmitYikeStoryboardJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
