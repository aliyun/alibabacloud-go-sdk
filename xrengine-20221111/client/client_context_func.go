// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 用户签署协议
//
// @param request - AcceptAgreementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptAgreementResponse
func (client *Client) AcceptAgreementWithContext(ctx context.Context, request *AcceptAgreementRequest, runtime *dara.RuntimeOptions) (_result *AcceptAgreementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptAgreement"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptAgreementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加白名单（云账号）
//
// @param request - AddWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWhitelistResponse
func (client *Client) AddWhitelistWithContext(ctx context.Context, request *AddWhitelistRequest, runtime *dara.RuntimeOptions) (_result *AddWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunUid) {
		body["AliyunUid"] = request.AliyunUid
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWhitelist"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请ai试衣服务
//
// @param request - ApplyForTryOnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyForTryOnResponse
func (client *Client) ApplyForTryOnWithContext(ctx context.Context, request *ApplyForTryOnRequest, runtime *dara.RuntimeOptions) (_result *ApplyForTryOnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyForTryOn"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyForTryOnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthUserResponse
func (client *Client) AuthUserWithContext(ctx context.Context, request *AuthUserRequest, runtime *dara.RuntimeOptions) (_result *AuthUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthUser"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - BatchCreateSvcMapBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateSvcMapBindResponse
func (client *Client) BatchCreateSvcMapBindWithContext(ctx context.Context, tmpReq *BatchCreateSvcMapBindRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateSvcMapBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateSvcMapBindShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MapIds) {
		request.MapIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MapIds, dara.String("MapIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapIdsShrink) {
		body["MapIds"] = request.MapIdsShrink
	}

	if !dara.IsNil(request.SvcId) {
		body["SvcId"] = request.SvcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateSvcMapBind"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateSvcMapBindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消关联
//
// @param tmpReq - BatchDeleteSvcMapBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteSvcMapBindResponse
func (client *Client) BatchDeleteSvcMapBindWithContext(ctx context.Context, tmpReq *BatchDeleteSvcMapBindRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteSvcMapBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteSvcMapBindShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteSvcMapBind"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteSvcMapBindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 构建项目
//
// @param request - BuildProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildProjectResponse
func (client *Client) BuildProjectWithContext(ctx context.Context, request *BuildProjectRequest, runtime *dara.RuntimeOptions) (_result *BuildProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VideoName) {
		query["VideoName"] = request.VideoName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建，同时创建空白的pai占位
//
// @param request - CreateLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLocationServiceResponse
func (client *Client) CreateLocationServiceWithContext(ctx context.Context, request *CreateLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateLocationServiceResponse, _err error) {
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

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Note) {
		body["Note"] = request.Note
	}

	if !dara.IsNil(request.Qps) {
		body["Qps"] = request.Qps
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLocationServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TryOnParams) {
		request.TryOnParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TryOnParams, dara.String("TryOnParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuild) {
		body["AutoBuild"] = request.AutoBuild
	}

	if !dara.IsNil(request.Dependencies) {
		body["Dependencies"] = request.Dependencies
	}

	if !dara.IsNil(request.ExtInfo) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.Gender) {
		body["Gender"] = request.Gender
	}

	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.Intro) {
		body["Intro"] = request.Intro
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapUuid) {
		body["MapUuid"] = request.MapUuid
	}

	if !dara.IsNil(request.Method) {
		body["Method"] = request.Method
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.TryOnParamsShrink) {
		body["TryOnParams"] = request.TryOnParamsShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Weight) {
		body["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建素材上传policy
//
// @param request - CreateSourcePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSourcePolicyResponse
func (client *Client) CreateSourcePolicyWithContext(ctx context.Context, request *CreateSourcePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateSourcePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSourcePolicy"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSourcePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目信息
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件
//
// @param request - DeleteSourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceFileResponse
func (client *Client) DeleteSourceFileWithContext(ctx context.Context, request *DeleteSourceFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteSourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSourceFile"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 1校验部署状态，一期不支持二次部署。相关关联记录里状态智能变更
//
// @param request - DeployLocationTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployLocationTreeResponse
func (client *Client) DeployLocationTreeWithContext(ctx context.Context, request *DeployLocationTreeRequest, runtime *dara.RuntimeOptions) (_result *DeployLocationTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceUpdate) {
		body["ForceUpdate"] = request.ForceUpdate
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.SvcId) {
		body["SvcId"] = request.SvcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployLocationTree"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployLocationTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查出绑定记录
//
// @param request - FindSvcMapBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindSvcMapBindResponse
func (client *Client) FindSvcMapBindWithContext(ctx context.Context, request *FindSvcMapBindRequest, runtime *dara.RuntimeOptions) (_result *FindSvcMapBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SvcId) {
		body["SvcId"] = request.SvcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindSvcMapBind"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindSvcMapBindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户申请状态
//
// @param request - GetApplyStatusForTryOnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplyStatusForTryOnResponse
func (client *Client) GetApplyStatusForTryOnWithContext(ctx context.Context, request *GetApplyStatusForTryOnRequest, runtime *dara.RuntimeOptions) (_result *GetApplyStatusForTryOnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplyStatusForTryOn"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplyStatusForTryOnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArEditCommonMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArEditCommonMaterialResponse
func (client *Client) GetArEditCommonMaterialWithContext(ctx context.Context, request *GetArEditCommonMaterialRequest, runtime *dara.RuntimeOptions) (_result *GetArEditCommonMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArEditCommonMaterial"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArEditCommonMaterialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArEditStsAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArEditStsAuthResponse
func (client *Client) GetArEditStsAuthWithContext(ctx context.Context, request *GetArEditStsAuthRequest, runtime *dara.RuntimeOptions) (_result *GetArEditStsAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArEditStsAuth"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArEditStsAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArEditUgcMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArEditUgcMaterialResponse
func (client *Client) GetArEditUgcMaterialWithContext(ctx context.Context, request *GetArEditUgcMaterialRequest, runtime *dara.RuntimeOptions) (_result *GetArEditUgcMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArEditUgcMaterial"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArEditUgcMaterialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目模型详情
//
// @param request - GetProjectDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectDatasetResponse
func (client *Client) GetProjectDatasetWithContext(ctx context.Context, request *GetProjectDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetProjectDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectDataset"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用同步算法
//
// @param tmpReq - InvokeSyncAlgorithmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeSyncAlgorithmResponse
func (client *Client) InvokeSyncAlgorithmWithContext(ctx context.Context, tmpReq *InvokeSyncAlgorithmRequest, runtime *dara.RuntimeOptions) (_result *InvokeSyncAlgorithmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvokeSyncAlgorithmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeSyncAlgorithm"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeSyncAlgorithmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAreaMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAreaMapResponse
func (client *Client) ListAreaMapWithContext(ctx context.Context, request *ListAreaMapRequest, runtime *dara.RuntimeOptions) (_result *ListAreaMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAreaMap"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAreaMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举服饰
//
// @param tmpReq - ListClothesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClothesResponse
func (client *Client) ListClothesWithContext(ctx context.Context, tmpReq *ListClothesRequest, runtime *dara.RuntimeOptions) (_result *ListClothesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListClothesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("Categories"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		query["Categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.ClothSize) {
		query["ClothSize"] = request.ClothSize
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClothes"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClothesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLocationPaiImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLocationPaiImageResponse
func (client *Client) ListLocationPaiImageWithContext(ctx context.Context, request *ListLocationPaiImageRequest, runtime *dara.RuntimeOptions) (_result *ListLocationPaiImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLocationPaiImage"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLocationPaiImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLocationServiceResponse
func (client *Client) ListLocationServiceWithContext(ctx context.Context, request *ListLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *ListLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLocationServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出当前用户项目列表
//
// @param request - ListProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithContext(ctx context.Context, request *ListProjectRequest, runtime *dara.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizUsage) {
		body["BizUsage"] = request.BizUsage
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.ExcludedBizUsage) {
		body["ExcludedBizUsage"] = request.ExcludedBizUsage
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WithSource) {
		body["WithSource"] = request.WithSource
	}

	if !dara.IsNil(request.WithUser) {
		body["WithUser"] = request.WithUser
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据项目Id查出依赖当前项目的项目
//
// @param request - ListProjectsByDependencyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsByDependencyIdResponse
func (client *Client) ListProjectsByDependencyIdWithContext(ctx context.Context, request *ListProjectsByDependencyIdRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsByDependencyIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DependencyProjectId) {
		body["DependencyProjectId"] = request.DependencyProjectId
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapUuid) {
		body["MapUuid"] = request.MapUuid
	}

	if !dara.IsNil(request.WithDataset) {
		body["WithDataset"] = request.WithDataset
	}

	if !dara.IsNil(request.WithSource) {
		body["WithSource"] = request.WithSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectsByDependencyId"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsByDependencyIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举项目中上传的文件列表
//
// @param request - ListSourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSourceFileResponse
func (client *Client) ListSourceFileWithContext(ctx context.Context, request *ListSourceFileRequest, runtime *dara.RuntimeOptions) (_result *ListSourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSourceFile"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSourceFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - LoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginResponse
func (client *Client) LoginWithContext(ctx context.Context, request *LoginRequest, runtime *dara.RuntimeOptions) (_result *LoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EmpId) {
		body["EmpId"] = request.EmpId
	}

	if !dara.IsNil(request.EmpName) {
		body["EmpName"] = request.EmpName
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Login"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 无线端APP登陆
//
// @param request - LoginAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginAppResponse
func (client *Client) LoginAppWithContext(ctx context.Context, request *LoginAppRequest, runtime *dara.RuntimeOptions) (_result *LoginAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EmpId) {
		body["EmpId"] = request.EmpId
	}

	if !dara.IsNil(request.EmpName) {
		body["EmpName"] = request.EmpName
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoginApp"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoginAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PublishArEditProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishArEditProjectResponse
func (client *Client) PublishArEditProjectWithContext(ctx context.Context, request *PublishArEditProjectRequest, runtime *dara.RuntimeOptions) (_result *PublishArEditProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishArEditProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishArEditProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目发布
//
// @param request - PublishProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishProjectResponse
func (client *Client) PublishProjectWithContext(ctx context.Context, request *PublishProjectRequest, runtime *dara.RuntimeOptions) (_result *PublishProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAreaMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAreaMapResponse
func (client *Client) QueryAreaMapWithContext(ctx context.Context, request *QueryAreaMapRequest, runtime *dara.RuntimeOptions) (_result *QueryAreaMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAreaMap"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAreaMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找项目构建失败时的信息
//
// @param request - QueryBuildBreakpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBuildBreakpointResponse
func (client *Client) QueryBuildBreakpointWithContext(ctx context.Context, request *QueryBuildBreakpointRequest, runtime *dara.RuntimeOptions) (_result *QueryBuildBreakpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBuildBreakpoint"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBuildBreakpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLocationServiceResponse
func (client *Client) QueryLocationServiceWithContext(ctx context.Context, request *QueryLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *QueryLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLocationServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目构建详情
//
// @param request - QueryProjectBuildDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectBuildDetailResponse
func (client *Client) QueryProjectBuildDetailWithContext(ctx context.Context, request *QueryProjectBuildDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryProjectBuildDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProjectBuildDetail"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProjectBuildDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目详情
//
// @param request - QueryProjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectDetailResponse
func (client *Client) QueryProjectDetailWithContext(ctx context.Context, request *QueryProjectDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryProjectDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProjectDetail"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProjectDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 该注册接口只用于oneconsole页面的注册
//
// @param request - RegisterUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterUserResponse
func (client *Client) RegisterUserWithContext(ctx context.Context, request *RegisterUserRequest, runtime *dara.RuntimeOptions) (_result *RegisterUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterUser"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResumeLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeLocationServiceResponse
func (client *Client) ResumeLocationServiceWithContext(ctx context.Context, request *ResumeLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *ResumeLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeLocationServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveArEditProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveArEditProjectResponse
func (client *Client) SaveArEditProjectWithContext(ctx context.Context, request *SaveArEditProjectRequest, runtime *dara.RuntimeOptions) (_result *SaveArEditProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveArEditProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveArEditProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存素材
//
// @param request - SaveSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSourceResponse
func (client *Client) SaveSourceWithContext(ctx context.Context, request *SaveSourceRequest, runtime *dara.RuntimeOptions) (_result *SaveSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeStatus) {
		query["ChangeStatus"] = request.ChangeStatus
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.NeedCheck) {
		query["NeedCheck"] = request.NeedCheck
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSource"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuspendLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendLocationServiceResponse
func (client *Client) SuspendLocationServiceWithContext(ctx context.Context, request *SuspendLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *SuspendLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendLocationServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目(发布状态改回隐私状态)
//
// @param request - UnPublishProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnPublishProjectResponse
func (client *Client) UnPublishProjectWithContext(ctx context.Context, request *UnPublishProjectRequest, runtime *dara.RuntimeOptions) (_result *UnPublishProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnPublishProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnPublishProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 如果不包含treeConfig则只是简单更新
//
// @param request - UpdateLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocationServiceResponse
func (client *Client) UpdateLocationServiceWithContext(ctx context.Context, request *UpdateLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Note) {
		body["Note"] = request.Note
	}

	if !dara.IsNil(request.Qps) {
		body["Qps"] = request.Qps
	}

	if !dara.IsNil(request.SvcName) {
		body["SvcName"] = request.SvcName
	}

	if !dara.IsNil(request.SvcState) {
		body["SvcState"] = request.SvcState
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocationServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂存接口比较复杂，单独拎出来
//
// @param request - UpdateLocationTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocationTreeResponse
func (client *Client) UpdateLocationTreeWithContext(ctx context.Context, request *UpdateLocationTreeRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocationTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.TreeConfig) {
		body["TreeConfig"] = request.TreeConfig
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocationTree"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocationTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新项目信息
//
// @param tmpReq - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, tmpReq *UpdateProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuild) {
		query["AutoBuild"] = request.AutoBuild
	}

	if !dara.IsNil(request.ExtShrink) {
		query["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Intro) {
		query["Intro"] = request.Intro
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
