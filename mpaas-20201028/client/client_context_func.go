// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AddMdsMiniConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMdsMiniConfigResponse
func (client *Client) AddMdsMiniConfigWithContext(ctx context.Context, request *AddMdsMiniConfigRequest, runtime *dara.RuntimeOptions) (_result *AddMdsMiniConfigResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMiniConfigAddJsonStr) {
		body["MpaasMappcenterMiniConfigAddJsonStr"] = request.MpaasMappcenterMiniConfigAddJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMdsMiniConfig"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMdsMiniConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelPushSchedulerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPushSchedulerResponse
func (client *Client) CancelPushSchedulerWithContext(ctx context.Context, request *CancelPushSchedulerRequest, runtime *dara.RuntimeOptions) (_result *CancelPushSchedulerResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.UniqueIds) {
		body["UniqueIds"] = request.UniqueIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPushScheduler"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPushSchedulerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeMcubeMiniTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMcubeMiniTaskStatusResponse
func (client *Client) ChangeMcubeMiniTaskStatusWithContext(ctx context.Context, request *ChangeMcubeMiniTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *ChangeMcubeMiniTaskStatusResponse, _err error) {
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

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeMcubeMiniTaskStatus"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeMcubeMiniTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeMcubeNebulaTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMcubeNebulaTaskStatusResponse
func (client *Client) ChangeMcubeNebulaTaskStatusWithContext(ctx context.Context, request *ChangeMcubeNebulaTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *ChangeMcubeNebulaTaskStatusResponse, _err error) {
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

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeMcubeNebulaTaskStatus"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeMcubeNebulaTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeMcubePublicTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMcubePublicTaskStatusResponse
func (client *Client) ChangeMcubePublicTaskStatusWithContext(ctx context.Context, request *ChangeMcubePublicTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *ChangeMcubePublicTaskStatusResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeMcubePublicTaskStatus"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeMcubePublicTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeMdsCubeTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMdsCubeTaskStatusResponse
func (client *Client) ChangeMdsCubeTaskStatusWithContext(ctx context.Context, request *ChangeMdsCubeTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *ChangeMdsCubeTaskStatusResponse, _err error) {
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

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TemplateResourceId) {
		body["TemplateResourceId"] = request.TemplateResourceId
	}

	if !dara.IsNil(request.TemplateTaskId) {
		body["TemplateTaskId"] = request.TemplateTaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeMdsCubeTaskStatus"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeMdsCubeTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CopyMcdpGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyMcdpGroupResponse
func (client *Client) CopyMcdpGroupWithContext(ctx context.Context, request *CopyMcdpGroupRequest, runtime *dara.RuntimeOptions) (_result *CopyMcdpGroupResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpGroupCopyJsonStr) {
		body["MpaasMappcenterMcdpGroupCopyJsonStr"] = request.MpaasMappcenterMcdpGroupCopyJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyMcdpGroup"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyMcdpGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短链
//
// @param request - CreateLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLinkResponse
func (client *Client) CreateLinkWithContext(ctx context.Context, request *CreateLinkRequest, runtime *dara.RuntimeOptions) (_result *CreateLinkResponse, _err error) {
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

	if !dara.IsNil(request.Cors) {
		body["Cors"] = request.Cors
	}

	if !dara.IsNil(request.Domain) {
		body["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Dynamicfield) {
		body["Dynamicfield"] = request.Dynamicfield
	}

	if !dara.IsNil(request.TargetUrl) {
		body["TargetUrl"] = request.TargetUrl
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLink"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcdpGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcdpGroupResponse
func (client *Client) CreateMcdpGroupWithContext(ctx context.Context, request *CreateMcdpGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMcdpGroupResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpGroupCreateJsonStr) {
		body["MpaasMappcenterMcdpGroupCreateJsonStr"] = request.MpaasMappcenterMcdpGroupCreateJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcdpGroup"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcdpGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcdpMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcdpMaterialResponse
func (client *Client) CreateMcdpMaterialWithContext(ctx context.Context, request *CreateMcdpMaterialRequest, runtime *dara.RuntimeOptions) (_result *CreateMcdpMaterialResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpMaterialCreateJsonStr) {
		body["MpaasMappcenterMcdpMaterialCreateJsonStr"] = request.MpaasMappcenterMcdpMaterialCreateJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcdpMaterial"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcdpMaterialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcdpZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcdpZoneResponse
func (client *Client) CreateMcdpZoneWithContext(ctx context.Context, request *CreateMcdpZoneRequest, runtime *dara.RuntimeOptions) (_result *CreateMcdpZoneResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpZoneCreateJsonStr) {
		body["MpaasMappcenterMcdpZoneCreateJsonStr"] = request.MpaasMappcenterMcdpZoneCreateJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcdpZone"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcdpZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热修复资源
//
// @param request - CreateMcubeHotpatchResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeHotpatchResourceResponse
func (client *Client) CreateMcubeHotpatchResourceWithContext(ctx context.Context, request *CreateMcubeHotpatchResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeHotpatchResourceResponse, _err error) {
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

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FixDesc) {
		body["FixDesc"] = request.FixDesc
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ProductVersion) {
		body["ProductVersion"] = request.ProductVersion
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeHotpatchResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeHotpatchResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热修复回滚任务
//
// @param request - CreateMcubeHotpatchRollbackTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeHotpatchRollbackTaskResponse
func (client *Client) CreateMcubeHotpatchRollbackTaskWithContext(ctx context.Context, request *CreateMcubeHotpatchRollbackTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeHotpatchRollbackTaskResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductVersion) {
		body["ProductVersion"] = request.ProductVersion
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeHotpatchRollbackTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeHotpatchRollbackTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热修复发布任务
//
// @param request - CreateMcubeHotpatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeHotpatchTaskResponse
func (client *Client) CreateMcubeHotpatchTaskWithContext(ctx context.Context, request *CreateMcubeHotpatchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeHotpatchTaskResponse, _err error) {
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

	if !dara.IsNil(request.GreyConfigInfo) {
		body["GreyConfigInfo"] = request.GreyConfigInfo
	}

	if !dara.IsNil(request.GreyEndtimeData) {
		body["GreyEndtimeData"] = request.GreyEndtimeData
	}

	if !dara.IsNil(request.GreyNum) {
		body["GreyNum"] = request.GreyNum
	}

	if !dara.IsNil(request.Memo) {
		body["Memo"] = request.Memo
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	if !dara.IsNil(request.PublishType) {
		body["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.SyncMode) {
		body["SyncMode"] = request.SyncMode
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistIds) {
		body["WhitelistIds"] = request.WhitelistIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeHotpatchTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeHotpatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeMiniAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeMiniAppResponse
func (client *Client) CreateMcubeMiniAppWithContext(ctx context.Context, request *CreateMcubeMiniAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeMiniAppResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.H5Name) {
		body["H5Name"] = request.H5Name
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeMiniApp"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeMiniAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeMiniTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeMiniTaskResponse
func (client *Client) CreateMcubeMiniTaskWithContext(ctx context.Context, request *CreateMcubeMiniTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeMiniTaskResponse, _err error) {
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

	if !dara.IsNil(request.GreyConfigInfo) {
		body["GreyConfigInfo"] = request.GreyConfigInfo
	}

	if !dara.IsNil(request.GreyEndtimeData) {
		body["GreyEndtimeData"] = request.GreyEndtimeData
	}

	if !dara.IsNil(request.GreyNum) {
		body["GreyNum"] = request.GreyNum
	}

	if !dara.IsNil(request.Memo) {
		body["Memo"] = request.Memo
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	if !dara.IsNil(request.PublishType) {
		body["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistIds) {
		body["WhitelistIds"] = request.WhitelistIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeMiniTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeMiniTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeNebulaAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeNebulaAppResponse
func (client *Client) CreateMcubeNebulaAppWithContext(ctx context.Context, request *CreateMcubeNebulaAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeNebulaAppResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.H5Name) {
		body["H5Name"] = request.H5Name
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeNebulaApp"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeNebulaAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeNebulaResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeNebulaResourceResponse
func (client *Client) CreateMcubeNebulaResourceWithContext(ctx context.Context, request *CreateMcubeNebulaResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeNebulaResourceResponse, _err error) {
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

	if !dara.IsNil(request.AutoInstall) {
		body["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.ClientVersionMax) {
		body["ClientVersionMax"] = request.ClientVersionMax
	}

	if !dara.IsNil(request.ClientVersionMin) {
		body["ClientVersionMin"] = request.ClientVersionMin
	}

	if !dara.IsNil(request.CustomDomainName) {
		body["CustomDomainName"] = request.CustomDomainName
	}

	if !dara.IsNil(request.ExtendInfo) {
		body["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.H5Name) {
		body["H5Name"] = request.H5Name
	}

	if !dara.IsNil(request.H5Version) {
		body["H5Version"] = request.H5Version
	}

	if !dara.IsNil(request.InstallType) {
		body["InstallType"] = request.InstallType
	}

	if !dara.IsNil(request.MainUrl) {
		body["MainUrl"] = request.MainUrl
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RepeatNebula) {
		body["RepeatNebula"] = request.RepeatNebula
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SubUrl) {
		body["SubUrl"] = request.SubUrl
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Vhost) {
		body["Vhost"] = request.Vhost
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeNebulaResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeNebulaResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeNebulaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeNebulaTaskResponse
func (client *Client) CreateMcubeNebulaTaskWithContext(ctx context.Context, request *CreateMcubeNebulaTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeNebulaTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		body["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Creator) {
		body["Creator"] = request.Creator
	}

	if !dara.IsNil(request.GmtCreate) {
		body["GmtCreate"] = request.GmtCreate
	}

	if !dara.IsNil(request.GmtModified) {
		body["GmtModified"] = request.GmtModified
	}

	if !dara.IsNil(request.GmtModifiedStr) {
		body["GmtModifiedStr"] = request.GmtModifiedStr
	}

	if !dara.IsNil(request.GreyConfigInfo) {
		body["GreyConfigInfo"] = request.GreyConfigInfo
	}

	if !dara.IsNil(request.GreyEndtime) {
		body["GreyEndtime"] = request.GreyEndtime
	}

	if !dara.IsNil(request.GreyEndtimeData) {
		body["GreyEndtimeData"] = request.GreyEndtimeData
	}

	if !dara.IsNil(request.GreyEndtimeStr) {
		body["GreyEndtimeStr"] = request.GreyEndtimeStr
	}

	if !dara.IsNil(request.GreyNum) {
		body["GreyNum"] = request.GreyNum
	}

	if !dara.IsNil(request.GreyUrl) {
		body["GreyUrl"] = request.GreyUrl
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Memo) {
		body["Memo"] = request.Memo
	}

	if !dara.IsNil(request.Modifier) {
		body["Modifier"] = request.Modifier
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.Percent) {
		body["Percent"] = request.Percent
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductVersion) {
		body["ProductVersion"] = request.ProductVersion
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	if !dara.IsNil(request.PublishType) {
		body["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["ReleaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.ResIds) {
		body["ResIds"] = request.ResIds
	}

	if !dara.IsNil(request.SerialVersionUID) {
		body["SerialVersionUID"] = request.SerialVersionUID
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.SyncMode) {
		body["SyncMode"] = request.SyncMode
	}

	if !dara.IsNil(request.SyncResult) {
		body["SyncResult"] = request.SyncResult
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TaskVersion) {
		body["TaskVersion"] = request.TaskVersion
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UpgradeNoticeNum) {
		body["UpgradeNoticeNum"] = request.UpgradeNoticeNum
	}

	if !dara.IsNil(request.UpgradeProgress) {
		body["UpgradeProgress"] = request.UpgradeProgress
	}

	if !dara.IsNil(request.WhitelistIds) {
		body["WhitelistIds"] = request.WhitelistIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeNebulaTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeNebulaTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeUpgradePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeUpgradePackageResponse
func (client *Client) CreateMcubeUpgradePackageWithContext(ctx context.Context, request *CreateMcubeUpgradePackageRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeUpgradePackageResponse, _err error) {
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

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.AppstoreUrl) {
		body["AppstoreUrl"] = request.AppstoreUrl
	}

	if !dara.IsNil(request.BundleId) {
		body["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.CustomDomainName) {
		body["CustomDomainName"] = request.CustomDomainName
	}

	if !dara.IsNil(request.Desc) {
		body["Desc"] = request.Desc
	}

	if !dara.IsNil(request.DownloadUrl) {
		body["DownloadUrl"] = request.DownloadUrl
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.HarmonyLabel) {
		body["HarmonyLabel"] = request.HarmonyLabel
	}

	if !dara.IsNil(request.IconFileUrl) {
		body["IconFileUrl"] = request.IconFileUrl
	}

	if !dara.IsNil(request.InstallAmount) {
		body["InstallAmount"] = request.InstallAmount
	}

	if !dara.IsNil(request.IosSymbolfileUrl) {
		body["IosSymbolfileUrl"] = request.IosSymbolfileUrl
	}

	if !dara.IsNil(request.IsEnterprise) {
		body["IsEnterprise"] = request.IsEnterprise
	}

	if !dara.IsNil(request.LargeIconUrl) {
		body["LargeIconUrl"] = request.LargeIconUrl
	}

	if !dara.IsNil(request.NeedCheck) {
		body["NeedCheck"] = request.NeedCheck
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ValidDays) {
		body["ValidDays"] = request.ValidDays
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeUpgradePackage"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeUpgradePackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeUpgradeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeUpgradeTaskResponse
func (client *Client) CreateMcubeUpgradeTaskWithContext(ctx context.Context, request *CreateMcubeUpgradeTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeUpgradeTaskResponse, _err error) {
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

	if !dara.IsNil(request.GreyConfigInfo) {
		body["GreyConfigInfo"] = request.GreyConfigInfo
	}

	if !dara.IsNil(request.GreyEndtimeData) {
		body["GreyEndtimeData"] = request.GreyEndtimeData
	}

	if !dara.IsNil(request.GreyNum) {
		body["GreyNum"] = request.GreyNum
	}

	if !dara.IsNil(request.HistoryForce) {
		body["HistoryForce"] = request.HistoryForce
	}

	if !dara.IsNil(request.Memo) {
		body["Memo"] = request.Memo
	}

	if !dara.IsNil(request.PackageInfoId) {
		body["PackageInfoId"] = request.PackageInfoId
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	if !dara.IsNil(request.PublishType) {
		body["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UpgradeContent) {
		body["UpgradeContent"] = request.UpgradeContent
	}

	if !dara.IsNil(request.UpgradeType) {
		body["UpgradeType"] = request.UpgradeType
	}

	if !dara.IsNil(request.WhitelistIds) {
		body["WhitelistIds"] = request.WhitelistIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeUpgradeTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeUpgradeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeVhostResponse
func (client *Client) CreateMcubeVhostWithContext(ctx context.Context, request *CreateMcubeVhostRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeVhostResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Vhost) {
		body["Vhost"] = request.Vhost
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeVhost"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeVhostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeWhitelistResponse
func (client *Client) CreateMcubeWhitelistWithContext(ctx context.Context, request *CreateMcubeWhitelistRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhiteListName) {
		body["WhiteListName"] = request.WhiteListName
	}

	if !dara.IsNil(request.WhitelistType) {
		body["WhitelistType"] = request.WhitelistType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeWhitelist"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMcubeWhitelistForIdeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeWhitelistForIdeResponse
func (client *Client) CreateMcubeWhitelistForIdeWithContext(ctx context.Context, request *CreateMcubeWhitelistForIdeRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeWhitelistForIdeResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WhitelistValue) {
		body["WhitelistValue"] = request.WhitelistValue
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcubeWhitelistForIde"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcubeWhitelistForIdeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMdsCubeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMdsCubeResourceResponse
func (client *Client) CreateMdsCubeResourceWithContext(ctx context.Context, request *CreateMdsCubeResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateMdsCubeResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidMaxVersion) {
		body["AndroidMaxVersion"] = request.AndroidMaxVersion
	}

	if !dara.IsNil(request.AndroidMinVersion) {
		body["AndroidMinVersion"] = request.AndroidMinVersion
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ExtendInfo) {
		body["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.IosMaxVersion) {
		body["IosMaxVersion"] = request.IosMaxVersion
	}

	if !dara.IsNil(request.IosMinVersion) {
		body["IosMinVersion"] = request.IosMinVersion
	}

	if !dara.IsNil(request.MockDataUrl) {
		body["MockDataUrl"] = request.MockDataUrl
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.PreviewPictureUrl) {
		body["PreviewPictureUrl"] = request.PreviewPictureUrl
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateResourceDesc) {
		body["TemplateResourceDesc"] = request.TemplateResourceDesc
	}

	if !dara.IsNil(request.TemplateResourceVersion) {
		body["TemplateResourceVersion"] = request.TemplateResourceVersion
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMdsCubeResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMdsCubeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMdsCubeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMdsCubeTaskResponse
func (client *Client) CreateMdsCubeTaskWithContext(ctx context.Context, request *CreateMdsCubeTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMdsCubeTaskResponse, _err error) {
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

	if !dara.IsNil(request.GreyConfigInfo) {
		body["GreyConfigInfo"] = request.GreyConfigInfo
	}

	if !dara.IsNil(request.GreyEndtimeData) {
		body["GreyEndtimeData"] = request.GreyEndtimeData
	}

	if !dara.IsNil(request.GreyNum) {
		body["GreyNum"] = request.GreyNum
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	if !dara.IsNil(request.PublishType) {
		body["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.TaskDesc) {
		body["TaskDesc"] = request.TaskDesc
	}

	if !dara.IsNil(request.TemplateResourceId) {
		body["TemplateResourceId"] = request.TemplateResourceId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistIds) {
		body["WhitelistIds"] = request.WhitelistIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMdsCubeTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMdsCubeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMdsCubeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMdsCubeTemplateResponse
func (client *Client) CreateMdsCubeTemplateWithContext(ctx context.Context, request *CreateMdsCubeTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateMdsCubeTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateDesc) {
		body["TemplateDesc"] = request.TemplateDesc
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMdsCubeTemplate"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMdsCubeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMdsMiniprogramTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMdsMiniprogramTaskResponse
func (client *Client) CreateMdsMiniprogramTaskWithContext(ctx context.Context, request *CreateMdsMiniprogramTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMdsMiniprogramTaskResponse, _err error) {
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

	if !dara.IsNil(request.GreyConfigInfo) {
		body["GreyConfigInfo"] = request.GreyConfigInfo
	}

	if !dara.IsNil(request.GreyEndtimeData) {
		body["GreyEndtimeData"] = request.GreyEndtimeData
	}

	if !dara.IsNil(request.GreyNum) {
		body["GreyNum"] = request.GreyNum
	}

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Memo) {
		body["Memo"] = request.Memo
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	if !dara.IsNil(request.PublishType) {
		body["PublishType"] = request.PublishType
	}

	if !dara.IsNil(request.SyncMode) {
		body["SyncMode"] = request.SyncMode
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistIds) {
		body["WhitelistIds"] = request.WhitelistIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMdsMiniprogramTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMdsMiniprogramTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateOpenGlobalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpenGlobalDataResponse
func (client *Client) CreateOpenGlobalDataWithContext(ctx context.Context, request *CreateOpenGlobalDataRequest, runtime *dara.RuntimeOptions) (_result *CreateOpenGlobalDataResponse, _err error) {
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

	if !dara.IsNil(request.AppMaxVersion) {
		body["AppMaxVersion"] = request.AppMaxVersion
	}

	if !dara.IsNil(request.AppMinVersion) {
		body["AppMinVersion"] = request.AppMinVersion
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ExtAttrStr) {
		body["ExtAttrStr"] = request.ExtAttrStr
	}

	if !dara.IsNil(request.MaxUid) {
		body["MaxUid"] = request.MaxUid
	}

	if !dara.IsNil(request.MinUid) {
		body["MinUid"] = request.MinUid
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdMsgId) {
		body["ThirdMsgId"] = request.ThirdMsgId
	}

	if !dara.IsNil(request.Uids) {
		body["Uids"] = request.Uids
	}

	if !dara.IsNil(request.ValidTimeEnd) {
		body["ValidTimeEnd"] = request.ValidTimeEnd
	}

	if !dara.IsNil(request.ValidTimeStart) {
		body["ValidTimeStart"] = request.ValidTimeStart
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOpenGlobalData"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOpenGlobalDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateOpenSingleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpenSingleDataResponse
func (client *Client) CreateOpenSingleDataWithContext(ctx context.Context, request *CreateOpenSingleDataRequest, runtime *dara.RuntimeOptions) (_result *CreateOpenSingleDataResponse, _err error) {
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

	if !dara.IsNil(request.AppMaxVersion) {
		body["AppMaxVersion"] = request.AppMaxVersion
	}

	if !dara.IsNil(request.AppMinVersion) {
		body["AppMinVersion"] = request.AppMinVersion
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CheckOnline) {
		body["CheckOnline"] = request.CheckOnline
	}

	if !dara.IsNil(request.ExtAttrStr) {
		body["ExtAttrStr"] = request.ExtAttrStr
	}

	if !dara.IsNil(request.LinkToken) {
		body["LinkToken"] = request.LinkToken
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdMsgId) {
		body["ThirdMsgId"] = request.ThirdMsgId
	}

	if !dara.IsNil(request.ValidTimeEnd) {
		body["ValidTimeEnd"] = request.ValidTimeEnd
	}

	if !dara.IsNil(request.ValidTimeStart) {
		body["ValidTimeStart"] = request.ValidTimeStart
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOpenSingleData"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOpenSingleDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模版
//
// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithContext(ctx context.Context, request *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DescInfo) {
		body["DescInfo"] = request.DescInfo
	}

	if !dara.IsNil(request.IconUrls) {
		body["IconUrls"] = request.IconUrls
	}

	if !dara.IsNil(request.ImageUrls) {
		body["ImageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.JumpAction) {
		body["JumpAction"] = request.JumpAction
	}

	if !dara.IsNil(request.PushStyle) {
		body["PushStyle"] = request.PushStyle
	}

	if !dara.IsNil(request.ShowStyle) {
		body["ShowStyle"] = request.ShowStyle
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	if !dara.IsNil(request.Variables) {
		body["Variables"] = request.Variables
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCubecardWhitelistContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCubecardWhitelistContentResponse
func (client *Client) DeleteCubecardWhitelistContentWithContext(ctx context.Context, request *DeleteCubecardWhitelistContentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCubecardWhitelistContentResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistId) {
		body["WhitelistId"] = request.WhitelistId
	}

	if !dara.IsNil(request.WhitelistValue) {
		body["WhitelistValue"] = request.WhitelistValue
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCubecardWhitelistContent"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCubecardWhitelistContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcdpAimRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcdpAimResponse
func (client *Client) DeleteMcdpAimWithContext(ctx context.Context, request *DeleteMcdpAimRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcdpAimResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpAimDeleteJsonStr) {
		body["MpaasMappcenterMcdpAimDeleteJsonStr"] = request.MpaasMappcenterMcdpAimDeleteJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcdpAim"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcdpAimResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcdpCrowdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcdpCrowdResponse
func (client *Client) DeleteMcdpCrowdWithContext(ctx context.Context, request *DeleteMcdpCrowdRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcdpCrowdResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpCrowdDeleteJsonStr) {
		body["MpaasMappcenterMcdpCrowdDeleteJsonStr"] = request.MpaasMappcenterMcdpCrowdDeleteJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcdpCrowd"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcdpCrowdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcdpZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcdpZoneResponse
func (client *Client) DeleteMcdpZoneWithContext(ctx context.Context, request *DeleteMcdpZoneRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcdpZoneResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMcdpZoneDeleteJsonStr) {
		body["MpaasMappcenterMcdpZoneDeleteJsonStr"] = request.MpaasMappcenterMcdpZoneDeleteJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcdpZone"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcdpZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除热修复资源包
//
// @param request - DeleteMcubeHotpatchResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcubeHotpatchResourceResponse
func (client *Client) DeleteMcubeHotpatchResourceWithContext(ctx context.Context, request *DeleteMcubeHotpatchResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcubeHotpatchResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		body["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcubeHotpatchResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcubeHotpatchResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcubeMiniAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcubeMiniAppResponse
func (client *Client) DeleteMcubeMiniAppWithContext(ctx context.Context, request *DeleteMcubeMiniAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcubeMiniAppResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcubeMiniApp"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcubeMiniAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcubeNebulaAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcubeNebulaAppResponse
func (client *Client) DeleteMcubeNebulaAppWithContext(ctx context.Context, request *DeleteMcubeNebulaAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcubeNebulaAppResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcubeNebulaApp"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcubeNebulaAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcubeUpgradeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcubeUpgradeResourceResponse
func (client *Client) DeleteMcubeUpgradeResourceWithContext(ctx context.Context, request *DeleteMcubeUpgradeResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcubeUpgradeResourceResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcubeUpgradeResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcubeUpgradeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMcubeWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcubeWhitelistResponse
func (client *Client) DeleteMcubeWhitelistWithContext(ctx context.Context, request *DeleteMcubeWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcubeWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcubeWhitelist"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcubeWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMdsCubeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMdsCubeTemplateResponse
func (client *Client) DeleteMdsCubeTemplateWithContext(ctx context.Context, request *DeleteMdsCubeTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteMdsCubeTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMdsCubeTemplate"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMdsCubeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMdsWhitelistContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMdsWhitelistContentResponse
func (client *Client) DeleteMdsWhitelistContentWithContext(ctx context.Context, request *DeleteMdsWhitelistContentRequest, runtime *dara.RuntimeOptions) (_result *DeleteMdsWhitelistContentResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistId) {
		body["WhitelistId"] = request.WhitelistId
	}

	if !dara.IsNil(request.WhitelistValue) {
		body["WhitelistValue"] = request.WhitelistValue
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMdsWhitelistContent"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMdsWhitelistContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模版
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithContext(ctx context.Context, request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExistMcubeRsaKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExistMcubeRsaKeyResponse
func (client *Client) ExistMcubeRsaKeyWithContext(ctx context.Context, request *ExistMcubeRsaKeyRequest, runtime *dara.RuntimeOptions) (_result *ExistMcubeRsaKeyResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExistMcubeRsaKey"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExistMcubeRsaKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExportMappCenterAppConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportMappCenterAppConfigResponse
func (client *Client) ExportMappCenterAppConfigWithContext(ctx context.Context, request *ExportMappCenterAppConfigRequest, runtime *dara.RuntimeOptions) (_result *ExportMappCenterAppConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApkFileUrl) {
		body["ApkFileUrl"] = request.ApkFileUrl
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CertRsaBase64) {
		body["CertRsaBase64"] = request.CertRsaBase64
	}

	if !dara.IsNil(request.Identifier) {
		body["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportMappCenterAppConfig"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportMappCenterAppConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetFileTokenForUploadToMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileTokenForUploadToMsaResponse
func (client *Client) GetFileTokenForUploadToMsaWithContext(ctx context.Context, request *GetFileTokenForUploadToMsaRequest, runtime *dara.RuntimeOptions) (_result *GetFileTokenForUploadToMsaResponse, _err error) {
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

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileTokenForUploadToMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileTokenForUploadToMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看全部审核记录
//
// @param request - GetGameReviewByStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGameReviewByStatusResponse
func (client *Client) GetGameReviewByStatusWithContext(ctx context.Context, request *GetGameReviewByStatusRequest, runtime *dara.RuntimeOptions) (_result *GetGameReviewByStatusResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReviewStatus) {
		body["ReviewStatus"] = request.ReviewStatus
	}

	if !dara.IsNil(request.SortMode) {
		body["SortMode"] = request.SortMode
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGameReviewByStatus"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGameReviewByStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLogUrlInMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogUrlInMsaResponse
func (client *Client) GetLogUrlInMsaWithContext(ctx context.Context, request *GetLogUrlInMsaRequest, runtime *dara.RuntimeOptions) (_result *GetLogUrlInMsaResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLogUrlInMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogUrlInMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMcubeFileTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcubeFileTokenResponse
func (client *Client) GetMcubeFileTokenWithContext(ctx context.Context, request *GetMcubeFileTokenRequest, runtime *dara.RuntimeOptions) (_result *GetMcubeFileTokenResponse, _err error) {
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

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcubeFileToken"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcubeFileTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMcubeNebulaResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcubeNebulaResourceResponse
func (client *Client) GetMcubeNebulaResourceWithContext(ctx context.Context, request *GetMcubeNebulaResourceRequest, runtime *dara.RuntimeOptions) (_result *GetMcubeNebulaResourceResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcubeNebulaResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcubeNebulaResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMcubeNebulaTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcubeNebulaTaskDetailResponse
func (client *Client) GetMcubeNebulaTaskDetailWithContext(ctx context.Context, request *GetMcubeNebulaTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *GetMcubeNebulaTaskDetailResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcubeNebulaTaskDetail"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcubeNebulaTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMcubeUpgradePackageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcubeUpgradePackageInfoResponse
func (client *Client) GetMcubeUpgradePackageInfoWithContext(ctx context.Context, request *GetMcubeUpgradePackageInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMcubeUpgradePackageInfoResponse, _err error) {
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

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcubeUpgradePackageInfo"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcubeUpgradePackageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMcubeUpgradeTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcubeUpgradeTaskInfoResponse
func (client *Client) GetMcubeUpgradeTaskInfoWithContext(ctx context.Context, request *GetMcubeUpgradeTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMcubeUpgradeTaskInfoResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcubeUpgradeTaskInfo"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcubeUpgradeTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMdsMiniConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMdsMiniConfigResponse
func (client *Client) GetMdsMiniConfigWithContext(ctx context.Context, request *GetMdsMiniConfigRequest, runtime *dara.RuntimeOptions) (_result *GetMdsMiniConfigResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMdsMiniConfig"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMdsMiniConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模版
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithContext(ctx context.Context, request *GetTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserAppDonwloadUrlInMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAppDonwloadUrlInMsaResponse
func (client *Client) GetUserAppDonwloadUrlInMsaWithContext(ctx context.Context, request *GetUserAppDonwloadUrlInMsaRequest, runtime *dara.RuntimeOptions) (_result *GetUserAppDonwloadUrlInMsaResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserAppDonwloadUrlInMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserAppDonwloadUrlInMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserAppEnhanceProcessInMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAppEnhanceProcessInMsaResponse
func (client *Client) GetUserAppEnhanceProcessInMsaWithContext(ctx context.Context, request *GetUserAppEnhanceProcessInMsaRequest, runtime *dara.RuntimeOptions) (_result *GetUserAppEnhanceProcessInMsaResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserAppEnhanceProcessInMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserAppEnhanceProcessInMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserAppUploadProcessInMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAppUploadProcessInMsaResponse
func (client *Client) GetUserAppUploadProcessInMsaWithContext(ctx context.Context, request *GetUserAppUploadProcessInMsaRequest, runtime *dara.RuntimeOptions) (_result *GetUserAppUploadProcessInMsaResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserAppUploadProcessInMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserAppUploadProcessInMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报表
//
// @param request - ListAnalysisCoreIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnalysisCoreIndexResponse
func (client *Client) ListAnalysisCoreIndexWithContext(ctx context.Context, request *ListAnalysisCoreIndexRequest, runtime *dara.RuntimeOptions) (_result *ListAnalysisCoreIndexResponse, _err error) {
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

	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnalysisCoreIndex"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnalysisCoreIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcdpAimRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcdpAimResponse
func (client *Client) ListMcdpAimWithContext(ctx context.Context, request *ListMcdpAimRequest, runtime *dara.RuntimeOptions) (_result *ListMcdpAimResponse, _err error) {
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

	if !dara.IsNil(request.EmptyTag) {
		body["EmptyTag"] = request.EmptyTag
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcdpAim"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcdpAimResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询热修复资源包列表
//
// @param request - ListMcubeHotpatchResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeHotpatchResourcesResponse
func (client *Client) ListMcubeHotpatchResourcesWithContext(ctx context.Context, request *ListMcubeHotpatchResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeHotpatchResourcesResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeHotpatchResources"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeHotpatchResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询热修复发布任务列表
//
// @param request - ListMcubeHotpatchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeHotpatchTasksResponse
func (client *Client) ListMcubeHotpatchTasksWithContext(ctx context.Context, request *ListMcubeHotpatchTasksRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeHotpatchTasksResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeHotpatchTasks"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeHotpatchTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeMiniAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeMiniAppsResponse
func (client *Client) ListMcubeMiniAppsWithContext(ctx context.Context, request *ListMcubeMiniAppsRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeMiniAppsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeMiniApps"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeMiniAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeMiniPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeMiniPackagesResponse
func (client *Client) ListMcubeMiniPackagesWithContext(ctx context.Context, request *ListMcubeMiniPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeMiniPackagesResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.PackageTypes) {
		body["PackageTypes"] = request.PackageTypes
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeMiniPackages"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeMiniPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeMiniTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeMiniTasksResponse
func (client *Client) ListMcubeMiniTasksWithContext(ctx context.Context, request *ListMcubeMiniTasksRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeMiniTasksResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeMiniTasks"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeMiniTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeNebulaAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeNebulaAppsResponse
func (client *Client) ListMcubeNebulaAppsWithContext(ctx context.Context, request *ListMcubeNebulaAppsRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeNebulaAppsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeNebulaApps"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeNebulaAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeNebulaResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeNebulaResourcesResponse
func (client *Client) ListMcubeNebulaResourcesWithContext(ctx context.Context, request *ListMcubeNebulaResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeNebulaResourcesResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeNebulaResources"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeNebulaResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeNebulaTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeNebulaTasksResponse
func (client *Client) ListMcubeNebulaTasksWithContext(ctx context.Context, request *ListMcubeNebulaTasksRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeNebulaTasksResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeNebulaTasks"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeNebulaTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeUpgradePackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeUpgradePackagesResponse
func (client *Client) ListMcubeUpgradePackagesWithContext(ctx context.Context, request *ListMcubeUpgradePackagesRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeUpgradePackagesResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeUpgradePackages"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeUpgradePackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeUpgradeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeUpgradeTasksResponse
func (client *Client) ListMcubeUpgradeTasksWithContext(ctx context.Context, request *ListMcubeUpgradeTasksRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeUpgradeTasksResponse, _err error) {
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

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeUpgradeTasks"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeUpgradeTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMcubeWhitelistsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeWhitelistsResponse
func (client *Client) ListMcubeWhitelistsWithContext(ctx context.Context, request *ListMcubeWhitelistsRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeWhitelistsResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WhitelistName) {
		body["WhitelistName"] = request.WhitelistName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcubeWhitelists"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcubeWhitelistsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMdsCubeResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMdsCubeResourcesResponse
func (client *Client) ListMdsCubeResourcesWithContext(ctx context.Context, request *ListMdsCubeResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListMdsCubeResourcesResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.Test) {
		body["test"] = request.Test
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMdsCubeResources"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMdsCubeResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMdsCubeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMdsCubeTasksResponse
func (client *Client) ListMdsCubeTasksWithContext(ctx context.Context, request *ListMdsCubeTasksRequest, runtime *dara.RuntimeOptions) (_result *ListMdsCubeTasksResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateResourceId) {
		body["TemplateResourceId"] = request.TemplateResourceId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMdsCubeTasks"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMdsCubeTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMdsCubeTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMdsCubeTemplatesResponse
func (client *Client) ListMdsCubeTemplatesWithContext(ctx context.Context, request *ListMdsCubeTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListMdsCubeTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMdsCubeTemplates"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMdsCubeTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMgsApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMgsApiResponse
func (client *Client) ListMgsApiWithContext(ctx context.Context, request *ListMgsApiRequest, runtime *dara.RuntimeOptions) (_result *ListMgsApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiStatus) {
		body["ApiStatus"] = request.ApiStatus
	}

	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Format) {
		body["Format"] = request.Format
	}

	if !dara.IsNil(request.Host) {
		body["Host"] = request.Host
	}

	if !dara.IsNil(request.NeedEncrypt) {
		body["NeedEncrypt"] = request.NeedEncrypt
	}

	if !dara.IsNil(request.NeedEtag) {
		body["NeedEtag"] = request.NeedEtag
	}

	if !dara.IsNil(request.NeedSign) {
		body["NeedSign"] = request.NeedSign
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OptFuzzy) {
		body["OptFuzzy"] = request.OptFuzzy
	}

	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SysId) {
		body["SysId"] = request.SysId
	}

	if !dara.IsNil(request.SysName) {
		body["SysName"] = request.SysName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMgsApi"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMgsApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询模版列表
//
// @param request - ListTemplatePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatePageResponse
func (client *Client) ListTemplatePageWithContext(ctx context.Context, request *ListTemplatePageRequest, runtime *dara.RuntimeOptions) (_result *ListTemplatePageResponse, _err error) {
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

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplatePage"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatePageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OCR通用接口
//
// @param request - MTRSOCRServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MTRSOCRServiceResponse
func (client *Client) MTRSOCRServiceWithContext(ctx context.Context, request *MTRSOCRServiceRequest, runtime *dara.RuntimeOptions) (_result *MTRSOCRServiceResponse, _err error) {
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

	if !dara.IsNil(request.ImageRaw) {
		body["ImageRaw"] = request.ImageRaw
	}

	if !dara.IsNil(request.Mask) {
		body["Mask"] = request.Mask
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MTRSOCRService"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MTRSOCRServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PushBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushBindResponse
func (client *Client) PushBindWithContext(ctx context.Context, request *PushBindRequest, runtime *dara.RuntimeOptions) (_result *PushBindResponse, _err error) {
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

	if !dara.IsNil(request.DeliveryToken) {
		body["DeliveryToken"] = request.DeliveryToken
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PhoneNumber) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushBind"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushBindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - PushBroadcastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushBroadcastResponse
func (client *Client) PushBroadcastWithContext(ctx context.Context, tmpReq *PushBroadcastRequest, runtime *dara.RuntimeOptions) (_result *PushBroadcastResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushBroadcastShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotifyLevel) {
		request.NotifyLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyLevel, dara.String("NotifyLevel"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ThirdChannelCategory) {
		request.ThirdChannelCategoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdChannelCategory, dara.String("ThirdChannelCategory"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidChannel) {
		body["AndroidChannel"] = request.AndroidChannel
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BindEndTime) {
		body["BindEndTime"] = request.BindEndTime
	}

	if !dara.IsNil(request.BindPeriod) {
		body["BindPeriod"] = request.BindPeriod
	}

	if !dara.IsNil(request.BindStartTime) {
		body["BindStartTime"] = request.BindStartTime
	}

	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Classification) {
		body["Classification"] = request.Classification
	}

	if !dara.IsNil(request.DeliveryType) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.ExpiredSeconds) {
		body["ExpiredSeconds"] = request.ExpiredSeconds
	}

	if !dara.IsNil(request.ExtendedParams) {
		body["ExtendedParams"] = request.ExtendedParams
	}

	if !dara.IsNil(request.MiChannelId) {
		body["MiChannelId"] = request.MiChannelId
	}

	if !dara.IsNil(request.Msgkey) {
		body["Msgkey"] = request.Msgkey
	}

	if !dara.IsNil(request.NotifyLevelShrink) {
		body["NotifyLevel"] = request.NotifyLevelShrink
	}

	if !dara.IsNil(request.NotifyType) {
		body["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.PushAction) {
		body["PushAction"] = request.PushAction
	}

	if !dara.IsNil(request.PushStatus) {
		body["PushStatus"] = request.PushStatus
	}

	if !dara.IsNil(request.Silent) {
		body["Silent"] = request.Silent
	}

	if !dara.IsNil(request.StrategyContent) {
		body["StrategyContent"] = request.StrategyContent
	}

	if !dara.IsNil(request.StrategyType) {
		body["StrategyType"] = request.StrategyType
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TemplateKeyValue) {
		body["TemplateKeyValue"] = request.TemplateKeyValue
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdChannelCategoryShrink) {
		body["ThirdChannelCategory"] = request.ThirdChannelCategoryShrink
	}

	if !dara.IsNil(request.TimeMode) {
		body["TimeMode"] = request.TimeMode
	}

	if !dara.IsNil(request.TransparentMessagePayload) {
		body["TransparentMessagePayload"] = request.TransparentMessagePayload
	}

	if !dara.IsNil(request.TransparentMessageUrgency) {
		body["TransparentMessageUrgency"] = request.TransparentMessageUrgency
	}

	if !dara.IsNil(request.UnBindEndTime) {
		body["UnBindEndTime"] = request.UnBindEndTime
	}

	if !dara.IsNil(request.UnBindPeriod) {
		body["UnBindPeriod"] = request.UnBindPeriod
	}

	if !dara.IsNil(request.UnBindStartTime) {
		body["UnBindStartTime"] = request.UnBindStartTime
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushBroadcast"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushBroadcastResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - PushMultipleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMultipleResponse
func (client *Client) PushMultipleWithContext(ctx context.Context, tmpReq *PushMultipleRequest, runtime *dara.RuntimeOptions) (_result *PushMultipleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushMultipleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotifyLevel) {
		request.NotifyLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyLevel, dara.String("NotifyLevel"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ThirdChannelCategory) {
		request.ThirdChannelCategoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdChannelCategory, dara.String("ThirdChannelCategory"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActivityContentState) {
		body["ActivityContentState"] = request.ActivityContentState
	}

	if !dara.IsNil(request.ActivityEvent) {
		body["ActivityEvent"] = request.ActivityEvent
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Classification) {
		body["Classification"] = request.Classification
	}

	if !dara.IsNil(request.DeliveryType) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.DismissalDate) {
		body["DismissalDate"] = request.DismissalDate
	}

	if !dara.IsNil(request.ExpiredSeconds) {
		body["ExpiredSeconds"] = request.ExpiredSeconds
	}

	if !dara.IsNil(request.ExtendedParams) {
		body["ExtendedParams"] = request.ExtendedParams
	}

	if !dara.IsNil(request.MiChannelId) {
		body["MiChannelId"] = request.MiChannelId
	}

	if !dara.IsNil(request.NotifyLevelShrink) {
		body["NotifyLevel"] = request.NotifyLevelShrink
	}

	if !dara.IsNil(request.NotifyType) {
		body["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.PushAction) {
		body["PushAction"] = request.PushAction
	}

	if !dara.IsNil(request.Silent) {
		body["Silent"] = request.Silent
	}

	if !dara.IsNil(request.StrategyContent) {
		body["StrategyContent"] = request.StrategyContent
	}

	if !dara.IsNil(request.StrategyType) {
		body["StrategyType"] = request.StrategyType
	}

	if !dara.IsNil(request.TargetMsg) {
		body["TargetMsg"] = request.TargetMsg
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdChannelCategoryShrink) {
		body["ThirdChannelCategory"] = request.ThirdChannelCategoryShrink
	}

	if !dara.IsNil(request.TransparentMessagePayload) {
		body["TransparentMessagePayload"] = request.TransparentMessagePayload
	}

	if !dara.IsNil(request.TransparentMessageUrgency) {
		body["TransparentMessageUrgency"] = request.TransparentMessageUrgency
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMultiple"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMultipleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备状态信息
//
// @param request - PushQueryDeviceStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushQueryDeviceStateResponse
func (client *Client) PushQueryDeviceStateWithContext(ctx context.Context, request *PushQueryDeviceStateRequest, runtime *dara.RuntimeOptions) (_result *PushQueryDeviceStateResponse, _err error) {
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

	if !dara.IsNil(request.Target) {
		body["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetType) {
		body["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushQueryDeviceState"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushQueryDeviceStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PushReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushReportResponse
func (client *Client) PushReportWithContext(ctx context.Context, request *PushReportRequest, runtime *dara.RuntimeOptions) (_result *PushReportResponse, _err error) {
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

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ConnectType) {
		body["ConnectType"] = request.ConnectType
	}

	if !dara.IsNil(request.DeliveryToken) {
		body["DeliveryToken"] = request.DeliveryToken
	}

	if !dara.IsNil(request.Imei) {
		body["Imei"] = request.Imei
	}

	if !dara.IsNil(request.Imsi) {
		body["Imsi"] = request.Imsi
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PushVersion) {
		body["PushVersion"] = request.PushVersion
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdChannel) {
		body["ThirdChannel"] = request.ThirdChannel
	}

	if !dara.IsNil(request.ThirdChannelDeviceToken) {
		body["ThirdChannelDeviceToken"] = request.ThirdChannelDeviceToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushReport"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 极简推送
//
// @param tmpReq - PushSimpleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushSimpleResponse
func (client *Client) PushSimpleWithContext(ctx context.Context, tmpReq *PushSimpleRequest, runtime *dara.RuntimeOptions) (_result *PushSimpleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushSimpleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotifyLevel) {
		request.NotifyLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyLevel, dara.String("NotifyLevel"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ThirdChannelCategory) {
		request.ThirdChannelCategoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdChannelCategory, dara.String("ThirdChannelCategory"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActivityContentState) {
		body["ActivityContentState"] = request.ActivityContentState
	}

	if !dara.IsNil(request.ActivityEvent) {
		body["ActivityEvent"] = request.ActivityEvent
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Classification) {
		body["Classification"] = request.Classification
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DeliveryType) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.DismissalDate) {
		body["DismissalDate"] = request.DismissalDate
	}

	if !dara.IsNil(request.ExpiredSeconds) {
		body["ExpiredSeconds"] = request.ExpiredSeconds
	}

	if !dara.IsNil(request.ExtendedParams) {
		body["ExtendedParams"] = request.ExtendedParams
	}

	if !dara.IsNil(request.IconUrls) {
		body["IconUrls"] = request.IconUrls
	}

	if !dara.IsNil(request.ImageUrls) {
		body["ImageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.MiChannelId) {
		body["MiChannelId"] = request.MiChannelId
	}

	if !dara.IsNil(request.NotifyLevelShrink) {
		body["NotifyLevel"] = request.NotifyLevelShrink
	}

	if !dara.IsNil(request.NotifyType) {
		body["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.PushAction) {
		body["PushAction"] = request.PushAction
	}

	if !dara.IsNil(request.PushStyle) {
		body["PushStyle"] = request.PushStyle
	}

	if !dara.IsNil(request.Silent) {
		body["Silent"] = request.Silent
	}

	if !dara.IsNil(request.SmsSignName) {
		body["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.SmsStrategy) {
		body["SmsStrategy"] = request.SmsStrategy
	}

	if !dara.IsNil(request.SmsTemplateCode) {
		body["SmsTemplateCode"] = request.SmsTemplateCode
	}

	if !dara.IsNil(request.SmsTemplateParam) {
		body["SmsTemplateParam"] = request.SmsTemplateParam
	}

	if !dara.IsNil(request.StrategyContent) {
		body["StrategyContent"] = request.StrategyContent
	}

	if !dara.IsNil(request.StrategyType) {
		body["StrategyType"] = request.StrategyType
	}

	if !dara.IsNil(request.TargetMsgkey) {
		body["TargetMsgkey"] = request.TargetMsgkey
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdChannelCategoryShrink) {
		body["ThirdChannelCategory"] = request.ThirdChannelCategoryShrink
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.TransparentMessagePayload) {
		body["TransparentMessagePayload"] = request.TransparentMessagePayload
	}

	if !dara.IsNil(request.TransparentMessageUrgency) {
		body["TransparentMessageUrgency"] = request.TransparentMessageUrgency
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushSimple"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushSimpleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - PushTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushTemplateResponse
func (client *Client) PushTemplateWithContext(ctx context.Context, tmpReq *PushTemplateRequest, runtime *dara.RuntimeOptions) (_result *PushTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotifyLevel) {
		request.NotifyLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyLevel, dara.String("NotifyLevel"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ThirdChannelCategory) {
		request.ThirdChannelCategoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ThirdChannelCategory, dara.String("ThirdChannelCategory"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActivityContentState) {
		body["ActivityContentState"] = request.ActivityContentState
	}

	if !dara.IsNil(request.ActivityEvent) {
		body["ActivityEvent"] = request.ActivityEvent
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Classification) {
		body["Classification"] = request.Classification
	}

	if !dara.IsNil(request.DeliveryType) {
		body["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.DismissalDate) {
		body["DismissalDate"] = request.DismissalDate
	}

	if !dara.IsNil(request.ExpiredSeconds) {
		body["ExpiredSeconds"] = request.ExpiredSeconds
	}

	if !dara.IsNil(request.ExtendedParams) {
		body["ExtendedParams"] = request.ExtendedParams
	}

	if !dara.IsNil(request.MiChannelId) {
		body["MiChannelId"] = request.MiChannelId
	}

	if !dara.IsNil(request.NotifyLevelShrink) {
		body["NotifyLevel"] = request.NotifyLevelShrink
	}

	if !dara.IsNil(request.NotifyType) {
		body["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.PushAction) {
		body["PushAction"] = request.PushAction
	}

	if !dara.IsNil(request.Silent) {
		body["Silent"] = request.Silent
	}

	if !dara.IsNil(request.SmsSignName) {
		body["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.SmsStrategy) {
		body["SmsStrategy"] = request.SmsStrategy
	}

	if !dara.IsNil(request.SmsTemplateCode) {
		body["SmsTemplateCode"] = request.SmsTemplateCode
	}

	if !dara.IsNil(request.SmsTemplateParam) {
		body["SmsTemplateParam"] = request.SmsTemplateParam
	}

	if !dara.IsNil(request.StrategyContent) {
		body["StrategyContent"] = request.StrategyContent
	}

	if !dara.IsNil(request.StrategyType) {
		body["StrategyType"] = request.StrategyType
	}

	if !dara.IsNil(request.TargetMsgkey) {
		body["TargetMsgkey"] = request.TargetMsgkey
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TemplateKeyValue) {
		body["TemplateKeyValue"] = request.TemplateKeyValue
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ThirdChannelCategoryShrink) {
		body["ThirdChannelCategory"] = request.ThirdChannelCategoryShrink
	}

	if !dara.IsNil(request.TransparentMessagePayload) {
		body["TransparentMessagePayload"] = request.TransparentMessagePayload
	}

	if !dara.IsNil(request.TransparentMessageUrgency) {
		body["TransparentMessageUrgency"] = request.TransparentMessageUrgency
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushTemplate"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PushUnBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushUnBindResponse
func (client *Client) PushUnBindWithContext(ctx context.Context, request *PushUnBindRequest, runtime *dara.RuntimeOptions) (_result *PushUnBindResponse, _err error) {
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

	if !dara.IsNil(request.DeliveryToken) {
		body["DeliveryToken"] = request.DeliveryToken
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushUnBind"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushUnBindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCubecardFiletokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubecardFiletokenResponse
func (client *Client) QueryCubecardFiletokenWithContext(ctx context.Context, request *QueryCubecardFiletokenRequest, runtime *dara.RuntimeOptions) (_result *QueryCubecardFiletokenResponse, _err error) {
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

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCubecardFiletoken"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCubecardFiletokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Device+服务的
//
// @param request - QueryInfoFromMdpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInfoFromMdpResponse
func (client *Client) QueryInfoFromMdpWithContext(ctx context.Context, request *QueryInfoFromMdpRequest, runtime *dara.RuntimeOptions) (_result *QueryInfoFromMdpResponse, _err error) {
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

	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.MobileMd5) {
		body["MobileMd5"] = request.MobileMd5
	}

	if !dara.IsNil(request.MobileSha256) {
		body["MobileSha256"] = request.MobileSha256
	}

	if !dara.IsNil(request.MobileSm3) {
		body["MobileSm3"] = request.MobileSm3
	}

	if !dara.IsNil(request.RiskScene) {
		body["RiskScene"] = request.RiskScene
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInfoFromMdp"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInfoFromMdpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询短链
//
// @param request - QueryLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLinkResponse
func (client *Client) QueryLinkWithContext(ctx context.Context, request *QueryLinkRequest, runtime *dara.RuntimeOptions) (_result *QueryLinkResponse, _err error) {
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

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLink"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMappCenterAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMappCenterAppResponse
func (client *Client) QueryMappCenterAppWithContext(ctx context.Context, request *QueryMappCenterAppRequest, runtime *dara.RuntimeOptions) (_result *QueryMappCenterAppResponse, _err error) {
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

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMappCenterApp"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMappCenterAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMcdpAimRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcdpAimResponse
func (client *Client) QueryMcdpAimWithContext(ctx context.Context, request *QueryMcdpAimRequest, runtime *dara.RuntimeOptions) (_result *QueryMcdpAimResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMcdpAim"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMcdpAimResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMcdpZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcdpZoneResponse
func (client *Client) QueryMcdpZoneWithContext(ctx context.Context, request *QueryMcdpZoneRequest, runtime *dara.RuntimeOptions) (_result *QueryMcdpZoneResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMcdpZone"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMcdpZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询热修复发布任务详情
//
// @param request - QueryMcubeHotpatchTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcubeHotpatchTaskDetailResponse
func (client *Client) QueryMcubeHotpatchTaskDetailWithContext(ctx context.Context, request *QueryMcubeHotpatchTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryMcubeHotpatchTaskDetailResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMcubeHotpatchTaskDetail"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMcubeHotpatchTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMcubeMiniPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcubeMiniPackageResponse
func (client *Client) QueryMcubeMiniPackageWithContext(ctx context.Context, request *QueryMcubeMiniPackageRequest, runtime *dara.RuntimeOptions) (_result *QueryMcubeMiniPackageResponse, _err error) {
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

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMcubeMiniPackage"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMcubeMiniPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMcubeMiniTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcubeMiniTaskResponse
func (client *Client) QueryMcubeMiniTaskWithContext(ctx context.Context, request *QueryMcubeMiniTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryMcubeMiniTaskResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMcubeMiniTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMcubeMiniTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMcubeVhostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcubeVhostResponse
func (client *Client) QueryMcubeVhostWithContext(ctx context.Context, request *QueryMcubeVhostRequest, runtime *dara.RuntimeOptions) (_result *QueryMcubeVhostResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMcubeVhost"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMcubeVhostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMdsUpgradeTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMdsUpgradeTaskDetailResponse
func (client *Client) QueryMdsUpgradeTaskDetailWithContext(ctx context.Context, request *QueryMdsUpgradeTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryMdsUpgradeTaskDetailResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMdsUpgradeTaskDetail"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMdsUpgradeTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMgsApipageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMgsApipageResponse
func (client *Client) QueryMgsApipageWithContext(ctx context.Context, request *QueryMgsApipageRequest, runtime *dara.RuntimeOptions) (_result *QueryMgsApipageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiStatus) {
		body["ApiStatus"] = request.ApiStatus
	}

	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Format) {
		body["Format"] = request.Format
	}

	if !dara.IsNil(request.Host) {
		body["Host"] = request.Host
	}

	if !dara.IsNil(request.NeedEncrypt) {
		body["NeedEncrypt"] = request.NeedEncrypt
	}

	if !dara.IsNil(request.NeedEtag) {
		body["NeedEtag"] = request.NeedEtag
	}

	if !dara.IsNil(request.NeedSign) {
		body["NeedSign"] = request.NeedSign
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OptFuzzy) {
		body["OptFuzzy"] = request.OptFuzzy
	}

	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SysId) {
		body["SysId"] = request.SysId
	}

	if !dara.IsNil(request.SysName) {
		body["SysName"] = request.SysName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMgsApipage"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMgsApipageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMgsApirestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMgsApirestResponse
func (client *Client) QueryMgsApirestWithContext(ctx context.Context, request *QueryMgsApirestRequest, runtime *dara.RuntimeOptions) (_result *QueryMgsApirestResponse, _err error) {
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

	if !dara.IsNil(request.Format) {
		body["Format"] = request.Format
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMgsApirest"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMgsApirestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMgsTestreqbodyautogenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMgsTestreqbodyautogenResponse
func (client *Client) QueryMgsTestreqbodyautogenWithContext(ctx context.Context, request *QueryMgsTestreqbodyautogenRequest, runtime *dara.RuntimeOptions) (_result *QueryMgsTestreqbodyautogenResponse, _err error) {
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

	if !dara.IsNil(request.Format) {
		body["Format"] = request.Format
	}

	if !dara.IsNil(request.MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr) {
		body["MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr"] = request.MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMgsTestreqbodyautogen"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMgsTestreqbodyautogenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMpsSchedulerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMpsSchedulerListResponse
func (client *Client) QueryMpsSchedulerListWithContext(ctx context.Context, request *QueryMpsSchedulerListRequest, runtime *dara.RuntimeOptions) (_result *QueryMpsSchedulerListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.UniqueId) {
		body["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMpsSchedulerList"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMpsSchedulerListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询风险信息
//
// @param request - QueryMscpRiskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMscpRiskInfoResponse
func (client *Client) QueryMscpRiskInfoWithContext(ctx context.Context, request *QueryMscpRiskInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryMscpRiskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApdidToken) {
		body["ApdidToken"] = request.ApdidToken
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.TerminalType) {
		body["TerminalType"] = request.TerminalType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMscpRiskInfo"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMscpRiskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPushAnalysisCoreIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushAnalysisCoreIndexResponse
func (client *Client) QueryPushAnalysisCoreIndexWithContext(ctx context.Context, request *QueryPushAnalysisCoreIndexRequest, runtime *dara.RuntimeOptions) (_result *QueryPushAnalysisCoreIndexResponse, _err error) {
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

	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushAnalysisCoreIndex"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushAnalysisCoreIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPushAnalysisTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushAnalysisTaskDetailResponse
func (client *Client) QueryPushAnalysisTaskDetailWithContext(ctx context.Context, request *QueryPushAnalysisTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryPushAnalysisTaskDetailResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushAnalysisTaskDetail"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushAnalysisTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPushAnalysisTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushAnalysisTaskListResponse
func (client *Client) QueryPushAnalysisTaskListWithContext(ctx context.Context, request *QueryPushAnalysisTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryPushAnalysisTaskListResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushAnalysisTaskList"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushAnalysisTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPushSchedulerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushSchedulerListResponse
func (client *Client) QueryPushSchedulerListWithContext(ctx context.Context, request *QueryPushSchedulerListRequest, runtime *dara.RuntimeOptions) (_result *QueryPushSchedulerListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.UniqueId) {
		body["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushSchedulerList"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushSchedulerListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RevokePushMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokePushMessageResponse
func (client *Client) RevokePushMessageWithContext(ctx context.Context, request *RevokePushMessageRequest, runtime *dara.RuntimeOptions) (_result *RevokePushMessageResponse, _err error) {
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

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.TargetId) {
		body["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokePushMessage"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokePushMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RevokePushTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokePushTaskResponse
func (client *Client) RevokePushTaskWithContext(ctx context.Context, request *RevokePushTaskRequest, runtime *dara.RuntimeOptions) (_result *RevokePushTaskResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokePushTask"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokePushTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RunMsaDiffRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMsaDiffResponse
func (client *Client) RunMsaDiffWithContext(ctx context.Context, request *RunMsaDiffRequest, runtime *dara.RuntimeOptions) (_result *RunMsaDiffResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMsaDiffRunJsonStr) {
		body["MpaasMappcenterMsaDiffRunJsonStr"] = request.MpaasMappcenterMsaDiffRunJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMsaDiff"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunMsaDiffResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveMgsApirestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveMgsApirestResponse
func (client *Client) SaveMgsApirestWithContext(ctx context.Context, request *SaveMgsApirestRequest, runtime *dara.RuntimeOptions) (_result *SaveMgsApirestResponse, _err error) {
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

	if !dara.IsNil(request.MpaasMappcenterMgsApirestSaveJsonStr) {
		body["MpaasMappcenterMgsApirestSaveJsonStr"] = request.MpaasMappcenterMgsApirestSaveJsonStr
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveMgsApirest"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveMgsApirestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartUserAppAsyncEnhanceInMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartUserAppAsyncEnhanceInMsaResponse
func (client *Client) StartUserAppAsyncEnhanceInMsaWithContext(ctx context.Context, request *StartUserAppAsyncEnhanceInMsaRequest, runtime *dara.RuntimeOptions) (_result *StartUserAppAsyncEnhanceInMsaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewShieldConfig) {
		query["NewShieldConfig"] = request.NewShieldConfig
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApkProtector) {
		body["ApkProtector"] = request.ApkProtector
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AssetsFileList) {
		body["AssetsFileList"] = request.AssetsFileList
	}

	if !dara.IsNil(request.Classes) {
		body["Classes"] = request.Classes
	}

	if !dara.IsNil(request.DalvikDebugger) {
		body["DalvikDebugger"] = request.DalvikDebugger
	}

	if !dara.IsNil(request.EmulatorEnvironment) {
		body["EmulatorEnvironment"] = request.EmulatorEnvironment
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JavaHook) {
		body["JavaHook"] = request.JavaHook
	}

	if !dara.IsNil(request.MemoryDump) {
		body["MemoryDump"] = request.MemoryDump
	}

	if !dara.IsNil(request.NativeDebugger) {
		body["NativeDebugger"] = request.NativeDebugger
	}

	if !dara.IsNil(request.NativeHook) {
		body["NativeHook"] = request.NativeHook
	}

	if !dara.IsNil(request.PackageTampered) {
		body["PackageTampered"] = request.PackageTampered
	}

	if !dara.IsNil(request.Root) {
		body["Root"] = request.Root
	}

	if !dara.IsNil(request.RunMode) {
		body["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.SoFileList) {
		body["SoFileList"] = request.SoFileList
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.TotalSwitch) {
		body["TotalSwitch"] = request.TotalSwitch
	}

	if !dara.IsNil(request.UseAShield) {
		body["UseAShield"] = request.UseAShield
	}

	if !dara.IsNil(request.UseYShield) {
		body["UseYShield"] = request.UseYShield
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartUserAppAsyncEnhanceInMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartUserAppAsyncEnhanceInMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新短链
//
// @param request - UpdateLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLinkResponse
func (client *Client) UpdateLinkWithContext(ctx context.Context, request *UpdateLinkRequest, runtime *dara.RuntimeOptions) (_result *UpdateLinkResponse, _err error) {
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

	if !dara.IsNil(request.Cors) {
		body["Cors"] = request.Cors
	}

	if !dara.IsNil(request.Domain) {
		body["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Dynamicfield) {
		body["Dynamicfield"] = request.Dynamicfield
	}

	if !dara.IsNil(request.TargetUrl) {
		body["TargetUrl"] = request.TargetUrl
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLink"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新热修复发布任务状态
//
// @param request - UpdateMcubeHotpatchTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcubeHotpatchTaskStatusResponse
func (client *Client) UpdateMcubeHotpatchTaskStatusWithContext(ctx context.Context, request *UpdateMcubeHotpatchTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcubeHotpatchTaskStatusResponse, _err error) {
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

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PackageId) {
		body["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcubeHotpatchTaskStatus"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcubeHotpatchTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMcubeWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcubeWhitelistResponse
func (client *Client) UpdateMcubeWhitelistWithContext(ctx context.Context, request *UpdateMcubeWhitelistRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcubeWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.KeyIds) {
		body["KeyIds"] = request.KeyIds
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.OssUrl) {
		body["OssUrl"] = request.OssUrl
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcubeWhitelist"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcubeWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMdsCubeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMdsCubeResourceResponse
func (client *Client) UpdateMdsCubeResourceWithContext(ctx context.Context, request *UpdateMdsCubeResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateMdsCubeResourceResponse, _err error) {
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

	if !dara.IsNil(request.MockDataUrl) {
		body["MockDataUrl"] = request.MockDataUrl
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.TemplateResourceId) {
		body["TemplateResourceId"] = request.TemplateResourceId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMdsCubeResource"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMdsCubeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMpaasAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMpaasAppInfoResponse
func (client *Client) UpdateMpaasAppInfoWithContext(ctx context.Context, request *UpdateMpaasAppInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateMpaasAppInfoResponse, _err error) {
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

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.IconFileUrl) {
		body["IconFileUrl"] = request.IconFileUrl
	}

	if !dara.IsNil(request.Identifier) {
		body["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMpaasAppInfo"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMpaasAppInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传字节码到msa进行加固
//
// @param request - UploadBitcodeToMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadBitcodeToMsaResponse
func (client *Client) UploadBitcodeToMsaWithContext(ctx context.Context, request *UploadBitcodeToMsaRequest, runtime *dara.RuntimeOptions) (_result *UploadBitcodeToMsaResponse, _err error) {
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

	if !dara.IsNil(request.Bitcode) {
		body["Bitcode"] = request.Bitcode
	}

	if !dara.IsNil(request.CodeVersion) {
		body["CodeVersion"] = request.CodeVersion
	}

	if !dara.IsNil(request.License) {
		body["License"] = request.License
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadBitcodeToMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadBitcodeToMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadMcubeMiniPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMcubeMiniPackageResponse
func (client *Client) UploadMcubeMiniPackageWithContext(ctx context.Context, request *UploadMcubeMiniPackageRequest, runtime *dara.RuntimeOptions) (_result *UploadMcubeMiniPackageResponse, _err error) {
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

	if !dara.IsNil(request.AutoInstall) {
		body["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.ClientVersionMax) {
		body["ClientVersionMax"] = request.ClientVersionMax
	}

	if !dara.IsNil(request.ClientVersionMin) {
		body["ClientVersionMin"] = request.ClientVersionMin
	}

	if !dara.IsNil(request.EnableKeepAlive) {
		body["EnableKeepAlive"] = request.EnableKeepAlive
	}

	if !dara.IsNil(request.EnableOptionMenu) {
		body["EnableOptionMenu"] = request.EnableOptionMenu
	}

	if !dara.IsNil(request.EnableTabBar) {
		body["EnableTabBar"] = request.EnableTabBar
	}

	if !dara.IsNil(request.ExtendInfo) {
		body["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.H5Id) {
		body["H5Id"] = request.H5Id
	}

	if !dara.IsNil(request.H5Name) {
		body["H5Name"] = request.H5Name
	}

	if !dara.IsNil(request.H5Version) {
		body["H5Version"] = request.H5Version
	}

	if !dara.IsNil(request.IconFileUrl) {
		body["IconFileUrl"] = request.IconFileUrl
	}

	if !dara.IsNil(request.IconUrl) {
		body["IconUrl"] = request.IconUrl
	}

	if !dara.IsNil(request.InstallType) {
		body["InstallType"] = request.InstallType
	}

	if !dara.IsNil(request.MainUrl) {
		body["MainUrl"] = request.MainUrl
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.PackageType) {
		body["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceFileUrl) {
		body["ResourceFileUrl"] = request.ResourceFileUrl
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.Vhost) {
		body["Vhost"] = request.Vhost
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadMcubeMiniPackage"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadMcubeMiniPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadMcubeRsaKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMcubeRsaKeyResponse
func (client *Client) UploadMcubeRsaKeyWithContext(ctx context.Context, request *UploadMcubeRsaKeyRequest, runtime *dara.RuntimeOptions) (_result *UploadMcubeRsaKeyResponse, _err error) {
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

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OnexFlag) {
		body["OnexFlag"] = request.OnexFlag
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadMcubeRsaKey"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadMcubeRsaKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadUserAppToMsaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadUserAppToMsaResponse
func (client *Client) UploadUserAppToMsaWithContext(ctx context.Context, request *UploadUserAppToMsaRequest, runtime *dara.RuntimeOptions) (_result *UploadUserAppToMsaResponse, _err error) {
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

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.UseYShield) {
		body["UseYShield"] = request.UseYShield
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadUserAppToMsa"),
		Version:     dara.String("2020-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadUserAppToMsaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
