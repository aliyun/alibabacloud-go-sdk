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
	client.Endpoint, _err = client.GetEndpoint(dara.String("cms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加记忆
//
// @param request - AddMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemoriesResponse
func (client *Client) AddMemoriesWithOptions(workspace *string, memoryStoreName *string, request *AddMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AsyncMode) {
		body["asyncMode"] = request.AsyncMode
	}

	if !dara.IsNil(request.CustomInstructions) {
		body["customInstructions"] = request.CustomInstructions
	}

	if !dara.IsNil(request.Infer) {
		body["infer"] = request.Infer
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.Timestamp) {
		body["timestamp"] = request.Timestamp
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMemories"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加记忆
//
// @param request - AddMemoriesRequest
//
// @return AddMemoriesResponse
func (client *Client) AddMemories(workspace *string, memoryStoreName *string, request *AddMemoriesRequest) (_result *AddMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddMemoriesResponse{}
	_body, _err := client.AddMemoriesWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改资源所属资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourcegroup"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改资源所属资源组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Install the access component, representing a single access attempt
//
// Description:
//
// # Used to create a site monitoring task
//
// @param request - CreateAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAddonReleaseResponse
func (client *Client) CreateAddonReleaseWithOptions(policyId *string, request *CreateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAddonReleaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		body["addonName"] = request.AddonName
	}

	if !dara.IsNil(request.AliyunLang) {
		body["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EntityRules) {
		body["entityRules"] = request.EntityRules
	}

	if !dara.IsNil(request.EnvType) {
		body["envType"] = request.EnvType
	}

	if !dara.IsNil(request.ParentAddonReleaseId) {
		body["parentAddonReleaseId"] = request.ParentAddonReleaseId
	}

	if !dara.IsNil(request.ReleaseName) {
		body["releaseName"] = request.ReleaseName
	}

	if !dara.IsNil(request.Values) {
		body["values"] = request.Values
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAddonRelease"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/addon-releases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAddonReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Install the access component, representing a single access attempt
//
// Description:
//
// # Used to create a site monitoring task
//
// @param request - CreateAddonReleaseRequest
//
// @return CreateAddonReleaseResponse
func (client *Client) CreateAddonRelease(policyId *string, request *CreateAddonReleaseRequest) (_result *CreateAddonReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAddonReleaseResponse{}
	_body, _err := client.CreateAddonReleaseWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Aggregation Task Group
//
// @param request - CreateAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggTaskGroupResponse
func (client *Client) CreateAggTaskGroupWithOptions(instanceId *string, request *CreateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAggTaskGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OverrideIfExists) {
		query["overrideIfExists"] = request.OverrideIfExists
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AggTaskGroupConfig) {
		body["aggTaskGroupConfig"] = request.AggTaskGroupConfig
	}

	if !dara.IsNil(request.AggTaskGroupConfigType) {
		body["aggTaskGroupConfigType"] = request.AggTaskGroupConfigType
	}

	if !dara.IsNil(request.AggTaskGroupName) {
		body["aggTaskGroupName"] = request.AggTaskGroupName
	}

	if !dara.IsNil(request.CronExpr) {
		body["cronExpr"] = request.CronExpr
	}

	if !dara.IsNil(request.Delay) {
		body["delay"] = request.Delay
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.FromTime) {
		body["fromTime"] = request.FromTime
	}

	if !dara.IsNil(request.MaxRetries) {
		body["maxRetries"] = request.MaxRetries
	}

	if !dara.IsNil(request.MaxRunTimeInSeconds) {
		body["maxRunTimeInSeconds"] = request.MaxRunTimeInSeconds
	}

	if !dara.IsNil(request.PrecheckString) {
		body["precheckString"] = request.PrecheckString
	}

	if !dara.IsNil(request.ScheduleMode) {
		body["scheduleMode"] = request.ScheduleMode
	}

	if !dara.IsNil(request.ScheduleTimeExpr) {
		body["scheduleTimeExpr"] = request.ScheduleTimeExpr
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TargetPrometheusId) {
		body["targetPrometheusId"] = request.TargetPrometheusId
	}

	if !dara.IsNil(request.ToTime) {
		body["toTime"] = request.ToTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAggTaskGroup"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/agg-task-groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAggTaskGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Aggregation Task Group
//
// @param request - CreateAggTaskGroupRequest
//
// @return CreateAggTaskGroupResponse
func (client *Client) CreateAggTaskGroup(instanceId *string, request *CreateAggTaskGroupRequest) (_result *CreateAggTaskGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAggTaskGroupResponse{}
	_body, _err := client.CreateAggTaskGroupWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Webhook
//
// @param request - CreateAlertWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertWebhookResponse
func (client *Client) CreateAlertWebhookWithOptions(request *CreateAlertWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlertWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.Headers) {
		body["headers"] = request.Headers
	}

	if !dara.IsNil(request.Lang) {
		body["lang"] = request.Lang
	}

	if !dara.IsNil(request.Method) {
		body["method"] = request.Method
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	if !dara.IsNil(request.WebhookId) {
		body["webhookId"] = request.WebhookId
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertWebhook"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webhook"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Webhook
//
// @param request - CreateAlertWebhookRequest
//
// @return CreateAlertWebhookResponse
func (client *Client) CreateAlertWebhook(request *CreateAlertWebhookRequest) (_result *CreateAlertWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlertWebhookResponse{}
	_body, _err := client.CreateAlertWebhookWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建业务链路
//
// @param request - CreateBizTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizTraceResponse
func (client *Client) CreateBizTraceWithOptions(request *CreateBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBizTraceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedConfig) {
		body["advancedConfig"] = request.AdvancedConfig
	}

	if !dara.IsNil(request.BizTraceCode) {
		body["bizTraceCode"] = request.BizTraceCode
	}

	if !dara.IsNil(request.BizTraceName) {
		body["bizTraceName"] = request.BizTraceName
	}

	if !dara.IsNil(request.RuleConfig) {
		body["ruleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizTrace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bizTrace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建业务链路
//
// @param request - CreateBizTraceRequest
//
// @return CreateBizTraceResponse
func (client *Client) CreateBizTrace(request *CreateBizTraceRequest) (_result *CreateBizTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBizTraceResponse{}
	_body, _err := client.CreateBizTraceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建云资源中心
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudResourceResponse
func (client *Client) CreateCloudResourceWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCloudResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudResource"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cloudresource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建云资源中心
//
// @return CreateCloudResourceResponse
func (client *Client) CreateCloudResource() (_result *CreateCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCloudResourceResponse{}
	_body, _err := client.CreateCloudResourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithOptions(workspace *string, request *CreateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		body["datasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Schema) {
		body["schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/dataset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDatasetRequest
//
// @return CreateDatasetResponse
func (client *Client) CreateDataset(workspace *string, request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据投递任务
//
// @param request - CreateDeliveryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryTaskResponse
func (client *Client) CreateDeliveryTaskWithOptions(request *CreateDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["dataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.ExternalLabels) {
		body["externalLabels"] = request.ExternalLabels
	}

	if !dara.IsNil(request.LabelFilters) {
		body["labelFilters"] = request.LabelFilters
	}

	if !dara.IsNil(request.LabelFiltersType) {
		body["labelFiltersType"] = request.LabelFiltersType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SinkList) {
		body["sinkList"] = request.SinkList
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	if !dara.IsNil(request.TaskName) {
		body["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeliveryTask"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/delivery-tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据投递任务
//
// @param request - CreateDeliveryTaskRequest
//
// @return CreateDeliveryTaskResponse
func (client *Client) CreateDeliveryTask(request *CreateDeliveryTaskRequest) (_result *CreateDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDeliveryTaskResponse{}
	_body, _err := client.CreateDeliveryTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create storage related to EntityStore
//
// @param request - CreateEntityStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStoreWithOptions(workspaceName *string, request *CreateEntityStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEntityStoreResponse, _err error) {
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
		Action:      dara.String("CreateEntityStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/entitystore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEntityStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create storage related to EntityStore
//
// @param request - CreateEntityStoreRequest
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStore(workspaceName *string, request *CreateEntityStoreRequest) (_result *CreateEntityStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEntityStoreResponse{}
	_body, _err := client.CreateEntityStoreWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Access Center Policy
//
// Description:
//
// This interface is used to support users in creating event integration.
//
// @param request - CreateIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntegrationPolicyResponse
func (client *Client) CreateIntegrationPolicyWithOptions(request *CreateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIntegrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityGroup) {
		body["entityGroup"] = request.EntityGroup
	}

	if !dara.IsNil(request.PolicyName) {
		body["policyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		body["policyType"] = request.PolicyType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntegrationPolicy"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntegrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Access Center Policy
//
// Description:
//
// This interface is used to support users in creating event integration.
//
// @param request - CreateIntegrationPolicyRequest
//
// @return CreateIntegrationPolicyResponse
func (client *Client) CreateIntegrationPolicy(request *CreateIntegrationPolicyRequest) (_result *CreateIntegrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIntegrationPolicyResponse{}
	_body, _err := client.CreateIntegrationPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建记忆库
//
// @param request - CreateMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryStoreResponse
func (client *Client) CreateMemoryStoreWithOptions(workspace *string, request *CreateMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomExtractionStrategies) {
		body["customExtractionStrategies"] = request.CustomExtractionStrategies
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ExtractionStrategies) {
		body["extractionStrategies"] = request.ExtractionStrategies
	}

	if !dara.IsNil(request.MemoryStoreName) {
		body["memoryStoreName"] = request.MemoryStoreName
	}

	if !dara.IsNil(request.ShortTermTtl) {
		body["shortTermTtl"] = request.ShortTermTtl
	}

	if !dara.IsNil(request.SourceType) {
		body["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.TraceSourceConfig) {
		body["traceSourceConfig"] = request.TraceSourceConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemoryStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建记忆库
//
// @param request - CreateMemoryStoreRequest
//
// @return CreateMemoryStoreResponse
func (client *Client) CreateMemoryStore(workspace *string, request *CreateMemoryStoreRequest) (_result *CreateMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryStoreResponse{}
	_body, _err := client.CreateMemoryStoreWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a Prometheus monitoring instance
//
// @param request - CreatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstanceWithOptions(request *CreatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArchiveDuration) {
		body["archiveDuration"] = request.ArchiveDuration
	}

	if !dara.IsNil(request.AuthFreeReadPolicy) {
		body["authFreeReadPolicy"] = request.AuthFreeReadPolicy
	}

	if !dara.IsNil(request.AuthFreeWritePolicy) {
		body["authFreeWritePolicy"] = request.AuthFreeWritePolicy
	}

	if !dara.IsNil(request.EnableAuthFreeRead) {
		body["enableAuthFreeRead"] = request.EnableAuthFreeRead
	}

	if !dara.IsNil(request.EnableAuthFreeWrite) {
		body["enableAuthFreeWrite"] = request.EnableAuthFreeWrite
	}

	if !dara.IsNil(request.EnableAuthToken) {
		body["enableAuthToken"] = request.EnableAuthToken
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.PrometheusInstanceName) {
		body["prometheusInstanceName"] = request.PrometheusInstanceName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.StorageDuration) {
		body["storageDuration"] = request.StorageDuration
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrometheusInstance"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrometheusInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a Prometheus monitoring instance
//
// @param request - CreatePrometheusInstanceRequest
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstance(request *CreatePrometheusInstanceRequest) (_result *CreatePrometheusInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePrometheusInstanceResponse{}
	_body, _err := client.CreatePrometheusInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Prometheus View
//
// Description:
//
// # Used to create a site monitoring task
//
// @param request - CreatePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusViewResponse
func (client *Client) CreatePrometheusViewWithOptions(request *CreatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthFreeReadPolicy) {
		body["authFreeReadPolicy"] = request.AuthFreeReadPolicy
	}

	if !dara.IsNil(request.EnableAuthFreeRead) {
		body["enableAuthFreeRead"] = request.EnableAuthFreeRead
	}

	if !dara.IsNil(request.EnableAuthToken) {
		body["enableAuthToken"] = request.EnableAuthToken
	}

	if !dara.IsNil(request.PrometheusInstances) {
		body["prometheusInstances"] = request.PrometheusInstances
	}

	if !dara.IsNil(request.PrometheusViewName) {
		body["prometheusViewName"] = request.PrometheusViewName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrometheusView"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-views"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrometheusViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Prometheus View
//
// Description:
//
// # Used to create a site monitoring task
//
// @param request - CreatePrometheusViewRequest
//
// @return CreatePrometheusViewResponse
func (client *Client) CreatePrometheusView(request *CreatePrometheusViewRequest) (_result *CreatePrometheusViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePrometheusViewResponse{}
	_body, _err := client.CreatePrometheusViewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Prometheus Monitoring Instance
//
// Description:
//
// Create a Prometheus monitoring virtual instance.
//
// @param request - CreatePrometheusVirtualInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusVirtualInstanceResponse
func (client *Client) CreatePrometheusVirtualInstanceWithOptions(request *CreatePrometheusVirtualInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusVirtualInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		body["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.TenantId) {
		body["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrometheusVirtualInstance"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/virtual-instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrometheusVirtualInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Prometheus Monitoring Instance
//
// Description:
//
// Create a Prometheus monitoring virtual instance.
//
// @param request - CreatePrometheusVirtualInstanceRequest
//
// @return CreatePrometheusVirtualInstanceResponse
func (client *Client) CreatePrometheusVirtualInstance(request *CreatePrometheusVirtualInstanceRequest) (_result *CreatePrometheusVirtualInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePrometheusVirtualInstanceResponse{}
	_body, _err := client.CreatePrometheusVirtualInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Service
//
// @param request - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(workspace *string, request *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Pid) {
		body["pid"] = request.Pid
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceName) {
		body["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceStatus) {
		body["serviceStatus"] = request.ServiceStatus
	}

	if !dara.IsNil(request.ServiceType) {
		body["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateService"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/service"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Service
//
// @param request - CreateServiceRequest
//
// @return CreateServiceResponse
func (client *Client) CreateService(workspace *string, request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用可观测实例
//
// @param request - CreateServiceObservabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceObservabilityResponse
func (client *Client) CreateServiceObservabilityWithOptions(workspace *string, _type *string, request *CreateServiceObservabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceObservabilityResponse, _err error) {
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
		Action:      dara.String("CreateServiceObservability"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/service-observability/" + dara.PercentEncode(dara.StringValue(_type))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceObservabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用可观测实例
//
// @param request - CreateServiceObservabilityRequest
//
// @return CreateServiceObservabilityResponse
func (client *Client) CreateServiceObservability(workspace *string, _type *string, request *CreateServiceObservabilityRequest) (_result *CreateServiceObservabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceObservabilityResponse{}
	_body, _err := client.CreateServiceObservabilityWithOptions(workspace, _type, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Ticket
//
// @param request - CreateTicketRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenExpirationTime) {
		query["accessTokenExpirationTime"] = request.AccessTokenExpirationTime
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["expirationTime"] = request.ExpirationTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tickets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Ticket
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Umodel configuration
//
// Description:
//
// # Create Umodel configuration in the specified workspace
//
// @param request - CreateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodelWithOptions(workspace *string, request *CreateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUmodelResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUmodel"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Umodel configuration
//
// Description:
//
// # Create Umodel configuration in the specified workspace
//
// @param request - CreateUmodelRequest
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodel(workspace *string, request *CreateUmodelRequest) (_result *CreateUmodelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUmodelResponse{}
	_body, _err := client.CreateUmodelWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete addon release information
//
// @param request - DeleteAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddonReleaseResponse
func (client *Client) DeleteAddonReleaseWithOptions(policyId *string, request *DeleteAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAddonReleaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["addonName"] = request.AddonName
	}

	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	if !dara.IsNil(request.ReleaseName) {
		query["releaseName"] = request.ReleaseName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAddonRelease"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/addon-releases"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAddonReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete addon release information
//
// @param request - DeleteAddonReleaseRequest
//
// @return DeleteAddonReleaseResponse
func (client *Client) DeleteAddonRelease(policyId *string, request *DeleteAddonReleaseRequest) (_result *DeleteAddonReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAddonReleaseResponse{}
	_body, _err := client.DeleteAddonReleaseWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Aggregation Task Group
//
// @param request - DeleteAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggTaskGroupResponse
func (client *Client) DeleteAggTaskGroupWithOptions(instanceId *string, groupId *string, request *DeleteAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAggTaskGroupResponse, _err error) {
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
		Action:      dara.String("DeleteAggTaskGroup"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/agg-task-groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAggTaskGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Aggregation Task Group
//
// @param request - DeleteAggTaskGroupRequest
//
// @return DeleteAggTaskGroupResponse
func (client *Client) DeleteAggTaskGroup(instanceId *string, groupId *string, request *DeleteAggTaskGroupRequest) (_result *DeleteAggTaskGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAggTaskGroupResponse{}
	_body, _err := client.DeleteAggTaskGroupWithOptions(instanceId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Webhook
//
// @param tmpReq - DeleteAlertWebhooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertWebhooksResponse
func (client *Client) DeleteAlertWebhooksWithOptions(tmpReq *DeleteAlertWebhooksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertWebhooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAlertWebhooksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WebhookIds) {
		request.WebhookIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WebhookIds, dara.String("webhookIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.WebhookIdsShrink) {
		query["webhookIds"] = request.WebhookIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertWebhooks"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webhooks"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertWebhooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Webhook
//
// @param request - DeleteAlertWebhooksRequest
//
// @return DeleteAlertWebhooksResponse
func (client *Client) DeleteAlertWebhooks(request *DeleteAlertWebhooksRequest) (_result *DeleteAlertWebhooksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertWebhooksResponse{}
	_body, _err := client.DeleteAlertWebhooksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除业务链路
//
// @param request - DeleteBizTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizTraceResponse
func (client *Client) DeleteBizTraceWithOptions(bizTraceId *string, request *DeleteBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBizTraceResponse, _err error) {
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
		Action:      dara.String("DeleteBizTrace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bizTrace/" + dara.PercentEncode(dara.StringValue(bizTraceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除业务链路
//
// @param request - DeleteBizTraceRequest
//
// @return DeleteBizTraceResponse
func (client *Client) DeleteBizTrace(bizTraceId *string, request *DeleteBizTraceRequest) (_result *DeleteBizTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBizTraceResponse{}
	_body, _err := client.DeleteBizTraceWithOptions(bizTraceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除云资源中心
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudResourceResponse
func (client *Client) DeleteCloudResourceWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCloudResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudResource"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cloudresource"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除云资源中心
//
// @return DeleteCloudResourceResponse
func (client *Client) DeleteCloudResource() (_result *DeleteCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCloudResourceResponse{}
	_body, _err := client.DeleteCloudResourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithOptions(workspace *string, datasetName *string, request *DeleteDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
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
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteDatasetRequest
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDataset(workspace *string, datasetName *string, request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(workspace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据投递任务
//
// @param request - DeleteDeliveryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeliveryTaskResponse
func (client *Client) DeleteDeliveryTaskWithOptions(taskId *string, request *DeleteDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDeliveryTaskResponse, _err error) {
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
		Action:      dara.String("DeleteDeliveryTask"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/delivery-task/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据投递任务
//
// @param request - DeleteDeliveryTaskRequest
//
// @return DeleteDeliveryTaskResponse
func (client *Client) DeleteDeliveryTask(taskId *string, request *DeleteDeliveryTaskRequest) (_result *DeleteDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDeliveryTaskResponse{}
	_body, _err := client.DeleteDeliveryTaskWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete EntityStore related storage
//
// @param request - DeleteEntityStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStoreWithOptions(workspaceName *string, request *DeleteEntityStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEntityStoreResponse, _err error) {
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
		Action:      dara.String("DeleteEntityStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/entitystore"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEntityStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete EntityStore related storage
//
// @param request - DeleteEntityStoreRequest
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStore(workspaceName *string, request *DeleteEntityStoreRequest) (_result *DeleteEntityStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEntityStoreResponse{}
	_body, _err := client.DeleteEntityStoreWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Access Center Policy
//
// @param request - DeleteIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntegrationPolicyResponse
func (client *Client) DeleteIntegrationPolicyWithOptions(policyId *string, request *DeleteIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIntegrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DeleteIntegrationPolicy"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIntegrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Access Center Policy
//
// @param request - DeleteIntegrationPolicyRequest
//
// @return DeleteIntegrationPolicyResponse
func (client *Client) DeleteIntegrationPolicy(policyId *string, request *DeleteIntegrationPolicyRequest) (_result *DeleteIntegrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIntegrationPolicyResponse{}
	_body, _err := client.DeleteIntegrationPolicyWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除记忆
//
// @param request - DeleteMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoriesResponse
func (client *Client) DeleteMemoriesWithOptions(workspace *string, memoryStoreName *string, request *DeleteMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AppId) {
		query["appId"] = request.AppId
	}

	if !dara.IsNil(request.RunId) {
		query["runId"] = request.RunId
	}

	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemories"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除记忆
//
// @param request - DeleteMemoriesRequest
//
// @return DeleteMemoriesResponse
func (client *Client) DeleteMemories(workspace *string, memoryStoreName *string, request *DeleteMemoriesRequest) (_result *DeleteMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoriesResponse{}
	_body, _err := client.DeleteMemoriesWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除记忆
//
// @param request - DeleteMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(workspace *string, memoryStoreName *string, memoryId *string, request *DeleteMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
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
		Action:      dara.String("DeleteMemory"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除记忆
//
// @param request - DeleteMemoryRequest
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(workspace *string, memoryStoreName *string, memoryId *string, request *DeleteMemoryRequest) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(workspace, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除记忆库
//
// @param request - DeleteMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryStoreResponse
func (client *Client) DeleteMemoryStoreWithOptions(workspace *string, memoryStoreName *string, request *DeleteMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryStoreResponse, _err error) {
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
		Action:      dara.String("DeleteMemoryStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除记忆库
//
// @param request - DeleteMemoryStoreRequest
//
// @return DeleteMemoryStoreResponse
func (client *Client) DeleteMemoryStore(workspace *string, memoryStoreName *string, request *DeleteMemoryStoreRequest) (_result *DeleteMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryStoreResponse{}
	_body, _err := client.DeleteMemoryStoreWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete prom instance
//
// Description:
//
// Delete a Prometheus instance.
//
// @param request - DeletePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusInstanceResponse
func (client *Client) DeletePrometheusInstanceWithOptions(prometheusInstanceId *string, request *DeletePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusInstanceResponse, _err error) {
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
		Action:      dara.String("DeletePrometheusInstance"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(prometheusInstanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete prom instance
//
// Description:
//
// Delete a Prometheus instance.
//
// @param request - DeletePrometheusInstanceRequest
//
// @return DeletePrometheusInstanceResponse
func (client *Client) DeletePrometheusInstance(prometheusInstanceId *string, request *DeletePrometheusInstanceRequest) (_result *DeletePrometheusInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePrometheusInstanceResponse{}
	_body, _err := client.DeletePrometheusInstanceWithOptions(prometheusInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete prometheus view instance
//
// Description:
//
// Delete prometheus view instance.
//
// @param request - DeletePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusViewResponse
func (client *Client) DeletePrometheusViewWithOptions(prometheusViewId *string, request *DeletePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusViewResponse, _err error) {
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
		Action:      dara.String("DeletePrometheusView"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-views/" + dara.PercentEncode(dara.StringValue(prometheusViewId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete prometheus view instance
//
// Description:
//
// Delete prometheus view instance.
//
// @param request - DeletePrometheusViewRequest
//
// @return DeletePrometheusViewResponse
func (client *Client) DeletePrometheusView(prometheusViewId *string, request *DeletePrometheusViewRequest) (_result *DeletePrometheusViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePrometheusViewResponse{}
	_body, _err := client.DeletePrometheusViewWithOptions(prometheusViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 Prometheus 虚拟实例
//
// @param request - DeletePrometheusVirtualInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusVirtualInstanceResponse
func (client *Client) DeletePrometheusVirtualInstanceWithOptions(prometheusInstanceId *string, request *DeletePrometheusVirtualInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusVirtualInstanceResponse, _err error) {
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
		Action:      dara.String("DeletePrometheusVirtualInstance"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/virtual-instances/" + dara.PercentEncode(dara.StringValue(prometheusInstanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrometheusVirtualInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 Prometheus 虚拟实例
//
// @param request - DeletePrometheusVirtualInstanceRequest
//
// @return DeletePrometheusVirtualInstanceResponse
func (client *Client) DeletePrometheusVirtualInstance(prometheusInstanceId *string, request *DeletePrometheusVirtualInstanceRequest) (_result *DeletePrometheusVirtualInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePrometheusVirtualInstanceResponse{}
	_body, _err := client.DeletePrometheusVirtualInstanceWithOptions(prometheusInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Service
//
// @param request - DeleteServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithOptions(workspace *string, serviceId *string, request *DeleteServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
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
		Action:      dara.String("DeleteService"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/service/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Service
//
// @param request - DeleteServiceRequest
//
// @return DeleteServiceResponse
func (client *Client) DeleteService(workspace *string, serviceId *string, request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(workspace, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Umodel configuration information
//
// Description:
//
// # Delete the Umodel under the specified workspace
//
// @param request - DeleteUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodelWithOptions(workspace *string, request *DeleteUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelResponse, _err error) {
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
		Action:      dara.String("DeleteUmodel"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Umodel configuration information
//
// Description:
//
// # Delete the Umodel under the specified workspace
//
// @param request - DeleteUmodelRequest
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodel(workspace *string, request *DeleteUmodelRequest) (_result *DeleteUmodelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUmodelResponse{}
	_body, _err := client.DeleteUmodelWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Umodel配置信息
//
// @param request - DeleteUmodelCommonSchemaRefRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelCommonSchemaRefResponse
func (client *Client) DeleteUmodelCommonSchemaRefWithOptions(workspace *string, request *DeleteUmodelCommonSchemaRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelCommonSchemaRefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Group) {
		query["group"] = request.Group
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUmodelCommonSchemaRef"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel/common-schema-ref"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUmodelCommonSchemaRefResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Umodel配置信息
//
// @param request - DeleteUmodelCommonSchemaRefRequest
//
// @return DeleteUmodelCommonSchemaRefResponse
func (client *Client) DeleteUmodelCommonSchemaRef(workspace *string, request *DeleteUmodelCommonSchemaRefRequest) (_result *DeleteUmodelCommonSchemaRefResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUmodelCommonSchemaRefResponse{}
	_body, _err := client.DeleteUmodelCommonSchemaRefWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Umodel Elements
//
// Description:
//
// # Delete the Umodel Data under a specified workspace
//
// @param request - DeleteUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelDataResponse
func (client *Client) DeleteUmodelDataWithOptions(workspace *string, request *DeleteUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["domain"] = request.Domain
	}

	if !dara.IsNil(request.Kind) {
		query["kind"] = request.Kind
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUmodelData"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel/data"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUmodelDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Umodel Elements
//
// Description:
//
// # Delete the Umodel Data under a specified workspace
//
// @param request - DeleteUmodelDataRequest
//
// @return DeleteUmodelDataResponse
func (client *Client) DeleteUmodelData(workspace *string, request *DeleteUmodelDataRequest) (_result *DeleteUmodelDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUmodelDataResponse{}
	_body, _err := client.DeleteUmodelDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Workspace
//
// @param request - DeleteWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceName *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspaceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Workspace
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceName *string, request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询地域信息列表
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询地域信息列表
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行查询语句
//
// @param request - ExecuteQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteQueryResponse
func (client *Client) ExecuteQueryWithOptions(workspace *string, datasetName *string, request *ExecuteQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteQuery"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName)) + "/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行查询语句
//
// @param request - ExecuteQueryRequest
//
// @return ExecuteQueryResponse
func (client *Client) ExecuteQuery(workspace *string, datasetName *string, request *ExecuteQueryRequest) (_result *ExecuteQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteQueryResponse{}
	_body, _err := client.ExecuteQueryWithOptions(workspace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插件详情(Addon)
//
// @param request - GetAddonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonResponse
func (client *Client) GetAddonWithOptions(addonName *string, request *GetAddonRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAddon"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addons/" + dara.PercentEncode(dara.StringValue(addonName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插件详情(Addon)
//
// @param request - GetAddonRequest
//
// @return GetAddonResponse
func (client *Client) GetAddon(addonName *string, request *GetAddonRequest) (_result *GetAddonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddonResponse{}
	_body, _err := client.GetAddonWithOptions(addonName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插件schema详情(Addon)
//
// @param request - GetAddonCodeTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonCodeTemplateResponse
func (client *Client) GetAddonCodeTemplateWithOptions(addonName *string, request *GetAddonCodeTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonCodeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentType) {
		query["environmentType"] = request.EnvironmentType
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAddonCodeTemplate"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addons/" + dara.PercentEncode(dara.StringValue(addonName)) + "/alert-code-template"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAddonCodeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插件schema详情(Addon)
//
// @param request - GetAddonCodeTemplateRequest
//
// @return GetAddonCodeTemplateResponse
func (client *Client) GetAddonCodeTemplate(addonName *string, request *GetAddonCodeTemplateRequest) (_result *GetAddonCodeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddonCodeTemplateResponse{}
	_body, _err := client.GetAddonCodeTemplateWithOptions(addonName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check addon release (view connection status)
//
// @param request - GetAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonReleaseResponse
func (client *Client) GetAddonReleaseWithOptions(releaseName *string, policyId *string, request *GetAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonReleaseResponse, _err error) {
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
		Action:      dara.String("GetAddonRelease"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/addon-releases/" + dara.PercentEncode(dara.StringValue(releaseName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAddonReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check addon release (view connection status)
//
// @param request - GetAddonReleaseRequest
//
// @return GetAddonReleaseResponse
func (client *Client) GetAddonRelease(releaseName *string, policyId *string, request *GetAddonReleaseRequest) (_result *GetAddonReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddonReleaseResponse{}
	_body, _err := client.GetAddonReleaseWithOptions(releaseName, policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插件schema详情(Addon)
//
// @param request - GetAddonSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonSchemaResponse
func (client *Client) GetAddonSchemaWithOptions(addonName *string, request *GetAddonSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EnvironmentType) {
		query["environmentType"] = request.EnvironmentType
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAddonSchema"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addons/" + dara.PercentEncode(dara.StringValue(addonName)) + "/schema"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAddonSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插件schema详情(Addon)
//
// @param request - GetAddonSchemaRequest
//
// @return GetAddonSchemaResponse
func (client *Client) GetAddonSchema(addonName *string, request *GetAddonSchemaRequest) (_result *GetAddonSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddonSchemaResponse{}
	_body, _err := client.GetAddonSchemaWithOptions(addonName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Describes the aggregation task group
//
// @param request - GetAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggTaskGroupResponse
func (client *Client) GetAggTaskGroupWithOptions(instanceId *string, groupId *string, request *GetAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAggTaskGroupResponse, _err error) {
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
		Action:      dara.String("GetAggTaskGroup"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/agg-task-groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAggTaskGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Describes the aggregation task group
//
// @param request - GetAggTaskGroupRequest
//
// @return GetAggTaskGroupResponse
func (client *Client) GetAggTaskGroup(instanceId *string, groupId *string, request *GetAggTaskGroupRequest) (_result *GetAggTaskGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAggTaskGroupResponse{}
	_body, _err := client.GetAggTaskGroupWithOptions(instanceId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询业务链路
//
// @param request - GetBizTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizTraceResponse
func (client *Client) GetBizTraceWithOptions(bizTraceId *string, request *GetBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBizTraceResponse, _err error) {
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
		Action:      dara.String("GetBizTrace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bizTrace/" + dara.PercentEncode(dara.StringValue(bizTraceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询业务链路
//
// @param request - GetBizTraceRequest
//
// @return GetBizTraceResponse
func (client *Client) GetBizTrace(bizTraceId *string, request *GetBizTraceRequest) (_result *GetBizTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBizTraceResponse{}
	_body, _err := client.GetBizTraceWithOptions(bizTraceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云资源中心
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudResourceResponse
func (client *Client) GetCloudResourceWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCloudResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCloudResource"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cloudresource"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云资源中心
//
// @return GetCloudResourceResponse
func (client *Client) GetCloudResource() (_result *GetCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCloudResourceResponse{}
	_body, _err := client.GetCloudResourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云资源中心数据
//
// @param request - GetCloudResourceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudResourceDataResponse
func (client *Client) GetCloudResourceDataWithOptions(request *GetCloudResourceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCloudResourceDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCloudResourceData"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cloudresource/data"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudResourceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云资源中心数据
//
// @param request - GetCloudResourceDataRequest
//
// @return GetCloudResourceDataResponse
func (client *Client) GetCloudResourceData(request *GetCloudResourceDataRequest) (_result *GetCloudResourceDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCloudResourceDataResponse{}
	_body, _err := client.GetCloudResourceDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云监控开通状态
//
// @param request - GetCmsServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCmsServiceResponse
func (client *Client) GetCmsServiceWithOptions(request *GetCmsServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCmsServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["product"] = request.Product
	}

	if !dara.IsNil(request.Service) {
		query["service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCmsService"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cmsservice"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCmsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云监控开通状态
//
// @param request - GetCmsServiceRequest
//
// @return GetCmsServiceResponse
func (client *Client) GetCmsService(request *GetCmsServiceRequest) (_result *GetCmsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCmsServiceResponse{}
	_body, _err := client.GetCmsServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据集
//
// @param request - GetDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithOptions(workspace *string, datasetName *string, request *GetDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
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
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集
//
// @param request - GetDatasetRequest
//
// @return GetDatasetResponse
func (client *Client) GetDataset(workspace *string, datasetName *string, request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(workspace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据投递任务详情
//
// @param request - GetDeliveryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryTaskResponse
func (client *Client) GetDeliveryTaskWithOptions(taskId *string, request *GetDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeliveryTaskResponse, _err error) {
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
		Action:      dara.String("GetDeliveryTask"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/delivery-task/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据投递任务详情
//
// @param request - GetDeliveryTaskRequest
//
// @return GetDeliveryTaskResponse
func (client *Client) GetDeliveryTask(taskId *string, request *GetDeliveryTaskRequest) (_result *GetDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeliveryTaskResponse{}
	_body, _err := client.GetDeliveryTaskWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get EntityStore related storage information
//
// @param request - GetEntityStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStoreWithOptions(workspaceName *string, request *GetEntityStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEntityStoreResponse, _err error) {
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
		Action:      dara.String("GetEntityStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/entitystore"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEntityStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get EntityStore related storage information
//
// @param request - GetEntityStoreRequest
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStore(workspaceName *string, request *GetEntityStoreRequest) (_result *GetEntityStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEntityStoreResponse{}
	_body, _err := client.GetEntityStoreWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the entity and relationship data under a specified Workspace, returning the entity data within a certain time range (the returned result is transmitted after compression).
//
// @param request - GetEntityStoreDataRequest
//
// @param headers - GetEntityStoreDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreDataResponse
func (client *Client) GetEntityStoreDataWithOptions(workspace *string, request *GetEntityStoreDataRequest, headers *GetEntityStoreDataHeaders, runtime *dara.RuntimeOptions) (_result *GetEntityStoreDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		body["from"] = request.From
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.To) {
		body["to"] = request.To
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AcceptEncoding) {
		realHeaders["acceptEncoding"] = dara.String(dara.ToString(dara.StringValue(headers.AcceptEncoding)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEntityStoreData"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/entitiesAndRelations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEntityStoreDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the entity and relationship data under a specified Workspace, returning the entity data within a certain time range (the returned result is transmitted after compression).
//
// @param request - GetEntityStoreDataRequest
//
// @return GetEntityStoreDataResponse
func (client *Client) GetEntityStoreData(workspace *string, request *GetEntityStoreDataRequest) (_result *GetEntityStoreDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetEntityStoreDataHeaders{}
	_result = &GetEntityStoreDataResponse{}
	_body, _err := client.GetEntityStoreDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of access center policies
//
// @param request - GetIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationPolicyResponse
func (client *Client) GetIntegrationPolicyWithOptions(policyId *string, request *GetIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIntegrationPolicyResponse, _err error) {
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
		Action:      dara.String("GetIntegrationPolicy"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntegrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of access center policies
//
// @param request - GetIntegrationPolicyRequest
//
// @return GetIntegrationPolicyResponse
func (client *Client) GetIntegrationPolicy(policyId *string, request *GetIntegrationPolicyRequest) (_result *GetIntegrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIntegrationPolicyResponse{}
	_body, _err := client.GetIntegrationPolicyWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入中心在CS的版本
//
// @param request - GetIntegrationVersionForCSRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationVersionForCSResponse
func (client *Client) GetIntegrationVersionForCSWithOptions(request *GetIntegrationVersionForCSRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIntegrationVersionForCSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["clusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterType) {
		query["clusterType"] = request.ClusterType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntegrationVersionForCS"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-version/cs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntegrationVersionForCSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入中心在CS的版本
//
// @param request - GetIntegrationVersionForCSRequest
//
// @return GetIntegrationVersionForCSResponse
func (client *Client) GetIntegrationVersionForCS(request *GetIntegrationVersionForCSRequest) (_result *GetIntegrationVersionForCSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIntegrationVersionForCSResponse{}
	_body, _err := client.GetIntegrationVersionForCSWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询全部记忆
//
// @param request - GetMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoriesResponse
func (client *Client) GetMemoriesWithOptions(workspace *string, memoryStoreName *string, request *GetMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Filters) {
		body["filters"] = request.Filters
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemories"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全部记忆
//
// @param request - GetMemoriesRequest
//
// @return GetMemoriesResponse
func (client *Client) GetMemories(workspace *string, memoryStoreName *string, request *GetMemoriesRequest) (_result *GetMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoriesResponse{}
	_body, _err := client.GetMemoriesWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询记忆
//
// @param request - GetMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(workspace *string, memoryStoreName *string, memoryId *string, request *GetMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
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
		Action:      dara.String("GetMemory"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆
//
// @param request - GetMemoryRequest
//
// @return GetMemoryResponse
func (client *Client) GetMemory(workspace *string, memoryStoreName *string, memoryId *string, request *GetMemoryRequest) (_result *GetMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(workspace, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询记忆历史记录
//
// @param request - GetMemoryHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryHistoryResponse
func (client *Client) GetMemoryHistoryWithOptions(workspace *string, memoryStoreName *string, memoryId *string, request *GetMemoryHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryHistoryResponse, _err error) {
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
		Action:      dara.String("GetMemoryHistory"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory/" + dara.PercentEncode(dara.StringValue(memoryId)) + "/history"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆历史记录
//
// @param request - GetMemoryHistoryRequest
//
// @return GetMemoryHistoryResponse
func (client *Client) GetMemoryHistory(workspace *string, memoryStoreName *string, memoryId *string, request *GetMemoryHistoryRequest) (_result *GetMemoryHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryHistoryResponse{}
	_body, _err := client.GetMemoryHistoryWithOptions(workspace, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询记忆库
//
// @param request - GetMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryStoreResponse
func (client *Client) GetMemoryStoreWithOptions(workspace *string, memoryStoreName *string, request *GetMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryStoreResponse, _err error) {
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
		Action:      dara.String("GetMemoryStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆库
//
// @param request - GetMemoryStoreRequest
//
// @return GetMemoryStoreResponse
func (client *Client) GetMemoryStore(workspace *string, memoryStoreName *string, request *GetMemoryStoreRequest) (_result *GetMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryStoreResponse{}
	_body, _err := client.GetMemoryStoreWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the instance in a specified environment
//
// Description:
//
// Retrieve details of a Prometheus instance.
//
// @param request - GetPrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusInstanceResponse
func (client *Client) GetPrometheusInstanceWithOptions(prometheusInstanceId *string, request *GetPrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrometheusInstance"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(prometheusInstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the instance in a specified environment
//
// Description:
//
// Retrieve details of a Prometheus instance.
//
// @param request - GetPrometheusInstanceRequest
//
// @return GetPrometheusInstanceResponse
func (client *Client) GetPrometheusInstance(prometheusInstanceId *string, request *GetPrometheusInstanceRequest) (_result *GetPrometheusInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPrometheusInstanceResponse{}
	_body, _err := client.GetPrometheusInstanceWithOptions(prometheusInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定环境实例
//
// @param request - GetPrometheusUserSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusUserSettingResponse
func (client *Client) GetPrometheusUserSettingWithOptions(request *GetPrometheusUserSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusUserSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrometheusUserSetting"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-user-setting"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusUserSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定环境实例
//
// @param request - GetPrometheusUserSettingRequest
//
// @return GetPrometheusUserSettingResponse
func (client *Client) GetPrometheusUserSetting(request *GetPrometheusUserSettingRequest) (_result *GetPrometheusUserSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPrometheusUserSettingResponse{}
	_body, _err := client.GetPrometheusUserSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a specified Prometheus view instance
//
// Description:
//
// Query a specified Prometheus view instance.
//
// @param request - GetPrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusViewResponse
func (client *Client) GetPrometheusViewWithOptions(prometheusViewId *string, request *GetPrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrometheusView"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-views/" + dara.PercentEncode(dara.StringValue(prometheusViewId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrometheusViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a specified Prometheus view instance
//
// Description:
//
// Query a specified Prometheus view instance.
//
// @param request - GetPrometheusViewRequest
//
// @return GetPrometheusViewResponse
func (client *Client) GetPrometheusView(prometheusViewId *string, request *GetPrometheusViewRequest) (_result *GetPrometheusViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPrometheusViewResponse{}
	_body, _err := client.GetPrometheusViewWithOptions(prometheusViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Service
//
// @param request - GetServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(workspace *string, serviceId *string, request *GetServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
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
		Action:      dara.String("GetService"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/service/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Service
//
// @param request - GetServiceRequest
//
// @return GetServiceResponse
func (client *Client) GetService(workspace *string, serviceId *string, request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(workspace, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Application Observability Instance
//
// @param request - GetServiceObservabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceObservabilityResponse
func (client *Client) GetServiceObservabilityWithOptions(workspace *string, _type *string, request *GetServiceObservabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceObservabilityResponse, _err error) {
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
		Action:      dara.String("GetServiceObservability"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/service-observability/" + dara.PercentEncode(dara.StringValue(_type))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceObservabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Application Observability Instance
//
// @param request - GetServiceObservabilityRequest
//
// @return GetServiceObservabilityResponse
func (client *Client) GetServiceObservability(workspace *string, _type *string, request *GetServiceObservabilityRequest) (_result *GetServiceObservabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceObservabilityResponse{}
	_body, _err := client.GetServiceObservabilityWithOptions(workspace, _type, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Umodel configuration information
//
// Description:
//
// # Get Umodel configuration information
//
// @param request - GetUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelResponse
func (client *Client) GetUmodelWithOptions(workspace *string, request *GetUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelResponse, _err error) {
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
		Action:      dara.String("GetUmodel"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Umodel configuration information
//
// Description:
//
// # Get Umodel configuration information
//
// @param request - GetUmodelRequest
//
// @return GetUmodelResponse
func (client *Client) GetUmodel(workspace *string, request *GetUmodelRequest) (_result *GetUmodelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUmodelResponse{}
	_body, _err := client.GetUmodelWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Umodel配置信息
//
// @param request - GetUmodelCommonSchemaRefRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelCommonSchemaRefResponse
func (client *Client) GetUmodelCommonSchemaRefWithOptions(workspace *string, request *GetUmodelCommonSchemaRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelCommonSchemaRefResponse, _err error) {
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
		Action:      dara.String("GetUmodelCommonSchemaRef"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel/common-schema-ref"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUmodelCommonSchemaRefResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Umodel配置信息
//
// @param request - GetUmodelCommonSchemaRefRequest
//
// @return GetUmodelCommonSchemaRefResponse
func (client *Client) GetUmodelCommonSchemaRef(workspace *string, request *GetUmodelCommonSchemaRefRequest) (_result *GetUmodelCommonSchemaRefResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUmodelCommonSchemaRefResponse{}
	_body, _err := client.GetUmodelCommonSchemaRefWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve associated Umodel graph data
//
// Description:
//
// # Find Umodel
//
// @param request - GetUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelDataResponse
func (client *Client) GetUmodelDataWithOptions(workspace *string, request *GetUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Method) {
		query["method"] = request.Method
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUmodelData"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel/graph"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUmodelDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve associated Umodel graph data
//
// Description:
//
// # Find Umodel
//
// @param request - GetUmodelDataRequest
//
// @return GetUmodelDataResponse
func (client *Client) GetUmodelData(workspace *string, request *GetUmodelDataRequest) (_result *GetUmodelDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUmodelDataResponse{}
	_body, _err := client.GetUmodelDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Workspace
//
// @param request - GetWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceName *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspaceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Workspace
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceName *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of addon releases
//
// Description:
//
// # Query the list of access configurations
//
// @param request - ListAddonReleasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonReleasesResponse
func (client *Client) ListAddonReleasesWithOptions(policyId *string, request *ListAddonReleasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonReleasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["addonName"] = request.AddonName
	}

	if !dara.IsNil(request.ParentAddonReleaseId) {
		query["parentAddonReleaseId"] = request.ParentAddonReleaseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddonReleases"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/addon-releases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddonReleasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of addon releases
//
// Description:
//
// # Query the list of access configurations
//
// @param request - ListAddonReleasesRequest
//
// @return ListAddonReleasesResponse
func (client *Client) ListAddonReleases(policyId *string, request *ListAddonReleasesRequest) (_result *ListAddonReleasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAddonReleasesResponse{}
	_body, _err := client.ListAddonReleasesWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新版接入中心产品列表(分组)
//
// @param request - ListAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithOptions(request *ListAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.Regexp) {
		query["regexp"] = request.Regexp
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddons"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addons"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddonsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新版接入中心产品列表(分组)
//
// @param request - ListAddonsRequest
//
// @return ListAddonsResponse
func (client *Client) ListAddons(request *ListAddonsRequest) (_result *ListAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAddonsResponse{}
	_body, _err := client.ListAddonsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Aggregation Task Groups
//
// @param tmpReq - ListAggTaskGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggTaskGroupsResponse
func (client *Client) ListAggTaskGroupsWithOptions(instanceId *string, tmpReq *ListAggTaskGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAggTaskGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAggTaskGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterAggTaskGroupIds) {
		query["filterAggTaskGroupIds"] = request.FilterAggTaskGroupIds
	}

	if !dara.IsNil(request.FilterAggTaskGroupNames) {
		query["filterAggTaskGroupNames"] = request.FilterAggTaskGroupNames
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetPrometheusId) {
		query["targetPrometheusId"] = request.TargetPrometheusId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAggTaskGroups"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/agg-task-groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAggTaskGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Aggregation Task Groups
//
// @param request - ListAggTaskGroupsRequest
//
// @return ListAggTaskGroupsResponse
func (client *Client) ListAggTaskGroups(instanceId *string, request *ListAggTaskGroupsRequest) (_result *ListAggTaskGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAggTaskGroupsResponse{}
	_body, _err := client.ListAggTaskGroupsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Alert Actions
//
// @param tmpReq - ListAlertActionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActionsWithOptions(tmpReq *ListAlertActionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertActionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertActionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertActionIds) {
		request.AlertActionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertActionIds, dara.String("alertActionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertActionIdsShrink) {
		query["alertActionIds"] = request.AlertActionIdsShrink
	}

	if !dara.IsNil(request.AlertActionName) {
		query["alertActionName"] = request.AlertActionName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertActions"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/alertActions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Alert Actions
//
// @param request - ListAlertActionsRequest
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActions(request *ListAlertActionsRequest) (_result *ListAlertActionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertActionsResponse{}
	_body, _err := client.ListAlertActionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询机器人
//
// @param tmpReq - ListAlertRobotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertRobotsResponse
func (client *Client) ListAlertRobotsWithOptions(tmpReq *ListAlertRobotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertRobotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertRobotsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RobotIds) {
		request.RobotIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RobotIds, dara.String("robotIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("types"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RobotIdsShrink) {
		query["robotIds"] = request.RobotIdsShrink
	}

	if !dara.IsNil(request.TypesShrink) {
		query["types"] = request.TypesShrink
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertRobots"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/robots"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertRobotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机器人
//
// @param request - ListAlertRobotsRequest
//
// @return ListAlertRobotsResponse
func (client *Client) ListAlertRobots(request *ListAlertRobotsRequest) (_result *ListAlertRobotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertRobotsResponse{}
	_body, _err := client.ListAlertRobotsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Webhook
//
// @param tmpReq - ListAlertWebhooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertWebhooksResponse
func (client *Client) ListAlertWebhooksWithOptions(tmpReq *ListAlertWebhooksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertWebhooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertWebhooksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WebhookIds) {
		request.WebhookIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WebhookIds, dara.String("webhookIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WebhookIdsShrink) {
		query["webhookIds"] = request.WebhookIdsShrink
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertWebhooks"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webhooks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertWebhooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Webhook
//
// @param request - ListAlertWebhooksRequest
//
// @return ListAlertWebhooksResponse
func (client *Client) ListAlertWebhooks(request *ListAlertWebhooksRequest) (_result *ListAlertWebhooksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertWebhooksResponse{}
	_body, _err := client.ListAlertWebhooksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 业务链路列表
//
// @param request - ListBizTracesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizTracesResponse
func (client *Client) ListBizTracesWithOptions(request *ListBizTracesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBizTracesResponse, _err error) {
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

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizTraces"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bizTraces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 业务链路列表
//
// @param request - ListBizTracesRequest
//
// @return ListBizTracesResponse
func (client *Client) ListBizTraces(request *ListBizTracesRequest) (_result *ListBizTracesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBizTracesResponse{}
	_body, _err := client.ListBizTracesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询联系人组
//
// @param tmpReq - ListContactGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactGroupsResponse
func (client *Client) ListContactGroupsWithOptions(tmpReq *ListContactGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContactGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListContactGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactGroupIds) {
		request.ContactGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactGroupIds, dara.String("contactGroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupIdsShrink) {
		query["contactGroupIds"] = request.ContactGroupIdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContactGroups"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/contactGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContactGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人组
//
// @param request - ListContactGroupsRequest
//
// @return ListContactGroupsResponse
func (client *Client) ListContactGroups(request *ListContactGroupsRequest) (_result *ListContactGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListContactGroupsResponse{}
	_body, _err := client.ListContactGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询联系人
//
// @param tmpReq - ListContactsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactsResponse
func (client *Client) ListContactsWithOptions(tmpReq *ListContactsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListContactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactIds) {
		request.ContactIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactIds, dara.String("contactIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIdsShrink) {
		query["contactIds"] = request.ContactIdsShrink
	}

	if !dara.IsNil(request.Email) {
		query["email"] = request.Email
	}

	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Phone) {
		query["phone"] = request.Phone
	}

	if !dara.IsNil(request.QueryUngroupedContacts) {
		query["queryUngroupedContacts"] = request.QueryUngroupedContacts
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContacts"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/contact"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人
//
// @param request - ListContactsRequest
//
// @return ListContactsResponse
func (client *Client) ListContacts(request *ListContactsRequest) (_result *ListContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListContactsResponse{}
	_body, _err := client.ListContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据集列表
//
// @param request - ListDatasetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithOptions(workspace *string, request *ListDatasetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["datasetName"] = request.DatasetName
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
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/dataset"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集列表
//
// @param request - ListDatasetsRequest
//
// @return ListDatasetsResponse
func (client *Client) ListDatasets(workspace *string, request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据投递任务列表
//
// @param tmpReq - ListDeliveryTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryTasksResponse
func (client *Client) ListDeliveryTasksWithOptions(tmpReq *ListDeliveryTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeliveryTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeliveryTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWords) {
		query["keyWords"] = request.KeyWords
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeliveryTasks"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/delivery-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeliveryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据投递任务列表
//
// @param request - ListDeliveryTasksRequest
//
// @return ListDeliveryTasksResponse
func (client *Client) ListDeliveryTasks(request *ListDeliveryTasksRequest) (_result *ListDeliveryTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDeliveryTasksResponse{}
	_body, _err := client.ListDeliveryTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Access Center Policy List Information
//
// Description:
//
// # Query integration list
//
// @param tmpReq - ListIntegrationPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPoliciesResponse
func (client *Client) ListIntegrationPoliciesWithOptions(tmpReq *ListIntegrationPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListIntegrationPoliciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["addonName"] = request.AddonName
	}

	if !dara.IsNil(request.BindResourceId) {
		query["bindResourceId"] = request.BindResourceId
	}

	if !dara.IsNil(request.EntityGroupIds) {
		query["entityGroupIds"] = request.EntityGroupIds
	}

	if !dara.IsNil(request.FilterRegionIds) {
		query["filterRegionIds"] = request.FilterRegionIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicyId) {
		query["policyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["policyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["policyType"] = request.PolicyType
	}

	if !dara.IsNil(request.PrometheusInstanceId) {
		query["prometheusInstanceId"] = request.PrometheusInstanceId
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicies"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Access Center Policy List Information
//
// Description:
//
// # Query integration list
//
// @param request - ListIntegrationPoliciesRequest
//
// @return ListIntegrationPoliciesResponse
func (client *Client) ListIntegrationPolicies(request *ListIntegrationPoliciesRequest) (_result *ListIntegrationPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPoliciesResponse{}
	_body, _err := client.ListIntegrationPoliciesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 策略addon列表
//
// @param request - ListIntegrationPolicyAddonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyAddonsResponse
func (client *Client) ListIntegrationPolicyAddonsWithOptions(policyId *string, request *ListIntegrationPolicyAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyAddonsResponse, _err error) {
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
		Action:      dara.String("ListIntegrationPolicyAddons"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/addons"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyAddonsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 策略addon列表
//
// @param request - ListIntegrationPolicyAddonsRequest
//
// @return ListIntegrationPolicyAddonsResponse
func (client *Client) ListIntegrationPolicyAddons(policyId *string, request *ListIntegrationPolicyAddonsRequest) (_result *ListIntegrationPolicyAddonsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyAddonsResponse{}
	_body, _err := client.ListIntegrationPolicyAddonsWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取接入中心策略的存储要求信息
//
// @param request - ListIntegrationPolicyCollectorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyCollectorsResponse
func (client *Client) ListIntegrationPolicyCollectorsWithOptions(policyId *string, request *ListIntegrationPolicyCollectorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyCollectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonReleaseName) {
		query["addonReleaseName"] = request.AddonReleaseName
	}

	if !dara.IsNil(request.CollectorType) {
		query["collectorType"] = request.CollectorType
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicyCollectors"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/collectors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyCollectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取接入中心策略的存储要求信息
//
// @param request - ListIntegrationPolicyCollectorsRequest
//
// @return ListIntegrationPolicyCollectorsResponse
func (client *Client) ListIntegrationPolicyCollectors(policyId *string, request *ListIntegrationPolicyCollectorsRequest) (_result *ListIntegrationPolicyCollectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyCollectorsResponse{}
	_body, _err := client.ListIntegrationPolicyCollectorsWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get storage requirement information for the access center policy
//
// @param request - ListIntegrationPolicyCustomScrapeJobRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyCustomScrapeJobRulesResponse
func (client *Client) ListIntegrationPolicyCustomScrapeJobRulesWithOptions(policyId *string, request *ListIntegrationPolicyCustomScrapeJobRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyCustomScrapeJobRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonReleaseName) {
		query["addonReleaseName"] = request.AddonReleaseName
	}

	if !dara.IsNil(request.EncryptYaml) {
		query["encryptYaml"] = request.EncryptYaml
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicyCustomScrapeJobRules"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/custom-scrape-job-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyCustomScrapeJobRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get storage requirement information for the access center policy
//
// @param request - ListIntegrationPolicyCustomScrapeJobRulesRequest
//
// @return ListIntegrationPolicyCustomScrapeJobRulesResponse
func (client *Client) ListIntegrationPolicyCustomScrapeJobRules(policyId *string, request *ListIntegrationPolicyCustomScrapeJobRulesRequest) (_result *ListIntegrationPolicyCustomScrapeJobRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyCustomScrapeJobRulesResponse{}
	_body, _err := client.ListIntegrationPolicyCustomScrapeJobRulesWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Policy Dashboard List
//
// Description:
//
// This article provides an example of querying the alarm template list. The result shows that there are 2 alarm templates in the list, which are `ECS_Template1` and `ECS_Template2`.
//
// @param request - ListIntegrationPolicyDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyDashboardsResponse
func (client *Client) ListIntegrationPolicyDashboardsWithOptions(policyId *string, request *ListIntegrationPolicyDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyDashboardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["addonName"] = request.AddonName
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicyDashboards"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/dashboards"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Policy Dashboard List
//
// Description:
//
// This article provides an example of querying the alarm template list. The result shows that there are 2 alarm templates in the list, which are `ECS_Template1` and `ECS_Template2`.
//
// @param request - ListIntegrationPolicyDashboardsRequest
//
// @return ListIntegrationPolicyDashboardsResponse
func (client *Client) ListIntegrationPolicyDashboards(policyId *string, request *ListIntegrationPolicyDashboardsRequest) (_result *ListIntegrationPolicyDashboardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyDashboardsResponse{}
	_body, _err := client.ListIntegrationPolicyDashboardsWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get PodMonitor Resources of Access Center Policy
//
// Description:
//
// This article provides an example to query the alarm template list. The result shows that there are 2 alarm templates in the alarm template list, which are `ECS_Template1` and `ECS_Template2`.
//
// @param request - ListIntegrationPolicyPodMonitorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyPodMonitorsResponse
func (client *Client) ListIntegrationPolicyPodMonitorsWithOptions(policyId *string, request *ListIntegrationPolicyPodMonitorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyPodMonitorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonReleaseName) {
		query["addonReleaseName"] = request.AddonReleaseName
	}

	if !dara.IsNil(request.EncryptYaml) {
		query["encryptYaml"] = request.EncryptYaml
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicyPodMonitors"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/pod-monitors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyPodMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get PodMonitor Resources of Access Center Policy
//
// Description:
//
// This article provides an example to query the alarm template list. The result shows that there are 2 alarm templates in the alarm template list, which are `ECS_Template1` and `ECS_Template2`.
//
// @param request - ListIntegrationPolicyPodMonitorsRequest
//
// @return ListIntegrationPolicyPodMonitorsResponse
func (client *Client) ListIntegrationPolicyPodMonitors(policyId *string, request *ListIntegrationPolicyPodMonitorsRequest) (_result *ListIntegrationPolicyPodMonitorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyPodMonitorsResponse{}
	_body, _err := client.ListIntegrationPolicyPodMonitorsWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取接入中心策略的存储要求信息
//
// @param request - ListIntegrationPolicyServiceMonitorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyServiceMonitorsResponse
func (client *Client) ListIntegrationPolicyServiceMonitorsWithOptions(policyId *string, request *ListIntegrationPolicyServiceMonitorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyServiceMonitorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonReleaseName) {
		query["addonReleaseName"] = request.AddonReleaseName
	}

	if !dara.IsNil(request.EncryptYaml) {
		query["encryptYaml"] = request.EncryptYaml
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicyServiceMonitors"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/service-monitors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyServiceMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取接入中心策略的存储要求信息
//
// @param request - ListIntegrationPolicyServiceMonitorsRequest
//
// @return ListIntegrationPolicyServiceMonitorsResponse
func (client *Client) ListIntegrationPolicyServiceMonitors(policyId *string, request *ListIntegrationPolicyServiceMonitorsRequest) (_result *ListIntegrationPolicyServiceMonitorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyServiceMonitorsResponse{}
	_body, _err := client.ListIntegrationPolicyServiceMonitorsWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Storage Requirements Information for Access Center Policy
//
// Description:
//
// During the effective period of the policy, all alarms within the application group will no longer send notifications.
//
// This article provides an example of creating a pause alarm notification policy `PauseNotify` for the application group `7301****`. This application group will pause alarms from `1622949300000` to `1623208500000` (Beijing Time `2021-06-06 11:15:00` to `2021-06-09 11:15:00`).
//
// @param request - ListIntegrationPolicyStorageRequirementsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyStorageRequirementsResponse
func (client *Client) ListIntegrationPolicyStorageRequirementsWithOptions(policyId *string, request *ListIntegrationPolicyStorageRequirementsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyStorageRequirementsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["addonName"] = request.AddonName
	}

	if !dara.IsNil(request.AddonReleaseName) {
		query["addonReleaseName"] = request.AddonReleaseName
	}

	if !dara.IsNil(request.StorageType) {
		query["storageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntegrationPolicyStorageRequirements"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/storage-requirements"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntegrationPolicyStorageRequirementsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Storage Requirements Information for Access Center Policy
//
// Description:
//
// During the effective period of the policy, all alarms within the application group will no longer send notifications.
//
// This article provides an example of creating a pause alarm notification policy `PauseNotify` for the application group `7301****`. This application group will pause alarms from `1622949300000` to `1623208500000` (Beijing Time `2021-06-06 11:15:00` to `2021-06-09 11:15:00`).
//
// @param request - ListIntegrationPolicyStorageRequirementsRequest
//
// @return ListIntegrationPolicyStorageRequirementsResponse
func (client *Client) ListIntegrationPolicyStorageRequirements(policyId *string, request *ListIntegrationPolicyStorageRequirementsRequest) (_result *ListIntegrationPolicyStorageRequirementsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIntegrationPolicyStorageRequirementsResponse{}
	_body, _err := client.ListIntegrationPolicyStorageRequirementsWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询记忆库列表
//
// @param request - ListMemoryStoresRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryStoresResponse
func (client *Client) ListMemoryStoresWithOptions(workspace *string, request *ListMemoryStoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryStoresResponse, _err error) {
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

	if !dara.IsNil(request.MemoryStoreName) {
		query["memoryStoreName"] = request.MemoryStoreName
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemoryStores"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoryStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆库列表
//
// @param request - ListMemoryStoresRequest
//
// @return ListMemoryStoresResponse
func (client *Client) ListMemoryStores(workspace *string, request *ListMemoryStoresRequest) (_result *ListMemoryStoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryStoresResponse{}
	_body, _err := client.ListMemoryStoresWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Prometheus Instance Dashboard List
//
// Description:
//
// Get the list of Prometheus instance dashboards.
//
// @param request - ListPrometheusDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusDashboardsResponse
func (client *Client) ListPrometheusDashboardsWithOptions(prometheusInstanceId *string, request *ListPrometheusDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusDashboardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["aliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusDashboards"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(prometheusInstanceId)) + "/dashboards"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Prometheus Instance Dashboard List
//
// Description:
//
// Get the list of Prometheus instance dashboards.
//
// @param request - ListPrometheusDashboardsRequest
//
// @return ListPrometheusDashboardsResponse
func (client *Client) ListPrometheusDashboards(prometheusInstanceId *string, request *ListPrometheusDashboardsRequest) (_result *ListPrometheusDashboardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPrometheusDashboardsResponse{}
	_body, _err := client.ListPrometheusDashboardsWithOptions(prometheusInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the list of Prometheus instance information
//
// Description:
//
// Get the list of Prometheus instances.
//
// @param tmpReq - ListPrometheusInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusInstancesResponse
func (client *Client) ListPrometheusInstancesWithOptions(tmpReq *ListPrometheusInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPrometheusInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterRegionIds) {
		query["filterRegionIds"] = request.FilterRegionIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PrometheusInstanceIds) {
		query["prometheusInstanceIds"] = request.PrometheusInstanceIds
	}

	if !dara.IsNil(request.PrometheusInstanceName) {
		query["prometheusInstanceName"] = request.PrometheusInstanceName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusInstances"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the list of Prometheus instance information
//
// Description:
//
// Get the list of Prometheus instances.
//
// @param request - ListPrometheusInstancesRequest
//
// @return ListPrometheusInstancesResponse
func (client *Client) ListPrometheusInstances(request *ListPrometheusInstancesRequest) (_result *ListPrometheusInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPrometheusInstancesResponse{}
	_body, _err := client.ListPrometheusInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve a list of Prometheus view instance information
//
// Description:
//
// Retrieve a list of Prometheus view instance information.
//
// @param tmpReq - ListPrometheusViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusViewsResponse
func (client *Client) ListPrometheusViewsWithOptions(tmpReq *ListPrometheusViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusViewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPrometheusViewsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterRegionIds) {
		query["filterRegionIds"] = request.FilterRegionIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PrometheusViewIds) {
		query["prometheusViewIds"] = request.PrometheusViewIds
	}

	if !dara.IsNil(request.PrometheusViewName) {
		query["prometheusViewName"] = request.PrometheusViewName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusViews"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-views"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusViewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve a list of Prometheus view instance information
//
// Description:
//
// Retrieve a list of Prometheus view instance information.
//
// @param request - ListPrometheusViewsRequest
//
// @return ListPrometheusViewsResponse
func (client *Client) ListPrometheusViews(request *ListPrometheusViewsRequest) (_result *ListPrometheusViewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPrometheusViewsResponse{}
	_body, _err := client.ListPrometheusViewsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Prometheus Virtual Instance
//
// Description:
//
// # Used for creating a site monitoring task
//
// @param request - ListPrometheusVirtualInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusVirtualInstancesResponse
func (client *Client) ListPrometheusVirtualInstancesWithOptions(request *ListPrometheusVirtualInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusVirtualInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrometheusVirtualInstances"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/virtual-instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrometheusVirtualInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Prometheus Virtual Instance
//
// Description:
//
// # Used for creating a site monitoring task
//
// @param request - ListPrometheusVirtualInstancesRequest
//
// @return ListPrometheusVirtualInstancesResponse
func (client *Client) ListPrometheusVirtualInstances(request *ListPrometheusVirtualInstancesRequest) (_result *ListPrometheusVirtualInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPrometheusVirtualInstancesResponse{}
	_body, _err := client.ListPrometheusVirtualInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Resource Services
//
// @param tmpReq - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(workspace *string, tmpReq *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListServicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Resource Services
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(workspace *string, request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查标签接口
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("resourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查标签接口
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Workspace List
//
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkspaceNameList) {
		request.WorkspaceNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkspaceNameList, dara.String("workspaceNameList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["workspaceName"] = request.WorkspaceName
	}

	if !dara.IsNil(request.WorkspaceNameListShrink) {
		query["workspaceNameList"] = request.WorkspaceNameListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Workspace List
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管理告警规则
//
// @param tmpReq - ManageAlertRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageAlertRulesResponse
func (client *Client) ManageAlertRulesWithOptions(tmpReq *ManageAlertRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ManageAlertRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ManageAlertRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManageAlertRules"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manageAlertRules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManageAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管理告警规则
//
// @param request - ManageAlertRulesRequest
//
// @return ManageAlertRulesResponse
func (client *Client) ManageAlertRules(request *ManageAlertRulesRequest) (_result *ManageAlertRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ManageAlertRulesResponse{}
	_body, _err := client.ManageAlertRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Workspace
//
// @param request - PutWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspaceWithOptions(workspaceName *string, request *PutWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutWorkspaceResponse, _err error) {
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

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.SlsProject) {
		body["slsProject"] = request.SlsProject
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutWorkspace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspaceName))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Workspace
//
// @param request - PutWorkspaceRequest
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspace(workspaceName *string, request *PutWorkspaceRequest) (_result *PutWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutWorkspaceResponse{}
	_body, _err := client.PutWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询告警规则
//
// @param tmpReq - QueryAlertRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAlertRulesResponse
func (client *Client) QueryAlertRulesWithOptions(tmpReq *QueryAlertRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAlertRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryAlertRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAlertRules"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queryAlertRules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询告警规则
//
// @param request - QueryAlertRulesRequest
//
// @return QueryAlertRulesResponse
func (client *Client) QueryAlertRules(request *QueryAlertRulesRequest) (_result *QueryAlertRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAlertRulesResponse{}
	_body, _err := client.QueryAlertRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索记忆
//
// @param request - SearchMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemoriesWithOptions(workspace *string, memoryStoreName *string, request *SearchMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Filters) {
		body["filters"] = request.Filters
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Rerank) {
		body["rerank"] = request.Rerank
	}

	if !dara.IsNil(request.RetrieveLevel) {
		body["retrieveLevel"] = request.RetrieveLevel
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.SearchType) {
		body["searchType"] = request.SearchType
	}

	if !dara.IsNil(request.Threshold) {
		body["threshold"] = request.Threshold
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMemories"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索记忆
//
// @param request - SearchMemoriesRequest
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemories(workspace *string, memoryStoreName *string, request *SearchMemoriesRequest) (_result *SearchMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchMemoriesResponse{}
	_body, _err := client.SearchMemoriesWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		body["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删标签接口
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("resourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKey) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, dara.String("tagKey"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["all"] = request.All
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeyShrink) {
		query["tagKey"] = request.TagKeyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删标签接口
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Upgrade Access Component
//
// @param request - UpdateAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAddonReleaseResponse
func (client *Client) UpdateAddonReleaseWithOptions(releaseName *string, policyId *string, request *UpdateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAddonReleaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddonVersion) {
		body["addonVersion"] = request.AddonVersion
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EntityRules) {
		body["entityRules"] = request.EntityRules
	}

	if !dara.IsNil(request.Values) {
		body["values"] = request.Values
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAddonRelease"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(policyId)) + "/addon-releases/" + dara.PercentEncode(dara.StringValue(releaseName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAddonReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Upgrade Access Component
//
// @param request - UpdateAddonReleaseRequest
//
// @return UpdateAddonReleaseResponse
func (client *Client) UpdateAddonRelease(releaseName *string, policyId *string, request *UpdateAddonReleaseRequest) (_result *UpdateAddonReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAddonReleaseResponse{}
	_body, _err := client.UpdateAddonReleaseWithOptions(releaseName, policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Apply Aggregation Task Group
//
// @param request - UpdateAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupResponse
func (client *Client) UpdateAggTaskGroupWithOptions(instanceId *string, groupId *string, request *UpdateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggTaskGroupConfig) {
		body["aggTaskGroupConfig"] = request.AggTaskGroupConfig
	}

	if !dara.IsNil(request.AggTaskGroupConfigType) {
		body["aggTaskGroupConfigType"] = request.AggTaskGroupConfigType
	}

	if !dara.IsNil(request.AggTaskGroupName) {
		body["aggTaskGroupName"] = request.AggTaskGroupName
	}

	if !dara.IsNil(request.CronExpr) {
		body["cronExpr"] = request.CronExpr
	}

	if !dara.IsNil(request.Delay) {
		body["delay"] = request.Delay
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.FromTime) {
		body["fromTime"] = request.FromTime
	}

	if !dara.IsNil(request.MaxRetries) {
		body["maxRetries"] = request.MaxRetries
	}

	if !dara.IsNil(request.MaxRunTimeInSeconds) {
		body["maxRunTimeInSeconds"] = request.MaxRunTimeInSeconds
	}

	if !dara.IsNil(request.PrecheckString) {
		body["precheckString"] = request.PrecheckString
	}

	if !dara.IsNil(request.ScheduleMode) {
		body["scheduleMode"] = request.ScheduleMode
	}

	if !dara.IsNil(request.ScheduleTimeExpr) {
		body["scheduleTimeExpr"] = request.ScheduleTimeExpr
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TargetPrometheusId) {
		body["targetPrometheusId"] = request.TargetPrometheusId
	}

	if !dara.IsNil(request.ToTime) {
		body["toTime"] = request.ToTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggTaskGroup"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/agg-task-groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggTaskGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Apply Aggregation Task Group
//
// @param request - UpdateAggTaskGroupRequest
//
// @return UpdateAggTaskGroupResponse
func (client *Client) UpdateAggTaskGroup(instanceId *string, groupId *string, request *UpdateAggTaskGroupRequest) (_result *UpdateAggTaskGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAggTaskGroupResponse{}
	_body, _err := client.UpdateAggTaskGroupWithOptions(instanceId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Aggregation Task Group Status
//
// @param request - UpdateAggTaskGroupStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupStatusResponse
func (client *Client) UpdateAggTaskGroupStatusWithOptions(instanceId *string, groupId *string, request *UpdateAggTaskGroupStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAggTaskGroupStatus"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/agg-task-groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/status"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAggTaskGroupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Aggregation Task Group Status
//
// @param request - UpdateAggTaskGroupStatusRequest
//
// @return UpdateAggTaskGroupStatusResponse
func (client *Client) UpdateAggTaskGroupStatus(instanceId *string, groupId *string, request *UpdateAggTaskGroupStatusRequest) (_result *UpdateAggTaskGroupStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAggTaskGroupStatusResponse{}
	_body, _err := client.UpdateAggTaskGroupStatusWithOptions(instanceId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改已存在的告警 Webhook 通知配置。
//
// @param request - UpdateAlertWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertWebhookResponse
func (client *Client) UpdateAlertWebhookWithOptions(webhookId *string, request *UpdateAlertWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.Headers) {
		body["headers"] = request.Headers
	}

	if !dara.IsNil(request.Lang) {
		body["lang"] = request.Lang
	}

	if !dara.IsNil(request.Method) {
		body["method"] = request.Method
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertWebhook"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webhook/" + dara.PercentEncode(dara.StringValue(webhookId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改已存在的告警 Webhook 通知配置。
//
// @param request - UpdateAlertWebhookRequest
//
// @return UpdateAlertWebhookResponse
func (client *Client) UpdateAlertWebhook(webhookId *string, request *UpdateAlertWebhookRequest) (_result *UpdateAlertWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertWebhookResponse{}
	_body, _err := client.UpdateAlertWebhookWithOptions(webhookId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改业务链路
//
// @param request - UpdateBizTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizTraceResponse
func (client *Client) UpdateBizTraceWithOptions(bizTraceId *string, request *UpdateBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBizTraceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedConfig) {
		body["advancedConfig"] = request.AdvancedConfig
	}

	if !dara.IsNil(request.BizTraceName) {
		body["bizTraceName"] = request.BizTraceName
	}

	if !dara.IsNil(request.RuleConfig) {
		body["ruleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizTrace"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bizTrace/" + dara.PercentEncode(dara.StringValue(bizTraceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改业务链路
//
// @param request - UpdateBizTraceRequest
//
// @return UpdateBizTraceResponse
func (client *Client) UpdateBizTrace(bizTraceId *string, request *UpdateBizTraceRequest) (_result *UpdateBizTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBizTraceResponse{}
	_body, _err := client.UpdateBizTraceWithOptions(bizTraceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集
//
// @param request - UpdateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithOptions(workspace *string, datasetName *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集
//
// @param request - UpdateDatasetRequest
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDataset(workspace *string, datasetName *string, request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(workspace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据投递任务
//
// @param request - UpdateDeliveryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeliveryTaskResponse
func (client *Client) UpdateDeliveryTaskWithOptions(taskId *string, request *UpdateDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDeliveryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["dataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.ExternalLabels) {
		body["externalLabels"] = request.ExternalLabels
	}

	if !dara.IsNil(request.LabelFilters) {
		body["labelFilters"] = request.LabelFilters
	}

	if !dara.IsNil(request.LabelFiltersType) {
		body["labelFiltersType"] = request.LabelFiltersType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SinkList) {
		body["sinkList"] = request.SinkList
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	if !dara.IsNil(request.TaskName) {
		body["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeliveryTask"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/delivery-task/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeliveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据投递任务
//
// @param request - UpdateDeliveryTaskRequest
//
// @return UpdateDeliveryTaskResponse
func (client *Client) UpdateDeliveryTask(taskId *string, request *UpdateDeliveryTaskRequest) (_result *UpdateDeliveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDeliveryTaskResponse{}
	_body, _err := client.UpdateDeliveryTaskWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update the specified policy
//
// @param request - UpdateIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegrationPolicyResponse
func (client *Client) UpdateIntegrationPolicyWithOptions(integrationPolicyId *string, request *UpdateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIntegrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FeePackage) {
		body["feePackage"] = request.FeePackage
	}

	if !dara.IsNil(request.PolicyName) {
		body["policyName"] = request.PolicyName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIntegrationPolicy"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/integration-policies/" + dara.PercentEncode(dara.StringValue(integrationPolicyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIntegrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update the specified policy
//
// @param request - UpdateIntegrationPolicyRequest
//
// @return UpdateIntegrationPolicyResponse
func (client *Client) UpdateIntegrationPolicy(integrationPolicyId *string, request *UpdateIntegrationPolicyRequest) (_result *UpdateIntegrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIntegrationPolicyResponse{}
	_body, _err := client.UpdateIntegrationPolicyWithOptions(integrationPolicyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改记忆
//
// @param request - UpdateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(workspace *string, memoryStoreName *string, memoryId *string, request *UpdateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemory"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName)) + "/memory/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改记忆
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(workspace *string, memoryStoreName *string, memoryId *string, request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(workspace, memoryStoreName, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改记忆库配置
//
// @param request - UpdateMemoryStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryStoreResponse
func (client *Client) UpdateMemoryStoreWithOptions(workspace *string, memoryStoreName *string, request *UpdateMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomExtractionStrategies) {
		body["customExtractionStrategies"] = request.CustomExtractionStrategies
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ExtractionStrategies) {
		body["extractionStrategies"] = request.ExtractionStrategies
	}

	if !dara.IsNil(request.ShortTermTtl) {
		body["shortTermTtl"] = request.ShortTermTtl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemoryStore"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/memorystore/" + dara.PercentEncode(dara.StringValue(memoryStoreName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改记忆库配置
//
// @param request - UpdateMemoryStoreRequest
//
// @return UpdateMemoryStoreResponse
func (client *Client) UpdateMemoryStore(workspace *string, memoryStoreName *string, request *UpdateMemoryStoreRequest) (_result *UpdateMemoryStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryStoreResponse{}
	_body, _err := client.UpdateMemoryStoreWithOptions(workspace, memoryStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改已存在的告警通知策略
//
// @param request - UpdateNotifyStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNotifyStrategyResponse
func (client *Client) UpdateNotifyStrategyWithOptions(notifyStrategyId *string, request *UpdateNotifyStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNotifyStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNotifyStrategy"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/notifyStrategies/" + dara.PercentEncode(dara.StringValue(notifyStrategyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNotifyStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改已存在的告警通知策略
//
// @param request - UpdateNotifyStrategyRequest
//
// @return UpdateNotifyStrategyResponse
func (client *Client) UpdateNotifyStrategy(notifyStrategyId *string, request *UpdateNotifyStrategyRequest) (_result *UpdateNotifyStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNotifyStrategyResponse{}
	_body, _err := client.UpdateNotifyStrategyWithOptions(notifyStrategyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Prometheus instance information
//
// Description:
//
// Update Prometheus instance information.
//
// @param request - UpdatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusInstanceResponse
func (client *Client) UpdatePrometheusInstanceWithOptions(prometheusInstanceId *string, request *UpdatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArchiveDuration) {
		body["archiveDuration"] = request.ArchiveDuration
	}

	if !dara.IsNil(request.AuthFreeReadPolicy) {
		body["authFreeReadPolicy"] = request.AuthFreeReadPolicy
	}

	if !dara.IsNil(request.AuthFreeWritePolicy) {
		body["authFreeWritePolicy"] = request.AuthFreeWritePolicy
	}

	if !dara.IsNil(request.EnableAuthFreeRead) {
		body["enableAuthFreeRead"] = request.EnableAuthFreeRead
	}

	if !dara.IsNil(request.EnableAuthFreeWrite) {
		body["enableAuthFreeWrite"] = request.EnableAuthFreeWrite
	}

	if !dara.IsNil(request.EnableAuthToken) {
		body["enableAuthToken"] = request.EnableAuthToken
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.PrometheusInstanceName) {
		body["prometheusInstanceName"] = request.PrometheusInstanceName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.StorageDuration) {
		body["storageDuration"] = request.StorageDuration
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusInstance"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-instances/" + dara.PercentEncode(dara.StringValue(prometheusInstanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Prometheus instance information
//
// Description:
//
// Update Prometheus instance information.
//
// @param request - UpdatePrometheusInstanceRequest
//
// @return UpdatePrometheusInstanceResponse
func (client *Client) UpdatePrometheusInstance(prometheusInstanceId *string, request *UpdatePrometheusInstanceRequest) (_result *UpdatePrometheusInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePrometheusInstanceResponse{}
	_body, _err := client.UpdatePrometheusInstanceWithOptions(prometheusInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Prom实例信息
//
// @param request - UpdatePrometheusUserSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusUserSettingResponse
func (client *Client) UpdatePrometheusUserSettingWithOptions(settingKey *string, request *UpdatePrometheusUserSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusUserSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SettingValue) {
		query["settingValue"] = request.SettingValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusUserSetting"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-user-setting/" + dara.PercentEncode(dara.StringValue(settingKey))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusUserSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Prom实例信息
//
// @param request - UpdatePrometheusUserSettingRequest
//
// @return UpdatePrometheusUserSettingResponse
func (client *Client) UpdatePrometheusUserSetting(settingKey *string, request *UpdatePrometheusUserSettingRequest) (_result *UpdatePrometheusUserSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePrometheusUserSettingResponse{}
	_body, _err := client.UpdatePrometheusUserSettingWithOptions(settingKey, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Prometheus view instance information
//
// Description:
//
// Update Prometheus view instance information.
//
// @param request - UpdatePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusViewResponse
func (client *Client) UpdatePrometheusViewWithOptions(prometheusViewId *string, request *UpdatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthFreeReadPolicy) {
		body["authFreeReadPolicy"] = request.AuthFreeReadPolicy
	}

	if !dara.IsNil(request.EnableAuthFreeRead) {
		body["enableAuthFreeRead"] = request.EnableAuthFreeRead
	}

	if !dara.IsNil(request.EnableAuthToken) {
		body["enableAuthToken"] = request.EnableAuthToken
	}

	if !dara.IsNil(request.PrometheusInstances) {
		body["prometheusInstances"] = request.PrometheusInstances
	}

	if !dara.IsNil(request.PrometheusViewName) {
		body["prometheusViewName"] = request.PrometheusViewName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrometheusView"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/prometheus-views/" + dara.PercentEncode(dara.StringValue(prometheusViewId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrometheusViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Prometheus view instance information
//
// Description:
//
// Update Prometheus view instance information.
//
// @param request - UpdatePrometheusViewRequest
//
// @return UpdatePrometheusViewResponse
func (client *Client) UpdatePrometheusView(prometheusViewId *string, request *UpdatePrometheusViewRequest) (_result *UpdatePrometheusViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePrometheusViewResponse{}
	_body, _err := client.UpdatePrometheusViewWithOptions(prometheusViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Service
//
// @param request - UpdateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithOptions(workspace *string, serviceId *string, request *UpdateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.ServiceStatus) {
		body["serviceStatus"] = request.ServiceStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateService"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/service/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Service
//
// @param request - UpdateServiceRequest
//
// @return UpdateServiceResponse
func (client *Client) UpdateService(workspace *string, serviceId *string, request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(workspace, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个已存在的订阅配置
//
// @param request - UpdateSubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithOptions(subscriptionId *string, request *UpdateSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Workspace) {
		query["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubscription"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/subscriptions/" + dara.PercentEncode(dara.StringValue(subscriptionId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个已存在的订阅配置
//
// @param request - UpdateSubscriptionRequest
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscription(subscriptionId *string, request *UpdateSubscriptionRequest) (_result *UpdateSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.UpdateSubscriptionWithOptions(subscriptionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Umodel configuration information
//
// Description:
//
// # Update Umodel configuration information
//
// @param request - UpdateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodelWithOptions(workspace *string, request *UpdateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUmodelResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUmodel"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUmodelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Umodel configuration information
//
// Description:
//
// # Update Umodel configuration information
//
// @param request - UpdateUmodelRequest
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodel(workspace *string, request *UpdateUmodelRequest) (_result *UpdateUmodelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUmodelResponse{}
	_body, _err := client.UpdateUmodelWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Umodel配置信息
//
// @param request - UpsertUmodelCommonSchemaRefRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelCommonSchemaRefResponse
func (client *Client) UpsertUmodelCommonSchemaRefWithOptions(workspace *string, request *UpsertUmodelCommonSchemaRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpsertUmodelCommonSchemaRefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Group) {
		query["group"] = request.Group
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertUmodelCommonSchemaRef"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel/common-schema-ref"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertUmodelCommonSchemaRefResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Umodel配置信息
//
// @param request - UpsertUmodelCommonSchemaRefRequest
//
// @return UpsertUmodelCommonSchemaRefResponse
func (client *Client) UpsertUmodelCommonSchemaRef(workspace *string, request *UpsertUmodelCommonSchemaRefRequest) (_result *UpsertUmodelCommonSchemaRefResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpsertUmodelCommonSchemaRefResponse{}
	_body, _err := client.UpsertUmodelCommonSchemaRefWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Write Umodel Elements
//
// @param request - UpsertUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelDataWithOptions(workspace *string, request *UpsertUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpsertUmodelDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Method) {
		query["method"] = request.Method
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Elements) {
		body["elements"] = request.Elements
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertUmodelData"),
		Version:     dara.String("2024-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/workspace/" + dara.PercentEncode(dara.StringValue(workspace)) + "/umodel/data"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertUmodelDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Write Umodel Elements
//
// @param request - UpsertUmodelDataRequest
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelData(workspace *string, request *UpsertUmodelDataRequest) (_result *UpsertUmodelDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpsertUmodelDataResponse{}
	_body, _err := client.UpsertUmodelDataWithOptions(workspace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
