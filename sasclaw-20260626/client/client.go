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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("sasclaw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Claw SSE Chat
//
// @param request - ChatUserSecAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatUserSecAgentResponse
func (client *Client) ChatUserSecAgentWithSSE(request *ChatUserSecAgentRequest, runtime *dara.RuntimeOptions, _yield chan *ChatUserSecAgentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatUserSecAgentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// # Claw SSE Chat
//
// @param request - ChatUserSecAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatUserSecAgentResponse
func (client *Client) ChatUserSecAgentWithOptions(request *ChatUserSecAgentRequest, runtime *dara.RuntimeOptions) (_result *ChatUserSecAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Agent) {
		body["Agent"] = request.Agent
	}

	if !dara.IsNil(request.AttachmentStagingId) {
		body["AttachmentStagingId"] = request.AttachmentStagingId
	}

	if !dara.IsNil(request.Attachments) {
		body["Attachments"] = request.Attachments
	}

	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ExecutionMode) {
		body["ExecutionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.ExtraParams) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !dara.IsNil(request.Memory) {
		body["Memory"] = request.Memory
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ResponseLanguage) {
		body["ResponseLanguage"] = request.ResponseLanguage
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Skill) {
		body["Skill"] = request.Skill
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TalkId) {
		body["TalkId"] = request.TalkId
	}

	if !dara.IsNil(request.Target) {
		body["Target"] = request.Target
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.UserInputInfo) {
		body["UserInputInfo"] = request.UserInputInfo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatUserSecAgent"),
		Version:     dara.String("2026-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &ChatUserSecAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Claw SSE Chat
//
// @param request - ChatUserSecAgentRequest
//
// @return ChatUserSecAgentResponse
func (client *Client) ChatUserSecAgent(request *ChatUserSecAgentRequest) (_result *ChatUserSecAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatUserSecAgentResponse{}
	_body, _err := client.ChatUserSecAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) chatUserSecAgentWithSSE_opYieldFunc(_yield chan *ChatUserSecAgentResponse, _yieldErr chan error, request *ChatUserSecAgentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Agent) {
		body["Agent"] = request.Agent
	}

	if !dara.IsNil(request.AttachmentStagingId) {
		body["AttachmentStagingId"] = request.AttachmentStagingId
	}

	if !dara.IsNil(request.Attachments) {
		body["Attachments"] = request.Attachments
	}

	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ExecutionMode) {
		body["ExecutionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.ExtraParams) {
		body["ExtraParams"] = request.ExtraParams
	}

	if !dara.IsNil(request.Memory) {
		body["Memory"] = request.Memory
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ResponseLanguage) {
		body["ResponseLanguage"] = request.ResponseLanguage
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Skill) {
		body["Skill"] = request.Skill
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TalkId) {
		body["TalkId"] = request.TalkId
	}

	if !dara.IsNil(request.Target) {
		body["Target"] = request.Target
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.UserInputInfo) {
		body["UserInputInfo"] = request.UserInputInfo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatUserSecAgent"),
		Version:     dara.String("2026-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
