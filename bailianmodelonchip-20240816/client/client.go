// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DeviceRegisterRequest struct {
	// This parameter is required.
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a64edd96296880f55aa61987b
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// This parameter is required.
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// This parameter is required.
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s DeviceRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterRequest) GoString() string {
	return s.String()
}

func (s *DeviceRegisterRequest) SetAppId(v string) *DeviceRegisterRequest {
	s.AppId = &v
	return s
}

func (s *DeviceRegisterRequest) SetNonce(v string) *DeviceRegisterRequest {
	s.Nonce = &v
	return s
}

func (s *DeviceRegisterRequest) SetRequestTime(v string) *DeviceRegisterRequest {
	s.RequestTime = &v
	return s
}

func (s *DeviceRegisterRequest) SetSignature(v string) *DeviceRegisterRequest {
	s.Signature = &v
	return s
}

type DeviceRegisterResponseBody struct {
	// example:
	//
	// success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeviceRegisterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 52548360-B3AA-55EA-893F-48C16470F64A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeviceRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponseBody) SetCode(v string) *DeviceRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetData(v *DeviceRegisterResponseBodyData) *DeviceRegisterResponseBody {
	s.Data = v
	return s
}

func (s *DeviceRegisterResponseBody) SetHttpStatusCode(v int32) *DeviceRegisterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetMessage(v string) *DeviceRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetRequestId(v string) *DeviceRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceRegisterResponseBody) SetSuccess(v bool) *DeviceRegisterResponseBody {
	s.Success = &v
	return s
}

type DeviceRegisterResponseBodyData struct {
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 991fa52b7935aaa33536e05d4f4b5003
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// e2e928e8244f45ab30ec6ba9f9
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// 1748312544852
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// example:
	//
	// s8wPO/w79jP7sz6OaHkixAje2GmgzmZiCuCZZ+J8w77ICTyqD30lL6rUhnXwwx4MyGF62DRPFnpXtJ6c5mlmt6QdML3FkjGn+i/wR5T6QMkVDW6YRPWsx3jkIWRQ2sDnmVNBtixo2s9w3RJrnddRzVCaR/WeLOfiVLWcrLcJditzO/1YXBZ9vuRKQ4iperfhZvw372N/m8/1qtjJl+FUe2+wxK6RMxr03K7R
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s DeviceRegisterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponseBodyData) SetAppId(v string) *DeviceRegisterResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetDeviceName(v string) *DeviceRegisterResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetNonce(v string) *DeviceRegisterResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetResponseTime(v string) *DeviceRegisterResponseBodyData {
	s.ResponseTime = &v
	return s
}

func (s *DeviceRegisterResponseBodyData) SetSignature(v string) *DeviceRegisterResponseBodyData {
	s.Signature = &v
	return s
}

type DeviceRegisterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeviceRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeviceRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceRegisterResponse) GoString() string {
	return s.String()
}

func (s *DeviceRegisterResponse) SetHeaders(v map[string]*string) *DeviceRegisterResponse {
	s.Headers = v
	return s
}

func (s *DeviceRegisterResponse) SetStatusCode(v int32) *DeviceRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceRegisterResponse) SetBody(v *DeviceRegisterResponseBody) *DeviceRegisterResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	// This parameter is required.
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a64edd96296880f55aa61987b
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1748413148546
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5/Smm8gnDSgZY2Blftq9eGYpVnpYCztoLJaJfhlH7id0lNlQxydRLtjUkGPr1qdbQq+oUn6Y1h0KJUdk0rf4am3MzdNr/Uhc47c8bbXnV8SlIC0agGo5skEQZNObzUD+sFxt8uN35/FYf7YRC8R61xY7+NPN2NLJrA/DPhewtVRRgAbb8RjergTcIG6oN1XTzLyC+76Az/3o2dPCxTfMXG3AFQc=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// sk-4asv136677d2411b876e536bc8xxxxx
	TokenKey *string `json:"tokenKey,omitempty" xml:"tokenKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	TokenType *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetAppId(v string) *GetTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetTokenRequest) SetDeviceName(v string) *GetTokenRequest {
	s.DeviceName = &v
	return s
}

func (s *GetTokenRequest) SetNonce(v string) *GetTokenRequest {
	s.Nonce = &v
	return s
}

func (s *GetTokenRequest) SetRequestTime(v string) *GetTokenRequest {
	s.RequestTime = &v
	return s
}

func (s *GetTokenRequest) SetSignature(v string) *GetTokenRequest {
	s.Signature = &v
	return s
}

func (s *GetTokenRequest) SetTokenKey(v string) *GetTokenRequest {
	s.TokenKey = &v
	return s
}

func (s *GetTokenRequest) SetTokenType(v string) *GetTokenRequest {
	s.TokenType = &v
	return s
}

type GetTokenResponseBody struct {
	// example:
	//
	// success
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// B08AAA14-AD93-51F6-82AE-82AFAE9375B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetHttpStatusCode(v string) *GetTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v string) *GetTokenResponseBody {
	s.Success = &v
	return s
}

type GetTokenResponseBodyData struct {
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// b79d692c315d6bfb28312edf15
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// example:
	//
	// 127.0.0.1
	RequestIp *string `json:"requestIp,omitempty" xml:"requestIp,omitempty"`
	// example:
	//
	// 1748413248360
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	// example:
	//
	// N1faAjFhhaRNFaZNC8woRpQyAzEfBaIoWQEgDfds/Fwm7nIyEDLlSK3Ttx2OFebiHZ/MpHRr/3MnI/jpVWB/xNYIQxm6sccHJENHNAz6gaW+itU5wUrh+46EpqySABV8kc2pQ0HmYlbePfjjOK6lCfQjEGpekSAgQ6tDhG1lXWfKdtggq58Ut5bImMxMhk4R/PFUWrJe4CDuFu072C+foI0JlUV9TnGtVQ58oz8VRndrGXyauS/xqg8iGSZn6FyprUf5p+0ow20E
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetAppId(v string) *GetTokenResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetTokenResponseBodyData) SetDeviceName(v string) *GetTokenResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetTokenResponseBodyData) SetNonce(v string) *GetTokenResponseBodyData {
	s.Nonce = &v
	return s
}

func (s *GetTokenResponseBodyData) SetRequestIp(v string) *GetTokenResponseBodyData {
	s.RequestIp = &v
	return s
}

func (s *GetTokenResponseBodyData) SetResponseTime(v string) *GetTokenResponseBodyData {
	s.ResponseTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSignature(v string) *GetTokenResponseBodyData {
	s.Signature = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailianmodelonchip"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 设备注册
//
// @param request - DeviceRegisterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceRegisterResponse
func (client *Client) DeviceRegisterWithOptions(request *DeviceRegisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeviceRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Nonce)) {
		body["nonce"] = request.Nonce
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTime)) {
		body["requestTime"] = request.RequestTime
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceRegister"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/device/v1/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - DeviceRegisterRequest
//
// @return DeviceRegisterResponse
func (client *Client) DeviceRegister(request *DeviceRegisterRequest) (_result *DeviceRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeviceRegisterResponse{}
	_body, _err := client.DeviceRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网关校验Token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.Nonce)) {
		body["nonce"] = request.Nonce
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTime)) {
		body["requestTime"] = request.RequestTime
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.TokenKey)) {
		body["tokenKey"] = request.TokenKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenType)) {
		body["tokenType"] = request.TokenType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2024-08-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/open/api/auth/v1/token/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网关校验Token
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
