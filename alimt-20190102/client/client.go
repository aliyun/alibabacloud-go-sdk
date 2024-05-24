// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetResourceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetCode(v string) *GetResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceResponseBody) SetData(v string) *GetResourceResponseBody {
	s.Data = &v
	return s
}

func (s *GetResourceResponseBody) SetMessage(v string) *GetResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type ListMtConnectorRequest struct {
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMtConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMtConnectorRequest) GoString() string {
	return s.String()
}

func (s *ListMtConnectorRequest) SetUserId(v int64) *ListMtConnectorRequest {
	s.UserId = &v
	return s
}

type ListMtConnectorResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BizErrorCode *int64                             `json:"bizErrorCode,omitempty" xml:"bizErrorCode,omitempty"`
	Data         []*ListMtConnectorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Success      *string                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMtConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMtConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *ListMtConnectorResponseBody) SetRequestId(v string) *ListMtConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMtConnectorResponseBody) SetBizErrorCode(v int64) *ListMtConnectorResponseBody {
	s.BizErrorCode = &v
	return s
}

func (s *ListMtConnectorResponseBody) SetData(v []*ListMtConnectorResponseBodyData) *ListMtConnectorResponseBody {
	s.Data = v
	return s
}

func (s *ListMtConnectorResponseBody) SetErrorMessage(v string) *ListMtConnectorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMtConnectorResponseBody) SetSuccess(v string) *ListMtConnectorResponseBody {
	s.Success = &v
	return s
}

type ListMtConnectorResponseBodyData struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMtConnectorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMtConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMtConnectorResponseBodyData) SetKey(v string) *ListMtConnectorResponseBodyData {
	s.Key = &v
	return s
}

func (s *ListMtConnectorResponseBodyData) SetValue(v string) *ListMtConnectorResponseBodyData {
	s.Value = &v
	return s
}

type ListMtConnectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMtConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMtConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMtConnectorResponse) GoString() string {
	return s.String()
}

func (s *ListMtConnectorResponse) SetHeaders(v map[string]*string) *ListMtConnectorResponse {
	s.Headers = v
	return s
}

func (s *ListMtConnectorResponse) SetStatusCode(v int32) *ListMtConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMtConnectorResponse) SetBody(v *ListMtConnectorResponseBody) *ListMtConnectorResponse {
	s.Body = v
	return s
}

type TranslateECommerceResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateECommerceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateECommerceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBody) SetCode(v string) *TranslateECommerceResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetData(v *TranslateECommerceResponseBodyData) *TranslateECommerceResponseBody {
	s.Data = v
	return s
}

func (s *TranslateECommerceResponseBody) SetMessage(v string) *TranslateECommerceResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetRequestId(v string) *TranslateECommerceResponseBody {
	s.RequestId = &v
	return s
}

type TranslateECommerceResponseBodyData struct {
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	WordCount  *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateECommerceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBodyData) SetTranslated(v string) *TranslateECommerceResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateECommerceResponseBodyData) SetWordCount(v string) *TranslateECommerceResponseBodyData {
	s.WordCount = &v
	return s
}

type TranslateECommerceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateECommerceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateECommerceResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponse) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponse) SetHeaders(v map[string]*string) *TranslateECommerceResponse {
	s.Headers = v
	return s
}

func (s *TranslateECommerceResponse) SetStatusCode(v int32) *TranslateECommerceResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateECommerceResponse) SetBody(v *TranslateECommerceResponseBody) *TranslateECommerceResponse {
	s.Body = v
	return s
}

type TranslateGeneralResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateGeneralResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBody) SetCode(v string) *TranslateGeneralResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetData(v *TranslateGeneralResponseBodyData) *TranslateGeneralResponseBody {
	s.Data = v
	return s
}

func (s *TranslateGeneralResponseBody) SetMessage(v string) *TranslateGeneralResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetRequestId(v string) *TranslateGeneralResponseBody {
	s.RequestId = &v
	return s
}

type TranslateGeneralResponseBodyData struct {
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	WordCount  *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateGeneralResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBodyData) SetTranslated(v string) *TranslateGeneralResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) SetWordCount(v string) *TranslateGeneralResponseBodyData {
	s.WordCount = &v
	return s
}

type TranslateGeneralResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponse) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponse) SetHeaders(v map[string]*string) *TranslateGeneralResponse {
	s.Headers = v
	return s
}

func (s *TranslateGeneralResponse) SetStatusCode(v int32) *TranslateGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateGeneralResponse) SetBody(v *TranslateGeneralResponseBody) *TranslateGeneralResponse {
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
		"cn-hangzhou":                 tea.String("mt.cn-hangzhou.aliyuncs.com"),
		"ap-northeast-1":              tea.String("mt.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("mt.aliyuncs.com"),
		"ap-south-1":                  tea.String("mt.aliyuncs.com"),
		"ap-southeast-1":              tea.String("mt.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("mt.aliyuncs.com"),
		"ap-southeast-3":              tea.String("mt.aliyuncs.com"),
		"ap-southeast-5":              tea.String("mt.aliyuncs.com"),
		"cn-beijing":                  tea.String("mt.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("mt.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("mt.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("mt.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("mt.aliyuncs.com"),
		"cn-chengdu":                  tea.String("mt.aliyuncs.com"),
		"cn-edge-1":                   tea.String("mt.aliyuncs.com"),
		"cn-fujian":                   tea.String("mt.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("mt.aliyuncs.com"),
		"cn-hongkong":                 tea.String("mt.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("mt.aliyuncs.com"),
		"cn-huhehaote":                tea.String("mt.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("mt.aliyuncs.com"),
		"cn-qingdao":                  tea.String("mt.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("mt.aliyuncs.com"),
		"cn-shanghai":                 tea.String("mt.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("mt.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("mt.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("mt.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("mt.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("mt.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("mt.aliyuncs.com"),
		"cn-wuhan":                    tea.String("mt.aliyuncs.com"),
		"cn-yushanfang":               tea.String("mt.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("mt.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("mt.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("mt.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("mt.aliyuncs.com"),
		"eu-central-1":                tea.String("mt.aliyuncs.com"),
		"eu-west-1":                   tea.String("mt.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("mt.aliyuncs.com"),
		"me-east-1":                   tea.String("mt.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("mt.aliyuncs.com"),
		"us-east-1":                   tea.String("mt.aliyuncs.com"),
		"us-west-1":                   tea.String("mt.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alimt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2019-01-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/resource"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return GetResourceResponse
func (client *Client) GetResource() (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取系统模型和AI模型
//
// @param request - ListMtConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMtConnectorResponse
func (client *Client) ListMtConnectorWithOptions(request *ListMtConnectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMtConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMtConnector"),
		Version:     tea.String("2019-01-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/alynx/listMtModels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMtConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统模型和AI模型
//
// @param request - ListMtConnectorRequest
//
// @return ListMtConnectorResponse
func (client *Client) ListMtConnector(request *ListMtConnectorRequest) (_result *ListMtConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMtConnectorResponse{}
	_body, _err := client.ListMtConnectorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateECommerceResponse
func (client *Client) TranslateECommerceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *TranslateECommerceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateECommerce"),
		Version:     tea.String("2019-01-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/translate/web/ecommerce"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return TranslateECommerceResponse
func (client *Client) TranslateECommerce() (_result *TranslateECommerceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TranslateECommerceResponse{}
	_body, _err := client.TranslateECommerceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralResponse
func (client *Client) TranslateGeneralWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *TranslateGeneralResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateGeneral"),
		Version:     tea.String("2019-01-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/translate/web/general"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return TranslateGeneralResponse
func (client *Client) TranslateGeneral() (_result *TranslateGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TranslateGeneralResponse{}
	_body, _err := client.TranslateGeneralWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
