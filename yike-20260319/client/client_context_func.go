// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds members to a Yike project.
//
// @param request - AddYikeProductionMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddYikeProductionMembersResponse
func (client *Client) AddYikeProductionMembersWithContext(ctx context.Context, request *AddYikeProductionMembersRequest, runtime *dara.RuntimeOptions) (_result *AddYikeProductionMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	if !dara.IsNil(request.YikeUserIds) {
		query["YikeUserIds"] = request.YikeUserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddYikeProductionMembers"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddYikeProductionMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Increases user credits.
//
// @param request - AddYikeUserCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddYikeUserCreditResponse
func (client *Client) AddYikeUserCreditWithContext(ctx context.Context, request *AddYikeUserCreditRequest, runtime *dara.RuntimeOptions) (_result *AddYikeUserCreditResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves multiple Yike AI application generation tasks in a batch.
//
// @param request - BatchGetYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetYikeAIAppJobResponse
func (client *Client) BatchGetYikeAIAppJobWithContext(ctx context.Context, request *BatchGetYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *BatchGetYikeAIAppJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves media asset information in batches.
//
// @param request - BatchGetYikeAssetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetYikeAssetMediaInfosResponse
func (client *Client) BatchGetYikeAssetMediaInfosWithContext(ctx context.Context, request *BatchGetYikeAssetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *BatchGetYikeAssetMediaInfosResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the upload credential for a media asset.
//
// @param request - CreateYikeAssetUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeAssetUploadResponse
func (client *Client) CreateYikeAssetUploadWithContext(ctx context.Context, request *CreateYikeAssetUploadRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeAssetUploadResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一刻云剪辑工程
//
// @param request - CreateYikeEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeEditingProjectResponse
func (client *Client) CreateYikeEditingProjectWithContext(ctx context.Context, request *CreateYikeEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeEditingProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.MaterialMaps) {
		query["MaterialMaps"] = request.MaterialMaps
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Timeline) {
		body["Timeline"] = request.Timeline
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateYikeEditingProject"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateYikeEditingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// @param request - CreateYikeProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeProductionResponse
func (client *Client) CreateYikeProductionWithContext(ctx context.Context, request *CreateYikeProductionRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeProductionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sub-user.
//
// @param request - CreateYikeUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeUserResponse
func (client *Client) CreateYikeUserWithContext(ctx context.Context, request *CreateYikeUserRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateYikeWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYikeWorkspaceResponse
func (client *Client) CreateYikeWorkspaceWithContext(ctx context.Context, request *CreateYikeWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateYikeWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - DeleteYikeAssetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteYikeAssetMediaInfosResponse
func (client *Client) DeleteYikeAssetMediaInfosWithContext(ctx context.Context, request *DeleteYikeAssetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *DeleteYikeAssetMediaInfosResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## Request description
//
// This API is used to generate a video narrated by a virtual human based on the provided text content and other parameters such as digital human information and common scenario type. You must specify key configuration items including the text type (original script or narration script), output dimensions, and resolution. You can also choose whether to add subtitles or specify the output language. In addition, you can pass custom parameters through the `UserData` field, which are returned as-is in the callback.
//
// @param request - GetImageGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageGenerationJobResponse
func (client *Client) GetImageGenerationJobWithContext(ctx context.Context, request *GetImageGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *GetImageGenerationJobResponse, _err error) {
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
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoGenerationJobResponse
func (client *Client) GetVideoGenerationJobWithContext(ctx context.Context, request *GetVideoGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *GetVideoGenerationJobResponse, _err error) {
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
		Action:      dara.String("GetVideoGenerationJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an AI application task from China International Communications Group (YIKE) AI.
//
// @param request - GetYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAIAppJobResponse
func (client *Client) GetYikeAIAppJobWithContext(ctx context.Context, request *GetYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAIAppJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAccountCreditResponse
func (client *Client) GetYikeAccountCreditWithContext(ctx context.Context, request *GetYikeAccountCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAccountCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeAccountCredit"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeAccountCreditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an intelligent agent task.
//
// @param request - GetYikeAgentJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAgentJobResponse
func (client *Client) GetYikeAgentJobWithContext(ctx context.Context, request *GetYikeAgentJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAgentJobResponse, _err error) {
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
		Action:      dara.String("GetYikeAgentJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeAgentJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the content information of a media asset.
//
// @param request - GetYikeAssetMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAssetMediaInfoResponse
func (client *Client) GetYikeAssetMediaInfoWithContext(ctx context.Context, request *GetYikeAssetMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAssetMediaInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information and results of an editing project export task.
//
// @param request - GetYikeProjectExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeProjectExportJobResponse
func (client *Client) GetYikeProjectExportJobWithContext(ctx context.Context, request *GetYikeProjectExportJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeProjectExportJobResponse, _err error) {
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
		Action:      dara.String("GetYikeProjectExportJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeProjectExportJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a Yike prompt enhancement and audio repair video generation task.
//
// @param request - GetYikePromptExpansionVoiceFixJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikePromptExpansionVoiceFixJobResponse
func (client *Client) GetYikePromptExpansionVoiceFixJobWithContext(ctx context.Context, request *GetYikePromptExpansionVoiceFixJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikePromptExpansionVoiceFixJobResponse, _err error) {
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
		Action:      dara.String("GetYikePromptExpansionVoiceFixJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikePromptExpansionVoiceFixJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a storyboard task.
//
// @param request - GetYikeStoryboardJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeStoryboardJobResponse
func (client *Client) GetYikeStoryboardJobWithContext(ctx context.Context, request *GetYikeStoryboardJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeStoryboardJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a Yike RAM user.
//
// @param request - GetYikeUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeUserResponse
func (client *Client) GetYikeUserWithContext(ctx context.Context, request *GetYikeUserRequest, runtime *dara.RuntimeOptions) (_result *GetYikeUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the credits of a Yike user.
//
// @param request - GetYikeUserCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeUserCreditResponse
func (client *Client) GetYikeUserCreditWithContext(ctx context.Context, request *GetYikeUserCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeUserCreditResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an intelligent video generation task for narration-only videos (without digital humans).
//
// @param request - GetYikeVoiceNarratorJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeVoiceNarratorJobResponse
func (client *Client) GetYikeVoiceNarratorJobWithContext(ctx context.Context, request *GetYikeVoiceNarratorJobRequest, runtime *dara.RuntimeOptions) (_result *GetYikeVoiceNarratorJobResponse, _err error) {
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
		Action:      dara.String("GetYikeVoiceNarratorJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeVoiceNarratorJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of folders.
//
// @param request - ListYikeAssetFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListYikeAssetFoldersResponse
func (client *Client) ListYikeAssetFoldersWithContext(ctx context.Context, request *ListYikeAssetFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListYikeAssetFoldersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of Yike projects.
//
// @param request - ListYikeProductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListYikeProductionsResponse
func (client *Client) ListYikeProductionsWithContext(ctx context.Context, request *ListYikeProductionsRequest, runtime *dara.RuntimeOptions) (_result *ListYikeProductionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of Yike workspaces.
//
// @param request - ListYikeWorkspacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListYikeWorkspacesResponse
func (client *Client) ListYikeWorkspacesWithContext(ctx context.Context, request *ListYikeWorkspacesRequest, runtime *dara.RuntimeOptions) (_result *ListYikeWorkspacesResponse, _err error) {
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
		Action:      dara.String("ListYikeWorkspaces"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListYikeWorkspacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the application parameters are valid.
//
// @param request - PrecheckYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckYikeAIAppJobResponse
func (client *Client) PrecheckYikeAIAppJobWithContext(ctx context.Context, request *PrecheckYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *PrecheckYikeAIAppJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a media asset.
//
// @param request - RegisterYikeAssetMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterYikeAssetMediaInfoResponse
func (client *Client) RegisterYikeAssetMediaInfoWithContext(ctx context.Context, request *RegisterYikeAssetMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *RegisterYikeAssetMediaInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes the execution of a storyboard task.
//
// @param request - ResumeYikeStoryboardJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeYikeStoryboardJobResponse
func (client *Client) ResumeYikeStoryboardJobWithContext(ctx context.Context, request *ResumeYikeStoryboardJobRequest, runtime *dara.RuntimeOptions) (_result *ResumeYikeStoryboardJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures event callbacks for the business system.
//
// @param request - SetYikeCallbackConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetYikeCallbackConfigResponse
func (client *Client) SetYikeCallbackConfigWithContext(ctx context.Context, request *SetYikeCallbackConfigRequest, runtime *dara.RuntimeOptions) (_result *SetYikeCallbackConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the role of a user.
//
// @param request - SetYikeUserRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetYikeUserRoleResponse
func (client *Client) SetYikeUserRoleWithContext(ctx context.Context, request *SetYikeUserRoleRequest, runtime *dara.RuntimeOptions) (_result *SetYikeUserRoleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deducts user credits.
//
// @param request - SubYikeUserCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubYikeUserCreditResponse
func (client *Client) SubYikeUserCreditWithContext(ctx context.Context, request *SubYikeUserCreditRequest, runtime *dara.RuntimeOptions) (_result *SubYikeUserCreditResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## Request description
//
// This API is used to generate a video narrated by a virtual human based on the provided text content and other parameters (such as digital human information and application scenario type). You must specify key configuration items such as the text type (original script or narration script), output dimensions, and resolution. You can also choose whether to add subtitles or specify the output language. In addition, you can pass custom parameters through the `UserData` field, which are returned as-is in the callback.
//
// @param request - SubmitImageGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageGenerationJobResponse
func (client *Client) SubmitImageGenerationJobWithContext(ctx context.Context, request *SubmitImageGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitImageGenerationJobResponse, _err error) {
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
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## Request description
//
// This API generates a video featuring a virtual human speaking based on the provided text content and other parameters (such as digital human information and common scenarios type). You must specify the text type (original script or spoken script), output dimensions, resolution, and other key configuration items. You can also choose whether to add subtitles or specify the output language. Additionally, you can pass custom parameters through the `UserData` field, which are returned as-is in the callback.
//
// @param request - SubmitVideoGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoGenerationJobResponse
func (client *Client) SubmitVideoGenerationJobWithContext(ctx context.Context, request *SubmitVideoGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoGenerationJobResponse, _err error) {
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
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an AI application task to Yike AI.
//
// @param request - SubmitYikeAIAppJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeAIAppJobResponse
func (client *Client) SubmitYikeAIAppJobWithContext(ctx context.Context, request *SubmitYikeAIAppJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeAIAppJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an intelligent video production task for a digital human oral broadcasting scenario. This operation is applicable to video scenarios such as influencer product promotion and knowledge explanation.
//
// Description:
//
// ## Description
//
// This operation generates a video in which a virtual human delivers an oral broadcast based on the provided text content and other parameters such as digital human information and application scenario type. You must specify key configuration items including the text type (raw script or oral broadcast script), video dimensions, and resolution. You can also choose whether to add subtitles or specify the output language. In addition, you can use the `UserData` field to pass custom parameters, which are returned as-is in the callback.
//
// Common scenarios: influencer product promotion, knowledge explanation, and similar video scenarios.
//
// @param request - SubmitYikeAvatarNarratorJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeAvatarNarratorJobResponse
func (client *Client) SubmitYikeAvatarNarratorJobWithContext(ctx context.Context, request *SubmitYikeAvatarNarratorJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeAvatarNarratorJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobParams) {
		body["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikeAvatarNarratorJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikeAvatarNarratorJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an online editing project export task that supports exporting pure audio and SRT subtitles.
//
// @param request - SubmitYikeProjectExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeProjectExportJobResponse
func (client *Client) SubmitYikeProjectExportJobWithContext(ctx context.Context, request *SubmitYikeProjectExportJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeProjectExportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikeProjectExportJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikeProjectExportJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video generation task with prompt enhancement and audio repair.
//
// @param request - SubmitYikePromptExpansionVoiceFixJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikePromptExpansionVoiceFixJobResponse
func (client *Client) SubmitYikePromptExpansionVoiceFixJobWithContext(ctx context.Context, request *SubmitYikePromptExpansionVoiceFixJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikePromptExpansionVoiceFixJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobParams) {
		body["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikePromptExpansionVoiceFixJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikePromptExpansionVoiceFixJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a storyboard generation task.
//
// Description:
//
// Ensure that your credits remain above 5,000 when calling this operation. Insufficient credits may cause the task to be interrupted.
//
// @param request - SubmitYikeStoryboardJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeStoryboardJobResponse
func (client *Client) SubmitYikeStoryboardJobWithContext(ctx context.Context, request *SubmitYikeStoryboardJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeStoryboardJobResponse, _err error) {
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

	if !dara.IsNil(request.KeepOriginDialogue) {
		query["KeepOriginDialogue"] = request.KeepOriginDialogue
	}

	if !dara.IsNil(request.ModelParams) {
		query["ModelParams"] = request.ModelParams
	}

	if !dara.IsNil(request.NarrationVoiceId) {
		query["NarrationVoiceId"] = request.NarrationVoiceId
	}

	if !dara.IsNil(request.NeedCaption) {
		query["NeedCaption"] = request.NeedCaption
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.ShotPromptLang) {
		query["ShotPromptLang"] = request.ShotPromptLang
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a video replication task that supports same-category content rewriting scenarios by preserving the original video framework and generating deduplicated variants through replacing partial elements (person/voice/image/text).
//
// @param request - SubmitYikeVideoCloneJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeVideoCloneJobResponse
func (client *Client) SubmitYikeVideoCloneJobWithContext(ctx context.Context, request *SubmitYikeVideoCloneJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeVideoCloneJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobParams) {
		body["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikeVideoCloneJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikeVideoCloneJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an intelligent video generation task for a voice-over-only scenario (without a digital human). This task is applicable to video scenarios such as product showcasing and news broadcasting.
//
// @param request - SubmitYikeVoiceNarratorJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitYikeVoiceNarratorJobResponse
func (client *Client) SubmitYikeVoiceNarratorJobWithContext(ctx context.Context, request *SubmitYikeVoiceNarratorJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitYikeVoiceNarratorJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobParams) {
		body["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitYikeVoiceNarratorJob"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitYikeVoiceNarratorJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update a Yike project
//
// @param request - UpdateYikeProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateYikeProductionResponse
func (client *Client) UpdateYikeProductionWithContext(ctx context.Context, request *UpdateYikeProductionRequest, runtime *dara.RuntimeOptions) (_result *UpdateYikeProductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateYikeProduction"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateYikeProductionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a China Short Video (Yike) project member.
//
// @param request - UpdateYikeProductionMemberAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateYikeProductionMemberAuthResponse
func (client *Client) UpdateYikeProductionMemberAuthWithContext(ctx context.Context, request *UpdateYikeProductionMemberAuthRequest, runtime *dara.RuntimeOptions) (_result *UpdateYikeProductionMemberAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["Auth"] = request.Auth
	}

	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	if !dara.IsNil(request.YikeUserId) {
		query["YikeUserId"] = request.YikeUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateYikeProductionMemberAuth"),
		Version:     dara.String("2026-03-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateYikeProductionMemberAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
