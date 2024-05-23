// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type MachineTranslateECommerceRequest struct {
	// This parameter is required.
	ContentFormat *string `json:"ContentFormat,omitempty" xml:"ContentFormat,omitempty"`
	// This parameter is required.
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s MachineTranslateECommerceRequest) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateECommerceRequest) GoString() string {
	return s.String()
}

func (s *MachineTranslateECommerceRequest) SetContentFormat(v string) *MachineTranslateECommerceRequest {
	s.ContentFormat = &v
	return s
}

func (s *MachineTranslateECommerceRequest) SetSourceLanguage(v string) *MachineTranslateECommerceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *MachineTranslateECommerceRequest) SetSourceText(v string) *MachineTranslateECommerceRequest {
	s.SourceText = &v
	return s
}

func (s *MachineTranslateECommerceRequest) SetTargetLanguage(v string) *MachineTranslateECommerceRequest {
	s.TargetLanguage = &v
	return s
}

type MachineTranslateECommerceResponseBody struct {
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *MachineTranslateECommerceResponseBodyResultCode `json:"ResultCode,omitempty" xml:"ResultCode,omitempty" type:"Struct"`
	Success       *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TranslateText *string                                          `json:"TranslateText,omitempty" xml:"TranslateText,omitempty"`
}

func (s MachineTranslateECommerceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateECommerceResponseBody) GoString() string {
	return s.String()
}

func (s *MachineTranslateECommerceResponseBody) SetRequestId(v string) *MachineTranslateECommerceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MachineTranslateECommerceResponseBody) SetResultCode(v *MachineTranslateECommerceResponseBodyResultCode) *MachineTranslateECommerceResponseBody {
	s.ResultCode = v
	return s
}

func (s *MachineTranslateECommerceResponseBody) SetSuccess(v bool) *MachineTranslateECommerceResponseBody {
	s.Success = &v
	return s
}

func (s *MachineTranslateECommerceResponseBody) SetTranslateText(v string) *MachineTranslateECommerceResponseBody {
	s.TranslateText = &v
	return s
}

type MachineTranslateECommerceResponseBodyResultCode struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MachineTranslateECommerceResponseBodyResultCode) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateECommerceResponseBodyResultCode) GoString() string {
	return s.String()
}

func (s *MachineTranslateECommerceResponseBodyResultCode) SetCode(v int32) *MachineTranslateECommerceResponseBodyResultCode {
	s.Code = &v
	return s
}

func (s *MachineTranslateECommerceResponseBodyResultCode) SetMessage(v string) *MachineTranslateECommerceResponseBodyResultCode {
	s.Message = &v
	return s
}

type MachineTranslateECommerceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MachineTranslateECommerceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MachineTranslateECommerceResponse) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateECommerceResponse) GoString() string {
	return s.String()
}

func (s *MachineTranslateECommerceResponse) SetHeaders(v map[string]*string) *MachineTranslateECommerceResponse {
	s.Headers = v
	return s
}

func (s *MachineTranslateECommerceResponse) SetStatusCode(v int32) *MachineTranslateECommerceResponse {
	s.StatusCode = &v
	return s
}

func (s *MachineTranslateECommerceResponse) SetBody(v *MachineTranslateECommerceResponseBody) *MachineTranslateECommerceResponse {
	s.Body = v
	return s
}

type MachineTranslateGeneralRequest struct {
	// This parameter is required.
	ContentFormat *string `json:"ContentFormat,omitempty" xml:"ContentFormat,omitempty"`
	// This parameter is required.
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s MachineTranslateGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateGeneralRequest) GoString() string {
	return s.String()
}

func (s *MachineTranslateGeneralRequest) SetContentFormat(v string) *MachineTranslateGeneralRequest {
	s.ContentFormat = &v
	return s
}

func (s *MachineTranslateGeneralRequest) SetSourceLanguage(v string) *MachineTranslateGeneralRequest {
	s.SourceLanguage = &v
	return s
}

func (s *MachineTranslateGeneralRequest) SetSourceText(v string) *MachineTranslateGeneralRequest {
	s.SourceText = &v
	return s
}

func (s *MachineTranslateGeneralRequest) SetTargetLanguage(v string) *MachineTranslateGeneralRequest {
	s.TargetLanguage = &v
	return s
}

type MachineTranslateGeneralResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *MachineTranslateGeneralResponseBodyResultCode `json:"ResultCode,omitempty" xml:"ResultCode,omitempty" type:"Struct"`
	Success       *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TranslateText *string                                        `json:"TranslateText,omitempty" xml:"TranslateText,omitempty"`
}

func (s MachineTranslateGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *MachineTranslateGeneralResponseBody) SetRequestId(v string) *MachineTranslateGeneralResponseBody {
	s.RequestId = &v
	return s
}

func (s *MachineTranslateGeneralResponseBody) SetResultCode(v *MachineTranslateGeneralResponseBodyResultCode) *MachineTranslateGeneralResponseBody {
	s.ResultCode = v
	return s
}

func (s *MachineTranslateGeneralResponseBody) SetSuccess(v bool) *MachineTranslateGeneralResponseBody {
	s.Success = &v
	return s
}

func (s *MachineTranslateGeneralResponseBody) SetTranslateText(v string) *MachineTranslateGeneralResponseBody {
	s.TranslateText = &v
	return s
}

type MachineTranslateGeneralResponseBodyResultCode struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MachineTranslateGeneralResponseBodyResultCode) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateGeneralResponseBodyResultCode) GoString() string {
	return s.String()
}

func (s *MachineTranslateGeneralResponseBodyResultCode) SetCode(v int32) *MachineTranslateGeneralResponseBodyResultCode {
	s.Code = &v
	return s
}

func (s *MachineTranslateGeneralResponseBodyResultCode) SetMessage(v string) *MachineTranslateGeneralResponseBodyResultCode {
	s.Message = &v
	return s
}

type MachineTranslateGeneralResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MachineTranslateGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MachineTranslateGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s MachineTranslateGeneralResponse) GoString() string {
	return s.String()
}

func (s *MachineTranslateGeneralResponse) SetHeaders(v map[string]*string) *MachineTranslateGeneralResponse {
	s.Headers = v
	return s
}

func (s *MachineTranslateGeneralResponse) SetStatusCode(v int32) *MachineTranslateGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *MachineTranslateGeneralResponse) SetBody(v *MachineTranslateGeneralResponseBody) *MachineTranslateGeneralResponse {
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

// @param request - MachineTranslateECommerceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MachineTranslateECommerceResponse
func (client *Client) MachineTranslateECommerceWithOptions(request *MachineTranslateECommerceRequest, runtime *util.RuntimeOptions) (_result *MachineTranslateECommerceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentFormat)) {
		body["ContentFormat"] = request.ContentFormat
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MachineTranslateECommerce"),
		Version:     tea.String("2018-08-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MachineTranslateECommerceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - MachineTranslateECommerceRequest
//
// @return MachineTranslateECommerceResponse
func (client *Client) MachineTranslateECommerce(request *MachineTranslateECommerceRequest) (_result *MachineTranslateECommerceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MachineTranslateECommerceResponse{}
	_body, _err := client.MachineTranslateECommerceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MachineTranslateGeneralRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MachineTranslateGeneralResponse
func (client *Client) MachineTranslateGeneralWithOptions(request *MachineTranslateGeneralRequest, runtime *util.RuntimeOptions) (_result *MachineTranslateGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentFormat)) {
		body["ContentFormat"] = request.ContentFormat
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MachineTranslateGeneral"),
		Version:     tea.String("2018-08-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MachineTranslateGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - MachineTranslateGeneralRequest
//
// @return MachineTranslateGeneralResponse
func (client *Client) MachineTranslateGeneral(request *MachineTranslateGeneralRequest) (_result *MachineTranslateGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MachineTranslateGeneralResponse{}
	_body, _err := client.MachineTranslateGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
