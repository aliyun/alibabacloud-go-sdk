// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建克隆音色
//
// @param request - CreateCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloneVoiceResponse
func (client *Client) CreateCloneVoiceWithContext(ctx context.Context, request *CreateCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *CreateCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloneVoiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// @param tmpReq - CreateFlashSmsAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlashSmsAccessProfileResponse
func (client *Client) CreateFlashSmsAccessProfileWithContext(ctx context.Context, tmpReq *CreateFlashSmsAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateFlashSmsAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFlashSmsAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessProfile) {
		request.AccessProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessProfile, dara.String("AccessProfile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileShrink) {
		body["AccessProfile"] = request.AccessProfileShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlashSmsAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlashSmsAccessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ServiceMode) {
		body["ServiceMode"] = request.ServiceMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// @param tmpReq - CreateOutboundCallRestrictionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOutboundCallRestrictionResponse
func (client *Client) CreateOutboundCallRestrictionWithContext(ctx context.Context, tmpReq *CreateOutboundCallRestrictionRequest, runtime *dara.RuntimeOptions) (_result *CreateOutboundCallRestrictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOutboundCallRestrictionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OutboundCallRestriction) {
		request.OutboundCallRestrictionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundCallRestriction, dara.String("OutboundCallRestriction"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundCallRestrictionShrink) {
		body["OutboundCallRestriction"] = request.OutboundCallRestrictionShrink
	}

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOutboundCallRestriction"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOutboundCallRestrictionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建场景
//
// @param request - CreateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptResponse
func (client *Client) CreateScriptWithContext(ctx context.Context, request *CreateScriptRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NluEngine) {
		body["NluEngine"] = request.NluEngine
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建场景配置
//
// @param tmpReq - CreateScriptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptVersionResponse
func (client *Client) CreateScriptVersionWithContext(ctx context.Context, tmpReq *CreateScriptVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScriptVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InteractionConfig) {
		request.InteractionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InteractionConfig, dara.String("InteractionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LabelConfigs) {
		request.LabelConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelConfigs, dara.String("LabelConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScriptProfile) {
		request.ScriptProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScriptProfile, dara.String("ScriptProfile"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SynthesizerConfig) {
		request.SynthesizerConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SynthesizerConfig, dara.String("SynthesizerConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TranscriberConfig) {
		request.TranscriberConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TranscriberConfig, dara.String("TranscriberConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InteractionConfigShrink) {
		body["InteractionConfig"] = request.InteractionConfigShrink
	}

	if !dara.IsNil(request.LabelConfigsShrink) {
		body["LabelConfigs"] = request.LabelConfigsShrink
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptProfileShrink) {
		body["ScriptProfile"] = request.ScriptProfileShrink
	}

	if !dara.IsNil(request.SourceVersionId) {
		body["SourceVersionId"] = request.SourceVersionId
	}

	if !dara.IsNil(request.SynthesizerConfigShrink) {
		body["SynthesizerConfig"] = request.SynthesizerConfigShrink
	}

	if !dara.IsNil(request.TranscriberConfigShrink) {
		body["TranscriberConfig"] = request.TranscriberConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScriptVersion"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建语音接入配置
//
// @param tmpReq - CreateVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVoiceAccessProfileResponse
func (client *Client) CreateVoiceAccessProfileWithContext(ctx context.Context, tmpReq *CreateVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVoiceAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Profile) {
		request.ProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Profile, dara.String("Profile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsEngine) {
		body["NlsEngine"] = request.NlsEngine
	}

	if !dara.IsNil(request.ProfileShrink) {
		body["Profile"] = request.ProfileShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVoiceAccessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除克隆音色
//
// @param request - DeleteCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloneVoiceResponse
func (client *Client) DeleteCloneVoiceWithContext(ctx context.Context, request *DeleteCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloneVoiceId) {
		body["CloneVoiceId"] = request.CloneVoiceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloneVoiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除闪信配置
//
// @param request - DeleteFlashSmsAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlashSmsAccessProfileResponse
func (client *Client) DeleteFlashSmsAccessProfileWithContext(ctx context.Context, request *DeleteFlashSmsAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlashSmsAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlashSmsAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlashSmsAccessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an instance.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes blacklists and whitelists.
//
// @param tmpReq - DeleteOutboundCallRestrictionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOutboundCallRestrictionResponse
func (client *Client) DeleteOutboundCallRestrictionWithContext(ctx context.Context, tmpReq *DeleteOutboundCallRestrictionRequest, runtime *dara.RuntimeOptions) (_result *DeleteOutboundCallRestrictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteOutboundCallRestrictionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RestrictionIdList) {
		request.RestrictionIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RestrictionIdList, dara.String("RestrictionIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RestrictionIdListShrink) {
		body["RestrictionIdList"] = request.RestrictionIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOutboundCallRestriction"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOutboundCallRestrictionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScriptResponse
func (client *Client) DeleteScriptWithContext(ctx context.Context, request *DeleteScriptRequest, runtime *dara.RuntimeOptions) (_result *DeleteScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除三方语音配置
//
// @param request - DeleteVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVoiceAccessProfileResponse
func (client *Client) DeleteVoiceAccessProfileWithContext(ctx context.Context, request *DeleteVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVoiceAccessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用消息订阅
//
// @param request - DisableSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSubscriptionResponse
func (client *Client) DisableSubscriptionWithContext(ctx context.Context, request *DisableSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *DisableSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSubscription"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an instance.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取prompt场景模版
//
// @param request - GetScriptProfileTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScriptProfileTemplateResponse
func (client *Client) GetScriptProfileTemplateWithContext(ctx context.Context, request *GetScriptProfileTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetScriptProfileTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScriptProfileTemplate"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScriptProfileTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取MQ配置
//
// @param request - GetSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionResponse
func (client *Client) GetSubscriptionWithContext(ctx context.Context, request *GetSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *GetSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscription"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取克隆音色可用模型列表
//
// @param request - ListCloneVoiceModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloneVoiceModelsResponse
func (client *Client) ListCloneVoiceModelsWithContext(ctx context.Context, request *ListCloneVoiceModelsRequest, runtime *dara.RuntimeOptions) (_result *ListCloneVoiceModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloneVoiceModels"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloneVoiceModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取克隆音色列表
//
// @param request - ListCloneVoicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloneVoicesResponse
func (client *Client) ListCloneVoicesWithContext(ctx context.Context, request *ListCloneVoicesRequest, runtime *dara.RuntimeOptions) (_result *ListCloneVoicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloneVoices"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloneVoicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取闪信配置列表
//
// @param request - ListFlashSmsAccessProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsAccessProfilesResponse
func (client *Client) ListFlashSmsAccessProfilesWithContext(ctx context.Context, request *ListFlashSmsAccessProfilesRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsAccessProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsAccessProfiles"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsAccessProfilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取闪信厂商列表
//
// @param request - ListFlashSmsProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsProvidersResponse
func (client *Client) ListFlashSmsProvidersWithContext(ctx context.Context, request *ListFlashSmsProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsProviders"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取闪信模版列表
//
// @param request - ListFlashSmsTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsTemplatesResponse
func (client *Client) ListFlashSmsTemplatesWithContext(ctx context.Context, request *ListFlashSmsTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsTemplates"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves instance details.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the blacklists and whitelists of an outbound robot.
//
// @param request - ListOutboundCallRestrictionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOutboundCallRestrictionsResponse
func (client *Client) ListOutboundCallRestrictionsWithContext(ctx context.Context, request *ListOutboundCallRestrictionsRequest, runtime *dara.RuntimeOptions) (_result *ListOutboundCallRestrictionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOutboundCallRestrictions"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOutboundCallRestrictionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景配置模板列表
//
// @param request - ListScriptProfileTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptProfileTemplatesResponse
func (client *Client) ListScriptProfileTemplatesWithContext(ctx context.Context, request *ListScriptProfileTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListScriptProfileTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NluEngine) {
		body["NluEngine"] = request.NluEngine
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptProfileTemplates"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptProfileTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景列表
//
// @param tmpReq - ListScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptsResponse
func (client *Client) ListScriptsWithContext(ctx context.Context, tmpReq *ListScriptsRequest, runtime *dara.RuntimeOptions) (_result *ListScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListScriptsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScriptIds) {
		request.ScriptIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScriptIds, dara.String("ScriptIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PublishOnly) {
		body["PublishOnly"] = request.PublishOnly
	}

	if !dara.IsNil(request.ScriptIdsShrink) {
		body["ScriptIds"] = request.ScriptIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScripts"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据流程ID获取场景列表
//
// @param request - ListScriptsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptsByFlowResponse
func (client *Client) ListScriptsByFlowWithContext(ctx context.Context, request *ListScriptsByFlowRequest, runtime *dara.RuntimeOptions) (_result *ListScriptsByFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		body["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptsByFlow"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptsByFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统配置列表
//
// @param request - ListSystemConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSystemConfigsResponse
func (client *Client) ListSystemConfigsWithContext(ctx context.Context, request *ListSystemConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListSystemConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSystemConfigs"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSystemConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取语音接入配置列表
//
// @param request - ListVoiceAccessProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoiceAccessProfilesResponse
func (client *Client) ListVoiceAccessProfilesWithContext(ctx context.Context, request *ListVoiceAccessProfilesRequest, runtime *dara.RuntimeOptions) (_result *ListVoiceAccessProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVoiceAccessProfiles"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoiceAccessProfilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布场景
//
// @param request - PublishScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishScriptResponse
func (client *Client) PublishScriptWithContext(ctx context.Context, request *PublishScriptRequest, runtime *dara.RuntimeOptions) (_result *PublishScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新克隆音色
//
// @param request - UpdateCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloneVoiceResponse
func (client *Client) UpdateCloneVoiceWithContext(ctx context.Context, request *UpdateCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloneVoiceId) {
		body["CloneVoiceId"] = request.CloneVoiceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloneVoiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新闪信配置
//
// @param tmpReq - UpdateFlashSmsAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlashSmsAccessProfileResponse
func (client *Client) UpdateFlashSmsAccessProfileWithContext(ctx context.Context, tmpReq *UpdateFlashSmsAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlashSmsAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFlashSmsAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessProfile) {
		request.AccessProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessProfile, dara.String("AccessProfile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileShrink) {
		body["AccessProfile"] = request.AccessProfileShrink
	}

	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlashSmsAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlashSmsAccessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an instance.
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新场景
//
// @param request - UpdateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScriptResponse
func (client *Client) UpdateScriptWithContext(ctx context.Context, request *UpdateScriptRequest, runtime *dara.RuntimeOptions) (_result *UpdateScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新MQ配置
//
// @param tmpReq - UpdateSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithContext(ctx context.Context, tmpReq *UpdateSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSubscriptionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventSubscriptions) {
		request.EventSubscriptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventSubscriptions, dara.String("EventSubscriptions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Endpoint) {
		body["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EventSubscriptionsShrink) {
		body["EventSubscriptions"] = request.EventSubscriptionsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqInstanceId) {
		body["MqInstanceId"] = request.MqInstanceId
	}

	if !dara.IsNil(request.MqType) {
		body["MqType"] = request.MqType
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.ProducerId) {
		body["ProducerId"] = request.ProducerId
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubscription"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新系统配置
//
// @param tmpReq - UpdateSystemConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSystemConfigsResponse
func (client *Client) UpdateSystemConfigsWithContext(ctx context.Context, tmpReq *UpdateSystemConfigsRequest, runtime *dara.RuntimeOptions) (_result *UpdateSystemConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSystemConfigsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configs) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, dara.String("Configs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigsShrink) {
		body["Configs"] = request.ConfigsShrink
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSystemConfigs"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSystemConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新语音接入配置
//
// @param tmpReq - UpdateVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVoiceAccessProfileResponse
func (client *Client) UpdateVoiceAccessProfileWithContext(ctx context.Context, tmpReq *UpdateVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateVoiceAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Profile) {
		request.ProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Profile, dara.String("Profile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsEngine) {
		body["NlsEngine"] = request.NlsEngine
	}

	if !dara.IsNil(request.ProfileShrink) {
		body["Profile"] = request.ProfileShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVoiceAccessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
