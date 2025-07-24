// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # AI搜索流式接口
//
// @param request - AiSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AiSearchResponse
func (client *Client) AiSearchWithContext(ctx context.Context, request *AiSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AiSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AiSearch"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v3/linkedRetrieval/commands/aiSearch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AiSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增强版通用搜索
//
// @param request - GenericAdvancedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenericAdvancedSearchResponse
func (client *Client) GenericAdvancedSearchWithContext(ctx context.Context, request *GenericAdvancedSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenericAdvancedSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenericAdvancedSearch"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v2/linkedRetrieval/commands/genericAdvancedSearch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenericAdvancedSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用搜索
//
// @param request - GenericSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenericSearchResponse
func (client *Client) GenericSearchWithContext(ctx context.Context, request *GenericSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenericSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableRerank) {
		query["enableRerank"] = request.EnableRerank
	}

	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.ReturnMainText) {
		query["returnMainText"] = request.ReturnMainText
	}

	if !dara.IsNil(request.ReturnMarkdownText) {
		query["returnMarkdownText"] = request.ReturnMarkdownText
	}

	if !dara.IsNil(request.ReturnSummary) {
		query["returnSummary"] = request.ReturnSummary
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenericSearch"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v2/linkedRetrieval/commands/genericSearch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenericSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 信息查询服务接口日维度使用量查询
//
// @param request - GetIqsUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIqsUsageResponse
func (client *Client) GetIqsUsageWithContext(ctx context.Context, request *GetIqsUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIqsUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["startDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIqsUsage"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/v1/iqs/usage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIqsUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通晓搜索-出海版(全球信息搜索)
//
// @param request - GlobalSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalSearchResponse
func (client *Client) GlobalSearchWithContext(ctx context.Context, request *GlobalSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GlobalSearch"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/search/global"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通晓统一搜索API
//
// @param request - UnifiedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnifiedSearchResponse
func (client *Client) UnifiedSearchWithContext(ctx context.Context, request *UnifiedSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnifiedSearchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnifiedSearch"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/search/unified"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnifiedSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
