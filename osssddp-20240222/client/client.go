// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetSddpVersionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s GetSddpVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSddpVersionRequest) GoString() string {
	return s.String()
}

func (s *GetSddpVersionRequest) SetClientToken(v string) *GetSddpVersionRequest {
	s.ClientToken = &v
	return s
}

type GetSddpVersionResponseBody struct {
	Content   *int32  `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSddpVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSddpVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSddpVersionResponseBody) SetContent(v int32) *GetSddpVersionResponseBody {
	s.Content = &v
	return s
}

func (s *GetSddpVersionResponseBody) SetRequestId(v string) *GetSddpVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetSddpVersionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSddpVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSddpVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSddpVersionResponse) GoString() string {
	return s.String()
}

func (s *GetSddpVersionResponse) SetHeaders(v map[string]*string) *GetSddpVersionResponse {
	s.Headers = v
	return s
}

func (s *GetSddpVersionResponse) SetStatusCode(v int32) *GetSddpVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSddpVersionResponse) SetBody(v *GetSddpVersionResponseBody) *GetSddpVersionResponse {
	s.Body = v
	return s
}

type UpgradeSddpVersionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OssVersion  *int32  `json:"OssVersion,omitempty" xml:"OssVersion,omitempty"`
}

func (s UpgradeSddpVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeSddpVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeSddpVersionRequest) SetClientToken(v string) *UpgradeSddpVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeSddpVersionRequest) SetOssVersion(v int32) *UpgradeSddpVersionRequest {
	s.OssVersion = &v
	return s
}

type UpgradeSddpVersionResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeSddpVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeSddpVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeSddpVersionResponseBody) SetContent(v string) *UpgradeSddpVersionResponseBody {
	s.Content = &v
	return s
}

func (s *UpgradeSddpVersionResponseBody) SetRequestId(v string) *UpgradeSddpVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeSddpVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeSddpVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeSddpVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeSddpVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeSddpVersionResponse) SetHeaders(v map[string]*string) *UpgradeSddpVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeSddpVersionResponse) SetStatusCode(v int32) *UpgradeSddpVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeSddpVersionResponse) SetBody(v *UpgradeSddpVersionResponseBody) *UpgradeSddpVersionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("osssddp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 查看用户的敏感数据保护版本信息
//
// @param request - GetSddpVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSddpVersionResponse
func (client *Client) GetSddpVersionWithOptions(request *GetSddpVersionRequest, runtime *util.RuntimeOptions) (_result *GetSddpVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSddpVersion"),
		Version:     tea.String("2024-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSddpVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看用户的敏感数据保护版本信息
//
// @param request - GetSddpVersionRequest
//
// @return GetSddpVersionResponse
func (client *Client) GetSddpVersion(request *GetSddpVersionRequest) (_result *GetSddpVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSddpVersionResponse{}
	_body, _err := client.GetSddpVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级敏感数据保护版本
//
// @param request - UpgradeSddpVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeSddpVersionResponse
func (client *Client) UpgradeSddpVersionWithOptions(request *UpgradeSddpVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeSddpVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OssVersion)) {
		query["OssVersion"] = request.OssVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeSddpVersion"),
		Version:     tea.String("2024-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeSddpVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级敏感数据保护版本
//
// @param request - UpgradeSddpVersionRequest
//
// @return UpgradeSddpVersionResponse
func (client *Client) UpgradeSddpVersion(request *UpgradeSddpVersionRequest) (_result *UpgradeSddpVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeSddpVersionResponse{}
	_body, _err := client.UpgradeSddpVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
