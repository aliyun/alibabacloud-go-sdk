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

type RegisterDeviceRequest struct {
	AppKey       *string                `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId     *string                `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Extend       map[string]interface{} `json:"Extend,omitempty" xml:"Extend,omitempty"`
	SdkCode      *string                `json:"SdkCode,omitempty" xml:"SdkCode,omitempty"`
	UserDeviceId *string                `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetAppKey(v string) *RegisterDeviceRequest {
	s.AppKey = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceId(v string) *RegisterDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterDeviceRequest) SetExtend(v map[string]interface{}) *RegisterDeviceRequest {
	s.Extend = v
	return s
}

func (s *RegisterDeviceRequest) SetSdkCode(v string) *RegisterDeviceRequest {
	s.SdkCode = &v
	return s
}

func (s *RegisterDeviceRequest) SetUserDeviceId(v string) *RegisterDeviceRequest {
	s.UserDeviceId = &v
	return s
}

type RegisterDeviceShrinkRequest struct {
	AppKey       *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExtendShrink *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	SdkCode      *string `json:"SdkCode,omitempty" xml:"SdkCode,omitempty"`
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s RegisterDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceShrinkRequest) SetAppKey(v string) *RegisterDeviceShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *RegisterDeviceShrinkRequest) SetDeviceId(v string) *RegisterDeviceShrinkRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterDeviceShrinkRequest) SetExtendShrink(v string) *RegisterDeviceShrinkRequest {
	s.ExtendShrink = &v
	return s
}

func (s *RegisterDeviceShrinkRequest) SetSdkCode(v string) *RegisterDeviceShrinkRequest {
	s.SdkCode = &v
	return s
}

func (s *RegisterDeviceShrinkRequest) SetUserDeviceId(v string) *RegisterDeviceShrinkRequest {
	s.UserDeviceId = &v
	return s
}

type RegisterDeviceResponseBody struct {
	Data         *RegisterDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *int32                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message      *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetData(v *RegisterDeviceResponseBodyData) *RegisterDeviceResponseBody {
	s.Data = v
	return s
}

func (s *RegisterDeviceResponseBody) SetErrorCode(v int32) *RegisterDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetErrorMessage(v string) *RegisterDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

type RegisterDeviceResponseBodyData struct {
	License   *string `json:"License,omitempty" xml:"License,omitempty"`
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	Rid       *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s RegisterDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyData) SetLicense(v string) *RegisterDeviceResponseBodyData {
	s.License = &v
	return s
}

func (s *RegisterDeviceResponseBodyData) SetPublicKey(v string) *RegisterDeviceResponseBodyData {
	s.PublicKey = &v
	return s
}

func (s *RegisterDeviceResponseBodyData) SetRid(v string) *RegisterDeviceResponseBodyData {
	s.Rid = &v
	return s
}

func (s *RegisterDeviceResponseBodyData) SetSignature(v string) *RegisterDeviceResponseBodyData {
	s.Signature = &v
	return s
}

type RegisterDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetStatusCode(v int32) *RegisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("moguan-sdk"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) RegisterDeviceWithOptions(tmpReq *RegisterDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RegisterDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extend)) {
		request.ExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extend, tea.String("Extend"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendShrink)) {
		body["Extend"] = request.ExtendShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SdkCode)) {
		body["SdkCode"] = request.SdkCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeviceId)) {
		body["UserDeviceId"] = request.UserDeviceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDevice"),
		Version:     tea.String("2021-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
