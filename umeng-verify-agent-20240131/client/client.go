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

type CreateRequest struct {
	AppName    *string `json:"appName,omitempty" xml:"appName,omitempty"`
	BundleId   *string `json:"bundleId,omitempty" xml:"bundleId,omitempty"`
	PackName   *string `json:"packName,omitempty" xml:"packName,omitempty"`
	Platform   *string `json:"platform,omitempty" xml:"platform,omitempty"`
	SchemeName *string `json:"schemeName,omitempty" xml:"schemeName,omitempty"`
	SignName   *string `json:"signName,omitempty" xml:"signName,omitempty"`
}

func (s CreateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRequest) GoString() string {
	return s.String()
}

func (s *CreateRequest) SetAppName(v string) *CreateRequest {
	s.AppName = &v
	return s
}

func (s *CreateRequest) SetBundleId(v string) *CreateRequest {
	s.BundleId = &v
	return s
}

func (s *CreateRequest) SetPackName(v string) *CreateRequest {
	s.PackName = &v
	return s
}

func (s *CreateRequest) SetPlatform(v string) *CreateRequest {
	s.Platform = &v
	return s
}

func (s *CreateRequest) SetSchemeName(v string) *CreateRequest {
	s.SchemeName = &v
	return s
}

func (s *CreateRequest) SetSignName(v string) *CreateRequest {
	s.SignName = &v
	return s
}

type CreateResponseBody struct {
	Code      *string                 `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CreateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Msg       *string                 `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResponseBody) SetCode(v string) *CreateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResponseBody) SetData(v *CreateResponseBodyData) *CreateResponseBody {
	s.Data = v
	return s
}

func (s *CreateResponseBody) SetMsg(v string) *CreateResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateResponseBody) SetRequestId(v string) *CreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResponseBody) SetSuccess(v bool) *CreateResponseBody {
	s.Success = &v
	return s
}

type CreateResponseBodyData struct {
	SchemeCode *string `json:"schemeCode,omitempty" xml:"schemeCode,omitempty"`
}

func (s CreateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateResponseBodyData) SetSchemeCode(v string) *CreateResponseBodyData {
	s.SchemeCode = &v
	return s
}

type CreateResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResponse) GoString() string {
	return s.String()
}

func (s *CreateResponse) SetHeaders(v map[string]*string) *CreateResponse {
	s.Headers = v
	return s
}

func (s *CreateResponse) SetStatusCode(v int32) *CreateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResponse) SetBody(v *CreateResponseBody) *CreateResponse {
	s.Body = v
	return s
}

type GetMobileWithTokenRequest struct {
	ApiCode    *int32  `json:"apiCode,omitempty" xml:"apiCode,omitempty"`
	OperatorId *int32  `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OsType     *string `json:"osType,omitempty" xml:"osType,omitempty"`
	SchemeCode *string `json:"schemeCode,omitempty" xml:"schemeCode,omitempty"`
	Token      *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetMobileWithTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMobileWithTokenRequest) GoString() string {
	return s.String()
}

func (s *GetMobileWithTokenRequest) SetApiCode(v int32) *GetMobileWithTokenRequest {
	s.ApiCode = &v
	return s
}

func (s *GetMobileWithTokenRequest) SetOperatorId(v int32) *GetMobileWithTokenRequest {
	s.OperatorId = &v
	return s
}

func (s *GetMobileWithTokenRequest) SetOsType(v string) *GetMobileWithTokenRequest {
	s.OsType = &v
	return s
}

func (s *GetMobileWithTokenRequest) SetSchemeCode(v string) *GetMobileWithTokenRequest {
	s.SchemeCode = &v
	return s
}

func (s *GetMobileWithTokenRequest) SetToken(v string) *GetMobileWithTokenRequest {
	s.Token = &v
	return s
}

type GetMobileWithTokenResponseBody struct {
	Code      *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data      *GetMobileWithTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Msg       *string                             `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMobileWithTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMobileWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetMobileWithTokenResponseBody) SetCode(v string) *GetMobileWithTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetMobileWithTokenResponseBody) SetData(v *GetMobileWithTokenResponseBodyData) *GetMobileWithTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetMobileWithTokenResponseBody) SetMsg(v string) *GetMobileWithTokenResponseBody {
	s.Msg = &v
	return s
}

func (s *GetMobileWithTokenResponseBody) SetRequestId(v string) *GetMobileWithTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMobileWithTokenResponseBody) SetSuccess(v bool) *GetMobileWithTokenResponseBody {
	s.Success = &v
	return s
}

type GetMobileWithTokenResponseBodyData struct {
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s GetMobileWithTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMobileWithTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMobileWithTokenResponseBodyData) SetMobile(v string) *GetMobileWithTokenResponseBodyData {
	s.Mobile = &v
	return s
}

type GetMobileWithTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMobileWithTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMobileWithTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMobileWithTokenResponse) GoString() string {
	return s.String()
}

func (s *GetMobileWithTokenResponse) SetHeaders(v map[string]*string) *GetMobileWithTokenResponse {
	s.Headers = v
	return s
}

func (s *GetMobileWithTokenResponse) SetStatusCode(v int32) *GetMobileWithTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMobileWithTokenResponse) SetBody(v *GetMobileWithTokenResponseBody) *GetMobileWithTokenResponse {
	s.Body = v
	return s
}

type QueryAppInfoBySchemeRequest struct {
	SchemeCode *string `json:"schemeCode,omitempty" xml:"schemeCode,omitempty"`
}

func (s QueryAppInfoBySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoBySchemeRequest) GoString() string {
	return s.String()
}

func (s *QueryAppInfoBySchemeRequest) SetSchemeCode(v string) *QueryAppInfoBySchemeRequest {
	s.SchemeCode = &v
	return s
}

type QueryAppInfoBySchemeResponseBody struct {
	Code      *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data      *QueryAppInfoBySchemeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Msg       *string                               `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryAppInfoBySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoBySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppInfoBySchemeResponseBody) SetCode(v string) *QueryAppInfoBySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBody) SetData(v *QueryAppInfoBySchemeResponseBodyData) *QueryAppInfoBySchemeResponseBody {
	s.Data = v
	return s
}

func (s *QueryAppInfoBySchemeResponseBody) SetMsg(v string) *QueryAppInfoBySchemeResponseBody {
	s.Msg = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBody) SetRequestId(v string) *QueryAppInfoBySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBody) SetSuccess(v bool) *QueryAppInfoBySchemeResponseBody {
	s.Success = &v
	return s
}

type QueryAppInfoBySchemeResponseBodyData struct {
	CmAppId         *string `json:"cmAppId,omitempty" xml:"cmAppId,omitempty"`
	CmAppKey        *string `json:"cmAppKey,omitempty" xml:"cmAppKey,omitempty"`
	CtAppId         *string `json:"ctAppId,omitempty" xml:"ctAppId,omitempty"`
	CtAppKey        *string `json:"ctAppKey,omitempty" xml:"ctAppKey,omitempty"`
	CuAppId         *string `json:"cuAppId,omitempty" xml:"cuAppId,omitempty"`
	CuAppKey        *string `json:"cuAppKey,omitempty" xml:"cuAppKey,omitempty"`
	CuRsaPublickKey *string `json:"cuRsaPublickKey,omitempty" xml:"cuRsaPublickKey,omitempty"`
}

func (s QueryAppInfoBySchemeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoBySchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCmAppId(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CmAppId = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCmAppKey(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CmAppKey = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCtAppId(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CtAppId = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCtAppKey(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CtAppKey = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCuAppId(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CuAppId = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCuAppKey(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CuAppKey = &v
	return s
}

func (s *QueryAppInfoBySchemeResponseBodyData) SetCuRsaPublickKey(v string) *QueryAppInfoBySchemeResponseBodyData {
	s.CuRsaPublickKey = &v
	return s
}

type QueryAppInfoBySchemeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppInfoBySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppInfoBySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoBySchemeResponse) GoString() string {
	return s.String()
}

func (s *QueryAppInfoBySchemeResponse) SetHeaders(v map[string]*string) *QueryAppInfoBySchemeResponse {
	s.Headers = v
	return s
}

func (s *QueryAppInfoBySchemeResponse) SetStatusCode(v int32) *QueryAppInfoBySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppInfoBySchemeResponse) SetBody(v *QueryAppInfoBySchemeResponseBody) *QueryAppInfoBySchemeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("umeng-verify-agent"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateWithOptions(request *CreateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		body["bundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.PackName)) {
		body["packName"] = request.PackName
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		body["schemeName"] = request.SchemeName
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		body["signName"] = request.SignName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Create"),
		Version:     tea.String("2024-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/Create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Create(request *CreateRequest) (_result *CreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResponse{}
	_body, _err := client.CreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMobileWithTokenWithOptions(request *GetMobileWithTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMobileWithTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiCode)) {
		body["apiCode"] = request.ApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		body["osType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		body["schemeCode"] = request.SchemeCode
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMobileWithToken"),
		Version:     tea.String("2024-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/GetMobileWithToken"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMobileWithTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMobileWithToken(request *GetMobileWithTokenRequest) (_result *GetMobileWithTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMobileWithTokenResponse{}
	_body, _err := client.GetMobileWithTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAppInfoBySchemeWithOptions(request *QueryAppInfoBySchemeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryAppInfoBySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		body["schemeCode"] = request.SchemeCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAppInfoByScheme"),
		Version:     tea.String("2024-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/QueryAppInfoByScheme"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAppInfoBySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAppInfoByScheme(request *QueryAppInfoBySchemeRequest) (_result *QueryAppInfoBySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAppInfoBySchemeResponse{}
	_body, _err := client.QueryAppInfoBySchemeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
