// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) ActivateZonesWithContext(ctx context.Context, InstanceId *string, request *ActivateZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Connects Elasticsearch clusters.
//
// @param request - AddConnectableClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddConnectableClusterResponse
func (client *Client) AddConnectableClusterWithContext(ctx context.Context, InstanceId *string, request *AddConnectableClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddConnectableClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call the AddSnapshotRepo to create a reference repository when configuring a cross-cluster OSS repository.
//
// @param request - AddSnapshotRepoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSnapshotRepoResponse
func (client *Client) AddSnapshotRepoWithContext(ctx context.Context, InstanceId *string, request *AddSnapshotRepoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddSnapshotRepoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Restores an Elasticsearch cluster that is frozen after it is released.
//
// @param request - CancelDeletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDeletionResponse
func (client *Client) CancelDeletionWithContext(ctx context.Context, InstanceId *string, request *CancelDeletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelDeletionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Restores a Logstash cluster that is frozen after it is released.
//
// @param request - CancelLogstashDeletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLogstashDeletionResponse
func (client *Client) CancelLogstashDeletionWithContext(ctx context.Context, InstanceId *string, request *CancelLogstashDeletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelLogstashDeletionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CancelTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, InstanceId *string, request *CancelTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # Capacity Planning
//
// @param request - CapacityPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CapacityPlanResponse
func (client *Client) CapacityPlanWithContext(ctx context.Context, request *CapacityPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CapacityPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 关闭实例的智能运维功能
//
// @param request - CloseDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDiagnosisResponse
func (client *Client) CloseDiagnosisWithContext(ctx context.Context, InstanceId *string, request *CloseDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseDiagnosisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CloseHttpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseHttpsResponse
func (client *Client) CloseHttpsWithContext(ctx context.Context, InstanceId *string, request *CloseHttpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseHttpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # Disable Managed Index
//
// @param request - CloseManagedIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseManagedIndexResponse
func (client *Client) CloseManagedIndexWithContext(ctx context.Context, InstanceId *string, Index *string, request *CloseManagedIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseManagedIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建收集器
//
// @param request - CreateCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCollectorResponse
func (client *Client) CreateCollectorWithContext(ctx context.Context, request *CreateCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCollectorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建Elasticsearch组合模板
//
// @param request - CreateComponentIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentIndexResponse
func (client *Client) CreateComponentIndexWithContext(ctx context.Context, InstanceId *string, name *string, request *CreateComponentIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComponentIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建数据流
//
// @param request - CreateDataStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataStreamResponse
func (client *Client) CreateDataStreamWithContext(ctx context.Context, InstanceId *string, request *CreateDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建索引生命周期策略
//
// @param request - CreateILMPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateILMPolicyResponse
func (client *Client) CreateILMPolicyWithContext(ctx context.Context, InstanceId *string, request *CreateILMPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateILMPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建索引模版
//
// @param request - CreateIndexTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexTemplateResponse
func (client *Client) CreateIndexTemplateWithContext(ctx context.Context, InstanceId *string, request *CreateIndexTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndexTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建logstash实例
//
// @param request - CreateLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogstashResponse
func (client *Client) CreateLogstashWithContext(ctx context.Context, request *CreateLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLogstashResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建Logstash管道任务
//
// @param request - CreatePipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelinesResponse
func (client *Client) CreatePipelinesWithContext(ctx context.Context, InstanceId *string, request *CreatePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePipelinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateSnapshotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithContext(ctx context.Context, InstanceId *string, request *CreateSnapshotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) CreateVpcEndpointWithContext(ctx context.Context, InstanceId *string, request *CreateVpcEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Invoke DeactivateZones to offline certain zones when there are multiple availability zones, and migrate nodes in the offline zones to other availability zones.
//
// @param request - DeactivateZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateZonesResponse
func (client *Client) DeactivateZonesWithContext(ctx context.Context, InstanceId *string, request *DeactivateZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeactivateZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a shipper.
//
// @param request - DeleteCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCollectorResponse
func (client *Client) DeleteCollectorWithContext(ctx context.Context, ResId *string, request *DeleteCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCollectorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除组合索引模板
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

// @param request - DeleteConnectedClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectedClusterResponse
func (client *Client) DeleteConnectedClusterWithContext(ctx context.Context, InstanceId *string, request *DeleteConnectedClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConnectedClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除数据流
//
// @param request - DeleteDataStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataStreamResponse
func (client *Client) DeleteDataStreamWithContext(ctx context.Context, InstanceId *string, DataStream *string, request *DeleteDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeleteDataTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataTaskResponse
func (client *Client) DeleteDataTaskWithContext(ctx context.Context, InstanceId *string, request *DeleteDataTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除历史索引模板
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
// 删除ES集群索引模版
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

// @param request - DeleteInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, InstanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) DeleteLogstashWithContext(ctx context.Context, InstanceId *string, request *DeleteLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLogstashResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a pipeline that is configured for a Logstash cluster.
//
// @param request - DeletePipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelinesResponse
func (client *Client) DeletePipelinesWithContext(ctx context.Context, InstanceId *string, request *DeletePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeleteSnapshotRepoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotRepoResponse
func (client *Client) DeleteSnapshotRepoWithContext(ctx context.Context, InstanceId *string, request *DeleteSnapshotRepoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotRepoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除服务账号vpc下的终端节点
//
// @param request - DeleteVpcEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpointWithContext(ctx context.Context, InstanceId *string, EndpointId *string, request *DeleteVpcEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # Describe APM
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApmResponse
func (client *Client) DescribeApmWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApmResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApm"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/apm/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// 查看组合索引模板详情
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

// @param request - DescribeConnectableClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConnectableClustersResponse
func (client *Client) DescribeConnectableClustersWithContext(ctx context.Context, InstanceId *string, request *DescribeConnectableClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeConnectableClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # DescribeDeprecatedTemplate
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

// @param request - DescribeDiagnoseReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnoseReportResponse
func (client *Client) DescribeDiagnoseReportWithContext(ctx context.Context, InstanceId *string, ReportId *string, request *DescribeDiagnoseReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDiagnoseReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDiagnosisSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosisSettingsResponse
func (client *Client) DescribeDiagnosisSettingsWithContext(ctx context.Context, InstanceId *string, request *DescribeDiagnosisSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosisSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取集群动态指标
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
// The name of the dictionary file.
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
// 获取Elasticsearch集群Kibana节点settings配置
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
// 查看Logstash实例详情
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
// Queries the management configurations of pipelines in a Logstash cluster.
//
// @param request - DescribePipelineManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePipelineManagementConfigResponse
func (client *Client) DescribePipelineManagementConfigWithContext(ctx context.Context, InstanceId *string, request *DescribePipelineManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribePipelineManagementConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查看备份设置
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
// Queries the configurations of the X-Pack Monitoring feature of a Logstash cluster.
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
// 触发ES实例智能诊断
//
// @param request - DiagnoseInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiagnoseInstanceResponse
func (client *Client) DiagnoseInstanceWithContext(ctx context.Context, InstanceId *string, request *DiagnoseInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DiagnoseInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 关闭kibana私网
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
// 开启v3 kibana私网
//
// @param request - EnableKibanaPvlNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableKibanaPvlNetworkResponse
func (client *Client) EnableKibanaPvlNetworkWithContext(ctx context.Context, InstanceId *string, request *EnableKibanaPvlNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableKibanaPvlNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedLogstashRestartTimeResponse
func (client *Client) EstimatedLogstashRestartTimeWithContext(ctx context.Context, InstanceId *string, request *EstimatedLogstashRestartTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EstimatedLogstashRestartTimeResponse, _err error) {
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
// Queries the estimated time that is required to restart an Elasticsearch cluster.
//
// @param request - EstimatedRestartTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstimatedRestartTimeResponse
func (client *Client) EstimatedRestartTimeWithContext(ctx context.Context, InstanceId *string, request *EstimatedRestartTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EstimatedRestartTimeResponse, _err error) {
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
// Call GetClusterDataInformation to obtain the data information about the cluster.
//
// @param request - GetClusterDataInformationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterDataInformationResponse
func (client *Client) GetClusterDataInformationWithContext(ctx context.Context, request *GetClusterDataInformationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterDataInformationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取高级监控报警自定义Grafana监控报警项
//
// @param request - GetEmonGrafanaAlertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonGrafanaAlertsResponse
func (client *Client) GetEmonGrafanaAlertsWithContext(ctx context.Context, ProjectId *string, request *GetEmonGrafanaAlertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonGrafanaAlertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取高级监控报警自定义Grafana监控大盘列表
//
// @param request - GetEmonGrafanaDashboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonGrafanaDashboardsResponse
func (client *Client) GetEmonGrafanaDashboardsWithContext(ctx context.Context, ProjectId *string, request *GetEmonGrafanaDashboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonGrafanaDashboardsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # GetEmonMonitorData
//
// @param request - GetEmonMonitorDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmonMonitorDataResponse
func (client *Client) GetEmonMonitorDataWithContext(ctx context.Context, ProjectId *string, request *GetEmonMonitorDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmonMonitorDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 统计OpenStore实例的存储容量和使用情况
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
// The maximum number of nodes.
//
// @param request - GetRegionConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionConfigurationResponse
func (client *Client) GetRegionConfigurationWithContext(ctx context.Context, request *GetRegionConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ES集群可缩容节点
//
// @param request - GetSuggestShrinkableNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuggestShrinkableNodesResponse
func (client *Client) GetSuggestShrinkableNodesWithContext(ctx context.Context, InstanceId *string, request *GetSuggestShrinkableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSuggestShrinkableNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取可数据迁移节点
//
// @param request - GetTransferableNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransferableNodesResponse
func (client *Client) GetTransferableNodesWithContext(ctx context.Context, InstanceId *string, request *GetTransferableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTransferableNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) InitializeOperationRoleWithContext(ctx context.Context, request *InitializeOperationRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitializeOperationRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) InstallAckOperatorWithContext(ctx context.Context, ClusterId *string, request *InstallAckOperatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAckOperatorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call InstallKibanaSystemPlugin to install the Kibana plug-in. The Kibana specification must be 2-Core 4 GB or higher.
//
// @param request - InstallKibanaSystemPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallKibanaSystemPluginResponse
func (client *Client) InstallKibanaSystemPluginWithContext(ctx context.Context, InstanceId *string, request *InstallKibanaSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallKibanaSystemPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) InstallLogstashSystemPluginWithContext(ctx context.Context, InstanceId *string, request *InstallLogstashSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallLogstashSystemPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call InstallSystemPlugin to install a system preset plug-in.
//
// @param request - InstallSystemPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallSystemPluginResponse
func (client *Client) InstallSystemPluginWithContext(ctx context.Context, InstanceId *string, request *InstallSystemPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallSystemPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Installs custom plug-ins that are uploaded to the Elasticsearch console.
//
// @param request - InstallUserPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallUserPluginsResponse
func (client *Client) InstallUserPluginsWithContext(ctx context.Context, InstanceId *string, request *InstallUserPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallUserPluginsResponse, _err error) {
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

// @param request - InterruptElasticsearchTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterruptElasticsearchTaskResponse
func (client *Client) InterruptElasticsearchTaskWithContext(ctx context.Context, InstanceId *string, request *InterruptElasticsearchTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InterruptElasticsearchTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// After the task is suspended, the Logstash cluster is in the suspended state.
//
// @param request - InterruptLogstashTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterruptLogstashTaskResponse
func (client *Client) InterruptLogstashTaskWithContext(ctx context.Context, InstanceId *string, request *InterruptLogstashTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InterruptLogstashTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries a list of Container Service for Kubernetes (ACK) clusters.
//
// @param request - ListAckClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAckClustersResponse
func (client *Client) ListAckClustersWithContext(ctx context.Context, request *ListAckClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAckClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// > When you install a shipper on an ACK cluster, you must specify a namespace. You can call this operation to query all namespaces in the ACK cluster, and select a namespace based on your business requirements.
//
// @param request - ListAckNamespacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAckNamespacesResponse
func (client *Client) ListAckNamespacesWithContext(ctx context.Context, ClusterId *string, request *ListAckNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAckNamespacesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 变更记录 变更详情
//
// @param request - ListActionRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionRecordsResponse
func (client *Client) ListActionRecordsWithContext(ctx context.Context, InstanceId *string, request *ListActionRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListActionRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// es-cn-tl32cpgwa002l\\*\\*\\*\\*
//
// @param request - ListAllNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllNodeResponse
func (client *Client) ListAllNodeWithContext(ctx context.Context, InstanceId *string, request *ListAllNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 实例可添加的OSS引用仓库
//
// @param request - ListAlternativeSnapshotReposRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlternativeSnapshotReposResponse
func (client *Client) ListAlternativeSnapshotReposWithContext(ctx context.Context, InstanceId *string, request *ListAlternativeSnapshotReposRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlternativeSnapshotReposResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ListApm
//
// @param request - ListApmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApmResponse
func (client *Client) ListApmWithContext(ctx context.Context, request *ListApmRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
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
		Action:      dara.String("ListApm"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/apm"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Queries shippers.
//
// @param request - ListCollectorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectorsResponse
func (client *Client) ListCollectorsWithContext(ctx context.Context, request *ListCollectorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCollectorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ES集群组合索引列表
//
// @param request - ListComponentIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentIndicesResponse
func (client *Client) ListComponentIndicesWithContext(ctx context.Context, InstanceId *string, request *ListComponentIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComponentIndicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取与当前实例进行网络互通的实例列表
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
// 查询数据流
//
// @param request - ListDataStreamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataStreamsResponse
func (client *Client) ListDataStreamsWithContext(ctx context.Context, InstanceId *string, request *ListDataStreamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataStreamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the default configuration files of shippers.
//
// @param request - ListDefaultCollectorConfigurationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDefaultCollectorConfigurationsResponse
func (client *Client) ListDefaultCollectorConfigurationsWithContext(ctx context.Context, request *ListDefaultCollectorConfigurationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDefaultCollectorConfigurationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ListDeprecatedTemplates
//
// @param request - ListDeprecatedTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeprecatedTemplatesResponse
func (client *Client) ListDeprecatedTemplatesWithContext(ctx context.Context, InstanceId *string, request *ListDeprecatedTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeprecatedTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the indexes for health diagnosis performed on an Elasticsearch cluster.
//
// @param request - ListDiagnoseIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseIndicesResponse
func (client *Client) ListDiagnoseIndicesWithContext(ctx context.Context, InstanceId *string, request *ListDiagnoseIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseIndicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取集群诊断报告列表
//
// @param request - ListDiagnoseReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseReportResponse
func (client *Client) ListDiagnoseReportWithContext(ctx context.Context, InstanceId *string, request *ListDiagnoseReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the IDs of the historical intelligent O&M reports of an Elasticsearch cluster.
//
// @param request - ListDiagnoseReportIdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnoseReportIdsResponse
func (client *Client) ListDiagnoseReportIdsWithContext(ctx context.Context, InstanceId *string, request *ListDiagnoseReportIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnoseReportIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// The diagnostic item is used to check whether data write requests of a cluster are accumulated. If data write requests are accumulated, a bulk rejection occurs. This may cause data loss and severely consume system resources.
//
// @param request - ListDiagnosisItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosisItemsResponse
func (client *Client) ListDiagnosisItemsWithContext(ctx context.Context, request *ListDiagnosisItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnosisItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDictInformationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictInformationResponse
func (client *Client) ListDictInformationWithContext(ctx context.Context, InstanceId *string, request *ListDictInformationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictInformationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details of a specified type of dictionary.
//
// @param request - ListDictsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictsResponse
func (client *Client) ListDictsWithContext(ctx context.Context, InstanceId *string, request *ListDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDictsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) ListEcsInstancesWithContext(ctx context.Context, request *ListEcsInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcsInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the driver files of a Logstash cluster.
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

// @param request - ListILMPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListILMPoliciesResponse
func (client *Client) ListILMPoliciesWithContext(ctx context.Context, InstanceId *string, request *ListILMPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListILMPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ListIndexTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexTemplatesResponse
func (client *Client) ListIndexTemplatesWithContext(ctx context.Context, InstanceId *string, request *ListIndexTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询Elasticsearch实例列表
//
// @param request - ListInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 集群触发的硬件运维事件列表
//
// @param tmpReq - ListInstanceHistoryEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHistoryEventsResponse
func (client *Client) ListInstanceHistoryEventsWithContext(ctx context.Context, tmpReq *ListInstanceHistoryEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceHistoryEventsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 获取当前实例先特定的索引列表
//
// @param request - ListInstanceIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceIndicesResponse
func (client *Client) ListInstanceIndicesWithContext(ctx context.Context, InstanceId *string, request *ListInstanceIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceIndicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries a list of Kibana plug-ins.
//
// @param request - ListKibanaPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKibanaPluginsResponse
func (client *Client) ListKibanaPluginsWithContext(ctx context.Context, InstanceId *string, request *ListKibanaPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKibanaPluginsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询kibana私网连接信息
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
// # Logstash集群列表
//
// @param request - ListLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashResponse
func (client *Client) ListLogstashWithContext(ctx context.Context, request *ListLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取Logstash日志
//
// @param request - ListLogstashLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashLogResponse
func (client *Client) ListLogstashLogWithContext(ctx context.Context, InstanceId *string, request *ListLogstashLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # Logstash插件列表
//
// @param request - ListLogstashPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogstashPluginsResponse
func (client *Client) ListLogstashPluginsWithContext(ctx context.Context, InstanceId *string, request *ListLogstashPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogstashPluginsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the statuses of Elastic Compute Service (ECS) instances on which a shipper is installed.
//
// @param request - ListNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, ResId *string, request *ListNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ListPipeline
//
// @param request - ListPipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineResponse
func (client *Client) ListPipelineWithContext(ctx context.Context, InstanceId *string, request *ListPipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// The error message returned.
//
// @param request - ListPipelineIdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineIdsResponse
func (client *Client) ListPipelineIdsWithContext(ctx context.Context, InstanceId *string, request *ListPipelineIdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ES系统插件列表
//
// @param request - ListPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsResponse
func (client *Client) ListPluginsWithContext(ctx context.Context, InstanceId *string, request *ListPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查看Elasticsearch集群各种类型的日志
//
// @param request - ListSearchLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchLogResponse
func (client *Client) ListSearchLogWithContext(ctx context.Context, InstanceId *string, request *ListSearchLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) ListShardRecoveriesWithContext(ctx context.Context, InstanceId *string, request *ListShardRecoveriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShardRecoveriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取跨集群索引仓库列表
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
// 查看资源和标签关系
//
// @param request - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查看所有已常见的标签
//
// @param request - ListTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the statuses of endpoints in the virtual private cloud (VPC) within the Elasticsearch service account.
//
// @param request - ListVpcEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpointsWithContext(ctx context.Context, InstanceId *string, request *ListVpcEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) MigrateToOtherZoneWithContext(ctx context.Context, InstanceId *string, request *MigrateToOtherZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Changes the Elastic Compute Service (ECS) instances on which a shipper is installed.
//
// @param request - ModifyDeployMachineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeployMachineResponse
func (client *Client) ModifyDeployMachineWithContext(ctx context.Context, ResId *string, request *ModifyDeployMachineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDeployMachineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyElastictaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElastictaskResponse
func (client *Client) ModifyElastictaskWithContext(ctx context.Context, InstanceId *string, request *ModifyElastictaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyElastictaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) ModifyInstanceMaintainTimeWithContext(ctx context.Context, InstanceId *string, request *ModifyInstanceMaintainTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) ModifyWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *ModifyWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyWhiteIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Migrates an Elasticsearch cluster to a specified resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithContext(ctx context.Context, InstanceId *string, request *MoveResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - OpenDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDiagnosisResponse
func (client *Client) OpenDiagnosisWithContext(ctx context.Context, InstanceId *string, request *OpenDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenDiagnosisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) OpenHttpsWithContext(ctx context.Context, InstanceId *string, request *OpenHttpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenHttpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # PostEmonTryAlarmRule
//
// @param request - PostEmonTryAlarmRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEmonTryAlarmRuleResponse
func (client *Client) PostEmonTryAlarmRuleWithContext(ctx context.Context, ProjectId *string, AlarmGroupId *string, request *PostEmonTryAlarmRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PostEmonTryAlarmRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - RecommendTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecommendTemplatesResponse
func (client *Client) RecommendTemplatesWithContext(ctx context.Context, InstanceId *string, request *RecommendTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecommendTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Installs a shipper that failed to be installed when you create the shipper.
//
// @param request - ReinstallCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinstallCollectorResponse
func (client *Client) ReinstallCollectorWithContext(ctx context.Context, ResId *string, request *ReinstallCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReinstallCollectorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # RemoveApm
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApmResponse
func (client *Client) RemoveApmWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveApmResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApm"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/apm/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, InstanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Renews a Logstash cluster.
//
// @param request - RenewLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewLogstashResponse
func (client *Client) RenewLogstashWithContext(ctx context.Context, InstanceId *string, request *RenewLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewLogstashResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Restarts a shipper.
//
// @param request - RestartCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartCollectorResponse
func (client *Client) RestartCollectorWithContext(ctx context.Context, ResId *string, request *RestartCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartCollectorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) RestartInstanceWithContext(ctx context.Context, InstanceId *string, request *RestartInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 重启Logstash集群
//
// @param request - RestartLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartLogstashResponse
func (client *Client) RestartLogstashWithContext(ctx context.Context, InstanceId *string, request *RestartLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartLogstashResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ResumeElasticsearchTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeElasticsearchTaskResponse
func (client *Client) ResumeElasticsearchTaskWithContext(ctx context.Context, InstanceId *string, request *ResumeElasticsearchTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeElasticsearchTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Resumes a change task of a Logstash cluster. After the task is resumed, the Logstash cluster is in the activating state.
//
// @param request - ResumeLogstashTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeLogstashTaskResponse
func (client *Client) ResumeLogstashTaskWithContext(ctx context.Context, InstanceId *string, request *ResumeLogstashTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeLogstashTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 滚动数据流，生成新索引
//
// @param request - RolloverDataStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RolloverDataStreamResponse
func (client *Client) RolloverDataStreamWithContext(ctx context.Context, InstanceId *string, DataStream *string, request *RolloverDataStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RolloverDataStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Runs pipelines in a Logstash cluster.
//
// @param request - RunPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPipelinesResponse
func (client *Client) RunPipelinesWithContext(ctx context.Context, InstanceId *string, request *RunPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunPipelinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # ES集群缩节点
//
// @param request - ShrinkNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShrinkNodeResponse
func (client *Client) ShrinkNodeWithContext(ctx context.Context, InstanceId *string, request *ShrinkNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ShrinkNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # StartApm
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartApmResponse
func (client *Client) StartApmWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartApmResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartApm"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/apm/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartApmResponse{}
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # StopApm
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopApmResponse
func (client *Client) StopApmWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopApmResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopApm"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/apm/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopApmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCollectorResponse
func (client *Client) StopCollectorWithContext(ctx context.Context, ResId *string, request *StopCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopCollectorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Stops pipelines in a Logstash cluster.
//
// @param request - StopPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPipelinesResponse
func (client *Client) StopPipelinesWithContext(ctx context.Context, InstanceId *string, request *StopPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopPipelinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// The information about the clusters and tags.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 缩节点，数据迁移
//
// @param request - TransferNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferNodeResponse
func (client *Client) TransferNodeWithContext(ctx context.Context, InstanceId *string, request *TransferNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 开关ES集群及Kibana节点公私网访问
//
// @param request - TriggerNetworkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerNetworkResponse
func (client *Client) TriggerNetworkWithContext(ctx context.Context, InstanceId *string, request *TriggerNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TriggerNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call the UninstallKibanaPlugin to uninstall the Kibana plug-in.
//
// @param request - UninstallKibanaPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallKibanaPluginResponse
func (client *Client) UninstallKibanaPluginWithContext(ctx context.Context, InstanceId *string, request *UninstallKibanaPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallKibanaPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 卸载Logstash实例已安装的插件
//
// @param request - UninstallLogstashPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallLogstashPluginResponse
func (client *Client) UninstallLogstashPluginWithContext(ctx context.Context, InstanceId *string, request *UninstallLogstashPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallLogstashPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call UninstallPlugin to uninstall the preset plug-in.
//
// @param request - UninstallPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallPluginResponse
func (client *Client) UninstallPluginWithContext(ctx context.Context, InstanceId *string, request *UninstallPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateAdminPasswordWithContext(ctx context.Context, InstanceId *string, request *UpdateAdminPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAdminPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call UpdateAdvancedSetting to change the garbage collector configuration for the specified instance.
//
// @param request - UpdateAdvancedSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdvancedSettingResponse
func (client *Client) UpdateAdvancedSettingWithContext(ctx context.Context, InstanceId *string, request *UpdateAdvancedSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAdvancedSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateAliwsDictWithContext(ctx context.Context, InstanceId *string, request *UpdateAliwsDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAliwsDictResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// Summary:
//
// 修改APM实规格配置
//
// @param request - UpdateApmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApmResponse
func (client *Client) UpdateApmWithContext(ctx context.Context, instanceId *string, request *UpdateApmRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.OutputES) {
		body["outputES"] = request.OutputES
	}

	if !dara.IsNil(request.OutputESPassword) {
		body["outputESPassword"] = request.OutputESPassword
	}

	if !dara.IsNil(request.OutputESUserName) {
		body["outputESUserName"] = request.OutputESUserName
	}

	if !dara.IsNil(request.Token) {
		body["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApm"),
		Version:     dara.String("2017-06-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/apm/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApmResponse{}
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
// 修改ES实例访问黑名单，已废弃
//
// @param request - UpdateBlackIpsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBlackIpsResponse
func (client *Client) UpdateBlackIpsWithContext(ctx context.Context, InstanceId *string, request *UpdateBlackIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBlackIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Updates the configurations of a shipper.
//
// @param request - UpdateCollectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectorResponse
func (client *Client) UpdateCollectorWithContext(ctx context.Context, ResId *string, request *UpdateCollectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCollectorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Changes the name of a shipper.
//
// @param request - UpdateCollectorNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCollectorNameResponse
func (client *Client) UpdateCollectorNameWithContext(ctx context.Context, ResId *string, request *UpdateCollectorNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCollectorNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改ES集群动态索引
//
// @param request - UpdateComponentIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentIndexResponse
func (client *Client) UpdateComponentIndexWithContext(ctx context.Context, InstanceId *string, name *string, request *UpdateComponentIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComponentIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改elasticsearch实例名称名称
//
// @param request - UpdateDescriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDescriptionResponse
func (client *Client) UpdateDescriptionWithContext(ctx context.Context, InstanceId *string, request *UpdateDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call UpdateDiagnosisSettings to update the instance of intelligent operation&maintenance (O&M) scene settings.
//
// @param request - UpdateDiagnosisSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDiagnosisSettingsResponse
func (client *Client) UpdateDiagnosisSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateDiagnosisSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDiagnosisSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateDictWithContext(ctx context.Context, InstanceId *string, request *UpdateDictRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDictResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改集群动态配置
//
// @param request - UpdateDynamicSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDynamicSettingsResponse
func (client *Client) UpdateDynamicSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateDynamicSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDynamicSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - UpdateExtendConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtendConfigResponse
func (client *Client) UpdateExtendConfigWithContext(ctx context.Context, InstanceId *string, request *UpdateExtendConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExtendConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateExtendfilesWithContext(ctx context.Context, InstanceId *string, request *UpdateExtendfilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExtendfilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateHotIkDictsWithContext(ctx context.Context, InstanceId *string, request *UpdateHotIkDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHotIkDictsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改ES集群索引生命周期策略
//
// @param request - UpdateILMPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateILMPolicyResponse
func (client *Client) UpdateILMPolicyWithContext(ctx context.Context, InstanceId *string, PolicyName *string, request *UpdateILMPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateILMPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改ES集群索引模版配置
//
// @param request - UpdateIndexTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIndexTemplateResponse
func (client *Client) UpdateIndexTemplateWithContext(ctx context.Context, InstanceId *string, IndexTemplate *string, request *UpdateIndexTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIndexTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateInstanceWithContext(ctx context.Context, InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceChargeTypeResponse
func (client *Client) UpdateInstanceChargeTypeWithContext(ctx context.Context, InstanceId *string, request *UpdateInstanceChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceChargeTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceSettingsResponse
func (client *Client) UpdateInstanceSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateInstanceSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaPvlNetworkResponse
func (client *Client) UpdateKibanaPvlNetworkWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaPvlNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaPvlNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKibanaSettingsResponse
func (client *Client) UpdateKibanaSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateKibanaWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdateKibanaWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKibanaWhiteIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改Logstash节点规格磁盘配置
//
// @param request - UpdateLogstashRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashResponse
func (client *Client) UpdateLogstashWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Switches the billing method of a Logstash cluster from pay-as-you-go to subscription.
//
// @param request - UpdateLogstashChargeTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogstashChargeTypeResponse
func (client *Client) UpdateLogstashChargeTypeWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashChargeTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateLogstashDescriptionWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateLogstashSettingsWithContext(ctx context.Context, InstanceId *string, request *UpdateLogstashSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLogstashSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改Logstash管道配置
//
// @param request - UpdatePipelineManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineManagementConfigResponse
func (client *Client) UpdatePipelineManagementConfigWithContext(ctx context.Context, InstanceId *string, request *UpdatePipelineManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineManagementConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Updates a pipeline of a Logstash cluster.
//
// @param request - UpdatePipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelinesResponse
func (client *Client) UpdatePipelinesWithContext(ctx context.Context, InstanceId *string, request *UpdatePipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdatePrivateNetworkWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdatePrivateNetworkWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePrivateNetworkWhiteIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdatePublicNetworkWithContext(ctx context.Context, InstanceId *string, request *UpdatePublicNetworkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdatePublicWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdatePublicWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicWhiteIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更改ES集群高可用策略
//
// @param request - UpdateReadWritePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReadWritePolicyResponse
func (client *Client) UpdateReadWritePolicyWithContext(ctx context.Context, InstanceId *string, request *UpdateReadWritePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateReadWritePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Call UpdateSnapshotSetting to update the data backup configuration of the specified instance.
//
// @param request - UpdateSnapshotSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSettingWithContext(ctx context.Context, InstanceId *string, request *UpdateSnapshotSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSnapshotSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateSynonymsDictsWithContext(ctx context.Context, InstanceId *string, request *UpdateSynonymsDictsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSynonymsDictsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithContext(ctx context.Context, InstanceId *string, TemplateName *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpdateWhiteIpsWithContext(ctx context.Context, InstanceId *string, request *UpdateWhiteIpsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWhiteIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改Logstash实例的X-Pack监控报警配置。
//
// @param request - UpdateXpackMonitorConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateXpackMonitorConfigResponse
func (client *Client) UpdateXpackMonitorConfigWithContext(ctx context.Context, InstanceId *string, request *UpdateXpackMonitorConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateXpackMonitorConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) UpgradeEngineVersionWithContext(ctx context.Context, InstanceId *string, request *UpgradeEngineVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeEngineVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) ValidateConnectionWithContext(ctx context.Context, InstanceId *string, request *ValidateConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 校验缩节点合法性
//
// @param request - ValidateShrinkNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateShrinkNodesResponse
func (client *Client) ValidateShrinkNodesWithContext(ctx context.Context, InstanceId *string, request *ValidateShrinkNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateShrinkNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ValidateSlrPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateSlrPermissionResponse
func (client *Client) ValidateSlrPermissionWithContext(ctx context.Context, request *ValidateSlrPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateSlrPermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 缩节点校验数据迁移合法性
//
// @param request - ValidateTransferableNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTransferableNodesResponse
func (client *Client) ValidateTransferableNodesWithContext(ctx context.Context, InstanceId *string, request *ValidateTransferableNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateTransferableNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
