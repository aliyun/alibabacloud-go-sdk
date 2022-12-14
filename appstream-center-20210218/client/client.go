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

type GetAuthCodeRequest struct {
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s GetAuthCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *GetAuthCodeRequest) SetEndUserId(v string) *GetAuthCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *GetAuthCodeRequest) SetExternalUserId(v string) *GetAuthCodeRequest {
	s.ExternalUserId = &v
	return s
}

func (s *GetAuthCodeRequest) SetPolicy(v string) *GetAuthCodeRequest {
	s.Policy = &v
	return s
}

type GetAuthCodeResponseBody struct {
	AuthModel *GetAuthCodeResponseBodyAuthModel `json:"AuthModel,omitempty" xml:"AuthModel,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponseBody) SetAuthModel(v *GetAuthCodeResponseBodyAuthModel) *GetAuthCodeResponseBody {
	s.AuthModel = v
	return s
}

func (s *GetAuthCodeResponseBody) SetRequestId(v string) *GetAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthCodeResponseBodyAuthModel struct {
	AuthCode   *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	EndUserId  *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s GetAuthCodeResponseBodyAuthModel) String() string {
	return tea.Prettify(s)
}

func (s GetAuthCodeResponseBodyAuthModel) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponseBodyAuthModel) SetAuthCode(v string) *GetAuthCodeResponseBodyAuthModel {
	s.AuthCode = &v
	return s
}

func (s *GetAuthCodeResponseBodyAuthModel) SetEndUserId(v string) *GetAuthCodeResponseBodyAuthModel {
	s.EndUserId = &v
	return s
}

func (s *GetAuthCodeResponseBodyAuthModel) SetExpireTime(v string) *GetAuthCodeResponseBodyAuthModel {
	s.ExpireTime = &v
	return s
}

type GetAuthCodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponse) SetHeaders(v map[string]*string) *GetAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAuthCodeResponse) SetStatusCode(v int32) *GetAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthCodeResponse) SetBody(v *GetAuthCodeResponseBody) *GetAuthCodeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetAuthCodeWithOptions(request *GetAuthCodeRequest, runtime *util.RuntimeOptions) (_result *GetAuthCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUserId)) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthCode"),
		Version:     tea.String("2021-02-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthCode(request *GetAuthCodeRequest) (_result *GetAuthCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.GetAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
