// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFileWithContext(ctx context.Context, WorkspaceId *string, request *CreateTextFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContractId) {
		body["ContractId"] = request.ContractId
	}

	if !dara.IsNil(request.CreateTime) {
		body["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.TextFileName) {
		body["TextFileName"] = request.TextFileName
	}

	if !dara.IsNil(request.TextFileUrl) {
		body["TextFileUrl"] = request.TextFileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTextFile"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/data/textFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTextFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合同审查结果
//
// @param tmpReq - RunContractResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGenerationWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunContractResultGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runContractResultGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 生成合同审查结果
//
// @param tmpReq - RunContractResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGenerationWithContext(ctx context.Context, workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunContractResultGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunContractResultGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/result/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunContractResultGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合同审查规则
//
// @param tmpReq - RunContractRuleGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGenerationWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunContractRuleGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runContractRuleGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 生成合同审查规则
//
// @param tmpReq - RunContractRuleGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGenerationWithContext(ctx context.Context, workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunContractRuleGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunContractRuleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractRuleGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/rule/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunContractRuleGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param tmpReq - RunLegalAdviceConsultationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultationWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunLegalAdviceConsultationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runLegalAdviceConsultationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 法律咨询
//
// @param tmpReq - RunLegalAdviceConsultationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultationWithContext(ctx context.Context, workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunLegalAdviceConsultationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunLegalAdviceConsultationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLegalAdviceConsultation"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/legalAdvice/consult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 案例检索
//
// @param tmpReq - RunSearchCaseFullTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchCaseFullTextResponse
func (client *Client) RunSearchCaseFullTextWithContext(ctx context.Context, workspaceId *string, tmpReq *RunSearchCaseFullTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunSearchCaseFullTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSearchCaseFullTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterCondition) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, dara.String("filterCondition"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PageParam) {
		request.PageParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageParam, dara.String("pageParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryKeywords) {
		request.QueryKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryKeywords, dara.String("queryKeywords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SortKeyAndDirection) {
		request.SortKeyAndDirectionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortKeyAndDirection, dara.String("sortKeyAndDirection"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.FilterConditionShrink) {
		body["filterCondition"] = request.FilterConditionShrink
	}

	if !dara.IsNil(request.PageParamShrink) {
		body["pageParam"] = request.PageParamShrink
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryKeywordsShrink) {
		body["queryKeywords"] = request.QueryKeywordsShrink
	}

	if !dara.IsNil(request.ReferLevel) {
		body["referLevel"] = request.ReferLevel
	}

	if !dara.IsNil(request.SortKeyAndDirectionShrink) {
		body["sortKeyAndDirection"] = request.SortKeyAndDirectionShrink
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchCaseFullText"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/search/case/fulltext"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSearchCaseFullTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法规搜索
//
// @param tmpReq - RunSearchLawQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchLawQueryResponse
func (client *Client) RunSearchLawQueryWithContext(ctx context.Context, workspaceId *string, tmpReq *RunSearchLawQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunSearchLawQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunSearchLawQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterCondition) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, dara.String("filterCondition"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PageParam) {
		request.PageParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageParam, dara.String("pageParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryKeywords) {
		request.QueryKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryKeywords, dara.String("queryKeywords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.FilterConditionShrink) {
		body["filterCondition"] = request.FilterConditionShrink
	}

	if !dara.IsNil(request.PageParamShrink) {
		body["pageParam"] = request.PageParamShrink
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryKeywordsShrink) {
		body["queryKeywords"] = request.QueryKeywordsShrink
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSearchLawQuery"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/search/law/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSearchLawQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) runContractResultGenerationWithSSECtx_opYieldFunc(_yield chan *RunContractResultGenerationResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunContractResultGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/result/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runContractRuleGenerationWithSSECtx_opYieldFunc(_yield chan *RunContractRuleGenerationResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunContractRuleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunContractRuleGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/contract/rule/genarate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runLegalAdviceConsultationWithSSECtx_opYieldFunc(_yield chan *RunLegalAdviceConsultationResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunLegalAdviceConsultationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assistant) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, dara.String("assistant"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Thread) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, dara.String("thread"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.AssistantShrink) {
		body["assistant"] = request.AssistantShrink
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadShrink) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLegalAdviceConsultation"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/farui/legalAdvice/consult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
