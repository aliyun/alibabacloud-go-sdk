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
// 增加用户积分
//
// @param request - AddYikeUserCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddYikeUserCreditResponse
func (client *Client) AddYikeUserCreditWithOptions(request *AddYikeUserCreditRequest, runtime *dara.RuntimeOptions) (_result *AddYikeUserCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Credit) {
		query["Credit"] = request.Credit
	}

	if !dara.IsNil(request.YikeUserId) {
		query["YikeUserId"] = request.YikeUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddYikeUserCredit"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddYikeUserCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加用户积分
//
// @param request - AddYikeUserCreditRequest
//
// @return AddYikeUserCreditResponse
func (client *Client) AddYikeUserCredit(request *AddYikeUserCreditRequest) (_result *AddYikeUserCreditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddYikeUserCreditResponse{}
	_body, _err := client.AddYikeUserCreditWithOptions(request, runtime)
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
// 批量获取媒资信息
//
// @param request - BatchGetYikeAssetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetYikeAssetMediaInfosResponse
func (client *Client) BatchGetYikeAssetMediaInfosWithOptions(request *BatchGetYikeAssetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *BatchGetYikeAssetMediaInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetYikeAssetMediaInfos"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetYikeAssetMediaInfosResponse{}
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
// @param request - BatchGetYikeAssetMediaInfosRequest
//
// @return BatchGetYikeAssetMediaInfosResponse
func (client *Client) BatchGetYikeAssetMediaInfos(request *BatchGetYikeAssetMediaInfosRequest) (_result *BatchGetYikeAssetMediaInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetYikeAssetMediaInfosResponse{}
	_body, _err := client.BatchGetYikeAssetMediaInfosWithOptions(request, runtime)
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
// 创建一刻项目
//
// @param request - CreateYikeProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeProductionResponse
func (client *Client) CreateYikeProductionWithOptions(request *CreateYikeProductionRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeProductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateYikeProduction"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateYikeProductionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一刻项目
//
// @param request - CreateYikeProductionRequest
//
// @return CreateYikeProductionResponse
func (client *Client) CreateYikeProduction(request *CreateYikeProductionRequest) (_result *CreateYikeProductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateYikeProductionResponse{}
	_body, _err := client.CreateYikeProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一刻子用户
//
// @param request - CreateYikeUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeUserResponse
func (client *Client) CreateYikeUserWithOptions(request *CreateYikeUserRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nickname) {
		query["Nickname"] = request.Nickname
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProductionIds) {
		query["ProductionIds"] = request.ProductionIds
	}

	if !dara.IsNil(request.UserNamePrefix) {
		query["UserNamePrefix"] = request.UserNamePrefix
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateYikeUser"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateYikeUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一刻子用户
//
// @param request - CreateYikeUserRequest
//
// @return CreateYikeUserResponse
func (client *Client) CreateYikeUser(request *CreateYikeUserRequest) (_result *CreateYikeUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateYikeUserResponse{}
	_body, _err := client.CreateYikeUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作室
//
// @param request - CreateYikeWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeWorkspaceResponse
func (client *Client) CreateYikeWorkspaceWithOptions(request *CreateYikeWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserCountLimit) {
		query["UserCountLimit"] = request.UserCountLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateYikeWorkspace"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateYikeWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作室
//
// @param request - CreateYikeWorkspaceRequest
//
// @return CreateYikeWorkspaceResponse
func (client *Client) CreateYikeWorkspace(request *CreateYikeWorkspaceRequest) (_result *CreateYikeWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateYikeWorkspaceResponse{}
	_body, _err := client.CreateYikeWorkspaceWithOptions(request, runtime)
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
// @param request - DeleteYikeAssetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteYikeAssetMediaInfosResponse
func (client *Client) DeleteYikeAssetMediaInfosWithOptions(request *DeleteYikeAssetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *DeleteYikeAssetMediaInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicDelete) {
		query["LogicDelete"] = request.LogicDelete
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteYikeAssetMediaInfos"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteYikeAssetMediaInfosResponse{}
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
// @param request - DeleteYikeAssetMediaInfosRequest
//
// @return DeleteYikeAssetMediaInfosResponse
func (client *Client) DeleteYikeAssetMediaInfos(request *DeleteYikeAssetMediaInfosRequest) (_result *DeleteYikeAssetMediaInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteYikeAssetMediaInfosResponse{}
	_body, _err := client.DeleteYikeAssetMediaInfosWithOptions(request, runtime)
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
// 获取一刻媒资内容信息
//
// @param request - GetYikeAssetMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAssetMediaInfoResponse
func (client *Client) GetYikeAssetMediaInfoWithOptions(request *GetYikeAssetMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAssetMediaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeAssetMediaInfo"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeAssetMediaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻媒资内容信息
//
// @param request - GetYikeAssetMediaInfoRequest
//
// @return GetYikeAssetMediaInfoResponse
func (client *Client) GetYikeAssetMediaInfo(request *GetYikeAssetMediaInfoRequest) (_result *GetYikeAssetMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeAssetMediaInfoResponse{}
	_body, _err := client.GetYikeAssetMediaInfoWithOptions(request, runtime)
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
// 获取一刻子用户信息
//
// @param request - GetYikeUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeUserResponse
func (client *Client) GetYikeUserWithOptions(request *GetYikeUserRequest, runtime *dara.RuntimeOptions) (_result *GetYikeUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeUser"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻子用户信息
//
// @param request - GetYikeUserRequest
//
// @return GetYikeUserResponse
func (client *Client) GetYikeUser(request *GetYikeUserRequest) (_result *GetYikeUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeUserResponse{}
	_body, _err := client.GetYikeUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一刻用户积分
//
// @param request - GetYikeUserCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeUserCreditResponse
func (client *Client) GetYikeUserCreditWithOptions(request *GetYikeUserCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeUserCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.YikeUserId) {
		query["YikeUserId"] = request.YikeUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeUserCredit"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeUserCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一刻用户积分
//
// @param request - GetYikeUserCreditRequest
//
// @return GetYikeUserCreditResponse
func (client *Client) GetYikeUserCredit(request *GetYikeUserCreditRequest) (_result *GetYikeUserCreditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetYikeUserCreditResponse{}
	_body, _err := client.GetYikeUserCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一刻文件夹列表
//
// @param request - ListYikeAssetFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListYikeAssetFoldersResponse
func (client *Client) ListYikeAssetFoldersWithOptions(request *ListYikeAssetFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListYikeAssetFoldersResponse, _err error) {
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

	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListYikeAssetFolders"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListYikeAssetFoldersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻文件夹列表
//
// @param request - ListYikeAssetFoldersRequest
//
// @return ListYikeAssetFoldersResponse
func (client *Client) ListYikeAssetFolders(request *ListYikeAssetFoldersRequest) (_result *ListYikeAssetFoldersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListYikeAssetFoldersResponse{}
	_body, _err := client.ListYikeAssetFoldersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一刻项目列表
//
// @param request - ListYikeProductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListYikeProductionsResponse
func (client *Client) ListYikeProductionsWithOptions(request *ListYikeProductionsRequest, runtime *dara.RuntimeOptions) (_result *ListYikeProductionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListYikeProductions"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListYikeProductionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一刻项目列表
//
// @param request - ListYikeProductionsRequest
//
// @return ListYikeProductionsResponse
func (client *Client) ListYikeProductions(request *ListYikeProductionsRequest) (_result *ListYikeProductionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListYikeProductionsResponse{}
	_body, _err := client.ListYikeProductionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查应用参数是否合法
//
// @param request - PrecheckYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckYikeAIAppJobResponse
func (client *Client) PrecheckYikeAIAppJobWithOptions(request *PrecheckYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *PrecheckYikeAIAppJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppParams) {
		query["AppParams"] = request.AppParams
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrecheckYikeAIAppJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrecheckYikeAIAppJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查应用参数是否合法
//
// @param request - PrecheckYikeAIAppJobRequest
//
// @return PrecheckYikeAIAppJobResponse
func (client *Client) PrecheckYikeAIAppJob(request *PrecheckYikeAIAppJobRequest) (_result *PrecheckYikeAIAppJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PrecheckYikeAIAppJobResponse{}
	_body, _err := client.PrecheckYikeAIAppJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册一刻媒资
//
// @param request - RegisterYikeAssetMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterYikeAssetMediaInfoResponse
func (client *Client) RegisterYikeAssetMediaInfoWithOptions(request *RegisterYikeAssetMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *RegisterYikeAssetMediaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterYikeAssetMediaInfo"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterYikeAssetMediaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册一刻媒资
//
// @param request - RegisterYikeAssetMediaInfoRequest
//
// @return RegisterYikeAssetMediaInfoResponse
func (client *Client) RegisterYikeAssetMediaInfo(request *RegisterYikeAssetMediaInfoRequest) (_result *RegisterYikeAssetMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterYikeAssetMediaInfoResponse{}
	_body, _err := client.RegisterYikeAssetMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 故事板任务恢复继续执行任务
//
// @param request - ResumeYikeStoryboardJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeYikeStoryboardJobResponse
func (client *Client) ResumeYikeStoryboardJobWithOptions(request *ResumeYikeStoryboardJobRequest, runtime *dara.RuntimeOptions) (_result *ResumeYikeStoryboardJobResponse, _err error) {
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
		Action:      dara.String("ResumeYikeStoryboardJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeYikeStoryboardJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 故事板任务恢复继续执行任务
//
// @param request - ResumeYikeStoryboardJobRequest
//
// @return ResumeYikeStoryboardJobResponse
func (client *Client) ResumeYikeStoryboardJob(request *ResumeYikeStoryboardJobRequest) (_result *ResumeYikeStoryboardJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeYikeStoryboardJobResponse{}
	_body, _err := client.ResumeYikeStoryboardJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置一刻事件回调
//
// @param request - SetYikeCallbackConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetYikeCallbackConfigResponse
func (client *Client) SetYikeCallbackConfigWithOptions(request *SetYikeCallbackConfigRequest, runtime *dara.RuntimeOptions) (_result *SetYikeCallbackConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackConfig) {
		query["CallbackConfig"] = request.CallbackConfig
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetYikeCallbackConfig"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetYikeCallbackConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置一刻事件回调
//
// @param request - SetYikeCallbackConfigRequest
//
// @return SetYikeCallbackConfigResponse
func (client *Client) SetYikeCallbackConfig(request *SetYikeCallbackConfigRequest) (_result *SetYikeCallbackConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetYikeCallbackConfigResponse{}
	_body, _err := client.SetYikeCallbackConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置用户角色
//
// @param request - SetYikeUserRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetYikeUserRoleResponse
func (client *Client) SetYikeUserRoleWithOptions(request *SetYikeUserRoleRequest, runtime *dara.RuntimeOptions) (_result *SetYikeUserRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.YikeUserId) {
		query["YikeUserId"] = request.YikeUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetYikeUserRole"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetYikeUserRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置用户角色
//
// @param request - SetYikeUserRoleRequest
//
// @return SetYikeUserRoleResponse
func (client *Client) SetYikeUserRole(request *SetYikeUserRoleRequest) (_result *SetYikeUserRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetYikeUserRoleResponse{}
	_body, _err := client.SetYikeUserRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 扣减用户积分
//
// @param request - SubYikeUserCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubYikeUserCreditResponse
func (client *Client) SubYikeUserCreditWithOptions(request *SubYikeUserCreditRequest, runtime *dara.RuntimeOptions) (_result *SubYikeUserCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Credit) {
		query["Credit"] = request.Credit
	}

	if !dara.IsNil(request.YikeUserId) {
		query["YikeUserId"] = request.YikeUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubYikeUserCredit"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubYikeUserCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 扣减用户积分
//
// @param request - SubYikeUserCreditRequest
//
// @return SubYikeUserCreditResponse
func (client *Client) SubYikeUserCredit(request *SubYikeUserCreditRequest) (_result *SubYikeUserCreditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubYikeUserCreditResponse{}
	_body, _err := client.SubYikeUserCreditWithOptions(request, runtime)
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
