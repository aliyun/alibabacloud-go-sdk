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
// 安装接入组件，代表进行一次接入
//
// @param request - CreateAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAddonReleaseResponse
func (client *Client) CreateAddonReleaseWithOptions(policyId *string, request *CreateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装接入组件，代表进行一次接入
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
// 创建聚合任务组
//
// @param request - CreateAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAggTaskGroupResponse
func (client *Client) CreateAggTaskGroupWithOptions(instanceId *string, request *CreateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAggTaskGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建EntityStore相关存储
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStoreWithOptions(workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEntityStoreResponse, _err error) {
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
// 创建EntityStore相关存储
//
// @return CreateEntityStoreResponse
func (client *Client) CreateEntityStore(workspaceName *string) (_result *CreateEntityStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEntityStoreResponse{}
	_body, _err := client.CreateEntityStoreWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateIntegrationPolicyWithOptions(request *CreateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建Prometheus监控实例
//
// @param request - CreatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusInstanceResponse
func (client *Client) CreatePrometheusInstanceWithOptions(request *CreatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建prometheus视图
//
// @param request - CreatePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusViewResponse
func (client *Client) CreatePrometheusViewWithOptions(request *CreatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建Prometheus监控实例
//
// @param request - CreatePrometheusVirtualInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrometheusVirtualInstanceResponse
func (client *Client) CreatePrometheusVirtualInstanceWithOptions(request *CreatePrometheusVirtualInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrometheusVirtualInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建Service
//
// @param request - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(workspace *string, request *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建票据
//
// @param request - CreateTicketRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 创建Umodel配置
//
// @param request - CreateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUmodelResponse
func (client *Client) CreateUmodelWithOptions(workspace *string, request *CreateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUmodelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 删除addon release信息
//
// @param request - DeleteAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddonReleaseResponse
func (client *Client) DeleteAddonReleaseWithOptions(policyId *string, request *DeleteAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 删除聚合任务组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAggTaskGroupResponse
func (client *Client) DeleteAggTaskGroupWithOptions(instanceId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAggTaskGroupResponse, _err error) {
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
// 删除聚合任务组
//
// @return DeleteAggTaskGroupResponse
func (client *Client) DeleteAggTaskGroup(instanceId *string, groupId *string) (_result *DeleteAggTaskGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAggTaskGroupResponse{}
	_body, _err := client.DeleteAggTaskGroupWithOptions(instanceId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEntityStoreWithOptions(workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEntityStoreResponse, _err error) {
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
// 删除EntityStore相关存储
//
// @return DeleteEntityStoreResponse
func (client *Client) DeleteEntityStore(workspaceName *string) (_result *DeleteEntityStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEntityStoreResponse{}
	_body, _err := client.DeleteEntityStoreWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteIntegrationPolicyWithOptions(policyId *string, request *DeleteIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 删除prom实例
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrometheusInstanceResponse
func (client *Client) DeletePrometheusInstanceWithOptions(prometheusInstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusInstanceResponse, _err error) {
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
// 删除prom实例
//
// @return DeletePrometheusInstanceResponse
func (client *Client) DeletePrometheusInstance(prometheusInstanceId *string) (_result *DeletePrometheusInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePrometheusInstanceResponse{}
	_body, _err := client.DeletePrometheusInstanceWithOptions(prometheusInstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePrometheusViewWithOptions(prometheusViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrometheusViewResponse, _err error) {
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
// 删除prometheus视图实例
//
// @return DeletePrometheusViewResponse
func (client *Client) DeletePrometheusView(prometheusViewId *string) (_result *DeletePrometheusViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePrometheusViewResponse{}
	_body, _err := client.DeletePrometheusViewWithOptions(prometheusViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteServiceWithOptions(workspace *string, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
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
// 删除Service
//
// @return DeleteServiceResponse
func (client *Client) DeleteService(workspace *string, serviceId *string) (_result *DeleteServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(workspace, serviceId, headers, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodelWithOptions(workspace *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelResponse, _err error) {
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
// 删除Umodel配置信息
//
// @return DeleteUmodelResponse
func (client *Client) DeleteUmodel(workspace *string) (_result *DeleteUmodelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUmodelResponse{}
	_body, _err := client.DeleteUmodelWithOptions(workspace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteUmodelDataWithOptions(workspace *string, request *DeleteUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUmodelDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 删除工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
// 删除工作空间
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceName *string) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAddonReleaseWithOptions(releaseName *string, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddonReleaseResponse, _err error) {
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
// 查看addon release(查看接入状态)
//
// @return GetAddonReleaseResponse
func (client *Client) GetAddonRelease(releaseName *string, policyId *string) (_result *GetAddonReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddonReleaseResponse{}
	_body, _err := client.GetAddonReleaseWithOptions(releaseName, policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAggTaskGroupWithOptions(instanceId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAggTaskGroupResponse, _err error) {
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
// 描述聚合任务组
//
// @return GetAggTaskGroupResponse
func (client *Client) GetAggTaskGroup(instanceId *string, groupId *string) (_result *GetAggTaskGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAggTaskGroupResponse{}
	_body, _err := client.GetAggTaskGroupWithOptions(instanceId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEntityStoreWithOptions(workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEntityStoreResponse, _err error) {
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
// 获取EntityStore相关存储信息
//
// @return GetEntityStoreResponse
func (client *Client) GetEntityStore(workspaceName *string) (_result *GetEntityStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEntityStoreResponse{}
	_body, _err := client.GetEntityStoreWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEntityStoreDataWithOptions(workspace *string, request *GetEntityStoreDataRequest, headers *GetEntityStoreDataHeaders, runtime *dara.RuntimeOptions) (_result *GetEntityStoreDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 查询接入中心策略列表信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntegrationPolicyResponse
func (client *Client) GetIntegrationPolicyWithOptions(policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIntegrationPolicyResponse, _err error) {
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
// 查询接入中心策略列表信息
//
// @return GetIntegrationPolicyResponse
func (client *Client) GetIntegrationPolicy(policyId *string) (_result *GetIntegrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIntegrationPolicyResponse{}
	_body, _err := client.GetIntegrationPolicyWithOptions(policyId, headers, runtime)
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
// @param request - GetPrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusInstanceResponse
func (client *Client) GetPrometheusInstanceWithOptions(prometheusInstanceId *string, request *GetPrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusInstanceResponse, _err error) {
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
// 查询指定Prometheus视图实例
//
// @param request - GetPrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrometheusViewResponse
func (client *Client) GetPrometheusViewWithOptions(prometheusViewId *string, request *GetPrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 查询 Service
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(workspace *string, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
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
// 查询 Service
//
// @return GetServiceResponse
func (client *Client) GetService(workspace *string, serviceId *string) (_result *GetServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(workspace, serviceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceObservabilityWithOptions(workspace *string, _type *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceObservabilityResponse, _err error) {
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
// 获取应用可观测实例
//
// @return GetServiceObservabilityResponse
func (client *Client) GetServiceObservability(workspace *string, _type *string) (_result *GetServiceObservabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceObservabilityResponse{}
	_body, _err := client.GetServiceObservabilityWithOptions(workspace, _type, headers, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUmodelResponse
func (client *Client) GetUmodelWithOptions(workspace *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelResponse, _err error) {
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
// 获取Umodel配置信息
//
// @return GetUmodelResponse
func (client *Client) GetUmodel(workspace *string) (_result *GetUmodelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUmodelResponse{}
	_body, _err := client.GetUmodelWithOptions(workspace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUmodelDataWithOptions(workspace *string, request *GetUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUmodelDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
// 获取工作空间
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceName *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAddonReleasesWithOptions(policyId *string, request *ListAddonReleasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddonReleasesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 列举聚合任务组
//
// @param tmpReq - ListAggTaskGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAggTaskGroupsResponse
func (client *Client) ListAggTaskGroupsWithOptions(instanceId *string, tmpReq *ListAggTaskGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAggTaskGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 查询告警动作
//
// @param tmpReq - ListAlertActionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertActionsResponse
func (client *Client) ListAlertActionsWithOptions(tmpReq *ListAlertActionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertActionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 查询接入中心策略列表信息
//
// @param tmpReq - ListIntegrationPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPoliciesResponse
func (client *Client) ListIntegrationPoliciesWithOptions(tmpReq *ListIntegrationPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPoliciesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询接入中心策略列表信息
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
// 获取接入中心策略的存储要求信息
//
// @param request - ListIntegrationPolicyCustomScrapeJobRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyCustomScrapeJobRulesResponse
func (client *Client) ListIntegrationPolicyCustomScrapeJobRulesWithOptions(policyId *string, request *ListIntegrationPolicyCustomScrapeJobRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyCustomScrapeJobRulesResponse, _err error) {
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
// 策略大盘列表
//
// @param request - ListIntegrationPolicyDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyDashboardsResponse
func (client *Client) ListIntegrationPolicyDashboardsWithOptions(policyId *string, request *ListIntegrationPolicyDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyDashboardsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddonName) {
		query["addonName"] = request.AddonName
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
// 策略大盘列表
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
// 获取接入中心策略的PodMonitor资源
//
// @param request - ListIntegrationPolicyPodMonitorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyPodMonitorsResponse
func (client *Client) ListIntegrationPolicyPodMonitorsWithOptions(policyId *string, request *ListIntegrationPolicyPodMonitorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyPodMonitorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListIntegrationPolicyStorageRequirementsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntegrationPolicyStorageRequirementsResponse
func (client *Client) ListIntegrationPolicyStorageRequirementsWithOptions(policyId *string, request *ListIntegrationPolicyStorageRequirementsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIntegrationPolicyStorageRequirementsResponse, _err error) {
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
// 获取Prometheus实例大盘列表
//
// @param request - ListPrometheusDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusDashboardsResponse
func (client *Client) ListPrometheusDashboardsWithOptions(prometheusInstanceId *string, request *ListPrometheusDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusDashboardsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取Prometheus实例信息列表
//
// @param tmpReq - ListPrometheusInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusInstancesResponse
func (client *Client) ListPrometheusInstancesWithOptions(tmpReq *ListPrometheusInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取Prometheus视图实例信息列表
//
// @param tmpReq - ListPrometheusViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusViewsResponse
func (client *Client) ListPrometheusViewsWithOptions(tmpReq *ListPrometheusViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusViewsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取Prometheus虚拟实例
//
// @param request - ListPrometheusVirtualInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrometheusVirtualInstancesResponse
func (client *Client) ListPrometheusVirtualInstancesWithOptions(request *ListPrometheusVirtualInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrometheusVirtualInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 列出资源Service
//
// @param request - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(workspace *string, request *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取工作空间列表
//
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
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
// 创建工作空间
//
// @param request - PutWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutWorkspaceResponse
func (client *Client) PutWorkspaceWithOptions(workspaceName *string, request *PutWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutWorkspaceResponse, _err error) {
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
// 升级接入组件
//
// @param request - UpdateAddonReleaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAddonReleaseResponse
func (client *Client) UpdateAddonReleaseWithOptions(releaseName *string, policyId *string, request *UpdateAddonReleaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAddonReleaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 应用聚合任务组
//
// @param request - UpdateAggTaskGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupResponse
func (client *Client) UpdateAggTaskGroupWithOptions(instanceId *string, groupId *string, request *UpdateAggTaskGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 更新聚合任务组状态
//
// @param request - UpdateAggTaskGroupStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAggTaskGroupStatusResponse
func (client *Client) UpdateAggTaskGroupStatusWithOptions(instanceId *string, groupId *string, request *UpdateAggTaskGroupStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAggTaskGroupStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 更新指定策略
//
// @param request - UpdateIntegrationPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegrationPolicyResponse
func (client *Client) UpdateIntegrationPolicyWithOptions(integrationPolicyId *string, request *UpdateIntegrationPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIntegrationPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 更新Prom实例信息
//
// @param request - UpdatePrometheusInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusInstanceResponse
func (client *Client) UpdatePrometheusInstanceWithOptions(prometheusInstanceId *string, request *UpdatePrometheusInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusInstanceResponse, _err error) {
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
// 更新Prom视图实例信息
//
// @param request - UpdatePrometheusViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrometheusViewResponse
func (client *Client) UpdatePrometheusViewWithOptions(prometheusViewId *string, request *UpdatePrometheusViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrometheusViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 更新UpdateService
//
// @param request - UpdateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithOptions(workspace *string, serviceId *string, request *UpdateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 更新Umodel配置信息
//
// @param request - UpdateUmodelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUmodelResponse
func (client *Client) UpdateUmodelWithOptions(workspace *string, request *UpdateUmodelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUmodelResponse, _err error) {
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
// 写入 Umodel Elements
//
// @param request - UpsertUmodelDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertUmodelDataResponse
func (client *Client) UpsertUmodelDataWithOptions(workspace *string, request *UpsertUmodelDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpsertUmodelDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
