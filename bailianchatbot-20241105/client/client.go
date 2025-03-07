// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SseChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-dDmF3jcdVf
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// TIMEOUT
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// uid129312098593
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// example:
	//
	// 15e04f27-acd7-489d-872f-1d68f7535e33
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	Utterance   *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	VendorParam *string `json:"VendorParam,omitempty" xml:"VendorParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-g7jspxljq8k909h3
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SseChatRequest) String() string {
	return tea.Prettify(s)
}

func (s SseChatRequest) GoString() string {
	return s.String()
}

func (s *SseChatRequest) SetAppId(v string) *SseChatRequest {
	s.AppId = &v
	return s
}

func (s *SseChatRequest) SetCommand(v string) *SseChatRequest {
	s.Command = &v
	return s
}

func (s *SseChatRequest) SetSenderId(v string) *SseChatRequest {
	s.SenderId = &v
	return s
}

func (s *SseChatRequest) SetSenderNick(v string) *SseChatRequest {
	s.SenderNick = &v
	return s
}

func (s *SseChatRequest) SetSessionId(v string) *SseChatRequest {
	s.SessionId = &v
	return s
}

func (s *SseChatRequest) SetUtterance(v string) *SseChatRequest {
	s.Utterance = &v
	return s
}

func (s *SseChatRequest) SetVendorParam(v string) *SseChatRequest {
	s.VendorParam = &v
	return s
}

func (s *SseChatRequest) SetWorkspaceId(v string) *SseChatRequest {
	s.WorkspaceId = &v
	return s
}

type SseChatResponseBody struct {
	// example:
	//
	// 200
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 30D68C4C-4512-58A7-A328-568015B39A02
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SseChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SseChatResponseBody) GoString() string {
	return s.String()
}

func (s *SseChatResponseBody) SetCode(v string) *SseChatResponseBody {
	s.Code = &v
	return s
}

func (s *SseChatResponseBody) SetData(v interface{}) *SseChatResponseBody {
	s.Data = v
	return s
}

func (s *SseChatResponseBody) SetMessage(v string) *SseChatResponseBody {
	s.Message = &v
	return s
}

func (s *SseChatResponseBody) SetRequestId(v string) *SseChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *SseChatResponseBody) SetSuccess(v bool) *SseChatResponseBody {
	s.Success = &v
	return s
}

type SseChatResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SseChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SseChatResponse) String() string {
	return tea.Prettify(s)
}

func (s SseChatResponse) GoString() string {
	return s.String()
}

func (s *SseChatResponse) SetHeaders(v map[string]*string) *SseChatResponse {
	s.Headers = v
	return s
}

func (s *SseChatResponse) SetStatusCode(v int32) *SseChatResponse {
	s.StatusCode = &v
	return s
}

func (s *SseChatResponse) SetBody(v *SseChatResponseBody) *SseChatResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailianchatbot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// SSE问答接口
//
// @param request - SseChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SseChatResponse
func (client *Client) SseChatWithOptions(request *SseChatRequest, runtime *util.RuntimeOptions) (_result *SseChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		query["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNick)) {
		query["SenderNick"] = request.SenderNick
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Utterance)) {
		query["Utterance"] = request.Utterance
	}

	if !tea.BoolValue(util.IsUnset(request.VendorParam)) {
		query["VendorParam"] = request.VendorParam
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SseChat"),
		Version:     tea.String("2024-11-05"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SseChatResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SseChatResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// SSE问答接口
//
// @param request - SseChatRequest
//
// @return SseChatResponse
func (client *Client) SseChat(request *SseChatRequest) (_result *SseChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SseChatResponse{}
	_body, _err := client.SseChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
