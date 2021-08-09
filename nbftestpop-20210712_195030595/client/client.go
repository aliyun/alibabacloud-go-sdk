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

type NewAuthTestRequest struct {
	X *int32 `json:"x,omitempty" xml:"x,omitempty"`
	Y *int32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s NewAuthTestRequest) String() string {
	return tea.Prettify(s)
}

func (s NewAuthTestRequest) GoString() string {
	return s.String()
}

func (s *NewAuthTestRequest) SetX(v int32) *NewAuthTestRequest {
	s.X = &v
	return s
}

func (s *NewAuthTestRequest) SetY(v int32) *NewAuthTestRequest {
	s.Y = &v
	return s
}

type NewAuthTestResponseBody struct {
	Sum *int32 `json:"sum,omitempty" xml:"sum,omitempty"`
}

func (s NewAuthTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NewAuthTestResponseBody) GoString() string {
	return s.String()
}

func (s *NewAuthTestResponseBody) SetSum(v int32) *NewAuthTestResponseBody {
	s.Sum = &v
	return s
}

type NewAuthTestResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NewAuthTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NewAuthTestResponse) String() string {
	return tea.Prettify(s)
}

func (s NewAuthTestResponse) GoString() string {
	return s.String()
}

func (s *NewAuthTestResponse) SetHeaders(v map[string]*string) *NewAuthTestResponse {
	s.Headers = v
	return s
}

func (s *NewAuthTestResponse) SetBody(v *NewAuthTestResponseBody) *NewAuthTestResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("nbftestpop"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) NewAuthTestWithOptions(request *NewAuthTestRequest, runtime *util.RuntimeOptions) (_result *NewAuthTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &NewAuthTestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("NewAuthTest"), tea.String("2021-07-12_19-50-30-595"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NewAuthTest(request *NewAuthTestRequest) (_result *NewAuthTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NewAuthTestResponse{}
	_body, _err := client.NewAuthTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
