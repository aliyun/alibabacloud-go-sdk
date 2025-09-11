// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds tags to an application group.
//
// Description:
//
// This topic provides an example on how to add a tag to an application group whose ID is `7301****`. In this example, the key of the tag is `key1` and the value of the tag is `value1`.
//
// @param request - AddTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTagsResponse
func (client *Client) AddTagsWithContext(ctx context.Context, request *AddTagsRequest, runtime *dara.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTags"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an alert template to an application group to generate an alert rule.
//
// Description:
//
// In this example, the `700****` alert template is applied to the `123456` application group. For the generated alert rule, the ID is `applyTemplate8ab74c6b-9f27-47ab-8841-de01dc08****`, and the name is `test123`.
//
// @param request - ApplyMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyMetricRuleTemplateResponse
func (client *Client) ApplyMetricRuleTemplateWithContext(ctx context.Context, request *ApplyMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *ApplyMetricRuleTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppendMode) {
		query["AppendMode"] = request.AppendMode
	}

	if !dara.IsNil(request.ApplyMode) {
		query["ApplyMode"] = request.ApplyMode
	}

	if !dara.IsNil(request.EnableEndTime) {
		query["EnableEndTime"] = request.EnableEndTime
	}

	if !dara.IsNil(request.EnableStartTime) {
		query["EnableStartTime"] = request.EnableStartTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NotifyLevel) {
		query["NotifyLevel"] = request.NotifyLevel
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyMetricRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates site monitoring tasks.
//
// Description:
//
// This topic provides an example on how to create a site monitoring task named `HangZhou_ECS1`. The URL that is monitored by the task is `https://www.aliyun.com` and the type of the task is `HTTP`. The returned result shows that the site monitoring task is created. The name of the site monitoring task is `HangZhou_ECS1` and the task ID is `679fbe4f-b80b-4706-91b2-5427b43e****`.
//
// @param request - BatchCreateInstantSiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateInstantSiteMonitorResponse
func (client *Client) BatchCreateInstantSiteMonitorWithContext(ctx context.Context, request *BatchCreateInstantSiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateInstantSiteMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskList) {
		query["TaskList"] = request.TaskList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateInstantSiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateInstantSiteMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports the monitoring data that is defined in the Cursor operation.
//
// Description:
//
// ### [](#)Prerequisites
//
// The `Cursor` information is returned by calling the [Cursor](https://help.aliyun.com/document_detail/2330730.html) operation.
//
// ### [](#)Description
//
// This topic provides an example on how to export the monitoring data of the `cpu_idle` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The `Cursor` information is specified. A maximum of 1,000 data entries can be returned in each response.
//
// @param tmpReq - BatchExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchExportResponse
func (client *Client) BatchExportWithContext(ctx context.Context, tmpReq *BatchExportRequest, runtime *dara.RuntimeOptions) (_result *BatchExportResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchExportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Measurements) {
		request.MeasurementsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Measurements, dara.String("Measurements"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.MeasurementsShrink) {
		body["Measurements"] = request.MeasurementsShrink
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchExport"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchExportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group based on the tags of cloud resources.
//
// Description:
//
// This operation is available for Elastic Compute Service (ECS), ApsaraDB RDS, and Server Load Balancer (SLB).
//
// This topic provides an example to show how to create an application group for resources whose tag key is `ecs_instance`. In this example, the alert contact group of the application group is `ECS_Group`.
//
// @param request - CreateDynamicTagGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDynamicTagGroupResponse
func (client *Client) CreateDynamicTagGroupWithContext(ctx context.Context, request *CreateDynamicTagGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDynamicTagGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupList) {
		query["ContactGroupList"] = request.ContactGroupList
	}

	if !dara.IsNil(request.EnableInstallAgent) {
		query["EnableInstallAgent"] = request.EnableInstallAgent
	}

	if !dara.IsNil(request.EnableSubscribeEvent) {
		query["EnableSubscribeEvent"] = request.EnableSubscribeEvent
	}

	if !dara.IsNil(request.MatchExpress) {
		query["MatchExpress"] = request.MatchExpress
	}

	if !dara.IsNil(request.MatchExpressFilterRelation) {
		query["MatchExpressFilterRelation"] = request.MatchExpressFilterRelation
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.TagRegionId) {
		query["TagRegionId"] = request.TagRegionId
	}

	if !dara.IsNil(request.TemplateIdList) {
		query["TemplateIdList"] = request.TemplateIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDynamicTagGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDynamicTagGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates one or more alert rules for a specified application group.
//
// Description:
//
// This topic provides an example to show how to create an alert rule for the `cpu_total` metric of Elastic Compute Service (ECS) in the `123456` application group. The ID of the alert rule is `456789`. The name of the alert rule is `ECS_Rule1`. The alert level is `Critical`. The statistical method is `Average`. The comparison operator is `GreaterThanOrEqualToThreshold`. The alert threshold is `90`. The number of alert retries is `3`. The response shows that the alert rule named `ECS_Rule1` is created.
//
// @param request - CreateGroupMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupMetricRulesResponse
func (client *Client) CreateGroupMetricRulesWithContext(ctx context.Context, request *CreateGroupMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupMetricRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupMetricRules) {
		query["GroupMetricRules"] = request.GroupMetricRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupMetricRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a process monitoring task for an application group.
//
// @param request - CreateGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupMonitoringAgentProcessResponse
func (client *Client) CreateGroupMonitoringAgentProcessWithContext(ctx context.Context, request *CreateGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupMonitoringAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MatchExpress) {
		query["MatchExpress"] = request.MatchExpress
	}

	if !dara.IsNil(request.MatchExpressFilterRelation) {
		query["MatchExpressFilterRelation"] = request.MatchExpressFilterRelation
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an availability monitoring task.
//
// Description:
//
// This topic provides an example on how to create an availability monitoring task named `task1` in an application group named `123456`. The TaskType parameter of the task is set to `HTTP`. After you start the task, the system sends alerts by using the specified email address and DingTalk chatbot.
//
// @param request - CreateHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostAvailabilityResponse
func (client *Client) CreateHostAvailabilityWithContext(ctx context.Context, request *CreateHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *CreateHostAvailabilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfigEscalationList) {
		query["AlertConfigEscalationList"] = request.AlertConfigEscalationList
	}

	if !dara.IsNil(request.AlertConfigTargetList) {
		query["AlertConfigTargetList"] = request.AlertConfigTargetList
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskScope) {
		query["TaskScope"] = request.TaskScope
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.TaskOption) {
		query["TaskOption"] = request.TaskOption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostAvailabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a namespace.
//
// Description:
//
// ## [](#)Prerequisites
//
// Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## [](#)Operation description
//
// This topic provides an example on how to create a namespace named `aliyun`. In this example, the data retention period of the namespace is set to `cms.s1.3xlarge`. The returned result indicates that the namespace is created.
//
// @param request - CreateHybridMonitorNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridMonitorNamespaceResponse
func (client *Client) CreateHybridMonitorNamespaceWithContext(ctx context.Context, request *CreateHybridMonitorNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridMonitorNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceRegion) {
		query["NamespaceRegion"] = request.NamespaceRegion
	}

	if !dara.IsNil(request.NamespaceType) {
		query["NamespaceType"] = request.NamespaceType
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridMonitorNamespace"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridMonitorNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Logstore group of Hybrid Cloud Monitoring.
//
// Description:
//
// ### [](#)Prerequisites
//
// Simple Log Service is activated. A project and a Logstore are created in Simple Log Service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/54604.html).
//
// ### [](#)Operation description
//
// This topic provides an example on how to create a Logstore group named `Logstore_test`. The region ID is `cn-hangzhou`. The project is `aliyun-project`. The Logstore is `Logstore-ECS`. The response shows that the Logstore group is created.
//
// @param request - CreateHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridMonitorSLSGroupResponse
func (client *Client) CreateHybridMonitorSLSGroupWithContext(ctx context.Context, request *CreateHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridMonitorSLSGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SLSGroupConfig) {
		query["SLSGroupConfig"] = request.SLSGroupConfig
	}

	if !dara.IsNil(request.SLSGroupDescription) {
		query["SLSGroupDescription"] = request.SLSGroupDescription
	}

	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a metric import task for an Alibaba Cloud service or creates a metric for logs imported from Simple Log Service.
//
// Description:
//
// # [](#)Prerequisites
//
//   - Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
//   - If you want to create a metric for logs imported from Simple Log Service, make sure that you have activated Simple Log Service and created a project and a Logstore. For more information, see [Getting Started](https://help.aliyun.com/document_detail/54604.html).
//
// # [](#)Description
//
// This topic provides an example on how to create a metric import task named `aliyun_task` for Elastic Compute Service (ECS). The task imports the `cpu_total` metric to the `aliyun` namespace. The response shows that the metric import task is created.
//
// @param request - CreateHybridMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridMonitorTaskResponse
func (client *Client) CreateHybridMonitorTaskWithContext(ctx context.Context, request *CreateHybridMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridMonitorTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachLabels) {
		query["AttachLabels"] = request.AttachLabels
	}

	if !dara.IsNil(request.CloudAccessId) {
		query["CloudAccessId"] = request.CloudAccessId
	}

	if !dara.IsNil(request.CollectInterval) {
		query["CollectInterval"] = request.CollectInterval
	}

	if !dara.IsNil(request.CollectTargetType) {
		query["CollectTargetType"] = request.CollectTargetType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SLSProcessConfig) {
		query["SLSProcessConfig"] = request.SLSProcessConfig
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TargetUserIdList) {
		query["TargetUserIdList"] = request.TargetUserIdList
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.YARMConfig) {
		query["YARMConfig"] = request.YARMConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridMonitorTask"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridMonitorTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instant test task.
//
// Description:
//
// You can create an instant test task only by using the Alibaba Cloud account that you used to enable Network Analysis and Monitoring.
//
// This topic provides an example to show how to create an instant test task. The name of the task is `task1`. The tested address is `http://www.aliyun.com`. The test type is `HTTP`. The number of detection points is `1`.
//
// @param request - CreateInstantSiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstantSiteMonitorResponse
func (client *Client) CreateInstantSiteMonitorWithContext(ctx context.Context, request *CreateInstantSiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateInstantSiteMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.IspCities) {
		query["IspCities"] = request.IspCities
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.RandomIspCity) {
		query["RandomIspCity"] = request.RandomIspCity
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstantSiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstantSiteMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a blacklist policy.
//
// Description:
//
// ### Background information
//
//   - CloudMonitor blocks alert notifications based on the blacklist policies that take effect. To block alert notifications when the value of a metric that belongs to a cloud service reaches the threshold that you specified, add the metric to a blacklist policy.
//
//   - CloudMonitor allows you to create blacklist policies only based on threshold metrics. You cannot create blacklist policies based on system events. For more information about the cloud services and the thresholds of the metrics that are supported by CloudMonitor, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
//
// @param request - CreateMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricRuleBlackListResponse
func (client *Client) CreateMetricRuleBlackListWithContext(ctx context.Context, request *CreateMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *CreateMetricRuleBlackListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnableEndTime) {
		query["EnableEndTime"] = request.EnableEndTime
	}

	if !dara.IsNil(request.EnableStartTime) {
		query["EnableStartTime"] = request.EnableStartTime
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetricRuleBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates resources with an alert rule.
//
// @param request - CreateMetricRuleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricRuleResourcesResponse
func (client *Client) CreateMetricRuleResourcesWithContext(ctx context.Context, request *CreateMetricRuleResourcesRequest, runtime *dara.RuntimeOptions) (_result *CreateMetricRuleResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetricRuleResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetricRuleResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert template.
//
// @param request - CreateMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricRuleTemplateResponse
func (client *Client) CreateMetricRuleTemplateWithContext(ctx context.Context, request *CreateMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateMetricRuleTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertTemplates) {
		query["AlertTemplates"] = request.AlertTemplates
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetricRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to monitor a process.
//
// @param request - CreateMonitorAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorAgentProcessResponse
func (client *Client) CreateMonitorAgentProcessWithContext(ctx context.Context, request *CreateMonitorAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	if !dara.IsNil(request.ProcessUser) {
		query["ProcessUser"] = request.ProcessUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group.
//
// Description:
//
// In this example, an application group named `ECS_Group` is created.
//
// @param request - CreateMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupResponse
func (client *Client) CreateMonitorGroupWithContext(ctx context.Context, request *CreateMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group by using a resource group.
//
// Description:
//
// This topic provides an example on how to create an application group by using the resource group `CloudMonitor` and the alert contact group `ECS_Group`. The region ID of the resource group is `cn-hangzhou`.
//
// @param request - CreateMonitorGroupByResourceGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupByResourceGroupIdResponse
func (client *Client) CreateMonitorGroupByResourceGroupIdWithContext(ctx context.Context, request *CreateMonitorGroupByResourceGroupIdRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupByResourceGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupList) {
		query["ContactGroupList"] = request.ContactGroupList
	}

	if !dara.IsNil(request.EnableInstallAgent) {
		query["EnableInstallAgent"] = request.EnableInstallAgent
	}

	if !dara.IsNil(request.EnableSubscribeEvent) {
		query["EnableSubscribeEvent"] = request.EnableSubscribeEvent
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupByResourceGroupId"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupByResourceGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds resources to an application group.
//
// Description:
//
// You can add a maximum of 1,000 instances to an application group at a time. You can add a maximum of 3,000 instances of an Alibaba Cloud service to an application group. The total number of instances that you can add to an application group is unlimited.
//
// In this example, an Elastic Compute Service (ECS) instance in the `China (Hangzhou)` region is added to the `3607****` application group. The instance ID is `i-2ze26xj5wwy12****` and the instance name is `test-instance-ecs`.
//
// @param request - CreateMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupInstancesResponse
func (client *Client) CreateMonitorGroupInstancesWithContext(ctx context.Context, request *CreateMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a policy to pause alert notifications for an application group.
//
// Description:
//
// If the policy is valid, no alert notifications are sent for the application group.
//
// This topic describes how to create a `PauseNotify` policy to pause alert notifications for the `7301****` application group. The StartTime parameter is set to `1622949300000` and the EndTime parameter is set to `1623208500000`. This indicates that the policy is valid from `2021-06-06 11:15:00 UTC+8` to `2021-06-09 11:15:00 UTC+8`.
//
// @param request - CreateMonitorGroupNotifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupNotifyPolicyResponse
func (client *Client) CreateMonitorGroupNotifyPolicyWithContext(ctx context.Context, request *CreateMonitorGroupNotifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupNotifyPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupNotifyPolicy"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to monitor a process.
//
// @param request - CreateMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitoringAgentProcessResponse
func (client *Client) CreateMonitoringAgentProcessWithContext(ctx context.Context, request *CreateMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitoringAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	if !dara.IsNil(request.ProcessUser) {
		query["ProcessUser"] = request.ProcessUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitoringAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a site monitoring task.
//
// Description:
//
// This topic provides an example on how to create a site monitoring task named `HanZhou_ECS1`. The URL that is monitored by the task is `https://www.aliyun.com` and the type of the task is `HTTPS`.
//
// @param request - CreateSiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteMonitorResponse
func (client *Client) CreateSiteMonitorWithContext(ctx context.Context, request *CreateSiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateSiteMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.AlertIds) {
		query["AlertIds"] = request.AlertIds
	}

	if !dara.IsNil(request.CustomSchedule) {
		query["CustomSchedule"] = request.CustomSchedule
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCities) {
		query["IspCities"] = request.IspCities
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.VpcConfig) {
		query["VpcConfig"] = request.VpcConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSiteMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Defines the range of monitoring data that you want to export. The Cursor information is returned. When you call the BatchExport operation for the first time, you must specify the Cursor information.
//
// Description:
//
// ### [](#)Prerequisites
//
// Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ### [](#)Background information
//
// You can call this operation to obtain the Cursor information and then call the [BatchExport](https://help.aliyun.com/document_detail/2329847.html) operation to export the monitoring data.
//
// ### [](#)Description
//
// This topic provides an example on how to define the monitoring data of a specified metric for a specified cloud service. In this example, the namespace of the cloud service is set to `acs_ecs_dashboard`, the metric is set to `cpu_idle`, the start time is set to `1641627000000`, and the end time is set to `1641645000000`. The number of idle CPU cores on your Elastic Compute Service (ECS) instances is measured every 60 seconds from 15:30:00, January 8, 2022 to 20:30:00, January 8, 2022. The `Cursor` information is returned.
//
// @param tmpReq - CursorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CursorResponse
func (client *Client) CursorWithContext(ctx context.Context, tmpReq *CursorRequest, runtime *dara.RuntimeOptions) (_result *CursorResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CursorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Matchers) {
		request.MatchersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Matchers, dara.String("Matchers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MatchersShrink) {
		body["Matchers"] = request.MatchersShrink
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Cursor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CursorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert contact.
//
// @param request - DeleteContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactResponse
func (client *Client) DeleteContactWithContext(ctx context.Context, request *DeleteContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContact"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert contact group.
//
// @param request - DeleteContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactGroupResponse
func (client *Client) DeleteContactGroupWithContext(ctx context.Context, request *DeleteContactGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the reported monitoring data of a metric.
//
// @param request - DeleteCustomMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomMetricResponse
func (client *Client) DeleteCustomMetricWithContext(ctx context.Context, request *DeleteCustomMetricRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.UUID) {
		query["UUID"] = request.UUID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomMetric"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tag rule.
//
// @param request - DeleteDynamicTagGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicTagGroupResponse
func (client *Client) DeleteDynamicTagGroupWithContext(ctx context.Context, request *DeleteDynamicTagGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDynamicTagGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicTagRuleId) {
		query["DynamicTagRuleId"] = request.DynamicTagRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDynamicTagGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDynamicTagGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the push channels of an event-triggered alert rule.
//
// @param request - DeleteEventRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRuleTargetsResponse
func (client *Client) DeleteEventRuleTargetsWithContext(ctx context.Context, request *DeleteEventRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRuleTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRuleTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes event-triggered alert rules.
//
// @param request - DeleteEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRulesResponse
func (client *Client) DeleteEventRulesWithContext(ctx context.Context, request *DeleteEventRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a configuration set that is used to export monitoring data.
//
// @param request - DeleteExporterOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExporterOutputResponse
func (client *Client) DeleteExporterOutputWithContext(ctx context.Context, request *DeleteExporterOutputRequest, runtime *dara.RuntimeOptions) (_result *DeleteExporterOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestName) {
		query["DestName"] = request.DestName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExporterOutput"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExporterOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data export rule.
//
// @param request - DeleteExporterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExporterRuleResponse
func (client *Client) DeleteExporterRuleWithContext(ctx context.Context, request *DeleteExporterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteExporterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExporterRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExporterRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a process monitoring task for an application group.
//
// @param request - DeleteGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupMonitoringAgentProcessResponse
func (client *Client) DeleteGroupMonitoringAgentProcessWithContext(ctx context.Context, request *DeleteGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupMonitoringAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes availability monitoring tasks.
//
// @param request - DeleteHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostAvailabilityResponse
func (client *Client) DeleteHostAvailabilityWithContext(ctx context.Context, request *DeleteHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostAvailabilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostAvailabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// Description:
//
// > If a metric import task is created for metrics in a namespace, you cannot delete the namespace unless you delete the task first.
//
// This topic provides an example on how to delete a namespace named `aliyun`. The response shows that the namespace is deleted.
//
// @param request - DeleteHybridMonitorNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridMonitorNamespaceResponse
func (client *Client) DeleteHybridMonitorNamespaceWithContext(ctx context.Context, request *DeleteHybridMonitorNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridMonitorNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridMonitorNamespace"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridMonitorNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Logstore group.
//
// Description:
//
// This topic provides an example on how to delete a Logstore group named `Logstore_test`. The response shows that the Logstore group is deleted.
//
// @param request - DeleteHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridMonitorSLSGroupResponse
func (client *Client) DeleteHybridMonitorSLSGroupWithContext(ctx context.Context, request *DeleteHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridMonitorSLSGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a metric import task for Alibaba Cloud services or a metric for the logs that are imported from Log Service.
//
// Description:
//
// This topic provides an example on how to delete a metric import task whose ID is `36****`. The returned result indicates that the metric import task is deleted.
//
// @param request - DeleteHybridMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridMonitorTaskResponse
func (client *Client) DeleteHybridMonitorTaskWithContext(ctx context.Context, request *DeleteHybridMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridMonitorTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridMonitorTask"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridMonitorTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a log monitoring metric.
//
// @param request - DeleteLogMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogMonitorResponse
func (client *Client) DeleteLogMonitorWithContext(ctx context.Context, request *DeleteLogMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogId) {
		query["LogId"] = request.LogId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple blacklist policies at a time.
//
// @param request - DeleteMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleBlackListResponse
func (client *Client) DeleteMetricRuleBlackListWithContext(ctx context.Context, request *DeleteMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleBlackListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates resources from an alert rule.
//
// @param request - DeleteMetricRuleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleResourcesResponse
func (client *Client) DeleteMetricRuleResourcesWithContext(ctx context.Context, request *DeleteMetricRuleResourcesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the push channels of an alert rule.
//
// @param request - DeleteMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleTargetsResponse
func (client *Client) DeleteMetricRuleTargetsWithContext(ctx context.Context, request *DeleteMetricRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TargetIds) {
		query["TargetIds"] = request.TargetIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert template.
//
// @param request - DeleteMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleTemplateResponse
func (client *Client) DeleteMetricRuleTemplateWithContext(ctx context.Context, request *DeleteMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more alert rules.
//
// @param request - DeleteMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRulesResponse
func (client *Client) DeleteMetricRulesWithContext(ctx context.Context, request *DeleteMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application group.
//
// @param request - DeleteMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupResponse
func (client *Client) DeleteMonitorGroupWithContext(ctx context.Context, request *DeleteMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule that is used to dynamically add the instances of a service to an application group.
//
// @param request - DeleteMonitorGroupDynamicRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupDynamicRuleResponse
func (client *Client) DeleteMonitorGroupDynamicRuleWithContext(ctx context.Context, request *DeleteMonitorGroupDynamicRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupDynamicRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupDynamicRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupDynamicRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes instances from an application group.
//
// @param request - DeleteMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupInstancesResponse
func (client *Client) DeleteMonitorGroupInstancesWithContext(ctx context.Context, request *DeleteMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy that is used to pause alert notifications for an application group.
//
// @param request - DeleteMonitorGroupNotifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupNotifyPolicyResponse
func (client *Client) DeleteMonitorGroupNotifyPolicyWithContext(ctx context.Context, request *DeleteMonitorGroupNotifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupNotifyPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupNotifyPolicy"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables monitoring on a process.
//
// @param request - DeleteMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitoringAgentProcessResponse
func (client *Client) DeleteMonitoringAgentProcessWithContext(ctx context.Context, request *DeleteMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitoringAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProcessId) {
		query["ProcessId"] = request.ProcessId
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitoringAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more site monitoring tasks.
//
// @param request - DeleteSiteMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteMonitorsResponse
func (client *Client) DeleteSiteMonitorsWithContext(ctx context.Context, request *DeleteSiteMonitorsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSiteMonitorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsDeleteAlarms) {
		query["IsDeleteAlarms"] = request.IsDeleteAlarms
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSiteMonitors"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSiteMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of initiative alert rules.
//
// @param request - DescribeActiveMetricRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveMetricRuleListResponse
func (client *Client) DescribeActiveMetricRuleListWithContext(ctx context.Context, request *DescribeActiveMetricRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveMetricRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveMetricRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveMetricRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeAlertHistoryList is deprecated, please use Cms::2019-01-01::DescribeAlertLogList instead.
//
// Summary:
//
// Queries historical alerts.
//
// Description:
//
// This API operation is no longer maintained. We recommend that you call the [DescribeAlertLogList](https://help.aliyun.com/document_detail/201087.html) operation.
//
// @param request - DescribeAlertHistoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertHistoryListResponse
func (client *Client) DescribeAlertHistoryListWithContext(ctx context.Context, request *DescribeAlertHistoryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertHistoryListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ascending) {
		query["Ascending"] = request.Ascending
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertHistoryList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of alert logs.
//
// Description:
//
// Queries the statistics of alert logs.
//
// This topic provides an example on how to query the statistics of alert logs for Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertLogCountResponse
func (client *Client) DescribeAlertLogCountWithContext(ctx context.Context, request *DescribeAlertLogCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertLogCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LastMin) {
		query["LastMin"] = request.LastMin
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertLogCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertLogCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of alert logs that are generated during each interval within a period of time.
//
// Description:
//
// This topic provides an example on how to query the number of alert logs for Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertLogHistogramResponse
func (client *Client) DescribeAlertLogHistogramWithContext(ctx context.Context, request *DescribeAlertLogHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertLogHistogramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LastMin) {
		query["LastMin"] = request.LastMin
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertLogHistogram"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertLogHistogramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert logs.
//
// Description:
//
// You can call the operation to query only the alert logs within the last year.
//
// This topic provides an example to show how to query the alert logs of Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertLogListResponse
func (client *Client) DescribeAlertLogListWithContext(ctx context.Context, request *DescribeAlertLogListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertLogListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LastMin) {
		query["LastMin"] = request.LastMin
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertLogList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertLogListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources for which active alerts are triggered based on an alert rule.
//
// @param request - DescribeAlertingMetricRuleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertingMetricRuleResourcesResponse
func (client *Client) DescribeAlertingMetricRuleResourcesWithContext(ctx context.Context, request *DescribeAlertingMetricRuleResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertingMetricRuleResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertingMetricRuleResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertingMetricRuleResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert groups.
//
// @param request - DescribeContactGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactGroupListResponse
func (client *Client) DescribeContactGroupListWithContext(ctx context.Context, request *DescribeContactGroupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactGroupListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContactGroupList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert contacts.
//
// @param request - DescribeContactListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactListResponse
func (client *Client) DescribeContactListWithContext(ctx context.Context, request *DescribeContactListRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChanelType) {
		query["ChanelType"] = request.ChanelType
	}

	if !dara.IsNil(request.ChanelValue) {
		query["ChanelValue"] = request.ChanelValue
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContactList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert contacts in an alert contact group.
//
// @param request - DescribeContactListByContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactListByContactGroupResponse
func (client *Client) DescribeContactListByContactGroupWithContext(ctx context.Context, request *DescribeContactListByContactGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactListByContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContactListByContactGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactListByContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom event.
//
// @param request - DescribeCustomEventAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEventAttributeResponse
func (client *Client) DescribeCustomEventAttributeWithContext(ctx context.Context, request *DescribeCustomEventAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEventAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEventAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEventAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a custom event occurred within a period of time.
//
// Description:
//
// >  This operation queries the number of times that a custom event occurred for each service.
//
// @param request - DescribeCustomEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEventCountResponse
func (client *Client) DescribeCustomEventCountWithContext(ctx context.Context, request *DescribeCustomEventCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEventCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEventCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEventCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a custom event occurred during each interval within a period of time.
//
// @param request - DescribeCustomEventHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEventHistogramResponse
func (client *Client) DescribeCustomEventHistogramWithContext(ctx context.Context, request *DescribeCustomEventHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEventHistogramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEventHistogram"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEventHistogramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the reported monitoring data.
//
// Description:
//
// >  You can call the DescribeMetricList operation to query the metrics of cloud services. For more information, see [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html).
//
// @param request - DescribeCustomMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomMetricListResponse
func (client *Client) DescribeCustomMetricListWithContext(ctx context.Context, request *DescribeCustomMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomMetricListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomMetricList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomMetricListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag rules.
//
// Description:
//
// This topic provides an example to show how to query tag rules that are related to `tagkey1`. The sample responses indicate that two tag rules are found. The rule IDs are `1536df65-a719-429d-8813-73cc40d7****` and `56e8cebb-b3d7-4a91-9880-78a8c84f****`.
//
// @param request - DescribeDynamicTagRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDynamicTagRuleListResponse
func (client *Client) DescribeDynamicTagRuleListWithContext(ctx context.Context, request *DescribeDynamicTagRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDynamicTagRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicTagRuleId) {
		query["DynamicTagRuleId"] = request.DynamicTagRuleId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.TagRegionId) {
		query["TagRegionId"] = request.TagRegionId
	}

	if !dara.IsNil(request.TagValue) {
		query["TagValue"] = request.TagValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDynamicTagRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDynamicTagRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an event-triggered alert rule.
//
// Description:
//
// This topic provides an example to show how to query the details of an event-triggered alert rule named `testRule`.
//
// @param request - DescribeEventRuleAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventRuleAttributeResponse
func (client *Client) DescribeEventRuleAttributeWithContext(ctx context.Context, request *DescribeEventRuleAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventRuleAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventRuleAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventRuleAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event-triggered alert rules.
//
// @param request - DescribeEventRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventRuleListResponse
func (client *Client) DescribeEventRuleListWithContext(ctx context.Context, request *DescribeEventRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.NamePrefix) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event-triggered alert rules.
//
// Description:
//
// This topic provides an example to show how to query the details of an event-triggered alert rule named `testRule`.
//
// @param request - DescribeEventRuleTargetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventRuleTargetListResponse
func (client *Client) DescribeEventRuleTargetListWithContext(ctx context.Context, request *DescribeEventRuleTargetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventRuleTargetListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventRuleTargetList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventRuleTargetListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configuration sets that are used to export monitoring data.
//
// @param request - DescribeExporterOutputListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExporterOutputListResponse
func (client *Client) DescribeExporterOutputListWithContext(ctx context.Context, request *DescribeExporterOutputListRequest, runtime *dara.RuntimeOptions) (_result *DescribeExporterOutputListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExporterOutputList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExporterOutputListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data export rules.
//
// @param request - DescribeExporterRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExporterRuleListResponse
func (client *Client) DescribeExporterRuleListWithContext(ctx context.Context, request *DescribeExporterRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeExporterRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExporterRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExporterRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the process monitoring tasks for an application group.
//
// Description:
//
// You can create a process monitoring task to monitor all or the specified Elastic Compute Service (ECS) instances in an application group and configure alert rules for the process monitoring task.
//
// @param request - DescribeGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupMonitoringAgentProcessResponse
func (client *Client) DescribeGroupMonitoringAgentProcessWithContext(ctx context.Context, request *DescribeGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupMonitoringAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries availability monitoring tasks.
//
// Description:
//
// This topic provides an example to show how to query all the availability monitoring tasks of your Alibaba Cloud account. The sample responses indicate that the account has one availability monitoring task named `ecs_instance`.
//
// @param request - DescribeHostAvailabilityListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHostAvailabilityListResponse
func (client *Client) DescribeHostAvailabilityListWithContext(ctx context.Context, request *DescribeHostAvailabilityListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHostAvailabilityListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHostAvailabilityList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHostAvailabilityListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data in a namespace.
//
// Description:
//
// ## [](#)Prerequisites
//
// Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## [](#)Limits
//
// The size of monitoring data that is returned in each call cannot exceed 1.5 MB. If the returned data reaches the upper limit, the query fails. You must reset the query conditions.
//
// ## [](#)Operation description
//
// This topic provides an example to show how to query the monitoring data of the `AliyunEcs_cpu_total` metric in the `default-aliyun` namespace from `1653804865` (14:14:25 on May 29, 2022) to `1653805225` (14:20:25 on May 29, 2022).
//
// @param request - DescribeHybridMonitorDataListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorDataListResponse
func (client *Client) DescribeHybridMonitorDataListWithContext(ctx context.Context, request *DescribeHybridMonitorDataListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorDataListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromSQL) {
		query["PromSQL"] = request.PromSQL
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorDataList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorDataListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries namespaces and the details of the related metric import tasks.
//
// Description:
//
// In this example, all namespaces within the current account are queried. The response shows that the current account has only one namespace named `aliyun-test`.
//
// @param request - DescribeHybridMonitorNamespaceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorNamespaceListResponse
func (client *Client) DescribeHybridMonitorNamespaceListWithContext(ctx context.Context, request *DescribeHybridMonitorNamespaceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorNamespaceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowTaskStatistic) {
		query["ShowTaskStatistic"] = request.ShowTaskStatistic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorNamespaceList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorNamespaceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Logstore groups.
//
// Description:
//
// In this example, all Logstore groups within the current account are queried. The response shows that the current account has two Logstore groups: `Logstore_test` and `Logstore_aliyun`.
//
// @param request - DescribeHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorSLSGroupResponse
func (client *Client) DescribeHybridMonitorSLSGroupWithContext(ctx context.Context, request *DescribeHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorSLSGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metric import tasks.
//
// Description:
//
// This topic provides an example on how to query all metric import tasks that belong to the current Alibaba Cloud account. The returned result indicates that the current account has only one metric import task. The metric import task is named `aliyun_task`.
//
// @param request - DescribeHybridMonitorTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorTaskListResponse
func (client *Client) DescribeHybridMonitorTaskListWithContext(ctx context.Context, request *DescribeHybridMonitorTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IncludeAliyunTask) {
		query["IncludeAliyunTask"] = request.IncludeAliyunTask
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorTaskList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a log monitoring metric.
//
// @param request - DescribeLogMonitorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogMonitorAttributeResponse
func (client *Client) DescribeLogMonitorAttributeWithContext(ctx context.Context, request *DescribeLogMonitorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogMonitorAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogMonitorAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogMonitorAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries log monitoring metrics.
//
// @param request - DescribeLogMonitorListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogMonitorListResponse
func (client *Client) DescribeLogMonitorListWithContext(ctx context.Context, request *DescribeLogMonitorListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogMonitorListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchValue) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogMonitorList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogMonitorListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of a metric for a cloud service.
//
// Description:
//
// ### [](#)Limits
//
//   - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. If the free quota is used up, you are automatically charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
//   - Each API operation can be called up to 10 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// ### [](#)Description
//
// >  Different from [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html), the DescribeMetricData operation provides statistical features. You can set the Dimension parameter to `{"instanceId": "i-abcdefgh12****"}` to aggregate all data of your Alibaba Cloud account.
//
// This topic provides an example on how to query the monitoring data of the `cpu_idle` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`.
//
// @param request - DescribeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricDataResponse
func (client *Client) DescribeMetricDataWithContext(ctx context.Context, request *DescribeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricData"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest monitoring data of a metric.
//
// Description:
//
// ### [](#)Limits
//
//   - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. After the free quota is used up, you are charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
//   - Each API operation can be called up to 50 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// >  If `Throttling.User` or `Request was denied due to user flow control` is returned when you call an API operation, the API operation is throttled. For more information about how to handle the issue, see [How do I handle the throttling of a query API?](https://help.aliyun.com/document_detail/2615031.html)
//
// ### [](#)Precautions
//
// The storage duration of the monitoring data of each cloud service is related to the `Period` parameter (statistical period). A larger value of the `Period` parameter indicates that the monitoring data is distributed in a larger time range and the storage duration of the monitoring data is longer. The following list describes the specific relationships:
//
//   - The storage duration is 7 days if the value of the `Period` parameter is less than 60 seconds.
//
//   - The storage duration is 31 days if the value of the `Period` parameter is 60 seconds.
//
//   - The storage duration is 91 days if the value of the `Period` parameter is greater than or equal to 300 seconds.
//
// ### [](#)Operation description
//
// This topic provides an example on how to query the latest monitoring data of the `CPUUtilization` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The returned result indicates that the monitoring data for the instance `i-abcdefgh12****` of the account `123456789876****` is queried at an interval of 60 seconds. The maximum, minimum, and average values of the metric are 100, 93.1, and 99.52.
//
// @param request - DescribeMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLastWithContext(ctx context.Context, request *DescribeMetricLastRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricLastResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricLast"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of a metric for a specified cloud service.
//
// Description:
//
// ## Limits
//
// Each API operation can be called up to 50 times per second. An Alibaba Cloud account and the RAM users within the account share the quota.
//
// >This topic provides an example to show how to query the monitoring data of the `cpu_idle` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The returned result indicates that the monitoring data for the instance `i-abcdefgh12****` of the account `120886317861****` is queried at an interval of 60 seconds. The maximum, minimum, and average values of the metric are 100, 93.1, and 99.52.
//
// @param request - DescribeMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricListWithContext(ctx context.Context, request *DescribeMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of metrics that are supported in CloudMonitor.
//
// Description:
//
// This operation is used together with DescribeMetricList and DescribeMetricLast. For more information, see [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html) and [DescribeMetricLast](https://help.aliyun.com/document_detail/51939.html).
//
// @param request - DescribeMetricMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricMetaListResponse
func (client *Client) DescribeMetricMetaListWithContext(ctx context.Context, request *DescribeMetricMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricMetaListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricMetaList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricMetaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleBlackListResponse
func (client *Client) DescribeMetricRuleBlackListWithContext(ctx context.Context, request *DescribeMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleBlackListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of alert rules in each state.
//
// @param request - DescribeMetricRuleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleCountResponse
func (client *Client) DescribeMetricRuleCountWithContext(ctx context.Context, request *DescribeMetricRuleCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert rules.
//
// Description:
//
// This topic provides an example on how to query all alert rules within your Alibaba Cloud account. The returned result shows that only one alert rule is found. The name of the alert rule is `Rule_01` and the ID is `applyTemplate344cfd42-0f32-4fd6-805a-88d7908a****`.
//
// @param request - DescribeMetricRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleListResponse
func (client *Client) DescribeMetricRuleListWithContext(ctx context.Context, request *DescribeMetricRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertState) {
		query["AlertState"] = request.AlertState
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EnableState) {
		query["EnableState"] = request.EnableState
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources that are associated with a specified alert rule.
//
// Description:
//
// ## Limit
//
// This operation supports only Message Service (MNS) resources.
//
// >This topic provides an example on how to query the resources that are associated with an alert rule whose ID is `ae06917_75a8c43178ab66****`.
//
// @param request - DescribeMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleTargetsResponse
func (client *Client) DescribeMetricRuleTargetsWithContext(ctx context.Context, request *DescribeMetricRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an alert template.
//
// Description:
//
// This topic provides an example on how to query the details of an alert template whose ID is `70****`.
//
// @param request - DescribeMetricRuleTemplateAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleTemplateAttributeResponse
func (client *Client) DescribeMetricRuleTemplateAttributeWithContext(ctx context.Context, request *DescribeMetricRuleTemplateAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleTemplateAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleTemplateAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleTemplateAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert templates.
//
// Description:
//
// This topic provides an example on how to query alert templates. In this example, the following alert templates are returned in the response: `ECS_Template1` and `ECS_Template2`.
//
// @param request - DescribeMetricRuleTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleTemplateListResponse
func (client *Client) DescribeMetricRuleTemplateListWithContext(ctx context.Context, request *DescribeMetricRuleTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleTemplateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.History) {
		query["History"] = request.History
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleTemplateList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest monitoring data of a metric for a cloud service. The data can be sorted by a specified order.
//
// Description:
//
// ### [](#)Limits
//
//   - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. After the free quota is used up, you are charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
//   - Each API operation can be called up to 10 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// ### [](#)Precautions
//
// The storage duration of the monitoring data of each cloud service is related to the `Period` parameter (statistical period). A larger value of the `Period` parameter indicates that the monitoring data is distributed in a larger time range and the storage duration of the monitoring data is longer. The following list describes the specific relationships:
//
//   - The storage duration is 7 days if the value of the `Period` parameter is less than 60 seconds.
//
//   - The storage duration is 31 days if the value of the `Period` parameter is 60 seconds.
//
//   - The storage duration is 91 days if the value of the `Period` is greater than or equal to 300 seconds.
//
// ### [](#)Operation description
//
// This topic provides an example on how to query the monitoring data of the `cpu_idle` metric in the last 60 seconds for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The monitoring data is sorted in descending order based on the `Average` field.
//
// @param request - DescribeMetricTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricTopResponse
func (client *Client) DescribeMetricTopWithContext(ctx context.Context, request *DescribeMetricTopRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricTopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OrderDesc) {
		query["OrderDesc"] = request.OrderDesc
	}

	if !dara.IsNil(request.Orderby) {
		query["Orderby"] = request.Orderby
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricTop"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud services to which the resources in an application group belong and the number of resources that belong to each cloud service in the application group.
//
// @param request - DescribeMonitorGroupCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupCategoriesResponse
func (client *Client) DescribeMonitorGroupCategoriesWithContext(ctx context.Context, request *DescribeMonitorGroupCategoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupCategoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupCategories"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupCategoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dynamic rules of an application group.
//
// @param request - DescribeMonitorGroupDynamicRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupDynamicRulesResponse
func (client *Client) DescribeMonitorGroupDynamicRulesWithContext(ctx context.Context, request *DescribeMonitorGroupDynamicRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupDynamicRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupDynamicRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupDynamicRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the resources in an application group.
//
// @param request - DescribeMonitorGroupInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupInstanceAttributeResponse
func (client *Client) DescribeMonitorGroupInstanceAttributeWithContext(ctx context.Context, request *DescribeMonitorGroupInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Total) {
		query["Total"] = request.Total
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupInstanceAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources in an application group.
//
// @param request - DescribeMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupInstancesResponse
func (client *Client) DescribeMonitorGroupInstancesWithContext(ctx context.Context, request *DescribeMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the policies that are used to pause alert notifications for an application group.
//
// @param request - DescribeMonitorGroupNotifyPolicyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupNotifyPolicyListResponse
func (client *Client) DescribeMonitorGroupNotifyPolicyListWithContext(ctx context.Context, request *DescribeMonitorGroupNotifyPolicyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupNotifyPolicyListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupNotifyPolicyList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupNotifyPolicyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries application groups.
//
// Description:
//
// This topic provides an example of how to query the application groups of the current account. The response shows that the current account has two application groups: `testGroup124` and `test123`.
//
// @param request - DescribeMonitorGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupsResponse
func (client *Client) DescribeMonitorGroupsWithContext(ctx context.Context, request *DescribeMonitorGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicTagRuleId) {
		query["DynamicTagRuleId"] = request.DynamicTagRuleId
	}

	if !dara.IsNil(request.GroupFounderTagKey) {
		query["GroupFounderTagKey"] = request.GroupFounderTagKey
	}

	if !dara.IsNil(request.GroupFounderTagValue) {
		query["GroupFounderTagValue"] = request.GroupFounderTagValue
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IncludeTemplateHistory) {
		query["IncludeTemplateHistory"] = request.IncludeTemplateHistory
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SelectContactGroups) {
		query["SelectContactGroups"] = request.SelectContactGroups
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroups"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource quotas of CloudMonitor.
//
// @param request - DescribeMonitorResourceQuotaAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorResourceQuotaAttributeResponse
func (client *Client) DescribeMonitorResourceQuotaAttributeWithContext(ctx context.Context, request *DescribeMonitorResourceQuotaAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorResourceQuotaAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ShowUsed) {
		query["ShowUsed"] = request.ShowUsed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorResourceQuotaAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorResourceQuotaAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey ID and AccessKey secret that are required to install the CloudMonitor agent on a third-party host.
//
// @param request - DescribeMonitoringAgentAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentAccessKeyResponse
func (client *Client) DescribeMonitoringAgentAccessKeyWithContext(ctx context.Context, request *DescribeMonitoringAgentAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentAccessKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentAccessKey"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentAccessKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the CloudMonitor agent.
//
// @param request - DescribeMonitoringAgentConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentConfigResponse
func (client *Client) DescribeMonitoringAgentConfigWithContext(ctx context.Context, request *DescribeMonitoringAgentConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all hosts within the current Alibaba Cloud account, including hosts on which the CloudMonitor agent is installed and uninstalled.
//
// @param request - DescribeMonitoringAgentHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentHostsResponse
func (client *Client) DescribeMonitoringAgentHostsWithContext(ctx context.Context, request *DescribeMonitoringAgentHostsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentHostsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunHost) {
		query["AliyunHost"] = request.AliyunHost
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceRegionId) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNumbers) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SysomStatus) {
		query["SysomStatus"] = request.SysomStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentHosts"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentHostsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定资源的进程数列表
//
// Description:
//
// >  Before you call this operation, call the CreateMonitoringAgentProcess operation to create processes. For more information, see [CreateMonitoringAgentProcess](https://help.aliyun.com/document_detail/114951.html~).
//
// This topic provides an example of how to query the processes of the `i-hp3hl3cx1pbahzy8****` instance. The response indicates the details of the `NGINX` and `HTTP` processes.
//
// @param request - DescribeMonitoringAgentProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentProcessesResponse
func (client *Client) DescribeMonitoringAgentProcessesWithContext(ctx context.Context, request *DescribeMonitoringAgentProcessesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentProcessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentProcesses"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the CloudMonitor agent.
//
// Description:
//
// This topic describes how to query the status of the CloudMonitor agent that is installed on the `i-hp3dunahluwajv6f****` instance. The result indicates that the CloudMonitor agent is in the `running` state.
//
// @param request - DescribeMonitoringAgentStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentStatusesResponse
func (client *Client) DescribeMonitoringAgentStatusesWithContext(ctx context.Context, request *DescribeMonitoringAgentStatusesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentStatusesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAvailabilityTaskId) {
		query["HostAvailabilityTaskId"] = request.HostAvailabilityTaskId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentStatuses"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentStatusesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global configurations of the CloudMonitor agent.
//
// @param request - DescribeMonitoringConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringConfigResponse
func (client *Client) DescribeMonitoringConfigWithContext(ctx context.Context, request *DescribeMonitoringConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the keys of all tags that are attached to cloud resources in a region.
//
// Description:
//
// >  If a tag is attached to multiple cloud resources in the region, the key of the tag is returned only once.
//
// @param request - DescribeProductResourceTagKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductResourceTagKeyListResponse
func (client *Client) DescribeProductResourceTagKeyListWithContext(ctx context.Context, request *DescribeProductResourceTagKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductResourceTagKeyListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductResourceTagKeyList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductResourceTagKeyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud services for which the initiative alert feature is enabled.
//
// @param request - DescribeProductsOfActiveMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductsOfActiveMetricRuleResponse
func (client *Client) DescribeProductsOfActiveMetricRuleWithContext(ctx context.Context, request *DescribeProductsOfActiveMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductsOfActiveMetricRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductsOfActiveMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductsOfActiveMetricRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about monitored services in CloudMonitor.
//
// Description:
//
// The information obtained by this operation includes the service description, namespace, and tags.
//
// @param request - DescribeProjectMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectMetaResponse
func (client *Client) DescribeProjectMetaWithContext(ctx context.Context, request *DescribeProjectMetaRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectMeta"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a site monitoring task.
//
// Description:
//
// This topic provides an example on how to query the details of a site monitoring task whose ID is `cc641dff-c19d-45f3-ad0a-818a0c4f****`. The returned result indicates that the task name is `test123`, the URL that is monitored by the task is `https://aliyun.com`, and the name of the carrier is `Alibaba`.
//
// @param request - DescribeSiteMonitorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorAttributeResponse
func (client *Client) DescribeSiteMonitorAttributeWithContext(ctx context.Context, request *DescribeSiteMonitorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeAlert) {
		query["IncludeAlert"] = request.IncludeAlert
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fine-grained monitoring data of a site monitoring task.
//
// @param request - DescribeSiteMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorDataResponse
func (client *Client) DescribeSiteMonitorDataWithContext(ctx context.Context, request *DescribeSiteMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorData"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detection points that are provided by carriers.
//
// Description:
//
// This topic provides an example on how to query the detection points that are provided by China Unicom in Guiyang.
//
// @param request - DescribeSiteMonitorISPCityListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorISPCityListResponse
func (client *Client) DescribeSiteMonitorISPCityListWithContext(ctx context.Context, request *DescribeSiteMonitorISPCityListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorISPCityListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.IPV4) {
		query["IPV4"] = request.IPV4
	}

	if !dara.IsNil(request.IPV6) {
		query["IPV6"] = request.IPV6
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.ViewAll) {
		query["ViewAll"] = request.ViewAll
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorISPCityList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorISPCityListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries site monitoring tasks.
//
// Description:
//
// This topic provides an example on how to query all the site monitoring tasks of your Alibaba Cloud account. In this example, the returned result indicates that the Alibaba Cloud account has one site monitoring task named `HanZhou_ECS2`.
//
// @param request - DescribeSiteMonitorListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorListResponse
func (client *Client) DescribeSiteMonitorListWithContext(ctx context.Context, request *DescribeSiteMonitorListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskState) {
		query["TaskState"] = request.TaskState
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of one or more instant test tasks.
//
// Description:
//
// You can create an instant test task only by using the Alibaba Cloud account that you used to enable Network Analysis and Monitoring.
//
// This topic provides an example to show how to query the logs of an instant test task whose ID is `afa5c3ce-f944-4363-9edb-ce919a29****`.
//
// @param request - DescribeSiteMonitorLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorLogResponse
func (client *Client) DescribeSiteMonitorLogWithContext(ctx context.Context, request *DescribeSiteMonitorLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Browser) {
		query["Browser"] = request.Browser
	}

	if !dara.IsNil(request.BrowserInfo) {
		query["BrowserInfo"] = request.BrowserInfo
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Device) {
		query["Device"] = request.Device
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorLog"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas and version of site monitoring.
//
// @param request - DescribeSiteMonitorQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorQuotaResponse
func (client *Client) DescribeSiteMonitorQuotaWithContext(ctx context.Context, request *DescribeSiteMonitorQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorQuota"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of a specified metric for a specified site monitoring task.
//
// Description:
//
// This topic provides an example on how to query the statistics of the `Availability` metric for a site monitoring task whose ID is `ef4cdc8b-9dc7-43e7-810e-f950e56c****`. The result indicates that the availability rate of the site is `100%`.
//
// @param request - DescribeSiteMonitorStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorStatisticsResponse
func (client *Client) DescribeSiteMonitorStatisticsWithContext(ctx context.Context, request *DescribeSiteMonitorStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorStatistics"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询拨测探测节点列表
//
// @param request - DescribeSyntheticProbeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyntheticProbeListResponse
func (client *Client) DescribeSyntheticProbeListWithContext(ctx context.Context, request *DescribeSyntheticProbeListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSyntheticProbeListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.IdcProbe) {
		query["IdcProbe"] = request.IdcProbe
	}

	if !dara.IsNil(request.Ipv4) {
		query["Ipv4"] = request.Ipv4
	}

	if !dara.IsNil(request.Ipv6) {
		query["Ipv6"] = request.Ipv6
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.LmProbe) {
		query["LmProbe"] = request.LmProbe
	}

	if !dara.IsNil(request.MbProbe) {
		query["MbProbe"] = request.MbProbe
	}

	if !dara.IsNil(request.ViewAll) {
		query["ViewAll"] = request.ViewAll
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSyntheticProbeList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSyntheticProbeListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a system event.
//
// @param request - DescribeSystemEventAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventAttributeResponse
func (client *Client) DescribeSystemEventAttributeWithContext(ctx context.Context, request *DescribeSystemEventAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a system event of a cloud service has occurred.
//
// Description:
//
// ### [](#)Background information
//
// You can call the [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) operation to query the cloud services supported by CloudMonitor and their system events.
//
// ### [](#)Description
//
// This topic provides an example on how to query the number of times that a system event of `Elastic Compute Service (ECS)` has occurred. The returned result shows that the specified system event has occurred three times.
//
// @param request - DescribeSystemEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventCountResponse
func (client *Client) DescribeSystemEventCountWithContext(ctx context.Context, request *DescribeSystemEventCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times a system event occurred during each interval within a period of time.
//
// @param request - DescribeSystemEventHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventHistogramResponse
func (client *Client) DescribeSystemEventHistogramWithContext(ctx context.Context, request *DescribeSystemEventHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventHistogramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventHistogram"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventHistogramResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the meta information about system events.
//
// @param request - DescribeSystemEventMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventMetaListResponse
func (client *Client) DescribeSystemEventMetaListWithContext(ctx context.Context, request *DescribeSystemEventMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventMetaListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventMetaList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventMetaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - DescribeTagKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagKeyListResponse
func (client *Client) DescribeTagKeyListWithContext(ctx context.Context, request *DescribeTagKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagKeyListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagKeyList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagKeyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values corresponding to a specified tag key.
//
// Description:
//
// This topic provides an example of how to query the tag values corresponding to `tagKey1`. The return results are `tagValue1` and `tagValue2`.
//
// @param request - DescribeTagValueListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagValueListResponse
func (client *Client) DescribeTagValueListWithContext(ctx context.Context, request *DescribeTagValueListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagValueListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagValueList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagValueListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries unhealthy instances detected by availability monitoring tasks.
//
// @param request - DescribeUnhealthyHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnhealthyHostAvailabilityResponse
func (client *Client) DescribeUnhealthyHostAvailabilityWithContext(ctx context.Context, request *DescribeUnhealthyHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnhealthyHostAvailabilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUnhealthyHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnhealthyHostAvailabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the initiative alert feature for a cloud service.
//
// @param request - DisableActiveMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableActiveMetricRuleResponse
func (client *Client) DisableActiveMetricRuleWithContext(ctx context.Context, request *DisableActiveMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableActiveMetricRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableActiveMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableActiveMetricRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableEventRulesResponse
func (client *Client) DisableEventRulesWithContext(ctx context.Context, request *DisableEventRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableEventRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableEventRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableEventRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables availability monitoring tasks.
//
// @param request - DisableHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableHostAvailabilityResponse
func (client *Client) DisableHostAvailabilityWithContext(ctx context.Context, request *DisableHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *DisableHostAvailabilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableHostAvailabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables alert rules.
//
// @param request - DisableMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableMetricRulesResponse
func (client *Client) DisableMetricRulesWithContext(ctx context.Context, request *DisableMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableMetricRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableMetricRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables site monitoring tasks.
//
// @param request - DisableSiteMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSiteMonitorsResponse
func (client *Client) DisableSiteMonitorsWithContext(ctx context.Context, request *DisableSiteMonitorsRequest, runtime *dara.RuntimeOptions) (_result *DisableSiteMonitorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSiteMonitors"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSiteMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the initiative alert feature for a cloud service.
//
// @param request - EnableActiveMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableActiveMetricRuleResponse
func (client *Client) EnableActiveMetricRuleWithContext(ctx context.Context, request *EnableActiveMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableActiveMetricRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableActiveMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableActiveMetricRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableEventRulesResponse
func (client *Client) EnableEventRulesWithContext(ctx context.Context, request *EnableEventRulesRequest, runtime *dara.RuntimeOptions) (_result *EnableEventRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableEventRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableEventRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables availability monitoring tasks.
//
// @param request - EnableHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHostAvailabilityResponse
func (client *Client) EnableHostAvailabilityWithContext(ctx context.Context, request *EnableHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *EnableHostAvailabilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHostAvailabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables multiple blacklist policies at a time.
//
// @param request - EnableMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMetricRuleBlackListResponse
func (client *Client) EnableMetricRuleBlackListWithContext(ctx context.Context, request *EnableMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *EnableMetricRuleBlackListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMetricRuleBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables alert rules.
//
// @param request - EnableMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMetricRulesResponse
func (client *Client) EnableMetricRulesWithContext(ctx context.Context, request *EnableMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *EnableMetricRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMetricRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables site monitoring tasks.
//
// @param request - EnableSiteMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSiteMonitorsResponse
func (client *Client) EnableSiteMonitorsWithContext(ctx context.Context, request *EnableSiteMonitorsRequest, runtime *dara.RuntimeOptions) (_result *EnableSiteMonitorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSiteMonitors"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSiteMonitorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # InstallMonitoringAgent
//
// Description:
//
// ## Prerequisites
//
// The Cloud Assistant client is installed on an ECS instance. For more information about how to install the Cloud Assistant client, see [Overview](https://help.aliyun.com/document_detail/64601.html).
//
// @param request - InstallMonitoringAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallMonitoringAgentResponse
func (client *Client) InstallMonitoringAgentWithContext(ctx context.Context, request *InstallMonitoringAgentRequest, runtime *dara.RuntimeOptions) (_result *InstallMonitoringAgentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstallCommand) {
		query["InstallCommand"] = request.InstallCommand
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallMonitoringAgent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallMonitoringAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用分组内的进程监控
//
// @param request - ModifyGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGroupMonitoringAgentProcessResponse
func (client *Client) ModifyGroupMonitoringAgentProcessWithContext(ctx context.Context, request *ModifyGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *ModifyGroupMonitoringAgentProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MatchExpressFilterRelation) {
		query["MatchExpressFilterRelation"] = request.MatchExpressFilterRelation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an availability monitoring task.
//
// Description:
//
// This topic provides an example on how to change the name of an availability monitoring task named `12345` in an application group named `123456` to `task2`.
//
// @param request - ModifyHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostAvailabilityResponse
func (client *Client) ModifyHostAvailabilityWithContext(ctx context.Context, request *ModifyHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostAvailabilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfigEscalationList) {
		query["AlertConfigEscalationList"] = request.AlertConfigEscalationList
	}

	if !dara.IsNil(request.AlertConfigTargetList) {
		query["AlertConfigTargetList"] = request.AlertConfigTargetList
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskScope) {
		query["TaskScope"] = request.TaskScope
	}

	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.TaskOption) {
		query["TaskOption"] = request.TaskOption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostAvailabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改非阿里云的主机显示信息
//
// Description:
//
// ***
//
// @param request - ModifyHostInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostInfoResponse
func (client *Client) ModifyHostInfoWithContext(ctx context.Context, request *ModifyHostInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostInfo"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a namespace.
//
// Description:
//
// This topic provides an example on how to change the data retention period of the `aliyun` namespace to `cms.s1.2xlarge`. The response shows that the namespace is modified.
//
// @param request - ModifyHybridMonitorNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridMonitorNamespaceResponse
func (client *Client) ModifyHybridMonitorNamespaceWithContext(ctx context.Context, request *ModifyHybridMonitorNamespaceRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridMonitorNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridMonitorNamespace"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridMonitorNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Logstore group.
//
// Description:
//
// In this example, a Logstore group named `Logstore_test` is modified. The Logstore of the `aliyun-project` project in the `cn-hangzhou` region is changed to `Logstore-aliyun-all`. The response shows that the Logstore group is modified.
//
// @param request - ModifyHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridMonitorSLSGroupResponse
func (client *Client) ModifyHybridMonitorSLSGroupWithContext(ctx context.Context, request *ModifyHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridMonitorSLSGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SLSGroupConfig) {
		query["SLSGroupConfig"] = request.SLSGroupConfig
	}

	if !dara.IsNil(request.SLSGroupDescription) {
		query["SLSGroupDescription"] = request.SLSGroupDescription
	}

	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a metric for the logs that are imported from Log Service.
//
// Description:
//
// This topic provides an example on how to change the collection period of a metric import task whose ID is `36****` to `15` seconds. The task is used to monitor the logs that are imported from Log Service. The returned result indicates that the metric is modified.
//
// @param request - ModifyHybridMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridMonitorTaskResponse
func (client *Client) ModifyHybridMonitorTaskWithContext(ctx context.Context, request *ModifyHybridMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridMonitorTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachLabels) {
		query["AttachLabels"] = request.AttachLabels
	}

	if !dara.IsNil(request.CollectInterval) {
		query["CollectInterval"] = request.CollectInterval
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SLSProcessConfig) {
		query["SLSProcessConfig"] = request.SLSProcessConfig
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridMonitorTask"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridMonitorTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a blacklist policy.
//
// @param request - ModifyMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMetricRuleBlackListResponse
func (client *Client) ModifyMetricRuleBlackListWithContext(ctx context.Context, request *ModifyMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *ModifyMetricRuleBlackListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnableEndTime) {
		query["EnableEndTime"] = request.EnableEndTime
	}

	if !dara.IsNil(request.EnableStartTime) {
		query["EnableStartTime"] = request.EnableStartTime
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMetricRuleBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an alert template.
//
// Description:
//
// This topic provides an example on how to modify an alert template whose version is `1` and ID is `123456`. The alert level is changed to `Critical`. The statistical method is changed to `Average`. The alert threshold comparator is changed to `GreaterThanOrEqualToThreshold`. The alert threshold is changed to `90`. The number of alert retries is changed to `3`. The response shows that the alert template is modified.
//
// @param request - ModifyMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMetricRuleTemplateResponse
func (client *Client) ModifyMetricRuleTemplateWithContext(ctx context.Context, request *ModifyMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyMetricRuleTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertTemplates) {
		query["AlertTemplates"] = request.AlertTemplates
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RestVersion) {
		query["RestVersion"] = request.RestVersion
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMetricRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用分组
//
// @param request - ModifyMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMonitorGroupResponse
func (client *Client) ModifyMonitorGroupWithContext(ctx context.Context, request *ModifyMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyMonitorGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMonitorGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMonitorGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用分组中的资源
//
// @param request - ModifyMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMonitorGroupInstancesResponse
func (client *Client) ModifyMonitorGroupInstancesWithContext(ctx context.Context, request *ModifyMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *ModifyMonitorGroupInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMonitorGroupInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a site monitoring task.
//
// Description:
//
// The number of site monitoring tasks.
//
// @param request - ModifySiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySiteMonitorResponse
func (client *Client) ModifySiteMonitorWithContext(ctx context.Context, request *ModifySiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *ModifySiteMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AlertIds) {
		query["AlertIds"] = request.AlertIds
	}

	if !dara.IsNil(request.CustomSchedule) {
		query["CustomSchedule"] = request.CustomSchedule
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IntervalUnit) {
		query["IntervalUnit"] = request.IntervalUnit
	}

	if !dara.IsNil(request.IspCities) {
		query["IspCities"] = request.IspCities
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySiteMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact.
//
// @param request - PutContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutContactResponse
func (client *Client) PutContactWithContext(ctx context.Context, request *PutContactRequest, runtime *dara.RuntimeOptions) (_result *PutContactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Channels) {
		query["Channels"] = request.Channels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutContact"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact group.
//
// Description:
//
// This topic provides an example on how to create an alert contact group named `ECS_Group`.
//
// @param request - PutContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutContactGroupResponse
func (client *Client) PutContactGroupWithContext(ctx context.Context, request *PutContactGroupRequest, runtime *dara.RuntimeOptions) (_result *PutContactGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactNames) {
		query["ContactNames"] = request.ContactNames
	}

	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.EnableSubscribed) {
		query["EnableSubscribed"] = request.EnableSubscribed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutContactGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports custom events.
//
// @param request - PutCustomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomEventResponse
func (client *Client) PutCustomEventWithContext(ctx context.Context, request *PutCustomEventRequest, runtime *dara.RuntimeOptions) (_result *PutCustomEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventInfo) {
		query["EventInfo"] = request.EventInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomEvent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Before you call this operation, call the PutCustomEvent operation to report the monitoring data of the custom event. For more information, see [PutCustomEvent](https://help.aliyun.com/document_detail/115012.html).
//
// @param request - PutCustomEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomEventRuleResponse
func (client *Client) PutCustomEventRuleWithContext(ctx context.Context, request *PutCustomEventRuleRequest, runtime *dara.RuntimeOptions) (_result *PutCustomEventRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomEventRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomEventRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports monitoring data.
//
// Description:
//
// >  We recommend that you call the [PutHybridMonitorMetricData](https://help.aliyun.com/document_detail/383455.html) operation of Hybrid Cloud Monitoring to report monitoring data.
//
// @param request - PutCustomMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomMetricResponse
func (client *Client) PutCustomMetricWithContext(ctx context.Context, request *PutCustomMetricRequest, runtime *dara.RuntimeOptions) (_result *PutCustomMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MetricList) {
		query["MetricList"] = request.MetricList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomMetric"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom alert rule.
//
// Description:
//
// Before you call this operation, call the PutCustomMetric operation to report custom monitoring data. For more information, see [PutCustomMetric](https://help.aliyun.com/document_detail/115004.html).
//
// @param request - PutCustomMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomMetricRuleResponse
func (client *Client) PutCustomMetricRuleWithContext(ctx context.Context, request *PutCustomMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *PutCustomMetricRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComparisonOperator) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.Statistics) {
		query["Statistics"] = request.Statistics
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomMetricRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或者修改事件监控
//
// Description:
//
// If the specified rule name does not exist, an event-triggered alert rule is created. If the specified rule name exists, the specified event-triggered alert rule is modified.
//
// In this example, the `myRuleName` alert rule is created for the `ecs` cloud service.
//
// @param request - PutEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventRuleResponse
func (client *Client) PutEventRuleWithContext(ctx context.Context, request *PutEventRuleRequest, runtime *dara.RuntimeOptions) (_result *PutEventRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventPattern) {
		query["EventPattern"] = request.EventPattern
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEventRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEventRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies the push channels of an event-triggered alert rule.
//
// @param request - PutEventRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventRuleTargetsResponse
func (client *Client) PutEventRuleTargetsWithContext(ctx context.Context, request *PutEventRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *PutEventRuleTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactParameters) {
		query["ContactParameters"] = request.ContactParameters
	}

	if !dara.IsNil(request.FcParameters) {
		query["FcParameters"] = request.FcParameters
	}

	if !dara.IsNil(request.MnsParameters) {
		query["MnsParameters"] = request.MnsParameters
	}

	if !dara.IsNil(request.OpenApiParameters) {
		query["OpenApiParameters"] = request.OpenApiParameters
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SlsParameters) {
		query["SlsParameters"] = request.SlsParameters
	}

	if !dara.IsNil(request.WebhookParameters) {
		query["WebhookParameters"] = request.WebhookParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEventRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEventRuleTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a configuration set for exporting monitoring data.
//
// Description:
//
// > The monitoring data can be exported only to Log Service. More services will be supported in the future.
//
// @param request - PutExporterOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutExporterOutputResponse
func (client *Client) PutExporterOutputWithContext(ctx context.Context, request *PutExporterOutputRequest, runtime *dara.RuntimeOptions) (_result *PutExporterOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigJson) {
		query["ConfigJson"] = request.ConfigJson
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.DestName) {
		query["DestName"] = request.DestName
	}

	if !dara.IsNil(request.DestType) {
		query["DestType"] = request.DestType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutExporterOutput"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutExporterOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a data export rule.
//
// @param request - PutExporterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutExporterRuleResponse
func (client *Client) PutExporterRuleWithContext(ctx context.Context, request *PutExporterRuleRequest, runtime *dara.RuntimeOptions) (_result *PutExporterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.DstNames) {
		query["DstNames"] = request.DstNames
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.TargetWindows) {
		query["TargetWindows"] = request.TargetWindows
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutExporterRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutExporterRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule for an application group.
//
// Description:
//
// This topic provides an example on how to create an alert rule for the `cpu_total` metric of Elastic Compute Service (ECS) in the `17285****` application group. The ID of the alert rule is `123456`. The name of the alert rule is `Rule_test`. The alert level is `Critical`. The statistical method is `Average`. The alert threshold comparator is `GreaterThanOrEqualToThreshold`. The alert threshold is `90`. The number of alert retries is `3`. The returned result shows that the alert rule is created and the alert rule ID is `123456`.
//
// @param request - PutGroupMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutGroupMetricRuleResponse
func (client *Client) PutGroupMetricRuleWithContext(ctx context.Context, request *PutGroupMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *PutGroupMetricRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.ExtraDimensionJson) {
		query["ExtraDimensionJson"] = request.ExtraDimensionJson
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NoDataPolicy) {
		query["NoDataPolicy"] = request.NoDataPolicy
	}

	if !dara.IsNil(request.NoEffectiveInterval) {
		query["NoEffectiveInterval"] = request.NoEffectiveInterval
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	if !dara.IsNil(request.Escalations) {
		query["Escalations"] = request.Escalations
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutGroupMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutGroupMetricRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports the monitoring data of a metric to a namespace of Hybrid Cloud Monitoring.
//
// Description:
//
// ## [](#)Prerequisites
//
// Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## [](#)Limits
//
// The size of the monitoring data that you import at a time must be less than or equal to 1 MB.
//
// ## [](#)Operation description
//
// This topic provides an example on how to import the monitoring data of the `CPU_Usage` metric to the `default-aliyun` namespace of Hybrid Cloud Monitoring.
//
// @param request - PutHybridMonitorMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutHybridMonitorMetricDataResponse
func (client *Client) PutHybridMonitorMetricDataWithContext(ctx context.Context, request *PutHybridMonitorMetricDataRequest, runtime *dara.RuntimeOptions) (_result *PutHybridMonitorMetricDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MetricList) {
		query["MetricList"] = request.MetricList
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutHybridMonitorMetricData"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutHybridMonitorMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a log monitoring metric.
//
// Description:
//
// In the example of this topic, the `cpu_total` log monitoring metric is created. The response shows that the log monitoring metric is created and the metric ID is `16****`.
//
// @param request - PutLogMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutLogMonitorResponse
func (client *Client) PutLogMonitorWithContext(ctx context.Context, request *PutLogMonitorRequest, runtime *dara.RuntimeOptions) (_result *PutLogMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aggregates) {
		query["Aggregates"] = request.Aggregates
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Groupbys) {
		query["Groupbys"] = request.Groupbys
	}

	if !dara.IsNil(request.LogId) {
		query["LogId"] = request.LogId
	}

	if !dara.IsNil(request.MetricExpress) {
		query["MetricExpress"] = request.MetricExpress
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.SlsLogstore) {
		query["SlsLogstore"] = request.SlsLogstore
	}

	if !dara.IsNil(request.SlsProject) {
		query["SlsProject"] = request.SlsProject
	}

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	if !dara.IsNil(request.Tumblingwindows) {
		query["Tumblingwindows"] = request.Tumblingwindows
	}

	if !dara.IsNil(request.Unit) {
		query["Unit"] = request.Unit
	}

	if !dara.IsNil(request.ValueFilter) {
		query["ValueFilter"] = request.ValueFilter
	}

	if !dara.IsNil(request.ValueFilterRelation) {
		query["ValueFilterRelation"] = request.ValueFilterRelation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutLogMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutLogMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies the push channels of an alert rule.
//
// Description:
//
// # [](#)
//
// This topic provides an example on how to associate an alert rule with a resource. In this example, the alert rule is `ae06917_75a8c43178ab66****`, the resource is `acs:mns:cn-hangzhou:120886317861****:/queues/test/message`, and the ID of the resource for which alerts are triggered is `1`. The response indicates that the resource is associated with the specified alert rule.
//
// @param request - PutMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMetricRuleTargetsResponse
func (client *Client) PutMetricRuleTargetsWithContext(ctx context.Context, request *PutMetricRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *PutMetricRuleTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMetricRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule to dynamically add instances that meet the rule to an application group.
//
// @param request - PutMonitorGroupDynamicRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMonitorGroupDynamicRuleResponse
func (client *Client) PutMonitorGroupDynamicRuleWithContext(ctx context.Context, request *PutMonitorGroupDynamicRuleRequest, runtime *dara.RuntimeOptions) (_result *PutMonitorGroupDynamicRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupRules) {
		query["GroupRules"] = request.GroupRules
	}

	if !dara.IsNil(request.IsAsync) {
		query["IsAsync"] = request.IsAsync
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMonitorGroupDynamicRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMonitorGroupDynamicRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures global settings for the CloudMonitor agent.
//
// @param request - PutMonitoringConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMonitoringConfigResponse
func (client *Client) PutMonitoringConfigWithContext(ctx context.Context, request *PutMonitoringConfigRequest, runtime *dara.RuntimeOptions) (_result *PutMonitoringConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoInstall) {
		query["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.EnableInstallAgentNewECS) {
		query["EnableInstallAgentNewECS"] = request.EnableInstallAgentNewECS
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMonitoringConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMonitoringConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an alert rule.
//
// Description:
//
// This topic provides an example on how to create a threshold-triggered alert rule for the `cpu_total` metric of an Elastic Compute Service (ECS) instance whose ID is `i-uf6j91r34rnwawoo****`. The namespace of ECS metrics is `acs_ecs_dashboard`. The alert contact group of the alert rule is `ECS_Group`. The name of the alert rule is `test123`. The ID of the alert rule is `a151cd6023eacee2f0978e03863cc1697c89508****`. The statistical method for Critical-level alerts is `Average`. The comparison operator for Critical-level alerts is `GreaterThanOrEqualToThreshold`. The threshold for Critical-level alerts is `90`. The consecutive number of times for which the metric value meets the trigger condition before a Critical-level alert is triggered is `3`.
//
// >  Statistics verification was added on August 15, 2024. Only the statistical value of the corresponding metric can be set for the Statistics parameter. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://www.alibabacloud.com/help/en/cms/support/appendix-1-metrics).
//
// @param tmpReq - PutResourceMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutResourceMetricRuleResponse
func (client *Client) PutResourceMetricRuleWithContext(ctx context.Context, tmpReq *PutResourceMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *PutResourceMetricRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PutResourceMetricRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CompositeExpression) {
		request.CompositeExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompositeExpression, dara.String("CompositeExpression"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Prometheus) {
		request.PrometheusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Prometheus, dara.String("Prometheus"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompositeExpressionShrink) {
		query["CompositeExpression"] = request.CompositeExpressionShrink
	}

	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NoDataPolicy) {
		query["NoDataPolicy"] = request.NoDataPolicy
	}

	if !dara.IsNil(request.NoEffectiveInterval) {
		query["NoEffectiveInterval"] = request.NoEffectiveInterval
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrometheusShrink) {
		query["Prometheus"] = request.PrometheusShrink
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	if !dara.IsNil(request.Escalations) {
		query["Escalations"] = request.Escalations
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutResourceMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple alert rules for the specified metric of a specified resource.
//
// Description:
//
// This topic provides an example on how to create a threshold-triggered alert rule for the `cpu_total` metric of an Elastic Compute Service (ECS) instance whose ID is `i-uf6j91r34rnwawoo****`. The namespace of ECS metrics is `acs_ecs_dashboard`. The alert contact group of the alert rule is `ECS_Group`. The name of the alert rule is `test123`. The ID of the alert rule is `a151cd6023eacee2f0978e03863cc1697c89508****`. The statistical method for Critical-level alerts is `Average`. The comparison operator for Critical-level alerts is `GreaterThanOrEqualToThreshold`. The threshold for Critical-level alerts is `90`. The consecutive number of times for which the metric value meets the trigger condition before a Critical-level alert is triggered is `3`.
//
// >  Statistics verification was added on August 15, 2024. Only the statistical value of the corresponding metric can be set for the Statistics parameter. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://www.alibabacloud.com/help/en/cms/support/appendix-1-metrics).
//
// @param request - PutResourceMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutResourceMetricRulesResponse
func (client *Client) PutResourceMetricRulesWithContext(ctx context.Context, request *PutResourceMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *PutResourceMetricRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutResourceMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutResourceMetricRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes tags.
//
// @param request - RemoveTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTagsResponse
func (client *Client) RemoveTagsWithContext(ctx context.Context, request *RemoveTagsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTags"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Debugs a system event of an Alibaba Cloud service.
//
// Description:
//
// This operation is used to test whether a system event can be triggered as expected. You can call this operation to simulate a system event and check whether an expected response is returned after the system event triggers an alert.
//
// @param request - SendDryRunSystemEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendDryRunSystemEventResponse
func (client *Client) SendDryRunSystemEventWithContext(ctx context.Context, request *SendDryRunSystemEventRequest, runtime *dara.RuntimeOptions) (_result *SendDryRunSystemEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventContent) {
		query["EventContent"] = request.EventContent
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendDryRunSystemEvent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendDryRunSystemEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls the CloudMonitor agent from a third-party host.
//
// Description:
//
// >  This API operation is not applicable to Elastic Compute Service (ECS) instances. To uninstall the agent from an ECS instance, see [Install and uninstall the CloudMonitor agent](https://help.aliyun.com/document_detail/183482.html).
//
// @param request - UninstallMonitoringAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallMonitoringAgentResponse
func (client *Client) UninstallMonitoringAgentWithContext(ctx context.Context, request *UninstallMonitoringAgentRequest, runtime *dara.RuntimeOptions) (_result *UninstallMonitoringAgentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallMonitoringAgent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallMonitoringAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
