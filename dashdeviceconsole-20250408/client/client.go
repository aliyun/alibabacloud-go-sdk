// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetPromptRequest struct {
	// example:
	//
	// chat
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s GetPromptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPromptRequest) GoString() string {
	return s.String()
}

func (s *GetPromptRequest) SetGroupId(v string) *GetPromptRequest {
	s.GroupId = &v
	return s
}

type GetPromptResponseBody struct {
	// example:
	//
	// {$PromptContent}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BDA59622-2114-5F68-A530-3FCACAF0F04F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GetPromptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPromptResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBody) SetData(v interface{}) *GetPromptResponseBody {
	s.Data = v
	return s
}

func (s *GetPromptResponseBody) SetErrorCode(v string) *GetPromptResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPromptResponseBody) SetMessage(v string) *GetPromptResponseBody {
	s.Message = &v
	return s
}

func (s *GetPromptResponseBody) SetRequestId(v string) *GetPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptResponseBody) SetStatusCode(v int32) *GetPromptResponseBody {
	s.StatusCode = &v
	return s
}

type GetPromptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPromptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPromptResponse) GoString() string {
	return s.String()
}

func (s *GetPromptResponse) SetHeaders(v map[string]*string) *GetPromptResponse {
	s.Headers = v
	return s
}

func (s *GetPromptResponse) SetStatusCode(v int32) *GetPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPromptResponse) SetBody(v *GetPromptResponseBody) *GetPromptResponse {
	s.Body = v
	return s
}

type PushPromptRequest struct {
	// example:
	//
	// chat
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// {}
	PromptContent *string `json:"promptContent,omitempty" xml:"promptContent,omitempty"`
}

func (s PushPromptRequest) String() string {
	return tea.Prettify(s)
}

func (s PushPromptRequest) GoString() string {
	return s.String()
}

func (s *PushPromptRequest) SetGroupId(v string) *PushPromptRequest {
	s.GroupId = &v
	return s
}

func (s *PushPromptRequest) SetPromptContent(v string) *PushPromptRequest {
	s.PromptContent = &v
	return s
}

type PushPromptResponseBody struct {
	// example:
	//
	// True
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5090DEE5-E7DB-59A8-B712-28918D3AAA8A
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StatusCode *int32  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PushPromptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushPromptResponseBody) GoString() string {
	return s.String()
}

func (s *PushPromptResponseBody) SetData(v interface{}) *PushPromptResponseBody {
	s.Data = v
	return s
}

func (s *PushPromptResponseBody) SetErrorCode(v string) *PushPromptResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PushPromptResponseBody) SetMessage(v string) *PushPromptResponseBody {
	s.Message = &v
	return s
}

func (s *PushPromptResponseBody) SetRequestId(v string) *PushPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushPromptResponseBody) SetStatusCode(v int32) *PushPromptResponseBody {
	s.StatusCode = &v
	return s
}

type PushPromptResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushPromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushPromptResponse) String() string {
	return tea.Prettify(s)
}

func (s PushPromptResponse) GoString() string {
	return s.String()
}

func (s *PushPromptResponse) SetHeaders(v map[string]*string) *PushPromptResponse {
	s.Headers = v
	return s
}

func (s *PushPromptResponse) SetStatusCode(v int32) *PushPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *PushPromptResponse) SetBody(v *PushPromptResponseBody) *PushPromptResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dashdeviceconsole"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// get prompt
//
// @param request - GetPromptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPromptResponse
func (client *Client) GetPromptWithOptions(request *GetPromptRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPromptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrompt"),
		Version:     tea.String("2025-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/prompt/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetPromptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetPromptResponse{}
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
// get prompt
//
// @param request - GetPromptRequest
//
// @return GetPromptResponse
func (client *Client) GetPrompt(request *GetPromptRequest) (_result *GetPromptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPromptResponse{}
	_body, _err := client.GetPromptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// push prompt
//
// @param request - PushPromptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushPromptResponse
func (client *Client) PushPromptWithOptions(request *PushPromptRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushPromptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PromptContent)) {
		body["promptContent"] = request.PromptContent
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushPrompt"),
		Version:     tea.String("2025-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/prompt/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PushPromptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PushPromptResponse{}
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
// push prompt
//
// @param request - PushPromptRequest
//
// @return PushPromptResponse
func (client *Client) PushPrompt(request *PushPromptRequest) (_result *PushPromptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushPromptResponse{}
	_body, _err := client.PushPromptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
