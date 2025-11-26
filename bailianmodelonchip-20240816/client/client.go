// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("bailianmodelonchip"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 主动交互消息传递
//
// @param request - ActiveInteractionCreateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveInteractionCreateResponse
func (client *Client) ActiveInteractionCreateWithOptions(request *ActiveInteractionCreateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActiveInteractionCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveInteractionCreate"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/v1/active/interaction/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveInteractionCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 主动交互消息传递
//
// @param request - ActiveInteractionCreateRequest
//
// @return ActiveInteractionCreateResponse
func (client *Client) ActiveInteractionCreate(request *ActiveInteractionCreateRequest) (_result *ActiveInteractionCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActiveInteractionCreateResponse{}
	_body, _err := client.ActiveInteractionCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 主动交互消息生成eu
//
// @param request - ActiveInteractionEuCreateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveInteractionEuCreateResponse
func (client *Client) ActiveInteractionEuCreateWithOptions(request *ActiveInteractionEuCreateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActiveInteractionEuCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveInteractionEuCreate"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/eu/active/interaction/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveInteractionEuCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 主动交互消息生成eu
//
// @param request - ActiveInteractionEuCreateRequest
//
// @return ActiveInteractionEuCreateResponse
func (client *Client) ActiveInteractionEuCreate(request *ActiveInteractionEuCreateRequest) (_result *ActiveInteractionEuCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActiveInteractionEuCreateResponse{}
	_body, _err := client.ActiveInteractionEuCreateWithOptions(request, headers, runtime)
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
func (client *Client) DeviceRegisterWithOptions(request *DeviceRegisterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeviceRegisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Nonce) {
		body["nonce"] = request.Nonce
	}

	if !dara.IsNil(request.RequestTime) {
		body["requestTime"] = request.RequestTime
	}

	if !dara.IsNil(request.Signature) {
		body["signature"] = request.Signature
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeviceRegister"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/device/v1/register"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeviceRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.DeviceName) {
		body["deviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.Nonce) {
		body["nonce"] = request.Nonce
	}

	if !dara.IsNil(request.RequestTime) {
		body["requestTime"] = request.RequestTime
	}

	if !dara.IsNil(request.Signature) {
		body["signature"] = request.Signature
	}

	if !dara.IsNil(request.TokenKey) {
		body["tokenKey"] = request.TokenKey
	}

	if !dara.IsNil(request.TokenType) {
		body["tokenType"] = request.TokenType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/auth/v1/token/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模型类型识别
//
// @param tmpReq - ModelTypeDetermineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelTypeDetermineResponse
func (client *Client) ModelTypeDetermineWithOptions(tmpReq *ModelTypeDetermineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelTypeDetermineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModelTypeDetermineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.History) {
		request.HistoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.History, dara.String("history"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HistoryShrink) {
		body["history"] = request.HistoryShrink
	}

	if !dara.IsNil(request.InputText) {
		body["inputText"] = request.InputText
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelTypeDetermine"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/v1/model/type/determine"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelTypeDetermineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型类型识别
//
// @param request - ModelTypeDetermineRequest
//
// @return ModelTypeDetermineResponse
func (client *Client) ModelTypeDetermine(request *ModelTypeDetermineRequest) (_result *ModelTypeDetermineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModelTypeDetermineResponse{}
	_body, _err := client.ModelTypeDetermineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 音频-供机械臂调用
//
// @param request - OmniRealtimeConversationEURequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OmniRealtimeConversationEUResponse
func (client *Client) OmniRealtimeConversationEUWithOptions(request *OmniRealtimeConversationEURequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OmniRealtimeConversationEUResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InputAudio) {
		body["inputAudio"] = request.InputAudio
	}

	if !dara.IsNil(request.UserPrompt) {
		body["userPrompt"] = request.UserPrompt
	}

	if !dara.IsNil(request.Voice) {
		body["voice"] = request.Voice
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OmniRealtimeConversationEU"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/eu/active/interaction/audio"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OmniRealtimeConversationEUResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 音频-供机械臂调用
//
// @param request - OmniRealtimeConversationEURequest
//
// @return OmniRealtimeConversationEUResponse
func (client *Client) OmniRealtimeConversationEU(request *OmniRealtimeConversationEURequest) (_result *OmniRealtimeConversationEUResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OmniRealtimeConversationEUResponse{}
	_body, _err := client.OmniRealtimeConversationEUWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
