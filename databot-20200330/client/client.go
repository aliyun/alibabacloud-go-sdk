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

type QueryDatabotRequest struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Query     *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s QueryDatabotRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDatabotRequest) GoString() string {
	return s.String()
}

func (s *QueryDatabotRequest) SetAccessId(v string) *QueryDatabotRequest {
	s.AccessId = &v
	return s
}

func (s *QueryDatabotRequest) SetQuery(v string) *QueryDatabotRequest {
	s.Query = &v
	return s
}

func (s *QueryDatabotRequest) SetSessionId(v string) *QueryDatabotRequest {
	s.SessionId = &v
	return s
}

type QueryDatabotResponseBody struct {
	CostTime  *int64                 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryDatabotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatabotResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatabotResponseBody) SetCostTime(v int64) *QueryDatabotResponseBody {
	s.CostTime = &v
	return s
}

func (s *QueryDatabotResponseBody) SetRequestId(v string) *QueryDatabotResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatabotResponseBody) SetData(v map[string]interface{}) *QueryDatabotResponseBody {
	s.Data = v
	return s
}

type QueryDatabotResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDatabotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDatabotResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDatabotResponse) GoString() string {
	return s.String()
}

func (s *QueryDatabotResponse) SetHeaders(v map[string]*string) *QueryDatabotResponse {
	s.Headers = v
	return s
}

func (s *QueryDatabotResponse) SetBody(v *QueryDatabotResponseBody) *QueryDatabotResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("databot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) QueryDatabotWithOptions(request *QueryDatabotRequest, runtime *util.RuntimeOptions) (_result *QueryDatabotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDatabotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDatabot"), tea.String("2020-03-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDatabot(request *QueryDatabotRequest) (_result *QueryDatabotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDatabotResponse{}
	_body, _err := client.QueryDatabotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
