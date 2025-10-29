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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("elasticsearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Restores nodes in disabled zones. This operation is available only for multi-zone Elasticsearch clusters.
//
// @param request - ActivateZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateZonesResponse
func (client *Client) ActivateZonesWithOptions(InstanceId *string, request *ActivateZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateZonesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateZones"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/recover-zones"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores nodes in disabled zones. This operation is available only for multi-zone Elasticsearch clusters.
//
// @param request - ActivateZonesRequest
//
// @return ActivateZonesResponse
func (client *Client) ActivateZones(InstanceId *string, request *ActivateZonesRequest) (_result *ActivateZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActivateZonesResponse{}
	_body, _err := client.ActivateZonesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Connects Elasticsearch clusters.
//
// @param request - AddConnectableClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddConnectableClusterResponse
func (client *Client) AddConnectableClusterWithOptions(InstanceId *string, request *AddConnectableClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddConnectableClusterResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddConnectableCluster"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/connected-clusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddConnectableClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Connects Elasticsearch clusters.
//
// @param request - AddConnectableClusterRequest
//
// @return AddConnectableClusterResponse
func (client *Client) AddConnectableCluster(InstanceId *string, request *AddConnectableClusterRequest) (_result *AddConnectableClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddConnectableClusterResponse{}
	_body, _err := client.AddConnectableClusterWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the AddSnapshotRepo to create a reference repository when configuring a cross-cluster OSS repository.
//
// @param request - AddSnapshotRepoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSnapshotRepoResponse
func (client *Client) AddSnapshotRepoWithOptions(InstanceId *string, request *AddSnapshotRepoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddSnapshotRepoResponse, _err error) {
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
		Action:      dara.String("AddSnapshotRepo"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshot-repos"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSnapshotRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the AddSnapshotRepo to create a reference repository when configuring a cross-cluster OSS repository.
//
// @param request - AddSnapshotRepoRequest
//
// @return AddSnapshotRepoResponse
func (client *Client) AddSnapshotRepo(InstanceId *string, request *AddSnapshotRepoRequest) (_result *AddSnapshotRepoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddSnapshotRepoResponse{}
	_body, _err := client.AddSnapshotRepoWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores an Elasticsearch cluster that is frozen after it is released.
//
// @param request - CancelDeletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDeletionResponse
func (client *Client) CancelDeletionWithOptions(InstanceId *string, request *CancelDeletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelDeletionResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDeletion"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/cancel-deletion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores an Elasticsearch cluster that is frozen after it is released.
//
// @param request - CancelDeletionRequest
//
// @return CancelDeletionResponse
func (client *Client) CancelDeletion(InstanceId *string, request *CancelDeletionRequest) (_result *CancelDeletionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelDeletionResponse{}
	_body, _err := client.CancelDeletionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a Logstash cluster that is frozen after it is released.
//
// @param request - CancelLogstashDeletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLogstashDeletionResponse
func (client *Client) CancelLogstashDeletionWithOptions(InstanceId *string, request *CancelLogstashDeletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelLogstashDeletionResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelLogstashDeletion"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/cancel-deletion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelLogstashDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a Logstash cluster that is frozen after it is released.
//
// @param request - CancelLogstashDeletionRequest
//
// @return CancelLogstashDeletionResponse
func (client *Client) CancelLogstashDeletion(InstanceId *string, request *CancelLogstashDeletionRequest) (_result *CancelLogstashDeletionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelLogstashDeletionResponse{}
	_body, _err := client.CancelLogstashDeletionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CancelTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithOptions(InstanceId *string, request *CancelTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
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

	if !dara.IsNil(request.TaskType) {
		query["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/cancel-task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelTaskRequest
//
// @return CancelTaskResponse
func (client *Client) CancelTask(InstanceId *string, request *CancelTaskRequest) (_result *CancelTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Capacity Planning
//
// @param request - CapacityPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CapacityPlanResponse
func (client *Client) CapacityPlanWithOptions(request *CapacityPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CapacityPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComplexQueryAvailable) {
		body["complexQueryAvailable"] = request.ComplexQueryAvailable
	}

	if !dara.IsNil(request.DataInfo) {
		body["dataInfo"] = request.DataInfo
	}

	if !dara.IsNil(request.Metric) {
		body["metric"] = request.Metric
	}

	if !dara.IsNil(request.UsageScenario) {
		body["usageScenario"] = request.UsageScenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CapacityPlan"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/assist/actions/capacity-plan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CapacityPlanResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Capacity Planning
//
// @param request - CapacityPlanRequest
//
// @return CapacityPlanResponse
func (client *Client) CapacityPlan(request *CapacityPlanRequest) (_result *CapacityPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CapacityPlanResponse{}
	_body, _err := client.CapacityPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭实例的智能运维功能
//
// @param request - CloseDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDiagnosisResponse
func (client *Client) CloseDiagnosisWithOptions(InstanceId *string, request *CloseDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseDiagnosis"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/close-diagnosis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭实例的智能运维功能
//
// @param request - CloseDiagnosisRequest
//
// @return CloseDiagnosisResponse
func (client *Client) CloseDiagnosis(InstanceId *string, request *CloseDiagnosisRequest) (_result *CloseDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseDiagnosisResponse{}
	_body, _err := client.CloseDiagnosisWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CloseHttpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseHttpsResponse
func (client *Client) CloseHttpsWithOptions(InstanceId *string, request *CloseHttpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseHttpsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseHttps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/close-https"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseHttpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloseHttpsRequest
//
// @return CloseHttpsResponse
func (client *Client) CloseHttps(InstanceId *string, request *CloseHttpsRequest) (_result *CloseHttpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseHttpsResponse{}
	_body, _err := client.CloseHttpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Disable Managed Index
//
// @param request - CloseManagedIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseManagedIndexResponse
func (client *Client) CloseManagedIndexWithOptions(InstanceId *string, Index *string, request *CloseManagedIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseManagedIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseManagedIndex"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/indices/" + dara.PercentEncode(dara.StringValue(Index)) + "/close-managed"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseManagedIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Disable Managed Index
//
// @param request - CloseManagedIndexRequest
//
// @return CloseManagedIndexResponse
func (client *Client) CloseManagedIndex(InstanceId *string, Index *string, request *CloseManagedIndexRequest) (_result *CloseManagedIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseManagedIndexResponse{}
	_body, _err := client.CloseManagedIndexWithOptions(InstanceId, Index, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建收集器
//
// @param request - CreateCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCollectorResponse
func (client *Client) CreateCollectorWithOptions(request *CreateCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCollectorResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.CollectorPaths) {
		body["collectorPaths"] = request.CollectorPaths
	}

	if !dara.IsNil(request.Configs) {
		body["configs"] = request.Configs
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ExtendConfigs) {
		body["extendConfigs"] = request.ExtendConfigs
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResType) {
		body["resType"] = request.ResType
	}

	if !dara.IsNil(request.ResVersion) {
		body["resVersion"] = request.ResVersion
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
		Action:      dara.String("CreateCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建收集器
//
// @param request - CreateCollectorRequest
//
// @return CreateCollectorResponse
func (client *Client) CreateCollector(request *CreateCollectorRequest) (_result *CreateCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCollectorResponse{}
	_body, _err := client.CreateCollectorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Elasticsearch组合模板
//
// @param request - CreateComponentIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentIndexResponse
func (client *Client) CreateComponentIndexWithOptions(InstanceId *string, name *string, request *CreateComponentIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComponentIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Meta) {
		body["_meta"] = request.Meta
	}

	if !dara.IsNil(request.Template) {
		body["template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComponentIndex"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/component-index/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Elasticsearch组合模板
//
// @param request - CreateComponentIndexRequest
//
// @return CreateComponentIndexResponse
func (client *Client) CreateComponentIndex(InstanceId *string, name *string, request *CreateComponentIndexRequest) (_result *CreateComponentIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComponentIndexResponse{}
	_body, _err := client.CreateComponentIndexWithOptions(InstanceId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据流
//
// @param request - CreateDataStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataStreamResponse
func (client *Client) CreateDataStreamWithOptions(InstanceId *string, request *CreateDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataStream"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/data-streams"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据流
//
// @param request - CreateDataStreamRequest
//
// @return CreateDataStreamResponse
func (client *Client) CreateDataStream(InstanceId *string, request *CreateDataStreamRequest) (_result *CreateDataStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataStreamResponse{}
	_body, _err := client.CreateDataStreamWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建索引生命周期策略
//
// @param request - CreateILMPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateILMPolicyResponse
func (client *Client) CreateILMPolicyWithOptions(InstanceId *string, request *CreateILMPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateILMPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateILMPolicy"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/ilm-policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建索引生命周期策略
//
// @param request - CreateILMPolicyRequest
//
// @return CreateILMPolicyResponse
func (client *Client) CreateILMPolicy(InstanceId *string, request *CreateILMPolicyRequest) (_result *CreateILMPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateILMPolicyResponse{}
	_body, _err := client.CreateILMPolicyWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建索引模版
//
// @param request - CreateIndexTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexTemplateResponse
func (client *Client) CreateIndexTemplateWithOptions(InstanceId *string, request *CreateIndexTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndexTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataStream) {
		body["dataStream"] = request.DataStream
	}

	if !dara.IsNil(request.IlmPolicy) {
		body["ilmPolicy"] = request.IlmPolicy
	}

	if !dara.IsNil(request.IndexPatterns) {
		body["indexPatterns"] = request.IndexPatterns
	}

	if !dara.IsNil(request.IndexTemplate) {
		body["indexTemplate"] = request.IndexTemplate
	}

	if !dara.IsNil(request.Priority) {
		body["priority"] = request.Priority
	}

	if !dara.IsNil(request.Template) {
		body["template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIndexTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/index-templates"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建索引模版
//
// @param request - CreateIndexTemplateRequest
//
// @return CreateIndexTemplateResponse
func (client *Client) CreateIndexTemplate(InstanceId *string, request *CreateIndexTemplateRequest) (_result *CreateIndexTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexTemplateResponse{}
	_body, _err := client.CreateIndexTemplateWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建logstash实例
//
// @param request - CreateLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogstashResponse
func (client *Client) CreateLogstashWithOptions(request *CreateLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLogstashResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.NetworkConfig) {
		body["networkConfig"] = request.NetworkConfig
	}

	if !dara.IsNil(request.NodeAmount) {
		body["nodeAmount"] = request.NodeAmount
	}

	if !dara.IsNil(request.NodeSpec) {
		body["nodeSpec"] = request.NodeSpec
	}

	if !dara.IsNil(request.PaymentInfo) {
		body["paymentInfo"] = request.PaymentInfo
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("CreateLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建logstash实例
//
// @param request - CreateLogstashRequest
//
// @return CreateLogstashResponse
func (client *Client) CreateLogstash(request *CreateLogstashRequest) (_result *CreateLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLogstashResponse{}
	_body, _err := client.CreateLogstashWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Logstash管道任务
//
// @param request - CreatePipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelinesResponse
func (client *Client) CreatePipelinesWithOptions(InstanceId *string, request *CreatePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePipelinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Trigger) {
		query["trigger"] = request.Trigger
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelines"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Logstash管道任务
//
// @param request - CreatePipelinesRequest
//
// @return CreatePipelinesResponse
func (client *Client) CreatePipelines(InstanceId *string, request *CreatePipelinesRequest) (_result *CreatePipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelinesResponse{}
	_body, _err := client.CreatePipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateSnapshotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithOptions(InstanceId *string, request *CreateSnapshotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnapshot"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshots"),
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

// @param request - CreateSnapshotRequest
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshot(InstanceId *string, request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建私网链接VPC终端节点
//
// Description:
//
// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D\\*\\*\\*
//
// @param request - CreateVpcEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpointWithOptions(InstanceId *string, request *CreateVpcEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceId) {
		body["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		body["zoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcEndpoint"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/vpc-endpoints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建私网链接VPC终端节点
//
// Description:
//
// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D\\*\\*\\*
//
// @param request - CreateVpcEndpointRequest
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpoint(InstanceId *string, request *CreateVpcEndpointRequest) (_result *CreateVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CreateVpcEndpointWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invoke DeactivateZones to offline certain zones when there are multiple availability zones, and migrate nodes in the offline zones to other availability zones.
//
// @param request - DeactivateZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateZonesResponse
func (client *Client) DeactivateZonesWithOptions(InstanceId *string, request *DeactivateZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeactivateZonesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactivateZones"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/down-zones"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactivateZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke DeactivateZones to offline certain zones when there are multiple availability zones, and migrate nodes in the offline zones to other availability zones.
//
// @param request - DeactivateZonesRequest
//
// @return DeactivateZonesResponse
func (client *Client) DeactivateZones(InstanceId *string, request *DeactivateZonesRequest) (_result *DeactivateZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeactivateZonesResponse{}
	_body, _err := client.DeactivateZonesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a shipper.
//
// @param request - DeleteCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCollectorResponse
func (client *Client) DeleteCollectorWithOptions(ResId *string, request *DeleteCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a shipper.
//
// @param request - DeleteCollectorRequest
//
// @return DeleteCollectorResponse
func (client *Client) DeleteCollector(ResId *string, request *DeleteCollectorRequest) (_result *DeleteCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCollectorResponse{}
	_body, _err := client.DeleteCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除组合索引模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentIndexResponse
func (client *Client) DeleteComponentIndexWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteComponentIndexResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponentIndex"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/component-index/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组合索引模板
//
// @return DeleteComponentIndexResponse
func (client *Client) DeleteComponentIndex(InstanceId *string, name *string) (_result *DeleteComponentIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteComponentIndexResponse{}
	_body, _err := client.DeleteComponentIndexWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteConnectedClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectedClusterResponse
func (client *Client) DeleteConnectedClusterWithOptions(InstanceId *string, request *DeleteConnectedClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConnectedClusterResponse, _err error) {
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

	if !dara.IsNil(request.ConnectedInstanceId) {
		query["connectedInstanceId"] = request.ConnectedInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnectedCluster"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/connected-clusters"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteConnectedClusterRequest
//
// @return DeleteConnectedClusterResponse
func (client *Client) DeleteConnectedCluster(InstanceId *string, request *DeleteConnectedClusterRequest) (_result *DeleteConnectedClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConnectedClusterResponse{}
	_body, _err := client.DeleteConnectedClusterWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据流
//
// @param request - DeleteDataStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataStreamResponse
func (client *Client) DeleteDataStreamWithOptions(InstanceId *string, DataStream *string, request *DeleteDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataStream"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/data-streams/" + dara.PercentEncode(dara.StringValue(DataStream))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据流
//
// @param request - DeleteDataStreamRequest
//
// @return DeleteDataStreamResponse
func (client *Client) DeleteDataStream(InstanceId *string, DataStream *string, request *DeleteDataStreamRequest) (_result *DeleteDataStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataStreamResponse{}
	_body, _err := client.DeleteDataStreamWithOptions(InstanceId, DataStream, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteDataTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataTaskResponse
func (client *Client) DeleteDataTaskWithOptions(InstanceId *string, request *DeleteDataTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataTask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/data-task"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteDataTaskRequest
//
// @return DeleteDataTaskResponse
func (client *Client) DeleteDataTask(InstanceId *string, request *DeleteDataTaskRequest) (_result *DeleteDataTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataTaskResponse{}
	_body, _err := client.DeleteDataTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除历史索引模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeprecatedTemplateResponse
func (client *Client) DeleteDeprecatedTemplateWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDeprecatedTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeprecatedTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/deprecated-templates/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeprecatedTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除历史索引模板
//
// @return DeleteDeprecatedTemplateResponse
func (client *Client) DeleteDeprecatedTemplate(InstanceId *string, name *string) (_result *DeleteDeprecatedTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDeprecatedTemplateResponse{}
	_body, _err := client.DeleteDeprecatedTemplateWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteILMPolicyResponse
func (client *Client) DeleteILMPolicyWithOptions(InstanceId *string, PolicyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteILMPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteILMPolicy"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/ilm-policies/" + dara.PercentEncode(dara.StringValue(PolicyName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return DeleteILMPolicyResponse
func (client *Client) DeleteILMPolicy(InstanceId *string, PolicyName *string) (_result *DeleteILMPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteILMPolicyResponse{}
	_body, _err := client.DeleteILMPolicyWithOptions(InstanceId, PolicyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除ES集群索引模版
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexTemplateResponse
func (client *Client) DeleteIndexTemplateWithOptions(InstanceId *string, IndexTemplate *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndexTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndexTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/index-templates/" + dara.PercentEncode(dara.StringValue(IndexTemplate))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除ES集群索引模版
//
// @return DeleteIndexTemplateResponse
func (client *Client) DeleteIndexTemplate(InstanceId *string, IndexTemplate *string) (_result *DeleteIndexTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexTemplateResponse{}
	_body, _err := client.DeleteIndexTemplateWithOptions(InstanceId, IndexTemplate, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(InstanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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

	if !dara.IsNil(request.DeleteType) {
		query["deleteType"] = request.DeleteType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(InstanceId *string, request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a Logstash cluster.
//
// Description:
//
// Before you call this operation, take note of the following information: After the cluster is released, the physical resources used by the cluster are reclaimed. The data stored in the cluster is deleted and cannot be recovered. The disks attached to the nodes in the cluster and the snapshots created for the cluster are released.
//
// @param request - DeleteLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogstashResponse
func (client *Client) DeleteLogstashWithOptions(InstanceId *string, request *DeleteLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLogstashResponse, _err error) {
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

	if !dara.IsNil(request.DeleteType) {
		query["deleteType"] = request.DeleteType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a Logstash cluster.
//
// Description:
//
// Before you call this operation, take note of the following information: After the cluster is released, the physical resources used by the cluster are reclaimed. The data stored in the cluster is deleted and cannot be recovered. The disks attached to the nodes in the cluster and the snapshots created for the cluster are released.
//
// @param request - DeleteLogstashRequest
//
// @return DeleteLogstashResponse
func (client *Client) DeleteLogstash(InstanceId *string, request *DeleteLogstashRequest) (_result *DeleteLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLogstashResponse{}
	_body, _err := client.DeleteLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a pipeline that is configured for a Logstash cluster.
//
// @param request - DeletePipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelinesResponse
func (client *Client) DeletePipelinesWithOptions(InstanceId *string, request *DeletePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.PipelineIds) {
		query["pipelineIds"] = request.PipelineIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePipelines"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a pipeline that is configured for a Logstash cluster.
//
// @param request - DeletePipelinesRequest
//
// @return DeletePipelinesResponse
func (client *Client) DeletePipelines(InstanceId *string, request *DeletePipelinesRequest) (_result *DeletePipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelinesResponse{}
	_body, _err := client.DeletePipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSnapshotRepoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotRepoResponse
func (client *Client) DeleteSnapshotRepoWithOptions(InstanceId *string, request *DeleteSnapshotRepoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotRepoResponse, _err error) {
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

	if !dara.IsNil(request.RepoPath) {
		query["repoPath"] = request.RepoPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshotRepo"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshot-repos"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSnapshotRepoRequest
//
// @return DeleteSnapshotRepoResponse
func (client *Client) DeleteSnapshotRepo(InstanceId *string, request *DeleteSnapshotRepoRequest) (_result *DeleteSnapshotRepoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSnapshotRepoResponse{}
	_body, _err := client.DeleteSnapshotRepoWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除服务账号vpc下的终端节点
//
// @param request - DeleteVpcEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpointWithOptions(InstanceId *string, EndpointId *string, request *DeleteVpcEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcEndpoint"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/vpc-endpoints/" + dara.PercentEncode(dara.StringValue(EndpointId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务账号vpc下的终端节点
//
// @param request - DeleteVpcEndpointRequest
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpoint(InstanceId *string, EndpointId *string, request *DeleteVpcEndpointRequest) (_result *DeleteVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.DeleteVpcEndpointWithOptions(InstanceId, EndpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of ES-operator that is installed for a specified Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > Before you install a shipper on an ACK cluster, you can call this operation to query the installation status of ES-operator for the ACK cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckOperatorResponse
func (client *Client) DescribeAckOperatorWithOptions(ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAckOperatorResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAckOperator"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ack-clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/operator"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAckOperatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of ES-operator that is installed for a specified Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > Before you install a shipper on an ACK cluster, you can call this operation to query the installation status of ES-operator for the ACK cluster.
//
// @return DescribeAckOperatorResponse
func (client *Client) DescribeAckOperator(ClusterId *string) (_result *DescribeAckOperatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAckOperatorResponse{}
	_body, _err := client.DescribeAckOperatorWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a shipper.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCollectorResponse
func (client *Client) DescribeCollectorWithOptions(ResId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeCollectorResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a shipper.
//
// @return DescribeCollectorResponse
func (client *Client) DescribeCollector(ResId *string) (_result *DescribeCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCollectorResponse{}
	_body, _err := client.DescribeCollectorWithOptions(ResId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看组合索引模板详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentIndexResponse
func (client *Client) DescribeComponentIndexWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeComponentIndexResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComponentIndex"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/component-index/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看组合索引模板详情
//
// @return DescribeComponentIndexResponse
func (client *Client) DescribeComponentIndex(InstanceId *string, name *string) (_result *DescribeComponentIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeComponentIndexResponse{}
	_body, _err := client.DescribeComponentIndexWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeConnectableClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConnectableClustersResponse
func (client *Client) DescribeConnectableClustersWithOptions(InstanceId *string, request *DescribeConnectableClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeConnectableClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlreadySetItems) {
		query["alreadySetItems"] = request.AlreadySetItems
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConnectableClusters"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/connectable-clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConnectableClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeConnectableClustersRequest
//
// @return DescribeConnectableClustersResponse
func (client *Client) DescribeConnectableClusters(InstanceId *string, request *DescribeConnectableClustersRequest) (_result *DescribeConnectableClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConnectableClustersResponse{}
	_body, _err := client.DescribeConnectableClustersWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeDeprecatedTemplate
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeprecatedTemplateResponse
func (client *Client) DescribeDeprecatedTemplateWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDeprecatedTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeprecatedTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/deprecated-templates/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeprecatedTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeDeprecatedTemplate
//
// @return DescribeDeprecatedTemplateResponse
func (client *Client) DescribeDeprecatedTemplate(InstanceId *string, name *string) (_result *DescribeDeprecatedTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDeprecatedTemplateResponse{}
	_body, _err := client.DescribeDeprecatedTemplateWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDiagnoseReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnoseReportResponse
func (client *Client) DescribeDiagnoseReportWithOptions(InstanceId *string, ReportId *string, request *DescribeDiagnoseReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDiagnoseReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnoseReport"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/reports/" + dara.PercentEncode(dara.StringValue(ReportId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnoseReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDiagnoseReportRequest
//
// @return DescribeDiagnoseReportResponse
func (client *Client) DescribeDiagnoseReport(InstanceId *string, ReportId *string, request *DescribeDiagnoseReportRequest) (_result *DescribeDiagnoseReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDiagnoseReportResponse{}
	_body, _err := client.DescribeDiagnoseReportWithOptions(InstanceId, ReportId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDiagnosisSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisSettingsResponse
func (client *Client) DescribeDiagnosisSettingsWithOptions(InstanceId *string, request *DescribeDiagnosisSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosisSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/settings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosisSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDiagnosisSettingsRequest
//
// @return DescribeDiagnosisSettingsResponse
func (client *Client) DescribeDiagnosisSettings(InstanceId *string, request *DescribeDiagnosisSettingsRequest) (_result *DescribeDiagnosisSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDiagnosisSettingsResponse{}
	_body, _err := client.DescribeDiagnosisSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群动态指标
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDynamicSettingsResponse
func (client *Client) DescribeDynamicSettingsWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDynamicSettingsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDynamicSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/dynamic-settings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDynamicSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群动态指标
//
// @return DescribeDynamicSettingsResponse
func (client *Client) DescribeDynamicSettings(InstanceId *string) (_result *DescribeDynamicSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDynamicSettingsResponse{}
	_body, _err := client.DescribeDynamicSettingsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health status of an Elasticsearch cluster.
//
// Description:
//
// An Elasticsearch cluster can be in a health state indicated by one of the following colors:
//
//   - GREEN: Primary shards and replica shards for the primary shards are normally allocated.
//
//   - YELLOW: Primary shards are normally allocated, but replica shards for the primary shards are not normally allocated.
//
//   - RED: Primary shards are not normally allocated.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticsearchHealthResponse
func (client *Client) DescribeElasticsearchHealthWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeElasticsearchHealthResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticsearchHealth"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/elasticsearch-health"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticsearchHealthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health status of an Elasticsearch cluster.
//
// Description:
//
// An Elasticsearch cluster can be in a health state indicated by one of the following colors:
//
//   - GREEN: Primary shards and replica shards for the primary shards are normally allocated.
//
//   - YELLOW: Primary shards are normally allocated, but replica shards for the primary shards are not normally allocated.
//
//   - RED: Primary shards are not normally allocated.
//
// @return DescribeElasticsearchHealthResponse
func (client *Client) DescribeElasticsearchHealth(InstanceId *string) (_result *DescribeElasticsearchHealthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeElasticsearchHealthResponse{}
	_body, _err := client.DescribeElasticsearchHealthWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeILMPolicyResponse
func (client *Client) DescribeILMPolicyWithOptions(InstanceId *string, PolicyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeILMPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeILMPolicy"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/ilm-policies/" + dara.PercentEncode(dara.StringValue(PolicyName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeILMPolicyResponse
func (client *Client) DescribeILMPolicy(InstanceId *string, PolicyName *string) (_result *DescribeILMPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeILMPolicyResponse{}
	_body, _err := client.DescribeILMPolicyWithOptions(InstanceId, PolicyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIndexTemplateResponse
func (client *Client) DescribeIndexTemplateWithOptions(InstanceId *string, IndexTemplate *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeIndexTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIndexTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/index-templates/" + dara.PercentEncode(dara.StringValue(IndexTemplate))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeIndexTemplateResponse
func (client *Client) DescribeIndexTemplate(InstanceId *string, IndexTemplate *string) (_result *DescribeIndexTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeIndexTemplateResponse{}
	_body, _err := client.DescribeIndexTemplateWithOptions(InstanceId, IndexTemplate, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The name of the dictionary file.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The name of the dictionary file.
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(InstanceId *string) (_result *DescribeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Elasticsearch集群Kibana节点settings配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKibanaSettingsResponse
func (client *Client) DescribeKibanaSettingsWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeKibanaSettingsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKibanaSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/kibana-settings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKibanaSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Elasticsearch集群Kibana节点settings配置
//
// @return DescribeKibanaSettingsResponse
func (client *Client) DescribeKibanaSettings(InstanceId *string) (_result *DescribeKibanaSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeKibanaSettingsResponse{}
	_body, _err := client.DescribeKibanaSettingsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看Logstash实例详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogstashResponse
func (client *Client) DescribeLogstashWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeLogstashResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Logstash实例详情
//
// @return DescribeLogstashResponse
func (client *Client) DescribeLogstash(InstanceId *string) (_result *DescribeLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeLogstashResponse{}
	_body, _err := client.DescribeLogstashWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePipelineResponse
func (client *Client) DescribePipelineWithOptions(InstanceId *string, PipelineId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePipelineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePipeline"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(PipelineId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return DescribePipelineResponse
func (client *Client) DescribePipeline(InstanceId *string, PipelineId *string) (_result *DescribePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePipelineResponse{}
	_body, _err := client.DescribePipelineWithOptions(InstanceId, PipelineId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the management configurations of pipelines in a Logstash cluster.
//
// @param request - DescribePipelineManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePipelineManagementConfigResponse
func (client *Client) DescribePipelineManagementConfigWithOptions(InstanceId *string, request *DescribePipelineManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePipelineManagementConfigResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePipelineManagementConfig"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipeline-management-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePipelineManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the management configurations of pipelines in a Logstash cluster.
//
// @param request - DescribePipelineManagementConfigRequest
//
// @return DescribePipelineManagementConfigResponse
func (client *Client) DescribePipelineManagementConfig(InstanceId *string, request *DescribePipelineManagementConfigRequest) (_result *DescribePipelineManagementConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePipelineManagementConfigResponse{}
	_body, _err := client.DescribePipelineManagementConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/regions"),
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

// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看备份设置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotSettingResponse
func (client *Client) DescribeSnapshotSettingWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotSettingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnapshotSetting"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshot-setting"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看备份设置
//
// @return DescribeSnapshotSettingResponse
func (client *Client) DescribeSnapshotSetting(InstanceId *string) (_result *DescribeSnapshotSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSnapshotSettingResponse{}
	_body, _err := client.DescribeSnapshotSettingWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplates"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplates(InstanceId *string) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of the X-Pack Monitoring feature of a Logstash cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeXpackMonitorConfigResponse
func (client *Client) DescribeXpackMonitorConfigWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeXpackMonitorConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeXpackMonitorConfig"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/xpack-monitor-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeXpackMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the X-Pack Monitoring feature of a Logstash cluster.
//
// @return DescribeXpackMonitorConfigResponse
func (client *Client) DescribeXpackMonitorConfig(InstanceId *string) (_result *DescribeXpackMonitorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeXpackMonitorConfigResponse{}
	_body, _err := client.DescribeXpackMonitorConfigWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 触发ES实例智能诊断
//
// @param request - DiagnoseInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiagnoseInstanceResponse
func (client *Client) DiagnoseInstanceWithOptions(InstanceId *string, request *DiagnoseInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DiagnoseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DiagnoseItems) {
		body["diagnoseItems"] = request.DiagnoseItems
	}

	if !dara.IsNil(request.Indices) {
		body["indices"] = request.Indices
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
		Action:      dara.String("DiagnoseInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/diagnose"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DiagnoseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发ES实例智能诊断
//
// @param request - DiagnoseInstanceRequest
//
// @return DiagnoseInstanceResponse
func (client *Client) DiagnoseInstance(InstanceId *string, request *DiagnoseInstanceRequest) (_result *DiagnoseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DiagnoseInstanceResponse{}
	_body, _err := client.DiagnoseInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭kibana私网
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableKibanaPvlNetworkResponse
func (client *Client) DisableKibanaPvlNetworkWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableKibanaPvlNetworkResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableKibanaPvlNetwork"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/disable-kibana-private"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableKibanaPvlNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭kibana私网
//
// @return DisableKibanaPvlNetworkResponse
func (client *Client) DisableKibanaPvlNetwork(InstanceId *string) (_result *DisableKibanaPvlNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableKibanaPvlNetworkResponse{}
	_body, _err := client.DisableKibanaPvlNetworkWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启v3 kibana私网
//
// @param request - EnableKibanaPvlNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableKibanaPvlNetworkResponse
func (client *Client) EnableKibanaPvlNetworkWithOptions(InstanceId *string, request *EnableKibanaPvlNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableKibanaPvlNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndpointName) {
		body["endpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.SecurityGroups) {
		body["securityGroups"] = request.SecurityGroups
	}

	if !dara.IsNil(request.VSwitchIdsZone) {
		body["vSwitchIdsZone"] = request.VSwitchIdsZone
	}

	if !dara.IsNil(request.VpcId) {
		body["vpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableKibanaPvlNetwork"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/enable-kibana-private"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableKibanaPvlNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启v3 kibana私网
//
// @param request - EnableKibanaPvlNetworkRequest
//
// @return EnableKibanaPvlNetworkResponse
func (client *Client) EnableKibanaPvlNetwork(InstanceId *string, request *EnableKibanaPvlNetworkRequest) (_result *EnableKibanaPvlNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableKibanaPvlNetworkResponse{}
	_body, _err := client.EnableKibanaPvlNetworkWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the estimated time that is required to restart a Logstash cluster.
//
// @param request - EstimatedLogstashRestartTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedLogstashRestartTimeResponse
func (client *Client) EstimatedLogstashRestartTimeWithOptions(InstanceId *string, request *EstimatedLogstashRestartTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EstimatedLogstashRestartTimeResponse, _err error) {
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
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EstimatedLogstashRestartTime"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/estimated-time/restart-time"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EstimatedLogstashRestartTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the estimated time that is required to restart a Logstash cluster.
//
// @param request - EstimatedLogstashRestartTimeRequest
//
// @return EstimatedLogstashRestartTimeResponse
func (client *Client) EstimatedLogstashRestartTime(InstanceId *string, request *EstimatedLogstashRestartTimeRequest) (_result *EstimatedLogstashRestartTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EstimatedLogstashRestartTimeResponse{}
	_body, _err := client.EstimatedLogstashRestartTimeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the estimated time that is required to restart an Elasticsearch cluster.
//
// @param request - EstimatedRestartTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedRestartTimeResponse
func (client *Client) EstimatedRestartTimeWithOptions(InstanceId *string, request *EstimatedRestartTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EstimatedRestartTimeResponse, _err error) {
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
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EstimatedRestartTime"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/estimated-time/restart-time"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EstimatedRestartTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the estimated time that is required to restart an Elasticsearch cluster.
//
// @param request - EstimatedRestartTimeRequest
//
// @return EstimatedRestartTimeResponse
func (client *Client) EstimatedRestartTime(InstanceId *string, request *EstimatedRestartTimeRequest) (_result *EstimatedRestartTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EstimatedRestartTimeResponse{}
	_body, _err := client.EstimatedRestartTimeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call GetClusterDataInformation to obtain the data information about the cluster.
//
// @param request - GetClusterDataInformationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDataInformationResponse
func (client *Client) GetClusterDataInformationWithOptions(request *GetClusterDataInformationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDataInformationResponse, _err error) {
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
		Action:      dara.String("GetClusterDataInformation"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/cluster/data-information"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterDataInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call GetClusterDataInformation to obtain the data information about the cluster.
//
// @param request - GetClusterDataInformationRequest
//
// @return GetClusterDataInformationResponse
func (client *Client) GetClusterDataInformation(request *GetClusterDataInformationRequest) (_result *GetClusterDataInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterDataInformationResponse{}
	_body, _err := client.GetClusterDataInformationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElastictaskResponse
func (client *Client) GetElastictaskWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetElastictaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetElastictask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/elastic-task"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetElastictaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return GetElastictaskResponse
func (client *Client) GetElastictask(InstanceId *string) (_result *GetElastictaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetElastictaskResponse{}
	_body, _err := client.GetElastictaskWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetEmonAlarmRecordStatisticsDistribute
//
// @param request - GetEmonAlarmRecordStatisticsDistributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonAlarmRecordStatisticsDistributeResponse
func (client *Client) GetEmonAlarmRecordStatisticsDistributeWithOptions(request *GetEmonAlarmRecordStatisticsDistributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonAlarmRecordStatisticsDistributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["body"] = request.Body
	}

	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.TimeEnd) {
		query["timeEnd"] = request.TimeEnd
	}

	if !dara.IsNil(request.TimeStart) {
		query["timeStart"] = request.TimeStart
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmonAlarmRecordStatisticsDistribute"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/emon/alarm-record-statistics/distribute"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmonAlarmRecordStatisticsDistributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetEmonAlarmRecordStatisticsDistribute
//
// @param request - GetEmonAlarmRecordStatisticsDistributeRequest
//
// @return GetEmonAlarmRecordStatisticsDistributeResponse
func (client *Client) GetEmonAlarmRecordStatisticsDistribute(request *GetEmonAlarmRecordStatisticsDistributeRequest) (_result *GetEmonAlarmRecordStatisticsDistributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonAlarmRecordStatisticsDistributeResponse{}
	_body, _err := client.GetEmonAlarmRecordStatisticsDistributeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取高级监控报警自定义Grafana监控报警项
//
// @param request - GetEmonGrafanaAlertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonGrafanaAlertsResponse
func (client *Client) GetEmonGrafanaAlertsWithOptions(ProjectId *string, request *GetEmonGrafanaAlertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonGrafanaAlertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmonGrafanaAlerts"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/emon/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/grafana/proxy/api/alerts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmonGrafanaAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取高级监控报警自定义Grafana监控报警项
//
// @param request - GetEmonGrafanaAlertsRequest
//
// @return GetEmonGrafanaAlertsResponse
func (client *Client) GetEmonGrafanaAlerts(ProjectId *string, request *GetEmonGrafanaAlertsRequest) (_result *GetEmonGrafanaAlertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonGrafanaAlertsResponse{}
	_body, _err := client.GetEmonGrafanaAlertsWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取高级监控报警自定义Grafana监控大盘列表
//
// @param request - GetEmonGrafanaDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonGrafanaDashboardsResponse
func (client *Client) GetEmonGrafanaDashboardsWithOptions(ProjectId *string, request *GetEmonGrafanaDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonGrafanaDashboardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmonGrafanaDashboards"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/emon/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/grafana/proxy/api/search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmonGrafanaDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取高级监控报警自定义Grafana监控大盘列表
//
// @param request - GetEmonGrafanaDashboardsRequest
//
// @return GetEmonGrafanaDashboardsResponse
func (client *Client) GetEmonGrafanaDashboards(ProjectId *string, request *GetEmonGrafanaDashboardsRequest) (_result *GetEmonGrafanaDashboardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonGrafanaDashboardsResponse{}
	_body, _err := client.GetEmonGrafanaDashboardsWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetEmonMonitorData
//
// @param request - GetEmonMonitorDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonMonitorDataResponse
func (client *Client) GetEmonMonitorDataWithOptions(ProjectId *string, request *GetEmonMonitorDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonMonitorDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmonMonitorData"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/emon/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/metrics/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmonMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetEmonMonitorData
//
// @param request - GetEmonMonitorDataRequest
//
// @return GetEmonMonitorDataResponse
func (client *Client) GetEmonMonitorData(ProjectId *string, request *GetEmonMonitorDataRequest) (_result *GetEmonMonitorDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonMonitorDataResponse{}
	_body, _err := client.GetEmonMonitorDataWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统计OpenStore实例的存储容量和使用情况
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenStoreUsageResponse
func (client *Client) GetOpenStoreUsageWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOpenStoreUsageResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpenStoreUsage"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/openstore/usage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpenStoreUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统计OpenStore实例的存储容量和使用情况
//
// @return GetOpenStoreUsageResponse
func (client *Client) GetOpenStoreUsage(InstanceId *string) (_result *GetOpenStoreUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOpenStoreUsageResponse{}
	_body, _err := client.GetOpenStoreUsageWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The maximum number of nodes.
//
// @param request - GetRegionConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionConfigurationResponse
func (client *Client) GetRegionConfigurationWithOptions(request *GetRegionConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ZoneId) {
		query["zoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegionConfiguration"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/region"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegionConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The maximum number of nodes.
//
// @param request - GetRegionConfigurationRequest
//
// @return GetRegionConfigurationResponse
func (client *Client) GetRegionConfiguration(request *GetRegionConfigurationRequest) (_result *GetRegionConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionConfigurationResponse{}
	_body, _err := client.GetRegionConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例区域商品化配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionalInstanceConfigResponse
func (client *Client) GetRegionalInstanceConfigWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionalInstanceConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegionalInstanceConfig"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/regions/instance-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegionalInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例区域商品化配置
//
// @return GetRegionalInstanceConfigResponse
func (client *Client) GetRegionalInstanceConfig() (_result *GetRegionalInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionalInstanceConfigResponse{}
	_body, _err := client.GetRegionalInstanceConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ES集群可缩容节点
//
// @param request - GetSuggestShrinkableNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuggestShrinkableNodesResponse
func (client *Client) GetSuggestShrinkableNodesWithOptions(InstanceId *string, request *GetSuggestShrinkableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSuggestShrinkableNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.IgnoreStatus) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !dara.IsNil(request.NodeType) {
		query["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSuggestShrinkableNodes"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/suggest-shrinkable-nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSuggestShrinkableNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ES集群可缩容节点
//
// @param request - GetSuggestShrinkableNodesRequest
//
// @return GetSuggestShrinkableNodesResponse
func (client *Client) GetSuggestShrinkableNodes(InstanceId *string, request *GetSuggestShrinkableNodesRequest) (_result *GetSuggestShrinkableNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSuggestShrinkableNodesResponse{}
	_body, _err := client.GetSuggestShrinkableNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取可数据迁移节点
//
// @param request - GetTransferableNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransferableNodesResponse
func (client *Client) GetTransferableNodesWithOptions(InstanceId *string, request *GetTransferableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTransferableNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.NodeType) {
		query["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTransferableNodes"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/transferable-nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTransferableNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可数据迁移节点
//
// @param request - GetTransferableNodesRequest
//
// @return GetTransferableNodesResponse
func (client *Client) GetTransferableNodes(InstanceId *string, request *GetTransferableNodesRequest) (_result *GetTransferableNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTransferableNodesResponse{}
	_body, _err := client.GetTransferableNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// Description:
//
// > Before you perform auto scaling for a cluster at the China site (aliyun.com) or you use shippers to collect logs, you must create a service-linked role.
//
// @param request - InitializeOperationRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeOperationRoleResponse
func (client *Client) InitializeOperationRoleWithOptions(request *InitializeOperationRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitializeOperationRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeOperationRole"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/user/slr"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeOperationRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// Description:
//
// > Before you perform auto scaling for a cluster at the China site (aliyun.com) or you use shippers to collect logs, you must create a service-linked role.
//
// @param request - InitializeOperationRoleRequest
//
// @return InitializeOperationRoleResponse
func (client *Client) InitializeOperationRole(request *InitializeOperationRoleRequest) (_result *InitializeOperationRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitializeOperationRoleResponse{}
	_body, _err := client.InitializeOperationRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs ES-operator for a Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > Before you install a shipper for an ACK cluster, you must call this operation to install ES-operator for the cluster.
//
// @param request - InstallAckOperatorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAckOperatorResponse
func (client *Client) InstallAckOperatorWithOptions(ClusterId *string, request *InstallAckOperatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAckOperatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAckOperator"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ack-clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/operator"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAckOperatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs ES-operator for a Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > Before you install a shipper for an ACK cluster, you must call this operation to install ES-operator for the cluster.
//
// @param request - InstallAckOperatorRequest
//
// @return InstallAckOperatorResponse
func (client *Client) InstallAckOperator(ClusterId *string, request *InstallAckOperatorRequest) (_result *InstallAckOperatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAckOperatorResponse{}
	_body, _err := client.InstallAckOperatorWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call InstallKibanaSystemPlugin to install the Kibana plug-in. The Kibana specification must be 2-Core 4 GB or higher.
//
// @param request - InstallKibanaSystemPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallKibanaSystemPluginResponse
func (client *Client) InstallKibanaSystemPluginWithOptions(InstanceId *string, request *InstallKibanaSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallKibanaSystemPluginResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallKibanaSystemPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/kibana-plugins/system/actions/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallKibanaSystemPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call InstallKibanaSystemPlugin to install the Kibana plug-in. The Kibana specification must be 2-Core 4 GB or higher.
//
// @param request - InstallKibanaSystemPluginRequest
//
// @return InstallKibanaSystemPluginResponse
func (client *Client) InstallKibanaSystemPlugin(InstanceId *string, request *InstallKibanaSystemPluginRequest) (_result *InstallKibanaSystemPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallKibanaSystemPluginResponse{}
	_body, _err := client.InstallKibanaSystemPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The returned data also contains **Headers*	- parameters, indicating that header information is returned.
//
// Description:
//
// ls-cn-oew1qbgl\\*\\*\\*\\*
//
// @param request - InstallLogstashSystemPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallLogstashSystemPluginResponse
func (client *Client) InstallLogstashSystemPluginWithOptions(InstanceId *string, request *InstallLogstashSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallLogstashSystemPluginResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallLogstashSystemPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins/system/actions/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallLogstashSystemPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The returned data also contains **Headers*	- parameters, indicating that header information is returned.
//
// Description:
//
// ls-cn-oew1qbgl\\*\\*\\*\\*
//
// @param request - InstallLogstashSystemPluginRequest
//
// @return InstallLogstashSystemPluginResponse
func (client *Client) InstallLogstashSystemPlugin(InstanceId *string, request *InstallLogstashSystemPluginRequest) (_result *InstallLogstashSystemPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallLogstashSystemPluginResponse{}
	_body, _err := client.InstallLogstashSystemPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call InstallSystemPlugin to install a system preset plug-in.
//
// @param request - InstallSystemPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallSystemPluginResponse
func (client *Client) InstallSystemPluginWithOptions(InstanceId *string, request *InstallSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallSystemPluginResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallSystemPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins/system/actions/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallSystemPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call InstallSystemPlugin to install a system preset plug-in.
//
// @param request - InstallSystemPluginRequest
//
// @return InstallSystemPluginResponse
func (client *Client) InstallSystemPlugin(InstanceId *string, request *InstallSystemPluginRequest) (_result *InstallSystemPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallSystemPluginResponse{}
	_body, _err := client.InstallSystemPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs custom plug-ins that are uploaded to the Elasticsearch console.
//
// @param request - InstallUserPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallUserPluginsResponse
func (client *Client) InstallUserPluginsWithOptions(InstanceId *string, request *InstallUserPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallUserPluginsResponse, _err error) {
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
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallUserPlugins"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins/user/actions/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallUserPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs custom plug-ins that are uploaded to the Elasticsearch console.
//
// @param request - InstallUserPluginsRequest
//
// @return InstallUserPluginsResponse
func (client *Client) InstallUserPlugins(InstanceId *string, request *InstallUserPluginsRequest) (_result *InstallUserPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallUserPluginsResponse{}
	_body, _err := client.InstallUserPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InterruptElasticsearchTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterruptElasticsearchTaskResponse
func (client *Client) InterruptElasticsearchTaskWithOptions(InstanceId *string, request *InterruptElasticsearchTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InterruptElasticsearchTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InterruptElasticsearchTask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/interrupt"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InterruptElasticsearchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - InterruptElasticsearchTaskRequest
//
// @return InterruptElasticsearchTaskResponse
func (client *Client) InterruptElasticsearchTask(InstanceId *string, request *InterruptElasticsearchTaskRequest) (_result *InterruptElasticsearchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InterruptElasticsearchTaskResponse{}
	_body, _err := client.InterruptElasticsearchTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After the task is suspended, the Logstash cluster is in the suspended state.
//
// @param request - InterruptLogstashTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterruptLogstashTaskResponse
func (client *Client) InterruptLogstashTaskWithOptions(InstanceId *string, request *InterruptLogstashTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InterruptLogstashTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InterruptLogstashTask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/interrupt"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InterruptLogstashTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After the task is suspended, the Logstash cluster is in the suspended state.
//
// @param request - InterruptLogstashTaskRequest
//
// @return InterruptLogstashTaskResponse
func (client *Client) InterruptLogstashTask(InstanceId *string, request *InterruptLogstashTaskRequest) (_result *InterruptLogstashTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InterruptLogstashTaskResponse{}
	_body, _err := client.InterruptLogstashTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Container Service for Kubernetes (ACK) clusters.
//
// @param request - ListAckClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAckClustersResponse
func (client *Client) ListAckClustersWithOptions(request *ListAckClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAckClustersResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.VpcId) {
		query["vpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAckClusters"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ack-clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAckClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Container Service for Kubernetes (ACK) clusters.
//
// @param request - ListAckClustersRequest
//
// @return ListAckClustersResponse
func (client *Client) ListAckClusters(request *ListAckClustersRequest) (_result *ListAckClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAckClustersResponse{}
	_body, _err := client.ListAckClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all namespaces in a specified Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > When you install a shipper on an ACK cluster, you must specify a namespace. You can call this operation to query all namespaces in the ACK cluster, and select a namespace based on your business requirements.
//
// @param request - ListAckNamespacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAckNamespacesResponse
func (client *Client) ListAckNamespacesWithOptions(ClusterId *string, request *ListAckNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAckNamespacesResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAckNamespaces"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ack-clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/namespaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAckNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all namespaces in a specified Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > When you install a shipper on an ACK cluster, you must specify a namespace. You can call this operation to query all namespaces in the ACK cluster, and select a namespace based on your business requirements.
//
// @param request - ListAckNamespacesRequest
//
// @return ListAckNamespacesResponse
func (client *Client) ListAckNamespaces(ClusterId *string, request *ListAckNamespacesRequest) (_result *ListAckNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAckNamespacesResponse{}
	_body, _err := client.ListAckNamespacesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更记录 变更详情
//
// @param request - ListActionRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionRecordsResponse
func (client *Client) ListActionRecordsWithOptions(InstanceId *string, request *ListActionRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListActionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionNames) {
		query["actionNames"] = request.ActionNames
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListActionRecords"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/action-records"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更记录 变更详情
//
// @param request - ListActionRecordsRequest
//
// @return ListActionRecordsResponse
func (client *Client) ListActionRecords(InstanceId *string, request *ListActionRecordsRequest) (_result *ListActionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListActionRecordsResponse{}
	_body, _err := client.ListActionRecordsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// es-cn-tl32cpgwa002l\\*\\*\\*\\*
//
// @param request - ListAllNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllNodeResponse
func (client *Client) ListAllNodeWithOptions(InstanceId *string, request *ListAllNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extended) {
		query["extended"] = request.Extended
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllNode"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// es-cn-tl32cpgwa002l\\*\\*\\*\\*
//
// @param request - ListAllNodeRequest
//
// @return ListAllNodeResponse
func (client *Client) ListAllNode(InstanceId *string, request *ListAllNodeRequest) (_result *ListAllNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAllNodeResponse{}
	_body, _err := client.ListAllNodeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例可添加的OSS引用仓库
//
// @param request - ListAlternativeSnapshotReposRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlternativeSnapshotReposResponse
func (client *Client) ListAlternativeSnapshotReposWithOptions(InstanceId *string, request *ListAlternativeSnapshotReposRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlternativeSnapshotReposResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlreadySetItems) {
		query["alreadySetItems"] = request.AlreadySetItems
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlternativeSnapshotRepos"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/alternative-snapshot-repos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlternativeSnapshotReposResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例可添加的OSS引用仓库
//
// @param request - ListAlternativeSnapshotReposRequest
//
// @return ListAlternativeSnapshotReposResponse
func (client *Client) ListAlternativeSnapshotRepos(InstanceId *string, request *ListAlternativeSnapshotReposRequest) (_result *ListAlternativeSnapshotReposResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlternativeSnapshotReposResponse{}
	_body, _err := client.ListAlternativeSnapshotReposWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Elasticsearch clusters that can be associated with a Logstash cluster when you configure the X-Pack Monitoring feature for the Logstash cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableEsInstanceIdsResponse
func (client *Client) ListAvailableEsInstanceIdsWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvailableEsInstanceIdsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableEsInstanceIds"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/available-elasticsearch-for-centralized-management"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableEsInstanceIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Elasticsearch clusters that can be associated with a Logstash cluster when you configure the X-Pack Monitoring feature for the Logstash cluster.
//
// @return ListAvailableEsInstanceIdsResponse
func (client *Client) ListAvailableEsInstanceIds(InstanceId *string) (_result *ListAvailableEsInstanceIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvailableEsInstanceIdsResponse{}
	_body, _err := client.ListAvailableEsInstanceIdsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries shippers.
//
// @param request - ListCollectorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectorsResponse
func (client *Client) ListCollectorsWithOptions(request *ListCollectorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCollectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.ResId) {
		query["resId"] = request.ResId
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCollectors"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCollectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries shippers.
//
// @param request - ListCollectorsRequest
//
// @return ListCollectorsResponse
func (client *Client) ListCollectors(request *ListCollectorsRequest) (_result *ListCollectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCollectorsResponse{}
	_body, _err := client.ListCollectorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ES集群组合索引列表
//
// @param request - ListComponentIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentIndicesResponse
func (client *Client) ListComponentIndicesWithOptions(InstanceId *string, request *ListComponentIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComponentIndicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponentIndices"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/component-index"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ES集群组合索引列表
//
// @param request - ListComponentIndicesRequest
//
// @return ListComponentIndicesResponse
func (client *Client) ListComponentIndices(InstanceId *string, request *ListComponentIndicesRequest) (_result *ListComponentIndicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComponentIndicesResponse{}
	_body, _err := client.ListComponentIndicesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取与当前实例进行网络互通的实例列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectedClustersResponse
func (client *Client) ListConnectedClustersWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConnectedClustersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnectedClusters"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/connected-clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectedClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取与当前实例进行网络互通的实例列表
//
// @return ListConnectedClustersResponse
func (client *Client) ListConnectedClusters(InstanceId *string) (_result *ListConnectedClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConnectedClustersResponse{}
	_body, _err := client.ListConnectedClustersWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据流
//
// @param request - ListDataStreamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataStreamsResponse
func (client *Client) ListDataStreamsWithOptions(InstanceId *string, request *ListDataStreamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataStreamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsManaged) {
		query["isManaged"] = request.IsManaged
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataStreams"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/data-streams"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据流
//
// @param request - ListDataStreamsRequest
//
// @return ListDataStreamsResponse
func (client *Client) ListDataStreams(InstanceId *string, request *ListDataStreamsRequest) (_result *ListDataStreamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataStreamsResponse{}
	_body, _err := client.ListDataStreamsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataTasksResponse
func (client *Client) ListDataTasksWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataTasksResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataTasks"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/data-task"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return ListDataTasksResponse
func (client *Client) ListDataTasks(InstanceId *string) (_result *ListDataTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataTasksResponse{}
	_body, _err := client.ListDataTasksWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default configuration files of shippers.
//
// @param request - ListDefaultCollectorConfigurationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDefaultCollectorConfigurationsResponse
func (client *Client) ListDefaultCollectorConfigurationsWithOptions(request *ListDefaultCollectorConfigurationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDefaultCollectorConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResType) {
		query["resType"] = request.ResType
	}

	if !dara.IsNil(request.ResVersion) {
		query["resVersion"] = request.ResVersion
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDefaultCollectorConfigurations"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/beats/default-configurations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDefaultCollectorConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default configuration files of shippers.
//
// @param request - ListDefaultCollectorConfigurationsRequest
//
// @return ListDefaultCollectorConfigurationsResponse
func (client *Client) ListDefaultCollectorConfigurations(request *ListDefaultCollectorConfigurationsRequest) (_result *ListDefaultCollectorConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDefaultCollectorConfigurationsResponse{}
	_body, _err := client.ListDefaultCollectorConfigurationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListDeprecatedTemplates
//
// @param request - ListDeprecatedTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeprecatedTemplatesResponse
func (client *Client) ListDeprecatedTemplatesWithOptions(InstanceId *string, request *ListDeprecatedTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeprecatedTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeprecatedTemplates"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/deprecated-templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeprecatedTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListDeprecatedTemplates
//
// @param request - ListDeprecatedTemplatesRequest
//
// @return ListDeprecatedTemplatesResponse
func (client *Client) ListDeprecatedTemplates(InstanceId *string, request *ListDeprecatedTemplatesRequest) (_result *ListDeprecatedTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDeprecatedTemplatesResponse{}
	_body, _err := client.ListDeprecatedTemplatesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the indexes for health diagnosis performed on an Elasticsearch cluster.
//
// @param request - ListDiagnoseIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseIndicesResponse
func (client *Client) ListDiagnoseIndicesWithOptions(InstanceId *string, request *ListDiagnoseIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseIndicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnoseIndices"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/indices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnoseIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the indexes for health diagnosis performed on an Elasticsearch cluster.
//
// @param request - ListDiagnoseIndicesRequest
//
// @return ListDiagnoseIndicesResponse
func (client *Client) ListDiagnoseIndices(InstanceId *string, request *ListDiagnoseIndicesRequest) (_result *ListDiagnoseIndicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnoseIndicesResponse{}
	_body, _err := client.ListDiagnoseIndicesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群诊断报告列表
//
// @param request - ListDiagnoseReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseReportResponse
func (client *Client) ListDiagnoseReportWithOptions(InstanceId *string, request *ListDiagnoseReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Detail) {
		query["detail"] = request.Detail
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Trigger) {
		query["trigger"] = request.Trigger
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnoseReport"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/reports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnoseReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群诊断报告列表
//
// @param request - ListDiagnoseReportRequest
//
// @return ListDiagnoseReportResponse
func (client *Client) ListDiagnoseReport(InstanceId *string, request *ListDiagnoseReportRequest) (_result *ListDiagnoseReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnoseReportResponse{}
	_body, _err := client.ListDiagnoseReportWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IDs of the historical intelligent O&M reports of an Elasticsearch cluster.
//
// @param request - ListDiagnoseReportIdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseReportIdsResponse
func (client *Client) ListDiagnoseReportIdsWithOptions(InstanceId *string, request *ListDiagnoseReportIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseReportIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Trigger) {
		query["trigger"] = request.Trigger
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnoseReportIds"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/report-ids"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnoseReportIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of the historical intelligent O&M reports of an Elasticsearch cluster.
//
// @param request - ListDiagnoseReportIdsRequest
//
// @return ListDiagnoseReportIdsResponse
func (client *Client) ListDiagnoseReportIds(InstanceId *string, request *ListDiagnoseReportIdsRequest) (_result *ListDiagnoseReportIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnoseReportIdsResponse{}
	_body, _err := client.ListDiagnoseReportIdsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The diagnostic item is used to check whether data write requests of a cluster are accumulated. If data write requests are accumulated, a bulk rejection occurs. This may cause data loss and severely consume system resources.
//
// @param request - ListDiagnosisItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosisItemsResponse
func (client *Client) ListDiagnosisItemsWithOptions(request *ListDiagnosisItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnosisItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnosisItems"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/items"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnosisItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The diagnostic item is used to check whether data write requests of a cluster are accumulated. If data write requests are accumulated, a bulk rejection occurs. This may cause data loss and severely consume system resources.
//
// @param request - ListDiagnosisItemsRequest
//
// @return ListDiagnosisItemsResponse
func (client *Client) ListDiagnosisItems(request *ListDiagnosisItemsRequest) (_result *ListDiagnosisItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnosisItemsResponse{}
	_body, _err := client.ListDiagnosisItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDictInformationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictInformationResponse
func (client *Client) ListDictInformationWithOptions(InstanceId *string, request *ListDictInformationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalyzerType) {
		query["analyzerType"] = request.AnalyzerType
	}

	if !dara.IsNil(request.BucketName) {
		query["bucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Key) {
		query["key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDictInformation"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/dict/_info"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDictInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDictInformationRequest
//
// @return ListDictInformationResponse
func (client *Client) ListDictInformation(InstanceId *string, request *ListDictInformationRequest) (_result *ListDictInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictInformationResponse{}
	_body, _err := client.ListDictInformationWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified type of dictionary.
//
// @param request - ListDictsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictsResponse
func (client *Client) ListDictsWithOptions(InstanceId *string, request *ListDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalyzerType) {
		query["analyzerType"] = request.AnalyzerType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDicts"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/dicts"),
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
// Queries the details of a specified type of dictionary.
//
// @param request - ListDictsRequest
//
// @return ListDictsResponse
func (client *Client) ListDicts(InstanceId *string, request *ListDictsRequest) (_result *ListDictsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictsResponse{}
	_body, _err := client.ListDictsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询ecs实例
//
// Description:
//
// *Important*	- To call this operation, you must create the Aliyun Elasticsearch AccessingOOSRole and the system service role AliyunOOSAccessingECS 4ESRole to Elasticsearch the service account to obtain the ECS access permissions of the primary account. For more information, see [Collect ECS service logs](https://help.aliyun.com/document_detail/146446.html).
//
// @param request - ListEcsInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcsInstancesResponse
func (client *Client) ListEcsInstancesWithOptions(request *ListEcsInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcsInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcsInstanceIds) {
		query["ecsInstanceIds"] = request.EcsInstanceIds
	}

	if !dara.IsNil(request.EcsInstanceName) {
		query["ecsInstanceName"] = request.EcsInstanceName
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	if !dara.IsNil(request.VpcId) {
		query["vpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEcsInstances"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ecs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEcsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询ecs实例
//
// Description:
//
// *Important*	- To call this operation, you must create the Aliyun Elasticsearch AccessingOOSRole and the system service role AliyunOOSAccessingECS 4ESRole to Elasticsearch the service account to obtain the ECS access permissions of the primary account. For more information, see [Collect ECS service logs](https://help.aliyun.com/document_detail/146446.html).
//
// @param request - ListEcsInstancesRequest
//
// @return ListEcsInstancesResponse
func (client *Client) ListEcsInstances(request *ListEcsInstancesRequest) (_result *ListEcsInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsInstancesResponse{}
	_body, _err := client.ListEcsInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the driver files of a Logstash cluster.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExtendfilesResponse
func (client *Client) ListExtendfilesWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExtendfilesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExtendfiles"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/extendfiles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExtendfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the driver files of a Logstash cluster.
//
// @return ListExtendfilesResponse
func (client *Client) ListExtendfiles(InstanceId *string) (_result *ListExtendfilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExtendfilesResponse{}
	_body, _err := client.ListExtendfilesWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListILMPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListILMPoliciesResponse
func (client *Client) ListILMPoliciesWithOptions(InstanceId *string, request *ListILMPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListILMPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["policyName"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListILMPolicies"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/ilm-policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListILMPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListILMPoliciesRequest
//
// @return ListILMPoliciesResponse
func (client *Client) ListILMPolicies(InstanceId *string, request *ListILMPoliciesRequest) (_result *ListILMPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListILMPoliciesResponse{}
	_body, _err := client.ListILMPoliciesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListIndexTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexTemplatesResponse
func (client *Client) ListIndexTemplatesWithOptions(InstanceId *string, request *ListIndexTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexTemplate) {
		query["indexTemplate"] = request.IndexTemplate
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIndexTemplates"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/index-templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndexTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIndexTemplatesRequest
//
// @return ListIndexTemplatesResponse
func (client *Client) ListIndexTemplates(InstanceId *string, request *ListIndexTemplatesRequest) (_result *ListIndexTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexTemplatesResponse{}
	_body, _err := client.ListIndexTemplatesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Elasticsearch实例列表
//
// @param request - ListInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EsVersion) {
		query["esVersion"] = request.EsVersion
	}

	if !dara.IsNil(request.InstanceCategory) {
		query["instanceCategory"] = request.InstanceCategory
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PaymentType) {
		query["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	if !dara.IsNil(request.VpcId) {
		query["vpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["zoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Elasticsearch实例列表
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集群触发的硬件运维事件列表
//
// @param tmpReq - ListInstanceHistoryEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHistoryEventsResponse
func (client *Client) ListInstanceHistoryEventsWithOptions(tmpReq *ListInstanceHistoryEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceHistoryEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListInstanceHistoryEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventCycleStatus) {
		request.EventCycleStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventCycleStatus, dara.String("eventCycleStatus"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.EventLevel) {
		request.EventLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventLevel, dara.String("eventLevel"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.EventType) {
		request.EventTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventType, dara.String("eventType"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventCreateEndTime) {
		query["eventCreateEndTime"] = request.EventCreateEndTime
	}

	if !dara.IsNil(request.EventCreateStartTime) {
		query["eventCreateStartTime"] = request.EventCreateStartTime
	}

	if !dara.IsNil(request.EventCycleStatusShrink) {
		query["eventCycleStatus"] = request.EventCycleStatusShrink
	}

	if !dara.IsNil(request.EventExecuteEndTime) {
		query["eventExecuteEndTime"] = request.EventExecuteEndTime
	}

	if !dara.IsNil(request.EventExecuteStartTime) {
		query["eventExecuteStartTime"] = request.EventExecuteStartTime
	}

	if !dara.IsNil(request.EventFinashEndTime) {
		query["eventFinashEndTime"] = request.EventFinashEndTime
	}

	if !dara.IsNil(request.EventFinashStartTime) {
		query["eventFinashStartTime"] = request.EventFinashStartTime
	}

	if !dara.IsNil(request.EventLevelShrink) {
		query["eventLevel"] = request.EventLevelShrink
	}

	if !dara.IsNil(request.EventTypeShrink) {
		query["eventType"] = request.EventTypeShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeIP) {
		query["nodeIP"] = request.NodeIP
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceHistoryEvents"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/events"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceHistoryEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集群触发的硬件运维事件列表
//
// @param request - ListInstanceHistoryEventsRequest
//
// @return ListInstanceHistoryEventsResponse
func (client *Client) ListInstanceHistoryEvents(request *ListInstanceHistoryEventsRequest) (_result *ListInstanceHistoryEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceHistoryEventsResponse{}
	_body, _err := client.ListInstanceHistoryEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前实例先特定的索引列表
//
// @param request - ListInstanceIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceIndicesResponse
func (client *Client) ListInstanceIndicesWithOptions(InstanceId *string, request *ListInstanceIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceIndicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["all"] = request.All
	}

	if !dara.IsNil(request.IsManaged) {
		query["isManaged"] = request.IsManaged
	}

	if !dara.IsNil(request.IsOpenstore) {
		query["isOpenstore"] = request.IsOpenstore
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceIndices"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/indices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前实例先特定的索引列表
//
// @param request - ListInstanceIndicesRequest
//
// @return ListInstanceIndicesResponse
func (client *Client) ListInstanceIndices(InstanceId *string, request *ListInstanceIndicesRequest) (_result *ListInstanceIndicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceIndicesResponse{}
	_body, _err := client.ListInstanceIndicesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Kibana plug-ins.
//
// @param request - ListKibanaPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKibanaPluginsResponse
func (client *Client) ListKibanaPluginsWithOptions(InstanceId *string, request *ListKibanaPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKibanaPluginsResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKibanaPlugins"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/kibana-plugins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKibanaPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Kibana plug-ins.
//
// @param request - ListKibanaPluginsRequest
//
// @return ListKibanaPluginsResponse
func (client *Client) ListKibanaPlugins(InstanceId *string, request *ListKibanaPluginsRequest) (_result *ListKibanaPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKibanaPluginsResponse{}
	_body, _err := client.ListKibanaPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询kibana私网连接信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKibanaPvlNetworkResponse
func (client *Client) ListKibanaPvlNetworkWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKibanaPvlNetworkResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKibanaPvlNetwork"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/get-kibana-private"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKibanaPvlNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询kibana私网连接信息
//
// @return ListKibanaPvlNetworkResponse
func (client *Client) ListKibanaPvlNetwork(InstanceId *string) (_result *ListKibanaPvlNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKibanaPvlNetworkResponse{}
	_body, _err := client.ListKibanaPvlNetworkWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Logstash集群列表
//
// @param request - ListLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashResponse
func (client *Client) ListLogstashWithOptions(request *ListLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Logstash集群列表
//
// @param request - ListLogstashRequest
//
// @return ListLogstashResponse
func (client *Client) ListLogstash(request *ListLogstashRequest) (_result *ListLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogstashResponse{}
	_body, _err := client.ListLogstashWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Logstash日志
//
// @param request - ListLogstashLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashLogResponse
func (client *Client) ListLogstashLogWithOptions(InstanceId *string, request *ListLogstashLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogstashLog"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/search-log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogstashLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Logstash日志
//
// @param request - ListLogstashLogRequest
//
// @return ListLogstashLogResponse
func (client *Client) ListLogstashLog(InstanceId *string, request *ListLogstashLogRequest) (_result *ListLogstashLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogstashLogResponse{}
	_body, _err := client.ListLogstashLogWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Logstash插件列表
//
// @param request - ListLogstashPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashPluginsResponse
func (client *Client) ListLogstashPluginsWithOptions(InstanceId *string, request *ListLogstashPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogstashPlugins"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogstashPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Logstash插件列表
//
// @param request - ListLogstashPluginsRequest
//
// @return ListLogstashPluginsResponse
func (client *Client) ListLogstashPlugins(InstanceId *string, request *ListLogstashPluginsRequest) (_result *ListLogstashPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogstashPluginsResponse{}
	_body, _err := client.ListLogstashPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statuses of Elastic Compute Service (ECS) instances on which a shipper is installed.
//
// @param request - ListNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(ResId *string, request *ListNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcsInstanceIds) {
		query["ecsInstanceIds"] = request.EcsInstanceIds
	}

	if !dara.IsNil(request.EcsInstanceName) {
		query["ecsInstanceName"] = request.EcsInstanceName
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statuses of Elastic Compute Service (ECS) instances on which a shipper is installed.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(ResId *string, request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListPipeline
//
// @param request - ListPipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineResponse
func (client *Client) ListPipelineWithOptions(InstanceId *string, request *ListPipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineResponse, _err error) {
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

	if !dara.IsNil(request.PipelineId) {
		query["pipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipeline"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListPipeline
//
// @param request - ListPipelineRequest
//
// @return ListPipelineResponse
func (client *Client) ListPipeline(InstanceId *string, request *ListPipelineRequest) (_result *ListPipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineResponse{}
	_body, _err := client.ListPipelineWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The error message returned.
//
// @param request - ListPipelineIdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineIdsResponse
func (client *Client) ListPipelineIdsWithOptions(InstanceId *string, request *ListPipelineIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineIdsResponse, _err error) {
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
		Action:      dara.String("ListPipelineIds"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipeline-ids"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The error message returned.
//
// @param request - ListPipelineIdsRequest
//
// @return ListPipelineIdsResponse
func (client *Client) ListPipelineIds(InstanceId *string, request *ListPipelineIdsRequest) (_result *ListPipelineIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineIdsResponse{}
	_body, _err := client.ListPipelineIdsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ES系统插件列表
//
// @param request - ListPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsResponse
func (client *Client) ListPluginsWithOptions(InstanceId *string, request *ListPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlugins"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ES系统插件列表
//
// @param request - ListPluginsRequest
//
// @return ListPluginsResponse
func (client *Client) ListPlugins(InstanceId *string, request *ListPluginsRequest) (_result *ListPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPluginsResponse{}
	_body, _err := client.ListPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看Elasticsearch集群各种类型的日志
//
// @param request - ListSearchLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchLogResponse
func (client *Client) ListSearchLogWithOptions(InstanceId *string, request *ListSearchLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchLog"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/search-log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Elasticsearch集群各种类型的日志
//
// @param request - ListSearchLogRequest
//
// @return ListSearchLogResponse
func (client *Client) ListSearchLog(InstanceId *string, request *ListSearchLogRequest) (_result *ListSearchLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSearchLogResponse{}
	_body, _err := client.ListSearchLogWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about shards that are being restored or shards that are restored. By default, this operation returns only the information about shards that are being restored after you call this operation.
//
// Description:
//
// > The restoration of a shard is a process of synchronizing data from a primary shard to a replica shard. After the restoration is complete, the replica shard is available for data searches.
//
// @param request - ListShardRecoveriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShardRecoveriesResponse
func (client *Client) ListShardRecoveriesWithOptions(InstanceId *string, request *ListShardRecoveriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShardRecoveriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveOnly) {
		query["activeOnly"] = request.ActiveOnly
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListShardRecoveries"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/cat-recovery"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListShardRecoveriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about shards that are being restored or shards that are restored. By default, this operation returns only the information about shards that are being restored after you call this operation.
//
// Description:
//
// > The restoration of a shard is a process of synchronizing data from a primary shard to a replica shard. After the restoration is complete, the replica shard is available for data searches.
//
// @param request - ListShardRecoveriesRequest
//
// @return ListShardRecoveriesResponse
func (client *Client) ListShardRecoveries(InstanceId *string, request *ListShardRecoveriesRequest) (_result *ListShardRecoveriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShardRecoveriesResponse{}
	_body, _err := client.ListShardRecoveriesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取跨集群索引仓库列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotReposByInstanceIdResponse
func (client *Client) ListSnapshotReposByInstanceIdWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSnapshotReposByInstanceIdResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshotReposByInstanceId"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshot-repos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotReposByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取跨集群索引仓库列表
//
// @return ListSnapshotReposByInstanceIdResponse
func (client *Client) ListSnapshotReposByInstanceId(InstanceId *string) (_result *ListSnapshotReposByInstanceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotReposByInstanceIdResponse{}
	_body, _err := client.ListSnapshotReposByInstanceIdWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看资源和标签关系
//
// @param request - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tags"),
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
// 查看资源和标签关系
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
// 查看所有已常见的标签
//
// @param request - ListTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithOptions(request *ListTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTags"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tags/all-tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看所有已常见的标签
//
// @param request - ListTagsRequest
//
// @return ListTagsResponse
func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statuses of endpoints in the virtual private cloud (VPC) within the Elasticsearch service account.
//
// @param request - ListVpcEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpointsWithOptions(InstanceId *string, request *ListVpcEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpoints"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/vpc-endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statuses of endpoints in the virtual private cloud (VPC) within the Elasticsearch service account.
//
// @param request - ListVpcEndpointsRequest
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpoints(InstanceId *string, request *ListVpcEndpointsRequest) (_result *ListVpcEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.ListVpcEndpointsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the MigrateToOtherZone to migrate the nodes in the specified zone to the destination zone.
//
// Description:
//
// If the specifications in your zone are insufficient, you can upgrade your instance to nodes in another zone. Before calling this interface, you must ensure that:
//
//   - The error message returned because the current account is in a zone that has sufficient resources.
//
//     After migrating nodes with current specifications to another zone, you need to manually [upgrade cluster](https://help.aliyun.com/document_detail/96650.html) because the cluster will not be upgraded during the migration process. Therefore, select a zone with sufficient resources to avoid cluster upgrade failure. We recommend that you choose new zones that are in lower alphabetical order. For example, for cn-hangzhou-e and cn-hangzhou-h zones, choose cn-hangzhou-h first.
//
//   - The cluster is in the healthy state.
//
//     Can be passed`  GET _cat/health?v  `command to view the health status of the cluster.
//
// @param request - MigrateToOtherZoneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateToOtherZoneResponse
func (client *Client) MigrateToOtherZoneWithOptions(InstanceId *string, request *MigrateToOtherZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateToOtherZone"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/migrate-zones"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the MigrateToOtherZone to migrate the nodes in the specified zone to the destination zone.
//
// Description:
//
// If the specifications in your zone are insufficient, you can upgrade your instance to nodes in another zone. Before calling this interface, you must ensure that:
//
//   - The error message returned because the current account is in a zone that has sufficient resources.
//
//     After migrating nodes with current specifications to another zone, you need to manually [upgrade cluster](https://help.aliyun.com/document_detail/96650.html) because the cluster will not be upgraded during the migration process. Therefore, select a zone with sufficient resources to avoid cluster upgrade failure. We recommend that you choose new zones that are in lower alphabetical order. For example, for cn-hangzhou-e and cn-hangzhou-h zones, choose cn-hangzhou-h first.
//
//   - The cluster is in the healthy state.
//
//     Can be passed`  GET _cat/health?v  `command to view the health status of the cluster.
//
// @param request - MigrateToOtherZoneRequest
//
// @return MigrateToOtherZoneResponse
func (client *Client) MigrateToOtherZone(InstanceId *string, request *MigrateToOtherZoneRequest) (_result *MigrateToOtherZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.MigrateToOtherZoneWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the Elastic Compute Service (ECS) instances on which a shipper is installed.
//
// @param request - ModifyDeployMachineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeployMachineResponse
func (client *Client) ModifyDeployMachineWithOptions(ResId *string, request *ModifyDeployMachineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDeployMachineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeployMachine"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/actions/modify-deploy-machines"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeployMachineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the Elastic Compute Service (ECS) instances on which a shipper is installed.
//
// @param request - ModifyDeployMachineRequest
//
// @return ModifyDeployMachineResponse
func (client *Client) ModifyDeployMachine(ResId *string, request *ModifyDeployMachineRequest) (_result *ModifyDeployMachineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDeployMachineResponse{}
	_body, _err := client.ModifyDeployMachineWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyElastictaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElastictaskResponse
func (client *Client) ModifyElastictaskWithOptions(InstanceId *string, request *ModifyElastictaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyElastictaskResponse, _err error) {
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
		Action:      dara.String("ModifyElastictask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/elastic-task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElastictaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyElastictaskRequest
//
// @return ModifyElastictaskResponse
func (client *Client) ModifyElastictask(InstanceId *string, request *ModifyElastictaskRequest) (_result *ModifyElastictaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyElastictaskResponse{}
	_body, _err := client.ModifyElastictaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ## RequestBody
//
// You must also specify the following parameters in the RequestBody parameter to specify the maintenance window information.
//
// | Parameter | Type | Required | Example | Description |
//
// | --------- | ---- | -------- | ------- | ----------- |
//
// | maintainStartTime | String | No | 02:00Z | The start time of the maintenance window. Specify the time in the HH:mmZ format. The time must be in UTC. |
//
// | maintainEndTime | String | No | 06:00Z | The end time of the maintenance window. Specify the time in the HH:mmZ format. The time must be displayed in UTC. |
//
// | openMaintainTime | boolean | Yes | true | Specifies whether to enable the maintenance window feature. Only **true*	- is supported, indicating that the feature is enabled. |
//
// Examples:
//
// ```
//
// {
//
//	"openMaintainTime": true,
//
//	"maintainStartTime": "03:00Z",
//
//	"maintainEndTime": "04:00Z"
//
// }
//
// ```
//
// Description:
//
// es-cn-n6w1o1x0w001c\\*\\*\\*\\*
//
// @param request - ModifyInstanceMaintainTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTimeWithOptions(InstanceId *string, request *ModifyInstanceMaintainTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceMaintainTime"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/modify-maintaintime"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ## RequestBody
//
// You must also specify the following parameters in the RequestBody parameter to specify the maintenance window information.
//
// | Parameter | Type | Required | Example | Description |
//
// | --------- | ---- | -------- | ------- | ----------- |
//
// | maintainStartTime | String | No | 02:00Z | The start time of the maintenance window. Specify the time in the HH:mmZ format. The time must be in UTC. |
//
// | maintainEndTime | String | No | 06:00Z | The end time of the maintenance window. Specify the time in the HH:mmZ format. The time must be displayed in UTC. |
//
// | openMaintainTime | boolean | Yes | true | Specifies whether to enable the maintenance window feature. Only **true*	- is supported, indicating that the feature is enabled. |
//
// Examples:
//
// ```
//
// {
//
//	"openMaintainTime": true,
//
//	"maintainStartTime": "03:00Z",
//
//	"maintainEndTime": "04:00Z"
//
// }
//
// ```
//
// Description:
//
// es-cn-n6w1o1x0w001c\\*\\*\\*\\*
//
// @param request - ModifyInstanceMaintainTimeRequest
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTime(InstanceId *string, request *ModifyInstanceMaintainTimeRequest) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyInstanceMaintainTimeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// >  If you want to add an IP address whitelist, you can set the modifyMode parameter only to Cover. If you set this parameter to Delete or Append, you can only update an IP address whitelist.
//
//   - If you set the modifyMode parameter to Cover and leave the ips parameter empty, the system deletes the specified whitelist. If the whitelist specified by using the groupName parameter does not exist, the system creates such a whitelist.
//
//   - If you set the modifyMode parameter to Delete, at least one IP address must be retained for the specified whitelist.
//
//   - If you set the modifyMode parameter to Append, you must make sure that the specified whitelist exists. Otherwise, the system reports the NotFound error.
//
// Description:
//
// The ID of the cluster.
//
// @param request - ModifyWhiteIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWhiteIpsResponse
func (client *Client) ModifyWhiteIpsWithOptions(InstanceId *string, request *ModifyWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyWhiteIpsResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.ModifyMode) {
		body["modifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.NetworkType) {
		body["networkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeType) {
		body["nodeType"] = request.NodeType
	}

	if !dara.IsNil(request.WhiteIpGroup) {
		body["whiteIpGroup"] = request.WhiteIpGroup
	}

	if !dara.IsNil(request.WhiteIpList) {
		body["whiteIpList"] = request.WhiteIpList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWhiteIps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/modify-white-ips"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// >  If you want to add an IP address whitelist, you can set the modifyMode parameter only to Cover. If you set this parameter to Delete or Append, you can only update an IP address whitelist.
//
//   - If you set the modifyMode parameter to Cover and leave the ips parameter empty, the system deletes the specified whitelist. If the whitelist specified by using the groupName parameter does not exist, the system creates such a whitelist.
//
//   - If you set the modifyMode parameter to Delete, at least one IP address must be retained for the specified whitelist.
//
//   - If you set the modifyMode parameter to Append, you must make sure that the specified whitelist exists. Otherwise, the system reports the NotFound error.
//
// Description:
//
// The ID of the cluster.
//
// @param request - ModifyWhiteIpsRequest
//
// @return ModifyWhiteIpsResponse
func (client *Client) ModifyWhiteIps(InstanceId *string, request *ModifyWhiteIpsRequest) (_result *ModifyWhiteIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyWhiteIpsResponse{}
	_body, _err := client.ModifyWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Migrates an Elasticsearch cluster to a specified resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(InstanceId *string, request *MoveResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resourcegroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates an Elasticsearch cluster to a specified resource group.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(InstanceId *string, request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OpenDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDiagnosisResponse
func (client *Client) OpenDiagnosisWithOptions(InstanceId *string, request *OpenDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenDiagnosis"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/open-diagnosis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - OpenDiagnosisRequest
//
// @return OpenDiagnosisResponse
func (client *Client) OpenDiagnosis(InstanceId *string, request *OpenDiagnosisRequest) (_result *OpenDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenDiagnosisResponse{}
	_body, _err := client.OpenDiagnosisWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// >  To ensure data security, we recommend that you enable HTTPS.
//
// @param request - OpenHttpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenHttpsResponse
func (client *Client) OpenHttpsWithOptions(InstanceId *string, request *OpenHttpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenHttpsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenHttps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/open-https"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenHttpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// >  To ensure data security, we recommend that you enable HTTPS.
//
// @param request - OpenHttpsRequest
//
// @return OpenHttpsResponse
func (client *Client) OpenHttps(InstanceId *string, request *OpenHttpsRequest) (_result *OpenHttpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenHttpsResponse{}
	_body, _err := client.OpenHttpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PostEmonTryAlarmRule
//
// @param request - PostEmonTryAlarmRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEmonTryAlarmRuleResponse
func (client *Client) PostEmonTryAlarmRuleWithOptions(ProjectId *string, AlarmGroupId *string, request *PostEmonTryAlarmRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PostEmonTryAlarmRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		query["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PostEmonTryAlarmRule"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/emon/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/alarm-groups/" + dara.PercentEncode(dara.StringValue(AlarmGroupId)) + "/alarm-rules/_test"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PostEmonTryAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PostEmonTryAlarmRule
//
// @param request - PostEmonTryAlarmRuleRequest
//
// @return PostEmonTryAlarmRuleResponse
func (client *Client) PostEmonTryAlarmRule(ProjectId *string, AlarmGroupId *string, request *PostEmonTryAlarmRuleRequest) (_result *PostEmonTryAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PostEmonTryAlarmRuleResponse{}
	_body, _err := client.PostEmonTryAlarmRuleWithOptions(ProjectId, AlarmGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RecommendTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecommendTemplatesResponse
func (client *Client) RecommendTemplatesWithOptions(InstanceId *string, request *RecommendTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecommendTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UsageScenario) {
		query["usageScenario"] = request.UsageScenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecommendTemplates"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/recommended-templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecommendTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RecommendTemplatesRequest
//
// @return RecommendTemplatesResponse
func (client *Client) RecommendTemplates(InstanceId *string, request *RecommendTemplatesRequest) (_result *RecommendTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecommendTemplatesResponse{}
	_body, _err := client.RecommendTemplatesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs a shipper that failed to be installed when you create the shipper.
//
// @param request - ReinstallCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinstallCollectorResponse
func (client *Client) ReinstallCollectorWithOptions(ResId *string, request *ReinstallCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReinstallCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReinstallCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/actions/reinstall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReinstallCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs a shipper that failed to be installed when you create the shipper.
//
// @param request - ReinstallCollectorRequest
//
// @return ReinstallCollectorResponse
func (client *Client) ReinstallCollector(ResId *string, request *ReinstallCollectorRequest) (_result *ReinstallCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReinstallCollectorResponse{}
	_body, _err := client.ReinstallCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call RenewInstance to renew a subscription instance.
//
// @param request - RenewInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(InstanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/renew"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call RenewInstance to renew a subscription instance.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(InstanceId *string, request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a Logstash cluster.
//
// @param request - RenewLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewLogstashResponse
func (client *Client) RenewLogstashWithOptions(InstanceId *string, request *RenewLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewLogstashResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/renew"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a Logstash cluster.
//
// @param request - RenewLogstashRequest
//
// @return RenewLogstashResponse
func (client *Client) RenewLogstash(InstanceId *string, request *RenewLogstashRequest) (_result *RenewLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewLogstashResponse{}
	_body, _err := client.RenewLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a shipper.
//
// @param request - RestartCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartCollectorResponse
func (client *Client) RestartCollectorWithOptions(ResId *string, request *RestartCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/actions/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a shipper.
//
// @param request - RestartCollectorRequest
//
// @return RestartCollectorResponse
func (client *Client) RestartCollector(ResId *string, request *RestartCollectorRequest) (_result *RestartCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartCollectorResponse{}
	_body, _err := client.RestartCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to restart a specified Elasticsearch instance.
//
// Description:
//
// >  After the instance is restarted, the instance enters the activating state. After the instance is restarted, its status changes to active. Alibaba Cloud Elasticsearch supports restarting a single node. Restarting a node can be divided into normal restart and blue-green restart.
//
// @param request - RestartInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(InstanceId *string, request *RestartInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to restart a specified Elasticsearch instance.
//
// Description:
//
// >  After the instance is restarted, the instance enters the activating state. After the instance is restarted, its status changes to active. Alibaba Cloud Elasticsearch supports restarting a single node. Restarting a node can be divided into normal restart and blue-green restart.
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(InstanceId *string, request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启Logstash集群
//
// @param request - RestartLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartLogstashResponse
func (client *Client) RestartLogstashWithOptions(InstanceId *string, request *RestartLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartLogstashResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BatchCount) {
		body["batchCount"] = request.BatchCount
	}

	if !dara.IsNil(request.BlueGreenDep) {
		body["blueGreenDep"] = request.BlueGreenDep
	}

	if !dara.IsNil(request.NodeTypes) {
		body["nodeTypes"] = request.NodeTypes
	}

	if !dara.IsNil(request.Nodes) {
		body["nodes"] = request.Nodes
	}

	if !dara.IsNil(request.RestartType) {
		body["restartType"] = request.RestartType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启Logstash集群
//
// @param request - RestartLogstashRequest
//
// @return RestartLogstashResponse
func (client *Client) RestartLogstash(InstanceId *string, request *RestartLogstashRequest) (_result *RestartLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartLogstashResponse{}
	_body, _err := client.RestartLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResumeElasticsearchTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeElasticsearchTaskResponse
func (client *Client) ResumeElasticsearchTaskWithOptions(InstanceId *string, request *ResumeElasticsearchTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeElasticsearchTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeElasticsearchTask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeElasticsearchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResumeElasticsearchTaskRequest
//
// @return ResumeElasticsearchTaskResponse
func (client *Client) ResumeElasticsearchTask(InstanceId *string, request *ResumeElasticsearchTaskRequest) (_result *ResumeElasticsearchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeElasticsearchTaskResponse{}
	_body, _err := client.ResumeElasticsearchTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a change task of a Logstash cluster. After the task is resumed, the Logstash cluster is in the activating state.
//
// @param request - ResumeLogstashTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeLogstashTaskResponse
func (client *Client) ResumeLogstashTaskWithOptions(InstanceId *string, request *ResumeLogstashTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeLogstashTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeLogstashTask"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeLogstashTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a change task of a Logstash cluster. After the task is resumed, the Logstash cluster is in the activating state.
//
// @param request - ResumeLogstashTaskRequest
//
// @return ResumeLogstashTaskResponse
func (client *Client) ResumeLogstashTask(InstanceId *string, request *ResumeLogstashTaskRequest) (_result *ResumeLogstashTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeLogstashTaskResponse{}
	_body, _err := client.ResumeLogstashTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 滚动数据流，生成新索引
//
// @param request - RolloverDataStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RolloverDataStreamResponse
func (client *Client) RolloverDataStreamWithOptions(InstanceId *string, DataStream *string, request *RolloverDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RolloverDataStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RolloverDataStream"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/data-streams/" + dara.PercentEncode(dara.StringValue(DataStream)) + "/rollover"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RolloverDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 滚动数据流，生成新索引
//
// @param request - RolloverDataStreamRequest
//
// @return RolloverDataStreamResponse
func (client *Client) RolloverDataStream(InstanceId *string, DataStream *string, request *RolloverDataStreamRequest) (_result *RolloverDataStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RolloverDataStreamResponse{}
	_body, _err := client.RolloverDataStreamWithOptions(InstanceId, DataStream, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs pipelines in a Logstash cluster.
//
// @param request - RunPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPipelinesResponse
func (client *Client) RunPipelinesWithOptions(InstanceId *string, request *RunPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunPipelinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPipelines"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines/action/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs pipelines in a Logstash cluster.
//
// @param request - RunPipelinesRequest
//
// @return RunPipelinesResponse
func (client *Client) RunPipelines(InstanceId *string, request *RunPipelinesRequest) (_result *RunPipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunPipelinesResponse{}
	_body, _err := client.RunPipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ES集群缩节点
//
// @param request - ShrinkNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShrinkNodeResponse
func (client *Client) ShrinkNodeWithOptions(InstanceId *string, request *ShrinkNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ShrinkNodeResponse, _err error) {
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

	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.IgnoreStatus) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !dara.IsNil(request.NodeType) {
		query["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShrinkNode"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/shrink"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ShrinkNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ES集群缩节点
//
// @param request - ShrinkNodeRequest
//
// @return ShrinkNodeResponse
func (client *Client) ShrinkNode(InstanceId *string, request *ShrinkNodeRequest) (_result *ShrinkNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ShrinkNodeResponse{}
	_body, _err := client.ShrinkNodeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a collector to collect data.
//
// @param request - StartCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCollectorResponse
func (client *Client) StartCollectorWithOptions(ResId *string, request *StartCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/actions/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a collector to collect data.
//
// @param request - StartCollectorRequest
//
// @return StartCollectorResponse
func (client *Client) StartCollector(ResId *string, request *StartCollectorRequest) (_result *StartCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartCollectorResponse{}
	_body, _err := client.StartCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a shipper.
//
// @param request - StopCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCollectorResponse
func (client *Client) StopCollectorWithOptions(ResId *string, request *StopCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/actions/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a shipper.
//
// @param request - StopCollectorRequest
//
// @return StopCollectorResponse
func (client *Client) StopCollector(ResId *string, request *StopCollectorRequest) (_result *StopCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopCollectorResponse{}
	_body, _err := client.StopCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops pipelines in a Logstash cluster.
//
// @param request - StopPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPipelinesResponse
func (client *Client) StopPipelinesWithOptions(InstanceId *string, request *StopPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopPipelinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPipelines"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops pipelines in a Logstash cluster.
//
// @param request - StopPipelinesRequest
//
// @return StopPipelinesResponse
func (client *Client) StopPipelines(InstanceId *string, request *StopPipelinesRequest) (_result *StopPipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopPipelinesResponse{}
	_body, _err := client.StopPipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The information about the clusters and tags.
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
	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tags"),
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
// The information about the clusters and tags.
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
// 缩节点，数据迁移
//
// @param request - TransferNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferNodeResponse
func (client *Client) TransferNodeWithOptions(InstanceId *string, request *TransferNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeType) {
		query["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferNode"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/transfer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 缩节点，数据迁移
//
// @param request - TransferNodeRequest
//
// @return TransferNodeResponse
func (client *Client) TransferNode(InstanceId *string, request *TransferNodeRequest) (_result *TransferNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransferNodeResponse{}
	_body, _err := client.TransferNodeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开关ES集群及Kibana节点公私网访问
//
// @param request - TriggerNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerNetworkResponse
func (client *Client) TriggerNetworkWithOptions(InstanceId *string, request *TriggerNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TriggerNetworkResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		body["actionType"] = request.ActionType
	}

	if !dara.IsNil(request.NetworkType) {
		body["networkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeType) {
		body["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerNetwork"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/network-trigger"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开关ES集群及Kibana节点公私网访问
//
// @param request - TriggerNetworkRequest
//
// @return TriggerNetworkResponse
func (client *Client) TriggerNetwork(InstanceId *string, request *TriggerNetworkRequest) (_result *TriggerNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TriggerNetworkResponse{}
	_body, _err := client.TriggerNetworkWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 可用区关机
//
// @param request - TurnOffZoneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TurnOffZoneResponse
func (client *Client) TurnOffZoneWithOptions(instanceId *string, request *TurnOffZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TurnOffZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Zone) {
		query["zone"] = request.Zone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TurnOffZone"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/turnOff-zone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TurnOffZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 可用区关机
//
// @param request - TurnOffZoneRequest
//
// @return TurnOffZoneResponse
func (client *Client) TurnOffZone(instanceId *string, request *TurnOffZoneRequest) (_result *TurnOffZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TurnOffZoneResponse{}
	_body, _err := client.TurnOffZoneWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 可用区开机
//
// @param request - TurnOnZoneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TurnOnZoneResponse
func (client *Client) TurnOnZoneWithOptions(instanceId *string, request *TurnOnZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TurnOnZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Zone) {
		query["zone"] = request.Zone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TurnOnZone"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/turnOn-zone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TurnOnZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 可用区开机
//
// @param request - TurnOnZoneRequest
//
// @return TurnOnZoneResponse
func (client *Client) TurnOnZone(instanceId *string, request *TurnOnZoneRequest) (_result *TurnOnZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TurnOnZoneResponse{}
	_body, _err := client.TurnOnZoneWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the UninstallKibanaPlugin to uninstall the Kibana plug-in.
//
// @param request - UninstallKibanaPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallKibanaPluginResponse
func (client *Client) UninstallKibanaPluginWithOptions(InstanceId *string, request *UninstallKibanaPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallKibanaPluginResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallKibanaPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/kibana-plugins/actions/uninstall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallKibanaPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the UninstallKibanaPlugin to uninstall the Kibana plug-in.
//
// @param request - UninstallKibanaPluginRequest
//
// @return UninstallKibanaPluginResponse
func (client *Client) UninstallKibanaPlugin(InstanceId *string, request *UninstallKibanaPluginRequest) (_result *UninstallKibanaPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallKibanaPluginResponse{}
	_body, _err := client.UninstallKibanaPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载Logstash实例已安装的插件
//
// @param request - UninstallLogstashPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallLogstashPluginResponse
func (client *Client) UninstallLogstashPluginWithOptions(InstanceId *string, request *UninstallLogstashPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallLogstashPluginResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallLogstashPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins/actions/uninstall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallLogstashPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载Logstash实例已安装的插件
//
// @param request - UninstallLogstashPluginRequest
//
// @return UninstallLogstashPluginResponse
func (client *Client) UninstallLogstashPlugin(InstanceId *string, request *UninstallLogstashPluginRequest) (_result *UninstallLogstashPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallLogstashPluginResponse{}
	_body, _err := client.UninstallLogstashPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UninstallPlugin to uninstall the preset plug-in.
//
// @param request - UninstallPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallPluginResponse
func (client *Client) UninstallPluginWithOptions(InstanceId *string, request *UninstallPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallPluginResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/plugins/actions/uninstall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UninstallPlugin to uninstall the preset plug-in.
//
// @param request - UninstallPluginRequest
//
// @return UninstallPluginResponse
func (client *Client) UninstallPlugin(InstanceId *string, request *UninstallPluginRequest) (_result *UninstallPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallPluginResponse{}
	_body, _err := client.UninstallPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除ES集群实例的用户可见标签
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - You can only delete user tags.
//
// > User labels are manually added to instances by users. A system Tag is a tag that Alibaba Cloud services add to instances. System labels are divided into visible labels and invisible labels.
//
//   - If you delete a resource tag relationship that is not associated with any resources, you must delete the tags.
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		query["TagKeys"] = request.TagKeys
	}

	if !dara.IsNil(request.Body) {
		query["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tags"),
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
// 删除ES集群实例的用户可见标签
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - You can only delete user tags.
//
// > User labels are manually added to instances by users. A system Tag is a tag that Alibaba Cloud services add to instances. System labels are divided into visible labels and invisible labels.
//
//   - If you delete a resource tag relationship that is not associated with any resources, you must delete the tags.
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
// 修改ES集群密码
//
// Description:
//
// 5A2CFF0E-5718-45B5-9D4D-70B3FF\\*\\*\\*\\*
//
// @param request - UpdateAdminPasswordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdminPasswordResponse
func (client *Client) UpdateAdminPasswordWithOptions(InstanceId *string, request *UpdateAdminPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAdminPasswordResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.EsAdminPassword) {
		body["esAdminPassword"] = request.EsAdminPassword
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdminPassword"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/admin-pwd"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdminPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ES集群密码
//
// Description:
//
// 5A2CFF0E-5718-45B5-9D4D-70B3FF\\*\\*\\*\\*
//
// @param request - UpdateAdminPasswordRequest
//
// @return UpdateAdminPasswordResponse
func (client *Client) UpdateAdminPassword(InstanceId *string, request *UpdateAdminPasswordRequest) (_result *UpdateAdminPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAdminPasswordResponse{}
	_body, _err := client.UpdateAdminPasswordWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdateAdvancedSetting to change the garbage collector configuration for the specified instance.
//
// @param request - UpdateAdvancedSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdvancedSettingResponse
func (client *Client) UpdateAdvancedSettingWithOptions(InstanceId *string, request *UpdateAdvancedSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAdvancedSettingResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdvancedSetting"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/update-advanced-setting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdvancedSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdateAdvancedSetting to change the garbage collector configuration for the specified instance.
//
// @param request - UpdateAdvancedSettingRequest
//
// @return UpdateAdvancedSettingResponse
func (client *Client) UpdateAdvancedSetting(InstanceId *string, request *UpdateAdvancedSettingRequest) (_result *UpdateAdvancedSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAdvancedSettingResponse{}
	_body, _err := client.UpdateAdvancedSettingWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the dictionary file of the analysis-aliws plug-in.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - Elasticsearch V5.X clusters do not support the analysis-aliws plug-in.
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateAliwsDictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliwsDictResponse
func (client *Client) UpdateAliwsDictWithOptions(InstanceId *string, request *UpdateAliwsDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAliwsDictResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAliwsDict"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/aliws-dict"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAliwsDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the dictionary file of the analysis-aliws plug-in.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - Elasticsearch V5.X clusters do not support the analysis-aliws plug-in.
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateAliwsDictRequest
//
// @return UpdateAliwsDictResponse
func (client *Client) UpdateAliwsDict(InstanceId *string, request *UpdateAliwsDictRequest) (_result *UpdateAliwsDictResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAliwsDictResponse{}
	_body, _err := client.UpdateAliwsDictWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateBlackIps is deprecated
//
// Summary:
//
// 修改ES实例访问黑名单，已废弃
//
// @param request - UpdateBlackIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBlackIpsResponse
func (client *Client) UpdateBlackIpsWithOptions(InstanceId *string, request *UpdateBlackIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBlackIpsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBlackIps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/black-ips"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBlackIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateBlackIps is deprecated
//
// Summary:
//
// 修改ES实例访问黑名单，已废弃
//
// @param request - UpdateBlackIpsRequest
//
// @return UpdateBlackIpsResponse
// Deprecated
func (client *Client) UpdateBlackIps(InstanceId *string, request *UpdateBlackIpsRequest) (_result *UpdateBlackIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBlackIpsResponse{}
	_body, _err := client.UpdateBlackIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a shipper.
//
// @param request - UpdateCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectorResponse
func (client *Client) UpdateCollectorWithOptions(ResId *string, request *UpdateCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCollectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCollector"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a shipper.
//
// @param request - UpdateCollectorRequest
//
// @return UpdateCollectorResponse
func (client *Client) UpdateCollector(ResId *string, request *UpdateCollectorRequest) (_result *UpdateCollectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCollectorResponse{}
	_body, _err := client.UpdateCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a shipper.
//
// @param request - UpdateCollectorNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectorNameResponse
func (client *Client) UpdateCollectorNameWithOptions(ResId *string, request *UpdateCollectorNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCollectorNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCollectorName"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/collectors/" + dara.PercentEncode(dara.StringValue(ResId)) + "/actions/rename"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCollectorNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a shipper.
//
// @param request - UpdateCollectorNameRequest
//
// @return UpdateCollectorNameResponse
func (client *Client) UpdateCollectorName(ResId *string, request *UpdateCollectorNameRequest) (_result *UpdateCollectorNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCollectorNameResponse{}
	_body, _err := client.UpdateCollectorNameWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改ES集群动态索引
//
// @param request - UpdateComponentIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentIndexResponse
func (client *Client) UpdateComponentIndexWithOptions(InstanceId *string, name *string, request *UpdateComponentIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComponentIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Meta) {
		body["_meta"] = request.Meta
	}

	if !dara.IsNil(request.Template) {
		body["template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComponentIndex"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/component-index/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ES集群动态索引
//
// @param request - UpdateComponentIndexRequest
//
// @return UpdateComponentIndexResponse
func (client *Client) UpdateComponentIndex(InstanceId *string, name *string, request *UpdateComponentIndexRequest) (_result *UpdateComponentIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateComponentIndexResponse{}
	_body, _err := client.UpdateComponentIndexWithOptions(InstanceId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改elasticsearch实例名称名称
//
// @param request - UpdateDescriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDescriptionResponse
func (client *Client) UpdateDescriptionWithOptions(InstanceId *string, request *UpdateDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDescriptionResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDescription"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/description"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改elasticsearch实例名称名称
//
// @param request - UpdateDescriptionRequest
//
// @return UpdateDescriptionResponse
func (client *Client) UpdateDescription(InstanceId *string, request *UpdateDescriptionRequest) (_result *UpdateDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDescriptionResponse{}
	_body, _err := client.UpdateDescriptionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdateDiagnosisSettings to update the instance of intelligent operation&maintenance (O&M) scene settings.
//
// @param request - UpdateDiagnosisSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDiagnosisSettingsResponse
func (client *Client) UpdateDiagnosisSettingsWithOptions(InstanceId *string, request *UpdateDiagnosisSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDiagnosisSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDiagnosisSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/diagnosis/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/settings"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDiagnosisSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdateDiagnosisSettings to update the instance of intelligent operation&maintenance (O&M) scene settings.
//
// @param request - UpdateDiagnosisSettingsRequest
//
// @return UpdateDiagnosisSettingsResponse
func (client *Client) UpdateDiagnosisSettings(InstanceId *string, request *UpdateDiagnosisSettingsRequest) (_result *UpdateDiagnosisSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDiagnosisSettingsResponse{}
	_body, _err := client.UpdateDiagnosisSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a dictionary of an Elasticsearch cluster.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateDictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDictResponse
func (client *Client) UpdateDictWithOptions(InstanceId *string, request *UpdateDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDictResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDict"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/dict"),
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
// Updates a dictionary of an Elasticsearch cluster.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateDictRequest
//
// @return UpdateDictResponse
func (client *Client) UpdateDict(InstanceId *string, request *UpdateDictRequest) (_result *UpdateDictResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDictResponse{}
	_body, _err := client.UpdateDictWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改集群动态配置
//
// @param request - UpdateDynamicSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDynamicSettingsResponse
func (client *Client) UpdateDynamicSettingsWithOptions(InstanceId *string, request *UpdateDynamicSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDynamicSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Mode) {
		query["mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDynamicSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/dynamic-settings"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDynamicSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改集群动态配置
//
// @param request - UpdateDynamicSettingsRequest
//
// @return UpdateDynamicSettingsResponse
func (client *Client) UpdateDynamicSettings(InstanceId *string, request *UpdateDynamicSettingsRequest) (_result *UpdateDynamicSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDynamicSettingsResponse{}
	_body, _err := client.UpdateDynamicSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateExtendConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtendConfigResponse
func (client *Client) UpdateExtendConfigWithOptions(InstanceId *string, request *UpdateExtendConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExtendConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExtendConfig"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/extend-configs/actions/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExtendConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateExtendConfigRequest
//
// @return UpdateExtendConfigResponse
func (client *Client) UpdateExtendConfig(InstanceId *string, request *UpdateExtendConfigRequest) (_result *UpdateExtendConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExtendConfigResponse{}
	_body, _err := client.UpdateExtendConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the driver files of a Logstash cluster.
//
// Description:
//
// When you call this operation, take note of the following items: You can call this operation only to delete the driver files that are uploaded to a Logstash cluster in the Alibaba Cloud Management Console. You can add or modify driver files only in the Alibaba Cloud Management Console.
//
// @param request - UpdateExtendfilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtendfilesResponse
func (client *Client) UpdateExtendfilesWithOptions(InstanceId *string, request *UpdateExtendfilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExtendfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExtendfiles"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/extendfiles"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExtendfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the driver files of a Logstash cluster.
//
// Description:
//
// When you call this operation, take note of the following items: You can call this operation only to delete the driver files that are uploaded to a Logstash cluster in the Alibaba Cloud Management Console. You can add or modify driver files only in the Alibaba Cloud Management Console.
//
// @param request - UpdateExtendfilesRequest
//
// @return UpdateExtendfilesResponse
func (client *Client) UpdateExtendfiles(InstanceId *string, request *UpdateExtendfilesRequest) (_result *UpdateExtendfilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExtendfilesResponse{}
	_body, _err := client.UpdateExtendfilesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a rolling update for the IK dictionaries of an Elasticsearch cluster.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateHotIkDictsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotIkDictsResponse
func (client *Client) UpdateHotIkDictsWithOptions(InstanceId *string, request *UpdateHotIkDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHotIkDictsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHotIkDicts"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/ik-hot-dict"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHotIkDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a rolling update for the IK dictionaries of an Elasticsearch cluster.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateHotIkDictsRequest
//
// @return UpdateHotIkDictsResponse
func (client *Client) UpdateHotIkDicts(InstanceId *string, request *UpdateHotIkDictsRequest) (_result *UpdateHotIkDictsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHotIkDictsResponse{}
	_body, _err := client.UpdateHotIkDictsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改ES集群索引生命周期策略
//
// @param request - UpdateILMPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateILMPolicyResponse
func (client *Client) UpdateILMPolicyWithOptions(InstanceId *string, PolicyName *string, request *UpdateILMPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateILMPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateILMPolicy"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/ilm-policies/" + dara.PercentEncode(dara.StringValue(PolicyName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ES集群索引生命周期策略
//
// @param request - UpdateILMPolicyRequest
//
// @return UpdateILMPolicyResponse
func (client *Client) UpdateILMPolicy(InstanceId *string, PolicyName *string, request *UpdateILMPolicyRequest) (_result *UpdateILMPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateILMPolicyResponse{}
	_body, _err := client.UpdateILMPolicyWithOptions(InstanceId, PolicyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改ES集群索引模版配置
//
// @param request - UpdateIndexTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIndexTemplateResponse
func (client *Client) UpdateIndexTemplateWithOptions(InstanceId *string, IndexTemplate *string, request *UpdateIndexTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIndexTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIndexTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/index-templates/" + dara.PercentEncode(dara.StringValue(IndexTemplate))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ES集群索引模版配置
//
// @param request - UpdateIndexTemplateRequest
//
// @return UpdateIndexTemplateResponse
func (client *Client) UpdateIndexTemplate(InstanceId *string, IndexTemplate *string, request *UpdateIndexTemplateRequest) (_result *UpdateIndexTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIndexTemplateResponse{}
	_body, _err := client.UpdateIndexTemplateWithOptions(InstanceId, IndexTemplate, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改ES集群节点配置
//
// Description:
//
// es-cn-n6w1ptcb30009\\*\\*\\*\\*
//
// @param request - UpdateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	if !dara.IsNil(request.OrderActionType) {
		query["orderActionType"] = request.OrderActionType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientNodeConfiguration) {
		body["clientNodeConfiguration"] = request.ClientNodeConfiguration
	}

	if !dara.IsNil(request.ElasticDataNodeConfiguration) {
		body["elasticDataNodeConfiguration"] = request.ElasticDataNodeConfiguration
	}

	if !dara.IsNil(request.InstanceCategory) {
		body["instanceCategory"] = request.InstanceCategory
	}

	if !dara.IsNil(request.KibanaConfiguration) {
		body["kibanaConfiguration"] = request.KibanaConfiguration
	}

	if !dara.IsNil(request.MasterConfiguration) {
		body["masterConfiguration"] = request.MasterConfiguration
	}

	if !dara.IsNil(request.NodeAmount) {
		body["nodeAmount"] = request.NodeAmount
	}

	if !dara.IsNil(request.NodeSpec) {
		body["nodeSpec"] = request.NodeSpec
	}

	if !dara.IsNil(request.WarmNodeConfiguration) {
		body["warmNodeConfiguration"] = request.WarmNodeConfiguration
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ES集群节点配置
//
// Description:
//
// es-cn-n6w1ptcb30009\\*\\*\\*\\*
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(InstanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdateInstanceChargeType to change the billing method of a pay-as-you-go instance to subscription.
//
// @param request - UpdateInstanceChargeTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceChargeTypeResponse
func (client *Client) UpdateInstanceChargeTypeWithOptions(InstanceId *string, request *UpdateInstanceChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceChargeTypeResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceChargeType"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/convert-pay-type"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdateInstanceChargeType to change the billing method of a pay-as-you-go instance to subscription.
//
// @param request - UpdateInstanceChargeTypeRequest
//
// @return UpdateInstanceChargeTypeResponse
func (client *Client) UpdateInstanceChargeType(InstanceId *string, request *UpdateInstanceChargeTypeRequest) (_result *UpdateInstanceChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceChargeTypeResponse{}
	_body, _err := client.UpdateInstanceChargeTypeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdateInstanceSettings to update the YML configuration of a specified instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// When the instance is in the activating, invalid, or inactive state, you cannot update the configuration.
//
// @param request - UpdateInstanceSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceSettingsResponse
func (client *Client) UpdateInstanceSettingsWithOptions(InstanceId *string, request *UpdateInstanceSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceSettingsResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	if !dara.IsNil(request.UpdateStrategy) {
		query["updateStrategy"] = request.UpdateStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/instance-settings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdateInstanceSettings to update the YML configuration of a specified instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// When the instance is in the activating, invalid, or inactive state, you cannot update the configuration.
//
// @param request - UpdateInstanceSettingsRequest
//
// @return UpdateInstanceSettingsResponse
func (client *Client) UpdateInstanceSettings(InstanceId *string, request *UpdateInstanceSettingsRequest) (_result *UpdateInstanceSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceSettingsResponse{}
	_body, _err := client.UpdateInstanceSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新kibana私网链接
//
// @param request - UpdateKibanaPvlNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaPvlNetworkResponse
func (client *Client) UpdateKibanaPvlNetworkWithOptions(InstanceId *string, request *UpdateKibanaPvlNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaPvlNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PvlId) {
		query["pvlId"] = request.PvlId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndpointName) {
		body["endpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.SecurityGroups) {
		body["securityGroups"] = request.SecurityGroups
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKibanaPvlNetwork"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/update-kibana-private"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKibanaPvlNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新kibana私网链接
//
// @param request - UpdateKibanaPvlNetworkRequest
//
// @return UpdateKibanaPvlNetworkResponse
func (client *Client) UpdateKibanaPvlNetwork(InstanceId *string, request *UpdateKibanaPvlNetworkRequest) (_result *UpdateKibanaPvlNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKibanaPvlNetworkResponse{}
	_body, _err := client.UpdateKibanaPvlNetworkWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdateKibanaSettings to modify the Kibana configuration. Currently, you can only modify the Kibana language configuration.
//
// @param request - UpdateKibanaSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaSettingsResponse
func (client *Client) UpdateKibanaSettingsWithOptions(InstanceId *string, request *UpdateKibanaSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaSettingsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKibanaSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/update-kibana-settings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKibanaSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdateKibanaSettings to modify the Kibana configuration. Currently, you can only modify the Kibana language configuration.
//
// @param request - UpdateKibanaSettingsRequest
//
// @return UpdateKibanaSettingsResponse
func (client *Client) UpdateKibanaSettings(InstanceId *string, request *UpdateKibanaSettingsRequest) (_result *UpdateKibanaSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKibanaSettingsResponse{}
	_body, _err := client.UpdateKibanaSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an IP address whitelist for access to the Kibana console of a specified Elasticsearch cluster.
//
// Description:
//
//	  Before you call this operation, you must make sure that the cluster is not in the activating, invalid, or inactive state.
//
//		- You can update an IP address whitelist by using the following parameters:
//
//	    	- kibanaIPWhitelist
//
//	    	- modifyMode and whiteIpGroup
//
//		- You cannot specify private IP addresses for public IP address whitelists and cannot specify public IP addresses for private IP address whitelists.
//
// @param request - UpdateKibanaWhiteIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaWhiteIpsResponse
func (client *Client) UpdateKibanaWhiteIpsWithOptions(InstanceId *string, request *UpdateKibanaWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaWhiteIpsResponse, _err error) {
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

	if !dara.IsNil(request.ModifyMode) {
		query["modifyMode"] = request.ModifyMode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KibanaIPWhitelist) {
		body["kibanaIPWhitelist"] = request.KibanaIPWhitelist
	}

	if !dara.IsNil(request.WhiteIpGroup) {
		body["whiteIpGroup"] = request.WhiteIpGroup
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKibanaWhiteIps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/kibana-white-ips"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKibanaWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an IP address whitelist for access to the Kibana console of a specified Elasticsearch cluster.
//
// Description:
//
//	  Before you call this operation, you must make sure that the cluster is not in the activating, invalid, or inactive state.
//
//		- You can update an IP address whitelist by using the following parameters:
//
//	    	- kibanaIPWhitelist
//
//	    	- modifyMode and whiteIpGroup
//
//		- You cannot specify private IP addresses for public IP address whitelists and cannot specify public IP addresses for private IP address whitelists.
//
// @param request - UpdateKibanaWhiteIpsRequest
//
// @return UpdateKibanaWhiteIpsResponse
func (client *Client) UpdateKibanaWhiteIps(InstanceId *string, request *UpdateKibanaWhiteIpsRequest) (_result *UpdateKibanaWhiteIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKibanaWhiteIpsResponse{}
	_body, _err := client.UpdateKibanaWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Logstash节点规格磁盘配置
//
// @param request - UpdateLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashResponse
func (client *Client) UpdateLogstashWithOptions(InstanceId *string, request *UpdateLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeAmount) {
		body["nodeAmount"] = request.NodeAmount
	}

	if !dara.IsNil(request.NodeSpec) {
		body["nodeSpec"] = request.NodeSpec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLogstash"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Logstash节点规格磁盘配置
//
// @param request - UpdateLogstashRequest
//
// @return UpdateLogstashResponse
func (client *Client) UpdateLogstash(InstanceId *string, request *UpdateLogstashRequest) (_result *UpdateLogstashResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashResponse{}
	_body, _err := client.UpdateLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches the billing method of a Logstash cluster from pay-as-you-go to subscription.
//
// @param request - UpdateLogstashChargeTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashChargeTypeResponse
func (client *Client) UpdateLogstashChargeTypeWithOptions(InstanceId *string, request *UpdateLogstashChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashChargeTypeResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLogstashChargeType"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/convert-pay-type"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLogstashChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the billing method of a Logstash cluster from pay-as-you-go to subscription.
//
// @param request - UpdateLogstashChargeTypeRequest
//
// @return UpdateLogstashChargeTypeResponse
func (client *Client) UpdateLogstashChargeType(InstanceId *string, request *UpdateLogstashChargeTypeRequest) (_result *UpdateLogstashChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashChargeTypeResponse{}
	_body, _err := client.UpdateLogstashChargeTypeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a specified Logstash cluster.
//
// Description:
//
// When you call this operation, take note of the following items: You cannot change the name of a cluster that is in the activating, invalid, or inactive state.
//
// @param request - UpdateLogstashDescriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashDescriptionResponse
func (client *Client) UpdateLogstashDescriptionWithOptions(InstanceId *string, request *UpdateLogstashDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashDescriptionResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLogstashDescription"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/description"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLogstashDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a specified Logstash cluster.
//
// Description:
//
// When you call this operation, take note of the following items: You cannot change the name of a cluster that is in the activating, invalid, or inactive state.
//
// @param request - UpdateLogstashDescriptionRequest
//
// @return UpdateLogstashDescriptionResponse
func (client *Client) UpdateLogstashDescription(InstanceId *string, request *UpdateLogstashDescriptionRequest) (_result *UpdateLogstashDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashDescriptionResponse{}
	_body, _err := client.UpdateLogstashDescriptionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified Logstash cluster.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// If the instance is in the Active (activating), Invalid (invalid), and Inactive (inactive) state, the information cannot be updated.
//
// @param request - UpdateLogstashSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashSettingsResponse
func (client *Client) UpdateLogstashSettingsWithOptions(InstanceId *string, request *UpdateLogstashSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashSettingsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLogstashSettings"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/instance-settings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLogstashSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified Logstash cluster.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// If the instance is in the Active (activating), Invalid (invalid), and Inactive (inactive) state, the information cannot be updated.
//
// @param request - UpdateLogstashSettingsRequest
//
// @return UpdateLogstashSettingsResponse
func (client *Client) UpdateLogstashSettings(InstanceId *string, request *UpdateLogstashSettingsRequest) (_result *UpdateLogstashSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashSettingsResponse{}
	_body, _err := client.UpdateLogstashSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Logstash管道配置
//
// @param request - UpdatePipelineManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineManagementConfigResponse
func (client *Client) UpdatePipelineManagementConfigWithOptions(InstanceId *string, request *UpdatePipelineManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineManagementConfigResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Endpoints) {
		body["endpoints"] = request.Endpoints
	}

	if !dara.IsNil(request.EsInstanceId) {
		body["esInstanceId"] = request.EsInstanceId
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.PipelineIds) {
		body["pipelineIds"] = request.PipelineIds
	}

	if !dara.IsNil(request.PipelineManagementType) {
		body["pipelineManagementType"] = request.PipelineManagementType
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipelineManagementConfig"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipeline-management-config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Logstash管道配置
//
// @param request - UpdatePipelineManagementConfigRequest
//
// @return UpdatePipelineManagementConfigResponse
func (client *Client) UpdatePipelineManagementConfig(InstanceId *string, request *UpdatePipelineManagementConfigRequest) (_result *UpdatePipelineManagementConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineManagementConfigResponse{}
	_body, _err := client.UpdatePipelineManagementConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a pipeline of a Logstash cluster.
//
// @param request - UpdatePipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelinesResponse
func (client *Client) UpdatePipelinesWithOptions(InstanceId *string, request *UpdatePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelinesResponse, _err error) {
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

	if !dara.IsNil(request.Trigger) {
		query["trigger"] = request.Trigger
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipelines"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/pipelines"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a pipeline of a Logstash cluster.
//
// @param request - UpdatePipelinesRequest
//
// @return UpdatePipelinesResponse
func (client *Client) UpdatePipelines(InstanceId *string, request *UpdatePipelinesRequest) (_result *UpdatePipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelinesResponse{}
	_body, _err := client.UpdatePipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ## RequestBody
//
// | Property | Type | Required | Example | Description |
//
// | -------- | ---- | -------- | ------- | ----------- |
//
// | privateNetworkIpWhiteList | List<String> | No | ["0.0.XX.XX","10.2.XX.XX","192.168.XX.XX/25"] | The list of IP address whitelists. This parameter is available if whiteIpGroup is left empty. The value of this parameter updates the IP address whitelist configurations in the Default whitelist group.
//
// You cannot configure both privateNetworkIpWhiteList and whiteIpGroup. |
//
// | whiteIpGroup | Object | No |  | You can update the whitelist configurations of an instance by using a whitelist group. You can update only one whitelist group.
//
// You cannot configure both privateNetworkIpWhiteList and whiteIpGroup. |
//
// | └ groupName | String | No | test_group_name | The group name of the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// | └ ips | List<String> | No | ["0.0.0.0", "10.2.XX.XX"] | The list of IP addresses in the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// > **Notice*	- The addition and deletion of whitelist groups are implemented by calling modifyMode to Cover. Delete and Append cannot add or delete whitelist groups at the same time. You can only modify the IP address list in the whitelist group. Take note of the following items: - If the modifyMode parameter is set to Cover, the whitelist group is deleted if ips is empty. If groupName is not in the list of existing whitelist group names, a whitelist group is created.
//
// - If the modifyMode parameter is set to Delete, you must retain at least one IP address for the deleted ips.
//
// - If the modifyMode parameter is set to Append, make sure that the whitelist group name has been created. Otherwise, the NotFound error message appears.
//
// Description:
//
// >  In the following returned example, only the parameters in the returned data list are guaranteed to be included, and the parameters not mentioned are for reference only. For more information about the parameters, see [ListInstance](https://help.aliyun.com/document_detail/142230.html). You cannot force a dependency in a program to get these parameters.
//
// @param request - UpdatePrivateNetworkWhiteIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateNetworkWhiteIpsResponse
func (client *Client) UpdatePrivateNetworkWhiteIpsWithOptions(InstanceId *string, request *UpdatePrivateNetworkWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrivateNetworkWhiteIpsResponse, _err error) {
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

	if !dara.IsNil(request.ModifyMode) {
		query["modifyMode"] = request.ModifyMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateNetworkWhiteIps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/private-network-white-ips"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateNetworkWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ## RequestBody
//
// | Property | Type | Required | Example | Description |
//
// | -------- | ---- | -------- | ------- | ----------- |
//
// | privateNetworkIpWhiteList | List<String> | No | ["0.0.XX.XX","10.2.XX.XX","192.168.XX.XX/25"] | The list of IP address whitelists. This parameter is available if whiteIpGroup is left empty. The value of this parameter updates the IP address whitelist configurations in the Default whitelist group.
//
// You cannot configure both privateNetworkIpWhiteList and whiteIpGroup. |
//
// | whiteIpGroup | Object | No |  | You can update the whitelist configurations of an instance by using a whitelist group. You can update only one whitelist group.
//
// You cannot configure both privateNetworkIpWhiteList and whiteIpGroup. |
//
// | └ groupName | String | No | test_group_name | The group name of the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// | └ ips | List<String> | No | ["0.0.0.0", "10.2.XX.XX"] | The list of IP addresses in the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// > **Notice*	- The addition and deletion of whitelist groups are implemented by calling modifyMode to Cover. Delete and Append cannot add or delete whitelist groups at the same time. You can only modify the IP address list in the whitelist group. Take note of the following items: - If the modifyMode parameter is set to Cover, the whitelist group is deleted if ips is empty. If groupName is not in the list of existing whitelist group names, a whitelist group is created.
//
// - If the modifyMode parameter is set to Delete, you must retain at least one IP address for the deleted ips.
//
// - If the modifyMode parameter is set to Append, make sure that the whitelist group name has been created. Otherwise, the NotFound error message appears.
//
// Description:
//
// >  In the following returned example, only the parameters in the returned data list are guaranteed to be included, and the parameters not mentioned are for reference only. For more information about the parameters, see [ListInstance](https://help.aliyun.com/document_detail/142230.html). You cannot force a dependency in a program to get these parameters.
//
// @param request - UpdatePrivateNetworkWhiteIpsRequest
//
// @return UpdatePrivateNetworkWhiteIpsResponse
func (client *Client) UpdatePrivateNetworkWhiteIps(InstanceId *string, request *UpdatePrivateNetworkWhiteIpsRequest) (_result *UpdatePrivateNetworkWhiteIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePrivateNetworkWhiteIpsResponse{}
	_body, _err := client.UpdatePrivateNetworkWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdatePublicNetwork to open or close the public network address of the specified elasticsearch instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// When the instance is in the activating, invalid, or inactive state, its configuration cannot be updated.
//
// @param request - UpdatePublicNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicNetworkResponse
func (client *Client) UpdatePublicNetworkWithOptions(InstanceId *string, request *UpdatePublicNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicNetworkResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePublicNetwork"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/public-network"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePublicNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call UpdatePublicNetwork to open or close the public network address of the specified elasticsearch instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// When the instance is in the activating, invalid, or inactive state, its configuration cannot be updated.
//
// @param request - UpdatePublicNetworkRequest
//
// @return UpdatePublicNetworkResponse
func (client *Client) UpdatePublicNetwork(InstanceId *string, request *UpdatePublicNetworkRequest) (_result *UpdatePublicNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePublicNetworkResponse{}
	_body, _err := client.UpdatePublicNetworkWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ## RequestBody
//
// | Property | Type | Required | Example | Description |
//
// | -------- | ---- | -------- | ------- | ----------- |
//
// | publicIpWhitelist | List<String> | Yes | ["0.0.0.0/0","0.0.0.0/1"] | The list of IP address whitelists. This parameter is available if whiteIpGroup is left empty. The value of this parameter updates the IP address whitelist configurations in the Default whitelist group.
//
// You cannot configure both publicIpWhitelist and whiteIpGroup. |
//
// | whiteIpGroup | Object | No |  | You can update the whitelist configurations of an instance by using a whitelist group. You can update only one whitelist group.
//
// You cannot configure both publicIpWhitelist and whiteIpGroup. |
//
// | └ groupName | String | No | test_group_name | The group name of the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// | └ ips | List<String> | No | ["0.0.0.0", "10.2.XX.XX"] | The list of IP addresses in the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// > **Notice*	- The addition and deletion of whitelist groups are implemented by calling modifyMode to Cover. Delete and Append cannot add or delete whitelist groups at the same time. You can only modify the IP address list in the whitelist group. Take note of the following items: - If the modifyMode parameter is set to Cover, the whitelist group is deleted if ips is empty. If groupName is not in the list of existing whitelist group names, a whitelist group is created.
//
// - If the modifyMode parameter is set to Delete, you must retain at least one IP address for the deleted ips.
//
// - If the modifyMode parameter is set to Append, make sure that the whitelist group name has been created. Otherwise, the NotFound error message appears.
//
// Description:
//
// >  In the following example, only the parameters in the returned data list are guaranteed to be included. The parameters that are not mentioned are for reference only. For more information about the parameters, see [ListInstance](https://help.aliyun.com/document_detail/142230.html). You cannot force a dependency in a program to get these parameters.
//
// @param request - UpdatePublicWhiteIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicWhiteIpsResponse
func (client *Client) UpdatePublicWhiteIpsWithOptions(InstanceId *string, request *UpdatePublicWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicWhiteIpsResponse, _err error) {
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

	if !dara.IsNil(request.ModifyMode) {
		query["modifyMode"] = request.ModifyMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePublicWhiteIps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/public-white-ips"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePublicWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ## RequestBody
//
// | Property | Type | Required | Example | Description |
//
// | -------- | ---- | -------- | ------- | ----------- |
//
// | publicIpWhitelist | List<String> | Yes | ["0.0.0.0/0","0.0.0.0/1"] | The list of IP address whitelists. This parameter is available if whiteIpGroup is left empty. The value of this parameter updates the IP address whitelist configurations in the Default whitelist group.
//
// You cannot configure both publicIpWhitelist and whiteIpGroup. |
//
// | whiteIpGroup | Object | No |  | You can update the whitelist configurations of an instance by using a whitelist group. You can update only one whitelist group.
//
// You cannot configure both publicIpWhitelist and whiteIpGroup. |
//
// | └ groupName | String | No | test_group_name | The group name of the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// | └ ips | List<String> | No | ["0.0.0.0", "10.2.XX.XX"] | The list of IP addresses in the whitelist group. This parameter is required if the whiteIpGroup parameter is optional. |
//
// > **Notice*	- The addition and deletion of whitelist groups are implemented by calling modifyMode to Cover. Delete and Append cannot add or delete whitelist groups at the same time. You can only modify the IP address list in the whitelist group. Take note of the following items: - If the modifyMode parameter is set to Cover, the whitelist group is deleted if ips is empty. If groupName is not in the list of existing whitelist group names, a whitelist group is created.
//
// - If the modifyMode parameter is set to Delete, you must retain at least one IP address for the deleted ips.
//
// - If the modifyMode parameter is set to Append, make sure that the whitelist group name has been created. Otherwise, the NotFound error message appears.
//
// Description:
//
// >  In the following example, only the parameters in the returned data list are guaranteed to be included. The parameters that are not mentioned are for reference only. For more information about the parameters, see [ListInstance](https://help.aliyun.com/document_detail/142230.html). You cannot force a dependency in a program to get these parameters.
//
// @param request - UpdatePublicWhiteIpsRequest
//
// @return UpdatePublicWhiteIpsResponse
func (client *Client) UpdatePublicWhiteIps(InstanceId *string, request *UpdatePublicWhiteIpsRequest) (_result *UpdatePublicWhiteIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePublicWhiteIpsResponse{}
	_body, _err := client.UpdatePublicWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更改ES集群高可用策略
//
// @param request - UpdateReadWritePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReadWritePolicyResponse
func (client *Client) UpdateReadWritePolicyWithOptions(InstanceId *string, request *UpdateReadWritePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateReadWritePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateReadWritePolicy"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/update-read-write-policy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateReadWritePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改ES集群高可用策略
//
// @param request - UpdateReadWritePolicyRequest
//
// @return UpdateReadWritePolicyResponse
func (client *Client) UpdateReadWritePolicy(InstanceId *string, request *UpdateReadWritePolicyRequest) (_result *UpdateReadWritePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateReadWritePolicyResponse{}
	_body, _err := client.UpdateReadWritePolicyWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call UpdateSnapshotSetting to update the data backup configuration of the specified instance.
//
// @param request - UpdateSnapshotSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSettingWithOptions(InstanceId *string, request *UpdateSnapshotSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSnapshotSettingResponse, _err error) {
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
		Action:      dara.String("UpdateSnapshotSetting"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshot-setting"),
		Method:      dara.String("POST"),
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
// Call UpdateSnapshotSetting to update the data backup configuration of the specified instance.
//
// @param request - UpdateSnapshotSettingRequest
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSetting(InstanceId *string, request *UpdateSnapshotSettingRequest) (_result *UpdateSnapshotSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.UpdateSnapshotSettingWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the synonym dictionaries of an Elasticsearch cluster.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateSynonymsDictsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSynonymsDictsResponse
func (client *Client) UpdateSynonymsDictsWithOptions(InstanceId *string, request *UpdateSynonymsDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSynonymsDictsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSynonymsDicts"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/synonymsDict"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSynonymsDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the synonym dictionaries of an Elasticsearch cluster.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the dictionary file is stored in an Object Storage Service (OSS) bucket, you must make sure that the access control list (ACL) of the bucket is public read.
//
//   - If you do not set sourceType to ORIGIN for an uploaded dictionary file, the file will be deleted after you call this operation.
//
// @param request - UpdateSynonymsDictsRequest
//
// @return UpdateSynonymsDictsResponse
func (client *Client) UpdateSynonymsDicts(InstanceId *string, request *UpdateSynonymsDictsRequest) (_result *UpdateSynonymsDictsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSynonymsDictsResponse{}
	_body, _err := client.UpdateSynonymsDictsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(InstanceId *string, TemplateName *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/templates/" + dara.PercentEncode(dara.StringValue(TemplateName))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(InstanceId *string, TemplateName *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(InstanceId, TemplateName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// >  If you want to add an IP address whitelist, you can set the modifyMode parameter only to Cover. If you set this parameter to Delete or Append, you can only update an IP address whitelist.
//
//   - If you set the modifyMode parameter to Cover and leave the ips parameter empty, the system deletes the specified whitelist. If the whitelist specified by using the groupName parameter does not exist, the system creates such a whitelist.
//
//   - If you set the modifyMode parameter to Delete, at least one IP address must be retained for the specified whitelist.
//
//   - If you set the modifyMode parameter to Append, you must make sure that the specified whitelist exists. Otherwise, the system reports the NotFound error.
//
// Description:
//
// > For more information about the parameters displayed in the following sample code but not provided in the preceding tables, see [ListInstance](https://help.aliyun.com/document_detail/142230.html). You cannot force your program to obtain these parameters.
//
// @param request - UpdateWhiteIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWhiteIpsResponse
func (client *Client) UpdateWhiteIpsWithOptions(InstanceId *string, request *UpdateWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWhiteIpsResponse, _err error) {
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

	if !dara.IsNil(request.ModifyMode) {
		query["modifyMode"] = request.ModifyMode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EsIPWhitelist) {
		body["esIPWhitelist"] = request.EsIPWhitelist
	}

	if !dara.IsNil(request.WhiteIpGroup) {
		body["whiteIpGroup"] = request.WhiteIpGroup
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWhiteIps"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/white-ips"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// >  If you want to add an IP address whitelist, you can set the modifyMode parameter only to Cover. If you set this parameter to Delete or Append, you can only update an IP address whitelist.
//
//   - If you set the modifyMode parameter to Cover and leave the ips parameter empty, the system deletes the specified whitelist. If the whitelist specified by using the groupName parameter does not exist, the system creates such a whitelist.
//
//   - If you set the modifyMode parameter to Delete, at least one IP address must be retained for the specified whitelist.
//
//   - If you set the modifyMode parameter to Append, you must make sure that the specified whitelist exists. Otherwise, the system reports the NotFound error.
//
// Description:
//
// > For more information about the parameters displayed in the following sample code but not provided in the preceding tables, see [ListInstance](https://help.aliyun.com/document_detail/142230.html). You cannot force your program to obtain these parameters.
//
// @param request - UpdateWhiteIpsRequest
//
// @return UpdateWhiteIpsResponse
func (client *Client) UpdateWhiteIps(InstanceId *string, request *UpdateWhiteIpsRequest) (_result *UpdateWhiteIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWhiteIpsResponse{}
	_body, _err := client.UpdateWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Logstash实例的X-Pack监控报警配置。
//
// @param request - UpdateXpackMonitorConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateXpackMonitorConfigResponse
func (client *Client) UpdateXpackMonitorConfigWithOptions(InstanceId *string, request *UpdateXpackMonitorConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateXpackMonitorConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Endpoints) {
		body["endpoints"] = request.Endpoints
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateXpackMonitorConfig"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/xpack-monitor-config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateXpackMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Logstash实例的X-Pack监控报警配置。
//
// @param request - UpdateXpackMonitorConfigRequest
//
// @return UpdateXpackMonitorConfigResponse
func (client *Client) UpdateXpackMonitorConfig(InstanceId *string, request *UpdateXpackMonitorConfigRequest) (_result *UpdateXpackMonitorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateXpackMonitorConfigResponse{}
	_body, _err := client.UpdateXpackMonitorConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ES集群版本升级
//
// Description:
//
// 5A2CFF0E-5718-45B5-9D4D-70B3FF\\*\\*\\*\\*
//
// @param request - UpgradeEngineVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeEngineVersionResponse
func (client *Client) UpgradeEngineVersionWithOptions(InstanceId *string, request *UpgradeEngineVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeEngineVersionResponse, _err error) {
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

	if !dara.IsNil(request.UpdateStrategy) {
		query["updateStrategy"] = request.UpdateStrategy
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Plugins) {
		body["plugins"] = request.Plugins
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
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
		Action:      dara.String("UpgradeEngineVersion"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/upgrade-version"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeEngineVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ES集群版本升级
//
// Description:
//
// 5A2CFF0E-5718-45B5-9D4D-70B3FF\\*\\*\\*\\*
//
// @param request - UpgradeEngineVersionRequest
//
// @return UpgradeEngineVersionResponse
func (client *Client) UpgradeEngineVersion(InstanceId *string, request *UpgradeEngineVersionRequest) (_result *UpgradeEngineVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeEngineVersionResponse{}
	_body, _err := client.UpgradeEngineVersionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the connectivity between a Logstash cluster and its associated Elasticsearch cluster when you configure the X-Pack Monitoring feature for the Logstash cluster.
//
// Description:
//
// > Before you enable the X-Pack Monitoring feature for a Logstash cluster, you must associate the Logstash cluster with an Elasticsearch cluster. This way, you can view the monitoring data of the Logstash cluster in the Kibana console of the Elasticsearch cluster.
//
// @param request - ValidateConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateConnectionResponse
func (client *Client) ValidateConnectionWithOptions(InstanceId *string, request *ValidateConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateConnection"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/logstashes/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/validate-connection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the connectivity between a Logstash cluster and its associated Elasticsearch cluster when you configure the X-Pack Monitoring feature for the Logstash cluster.
//
// Description:
//
// > Before you enable the X-Pack Monitoring feature for a Logstash cluster, you must associate the Logstash cluster with an Elasticsearch cluster. This way, you can view the monitoring data of the Logstash cluster in the Kibana console of the Elasticsearch cluster.
//
// @param request - ValidateConnectionRequest
//
// @return ValidateConnectionResponse
func (client *Client) ValidateConnection(InstanceId *string, request *ValidateConnectionRequest) (_result *ValidateConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateConnectionResponse{}
	_body, _err := client.ValidateConnectionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验缩节点合法性
//
// @param request - ValidateShrinkNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateShrinkNodesResponse
func (client *Client) ValidateShrinkNodesWithOptions(InstanceId *string, request *ValidateShrinkNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateShrinkNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.IgnoreStatus) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !dara.IsNil(request.NodeType) {
		query["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateShrinkNodes"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/validate-shrink-nodes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateShrinkNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验缩节点合法性
//
// @param request - ValidateShrinkNodesRequest
//
// @return ValidateShrinkNodesResponse
func (client *Client) ValidateShrinkNodes(InstanceId *string, request *ValidateShrinkNodesRequest) (_result *ValidateShrinkNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateShrinkNodesResponse{}
	_body, _err := client.ValidateShrinkNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ValidateSlrPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateSlrPermissionResponse
func (client *Client) ValidateSlrPermissionWithOptions(request *ValidateSlrPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateSlrPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Rolename) {
		query["rolename"] = request.Rolename
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateSlrPermission"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/user/servicerolepermission"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateSlrPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ValidateSlrPermissionRequest
//
// @return ValidateSlrPermissionResponse
func (client *Client) ValidateSlrPermission(request *ValidateSlrPermissionRequest) (_result *ValidateSlrPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateSlrPermissionResponse{}
	_body, _err := client.ValidateSlrPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 缩节点校验数据迁移合法性
//
// @param request - ValidateTransferableNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTransferableNodesResponse
func (client *Client) ValidateTransferableNodesWithOptions(InstanceId *string, request *ValidateTransferableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateTransferableNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeType) {
		query["nodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateTransferableNodes"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/validate-transfer-nodes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateTransferableNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 缩节点校验数据迁移合法性
//
// @param request - ValidateTransferableNodesRequest
//
// @return ValidateTransferableNodesResponse
func (client *Client) ValidateTransferableNodes(InstanceId *string, request *ValidateTransferableNodesRequest) (_result *ValidateTransferableNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateTransferableNodesResponse{}
	_body, _err := client.ValidateTransferableNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The configurations of dedicated master nodes.
//
// Description:
//
// The configurations of warm nodes.
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientNodeConfiguration) {
		body["clientNodeConfiguration"] = request.ClientNodeConfiguration
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ElasticDataNodeConfiguration) {
		body["elasticDataNodeConfiguration"] = request.ElasticDataNodeConfiguration
	}

	if !dara.IsNil(request.EsAdminPassword) {
		body["esAdminPassword"] = request.EsAdminPassword
	}

	if !dara.IsNil(request.EsVersion) {
		body["esVersion"] = request.EsVersion
	}

	if !dara.IsNil(request.InstanceCategory) {
		body["instanceCategory"] = request.InstanceCategory
	}

	if !dara.IsNil(request.KibanaConfiguration) {
		body["kibanaConfiguration"] = request.KibanaConfiguration
	}

	if !dara.IsNil(request.MasterConfiguration) {
		body["masterConfiguration"] = request.MasterConfiguration
	}

	if !dara.IsNil(request.NetworkConfig) {
		body["networkConfig"] = request.NetworkConfig
	}

	if !dara.IsNil(request.NodeAmount) {
		body["nodeAmount"] = request.NodeAmount
	}

	if !dara.IsNil(request.NodeSpec) {
		body["nodeSpec"] = request.NodeSpec
	}

	if !dara.IsNil(request.PaymentInfo) {
		body["paymentInfo"] = request.PaymentInfo
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.WarmNodeConfiguration) {
		body["warmNodeConfiguration"] = request.WarmNodeConfiguration
	}

	if !dara.IsNil(request.ZoneCount) {
		body["zoneCount"] = request.ZoneCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("createInstance"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The configurations of dedicated master nodes.
//
// Description:
//
// The configurations of warm nodes.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
