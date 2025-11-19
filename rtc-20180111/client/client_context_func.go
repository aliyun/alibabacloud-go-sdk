// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AddRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordTemplateResponse
func (client *Client) AddRecordTemplateWithContext(ctx context.Context, request *AddRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.DelayStopTime) {
		query["DelayStopTime"] = request.DelayStopTime
	}

	if !dara.IsNil(request.EnableM3u8DateTime) {
		query["EnableM3u8DateTime"] = request.EnableM3u8DateTime
	}

	if !dara.IsNil(request.FileSplitInterval) {
		query["FileSplitInterval"] = request.FileSplitInterval
	}

	if !dara.IsNil(request.Formats) {
		query["Formats"] = request.Formats
	}

	if !dara.IsNil(request.HttpCallbackUrl) {
		query["HttpCallbackUrl"] = request.HttpCallbackUrl
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MnsQueue) {
		query["MnsQueue"] = request.MnsQueue
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssFilePrefix) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecordTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用智能体模版
//
// @param tmpReq - CreateAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppAgentTemplateResponse
func (client *Client) CreateAppAgentTemplateWithContext(ctx context.Context, tmpReq *CreateAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppAgentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppAgentTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentSilenceConfig) {
		request.AgentSilenceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentSilenceConfig, dara.String("AgentSilenceConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AmbientSoundConfig) {
		request.AmbientSoundConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AmbientSoundConfig, dara.String("AmbientSoundConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AsrConfig) {
		request.AsrConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrConfig, dara.String("AsrConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BackChannelConfig) {
		request.BackChannelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackChannelConfig, dara.String("BackChannelConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InterruptConfig) {
		request.InterruptConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterruptConfig, dara.String("InterruptConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LlmConfig) {
		request.LlmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmConfig, dara.String("LlmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSilenceConfigShrink) {
		query["AgentSilenceConfig"] = request.AgentSilenceConfigShrink
	}

	if !dara.IsNil(request.AmbientSoundConfigShrink) {
		query["AmbientSoundConfig"] = request.AmbientSoundConfigShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AsrConfigShrink) {
		query["AsrConfig"] = request.AsrConfigShrink
	}

	if !dara.IsNil(request.BackChannelConfigShrink) {
		query["BackChannelConfig"] = request.BackChannelConfigShrink
	}

	if !dara.IsNil(request.ChatMode) {
		query["ChatMode"] = request.ChatMode
	}

	if !dara.IsNil(request.Greeting) {
		query["Greeting"] = request.Greeting
	}

	if !dara.IsNil(request.InterruptConfigShrink) {
		query["InterruptConfig"] = request.InterruptConfigShrink
	}

	if !dara.IsNil(request.InterruptMode) {
		query["InterruptMode"] = request.InterruptMode
	}

	if !dara.IsNil(request.LlmConfigShrink) {
		query["LlmConfig"] = request.LlmConfigShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TtsConfigShrink) {
		query["TtsConfig"] = request.TtsConfigShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppAgentTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppAgentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增app自定义布局
//
// @param tmpReq - CreateAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppLayoutResponse
func (client *Client) CreateAppLayoutWithContext(ctx context.Context, tmpReq *CreateAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *CreateAppLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Layout) {
		request.LayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Layout, dara.String("Layout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.LayoutShrink) {
		query["Layout"] = request.LayoutShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加应用录制模版
//
// @param tmpReq - CreateAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppRecordTemplateResponse
func (client *Client) CreateAppRecordTemplateWithContext(ctx context.Context, tmpReq *CreateAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordTemplate) {
		request.RecordTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordTemplate, dara.String("RecordTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordTemplateShrink) {
		query["RecordTemplate"] = request.RecordTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppRecordTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用推流模版
//
// @param tmpReq - CreateAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppStreamingOutTemplateResponse
func (client *Client) CreateAppStreamingOutTemplateWithContext(ctx context.Context, tmpReq *CreateAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppStreamingOutTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamingOutTemplate) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, dara.String("StreamingOutTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StreamingOutTemplateShrink) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppStreamingOutTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用合流模版
//
// @param tmpReq - CreateAppViewTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppViewTemplateResponse
func (client *Client) CreateAppViewTemplateWithContext(ctx context.Context, tmpReq *CreateAppViewTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppViewTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppViewTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppViewTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppViewTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoLiveStreamRuleResponse
func (client *Client) CreateAutoLiveStreamRuleWithContext(ctx context.Context, request *CreateAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.ChannelIdPrefixes) {
		query["ChannelIdPrefixes"] = request.ChannelIdPrefixes
	}

	if !dara.IsNil(request.ChannelIds) {
		query["ChannelIds"] = request.ChannelIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加纪要热词表
//
// @param tmpReq - CreateCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudNotePhrasesResponse
func (client *Client) CreateCloudNotePhrasesWithContext(ctx context.Context, tmpReq *CreateCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Phrase) {
		request.PhraseShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Phrase, dara.String("Phrase"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PhraseShrink) {
		query["Phrase"] = request.PhraseShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudNotePhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventSubscribeResponse
func (client *Client) CreateEventSubscribeWithContext(ctx context.Context, request *CreateEventSubscribeRequest, runtime *dara.RuntimeOptions) (_result *CreateEventSubscribeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Events) {
		query["Events"] = request.Events
	}

	if !dara.IsNil(request.NeedCallbackAuth) {
		query["NeedCallbackAuth"] = request.NeedCallbackAuth
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventSubscribe"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventSubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMPULayoutResponse
func (client *Client) CreateMPULayoutWithContext(ctx context.Context, request *CreateMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *CreateMPULayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AudioMixCount) {
		query["AudioMixCount"] = request.AudioMixCount
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMPULayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMPULayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用智能体模版
//
// @param request - DeleteAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppAgentTemplateResponse
func (client *Client) DeleteAppAgentTemplateWithContext(ctx context.Context, request *DeleteAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppAgentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppAgentTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppAgentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除app自定义布局
//
// @param tmpReq - DeleteAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppLayoutResponse
func (client *Client) DeleteAppLayoutWithContext(ctx context.Context, tmpReq *DeleteAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAppLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Layout) {
		request.LayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Layout, dara.String("Layout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.LayoutShrink) {
		query["Layout"] = request.LayoutShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用录制模版
//
// @param tmpReq - DeleteAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppRecordTemplateResponse
func (client *Client) DeleteAppRecordTemplateWithContext(ctx context.Context, tmpReq *DeleteAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAppRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppRecordTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用推流模版
//
// @param tmpReq - DeleteAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppStreamingOutTemplateResponse
func (client *Client) DeleteAppStreamingOutTemplateWithContext(ctx context.Context, tmpReq *DeleteAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppStreamingOutTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamingOutTemplate) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, dara.String("StreamingOutTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StreamingOutTemplateShrink) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppStreamingOutTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用合流模版
//
// @param tmpReq - DeleteAppViewTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppViewTemplateResponse
func (client *Client) DeleteAppViewTemplateWithContext(ctx context.Context, tmpReq *DeleteAppViewTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppViewTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAppViewTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppViewTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppViewTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoLiveStreamRuleResponse
func (client *Client) DeleteAutoLiveStreamRuleWithContext(ctx context.Context, request *DeleteAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannelWithContext(ctx context.Context, request *DeleteChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除纪要热词表
//
// @param tmpReq - DeleteCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudNotePhrasesResponse
func (client *Client) DeleteCloudNotePhrasesWithContext(ctx context.Context, tmpReq *DeleteCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Phrase) {
		request.PhraseShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Phrase, dara.String("Phrase"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PhraseShrink) {
		query["Phrase"] = request.PhraseShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudNotePhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventSubscribeResponse
func (client *Client) DeleteEventSubscribeWithContext(ctx context.Context, request *DeleteEventSubscribeRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventSubscribeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SubscribeId) {
		query["SubscribeId"] = request.SubscribeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventSubscribe"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventSubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMPULayoutResponse
func (client *Client) DeleteMPULayoutWithContext(ctx context.Context, request *DeleteMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteMPULayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMPULayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMPULayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordTemplateResponse
func (client *Client) DeleteRecordTemplateWithContext(ctx context.Context, request *DeleteRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecordTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用智能体开关
//
// @param request - DescribeAppAgentFunctionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAgentFunctionStatusResponse
func (client *Client) DescribeAppAgentFunctionStatusWithContext(ctx context.Context, request *DescribeAppAgentFunctionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAgentFunctionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppAgentFunctionStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppAgentFunctionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用智能体模版列表
//
// @param request - DescribeAppAgentTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAgentTemplatesResponse
func (client *Client) DescribeAppAgentTemplatesWithContext(ctx context.Context, request *DescribeAppAgentTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAgentTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppAgentTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppAgentTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看app回调开关
//
// @param request - DescribeAppCallStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppCallStatusResponse
func (client *Client) DescribeAppCallStatusWithContext(ctx context.Context, request *DescribeAppCallStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppCallStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppCallStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppCallStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取app回调密钥
//
// @param request - DescribeAppCallbackSecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppCallbackSecretKeyResponse
func (client *Client) DescribeAppCallbackSecretKeyWithContext(ctx context.Context, request *DescribeAppCallbackSecretKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppCallbackSecretKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppCallbackSecretKey"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppCallbackSecretKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看AppKey
//
// @param request - DescribeAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppKeyResponse
func (client *Client) DescribeAppKeyWithContext(ctx context.Context, request *DescribeAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppKey"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询app自定义布局
//
// @param tmpReq - DescribeAppLayoutsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppLayoutsResponse
func (client *Client) DescribeAppLayoutsWithContext(ctx context.Context, tmpReq *DescribeAppLayoutsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppLayoutsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAppLayoutsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppLayouts"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppLayoutsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看应用旁路开关
//
// @param request - DescribeAppLiveStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppLiveStreamStatusResponse
func (client *Client) DescribeAppLiveStreamStatusWithContext(ctx context.Context, request *DescribeAppLiveStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppLiveStreamStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppLiveStreamStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppLiveStreamStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用录制开关
//
// @param request - DescribeAppRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordStatusResponse
func (client *Client) DescribeAppRecordStatusWithContext(ctx context.Context, request *DescribeAppRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppRecordStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppRecordStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用录制模版列表
//
// @param tmpReq - DescribeAppRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordTemplatesResponse
func (client *Client) DescribeAppRecordTemplatesWithContext(ctx context.Context, tmpReq *DescribeAppRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAppRecordTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppRecordTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppRecordTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询录制列表
//
// @param tmpReq - DescribeAppRecordingFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordingFilesResponse
func (client *Client) DescribeAppRecordingFilesWithContext(ctx context.Context, tmpReq *DescribeAppRecordingFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordingFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAppRecordingFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppRecordingFiles"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppRecordingFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用推流模版列表
//
// @param tmpReq - DescribeAppStreamingOutTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppStreamingOutTemplatesResponse
func (client *Client) DescribeAppStreamingOutTemplatesWithContext(ctx context.Context, tmpReq *DescribeAppStreamingOutTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppStreamingOutTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAppStreamingOutTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ConditionShrink) {
		query["Condition"] = request.ConditionShrink
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppStreamingOutTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppStreamingOutTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看合流开关
//
// @param request - DescribeAppViewStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppViewStatusResponse
func (client *Client) DescribeAppViewStatusWithContext(ctx context.Context, request *DescribeAppViewStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppViewStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppViewStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppViewStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用合流模版列表
//
// @param tmpReq - DescribeAppViewTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppViewTemplatesResponse
func (client *Client) DescribeAppViewTemplatesWithContext(ctx context.Context, tmpReq *DescribeAppViewTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppViewTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAppViewTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ConditionShrink) {
		query["Condition"] = request.ConditionShrink
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppViewTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppViewTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # App列表
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithContext(ctx context.Context, request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoLiveStreamRuleResponse
func (client *Client) DescribeAutoLiveStreamRuleWithContext(ctx context.Context, request *DescribeAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeCall获取单次通信详情。
//
// @param request - DescribeCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallResponse
func (client *Client) DescribeCallWithContext(ctx context.Context, request *DescribeCallRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.ExtDataType) {
		query["ExtDataType"] = request.ExtDataType
	}

	if !dara.IsNil(request.QueryExpInfo) {
		query["QueryExpInfo"] = request.QueryExpInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCall"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeCallList分页查询时间范围内创建的通信信息。
//
// @param request - DescribeCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallListResponse
func (client *Client) DescribeCallListWithContext(ctx context.Context, request *DescribeCallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallStatus) {
		query["CallStatus"] = request.CallStatus
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryMode) {
		query["QueryMode"] = request.QueryMode
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCallList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCallListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// app事件回调列表
//
// @param request - DescribeCallbacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallbacksResponse
func (client *Client) DescribeCallbacksWithContext(ctx context.Context, request *DescribeCallbacksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallbacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCallbacks"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCallbacksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeChannel
//
// @param request - DescribeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelResponse
func (client *Client) DescribeChannelWithContext(ctx context.Context, request *DescribeChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询频道的所有参会者
//
// @param request - DescribeChannelAllUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAllUsersResponse
func (client *Client) DescribeChannelAllUsersWithContext(ctx context.Context, request *DescribeChannelAllUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAllUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelAllUsers"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelAllUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelAreaDistributionStatData获取频道地区分布统计数据。
//
// @param request - DescribeChannelAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAreaDistributionStatDataResponse
func (client *Client) DescribeChannelAreaDistributionStatDataWithContext(ctx context.Context, request *DescribeChannelAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.ParentArea) {
		query["ParentArea"] = request.ParentArea
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelAreaDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelDistributionStatData获取频道分布统计数据。
//
// @param request - DescribeChannelDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelDistributionStatDataResponse
func (client *Client) DescribeChannelDistributionStatDataWithContext(ctx context.Context, request *DescribeChannelDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelOverallData查询频道概览数据。
//
// @param request - DescribeChannelOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelOverallDataResponse
func (client *Client) DescribeChannelOverallDataWithContext(ctx context.Context, request *DescribeChannelOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeChannelParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelParticipantsResponse
func (client *Client) DescribeChannelParticipantsWithContext(ctx context.Context, request *DescribeChannelParticipantsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelParticipantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelParticipants"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelTopPubUserList获取频道内发布端的用户列表（按用户在线时长降序）。
//
// @param request - DescribeChannelTopPubUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelTopPubUserListResponse
func (client *Client) DescribeChannelTopPubUserListWithContext(ctx context.Context, request *DescribeChannelTopPubUserListRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelTopPubUserList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeChannelUser
//
// @param request - DescribeChannelUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserResponse
func (client *Client) DescribeChannelUserWithContext(ctx context.Context, request *DescribeChannelUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUser"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelUserMetrics查询频道总览中的用户数据。
//
// @param request - DescribeChannelUserMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserMetricsResponse
func (client *Client) DescribeChannelUserMetricsWithContext(ctx context.Context, request *DescribeChannelUserMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUserMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUserMetrics"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeChannelUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUsersResponse
func (client *Client) DescribeChannelUsersWithContext(ctx context.Context, request *DescribeChannelUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUsers"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询在线频道列表
//
// @param request - DescribeChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelsResponse
func (client *Client) DescribeChannelsWithContext(ctx context.Context, request *DescribeChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannels"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 纪要热词列表
//
// @param tmpReq - DescribeCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudNotePhrasesResponse
func (client *Client) DescribeCloudNotePhrasesWithContext(ctx context.Context, tmpReq *DescribeCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ConditionShrink) {
		query["Condition"] = request.ConditionShrink
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudNotePhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 纪要列表
//
// @param tmpReq - DescribeCloudNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudNotesResponse
func (client *Client) DescribeCloudNotesWithContext(ctx context.Context, tmpReq *DescribeCloudNotesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCloudNotesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudNotes"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudNotesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询录制任务状态
//
// @param request - DescribeCloudRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudRecordStatusResponse
func (client *Client) DescribeCloudRecordStatusWithContext(ctx context.Context, request *DescribeCloudRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudRecordStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudRecordStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointEventList获取端对端用户事件列表。
//
// @param request - DescribeEndPointEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointEventListResponse
func (client *Client) DescribeEndPointEventListWithContext(ctx context.Context, request *DescribeEndPointEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndPointEventListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndPointEventList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointMetricData获取端对端指标数据。
//
// @param request - DescribeEndPointMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointMetricDataResponse
func (client *Client) DescribeEndPointMetricDataWithContext(ctx context.Context, request *DescribeEndPointMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndPointMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.PubCallIdList) {
		query["PubCallIdList"] = request.PubCallIdList
	}

	if !dara.IsNil(request.PubUserId) {
		query["PubUserId"] = request.PubUserId
	}

	if !dara.IsNil(request.SubUserId) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndPointMetricData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断影响因素分布
//
// @param request - DescribeFaultDiagnosisFactorDistributionStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisFactorDistributionStatResponse
func (client *Client) DescribeFaultDiagnosisFactorDistributionStatWithContext(ctx context.Context, request *DescribeFaultDiagnosisFactorDistributionStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisFactorDistributionStat"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断总览数据
//
// @param request - DescribeFaultDiagnosisOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisOverallDataResponse
func (client *Client) DescribeFaultDiagnosisOverallDataWithContext(ctx context.Context, request *DescribeFaultDiagnosisOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户详情
//
// @param request - DescribeFaultDiagnosisUserDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserDetailResponse
func (client *Client) DescribeFaultDiagnosisUserDetailWithContext(ctx context.Context, request *DescribeFaultDiagnosisUserDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.FaultType) {
		query["FaultType"] = request.FaultType
	}

	if !dara.IsNil(request.QueryCallUserInfo) {
		query["QueryCallUserInfo"] = request.QueryCallUserInfo
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisUserDetail"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户明细列表
//
// @param request - DescribeFaultDiagnosisUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserListResponse
func (client *Client) DescribeFaultDiagnosisUserListWithContext(ctx context.Context, request *DescribeFaultDiagnosisUserListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.FaultTypes) {
		query["FaultTypes"] = request.FaultTypes
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisUserList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMPULayoutInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMPULayoutInfoListResponse
func (client *Client) DescribeMPULayoutInfoListWithContext(ctx context.Context, request *DescribeMPULayoutInfoListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMPULayoutInfoList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMPULayoutInfoListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribePubUserListBySubUser根据订阅端获取通信中发布端用户列表。
//
// @param request - DescribePubUserListBySubUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePubUserListBySubUserResponse
func (client *Client) DescribePubUserListBySubUserWithContext(ctx context.Context, request *DescribePubUserListBySubUserRequest, runtime *dara.RuntimeOptions) (_result *DescribePubUserListBySubUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.SubUserId) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePubUserListBySubUser"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeQoeMetricData获取单次通信中用户的下行体验质量指标。
//
// @param request - DescribeQoeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQoeMetricDataResponse
func (client *Client) DescribeQoeMetricDataWithContext(ctx context.Context, request *DescribeQoeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQoeMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQoeMetricData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计区域分布统计数据
//
// @param request - DescribeQualityAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityAreaDistributionStatDataResponse
func (client *Client) DescribeQualityAreaDistributionStatDataWithContext(ctx context.Context, request *DescribeQualityAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ParentArea) {
		query["ParentArea"] = request.ParentArea
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityAreaDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计分布数据
//
// @param request - DescribeQualityDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityDistributionStatDataResponse
func (client *Client) DescribeQualityDistributionStatDataWithContext(ctx context.Context, request *DescribeQualityDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeQualityOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeQualityOsSdkVersionDistributionStatDataWithContext(ctx context.Context, request *DescribeQualityOsSdkVersionDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityOsSdkVersionDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计概览数据
//
// @param request - DescribeQualityOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOverallDataResponse
func (client *Client) DescribeQualityOverallDataWithContext(ctx context.Context, request *DescribeQualityOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordFilesResponse
func (client *Client) DescribeRecordFilesWithContext(ctx context.Context, request *DescribeRecordFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordFiles"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordTemplatesResponse
func (client *Client) DescribeRecordTemplatesWithContext(ctx context.Context, request *DescribeRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcChannelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelListResponse
func (client *Client) DescribeRtcChannelListWithContext(ctx context.Context, request *DescribeRtcChannelListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcChannelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.TimePoint) {
		query["TimePoint"] = request.TimePoint
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcChannelList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcChannelListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcChannelMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelMetricResponse
func (client *Client) DescribeRtcChannelMetricWithContext(ctx context.Context, request *DescribeRtcChannelMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcChannelMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TimePoint) {
		query["TimePoint"] = request.TimePoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcChannelMetric"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcChannelMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcDurationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcDurationDataResponse
func (client *Client) DescribeRtcDurationDataWithContext(ctx context.Context, request *DescribeRtcDurationDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcDurationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcDurationData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcDurationDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcPeakChannelCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcPeakChannelCntDataResponse
func (client *Client) DescribeRtcPeakChannelCntDataWithContext(ctx context.Context, request *DescribeRtcPeakChannelCntDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcPeakChannelCntData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcPeakChannelCntDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcUserCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcUserCntDataResponse
func (client *Client) DescribeRtcUserCntDataWithContext(ctx context.Context, request *DescribeRtcUserCntDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcUserCntDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcUserCntData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcUserCntDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询旁路推流状态
//
// @param request - DescribeStreamingOutStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamingOutStatusResponse
func (client *Client) DescribeStreamingOutStatusWithContext(ctx context.Context, request *DescribeStreamingOutStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamingOutStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamingOutStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamingOutStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 系统内置布局
//
// @param request - DescribeSystemLayoutListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemLayoutListResponse
func (client *Client) DescribeSystemLayoutListWithContext(ctx context.Context, request *DescribeSystemLayoutListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemLayoutListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemLayoutList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemLayoutListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计地域分布数据
//
// @param request - DescribeUsageAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageAreaDistributionStatDataResponse
func (client *Client) DescribeUsageAreaDistributionStatDataWithContext(ctx context.Context, request *DescribeUsageAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ParentArea) {
		query["ParentArea"] = request.ParentArea
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageAreaDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计分布数据
//
// @param request - DescribeUsageDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageDistributionStatDataResponse
func (client *Client) DescribeUsageDistributionStatDataWithContext(ctx context.Context, request *DescribeUsageDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeUsageOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeUsageOsSdkVersionDistributionStatDataWithContext(ctx context.Context, request *DescribeUsageOsSdkVersionDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageOsSdkVersionDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计概览数据
//
// @param request - DescribeUsageOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOverallDataResponse
func (client *Client) DescribeUsageOverallDataWithContext(ctx context.Context, request *DescribeUsageOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeUserInfoInChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserInfoInChannelResponse
func (client *Client) DescribeUserInfoInChannelWithContext(ctx context.Context, request *DescribeUserInfoInChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserInfoInChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserInfoInChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserInfoInChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoLiveStreamRuleResponse
func (client *Client) DisableAutoLiveStreamRuleWithContext(ctx context.Context, request *DisableAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoLiveStreamRuleResponse
func (client *Client) EnableAutoLiveStreamRuleWithContext(ctx context.Context, request *EnableAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetAgent。
//
// @param request - GetAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithContext(ctx context.Context, request *GetAgentRequest, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMPUTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMPUTaskStatusResponse
func (client *Client) GetMPUTaskStatusWithContext(ctx context.Context, request *GetMPUTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMPUTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMPUTaskStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMPUTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改App信息
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApp"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用智能体开关
//
// @param request - ModifyAppAgentFunctionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppAgentFunctionStatusResponse
func (client *Client) ModifyAppAgentFunctionStatusWithContext(ctx context.Context, request *ModifyAppAgentFunctionStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppAgentFunctionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppAgentFunctionStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppAgentFunctionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用智能体模版
//
// @param tmpReq - ModifyAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppAgentTemplateResponse
func (client *Client) ModifyAppAgentTemplateWithContext(ctx context.Context, tmpReq *ModifyAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppAgentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppAgentTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentSilenceConfig) {
		request.AgentSilenceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentSilenceConfig, dara.String("AgentSilenceConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AmbientSoundConfig) {
		request.AmbientSoundConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AmbientSoundConfig, dara.String("AmbientSoundConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AsrConfig) {
		request.AsrConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrConfig, dara.String("AsrConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BackChannelConfig) {
		request.BackChannelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackChannelConfig, dara.String("BackChannelConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InterruptConfig) {
		request.InterruptConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterruptConfig, dara.String("InterruptConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LlmConfig) {
		request.LlmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmConfig, dara.String("LlmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSilenceConfigShrink) {
		query["AgentSilenceConfig"] = request.AgentSilenceConfigShrink
	}

	if !dara.IsNil(request.AmbientSoundConfigShrink) {
		query["AmbientSoundConfig"] = request.AmbientSoundConfigShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AsrConfigShrink) {
		query["AsrConfig"] = request.AsrConfigShrink
	}

	if !dara.IsNil(request.BackChannelConfigShrink) {
		query["BackChannelConfig"] = request.BackChannelConfigShrink
	}

	if !dara.IsNil(request.ChatMode) {
		query["ChatMode"] = request.ChatMode
	}

	if !dara.IsNil(request.Greeting) {
		query["Greeting"] = request.Greeting
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InterruptConfigShrink) {
		query["InterruptConfig"] = request.InterruptConfigShrink
	}

	if !dara.IsNil(request.InterruptMode) {
		query["InterruptMode"] = request.InterruptMode
	}

	if !dara.IsNil(request.LlmConfigShrink) {
		query["LlmConfig"] = request.LlmConfigShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TtsConfigShrink) {
		query["TtsConfig"] = request.TtsConfigShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppAgentTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppAgentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新app回调事件开关
//
// @param request - ModifyAppCallbackStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppCallbackStatusResponse
func (client *Client) ModifyAppCallbackStatusWithContext(ctx context.Context, request *ModifyAppCallbackStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppCallbackStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppCallbackStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppCallbackStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改app自定义布局
//
// @param tmpReq - ModifyAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppLayoutResponse
func (client *Client) ModifyAppLayoutWithContext(ctx context.Context, tmpReq *ModifyAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Layout) {
		request.LayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Layout, dara.String("Layout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.LayoutShrink) {
		query["Layout"] = request.LayoutShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用旁路开关
//
// @param request - ModifyAppLiveStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppLiveStreamStatusResponse
func (client *Client) ModifyAppLiveStreamStatusWithContext(ctx context.Context, request *ModifyAppLiveStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppLiveStreamStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppLiveStreamStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppLiveStreamStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用录制开关
//
// @param request - ModifyAppRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppRecordStatusResponse
func (client *Client) ModifyAppRecordStatusWithContext(ctx context.Context, request *ModifyAppRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppRecordStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppRecordStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用录制模版
//
// @param tmpReq - ModifyAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppRecordTemplateResponse
func (client *Client) ModifyAppRecordTemplateWithContext(ctx context.Context, tmpReq *ModifyAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordTemplate) {
		request.RecordTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordTemplate, dara.String("RecordTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordTemplateShrink) {
		query["RecordTemplate"] = request.RecordTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppRecordTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用推流模版
//
// @param tmpReq - ModifyAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppStreamingOutTemplateResponse
func (client *Client) ModifyAppStreamingOutTemplateWithContext(ctx context.Context, tmpReq *ModifyAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppStreamingOutTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamingOutTemplate) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, dara.String("StreamingOutTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StreamingOutTemplateShrink) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppStreamingOutTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改合流开关
//
// @param request - ModifyAppViewStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppViewStatusResponse
func (client *Client) ModifyAppViewStatusWithContext(ctx context.Context, request *ModifyAppViewStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppViewStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppViewStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppViewStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用合流模版
//
// @param tmpReq - ModifyAppViewTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppViewTemplateResponse
func (client *Client) ModifyAppViewTemplateWithContext(ctx context.Context, tmpReq *ModifyAppViewTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppViewTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppViewTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppViewTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppViewTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新app回调
//
// @param tmpReq - ModifyCallbackMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCallbackMetaResponse
func (client *Client) ModifyCallbackMetaWithContext(ctx context.Context, tmpReq *ModifyCallbackMetaRequest, runtime *dara.RuntimeOptions) (_result *ModifyCallbackMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyCallbackMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Callback) {
		request.CallbackShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Callback, dara.String("Callback"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackShrink) {
		query["Callback"] = request.CallbackShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCallbackMeta"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCallbackMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新纪要热词表
//
// @param tmpReq - ModifyCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudNotePhrasesResponse
func (client *Client) ModifyCloudNotePhrasesWithContext(ctx context.Context, tmpReq *ModifyCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Phrase) {
		request.PhraseShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Phrase, dara.String("Phrase"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PhraseShrink) {
		query["Phrase"] = request.PhraseShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudNotePhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMPULayoutResponse
func (client *Client) ModifyMPULayoutWithContext(ctx context.Context, request *ModifyMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyMPULayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AudioMixCount) {
		query["AudioMixCount"] = request.AudioMixCount
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMPULayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMPULayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置流属性
//
// @param tmpReq - ModifyStreamingPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStreamingPropertyResponse
func (client *Client) ModifyStreamingPropertyWithContext(ctx context.Context, tmpReq *ModifyStreamingPropertyRequest, runtime *dara.RuntimeOptions) (_result *ModifyStreamingPropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyStreamingPropertyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ViewSubscribers) {
		request.ViewSubscribersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewSubscribers, dara.String("ViewSubscribers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.ViewContent) {
		query["ViewContent"] = request.ViewContent
	}

	if !dara.IsNil(request.ViewSubscribersShrink) {
		query["ViewSubscribers"] = request.ViewSubscribersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStreamingProperty"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStreamingPropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置合流布局
//
// @param tmpReq - ModifyViewLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyViewLayoutResponse
func (client *Client) ModifyViewLayoutWithContext(ctx context.Context, tmpReq *ModifyViewLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyViewLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyViewLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyViewLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyViewLayoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NotifyAgent
//
// @param tmpReq - NotifyAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyAgentResponse
func (client *Client) NotifyAgentWithContext(ctx context.Context, tmpReq *NotifyAgentRequest, runtime *dara.RuntimeOptions) (_result *NotifyAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &NotifyAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BackgroundMusic) {
		request.BackgroundMusicShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackgroundMusic, dara.String("BackgroundMusic"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundMusicShrink) {
		query["BackgroundMusic"] = request.BackgroundMusicShrink
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CustomAttribute) {
		query["CustomAttribute"] = request.CustomAttribute
	}

	if !dara.IsNil(request.Interruptable) {
		query["Interruptable"] = request.Interruptable
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NotifyAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NotifyAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTerminalsResponse
func (client *Client) RemoveTerminalsWithContext(ctx context.Context, request *RemoveTerminalsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTerminalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TerminalIds) {
		query["TerminalIds"] = request.TerminalIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTerminals"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RemoveUsers
//
// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithContext(ctx context.Context, request *RemoveUsersRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsers"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消Sip邀请
//
// @param request - RtcCancelSipInviteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcCancelSipInviteResponse
func (client *Client) RtcCancelSipInviteWithContext(ctx context.Context, request *RtcCancelSipInviteRequest, runtime *dara.RuntimeOptions) (_result *RtcCancelSipInviteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcCancelSipInvite"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcCancelSipInviteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 邀请SIP加入频道
//
// @param request - RtcSipInviteMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcSipInviteMemberResponse
func (client *Client) RtcSipInviteMemberWithContext(ctx context.Context, request *RtcSipInviteMemberRequest, runtime *dara.RuntimeOptions) (_result *RtcSipInviteMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppToken) {
		query["AppToken"] = request.AppToken
	}

	if !dara.IsNil(request.CallNumber) {
		query["CallNumber"] = request.CallNumber
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Registered) {
		query["Registered"] = request.Registered
	}

	if !dara.IsNil(request.ServerAddress) {
		query["ServerAddress"] = request.ServerAddress
	}

	if !dara.IsNil(request.SipDisplayName) {
		query["SipDisplayName"] = request.SipDisplayName
	}

	if !dara.IsNil(request.SipRoomId) {
		query["SipRoomId"] = request.SipRoomId
	}

	if !dara.IsNil(request.SipUri) {
		query["SipUri"] = request.SipUri
	}

	if !dara.IsNil(request.SipUserAgent) {
		query["SipUserAgent"] = request.SipUserAgent
	}

	if !dara.IsNil(request.SipUserId) {
		query["SipUserId"] = request.SipUserId
	}

	if !dara.IsNil(request.SipUserPassword) {
		query["SipUserPassword"] = request.SipUserPassword
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcSipInviteMember"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcSipInviteMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Mute操作
//
// @param request - RtcSipMuteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcSipMuteResponse
func (client *Client) RtcSipMuteWithContext(ctx context.Context, request *RtcSipMuteRequest, runtime *dara.RuntimeOptions) (_result *RtcSipMuteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Operations) {
		query["Operations"] = request.Operations
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcSipMute"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcSipMuteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动AI Agent
//
// @param tmpReq - StartAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAgentResponse
func (client *Client) StartAgentWithContext(ctx context.Context, tmpReq *StartAgentRequest, runtime *dara.RuntimeOptions) (_result *StartAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RtcConfig) {
		request.RtcConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RtcConfig, dara.String("RtcConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VoiceChatConfig) {
		request.VoiceChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VoiceChatConfig, dara.String("VoiceChatConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.RtcConfigShrink) {
		query["RtcConfig"] = request.RtcConfigShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.VoiceChatConfigShrink) {
		query["VoiceChatConfig"] = request.VoiceChatConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启某个事件回调
//
// @param tmpReq - StartCategoryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCategoryCallbackResponse
func (client *Client) StartCategoryCallbackWithContext(ctx context.Context, tmpReq *StartCategoryCallbackRequest, runtime *dara.RuntimeOptions) (_result *StartCategoryCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartCategoryCallbackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Callback) {
		request.CallbackShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Callback, dara.String("Callback"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackShrink) {
		query["Callback"] = request.CallbackShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCategoryCallback"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCategoryCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启智能纪要
//
// @param tmpReq - StartCloudNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudNoteResponse
func (client *Client) StartCloudNoteWithContext(ctx context.Context, tmpReq *StartCloudNoteRequest, runtime *dara.RuntimeOptions) (_result *StartCloudNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartCloudNoteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AutoChapters) {
		request.AutoChaptersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoChapters, dara.String("AutoChapters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomPrompt) {
		request.CustomPromptShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomPrompt, dara.String("CustomPrompt"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MeetingAssistance) {
		request.MeetingAssistanceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MeetingAssistance, dara.String("MeetingAssistance"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RealtimeSubtitle) {
		request.RealtimeSubtitleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeSubtitle, dara.String("RealtimeSubtitle"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServiceInspection) {
		request.ServiceInspectionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceInspection, dara.String("ServiceInspection"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Summarization) {
		request.SummarizationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Summarization, dara.String("Summarization"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TextPolish) {
		request.TextPolishShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextPolish, dara.String("TextPolish"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Transcription) {
		request.TranscriptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Transcription, dara.String("Transcription"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoChaptersShrink) {
		query["AutoChapters"] = request.AutoChaptersShrink
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CustomPromptShrink) {
		query["CustomPrompt"] = request.CustomPromptShrink
	}

	if !dara.IsNil(request.LanguageHints) {
		query["LanguageHints"] = request.LanguageHints
	}

	if !dara.IsNil(request.MeetingAssistanceShrink) {
		query["MeetingAssistance"] = request.MeetingAssistanceShrink
	}

	if !dara.IsNil(request.RealtimeSubtitleShrink) {
		query["RealtimeSubtitle"] = request.RealtimeSubtitleShrink
	}

	if !dara.IsNil(request.ServiceInspectionShrink) {
		query["ServiceInspection"] = request.ServiceInspectionShrink
	}

	if !dara.IsNil(request.SourceLanguage) {
		query["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.StorageConfig) {
		query["StorageConfig"] = request.StorageConfig
	}

	if !dara.IsNil(request.SummarizationShrink) {
		query["Summarization"] = request.SummarizationShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TextPolishShrink) {
		query["TextPolish"] = request.TextPolishShrink
	}

	if !dara.IsNil(request.TranscriptionShrink) {
		query["Transcription"] = request.TranscriptionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCloudNote"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCloudNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StartCloudRecord
//
// @param tmpReq - StartCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecordWithContext(ctx context.Context, tmpReq *StartCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartCloudRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SingleStreamingRecord) {
		request.SingleStreamingRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SingleStreamingRecord, dara.String("SingleStreamingRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotation) {
		query["Annotation"] = request.Annotation
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.RecordMode) {
		query["RecordMode"] = request.RecordMode
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.ReservePaneForNoCameraUser) {
		query["ReservePaneForNoCameraUser"] = request.ReservePaneForNoCameraUser
	}

	if !dara.IsNil(request.ShowDefaultBackgroundOnMute) {
		query["ShowDefaultBackgroundOnMute"] = request.ShowDefaultBackgroundOnMute
	}

	if !dara.IsNil(request.SingleStreamingRecordShrink) {
		query["SingleStreamingRecord"] = request.SingleStreamingRecordShrink
	}

	if !dara.IsNil(request.StartWithoutChannel) {
		query["StartWithoutChannel"] = request.StartWithoutChannel
	}

	if !dara.IsNil(request.StartWithoutChannelWaitTime) {
		query["StartWithoutChannelWaitTime"] = request.StartWithoutChannelWaitTime
	}

	if !dara.IsNil(request.StorageConfig) {
		query["StorageConfig"] = request.StorageConfig
	}

	if !dara.IsNil(request.SubHighResolutionStream) {
		query["SubHighResolutionStream"] = request.SubHighResolutionStream
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCloudRecord"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMPUTaskResponse
func (client *Client) StartMPUTaskWithContext(ctx context.Context, request *StartMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StartMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayloadType) {
		query["PayloadType"] = request.PayloadType
	}

	if !dara.IsNil(request.ReportVad) {
		query["ReportVad"] = request.ReportVad
	}

	if !dara.IsNil(request.RtpExtInfo) {
		query["RtpExtInfo"] = request.RtpExtInfo
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TimeStampRef) {
		query["TimeStampRef"] = request.TimeStampRef
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	if !dara.IsNil(request.VadInterval) {
		query["VadInterval"] = request.VadInterval
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EnhancedParam) {
		bodyFlat["EnhancedParam"] = request.EnhancedParam
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMPUTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMPUTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordTaskResponse
func (client *Client) StartRecordTaskWithContext(ctx context.Context, request *StartRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *StartRecordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRecordTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRecordTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StartStreamingOut
//
// @param tmpReq - StartStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamingOutResponse
func (client *Client) StartStreamingOutWithContext(ctx context.Context, tmpReq *StartStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *StartStreamingOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartStreamingOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotation) {
		query["Annotation"] = request.Annotation
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.ReservePaneForNoCameraUser) {
		query["ReservePaneForNoCameraUser"] = request.ReservePaneForNoCameraUser
	}

	if !dara.IsNil(request.ShowDefaultBackgroundOnMute) {
		query["ShowDefaultBackgroundOnMute"] = request.ShowDefaultBackgroundOnMute
	}

	if !dara.IsNil(request.SpecMixedUserList) {
		query["SpecMixedUserList"] = request.SpecMixedUserList
	}

	if !dara.IsNil(request.StartWithoutChannel) {
		query["StartWithoutChannel"] = request.StartWithoutChannel
	}

	if !dara.IsNil(request.StartWithoutChannelWaitTime) {
		query["StartWithoutChannelWaitTime"] = request.StartWithoutChannelWaitTime
	}

	if !dara.IsNil(request.SubHighResolutionStream) {
		query["SubHighResolutionStream"] = request.SubHighResolutionStream
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartStreamingOut"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartStreamingOutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始合流
//
// @param tmpReq - StartViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartViewResponse
func (client *Client) StartViewWithContext(ctx context.Context, tmpReq *StartViewRequest, runtime *dara.RuntimeOptions) (_result *StartViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartViewShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ViewSubscribers) {
		request.ViewSubscribersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewSubscribers, dara.String("ViewSubscribers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.StartWithoutChannel) {
		query["StartWithoutChannel"] = request.StartWithoutChannel
	}

	if !dara.IsNil(request.StartWithoutChannelWaitTime) {
		query["StartWithoutChannelWaitTime"] = request.StartWithoutChannelWaitTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.ViewContent) {
		query["ViewContent"] = request.ViewContent
	}

	if !dara.IsNil(request.ViewSubscribersShrink) {
		query["ViewSubscribers"] = request.ViewSubscribersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartView"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止AI Agent
//
// @param request - StopAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAgentResponse
func (client *Client) StopAgentWithContext(ctx context.Context, request *StopAgentRequest, runtime *dara.RuntimeOptions) (_result *StopAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭某个事件回调
//
// @param tmpReq - StopCategoryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCategoryCallbackResponse
func (client *Client) StopCategoryCallbackWithContext(ctx context.Context, tmpReq *StopCategoryCallbackRequest, runtime *dara.RuntimeOptions) (_result *StopCategoryCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopCategoryCallbackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Callback) {
		request.CallbackShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Callback, dara.String("Callback"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackShrink) {
		query["Callback"] = request.CallbackShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCategoryCallback"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCategoryCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除频道
//
// @param request - StopChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopChannelResponse
func (client *Client) StopChannelWithContext(ctx context.Context, request *StopChannelRequest, runtime *dara.RuntimeOptions) (_result *StopChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止智能纪要
//
// @param request - StopCloudNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudNoteResponse
func (client *Client) StopCloudNoteWithContext(ctx context.Context, request *StopCloudNoteRequest, runtime *dara.RuntimeOptions) (_result *StopCloudNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCloudNote"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCloudNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StopCloudRecord
//
// @param request - StopCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecordWithContext(ctx context.Context, request *StopCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCloudRecord"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMPUTaskResponse
func (client *Client) StopMPUTaskWithContext(ctx context.Context, request *StopMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StopMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMPUTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMPUTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordTaskResponse
func (client *Client) StopRecordTaskWithContext(ctx context.Context, request *StopRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *StopRecordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRecordTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRecordTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StopStreamingOut
//
// @param request - StopStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamingOutResponse
func (client *Client) StopStreamingOutWithContext(ctx context.Context, request *StopStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *StopStreamingOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopStreamingOut"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopStreamingOutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止合流
//
// @param request - StopViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopViewResponse
func (client *Client) StopViewWithContext(ctx context.Context, request *StopViewRequest, runtime *dara.RuntimeOptions) (_result *StopViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopView"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AI Agent
//
// @param tmpReq - UpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithContext(ctx context.Context, tmpReq *UpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VoiceChatConfig) {
		request.VoiceChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VoiceChatConfig, dara.String("VoiceChatConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.VoiceChatConfigShrink) {
		query["VoiceChatConfig"] = request.VoiceChatConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoLiveStreamRuleResponse
func (client *Client) UpdateAutoLiveStreamRuleWithContext(ctx context.Context, request *UpdateAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.ChannelIdPrefixes) {
		query["ChannelIdPrefixes"] = request.ChannelIdPrefixes
	}

	if !dara.IsNil(request.ChannelIds) {
		query["ChannelIds"] = request.ChannelIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新云端录制任务
//
// @param tmpReq - UpdateCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudRecordResponse
func (client *Client) UpdateCloudRecordWithContext(ctx context.Context, tmpReq *UpdateCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudRecord"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMPUTaskResponse
func (client *Client) UpdateMPUTaskWithContext(ctx context.Context, request *UpdateMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMPUTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMPUTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTaskResponse
func (client *Client) UpdateRecordTaskWithContext(ctx context.Context, request *UpdateRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecordTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTemplateResponse
func (client *Client) UpdateRecordTemplateWithContext(ctx context.Context, request *UpdateRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.DelayStopTime) {
		query["DelayStopTime"] = request.DelayStopTime
	}

	if !dara.IsNil(request.EnableM3u8DateTime) {
		query["EnableM3u8DateTime"] = request.EnableM3u8DateTime
	}

	if !dara.IsNil(request.FileSplitInterval) {
		query["FileSplitInterval"] = request.FileSplitInterval
	}

	if !dara.IsNil(request.Formats) {
		query["Formats"] = request.Formats
	}

	if !dara.IsNil(request.HttpCallbackUrl) {
		query["HttpCallbackUrl"] = request.HttpCallbackUrl
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MnsQueue) {
		query["MnsQueue"] = request.MnsQueue
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssFilePrefix) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新旁路推流任务
//
// @param tmpReq - UpdateStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStreamingOutResponse
func (client *Client) UpdateStreamingOutWithContext(ctx context.Context, tmpReq *UpdateStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *UpdateStreamingOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStreamingOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.SpecMixedUserList) {
		query["SpecMixedUserList"] = request.SpecMixedUserList
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStreamingOut"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStreamingOutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
