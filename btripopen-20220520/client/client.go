// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccessTokenRequest struct {
	// 应用的appKey
	AppKey *string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 应用的appSecret
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s AccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenRequest) GoString() string {
	return s.String()
}

func (s *AccessTokenRequest) SetAppKey(v string) *AccessTokenRequest {
	s.AppKey = &v
	return s
}

func (s *AccessTokenRequest) SetAppSecret(v string) *AccessTokenRequest {
	s.AppSecret = &v
	return s
}

type AccessTokenResponseBody struct {
	// 错误码
	Code *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *AccessTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 错误信息
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBody) SetCode(v string) *AccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *AccessTokenResponseBody) SetData(v *AccessTokenResponseBodyData) *AccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *AccessTokenResponseBody) SetMessage(v string) *AccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *AccessTokenResponseBody) SetRequestId(v string) *AccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessTokenResponseBody) SetTraceId(v string) *AccessTokenResponseBody {
	s.TraceId = &v
	return s
}

type AccessTokenResponseBodyData struct {
	// 过期时间，单位ms
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// 令牌
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s AccessTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccessTokenResponseBodyData) SetExpire(v int64) *AccessTokenResponseBodyData {
	s.Expire = &v
	return s
}

func (s *AccessTokenResponseBodyData) SetToken(v string) *AccessTokenResponseBodyData {
	s.Token = &v
	return s
}

type AccessTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessTokenResponse) GoString() string {
	return s.String()
}

func (s *AccessTokenResponse) SetHeaders(v map[string]*string) *AccessTokenResponse {
	s.Headers = v
	return s
}

func (s *AccessTokenResponse) SetStatusCode(v int32) *AccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessTokenResponse) SetBody(v *AccessTokenResponseBody) *AccessTokenResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("btripopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AccessToken(request *AccessTokenRequest) (_result *AccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AccessTokenResponse{}
	_body, _err := client.AccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AccessTokenWithOptions(request *AccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["app_key"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["app_secret"] = request.AppSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AccessToken"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/btrip-open-auth/v1/access-token/action/take"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
