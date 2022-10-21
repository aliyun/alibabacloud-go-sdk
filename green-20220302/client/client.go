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

type TextModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s TextModerationRequest) GoString() string {
	return s.String()
}

func (s *TextModerationRequest) SetService(v string) *TextModerationRequest {
	s.Service = &v
	return s
}

func (s *TextModerationRequest) SetServiceParameters(v string) *TextModerationRequest {
	s.ServiceParameters = &v
	return s
}

type TextModerationResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TextModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponseBody) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBody) SetCode(v string) *TextModerationResponseBody {
	s.Code = &v
	return s
}

func (s *TextModerationResponseBody) SetData(v *TextModerationResponseBodyData) *TextModerationResponseBody {
	s.Data = v
	return s
}

func (s *TextModerationResponseBody) SetMessage(v string) *TextModerationResponseBody {
	s.Message = &v
	return s
}

func (s *TextModerationResponseBody) SetRequestId(v string) *TextModerationResponseBody {
	s.RequestId = &v
	return s
}

type TextModerationResponseBodyData struct {
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s TextModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBodyData) SetLabels(v string) *TextModerationResponseBodyData {
	s.Labels = &v
	return s
}

func (s *TextModerationResponseBodyData) SetReason(v string) *TextModerationResponseBodyData {
	s.Reason = &v
	return s
}

type TextModerationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TextModerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TextModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponse) GoString() string {
	return s.String()
}

func (s *TextModerationResponse) SetHeaders(v map[string]*string) *TextModerationResponse {
	s.Headers = v
	return s
}

func (s *TextModerationResponse) SetStatusCode(v int32) *TextModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *TextModerationResponse) SetBody(v *TextModerationResponseBody) *TextModerationResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("green.aliyuncs.com"),
		"cn-hongkong":           tea.String("green.aliyuncs.com"),
		"cn-huhehaote":          tea.String("green.aliyuncs.com"),
		"cn-qingdao":            tea.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("green.aliyuncs.com"),
		"eu-central-1":          tea.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) TextModerationWithOptions(request *TextModerationRequest, runtime *util.RuntimeOptions) (_result *TextModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TextModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TextModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextModeration(request *TextModerationRequest) (_result *TextModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TextModerationResponse{}
	_body, _err := client.TextModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
