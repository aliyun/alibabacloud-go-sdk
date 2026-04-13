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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("paicopilot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # CreateChat
//
// @param request - CreateChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatResponse
func (client *Client) CreateChatWithSSE(SessionId *string, request *CreateChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *CreateChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createChatWithSSE_opYieldFunc(_yield, _yieldErr, SessionId, request, headers, runtime)
	return
}

// Summary:
//
// # CreateChat
//
// @param request - CreateChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatResponse
func (client *Client) CreateChatWithOptions(SessionId *string, request *CreateChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtraData) {
		body["ExtraData"] = request.ExtraData
	}

	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	if !dara.IsNil(request.Question) {
		body["Question"] = request.Question
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChat"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId)) + "/chats"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateChat
//
// @param request - CreateChatRequest
//
// @return CreateChatResponse
func (client *Client) CreateChat(SessionId *string, request *CreateChatRequest) (_result *CreateChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateChatResponse{}
	_body, _err := client.CreateChatWithOptions(SessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// /api/v1/sessions
//
// @param request - CreateSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionResponse
func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSession"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// /api/v1/sessions
//
// @param request - CreateSessionRequest
//
// @return CreateSessionResponse
func (client *Client) CreateSession(request *CreateSessionRequest) (_result *CreateSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSessionResponse{}
	_body, _err := client.CreateSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeleteChat
//
// @param request - DeleteChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatResponse
func (client *Client) DeleteChatWithOptions(SessionId *string, ChatId *string, request *DeleteChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChat"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId)) + "/chats/" + dara.PercentEncode(dara.StringValue(ChatId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteChat
//
// @param request - DeleteChatRequest
//
// @return DeleteChatResponse
func (client *Client) DeleteChat(SessionId *string, ChatId *string, request *DeleteChatRequest) (_result *DeleteChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteChatResponse{}
	_body, _err := client.DeleteChatWithOptions(SessionId, ChatId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeleteSession
//
// @param request - DeleteSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSessionResponse
func (client *Client) DeleteSessionWithOptions(SessionId *string, request *DeleteSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSession"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteSession
//
// @param request - DeleteSessionRequest
//
// @return DeleteSessionResponse
func (client *Client) DeleteSession(SessionId *string, request *DeleteSessionRequest) (_result *DeleteSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSessionResponse{}
	_body, _err := client.DeleteSessionWithOptions(SessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetChat
//
// @param request - GetChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatResponse
func (client *Client) GetChatWithOptions(ChatId *string, SessionId *string, request *GetChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChat"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId)) + "/chats/" + dara.PercentEncode(dara.StringValue(ChatId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetChat
//
// @param request - GetChatRequest
//
// @return GetChatResponse
func (client *Client) GetChat(ChatId *string, SessionId *string, request *GetChatRequest) (_result *GetChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatResponse{}
	_body, _err := client.GetChatWithOptions(ChatId, SessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetSession
//
// @param request - GetSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionResponse
func (client *Client) GetSessionWithOptions(SessionId *string, request *GetSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSession"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetSession
//
// @param request - GetSessionRequest
//
// @return GetSessionResponse
func (client *Client) GetSession(SessionId *string, request *GetSessionRequest) (_result *GetSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSessionResponse{}
	_body, _err := client.GetSessionWithOptions(SessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListChats
//
// @param request - ListChatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatsResponse
func (client *Client) ListChatsWithOptions(SessionId *string, request *ListChatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListChatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChats"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId)) + "/chats"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListChats
//
// @param request - ListChatsRequest
//
// @return ListChatsResponse
func (client *Client) ListChats(SessionId *string, request *ListChatsRequest) (_result *ListChatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChatsResponse{}
	_body, _err := client.ListChatsWithOptions(SessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetChat
//
// @param request - ListSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionsResponse
func (client *Client) ListSessionsWithOptions(request *ListSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessions"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetChat
//
// @param request - ListSessionsRequest
//
// @return ListSessionsResponse
func (client *Client) ListSessions(request *ListSessionsRequest) (_result *ListSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSessionsResponse{}
	_body, _err := client.ListSessionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// /api/v1/sessions
//
// @param request - SearchInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchInfoResponse
func (client *Client) SearchInfoWithOptions(request *SearchInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeBaseFilters) {
		body["KnowledgeBaseFilters"] = request.KnowledgeBaseFilters
	}

	if !dara.IsNil(request.WebFilters) {
		body["WebFilters"] = request.WebFilters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchInfo"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/searches"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// /api/v1/sessions
//
// @param request - SearchInfoRequest
//
// @return SearchInfoResponse
func (client *Client) SearchInfo(request *SearchInfoRequest) (_result *SearchInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchInfoResponse{}
	_body, _err := client.SearchInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) createChatWithSSE_opYieldFunc(_yield chan *CreateChatResponse, _yieldErr chan error, SessionId *string, request *CreateChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtraData) {
		body["ExtraData"] = request.ExtraData
	}

	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	if !dara.IsNil(request.Question) {
		body["Question"] = request.Question
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChat"),
		Version:     dara.String("2025-07-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sessions/" + dara.PercentEncode(dara.StringValue(SessionId)) + "/chats"),
		Method:      dara.String("POST"),
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
