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
	client.EndpointMap = map[string]*string{
		"us-west-1":             dara.String("elasticsearch.us-west-1.aliyuncs.com"),
		"us-east-1":             dara.String("elasticsearch.us-east-1.aliyuncs.com"),
		"eu-west-1":             dara.String("elasticsearch.eu-west-1.aliyuncs.com"),
		"eu-central-1":          dara.String("elasticsearch.eu-central-1.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("elasticsearch.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("elasticsearch.cn-wulanchabu.aliyuncs.com"),
		"cn-shenzhen":           dara.String("elasticsearch.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("elasticsearch.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":           dara.String("elasticsearch.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("elasticsearch.cn-qingdao.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("elasticsearch.cn-north-2-gov-1.aliyuncs.com"),
		"cn-hongkong":           dara.String("elasticsearch.cn-hongkong.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("elasticsearch.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hangzhou":           dara.String("elasticsearch.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":          dara.String("elasticsearch.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("elasticsearch.cn-chengdu.aliyuncs.com"),
		"cn-beijing":            dara.String("elasticsearch.cn-beijing.aliyuncs.com"),
		"ap-southeast-5":        dara.String("elasticsearch.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3":        dara.String("elasticsearch.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-2":        dara.String("elasticsearch.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-1":        dara.String("elasticsearch.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("elasticsearch.ap-south-1.aliyuncs.com"),
		"ap-northeast-1":        dara.String("elasticsearch.ap-northeast-1.aliyuncs.com"),
	}
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
// Resumes an offline zone. This operation is valid only for multi-zone instances.
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
// Resumes an offline zone. This operation is valid only for multi-zone instances.
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
// Configures network connectivity to establish a connection between different instances.
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
// Configures network connectivity to establish a connection between different instances.
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
// Creates a reference repository when setting up a cross-cluster OSS repository.
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
// Creates a reference repository when setting up a cross-cluster OSS repository.
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
// Recovers a frozen Elasticsearch instance that was released.
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
// Recovers a frozen Elasticsearch instance that was released.
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
// Resumes a frozen Logstash instance that was frozen after release.
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
// Resumes a frozen Logstash instance that was frozen after release.
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

// Summary:
//
// Cancels a running data migration task.
//
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

// Summary:
//
// Cancels a running data migration task.
//
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
// Recommends optimal cluster capacity planning configurations based on business scenarios, QPS, and log generation volume.
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
// Recommends optimal cluster capacity planning configurations based on business scenarios, QPS, and log generation volume.
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
// Disables the intelligent O&M feature for an instance.
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
// Disables the intelligent O&M feature for an instance.
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

// Summary:
//
// Disables the HTTPS protocol for a cluster.
//
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

// Summary:
//
// Disables the HTTPS protocol for a cluster.
//
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
// Disables the cloud managed feature for a specified index in an Indexing Service cluster. This operation is irreversible. After the feature is disabled, it cannot be enabled again.
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
// Disables the cloud managed feature for a specified index in an Indexing Service cluster. This operation is irreversible. After the feature is disabled, it cannot be enabled again.
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
// Creates a collector to collect data from a specified service.
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
// Creates a collector to collect data from a specified service.
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
// Creates an Elasticsearch composable template.
//
// Description:
//
// For more information, see [Store large volumes of data by using OpenStore](https://help.aliyun.com/document_detail/317694.html).
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
// Creates an Elasticsearch composable template.
//
// Description:
//
// For more information, see [Store large volumes of data by using OpenStore](https://help.aliyun.com/document_detail/317694.html).
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
// Creates a data stream to manage a set of indexes.
//
// Description:
//
// > The data stream name you create must have a one-to-one correspondence with the index pattern in the index template, and the index template must have the data stream feature enabled. For example, if the index pattern in the index template is ds-\\*, the corresponding data stream name should be ds-.
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
// Creates a data stream to manage a set of indexes.
//
// Description:
//
// > The data stream name you create must have a one-to-one correspondence with the index pattern in the index template, and the index template must have the data stream feature enabled. For example, if the index pattern in the index template is ds-\\*, the corresponding data stream name should be ds-.
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
// Creates an index lifecycle policy. If a policy with the specified name already exists, the existing policy is replaced and its version is incremented.
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
// Creates an index lifecycle policy. If a policy with the specified name already exists, the existing policy is replaced and its version is incremented.
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
// Creates a cluster index template that can be used for component-based index template settings.
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
// Creates a cluster index template that can be used for component-based index template settings.
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
// Creates a Logstash instance by calling CreateLogstash.
//
// Description:
//
// Before calling this operation, note the following:
//
// - Make sure that you are familiar with the billing method and pricing of Logstash. <props="china"><ph>For more information, see [Billing](https://help.aliyun.com/document_detail/260882.html).</ph>
//
// - To create an instance, complete real-name verification. <props="china"><ph>For more information, see [Real-name verification](https://help.aliyun.com/document_detail/37175.html).</ph>.
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
// Creates a Logstash instance by calling CreateLogstash.
//
// Description:
//
// Before calling this operation, note the following:
//
// - Make sure that you are familiar with the billing method and pricing of Logstash. <props="china"><ph>For more information, see [Billing](https://help.aliyun.com/document_detail/260882.html).</ph>
//
// - To create an instance, complete real-name verification. <props="china"><ph>For more information, see [Real-name verification](https://help.aliyun.com/document_detail/37175.html).</ph>.
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
// Creates a Logstash pipeline to collect data.
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
// Creates a Logstash pipeline to collect data.
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

// Summary:
//
// Calls CreateSnapshot to manually create a snapshot backup of a cluster.
//
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

// Summary:
//
// Calls CreateSnapshot to manually create a snapshot backup of a cluster.
//
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
// Creates a PrivateLink VPC endpoint to connect to an endpoint service created in a user VPC.
//
// Description:
//
// For more information about this API operation, see [Configure private connectivity for an instance](https://help.aliyun.com/document_detail/279559.html).
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
// Creates a PrivateLink VPC endpoint to connect to an endpoint service created in a user VPC.
//
// Description:
//
// For more information about this API operation, see [Configure private connectivity for an instance](https://help.aliyun.com/document_detail/279559.html).
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
// Takes part of the zones offline when multiple zones are available, and migrates the nodes in the offline zones to other zones.
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
// Takes part of the zones offline when multiple zones are available, and migrates the nodes in the offline zones to other zones.
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
// Deletes a specified collector.
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
// Deletes a specified collector.
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
// Deletes a component index template of Elasticsearch.
//
// Description:
//
// For more information, see [Store massive amounts of data by using OpenStore](https://help.aliyun.com/document_detail/317694.html).
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
// Deletes a component index template of Elasticsearch.
//
// Description:
//
// For more information, see [Store massive amounts of data by using OpenStore](https://help.aliyun.com/document_detail/317694.html).
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

// Summary:
//
// Deletes the network connectivity between two instances.
//
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

// Summary:
//
// Deletes the network connectivity between two instances.
//
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
// Deletes a specified cluster data stream.
//
// Description:
//
// > - Deleting a data stream also deletes its backing indexes. Proceed with caution.- When an index template has associated data streams, you must delete the data streams associated with the index template before you can delete the index template. On the data stream list page, view the data stream details to find the index template that matches the data stream.
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
// Deletes a specified cluster data stream.
//
// Description:
//
// > - Deleting a data stream also deletes its backing indexes. Proceed with caution.- When an index template has associated data streams, you must delete the data streams associated with the index template before you can delete the index template. On the data stream list page, view the data stream details to find the index template that matches the data stream.
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

// Summary:
//
// Deletes an Elasticsearch index migration task.
//
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

// Summary:
//
// Deletes an Elasticsearch index migration task.
//
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
// Deletes a historical index template.
//
// Description:
//
// For more information, see [Store massive amounts of data through OpenStore](https://help.aliyun.com/document_detail/317694.html).
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
// Deletes a historical index template.
//
// Description:
//
// For more information, see [Store massive amounts of data through OpenStore](https://help.aliyun.com/document_detail/317694.html).
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

// Summary:
//
// Deletes a specified index lifecycle policy.
//
// Description:
//
// > You cannot delete a policy that is currently in use. If the policy is being used to manage any index, the request fails and returns an error.
//
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

// Summary:
//
// Deletes a specified index lifecycle policy.
//
// Description:
//
// > You cannot delete a policy that is currently in use. If the policy is being used to manage any index, the request fails and returns an error.
//
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
// Deletes a specified index template.
//
// Description:
//
// > Before deleting an index template, delete the data streams associated with the index template. Otherwise, the index template cannot be deleted.
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
// Deletes a specified index template.
//
// Description:
//
// > Before deleting an index template, delete the data streams associated with the index template. Otherwise, the index template cannot be deleted.
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

// Summary:
//
// All physical resources used by the instance are reclaimed, all related data is permanently lost and cannot be recovered, and the cloud disks mounted to the instance nodes along with their corresponding snapshots are released.
//
// Description:
//
// Before you invoke this operation, note the following:
//
// Data cannot be recovered after the instance is released. Back up your data before releasing the instance. For more information, see [Snapshot backup and recovery commands](https://help.aliyun.com/document_detail/65675.html).
//
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

// Summary:
//
// All physical resources used by the instance are reclaimed, all related data is permanently lost and cannot be recovered, and the cloud disks mounted to the instance nodes along with their corresponding snapshots are released.
//
// Description:
//
// Before you invoke this operation, note the following:
//
// Data cannot be recovered after the instance is released. Back up your data before releasing the instance. For more information, see [Snapshot backup and recovery commands](https://help.aliyun.com/document_detail/65675.html).
//
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
// Proactively releases a Logstash instance.
//
// Description:
//
// Before calling this operation, note the following:
//
// After the instance is released, all physical resources used by the instance are reclaimed, all related data is permanently lost and cannot be recovered, cloud disks mounted to the instance nodes are also released, and the corresponding snapshots are deleted.
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
// Proactively releases a Logstash instance.
//
// Description:
//
// Before calling this operation, note the following:
//
// After the instance is released, all physical resources used by the instance are reclaimed, all related data is permanently lost and cannot be recovered, cloud disks mounted to the instance nodes are also released, and the corresponding snapshots are deleted.
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
// Deletes pipelines configured for a Logstash instance.
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
// Deletes pipelines configured for a Logstash instance.
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

// Summary:
//
// Deletes a cross-cluster OSS reference repository from an instance.
//
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

// Summary:
//
// Deletes a cross-cluster OSS reference repository from an instance.
//
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
// Calls DeleteVpcEndpoint to delete a VPC endpoint under a service account.
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
// Calls DeleteVpcEndpoint to delete a VPC endpoint under a service account.
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
// Calls the DescribeAckOperator operation to query the Elasticsearch Operator information installed on a specified Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > Before installing a collector on an ACK cluster, you can call this operation to check the installation status of the Elasticsearch Operator on the target cluster.
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
// Calls the DescribeAckOperator operation to query the Elasticsearch Operator information installed on a specified Container Service for Kubernetes (ACK) cluster.
//
// Description:
//
// > Before installing a collector on an ACK cluster, you can call this operation to check the installation status of the Elasticsearch Operator on the target cluster.
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
// Retrieves the details of a collector instance.
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
// Retrieves the details of a collector instance.
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
// Queries the details of a composable index template in Elasticsearch.
//
// Description:
//
// For more information, see [Use OpenStore to store massive amounts of data](https://help.aliyun.com/document_detail/317694.html).
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
// Queries the details of a composable index template in Elasticsearch.
//
// Description:
//
// For more information, see [Use OpenStore to store massive amounts of data](https://help.aliyun.com/document_detail/317694.html).
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

// Summary:
//
// Retrieves a list of instances that can establish private network peering with the current instance. Instances that are already connected are not included.
//
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

// Summary:
//
// Retrieves a list of instances that can establish private network peering with the current instance. Instances that are already connected are not included.
//
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
// Queries the details of a historical index template.
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
// Queries the details of a historical index template.
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

// Summary:
//
// Calls the DescribeDiagnoseReport operation to view historical reports of intelligent O&M.
//
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

// Summary:
//
// Calls the DescribeDiagnoseReport operation to view historical reports of intelligent O&M.
//
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

// Summary:
//
// Calls the DescribeDiagnosisSettings operation to obtain the scenario settings of intelligent O&M.
//
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

// Summary:
//
// Calls the DescribeDiagnosisSettings operation to obtain the scenario settings of intelligent O&M.
//
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
// Retrieves dynamic metrics of a cluster.
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
// Retrieves dynamic metrics of a cluster.
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
// Queries the health status of a cluster to check whether it is running properly.
//
// Description:
//
// The instance health status. The following three states are supported:
//
// - GREEN: Primary and replica shards are allocated properly.
//
// - YELLOW: Primary shards are allocated properly, but replica shards are not allocated properly.
//
// - RED: Primary shards are not allocated properly.
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
// Queries the health status of a cluster to check whether it is running properly.
//
// Description:
//
// The instance health status. The following three states are supported:
//
// - GREEN: Primary and replica shards are allocated properly.
//
// - YELLOW: Primary shards are allocated properly, but replica shards are not allocated properly.
//
// - RED: Primary shards are not allocated properly.
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

// Summary:
//
// Queries the details of a specified index lifecycle policy.
//
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

// Summary:
//
// Queries the details of a specified index lifecycle policy.
//
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

// Summary:
//
// Returns information about an index template.
//
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

// Summary:
//
// Returns information about an index template.
//
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
// Queries the details of a specified instance.
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
// Queries the details of a specified instance.
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
// Retrieves the Kibana node configuration of an Elasticsearch instance.
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
// Retrieves the Kibana node configuration of an Elasticsearch instance.
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
// Queries the details of a Logstash instance.
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
// Queries the details of a Logstash instance.
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

// Summary:
//
// Retrieves the pipeline information of a Logstash instance.
//
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

// Summary:
//
// Retrieves the pipeline information of a Logstash instance.
//
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
// Calls DescribePipelineManagementConfig to retrieve the pipeline management configuration of a Logstash instance.
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
// Calls DescribePipelineManagementConfig to retrieve the pipeline management configuration of a Logstash instance.
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

// Summary:
//
// Retrieves the region information of Alibaba Cloud Elasticsearch.
//
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

// Summary:
//
// Retrieves the region information of Alibaba Cloud Elasticsearch.
//
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
// Retrieves the snapshot backup settings and backup cycle of a cluster.
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
// Retrieves the snapshot backup settings and backup cycle of a cluster.
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

// Summary:
//
// Retrieves the scenario-specific template configuration and cluster settings of an instance.
//
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

// Summary:
//
// Retrieves the scenario-specific template configuration and cluster settings of an instance.
//
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
// Retrieves the X-Pack monitoring configuration of a Logstash instance.
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
// Retrieves the X-Pack monitoring configuration of a Logstash instance.
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
// Calls DiagnoseInstance to immediately diagnose an instance.
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
// Calls DiagnoseInstance to immediately diagnose an instance.
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
// Disables Kibana private network access.
//
// Description:
//
// This API operation supports only cloud-native instances. For legacy architecture instances, use the TriggerNetwork method.
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
// Disables Kibana private network access.
//
// Description:
//
// This API operation supports only cloud-native instances. For legacy architecture instances, use the TriggerNetwork method.
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
// Invokes the EnableKibanaPvlNetwork operation to enable private network access for Kibana.
//
// Description:
//
// 1. This API operation is supported only for cloud-native instances. For legacy architecture instances, use the TriggerNetwork method.
//
// 2. The Kibana specification must be greater than 1 vCPU and 2 GB of memory.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
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
		Query:   openapiutil.Query(query),
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
// Invokes the EnableKibanaPvlNetwork operation to enable private network access for Kibana.
//
// Description:
//
// 1. This API operation is supported only for cloud-native instances. For legacy architecture instances, use the TriggerNetwork method.
//
// 2. The Kibana specification must be greater than 1 vCPU and 2 GB of memory.
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
// Retrieves the estimated restart time of a Logstash instance.
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
// Retrieves the estimated restart time of a Logstash instance.
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
// Retrieves the estimated restart time for an instance.
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
// Retrieves the estimated restart time for an instance.
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
// Retrieves index migration data information.
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
// Retrieves index migration data information.
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

// Summary:
//
// Retrieves the elastic scaling rules of a cluster. Elastic nodes must be purchased when the instance is created.
//
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

// Summary:
//
// Retrieves the elastic scaling rules of a cluster. Elastic nodes must be purchased when the instance is created.
//
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
// Calls GetEmonGrafanaAlerts to retrieve the Grafana alert list.
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
// Calls GetEmonGrafanaAlerts to retrieve the Grafana alert list.
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
// Calls GetEmonGrafanaDashboards to retrieve the list of Grafana dashboards.
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
// Calls GetEmonGrafanaDashboards to retrieve the list of Grafana dashboards.
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
// Queries the Grafana metric monitoring data of an Elasticsearch instance.
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
// Queries the Grafana metric monitoring data of an Elasticsearch instance.
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
// # Retrieve keystore information
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeystoresResponse
func (client *Client) GetKeystoresWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKeystoresResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKeystores"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/keystores"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKeystoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve keystore information
//
// @return GetKeystoresResponse
func (client *Client) GetKeystores(InstanceId *string) (_result *GetKeystoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKeystoresResponse{}
	_body, _err := client.GetKeystoresWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage capacity and usage of an OpenStore instance.
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
// Queries the storage capacity and usage of an OpenStore instance.
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
// Retrieves the current region information.
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
// Retrieves the current region information.
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
// Retrieves the nodes that can be removed based on the specified node type and quantity.
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
// Retrieves the nodes that can be removed based on the specified node type and quantity.
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
// Retrieves the nodes available for data migration based on the specified node type and count.
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
// Retrieves the nodes available for data migration based on the specified node type and count.
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
// Continue restarting the remaining edge zones of the Elasticsearch instance after the phased release is completed.
//
// @param request - GrayPublishRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrayPublishResponse
func (client *Client) GrayPublishWithOptions(InstanceId *string, request *GrayPublishRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrayPublishResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.XRequestChangeId) {
		query["X-Request-ChangeId"] = request.XRequestChangeId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrayPublish"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/grayPublish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GrayPublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Continue restarting the remaining edge zones of the Elasticsearch instance after the phased release is completed.
//
// @param request - GrayPublishRequest
//
// @return GrayPublishResponse
func (client *Client) GrayPublish(InstanceId *string, request *GrayPublishRequest) (_result *GrayPublishResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrayPublishResponse{}
	_body, _err := client.GrayPublishWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Initialize AI model
//
// @param request - InitModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitModelResponse
func (client *Client) InitModelWithOptions(InstanceId *string, request *InitModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["api_key"] = request.ApiKey
	}

	if !dara.IsNil(request.Host) {
		body["host"] = request.Host
	}

	if !dara.IsNil(request.HttpSchema) {
		body["http_schema"] = request.HttpSchema
	}

	if !dara.IsNil(request.Models) {
		body["models"] = request.Models
	}

	if !dara.IsNil(request.Workspace) {
		body["workspace"] = request.Workspace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitModel"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/initModel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InitModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Initialize AI model
//
// @param request - InitModelRequest
//
// @return InitModelResponse
func (client *Client) InitModel(InstanceId *string, request *InitModelRequest) (_result *InitModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitModelResponse{}
	_body, _err := client.InitModelWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the InitializeOperationRole operation to create a service-linked role.
//
// Description:
//
// > Before you use a collector to collect logs from different data sources or perform elastic scaling tasks for a cluster (applicable only to the China site), you must create a service-linked role.
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
// Calls the InitializeOperationRole operation to create a service-linked role.
//
// Description:
//
// > Before you use a collector to collect logs from different data sources or perform elastic scaling tasks for a cluster (applicable only to the China site), you must create a service-linked role.
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
// Installs the ACK Operator on a specified Container Service cluster.
//
// Description:
//
// > Before installing a collector on an ACK cluster, call this operation to install the Elasticsearch Operator on the target cluster.
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
// Installs the ACK Operator on a specified Container Service cluster.
//
// Description:
//
// > Before installing a collector on an ACK cluster, call this operation to install the Elasticsearch Operator on the target cluster.
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
// Installs preset plug-ins for Kibana. The Kibana instance must have specifications of 2 vCPUs and 4 GB of memory or higher.
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
// Installs preset plug-ins for Kibana. The Kibana instance must have specifications of 2 vCPUs and 4 GB of memory or higher.
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
// Installs system plugins for a specified Logstash instance.
//
// Description:
//
// Before calling this operation, note the following:
//
// The plugins to be installed must be included in the Alibaba Cloud Logstash [default system plugin list](https://help.aliyun.com/document_detail/139626.html). External open source plugins are not supported.
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
// Installs system plugins for a specified Logstash instance.
//
// Description:
//
// Before calling this operation, note the following:
//
// The plugins to be installed must be included in the Alibaba Cloud Logstash [default system plugin list](https://help.aliyun.com/document_detail/139626.html). External open source plugins are not supported.
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
// Installs system plug-ins on an Elasticsearch instance.
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
// Installs system plug-ins on an Elasticsearch instance.
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
// Installs custom plugins that have been uploaded to the Elasticsearch console.
//
// Description:
//
// > The custom plugin installation feature is being upgraded internally and is temporarily unavailable. If you urgently need this feature, submit a ticket to contact us.
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
// Installs custom plugins that have been uploaded to the Elasticsearch console.
//
// Description:
//
// > The custom plugin installation feature is being upgraded internally and is temporarily unavailable. If you urgently need this feature, submit a ticket to contact us.
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

// Summary:
//
// Interrupts an instance change task. This operation is valid only for instances in the Effecting state. After the interruption, the instance enters the suspended state.
//
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

// Summary:
//
// Interrupts an instance change task. This operation is valid only for instances in the Effecting state. After the interruption, the instance enters the suspended state.
//
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
// After the interruption, the instance enters the suspended state.
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
// After the interruption, the instance enters the suspended state.
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
// Retrieves the list of Container Service for Kubernetes (ACK) clusters.
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
// Retrieves the list of Container Service for Kubernetes (ACK) clusters.
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
// This operation is deprecated and will be taken offline soon.
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
// This operation is deprecated and will be taken offline soon.
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
// Release notes Release notes details.
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
// Release notes Release notes details.
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
// Retrieves information about all nodes in an Elasticsearch cluster.
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
// Retrieves information about all nodes in an Elasticsearch cluster.
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
// Retrieves the OSS reference repositories that can be added to the current instance.
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
// Retrieves the OSS reference repositories that can be added to the current instance.
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
// Retrieves a list of available Elasticsearch instances when configuring X-Pack monitoring for a Logstash instance.
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
// Retrieves a list of available Elasticsearch instances when configuring X-Pack monitoring for a Logstash instance.
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
// Retrieves a list of collectors.
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
// Retrieves a list of collectors.
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
// Retrieves the list of composable templates for an Elasticsearch instance.
//
// Description:
//
// For more information, see [Store massive amounts of data through OpenStore](https://help.aliyun.com/document_detail/317694.html).
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
// Retrieves the list of composable templates for an Elasticsearch instance.
//
// Description:
//
// For more information, see [Store massive amounts of data through OpenStore](https://help.aliyun.com/document_detail/317694.html).
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
// Retrieves a list of instances that have established private network peering with the current instance.
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
// Retrieves a list of instances that have established private network peering with the current instance.
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
// Retrieves the list of index data streams in an Elasticsearch cluster.
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
// Retrieves the list of index data streams in an Elasticsearch cluster.
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

// Summary:
//
// Retrieves a list of data migration tasks between different Elasticsearch clusters.
//
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

// Summary:
//
// Retrieves a list of data migration tasks between different Elasticsearch clusters.
//
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
// Invokes the ListDefaultCollectorConfigurations operation to retrieve the default configuration file of a collector.
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
// Invokes the ListDefaultCollectorConfigurations operation to retrieve the default configuration file of a collector.
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
// Queries the list of historical index templates.
//
// Description:
//
// For more information, see [Use OpenStore to store large volumes of data](https://help.aliyun.com/document_detail/317694.html).
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
// Queries the list of historical index templates.
//
// Description:
//
// For more information, see [Use OpenStore to store large volumes of data](https://help.aliyun.com/document_detail/317694.html).
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
// Retrieves the diagnostic indexes from the intelligent O&M module for a specified instance.
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
// Retrieves the diagnostic indexes from the intelligent O&M module for a specified instance.
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
// Calls ListDiagnoseReport to retrieve historical reports of intelligent O&M.
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
// Calls ListDiagnoseReport to retrieve historical reports of intelligent O&M.
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
// Retrieves all IDs of Intelligent O&M Center historical reports.
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
// Retrieves all IDs of Intelligent O&M Center historical reports.
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
// Lists the intelligent diagnostic items for an Elasticsearch instance.
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
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

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
// Lists the intelligent diagnostic items for an Elasticsearch instance.
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

// Summary:
//
// Queries the information of a specified dictionary.
//
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

// Summary:
//
// Queries the information of a specified dictionary.
//
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
// Queries the details of the dictionary list for a specified type.
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
// Queries the details of the dictionary list for a specified type.
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
// Retrieves the list of ECS instances under the current user\\"s VPC and vSwitch.
//
// Description:
//
//	Notice:  Before calling this operation, create the AliyunElasticsearchAccessingOOSRole and AliyunOOSAccessingECS4ESRole service-linked roles. These roles allow the Elasticsearch service account to obtain ECS access permissions of the Alibaba Cloud account. For more information, see [Collect ECS service logs](https://help.aliyun.com/document_detail/146446.html).
//
// .
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
// Retrieves the list of ECS instances under the current user\\"s VPC and vSwitch.
//
// Description:
//
//	Notice:  Before calling this operation, create the AliyunElasticsearchAccessingOOSRole and AliyunOOSAccessingECS4ESRole service-linked roles. These roles allow the Elasticsearch service account to obtain ECS access permissions of the Alibaba Cloud account. For more information, see [Collect ECS service logs](https://help.aliyun.com/document_detail/146446.html).
//
// .
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
// # Obtain Event List
//
// @param request - ListEventRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventRecordsResponse
func (client *Client) ListEventRecordsWithOptions(eventType *string, request *ListEventRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEventRecordsResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.TermContent) {
		query["termContent"] = request.TermContent
	}

	if !dara.IsNil(request.TermType) {
		query["termType"] = request.TermType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEventRecords"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/" + dara.PercentEncode(dara.StringValue(eventType)) + "/listEventRecords"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain Event List
//
// @param request - ListEventRecordsRequest
//
// @return ListEventRecordsResponse
func (client *Client) ListEventRecords(eventType *string, request *ListEventRecordsRequest) (_result *ListEventRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEventRecordsResponse{}
	_body, _err := client.ListEventRecordsWithOptions(eventType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the extension file configuration of a Logstash instance.
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
// Retrieves the extension file configuration of a Logstash instance.
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

// Summary:
//
// Queries the list of index lifecycle policies that have been created for a cluster.
//
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

// Summary:
//
// Queries the list of index lifecycle policies that have been created for a cluster.
//
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

// Summary:
//
// Queries a list of index templates.
//
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

// Summary:
//
// Queries a list of index templates.
//
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
// Queries the information about Elasticsearch instances.
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
// Queries the information about Elasticsearch instances.
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
// Queries the list of hardware O&M events triggered by an Elasticsearch cluster.
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
// Queries the list of hardware O&M events triggered by an Elasticsearch cluster.
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
// Filters system indexes from the index list of a cluster.
//
// Description:
//
// The ListInstanceIndices operation is applicable only to Elasticsearch instances that have the indexing service enabled. Query index information by using the Elasticsearch API. For more information, see [cat indices API
//
// ](https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html).
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
// Filters system indexes from the index list of a cluster.
//
// Description:
//
// The ListInstanceIndices operation is applicable only to Elasticsearch instances that have the indexing service enabled. Query index information by using the Elasticsearch API. For more information, see [cat indices API
//
// ](https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html).
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
// Retrieves the list of plugins installed on the Kibana node of an Elasticsearch instance.
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
// Retrieves the list of plugins installed on the Kibana node of an Elasticsearch instance.
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
// Queries the details of the Kibana private network connection.
//
// Description:
//
// This API operation supports only cloud-native instances.
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
// Queries the details of the Kibana private network connection.
//
// Description:
//
// This API operation supports only cloud-native instances.
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
// Displays the details of all or specified Logstash instances in a list.
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
// Displays the details of all or specified Logstash instances in a list.
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
// Queries the logs of a Logstash instance.
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
// Queries the logs of a Logstash instance.
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
// Calls ListLogstashPlugins to retrieve detailed information about all or specified plugins.
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
// Calls ListLogstashPlugins to retrieve detailed information about all or specified plugins.
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
// Historical report list of intelligent O&M.
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
// Historical report list of intelligent O&M.
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
// Retrieves the pipeline list of a Logstash instance.
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
// Retrieves the pipeline list of a Logstash instance.
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
// Retrieves the list of pipeline IDs for a Logstash instance.
//
// Description:
//
// > Pipeline management is divided into configuration file management and Kibana pipeline management. Kibana pipeline management is not available in the console for some regions.
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
// Retrieves the list of pipeline IDs for a Logstash instance.
//
// Description:
//
// > Pipeline management is divided into configuration file management and Kibana pipeline management. Kibana pipeline management is not available in the console for some regions.
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
// Retrieves the plugin list of a specified Alibaba Cloud Elasticsearch instance.
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
// Retrieves the plugin list of a specified Alibaba Cloud Elasticsearch instance.
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
// Queries logs of different types for an Elasticsearch instance.
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
// Queries logs of different types for an Elasticsearch instance.
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
// Queries the data progress list of ongoing and completed shard recoveries. By default, only ongoing shard recovery information is returned.
//
// Description:
//
// > Shard recovery is the process of synchronizing data from a primary shard to a replica shard. After recovery is complete, the replica shard becomes available for search.
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
// Queries the data progress list of ongoing and completed shard recoveries. By default, only ongoing shard recovery information is returned.
//
// Description:
//
// > Shard recovery is the process of synchronizing data from a primary shard to a replica shard. After recovery is complete, the replica shard becomes available for search.
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
// Retrieves the list of cross-cluster OSS repository settings for the current instance.
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
// Retrieves the list of cross-cluster OSS repository settings for the current instance.
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
// # Statistics of management event records
//
// @param request - ListStatsEventRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStatsEventRecordsResponse
func (client *Client) ListStatsEventRecordsWithOptions(request *ListStatsEventRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStatsEventRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventType) {
		query["eventType"] = request.EventType
	}

	if !dara.IsNil(request.Level) {
		query["level"] = request.Level
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStatsEventRecords"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/event/statsEventRecords"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStatsEventRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Statistics of management event records
//
// @param request - ListStatsEventRecordsRequest
//
// @return ListStatsEventRecordsResponse
func (client *Client) ListStatsEventRecords(request *ListStatsEventRecordsRequest) (_result *ListStatsEventRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStatsEventRecordsResponse{}
	_body, _err := client.ListStatsEventRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the relationships between all instances and tags.
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
// Retrieves the relationships between all instances and tags.
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
// Queries all labels created by the user in the current region.
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
// Queries all labels created by the user in the current region.
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
// # Custom plugin list
//
// @param request - ListUserPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPluginResponse
func (client *Client) ListUserPluginWithOptions(instanceId *string, request *ListUserPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserPluginResponse, _err error) {
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
		Action:      dara.String("ListUserPlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/userPlugins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Custom plugin list
//
// @param request - ListUserPluginRequest
//
// @return ListUserPluginResponse
func (client *Client) ListUserPlugin(instanceId *string, request *ListUserPluginRequest) (_result *ListUserPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserPluginResponse{}
	_body, _err := client.ListUserPluginWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of endpoints in the VPC of a service account.
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
// Queries the status of endpoints in the VPC of a service account.
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
// Migrates nodes in a specified zone to a destination zone.
//
// Description:
//
// When you upgrade the specifications of an instance and encounter insufficient inventory in the current zone, you can resolve this issue by migrating zone nodes. Before calling this operation, make sure that:
//
// - Your account has a zone with sufficient resources.
//
//	After migrating nodes of the current specifications to another zone, you must manually [upgrade the cluster](https://help.aliyun.com/document_detail/96650.html). The cluster is not upgraded during the migration. Therefore, select a zone with sufficient resources to avoid cluster upgrade failures. Select a zone with a later alphabetical order first. For example, between ap-southeast-1e and ap-southeast-1h, select ap-southeast-1h first.
//
// - The cluster is in a healthy state.
//
//	You can run the `GET _cat/health?v` command to check the cluster health status.
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
// Migrates nodes in a specified zone to a destination zone.
//
// Description:
//
// When you upgrade the specifications of an instance and encounter insufficient inventory in the current zone, you can resolve this issue by migrating zone nodes. Before calling this operation, make sure that:
//
// - Your account has a zone with sufficient resources.
//
//	After migrating nodes of the current specifications to another zone, you must manually [upgrade the cluster](https://help.aliyun.com/document_detail/96650.html). The cluster is not upgraded during the migration. Therefore, select a zone with sufficient resources to avoid cluster upgrade failures. Select a zone with a later alphabetical order first. For example, between ap-southeast-1e and ap-southeast-1h, select ap-southeast-1h first.
//
// - The cluster is in a healthy state.
//
//	You can run the `GET _cat/health?v` command to check the cluster health status.
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
// Updates the ECS instances on which a collector is installed.
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
// Updates the ECS instances on which a collector is installed.
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

// Summary:
//
// Updates the elastic scaling rules of a cluster.
//
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

// Summary:
//
// Updates the elastic scaling rules of a cluster.
//
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
// Modifies and enables the maintenance window for an instance.
//
// Description:
//
// Before calling this operation, note the following:
//
// - Before the scheduled maintenance, Alibaba Cloud sends SMS messages and emails to the contacts configured in your Alibaba Cloud account. Check your messages promptly.
//
// - On the day of instance maintenance, to ensure stability throughout the maintenance process, the instance enters the Effective state before the maintenance window begins. While the instance is in this state, access to the cluster and query operations such as performance monitoring are not affected. However, cluster change operations such as cluster upgrades and restarts are temporarily unavailable.
//
// - During the maintenance window, transient disconnections may occur. Ensure that your application has a reconnection mechanism.
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
// Modifies and enables the maintenance window for an instance.
//
// Description:
//
// Before calling this operation, note the following:
//
// - Before the scheduled maintenance, Alibaba Cloud sends SMS messages and emails to the contacts configured in your Alibaba Cloud account. Check your messages promptly.
//
// - On the day of instance maintenance, to ensure stability throughout the maintenance process, the instance enters the Effective state before the maintenance window begins. While the instance is in this state, access to the cluster and query operations such as performance monitoring are not affected. However, cluster change operations such as cluster upgrades and restarts are temporarily unavailable.
//
// - During the maintenance window, transient disconnections may occur. Ensure that your application has a reconnection mechanism.
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
// For O&M events in the Event Center, you can specify a restart event, and the system will restart the specified edge zone of the relevant instance at the scheduled time.
//
// @param request - ModifyScheduleExecuteTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduleExecuteTimeResponse
func (client *Client) ModifyScheduleExecuteTimeWithOptions(instanceId *string, request *ModifyScheduleExecuteTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyScheduleExecuteTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.ScheduleExecuteTime) {
		query["scheduleExecuteTime"] = request.ScheduleExecuteTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScheduleExecuteTime"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/event/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/modify-execute-time"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScheduleExecuteTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For O&M events in the Event Center, you can specify a restart event, and the system will restart the specified edge zone of the relevant instance at the scheduled time.
//
// @param request - ModifyScheduleExecuteTimeRequest
//
// @return ModifyScheduleExecuteTimeResponse
func (client *Client) ModifyScheduleExecuteTime(instanceId *string, request *ModifyScheduleExecuteTimeRequest) (_result *ModifyScheduleExecuteTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyScheduleExecuteTimeResponse{}
	_body, _err := client.ModifyScheduleExecuteTimeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ModifyWhiteIps to update the access whitelist of a specified instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update information for an instance whose instance status is activating, invalid, or freeze (inactive).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters:
//
//   - IP whitelist list: whiteIpList, nodeType, networkType
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private network IP addresses, and the internal-facing whitelist does not support public IP addresses.
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
// Calls ModifyWhiteIps to update the access whitelist of a specified instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update information for an instance whose instance status is activating, invalid, or freeze (inactive).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters:
//
//   - IP whitelist list: whiteIpList, nodeType, networkType
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private network IP addresses, and the internal-facing whitelist does not support public IP addresses.
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
// Modifies the resource group to which an instance belongs.
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
// Modifies the resource group to which an instance belongs.
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

// Summary:
//
// Enables the intelligent O&M feature for an instance.
//
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

// Summary:
//
// Enables the intelligent O&M feature for an instance.
//
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

// Summary:
//
// Enables the HTTPS protocol. Before enabling HTTPS, make sure that you have purchased client nodes.
//
// Description:
//
// > - To ensure data security, enable the HTTPS protocol.
//
// - Except for versions 8.5 and 7.16<props="china"><ph> and version 7.10 in some regions</ph>, make sure that you have purchased client nodes before enabling HTTPS.
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

// Summary:
//
// Enables the HTTPS protocol. Before enabling HTTPS, make sure that you have purchased client nodes.
//
// Description:
//
// > - To ensure data security, enable the HTTPS protocol.
//
// - Except for versions 8.5 and 7.16<props="china"><ph> and version 7.10 in some regions</ph>, make sure that you have purchased client nodes before enabling HTTPS.
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
// Upload a custom plugin to the plugin repository. After uploading, the plugin is in the pending installation status.
//
// @param request - PluginAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PluginAnalysisResponse
func (client *Client) PluginAnalysisWithOptions(instanceId *string, request *PluginAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PluginAnalysisResponse, _err error) {
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
		Action:      dara.String("PluginAnalysis"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/plugins/actions/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PluginAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upload a custom plugin to the plugin repository. After uploading, the plugin is in the pending installation status.
//
// @param request - PluginAnalysisRequest
//
// @return PluginAnalysisResponse
func (client *Client) PluginAnalysis(instanceId *string, request *PluginAnalysisRequest) (_result *PluginAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PluginAnalysisResponse{}
	_body, _err := client.PluginAnalysisWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a test alert message by calling PostEmonTryAlarmRule.
//
// Description:
//
// > This API operation can be called up to 10 times per hour.
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
// Sends a test alert message by calling PostEmonTryAlarmRule.
//
// Description:
//
// > This API operation can be called up to 10 times per hour.
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

// Summary:
//
// Calls RecommendTemplates to retrieve recommended cluster configurations.
//
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

// Summary:
//
// Calls RecommendTemplates to retrieve recommended cluster configurations.
//
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
// Retries the installation of a collector that failed to install during creation.
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
// Retries the installation of a collector that failed to install during creation.
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
// You can delete uploaded but uninstalled plugins from the plugin library.
//
// @param request - RemovePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePluginResponse
func (client *Client) RemovePluginWithOptions(instanceId *string, request *RemovePluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemovePluginResponse, _err error) {
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
		Action:      dara.String("RemovePlugin"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/plugins/actions/remove"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can delete uploaded but uninstalled plugins from the plugin library.
//
// @param request - RemovePluginRequest
//
// @return RemovePluginResponse
func (client *Client) RemovePlugin(instanceId *string, request *RemovePluginRequest) (_result *RemovePluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemovePluginResponse{}
	_body, _err := client.RemovePluginWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a subscription Elasticsearch instance.
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
// Renews a subscription Elasticsearch instance.
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
// Renews a specified Logstash instance.
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
// Renews a specified Logstash instance.
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
// Restarts a collector to perform data collection.
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
// Restarts a collector to perform data collection.
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
// Restarts an Elasticsearch cluster.
//
// Description:
//
// > After the restart, the instance enters the activating state. After the restart is complete, the instance status changes to active. Alibaba Cloud Elasticsearch supports single-node restarts. Node restarts are classified into normal restarts and blue-green restarts.
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
// Restarts an Elasticsearch cluster.
//
// Description:
//
// > After the restart, the instance enters the activating state. After the restart is complete, the instance status changes to active. Alibaba Cloud Elasticsearch supports single-node restarts. Node restarts are classified into normal restarts and blue-green restarts.
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
// Restarts a specified instance. After the restart, the instance enters the activating (activing) state.
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
// Restarts a specified instance. After the restart, the instance enters the activating (activing) state.
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

// Summary:
//
// Resumes an interrupted change task for an instance.
//
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

// Summary:
//
// Resumes an interrupted change task for an instance.
//
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
// Resumes an interrupted instance change task. After the task is resumed, the instance enters the activating state.
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
// Resumes an interrupted instance change task. After the task is resumed, the instance enters the activating state.
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
// Creates a new index for a data stream or index alias.
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
// Creates a new index for a data stream or index alias.
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
// Deploys Logstash pipelines immediately.
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
// Deploys Logstash pipelines immediately.
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
// Scales in nodes of a specified role in an Elasticsearch cluster.
//
// Description:
//
// Note the following when you invoke this operation:
//
// Before scaling in data nodes of a cluster, perform data migration from the nodes to be removed to other nodes. After you confirm that the nodes to be removed contain no data, proceed with the scale-in operation.
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
// Scales in nodes of a specified role in an Elasticsearch cluster.
//
// Description:
//
// Note the following when you invoke this operation:
//
// Before scaling in data nodes of a cluster, perform data migration from the nodes to be removed to other nodes. After you confirm that the nodes to be removed contain no data, proceed with the scale-in operation.
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
// Calls StopCollector to stop a running collector.
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
// Calls StopCollector to stop a running collector.
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
// Stops Logstash pipelines by calling StopPipelines.
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
// Stops Logstash pipelines by calling StopPipelines.
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
// Creates tag-resource relationships for a specified instance.
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
// Creates tag-resource relationships for a specified instance.
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
// Performs data migration on a node to facilitate node scale-in operations.
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
// Performs data migration on a node to facilitate node scale-in operations.
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
// Enables or shuts down public or private network access for an Elasticsearch or Kibana cluster.
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
// Enables or shuts down public or private network access for an Elasticsearch or Kibana cluster.
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
// Disables an existing zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
//
// Description:
//
// Disables an existing zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
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
	if !dara.IsNil(request.HpAlbZoneDrained) {
		query["hpAlbZoneDrained"] = request.HpAlbZoneDrained
	}

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
// Disables an existing zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
//
// Description:
//
// Disables an existing zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
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
// Reopens an offline zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
//
// Description:
//
// Reopens an offline zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
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
	if !dara.IsNil(request.HpAlbZoneDrained) {
		query["hpAlbZoneDrained"] = request.HpAlbZoneDrained
	}

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
// Reopens an offline zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
//
// Description:
//
// Reopens an offline zone for a multi-zone instance. This operation is intended only for disaster recovery drills. Proceed with caution.
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
// Uninstalls plug-ins from the Kibana node of an Elasticsearch instance.
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
// Uninstalls plug-ins from the Kibana node of an Elasticsearch instance.
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
// Uninstalls installed plug-ins from a Logstash instance.
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
// Uninstalls installed plug-ins from a Logstash instance.
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
// Uninstalls system plug-ins from an Elasticsearch instance.
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
// Uninstalls system plug-ins from an Elasticsearch instance.
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
// Deletes user resource tag associations for a specified instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - Only user tags can be deleted.
//
//	> User tags are tags that users manually add to instances. System tags are tags that Alibaba Cloud services add to user instances. System tags are classified into visible tags and invisible tags.
//
// - If a tag is not associated with any resource, the tag is also deleted when the resource tag association is deleted.
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
// Deletes user resource tag associations for a specified instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - Only user tags can be deleted.
//
//	> User tags are tags that users manually add to instances. System tags are tags that Alibaba Cloud services add to user instances. System tags are classified into visible tags and invisible tags.
//
// - If a tag is not associated with any resource, the tag is also deleted when the resource tag association is deleted.
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
// Updates the password of the elastic account for a specified Elasticsearch instance.
//
// Description:
//
// When you invoke this operation, note the following:
//
// You cannot update information when the instance status is activating, invalid, or freeze (inactive).
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
// Updates the password of the elastic account for a specified Elasticsearch instance.
//
// Description:
//
// When you invoke this operation, note the following:
//
// You cannot update information when the instance status is activating, invalid, or freeze (inactive).
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
// Changes the garbage collector configuration of a specified Elasticsearch instance.
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
// Changes the garbage collector configuration of a specified Elasticsearch instance.
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
// Updates the dictionary of the AliNLP tokenizer plugin (analysis-aliws).
//
// Description:
//
// When calling this operation, note the following:
//
// - Instances of version 5.x do not support the AliNLP tokenizer plugin.
//
// - If the dictionary file is sourced from OSS, ensure that the OSS bucket has public-read permission.
//
// - If a previously uploaded dictionary is not configured with ORIGIN, calling this operation will delete the dictionary file.
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
// Updates the dictionary of the AliNLP tokenizer plugin (analysis-aliws).
//
// Description:
//
// When calling this operation, note the following:
//
// - Instances of version 5.x do not support the AliNLP tokenizer plugin.
//
// - If the dictionary file is sourced from OSS, ensure that the OSS bucket has public-read permission.
//
// - If a previously uploaded dictionary is not configured with ORIGIN, calling this operation will delete the dictionary file.
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
// Modifies the access blacklist of an Elasticsearch instance. This operation is deprecated.
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
// Modifies the access blacklist of an Elasticsearch instance. This operation is deprecated.
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
// Modifies the configuration of a collector.
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
// Modifies the configuration of a collector.
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
// Calls UpdateCollectorName to modify the name of a collector.
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
// Calls UpdateCollectorName to modify the name of a collector.
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
// Updates a composable index template for an Elasticsearch instance.
//
// Description:
//
// For more information, see [Use OpenStore to store massive amounts of data](https://help.aliyun.com/document_detail/317694.html).
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
// Updates a composable index template for an Elasticsearch instance.
//
// Description:
//
// For more information, see [Use OpenStore to store massive amounts of data](https://help.aliyun.com/document_detail/317694.html).
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
// Changes the name of an Elasticsearch instance.
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
// Changes the name of an Elasticsearch instance.
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
// Modifies the intelligent O&M scenario settings of a specified Elasticsearch instance.
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
// Modifies the intelligent O&M scenario settings of a specified Elasticsearch instance.
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
// Updates the user dictionary of an Elasticsearch instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - If the dictionary file originates from OSS, ensure that the OSS storage space is publicly readable.
//
// - If previously uploaded dictionaries are not configured with ORIGIN, the dictionary files will be deleted after this operation is called.
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
// Updates the user dictionary of an Elasticsearch instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - If the dictionary file originates from OSS, ensure that the OSS storage space is publicly readable.
//
// - If previously uploaded dictionaries are not configured with ORIGIN, the dictionary files will be deleted after this operation is called.
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
// # Modify Cluster Dynamic Configuration
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
// # Modify Cluster Dynamic Configuration
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

// Summary:
//
// Modifies the scenario-based configuration template of an Elasticsearch instance.
//
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

// Summary:
//
// Modifies the scenario-based configuration template of an Elasticsearch instance.
//
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
// Updates the extension file configuration of a Logstash instance.
//
// Description:
//
// When calling this operation, note the following: Currently, this operation only supports deleting extension files that have been uploaded through the console. To add or modify extension files, perform the operations in the console.
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
// Updates the extension file configuration of a Logstash instance.
//
// Description:
//
// When calling this operation, note the following: Currently, this operation only supports deleting extension files that have been uploaded through the console. To add or modify extension files, perform the operations in the console.
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
// Toggle the FalconSeek cloud-native kernel attribute for instances of Version 8.17.
//
// @param request - UpdateFalconSeekRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFalconSeekResponse
func (client *Client) UpdateFalconSeekWithOptions(InstanceId *string, request *UpdateFalconSeekRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFalconSeekResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["enable"] = request.Enable
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFalconSeek"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/falconseek"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFalconSeekResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Toggle the FalconSeek cloud-native kernel attribute for instances of Version 8.17.
//
// @param request - UpdateFalconSeekRequest
//
// @return UpdateFalconSeekResponse
func (client *Client) UpdateFalconSeek(InstanceId *string, request *UpdateFalconSeekRequest) (_result *UpdateFalconSeekResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFalconSeekResponse{}
	_body, _err := client.UpdateFalconSeekWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the IK hot-word dictionary of an Alibaba Cloud Elasticsearch instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - If the dictionary file is sourced from OSS, make sure the OSS bucket has public-read permission.
//
// - If a previously uploaded dictionary is not configured with ORIGIN, the dictionary file will be deleted after this operation is called.
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
// Updates the IK hot-word dictionary of an Alibaba Cloud Elasticsearch instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - If the dictionary file is sourced from OSS, make sure the OSS bucket has public-read permission.
//
// - If a previously uploaded dictionary is not configured with ORIGIN, the dictionary file will be deleted after this operation is called.
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
// Modifies the lifecycle policy of an Elasticsearch index.
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
// Modifies the lifecycle policy of an Elasticsearch index.
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
// Modifies the template configuration of an Elasticsearch instance.
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
// Modifies the template configuration of an Elasticsearch instance.
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
// Upgrades the configuration of an Elasticsearch cluster, including the number of nodes, roles, specifications, and disk configurations.
//
// Description:
//
// When you call this operation, note the following items:
//
// For more precautions, see [Upgrade cluster configuration](https://help.aliyun.com/document_detail/96650.html) and [Downgrade cluster configuration](https://help.aliyun.com/document_detail/198887.html).
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

	if !dara.IsNil(request.UpdateType) {
		body["updateType"] = request.UpdateType
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
// Upgrades the configuration of an Elasticsearch cluster, including the number of nodes, roles, specifications, and disk configurations.
//
// Description:
//
// When you call this operation, note the following items:
//
// For more precautions, see [Upgrade cluster configuration](https://help.aliyun.com/document_detail/96650.html) and [Downgrade cluster configuration](https://help.aliyun.com/document_detail/198887.html).
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
// Transforms the billing method of an Elasticsearch instance from pay-as-you-go to a subscription instance.
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.PaymentInfo) {
		body["paymentInfo"] = request.PaymentInfo
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
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
// Transforms the billing method of an Elasticsearch instance from pay-as-you-go to a subscription instance.
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
// Modifies the YML parameter settings of a specified Elasticsearch instance.
//
// Description:
//
// When you invoke this operation, note the following:
//
// You cannot update the configuration when the instance status is activating, invalid, or inactive (freeze).
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.EsConfig) {
		body["esConfig"] = request.EsConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
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
// Modifies the YML parameter settings of a specified Elasticsearch instance.
//
// Description:
//
// When you invoke this operation, note the following:
//
// You cannot update the configuration when the instance status is activating, invalid, or inactive (freeze).
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
// # Update keystore
//
// @param request - UpdateKeystoresRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKeystoresResponse
func (client *Client) UpdateKeystoresWithOptions(InstanceId *string, request *UpdateKeystoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKeystoresResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Remove) {
		body["remove"] = request.Remove
	}

	if !dara.IsNil(request.Update) {
		body["update"] = request.Update
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKeystores"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/keystores"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKeystoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update keystore
//
// @param request - UpdateKeystoresRequest
//
// @return UpdateKeystoresResponse
func (client *Client) UpdateKeystores(InstanceId *string, request *UpdateKeystoresRequest) (_result *UpdateKeystoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKeystoresResponse{}
	_body, _err := client.UpdateKeystoresWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Kibana private network access
//
// Description:
//
// 1. This API operation supports only cloud-native instances. For instances of the legacy architecture, use the TriggerNetwork operation.
//
// 2. The Kibana specifications must be greater than 1 vCPU and 2 GB of memory.
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
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

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
// # Update Kibana private network access
//
// Description:
//
// 1. This API operation supports only cloud-native instances. For instances of the legacy architecture, use the TriggerNetwork operation.
//
// 2. The Kibana specifications must be greater than 1 vCPU and 2 GB of memory.
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
// Modifies the Kibana configuration. Currently, only the Kibana language configuration can be modified.
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
// Modifies the Kibana configuration. Currently, only the Kibana language configuration can be modified.
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
// Enable or disable Alibaba Cloud account authentication for Kibana. After Alibaba Cloud account authentication is enabled, you must log on with your Alibaba Cloud account before you can use Kibana features.
//
// @param request - UpdateKibanaSsoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaSsoResponse
func (client *Client) UpdateKibanaSsoWithOptions(InstanceId *string, request *UpdateKibanaSsoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaSsoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["enable"] = request.Enable
	}

	if !dara.IsNil(request.NetworkType) {
		query["networkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKibanaSso"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/actions/kibana-sso"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKibanaSsoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enable or disable Alibaba Cloud account authentication for Kibana. After Alibaba Cloud account authentication is enabled, you must log on with your Alibaba Cloud account before you can use Kibana features.
//
// @param request - UpdateKibanaSsoRequest
//
// @return UpdateKibanaSsoResponse
func (client *Client) UpdateKibanaSso(InstanceId *string, request *UpdateKibanaSsoRequest) (_result *UpdateKibanaSsoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKibanaSsoResponse{}
	_body, _err := client.UpdateKibanaSsoWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Kibana access whitelist of a specified Alibaba Cloud Elasticsearch instance.
//
// Description:
//
// ## Before you begin
//
// - When you invoke this operation, you cannot update information if the instance status is activating, invalid, or freeze (inactive).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters, as follows:
//
//   - IP whitelist list: kibanaIPWhitelist
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private IP addresses, and the internal-facing whitelist does not support public IP addresses.
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
// Updates the Kibana access whitelist of a specified Alibaba Cloud Elasticsearch instance.
//
// Description:
//
// ## Before you begin
//
// - When you invoke this operation, you cannot update information if the instance status is activating, invalid, or freeze (inactive).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters, as follows:
//
//   - IP whitelist list: kibanaIPWhitelist
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private IP addresses, and the internal-facing whitelist does not support public IP addresses.
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
// Modifies some information about a specified instance, such as the number of nodes, quota, name, and disk size.
//
// Description:
//
// ### Before you begin
//
// You cannot modify instance information when the instance status is activating, invalid, or freeze (inactive).
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
// Modifies some information about a specified instance, such as the number of nodes, quota, name, and disk size.
//
// Description:
//
// ### Before you begin
//
// You cannot modify instance information when the instance status is activating, invalid, or freeze (inactive).
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
// Converts a pay-as-you-go Alibaba Cloud Logstash instance to a subscription instance.
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
// Converts a pay-as-you-go Alibaba Cloud Logstash instance to a subscription instance.
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
// Modifies the name of a specified Logstash instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// You cannot modify the instance name when the instance status is activating, invalid, or freeze (inactive).
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
// Modifies the name of a specified Logstash instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
// You cannot modify the instance name when the instance status is activating, invalid, or freeze (inactive).
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
// Updates the configuration of a specified Logstash instance.
//
// Description:
//
// When you invoke this operation, note the following: The instance configuration cannot be updated when the instance status is activating, invalid, or freeze (inactive).
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
// Updates the configuration of a specified Logstash instance.
//
// Description:
//
// When you invoke this operation, note the following: The instance configuration cannot be updated when the instance status is activating, invalid, or freeze (inactive).
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
// Modifies the pipeline management method for a specified Logstash instance.
//
// Description:
//
// > Pipeline management methods include configuration file management and Kibana pipeline management. The console no longer supports Kibana pipeline management. You can use this feature only through the API.
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
// Modifies the pipeline management method for a specified Logstash instance.
//
// Description:
//
// > Pipeline management methods include configuration file management and Kibana pipeline management. The console no longer supports Kibana pipeline management. You can use this feature only through the API.
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
// Calls UpdatePipelines to update Logstash pipeline information.
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
// Calls UpdatePipelines to update Logstash pipeline information.
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
// Updates the VPC private network access whitelist of a specified instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update the VPC private network access whitelist of an instance when the instance status is Activating (activating), Invalid (invalid), or Freeze (inactive).
//
// - You can update the whitelist in two ways: by using an IP whitelist list or by using an IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters:
//
//   - IP whitelist list: privateNetworkIpWhiteList
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private IP addresses, and the private network access whitelist does not support public IP addresses.
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
// Updates the VPC private network access whitelist of a specified instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update the VPC private network access whitelist of an instance when the instance status is Activating (activating), Invalid (invalid), or Freeze (inactive).
//
// - You can update the whitelist in two ways: by using an IP whitelist list or by using an IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters:
//
//   - IP whitelist list: privateNetworkIpWhiteList
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private IP addresses, and the private network access whitelist does not support public IP addresses.
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
// Enables or disables the public network address for a specified Elasticsearch instance.
//
// Description:
//
// When you call this operation, note the following:
//
// You cannot update information when the instance status is activating, invalid, or freeze (inactive).
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
// Enables or disables the public network address for a specified Elasticsearch instance.
//
// Description:
//
// When you call this operation, note the following:
//
// You cannot update information when the instance status is activating, invalid, or freeze (inactive).
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
// Updates the public endpoint access whitelist of a specified Elasticsearch instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update the public endpoint access whitelist of an instance when the instance status is activating, invalid, or inactive (freeze).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters:
//
//   - IP whitelist list: publicIpWhitelist
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private network IP addresses, and the private network access whitelist does not support public network IP addresses.
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
// Updates the public endpoint access whitelist of a specified Elasticsearch instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update the public endpoint access whitelist of an instance when the instance status is activating, invalid, or inactive (freeze).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time. In addition to InstanceId and clientToken, the two methods support different parameters:
//
//   - IP whitelist list: publicIpWhitelist
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public network access whitelist does not support private network IP addresses, and the private network access whitelist does not support public network IP addresses.
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
// Enables or disables the write high availability feature for a cluster. Currently, only instances in the China (Beijing) region are supported.
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
// Enables or disables the write high availability feature for a cluster. Currently, only instances in the China (Beijing) region are supported.
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
// Updates the data backup configuration of a specified instance.
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
// Updates the data backup configuration of a specified instance.
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
// Updates the synonym dictionary of an Alibaba Cloud Elasticsearch instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - If the dictionary file is sourced from OSS, make sure the OSS bucket has public-read permission.
//
// - If a previously uploaded dictionary is not configured with ORIGIN, the dictionary file will be deleted after this operation is called.
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
// Updates the synonym dictionary of an Alibaba Cloud Elasticsearch instance.
//
// Description:
//
// When calling this operation, note the following:
//
// - If the dictionary file is sourced from OSS, make sure the OSS bucket has public-read permission.
//
// - If a previously uploaded dictionary is not configured with ORIGIN, the dictionary file will be deleted after this operation is called.
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

// Summary:
//
// Modifies the scenario-specific template configuration of a cluster.
//
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

// Summary:
//
// Modifies the scenario-specific template configuration of a cluster.
//
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
// Modifies the VPC internal-facing access whitelist of an Elasticsearch instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update information when the instance status is activating, invalid, or freeze (inactive).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time, and they support different parameters besides InstanceId and clientToken. The details are as follows:
//
//   - IP whitelist list: esIPWhitelist
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public access whitelist does not support private network IP addresses, and the internal-facing access whitelist does not support public IP addresses.
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
// Modifies the VPC internal-facing access whitelist of an Elasticsearch instance.
//
// Description:
//
// ## Before you begin
//
// - You cannot update information when the instance status is activating, invalid, or freeze (inactive).
//
// - You can update the whitelist in two ways: IP whitelist list and IP whitelist group. The two methods cannot be used at the same time, and they support different parameters besides InstanceId and clientToken. The details are as follows:
//
//   - IP whitelist list: esIPWhitelist
//
//   - IP whitelist group: modifyMode, whiteIpGroup
//
// - The public access whitelist does not support private network IP addresses, and the internal-facing access whitelist does not support public IP addresses.
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
// Updates the X-Pack monitoring and alerting configuration of a Logstash instance.
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
// Updates the X-Pack monitoring and alerting configuration of a Logstash instance.
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
// Upgrades the version of an Elasticsearch instance. Both major version upgrades and kernel version upgrades are supported.
//
// Description:
//
// > The version upgrade feature currently supports only the following upgrade paths: 5.5.3 to 5.6.16, 5.6.16 to 6.3.2, and 6.3.2 to 6.7.0. Upgrades between other versions are not supported. For more information, see [Upgrade version](https://help.aliyun.com/document_detail/148786.html).
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
// Upgrades the version of an Elasticsearch instance. Both major version upgrades and kernel version upgrades are supported.
//
// Description:
//
// > The version upgrade feature currently supports only the following upgrade paths: 5.5.3 to 5.6.16, 5.6.16 to 6.3.2, and 6.3.2 to 6.7.0. Upgrades between other versions are not supported. For more information, see [Upgrade version](https://help.aliyun.com/document_detail/148786.html).
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
// Query whether a minor version is available for upgrade.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeInfoResponse
func (client *Client) UpgradeInfoWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeInfo"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgradeInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query whether a minor version is available for upgrade.
//
// @return UpgradeInfoResponse
func (client *Client) UpgradeInfo(instanceId *string) (_result *UpgradeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeInfoResponse{}
	_body, _err := client.UpgradeInfoWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates the connectivity of an Elasticsearch instance that provides X-Pack monitoring.
//
// Description:
//
// > To enable X-Pack monitoring for Logstash, configure an Elasticsearch instance. After the configuration, you can monitor the Logstash instance in the Kibana console of the corresponding Elasticsearch instance.
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
// Validates the connectivity of an Elasticsearch instance that provides X-Pack monitoring.
//
// Description:
//
// > To enable X-Pack monitoring for Logstash, configure an Elasticsearch instance. After the configuration, you can monitor the Logstash instance in the Kibana console of the corresponding Elasticsearch instance.
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
// Checks whether specific nodes in a specified instance can be scaled in.
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
// Checks whether specific nodes in a specified instance can be scaled in.
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

// Summary:
//
// Verifies the service-linked role (SLR) permission of the current account.
//
// Description:
//
// > When you use a collector to collect logs from different data sources, you must first authorize the creation of a service-linked role. You can call this operation to check whether the service-linked role has been created.
//
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

// Summary:
//
// Verifies the service-linked role (SLR) permission of the current account.
//
// Description:
//
// > When you use a collector to collect logs from different data sources, you must first authorize the creation of a service-linked role. You can call this operation to check whether the service-linked role has been created.
//
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
// Validates whether data on specific nodes in a specified instance can be migrated.
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
// Validates whether data on specific nodes in a specified instance can be migrated.
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
// Creates an Elasticsearch instance.
//
// Description:
//
// ### Precautions
//
// - Before using this operation, make sure that you fully understand the billing methods and pricing of Elasticsearch. For more information, see [Alibaba Cloud Elasticsearch pricing](https://www.aliyun.com/price/product?spm=a2c4g.11186623.2.7.657d2cbeRoSPCd#/elasticsearch/detail).
//
// - Real-name verification is required to create instances.<props="china"><ph> For more information, see [Real-name verification](https://help.aliyun.com/document_detail/37175.html).</ph>
//
// - You do not need to specify a zone when creating an instance. By default, the instance is in the same zone as the selected VPC.
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
// Creates an Elasticsearch instance.
//
// Description:
//
// ### Precautions
//
// - Before using this operation, make sure that you fully understand the billing methods and pricing of Elasticsearch. For more information, see [Alibaba Cloud Elasticsearch pricing](https://www.aliyun.com/price/product?spm=a2c4g.11186623.2.7.657d2cbeRoSPCd#/elasticsearch/detail).
//
// - Real-name verification is required to create instances.<props="china"><ph> For more information, see [Real-name verification](https://help.aliyun.com/document_detail/37175.html).</ph>
//
// - You do not need to specify a zone when creating an instance. By default, the instance is in the same zone as the selected VPC.
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
