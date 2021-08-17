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

type ArrayOutputRequest struct {
	Input *ArrayOutputRequestInput `json:"input,omitempty" xml:"input,omitempty" type:"Struct"`
}

func (s ArrayOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s ArrayOutputRequest) GoString() string {
	return s.String()
}

func (s *ArrayOutputRequest) SetInput(v *ArrayOutputRequestInput) *ArrayOutputRequest {
	s.Input = v
	return s
}

type ArrayOutputRequestInput struct {
	Request   *string `json:"request,omitempty" xml:"request,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ArrayOutputRequestInput) String() string {
	return tea.Prettify(s)
}

func (s ArrayOutputRequestInput) GoString() string {
	return s.String()
}

func (s *ArrayOutputRequestInput) SetRequest(v string) *ArrayOutputRequestInput {
	s.Request = &v
	return s
}

func (s *ArrayOutputRequestInput) SetRequestId(v string) *ArrayOutputRequestInput {
	s.RequestId = &v
	return s
}

type ArrayOutputShrinkRequest struct {
	InputShrink *string `json:"input,omitempty" xml:"input,omitempty"`
}

func (s ArrayOutputShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ArrayOutputShrinkRequest) GoString() string {
	return s.String()
}

func (s *ArrayOutputShrinkRequest) SetInputShrink(v string) *ArrayOutputShrinkRequest {
	s.InputShrink = &v
	return s
}

type ArrayOutputResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    []*string          `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ArrayOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s ArrayOutputResponse) GoString() string {
	return s.String()
}

func (s *ArrayOutputResponse) SetHeaders(v map[string]*string) *ArrayOutputResponse {
	s.Headers = v
	return s
}

func (s *ArrayOutputResponse) SetBody(v []*string) *ArrayOutputResponse {
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

func (client *Client) ArrayOutput(request *ArrayOutputRequest) (_result *ArrayOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ArrayOutputResponse{}
	_body, _err := client.ArrayOutputWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ArrayOutputWithOptions(tmpReq *ArrayOutputRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ArrayOutputResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ArrayOutputShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Input))) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Input), tea.String("input"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		query["input"] = request.InputShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ArrayOutputResponse{}
	_body, _err := client.DoROARequest(tea.String("ArrayOutput"), tea.String("2021-08-04_14-46-22-100"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/kxo/arrayoutput"), tea.String("array"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
