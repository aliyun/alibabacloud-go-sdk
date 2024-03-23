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

type RefreshLoginTokenRequest struct {
	ClientId        *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType      *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	EndUserId       *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginIdentifier *string `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	LoginToken      *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OfficeSiteId    *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	SessionId       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RefreshLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenRequest) SetClientId(v string) *RefreshLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetClientType(v string) *RefreshLoginTokenRequest {
	s.ClientType = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetEndUserId(v string) *RefreshLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginIdentifier(v string) *RefreshLoginTokenRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginToken(v string) *RefreshLoginTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetOfficeSiteId(v string) *RefreshLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetSessionId(v string) *RefreshLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetUuid(v string) *RefreshLoginTokenRequest {
	s.Uuid = &v
	return s
}

type RefreshLoginTokenResponseBody struct {
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponseBody) SetLoginToken(v string) *RefreshLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenResponseBody) SetRequestId(v string) *RefreshLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshLoginTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponse) SetHeaders(v map[string]*string) *RefreshLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLoginTokenResponse) SetStatusCode(v int32) *RefreshLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLoginTokenResponse) SetBody(v *RefreshLoginTokenResponseBody) *RefreshLoginTokenResponse {
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

func (client *Client) RefreshLoginTokenWithOptions(request *RefreshLoginTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginIdentifier)) {
		query["LoginIdentifier"] = request.LoginIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshLoginToken"),
		Version:     tea.String("2021-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshLoginToken(request *RefreshLoginTokenRequest) (_result *RefreshLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.RefreshLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
