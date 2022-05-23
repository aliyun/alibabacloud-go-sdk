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

type TakeAccessTokenRequest struct {
	AppKey    *string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s TakeAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenRequest) SetAppKey(v string) *TakeAccessTokenRequest {
	s.AppKey = &v
	return s
}

func (s *TakeAccessTokenRequest) SetAppSecret(v string) *TakeAccessTokenRequest {
	s.AppSecret = &v
	return s
}

type TakeAccessTokenResponseBody struct {
	Code    *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data    *TakeAccessTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                          `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TakeAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenResponseBody) SetCode(v string) *TakeAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *TakeAccessTokenResponseBody) SetData(v *TakeAccessTokenResponseBodyData) *TakeAccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *TakeAccessTokenResponseBody) SetMessage(v string) *TakeAccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *TakeAccessTokenResponseBody) SetRequestId(v string) *TakeAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *TakeAccessTokenResponseBody) SetSuccess(v string) *TakeAccessTokenResponseBody {
	s.Success = &v
	return s
}

type TakeAccessTokenResponseBodyData struct {
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Expire      *int64  `json:"expire,omitempty" xml:"expire,omitempty"`
}

func (s TakeAccessTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenResponseBodyData) SetAccessToken(v string) *TakeAccessTokenResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *TakeAccessTokenResponseBodyData) SetExpire(v int64) *TakeAccessTokenResponseBodyData {
	s.Expire = &v
	return s
}

type TakeAccessTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TakeAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TakeAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s TakeAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *TakeAccessTokenResponse) SetHeaders(v map[string]*string) *TakeAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *TakeAccessTokenResponse) SetStatusCode(v int32) *TakeAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TakeAccessTokenResponse) SetBody(v *TakeAccessTokenResponseBody) *TakeAccessTokenResponse {
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

func (client *Client) TakeAccessToken(request *TakeAccessTokenRequest) (_result *TakeAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TakeAccessTokenResponse{}
	_body, _err := client.TakeAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TakeAccessTokenWithOptions(request *TakeAccessTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TakeAccessTokenResponse, _err error) {
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
		Action:      tea.String("TakeAccessToken"),
		Version:     tea.String("2022-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/btrip/open/access-token/take"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TakeAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
