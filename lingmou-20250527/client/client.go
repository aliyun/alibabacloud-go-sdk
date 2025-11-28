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
	client.Endpoint, _err = client.GetEndpoint(dara.String("lingmou"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 关闭会话实例session
//
// @param tmpReq - CloseChatInstanceSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseChatInstanceSessionsResponse
func (client *Client) CloseChatInstanceSessionsWithOptions(instanceId *string, tmpReq *CloseChatInstanceSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseChatInstanceSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloseChatInstanceSessionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionIds) {
		request.SessionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionIds, dara.String("sessionIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionIdsShrink) {
		body["sessionIds"] = request.SessionIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseChatInstanceSessions"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sessions/close"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseChatInstanceSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭会话实例session
//
// @param request - CloseChatInstanceSessionsRequest
//
// @return CloseChatInstanceSessionsResponse
func (client *Client) CloseChatInstanceSessions(instanceId *string, request *CloseChatInstanceSessionsRequest) (_result *CloseChatInstanceSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseChatInstanceSessionsResponse{}
	_body, _err := client.CloseChatInstanceSessionsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户确认
//
// @param request - ConfirmTrainPicAvatarRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmTrainPicAvatarResponse
func (client *Client) ConfirmTrainPicAvatarWithOptions(request *ConfirmTrainPicAvatarRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConfirmTrainPicAvatarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarId) {
		query["avatarId"] = request.AvatarId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmTrainPicAvatar"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/train/confirmTrainPicAvatar"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmTrainPicAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户确认
//
// @param request - ConfirmTrainPicAvatarRequest
//
// @return ConfirmTrainPicAvatarResponse
func (client *Client) ConfirmTrainPicAvatar(request *ConfirmTrainPicAvatarRequest) (_result *ConfirmTrainPicAvatarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmTrainPicAvatarResponse{}
	_body, _err := client.ConfirmTrainPicAvatarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建背景素材
//
// @param request - CreateBackgroundPicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackgroundPicResponse
func (client *Client) CreateBackgroundPicWithOptions(request *CreateBackgroundPicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBackgroundPicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filename) {
		query["filename"] = request.Filename
	}

	if !dara.IsNil(request.OssKey) {
		query["ossKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackgroundPic"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/createBackgroundPic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackgroundPicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建背景素材
//
// @param request - CreateBackgroundPicRequest
//
// @return CreateBackgroundPicResponse
func (client *Client) CreateBackgroundPic(request *CreateBackgroundPicRequest) (_result *CreateBackgroundPicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBackgroundPicResponse{}
	_body, _err := client.CreateBackgroundPicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建播报贴图
//
// @param request - CreateBroadcastStickerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBroadcastStickerResponse
func (client *Client) CreateBroadcastStickerWithOptions(request *CreateBroadcastStickerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBroadcastStickerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.OssKey) {
		body["ossKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBroadcastSticker"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/customer/broadcast/material/sticker/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBroadcastStickerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建播报贴图
//
// @param request - CreateBroadcastStickerRequest
//
// @return CreateBroadcastStickerResponse
func (client *Client) CreateBroadcastSticker(request *CreateBroadcastStickerRequest) (_result *CreateBroadcastStickerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBroadcastStickerResponse{}
	_body, _err := client.CreateBroadcastStickerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 基于模板创建播报视频
//
// @param request - CreateBroadcastVideoFromTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBroadcastVideoFromTemplateResponse
func (client *Client) CreateBroadcastVideoFromTemplateWithOptions(request *CreateBroadcastVideoFromTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBroadcastVideoFromTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		body["templateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	if !dara.IsNil(request.VideoOptions) {
		body["videoOptions"] = request.VideoOptions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBroadcastVideoFromTemplate"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/amp/customer/broadcast/video/createFromTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBroadcastVideoFromTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 基于模板创建播报视频
//
// @param request - CreateBroadcastVideoFromTemplateRequest
//
// @return CreateBroadcastVideoFromTemplateResponse
func (client *Client) CreateBroadcastVideoFromTemplate(request *CreateBroadcastVideoFromTemplateRequest) (_result *CreateBroadcastVideoFromTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBroadcastVideoFromTemplateResponse{}
	_body, _err := client.CreateBroadcastVideoFromTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 背景配置
//
// @param request - CreateChatConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatConfigResponse
func (client *Client) CreateChatConfigWithOptions(request *CreateChatConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarId) {
		query["avatarId"] = request.AvatarId
	}

	if !dara.IsNil(request.BackgroundId) {
		query["backgroundId"] = request.BackgroundId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatConfig"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/createChatConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 背景配置
//
// @param request - CreateChatConfigRequest
//
// @return CreateChatConfigResponse
func (client *Client) CreateChatConfig(request *CreateChatConfigRequest) (_result *CreateChatConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateChatConfigResponse{}
	_body, _err := client.CreateChatConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数字人会话
//
// @param request - CreateChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatSessionResponse
func (client *Client) CreateChatSessionWithOptions(id *string, request *CreateChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.License) {
		query["license"] = request.License
	}

	if !dara.IsNil(request.Platform) {
		query["platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatSession"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/init/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数字人会话
//
// @param request - CreateChatSessionRequest
//
// @return CreateChatSessionResponse
func (client *Client) CreateChatSession(id *string, request *CreateChatSessionRequest) (_result *CreateChatSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateChatSessionResponse{}
	_body, _err := client.CreateChatSessionWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建对话免训照片数字人
//
// @param request - CreateNoTrainPicAvatarRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNoTrainPicAvatarResponse
func (client *Client) CreateNoTrainPicAvatarWithOptions(request *CreateNoTrainPicAvatarRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateNoTrainPicAvatarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expressiveness) {
		query["expressiveness"] = request.Expressiveness
	}

	if !dara.IsNil(request.Gender) {
		query["gender"] = request.Gender
	}

	if !dara.IsNil(request.GenerateAssets) {
		query["generateAssets"] = request.GenerateAssets
	}

	if !dara.IsNil(request.ImageOssPath) {
		query["imageOssPath"] = request.ImageOssPath
	}

	if !dara.IsNil(request.JwtToken) {
		query["jwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Transparent) {
		query["transparent"] = request.Transparent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNoTrainPicAvatar"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/createNoTrainPicAvatar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNoTrainPicAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建对话免训照片数字人
//
// @param request - CreateNoTrainPicAvatarRequest
//
// @return CreateNoTrainPicAvatarResponse
func (client *Client) CreateNoTrainPicAvatar(request *CreateNoTrainPicAvatarRequest) (_result *CreateNoTrainPicAvatarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateNoTrainPicAvatarResponse{}
	_body, _err := client.CreateNoTrainPicAvatarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建TTS音色
//
// @param request - CreateTTSVoiceCustomRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTTSVoiceCustomResponse
func (client *Client) CreateTTSVoiceCustomWithOptions(request *CreateTTSVoiceCustomRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTTSVoiceCustomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.Gender) {
		query["gender"] = request.Gender
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OssKey) {
		query["ossKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTTSVoiceCustom"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/voice/createTTSVoiceCustom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTTSVoiceCustomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建TTS音色
//
// @param request - CreateTTSVoiceCustomRequest
//
// @return CreateTTSVoiceCustomResponse
func (client *Client) CreateTTSVoiceCustom(request *CreateTTSVoiceCustomRequest) (_result *CreateTTSVoiceCustomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTTSVoiceCustomResponse{}
	_body, _err := client.CreateTTSVoiceCustomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建图片训练数字人
//
// @param request - CreateTrainPicAvatarRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrainPicAvatarResponse
func (client *Client) CreateTrainPicAvatarWithOptions(request *CreateTrainPicAvatarRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrainPicAvatarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Gender) {
		query["gender"] = request.Gender
	}

	if !dara.IsNil(request.GenerateAssets) {
		query["generateAssets"] = request.GenerateAssets
	}

	if !dara.IsNil(request.ImageOssPath) {
		query["imageOssPath"] = request.ImageOssPath
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Transparent) {
		query["transparent"] = request.Transparent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrainPicAvatar"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/train/createTrainPicAvatar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrainPicAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片训练数字人
//
// @param request - CreateTrainPicAvatarRequest
//
// @return CreateTrainPicAvatarResponse
func (client *Client) CreateTrainPicAvatar(request *CreateTrainPicAvatarRequest) (_result *CreateTrainPicAvatarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrainPicAvatarResponse{}
	_body, _err := client.CreateTrainPicAvatarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询播报模板详情
//
// @param request - GetBroadcastTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBroadcastTemplateResponse
func (client *Client) GetBroadcastTemplateWithOptions(request *GetBroadcastTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBroadcastTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBroadcastTemplate"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/customer/broadcast/template/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBroadcastTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询播报模板详情
//
// @param request - GetBroadcastTemplateRequest
//
// @return GetBroadcastTemplateResponse
func (client *Client) GetBroadcastTemplate(request *GetBroadcastTemplateRequest) (_result *GetBroadcastTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBroadcastTemplateResponse{}
	_body, _err := client.GetBroadcastTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询图片训练数字人的状态
//
// @param request - GetTrainPicAvatarStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainPicAvatarStatusResponse
func (client *Client) GetTrainPicAvatarStatusWithOptions(request *GetTrainPicAvatarStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainPicAvatarStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarId) {
		query["avatarId"] = request.AvatarId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrainPicAvatarStatus"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/train/getTrainPicAvatarStatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrainPicAvatarStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询图片训练数字人的状态
//
// @param request - GetTrainPicAvatarStatusRequest
//
// @return GetTrainPicAvatarStatusResponse
func (client *Client) GetTrainPicAvatarStatus(request *GetTrainPicAvatarStatusRequest) (_result *GetTrainPicAvatarStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainPicAvatarStatusResponse{}
	_body, _err := client.GetTrainPicAvatarStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取对话免训图片素材上传凭证
//
// @param request - GetUploadPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadPolicyResponse
func (client *Client) GetUploadPolicyWithOptions(request *GetUploadPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUploadPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["jwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadPolicy"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/getUploadPolicy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取对话免训图片素材上传凭证
//
// @param request - GetUploadPolicyRequest
//
// @return GetUploadPolicyResponse
func (client *Client) GetUploadPolicy(request *GetUploadPolicyRequest) (_result *GetUploadPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUploadPolicyResponse{}
	_body, _err := client.GetUploadPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举播报模板
//
// @param request - ListBroadcastTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBroadcastTemplatesResponse
func (client *Client) ListBroadcastTemplatesWithOptions(request *ListBroadcastTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBroadcastTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBroadcastTemplates"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/customer/broadcast/template/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBroadcastTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举播报模板
//
// @param request - ListBroadcastTemplatesRequest
//
// @return ListBroadcastTemplatesResponse
func (client *Client) ListBroadcastTemplates(request *ListBroadcastTemplatesRequest) (_result *ListBroadcastTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBroadcastTemplatesResponse{}
	_body, _err := client.ListBroadcastTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据视频id批量查询播报视频（最多查询100个）
//
// @param tmpReq - ListBroadcastVideosByIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBroadcastVideosByIdResponse
func (client *Client) ListBroadcastVideosByIdWithOptions(tmpReq *ListBroadcastVideosByIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBroadcastVideosByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListBroadcastVideosByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VideoIds) {
		request.VideoIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoIds, dara.String("videoIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoIdsShrink) {
		query["videoIds"] = request.VideoIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBroadcastVideosById"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/amp/customer/broadcast/video/batchQuery"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBroadcastVideosByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据视频id批量查询播报视频（最多查询100个）
//
// @param request - ListBroadcastVideosByIdRequest
//
// @return ListBroadcastVideosByIdResponse
func (client *Client) ListBroadcastVideosById(request *ListBroadcastVideosByIdRequest) (_result *ListBroadcastVideosByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBroadcastVideosByIdResponse{}
	_body, _err := client.ListBroadcastVideosByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举私有TTS音色
//
// @param request - ListPrivateTTSVoicesCustomRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateTTSVoicesCustomResponse
func (client *Client) ListPrivateTTSVoicesCustomWithOptions(request *ListPrivateTTSVoicesCustomRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrivateTTSVoicesCustomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateTTSVoicesCustom"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/voice/listPrivateTTSVoicesCustom"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateTTSVoicesCustomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举私有TTS音色
//
// @param request - ListPrivateTTSVoicesCustomRequest
//
// @return ListPrivateTTSVoicesCustomResponse
func (client *Client) ListPrivateTTSVoicesCustom(request *ListPrivateTTSVoicesCustomRequest) (_result *ListPrivateTTSVoicesCustomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPrivateTTSVoicesCustomResponse{}
	_body, _err := client.ListPrivateTTSVoicesCustomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询底板素材
//
// @param request - ListTemplateMaterialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateMaterialResponse
func (client *Client) ListTemplateMaterialWithOptions(request *ListTemplateMaterialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTemplateMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.TemplateIds) {
		query["templateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplateMaterial"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/train/listTemplateMaterial"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询底板素材
//
// @param request - ListTemplateMaterialRequest
//
// @return ListTemplateMaterialResponse
func (client *Client) ListTemplateMaterial(request *ListTemplateMaterialRequest) (_result *ListTemplateMaterialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplateMaterialResponse{}
	_body, _err := client.ListTemplateMaterialWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话实例session
//
// @param tmpReq - QueryChatInstanceSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChatInstanceSessionsResponse
func (client *Client) QueryChatInstanceSessionsWithOptions(instanceId *string, tmpReq *QueryChatInstanceSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryChatInstanceSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryChatInstanceSessionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionIds) {
		request.SessionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionIds, dara.String("sessionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionIdsShrink) {
		query["sessionIds"] = request.SessionIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryChatInstanceSessions"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryChatInstanceSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话实例session
//
// @param request - QueryChatInstanceSessionsRequest
//
// @return QueryChatInstanceSessionsResponse
func (client *Client) QueryChatInstanceSessions(instanceId *string, request *QueryChatInstanceSessionsRequest) (_result *QueryChatInstanceSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryChatInstanceSessionsResponse{}
	_body, _err := client.QueryChatInstanceSessionsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
