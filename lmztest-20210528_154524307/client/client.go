// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type YxTestAPIRequest struct {
	// param0
	F0 *int32 `json:"f0,omitempty" xml:"f0,omitempty"`
	// param1
	F1 *int32 `json:"f1,omitempty" xml:"f1,omitempty"`
}

func (s YxTestAPIRequest) String() string {
	return tea.Prettify(s)
}

func (s YxTestAPIRequest) GoString() string {
	return s.String()
}

func (s *YxTestAPIRequest) SetF0(v int32) *YxTestAPIRequest {
	s.F0 = &v
	return s
}

func (s *YxTestAPIRequest) SetF1(v int32) *YxTestAPIRequest {
	s.F1 = &v
	return s
}

type YxTestAPIResponseBody struct {
	// sum
	Sum *int32 `json:"sum,omitempty" xml:"sum,omitempty"`
}

func (s YxTestAPIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s YxTestAPIResponseBody) GoString() string {
	return s.String()
}

func (s *YxTestAPIResponseBody) SetSum(v int32) *YxTestAPIResponseBody {
	s.Sum = &v
	return s
}

type YxTestAPIResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *YxTestAPIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s YxTestAPIResponse) String() string {
	return tea.Prettify(s)
}

func (s YxTestAPIResponse) GoString() string {
	return s.String()
}

func (s *YxTestAPIResponse) SetHeaders(v map[string]*string) *YxTestAPIResponse {
	s.Headers = v
	return s
}

func (s *YxTestAPIResponse) SetBody(v *YxTestAPIResponseBody) *YxTestAPIResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("lmztest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) YxTestAPIWithOptions(request *YxTestAPIRequest, runtime *util.RuntimeOptions) (_result *YxTestAPIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &YxTestAPIResponse{}
	_body, _err := client.DoRPCRequest(tea.String("YxTestAPI"), tea.String("2021-05-28_15-45-24-307"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) YxTestAPI(request *YxTestAPIRequest) (_result *YxTestAPIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &YxTestAPIResponse{}
	_body, _err := client.YxTestAPIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
