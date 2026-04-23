// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AddMemoriesWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *AddMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMemoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAddonReleaseResponse
func (client *Client) CreateAddonReleaseWithContext(ctx context.Context, policyId *string, request *CreateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggTaskGroupResponse
func (client *Client) CreateAggTaskGroupWithContext(ctx context.Context, instanceId *string, request *CreateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAggTaskGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertWebhookResponse
func (client *Client) CreateAlertWebhookWithContext(ctx context.Context, request *CreateAlertWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlertWebhookResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizTraceResponse
func (client *Client) CreateBizTraceWithContext(ctx context.Context, request *CreateBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBizTraceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudResourceResponse
func (client *Client) CreateCloudResourceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCloudResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, workspace *string, request *CreateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryTaskResponse
func (client *Client) CreateDeliveryTaskWithContext(ctx context.Context, request *CreateDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDeliveryTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStoreWithContext(ctx context.Context, workspaceName *string, request *CreateEntityStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEntityStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntegrationPolicyResponse
func (client *Client) CreateIntegrationPolicyWithContext(ctx context.Context, request *CreateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryStoreResponse
func (client *Client) CreateMemoryStoreWithContext(ctx context.Context, workspace *string, request *CreateMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstanceWithContext(ctx context.Context, request *CreatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusViewResponse
func (client *Client) CreatePrometheusViewWithContext(ctx context.Context, request *CreatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusVirtualInstanceResponse
func (client *Client) CreatePrometheusVirtualInstanceWithContext(ctx context.Context, request *CreatePrometheusVirtualInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusVirtualInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithContext(ctx context.Context, workspace *string, request *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceObservabilityResponse
func (client *Client) CreateServiceObservabilityWithContext(ctx context.Context, workspace *string, _type *string, request *CreateServiceObservabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceObservabilityResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithContext(ctx context.Context, request *CreateTicketRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodelWithContext(ctx context.Context, workspace *string, request *CreateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUmodelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddonReleaseResponse
func (client *Client) DeleteAddonReleaseWithContext(ctx context.Context, policyId *string, request *DeleteAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggTaskGroupResponse
func (client *Client) DeleteAggTaskGroupWithContext(ctx context.Context, instanceId *string, groupId *string, request *DeleteAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAggTaskGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteAlertWebhooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertWebhooksResponse
func (client *Client) DeleteAlertWebhooksWithContext(ctx context.Context, tmpReq *DeleteAlertWebhooksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertWebhooksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizTraceResponse
func (client *Client) DeleteBizTraceWithContext(ctx context.Context, bizTraceId *string, request *DeleteBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBizTraceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudResourceResponse
func (client *Client) DeleteCloudResourceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCloudResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, workspace *string, datasetName *string, request *DeleteDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeliveryTaskResponse
func (client *Client) DeleteDeliveryTaskWithContext(ctx context.Context, taskId *string, request *DeleteDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDeliveryTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStoreWithContext(ctx context.Context, workspaceName *string, request *DeleteEntityStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEntityStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntegrationPolicyResponse
func (client *Client) DeleteIntegrationPolicyWithContext(ctx context.Context, policyId *string, request *DeleteIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoriesResponse
func (client *Client) DeleteMemoriesWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *DeleteMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithContext(ctx context.Context, workspace *string, memoryStoreName *string, memoryId *string, request *DeleteMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryStoreResponse
func (client *Client) DeleteMemoryStoreWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *DeleteMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusInstanceResponse
func (client *Client) DeletePrometheusInstanceWithContext(ctx context.Context, prometheusInstanceId *string, request *DeletePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusViewResponse
func (client *Client) DeletePrometheusViewWithContext(ctx context.Context, prometheusViewId *string, request *DeletePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusVirtualInstanceResponse
func (client *Client) DeletePrometheusVirtualInstanceWithContext(ctx context.Context, prometheusInstanceId *string, request *DeletePrometheusVirtualInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusVirtualInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithContext(ctx context.Context, workspace *string, serviceId *string, request *DeleteServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodelWithContext(ctx context.Context, workspace *string, request *DeleteUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelCommonSchemaRefResponse
func (client *Client) DeleteUmodelCommonSchemaRefWithContext(ctx context.Context, workspace *string, request *DeleteUmodelCommonSchemaRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelCommonSchemaRefResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelDataResponse
func (client *Client) DeleteUmodelDataWithContext(ctx context.Context, workspace *string, request *DeleteUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithContext(ctx context.Context, workspaceName *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteQueryResponse
func (client *Client) ExecuteQueryWithContext(ctx context.Context, workspace *string, datasetName *string, request *ExecuteQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteQueryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonResponse
func (client *Client) GetAddonWithContext(ctx context.Context, addonName *string, request *GetAddonRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonCodeTemplateResponse
func (client *Client) GetAddonCodeTemplateWithContext(ctx context.Context, addonName *string, request *GetAddonCodeTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonCodeTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonReleaseResponse
func (client *Client) GetAddonReleaseWithContext(ctx context.Context, releaseName *string, policyId *string, request *GetAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonSchemaResponse
func (client *Client) GetAddonSchemaWithContext(ctx context.Context, addonName *string, request *GetAddonSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonSchemaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggTaskGroupResponse
func (client *Client) GetAggTaskGroupWithContext(ctx context.Context, instanceId *string, groupId *string, request *GetAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAggTaskGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizTraceResponse
func (client *Client) GetBizTraceWithContext(ctx context.Context, bizTraceId *string, request *GetBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBizTraceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudResourceResponse
func (client *Client) GetCloudResourceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCloudResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudResourceDataResponse
func (client *Client) GetCloudResourceDataWithContext(ctx context.Context, request *GetCloudResourceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCloudResourceDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCmsServiceResponse
func (client *Client) GetCmsServiceWithContext(ctx context.Context, request *GetCmsServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCmsServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithContext(ctx context.Context, workspace *string, datasetName *string, request *GetDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryTaskResponse
func (client *Client) GetDeliveryTaskWithContext(ctx context.Context, taskId *string, request *GetDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeliveryTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStoreWithContext(ctx context.Context, workspaceName *string, request *GetEntityStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEntityStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - GetEntityStoreDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreDataResponse
func (client *Client) GetEntityStoreDataWithContext(ctx context.Context, workspace *string, request *GetEntityStoreDataRequest, headers *GetEntityStoreDataHeaders, runtime *dara.RuntimeOptions) (_result *GetEntityStoreDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationPolicyResponse
func (client *Client) GetIntegrationPolicyWithContext(ctx context.Context, policyId *string, request *GetIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationVersionForCSResponse
func (client *Client) GetIntegrationVersionForCSWithContext(ctx context.Context, request *GetIntegrationVersionForCSRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIntegrationVersionForCSResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoriesResponse
func (client *Client) GetMemoriesWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *GetMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithContext(ctx context.Context, workspace *string, memoryStoreName *string, memoryId *string, request *GetMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryHistoryResponse
func (client *Client) GetMemoryHistoryWithContext(ctx context.Context, workspace *string, memoryStoreName *string, memoryId *string, request *GetMemoryHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryHistoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryStoreResponse
func (client *Client) GetMemoryStoreWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *GetMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusInstanceResponse
func (client *Client) GetPrometheusInstanceWithContext(ctx context.Context, prometheusInstanceId *string, request *GetPrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusUserSettingResponse
func (client *Client) GetPrometheusUserSettingWithContext(ctx context.Context, request *GetPrometheusUserSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusUserSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusViewResponse
func (client *Client) GetPrometheusViewWithContext(ctx context.Context, prometheusViewId *string, request *GetPrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithContext(ctx context.Context, workspace *string, serviceId *string, request *GetServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceObservabilityResponse
func (client *Client) GetServiceObservabilityWithContext(ctx context.Context, workspace *string, _type *string, request *GetServiceObservabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceObservabilityResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelResponse
func (client *Client) GetUmodelWithContext(ctx context.Context, workspace *string, request *GetUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelCommonSchemaRefResponse
func (client *Client) GetUmodelCommonSchemaRefWithContext(ctx context.Context, workspace *string, request *GetUmodelCommonSchemaRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelCommonSchemaRefResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelDataResponse
func (client *Client) GetUmodelDataWithContext(ctx context.Context, workspace *string, request *GetUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithContext(ctx context.Context, workspaceName *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonReleasesResponse
func (client *Client) ListAddonReleasesWithContext(ctx context.Context, policyId *string, request *ListAddonReleasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonReleasesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithContext(ctx context.Context, request *ListAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListAggTaskGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggTaskGroupsResponse
func (client *Client) ListAggTaskGroupsWithContext(ctx context.Context, instanceId *string, tmpReq *ListAggTaskGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAggTaskGroupsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListAlertActionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActionsWithContext(ctx context.Context, tmpReq *ListAlertActionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertActionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListAlertRobotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertRobotsResponse
func (client *Client) ListAlertRobotsWithContext(ctx context.Context, tmpReq *ListAlertRobotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertRobotsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListAlertWebhooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertWebhooksResponse
func (client *Client) ListAlertWebhooksWithContext(ctx context.Context, tmpReq *ListAlertWebhooksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertWebhooksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizTracesResponse
func (client *Client) ListBizTracesWithContext(ctx context.Context, request *ListBizTracesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBizTracesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListContactGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactGroupsResponse
func (client *Client) ListContactGroupsWithContext(ctx context.Context, tmpReq *ListContactGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContactGroupsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListContactsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactsResponse
func (client *Client) ListContactsWithContext(ctx context.Context, tmpReq *ListContactsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContactsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithContext(ctx context.Context, workspace *string, request *ListDatasetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListDeliveryTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryTasksResponse
func (client *Client) ListDeliveryTasksWithContext(ctx context.Context, tmpReq *ListDeliveryTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeliveryTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListIntegrationPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPoliciesResponse
func (client *Client) ListIntegrationPoliciesWithContext(ctx context.Context, tmpReq *ListIntegrationPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPoliciesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyAddonsResponse
func (client *Client) ListIntegrationPolicyAddonsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyAddonsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyAddonsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyCollectorsResponse
func (client *Client) ListIntegrationPolicyCollectorsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyCollectorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyCollectorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyCustomScrapeJobRulesResponse
func (client *Client) ListIntegrationPolicyCustomScrapeJobRulesWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyCustomScrapeJobRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyCustomScrapeJobRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyDashboardsResponse
func (client *Client) ListIntegrationPolicyDashboardsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyDashboardsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyPodMonitorsResponse
func (client *Client) ListIntegrationPolicyPodMonitorsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyPodMonitorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyPodMonitorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyServiceMonitorsResponse
func (client *Client) ListIntegrationPolicyServiceMonitorsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyServiceMonitorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyServiceMonitorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyStorageRequirementsResponse
func (client *Client) ListIntegrationPolicyStorageRequirementsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyStorageRequirementsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyStorageRequirementsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryStoresResponse
func (client *Client) ListMemoryStoresWithContext(ctx context.Context, workspace *string, request *ListMemoryStoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryStoresResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusDashboardsResponse
func (client *Client) ListPrometheusDashboardsWithContext(ctx context.Context, prometheusInstanceId *string, request *ListPrometheusDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusDashboardsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListPrometheusInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusInstancesResponse
func (client *Client) ListPrometheusInstancesWithContext(ctx context.Context, tmpReq *ListPrometheusInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListPrometheusViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusViewsResponse
func (client *Client) ListPrometheusViewsWithContext(ctx context.Context, tmpReq *ListPrometheusViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusViewsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusVirtualInstancesResponse
func (client *Client) ListPrometheusVirtualInstancesWithContext(ctx context.Context, request *ListPrometheusVirtualInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusVirtualInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithContext(ctx context.Context, workspace *string, tmpReq *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithContext(ctx context.Context, tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ManageAlertRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageAlertRulesResponse
func (client *Client) ManageAlertRulesWithContext(ctx context.Context, tmpReq *ManageAlertRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ManageAlertRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspaceWithContext(ctx context.Context, workspaceName *string, request *PutWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - QueryAlertRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAlertRulesResponse
func (client *Client) QueryAlertRulesWithContext(ctx context.Context, tmpReq *QueryAlertRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAlertRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMemoriesResponse
func (client *Client) SearchMemoriesWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *SearchMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMemoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAddonReleaseResponse
func (client *Client) UpdateAddonReleaseWithContext(ctx context.Context, releaseName *string, policyId *string, request *UpdateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupResponse
func (client *Client) UpdateAggTaskGroupWithContext(ctx context.Context, instanceId *string, groupId *string, request *UpdateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupStatusResponse
func (client *Client) UpdateAggTaskGroupStatusWithContext(ctx context.Context, instanceId *string, groupId *string, request *UpdateAggTaskGroupStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertWebhookResponse
func (client *Client) UpdateAlertWebhookWithContext(ctx context.Context, webhookId *string, request *UpdateAlertWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertWebhookResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizTraceResponse
func (client *Client) UpdateBizTraceWithContext(ctx context.Context, bizTraceId *string, request *UpdateBizTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBizTraceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithContext(ctx context.Context, workspace *string, datasetName *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeliveryTaskResponse
func (client *Client) UpdateDeliveryTaskWithContext(ctx context.Context, taskId *string, request *UpdateDeliveryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDeliveryTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegrationPolicyResponse
func (client *Client) UpdateIntegrationPolicyWithContext(ctx context.Context, integrationPolicyId *string, request *UpdateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithContext(ctx context.Context, workspace *string, memoryStoreName *string, memoryId *string, request *UpdateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryStoreResponse
func (client *Client) UpdateMemoryStoreWithContext(ctx context.Context, workspace *string, memoryStoreName *string, request *UpdateMemoryStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryStoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNotifyStrategyResponse
func (client *Client) UpdateNotifyStrategyWithContext(ctx context.Context, notifyStrategyId *string, request *UpdateNotifyStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNotifyStrategyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusInstanceResponse
func (client *Client) UpdatePrometheusInstanceWithContext(ctx context.Context, prometheusInstanceId *string, request *UpdatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusUserSettingResponse
func (client *Client) UpdatePrometheusUserSettingWithContext(ctx context.Context, settingKey *string, request *UpdatePrometheusUserSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusUserSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusViewResponse
func (client *Client) UpdatePrometheusViewWithContext(ctx context.Context, prometheusViewId *string, request *UpdatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithContext(ctx context.Context, workspace *string, serviceId *string, request *UpdateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithContext(ctx context.Context, subscriptionId *string, request *UpdateSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodelWithContext(ctx context.Context, workspace *string, request *UpdateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUmodelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelCommonSchemaRefResponse
func (client *Client) UpsertUmodelCommonSchemaRefWithContext(ctx context.Context, workspace *string, request *UpsertUmodelCommonSchemaRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpsertUmodelCommonSchemaRefResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelDataWithContext(ctx context.Context, workspace *string, request *UpsertUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpsertUmodelDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
