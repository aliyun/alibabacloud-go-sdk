// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 关闭舆情产品
//
// @param request - CloseProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseProductResponse
func (client *Client) CloseProductWithContext(ctx context.Context, request *CloseProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductInstance) {
		body["productInstance"] = request.ProductInstance
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseProduct"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/closeProduct.json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制台统一代理API
//
// @param request - ConsoleApiProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleApiProxyResponse
func (client *Client) ConsoleApiProxyWithContext(ctx context.Context, request *ConsoleApiProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConsoleApiProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConsoleApiProxy"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/consoleApiProxy.json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConsoleApiProxyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ConsoleProxy is deprecated
//
// Summary:
//
// 控制台统一代理API
//
// @param request - ConsoleProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConsoleProxyResponse
func (client *Client) ConsoleProxyWithContext(ctx context.Context, request *ConsoleProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConsoleProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		body["appCode"] = request.AppCode
	}

	if !dara.IsNil(request.InterfaceName) {
		body["interfaceName"] = request.InterfaceName
	}

	if !dara.IsNil(request.ParamJson) {
		body["paramJson"] = request.ParamJson
	}

	if !dara.IsNil(request.TeamHashId) {
		body["teamHashId"] = request.TeamHashId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConsoleProxy"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/consoleProxy.json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConsoleProxyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取分析组件计算任务结果
//
// @param request - GetAnalysisTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnalysisTaskResultResponse
func (client *Client) GetAnalysisTaskResultWithContext(ctx context.Context, request *GetAnalysisTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAnalysisTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisId) {
		query["analysisId"] = request.AnalysisId
	}

	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.TeamHashId) {
		query["teamHashId"] = request.TeamHashId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAnalysisTaskResult"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/getAnalysisComponentResult.json"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAnalysisTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通舆情产品
//
// @param request - OpenProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenProductResponse
func (client *Client) OpenProductWithContext(ctx context.Context, request *OpenProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProductInstance) {
		body["productInstance"] = request.ProductInstance
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenProduct"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/openProduct.json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品开通实例列表
//
// @param request - QueryProductInstanceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductInstanceListResponse
func (client *Client) QueryProductInstanceListWithContext(ctx context.Context, request *QueryProductInstanceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryProductInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["appCode"] = request.AppCode
	}

	if !dara.IsNil(request.FromTime) {
		query["fromTime"] = request.FromTime
	}

	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.TenantUid) {
		query["tenantUid"] = request.TenantUid
	}

	if !dara.IsNil(request.ToTime) {
		query["toTime"] = request.ToTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProductInstanceList"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/queryProductInstanceList.json"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProductInstanceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询舆情文章列表
//
// @param request - QueryYuqingMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryYuqingMessageResponse
func (client *Client) QueryYuqingMessageWithContext(ctx context.Context, request *QueryYuqingMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryYuqingMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SearchCondition) {
		body["searchCondition"] = request.SearchCondition
	}

	if !dara.IsNil(request.TeamHashId) {
		body["teamHashId"] = request.TeamHashId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryYuqingMessage"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/queryYuqingMessage.json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryYuqingMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交分析组件计算任务
//
// @param request - SubmitAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAnalysisTaskResponse
func (client *Client) SubmitAnalysisTaskWithContext(ctx context.Context, request *SubmitAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalyseType) {
		body["analyseType"] = request.AnalyseType
	}

	if !dara.IsNil(request.SearchCondition) {
		body["searchCondition"] = request.SearchCondition
	}

	if !dara.IsNil(request.TeamHashId) {
		body["teamHashId"] = request.TeamHashId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAnalysisTask"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/aliyun/submitAnalysisComponent.json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
