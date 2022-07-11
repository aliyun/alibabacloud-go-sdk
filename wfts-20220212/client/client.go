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

type GetLjxAccountInfoRequest struct {
	// 实例 ID。
	LjxAccountInfoId *string `json:"LjxAccountInfoId,omitempty" xml:"LjxAccountInfoId,omitempty"`
}

func (s GetLjxAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLjxAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLjxAccountInfoRequest) SetLjxAccountInfoId(v string) *GetLjxAccountInfoRequest {
	s.LjxAccountInfoId = &v
	return s
}

type GetLjxAccountInfoResponseBody struct {
	Apple *string `json:"Apple,omitempty" xml:"Apple,omitempty"`
	// 资源一级ID
	LjxAccountInfoId *string `json:"LjxAccountInfoId,omitempty" xml:"LjxAccountInfoId,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLjxAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLjxAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLjxAccountInfoResponseBody) SetApple(v string) *GetLjxAccountInfoResponseBody {
	s.Apple = &v
	return s
}

func (s *GetLjxAccountInfoResponseBody) SetLjxAccountInfoId(v string) *GetLjxAccountInfoResponseBody {
	s.LjxAccountInfoId = &v
	return s
}

func (s *GetLjxAccountInfoResponseBody) SetRequestId(v string) *GetLjxAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetLjxAccountInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLjxAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLjxAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLjxAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLjxAccountInfoResponse) SetHeaders(v map[string]*string) *GetLjxAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLjxAccountInfoResponse) SetStatusCode(v int32) *GetLjxAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLjxAccountInfoResponse) SetBody(v *GetLjxAccountInfoResponseBody) *GetLjxAccountInfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("wfts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetLjxAccountInfo(request *GetLjxAccountInfoRequest) (_result *GetLjxAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLjxAccountInfoResponse{}
	_body, _err := client.GetLjxAccountInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLjxAccountInfoWithOptions(request *GetLjxAccountInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLjxAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LjxAccountInfoId)) {
		query["LjxAccountInfoId"] = request.LjxAccountInfoId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLjxAccountInfo"),
		Version:     tea.String("2022-02-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/get/ljx/acc"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLjxAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
