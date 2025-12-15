// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建语音转录异步任务
//
// @param request - CreateAudioAsrTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAudioAsrTaskResponse
func (client *Client) CreateAudioAsrTaskWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *CreateAudioAsrTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAudioAsrTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAudioAsrTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/audio-asr/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAudioAsrTaskResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档解析异步提取任务
//
// @param request - CreateDocumentAnalyzeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocumentAnalyzeTaskResponse
func (client *Client) CreateDocumentAnalyzeTaskWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *CreateDocumentAnalyzeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDocumentAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocumentAnalyzeTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/document-analyze/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocumentAnalyzeTaskResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片解析异步提取任务
//
// @param request - CreateImageAnalyzeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageAnalyzeTaskResponse
func (client *Client) CreateImageAnalyzeTaskWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *CreateImageAnalyzeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateImageAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageAnalyzeTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/image-analyze/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageAnalyzeTaskResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建语音转录异步任务
//
// @param request - CreateVideoSnapshotTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoSnapshotTaskResponse
func (client *Client) CreateVideoSnapshotTaskWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *CreateVideoSnapshotTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVideoSnapshotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoSnapshotTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-snapshot/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoSnapshotTaskResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频截帧异步提取任务状态
//
// @param request - GetAudioAsrTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioAsrTaskStatusResponse
func (client *Client) GetAudioAsrTaskStatusWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetAudioAsrTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAudioAsrTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAudioAsrTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/audio-asr/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAudioAsrTaskStatusResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档解析异步提取任务状态
//
// @param request - GetDocumentAnalyzeTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentAnalyzeTaskStatusResponse
func (client *Client) GetDocumentAnalyzeTaskStatusWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetDocumentAnalyzeTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentAnalyzeTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentAnalyzeTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/document-analyze/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentAnalyzeTaskStatusResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档相关性打分
//
// @param request - GetDocumentRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentRankResponse
func (client *Client) GetDocumentRankWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetDocumentRankRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Docs) {
		body["docs"] = request.Docs
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentRank"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/ranker/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentRankResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档切片
//
// @param request - GetDocumentSplitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentSplitResponse
func (client *Client) GetDocumentSplitWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetDocumentSplitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentSplitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentSplit"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/document-split/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentSplitResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向量微调
//
// @param request - GetEmbeddingTuningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmbeddingTuningResponse
func (client *Client) GetEmbeddingTuningWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetEmbeddingTuningRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmbeddingTuningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmbeddingTuning"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/embedding-tuning/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmbeddingTuningResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片解析异步提取任务状态
//
// @param request - GetImageAnalyzeTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageAnalyzeTaskStatusResponse
func (client *Client) GetImageAnalyzeTaskStatusWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetImageAnalyzeTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageAnalyzeTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageAnalyzeTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/image-analyze/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageAnalyzeTaskStatusResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片主体检测
//
// @param request - GetImageObjectDetectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageObjectDetectionResponse
func (client *Client) GetImageObjectDetectionWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetImageObjectDetectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageObjectDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageObjectDetection"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/image-object-detection/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageObjectDetectionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态向量化
//
// @param request - GetMultiModalEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiModalEmbeddingResponse
func (client *Client) GetMultiModalEmbeddingWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetMultiModalEmbeddingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMultiModalEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiModalEmbedding"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/multi-modal-embedding/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiModalEmbeddingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态 Reranker
//
// @param request - GetMultiModalRerankerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiModalRerankerResponse
func (client *Client) GetMultiModalRerankerWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetMultiModalRerankerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMultiModalRerankerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Docs) {
		body["docs"] = request.Docs
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiModalReranker"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/multi-modal-reranker/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiModalRerankerResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推理结果
//
// @param request - GetPredictionRequest
//
// @param headers - GetPredictionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPredictionResponse
func (client *Client) GetPredictionWithContext(ctx context.Context, deploymentId *string, request *GetPredictionRequest, headers *GetPredictionHeaders, runtime *dara.RuntimeOptions) (_result *GetPredictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Token) {
		realHeaders["Token"] = dara.String(dara.ToString(dara.StringValue(headers.Token)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrediction"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/deployments/" + dara.StringValue(deploymentId) + "/predict"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &GetPredictionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取query分析结果
//
// @param request - GetQueryAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryAnalysisResponse
func (client *Client) GetQueryAnalysisWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetQueryAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQueryAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Functions) {
		body["functions"] = request.Functions
	}

	if !dara.IsNil(request.History) {
		body["history"] = request.History
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueryAnalysis"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/query-analyze/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryAnalysisResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本向量化
//
// @param request - GetTextEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextEmbeddingResponse
func (client *Client) GetTextEmbeddingWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetTextEmbeddingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.InputType) {
		body["input_type"] = request.InputType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextEmbedding"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/text-embedding/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextEmbeddingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大模型问答
//
// @param request - GetTextGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextGenerationResponse
func (client *Client) GetTextGenerationWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetTextGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CsiLevel) {
		body["csi_level"] = request.CsiLevel
	}

	if !dara.IsNil(request.EnableSearch) {
		body["enable_search"] = request.EnableSearch
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextGeneration"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/text-generation/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextGenerationResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本稀疏向量化
//
// @param request - GetTextSparseEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextSparseEmbeddingResponse
func (client *Client) GetTextSparseEmbeddingWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetTextSparseEmbeddingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextSparseEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.InputType) {
		body["input_type"] = request.InputType
	}

	if !dara.IsNil(request.ReturnToken) {
		body["return_token"] = request.ReturnToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextSparseEmbedding"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/text-sparse-embedding/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextSparseEmbeddingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频截帧异步提取任务状态
//
// @param request - GetVideoSnapshotTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoSnapshotTaskStatusResponse
func (client *Client) GetVideoSnapshotTaskStatusWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetVideoSnapshotTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoSnapshotTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoSnapshotTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-snapshot/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoSnapshotTaskStatusResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 联网搜索
//
// @param request - GetWebSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebSearchResponse
func (client *Client) GetWebSearchWithContext(ctx context.Context, workspaceName *string, serviceId *string, request *GetWebSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWebSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["content_type"] = request.ContentType
	}

	if !dara.IsNil(request.History) {
		body["history"] = request.History
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryRewrite) {
		body["query_rewrite"] = request.QueryRewrite
	}

	if !dara.IsNil(request.TopK) {
		body["top_k"] = request.TopK
	}

	if !dara.IsNil(request.Way) {
		body["way"] = request.Way
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWebSearch"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/web-search/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWebSearchResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
