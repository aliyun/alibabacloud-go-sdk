// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 新增共享账号信息
//
// @param request - AddSharedAccountsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSharedAccountsResponse
func (client *Client) AddSharedAccountsWithContext(ctx context.Context, request *AddSharedAccountsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddSharedAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		body["accountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.ResourceId) {
		body["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSharedAccounts"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/sharedAccounts"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSharedAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将参数集关联资源
//
// @param request - AssociateDetectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateDetectConfigResponse
func (client *Client) AssociateDetectConfigWithContext(ctx context.Context, request *AssociateDetectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AssociateDetectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DetectConfigId) {
		body["detectConfigId"] = request.DetectConfigId
	}

	if !dara.IsNil(request.TargetId) {
		body["targetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		body["targetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateDetectConfig"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig/operations/associate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateDetectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分组关联
//
// @param request - AssociateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateGroupResponse
func (client *Client) AssociateGroupWithContext(ctx context.Context, groupId *string, request *AssociateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AssociateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceIds) {
		body["resourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group/" + dara.PercentEncode(dara.StringValue(groupId)) + "/associate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将参数集关联资源
//
// @param request - AssociateParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateParameterSetResponse
func (client *Client) AssociateParameterSetWithContext(ctx context.Context, request *AssociateParameterSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AssociateParameterSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ParameterSetIds) {
		body["parameterSetIds"] = request.ParameterSetIds
	}

	if !dara.IsNil(request.ResourceId) {
		body["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateParameterSet"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets/operations/associate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateParameterSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消资源导出任务
//
// @param request - CancelResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelResourceExportTaskResponse
func (client *Client) CancelResourceExportTaskWithContext(ctx context.Context, exportTaskId *string, request *CancelResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelResourceExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelResourceExportTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks/cancel/" + dara.PercentEncode(dara.StringValue(exportTaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelResourceExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建偏差检测配置
//
// @param request - CreateDetectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDetectConfigResponse
func (client *Client) CreateDetectConfigWithContext(ctx context.Context, request *CreateDetectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDetectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmConfigs) {
		body["alarmConfigs"] = request.AlarmConfigs
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CronExpression) {
		body["cronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DetectConfigName) {
		body["detectConfigName"] = request.DetectConfigName
	}

	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.TriggerType) {
		body["triggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDetectConfig"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDetectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建分组
//
// @param request - CreateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoDestroy) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !dara.IsNil(request.AutoTrigger) {
		body["autoTrigger"] = request.AutoTrigger
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ForcedSetting) {
		body["forcedSetting"] = request.ForcedSetting
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NotifyConfig) {
		body["notifyConfig"] = request.NotifyConfig
	}

	if !dara.IsNil(request.NotifyOperationTypes) {
		body["notifyOperationTypes"] = request.NotifyOperationTypes
	}

	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.ReportExportField) {
		body["reportExportField"] = request.ReportExportField
	}

	if !dara.IsNil(request.ReportExportPath) {
		body["reportExportPath"] = request.ReportExportPath
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !dara.IsNil(request.TriggerConfig) {
		body["triggerConfig"] = request.TriggerConfig
	}

	if !dara.IsNil(request.TriggerResourceType) {
		body["triggerResourceType"] = request.TriggerResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建作业
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithContext(ctx context.Context, taskId *string, request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.SubCommand) {
		body["subCommand"] = request.SubCommand
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Module
//
// @param request - CreateModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModuleResponse
func (client *Client) CreateModuleWithContext(ctx context.Context, request *CreateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GroupInfo) {
		body["groupInfo"] = request.GroupInfo
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.SourcePath) {
		body["sourcePath"] = request.SourcePath
	}

	if !dara.IsNil(request.StatePath) {
		body["statePath"] = request.StatePath
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.VersionStrategy) {
		body["versionStrategy"] = request.VersionStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publish a template version.
//
// @param request - CreateModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModuleVersionResponse
func (client *Client) CreateModuleVersionWithContext(ctx context.Context, moduleId *string, request *CreateModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModuleVersion"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/" + dara.PercentEncode(dara.StringValue(moduleId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModuleVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建参数集
//
// @param request - CreateParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterSetResponse
func (client *Client) CreateParameterSetWithContext(ctx context.Context, request *CreateParameterSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateParameterSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParameterSet"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParameterSetResponse{}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/project"),
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
// 创建RegistryModule
//
// @param request - CreateRegistryModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRegistryModuleResponse
func (client *Client) CreateRegistryModuleWithContext(ctx context.Context, request *CreateRegistryModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRegistryModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Acl) {
		body["acl"] = request.Acl
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ModuleName) {
		body["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.NamespaceName) {
		body["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.Provider) {
		body["provider"] = request.Provider
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRegistryModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRegistryModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateRegistryNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRegistryNamespaceResponse
func (client *Client) CreateRegistryNamespaceWithContext(ctx context.Context, request *CreateRegistryNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRegistryNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Acl) {
		body["acl"] = request.Acl
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Maintainer) {
		body["maintainer"] = request.Maintainer
	}

	if !dara.IsNil(request.NamespaceName) {
		body["namespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRegistryNamespace"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryNamespace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRegistryNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建导出任务
//
// @param request - CreateResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceExportTaskResponse
func (client *Client) CreateResourceExportTaskWithContext(ctx context.Context, request *CreateResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ExportToModule) {
		body["exportToModule"] = request.ExportToModule
	}

	if !dara.IsNil(request.IncludeRules) {
		body["includeRules"] = request.IncludeRules
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !dara.IsNil(request.TerraformVersion) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !dara.IsNil(request.TriggerStrategy) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceExportTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源栈
//
// @param request - CreateStackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStackResponse
func (client *Client) CreateStackWithContext(ctx context.Context, request *CreateStackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.SourcePath) {
		body["sourcePath"] = request.SourcePath
	}

	if !dara.IsNil(request.WorkingDirectory) {
		body["workingDirectory"] = request.WorkingDirectory
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStack"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStackResponse{}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoApply) {
		body["autoApply"] = request.AutoApply
	}

	if !dara.IsNil(request.AutoDestroy) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GroupInfo) {
		body["groupInfo"] = request.GroupInfo
	}

	if !dara.IsNil(request.InitModuleState) {
		body["initModuleState"] = request.InitModuleState
	}

	if !dara.IsNil(request.ModuleId) {
		body["moduleId"] = request.ModuleId
	}

	if !dara.IsNil(request.ModuleVersion) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ProtectionStrategy) {
		body["protectionStrategy"] = request.ProtectionStrategy
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.SkipPropertyValidation) {
		body["skipPropertyValidation"] = request.SkipPropertyValidation
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TaskBackend) {
		body["taskBackend"] = request.TaskBackend
	}

	if !dara.IsNil(request.TerraformVersion) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !dara.IsNil(request.TriggerStrategy) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks"),
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
// 删除偏差检测配置
//
// @param request - DeleteDetectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDetectConfigResponse
func (client *Client) DeleteDetectConfigWithContext(ctx context.Context, detectConfigId *string, request *DeleteDetectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDetectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDetectConfig"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig/" + dara.PercentEncode(dara.StringValue(detectConfigId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDetectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @param request - DeleteGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithContext(ctx context.Context, groupId *string, request *DeleteGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模板
//
// @param request - DeleteModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModuleResponse
func (client *Client) DeleteModuleWithContext(ctx context.Context, moduleId *string, request *DeleteModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/" + dara.PercentEncode(dara.StringValue(moduleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除参数集
//
// @param request - DeleteParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterSetResponse
func (client *Client) DeleteParameterSetWithContext(ctx context.Context, parameterSetId *string, request *DeleteParameterSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteParameterSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParameterSet"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets/" + dara.PercentEncode(dara.StringValue(parameterSetId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParameterSetResponse{}
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
func (client *Client) DeleteProjectWithContext(ctx context.Context, projectId *string, request *DeleteProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/project/" + dara.PercentEncode(dara.StringValue(projectId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 删除RegistryModule
//
// @param request - DeleteRegistryModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistryModuleResponse
func (client *Client) DeleteRegistryModuleWithContext(ctx context.Context, namespaceName *string, moduleName *string, request *DeleteRegistryModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRegistryModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegistryModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModule/" + dara.PercentEncode(dara.StringValue(namespaceName)) + "/" + dara.PercentEncode(dara.StringValue(moduleName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegistryModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除RegistryModule版本
//
// @param request - DeleteRegistryModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistryModuleVersionResponse
func (client *Client) DeleteRegistryModuleVersionWithContext(ctx context.Context, namespaceName *string, moduleName *string, version *string, request *DeleteRegistryModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRegistryModuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegistryModuleVersion"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModuleVersion/" + dara.PercentEncode(dara.StringValue(namespaceName)) + "/" + dara.PercentEncode(dara.StringValue(moduleName)) + "/" + dara.PercentEncode(dara.StringValue(version))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegistryModuleVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @param request - DeleteRegistryNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistryNamespaceResponse
func (client *Client) DeleteRegistryNamespaceWithContext(ctx context.Context, namespaceName *string, request *DeleteRegistryNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRegistryNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegistryNamespace"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryNamespace/" + dara.PercentEncode(dara.StringValue(namespaceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegistryNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源导出任务
//
// @param request - DeleteResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceExportTaskResponse
func (client *Client) DeleteResourceExportTaskWithContext(ctx context.Context, exportTaskId *string, request *DeleteResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceExportTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks/" + dara.PercentEncode(dara.StringValue(exportTaskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源栈
//
// @param request - DeleteStackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStackResponse
func (client *Client) DeleteStackWithContext(ctx context.Context, stackId *string, request *DeleteStackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CleanResources) {
		query["cleanResources"] = request.CleanResources
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStack"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/" + dara.PercentEncode(dara.StringValue(stackId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param request - DeleteTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithContext(ctx context.Context, taskId *string, request *DeleteTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起状态文件一致性检测
//
// @param request - DetectTerraformStateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectTerraformStateResponse
func (client *Client) DetectTerraformStateWithContext(ctx context.Context, request *DetectTerraformStateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DetectTerraformStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Identifier) {
		body["identifier"] = request.Identifier
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectTerraformState"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detect"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectTerraformStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除参数集关联资源关系
//
// @param request - DissociateDetectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateDetectConfigResponse
func (client *Client) DissociateDetectConfigWithContext(ctx context.Context, request *DissociateDetectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DissociateDetectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DetectConfigId) {
		body["detectConfigId"] = request.DetectConfigId
	}

	if !dara.IsNil(request.TargetId) {
		body["targetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		body["targetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateDetectConfig"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig/operations/dissociate"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateDetectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消关联分组
//
// @param request - DissociateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateGroupResponse
func (client *Client) DissociateGroupWithContext(ctx context.Context, projectId *string, groupId *string, request *DissociateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DissociateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceIds) {
		body["resourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group/" + dara.PercentEncode(dara.StringValue(groupId)) + "/dissociate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除参数集关联资源关系
//
// @param request - DissociateParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateParameterSetResponse
func (client *Client) DissociateParameterSetWithContext(ctx context.Context, request *DissociateParameterSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DissociateParameterSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ParameterSetIds) {
		body["parameterSetIds"] = request.ParameterSetIds
	}

	if !dara.IsNil(request.ResourceId) {
		body["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateParameterSet"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets/operations/dissociate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateParameterSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行RegistryModule
//
// @param request - ExecuteRegistryModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteRegistryModuleResponse
func (client *Client) ExecuteRegistryModuleWithContext(ctx context.Context, namespaceName *string, moduleName *string, request *ExecuteRegistryModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteRegistryModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteRegistryModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModule/" + dara.PercentEncode(dara.StringValue(namespaceName)) + "/" + dara.PercentEncode(dara.StringValue(moduleName)) + "/execution"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteRegistryModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行资源导出任务
//
// @param request - ExecuteResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteResourceExportTaskResponse
func (client *Client) ExecuteResourceExportTaskWithContext(ctx context.Context, exportTaskId *string, request *ExecuteResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteResourceExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteResourceExportTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks/execute/" + dara.PercentEncode(dara.StringValue(exportTaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteResourceExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行TerraformApply
//
// @param request - ExecuteTerraformApplyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTerraformApplyResponse
func (client *Client) ExecuteTerraformApplyWithContext(ctx context.Context, request *ExecuteTerraformApplyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTerraformApplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.StateId) {
		body["stateId"] = request.StateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTerraformApply"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraform/execution/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTerraformApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行TerraformDestroy
//
// @param request - ExecuteTerraformDestroyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTerraformDestroyResponse
func (client *Client) ExecuteTerraformDestroyWithContext(ctx context.Context, request *ExecuteTerraformDestroyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTerraformDestroyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.StateId) {
		body["stateId"] = request.StateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTerraformDestroy"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraform/execution/destroy"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTerraformDestroyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行TerraformPlan
//
// @param request - ExecuteTerraformPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTerraformPlanResponse
func (client *Client) ExecuteTerraformPlanWithContext(ctx context.Context, request *ExecuteTerraformPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTerraformPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.StateId) {
		body["stateId"] = request.StateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTerraformPlan"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraform/execution/plan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTerraformPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成模板
//
// @param request - GenerateModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateModuleResponse
func (client *Client) GenerateModuleWithContext(ctx context.Context, request *GenerateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GenerateSource) {
		body["generateSource"] = request.GenerateSource
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		body["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.Syntax) {
		body["syntax"] = request.Syntax
	}

	if !dara.IsNil(request.Template) {
		body["template"] = request.Template
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !dara.IsNil(request.TerraformResourceType) {
		body["terraformResourceType"] = request.TerraformResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/explorer/generate/module"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateModuleResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 偏差检测配置详情
//
// @param request - GetDetectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectConfigResponse
func (client *Client) GetDetectConfigWithContext(ctx context.Context, detectConfigId *string, request *GetDetectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDetectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetectConfig"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig/" + dara.PercentEncode(dara.StringValue(detectConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Terraform运行结果
//
// @param request - GetExecuteStateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteStateResponse
func (client *Client) GetExecuteStateWithContext(ctx context.Context, stateId *string, request *GetExecuteStateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExecuteStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExecuteState"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraform/execution/" + dara.PercentEncode(dara.StringValue(stateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecuteStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分组
//
// @param request - GetGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithContext(ctx context.Context, groupId *string, request *GetGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作业详情
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithContext(ctx context.Context, taskId *string, jobId *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskType) {
		query["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Module Details
//
// @param request - GetModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModuleResponse
func (client *Client) GetModuleWithContext(ctx context.Context, moduleId *string, request *GetModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/" + dara.PercentEncode(dara.StringValue(moduleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模板版本详情
//
// @param request - GetModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModuleVersionResponse
func (client *Client) GetModuleVersionWithContext(ctx context.Context, moduleId *string, moduleVersion *string, request *GetModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModuleVersion"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/" + dara.PercentEncode(dara.StringValue(moduleId)) + "/versions/" + dara.PercentEncode(dara.StringValue(moduleVersion))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModuleVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 参数集详情
//
// @param request - GetParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParameterSetResponse
func (client *Client) GetParameterSetWithContext(ctx context.Context, parameterSetId *string, request *GetParameterSetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetParameterSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParameterSet"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets/" + dara.PercentEncode(dara.StringValue(parameterSetId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParameterSetResponse{}
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
// @param request - GetProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, projectId *string, request *GetProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/project/" + dara.PercentEncode(dara.StringValue(projectId))),
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
// 获取RegistryModule信息
//
// @param request - GetRegistryModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistryModuleResponse
func (client *Client) GetRegistryModuleWithContext(ctx context.Context, namespaceName *string, moduleName *string, request *GetRegistryModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegistryModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegistryModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModule/" + dara.PercentEncode(dara.StringValue(namespaceName)) + "/" + dara.PercentEncode(dara.StringValue(moduleName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegistryModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RegistryModule版本信息
//
// @param request - GetRegistryModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistryModuleVersionResponse
func (client *Client) GetRegistryModuleVersionWithContext(ctx context.Context, namespaceName *string, moduleName *string, version *string, request *GetRegistryModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegistryModuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegistryModuleVersion"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModuleVersion/" + dara.PercentEncode(dara.StringValue(namespaceName)) + "/" + dara.PercentEncode(dara.StringValue(moduleName)) + "/" + dara.PercentEncode(dara.StringValue(version))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegistryModuleVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作空间信息
//
// @param request - GetRegistryNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistryNamespaceResponse
func (client *Client) GetRegistryNamespaceWithContext(ctx context.Context, namespaceName *string, request *GetRegistryNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegistryNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegistryNamespace"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryNamespace/" + dara.PercentEncode(dara.StringValue(namespaceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegistryNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询导出任务详情
//
// @param request - GetResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceExportTaskResponse
func (client *Client) GetResourceExportTaskWithContext(ctx context.Context, exportTaskId *string, request *GetResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportVersion) {
		query["exportVersion"] = request.ExportVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceExportTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks/" + dara.PercentEncode(dara.StringValue(exportTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源类型信息
//
// @param request - GetResourceTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceTypeWithContext(ctx context.Context, resourceType *string, request *GetResourceTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FilterReadOnly) {
		query["filterReadOnly"] = request.FilterReadOnly
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		query["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceType"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourceType/" + dara.PercentEncode(dara.StringValue(resourceType))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源栈
//
// @param request - GetStackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackResponse
func (client *Client) GetStackWithContext(ctx context.Context, stackId *string, request *GetStackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStack"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/" + dara.PercentEncode(dara.StringValue(stackId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 部署详情接口
//
// @param request - GetStackDeploymentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackDeploymentsResponse
func (client *Client) GetStackDeploymentsWithContext(ctx context.Context, stackId *string, request *GetStackDeploymentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStackDeploymentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigVersion) {
		query["configVersion"] = request.ConfigVersion
	}

	if !dara.IsNil(request.DeploymentName) {
		query["deploymentName"] = request.DeploymentName
	}

	if !dara.IsNil(request.DeploymentNo) {
		query["deploymentNo"] = request.DeploymentNo
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackDeployments"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/" + dara.PercentEncode(dara.StringValue(stackId)) + "/deployments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackDeploymentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源栈部署结果
//
// @param request - GetStackExecutionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackExecutionResultResponse
func (client *Client) GetStackExecutionResultWithContext(ctx context.Context, triggerId *string, request *GetStackExecutionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStackExecutionResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackExecutionResult"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/trigger/" + dara.PercentEncode(dara.StringValue(triggerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackExecutionResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务详情
//
// @param request - GetTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, taskId *string, request *GetTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
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
// 获取状态文件检测结果
//
// @param request - GetTerraformStateDetectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTerraformStateDetectionResponse
func (client *Client) GetTerraformStateDetectionWithContext(ctx context.Context, detectionId *string, request *GetTerraformStateDetectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTerraformStateDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTerraformStateDetection"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detect/" + dara.PercentEncode(dara.StringValue(detectionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTerraformStateDetectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关联到资源的偏差检测配置列表
//
// @param request - ListDetectConfigRelationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDetectConfigRelationsResponse
func (client *Client) ListDetectConfigRelationsWithContext(ctx context.Context, request *ListDetectConfigRelationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDetectConfigRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DetectConfigId) {
		query["detectConfigId"] = request.DetectConfigId
	}

	if !dara.IsNil(request.TargetId) {
		query["targetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["targetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDetectConfigRelations"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig/operations/relation"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDetectConfigRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 偏差检测配置列表
//
// @param request - ListDetectConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDetectConfigsResponse
func (client *Client) ListDetectConfigsWithContext(ctx context.Context, request *ListDetectConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDetectConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DetectConfigName) {
		query["detectConfigName"] = request.DetectConfigName
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDetectConfigs"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDetectConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Explorer的egistryModule版本示例列表
//
// @param request - ListExplorerRegistryModuleExamplesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExplorerRegistryModuleExamplesResponse
func (client *Client) ListExplorerRegistryModuleExamplesWithContext(ctx context.Context, request *ListExplorerRegistryModuleExamplesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExplorerRegistryModuleExamplesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExampleName) {
		query["exampleName"] = request.ExampleName
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModuleName) {
		query["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.ModuleVersion) {
		query["moduleVersion"] = request.ModuleVersion
	}

	if !dara.IsNil(request.NamespaceName) {
		query["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExplorerRegistryModuleExamples"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/explorerRegistryModule/example"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExplorerRegistryModuleExamplesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Explorer的egistryModule版本列表
//
// @param request - ListExplorerRegistryModuleVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExplorerRegistryModuleVersionsResponse
func (client *Client) ListExplorerRegistryModuleVersionsWithContext(ctx context.Context, request *ListExplorerRegistryModuleVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExplorerRegistryModuleVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModuleName) {
		query["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.ModuleVersion) {
		query["moduleVersion"] = request.ModuleVersion
	}

	if !dara.IsNil(request.NamespaceName) {
		query["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExplorerRegistryModuleVersions"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/explorerRegistryModule/version"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExplorerRegistryModuleVersionsResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Explorer的Registry Module列表
//
// @param request - ListExplorerRegistryModulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExplorerRegistryModulesResponse
func (client *Client) ListExplorerRegistryModulesWithContext(ctx context.Context, request *ListExplorerRegistryModulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExplorerRegistryModulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModuleName) {
		query["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExplorerRegistryModules"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/explorerRegistryModule"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExplorerRegistryModulesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分组列表
//
// @param tmpReq - ListGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupResponse
func (client *Client) ListGroupWithContext(ctx context.Context, tmpReq *ListGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作业列表
//
// @param request - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, taskId *string, request *ListJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobType) {
		query["jobType"] = request.JobType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TaskType) {
		query["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模板版本列表
//
// @param request - ListModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModuleVersionResponse
func (client *Client) ListModuleVersionWithContext(ctx context.Context, moduleId *string, request *ListModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
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
		Action:      dara.String("ListModuleVersion"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/" + dara.PercentEncode(dara.StringValue(moduleId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModuleVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举模板
//
// @param tmpReq - ListModulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModulesResponse
func (client *Client) ListModulesWithContext(ctx context.Context, tmpReq *ListModulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListModulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ModuleName) {
		query["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModules"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关联到资源的参数集列表
//
// @param request - ListParameterSetRelationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterSetRelationResponse
func (client *Client) ListParameterSetRelationWithContext(ctx context.Context, request *ListParameterSetRelationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListParameterSetRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParameterSetRelation"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets/operations/relation"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParameterSetRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 参数集列表
//
// @param request - ListParameterSetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterSetsResponse
func (client *Client) ListParameterSetsWithContext(ctx context.Context, request *ListParameterSetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListParameterSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
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
		Action:      dara.String("ListParameterSets"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParameterSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 所有产品列表
//
// @param request - ListProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithContext(ctx context.Context, request *ListProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.SupportTerraformer) {
		query["supportTerraformer"] = request.SupportTerraformer
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		query["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目列表
//
// @param tmpReq - ListProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithContext(ctx context.Context, tmpReq *ListProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProject"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/project"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 获取RegistryModule版本列表
//
// @param request - ListRegistryModuleVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistryModuleVersionsResponse
func (client *Client) ListRegistryModuleVersionsWithContext(ctx context.Context, request *ListRegistryModuleVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegistryModuleVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModuleName) {
		query["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.NamespaceName) {
		query["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistryModuleVersions"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModuleVersion"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistryModuleVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RegistryModule列表
//
// @param request - ListRegistryModulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistryModulesResponse
func (client *Client) ListRegistryModulesWithContext(ctx context.Context, request *ListRegistryModulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegistryModulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NamespaceName) {
		query["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistryModules"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistryModulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作空间列表
//
// @param request - ListRegistryNamespacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistryNamespacesResponse
func (client *Client) ListRegistryNamespacesWithContext(ctx context.Context, request *ListRegistryNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegistryNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistryNamespaces"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryNamespace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistryNamespacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务版本列表
//
// @param request - ListResourceExportTaskVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceExportTaskVersionsResponse
func (client *Client) ListResourceExportTaskVersionsWithContext(ctx context.Context, exportTaskId *string, request *ListResourceExportTaskVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceExportTaskVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportVersion) {
		query["exportVersion"] = request.ExportVersion
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceExportTaskVersions"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks/" + dara.PercentEncode(dara.StringValue(exportTaskId)) + "/exportVersions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceExportTaskVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询导出任务列表
//
// @param request - ListResourceExportTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceExportTasksResponse
func (client *Client) ListResourceExportTasksWithContext(ctx context.Context, request *ListResourceExportTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceExportTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportTaskId) {
		query["exportTaskId"] = request.ExportTaskId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
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
		Action:      dara.String("ListResourceExportTasks"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceExportTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源类型列表
//
// @param tmpReq - ListResourceTypesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithContext(ctx context.Context, tmpReq *ListResourceTypesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourceTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TerraformResourceTypes) {
		request.TerraformResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TerraformResourceTypes, dara.String("terraformResourceTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Product) {
		query["product"] = request.Product
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Subcategory) {
		query["subcategory"] = request.Subcategory
	}

	if !dara.IsNil(request.SupportTerraformer) {
		query["supportTerraformer"] = request.SupportTerraformer
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		query["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !dara.IsNil(request.TerraformResourceTypesShrink) {
		query["terraformResourceTypes"] = request.TerraformResourceTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypes"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourceTypes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源列表
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SourceValue) {
		query["sourceValue"] = request.SourceValue
	}

	if !dara.IsNil(request.SpecType) {
		query["specType"] = request.SpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResources"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resources/stateparser"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源栈配置列表
//
// @param request - ListStackConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackConfigsResponse
func (client *Client) ListStackConfigsWithContext(ctx context.Context, stackId *string, request *ListStackConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStackConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackConfigs"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/" + dara.PercentEncode(dara.StringValue(stackId)) + "/configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举资源栈
//
// @param request - ListStacksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStacksResponse
func (client *Client) ListStacksWithContext(ctx context.Context, request *ListStacksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStacks"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStacksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务列表
//
// @param tmpReq - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, tmpReq *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ModuleId) {
		query["moduleId"] = request.ModuleId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
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
// terraformProvider版本
//
// @param request - ListTerraformProviderVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerraformProviderVersionsResponse
func (client *Client) ListTerraformProviderVersionsWithContext(ctx context.Context, request *ListTerraformProviderVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTerraformProviderVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Usage) {
		query["usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTerraformProviderVersions"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/version/terraform/provider"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTerraformProviderVersionsResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 支持状态文件的资源导入和移除
//
// @param request - ManageTerraformStateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageTerraformStateResponse
func (client *Client) ManageTerraformStateWithContext(ctx context.Context, request *ManageTerraformStateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ManageTerraformStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Identifier) {
		body["identifier"] = request.Identifier
	}

	if !dara.IsNil(request.ImportResourceId) {
		body["importResourceId"] = request.ImportResourceId
	}

	if !dara.IsNil(request.ResourceIdentifier) {
		body["resourceIdentifier"] = request.ResourceIdentifier
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManageTerraformState"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/manage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ManageTerraformStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制作业
//
// @param request - OperateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateJobResponse
func (client *Client) OperateJobWithContext(ctx context.Context, taskId *string, jobId *string, operationType *string, request *OperateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OperateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["comment"] = request.Comment
	}

	if !dara.IsNil(request.TaskType) {
		query["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateJob"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/operation/" + dara.PercentEncode(dara.StringValue(operationType))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布RegistryModule版本
//
// @param request - PublishRegistryModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRegistryModuleVersionResponse
func (client *Client) PublishRegistryModuleVersionWithContext(ctx context.Context, request *PublishRegistryModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishRegistryModuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ModuleName) {
		body["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.NamespaceName) {
		body["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishRegistryModuleVersion"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModuleVersion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRegistryModuleVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除共享账号信息
//
// @param tmpReq - RemoveSharedAccountsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSharedAccountsResponse
func (client *Client) RemoveSharedAccountsWithContext(ctx context.Context, tmpReq *RemoveSharedAccountsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveSharedAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveSharedAccountsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("accountIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["accountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSharedAccounts"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/sharedAccounts"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSharedAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发资源栈部署
//
// @param request - TriggerStackExecutionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerStackExecutionResponse
func (client *Client) TriggerStackExecutionWithContext(ctx context.Context, request *TriggerStackExecutionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TriggerStackExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.ChangedFolders) {
		body["changedFolders"] = request.ChangedFolders
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CodePackagePath) {
		body["codePackagePath"] = request.CodePackagePath
	}

	if !dara.IsNil(request.CodeVersionId) {
		body["codeVersionId"] = request.CodeVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerStackExecution"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/trigger"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerStackExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新偏差检测配置
//
// @param request - UpdateDetectConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDetectConfigResponse
func (client *Client) UpdateDetectConfigWithContext(ctx context.Context, detectConfigId *string, request *UpdateDetectConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDetectConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmConfigs) {
		body["alarmConfigs"] = request.AlarmConfigs
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CronExpression) {
		body["cronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DetectConfigName) {
		body["detectConfigName"] = request.DetectConfigName
	}

	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.TriggerType) {
		body["triggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDetectConfig"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraformState/detectConfig/" + dara.PercentEncode(dara.StringValue(detectConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDetectConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ExplorerModule
//
// @param request - UpdateExplorerModuleAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExplorerModuleAttributeResponse
func (client *Client) UpdateExplorerModuleAttributeWithContext(ctx context.Context, explorerModuleId *string, request *UpdateExplorerModuleAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExplorerModuleAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExplorerModuleAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/explorerModule/" + dara.PercentEncode(dara.StringValue(explorerModuleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExplorerModuleAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改分组
//
// @param request - UpdateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithContext(ctx context.Context, groupId *string, request *UpdateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoDestroy) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !dara.IsNil(request.AutoTrigger) {
		body["autoTrigger"] = request.AutoTrigger
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ForcedSetting) {
		body["forcedSetting"] = request.ForcedSetting
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NotifyConfig) {
		body["notifyConfig"] = request.NotifyConfig
	}

	if !dara.IsNil(request.NotifyOperationTypes) {
		body["notifyOperationTypes"] = request.NotifyOperationTypes
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.ReportExportField) {
		body["reportExportField"] = request.ReportExportField
	}

	if !dara.IsNil(request.ReportExportPath) {
		body["reportExportPath"] = request.ReportExportPath
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !dara.IsNil(request.TriggerConfig) {
		body["triggerConfig"] = request.TriggerConfig
	}

	if !dara.IsNil(request.TriggerResourceType) {
		body["triggerResourceType"] = request.TriggerResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/group/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Module
//
// @param request - UpdateModuleAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModuleAttributeResponse
func (client *Client) UpdateModuleAttributeWithContext(ctx context.Context, moduleId *string, request *UpdateModuleAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModuleAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GroupInfo) {
		body["groupInfo"] = request.GroupInfo
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SourcePath) {
		body["sourcePath"] = request.SourcePath
	}

	if !dara.IsNil(request.StatePath) {
		body["statePath"] = request.StatePath
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.VersionStrategy) {
		body["versionStrategy"] = request.VersionStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModuleAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/" + dara.PercentEncode(dara.StringValue(moduleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModuleAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新参数集
//
// @param request - UpdateParameterSetAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParameterSetAttributeResponse
func (client *Client) UpdateParameterSetAttributeWithContext(ctx context.Context, parameterSetId *string, request *UpdateParameterSetAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateParameterSetAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateParameterSetAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/parameterSets/" + dara.PercentEncode(dara.StringValue(parameterSetId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateParameterSetAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改项目
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, projectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/project/" + dara.PercentEncode(dara.StringValue(projectId))),
		Method:      dara.String("POST"),
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
// 修改RegistryModule
//
// @param request - UpdateRegistryModuleAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRegistryModuleAttributeResponse
func (client *Client) UpdateRegistryModuleAttributeWithContext(ctx context.Context, namespaceName *string, moduleName *string, request *UpdateRegistryModuleAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRegistryModuleAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Acl) {
		body["acl"] = request.Acl
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRegistryModuleAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryModule/" + dara.PercentEncode(dara.StringValue(namespaceName)) + "/" + dara.PercentEncode(dara.StringValue(moduleName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRegistryModuleAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改工作空间
//
// @param request - UpdateRegistryNamespaceAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRegistryNamespaceAttributeResponse
func (client *Client) UpdateRegistryNamespaceAttributeWithContext(ctx context.Context, namespaceName *string, request *UpdateRegistryNamespaceAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRegistryNamespaceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Acl) {
		body["acl"] = request.Acl
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRegistryNamespaceAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/registryNamespace/" + dara.PercentEncode(dara.StringValue(namespaceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRegistryNamespaceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新导出任务
//
// @param request - UpdateResourceExportTaskAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceExportTaskAttributeResponse
func (client *Client) UpdateResourceExportTaskAttributeWithContext(ctx context.Context, exportTaskId *string, request *UpdateResourceExportTaskAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceExportTaskAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ExportToModule) {
		body["exportToModule"] = request.ExportToModule
	}

	if !dara.IsNil(request.IncludeRules) {
		body["includeRules"] = request.IncludeRules
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.TerraformProviderVersion) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !dara.IsNil(request.TerraformVersion) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !dara.IsNil(request.TriggerStrategy) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceExportTaskAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/exportTasks/" + dara.PercentEncode(dara.StringValue(exportTaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceExportTaskAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源栈
//
// @param request - UpdateStackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStackResponse
func (client *Client) UpdateStackWithContext(ctx context.Context, stackId *string, request *UpdateStackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.SourcePath) {
		body["sourcePath"] = request.SourcePath
	}

	if !dara.IsNil(request.WorkingDirectory) {
		body["workingDirectory"] = request.WorkingDirectory
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStack"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/stacks/" + dara.PercentEncode(dara.StringValue(stackId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改任务
//
// @param request - UpdateTaskAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskAttributeResponse
func (client *Client) UpdateTaskAttributeWithContext(ctx context.Context, taskId *string, request *UpdateTaskAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTaskAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoApply) {
		body["autoApply"] = request.AutoApply
	}

	if !dara.IsNil(request.AutoDestroy) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GroupInfo) {
		body["groupInfo"] = request.GroupInfo
	}

	if !dara.IsNil(request.InitModuleState) {
		body["initModuleState"] = request.InitModuleState
	}

	if !dara.IsNil(request.ModuleVersion) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ProtectionStrategy) {
		body["protectionStrategy"] = request.ProtectionStrategy
	}

	if !dara.IsNil(request.RamRole) {
		body["ramRole"] = request.RamRole
	}

	if !dara.IsNil(request.SkipPropertyValidation) {
		body["skipPropertyValidation"] = request.SkipPropertyValidation
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TerraformVersion) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !dara.IsNil(request.TriggerStrategy) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskAttribute"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模版上传
//
// @param request - UploadModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadModuleResponse
func (client *Client) UploadModuleWithContext(ctx context.Context, resourceType *string, request *UploadModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModuleId) {
		query["moduleId"] = request.ModuleId
	}

	if !dara.IsNil(request.ModuleName) {
		query["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.NamespaceName) {
		query["namespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.Url) {
		query["url"] = request.Url
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modules/upload/" + dara.PercentEncode(dara.StringValue(resourceType))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模版预检
//
// @param request - ValidateModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateModuleResponse
func (client *Client) ValidateModuleWithSSECtx(ctx context.Context, request *ValidateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ValidateModuleResponse, _yieldErr chan error) {
	defer close(_yield)
	client.validateModuleWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 模版预检
//
// @param request - ValidateModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateModuleResponse
func (client *Client) ValidateModuleWithContext(ctx context.Context, request *ValidateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.CodeMap) {
		body["codeMap"] = request.CodeMap
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.SourcePath) {
		body["sourcePath"] = request.SourcePath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/module/validation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateModuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) validateModuleWithSSECtx_opYieldFunc(_yield chan *ValidateModuleResponse, _yieldErr chan error, ctx context.Context, request *ValidateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.CodeMap) {
		body["codeMap"] = request.CodeMap
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.SourcePath) {
		body["sourcePath"] = request.SourcePath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateModule"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/module/validation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
