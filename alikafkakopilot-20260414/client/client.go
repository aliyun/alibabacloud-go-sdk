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
	client.Endpoint, _err = client.GetEndpoint(dara.String("alikafkakopilot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 智能体 stream chat
//
// @param request - KopilotChatStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KopilotChatStreamResponse
func (client *Client) KopilotChatStreamWithSSE(request *KopilotChatStreamRequest, runtime *dara.RuntimeOptions, _yield chan *KopilotChatStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.kopilotChatStreamWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// 智能体 stream chat
//
// @param request - KopilotChatStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KopilotChatStreamResponse
func (client *Client) KopilotChatStreamWithOptions(request *KopilotChatStreamRequest, runtime *dara.RuntimeOptions) (_result *KopilotChatStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KopilotChatStream"),
		Version:     dara.String("2026-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KopilotChatStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能体 stream chat
//
// @param request - KopilotChatStreamRequest
//
// @return KopilotChatStreamResponse
func (client *Client) KopilotChatStream(request *KopilotChatStreamRequest) (_result *KopilotChatStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KopilotChatStreamResponse{}
	_body, _err := client.KopilotChatStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 评价
//
// @param request - KopilotFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KopilotFeedbackResponse
func (client *Client) KopilotFeedbackWithOptions(request *KopilotFeedbackRequest, runtime *dara.RuntimeOptions) (_result *KopilotFeedbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Feedback) {
		query["Feedback"] = request.Feedback
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TurnId) {
		query["TurnId"] = request.TurnId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KopilotFeedback"),
		Version:     dara.String("2026-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KopilotFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 评价
//
// @param request - KopilotFeedbackRequest
//
// @return KopilotFeedbackResponse
func (client *Client) KopilotFeedback(request *KopilotFeedbackRequest) (_result *KopilotFeedbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KopilotFeedbackResponse{}
	_body, _err := client.KopilotFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 历史会话
//
// @param request - KopilotListConversationChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KopilotListConversationChatMessagesResponse
func (client *Client) KopilotListConversationChatMessagesWithOptions(request *KopilotListConversationChatMessagesRequest, runtime *dara.RuntimeOptions) (_result *KopilotListConversationChatMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeforeTurnId) {
		query["BeforeTurnId"] = request.BeforeTurnId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KopilotListConversationChatMessages"),
		Version:     dara.String("2026-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KopilotListConversationChatMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 历史会话
//
// @param request - KopilotListConversationChatMessagesRequest
//
// @return KopilotListConversationChatMessagesResponse
func (client *Client) KopilotListConversationChatMessages(request *KopilotListConversationChatMessagesRequest) (_result *KopilotListConversationChatMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KopilotListConversationChatMessagesResponse{}
	_body, _err := client.KopilotListConversationChatMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能体
//
// @param request - KopilotListConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KopilotListConversationsResponse
func (client *Client) KopilotListConversationsWithOptions(request *KopilotListConversationsRequest, runtime *dara.RuntimeOptions) (_result *KopilotListConversationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KopilotListConversations"),
		Version:     dara.String("2026-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KopilotListConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能体
//
// @param request - KopilotListConversationsRequest
//
// @return KopilotListConversationsResponse
func (client *Client) KopilotListConversations(request *KopilotListConversationsRequest) (_result *KopilotListConversationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KopilotListConversationsResponse{}
	_body, _err := client.KopilotListConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 状态
//
// @param request - KopilotQueryStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KopilotQueryStatusResponse
func (client *Client) KopilotQueryStatusWithOptions(request *KopilotQueryStatusRequest, runtime *dara.RuntimeOptions) (_result *KopilotQueryStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KopilotQueryStatus"),
		Version:     dara.String("2026-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KopilotQueryStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 状态
//
// @param request - KopilotQueryStatusRequest
//
// @return KopilotQueryStatusResponse
func (client *Client) KopilotQueryStatus(request *KopilotQueryStatusRequest) (_result *KopilotQueryStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KopilotQueryStatusResponse{}
	_body, _err := client.KopilotQueryStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) kopilotChatStreamWithSSE_opYieldFunc(_yield chan *KopilotChatStreamResponse, _yieldErr chan error, request *KopilotChatStreamRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KopilotChatStream"),
		Version:     dara.String("2026-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
