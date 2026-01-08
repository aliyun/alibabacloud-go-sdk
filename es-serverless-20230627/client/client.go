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
	client.Endpoint, _err = client.GetEndpoint(dara.String("es-serverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 撤销规格审批
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSpecReviewTaskResponse
func (client *Client) CancelSpecReviewTaskWithOptions(appName *string, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelSpecReviewTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelSpecReviewTask"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/spec-review-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelSpecReviewTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销规格审批
//
// @return CancelSpecReviewTaskResponse
func (client *Client) CancelSpecReviewTask(appName *string, taskId *string) (_result *CancelSpecReviewTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelSpecReviewTaskResponse{}
	_body, _err := client.CancelSpecReviewTaskWithOptions(appName, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Serverless应用
//
// @param request - CreateAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["appName"] = request.AppName
	}

	if !dara.IsNil(request.Authentication) {
		body["authentication"] = request.Authentication
	}

	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Network) {
		body["network"] = request.Network
	}

	if !dara.IsNil(request.PrivateNetwork) {
		body["privateNetwork"] = request.PrivateNetwork
	}

	if !dara.IsNil(request.QuotaInfo) {
		body["quotaInfo"] = request.QuotaInfo
	}

	if !dara.IsNil(request.RegionId) {
		body["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Serverless应用
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建端点
//
// @param request - CreateEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpointWithOptions(request *CreateEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndpointZones) {
		body["endpointZones"] = request.EndpointZones
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.VpcId) {
		body["vpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEndpoint"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/endpoints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建端点
//
// @param request - CreateEndpointRequest
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpoint(request *CreateEndpointRequest) (_result *CreateEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEndpointResponse{}
	_body, _err := client.CreateEndpointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建快照
//
// @param request - CreateSnapshotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithOptions(appName *string, repository *string, request *CreateSnapshotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Indices) {
		body["indices"] = request.Indices
	}

	if !dara.IsNil(request.Snapshot) {
		body["snapshot"] = request.Snapshot
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnapshot"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/snapshot-repositories/" + dara.PercentEncode(dara.StringValue(repository)) + "/snapshots"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建快照
//
// @param request - CreateSnapshotRequest
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshot(appName *string, repository *string, request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(appName, repository, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Serverless应用。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(appName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApp"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Serverless应用。
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(appName *string) (_result *DeleteAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除词典
//
// @param request - DeleteDictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDictResponse
func (client *Client) DeleteDictWithOptions(appName *string, request *DeleteDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDictResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDict"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/dicts/actions/remove"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除词典
//
// @param request - DeleteDictRequest
//
// @return DeleteDictResponse
func (client *Client) DeleteDict(appName *string, request *DeleteDictRequest) (_result *DeleteDictResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDictResponse{}
	_body, _err := client.DeleteDictWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除端点
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpointWithOptions(endpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEndpointResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEndpoint"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/endpoints/" + dara.PercentEncode(dara.StringValue(endpointId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除端点
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpoint(endpointId *string) (_result *DeleteEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEndpointResponse{}
	_body, _err := client.DeleteEndpointWithOptions(endpointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除快照
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithOptions(appName *string, repository *string, snapshot *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshot"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/snapshot-repositories/" + dara.PercentEncode(dara.StringValue(repository)) + "/snapshots/" + dara.PercentEncode(dara.StringValue(snapshot))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除快照
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshot(appName *string, repository *string, snapshot *string) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(appName, repository, snapshot, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Serverless应用详情
//
// @param request - GetAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppResponse
func (client *Client) GetAppWithOptions(appName *string, request *GetAppRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Detailed) {
		query["detailed"] = request.Detailed
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApp"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Serverless应用详情
//
// @param request - GetAppRequest
//
// @return GetAppResponse
func (client *Client) GetApp(appName *string, request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Serverless应用配额详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppQuotaResponse
func (client *Client) GetAppQuotaWithOptions(appName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAppQuotaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppQuota"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/quota"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Serverless应用配额详情
//
// @return GetAppQuotaResponse
func (client *Client) GetAppQuota(appName *string) (_result *GetAppQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppQuotaResponse{}
	_body, _err := client.GetAppQuotaWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取监控数据
//
// @param request - GetMonitorDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorDataResponse
func (client *Client) GetMonitorDataWithOptions(request *GetMonitorDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMonitorDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonitorData"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/emon/metrics/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取监控数据
//
// @param request - GetMonitorDataRequest
//
// @return GetMonitorDataResponse
func (client *Client) GetMonitorData(request *GetMonitorDataRequest) (_result *GetMonitorDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMonitorDataResponse{}
	_body, _err := client.GetMonitorDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自动备份配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSnapshotSettingResponse
func (client *Client) GetSnapshotSettingWithOptions(appName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSnapshotSettingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSnapshotSetting"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/auto-snapshot-setting"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自动备份配置
//
// @return GetSnapshotSettingResponse
func (client *Client) GetSnapshotSetting(appName *string) (_result *GetSnapshotSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSnapshotSettingResponse{}
	_body, _err := client.GetSnapshotSettingWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取配额审批详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpecReviewTaskResponse
func (client *Client) GetSpecReviewTaskWithOptions(appName *string, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSpecReviewTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSpecReviewTask"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/spec-review-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSpecReviewTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取配额审批详情
//
// @return GetSpecReviewTaskResponse
func (client *Client) GetSpecReviewTask(appName *string, taskId *string) (_result *GetSpecReviewTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSpecReviewTaskResponse{}
	_body, _err := client.GetSpecReviewTaskWithOptions(appName, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看Serverless应用列表
//
// @param request - ListAppsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsResponse
func (client *Client) ListAppsWithOptions(request *ListAppsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["appName"] = request.AppName
	}

	if !dara.IsNil(request.CreateTime) {
		query["createTime"] = request.CreateTime
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.OrderType) {
		query["orderType"] = request.OrderType
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

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApps"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Serverless应用列表
//
// @param request - ListAppsRequest
//
// @return ListAppsResponse
func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取词典列表
//
// @param request - ListDictsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictsResponse
func (client *Client) ListDictsWithOptions(appName *string, request *ListDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDicts"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/dicts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取词典列表
//
// @param request - ListDictsRequest
//
// @return ListDictsResponse
func (client *Client) ListDicts(appName *string, request *ListDictsRequest) (_result *ListDictsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictsResponse{}
	_body, _err := client.ListDictsWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取端点信息列表
//
// @param request - ListEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEndpointsResponse
func (client *Client) ListEndpointsWithOptions(request *ListEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.VpcId) {
		query["vpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEndpoints"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取端点信息列表
//
// @param request - ListEndpointsRequest
//
// @return ListEndpointsResponse
func (client *Client) ListEndpoints(request *ListEndpointsRequest) (_result *ListEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEndpointsResponse{}
	_body, _err := client.ListEndpointsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看索引列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndicesResponse
func (client *Client) ListIndicesWithOptions(appName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndicesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIndices"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/indices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看索引列表
//
// @return ListIndicesResponse
func (client *Client) ListIndices(appName *string) (_result *ListIndicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndicesResponse{}
	_body, _err := client.ListIndicesWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取快照仓库列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotRepositoriesResponse
func (client *Client) ListSnapshotRepositoriesWithOptions(appName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSnapshotRepositoriesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshotRepositories"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/snapshot-repositories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取快照仓库列表
//
// @return ListSnapshotRepositoriesResponse
func (client *Client) ListSnapshotRepositories(appName *string) (_result *ListSnapshotRepositoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotRepositoriesResponse{}
	_body, _err := client.ListSnapshotRepositoriesWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取仓库的快照列表
//
// @param request - ListSnapshotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithOptions(appName *string, request *ListSnapshotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
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

	if !dara.IsNil(request.Repository) {
		query["repository"] = request.Repository
	}

	if !dara.IsNil(request.Snapshot) {
		query["snapshot"] = request.Snapshot
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshots"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/snapshots"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取仓库的快照列表
//
// @param request - ListSnapshotsRequest
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshots(appName *string, request *ListSnapshotsRequest) (_result *ListSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotsResponse{}
	_body, _err := client.ListSnapshotsWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取规格审批列表
//
// @param request - ListSpecReviewTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSpecReviewTasksResponse
func (client *Client) ListSpecReviewTasksWithOptions(appName *string, request *ListSpecReviewTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSpecReviewTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
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
		Action:      dara.String("ListSpecReviewTasks"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/spec-review-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSpecReviewTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取规格审批列表
//
// @param request - ListSpecReviewTasksRequest
//
// @return ListSpecReviewTasksResponse
func (client *Client) ListSpecReviewTasks(appName *string, request *ListSpecReviewTasksRequest) (_result *ListSpecReviewTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSpecReviewTasksResponse{}
	_body, _err := client.ListSpecReviewTasksWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑Serverless应用
//
// @param request - UpdateAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppResponse
func (client *Client) UpdateAppWithOptions(appName *string, request *UpdateAppRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyReason) {
		body["applyReason"] = request.ApplyReason
	}

	if !dara.IsNil(request.Authentication) {
		body["authentication"] = request.Authentication
	}

	if !dara.IsNil(request.ContactInfo) {
		body["contactInfo"] = request.ContactInfo
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.LimiterInfo) {
		body["limiterInfo"] = request.LimiterInfo
	}

	if !dara.IsNil(request.Network) {
		body["network"] = request.Network
	}

	if !dara.IsNil(request.PrivateNetwork) {
		body["privateNetwork"] = request.PrivateNetwork
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApp"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑Serverless应用
//
// @param request - UpdateAppRequest
//
// @return UpdateAppResponse
func (client *Client) UpdateApp(appName *string, request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新词典
//
// @param request - UpdateDictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDictResponse
func (client *Client) UpdateDictWithOptions(appName *string, request *UpdateDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDictResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowCover) {
		query["allowCover"] = request.AllowCover
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.SourceType) {
		body["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDict"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/dicts"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新词典
//
// @param request - UpdateDictRequest
//
// @return UpdateDictResponse
func (client *Client) UpdateDict(appName *string, request *UpdateDictRequest) (_result *UpdateDictResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDictResponse{}
	_body, _err := client.UpdateDictWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改端点信息
//
// @param request - UpdateEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointResponse
func (client *Client) UpdateEndpointWithOptions(endpointId *string, request *UpdateEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndpointZones) {
		body["endpointZones"] = request.EndpointZones
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEndpoint"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/endpoints/" + dara.PercentEncode(dara.StringValue(endpointId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改端点信息
//
// @param request - UpdateEndpointRequest
//
// @return UpdateEndpointResponse
func (client *Client) UpdateEndpoint(endpointId *string, request *UpdateEndpointRequest) (_result *UpdateEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEndpointResponse{}
	_body, _err := client.UpdateEndpointWithOptions(endpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用公网配置。
//
// @param request - UpdateNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkResponse
func (client *Client) UpdateNetworkWithOptions(appName *string, request *UpdateNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNetwork"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/networks"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用公网配置。
//
// @param request - UpdateNetworkRequest
//
// @return UpdateNetworkResponse
func (client *Client) UpdateNetwork(appName *string, request *UpdateNetworkRequest) (_result *UpdateNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNetworkResponse{}
	_body, _err := client.UpdateNetworkWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用公网信息。
//
// @param request - UpdatePrivateNetwrokRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateNetwrokResponse
func (client *Client) UpdatePrivateNetwrokWithOptions(appName *string, request *UpdatePrivateNetwrokRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrivateNetwrokResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateNetwrok"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/private-networks"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateNetwrokResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用公网信息。
//
// @param request - UpdatePrivateNetwrokRequest
//
// @return UpdatePrivateNetwrokResponse
func (client *Client) UpdatePrivateNetwrok(appName *string, request *UpdatePrivateNetwrokRequest) (_result *UpdatePrivateNetwrokResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePrivateNetwrokResponse{}
	_body, _err := client.UpdatePrivateNetwrokWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自动备份配置
//
// @param request - UpdateSnapshotSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSettingWithOptions(appName *string, request *UpdateSnapshotSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSnapshotSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.QuartzRegex) {
		body["quartzRegex"] = request.QuartzRegex
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSnapshotSetting"),
		Version:     dara.String("2023-06-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/es-serverless/instances/" + dara.PercentEncode(dara.StringValue(appName)) + "/auto-snapshot-setting"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自动备份配置
//
// @param request - UpdateSnapshotSettingRequest
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSetting(appName *string, request *UpdateSnapshotSettingRequest) (_result *UpdateSnapshotSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.UpdateSnapshotSettingWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
