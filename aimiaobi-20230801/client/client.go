// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateTokenRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetAgentKey(v string) *CreateTokenRequest {
	s.AgentKey = &v
	return s
}

type CreateTokenResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) SetCode(v string) *CreateTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTokenResponseBody) SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenResponseBody) SetHttpStatusCode(v string) *CreateTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTokenResponseBody) SetMessage(v string) *CreateTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetSuccess(v bool) *CreateTokenResponseBody {
	s.Success = &v
	return s
}

type CreateTokenResponseBodyData struct {
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyData) SetExpiredTime(v int64) *CreateTokenResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *CreateTokenResponseBodyData) SetToken(v string) *CreateTokenResponseBodyData {
	s.Token = &v
	return s
}

type CreateTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetStatusCode(v int32) *CreateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aimiaobi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateToken"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
