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

type EdasRefundRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s EdasRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s EdasRefundRequest) GoString() string {
	return s.String()
}

func (s *EdasRefundRequest) SetData(v string) *EdasRefundRequest {
	s.Data = &v
	return s
}

type EdasRefundResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EdasRefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EdasRefundResponseBody) GoString() string {
	return s.String()
}

func (s *EdasRefundResponseBody) SetCode(v string) *EdasRefundResponseBody {
	s.Code = &v
	return s
}

func (s *EdasRefundResponseBody) SetData(v string) *EdasRefundResponseBody {
	s.Data = &v
	return s
}

func (s *EdasRefundResponseBody) SetMessage(v string) *EdasRefundResponseBody {
	s.Message = &v
	return s
}

func (s *EdasRefundResponseBody) SetRequestId(v string) *EdasRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *EdasRefundResponseBody) SetSuccess(v bool) *EdasRefundResponseBody {
	s.Success = &v
	return s
}

type EdasRefundResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EdasRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EdasRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s EdasRefundResponse) GoString() string {
	return s.String()
}

func (s *EdasRefundResponse) SetHeaders(v map[string]*string) *EdasRefundResponse {
	s.Headers = v
	return s
}

func (s *EdasRefundResponse) SetStatusCode(v int32) *EdasRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *EdasRefundResponse) SetBody(v *EdasRefundResponseBody) *EdasRefundResponse {
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
		"ap-northeast-2-pop":          tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("edas.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("edas.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("edas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("edas.aliyuncs.com"),
		"cn-chengdu":                  tea.String("edas.aliyuncs.com"),
		"cn-edge-1":                   tea.String("edas.aliyuncs.com"),
		"cn-fujian":                   tea.String("edas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("edas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("edas.aliyuncs.com"),
		"cn-huhehaote":                tea.String("edas.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("edas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("edas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("edas.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("edas.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("edas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("edas.aliyuncs.com"),
		"cn-wuhan":                    tea.String("edas.aliyuncs.com"),
		"cn-yushanfang":               tea.String("edas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("edas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("edas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("edas.aliyuncs.com"),
		"eu-west-1":                   tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("edas.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("edas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) EdasRefundWithOptions(request *EdasRefundRequest, runtime *util.RuntimeOptions) (_result *EdasRefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EdasRefund"),
		Version:     tea.String("2016-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EdasRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EdasRefund(request *EdasRefundRequest) (_result *EdasRefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EdasRefundResponse{}
	_body, _err := client.EdasRefundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
