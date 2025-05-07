// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SendOpsMessageToTerminalRequest struct {
	// example:
	//
	// text
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// example:
	//
	// cn-hangzhou
	OfficeRegionId *string `json:"OfficeRegionId,omitempty" xml:"OfficeRegionId,omitempty"`
	// example:
	//
	// reboot
	OpsAction *string `json:"OpsAction,omitempty" xml:"OpsAction,omitempty"`
	// example:
	//
	// AAAAAAAAAAAA
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// example:
	//
	// true
	WaitForAck *bool `json:"WaitForAck,omitempty" xml:"WaitForAck,omitempty"`
	// example:
	//
	// false
	WaitForMsg *bool `json:"WaitForMsg,omitempty" xml:"WaitForMsg,omitempty"`
}

func (s SendOpsMessageToTerminalRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOpsMessageToTerminalRequest) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalRequest) SetMessageBody(v string) *SendOpsMessageToTerminalRequest {
	s.MessageBody = &v
	return s
}

func (s *SendOpsMessageToTerminalRequest) SetOfficeRegionId(v string) *SendOpsMessageToTerminalRequest {
	s.OfficeRegionId = &v
	return s
}

func (s *SendOpsMessageToTerminalRequest) SetOpsAction(v string) *SendOpsMessageToTerminalRequest {
	s.OpsAction = &v
	return s
}

func (s *SendOpsMessageToTerminalRequest) SetSerialNo(v string) *SendOpsMessageToTerminalRequest {
	s.SerialNo = &v
	return s
}

func (s *SendOpsMessageToTerminalRequest) SetWaitForAck(v bool) *SendOpsMessageToTerminalRequest {
	s.WaitForAck = &v
	return s
}

func (s *SendOpsMessageToTerminalRequest) SetWaitForMsg(v bool) *SendOpsMessageToTerminalRequest {
	s.WaitForMsg = &v
	return s
}

type SendOpsMessageToTerminalResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2F015A9B-C457-5E5A-A58D-215CCEBC4A80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendOpsMessageToTerminalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOpsMessageToTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalResponseBody) SetCode(v string) *SendOpsMessageToTerminalResponseBody {
	s.Code = &v
	return s
}

func (s *SendOpsMessageToTerminalResponseBody) SetMessage(v string) *SendOpsMessageToTerminalResponseBody {
	s.Message = &v
	return s
}

func (s *SendOpsMessageToTerminalResponseBody) SetRequestId(v string) *SendOpsMessageToTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOpsMessageToTerminalResponseBody) SetSuccess(v bool) *SendOpsMessageToTerminalResponseBody {
	s.Success = &v
	return s
}

type SendOpsMessageToTerminalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendOpsMessageToTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendOpsMessageToTerminalResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOpsMessageToTerminalResponse) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalResponse) SetHeaders(v map[string]*string) *SendOpsMessageToTerminalResponse {
	s.Headers = v
	return s
}

func (s *SendOpsMessageToTerminalResponse) SetStatusCode(v int32) *SendOpsMessageToTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOpsMessageToTerminalResponse) SetBody(v *SendOpsMessageToTerminalResponseBody) *SendOpsMessageToTerminalResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("wuyingsolutionframework"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 发送运维消息
//
// @param request - SendOpsMessageToTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOpsMessageToTerminalResponse
func (client *Client) SendOpsMessageToTerminalWithOptions(request *SendOpsMessageToTerminalRequest, runtime *util.RuntimeOptions) (_result *SendOpsMessageToTerminalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageBody)) {
		query["MessageBody"] = request.MessageBody
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeRegionId)) {
		query["OfficeRegionId"] = request.OfficeRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpsAction)) {
		query["OpsAction"] = request.OpsAction
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		query["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.WaitForAck)) {
		query["WaitForAck"] = request.WaitForAck
	}

	if !tea.BoolValue(util.IsUnset(request.WaitForMsg)) {
		query["WaitForMsg"] = request.WaitForMsg
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendOpsMessageToTerminal"),
		Version:     tea.String("2023-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendOpsMessageToTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送运维消息
//
// @param request - SendOpsMessageToTerminalRequest
//
// @return SendOpsMessageToTerminalResponse
func (client *Client) SendOpsMessageToTerminal(request *SendOpsMessageToTerminalRequest) (_result *SendOpsMessageToTerminalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendOpsMessageToTerminalResponse{}
	_body, _err := client.SendOpsMessageToTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
