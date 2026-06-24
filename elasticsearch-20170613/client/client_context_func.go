// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) ActivateZonesWithContext(ctx context.Context, InstanceId *string, request *ActivateZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateZonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddConnectableClusterResponse
func (client *Client) AddConnectableClusterWithContext(ctx context.Context, InstanceId *string, request *AddConnectableClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddConnectableClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSnapshotRepoResponse
func (client *Client) AddSnapshotRepoWithContext(ctx context.Context, InstanceId *string, request *AddSnapshotRepoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddSnapshotRepoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDeletionResponse
func (client *Client) CancelDeletionWithContext(ctx context.Context, InstanceId *string, request *CancelDeletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelDeletionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLogstashDeletionResponse
func (client *Client) CancelLogstashDeletionWithContext(ctx context.Context, InstanceId *string, request *CancelLogstashDeletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelLogstashDeletionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, InstanceId *string, request *CancelTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CapacityPlanResponse
func (client *Client) CapacityPlanWithContext(ctx context.Context, request *CapacityPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CapacityPlanResponse, _err error) {
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
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDiagnosisResponse
func (client *Client) CloseDiagnosisWithContext(ctx context.Context, InstanceId *string, request *CloseDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseHttpsResponse
func (client *Client) CloseHttpsWithContext(ctx context.Context, InstanceId *string, request *CloseHttpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseHttpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseManagedIndexResponse
func (client *Client) CloseManagedIndexWithContext(ctx context.Context, InstanceId *string, Index *string, request *CloseManagedIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseManagedIndexResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCollectorResponse
func (client *Client) CreateCollectorWithContext(ctx context.Context, request *CreateCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentIndexResponse
func (client *Client) CreateComponentIndexWithContext(ctx context.Context, InstanceId *string, name *string, request *CreateComponentIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComponentIndexResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataStreamResponse
func (client *Client) CreateDataStreamWithContext(ctx context.Context, InstanceId *string, request *CreateDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataStreamResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateILMPolicyResponse
func (client *Client) CreateILMPolicyWithContext(ctx context.Context, InstanceId *string, request *CreateILMPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateILMPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexTemplateResponse
func (client *Client) CreateIndexTemplateWithContext(ctx context.Context, InstanceId *string, request *CreateIndexTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndexTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogstashResponse
func (client *Client) CreateLogstashWithContext(ctx context.Context, request *CreateLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelinesResponse
func (client *Client) CreatePipelinesWithContext(ctx context.Context, InstanceId *string, request *CreatePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePipelinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithContext(ctx context.Context, InstanceId *string, request *CreateSnapshotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpointWithContext(ctx context.Context, InstanceId *string, request *CreateVpcEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateZonesResponse
func (client *Client) DeactivateZonesWithContext(ctx context.Context, InstanceId *string, request *DeactivateZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeactivateZonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCollectorResponse
func (client *Client) DeleteCollectorWithContext(ctx context.Context, ResId *string, request *DeleteCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentIndexResponse
func (client *Client) DeleteComponentIndexWithContext(ctx context.Context, InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteComponentIndexResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectedClusterResponse
func (client *Client) DeleteConnectedClusterWithContext(ctx context.Context, InstanceId *string, request *DeleteConnectedClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConnectedClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataStreamResponse
func (client *Client) DeleteDataStreamWithContext(ctx context.Context, InstanceId *string, DataStream *string, request *DeleteDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataStreamResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataTaskResponse
func (client *Client) DeleteDataTaskWithContext(ctx context.Context, InstanceId *string, request *DeleteDataTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeprecatedTemplateResponse
func (client *Client) DeleteDeprecatedTemplateWithContext(ctx context.Context, InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDeprecatedTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteILMPolicyResponse
func (client *Client) DeleteILMPolicyWithContext(ctx context.Context, InstanceId *string, PolicyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteILMPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexTemplateResponse
func (client *Client) DeleteIndexTemplateWithContext(ctx context.Context, InstanceId *string, IndexTemplate *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndexTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, InstanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogstashResponse
func (client *Client) DeleteLogstashWithContext(ctx context.Context, InstanceId *string, request *DeleteLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelinesResponse
func (client *Client) DeletePipelinesWithContext(ctx context.Context, InstanceId *string, request *DeletePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotRepoResponse
func (client *Client) DeleteSnapshotRepoWithContext(ctx context.Context, InstanceId *string, request *DeleteSnapshotRepoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotRepoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpointWithContext(ctx context.Context, InstanceId *string, EndpointId *string, request *DeleteVpcEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckOperatorResponse
func (client *Client) DescribeAckOperatorWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAckOperatorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCollectorResponse
func (client *Client) DescribeCollectorWithContext(ctx context.Context, ResId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentIndexResponse
func (client *Client) DescribeComponentIndexWithContext(ctx context.Context, InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeComponentIndexResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConnectableClustersResponse
func (client *Client) DescribeConnectableClustersWithContext(ctx context.Context, InstanceId *string, request *DescribeConnectableClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeConnectableClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeprecatedTemplateResponse
func (client *Client) DescribeDeprecatedTemplateWithContext(ctx context.Context, InstanceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDeprecatedTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnoseReportResponse
func (client *Client) DescribeDiagnoseReportWithContext(ctx context.Context, InstanceId *string, ReportId *string, request *DescribeDiagnoseReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDiagnoseReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisSettingsResponse
func (client *Client) DescribeDiagnosisSettingsWithContext(ctx context.Context, InstanceId *string, request *DescribeDiagnosisSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDynamicSettingsResponse
func (client *Client) DescribeDynamicSettingsWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDynamicSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticsearchHealthResponse
func (client *Client) DescribeElasticsearchHealthWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeElasticsearchHealthResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeILMPolicyResponse
func (client *Client) DescribeILMPolicyWithContext(ctx context.Context, InstanceId *string, PolicyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeILMPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIndexTemplateResponse
func (client *Client) DescribeIndexTemplateWithContext(ctx context.Context, InstanceId *string, IndexTemplate *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeIndexTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKibanaSettingsResponse
func (client *Client) DescribeKibanaSettingsWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeKibanaSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogstashResponse
func (client *Client) DescribeLogstashWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePipelineResponse
func (client *Client) DescribePipelineWithContext(ctx context.Context, InstanceId *string, PipelineId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePipelineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePipelineManagementConfigResponse
func (client *Client) DescribePipelineManagementConfigWithContext(ctx context.Context, InstanceId *string, request *DescribePipelineManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePipelineManagementConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotSettingResponse
func (client *Client) DescribeSnapshotSettingWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeXpackMonitorConfigResponse
func (client *Client) DescribeXpackMonitorConfigWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeXpackMonitorConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiagnoseInstanceResponse
func (client *Client) DiagnoseInstanceWithContext(ctx context.Context, InstanceId *string, request *DiagnoseInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DiagnoseInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableKibanaPvlNetworkResponse
func (client *Client) DisableKibanaPvlNetworkWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableKibanaPvlNetworkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableKibanaPvlNetworkResponse
func (client *Client) EnableKibanaPvlNetworkWithContext(ctx context.Context, InstanceId *string, request *EnableKibanaPvlNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableKibanaPvlNetworkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedLogstashRestartTimeResponse
func (client *Client) EstimatedLogstashRestartTimeWithContext(ctx context.Context, InstanceId *string, request *EstimatedLogstashRestartTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EstimatedLogstashRestartTimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedRestartTimeResponse
func (client *Client) EstimatedRestartTimeWithContext(ctx context.Context, InstanceId *string, request *EstimatedRestartTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EstimatedRestartTimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDataInformationResponse
func (client *Client) GetClusterDataInformationWithContext(ctx context.Context, request *GetClusterDataInformationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDataInformationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElastictaskResponse
func (client *Client) GetElastictaskWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetElastictaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonAlarmRecordStatisticsDistributeResponse
func (client *Client) GetEmonAlarmRecordStatisticsDistributeWithContext(ctx context.Context, request *GetEmonAlarmRecordStatisticsDistributeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonAlarmRecordStatisticsDistributeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonGrafanaAlertsResponse
func (client *Client) GetEmonGrafanaAlertsWithContext(ctx context.Context, ProjectId *string, request *GetEmonGrafanaAlertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonGrafanaAlertsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonGrafanaDashboardsResponse
func (client *Client) GetEmonGrafanaDashboardsWithContext(ctx context.Context, ProjectId *string, request *GetEmonGrafanaDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonGrafanaDashboardsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonMonitorDataResponse
func (client *Client) GetEmonMonitorDataWithContext(ctx context.Context, ProjectId *string, request *GetEmonMonitorDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonMonitorDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeystoresResponse
func (client *Client) GetKeystoresWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKeystoresResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenStoreUsageResponse
func (client *Client) GetOpenStoreUsageWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOpenStoreUsageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionConfigurationResponse
func (client *Client) GetRegionConfigurationWithContext(ctx context.Context, request *GetRegionConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionConfigurationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionalInstanceConfigResponse
func (client *Client) GetRegionalInstanceConfigWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionalInstanceConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuggestShrinkableNodesResponse
func (client *Client) GetSuggestShrinkableNodesWithContext(ctx context.Context, InstanceId *string, request *GetSuggestShrinkableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSuggestShrinkableNodesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransferableNodesResponse
func (client *Client) GetTransferableNodesWithContext(ctx context.Context, InstanceId *string, request *GetTransferableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTransferableNodesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrayPublishResponse
func (client *Client) GrayPublishWithContext(ctx context.Context, InstanceId *string, request *GrayPublishRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrayPublishResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitModelResponse
func (client *Client) InitModelWithContext(ctx context.Context, InstanceId *string, request *InitModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitModelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeOperationRoleResponse
func (client *Client) InitializeOperationRoleWithContext(ctx context.Context, request *InitializeOperationRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitializeOperationRoleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAckOperatorResponse
func (client *Client) InstallAckOperatorWithContext(ctx context.Context, ClusterId *string, request *InstallAckOperatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAckOperatorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallKibanaSystemPluginResponse
func (client *Client) InstallKibanaSystemPluginWithContext(ctx context.Context, InstanceId *string, request *InstallKibanaSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallKibanaSystemPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallLogstashSystemPluginResponse
func (client *Client) InstallLogstashSystemPluginWithContext(ctx context.Context, InstanceId *string, request *InstallLogstashSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallLogstashSystemPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallSystemPluginResponse
func (client *Client) InstallSystemPluginWithContext(ctx context.Context, InstanceId *string, request *InstallSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallSystemPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallUserPluginsResponse
func (client *Client) InstallUserPluginsWithContext(ctx context.Context, InstanceId *string, request *InstallUserPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallUserPluginsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterruptElasticsearchTaskResponse
func (client *Client) InterruptElasticsearchTaskWithContext(ctx context.Context, InstanceId *string, request *InterruptElasticsearchTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InterruptElasticsearchTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterruptLogstashTaskResponse
func (client *Client) InterruptLogstashTaskWithContext(ctx context.Context, InstanceId *string, request *InterruptLogstashTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InterruptLogstashTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAckClustersResponse
func (client *Client) ListAckClustersWithContext(ctx context.Context, request *ListAckClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAckClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAckNamespacesResponse
func (client *Client) ListAckNamespacesWithContext(ctx context.Context, ClusterId *string, request *ListAckNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAckNamespacesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionRecordsResponse
func (client *Client) ListActionRecordsWithContext(ctx context.Context, InstanceId *string, request *ListActionRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListActionRecordsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllNodeResponse
func (client *Client) ListAllNodeWithContext(ctx context.Context, InstanceId *string, request *ListAllNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllNodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlternativeSnapshotReposResponse
func (client *Client) ListAlternativeSnapshotReposWithContext(ctx context.Context, InstanceId *string, request *ListAlternativeSnapshotReposRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlternativeSnapshotReposResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableEsInstanceIdsResponse
func (client *Client) ListAvailableEsInstanceIdsWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvailableEsInstanceIdsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectorsResponse
func (client *Client) ListCollectorsWithContext(ctx context.Context, request *ListCollectorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCollectorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentIndicesResponse
func (client *Client) ListComponentIndicesWithContext(ctx context.Context, InstanceId *string, request *ListComponentIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComponentIndicesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectedClustersResponse
func (client *Client) ListConnectedClustersWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConnectedClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataStreamsResponse
func (client *Client) ListDataStreamsWithContext(ctx context.Context, InstanceId *string, request *ListDataStreamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataStreamsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataTasksResponse
func (client *Client) ListDataTasksWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDefaultCollectorConfigurationsResponse
func (client *Client) ListDefaultCollectorConfigurationsWithContext(ctx context.Context, request *ListDefaultCollectorConfigurationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDefaultCollectorConfigurationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeprecatedTemplatesResponse
func (client *Client) ListDeprecatedTemplatesWithContext(ctx context.Context, InstanceId *string, request *ListDeprecatedTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeprecatedTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseIndicesResponse
func (client *Client) ListDiagnoseIndicesWithContext(ctx context.Context, InstanceId *string, request *ListDiagnoseIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseIndicesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseReportResponse
func (client *Client) ListDiagnoseReportWithContext(ctx context.Context, InstanceId *string, request *ListDiagnoseReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseReportIdsResponse
func (client *Client) ListDiagnoseReportIdsWithContext(ctx context.Context, InstanceId *string, request *ListDiagnoseReportIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseReportIdsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosisItemsResponse
func (client *Client) ListDiagnosisItemsWithContext(ctx context.Context, request *ListDiagnosisItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnosisItemsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictInformationResponse
func (client *Client) ListDictInformationWithContext(ctx context.Context, InstanceId *string, request *ListDictInformationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictInformationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictsResponse
func (client *Client) ListDictsWithContext(ctx context.Context, InstanceId *string, request *ListDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcsInstancesResponse
func (client *Client) ListEcsInstancesWithContext(ctx context.Context, request *ListEcsInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcsInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventRecordsResponse
func (client *Client) ListEventRecordsWithContext(ctx context.Context, eventType *string, request *ListEventRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEventRecordsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExtendfilesResponse
func (client *Client) ListExtendfilesWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExtendfilesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListILMPoliciesResponse
func (client *Client) ListILMPoliciesWithContext(ctx context.Context, InstanceId *string, request *ListILMPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListILMPoliciesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexTemplatesResponse
func (client *Client) ListIndexTemplatesWithContext(ctx context.Context, InstanceId *string, request *ListIndexTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListInstanceHistoryEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHistoryEventsResponse
func (client *Client) ListInstanceHistoryEventsWithContext(ctx context.Context, tmpReq *ListInstanceHistoryEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceHistoryEventsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceIndicesResponse
func (client *Client) ListInstanceIndicesWithContext(ctx context.Context, InstanceId *string, request *ListInstanceIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceIndicesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKibanaPluginsResponse
func (client *Client) ListKibanaPluginsWithContext(ctx context.Context, InstanceId *string, request *ListKibanaPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKibanaPluginsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKibanaPvlNetworkResponse
func (client *Client) ListKibanaPvlNetworkWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKibanaPvlNetworkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashResponse
func (client *Client) ListLogstashWithContext(ctx context.Context, request *ListLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashLogResponse
func (client *Client) ListLogstashLogWithContext(ctx context.Context, InstanceId *string, request *ListLogstashLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashLogResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashPluginsResponse
func (client *Client) ListLogstashPluginsWithContext(ctx context.Context, InstanceId *string, request *ListLogstashPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashPluginsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, ResId *string, request *ListNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineResponse
func (client *Client) ListPipelineWithContext(ctx context.Context, InstanceId *string, request *ListPipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineIdsResponse
func (client *Client) ListPipelineIdsWithContext(ctx context.Context, InstanceId *string, request *ListPipelineIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineIdsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsResponse
func (client *Client) ListPluginsWithContext(ctx context.Context, InstanceId *string, request *ListPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchLogResponse
func (client *Client) ListSearchLogWithContext(ctx context.Context, InstanceId *string, request *ListSearchLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchLogResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShardRecoveriesResponse
func (client *Client) ListShardRecoveriesWithContext(ctx context.Context, InstanceId *string, request *ListShardRecoveriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShardRecoveriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotReposByInstanceIdResponse
func (client *Client) ListSnapshotReposByInstanceIdWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSnapshotReposByInstanceIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStatsEventRecordsResponse
func (client *Client) ListStatsEventRecordsWithContext(ctx context.Context, request *ListStatsEventRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStatsEventRecordsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPluginResponse
func (client *Client) ListUserPluginWithContext(ctx context.Context, instanceId *string, request *ListUserPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpointsWithContext(ctx context.Context, InstanceId *string, request *ListVpcEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateToOtherZoneResponse
func (client *Client) MigrateToOtherZoneWithContext(ctx context.Context, InstanceId *string, request *MigrateToOtherZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeployMachineResponse
func (client *Client) ModifyDeployMachineWithContext(ctx context.Context, ResId *string, request *ModifyDeployMachineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDeployMachineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElastictaskResponse
func (client *Client) ModifyElastictaskWithContext(ctx context.Context, InstanceId *string, request *ModifyElastictaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyElastictaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTimeWithContext(ctx context.Context, InstanceId *string, request *ModifyInstanceMaintainTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduleExecuteTimeResponse
func (client *Client) ModifyScheduleExecuteTimeWithContext(ctx context.Context, instanceId *string, request *ModifyScheduleExecuteTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyScheduleExecuteTimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWhiteIpsResponse
func (client *Client) ModifyWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *ModifyWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyWhiteIpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithContext(ctx context.Context, InstanceId *string, request *MoveResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDiagnosisResponse
func (client *Client) OpenDiagnosisWithContext(ctx context.Context, InstanceId *string, request *OpenDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenHttpsResponse
func (client *Client) OpenHttpsWithContext(ctx context.Context, InstanceId *string, request *OpenHttpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenHttpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PluginAnalysisResponse
func (client *Client) PluginAnalysisWithContext(ctx context.Context, instanceId *string, request *PluginAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PluginAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEmonTryAlarmRuleResponse
func (client *Client) PostEmonTryAlarmRuleWithContext(ctx context.Context, ProjectId *string, AlarmGroupId *string, request *PostEmonTryAlarmRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PostEmonTryAlarmRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecommendTemplatesResponse
func (client *Client) RecommendTemplatesWithContext(ctx context.Context, InstanceId *string, request *RecommendTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecommendTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinstallCollectorResponse
func (client *Client) ReinstallCollectorWithContext(ctx context.Context, ResId *string, request *ReinstallCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReinstallCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePluginResponse
func (client *Client) RemovePluginWithContext(ctx context.Context, instanceId *string, request *RemovePluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemovePluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, InstanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewLogstashResponse
func (client *Client) RenewLogstashWithContext(ctx context.Context, InstanceId *string, request *RenewLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartCollectorResponse
func (client *Client) RestartCollectorWithContext(ctx context.Context, ResId *string, request *RestartCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithContext(ctx context.Context, InstanceId *string, request *RestartInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartLogstashResponse
func (client *Client) RestartLogstashWithContext(ctx context.Context, InstanceId *string, request *RestartLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeElasticsearchTaskResponse
func (client *Client) ResumeElasticsearchTaskWithContext(ctx context.Context, InstanceId *string, request *ResumeElasticsearchTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeElasticsearchTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeLogstashTaskResponse
func (client *Client) ResumeLogstashTaskWithContext(ctx context.Context, InstanceId *string, request *ResumeLogstashTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeLogstashTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RolloverDataStreamResponse
func (client *Client) RolloverDataStreamWithContext(ctx context.Context, InstanceId *string, DataStream *string, request *RolloverDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RolloverDataStreamResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPipelinesResponse
func (client *Client) RunPipelinesWithContext(ctx context.Context, InstanceId *string, request *RunPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunPipelinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShrinkNodeResponse
func (client *Client) ShrinkNodeWithContext(ctx context.Context, InstanceId *string, request *ShrinkNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ShrinkNodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCollectorResponse
func (client *Client) StartCollectorWithContext(ctx context.Context, ResId *string, request *StartCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCollectorResponse
func (client *Client) StopCollectorWithContext(ctx context.Context, ResId *string, request *StopCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPipelinesResponse
func (client *Client) StopPipelinesWithContext(ctx context.Context, InstanceId *string, request *StopPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopPipelinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferNodeResponse
func (client *Client) TransferNodeWithContext(ctx context.Context, InstanceId *string, request *TransferNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferNodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerNetworkResponse
func (client *Client) TriggerNetworkWithContext(ctx context.Context, InstanceId *string, request *TriggerNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TriggerNetworkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TurnOffZoneResponse
func (client *Client) TurnOffZoneWithContext(ctx context.Context, instanceId *string, request *TurnOffZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TurnOffZoneResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TurnOnZoneResponse
func (client *Client) TurnOnZoneWithContext(ctx context.Context, instanceId *string, request *TurnOnZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TurnOnZoneResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallKibanaPluginResponse
func (client *Client) UninstallKibanaPluginWithContext(ctx context.Context, InstanceId *string, request *UninstallKibanaPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallKibanaPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallLogstashPluginResponse
func (client *Client) UninstallLogstashPluginWithContext(ctx context.Context, InstanceId *string, request *UninstallLogstashPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallLogstashPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallPluginResponse
func (client *Client) UninstallPluginWithContext(ctx context.Context, InstanceId *string, request *UninstallPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallPluginResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdminPasswordResponse
func (client *Client) UpdateAdminPasswordWithContext(ctx context.Context, InstanceId *string, request *UpdateAdminPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAdminPasswordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdvancedSettingResponse
func (client *Client) UpdateAdvancedSettingWithContext(ctx context.Context, InstanceId *string, request *UpdateAdvancedSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAdvancedSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliwsDictResponse
func (client *Client) UpdateAliwsDictWithContext(ctx context.Context, InstanceId *string, request *UpdateAliwsDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAliwsDictResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBlackIpsResponse
func (client *Client) UpdateBlackIpsWithContext(ctx context.Context, InstanceId *string, request *UpdateBlackIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBlackIpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectorResponse
func (client *Client) UpdateCollectorWithContext(ctx context.Context, ResId *string, request *UpdateCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCollectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectorNameResponse
func (client *Client) UpdateCollectorNameWithContext(ctx context.Context, ResId *string, request *UpdateCollectorNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCollectorNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentIndexResponse
func (client *Client) UpdateComponentIndexWithContext(ctx context.Context, InstanceId *string, name *string, request *UpdateComponentIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComponentIndexResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDescriptionResponse
func (client *Client) UpdateDescriptionWithContext(ctx context.Context, InstanceId *string, request *UpdateDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDescriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDiagnosisSettingsResponse
func (client *Client) UpdateDiagnosisSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateDiagnosisSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDiagnosisSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDictResponse
func (client *Client) UpdateDictWithContext(ctx context.Context, InstanceId *string, request *UpdateDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDictResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDynamicSettingsResponse
func (client *Client) UpdateDynamicSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateDynamicSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDynamicSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtendConfigResponse
func (client *Client) UpdateExtendConfigWithContext(ctx context.Context, InstanceId *string, request *UpdateExtendConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExtendConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtendfilesResponse
func (client *Client) UpdateExtendfilesWithContext(ctx context.Context, InstanceId *string, request *UpdateExtendfilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExtendfilesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFalconSeekResponse
func (client *Client) UpdateFalconSeekWithContext(ctx context.Context, InstanceId *string, request *UpdateFalconSeekRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFalconSeekResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotIkDictsResponse
func (client *Client) UpdateHotIkDictsWithContext(ctx context.Context, InstanceId *string, request *UpdateHotIkDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHotIkDictsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateILMPolicyResponse
func (client *Client) UpdateILMPolicyWithContext(ctx context.Context, InstanceId *string, PolicyName *string, request *UpdateILMPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateILMPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIndexTemplateResponse
func (client *Client) UpdateIndexTemplateWithContext(ctx context.Context, InstanceId *string, IndexTemplate *string, request *UpdateIndexTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIndexTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceChargeTypeResponse
func (client *Client) UpdateInstanceChargeTypeWithContext(ctx context.Context, InstanceId *string, request *UpdateInstanceChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceChargeTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceSettingsResponse
func (client *Client) UpdateInstanceSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateInstanceSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKeystoresResponse
func (client *Client) UpdateKeystoresWithContext(ctx context.Context, InstanceId *string, request *UpdateKeystoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKeystoresResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaPvlNetworkResponse
func (client *Client) UpdateKibanaPvlNetworkWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaPvlNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaPvlNetworkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaSettingsResponse
func (client *Client) UpdateKibanaSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaSsoResponse
func (client *Client) UpdateKibanaSsoWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaSsoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaSsoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaWhiteIpsResponse
func (client *Client) UpdateKibanaWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaWhiteIpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashResponse
func (client *Client) UpdateLogstashWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashChargeTypeResponse
func (client *Client) UpdateLogstashChargeTypeWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashChargeTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashDescriptionResponse
func (client *Client) UpdateLogstashDescriptionWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashDescriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashSettingsResponse
func (client *Client) UpdateLogstashSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineManagementConfigResponse
func (client *Client) UpdatePipelineManagementConfigWithContext(ctx context.Context, InstanceId *string, request *UpdatePipelineManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineManagementConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelinesResponse
func (client *Client) UpdatePipelinesWithContext(ctx context.Context, InstanceId *string, request *UpdatePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateNetworkWhiteIpsResponse
func (client *Client) UpdatePrivateNetworkWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdatePrivateNetworkWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrivateNetworkWhiteIpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicNetworkResponse
func (client *Client) UpdatePublicNetworkWithContext(ctx context.Context, InstanceId *string, request *UpdatePublicNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicNetworkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicWhiteIpsResponse
func (client *Client) UpdatePublicWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdatePublicWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicWhiteIpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReadWritePolicyResponse
func (client *Client) UpdateReadWritePolicyWithContext(ctx context.Context, InstanceId *string, request *UpdateReadWritePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateReadWritePolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSettingWithContext(ctx context.Context, InstanceId *string, request *UpdateSnapshotSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSnapshotSettingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSynonymsDictsResponse
func (client *Client) UpdateSynonymsDictsWithContext(ctx context.Context, InstanceId *string, request *UpdateSynonymsDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSynonymsDictsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithContext(ctx context.Context, InstanceId *string, TemplateName *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWhiteIpsResponse
func (client *Client) UpdateWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdateWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWhiteIpsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateXpackMonitorConfigResponse
func (client *Client) UpdateXpackMonitorConfigWithContext(ctx context.Context, InstanceId *string, request *UpdateXpackMonitorConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateXpackMonitorConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeEngineVersionResponse
func (client *Client) UpgradeEngineVersionWithContext(ctx context.Context, InstanceId *string, request *UpgradeEngineVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeEngineVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeInfoResponse
func (client *Client) UpgradeInfoWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateConnectionResponse
func (client *Client) ValidateConnectionWithContext(ctx context.Context, InstanceId *string, request *ValidateConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateConnectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateShrinkNodesResponse
func (client *Client) ValidateShrinkNodesWithContext(ctx context.Context, InstanceId *string, request *ValidateShrinkNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateShrinkNodesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateSlrPermissionResponse
func (client *Client) ValidateSlrPermissionWithContext(ctx context.Context, request *ValidateSlrPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateSlrPermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTransferableNodesResponse
func (client *Client) ValidateTransferableNodesWithContext(ctx context.Context, InstanceId *string, request *ValidateTransferableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateTransferableNodesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
