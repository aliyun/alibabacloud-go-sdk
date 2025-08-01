// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a trace evaluation task. The system will sample some data from the user\\"s trace data based on the task\\"s configuration. Then, an LLM is used to evaluate the performance of these traces, and the evaluation results are recorded.
//
// @param tmpReq - CreateOnlineEvalTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOnlineEvalTaskResponse
func (client *Client) CreateOnlineEvalTaskWithContext(ctx context.Context, tmpReq *CreateOnlineEvalTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOnlineEvalTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateOnlineEvalTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		query["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOnlineEvalTask"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOnlineEvalTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role required for the PaiLLMTrace service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRoleWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceIdentityRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceIdentityRole"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/ServiceIdentityRole"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete an online evaluation task
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOnlineEvalTaskResponse
func (client *Client) DeleteOnlineEvalTaskWithContext(ctx context.Context, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteOnlineEvalTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOnlineEvalTask"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOnlineEvalTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Evaluates a specified piece of trace data.
//
// @param request - EvaluateTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateTraceResponse
func (client *Client) EvaluateTraceWithContext(ctx context.Context, TraceId *string, request *EvaluateTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EvaluateTraceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.EvaluationConfig) {
		body["EvaluationConfig"] = request.EvaluationConfig
	}

	if !dara.IsNil(request.EvaluationId) {
		body["EvaluationId"] = request.EvaluationId
	}

	if !dara.IsNil(request.MaxTime) {
		body["MaxTime"] = request.MaxTime
	}

	if !dara.IsNil(request.MinTime) {
		body["MinTime"] = request.MinTime
	}

	if !dara.IsNil(request.ModelConfig) {
		body["ModelConfig"] = request.ModelConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateTrace"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/eval/trace/" + dara.PercentEncode(dara.StringValue(TraceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateTraceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the content of prompt templates used for evaluation
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEvaluationTemplatesResponse
func (client *Client) GetEvaluationTemplatesWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEvaluationTemplatesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEvaluationTemplates"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/eval/templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEvaluationTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the details of an online evaluation task
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOnlineEvalTaskResponse
func (client *Client) GetOnlineEvalTaskWithContext(ctx context.Context, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOnlineEvalTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOnlineEvalTask"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOnlineEvalTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information related to the service-linked role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRoleWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceIdentityRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceIdentityRole"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/ServiceIdentityRole"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the token used in the Xtrace service and the endpoint required for uploading trace data.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXtraceTokenResponse
func (client *Client) GetXtraceTokenWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetXtraceTokenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetXtraceToken"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/XtraceToken"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetXtraceTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of results for trace evaluation. This API is used together with EvaluateTrace. EvaluateTrace starts the evaluation. ListEvalResults obtains the results.
//
// @param tmpReq - ListEvalResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvalResultsResponse
func (client *Client) ListEvalResultsWithContext(ctx context.Context, tmpReq *ListEvalResultsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEvalResultsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListEvalResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordIds) {
		request.RecordIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordIds, dara.String("RecordIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EvaluationId) {
		query["EvaluationId"] = request.EvaluationId
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

	if !dara.IsNil(request.RecordIdsShrink) {
		query["RecordIds"] = request.RecordIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvalResults"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/eval/results"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvalResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List the results of online evaluation tasks that meet the criteria
//
// @param tmpReq - ListOnlineEvalTaskResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnlineEvalTaskResultsResponse
func (client *Client) ListOnlineEvalTaskResultsWithContext(ctx context.Context, tmpReq *ListOnlineEvalTaskResultsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOnlineEvalTaskResultsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListOnlineEvalTaskResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TraceIds) {
		request.TraceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TraceIds, dara.String("TraceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EvaluationId) {
		query["EvaluationId"] = request.EvaluationId
	}

	if !dara.IsNil(request.MostRecentResultsOnly) {
		query["MostRecentResultsOnly"] = request.MostRecentResultsOnly
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TraceIdsShrink) {
		query["TraceIds"] = request.TraceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOnlineEvalTaskResults"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltaskresults"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOnlineEvalTaskResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View online evaluation tasks that meet the criteria
//
// @param request - ListOnlineEvalTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnlineEvalTasksResponse
func (client *Client) ListOnlineEvalTasksWithContext(ctx context.Context, request *ListOnlineEvalTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOnlineEvalTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxTime) {
		query["MaxTime"] = request.MaxTime
	}

	if !dara.IsNil(request.MinTime) {
		query["MinTime"] = request.MinTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOnlineEvalTasks"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOnlineEvalTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of trace data based on specified criteria.
//
// @param tmpReq - ListTracesDatasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTracesDatasResponse
func (client *Client) ListTracesDatasWithContext(ctx context.Context, tmpReq *ListTracesDatasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTracesDatasResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTracesDatasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filters) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, dara.String("Filters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SpanIds) {
		request.SpanIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpanIds, dara.String("SpanIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.TraceIds) {
		request.TraceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TraceIds, dara.String("TraceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FiltersShrink) {
		query["Filters"] = request.FiltersShrink
	}

	if !dara.IsNil(request.HasEvents) {
		query["HasEvents"] = request.HasEvents
	}

	if !dara.IsNil(request.HasStatusMessage) {
		query["HasStatusMessage"] = request.HasStatusMessage
	}

	if !dara.IsNil(request.LlmAppName) {
		query["LlmAppName"] = request.LlmAppName
	}

	if !dara.IsNil(request.MaxDuration) {
		query["MaxDuration"] = request.MaxDuration
	}

	if !dara.IsNil(request.MaxTime) {
		query["MaxTime"] = request.MaxTime
	}

	if !dara.IsNil(request.MinDuration) {
		query["MinDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.MinTime) {
		query["MinTime"] = request.MinTime
	}

	if !dara.IsNil(request.OpentelemetryCompatible) {
		query["OpentelemetryCompatible"] = request.OpentelemetryCompatible
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.OwnerSubId) {
		query["OwnerSubId"] = request.OwnerSubId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.SpanIdsShrink) {
		query["SpanIds"] = request.SpanIdsShrink
	}

	if !dara.IsNil(request.SpanName) {
		query["SpanName"] = request.SpanName
	}

	if !dara.IsNil(request.TraceIdsShrink) {
		query["TraceIds"] = request.TraceIdsShrink
	}

	if !dara.IsNil(request.TraceReduceMethod) {
		query["TraceReduceMethod"] = request.TraceReduceMethod
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTracesDatas"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/TracesDatas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTracesDatasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Stop the execution of an online evaluation task
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOnlineEvalTaskResponse
func (client *Client) StopOnlineEvalTaskWithContext(ctx context.Context, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopOnlineEvalTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopOnlineEvalTask"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopOnlineEvalTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configuration of a trace evaluation task.
//
// @param request - UpdateOnlineEvalTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOnlineEvalTaskResponse
func (client *Client) UpdateOnlineEvalTaskWithContext(ctx context.Context, TaskId *string, request *UpdateOnlineEvalTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateOnlineEvalTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EvaluationConfig) {
		body["EvaluationConfig"] = request.EvaluationConfig
	}

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.ModelConfig) {
		body["ModelConfig"] = request.ModelConfig
	}

	if !dara.IsNil(request.SamplingFrequencyMinutes) {
		body["SamplingFrequencyMinutes"] = request.SamplingFrequencyMinutes
	}

	if !dara.IsNil(request.SamplingRatio) {
		body["SamplingRatio"] = request.SamplingRatio
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOnlineEvalTask"),
		Version:     dara.String("2024-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/PAILLMTrace/onlineevaltasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOnlineEvalTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
