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
		"cn-zhangjiakou": dara.String("iqs.cn-zhangjiakou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("iqs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// This topic describes the method to invoke the multi-stage streaming API V3 for general search by using the Alibaba Cloud OpenAPI SDK, including parameter descriptions.
//
// Description:
//
// Provides streaming results in two stages, common_search and post_retrieval, for on-demand use.
//
// General search results (common_search): The raw search results. Covers key fields such as web page title, dynamic summary, body text, source website, and publication time. Post-retrieval processing (post_retrieval): Uses a rerank model to rerank and filter the common_search results from the previous stage. The mAP metric for context relevancy improves by approximately 5%, with an additional latency of approximately 110 ms.
//
// @param request - AiSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AiSearchResponse
func (client *Client) AiSearchWithSSE(request *AiSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *AiSearchResponse, _yieldErr chan error) {
	defer close(_yield)
	client.aiSearchWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// This topic describes the method to invoke the multi-stage streaming API V3 for general search by using the Alibaba Cloud OpenAPI SDK, including parameter descriptions.
//
// Description:
//
// Provides streaming results in two stages, common_search and post_retrieval, for on-demand use.
//
// General search results (common_search): The raw search results. Covers key fields such as web page title, dynamic summary, body text, source website, and publication time. Post-retrieval processing (post_retrieval): Uses a rerank model to rerank and filter the common_search results from the previous stage. The mAP metric for context relevancy improves by approximately 5%, with an additional latency of approximately 110 ms.
//
// @param request - AiSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AiSearchResponse
func (client *Client) AiSearchWithOptions(request *AiSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AiSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the method to invoke the multi-stage streaming API V3 for general search by using the Alibaba Cloud OpenAPI SDK, including parameter descriptions.
//
// Description:
//
// Provides streaming results in two stages, common_search and post_retrieval, for on-demand use.
//
// General search results (common_search): The raw search results. Covers key fields such as web page title, dynamic summary, body text, source website, and publication time. Post-retrieval processing (post_retrieval): Uses a rerank model to rerank and filter the common_search results from the previous stage. The mAP metric for context relevancy improves by approximately 5%, with an additional latency of approximately 110 ms.
//
// @param request - AiSearchRequest
//
// @return AiSearchResponse
func (client *Client) AiSearch(request *AiSearchRequest) (_result *AiSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AiSearchResponse{}
	_body, _err := client.AiSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自然语言通用查询
//
// Description:
//
// 自然语言搜索通用接口
//
// @param request - CommonQueryBySceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommonQueryBySceneResponse
func (client *Client) CommonQueryBySceneWithOptions(request *CommonQueryBySceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CommonQueryBySceneResponse, _err error) {
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
		Action:      dara.String("CommonQueryByScene"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/amap-function-call-agent/iqs-agent-service/v2/nl/common"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CommonQueryBySceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自然语言通用查询
//
// Description:
//
// 自然语言搜索通用接口
//
// @param request - CommonQueryBySceneRequest
//
// @return CommonQueryBySceneResponse
func (client *Client) CommonQueryByScene(request *CommonQueryBySceneRequest) (_result *CommonQueryBySceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CommonQueryBySceneResponse{}
	_body, _err := client.CommonQueryBySceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes how to invoke the enhanced search operation (GenericAdvancedSearch) by using the Alibaba Cloud OpenAPI SDK, including the method and metric description. Compared with GenericSearch, GenericAdvancedSearch provides better recall of authoritative websites, with a maximum recall count of 40, delivering improved authoritativeness and data diversity. The response parameters and format of the enhanced operation are consistent with those of GenericAdvancedSearch.
//
// Description:
//
// GenericAdvancedSearch is currently in the testing phase. New specifications will be available for purchase in the future.
//
// @param request - GenericAdvancedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenericAdvancedSearchResponse
func (client *Client) GenericAdvancedSearchWithOptions(request *GenericAdvancedSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenericAdvancedSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes how to invoke the enhanced search operation (GenericAdvancedSearch) by using the Alibaba Cloud OpenAPI SDK, including the method and metric description. Compared with GenericSearch, GenericAdvancedSearch provides better recall of authoritative websites, with a maximum recall count of 40, delivering improved authoritativeness and data diversity. The response parameters and format of the enhanced operation are consistent with those of GenericAdvancedSearch.
//
// Description:
//
// GenericAdvancedSearch is currently in the testing phase. New specifications will be available for purchase in the future.
//
// @param request - GenericAdvancedSearchRequest
//
// @return GenericAdvancedSearchResponse
func (client *Client) GenericAdvancedSearch(request *GenericAdvancedSearchRequest) (_result *GenericAdvancedSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenericAdvancedSearchResponse{}
	_body, _err := client.GenericAdvancedSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a general-purpose search.
//
// @param tmpReq - GenericSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenericSearchResponse
func (client *Client) GenericSearchWithOptions(tmpReq *GenericSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenericSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenericSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdvancedParams) {
		request.AdvancedParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvancedParams, dara.String("advancedParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedParamsShrink) {
		query["advancedParams"] = request.AdvancedParamsShrink
	}

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

	if !dara.IsNil(request.ReturnRichMainBody) {
		query["returnRichMainBody"] = request.ReturnRichMainBody
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a general-purpose search.
//
// @param request - GenericSearchRequest
//
// @return GenericSearchResponse
func (client *Client) GenericSearch(request *GenericSearchRequest) (_result *GenericSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenericSearchResponse{}
	_body, _err := client.GenericSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query daily usage by dimension for the Information Query Service API
//
// @param request - GetIqsUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIqsUsageResponse
func (client *Client) GetIqsUsageWithOptions(request *GetIqsUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIqsUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerId) {
		query["callerId"] = request.CallerId
	}

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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query daily usage by dimension for the Information Query Service API
//
// @param request - GetIqsUsageRequest
//
// @return GetIqsUsageResponse
func (client *Client) GetIqsUsage(request *GetIqsUsageRequest) (_result *GetIqsUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIqsUsageResponse{}
	_body, _err := client.GetIqsUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// IQS Search - Global Edition (Global Information Search)
//
// Description:
//
// This document describes the usage and parameter specifications of GlobalSearch, the global edition of IQS Search.
//
// @param request - GlobalSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalSearchResponse
func (client *Client) GlobalSearchWithOptions(request *GlobalSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalSearchResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// IQS Search - Global Edition (Global Information Search)
//
// Description:
//
// This document describes the usage and parameter specifications of GlobalSearch, the global edition of IQS Search.
//
// @param request - GlobalSearchRequest
//
// @return GlobalSearchResponse
func (client *Client) GlobalSearch(request *GlobalSearchRequest) (_result *GlobalSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalSearchResponse{}
	_body, _err := client.GlobalSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Medical Q&A
//
// @param request - MedicalAnswerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MedicalAnswerResponse
func (client *Client) MedicalAnswerWithOptions(request *MedicalAnswerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MedicalAnswerResponse, _err error) {
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
		Action:      dara.String("MedicalAnswer"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/domain/medical/answer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MedicalAnswerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Medical Q&A
//
// @param request - MedicalAnswerRequest
//
// @return MedicalAnswerResponse
func (client *Client) MedicalAnswer(request *MedicalAnswerRequest) (_result *MedicalAnswerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MedicalAnswerResponse{}
	_body, _err := client.MedicalAnswerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Health Graph
//
// @param request - MedicalKnowledgeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MedicalKnowledgeResponse
func (client *Client) MedicalKnowledgeWithOptions(request *MedicalKnowledgeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MedicalKnowledgeResponse, _err error) {
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
		Action:      dara.String("MedicalKnowledge"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/domain/medical/know"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MedicalKnowledgeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Health Graph
//
// @param request - MedicalKnowledgeRequest
//
// @return MedicalKnowledgeResponse
func (client *Client) MedicalKnowledge(request *MedicalKnowledgeRequest) (_result *MedicalKnowledgeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MedicalKnowledgeResponse{}
	_body, _err := client.MedicalKnowledgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Multimodal search
//
// @param request - MultimodalSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultimodalSearchResponse
func (client *Client) MultimodalSearchWithOptions(request *MultimodalSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MultimodalSearchResponse, _err error) {
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
		Action:      dara.String("MultimodalSearch"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/multimodal/unified"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MultimodalSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Multimodal search
//
// @param request - MultimodalSearchRequest
//
// @return MultimodalSearchResponse
func (client *Client) MultimodalSearch(request *MultimodalSearchRequest) (_result *MultimodalSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MultimodalSearchResponse{}
	_body, _err := client.MultimodalSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the intelligent search Q&A streaming API (SearchStream). Based on retrieval-augmented generation (RAG) and large language model technologies, this API accepts natural language questions from users and automatically performs intent recognition, query rewrite, multi-source retrieval, and content generation. The API returns data using the Server-Sent Events (SSE) streaming protocol, supporting real-time output of inference status, reference sources, and token-by-token generated answers. It is suitable for AI chat and search scenarios that require low latency and high interpretability.
//
// @param request - OmniAnswerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OmniAnswerResponse
func (client *Client) OmniAnswerWithSSE(request *OmniAnswerRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *OmniAnswerResponse, _yieldErr chan error) {
	defer close(_yield)
	client.omniAnswerWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// This topic describes the intelligent search Q&A streaming API (SearchStream). Based on retrieval-augmented generation (RAG) and large language model technologies, this API accepts natural language questions from users and automatically performs intent recognition, query rewrite, multi-source retrieval, and content generation. The API returns data using the Server-Sent Events (SSE) streaming protocol, supporting real-time output of inference status, reference sources, and token-by-token generated answers. It is suitable for AI chat and search scenarios that require low latency and high interpretability.
//
// @param request - OmniAnswerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OmniAnswerResponse
func (client *Client) OmniAnswerWithOptions(request *OmniAnswerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OmniAnswerResponse, _err error) {
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
		Action:      dara.String("OmniAnswer"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/answer/omni/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &OmniAnswerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the intelligent search Q&A streaming API (SearchStream). Based on retrieval-augmented generation (RAG) and large language model technologies, this API accepts natural language questions from users and automatically performs intent recognition, query rewrite, multi-source retrieval, and content generation. The API returns data using the Server-Sent Events (SSE) streaming protocol, supporting real-time output of inference status, reference sources, and token-by-token generated answers. It is suitable for AI chat and search scenarios that require low latency and high interpretability.
//
// @param request - OmniAnswerRequest
//
// @return OmniAnswerResponse
func (client *Client) OmniAnswer(request *OmniAnswerRequest) (_result *OmniAnswerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OmniAnswerResponse{}
	_body, _err := client.OmniAnswerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Quickly retrieve HTML and parse static web page content.
//
// Description:
//
// 1. When the HTTP status code (httpcode) of the target address is less than 500, it is counted as one valid request.
//
// 2. If the content type (Content-Type) in the response header of the target address is application/pdf, the system automatically triggers PDF parsing (PDF file size must not exceed 20 MB). This operation is counted as an additional valid request.
//
// 3. Trial terms: During the trial period, the API is limited to 5 queries per second (QPS); the trial quota is 1,000 requests per 30 days.
//
// @param request - ReadPageBasicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadPageBasicResponse
func (client *Client) ReadPageBasicWithOptions(request *ReadPageBasicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReadPageBasicResponse, _err error) {
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
		Action:      dara.String("ReadPageBasic"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/readpage/basic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadPageBasicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Quickly retrieve HTML and parse static web page content.
//
// Description:
//
// 1. When the HTTP status code (httpcode) of the target address is less than 500, it is counted as one valid request.
//
// 2. If the content type (Content-Type) in the response header of the target address is application/pdf, the system automatically triggers PDF parsing (PDF file size must not exceed 20 MB). This operation is counted as an additional valid request.
//
// 3. Trial terms: During the trial period, the API is limited to 5 queries per second (QPS); the trial quota is 1,000 requests per 30 days.
//
// @param request - ReadPageBasicRequest
//
// @return ReadPageBasicResponse
func (client *Client) ReadPageBasic(request *ReadPageBasicRequest) (_result *ReadPageBasicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReadPageBasicResponse{}
	_body, _err := client.ReadPageBasicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 1. Read HTML and parse web page content through the browser sandbox environment.
//
// 2. The API starts parsing after all resources on the target page are fully loaded (the maximum waiting duration can be adjusted via the pageTimeout parameter). The overall Duration of the API call is significantly affected by the resource loading performance of the target site.
//
// Description:
//
// 1. A request is counted as valid when the HTTP status code (httpcode) of the target URL is less than 500.
//
// 2. If the content type (Content-Type) in the response header of the target URL is application/pdf, the system automatically triggers PDF parsing (PDF files up to 20 MB are supported). This operation is counted as an additional valid request.
//
// 3. Trial terms: During the trial period, the API is limited to 5 queries per second (QPS); the trial quota is 1,000 requests per 30 days.
//
// @param request - ReadPageScrapeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadPageScrapeResponse
func (client *Client) ReadPageScrapeWithOptions(request *ReadPageScrapeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReadPageScrapeResponse, _err error) {
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
		Action:      dara.String("ReadPageScrape"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/readpage/scrape"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadPageScrapeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 1. Read HTML and parse web page content through the browser sandbox environment.
//
// 2. The API starts parsing after all resources on the target page are fully loaded (the maximum waiting duration can be adjusted via the pageTimeout parameter). The overall Duration of the API call is significantly affected by the resource loading performance of the target site.
//
// Description:
//
// 1. A request is counted as valid when the HTTP status code (httpcode) of the target URL is less than 500.
//
// 2. If the content type (Content-Type) in the response header of the target URL is application/pdf, the system automatically triggers PDF parsing (PDF files up to 20 MB are supported). This operation is counted as an additional valid request.
//
// 3. Trial terms: During the trial period, the API is limited to 5 queries per second (QPS); the trial quota is 1,000 requests per 30 days.
//
// @param request - ReadPageScrapeRequest
//
// @return ReadPageScrapeResponse
func (client *Client) ReadPageScrape(request *ReadPageScrapeRequest) (_result *ReadPageScrapeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReadPageScrapeResponse{}
	_body, _err := client.ReadPageScrapeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 扫描文件
//
// @param request - ScanFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanFileResponse
func (client *Client) ScanFileWithOptions(request *ScanFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScanFileResponse, _err error) {
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
		Action:      dara.String("ScanFile"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/domain/scan/file"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScanFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 扫描文件
//
// @param request - ScanFileRequest
//
// @return ScanFileResponse
func (client *Client) ScanFile(request *ScanFileRequest) (_result *ScanFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScanFileResponse{}
	_body, _err := client.ScanFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides access to the Tongxiao unified search API, enabling quick and easy integration of web-wide general search capabilities.
//
// @param request - UnifiedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnifiedSearchResponse
func (client *Client) UnifiedSearchWithOptions(request *UnifiedSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnifiedSearchResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides access to the Tongxiao unified search API, enabling quick and easy integration of web-wide general search capabilities.
//
// @param request - UnifiedSearchRequest
//
// @return UnifiedSearchResponse
func (client *Client) UnifiedSearch(request *UnifiedSearchRequest) (_result *UnifiedSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnifiedSearchResponse{}
	_body, _err := client.UnifiedSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) aiSearchWithSSE_opYieldFunc(_yield chan *AiSearchResponse, _yieldErr chan error, request *AiSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
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
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}

func (client *Client) omniAnswerWithSSE_opYieldFunc(_yield chan *OmniAnswerResponse, _yieldErr chan error, request *OmniAnswerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OmniAnswer"),
		Version:     dara.String("2024-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/answer/omni/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.StringValue(resp.Event.Data)
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
