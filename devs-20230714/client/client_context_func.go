// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 激活身份绑定,完成OAuth授权
//
// @param request - ActivateConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateConnectionResponse
func (client *Client) ActivateConnectionWithContext(ctx context.Context, name *string, request *ActivateConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		body["account"] = request.Account
	}

	if !dara.IsNil(request.Credential) {
		body["credential"] = request.Credential
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateConnection"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/connections/" + dara.PercentEncode(dara.StringValue(name)) + "/activate"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPipelineResponse
func (client *Client) CancelPipelineWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelPipelineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPipeline"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/pipelines/" + dara.PercentEncode(dara.StringValue(name)) + "/cancel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks/" + dara.PercentEncode(dara.StringValue(name)) + "/cancel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建交付物存储
//
// @param request - CreateArtifactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactResponse
func (client *Client) CreateArtifactWithContext(ctx context.Context, request *CreateArtifactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateArtifactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArtifact"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/artifacts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建环境
//
// @param request - CreateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithContext(ctx context.Context, projectName *string, request *CreateEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvironment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线
//
// @param request - CreatePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineResponse
func (client *Client) CreatePipelineWithContext(ctx context.Context, request *CreatePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipeline"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/pipelines"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineResponse{}
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
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 创建任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工具集
//
// @param request - CreateToolsetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateToolsetResponse
func (client *Client) CreateToolsetWithContext(ctx context.Context, request *CreateToolsetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateToolsetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateToolset"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/toolsets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateToolsetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除交付物
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactResponse
func (client *Client) DeleteArtifactWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteArtifactResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteArtifact"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/artifacts/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除身份绑定
//
// @param request - DeleteConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnectionWithContext(ctx context.Context, name *string, request *DeleteConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnection"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/connections/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除环境
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithContext(ctx context.Context, projectName *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvironment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param request - DeleteProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, name *string, request *DeleteProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
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
// 删除工具集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteToolsetResponse
func (client *Client) DeleteToolsetWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteToolsetResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteToolset"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/toolsets/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteToolsetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手动触发环境部署
//
// @param request - DeployEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployEnvironmentResponse
func (client *Client) DeployEnvironmentWithContext(ctx context.Context, projectName *string, name *string, request *DeployEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployEnvironment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments/" + dara.PercentEncode(dara.StringValue(name)) + "/deploy"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取交付物的zip包临时下载地址url
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchArtifactDownloadUrlResponse
func (client *Client) FetchArtifactDownloadUrlWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FetchArtifactDownloadUrlResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchArtifactDownloadUrl"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/artifacts/" + dara.PercentEncode(dara.StringValue(name)) + "/fetchCode"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchArtifactDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取交付物临时上传的bucket、object和临时sts
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchArtifactTempBucketTokenResponse
func (client *Client) FetchArtifactTempBucketTokenWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FetchArtifactTempBucketTokenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchArtifactTempBucketToken"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/artifacts/action/fetchTempBucketToken"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchArtifactTempBucketTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份绑定中的凭证信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchConnectionCredentialResponse
func (client *Client) FetchConnectionCredentialWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FetchConnectionCredentialResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchConnectionCredential"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/connections/" + dara.PercentEncode(dara.StringValue(name)) + "/fetchCredential"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchConnectionCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询交付物
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactResponse
func (client *Client) GetArtifactWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArtifact"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/artifacts/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取环境信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironmentWithContext(ctx context.Context, projectName *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnvironment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询环境部署信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentDeploymentResponse
func (client *Client) GetEnvironmentDeploymentWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnvironmentDeploymentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnvironmentDeployment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/environmentdeployments/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnvironmentDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineResponse
func (client *Client) GetPipelineWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipeline"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/pipelines/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询仓库绑定
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryResponse
func (client *Client) GetRepositoryWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepository"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/repositories/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务部署信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceDeploymentResponse
func (client *Client) GetServiceDeploymentWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceDeploymentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceDeployment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/servicedeployments/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工具集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetToolsetResponse
func (client *Client) GetToolsetWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetToolsetResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToolset"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/toolsets/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetToolsetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询身份绑定
//
// @param tmpReq - ListConnectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectionsResponse
func (client *Client) ListConnectionsWithContext(ctx context.Context, tmpReq *ListConnectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConnectionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListConnectionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnections"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/connections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询环境列表
//
// @param tmpReq - ListEnvironmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithContext(ctx context.Context, projectName *string, tmpReq *ListEnvironmentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListEnvironmentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironments"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询流水线
//
// @param tmpReq - ListPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelinesResponse
func (client *Client) ListPipelinesWithContext(ctx context.Context, tmpReq *ListPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListPipelinesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelines"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/pipelines"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询项目
//
// @param tmpReq - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, tmpReq *ListProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询服务部署信息
//
// @param tmpReq - ListServiceDeploymentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceDeploymentsResponse
func (client *Client) ListServiceDeploymentsWithContext(ctx context.Context, tmpReq *ListServiceDeploymentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceDeploymentsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListServiceDeploymentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceDeployments"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/servicedeployments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceDeploymentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询任务
//
// @param tmpReq - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, tmpReq *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询工具集
//
// @param tmpReq - ListToolsetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListToolsetsResponse
func (client *Client) ListToolsetsWithContext(ctx context.Context, tmpReq *ListToolsetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListToolsetsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListToolsetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSelector) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, dara.String("labelSelector"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LabelSelectorShrink) {
		query["labelSelector"] = request.LabelSelectorShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListToolsets"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/toolsets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListToolsetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预览环境变更信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewEnvironmentResponse
func (client *Client) PreviewEnvironmentWithContext(ctx context.Context, projectName *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PreviewEnvironmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewEnvironment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments/" + dara.PercentEncode(dara.StringValue(name)) + "/preview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新交付物
//
// @param request - PutArtifactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutArtifactResponse
func (client *Client) PutArtifactWithContext(ctx context.Context, name *string, request *PutArtifactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutArtifactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutArtifact"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/artifacts/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流水线状态
//
// @param request - PutPipelineStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutPipelineStatusResponse
func (client *Client) PutPipelineStatusWithContext(ctx context.Context, name *string, request *PutPipelineStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutPipelineStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutPipelineStatus"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/pipelines/" + dara.PercentEncode(dara.StringValue(name)) + "/status"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutPipelineStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换任务状态
//
// @param request - PutTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutTaskStatusResponse
func (client *Client) PutTaskStatusWithContext(ctx context.Context, name *string, request *PutTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutTaskStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutTaskStatus"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks/" + dara.PercentEncode(dara.StringValue(name)) + "/status"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解析模板中的服务、变量配置
//
// @param request - RenderServicesByTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderServicesByTemplateResponse
func (client *Client) RenderServicesByTemplateWithContext(ctx context.Context, request *RenderServicesByTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenderServicesByTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.ProjectName) {
		body["projectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ServiceNameChanges) {
		body["serviceNameChanges"] = request.ServiceNameChanges
	}

	if !dara.IsNil(request.TemplateName) {
		body["templateName"] = request.TemplateName
	}

	if !dara.IsNil(request.VariableValues) {
		body["variableValues"] = request.VariableValues
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenderServicesByTemplate"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/templates/action/renderServices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenderServicesByTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认并继续执行任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskResponse
func (client *Client) ResumeTaskWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeTask"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks/" + dara.PercentEncode(dara.StringValue(name)) + "/resume"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重试执行任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryTaskResponse
func (client *Client) RetryTaskWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryTask"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks/" + dara.PercentEncode(dara.StringValue(name)) + "/retry"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始执行流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPipelineResponse
func (client *Client) StartPipelineWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartPipelineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPipeline"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/pipelines/" + dara.PercentEncode(dara.StringValue(name)) + "/start"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始执行任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTaskResponse
func (client *Client) StartTaskWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTask"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/tasks/" + dara.PercentEncode(dara.StringValue(name)) + "/start"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新环境（局部更新）
//
// @param request - UpdateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironmentWithContext(ctx context.Context, projectName *string, name *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvironment"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/environments/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新替换应用
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, name *string, request *UpdateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/projects/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// 更新工具集
//
// @param request - UpdateToolsetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateToolsetResponse
func (client *Client) UpdateToolsetWithContext(ctx context.Context, name *string, request *UpdateToolsetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateToolsetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateToolset"),
		Version:     dara.String("2023-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-07-14/toolsets/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateToolsetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
