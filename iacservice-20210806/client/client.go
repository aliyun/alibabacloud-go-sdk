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
	client.Endpoint, _err = client.GetEndpoint(dara.String("iacservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 新增共享账号信息
//
// @param request - AddSharedAccountsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSharedAccountsResponse
func (client *Client) AddSharedAccountsWithOptions(request *AddSharedAccountsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddSharedAccountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增共享账号信息
//
// @param request - AddSharedAccountsRequest
//
// @return AddSharedAccountsResponse
func (client *Client) AddSharedAccounts(request *AddSharedAccountsRequest) (_result *AddSharedAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddSharedAccountsResponse{}
	_body, _err := client.AddSharedAccountsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AssociateGroupWithOptions(groupId *string, request *AssociateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AssociateGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AssociateGroupResponse
func (client *Client) AssociateGroup(groupId *string, request *AssociateGroupRequest) (_result *AssociateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AssociateGroupResponse{}
	_body, _err := client.AssociateGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelResourceExportTaskWithOptions(exportTaskId *string, request *CancelResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelResourceExportTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelResourceExportTaskResponse
func (client *Client) CancelResourceExportTask(exportTaskId *string, request *CancelResourceExportTaskRequest) (_result *CancelResourceExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelResourceExportTaskResponse{}
	_body, _err := client.CancelResourceExportTaskWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateJobWithOptions(taskId *string, request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateJobResponse
func (client *Client) CreateJob(taskId *string, request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模板
//
// @param request - CreateModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModuleResponse
func (client *Client) CreateModuleWithOptions(request *CreateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模板
//
// @param request - CreateModuleRequest
//
// @return CreateModuleResponse
func (client *Client) CreateModule(request *CreateModuleRequest) (_result *CreateModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModuleResponse{}
	_body, _err := client.CreateModuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模板版本
//
// @param request - CreateModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModuleVersionResponse
func (client *Client) CreateModuleVersionWithOptions(moduleId *string, request *CreateModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModuleVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模板版本
//
// @param request - CreateModuleVersionRequest
//
// @return CreateModuleVersionResponse
func (client *Client) CreateModuleVersion(moduleId *string, request *CreateModuleVersionRequest) (_result *CreateModuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModuleVersionResponse{}
	_body, _err := client.CreateModuleVersionWithOptions(moduleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRegistryModuleWithOptions(request *CreateRegistryModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRegistryModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRegistryModuleResponse
func (client *Client) CreateRegistryModule(request *CreateRegistryModuleRequest) (_result *CreateRegistryModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRegistryModuleResponse{}
	_body, _err := client.CreateRegistryModuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRegistryNamespaceWithOptions(request *CreateRegistryNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRegistryNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRegistryNamespaceResponse
func (client *Client) CreateRegistryNamespace(request *CreateRegistryNamespaceRequest) (_result *CreateRegistryNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRegistryNamespaceResponse{}
	_body, _err := client.CreateRegistryNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateResourceExportTaskWithOptions(request *CreateResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceExportTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateResourceExportTaskResponse
func (client *Client) CreateResourceExportTask(request *CreateResourceExportTaskRequest) (_result *CreateResourceExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceExportTaskResponse{}
	_body, _err := client.CreateResourceExportTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTaskResponse
func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(groupId *string) (_result *DeleteGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModuleResponse
func (client *Client) DeleteModuleWithOptions(moduleId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteModuleResponse
func (client *Client) DeleteModule(moduleId *string) (_result *DeleteModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModuleResponse{}
	_body, _err := client.DeleteModuleWithOptions(moduleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(projectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteProjectResponse
func (client *Client) DeleteProject(projectId *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除RegistryModule
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistryModuleResponse
func (client *Client) DeleteRegistryModuleWithOptions(namespaceName *string, moduleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRegistryModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRegistryModuleResponse
func (client *Client) DeleteRegistryModule(namespaceName *string, moduleName *string) (_result *DeleteRegistryModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRegistryModuleResponse{}
	_body, _err := client.DeleteRegistryModuleWithOptions(namespaceName, moduleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除RegistryModule版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistryModuleVersionResponse
func (client *Client) DeleteRegistryModuleVersionWithOptions(namespaceName *string, moduleName *string, version *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRegistryModuleVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRegistryModuleVersionResponse
func (client *Client) DeleteRegistryModuleVersion(namespaceName *string, moduleName *string, version *string) (_result *DeleteRegistryModuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRegistryModuleVersionResponse{}
	_body, _err := client.DeleteRegistryModuleVersionWithOptions(namespaceName, moduleName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistryNamespaceResponse
func (client *Client) DeleteRegistryNamespaceWithOptions(namespaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRegistryNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRegistryNamespaceResponse
func (client *Client) DeleteRegistryNamespace(namespaceName *string) (_result *DeleteRegistryNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRegistryNamespaceResponse{}
	_body, _err := client.DeleteRegistryNamespaceWithOptions(namespaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源导出任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceExportTaskResponse
func (client *Client) DeleteResourceExportTaskWithOptions(exportTaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceExportTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteResourceExportTaskResponse
func (client *Client) DeleteResourceExportTask(exportTaskId *string) (_result *DeleteResourceExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceExportTaskResponse{}
	_body, _err := client.DeleteResourceExportTaskWithOptions(exportTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteTaskResponse
func (client *Client) DeleteTask(taskId *string) (_result *DeleteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTaskResponse{}
	_body, _err := client.DeleteTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DissociateGroupWithOptions(projectId *string, groupId *string, request *DissociateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DissociateGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DissociateGroupResponse
func (client *Client) DissociateGroup(projectId *string, groupId *string, request *DissociateGroupRequest) (_result *DissociateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DissociateGroupResponse{}
	_body, _err := client.DissociateGroupWithOptions(projectId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteRegistryModuleWithOptions(namespaceName *string, moduleName *string, request *ExecuteRegistryModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteRegistryModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteRegistryModuleResponse
func (client *Client) ExecuteRegistryModule(namespaceName *string, moduleName *string, request *ExecuteRegistryModuleRequest) (_result *ExecuteRegistryModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteRegistryModuleResponse{}
	_body, _err := client.ExecuteRegistryModuleWithOptions(namespaceName, moduleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteResourceExportTaskWithOptions(exportTaskId *string, request *ExecuteResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteResourceExportTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteResourceExportTaskResponse
func (client *Client) ExecuteResourceExportTask(exportTaskId *string, request *ExecuteResourceExportTaskRequest) (_result *ExecuteResourceExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteResourceExportTaskResponse{}
	_body, _err := client.ExecuteResourceExportTaskWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteTerraformApplyWithOptions(request *ExecuteTerraformApplyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTerraformApplyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteTerraformApplyResponse
func (client *Client) ExecuteTerraformApply(request *ExecuteTerraformApplyRequest) (_result *ExecuteTerraformApplyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTerraformApplyResponse{}
	_body, _err := client.ExecuteTerraformApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteTerraformDestroyWithOptions(request *ExecuteTerraformDestroyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTerraformDestroyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteTerraformDestroyResponse
func (client *Client) ExecuteTerraformDestroy(request *ExecuteTerraformDestroyRequest) (_result *ExecuteTerraformDestroyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTerraformDestroyResponse{}
	_body, _err := client.ExecuteTerraformDestroyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteTerraformPlanWithOptions(request *ExecuteTerraformPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTerraformPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteTerraformPlanResponse
func (client *Client) ExecuteTerraformPlan(request *ExecuteTerraformPlanRequest) (_result *ExecuteTerraformPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTerraformPlanResponse{}
	_body, _err := client.ExecuteTerraformPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateModuleWithOptions(request *GenerateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateModuleResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return GenerateModuleResponse
func (client *Client) GenerateModule(request *GenerateModuleRequest) (_result *GenerateModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateModuleResponse{}
	_body, _err := client.GenerateModuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Terraform运行结果
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteStateResponse
func (client *Client) GetExecuteStateWithOptions(stateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExecuteStateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetExecuteStateResponse
func (client *Client) GetExecuteState(stateId *string) (_result *GetExecuteStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExecuteStateResponse{}
	_body, _err := client.GetExecuteStateWithOptions(stateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGroupResponse
func (client *Client) GetGroup(groupId *string) (_result *GetGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetJobWithOptions(taskId *string, jobId *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetJobResponse
func (client *Client) GetJob(taskId *string, jobId *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(taskId, jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模板详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModuleResponse
func (client *Client) GetModuleWithOptions(moduleId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模板详情
//
// @return GetModuleResponse
func (client *Client) GetModule(moduleId *string) (_result *GetModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModuleResponse{}
	_body, _err := client.GetModuleWithOptions(moduleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模板版本详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModuleVersionResponse
func (client *Client) GetModuleVersionWithOptions(moduleId *string, moduleVersion *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModuleVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetModuleVersionResponse
func (client *Client) GetModuleVersion(moduleId *string, moduleVersion *string) (_result *GetModuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModuleVersionResponse{}
	_body, _err := client.GetModuleVersionWithOptions(moduleId, moduleVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetProjectWithOptions(projectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetProjectResponse
func (client *Client) GetProject(projectId *string) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RegistryModule信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistryModuleResponse
func (client *Client) GetRegistryModuleWithOptions(namespaceName *string, moduleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegistryModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRegistryModuleResponse
func (client *Client) GetRegistryModule(namespaceName *string, moduleName *string) (_result *GetRegistryModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegistryModuleResponse{}
	_body, _err := client.GetRegistryModuleWithOptions(namespaceName, moduleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RegistryModule版本信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistryModuleVersionResponse
func (client *Client) GetRegistryModuleVersionWithOptions(namespaceName *string, moduleName *string, version *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegistryModuleVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRegistryModuleVersionResponse
func (client *Client) GetRegistryModuleVersion(namespaceName *string, moduleName *string, version *string) (_result *GetRegistryModuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegistryModuleVersionResponse{}
	_body, _err := client.GetRegistryModuleVersionWithOptions(namespaceName, moduleName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作空间信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistryNamespaceResponse
func (client *Client) GetRegistryNamespaceWithOptions(namespaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegistryNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRegistryNamespaceResponse
func (client *Client) GetRegistryNamespace(namespaceName *string) (_result *GetRegistryNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegistryNamespaceResponse{}
	_body, _err := client.GetRegistryNamespaceWithOptions(namespaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceExportTaskWithOptions(exportTaskId *string, request *GetResourceExportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceExportTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetResourceExportTaskResponse
func (client *Client) GetResourceExportTask(exportTaskId *string, request *GetResourceExportTaskRequest) (_result *GetResourceExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceExportTaskResponse{}
	_body, _err := client.GetResourceExportTaskWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceTypeWithOptions(resourceType *string, request *GetResourceTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return GetResourceTypeResponse
func (client *Client) GetResourceType(resourceType *string, request *GetResourceTypeRequest) (_result *GetResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceTypeResponse{}
	_body, _err := client.GetResourceTypeWithOptions(resourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTaskResponse
func (client *Client) GetTask(taskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListExplorerRegistryModuleExamplesWithOptions(request *ListExplorerRegistryModuleExamplesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExplorerRegistryModuleExamplesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return ListExplorerRegistryModuleExamplesResponse
func (client *Client) ListExplorerRegistryModuleExamples(request *ListExplorerRegistryModuleExamplesRequest) (_result *ListExplorerRegistryModuleExamplesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExplorerRegistryModuleExamplesResponse{}
	_body, _err := client.ListExplorerRegistryModuleExamplesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListExplorerRegistryModuleVersionsWithOptions(request *ListExplorerRegistryModuleVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExplorerRegistryModuleVersionsResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return ListExplorerRegistryModuleVersionsResponse
func (client *Client) ListExplorerRegistryModuleVersions(request *ListExplorerRegistryModuleVersionsRequest) (_result *ListExplorerRegistryModuleVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExplorerRegistryModuleVersionsResponse{}
	_body, _err := client.ListExplorerRegistryModuleVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListExplorerRegistryModulesWithOptions(request *ListExplorerRegistryModulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExplorerRegistryModulesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return ListExplorerRegistryModulesResponse
func (client *Client) ListExplorerRegistryModules(request *ListExplorerRegistryModulesRequest) (_result *ListExplorerRegistryModulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExplorerRegistryModulesResponse{}
	_body, _err := client.ListExplorerRegistryModulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListGroupWithOptions(tmpReq *ListGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListGroupRequest
//
// @return ListGroupResponse
func (client *Client) ListGroup(request *ListGroupRequest) (_result *ListGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupResponse{}
	_body, _err := client.ListGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListJobsWithOptions(taskId *string, request *ListJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListJobsResponse
func (client *Client) ListJobs(taskId *string, request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListModuleVersionWithOptions(moduleId *string, request *ListModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModuleVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListModuleVersionResponse
func (client *Client) ListModuleVersion(moduleId *string, request *ListModuleVersionRequest) (_result *ListModuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModuleVersionResponse{}
	_body, _err := client.ListModuleVersionWithOptions(moduleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListModulesWithOptions(tmpReq *ListModulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListModulesRequest
//
// @return ListModulesResponse
func (client *Client) ListModules(request *ListModulesRequest) (_result *ListModulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModulesResponse{}
	_body, _err := client.ListModulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProductsWithOptions(request *ListProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return ListProductsResponse
func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProjectWithOptions(tmpReq *ListProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListProjectRequest
//
// @return ListProjectResponse
func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRegistryModuleVersionsWithOptions(request *ListRegistryModuleVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegistryModuleVersionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRegistryModuleVersionsResponse
func (client *Client) ListRegistryModuleVersions(request *ListRegistryModuleVersionsRequest) (_result *ListRegistryModuleVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegistryModuleVersionsResponse{}
	_body, _err := client.ListRegistryModuleVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRegistryModulesWithOptions(request *ListRegistryModulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegistryModulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRegistryModulesResponse
func (client *Client) ListRegistryModules(request *ListRegistryModulesRequest) (_result *ListRegistryModulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegistryModulesResponse{}
	_body, _err := client.ListRegistryModulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRegistryNamespacesWithOptions(request *ListRegistryNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegistryNamespacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRegistryNamespacesResponse
func (client *Client) ListRegistryNamespaces(request *ListRegistryNamespacesRequest) (_result *ListRegistryNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegistryNamespacesResponse{}
	_body, _err := client.ListRegistryNamespacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResourceExportTaskVersionsWithOptions(exportTaskId *string, request *ListResourceExportTaskVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceExportTaskVersionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListResourceExportTaskVersionsResponse
func (client *Client) ListResourceExportTaskVersions(exportTaskId *string, request *ListResourceExportTaskVersionsRequest) (_result *ListResourceExportTaskVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceExportTaskVersionsResponse{}
	_body, _err := client.ListResourceExportTaskVersionsWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResourceExportTasksWithOptions(request *ListResourceExportTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceExportTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListResourceExportTasksResponse
func (client *Client) ListResourceExportTasks(request *ListResourceExportTasksRequest) (_result *ListResourceExportTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceExportTasksResponse{}
	_body, _err := client.ListResourceExportTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResourceTypesWithOptions(tmpReq *ListResourceTypesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @param request - ListResourceTypesRequest
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTasksWithOptions(tmpReq *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTerraformProviderVersionsWithOptions(request *ListTerraformProviderVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTerraformProviderVersionsResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return ListTerraformProviderVersionsResponse
func (client *Client) ListTerraformProviderVersions(request *ListTerraformProviderVersionsRequest) (_result *ListTerraformProviderVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTerraformProviderVersionsResponse{}
	_body, _err := client.ListTerraformProviderVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateJobWithOptions(taskId *string, jobId *string, operationType *string, request *OperateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OperateJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateJobResponse
func (client *Client) OperateJob(taskId *string, jobId *string, operationType *string, request *OperateJobRequest) (_result *OperateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OperateJobResponse{}
	_body, _err := client.OperateJobWithOptions(taskId, jobId, operationType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PublishRegistryModuleVersionWithOptions(request *PublishRegistryModuleVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishRegistryModuleVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PublishRegistryModuleVersionResponse
func (client *Client) PublishRegistryModuleVersion(request *PublishRegistryModuleVersionRequest) (_result *PublishRegistryModuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishRegistryModuleVersionResponse{}
	_body, _err := client.PublishRegistryModuleVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveSharedAccountsWithOptions(tmpReq *RemoveSharedAccountsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveSharedAccountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RemoveSharedAccountsRequest
//
// @return RemoveSharedAccountsResponse
func (client *Client) RemoveSharedAccounts(request *RemoveSharedAccountsRequest) (_result *RemoveSharedAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveSharedAccountsResponse{}
	_body, _err := client.RemoveSharedAccountsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateExplorerModuleAttributeWithOptions(explorerModuleId *string, request *UpdateExplorerModuleAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExplorerModuleAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateExplorerModuleAttributeResponse
func (client *Client) UpdateExplorerModuleAttribute(explorerModuleId *string, request *UpdateExplorerModuleAttributeRequest) (_result *UpdateExplorerModuleAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExplorerModuleAttributeResponse{}
	_body, _err := client.UpdateExplorerModuleAttributeWithOptions(explorerModuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateGroupWithOptions(groupId *string, request *UpdateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(groupId *string, request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新模板
//
// @param request - UpdateModuleAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModuleAttributeResponse
func (client *Client) UpdateModuleAttributeWithOptions(moduleId *string, request *UpdateModuleAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModuleAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模板
//
// @param request - UpdateModuleAttributeRequest
//
// @return UpdateModuleAttributeResponse
func (client *Client) UpdateModuleAttribute(moduleId *string, request *UpdateModuleAttributeRequest) (_result *UpdateModuleAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModuleAttributeResponse{}
	_body, _err := client.UpdateModuleAttributeWithOptions(moduleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateProjectWithOptions(projectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateProjectResponse
func (client *Client) UpdateProject(projectId *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRegistryModuleAttributeWithOptions(namespaceName *string, moduleName *string, request *UpdateRegistryModuleAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRegistryModuleAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRegistryModuleAttributeResponse
func (client *Client) UpdateRegistryModuleAttribute(namespaceName *string, moduleName *string, request *UpdateRegistryModuleAttributeRequest) (_result *UpdateRegistryModuleAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRegistryModuleAttributeResponse{}
	_body, _err := client.UpdateRegistryModuleAttributeWithOptions(namespaceName, moduleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRegistryNamespaceAttributeWithOptions(namespaceName *string, request *UpdateRegistryNamespaceAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRegistryNamespaceAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRegistryNamespaceAttributeResponse
func (client *Client) UpdateRegistryNamespaceAttribute(namespaceName *string, request *UpdateRegistryNamespaceAttributeRequest) (_result *UpdateRegistryNamespaceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRegistryNamespaceAttributeResponse{}
	_body, _err := client.UpdateRegistryNamespaceAttributeWithOptions(namespaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateResourceExportTaskAttributeWithOptions(exportTaskId *string, request *UpdateResourceExportTaskAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceExportTaskAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateResourceExportTaskAttributeResponse
func (client *Client) UpdateResourceExportTaskAttribute(exportTaskId *string, request *UpdateResourceExportTaskAttributeRequest) (_result *UpdateResourceExportTaskAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceExportTaskAttributeResponse{}
	_body, _err := client.UpdateResourceExportTaskAttributeWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTaskAttributeWithOptions(taskId *string, request *UpdateTaskAttributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTaskAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateTaskAttributeResponse
func (client *Client) UpdateTaskAttribute(taskId *string, request *UpdateTaskAttributeRequest) (_result *UpdateTaskAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTaskAttributeResponse{}
	_body, _err := client.UpdateTaskAttributeWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ValidateModuleWithSSE(request *ValidateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ValidateModuleResponse, _yieldErr chan error) {
	defer close(_yield)
	client.validateModuleWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
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
func (client *Client) ValidateModuleWithOptions(request *ValidateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateModuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ValidateModuleResponse
func (client *Client) ValidateModule(request *ValidateModuleRequest) (_result *ValidateModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateModuleResponse{}
	_body, _err := client.ValidateModuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) validateModuleWithSSE_opYieldFunc(_yield chan *ValidateModuleResponse, _yieldErr chan error, request *ValidateModuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
