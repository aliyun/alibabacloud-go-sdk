// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GenerateUserSessionTokenRequest struct {
	// AI Assistant ID
	//
	// example:
	//
	// cb-069d508f9ab341b1****
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Expiration Time, in seconds, default 24 hours
	//
	// example:
	//
	// 6000
	ExpireSecond *int64 `json:"ExpireSecond,omitempty" xml:"ExpireSecond,omitempty"`
	// Integration ID
	//
	// example:
	//
	// cit-960f499au184m7****
	IntegrateId *string `json:"IntegrateId,omitempty" xml:"IntegrateId,omitempty"`
	// User Avatar (URL)
	//
	// example:
	//
	// https://xxxx.com/xxx
	UserAvatar *string `json:"UserAvatar,omitempty" xml:"UserAvatar,omitempty"`
	// User ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 929293312213****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// User Nickname
	//
	// example:
	//
	// testxxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GenerateUserSessionTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUserSessionTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateUserSessionTokenRequest) SetChatbotId(v string) *GenerateUserSessionTokenRequest {
	s.ChatbotId = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetExpireSecond(v int64) *GenerateUserSessionTokenRequest {
	s.ExpireSecond = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetIntegrateId(v string) *GenerateUserSessionTokenRequest {
	s.IntegrateId = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetUserAvatar(v string) *GenerateUserSessionTokenRequest {
	s.UserAvatar = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetUserId(v string) *GenerateUserSessionTokenRequest {
	s.UserId = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetUserName(v string) *GenerateUserSessionTokenRequest {
	s.UserName = &v
	return s
}

type GenerateUserSessionTokenResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 34C2713A-2270-5EEC-825E-115F9AD3BAC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Token.
	//
	// example:
	//
	// 960f499au184m7****
	UserSessionToken *string `json:"UserSessionToken,omitempty" xml:"UserSessionToken,omitempty"`
}

func (s GenerateUserSessionTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUserSessionTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUserSessionTokenResponseBody) SetRequestId(v string) *GenerateUserSessionTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUserSessionTokenResponseBody) SetUserSessionToken(v string) *GenerateUserSessionTokenResponseBody {
	s.UserSessionToken = &v
	return s
}

type GenerateUserSessionTokenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUserSessionTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUserSessionTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUserSessionTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateUserSessionTokenResponse) SetHeaders(v map[string]*string) *GenerateUserSessionTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateUserSessionTokenResponse) SetStatusCode(v int32) *GenerateUserSessionTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUserSessionTokenResponse) SetBody(v *GenerateUserSessionTokenResponseBody) *GenerateUserSessionTokenResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("appflow"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Generate Login Session Token
//
// @param request - GenerateUserSessionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionTokenWithOptions(request *GenerateUserSessionTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateUserSessionTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatbotId)) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireSecond)) {
		query["ExpireSecond"] = request.ExpireSecond
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrateId)) {
		query["IntegrateId"] = request.IntegrateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAvatar)) {
		query["UserAvatar"] = request.UserAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUserSessionToken"),
		Version:     tea.String("2023-09-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Generate Login Session Token
//
// @param request - GenerateUserSessionTokenRequest
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionToken(request *GenerateUserSessionTokenRequest) (_result *GenerateUserSessionTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.GenerateUserSessionTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
