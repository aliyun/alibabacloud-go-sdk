// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AdadaAResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AdadaAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdadaAResponseBody) GoString() string {
	return s.String()
}

func (s *AdadaAResponseBody) SetErrorCode(v string) *AdadaAResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AdadaAResponseBody) SetErrorMsg(v string) *AdadaAResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AdadaAResponseBody) SetResult(v string) *AdadaAResponseBody {
	s.Result = &v
	return s
}

func (s *AdadaAResponseBody) SetSuccess(v bool) *AdadaAResponseBody {
	s.Success = &v
	return s
}

type AdadaAResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdadaAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdadaAResponse) String() string {
	return tea.Prettify(s)
}

func (s AdadaAResponse) GoString() string {
	return s.String()
}

func (s *AdadaAResponse) SetHeaders(v map[string]*string) *AdadaAResponse {
	s.Headers = v
	return s
}

func (s *AdadaAResponse) SetStatusCode(v int32) *AdadaAResponse {
	s.StatusCode = &v
	return s
}

func (s *AdadaAResponse) SetBody(v *AdadaAResponseBody) *AdadaAResponse {
	s.Body = v
	return s
}

type YxTestApiResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s YxTestApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s YxTestApiResponseBody) GoString() string {
	return s.String()
}

func (s *YxTestApiResponseBody) SetErrorCode(v string) *YxTestApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *YxTestApiResponseBody) SetErrorMsg(v string) *YxTestApiResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *YxTestApiResponseBody) SetResult(v string) *YxTestApiResponseBody {
	s.Result = &v
	return s
}

func (s *YxTestApiResponseBody) SetSuccess(v bool) *YxTestApiResponseBody {
	s.Success = &v
	return s
}

type YxTestApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *YxTestApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s YxTestApiResponse) String() string {
	return tea.Prettify(s)
}

func (s YxTestApiResponse) GoString() string {
	return s.String()
}

func (s *YxTestApiResponse) SetHeaders(v map[string]*string) *YxTestApiResponse {
	s.Headers = v
	return s
}

func (s *YxTestApiResponse) SetStatusCode(v int32) *YxTestApiResponse {
	s.StatusCode = &v
	return s
}

func (s *YxTestApiResponse) SetBody(v *YxTestApiResponseBody) *YxTestApiResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("nbf-vpc-cloud"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AdadaAWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AdadaAResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AdadaA"),
		Version:     tea.String("2021-11-15_13-11-23-360"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/caihe_cloud_product_1/1_0_0/adadaA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AdadaAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AdadaA() (_result *AdadaAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AdadaAResponse{}
	_body, _err := client.AdadaAWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) YxTestApiWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *YxTestApiResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("YxTestApi"),
		Version:     tea.String("2021-11-15_13-11-23-360"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/caihe_cloud_product_1/1_0_0/yxTestApi"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &YxTestApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) YxTestApi() (_result *YxTestApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &YxTestApiResponse{}
	_body, _err := client.YxTestApiWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
