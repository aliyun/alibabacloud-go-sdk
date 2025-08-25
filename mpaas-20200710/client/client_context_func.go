// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 取消定时任务
//
// @param request - CancelMpsSchedulerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelMpsSchedulerResponse
func (client *Client) CancelMpsSchedulerWithContext(ctx context.Context, request *CancelMpsSchedulerRequest, runtime *dara.RuntimeOptions) (_result *CancelMpsSchedulerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
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
		Action:      dara.String("CancelMpsScheduler"),
		Version:     dara.String("2020-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelMpsSchedulerResponse{}
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - CreateMcubeMiniAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcubeMiniAppResponse
func (client *Client) CreateMcubeMiniAppWithContext(ctx context.Context, request *CreateMcubeMiniAppRequest, runtime *dara.RuntimeOptions) (_result *CreateMcubeMiniAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - CreateOpenGlobalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpenGlobalDataResponse
func (client *Client) CreateOpenGlobalDataWithContext(ctx context.Context, request *CreateOpenGlobalDataRequest, runtime *dara.RuntimeOptions) (_result *CreateOpenGlobalDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - DeleteMcubeMiniAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcubeMiniAppResponse
func (client *Client) DeleteMcubeMiniAppWithContext(ctx context.Context, request *DeleteMcubeMiniAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcubeMiniAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - DeleteMdsWhitelistContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMdsWhitelistContentResponse
func (client *Client) DeleteMdsWhitelistContentWithContext(ctx context.Context, request *DeleteMdsWhitelistContentRequest, runtime *dara.RuntimeOptions) (_result *DeleteMdsWhitelistContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - GetMcubeFileTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcubeFileTokenResponse
func (client *Client) GetMcubeFileTokenWithContext(ctx context.Context, request *GetMcubeFileTokenRequest, runtime *dara.RuntimeOptions) (_result *GetMcubeFileTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - ListMcubeMiniAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcubeMiniAppsResponse
func (client *Client) ListMcubeMiniAppsWithContext(ctx context.Context, request *ListMcubeMiniAppsRequest, runtime *dara.RuntimeOptions) (_result *ListMcubeMiniAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param tmpReq - PushBroadcastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushBroadcastResponse
func (client *Client) PushBroadcastWithContext(ctx context.Context, tmpReq *PushBroadcastRequest, runtime *dara.RuntimeOptions) (_result *PushBroadcastResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param tmpReq - PushSimpleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushSimpleResponse
func (client *Client) PushSimpleWithContext(ctx context.Context, tmpReq *PushSimpleRequest, runtime *dara.RuntimeOptions) (_result *PushSimpleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - QueryMcubeMiniPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMcubeMiniPackageResponse
func (client *Client) QueryMcubeMiniPackageWithContext(ctx context.Context, request *QueryMcubeMiniPackageRequest, runtime *dara.RuntimeOptions) (_result *QueryMcubeMiniPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// Summary:
//
// 查询定时任务列表
//
// @param request - QueryMpsSchedulerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMpsSchedulerListResponse
func (client *Client) QueryMpsSchedulerListWithContext(ctx context.Context, request *QueryMpsSchedulerListRequest, runtime *dara.RuntimeOptions) (_result *QueryMpsSchedulerListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - QueryPushAnalysisCoreIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushAnalysisCoreIndexResponse
func (client *Client) QueryPushAnalysisCoreIndexWithContext(ctx context.Context, request *QueryPushAnalysisCoreIndexRequest, runtime *dara.RuntimeOptions) (_result *QueryPushAnalysisCoreIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - UpdateMcubeWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcubeWhitelistResponse
func (client *Client) UpdateMcubeWhitelistWithContext(ctx context.Context, request *UpdateMcubeWhitelistRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcubeWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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

// @param request - UploadMcubeMiniPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMcubeMiniPackageResponse
func (client *Client) UploadMcubeMiniPackageWithContext(ctx context.Context, request *UploadMcubeMiniPackageRequest, runtime *dara.RuntimeOptions) (_result *UploadMcubeMiniPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2020-07-10"),
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
