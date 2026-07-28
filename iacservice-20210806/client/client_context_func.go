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
// Adds shared accounts.
//
// Description:
//
// Per-user call frequency: 100 calls per second.
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
// # Associate drift detection configuration
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
// Associates resources with a group.
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
// Associates parameter sets.
//
// Description:
//
// After creating a parameter set, you need to associate it with a resource. Valid values for the resource type:
//
// - Module: template
//
// - ModuleVersion: template version
//
// - Task: node.
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
// Cancels a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Creates a drift detection configuration that supports manual or scheduled triggering.
//
// Description:
//
// ## Request Description
//
// - When `triggerType` is set to `Cron`, a valid `cronExpression` must be provided.
//
// - Each element in the `alarmConfigs` list must specify the alerting method `type` and the corresponding alerting address `address`.
//
// - If the `enabled` parameter is not explicitly set, its default value is `true`, meaning newly created detection configurations are enabled by default.
//
// - It is recommended to use a UUID as the value of `clientToken` to ensure request idempotence.
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
// Creates a group.
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
// Creates a job and runs a task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Creates a Terraform template. Multiple source methods are supported, such as OSS import, Registry import, file upload, and online editing.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Publishes a new version for a specified template.
//
// Description:
//
// ## Operation description
//
// - Use the `clientToken` parameter to ensure idempotence of the request and prevent duplicate submissions caused by network retries.
//
// - Use semantic versioning (such as `v1.0.0`).
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
// Adds a new parameter set. You can set the name, description, and parameter list.
//
// Description:
//
// ## Operation description
//
// - This operation creates a new parameter set.
//
// - The name field is required and can be up to 128 characters in length.
//
// - Each element in the parameters array must contain the name field. Other fields are optional.
//
// - Use the clientToken field to ensure the idempotence of the request.
//
// - The request header must contain authentication information to ensure secure access.
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
// Creates a project.
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
// Creates a Registry template.
//
// Description:
//
// Per-user call frequency: 100 calls per second.
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
// Creates a workspace.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Creates a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Creates a resource stack and triggers deployment.
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

	if !dara.IsNil(request.ParameterSetIds) {
		body["parameterSetIds"] = request.ParameterSetIds
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
// Creates a node.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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

	if !dara.IsNil(request.ParameterSetIds) {
		body["parameterSetIds"] = request.ParameterSetIds
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

	if !dara.IsNil(request.SkipRegionValidation) {
		body["skipRegionValidation"] = request.SkipRegionValidation
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TaskBackend) {
		body["taskBackend"] = request.TaskBackend
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
// # Delete drift detection configuration
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
// Deletes a group.
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
// Deletes a specified template and all its versions.
//
// Description:
//
// ## Operation description
//
// - This operation deletes a specified template.
//
// - Deletion is irreversible. Proceed with caution.
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
// Deletes a specified parameter set by parameter set ID.
//
// Description:
//
// Deletes a specified parameter set.
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
// Deletes a project.
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
// Deletes a Registry template.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Deletes a Registry template version.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Deletes a workspace.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Deletes a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Deletes a stack.
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
// Deletes a node.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
//
// Deletes a node. If the node has resources that have not been destroyed, the node cannot be deleted.
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
// Initiates a state file consistency check.
//
// Description:
//
// This API is used to perform drift detection on the state files of resource orchestration tasks and stack tasks in the automated service desk.
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
// # Disassociate drift detection configuration
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
// Dissociates a resource group.
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
// Dissociates a parameter set from other resources.
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
// Executes a Module officially provided by Alibaba Cloud Terraform.
//
// Description:
//
// This API operation is used to execute Terraform Module code to create or update cloud resources. Before using this API operation, make sure that all required authentication information is correctly configured and that the Terraform code corresponding to the Module meets the expected functional requirements.
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
// Runs a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Executes TerraformApply.
//
// Description:
//
// Executes the Terraform Apply command to create or update cloud resources based on the provided Terraform code. This API can handle complex scenarios such as operations that depend on a previous state.
//
// Before calling this API, ensure that all required authentication information is properly configured and that the Terraform code meets the expected functional requirements.
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
// Executes Terraform Destroy.
//
// Description:
//
// Executes the Terraform Destroy command to destroy resources created by Terraform.
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
// Executes a Terraform plan.
//
// Description:
//
// Executes a Terraform Plan command by using the provided Terraform code to create or update cloud resources. This API operation can handle complex scenarios such as operations that depend on a previous state.
//
// Before calling this API operation, ensure that all required authentication information is properly configured and that the Terraform code meets the expected functional requirements.
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
// Generates Terraform HCL template code.
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
// # Retrieve drift detection configuration
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
// Retrieves the result of a Terraform run.
//
// Description:
//
// Retrieves the result of a Terraform run.
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
// Queries a group.
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
// Retrieves job information.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Queries the details of a specified template.
//
// Description:
//
// ## Operation description
//
// You can call this operation to query the details of a specified template, including but not limited to the template name, description, source, status, and latest version. You must specify the template ID and include authentication information in the request.
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
// Queries the details of a specific version of a specified template.
//
// Description:
//
// ## Operation description
//
// You can call this operation to query the details of a specific version of a specified template, including the version number, description, and release time. Make sure that the template ID and version number are correct.
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
// Retrieves the details of a parameter set by parameter set ID.
//
// Description:
//
// ## Description
//
// - This operation retrieves detailed parameter set information by specifying a parameterSetId.
//
// - Authentication is required to call this operation.
//
// - If the request succeeds, the response includes detailed data such as the parameter set name, description, and parameter list.
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
// Queries a project.
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
// Retrieves the resource documentation of a Terraform provider.
//
// @param request - GetProviderDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProviderDocumentResponse
func (client *Client) GetProviderDocumentWithContext(ctx context.Context, request *GetProviderDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProviderDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProviderVersion) {
		query["providerVersion"] = request.ProviderVersion
	}

	if !dara.IsNil(request.TerraformResourceType) {
		query["terraformResourceType"] = request.TerraformResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProviderDocument"),
		Version:     dara.String("2021-08-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/version/terraform/provider/document"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProviderDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a Registry module.
//
// Description:
//
// Single-user call frequency: 200 calls per second.
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
// Queries a Registry template version.
//
// Description:
//
// Single-user call frequency: 200 calls per second.
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
// Queries a workspace.
//
// Description:
//
// Single-user call frequency: 200 calls per second.
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
// Queries the details of a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Retrieves resource type information.
//
// Description:
//
// ## Request description.
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
// Queries a stack.
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
// Queries the list of deployments for a stack.
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
// Retrieves the trigger result of a stack.
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
// Retrieves the details of a task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Retrieves the detection result of a state file.
//
// Description:
//
// This API is used to retrieve the detection results of state files for resource orchestration tasks and stack tasks on the automation service desk.
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
// # List drift detection associations
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
// # List drift detection configurations
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
// Retrieves the list of official Terraform Module examples.
//
// Description:
//
// This operation queries the example information of Terraform Modules officially provided by Alibaba Cloud.
//
// You can use the `maxResults` parameter to adjust the maximum number of entries to return.
//
// - If `nextToken` is not included in the response, no more data is available. Otherwise, more data is available. To query the next page, set the `nextToken` parameter of the ListExplorerRegistryModuleExamples operation to the `nextToken` value returned in the previous response. If the `NextToken` parameter is not specified, the first page of data is returned by default.
//
// - You can use keyword, namespaceName, moduleName, moduleVersion, and exampleName as conditional filter settings to narrow down the search scope. Multiple filter conditions have a logical `AND` relationship, and only resources that meet all filter conditions are returned.
//
//   - keyword: optional. Searches by keyword and supports fuzzy match on exampleName. For example, if keyword is set to ecs, module examples whose names contain ecs are returned.
//
//   - namespaceName: optional. Filters module examples by a specific workspace. For example, if namespaceName is set to alibaba, module examples in the alibaba workspace are returned.
//
//   - moduleName: optional. Filters module examples by a specific module name. For example, if moduleName is set to ecs, module examples whose module name is ecs are returned.
//
//   - moduleVersion: optional. Filters module examples by a specific module version. For example, if moduleVersion is set to 1.0.0, module examples whose module version is 1.0.0 are returned.
//
//   - exampleName: optional. Filters module examples by a specific example name. For example, if exampleName is set to ecs, module examples whose example name is ecs are returned.
//
// The response contains the request ID, total number of entries, data of the current page, and pagination information, which facilitates processing of query results.
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
// Lists the version information of official Terraform modules provided by Alibaba Cloud.
//
// Description:
//
// This operation queries the version information of official Terraform modules provided by Alibaba Cloud.
//
// You can use the `maxResults` parameter to adjust the maximum number of entries to return.
//
// - If `nextToken` is not included in the response, no more data is available. Otherwise, more data is available. To query the next page, set the `nextToken` parameter of the ListExplorerRegistryModules operation to the `nextToken` value returned in the previous response. If the `NextToken` parameter is not specified, the first page of data is returned by default.
//
// - You can use keyword, namespaceName, moduleName, and moduleVersion as conditional filter Settings to narrow the search scope. Multiple filter conditions have a logical `AND` relationship. Only resources that meet all filter conditions are returned.
//
//   - keyword: optional. Performs a fuzzy match on the module name. For example, if keyword is set to ecs, modules whose names contain ecs are returned.
//
//   - namespaceName: optional. Filters modules by a specific workspace. For example, if namespaceName is set to alibaba, modules whose workspace is alibaba are returned. When moduleName is specified, namespaceName must also be specified. You can call the ListExplorerRegistryModule operation to obtain the namespaceName information.
//
//   - moduleName: optional. Filters modules by a specific name. For example, if moduleName is set to ecs, modules whose name is ecs are returned.
//
//   - moduleVersion: optional. Filters modules by a specific version. For example, if moduleVersion is set to 1.0.0, modules whose version is 1.0.0 are returned.
//
// The response contains the request ID, total number of entries, data on the current page, and pagination information, which facilitates the processing of query results.
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
// Lists information about official Terraform modules provided by Alibaba Cloud.
//
// Description:
//
// This operation queries information about official Terraform modules provided by Alibaba Cloud.
//
// You can use the `maxResults` parameter to adjust the maximum number of entries to return.
//
// - If the `nextToken` parameter is not included in the response, no more data is available. Otherwise, more data is available. To query the next page, set the `nextToken` parameter of the ListExplorerRegistryModules operation to the `nextToken` value returned in the previous response. If you do not specify the `NextToken` parameter, the first page of data is returned by default.
//
// - You can use keyword and moduleName as filter conditions to narrow the search scope. Multiple filter conditions are evaluated by using a logical `AND`. Only resources that meet all filter conditions are returned.
//
//   - keyword: optional. Searches by keyword through fuzzy matching against ModuleName. For example, if keyword is set to ecs, modules whose names contain ecs are returned.
//
//   - moduleName: optional. Filters modules by a specific name. For example, if moduleName is set to ecs, only the module whose name is exactly ecs is returned.
//
// The response contains the request ID, total number of entries, data of the current page, and pagination information, which facilitates the processing of query results.
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
// Queries the list of groups.
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
// Queries a list of jobs.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Retrieves a list of template versions.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Retrieves a list of templates for the current user, with support for pagination and conditional filtering.
//
// Description:
//
// ## Operation description
//
// This operation lists all Terraform templates for the current user. You can specify query parameters to implement pagination, fuzzy match template names, and filter templates by source or status. You can also filter templates by tag for more granular results.
//
// ### Notes
//
// - Use the pageNumber and pageSize parameters to control the number of returned results.
//
// - Use the name parameter to perform a fuzzy match on template names.
//
// - Use the source parameter to filter templates by source, such as OSS import or file upload.
//
// - Use the status parameter to filter templates by status, such as Created or Published.
//
// - Tag-based filtering requires a JSON-formatted string, for example, `[{"key":"env","value":"prod"}]`.
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
// Lists the parameter sets associated with a resource.
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
// Queries and retrieves a paginated list of parameter sets with keyword search support.
//
// Description:
//
// ## Operation description
//
// This operation queries all parameter sets in the system. You can filter results by keyword and paginate the results. Authentication information is required.
//
// ### Notes
//
// - The keyword parameter can be used to perform a fuzzy match on parameter sets by name or description.
//
// - Pagination is controlled by pageNumber and pageSize. Results start from the first page by default. Set pageSize to a reasonable value to avoid performance issues.
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

	if !dara.IsNil(request.KmsKeyId) {
		query["kmsKeyId"] = request.KmsKeyId
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
// Queries the list of all products.
//
// Description:
//
// ## Operation description
//
// - **Keyword search**: Use the `keyword` parameter for fuzzy matching.
//
// - **Paged query**: Use `nextToken` for pagination and `maxResults` to specify the maximum number of results per page (default: 100, maximum: 200).
//
// - **Terraform Provider version**: The optional `terraformProviderVersion` parameter filters products associated with a specific Provider version.
//
// - **Response structure**: The response contains the request ID, total number of entries, data of the current page, and pagination information for easy processing of query results.
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
// Queries the list of projects.
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
// Queries the list of Registry template versions.
//
// Description:
//
// Single-user call frequency: 200 calls per second.
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
// Queries a list of registry modules.
//
// Description:
//
// Single-user call frequency: 200 calls per second.
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
// Queries the list of workspaces.
//
// Description:
//
// Single-user call frequency: 200 calls per second.
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
// Retrieves the list of versions for a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Queries the list of resource export tasks.
//
// Description:
//
// Rate limit per user: 100 calls per second.
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
// Queries a list of resource types by filter conditions with pagination support.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to perform a conditional query for a list of resource types based on conditions such as product code, Terraform provider version, child class, status, and keyword. The results include detailed information about each resource, such as the product code, status, status effective version, child class, Terraform provider version, and resource type code. Paging is supported to facilitate handling large amounts of data.
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
// Retrieves the resources of a node.
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
// Queries the list of stack configurations.
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
// Queries the list of stacks.
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

	if !dara.IsNil(request.KmsKeyId) {
		query["kmsKeyId"] = request.KmsKeyId
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
// Queries a list of tasks.
//
// Description:
//
// The maximum number of times that a single user can call this operation per second: 100.
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

	if !dara.IsNil(request.KmsKeyId) {
		query["kmsKeyId"] = request.KmsKeyId
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
// Retrieves the list of Terraform provider versions.
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
// Supports resource import and removal for state files.
//
// Description:
//
// This API is used to manage state files for resource orchestration tasks and stack tasks on the automated service desk.
//
// Before using this API, make sure that all required authentication information is correctly configured and that the Terraform code meets the expected functional requirements.
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
// After a job is created, you can perform the **Cancel*	- operation to stop the job while it is running.
//
// After a job reaches the pending confirmation state, you can perform the **Abolish*	- operation to stop the job, or perform the **Execute*	- operation to continue the job execution.
//
// Description:
//
// Per-user call frequency: 100 calls per second.
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
// Publishes a Registry template version.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Removes a shared account.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// # Trigger Stack execution
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
// Updates the drift detection configuration information for the specified ID.
//
// Description:
//
// ## Request Description
//
// - `detectConfigId` is a required parameter used to identify the specific detection configuration to update.
//
// - When `triggerType` is set to `Cron`, a valid `cronExpression` must be provided.
//
// - Each element in the `alarmConfigs` list must include an alert type (`type`) and an address (`address`).
//
// - If you do not want to change certain properties (such as `name`, `description`, etc.), you can omit these fields from the request body.
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
// Updates an Explorer template.
//
// Description:
//
// Updates an Explorer template.
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
// Modifies a group.
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
// Updates the name, description, tags, and other information of a specified template.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to modify the basic attributes of an existing template, including but not limited to the template name, description, and tags.
//
// - The update operation does not affect the content or version information of the template.
//
// - To enable or disable deletion protection, use the deletionProtection parameter.
//
// - Use clientToken to ensure the idempotence of the request and avoid duplicate submissions caused by network issues.
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
// Updates the attributes of a specified parameter set, such as the name and description.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to modify the basic information of an existing parameter set, including the name and description.
//
// - If the request includes the parameters field, the parameter list in the parameter set is updated.
//
// - The clientToken field can be used to ensure the idempotence of the request.
//
// - The update operation requires a valid parameterSetId as a path parameter.
//
// - The request must include authentication information to pass identity verification.
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
// Updates project information.
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
// Updates a Registry template.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Modifies a workspace.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Modifies a resource export task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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
// Modifies a stack. When the configuration changes, a stack deployment is triggered.
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
// Updates the properties of a task.
//
// Description:
//
// Single-user call frequency: 100 calls per second.
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

	if !dara.IsNil(request.SkipRegionValidation) {
		body["skipRegionValidation"] = request.SkipRegionValidation
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
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
// Uploads a template.
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
// Performs a dry run on a template.
//
// Description:
//
// Performs a dry run on the content of a Terraform configuration file.
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
// Performs a dry run on a template.
//
// Description:
//
// Performs a dry run on the content of a Terraform configuration file.
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
