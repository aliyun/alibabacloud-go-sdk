// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 安装接入组件，代表进行一次接入
//
// @param request - CreateAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAddonReleaseResponse
func (client *Client) CreateAddonReleaseWithContext(ctx context.Context, policyId *string, request *CreateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAddonReleaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建聚合任务组
//
// @param request - CreateAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggTaskGroupResponse
func (client *Client) CreateAggTaskGroupWithContext(ctx context.Context, instanceId *string, request *CreateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAggTaskGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建EntityStore相关存储
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStoreWithContext(ctx context.Context, workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEntityStoreResponse, _err error) {
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
// 创建接入中心策略
//
// @param request - CreateIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntegrationPolicyResponse
func (client *Client) CreateIntegrationPolicyWithContext(ctx context.Context, request *CreateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIntegrationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建Prometheus监控实例
//
// @param request - CreatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstanceWithContext(ctx context.Context, request *CreatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建prometheus视图
//
// @param request - CreatePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusViewResponse
func (client *Client) CreatePrometheusViewWithContext(ctx context.Context, request *CreatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建Prometheus监控实例
//
// @param request - CreatePrometheusVirtualInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusVirtualInstanceResponse
func (client *Client) CreatePrometheusVirtualInstanceWithContext(ctx context.Context, request *CreatePrometheusVirtualInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusVirtualInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		body["namespace"] = request.Namespace
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
// 创建Service
//
// @param request - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithContext(ctx context.Context, workspace *string, request *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.ServiceName) {
		body["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceStatus) {
		body["serviceStatus"] = request.ServiceStatus
	}

	if !dara.IsNil(request.ServiceType) {
		body["serviceType"] = request.ServiceType
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
// 创建票据
//
// @param request - CreateTicketRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithContext(ctx context.Context, request *CreateTicketRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建Umodel配置
//
// @param request - CreateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodelWithContext(ctx context.Context, workspace *string, request *CreateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUmodelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除addon release信息
//
// @param request - DeleteAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddonReleaseResponse
func (client *Client) DeleteAddonReleaseWithContext(ctx context.Context, policyId *string, request *DeleteAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAddonReleaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除聚合任务组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggTaskGroupResponse
func (client *Client) DeleteAggTaskGroupWithContext(ctx context.Context, instanceId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAggTaskGroupResponse, _err error) {
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
// 删除EntityStore相关存储
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStoreWithContext(ctx context.Context, workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEntityStoreResponse, _err error) {
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
// 删除接入中心策略
//
// @param request - DeleteIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntegrationPolicyResponse
func (client *Client) DeleteIntegrationPolicyWithContext(ctx context.Context, policyId *string, request *DeleteIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIntegrationPolicyResponse, _err error) {
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
// 删除prom实例
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusInstanceResponse
func (client *Client) DeletePrometheusInstanceWithContext(ctx context.Context, prometheusInstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusInstanceResponse, _err error) {
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
// 删除prometheus视图实例
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusViewResponse
func (client *Client) DeletePrometheusViewWithContext(ctx context.Context, prometheusViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusViewResponse, _err error) {
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
// 删除Service
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithContext(ctx context.Context, workspace *string, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
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
// 删除Umodel配置信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodelWithContext(ctx context.Context, workspace *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelResponse, _err error) {
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
// 删除 Umodel Elements
//
// @param request - DeleteUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelDataResponse
func (client *Client) DeleteUmodelDataWithContext(ctx context.Context, workspace *string, request *DeleteUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithContext(ctx context.Context, workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
// 查看addon release(查看接入状态)
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonReleaseResponse
func (client *Client) GetAddonReleaseWithContext(ctx context.Context, releaseName *string, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonReleaseResponse, _err error) {
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
// 描述聚合任务组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAggTaskGroupResponse
func (client *Client) GetAggTaskGroupWithContext(ctx context.Context, instanceId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAggTaskGroupResponse, _err error) {
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
// 获取EntityStore相关存储信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStoreWithContext(ctx context.Context, workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEntityStoreResponse, _err error) {
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
// 查询指定Workspace下的实体和关系数据，返回结果显示某时间区间中的实体数据（返回结果压缩后传输）。
//
// @param request - GetEntityStoreDataRequest
//
// @param headers - GetEntityStoreDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityStoreDataResponse
func (client *Client) GetEntityStoreDataWithContext(ctx context.Context, workspace *string, request *GetEntityStoreDataRequest, headers *GetEntityStoreDataHeaders, runtime *dara.RuntimeOptions) (_result *GetEntityStoreDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询接入中心策略列表信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationPolicyResponse
func (client *Client) GetIntegrationPolicyWithContext(ctx context.Context, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIntegrationPolicyResponse, _err error) {
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
// 查询指定环境实例
//
// @param request - GetPrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusInstanceResponse
func (client *Client) GetPrometheusInstanceWithContext(ctx context.Context, prometheusInstanceId *string, request *GetPrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询指定Prometheus视图实例
//
// @param request - GetPrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusViewResponse
func (client *Client) GetPrometheusViewWithContext(ctx context.Context, prometheusViewId *string, request *GetPrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询 Service
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithContext(ctx context.Context, workspace *string, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
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
// 获取应用可观测实例
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceObservabilityResponse
func (client *Client) GetServiceObservabilityWithContext(ctx context.Context, workspace *string, _type *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceObservabilityResponse, _err error) {
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
// 获取Umodel配置信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelResponse
func (client *Client) GetUmodelWithContext(ctx context.Context, workspace *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelResponse, _err error) {
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
// 获取相关联的 Umodel 图数据
//
// @param request - GetUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelDataResponse
func (client *Client) GetUmodelDataWithContext(ctx context.Context, workspace *string, request *GetUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithContext(ctx context.Context, workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
// addon的release列表
//
// @param request - ListAddonReleasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonReleasesResponse
func (client *Client) ListAddonReleasesWithContext(ctx context.Context, policyId *string, request *ListAddonReleasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonReleasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 列举聚合任务组
//
// @param tmpReq - ListAggTaskGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggTaskGroupsResponse
func (client *Client) ListAggTaskGroupsWithContext(ctx context.Context, instanceId *string, tmpReq *ListAggTaskGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAggTaskGroupsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询告警动作
//
// @param tmpReq - ListAlertActionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActionsWithContext(ctx context.Context, tmpReq *ListAlertActionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertActionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取接入中心策略的存储要求信息
//
// @param request - ListIntegrationPolicyCustomScrapeJobRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyCustomScrapeJobRulesResponse
func (client *Client) ListIntegrationPolicyCustomScrapeJobRulesWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyCustomScrapeJobRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyCustomScrapeJobRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取接入中心策略的PodMonitor资源
//
// @param request - ListIntegrationPolicyPodMonitorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyPodMonitorsResponse
func (client *Client) ListIntegrationPolicyPodMonitorsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyPodMonitorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyPodMonitorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// @param request - ListIntegrationPolicyStorageRequirementsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyStorageRequirementsResponse
func (client *Client) ListIntegrationPolicyStorageRequirementsWithContext(ctx context.Context, policyId *string, request *ListIntegrationPolicyStorageRequirementsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyStorageRequirementsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取Prometheus实例大盘列表
//
// @param request - ListPrometheusDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusDashboardsResponse
func (client *Client) ListPrometheusDashboardsWithContext(ctx context.Context, prometheusInstanceId *string, request *ListPrometheusDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusDashboardsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取Prometheus实例信息列表
//
// @param tmpReq - ListPrometheusInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusInstancesResponse
func (client *Client) ListPrometheusInstancesWithContext(ctx context.Context, tmpReq *ListPrometheusInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取Prometheus视图实例信息列表
//
// @param tmpReq - ListPrometheusViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusViewsResponse
func (client *Client) ListPrometheusViewsWithContext(ctx context.Context, tmpReq *ListPrometheusViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusViewsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取Prometheus虚拟实例
//
// @param request - ListPrometheusVirtualInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusVirtualInstancesResponse
func (client *Client) ListPrometheusVirtualInstancesWithContext(ctx context.Context, request *ListPrometheusVirtualInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusVirtualInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
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
// 列出资源Service
//
// @param request - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithContext(ctx context.Context, workspace *string, request *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
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
// 获取工作空间列表
//
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithContext(ctx context.Context, tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建工作空间
//
// @param request - PutWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspaceWithContext(ctx context.Context, workspaceName *string, request *PutWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 升级接入组件
//
// @param request - UpdateAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAddonReleaseResponse
func (client *Client) UpdateAddonReleaseWithContext(ctx context.Context, releaseName *string, policyId *string, request *UpdateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAddonReleaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 应用聚合任务组
//
// @param request - UpdateAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupResponse
func (client *Client) UpdateAggTaskGroupWithContext(ctx context.Context, instanceId *string, groupId *string, request *UpdateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新聚合任务组状态
//
// @param request - UpdateAggTaskGroupStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupStatusResponse
func (client *Client) UpdateAggTaskGroupStatusWithContext(ctx context.Context, instanceId *string, groupId *string, request *UpdateAggTaskGroupStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新指定策略
//
// @param request - UpdateIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegrationPolicyResponse
func (client *Client) UpdateIntegrationPolicyWithContext(ctx context.Context, integrationPolicyId *string, request *UpdateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIntegrationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新Prom实例信息
//
// @param request - UpdatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusInstanceResponse
func (client *Client) UpdatePrometheusInstanceWithContext(ctx context.Context, prometheusInstanceId *string, request *UpdatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新Prom视图实例信息
//
// @param request - UpdatePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusViewResponse
func (client *Client) UpdatePrometheusViewWithContext(ctx context.Context, prometheusViewId *string, request *UpdatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新UpdateService
//
// @param request - UpdateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithContext(ctx context.Context, workspace *string, serviceId *string, request *UpdateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新Umodel配置信息
//
// @param request - UpdateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodelWithContext(ctx context.Context, workspace *string, request *UpdateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUmodelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 写入 Umodel Elements
//
// @param request - UpsertUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelDataWithContext(ctx context.Context, workspace *string, request *UpsertUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpsertUmodelDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
