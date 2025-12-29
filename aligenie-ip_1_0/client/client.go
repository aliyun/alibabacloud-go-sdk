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
	client.Endpoint, _err = client.GetEndpoint(dara.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加动画
//
// @param request - AddCartoonRequest
//
// @param headers - AddCartoonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCartoonResponse
func (client *Client) AddCartoonWithOptions(request *AddCartoonRequest, headers *AddCartoonHeaders, runtime *dara.RuntimeOptions) (_result *AddCartoonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.StartVideoMd5) {
		body["StartVideoMd5"] = request.StartVideoMd5
	}

	if !dara.IsNil(request.StartVideoUrl) {
		body["StartVideoUrl"] = request.StartVideoUrl
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCartoon"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addCartoon"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCartoonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加动画
//
// @param request - AddCartoonRequest
//
// @return AddCartoonResponse
func (client *Client) AddCartoon(request *AddCartoonRequest) (_result *AddCartoonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddCartoonHeaders{}
	_result = &AddCartoonResponse{}
	_body, _err := client.AddCartoonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增自定义问答
//
// @param tmpReq - AddCustomQARequest
//
// @param headers - AddCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomQAResponse
func (client *Client) AddCustomQAWithOptions(tmpReq *AddCustomQARequest, headers *AddCustomQAHeaders, runtime *dara.RuntimeOptions) (_result *AddCustomQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Answers) {
		request.AnswersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Answers, dara.String("Answers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KeyWords) {
		request.KeyWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KeyWords, dara.String("KeyWords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SupplementaryQuestions) {
		request.SupplementaryQuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupplementaryQuestions, dara.String("SupplementaryQuestions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnswersShrink) {
		body["Answers"] = request.AnswersShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.KeyWordsShrink) {
		body["KeyWords"] = request.KeyWordsShrink
	}

	if !dara.IsNil(request.MajorQuestion) {
		body["MajorQuestion"] = request.MajorQuestion
	}

	if !dara.IsNil(request.SupplementaryQuestionsShrink) {
		body["SupplementaryQuestions"] = request.SupplementaryQuestionsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addCustomQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增自定义问答
//
// @param request - AddCustomQARequest
//
// @return AddCustomQAResponse
func (client *Client) AddCustomQA(request *AddCustomQARequest) (_result *AddCustomQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddCustomQAHeaders{}
	_result = &AddCustomQAResponse{}
	_body, _err := client.AddCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加问答V2
//
// @param tmpReq - AddCustomQAV2Request
//
// @param headers - AddCustomQAV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomQAV2Response
func (client *Client) AddCustomQAV2WithOptions(tmpReq *AddCustomQAV2Request, headers *AddCustomQAV2Headers, runtime *dara.RuntimeOptions) (_result *AddCustomQAV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddCustomQAV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Answers) {
		request.AnswersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Answers, dara.String("Answers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KeyWords) {
		request.KeyWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KeyWords, dara.String("KeyWords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SupplementaryQuestions) {
		request.SupplementaryQuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupplementaryQuestions, dara.String("SupplementaryQuestions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnswersShrink) {
		body["Answers"] = request.AnswersShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.KeyWordsShrink) {
		body["KeyWords"] = request.KeyWordsShrink
	}

	if !dara.IsNil(request.MajorQuestion) {
		body["MajorQuestion"] = request.MajorQuestion
	}

	if !dara.IsNil(request.SupplementaryQuestionsShrink) {
		body["SupplementaryQuestions"] = request.SupplementaryQuestionsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomQAV2"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addQAV2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomQAV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加问答V2
//
// @param request - AddCustomQAV2Request
//
// @return AddCustomQAV2Response
func (client *Client) AddCustomQAV2(request *AddCustomQAV2Request) (_result *AddCustomQAV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddCustomQAV2Headers{}
	_result = &AddCustomQAV2Response{}
	_body, _err := client.AddCustomQAV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加消息模板
//
// @param request - AddMessageTemplateRequest
//
// @param headers - AddMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMessageTemplateResponse
func (client *Client) AddMessageTemplateWithOptions(request *AddMessageTemplateRequest, headers *AddMessageTemplateHeaders, runtime *dara.RuntimeOptions) (_result *AddMessageTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateDetail) {
		body["TemplateDetail"] = request.TemplateDetail
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMessageTemplate"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addMessageTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加消息模板
//
// @param request - AddMessageTemplateRequest
//
// @return AddMessageTemplateResponse
func (client *Client) AddMessageTemplate(request *AddMessageTemplateRequest) (_result *AddMessageTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddMessageTemplateHeaders{}
	_result = &AddMessageTemplateResponse{}
	_body, _err := client.AddMessageTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏展示模式
//
// @param tmpReq - AddOrUpdateDisPlayModesRequest
//
// @param headers - AddOrUpdateDisPlayModesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateDisPlayModesResponse
func (client *Client) AddOrUpdateDisPlayModesWithOptions(tmpReq *AddOrUpdateDisPlayModesRequest, headers *AddOrUpdateDisPlayModesHeaders, runtime *dara.RuntimeOptions) (_result *AddOrUpdateDisPlayModesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddOrUpdateDisPlayModesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelDeviceModeList) {
		request.HotelDeviceModeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelDeviceModeList, dara.String("HotelDeviceModeList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelDeviceModeListShrink) {
		body["HotelDeviceModeList"] = request.HotelDeviceModeListShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrUpdateDisPlayModes"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addOrUpdateDisPlayModes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrUpdateDisPlayModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏展示模式
//
// @param request - AddOrUpdateDisPlayModesRequest
//
// @return AddOrUpdateDisPlayModesResponse
func (client *Client) AddOrUpdateDisPlayModes(request *AddOrUpdateDisPlayModesRequest) (_result *AddOrUpdateDisPlayModesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddOrUpdateDisPlayModesHeaders{}
	_result = &AddOrUpdateDisPlayModesResponse{}
	_body, _err := client.AddOrUpdateDisPlayModesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或者编辑定制配置
//
// @param tmpReq - AddOrUpdateHotelSettingRequest
//
// @param headers - AddOrUpdateHotelSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateHotelSettingResponse
func (client *Client) AddOrUpdateHotelSettingWithOptions(tmpReq *AddOrUpdateHotelSettingRequest, headers *AddOrUpdateHotelSettingHeaders, runtime *dara.RuntimeOptions) (_result *AddOrUpdateHotelSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddOrUpdateHotelSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelDeviceModeList) {
		request.HotelDeviceModeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelDeviceModeList, dara.String("HotelDeviceModeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HotelScreenSaver) {
		request.HotelScreenSaverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelScreenSaver, dara.String("HotelScreenSaver"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NightMode) {
		request.NightModeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NightMode, dara.String("NightMode"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelDeviceModeListShrink) {
		body["HotelDeviceModeList"] = request.HotelDeviceModeListShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.HotelScreenSaverShrink) {
		body["HotelScreenSaver"] = request.HotelScreenSaverShrink
	}

	if !dara.IsNil(request.NightModeShrink) {
		body["NightMode"] = request.NightModeShrink
	}

	if !dara.IsNil(request.SettingType) {
		body["SettingType"] = request.SettingType
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrUpdateHotelSetting"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addOrUpdateHotelSetting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrUpdateHotelSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或者编辑定制配置
//
// @param request - AddOrUpdateHotelSettingRequest
//
// @return AddOrUpdateHotelSettingResponse
func (client *Client) AddOrUpdateHotelSetting(request *AddOrUpdateHotelSettingRequest) (_result *AddOrUpdateHotelSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddOrUpdateHotelSettingHeaders{}
	_result = &AddOrUpdateHotelSettingResponse{}
	_body, _err := client.AddOrUpdateHotelSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏屏保
//
// @param tmpReq - AddOrUpdateScreenSaverRequest
//
// @param headers - AddOrUpdateScreenSaverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateScreenSaverResponse
func (client *Client) AddOrUpdateScreenSaverWithOptions(tmpReq *AddOrUpdateScreenSaverRequest, headers *AddOrUpdateScreenSaverHeaders, runtime *dara.RuntimeOptions) (_result *AddOrUpdateScreenSaverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddOrUpdateScreenSaverShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelScreenSaver) {
		request.HotelScreenSaverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelScreenSaver, dara.String("HotelScreenSaver"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.HotelScreenSaverShrink) {
		body["HotelScreenSaver"] = request.HotelScreenSaverShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrUpdateScreenSaver"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addOrUpdateScreenSaver"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrUpdateScreenSaverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或者编辑带屏屏保
//
// @param request - AddOrUpdateScreenSaverRequest
//
// @return AddOrUpdateScreenSaverResponse
func (client *Client) AddOrUpdateScreenSaver(request *AddOrUpdateScreenSaverRequest) (_result *AddOrUpdateScreenSaverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddOrUpdateScreenSaverHeaders{}
	_result = &AddOrUpdateScreenSaverResponse{}
	_body, _err := client.AddOrUpdateScreenSaverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加/更新欢迎语信息
//
// @param request - AddOrUpdateWelcomeTextRequest
//
// @param headers - AddOrUpdateWelcomeTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateWelcomeTextResponse
func (client *Client) AddOrUpdateWelcomeTextWithOptions(request *AddOrUpdateWelcomeTextRequest, headers *AddOrUpdateWelcomeTextHeaders, runtime *dara.RuntimeOptions) (_result *AddOrUpdateWelcomeTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.MusicUrl) {
		body["MusicUrl"] = request.MusicUrl
	}

	if !dara.IsNil(request.WelcomeText) {
		body["WelcomeText"] = request.WelcomeText
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrUpdateWelcomeText"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/addOrUpdateWelcomeText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrUpdateWelcomeTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加/更新欢迎语信息
//
// @param request - AddOrUpdateWelcomeTextRequest
//
// @return AddOrUpdateWelcomeTextResponse
func (client *Client) AddOrUpdateWelcomeText(request *AddOrUpdateWelcomeTextRequest) (_result *AddOrUpdateWelcomeTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddOrUpdateWelcomeTextHeaders{}
	_result = &AddOrUpdateWelcomeTextResponse{}
	_body, _err := client.AddOrUpdateWelcomeTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 审批酒店
//
// @param tmpReq - AuditHotelRequest
//
// @param headers - AuditHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditHotelResponse
func (client *Client) AuditHotelWithOptions(tmpReq *AuditHotelRequest, headers *AuditHotelHeaders, runtime *dara.RuntimeOptions) (_result *AuditHotelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuditHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuditHotelReq) {
		request.AuditHotelReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuditHotelReq, dara.String("AuditHotelReq"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditHotelReqShrink) {
		query["AuditHotelReq"] = request.AuditHotelReqShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuditHotel"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/auditHotel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuditHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 审批酒店
//
// @param request - AuditHotelRequest
//
// @return AuditHotelResponse
func (client *Client) AuditHotel(request *AuditHotelRequest) (_result *AuditHotelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AuditHotelHeaders{}
	_result = &AuditHotelResponse{}
	_body, _err := client.AuditHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建房间
//
// @param tmpReq - BatchAddHotelRoomRequest
//
// @param headers - BatchAddHotelRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddHotelRoomResponse
func (client *Client) BatchAddHotelRoomWithOptions(tmpReq *BatchAddHotelRoomRequest, headers *BatchAddHotelRoomHeaders, runtime *dara.RuntimeOptions) (_result *BatchAddHotelRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchAddHotelRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoomNoList) {
		request.RoomNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomNoList, dara.String("RoomNoList"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNoListShrink) {
		body["RoomNoList"] = request.RoomNoListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddHotelRoom"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/batchAddHotelRoom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddHotelRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建房间
//
// @param request - BatchAddHotelRoomRequest
//
// @return BatchAddHotelRoomResponse
func (client *Client) BatchAddHotelRoom(request *BatchAddHotelRoomRequest) (_result *BatchAddHotelRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchAddHotelRoomHeaders{}
	_result = &BatchAddHotelRoomResponse{}
	_body, _err := client.BatchAddHotelRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除房间
//
// @param tmpReq - BatchDeleteHotelRoomRequest
//
// @param headers - BatchDeleteHotelRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteHotelRoomResponse
func (client *Client) BatchDeleteHotelRoomWithOptions(tmpReq *BatchDeleteHotelRoomRequest, headers *BatchDeleteHotelRoomHeaders, runtime *dara.RuntimeOptions) (_result *BatchDeleteHotelRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteHotelRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoomNoList) {
		request.RoomNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomNoList, dara.String("RoomNoList"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNoListShrink) {
		body["RoomNoList"] = request.RoomNoListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteHotelRoom"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/batchDeleteHotelRoom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteHotelRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除房间
//
// @param request - BatchDeleteHotelRoomRequest
//
// @return BatchDeleteHotelRoomResponse
func (client *Client) BatchDeleteHotelRoom(request *BatchDeleteHotelRoomRequest) (_result *BatchDeleteHotelRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchDeleteHotelRoomHeaders{}
	_result = &BatchDeleteHotelRoomResponse{}
	_body, _err := client.BatchDeleteHotelRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店退房，清楚例如闹钟等定时信息
//
// @param request - CheckoutWithAKRequest
//
// @param headers - CheckoutWithAKHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckoutWithAKResponse
func (client *Client) CheckoutWithAKWithOptions(request *CheckoutWithAKRequest, headers *CheckoutWithAKHeaders, runtime *dara.RuntimeOptions) (_result *CheckoutWithAKResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckoutWithAK"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/checkoutWithAK"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckoutWithAKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店退房，清楚例如闹钟等定时信息
//
// @param request - CheckoutWithAKRequest
//
// @return CheckoutWithAKResponse
func (client *Client) CheckoutWithAK(request *CheckoutWithAKRequest) (_result *CheckoutWithAKResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CheckoutWithAKHeaders{}
	_result = &CheckoutWithAKResponse{}
	_body, _err := client.CheckoutWithAKWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 子账号授权
//
// @param request - ChildAccountAuthRequest
//
// @param headers - ChildAccountAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChildAccountAuthResponse
func (client *Client) ChildAccountAuthWithOptions(request *ChildAccountAuthRequest, headers *ChildAccountAuthHeaders, runtime *dara.RuntimeOptions) (_result *ChildAccountAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		body["Account"] = request.Account
	}

	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.TbOpenId) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChildAccountAuth"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/childAccountAuth"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChildAccountAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 子账号授权
//
// @param request - ChildAccountAuthRequest
//
// @return ChildAccountAuthResponse
func (client *Client) ChildAccountAuth(request *ChildAccountAuthRequest) (_result *ChildAccountAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChildAccountAuthHeaders{}
	_result = &ChildAccountAuthResponse{}
	_body, _err := client.ChildAccountAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 控制房间内设备
//
// @param tmpReq - ControlRoomDeviceRequest
//
// @param headers - ControlRoomDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ControlRoomDeviceResponse
func (client *Client) ControlRoomDeviceWithOptions(tmpReq *ControlRoomDeviceRequest, headers *ControlRoomDeviceHeaders, runtime *dara.RuntimeOptions) (_result *ControlRoomDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ControlRoomDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Properties) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, dara.String("Properties"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cmd) {
		body["Cmd"] = request.Cmd
	}

	if !dara.IsNil(request.DeviceIndex) {
		body["DeviceIndex"] = request.DeviceIndex
	}

	if !dara.IsNil(request.DeviceNumber) {
		body["DeviceNumber"] = request.DeviceNumber
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.PropertiesShrink) {
		body["Properties"] = request.PropertiesShrink
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ControlRoomDevice"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/controlRoomDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ControlRoomDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制房间内设备
//
// @param request - ControlRoomDeviceRequest
//
// @return ControlRoomDeviceResponse
func (client *Client) ControlRoomDevice(request *ControlRoomDeviceRequest) (_result *ControlRoomDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ControlRoomDeviceHeaders{}
	_result = &ControlRoomDeviceResponse{}
	_body, _err := client.ControlRoomDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建酒店项目
//
// @param tmpReq - CreateHotelRequest
//
// @param headers - CreateHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHotelResponse
func (client *Client) CreateHotelWithOptions(tmpReq *CreateHotelRequest, headers *CreateHotelHeaders, runtime *dara.RuntimeOptions) (_result *CreateHotelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RelatedPks) {
		request.RelatedPksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedPks, dara.String("RelatedPks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EstOpenTime) {
		body["EstOpenTime"] = request.EstOpenTime
	}

	if !dara.IsNil(request.HotelAddress) {
		body["HotelAddress"] = request.HotelAddress
	}

	if !dara.IsNil(request.HotelEmail) {
		body["HotelEmail"] = request.HotelEmail
	}

	if !dara.IsNil(request.HotelName) {
		body["HotelName"] = request.HotelName
	}

	if !dara.IsNil(request.PhoneNumber) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.RelatedPk) {
		body["RelatedPk"] = request.RelatedPk
	}

	if !dara.IsNil(request.RelatedPksShrink) {
		body["RelatedPks"] = request.RelatedPksShrink
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RoomNum) {
		body["RoomNum"] = request.RoomNum
	}

	if !dara.IsNil(request.TbOpenId) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHotel"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/createHotel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建酒店项目
//
// @param request - CreateHotelRequest
//
// @return CreateHotelResponse
func (client *Client) CreateHotel(request *CreateHotelRequest) (_result *CreateHotelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateHotelHeaders{}
	_result = &CreateHotelResponse{}
	_body, _err := client.CreateHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建闹钟
//
// @param tmpReq - CreateHotelAlarmRequest
//
// @param headers - CreateHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHotelAlarmResponse
func (client *Client) CreateHotelAlarmWithOptions(tmpReq *CreateHotelAlarmRequest, headers *CreateHotelAlarmHeaders, runtime *dara.RuntimeOptions) (_result *CreateHotelAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rooms) {
		request.RoomsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rooms, dara.String("Rooms"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleInfo) {
		request.ScheduleInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleInfo, dara.String("ScheduleInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.MusicType) {
		body["MusicType"] = request.MusicType
	}

	if !dara.IsNil(request.RoomsShrink) {
		body["Rooms"] = request.RoomsShrink
	}

	if !dara.IsNil(request.ScheduleInfoShrink) {
		body["ScheduleInfo"] = request.ScheduleInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHotelAlarm"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/createHotelAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建闹钟
//
// @param request - CreateHotelAlarmRequest
//
// @return CreateHotelAlarmResponse
func (client *Client) CreateHotelAlarm(request *CreateHotelAlarmRequest) (_result *CreateHotelAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateHotelAlarmHeaders{}
	_result = &CreateHotelAlarmResponse{}
	_body, _err := client.CreateHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店rcu自定义场景创建
//
// @param tmpReq - CreateRcuSceneRequest
//
// @param headers - CreateRcuSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRcuSceneResponse
func (client *Client) CreateRcuSceneWithOptions(tmpReq *CreateRcuSceneRequest, headers *CreateRcuSceneHeaders, runtime *dara.RuntimeOptions) (_result *CreateRcuSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRcuSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SceneRelationExtDTO) {
		request.SceneRelationExtDTOShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneRelationExtDTO, dara.String("SceneRelationExtDTO"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SceneRelationExtDTOShrink) {
		body["SceneRelationExtDTO"] = request.SceneRelationExtDTOShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRcuScene"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/createRcuScene"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRcuSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店rcu自定义场景创建
//
// @param request - CreateRcuSceneRequest
//
// @return CreateRcuSceneResponse
func (client *Client) CreateRcuScene(request *CreateRcuSceneRequest) (_result *CreateRcuSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateRcuSceneHeaders{}
	_result = &CreateRcuSceneResponse{}
	_body, _err := client.CreateRcuSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除动画
//
// @param request - DeleteCartoonRequest
//
// @param headers - DeleteCartoonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCartoonResponse
func (client *Client) DeleteCartoonWithOptions(request *DeleteCartoonRequest, headers *DeleteCartoonHeaders, runtime *dara.RuntimeOptions) (_result *DeleteCartoonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCartoon"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteCartoon"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCartoonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除动画
//
// @param request - DeleteCartoonRequest
//
// @return DeleteCartoonResponse
func (client *Client) DeleteCartoon(request *DeleteCartoonRequest) (_result *DeleteCartoonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteCartoonHeaders{}
	_result = &DeleteCartoonResponse{}
	_body, _err := client.DeleteCartoonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义问答
//
// @param tmpReq - DeleteCustomQARequest
//
// @param headers - DeleteCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomQAResponse
func (client *Client) DeleteCustomQAWithOptions(tmpReq *DeleteCustomQARequest, headers *DeleteCustomQAHeaders, runtime *dara.RuntimeOptions) (_result *DeleteCustomQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomQAIds) {
		request.CustomQAIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomQAIds, dara.String("CustomQAIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomQAIdsShrink) {
		body["CustomQAIds"] = request.CustomQAIdsShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteCustomQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义问答
//
// @param request - DeleteCustomQARequest
//
// @return DeleteCustomQAResponse
func (client *Client) DeleteCustomQA(request *DeleteCustomQARequest) (_result *DeleteCustomQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteCustomQAHeaders{}
	_result = &DeleteCustomQAResponse{}
	_body, _err := client.DeleteCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除酒店闹钟
//
// @param tmpReq - DeleteHotelAlarmRequest
//
// @param headers - DeleteHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotelAlarmResponse
func (client *Client) DeleteHotelAlarmWithOptions(tmpReq *DeleteHotelAlarmRequest, headers *DeleteHotelAlarmHeaders, runtime *dara.RuntimeOptions) (_result *DeleteHotelAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Alarms) {
		request.AlarmsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Alarms, dara.String("Alarms"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmsShrink) {
		body["Alarms"] = request.AlarmsShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHotelAlarm"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteHotelAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除酒店闹钟
//
// @param request - DeleteHotelAlarmRequest
//
// @return DeleteHotelAlarmResponse
func (client *Client) DeleteHotelAlarm(request *DeleteHotelAlarmRequest) (_result *DeleteHotelAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteHotelAlarmHeaders{}
	_result = &DeleteHotelAlarmResponse{}
	_body, _err := client.DeleteHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订删除
//
// @param request - DeleteHotelSceneBookItemRequest
//
// @param headers - DeleteHotelSceneBookItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotelSceneBookItemResponse
func (client *Client) DeleteHotelSceneBookItemWithOptions(request *DeleteHotelSceneBookItemRequest, headers *DeleteHotelSceneBookItemHeaders, runtime *dara.RuntimeOptions) (_result *DeleteHotelSceneBookItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHotelSceneBookItem"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteHotelSceneBookItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHotelSceneBookItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订删除
//
// @param request - DeleteHotelSceneBookItemRequest
//
// @return DeleteHotelSceneBookItemResponse
func (client *Client) DeleteHotelSceneBookItem(request *DeleteHotelSceneBookItemRequest) (_result *DeleteHotelSceneBookItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteHotelSceneBookItemHeaders{}
	_result = &DeleteHotelSceneBookItemResponse{}
	_body, _err := client.DeleteHotelSceneBookItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除定制配置
//
// @param request - DeleteHotelSettingRequest
//
// @param headers - DeleteHotelSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotelSettingResponse
func (client *Client) DeleteHotelSettingWithOptions(request *DeleteHotelSettingRequest, headers *DeleteHotelSettingHeaders, runtime *dara.RuntimeOptions) (_result *DeleteHotelSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.SettingType) {
		body["SettingType"] = request.SettingType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHotelSetting"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteHotelSetting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHotelSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除定制配置
//
// @param request - DeleteHotelSettingRequest
//
// @return DeleteHotelSettingResponse
func (client *Client) DeleteHotelSetting(request *DeleteHotelSettingRequest) (_result *DeleteHotelSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteHotelSettingHeaders{}
	_result = &DeleteHotelSettingResponse{}
	_body, _err := client.DeleteHotelSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消息通知模板
//
// @param request - DeleteMessageTemplateRequest
//
// @param headers - DeleteMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageTemplateResponse
func (client *Client) DeleteMessageTemplateWithOptions(request *DeleteMessageTemplateRequest, headers *DeleteMessageTemplateHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMessageTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessageTemplate"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteMessageTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消息通知模板
//
// @param request - DeleteMessageTemplateRequest
//
// @return DeleteMessageTemplateResponse
func (client *Client) DeleteMessageTemplate(request *DeleteMessageTemplateRequest) (_result *DeleteMessageTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteMessageTemplateHeaders{}
	_result = &DeleteMessageTemplateResponse{}
	_body, _err := client.DeleteMessageTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除酒店自定义rcu场景
//
// @param request - DeleteRcuSceneRequest
//
// @param headers - DeleteRcuSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRcuSceneResponse
func (client *Client) DeleteRcuSceneWithOptions(request *DeleteRcuSceneRequest, headers *DeleteRcuSceneHeaders, runtime *dara.RuntimeOptions) (_result *DeleteRcuSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRcuScene"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deleteRcuScene"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRcuSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除酒店自定义rcu场景
//
// @param request - DeleteRcuSceneRequest
//
// @return DeleteRcuSceneResponse
func (client *Client) DeleteRcuScene(request *DeleteRcuSceneRequest) (_result *DeleteRcuSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteRcuSceneHeaders{}
	_result = &DeleteRcuSceneResponse{}
	_body, _err := client.DeleteRcuSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备控制
//
// @param tmpReq - DeviceControlRequest
//
// @param headers - DeviceControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceControlResponse
func (client *Client) DeviceControlWithOptions(tmpReq *DeviceControlRequest, headers *DeviceControlHeaders, runtime *dara.RuntimeOptions) (_result *DeviceControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeviceControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeviceControl"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/deviceControl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeviceControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备控制
//
// @param request - DeviceControlRequest
//
// @return DeviceControlResponse
func (client *Client) DeviceControl(request *DeviceControlRequest) (_result *DeviceControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeviceControlHeaders{}
	_result = &DeviceControlResponse{}
	_body, _err := client.DeviceControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 控制房间场景
//
// @param request - ExecuteSceneRequest
//
// @param headers - ExecuteSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSceneResponse
func (client *Client) ExecuteSceneWithOptions(request *ExecuteSceneRequest, headers *ExecuteSceneHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.SceneName) {
		body["SceneName"] = request.SceneName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteScene"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/executeScene"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制房间场景
//
// @param request - ExecuteSceneRequest
//
// @return ExecuteSceneResponse
func (client *Client) ExecuteScene(request *ExecuteSceneRequest) (_result *ExecuteSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ExecuteSceneHeaders{}
	_result = &ExecuteSceneResponse{}
	_body, _err := client.ExecuteSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取基础信息问答
//
// @param request - GetBasicInfoQARequest
//
// @param headers - GetBasicInfoQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicInfoQAResponse
func (client *Client) GetBasicInfoQAWithOptions(request *GetBasicInfoQARequest, headers *GetBasicInfoQAHeaders, runtime *dara.RuntimeOptions) (_result *GetBasicInfoQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicInfoQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getBasicInfoQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicInfoQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取基础信息问答
//
// @param request - GetBasicInfoQARequest
//
// @return GetBasicInfoQAResponse
func (client *Client) GetBasicInfoQA(request *GetBasicInfoQARequest) (_result *GetBasicInfoQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetBasicInfoQAHeaders{}
	_result = &GetBasicInfoQAResponse{}
	_body, _err := client.GetBasicInfoQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询动画
//
// @param request - GetCartoonRequest
//
// @param headers - GetCartoonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCartoonResponse
func (client *Client) GetCartoonWithOptions(request *GetCartoonRequest, headers *GetCartoonHeaders, runtime *dara.RuntimeOptions) (_result *GetCartoonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCartoon"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getCartoon"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCartoonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询动画
//
// @param request - GetCartoonRequest
//
// @return GetCartoonResponse
func (client *Client) GetCartoon(request *GetCartoonRequest) (_result *GetCartoonResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCartoonHeaders{}
	_result = &GetCartoonResponse{}
	_body, _err := client.GetCartoonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前设备的通话信息
//
// @param tmpReq - GetHotelContactByGenieDeviceRequest
//
// @param headers - GetHotelContactByGenieDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelContactByGenieDeviceResponse
func (client *Client) GetHotelContactByGenieDeviceWithOptions(tmpReq *GetHotelContactByGenieDeviceRequest, headers *GetHotelContactByGenieDeviceHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelContactByGenieDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelContactByGenieDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelContactByGenieDevice"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelContactByGenieDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelContactByGenieDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前设备的通话信息
//
// @param request - GetHotelContactByGenieDeviceRequest
//
// @return GetHotelContactByGenieDeviceResponse
func (client *Client) GetHotelContactByGenieDevice(request *GetHotelContactByGenieDeviceRequest) (_result *GetHotelContactByGenieDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelContactByGenieDeviceHeaders{}
	_result = &GetHotelContactByGenieDeviceResponse{}
	_body, _err := client.GetHotelContactByGenieDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据号码获取酒店联系人
//
// @param tmpReq - GetHotelContactByNumberRequest
//
// @param headers - GetHotelContactByNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelContactByNumberResponse
func (client *Client) GetHotelContactByNumberWithOptions(tmpReq *GetHotelContactByNumberRequest, headers *GetHotelContactByNumberHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelContactByNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelContactByNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelContactByNumber"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelContactByNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelContactByNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据号码获取酒店联系人
//
// @param request - GetHotelContactByNumberRequest
//
// @return GetHotelContactByNumberResponse
func (client *Client) GetHotelContactByNumber(request *GetHotelContactByNumberRequest) (_result *GetHotelContactByNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelContactByNumberHeaders{}
	_result = &GetHotelContactByNumberResponse{}
	_body, _err := client.GetHotelContactByNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店联系人
//
// @param tmpReq - GetHotelContactsRequest
//
// @param headers - GetHotelContactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelContactsResponse
func (client *Client) GetHotelContactsWithOptions(tmpReq *GetHotelContactsRequest, headers *GetHotelContactsHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelContactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelContacts"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelContacts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店联系人
//
// @param request - GetHotelContactsRequest
//
// @return GetHotelContactsResponse
func (client *Client) GetHotelContacts(request *GetHotelContactsRequest) (_result *GetHotelContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelContactsHeaders{}
	_result = &GetHotelContactsResponse{}
	_body, _err := client.GetHotelContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取首页背景图和场景模式
//
// @param tmpReq - GetHotelHomeBackImageAndModesRequest
//
// @param headers - GetHotelHomeBackImageAndModesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelHomeBackImageAndModesResponse
func (client *Client) GetHotelHomeBackImageAndModesWithOptions(tmpReq *GetHotelHomeBackImageAndModesRequest, headers *GetHotelHomeBackImageAndModesHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelHomeBackImageAndModesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelHomeBackImageAndModesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelHomeBackImageAndModes"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelHomeBackImageAndModes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelHomeBackImageAndModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取首页背景图和场景模式
//
// @param request - GetHotelHomeBackImageAndModesRequest
//
// @return GetHotelHomeBackImageAndModesResponse
func (client *Client) GetHotelHomeBackImageAndModes(request *GetHotelHomeBackImageAndModesRequest) (_result *GetHotelHomeBackImageAndModesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelHomeBackImageAndModesHeaders{}
	_result = &GetHotelHomeBackImageAndModesResponse{}
	_body, _err := client.GetHotelHomeBackImageAndModesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param tmpReq - GetHotelNoticeRequest
//
// @param headers - GetHotelNoticeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelNoticeResponse
func (client *Client) GetHotelNoticeWithOptions(tmpReq *GetHotelNoticeRequest, headers *GetHotelNoticeHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelNoticeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelNoticeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelNotice"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelNotice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param request - GetHotelNoticeRequest
//
// @return GetHotelNoticeResponse
func (client *Client) GetHotelNotice(request *GetHotelNoticeRequest) (_result *GetHotelNoticeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelNoticeHeaders{}
	_result = &GetHotelNoticeResponse{}
	_body, _err := client.GetHotelNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param tmpReq - GetHotelNoticeV2Request
//
// @param headers - GetHotelNoticeV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelNoticeV2Response
func (client *Client) GetHotelNoticeV2WithOptions(tmpReq *GetHotelNoticeV2Request, headers *GetHotelNoticeV2Headers, runtime *dara.RuntimeOptions) (_result *GetHotelNoticeV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelNoticeV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelNoticeV2"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelNoticeV2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelNoticeV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店通知
//
// @param request - GetHotelNoticeV2Request
//
// @return GetHotelNoticeV2Response
func (client *Client) GetHotelNoticeV2(request *GetHotelNoticeV2Request) (_result *GetHotelNoticeV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelNoticeV2Headers{}
	_result = &GetHotelNoticeV2Response{}
	_body, _err := client.GetHotelNoticeV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店订单详情
//
// @param tmpReq - GetHotelOrderDetailRequest
//
// @param headers - GetHotelOrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelOrderDetailResponse
func (client *Client) GetHotelOrderDetailWithOptions(tmpReq *GetHotelOrderDetailRequest, headers *GetHotelOrderDetailHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelOrderDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelOrderDetail"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelOrderDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店订单详情
//
// @param request - GetHotelOrderDetailRequest
//
// @return GetHotelOrderDetailResponse
func (client *Client) GetHotelOrderDetail(request *GetHotelOrderDetailRequest) (_result *GetHotelOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelOrderDetailHeaders{}
	_result = &GetHotelOrderDetailResponse{}
	_body, _err := client.GetHotelOrderDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店房间猫精设备信息
//
// @param request - GetHotelRoomDeviceRequest
//
// @param headers - GetHotelRoomDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelRoomDeviceResponse
func (client *Client) GetHotelRoomDeviceWithOptions(request *GetHotelRoomDeviceRequest, headers *GetHotelRoomDeviceHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelRoomDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		query["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		query["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelRoomDevice"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelRoomDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelRoomDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店房间猫精设备信息
//
// @param request - GetHotelRoomDeviceRequest
//
// @return GetHotelRoomDeviceResponse
func (client *Client) GetHotelRoomDevice(request *GetHotelRoomDeviceRequest) (_result *GetHotelRoomDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelRoomDeviceHeaders{}
	_result = &GetHotelRoomDeviceResponse{}
	_body, _err := client.GetHotelRoomDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取推荐语料
//
// @param tmpReq - GetHotelSampleUtterancesRequest
//
// @param headers - GetHotelSampleUtterancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelSampleUtterancesResponse
func (client *Client) GetHotelSampleUtterancesWithOptions(tmpReq *GetHotelSampleUtterancesRequest, headers *GetHotelSampleUtterancesHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelSampleUtterancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelSampleUtterancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelSampleUtterances"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelSampleUtterances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelSampleUtterancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推荐语料
//
// @param request - GetHotelSampleUtterancesRequest
//
// @return GetHotelSampleUtterancesResponse
func (client *Client) GetHotelSampleUtterances(request *GetHotelSampleUtterancesRequest) (_result *GetHotelSampleUtterancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelSampleUtterancesHeaders{}
	_result = &GetHotelSampleUtterancesResponse{}
	_body, _err := client.GetHotelSampleUtterancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景详情
//
// @param request - GetHotelSceneItemDetailRequest
//
// @param headers - GetHotelSceneItemDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelSceneItemDetailResponse
func (client *Client) GetHotelSceneItemDetailWithOptions(request *GetHotelSceneItemDetailRequest, headers *GetHotelSceneItemDetailHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelSceneItemDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.ItemId) {
		body["ItemId"] = request.ItemId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelSceneItemDetail"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelSceneItemDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelSceneItemDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景详情
//
// @param request - GetHotelSceneItemDetailRequest
//
// @return GetHotelSceneItemDetailResponse
func (client *Client) GetHotelSceneItemDetail(request *GetHotelSceneItemDetailRequest) (_result *GetHotelSceneItemDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelSceneItemDetailHeaders{}
	_result = &GetHotelSceneItemDetailResponse{}
	_body, _err := client.GetHotelSceneItemDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店屏保
//
// @param tmpReq - GetHotelScreenSaverRequest
//
// @param headers - GetHotelScreenSaverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelScreenSaverResponse
func (client *Client) GetHotelScreenSaverWithOptions(tmpReq *GetHotelScreenSaverRequest, headers *GetHotelScreenSaverHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelScreenSaverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotelScreenSaverShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelScreenSaver"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelScreenSaver"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelScreenSaverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店屏保
//
// @param request - GetHotelScreenSaverRequest
//
// @return GetHotelScreenSaverResponse
func (client *Client) GetHotelScreenSaver(request *GetHotelScreenSaverRequest) (_result *GetHotelScreenSaverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelScreenSaverHeaders{}
	_result = &GetHotelScreenSaverResponse{}
	_body, _err := client.GetHotelScreenSaverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取屏保列表
//
// @param request - GetHotelScreenSaverStyleRequest
//
// @param headers - GetHotelScreenSaverStyleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelScreenSaverStyleResponse
func (client *Client) GetHotelScreenSaverStyleWithOptions(request *GetHotelScreenSaverStyleRequest, headers *GetHotelScreenSaverStyleHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelScreenSaverStyleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelScreenSaverStyle"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelScreenSaverStyle"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelScreenSaverStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取屏保列表
//
// @param request - GetHotelScreenSaverStyleRequest
//
// @return GetHotelScreenSaverStyleResponse
func (client *Client) GetHotelScreenSaverStyle(request *GetHotelScreenSaverStyleRequest) (_result *GetHotelScreenSaverStyleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelScreenSaverStyleHeaders{}
	_result = &GetHotelScreenSaverStyleResponse{}
	_body, _err := client.GetHotelScreenSaverStyleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询定制配置
//
// @param request - GetHotelSettingRequest
//
// @param headers - GetHotelSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotelSettingResponse
func (client *Client) GetHotelSettingWithOptions(request *GetHotelSettingRequest, headers *GetHotelSettingHeaders, runtime *dara.RuntimeOptions) (_result *GetHotelSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.SettingType) {
		body["SettingType"] = request.SettingType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotelSetting"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelSetting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotelSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询定制配置
//
// @param request - GetHotelSettingRequest
//
// @return GetHotelSettingResponse
func (client *Client) GetHotelSetting(request *GetHotelSettingRequest) (_result *GetHotelSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetHotelSettingHeaders{}
	_result = &GetHotelSettingResponse{}
	_body, _err := client.GetHotelSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关联产品列表查看
//
// @param headers - GetRelationProductListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRelationProductListResponse
func (client *Client) GetRelationProductListWithOptions(headers *GetRelationProductListHeaders, runtime *dara.RuntimeOptions) (_result *GetRelationProductListResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRelationProductList"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getRelationProductList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRelationProductListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关联产品列表查看
//
// @return GetRelationProductListResponse
func (client *Client) GetRelationProductList() (_result *GetRelationProductListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetRelationProductListHeaders{}
	_result = &GetRelationProductListResponse{}
	_body, _err := client.GetRelationProductListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取组织下unionId列表
//
// @param request - GetUnionIdRequest
//
// @param headers - GetUnionIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUnionIdResponse
func (client *Client) GetUnionIdWithOptions(request *GetUnionIdRequest, headers *GetUnionIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUnionIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EncodeKey) {
		body["EncodeKey"] = request.EncodeKey
	}

	if !dara.IsNil(request.EncodeType) {
		body["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.IdType) {
		body["IdType"] = request.IdType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUnionId"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getUnionId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUnionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组织下unionId列表
//
// @param request - GetUnionIdRequest
//
// @return GetUnionIdResponse
func (client *Client) GetUnionId(request *GetUnionIdRequest) (_result *GetUnionIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUnionIdHeaders{}
	_result = &GetUnionIdResponse{}
	_body, _err := client.GetUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询欢迎语信息
//
// @param request - GetWelcomeTextAndMusicRequest
//
// @param headers - GetWelcomeTextAndMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWelcomeTextAndMusicResponse
func (client *Client) GetWelcomeTextAndMusicWithOptions(request *GetWelcomeTextAndMusicRequest, headers *GetWelcomeTextAndMusicHeaders, runtime *dara.RuntimeOptions) (_result *GetWelcomeTextAndMusicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWelcomeTextAndMusic"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getWelcomeTextAndMusic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWelcomeTextAndMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询欢迎语信息
//
// @param request - GetWelcomeTextAndMusicRequest
//
// @return GetWelcomeTextAndMusicResponse
func (client *Client) GetWelcomeTextAndMusic(request *GetWelcomeTextAndMusicRequest) (_result *GetWelcomeTextAndMusicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetWelcomeTextAndMusicHeaders{}
	_result = &GetWelcomeTextAndMusicResponse{}
	_body, _err := client.GetWelcomeTextAndMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店带屏设备扫码绑定
//
// @param request - HotelQrBindRequest
//
// @param headers - HotelQrBindHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotelQrBindResponse
func (client *Client) HotelQrBindWithOptions(request *HotelQrBindRequest, headers *HotelQrBindHeaders, runtime *dara.RuntimeOptions) (_result *HotelQrBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.ExtInfo) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotelQrBind"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/hotelQrBind"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotelQrBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店带屏设备扫码绑定
//
// @param request - HotelQrBindRequest
//
// @return HotelQrBindResponse
func (client *Client) HotelQrBind(request *HotelQrBindRequest) (_result *HotelQrBindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &HotelQrBindHeaders{}
	_result = &HotelQrBindResponse{}
	_body, _err := client.HotelQrBindWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量导入酒店配置
//
// @param tmpReq - ImportHotelConfigRequest
//
// @param headers - ImportHotelConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportHotelConfigResponse
func (client *Client) ImportHotelConfigWithOptions(tmpReq *ImportHotelConfigRequest, headers *ImportHotelConfigHeaders, runtime *dara.RuntimeOptions) (_result *ImportHotelConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportHotelConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImportHotelConfig) {
		request.ImportHotelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImportHotelConfig, dara.String("ImportHotelConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.ImportHotelConfigShrink) {
		body["ImportHotelConfig"] = request.ImportHotelConfigShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportHotelConfig"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/importHotelConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportHotelConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入酒店配置
//
// @param request - ImportHotelConfigRequest
//
// @return ImportHotelConfigResponse
func (client *Client) ImportHotelConfig(request *ImportHotelConfigRequest) (_result *ImportHotelConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ImportHotelConfigHeaders{}
	_result = &ImportHotelConfigResponse{}
	_body, _err := client.ImportHotelConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量导入设备（同时补充房型）
//
// @param tmpReq - ImportRoomControlDevicesRequest
//
// @param headers - ImportRoomControlDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportRoomControlDevicesResponse
func (client *Client) ImportRoomControlDevicesWithOptions(tmpReq *ImportRoomControlDevicesRequest, headers *ImportRoomControlDevicesHeaders, runtime *dara.RuntimeOptions) (_result *ImportRoomControlDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportRoomControlDevicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LocationDevices) {
		request.LocationDevicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDevices, dara.String("LocationDevices"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableInfraredDeviceImport) {
		body["EnableInfraredDeviceImport"] = request.EnableInfraredDeviceImport
	}

	if !dara.IsNil(request.EnableMeshDeviceImport) {
		body["EnableMeshDeviceImport"] = request.EnableMeshDeviceImport
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.LocationDevicesShrink) {
		body["LocationDevices"] = request.LocationDevicesShrink
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportRoomControlDevices"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/importRoomControlDevices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportRoomControlDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入设备（同时补充房型）
//
// @param request - ImportRoomControlDevicesRequest
//
// @return ImportRoomControlDevicesResponse
func (client *Client) ImportRoomControlDevices(request *ImportRoomControlDevicesRequest) (_result *ImportRoomControlDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ImportRoomControlDevicesHeaders{}
	_result = &ImportRoomControlDevicesResponse{}
	_body, _err := client.ImportRoomControlDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入房间内精灵场景
//
// @param tmpReq - ImportRoomGenieScenesRequest
//
// @param headers - ImportRoomGenieScenesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportRoomGenieScenesResponse
func (client *Client) ImportRoomGenieScenesWithOptions(tmpReq *ImportRoomGenieScenesRequest, headers *ImportRoomGenieScenesHeaders, runtime *dara.RuntimeOptions) (_result *ImportRoomGenieScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportRoomGenieScenesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SceneList) {
		request.SceneListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneList, dara.String("SceneList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.SceneListShrink) {
		body["SceneList"] = request.SceneListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportRoomGenieScenes"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/importRoomGenieScenes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportRoomGenieScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入房间内精灵场景
//
// @param request - ImportRoomGenieScenesRequest
//
// @return ImportRoomGenieScenesResponse
func (client *Client) ImportRoomGenieScenes(request *ImportRoomGenieScenesRequest) (_result *ImportRoomGenieScenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ImportRoomGenieScenesHeaders{}
	_result = &ImportRoomGenieScenesResponse{}
	_body, _err := client.ImportRoomGenieScenesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订新增
//
// @param tmpReq - InsertHotelSceneBookItemRequest
//
// @param headers - InsertHotelSceneBookItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertHotelSceneBookItemResponse
func (client *Client) InsertHotelSceneBookItemWithOptions(tmpReq *InsertHotelSceneBookItemRequest, headers *InsertHotelSceneBookItemHeaders, runtime *dara.RuntimeOptions) (_result *InsertHotelSceneBookItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertHotelSceneBookItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddHotelSceneItemReq) {
		request.AddHotelSceneItemReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddHotelSceneItemReq, dara.String("AddHotelSceneItemReq"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddHotelSceneItemReqShrink) {
		query["AddHotelSceneItemReq"] = request.AddHotelSceneItemReqShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertHotelSceneBookItem"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/insertHotelSceneBookItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertHotelSceneBookItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订新增
//
// @param request - InsertHotelSceneBookItemRequest
//
// @return InsertHotelSceneBookItemResponse
func (client *Client) InsertHotelSceneBookItem(request *InsertHotelSceneBookItemRequest) (_result *InsertHotelSceneBookItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InsertHotelSceneBookItemHeaders{}
	_result = &InsertHotelSceneBookItemResponse{}
	_body, _err := client.InsertHotelSceneBookItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人服务，消息推送
//
// @param request - InvokeRobotPushRequest
//
// @param headers - InvokeRobotPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeRobotPushResponse
func (client *Client) InvokeRobotPushWithOptions(request *InvokeRobotPushRequest, headers *InvokeRobotPushHeaders, runtime *dara.RuntimeOptions) (_result *InvokeRobotPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.PushType) {
		body["PushType"] = request.PushType
	}

	if !dara.IsNil(request.RoomName) {
		body["RoomName"] = request.RoomName
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeRobotPush"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/invokeRobotPush"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeRobotPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人服务，消息推送
//
// @param request - InvokeRobotPushRequest
//
// @return InvokeRobotPushResponse
func (client *Client) InvokeRobotPush(request *InvokeRobotPushRequest) (_result *InvokeRobotPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InvokeRobotPushHeaders{}
	_result = &InvokeRobotPushResponse{}
	_body, _err := client.InvokeRobotPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询省份
//
// @param headers - ListAllProvincesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllProvincesResponse
func (client *Client) ListAllProvincesWithOptions(headers *ListAllProvincesHeaders, runtime *dara.RuntimeOptions) (_result *ListAllProvincesResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllProvinces"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listAllProvinces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllProvincesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询省份
//
// @return ListAllProvincesResponse
func (client *Client) ListAllProvinces() (_result *ListAllProvincesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListAllProvincesHeaders{}
	_result = &ListAllProvincesResponse{}
	_body, _err := client.ListAllProvincesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询城市
//
// @param request - ListCitiesByProvinceRequest
//
// @param headers - ListCitiesByProvinceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCitiesByProvinceResponse
func (client *Client) ListCitiesByProvinceWithOptions(request *ListCitiesByProvinceRequest, headers *ListCitiesByProvinceHeaders, runtime *dara.RuntimeOptions) (_result *ListCitiesByProvinceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Province) {
		body["Province"] = request.Province
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCitiesByProvince"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listCitiesByProvince"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCitiesByProvinceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询城市
//
// @param request - ListCitiesByProvinceRequest
//
// @return ListCitiesByProvinceResponse
func (client *Client) ListCitiesByProvince(request *ListCitiesByProvinceRequest) (_result *ListCitiesByProvinceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCitiesByProvinceHeaders{}
	_result = &ListCitiesByProvinceResponse{}
	_body, _err := client.ListCitiesByProvinceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义问答列表
//
// @param tmpReq - ListCustomQARequest
//
// @param headers - ListCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomQAResponse
func (client *Client) ListCustomQAWithOptions(tmpReq *ListCustomQARequest, headers *ListCustomQAHeaders, runtime *dara.RuntimeOptions) (_result *ListCustomQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageShrink) {
		body["Page"] = request.PageShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listCustomQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义问答列表
//
// @param request - ListCustomQARequest
//
// @return ListCustomQAResponse
func (client *Client) ListCustomQA(request *ListCustomQARequest) (_result *ListCustomQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCustomQAHeaders{}
	_result = &ListCustomQAResponse{}
	_body, _err := client.ListCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景对话模板
//
// @param request - ListDialogueTemplateRequest
//
// @param headers - ListDialogueTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDialogueTemplateResponse
func (client *Client) ListDialogueTemplateWithOptions(request *ListDialogueTemplateRequest, headers *ListDialogueTemplateHeaders, runtime *dara.RuntimeOptions) (_result *ListDialogueTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDialogueTemplate"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listDialogueTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDialogueTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景对话模板
//
// @param request - ListDialogueTemplateRequest
//
// @return ListDialogueTemplateResponse
func (client *Client) ListDialogueTemplate(request *ListDialogueTemplateRequest) (_result *ListDialogueTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDialogueTemplateHeaders{}
	_result = &ListDialogueTemplateResponse{}
	_body, _err := client.ListDialogueTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店闹钟
//
// @param tmpReq - ListHotelAlarmRequest
//
// @param headers - ListHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelAlarmResponse
func (client *Client) ListHotelAlarmWithOptions(tmpReq *ListHotelAlarmRequest, headers *ListHotelAlarmHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rooms) {
		request.RoomsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rooms, dara.String("Rooms"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomsShrink) {
		body["Rooms"] = request.RoomsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelAlarm"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/getHotelAlarmList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店闹钟
//
// @param request - ListHotelAlarmRequest
//
// @return ListHotelAlarmResponse
func (client *Client) ListHotelAlarm(request *ListHotelAlarmRequest) (_result *ListHotelAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelAlarmHeaders{}
	_result = &ListHotelAlarmResponse{}
	_body, _err := client.ListHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店设备列表
//
// @param tmpReq - ListHotelControlDeviceRequest
//
// @param headers - ListHotelControlDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelControlDeviceResponse
func (client *Client) ListHotelControlDeviceWithOptions(tmpReq *ListHotelControlDeviceRequest, headers *ListHotelControlDeviceHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelControlDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelControlDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelControlDevice"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelControlDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelControlDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店设备列表
//
// @param request - ListHotelControlDeviceRequest
//
// @return ListHotelControlDeviceResponse
func (client *Client) ListHotelControlDevice(request *ListHotelControlDeviceRequest) (_result *ListHotelControlDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelControlDeviceHeaders{}
	_result = &ListHotelControlDeviceResponse{}
	_body, _err := client.ListHotelControlDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店列表
//
// @param headers - ListHotelInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelInfoResponse
func (client *Client) ListHotelInfoWithOptions(headers *ListHotelInfoHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelInfo"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店列表
//
// @return ListHotelInfoResponse
func (client *Client) ListHotelInfo() (_result *ListHotelInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelInfoHeaders{}
	_result = &ListHotelInfoResponse{}
	_body, _err := client.ListHotelInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取消息模板
//
// @param headers - ListHotelMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelMessageTemplateResponse
func (client *Client) ListHotelMessageTemplateWithOptions(headers *ListHotelMessageTemplateHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelMessageTemplateResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelMessageTemplate"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelMessageTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取消息模板
//
// @return ListHotelMessageTemplateResponse
func (client *Client) ListHotelMessageTemplate() (_result *ListHotelMessageTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelMessageTemplateHeaders{}
	_result = &ListHotelMessageTemplateResponse{}
	_body, _err := client.ListHotelMessageTemplateWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店订单列表
//
// @param tmpReq - ListHotelOrderRequest
//
// @param headers - ListHotelOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelOrderResponse
func (client *Client) ListHotelOrderWithOptions(tmpReq *ListHotelOrderRequest, headers *ListHotelOrderHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelOrder"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店订单列表
//
// @param request - ListHotelOrderRequest
//
// @return ListHotelOrderResponse
func (client *Client) ListHotelOrder(request *ListHotelOrderRequest) (_result *ListHotelOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelOrderHeaders{}
	_result = &ListHotelOrderResponse{}
	_body, _err := client.ListHotelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取酒店的所有房间
//
// @param tmpReq - ListHotelRoomsRequest
//
// @param headers - ListHotelRoomsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelRoomsResponse
func (client *Client) ListHotelRoomsWithOptions(tmpReq *ListHotelRoomsRequest, headers *ListHotelRoomsHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelRoomsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelRoomsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelAdminRoom) {
		request.HotelAdminRoomShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelAdminRoom, dara.String("HotelAdminRoom"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelAdminRoomShrink) {
		body["HotelAdminRoom"] = request.HotelAdminRoomShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelRooms"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelRooms"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取酒店的所有房间
//
// @param request - ListHotelRoomsRequest
//
// @return ListHotelRoomsResponse
func (client *Client) ListHotelRooms(request *ListHotelRoomsRequest) (_result *ListHotelRoomsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelRoomsHeaders{}
	_result = &ListHotelRoomsResponse{}
	_body, _err := client.ListHotelRoomsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订列表（餐饮/SPA休闲/打车）
//
// @param tmpReq - ListHotelSceneBookItemsRequest
//
// @param headers - ListHotelSceneBookItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelSceneBookItemsResponse
func (client *Client) ListHotelSceneBookItemsWithOptions(tmpReq *ListHotelSceneBookItemsRequest, headers *ListHotelSceneBookItemsHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelSceneBookItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelSceneBookItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelSceneBookItems"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelSceneBookItems"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelSceneBookItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订列表（餐饮/SPA休闲/打车）
//
// @param request - ListHotelSceneBookItemsRequest
//
// @return ListHotelSceneBookItemsResponse
func (client *Client) ListHotelSceneBookItems(request *ListHotelSceneBookItemsRequest) (_result *ListHotelSceneBookItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelSceneBookItemsHeaders{}
	_result = &ListHotelSceneBookItemsResponse{}
	_body, _err := client.ListHotelSceneBookItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务项目
//
// @param tmpReq - ListHotelSceneItemRequest
//
// @param headers - ListHotelSceneItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelSceneItemResponse
func (client *Client) ListHotelSceneItemWithOptions(tmpReq *ListHotelSceneItemRequest, headers *ListHotelSceneItemHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelSceneItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelSceneItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelSceneItem"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelSceneItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelSceneItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务项目
//
// @param request - ListHotelSceneItemRequest
//
// @return ListHotelSceneItemResponse
func (client *Client) ListHotelSceneItem(request *ListHotelSceneItemRequest) (_result *ListHotelSceneItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelSceneItemHeaders{}
	_result = &ListHotelSceneItemResponse{}
	_body, _err := client.ListHotelSceneItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景列表（物品/服务/维修）
//
// @param tmpReq - ListHotelSceneItemsRequest
//
// @param headers - ListHotelSceneItemsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelSceneItemsResponse
func (client *Client) ListHotelSceneItemsWithOptions(tmpReq *ListHotelSceneItemsRequest, headers *ListHotelSceneItemsHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelSceneItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelSceneItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListHotelSceneReq) {
		request.ListHotelSceneReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListHotelSceneReq, dara.String("ListHotelSceneReq"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ListHotelSceneReqShrink) {
		query["ListHotelSceneReq"] = request.ListHotelSceneReqShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelSceneItems"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelSceneItems"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelSceneItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景列表（物品/服务/维修）
//
// @param request - ListHotelSceneItemsRequest
//
// @return ListHotelSceneItemsResponse
func (client *Client) ListHotelSceneItems(request *ListHotelSceneItemsRequest) (_result *ListHotelSceneItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelSceneItemsHeaders{}
	_result = &ListHotelSceneItemsResponse{}
	_body, _err := client.ListHotelSceneItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务分类列表
//
// @param tmpReq - ListHotelServiceCategoryRequest
//
// @param headers - ListHotelServiceCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelServiceCategoryResponse
func (client *Client) ListHotelServiceCategoryWithOptions(tmpReq *ListHotelServiceCategoryRequest, headers *ListHotelServiceCategoryHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelServiceCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelServiceCategoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotelServiceCategory"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotelServiceCategory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelServiceCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务分类列表
//
// @param request - ListHotelServiceCategoryRequest
//
// @return ListHotelServiceCategoryResponse
func (client *Client) ListHotelServiceCategory(request *ListHotelServiceCategoryRequest) (_result *ListHotelServiceCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelServiceCategoryHeaders{}
	_result = &ListHotelServiceCategoryResponse{}
	_body, _err := client.ListHotelServiceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店列表(待审批/已拒绝/已通过)
//
// @param tmpReq - ListHotelsRequest
//
// @param headers - ListHotelsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotelsResponse
func (client *Client) ListHotelsWithOptions(tmpReq *ListHotelsRequest, headers *ListHotelsHeaders, runtime *dara.RuntimeOptions) (_result *ListHotelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListHotelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HotelRequest) {
		request.HotelRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotelRequest, dara.String("HotelRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HotelRequestShrink) {
		query["HotelRequest"] = request.HotelRequestShrink
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotels"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listHotels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店列表(待审批/已拒绝/已通过)
//
// @param request - ListHotelsRequest
//
// @return ListHotelsResponse
func (client *Client) ListHotels(request *ListHotelsRequest) (_result *ListHotelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListHotelsHeaders{}
	_result = &ListHotelsResponse{}
	_body, _err := client.ListHotelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询红外品牌列表
//
// @param request - ListInfraredDeviceBrandsRequest
//
// @param headers - ListInfraredDeviceBrandsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInfraredDeviceBrandsResponse
func (client *Client) ListInfraredDeviceBrandsWithOptions(request *ListInfraredDeviceBrandsRequest, headers *ListInfraredDeviceBrandsHeaders, runtime *dara.RuntimeOptions) (_result *ListInfraredDeviceBrandsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.ServiceProvider) {
		body["ServiceProvider"] = request.ServiceProvider
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInfraredDeviceBrands"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listInfraredDeviceBrands"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInfraredDeviceBrandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询红外品牌列表
//
// @param request - ListInfraredDeviceBrandsRequest
//
// @return ListInfraredDeviceBrandsResponse
func (client *Client) ListInfraredDeviceBrands(request *ListInfraredDeviceBrandsRequest) (_result *ListInfraredDeviceBrandsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListInfraredDeviceBrandsHeaders{}
	_result = &ListInfraredDeviceBrandsResponse{}
	_body, _err := client.ListInfraredDeviceBrandsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询红外码库列表
//
// @param request - ListInfraredRemoteControllersRequest
//
// @param headers - ListInfraredRemoteControllersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInfraredRemoteControllersResponse
func (client *Client) ListInfraredRemoteControllersWithOptions(request *ListInfraredRemoteControllersRequest, headers *ListInfraredRemoteControllersHeaders, runtime *dara.RuntimeOptions) (_result *ListInfraredRemoteControllersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Brand) {
		body["Brand"] = request.Brand
	}

	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.City) {
		body["City"] = request.City
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Province) {
		body["Province"] = request.Province
	}

	if !dara.IsNil(request.ServiceProvider) {
		body["ServiceProvider"] = request.ServiceProvider
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInfraredRemoteControllers"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listInfraredRemoteControllers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInfraredRemoteControllersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询红外码库列表
//
// @param request - ListInfraredRemoteControllersRequest
//
// @return ListInfraredRemoteControllersResponse
func (client *Client) ListInfraredRemoteControllers(request *ListInfraredRemoteControllersRequest) (_result *ListInfraredRemoteControllersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListInfraredRemoteControllersHeaders{}
	_result = &ListInfraredRemoteControllersResponse{}
	_body, _err := client.ListInfraredRemoteControllersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务提供商
//
// @param request - ListSTBServiceProvidersRequest
//
// @param headers - ListSTBServiceProvidersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSTBServiceProvidersResponse
func (client *Client) ListSTBServiceProvidersWithOptions(request *ListSTBServiceProvidersRequest, headers *ListSTBServiceProvidersHeaders, runtime *dara.RuntimeOptions) (_result *ListSTBServiceProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		body["City"] = request.City
	}

	if !dara.IsNil(request.Province) {
		body["Province"] = request.Province
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSTBServiceProviders"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listSTBServiceProviders"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSTBServiceProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务提供商
//
// @param request - ListSTBServiceProvidersRequest
//
// @return ListSTBServiceProvidersResponse
func (client *Client) ListSTBServiceProviders(request *ListSTBServiceProvidersRequest) (_result *ListSTBServiceProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSTBServiceProvidersHeaders{}
	_result = &ListSTBServiceProvidersResponse{}
	_body, _err := client.ListSTBServiceProvidersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景分类
//
// @param request - ListSceneCategoryRequest
//
// @param headers - ListSceneCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSceneCategoryResponse
func (client *Client) ListSceneCategoryWithOptions(request *ListSceneCategoryRequest, headers *ListSceneCategoryHeaders, runtime *dara.RuntimeOptions) (_result *ListSceneCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSceneCategory"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listSceneCategory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSceneCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景分类
//
// @param request - ListSceneCategoryRequest
//
// @return ListSceneCategoryResponse
func (client *Client) ListSceneCategory(request *ListSceneCategoryRequest) (_result *ListSceneCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSceneCategoryHeaders{}
	_result = &ListSceneCategoryResponse{}
	_body, _err := client.ListSceneCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务设施问答列表
//
// @param tmpReq - ListServiceQARequest
//
// @param headers - ListServiceQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceQAResponse
func (client *Client) ListServiceQAWithOptions(tmpReq *ListServiceQARequest, headers *ListServiceQAHeaders, runtime *dara.RuntimeOptions) (_result *ListServiceQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListServiceQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		body["Active"] = request.Active
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageShrink) {
		body["Page"] = request.PageShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listServiceQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务设施问答列表
//
// @param request - ListServiceQARequest
//
// @return ListServiceQAResponse
func (client *Client) ListServiceQA(request *ListServiceQARequest) (_result *ListServiceQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListServiceQAHeaders{}
	_result = &ListServiceQAResponse{}
	_body, _err := client.ListServiceQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工单列表
//
// @param tmpReq - ListTicketsRequest
//
// @param headers - ListTicketsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketsResponse
func (client *Client) ListTicketsWithOptions(tmpReq *ListTicketsRequest, headers *ListTicketsHeaders, runtime *dara.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.IsDesc) {
		body["IsDesc"] = request.IsDesc
	}

	if !dara.IsNil(request.IsNeedCallback) {
		body["IsNeedCallback"] = request.IsNeedCallback
	}

	if !dara.IsNil(request.IsNeedCharges) {
		body["IsNeedCharges"] = request.IsNeedCharges
	}

	if !dara.IsNil(request.PageShrink) {
		body["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTickets"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/listTickets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工单列表
//
// @param request - ListTicketsRequest
//
// @return ListTicketsResponse
func (client *Client) ListTickets(request *ListTicketsRequest) (_result *ListTicketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListTicketsHeaders{}
	_result = &ListTicketsResponse{}
	_body, _err := client.ListTicketsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询酒店房间主控设备
//
// @param request - PageGetHotelRoomDevicesRequest
//
// @param headers - PageGetHotelRoomDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageGetHotelRoomDevicesResponse
func (client *Client) PageGetHotelRoomDevicesWithOptions(request *PageGetHotelRoomDevicesRequest, headers *PageGetHotelRoomDevicesHeaders, runtime *dara.RuntimeOptions) (_result *PageGetHotelRoomDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageGetHotelRoomDevices"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/pageGetHotelRoomDevices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageGetHotelRoomDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询酒店房间主控设备
//
// @param request - PageGetHotelRoomDevicesRequest
//
// @return PageGetHotelRoomDevicesResponse
func (client *Client) PageGetHotelRoomDevices(request *PageGetHotelRoomDevicesRequest) (_result *PageGetHotelRoomDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PageGetHotelRoomDevicesHeaders{}
	_result = &PageGetHotelRoomDevicesResponse{}
	_body, _err := client.PageGetHotelRoomDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// pms事件上报
//
// @param request - PmsEventReportRequest
//
// @param headers - PmsEventReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PmsEventReportResponse
func (client *Client) PmsEventReportWithOptions(request *PmsEventReportRequest, headers *PmsEventReportHeaders, runtime *dara.RuntimeOptions) (_result *PmsEventReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PmsEventReport"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/pmsEventReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PmsEventReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pms事件上报
//
// @param request - PmsEventReportRequest
//
// @return PmsEventReportResponse
func (client *Client) PmsEventReport(request *PmsEventReportRequest) (_result *PmsEventReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PmsEventReportHeaders{}
	_result = &PmsEventReportResponse{}
	_body, _err := client.PmsEventReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送酒店消息
//
// @param tmpReq - PushHotelMessageRequest
//
// @param headers - PushHotelMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushHotelMessageResponse
func (client *Client) PushHotelMessageWithOptions(tmpReq *PushHotelMessageRequest, headers *PushHotelMessageHeaders, runtime *dara.RuntimeOptions) (_result *PushHotelMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushHotelMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PushHotelMessageReq) {
		request.PushHotelMessageReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PushHotelMessageReq, dara.String("PushHotelMessageReq"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PushHotelMessageReqShrink) {
		query["PushHotelMessageReq"] = request.PushHotelMessageReqShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushHotelMessage"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/pushHotelMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushHotelMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送酒店消息
//
// @param request - PushHotelMessageRequest
//
// @return PushHotelMessageResponse
func (client *Client) PushHotelMessage(request *PushHotelMessageRequest) (_result *PushHotelMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PushHotelMessageHeaders{}
	_result = &PushHotelMessageResponse{}
	_body, _err := client.PushHotelMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送音箱指令
//
// @param tmpReq - PushVoiceBoxCommandsRequest
//
// @param headers - PushVoiceBoxCommandsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushVoiceBoxCommandsResponse
func (client *Client) PushVoiceBoxCommandsWithOptions(tmpReq *PushVoiceBoxCommandsRequest, headers *PushVoiceBoxCommandsHeaders, runtime *dara.RuntimeOptions) (_result *PushVoiceBoxCommandsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushVoiceBoxCommandsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Commands) {
		request.CommandsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commands, dara.String("Commands"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommandsShrink) {
		body["Commands"] = request.CommandsShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushVoiceBoxCommands"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/pushVoiceBoxCommands"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushVoiceBoxCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送音箱指令
//
// @param request - PushVoiceBoxCommandsRequest
//
// @return PushVoiceBoxCommandsResponse
func (client *Client) PushVoiceBoxCommands(request *PushVoiceBoxCommandsRequest) (_result *PushVoiceBoxCommandsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PushVoiceBoxCommandsHeaders{}
	_result = &PushVoiceBoxCommandsResponse{}
	_body, _err := client.PushVoiceBoxCommandsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 直接推送欢迎语
//
// @param request - PushWelcomeRequest
//
// @param headers - PushWelcomeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushWelcomeResponse
func (client *Client) PushWelcomeWithOptions(request *PushWelcomeRequest, headers *PushWelcomeHeaders, runtime *dara.RuntimeOptions) (_result *PushWelcomeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomName) {
		body["RoomName"] = request.RoomName
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.WelcomeMusicUrl) {
		body["WelcomeMusicUrl"] = request.WelcomeMusicUrl
	}

	if !dara.IsNil(request.WelcomeText) {
		body["WelcomeText"] = request.WelcomeText
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushWelcome"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/pushWelcome"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushWelcomeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直接推送欢迎语
//
// @param request - PushWelcomeRequest
//
// @return PushWelcomeResponse
func (client *Client) PushWelcome(request *PushWelcomeRequest) (_result *PushWelcomeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PushWelcomeHeaders{}
	_result = &PushWelcomeResponse{}
	_body, _err := client.PushWelcomeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送欢迎语
//
// @param tmpReq - PushWelcomeTextAndMusicRequest
//
// @param headers - PushWelcomeTextAndMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushWelcomeTextAndMusicResponse
func (client *Client) PushWelcomeTextAndMusicWithOptions(tmpReq *PushWelcomeTextAndMusicRequest, headers *PushWelcomeTextAndMusicHeaders, runtime *dara.RuntimeOptions) (_result *PushWelcomeTextAndMusicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushWelcomeTextAndMusicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TemplateVariable) {
		request.TemplateVariableShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateVariable, dara.String("TemplateVariable"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomName) {
		body["RoomName"] = request.RoomName
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.TemplateVariableShrink) {
		body["TemplateVariable"] = request.TemplateVariableShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushWelcomeTextAndMusic"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/pushWelcomeTextAndMusic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushWelcomeTextAndMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送欢迎语
//
// @param request - PushWelcomeTextAndMusicRequest
//
// @return PushWelcomeTextAndMusicResponse
func (client *Client) PushWelcomeTextAndMusic(request *PushWelcomeTextAndMusicRequest) (_result *PushWelcomeTextAndMusicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PushWelcomeTextAndMusicHeaders{}
	_result = &PushWelcomeTextAndMusicResponse{}
	_body, _err := client.PushWelcomeTextAndMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店设备状态/模式状态查询
//
// @param tmpReq - QueryDeviceStatusRequest
//
// @param headers - QueryDeviceStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceStatusResponse
func (client *Client) QueryDeviceStatusWithOptions(tmpReq *QueryDeviceStatusRequest, headers *QueryDeviceStatusHeaders, runtime *dara.RuntimeOptions) (_result *QueryDeviceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryDeviceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceStatus"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/queryDeviceStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店设备状态/模式状态查询
//
// @param request - QueryDeviceStatusRequest
//
// @return QueryDeviceStatusResponse
func (client *Client) QueryDeviceStatus(request *QueryDeviceStatusRequest) (_result *QueryDeviceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryDeviceStatusHeaders{}
	_result = &QueryDeviceStatusResponse{}
	_body, _err := client.QueryDeviceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询房间详细信息
//
// @param request - QueryHotelRoomDetailRequest
//
// @param headers - QueryHotelRoomDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotelRoomDetailResponse
func (client *Client) QueryHotelRoomDetailWithOptions(request *QueryHotelRoomDetailRequest, headers *QueryHotelRoomDetailHeaders, runtime *dara.RuntimeOptions) (_result *QueryHotelRoomDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Mac) {
		body["Mac"] = request.Mac
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.Sn) {
		body["Sn"] = request.Sn
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryHotelRoomDetail"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/queryHotelRoomDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryHotelRoomDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询房间详细信息
//
// @param request - QueryHotelRoomDetailRequest
//
// @return QueryHotelRoomDetailResponse
func (client *Client) QueryHotelRoomDetail(request *QueryHotelRoomDetailRequest) (_result *QueryHotelRoomDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryHotelRoomDetailHeaders{}
	_result = &QueryHotelRoomDetailResponse{}
	_body, _err := client.QueryHotelRoomDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店房间客控设备
//
// @param request - QueryRoomControlDevicesRequest
//
// @param headers - QueryRoomControlDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRoomControlDevicesResponse
func (client *Client) QueryRoomControlDevicesWithOptions(request *QueryRoomControlDevicesRequest, headers *QueryRoomControlDevicesHeaders, runtime *dara.RuntimeOptions) (_result *QueryRoomControlDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		query["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		query["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRoomControlDevices"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/queryRoomControlDevices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRoomControlDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店房间客控设备
//
// @param request - QueryRoomControlDevicesRequest
//
// @return QueryRoomControlDevicesResponse
func (client *Client) QueryRoomControlDevices(request *QueryRoomControlDevicesRequest) (_result *QueryRoomControlDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryRoomControlDevicesHeaders{}
	_result = &QueryRoomControlDevicesResponse{}
	_body, _err := client.QueryRoomControlDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询房间被控设备包含设备状态
//
// @param request - QueryRoomControlDevicesAndStatusRequest
//
// @param headers - QueryRoomControlDevicesAndStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRoomControlDevicesAndStatusResponse
func (client *Client) QueryRoomControlDevicesAndStatusWithOptions(request *QueryRoomControlDevicesAndStatusRequest, headers *QueryRoomControlDevicesAndStatusHeaders, runtime *dara.RuntimeOptions) (_result *QueryRoomControlDevicesAndStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRoomControlDevicesAndStatus"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/queryRoomControlDevicesAndStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRoomControlDevicesAndStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询房间被控设备包含设备状态
//
// @param request - QueryRoomControlDevicesAndStatusRequest
//
// @return QueryRoomControlDevicesAndStatusResponse
func (client *Client) QueryRoomControlDevicesAndStatus(request *QueryRoomControlDevicesAndStatusRequest) (_result *QueryRoomControlDevicesAndStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryRoomControlDevicesAndStatusHeaders{}
	_result = &QueryRoomControlDevicesAndStatusResponse{}
	_body, _err := client.QueryRoomControlDevicesAndStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询房态信息
//
// @param request - QueryRoomStatusRequest
//
// @param headers - QueryRoomStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRoomStatusResponse
func (client *Client) QueryRoomStatusWithOptions(request *QueryRoomStatusRequest, headers *QueryRoomStatusHeaders, runtime *dara.RuntimeOptions) (_result *QueryRoomStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRoomStatus"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/queryRoomStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRoomStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询房态信息
//
// @param request - QueryRoomStatusRequest
//
// @return QueryRoomStatusResponse
func (client *Client) QueryRoomStatus(request *QueryRoomStatusRequest) (_result *QueryRoomStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryRoomStatusHeaders{}
	_result = &QueryRoomStatusResponse{}
	_body, _err := client.QueryRoomStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店场景列表
//
// @param tmpReq - QuerySceneListRequest
//
// @param headers - QuerySceneListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySceneListResponse
func (client *Client) QuerySceneListWithOptions(tmpReq *QuerySceneListRequest, headers *QuerySceneListHeaders, runtime *dara.RuntimeOptions) (_result *QuerySceneListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuerySceneListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SceneStates) {
		request.SceneStatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneStates, dara.String("SceneStates"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SceneTypes) {
		request.SceneTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneTypes, dara.String("SceneTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateInfoIds) {
		request.TemplateInfoIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateInfoIds, dara.String("TemplateInfoIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.SceneStatesShrink) {
		body["SceneStates"] = request.SceneStatesShrink
	}

	if !dara.IsNil(request.SceneTypesShrink) {
		body["SceneTypes"] = request.SceneTypesShrink
	}

	if !dara.IsNil(request.TemplateInfoIdsShrink) {
		body["TemplateInfoIds"] = request.TemplateInfoIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySceneList"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/querySceneList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店场景列表
//
// @param request - QuerySceneListRequest
//
// @return QuerySceneListResponse
func (client *Client) QuerySceneList(request *QuerySceneListRequest) (_result *QuerySceneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QuerySceneListHeaders{}
	_result = &QuerySceneListResponse{}
	_body, _err := client.QuerySceneListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除子账号授权
//
// @param request - RemoveChildAccountAuthRequest
//
// @param headers - RemoveChildAccountAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveChildAccountAuthResponse
func (client *Client) RemoveChildAccountAuthWithOptions(request *RemoveChildAccountAuthRequest, headers *RemoveChildAccountAuthHeaders, runtime *dara.RuntimeOptions) (_result *RemoveChildAccountAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ChildAccountName) {
		body["ChildAccountName"] = request.ChildAccountName
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.TbOpenId) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveChildAccountAuth"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/removeChildAccountAuth"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveChildAccountAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除子账号授权
//
// @param request - RemoveChildAccountAuthRequest
//
// @return RemoveChildAccountAuthResponse
func (client *Client) RemoveChildAccountAuth(request *RemoveChildAccountAuthRequest) (_result *RemoveChildAccountAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RemoveChildAccountAuthHeaders{}
	_result = &RemoveChildAccountAuthResponse{}
	_body, _err := client.RemoveChildAccountAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除酒店项目
//
// @param request - RemoveHotelRequest
//
// @param headers - RemoveHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveHotelResponse
func (client *Client) RemoveHotelWithOptions(request *RemoveHotelRequest, headers *RemoveHotelHeaders, runtime *dara.RuntimeOptions) (_result *RemoveHotelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.TbOpenId) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveHotel"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/removeHotel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除酒店项目
//
// @param request - RemoveHotelRequest
//
// @return RemoveHotelResponse
func (client *Client) RemoveHotel(request *RemoveHotelRequest) (_result *RemoveHotelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RemoveHotelHeaders{}
	_result = &RemoveHotelResponse{}
	_body, _err := client.RemoveHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置欢迎语信息
//
// @param request - ResetWelcomeTextAndMusicRequest
//
// @param headers - ResetWelcomeTextAndMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetWelcomeTextAndMusicResponse
func (client *Client) ResetWelcomeTextAndMusicWithOptions(request *ResetWelcomeTextAndMusicRequest, headers *ResetWelcomeTextAndMusicHeaders, runtime *dara.RuntimeOptions) (_result *ResetWelcomeTextAndMusicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetWelcomeTextAndMusic"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/resetWelcomeTextAndMusic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetWelcomeTextAndMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置欢迎语信息
//
// @param request - ResetWelcomeTextAndMusicRequest
//
// @return ResetWelcomeTextAndMusicResponse
func (client *Client) ResetWelcomeTextAndMusic(request *ResetWelcomeTextAndMusicRequest) (_result *ResetWelcomeTextAndMusicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ResetWelcomeTextAndMusicHeaders{}
	_result = &ResetWelcomeTextAndMusicResponse{}
	_body, _err := client.ResetWelcomeTextAndMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 退房
//
// @param tmpReq - RoomCheckOutRequest
//
// @param headers - RoomCheckOutHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoomCheckOutResponse
func (client *Client) RoomCheckOutWithOptions(tmpReq *RoomCheckOutRequest, headers *RoomCheckOutHeaders, runtime *dara.RuntimeOptions) (_result *RoomCheckOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RoomCheckOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RoomCheckOut"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/roomCheckOut"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RoomCheckOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退房
//
// @param request - RoomCheckOutRequest
//
// @return RoomCheckOutResponse
func (client *Client) RoomCheckOut(request *RoomCheckOutRequest) (_result *RoomCheckOutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RoomCheckOutHeaders{}
	_result = &RoomCheckOutResponse{}
	_body, _err := client.RoomCheckOutWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交酒店订单
//
// @param tmpReq - SubmitHotelOrderRequest
//
// @param headers - SubmitHotelOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHotelOrderResponse
func (client *Client) SubmitHotelOrderWithOptions(tmpReq *SubmitHotelOrderRequest, headers *SubmitHotelOrderHeaders, runtime *dara.RuntimeOptions) (_result *SubmitHotelOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitHotelOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitHotelOrder"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/submitHotelOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitHotelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交酒店订单
//
// @param request - SubmitHotelOrderRequest
//
// @return SubmitHotelOrderResponse
func (client *Client) SubmitHotelOrder(request *SubmitHotelOrderRequest) (_result *SubmitHotelOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SubmitHotelOrderHeaders{}
	_result = &SubmitHotelOrderResponse{}
	_body, _err := client.SubmitHotelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步客控设备状态到主控设备
//
// @param request - SyncDeviceStatusWithAkRequest
//
// @param headers - SyncDeviceStatusWithAkHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDeviceStatusWithAkResponse
func (client *Client) SyncDeviceStatusWithAkWithOptions(request *SyncDeviceStatusWithAkRequest, headers *SyncDeviceStatusWithAkHeaders, runtime *dara.RuntimeOptions) (_result *SyncDeviceStatusWithAkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryCnName) {
		body["CategoryCnName"] = request.CategoryCnName
	}

	if !dara.IsNil(request.CategoryEnName) {
		body["CategoryEnName"] = request.CategoryEnName
	}

	if !dara.IsNil(request.DeviceName) {
		body["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Location) {
		body["Location"] = request.Location
	}

	if !dara.IsNil(request.LocationCnName) {
		body["LocationCnName"] = request.LocationCnName
	}

	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.RoomNo) {
		body["RoomNo"] = request.RoomNo
	}

	if !dara.IsNil(request.Switch) {
		body["Switch"] = request.Switch
	}

	if !dara.IsNil(request.FanSpeed) {
		body["fanSpeed"] = request.FanSpeed
	}

	if !dara.IsNil(request.Mode) {
		body["mode"] = request.Mode
	}

	if !dara.IsNil(request.RoomTemperature) {
		body["roomTemperature"] = request.RoomTemperature
	}

	if !dara.IsNil(request.Temperature) {
		body["temperature"] = request.Temperature
	}

	if !dara.IsNil(request.Value) {
		body["value"] = request.Value
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDeviceStatusWithAk"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/syncDeviceStatusWithAk"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDeviceStatusWithAkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步客控设备状态到主控设备
//
// @param request - SyncDeviceStatusWithAkRequest
//
// @return SyncDeviceStatusWithAkResponse
func (client *Client) SyncDeviceStatusWithAk(request *SyncDeviceStatusWithAkRequest) (_result *SyncDeviceStatusWithAkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SyncDeviceStatusWithAkHeaders{}
	_result = &SyncDeviceStatusWithAkResponse{}
	_body, _err := client.SyncDeviceStatusWithAkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改基础信息问答
//
// @param request - UpdateBasicInfoQARequest
//
// @param headers - UpdateBasicInfoQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicInfoQAResponse
func (client *Client) UpdateBasicInfoQAWithOptions(request *UpdateBasicInfoQARequest, headers *UpdateBasicInfoQAHeaders, runtime *dara.RuntimeOptions) (_result *UpdateBasicInfoQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckInTime) {
		body["CheckInTime"] = request.CheckInTime
	}

	if !dara.IsNil(request.CheckOutTime) {
		body["CheckOutTime"] = request.CheckOutTime
	}

	if !dara.IsNil(request.HotelAddress) {
		body["HotelAddress"] = request.HotelAddress
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.HotelIntroduction) {
		body["HotelIntroduction"] = request.HotelIntroduction
	}

	if !dara.IsNil(request.HotelMember) {
		body["HotelMember"] = request.HotelMember
	}

	if !dara.IsNil(request.HotelService) {
		body["HotelService"] = request.HotelService
	}

	if !dara.IsNil(request.ParkingExpenses) {
		body["ParkingExpenses"] = request.ParkingExpenses
	}

	if !dara.IsNil(request.ParkingPosition) {
		body["ParkingPosition"] = request.ParkingPosition
	}

	if !dara.IsNil(request.PhoneNumber) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.WifiName) {
		body["WifiName"] = request.WifiName
	}

	if !dara.IsNil(request.WifiPassword) {
		body["WifiPassword"] = request.WifiPassword
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBasicInfoQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateBasicInfoQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBasicInfoQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改基础信息问答
//
// @param request - UpdateBasicInfoQARequest
//
// @return UpdateBasicInfoQAResponse
func (client *Client) UpdateBasicInfoQA(request *UpdateBasicInfoQARequest) (_result *UpdateBasicInfoQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateBasicInfoQAHeaders{}
	_result = &UpdateBasicInfoQAResponse{}
	_body, _err := client.UpdateBasicInfoQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自定义问答
//
// @param tmpReq - UpdateCustomQARequest
//
// @param headers - UpdateCustomQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomQAResponse
func (client *Client) UpdateCustomQAWithOptions(tmpReq *UpdateCustomQARequest, headers *UpdateCustomQAHeaders, runtime *dara.RuntimeOptions) (_result *UpdateCustomQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCustomQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Answers) {
		request.AnswersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Answers, dara.String("Answers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KeyWords) {
		request.KeyWordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KeyWords, dara.String("KeyWords"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SupplementaryQuestions) {
		request.SupplementaryQuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupplementaryQuestions, dara.String("SupplementaryQuestions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnswersShrink) {
		body["Answers"] = request.AnswersShrink
	}

	if !dara.IsNil(request.CustomQAId) {
		body["CustomQAId"] = request.CustomQAId
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.KeyWordsShrink) {
		body["KeyWords"] = request.KeyWordsShrink
	}

	if !dara.IsNil(request.MajorQuestion) {
		body["MajorQuestion"] = request.MajorQuestion
	}

	if !dara.IsNil(request.SupplementaryQuestionsShrink) {
		body["SupplementaryQuestions"] = request.SupplementaryQuestionsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateCustomQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义问答
//
// @param request - UpdateCustomQARequest
//
// @return UpdateCustomQAResponse
func (client *Client) UpdateCustomQA(request *UpdateCustomQARequest) (_result *UpdateCustomQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateCustomQAHeaders{}
	_result = &UpdateCustomQAResponse{}
	_body, _err := client.UpdateCustomQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改酒店项目
//
// @param tmpReq - UpdateHotelRequest
//
// @param headers - UpdateHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelResponse
func (client *Client) UpdateHotelWithOptions(tmpReq *UpdateHotelRequest, headers *UpdateHotelHeaders, runtime *dara.RuntimeOptions) (_result *UpdateHotelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RelatedPks) {
		request.RelatedPksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedPks, dara.String("RelatedPks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EstOpenTime) {
		body["EstOpenTime"] = request.EstOpenTime
	}

	if !dara.IsNil(request.HotelAddress) {
		body["HotelAddress"] = request.HotelAddress
	}

	if !dara.IsNil(request.HotelEmail) {
		body["HotelEmail"] = request.HotelEmail
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.HotelName) {
		body["HotelName"] = request.HotelName
	}

	if !dara.IsNil(request.PhoneNumber) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.RelatedPksShrink) {
		body["RelatedPks"] = request.RelatedPksShrink
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RoomNum) {
		body["RoomNum"] = request.RoomNum
	}

	if !dara.IsNil(request.TbOpenId) {
		body["TbOpenId"] = request.TbOpenId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHotel"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateHotel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改酒店项目
//
// @param request - UpdateHotelRequest
//
// @return UpdateHotelResponse
func (client *Client) UpdateHotel(request *UpdateHotelRequest) (_result *UpdateHotelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateHotelHeaders{}
	_result = &UpdateHotelResponse{}
	_body, _err := client.UpdateHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改酒店闹钟
//
// @param tmpReq - UpdateHotelAlarmRequest
//
// @param headers - UpdateHotelAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelAlarmResponse
func (client *Client) UpdateHotelAlarmWithOptions(tmpReq *UpdateHotelAlarmRequest, headers *UpdateHotelAlarmHeaders, runtime *dara.RuntimeOptions) (_result *UpdateHotelAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHotelAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Alarms) {
		request.AlarmsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Alarms, dara.String("Alarms"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleInfo) {
		request.ScheduleInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleInfo, dara.String("ScheduleInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmsShrink) {
		body["Alarms"] = request.AlarmsShrink
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.ScheduleInfoShrink) {
		body["ScheduleInfo"] = request.ScheduleInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHotelAlarm"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateHotelAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHotelAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改酒店闹钟
//
// @param request - UpdateHotelAlarmRequest
//
// @return UpdateHotelAlarmResponse
func (client *Client) UpdateHotelAlarm(request *UpdateHotelAlarmRequest) (_result *UpdateHotelAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateHotelAlarmHeaders{}
	_result = &UpdateHotelAlarmResponse{}
	_body, _err := client.UpdateHotelAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景预订编辑
//
// @param tmpReq - UpdateHotelSceneBookItemRequest
//
// @param headers - UpdateHotelSceneBookItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelSceneBookItemResponse
func (client *Client) UpdateHotelSceneBookItemWithOptions(tmpReq *UpdateHotelSceneBookItemRequest, headers *UpdateHotelSceneBookItemHeaders, runtime *dara.RuntimeOptions) (_result *UpdateHotelSceneBookItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHotelSceneBookItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateHotelSceneBookReq) {
		request.UpdateHotelSceneBookReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateHotelSceneBookReq, dara.String("UpdateHotelSceneBookReq"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UpdateHotelSceneBookReqShrink) {
		query["UpdateHotelSceneBookReq"] = request.UpdateHotelSceneBookReqShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHotelSceneBookItem"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateHotelSceneBookItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHotelSceneBookItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景预订编辑
//
// @param request - UpdateHotelSceneBookItemRequest
//
// @return UpdateHotelSceneBookItemResponse
func (client *Client) UpdateHotelSceneBookItem(request *UpdateHotelSceneBookItemRequest) (_result *UpdateHotelSceneBookItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateHotelSceneBookItemHeaders{}
	_result = &UpdateHotelSceneBookItemResponse{}
	_body, _err := client.UpdateHotelSceneBookItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店场景修改（开启/关闭/编辑）
//
// @param tmpReq - UpdateHotelSceneItemRequest
//
// @param headers - UpdateHotelSceneItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotelSceneItemResponse
func (client *Client) UpdateHotelSceneItemWithOptions(tmpReq *UpdateHotelSceneItemRequest, headers *UpdateHotelSceneItemHeaders, runtime *dara.RuntimeOptions) (_result *UpdateHotelSceneItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateHotelSceneItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateHotelSceneOperateReq) {
		request.UpdateHotelSceneOperateReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateHotelSceneOperateReq, dara.String("UpdateHotelSceneOperateReq"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateHotelSceneReq) {
		request.UpdateHotelSceneReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateHotelSceneReq, dara.String("UpdateHotelSceneReq"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UpdateHotelSceneOperateReqShrink) {
		query["UpdateHotelSceneOperateReq"] = request.UpdateHotelSceneOperateReqShrink
	}

	if !dara.IsNil(request.UpdateHotelSceneReqShrink) {
		query["UpdateHotelSceneReq"] = request.UpdateHotelSceneReqShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHotelSceneItem"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateHotelSceneItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHotelSceneItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店场景修改（开启/关闭/编辑）
//
// @param request - UpdateHotelSceneItemRequest
//
// @return UpdateHotelSceneItemResponse
func (client *Client) UpdateHotelSceneItem(request *UpdateHotelSceneItemRequest) (_result *UpdateHotelSceneItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateHotelSceneItemHeaders{}
	_result = &UpdateHotelSceneItemResponse{}
	_body, _err := client.UpdateHotelSceneItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新消息通知模板
//
// @param request - UpdateMessageTemplateRequest
//
// @param headers - UpdateMessageTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageTemplateResponse
func (client *Client) UpdateMessageTemplateWithOptions(request *UpdateMessageTemplateRequest, headers *UpdateMessageTemplateHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMessageTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateDetail) {
		body["TemplateDetail"] = request.TemplateDetail
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMessageTemplate"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateMessageTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMessageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消息通知模板
//
// @param request - UpdateMessageTemplateRequest
//
// @return UpdateMessageTemplateResponse
func (client *Client) UpdateMessageTemplate(request *UpdateMessageTemplateRequest) (_result *UpdateMessageTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMessageTemplateHeaders{}
	_result = &UpdateMessageTemplateResponse{}
	_body, _err := client.UpdateMessageTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改酒店自定义rcu场景
//
// @param tmpReq - UpdateRcuSceneRequest
//
// @param headers - UpdateRcuSceneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRcuSceneResponse
func (client *Client) UpdateRcuSceneWithOptions(tmpReq *UpdateRcuSceneRequest, headers *UpdateRcuSceneHeaders, runtime *dara.RuntimeOptions) (_result *UpdateRcuSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRcuSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SceneRelationExtDTO) {
		request.SceneRelationExtDTOShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneRelationExtDTO, dara.String("SceneRelationExtDTO"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SceneRelationExtDTOShrink) {
		body["SceneRelationExtDTO"] = request.SceneRelationExtDTOShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRcuScene"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateRcuScene"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRcuSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改酒店自定义rcu场景
//
// @param request - UpdateRcuSceneRequest
//
// @return UpdateRcuSceneResponse
func (client *Client) UpdateRcuScene(request *UpdateRcuSceneRequest) (_result *UpdateRcuSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateRcuSceneHeaders{}
	_result = &UpdateRcuSceneResponse{}
	_body, _err := client.UpdateRcuSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改服务设施问答
//
// @param request - UpdateServiceQARequest
//
// @param headers - UpdateServiceQAHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceQAResponse
func (client *Client) UpdateServiceQAWithOptions(request *UpdateServiceQARequest, headers *UpdateServiceQAHeaders, runtime *dara.RuntimeOptions) (_result *UpdateServiceQAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Answer) {
		body["Answer"] = request.Answer
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.ServiceQAId) {
		body["ServiceQAId"] = request.ServiceQAId
	}

	if !dara.IsNil(request.IsActive) {
		body["isActive"] = request.IsActive
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceQA"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateServiceQA"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改服务设施问答
//
// @param request - UpdateServiceQARequest
//
// @return UpdateServiceQAResponse
func (client *Client) UpdateServiceQA(request *UpdateServiceQARequest) (_result *UpdateServiceQAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateServiceQAHeaders{}
	_result = &UpdateServiceQAResponse{}
	_body, _err := client.UpdateServiceQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改工单
//
// @param request - UpdateTicketRequest
//
// @param headers - UpdateTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketResponse
func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, headers *UpdateTicketHeaders, runtime *dara.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupKey) {
		body["GroupKey"] = request.GroupKey
	}

	if !dara.IsNil(request.HotelId) {
		body["HotelId"] = request.HotelId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTicket"),
		Version:     dara.String("ip_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ip/updateTicket"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改工单
//
// @param request - UpdateTicketRequest
//
// @return UpdateTicketResponse
func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateTicketHeaders{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
