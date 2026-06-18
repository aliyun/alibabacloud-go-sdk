// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Invoke the AddHotlineNumber API to add a hotline number.
//
// Description:
//
// > Hotline numbers are for inbound and outbound calls only.
//
// @param tmpReq - AddHotlineNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHotlineNumberResponse
func (client *Client) AddHotlineNumberWithContext(ctx context.Context, tmpReq *AddHotlineNumberRequest, runtime *dara.RuntimeOptions) (_result *AddHotlineNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddHotlineNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OutboundRangeList) {
		request.OutboundRangeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundRangeList, dara.String("OutboundRangeList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableInbound) {
		body["EnableInbound"] = request.EnableInbound
	}

	if !dara.IsNil(request.EnableInboundEvaluation) {
		body["EnableInboundEvaluation"] = request.EnableInboundEvaluation
	}

	if !dara.IsNil(request.EnableOutbound) {
		body["EnableOutbound"] = request.EnableOutbound
	}

	if !dara.IsNil(request.EnableOutboundEvaluation) {
		body["EnableOutboundEvaluation"] = request.EnableOutboundEvaluation
	}

	if !dara.IsNil(request.EvaluationLevel) {
		body["EvaluationLevel"] = request.EvaluationLevel
	}

	if !dara.IsNil(request.HotlineNumber) {
		body["HotlineNumber"] = request.HotlineNumber
	}

	if !dara.IsNil(request.InboundFlowId) {
		body["InboundFlowId"] = request.InboundFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundAllDepart) {
		body["OutboundAllDepart"] = request.OutboundAllDepart
	}

	if !dara.IsNil(request.OutboundRangeListShrink) {
		body["OutboundRangeList"] = request.OutboundRangeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddHotlineNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddHotlineNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more inbound numbers.
//
// @param tmpReq - AddInboundNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInboundNumberResponse
func (client *Client) AddInboundNumberWithContext(ctx context.Context, tmpReq *AddInboundNumberRequest, runtime *dara.RuntimeOptions) (_result *AddInboundNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddInboundNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InboundNumbers) {
		request.InboundNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InboundNumbers, dara.String("InboundNumbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.InboundNumbersShrink) {
		query["InboundNumbers"] = request.InboundNumbersShrink
	}

	if !dara.IsNil(request.InboundType) {
		query["InboundType"] = request.InboundType
	}

	if !dara.IsNil(request.LineCode) {
		query["LineCode"] = request.LineCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInboundNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInboundNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a large language model.
//
// @param tmpReq - AddLargeModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLargeModelResponse
func (client *Client) AddLargeModelWithContext(ctx context.Context, tmpReq *AddLargeModelRequest, runtime *dara.RuntimeOptions) (_result *AddLargeModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddLargeModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BaseModel) {
		request.BaseModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseModel, dara.String("BaseModel"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.BaseModelShrink) {
		query["BaseModel"] = request.BaseModelShrink
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelUrl) {
		query["ModelUrl"] = request.ModelUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Temperature) {
		query["Temperature"] = request.Temperature
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.TopP) {
		query["TopP"] = request.TopP
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLargeModel"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLargeModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add a model application
//
// @param tmpReq - AddModelApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddModelApplicationResponse
func (client *Client) AddModelApplicationWithContext(ctx context.Context, tmpReq *AddModelApplicationRequest, runtime *dara.RuntimeOptions) (_result *AddModelApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddModelApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCps) {
		query["ApplicationCps"] = request.ApplicationCps
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.CallConnectedTriggerModel) {
		query["CallConnectedTriggerModel"] = request.CallConnectedTriggerModel
	}

	if !dara.IsNil(request.DyvmsSceneName) {
		query["DyvmsSceneName"] = request.DyvmsSceneName
	}

	if !dara.IsNil(request.ModelCode) {
		query["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.MuteActive) {
		query["MuteActive"] = request.MuteActive
	}

	if !dara.IsNil(request.MuteDuration) {
		query["MuteDuration"] = request.MuteDuration
	}

	if !dara.IsNil(request.MuteHangupNum) {
		query["MuteHangupNum"] = request.MuteHangupNum
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.QualificationName) {
		query["QualificationName"] = request.QualificationName
	}

	if !dara.IsNil(request.RecordingFile) {
		query["RecordingFile"] = request.RecordingFile
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SpeechContent) {
		query["SpeechContent"] = request.SpeechContent
	}

	if !dara.IsNil(request.SpeechId) {
		query["SpeechId"] = request.SpeechId
	}

	if !dara.IsNil(request.StartWord) {
		query["StartWord"] = request.StartWord
	}

	if !dara.IsNil(request.StartWordType) {
		query["StartWordType"] = request.StartWordType
	}

	if !dara.IsNil(request.TtsConfigShrink) {
		query["TtsConfig"] = request.TtsConfigShrink
	}

	if !dara.IsNil(request.UsageDesc) {
		query["UsageDesc"] = request.UsageDesc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddModelApplication"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddModelApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke AddOuterAccount to add an external account.
//
// @param request - AddOuterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOuterAccountResponse
func (client *Client) AddOuterAccountWithContext(ctx context.Context, request *AddOuterAccountRequest, runtime *dara.RuntimeOptions) (_result *AddOuterAccountResponse, _err error) {
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
		Action:      dara.String("AddOuterAccount"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOuterAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke AddSkillGroup to create an external skill group.
//
// @param request - AddSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSkillGroupResponse
func (client *Client) AddSkillGroupWithContext(ctx context.Context, request *AddSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *AddSkillGroupResponse, _err error) {
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
		Action:      dara.String("AddSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the AiccsSmartCall API to initiate an Intelligent Speech Interaction call.
//
// @param request - AiccsSmartCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AiccsSmartCallResponse
func (client *Client) AiccsSmartCallWithContext(ctx context.Context, request *AiccsSmartCallRequest, runtime *dara.RuntimeOptions) (_result *AiccsSmartCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionCodeBreak) {
		query["ActionCodeBreak"] = request.ActionCodeBreak
	}

	if !dara.IsNil(request.ActionCodeTimeBreak) {
		query["ActionCodeTimeBreak"] = request.ActionCodeTimeBreak
	}

	if !dara.IsNil(request.AsrAlsAmId) {
		query["AsrAlsAmId"] = request.AsrAlsAmId
	}

	if !dara.IsNil(request.AsrBaseId) {
		query["AsrBaseId"] = request.AsrBaseId
	}

	if !dara.IsNil(request.AsrModelId) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.AsrVocabularyId) {
		query["AsrVocabularyId"] = request.AsrVocabularyId
	}

	if !dara.IsNil(request.BackgroundFileCode) {
		query["BackgroundFileCode"] = request.BackgroundFileCode
	}

	if !dara.IsNil(request.BackgroundSpeed) {
		query["BackgroundSpeed"] = request.BackgroundSpeed
	}

	if !dara.IsNil(request.BackgroundVolume) {
		query["BackgroundVolume"] = request.BackgroundVolume
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.DynamicId) {
		query["DynamicId"] = request.DynamicId
	}

	if !dara.IsNil(request.EarlyMediaAsr) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !dara.IsNil(request.EnableITN) {
		query["EnableITN"] = request.EnableITN
	}

	if !dara.IsNil(request.MuteTime) {
		query["MuteTime"] = request.MuteTime
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PauseTime) {
		query["PauseTime"] = request.PauseTime
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RecordFlag) {
		query["RecordFlag"] = request.RecordFlag
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.TtsConf) {
		query["TtsConf"] = request.TtsConf
	}

	if !dara.IsNil(request.TtsSpeed) {
		query["TtsSpeed"] = request.TtsSpeed
	}

	if !dara.IsNil(request.TtsStyle) {
		query["TtsStyle"] = request.TtsStyle
	}

	if !dara.IsNil(request.TtsVolume) {
		query["TtsVolume"] = request.TtsVolume
	}

	if !dara.IsNil(request.VoiceCode) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.VoiceCodeParam) {
		query["VoiceCodeParam"] = request.VoiceCodeParam
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AiccsSmartCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AiccsSmartCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the AiccsSmartCallOperate API to initiate a specified action during an Intelligent outbound call. This API is only used for scenarios such as parallel transfer to a human agent or allowing a human agent to listen in on the man-machine dialogue.
//
// @param request - AiccsSmartCallOperateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AiccsSmartCallOperateResponse
func (client *Client) AiccsSmartCallOperateWithContext(ctx context.Context, request *AiccsSmartCallOperateRequest, runtime *dara.RuntimeOptions) (_result *AiccsSmartCallOperateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AiccsSmartCallOperate"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AiccsSmartCallOperateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke AnswerCall to answer an incoming call.
//
// @param request - AnswerCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnswerCallResponse
func (client *Client) AnswerCallWithContext(ctx context.Context, request *AnswerCallRequest, runtime *dara.RuntimeOptions) (_result *AnswerCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnswerCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnswerCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Append job details.
//
// Description:
//
// - After creating an Intelligent Contact Robot calling job, you can invoke this API to append job details.
//
// - Before invoking this API, ensure that you already have a successfully created Intelligent Contact Robot calling job.
//
// - If you do not have a successfully created Intelligent Contact Robot calling job, you can click **Create Job*	- on the [Task Management](https://aiccs.console.aliyun.com/job/list) interface or create a job by using the [CreateTask](https://help.aliyun.com/document_detail/2718003.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - AttachTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachTaskResponse
func (client *Client) AttachTaskWithContext(ctx context.Context, request *AttachTaskRequest, runtime *dara.RuntimeOptions) (_result *AttachTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallString) {
		query["CallString"] = request.CallString
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchCreateQualityProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateQualityProjectsResponse
func (client *Client) BatchCreateQualityProjectsWithContext(ctx context.Context, request *BatchCreateQualityProjectsRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateQualityProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisIds) {
		query["AnalysisIds"] = request.AnalysisIds
	}

	if !dara.IsNil(request.ChannelTouchType) {
		query["ChannelTouchType"] = request.ChannelTouchType
	}

	if !dara.IsNil(request.CheckFreqType) {
		query["CheckFreqType"] = request.CheckFreqType
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TimeRangeEnd) {
		query["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !dara.IsNil(request.TimeRangeStart) {
		query["TimeRangeStart"] = request.TimeRangeStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateQualityProjects"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateQualityProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation cancels calls from a call task. You cannot cancel a call if its detail record is already in the pending call queue or is in progress.
//
// Description:
//
// - Use this operation to cancel calls. Alternatively, you can manually cancel calls in the console by navigating to **Call Task Management*	- > **Manage*	- > **Execution Records*	- > **Pending**.
//
// - Before calling this operation, ensure you have created a call task and imported callee data.
//
// - If you have not created a call task, you can create one and import callee data on the **Call Task Management*	- page. You can also call the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926815.html) and [ImportTaskNumberDatas]() operations.
//
// - Canceling a call task may affect your business. Please proceed with caution.
//
// @param tmpReq - CancelAiCallDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAiCallDetailsResponse
func (client *Client) CancelAiCallDetailsWithContext(ctx context.Context, tmpReq *CancelAiCallDetailsRequest, runtime *dara.RuntimeOptions) (_result *CancelAiCallDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CancelAiCallDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DetailIdList) {
		request.DetailIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetailIdList, dara.String("DetailIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PhoneNumbers) {
		request.PhoneNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumbers, dara.String("PhoneNumbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		query["BatchId"] = request.BatchId
	}

	if !dara.IsNil(request.DetailIdListShrink) {
		query["DetailIdList"] = request.DetailIdListShrink
	}

	if !dara.IsNil(request.EncryptionType) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumbersShrink) {
		query["PhoneNumbers"] = request.PhoneNumbersShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAiCallDetails"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAiCallDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancel an Intelligent Contact Robot calling job.
//
// Description:
//
// - You can invoke this API to cancel an Intelligent Contact Robot calling job, or manually cancel the job in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface.
//
// - After an Intelligent Contact Robot calling job is canceled, it cannot be started again. Proceed with caution.
//
// - If you want to pause a job and restart it later, you can manually pause the job in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface or pause it by using the [StopTask](https://help.aliyun.com/document_detail/2718006.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - CancelTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke ChangeChatAgentStatus to modify the Live Support status.
//
// Description:
//
// > Currently, only changing the Live Support status to offline is supported.
//
// @param request - ChangeChatAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeChatAgentStatusResponse
func (client *Client) ChangeChatAgentStatusWithContext(ctx context.Context, request *ChangeChatAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *ChangeChatAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Method) {
		body["Method"] = request.Method
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeChatAgentStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeQualityProjectStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeQualityProjectStatusResponse
func (client *Client) ChangeQualityProjectStatusWithContext(ctx context.Context, request *ChangeQualityProjectStatusRequest, runtime *dara.RuntimeOptions) (_result *ChangeQualityProjectStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeQualityProjectStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeQualityProjectStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an agent account in the Cloud Customer Service System and return the agent ID.
//
// Description:
//
// - Before calling this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify skill group information, refer to the guidance in [Request Parameters](#api-detail-35).
//
// - You can manage agents by calling [DeleteAgent](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-deleteagent) to delete an agent or [UpdateAgent](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-updateagent) to update agent data.
//
// ### QPS Limit
//
// - Per-user call frequency: No rate limiting.
//
// - API call frequency: 100 queries per second (QPS).
//
// > If the total calls from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - CreateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithContext(ctx context.Context, request *CreateAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.SkillGroupId) {
		bodyFlat["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		bodyFlat["SkillGroupIdList"] = request.SkillGroupIdList
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgent"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a call task.
//
// Description:
//
// Before creating a call task, make a test call with an agent to ensure the results meet your requirements.
//
// @param tmpReq - CreateAiCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAiCallTaskResponse
func (client *Client) CreateAiCallTaskWithContext(ctx context.Context, tmpReq *CreateAiCallTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAiCallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAiCallTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallDay) {
		request.CallDayShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallDay, dara.String("CallDay"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallRetryReason) {
		request.CallRetryReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallRetryReason, dara.String("CallRetryReason"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallTime) {
		request.CallTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTime, dara.String("CallTime"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.CallDayShrink) {
		query["CallDay"] = request.CallDayShrink
	}

	if !dara.IsNil(request.CallRetryInterval) {
		query["CallRetryInterval"] = request.CallRetryInterval
	}

	if !dara.IsNil(request.CallRetryReasonShrink) {
		query["CallRetryReason"] = request.CallRetryReasonShrink
	}

	if !dara.IsNil(request.CallRetryTimes) {
		query["CallRetryTimes"] = request.CallRetryTimes
	}

	if !dara.IsNil(request.CallTimeShrink) {
		query["CallTime"] = request.CallTimeShrink
	}

	if !dara.IsNil(request.LineEncoding) {
		query["LineEncoding"] = request.LineEncoding
	}

	if !dara.IsNil(request.LinePhoneNum) {
		query["LinePhoneNum"] = request.LinePhoneNum
	}

	if !dara.IsNil(request.MissCallRetry) {
		query["MissCallRetry"] = request.MissCallRetry
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneType) {
		query["PhoneType"] = request.PhoneType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.StartType) {
		query["StartType"] = request.StartType
	}

	if !dara.IsNil(request.TaskCps) {
		query["TaskCps"] = request.TaskCps
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStartTime) {
		query["TaskStartTime"] = request.TaskStartTime
	}

	if !dara.IsNil(request.VirtualNumber) {
		query["VirtualNumber"] = request.VirtualNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAiCallTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAiCallTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an Intelligent Outbound Call Job. You can configure the Task Type, job name, outbound caller ID, callee number deduplication policy, and other settings when creating the job.
//
// Description:
//
// - The **Data*	- field in the response parameters of this API is the job ID.
//
// - After creating an Intelligent Outbound Call Job, if you need to make updates, you can invoke the [UpdateAiOutboundTask](https://help.aliyun.com/document_detail/2718021.html) API to update the outbound call job.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 20 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param tmpReq - CreateAiOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAiOutboundTaskResponse
func (client *Client) CreateAiOutboundTaskWithContext(ctx context.Context, tmpReq *CreateAiOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAiOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAiOutboundTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OutboundNums) {
		request.OutboundNumsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundNums, dara.String("OutboundNums"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecallRule) {
		request.RecallRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecallRule, dara.String("RecallRule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrentRate) {
		query["ConcurrentRate"] = request.ConcurrentRate
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionTime) {
		query["ExecutionTime"] = request.ExecutionTime
	}

	if !dara.IsNil(request.ForecastCallRate) {
		query["ForecastCallRate"] = request.ForecastCallRate
	}

	if !dara.IsNil(request.HandlerId) {
		query["HandlerId"] = request.HandlerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NumRepeated) {
		query["NumRepeated"] = request.NumRepeated
	}

	if !dara.IsNil(request.OutboundNumsShrink) {
		query["OutboundNums"] = request.OutboundNumsShrink
	}

	if !dara.IsNil(request.RecallRuleShrink) {
		query["RecallRule"] = request.RecallRuleShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAiOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAiOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a batch for an Intelligent Outbound Calling job based on the instance ID and job ID, enabling data under the job to be queried by batch.
//
// Description:
//
// Before invoking this API, we recommend that you confirm the instance ID and job ID. For guidance on how to obtain them, see the instructions in [Request Parameters](#api-detail-35).
//
// ### QPS limit
//
// - Per-user invocation frequency: 20 queries per second (QPS).
//
// - API-wide invocation frequency: 20 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param request - CreateAiOutboundTaskBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAiOutboundTaskBatchResponse
func (client *Client) CreateAiOutboundTaskBatchWithContext(ctx context.Context, request *CreateAiOutboundTaskBatchRequest, runtime *dara.RuntimeOptions) (_result *CreateAiOutboundTaskBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAiOutboundTaskBatch"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAiOutboundTaskBatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates department information based on the Artificial Intelligence Cloud Call Service (AICCS) instance ID and department name. Upon successful creation, the department ID is returned.
//
// Description:
//
// - If you need to update department information, you can invoke the [UpdateDepartment](https://help.aliyun.com/document_detail/2717977.html) API.
//
// - After successfully creating department information by invoking this API, the **Data*	- field in the response contains the department ID. If you need to query the department ID later, you can invoke the [GetAllDepartment](https://help.aliyun.com/document_detail/2717975.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API invocation frequency: 100 times/second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param request - CreateDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDepartmentResponse
func (client *Client) CreateDepartmentWithContext(ctx context.Context, request *CreateDepartmentRequest, runtime *dara.RuntimeOptions) (_result *CreateDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentName) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDepartment"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOutboundTaskResponse
func (client *Client) CreateOutboundTaskWithContext(ctx context.Context, request *CreateOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ani) {
		query["Ani"] = request.Ani
	}

	if !dara.IsNil(request.CallInfos) {
		query["CallInfos"] = request.CallInfos
	}

	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExtAttrs) {
		query["ExtAttrs"] = request.ExtAttrs
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.RetryInterval) {
		query["RetryInterval"] = request.RetryInterval
	}

	if !dara.IsNil(request.RetryTime) {
		query["RetryTime"] = request.RetryTime
	}

	if !dara.IsNil(request.SkillGroup) {
		query["SkillGroup"] = request.SkillGroup
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateQualityProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityProjectResponse
func (client *Client) CreateQualityProjectWithContext(ctx context.Context, request *CreateQualityProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisIds) {
		body["AnalysisIds"] = request.AnalysisIds
	}

	if !dara.IsNil(request.ChannelTouchType) {
		body["ChannelTouchType"] = request.ChannelTouchType
	}

	if !dara.IsNil(request.CheckFreqType) {
		body["CheckFreqType"] = request.CheckFreqType
	}

	if !dara.IsNil(request.DepList) {
		body["DepList"] = request.DepList
	}

	if !dara.IsNil(request.GroupList) {
		body["GroupList"] = request.GroupList
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ScopeType) {
		body["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.ServicerList) {
		body["ServicerList"] = request.ServicerList
	}

	if !dara.IsNil(request.TimeRangeEnd) {
		body["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !dara.IsNil(request.TimeRangeStart) {
		body["TimeRangeStart"] = request.TimeRangeStart
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityProject"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityRuleResponse
func (client *Client) CreateQualityRuleWithContext(ctx context.Context, request *CreateQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityRuleResponse, _err error) {
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

	if !dara.IsNil(request.KeyWords) {
		body["KeyWords"] = request.KeyWords
	}

	if !dara.IsNil(request.MatchType) {
		body["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RuleTag) {
		body["RuleTag"] = request.RuleTag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityRule"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a skill group based on the skill group name and channel type.
//
// Description:
//
// - This API allows you to define information such as the external display name and description of the skill group. For details, see [Request Parameters](#api-detail-35).
//
// - The **Data*	- parameter returned by this API is the ID of the successfully created skill group.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 1000 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - CreateSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillGroupResponse
func (client *Client) CreateSkillGroupWithContext(ctx context.Context, request *CreateSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		body["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DepartmentId) {
		body["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupName) {
		body["SkillGroupName"] = request.SkillGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an Artificial Intelligence Cloud Call Service robot outbound calling job.
//
// Description:
//
// - You can invoke this API to create a job, or create one in the **Artificial Intelligence Cloud Call Service console*	- > **Outbound Robot (Standard Edition)*	- > [Task Management](https://aiccs.console.aliyun.com/job/list) by clicking **Create Job**.
//
// - After invoking this API, the **Data*	- field in the response contains the unique job ID of the robot outbound calling task.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - CreateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallString) {
		query["CallString"] = request.CallString
	}

	if !dara.IsNil(request.CallStringType) {
		query["CallStringType"] = request.CallStringType
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetryCount) {
		query["RetryCount"] = request.RetryCount
	}

	if !dara.IsNil(request.RetryFlag) {
		query["RetryFlag"] = request.RetryFlag
	}

	if !dara.IsNil(request.RetryInterval) {
		query["RetryInterval"] = request.RetryInterval
	}

	if !dara.IsNil(request.RetryStatusCode) {
		query["RetryStatusCode"] = request.RetryStatusCode
	}

	if !dara.IsNil(request.RobotId) {
		query["RobotId"] = request.RobotId
	}

	if !dara.IsNil(request.SeatCount) {
		query["SeatCount"] = request.SeatCount
	}

	if !dara.IsNil(request.StartNow) {
		query["StartNow"] = request.StartNow
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.WorkDay) {
		query["WorkDay"] = request.WorkDay
	}

	if !dara.IsNil(request.WorkTimeList) {
		query["WorkTimeList"] = request.WorkTimeList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an agent that enables password-free login to the Cloud Customer Service System based on a User Account.
//
// @param request - CreateThirdSsoAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateThirdSsoAgentResponse
func (client *Client) CreateThirdSsoAgentWithContext(ctx context.Context, request *CreateThirdSsoAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateThirdSsoAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.RoleIds) {
		bodyFlat["RoleIds"] = request.RoleIds
	}

	if !dara.IsNil(request.SkillGroupIds) {
		bodyFlat["SkillGroupIds"] = request.SkillGroupIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateThirdSsoAgent"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an agent account in the Cloud Customer Service System based on the instance ID and agent account name.
//
// Description:
//
// - Before deletion, we recommend that you confirm the agent account name and instance ID to be deleted. For guidance on how to obtain them, see the description of [Request Parameters](#api-detail-35).
//
// - If an agent is accidentally deleted, you can invoke the [CreateAgent](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-createagent) API to recreate the agent.
//
// > If an account is re-added after deletion, the agent ID remains unchanged.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param request - DeleteAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithContext(ctx context.Context, request *DeleteAgentRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Intelligent Outbound Calling job by instance ID and job ID. After deletion, the job will no longer appear in the outbound calling job list.
//
// Description:
//
//	Notice: Deletion is a sensitive operation. Proceed with caution.
//
// - Before deletion, we recommend that you confirm the job ID and related information. You can call the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API to view the outbound calling job list and verify the task name, description, and corresponding job ID.
//
// - If you need to recreate an Intelligent Outbound Calling job, you can call the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 20 times/second.
//
// - API invocation frequency: 20 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - DeleteAiOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAiOutboundTaskResponse
func (client *Client) DeleteAiOutboundTaskWithContext(ctx context.Context, request *DeleteAiOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteAiOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAiOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAiOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes department information based on the Artificial Intelligence Cloud Call Service (AICCS) instance ID and department ID.
//
// Description:
//
// - Deletion is a sensitive operation. Proceed with caution.
//
// - Before invoking this API, we recommend that you confirm the AICCS instance ID and department ID. For guidance on how to obtain them, see the instructions in [Request Parameters](#api-detail-35).
//
// - If you accidentally delete department information, you can call the [CreateDepartment](https://help.aliyun.com/document_detail/2717974.html) API to recreate it.
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: 100 times/second.
//
// - API invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - DeleteDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDepartmentResponse
func (client *Client) DeleteDepartmentWithContext(ctx context.Context, request *DeleteDepartmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDepartment"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the DeleteHotlineNumber API to delete a configured hotline number.
//
// @param request - DeleteHotlineNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotlineNumberResponse
func (client *Client) DeleteHotlineNumberWithContext(ctx context.Context, request *DeleteHotlineNumberRequest, runtime *dara.RuntimeOptions) (_result *DeleteHotlineNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HotlineNumber) {
		body["HotlineNumber"] = request.HotlineNumber
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHotlineNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHotlineNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOutboundTaskResponse
func (client *Client) DeleteOutboundTaskWithContext(ctx context.Context, request *DeleteOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundTaskId) {
		query["OutboundTaskId"] = request.OutboundTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke DeleteOuterAccount to delete an external account by its external Account ID.
//
// @param request - DeleteOuterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOuterAccountResponse
func (client *Client) DeleteOuterAccountWithContext(ctx context.Context, request *DeleteOuterAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteOuterAccountResponse, _err error) {
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
		Action:      dara.String("DeleteOuterAccount"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOuterAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteQualityProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityProjectResponse
func (client *Client) DeleteQualityProjectWithContext(ctx context.Context, request *DeleteQualityProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityProject"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityRuleResponse
func (client *Client) DeleteQualityRuleWithContext(ctx context.Context, request *DeleteQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityRule"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an external skill group based on the skill group ID and skill group channel type.
//
// @param request - DeleteSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillGroupResponse
func (client *Client) DeleteSkillGroupWithContext(ctx context.Context, request *DeleteSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillGroupResponse, _err error) {
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
		Action:      dara.String("DeleteSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the DescribeRecordData API to retrieve call information.
//
// @param request - DescribeRecordDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordDataResponse
func (client *Client) DescribeRecordDataWithContext(ctx context.Context, request *DescribeRecordDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.Acid) {
		query["Acid"] = request.Acid
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecLevel) {
		query["SecLevel"] = request.SecLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordData"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EditQualityProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditQualityProjectResponse
func (client *Client) EditQualityProjectWithContext(ctx context.Context, request *EditQualityProjectRequest, runtime *dara.RuntimeOptions) (_result *EditQualityProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisIds) {
		query["AnalysisIds"] = request.AnalysisIds
	}

	if !dara.IsNil(request.ChannelTouchType) {
		query["ChannelTouchType"] = request.ChannelTouchType
	}

	if !dara.IsNil(request.CheckFreqType) {
		query["CheckFreqType"] = request.CheckFreqType
	}

	if !dara.IsNil(request.DepList) {
		query["DepList"] = request.DepList
	}

	if !dara.IsNil(request.GroupList) {
		query["GroupList"] = request.GroupList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectVersion) {
		query["ProjectVersion"] = request.ProjectVersion
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.ServicerList) {
		query["ServicerList"] = request.ServicerList
	}

	if !dara.IsNil(request.TimeRangeEnd) {
		query["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !dara.IsNil(request.TimeRangeStart) {
		query["TimeRangeStart"] = request.TimeRangeStart
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditQualityProject"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditQualityProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EditQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditQualityRuleResponse
func (client *Client) EditQualityRuleWithContext(ctx context.Context, request *EditQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *EditQualityRuleResponse, _err error) {
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

	if !dara.IsNil(request.KeyWords) {
		body["KeyWords"] = request.KeyWords
	}

	if !dara.IsNil(request.MatchType) {
		body["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.QualityRuleId) {
		body["QualityRuleId"] = request.QualityRuleId
	}

	if !dara.IsNil(request.RuleTag) {
		body["RuleTag"] = request.RuleTag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditQualityRule"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EditQualityRuleTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditQualityRuleTagResponse
func (client *Client) EditQualityRuleTagWithContext(ctx context.Context, request *EditQualityRuleTagRequest, runtime *dara.RuntimeOptions) (_result *EditQualityRuleTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisTypes) {
		query["AnalysisTypes"] = request.AnalysisTypes
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditQualityRuleTag"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditQualityRuleTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke EncryptPhoneNum to encrypt the User\\"s Phone number.
//
// @param request - EncryptPhoneNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncryptPhoneNumResponse
func (client *Client) EncryptPhoneNumWithContext(ctx context.Context, request *EncryptPhoneNumRequest, runtime *dara.RuntimeOptions) (_result *EncryptPhoneNumResponse, _err error) {
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
		Action:      dara.String("EncryptPhoneNum"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EncryptPhoneNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve or recover a call.
//
// @param request - FetchCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchCallResponse
func (client *Client) FetchCallWithContext(ctx context.Context, request *FetchCallRequest, runtime *dara.RuntimeOptions) (_result *FetchCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.HoldConnectionId) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// End the hotline service by instance ID and agent account name.
//
// Description:
//
// ### Queries per second (QPS) limit
//
// - Per-user API call frequency: No rate limiting.
//
// - API frequency: 100 calls per second.
//
// > If the total number of calls from multiple users exceeds the API frequency limit, throttling will be triggered.
//
// @param request - FinishHotlineServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishHotlineServiceResponse
func (client *Client) FinishHotlineServiceWithContext(ctx context.Context, request *FinishHotlineServiceRequest, runtime *dara.RuntimeOptions) (_result *FinishHotlineServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishHotlineService"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate a message channel access signature based on the Artificial Intelligence Cloud Call Service (AICCS) instance ID and agent account name.
//
// Description:
//
// - You can invoke this API to generate a message channel access signature. The **Data*	- field in the response contains the MessageBox message channel signature code.
//
// - Before invoking the API, we recommend that you confirm your AICCS instance ID. For instructions on how to obtain it, see the description of [Request Parameters](#api-detail-35).
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API-wide frequency: 100 queries per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param request - GenerateWebSocketSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateWebSocketSignResponse
func (client *Client) GenerateWebSocketSignWithContext(ctx context.Context, request *GenerateWebSocketSignRequest, runtime *dara.RuntimeOptions) (_result *GenerateWebSocketSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateWebSocketSign"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query agent information in the Cloud Customer Service System by instance ID and agent account name, such as agent ID.
//
// Description:
//
// - Before invoking this API, confirm the AICCS instance information and the agent account name. For guidance on how to obtain these details, refer to the instructions in [Request Parameters](#api-detail-35).
//
// - Deleted agents can also be queried. Check the **Status*	- parameter in the response. If its value is 2, it indicates that the agent has been deleted.
//
// ### queries per second (QPS) Limit
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
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
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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

// Summary:
//
// Obtain the agent-level status metrics under hotline monitoring. Supports queries by instance, time range, agent, and department group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to provide agent or department information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetAgentBasisStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentBasisStatusResponse
func (client *Client) GetAgentBasisStatusWithContext(ctx context.Context, tmpReq *GetAgentBasisStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAgentBasisStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentBasisStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentBasisStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentBasisStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAgentByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentByIdResponse
func (client *Client) GetAgentByIdWithContext(ctx context.Context, request *GetAgentByIdRequest, runtime *dara.RuntimeOptions) (_result *GetAgentByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentById"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the summary data of hotline agent details under hotline reports. Supports queries by instance, time range, agent, and department group.
//
// Description:
//
// - Before you invoke this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to provide agent or department information, refer to the instructions in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetAgentDetailReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentDetailReportResponse
func (client *Client) GetAgentDetailReportWithContext(ctx context.Context, tmpReq *GetAgentDetailReportRequest, runtime *dara.RuntimeOptions) (_result *GetAgentDetailReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentDetailReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentDetailReport"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentDetailReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain real-time agent detail data grouped by instance, department, and skill group (real-time data).
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department or skill group information, refer to the instructions in the [Request Parameters](#api-detail-35) section.
//
// ### queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 10 requests per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetAgentIndexRealTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentIndexRealTimeResponse
func (client *Client) GetAgentIndexRealTimeWithContext(ctx context.Context, request *GetAgentIndexRealTimeRequest, runtime *dara.RuntimeOptions) (_result *GetAgentIndexRealTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DepIds) {
		query["DepIds"] = request.DepIds
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentIndexRealTime"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentIndexRealTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the agent service status metrics from online reports. Supports queries by instance, time range, agent, and department group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to provide agent or department information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 queries per second.
//
// - API-wide invocation frequency: 100 queries per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - GetAgentServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentServiceStatusResponse
func (client *Client) GetAgentServiceStatusWithContext(ctx context.Context, tmpReq *GetAgentServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAgentServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentServiceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentServiceStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain agent-level statistics under hotline monitoring. Supports querying by instance, time range, agent, and department group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify agent or department information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### queries per second (QPS) limit
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - GetAgentStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentStatisticsResponse
func (client *Client) GetAgentStatisticsWithContext(ctx context.Context, tmpReq *GetAgentStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetAgentStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentStatistics"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the business information associated with this call by instance ID and session ID.
//
// Description:
//
// ### Queries per second (QPS) limits
//
// - Call frequency per user: 100 times/second.
//
// - API call frequency: 100 times/second.
//
// > If the total calls from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetAiOutboundTaskBizDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiOutboundTaskBizDataResponse
func (client *Client) GetAiOutboundTaskBizDataWithContext(ctx context.Context, request *GetAiOutboundTaskBizDataRequest, runtime *dara.RuntimeOptions) (_result *GetAiOutboundTaskBizDataResponse, _err error) {
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
		Action:      dara.String("GetAiOutboundTaskBizData"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiOutboundTaskBizDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the outbound call job details by instance ID and job ID.
//
// Description:
//
// The outbound call job details include the job ID, job status, task type, outbound caller number, callee number repetition policy, and other information. For more information, see [Response parameters](#api-detail-40).
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetAiOutboundTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiOutboundTaskDetailResponse
func (client *Client) GetAiOutboundTaskDetailWithContext(ctx context.Context, request *GetAiOutboundTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *GetAiOutboundTaskDetailResponse, _err error) {
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
		Action:      dara.String("GetAiOutboundTaskDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiOutboundTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain task execution details by instance ID and job ID.
//
// Description:
//
// The task execution details include the total number of jobs, job batches, outbound call numbers, corresponding call counts, execution status, and other information.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetAiOutboundTaskExecDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiOutboundTaskExecDetailResponse
func (client *Client) GetAiOutboundTaskExecDetailWithContext(ctx context.Context, request *GetAiOutboundTaskExecDetailRequest, runtime *dara.RuntimeOptions) (_result *GetAiOutboundTaskExecDetailResponse, _err error) {
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
		Action:      dara.String("GetAiOutboundTaskExecDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiOutboundTaskExecDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the outbound call job list by instance ID and Task Type.
//
// Description:
//
// - The job list contains job information, including job ID, Task Status, Task Name, task completion rate, and more. For details, see [Response parameters](#api-detail-40).
//
// - If you need to update job information, you can invoke the [UpdateAiOutboundTask](https://help.aliyun.com/document_detail/2718021.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API invocation frequency: 100 times/second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param request - GetAiOutboundTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiOutboundTaskListResponse
func (client *Client) GetAiOutboundTaskListWithContext(ctx context.Context, request *GetAiOutboundTaskListRequest, runtime *dara.RuntimeOptions) (_result *GetAiOutboundTaskListResponse, _err error) {
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
		Action:      dara.String("GetAiOutboundTaskList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiOutboundTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the outbound call job progress by instance ID and job ID.
//
// Description:
//
// - Before invoking this API, we recommend that you confirm the instance ID and job ID. For more information, see [Request Parameters](#api-detail-35).
//
// - The outbound call job progress includes information such as job ID, Task Type, job completion rate, agent connection rate, and customer connection rate. For details, see [Response Parameters](#api-detail-40).
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API-wide invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param request - GetAiOutboundTaskProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiOutboundTaskProgressResponse
func (client *Client) GetAiOutboundTaskProgressWithContext(ctx context.Context, request *GetAiOutboundTaskProgressRequest, runtime *dara.RuntimeOptions) (_result *GetAiOutboundTaskProgressResponse, _err error) {
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
		Action:      dara.String("GetAiOutboundTaskProgress"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiOutboundTaskProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// getAllDepartment
//
// Description:
//
// - You can invoke this API to obtain department IDs for department group queries in certain data query APIs (such as [GetHotlineServiceStatistics](https://help.aliyun.com/document_detail/2717938.html)).
//
// - After creating, deleting, or updating department information, you can invoke this API to confirm whether the department information matches your expectations.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetAllDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllDepartmentResponse
func (client *Client) GetAllDepartmentWithContext(ctx context.Context, request *GetAllDepartmentRequest, runtime *dara.RuntimeOptions) (_result *GetAllDepartmentResponse, _err error) {
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
		Action:      dara.String("GetAllDepartment"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAllDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the upload address for a voice memo recording.
//
// @param request - GetAudioNoteUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioNoteUploadUrlResponse
func (client *Client) GetAudioNoteUploadUrlWithContext(ctx context.Context, request *GetAudioNoteUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAudioNoteUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAudioNoteUploadUrl"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAudioNoteUploadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the dialog content for a call by using its call ID. You can retrieve content for calls completed within the last 30 days.
//
// @param request - GetCallDialogContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallDialogContentResponse
func (client *Client) GetCallDialogContentWithContext(ctx context.Context, request *GetCallDialogContentRequest, runtime *dara.RuntimeOptions) (_result *GetCallDialogContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallDate) {
		query["CallDate"] = request.CallDate
	}

	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCallDialogContent"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCallDialogContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the URL of a call recording file based on the call creation time and the unique call ID, and then retrieve the call recording file using the URL.
//
// Description:
//
// ### Prerequisites
//
// Before invoking this API, ensure that call recording was enabled during the invocation of the [RobotCall](https://help.aliyun.com/document_detail/223270.html) API and that you successfully received the recording receipt, indicating that the recording file has been generated. Otherwise, an invalid URL will be returned.
//
// ### How-To
//
// This API serves as a supplementary method to the recording receipt. If the URL in the recording receipt message expires, you can use this API to obtain a new recording URL. By default, the validity period of the recording receipt URL is three days.
//
// > We recommend that you directly download the recording content using the recording receipt URL and save it locally, rather than relying on the receipt URL, to avoid issues caused by expiration.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > Throttling will be triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param request - GetCallSoundRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallSoundRecordResponse
func (client *Client) GetCallSoundRecordWithContext(ctx context.Context, request *GetCallSoundRecordRequest, runtime *dara.RuntimeOptions) (_result *GetCallSoundRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCallSoundRecord"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCallSoundRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetConfigNumList API to obtain the hotline settings number list.
//
// @param request - GetConfigNumListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigNumListResponse
func (client *Client) GetConfigNumListWithContext(ctx context.Context, request *GetConfigNumListRequest, runtime *dara.RuntimeOptions) (_result *GetConfigNumListResponse, _err error) {
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
		Action:      dara.String("GetConfigNumList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigNumListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can invoke the GetCustomerInfo API to obtain membership details by Workbench membership ID.
//
// @param request - GetCustomerInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerInfoResponse
func (client *Client) GetCustomerInfoWithContext(ctx context.Context, request *GetCustomerInfoRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerInfoResponse, _err error) {
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
		Action:      dara.String("GetCustomerInfo"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries skill group categories and the skill group list by AICCS instance ID and agent ID.
//
// Description:
//
// - The skill groups returned by this API are grouped by department ID, which can be specified when you [create a skill group](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-createskillgroup).
//
// - To query detailed skill group information, you can invoke the [QuerySkillGroups](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-queryskillgroups) API.
//
// - Before invoking this API, you should confirm the AICCS instance ID and agent ID. For guidance on how to obtain them, see the description of [Request Parameters](#api-detail-35).
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 1000 queries per second.
//
// - API-wide invocation frequency: 1000 queries per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param request - GetDepGroupTreeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDepGroupTreeDataResponse
func (client *Client) GetDepGroupTreeDataWithContext(ctx context.Context, request *GetDepGroupTreeDataRequest, runtime *dara.RuntimeOptions) (_result *GetDepGroupTreeDataResponse, _err error) {
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
		Action:      dara.String("GetDepGroupTreeData"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDepGroupTreeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain department-level status metrics under hotline monitoring. Supports queries by instance, time range, and department group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API-wide invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetDepartmentalLatitudeAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDepartmentalLatitudeAgentStatusResponse
func (client *Client) GetDepartmentalLatitudeAgentStatusWithContext(ctx context.Context, tmpReq *GetDepartmentalLatitudeAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDepartmentalLatitudeAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDepartmentalLatitudeAgentStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDepartmentalLatitudeAgentStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDepartmentalLatitudeAgentStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the current service details of a hotline agent based on the instance ID and agent account name.
//
// Description:
//
// The hotline agent details include information such as agent posture status, agent ID, and heartbeat signature. For more information, see [Response parameters](#api-detail-40).
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: Rate Limiting is not applied.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetHotlineAgentDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineAgentDetailResponse
func (client *Client) GetHotlineAgentDetailWithContext(ctx context.Context, request *GetHotlineAgentDetailRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineAgentDetailResponse, _err error) {
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
		Action:      dara.String("GetHotlineAgentDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain detailed data metrics by hotline agent dimension (T+1). Supports grouping queries by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/document_detail/276009.html) and [created an instance](https://help.aliyun.com/document_detail/276011.html).
//
// - If you need to specify department or skill group information, refer to the instructions in the [Request Parameters](#api-detail-35) section.
//
// > Query logic priority:
//
// > - If GroupIds is not empty, query data metrics under the specified skill groups.
//
// > - Otherwise, if DepIds is not empty, query data metrics under the corresponding department groups.
//
// > - Otherwise, query data metrics under the AICCS instance.
//
// ### QPS Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 10 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetHotlineAgentDetailReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineAgentDetailReportResponse
func (client *Client) GetHotlineAgentDetailReportWithContext(ctx context.Context, request *GetHotlineAgentDetailReportRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DepIds) {
		query["DepIds"] = request.DepIds
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotlineAgentDetailReport"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hotline agent status in the Cloud Customer Service System by AICCS instance ID and agent account name.
//
// Description:
//
// - Before invoking this API, confirm the AICCS instance information and the agent account name. For guidance on how to obtain these details, see the description in [Request Parameters](#api-detail-35).
//
// - To retrieve detailed information about a hotline agent, you can invoke the [GetHotlineAgentDetail](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-gethotlineagentdetail) API.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API-wide frequency: 200 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetHotlineAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineAgentStatusResponse
func (client *Client) GetHotlineAgentStatusWithContext(ctx context.Context, request *GetHotlineAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotlineAgentStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetHotlineCallAction API to query the result data of call actions.
//
// Description:
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 50 times/second.
//
// - API invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetHotlineCallActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineCallActionResponse
func (client *Client) GetHotlineCallActionWithContext(ctx context.Context, request *GetHotlineCallActionRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineCallActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Acc) {
		body["Acc"] = request.Acc
	}

	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.Act) {
		body["Act"] = request.Act
	}

	if !dara.IsNil(request.Biz) {
		body["Biz"] = request.Biz
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FromSource) {
		body["FromSource"] = request.FromSource
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotlineCallAction"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineCallActionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve detailed (T+1) data by hotline skill group dimension. Supports grouping by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/document_detail/276009.html) and [created an instance](https://help.aliyun.com/document_detail/276011.html).
//
// - If you need to specify agent, department, or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 10 queries per second.
//
// > Throttle will be triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param request - GetHotlineGroupDetailReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineGroupDetailReportResponse
func (client *Client) GetHotlineGroupDetailReportWithContext(ctx context.Context, request *GetHotlineGroupDetailReportRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DepIds) {
		query["DepIds"] = request.DepIds
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotlineGroupDetailReport"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetHotlineMessageLog API to retrieve hotline message records.
//
// @param request - GetHotlineMessageLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineMessageLogResponse
func (client *Client) GetHotlineMessageLogWithContext(ctx context.Context, request *GetHotlineMessageLogRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineMessageLogResponse, _err error) {
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
		Action:      dara.String("GetHotlineMessageLog"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineMessageLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetHotlineRuntimeInfo API to query hotline runtime information.
//
// @param request - GetHotlineRuntimeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineRuntimeInfoResponse
func (client *Client) GetHotlineRuntimeInfoWithContext(ctx context.Context, request *GetHotlineRuntimeInfoRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineRuntimeInfoResponse, _err error) {
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
		Action:      dara.String("GetHotlineRuntimeInfo"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineRuntimeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain service statistics under hotline monitoring. Supports grouped queries by instance, time dimension, agent, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to provide agent, department, or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - GetHotlineServiceStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineServiceStatisticsResponse
func (client *Client) GetHotlineServiceStatisticsWithContext(ctx context.Context, tmpReq *GetHotlineServiceStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineServiceStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetHotlineServiceStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotlineServiceStatistics"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineServiceStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the queue length of hotline members.
//
// @param request - GetHotlineWaitingNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineWaitingNumberResponse
func (client *Client) GetHotlineWaitingNumberWithContext(ctx context.Context, request *GetHotlineWaitingNumberRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineWaitingNumberResponse, _err error) {
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
		Action:      dara.String("GetHotlineWaitingNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can obtain the current statistical metrics (real-time data) and query them by instance ID, department, or skill group.
//
// Description:
//
// - The current data statistics metrics include cumulative metrics for the day and real-time metrics.
//
// - If you need to specify department or skill group information, refer to the instructions in the [Request Parameters](#api-detail-35) section.
//
// > Query logic priority:
//
// > - If GroupIds is not empty, query the data metrics for the specified skill groups.
//
// > - Otherwise, if DepIds is not empty, query the data metrics for the specified departments.
//
// > - Otherwise, query the data metrics for the Artificial Intelligence Cloud Call Service (AICCS) instance.
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetIndexCurrentValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexCurrentValueResponse
func (client *Client) GetIndexCurrentValueWithContext(ctx context.Context, request *GetIndexCurrentValueRequest, runtime *dara.RuntimeOptions) (_result *GetIndexCurrentValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepIds) {
		query["DepIds"] = request.DepIds
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIndexCurrentValue"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIndexCurrentValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceListResponse
func (client *Client) GetInstanceListWithContext(ctx context.Context, request *GetInstanceListRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceListResponse, _err error) {
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
		Action:      dara.String("GetInstanceList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetMcuLvsIp API to query the hotline server IP address.
//
// @param request - GetMcuLvsIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcuLvsIpResponse
func (client *Client) GetMcuLvsIpWithContext(ctx context.Context, request *GetMcuLvsIpRequest, runtime *dara.RuntimeOptions) (_result *GetMcuLvsIpResponse, _err error) {
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
		Action:      dara.String("GetMcuLvsIp"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcuLvsIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the number\\"s归属地 information based on the instance ID and phone number.
//
// Description:
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - GetNumLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNumLocationResponse
func (client *Client) GetNumLocationWithContext(ctx context.Context, request *GetNumLocationRequest, runtime *dara.RuntimeOptions) (_result *GetNumLocationResponse, _err error) {
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
		Action:      dara.String("GetNumLocation"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNumLocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain online agent information data under real-time monitoring. Support queries by instance, time range, agent, and department group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify agent or department information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - GetOnlineSeatInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOnlineSeatInformationResponse
func (client *Client) GetOnlineSeatInformationWithContext(ctx context.Context, tmpReq *GetOnlineSeatInformationRequest, runtime *dara.RuntimeOptions) (_result *GetOnlineSeatInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOnlineSeatInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOnlineSeatInformation"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOnlineSeatInformationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the volume of Online Services under online monitoring. Support grouped queries by instance, time range, agent, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify agent, department, or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetOnlineServiceVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOnlineServiceVolumeResponse
func (client *Client) GetOnlineServiceVolumeWithContext(ctx context.Context, tmpReq *GetOnlineServiceVolumeRequest, runtime *dara.RuntimeOptions) (_result *GetOnlineServiceVolumeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOnlineServiceVolumeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOnlineServiceVolume"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOnlineServiceVolumeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of external hotline numbers.
//
// @param request - GetOutbounNumListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOutbounNumListResponse
func (client *Client) GetOutbounNumListWithContext(ctx context.Context, request *GetOutbounNumListRequest, runtime *dara.RuntimeOptions) (_result *GetOutbounNumListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOutbounNumList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke GetQualityProjectDetail to obtain the quality inspection job details.
//
// @param request - GetQualityProjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityProjectDetailResponse
func (client *Client) GetQualityProjectDetailWithContext(ctx context.Context, request *GetQualityProjectDetailRequest, runtime *dara.RuntimeOptions) (_result *GetQualityProjectDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityProjectDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityProjectDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke GetQualityProjectList to obtain the quality inspection job list.
//
// @param request - GetQualityProjectListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityProjectListResponse
func (client *Client) GetQualityProjectListWithContext(ctx context.Context, request *GetQualityProjectListRequest, runtime *dara.RuntimeOptions) (_result *GetQualityProjectListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.CheckFreqType) {
		query["checkFreqType"] = request.CheckFreqType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityProjectList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityProjectListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualityProjectLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityProjectLogResponse
func (client *Client) GetQualityProjectLogWithContext(ctx context.Context, request *GetQualityProjectLogRequest, runtime *dara.RuntimeOptions) (_result *GetQualityProjectLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityProjectLog"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityProjectLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualityResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityResultResponse
func (client *Client) GetQualityResultWithContext(ctx context.Context, request *GetQualityResultRequest, runtime *dara.RuntimeOptions) (_result *GetQualityResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.HitStatus) {
		query["HitStatus"] = request.HitStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.QualityRuleIds) {
		query["QualityRuleIds"] = request.QualityRuleIds
	}

	if !dara.IsNil(request.TouchEndTime) {
		query["TouchEndTime"] = request.TouchEndTime
	}

	if !dara.IsNil(request.TouchStartTime) {
		query["TouchStartTime"] = request.TouchStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityResult"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualityRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleDetailResponse
func (client *Client) GetQualityRuleDetailWithContext(ctx context.Context, request *GetQualityRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QualityRuleId) {
		query["QualityRuleId"] = request.QualityRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualityRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleListResponse
func (client *Client) GetQualityRuleListWithContext(ctx context.Context, request *GetQualityRuleListRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualityRuleTagListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleTagListResponse
func (client *Client) GetQualityRuleTagListWithContext(ctx context.Context, request *GetQualityRuleTagListRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleTagListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleTagList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleTagListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain queue information under the skill group dimension in online monitoring. Supports queries grouped by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetQueueInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueInformationResponse
func (client *Client) GetQueueInformationWithContext(ctx context.Context, tmpReq *GetQueueInformationRequest, runtime *dara.RuntimeOptions) (_result *GetQueueInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetQueueInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueInformation"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueInformationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetRecordData API to obtain a recording file.
//
// @param request - GetRecordDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecordDataResponse
func (client *Client) GetRecordDataWithContext(ctx context.Context, request *GetRecordDataRequest, runtime *dara.RuntimeOptions) (_result *GetRecordDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Acid) {
		query["Acid"] = request.Acid
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecordData"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecordDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the GetRecordUrl API to obtain the incoming and outgoing calls recording link.
//
// @param request - GetRecordUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecordUrlResponse
func (client *Client) GetRecordUrlWithContext(ctx context.Context, request *GetRecordUrlRequest, runtime *dara.RuntimeOptions) (_result *GetRecordUrlResponse, _err error) {
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
		Action:      dara.String("GetRecordUrl"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecordUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke GetRtcToken to obtain the token for a shift agent.
//
// @param request - GetRtcTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRtcTokenResponse
func (client *Client) GetRtcTokenWithContext(ctx context.Context, request *GetRtcTokenRequest, runtime *dara.RuntimeOptions) (_result *GetRtcTokenResponse, _err error) {
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
		Action:      dara.String("GetRtcToken"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRtcTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain agent service capacity data under online monitoring. Supports queries by instance, time range, and department group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetSeatInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSeatInformationResponse
func (client *Client) GetSeatInformationWithContext(ctx context.Context, tmpReq *GetSeatInformationRequest, runtime *dara.RuntimeOptions) (_result *GetSeatInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSeatInformationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("depIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSeatInformation"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSeatInformationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the service status metrics of skill groups. Support grouped queries by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param tmpReq - GetSkillGroupAgentStatusDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupAgentStatusDetailsResponse
func (client *Client) GetSkillGroupAgentStatusDetailsWithContext(ctx context.Context, tmpReq *GetSkillGroupAgentStatusDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupAgentStatusDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSkillGroupAgentStatusDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupAgentStatusDetails"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupAgentStatusDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the aggregated agent status metrics for skill groups under hotline monitoring. Supports grouped queries by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - GetSkillGroupAndAgentStatusSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupAndAgentStatusSummaryResponse
func (client *Client) GetSkillGroupAndAgentStatusSummaryWithContext(ctx context.Context, tmpReq *GetSkillGroupAndAgentStatusSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupAndAgentStatusSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSkillGroupAndAgentStatusSummaryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupAndAgentStatusSummary"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupAndAgentStatusSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the hotline monitoring status metrics by skill group dimension. Supports queries grouped by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param tmpReq - GetSkillGroupLatitudeStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupLatitudeStateResponse
func (client *Client) GetSkillGroupLatitudeStateWithContext(ctx context.Context, tmpReq *GetSkillGroupLatitudeStateRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupLatitudeStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSkillGroupLatitudeStateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupLatitudeState"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupLatitudeStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve skill group–level service capacity data under online monitoring. Supports queries grouped by instance, time range, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify department or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - GetSkillGroupServiceCapabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupServiceCapabilityResponse
func (client *Client) GetSkillGroupServiceCapabilityWithContext(ctx context.Context, tmpReq *GetSkillGroupServiceCapabilityRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupServiceCapabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSkillGroupServiceCapabilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupServiceCapability"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupServiceCapabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the service status metrics of skill groups in online reports.
//
// Description:
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param tmpReq - GetSkillGroupServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupServiceStatusResponse
func (client *Client) GetSkillGroupServiceStatusWithContext(ctx context.Context, tmpReq *GetSkillGroupServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSkillGroupServiceStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupServiceStatus"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the total status count of skill groups under hotline reports. Supports grouped queries by instance, time range, agent, department, and skill group.
//
// Description:
//
// - Before invoking this API, ensure that you have [activated the service](https://help.aliyun.com/zh/aiccs/user-guide/activate-aiccs?spm=a2c4g.11186623.0.0.38365923RQDwdH) and [created an instance](https://help.aliyun.com/zh/aiccs/user-guide/create-an-instance?spm=a2c4g.11186623.0.0.8e0b5a2fWNeRUn).
//
// - If you need to specify agent, department, or skill group information, refer to the guidance in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times per second.
//
// - API-wide invocation frequency: 100 times per second.
//
// > Throttling will be triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param tmpReq - GetSkillGroupStatusTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupStatusTotalResponse
func (client *Client) GetSkillGroupStatusTotalWithContext(ctx context.Context, tmpReq *GetSkillGroupStatusTotalRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupStatusTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSkillGroupStatusTotalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DepIds) {
		request.DepIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepIds, dara.String("DepIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupStatusTotal"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupStatusTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancel dual-call.
//
// @param request - HangUpDoubleCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HangUpDoubleCallResponse
func (client *Client) HangUpDoubleCallWithContext(ctx context.Context, request *HangUpDoubleCallRequest, runtime *dara.RuntimeOptions) (_result *HangUpDoubleCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Acid) {
		query["Acid"] = request.Acid
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HangUpDoubleCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HangUpDoubleCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke HangupCall to execute the agent hang-up operation.
//
// @param request - HangupCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HangupCallResponse
func (client *Client) HangupCallWithContext(ctx context.Context, request *HangupCallRequest, runtime *dara.RuntimeOptions) (_result *HangupCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HangupCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HangupCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Hangs up a call in Communication Intelligence Engine.
//
// Description:
//
// - Before you hang up a call in Communication Intelligence Engine, ensure that a call has been initiated by a large model.
//
// - If a call has not been initiated, use the [LlmSmartCall](https://help.aliyun.com/document_detail/2862828.html) or [LlmSmartCallEncrypt](https://help.aliyun.com/document_detail/2881065.html) operation to do so.
//
// @param request - HangupOperateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HangupOperateResponse
func (client *Client) HangupOperateWithContext(ctx context.Context, request *HangupOperateRequest, runtime *dara.RuntimeOptions) (_result *HangupOperateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ImmediateHangup) {
		query["ImmediateHangup"] = request.ImmediateHangup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HangupOperate"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HangupOperateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Hang up a third-party call.
//
// @param request - HangupThirdCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HangupThirdCallResponse
func (client *Client) HangupThirdCallWithContext(ctx context.Context, request *HangupThirdCallRequest, runtime *dara.RuntimeOptions) (_result *HangupThirdCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HangupThirdCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Hold the call.
//
// @param request - HoldCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HoldCallResponse
func (client *Client) HoldCallWithContext(ctx context.Context, request *HoldCallRequest, runtime *dara.RuntimeOptions) (_result *HoldCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HoldCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HoldCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query call details by instance ID.
//
// Description:
//
// Call details include the total number of records, call result, agent name, call time, and other information.
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 80 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency, throttling will be triggered.
//
// @param request - HotlineSessionQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotlineSessionQueryResponse
func (client *Client) HotlineSessionQueryWithContext(ctx context.Context, request *HotlineSessionQueryRequest, runtime *dara.RuntimeOptions) (_result *HotlineSessionQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Acid) {
		query["Acid"] = request.Acid
	}

	if !dara.IsNil(request.AcidList) {
		query["AcidList"] = request.AcidList
	}

	if !dara.IsNil(request.CallResult) {
		query["CallResult"] = request.CallResult
	}

	if !dara.IsNil(request.CallResultList) {
		query["CallResultList"] = request.CallResultList
	}

	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.CallTypeList) {
		query["CallTypeList"] = request.CallTypeList
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledNumberList) {
		query["CalledNumberList"] = request.CalledNumberList
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.CallingNumberList) {
		query["CallingNumberList"] = request.CallingNumberList
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupIdList) {
		query["GroupIdList"] = request.GroupIdList
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.MemberIdList) {
		query["MemberIdList"] = request.MemberIdList
	}

	if !dara.IsNil(request.MemberName) {
		query["MemberName"] = request.MemberName
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.QueryEndTime) {
		query["QueryEndTime"] = request.QueryEndTime
	}

	if !dara.IsNil(request.QueryStartTime) {
		query["QueryStartTime"] = request.QueryStartTime
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.ServicerId) {
		query["ServicerId"] = request.ServicerId
	}

	if !dara.IsNil(request.ServicerIdList) {
		query["ServicerIdList"] = request.ServicerIdList
	}

	if !dara.IsNil(request.ServicerName) {
		query["ServicerName"] = request.ServicerName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotlineSessionQuery"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotlineSessionQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a single phone number to a task.
//
// @param tmpReq - ImportOneTaskPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportOneTaskPhoneNumberResponse
func (client *Client) ImportOneTaskPhoneNumberWithContext(ctx context.Context, tmpReq *ImportOneTaskPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *ImportOneTaskPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportOneTaskPhoneNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Variables) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, dara.String("Variables"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptionType) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.VariablesShrink) {
		query["Variables"] = request.VariablesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportOneTaskPhoneNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportOneTaskPhoneNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports callee data for a call task.
//
// Description:
//
// - You can import callee data by calling this operation. Alternatively, go to the **call task management*	- page, click **Import Callee Data**, download the template, and then upload your file.
//
// - This API operation currently supports only the JSON data type for importing callee data.
//
// - Ensure that you have created a call task before you call this operation.
//
// - To create a call task, go to the **call task management*	- page and click **Create Call Task**, or call the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926796.html) operation.
//
// @param tmpReq - ImportTaskNumberDatasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportTaskNumberDatasResponse
func (client *Client) ImportTaskNumberDatasWithContext(ctx context.Context, tmpReq *ImportTaskNumberDatasRequest, runtime *dara.RuntimeOptions) (_result *ImportTaskNumberDatasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportTaskNumberDatasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PhoneNumberList) {
		request.PhoneNumberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumberList, dara.String("PhoneNumberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EncryptionType) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !dara.IsNil(request.OssFileName) {
		query["OssFileName"] = request.OssFileName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumberListShrink) {
		body["PhoneNumberList"] = request.PhoneNumberListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportTaskNumberDatas"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportTaskNumberDatasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Import outbound call callee numbers based on the instance ID and job ID.
//
// Description:
//
// - After importing outbound call callee numbers, the outbound calling job can operate normally.
//
// - Before invoking this API, we recommend that you confirm the instance ID and job ID. For more information, see [Request Parameters](#api-detail-35).
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 20 times/second.
//
// - API invocation frequency: 50 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param tmpReq - InsertAiOutboundPhoneNumsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertAiOutboundPhoneNumsResponse
func (client *Client) InsertAiOutboundPhoneNumsWithContext(ctx context.Context, tmpReq *InsertAiOutboundPhoneNumsRequest, runtime *dara.RuntimeOptions) (_result *InsertAiOutboundPhoneNumsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertAiOutboundPhoneNumsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Details) {
		request.DetailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Details, dara.String("Details"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchVersion) {
		query["BatchVersion"] = request.BatchVersion
	}

	if !dara.IsNil(request.DetailsShrink) {
		query["Details"] = request.DetailsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertAiOutboundPhoneNums"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertAiOutboundPhoneNumsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - InsertTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertTaskDetailResponse
func (client *Client) InsertTaskDetailWithContext(ctx context.Context, request *InsertTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *InsertTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallInfos) {
		query["CallInfos"] = request.CallInfos
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundTaskId) {
		query["OutboundTaskId"] = request.OutboundTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertTaskDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a third party to the call.
//
// @param request - JoinThirdCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinThirdCallResponse
func (client *Client) JoinThirdCallWithContext(ctx context.Context, request *JoinThirdCallRequest, runtime *dara.RuntimeOptions) (_result *JoinThirdCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.HoldConnectionId) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinThirdCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the agent list in the Cloud Customer Service System by AICCS instance ID and skill group ID.
//
// Description:
//
// - Before invoking this API, confirm the AICCS instance information and skill group information. For guidance on how to obtain these details, refer to the description of [Request Parameters](#api-detail-35).
//
// - If an agent is not assigned to any skill group, you can invoke [GetAgent](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-getagent) to query the agent information.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListAgentBySkillGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentBySkillGroupIdResponse
func (client *Client) ListAgentBySkillGroupIdWithContext(ctx context.Context, request *ListAgentBySkillGroupIdRequest, runtime *dara.RuntimeOptions) (_result *ListAgentBySkillGroupIdResponse, _err error) {
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
		Action:      dara.String("ListAgentBySkillGroupId"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of robot scripts, including robot type, robot name, robot ID, associated business, and industry information.
//
// Description:
//
// - Before invoking this API, ensure that you already have scripts that have passed the Review.
//
// - If you do not have any scripts that have passed the Review, add a script and submit it for Review in the [Script Management](https://aiccs.console.aliyun.com/patter/list) interface first.
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 500 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttle will be triggered.
//
// @param request - ListAiccsRobotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAiccsRobotResponse
func (client *Client) ListAiccsRobotWithContext(ctx context.Context, request *ListAiccsRobotRequest, runtime *dara.RuntimeOptions) (_result *ListAiccsRobotResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RobotName) {
		query["RobotName"] = request.RobotName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAiccsRobot"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAiccsRobotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the system and custom TTS voices available for large model-based outbound calls.
//
// Description:
//
// If you have not created and published any custom voices on the [Custom Voice](https://aiccs.console.aliyun.com/engine/voiceprint) page, the operation returns only system voices.
//
// @param request - ListAvailableTtsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableTtsResponse
func (client *Client) ListAvailableTtsWithContext(ctx context.Context, request *ListAvailableTtsRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableTtsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TtsVoiceCode) {
		query["TtsVoiceCode"] = request.TtsVoiceCode
	}

	if !dara.IsNil(request.VoiceType) {
		query["VoiceType"] = request.VoiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableTts"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableTtsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke ListChatRecordDetail to query online session details by time period.
//
// Description:
//
// This API queries information about completed online sessions within a specified time range, including session content. The query rules are as follows:
//
// - The maximum time span for the query is 1 Day.
//
// - If only the query End Time is provided, the query Start Time is set to 1 hour before the End Time.
//
// - If only the query Start Time is provided, the End Time is set to 1 hour after the Start Time.
//
// - If neither time is provided, the End Time defaults to the current time, and the Start Time is set to 1 hour before the End Time.
//
// - If both times are provided but the time span exceeds 1 Day, an abnormal response is returned.
//
// @param request - ListChatRecordDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatRecordDetailResponse
func (client *Client) ListChatRecordDetailWithContext(ctx context.Context, request *ListChatRecordDetailRequest, runtime *dara.RuntimeOptions) (_result *ListChatRecordDetailResponse, _err error) {
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
		Action:      dara.String("ListChatRecordDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatRecordDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Intelligent Contact Bot conversation records.
//
// Description:
//
// - You can invoke this API to query Intelligent Contact Bot conversation records, or view them in the **Task Management*	- > **Details*	- > **View Conversation Records*	- interface.
//
// - Before invoking this API, ensure that your created Intelligent Contact Bot calling job has successfully connected to at least one phone number.
//
// - If you do not have an existing Intelligent Contact Bot calling job, you can create and start a job in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface, or use the [CreateTask](https://help.aliyun.com/document_detail/2718003.html) and [StartTask](https://help.aliyun.com/document_detail/2718005.html) APIs to create and start a job.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListDialogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDialogResponse
func (client *Client) ListDialogWithContext(ctx context.Context, request *ListDialogRequest, runtime *dara.RuntimeOptions) (_result *ListDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Called) {
		query["Called"] = request.Called
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDialog"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDialogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke ListHotlineRecord to query the hotline recording list by hotline session ID.
//
// @param request - ListHotlineRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotlineRecordResponse
func (client *Client) ListHotlineRecordWithContext(ctx context.Context, request *ListHotlineRecordRequest, runtime *dara.RuntimeOptions) (_result *ListHotlineRecordResponse, _err error) {
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
		Action:      dara.String("ListHotlineRecord"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of hotline details by time period.
//
// Description:
//
// This API queries detailed hotline information within a specified time range, including hotline call information. The query rules are as follows:
//
// - The maximum time span for the query is 1 Day.
//
// - If only the query End Time is provided, the query Start Time is set to 1 hour before the End Time.
//
// - If only the query Start Time is provided, the End Time is set to 1 hour after the Start Time.
//
// - If neither time is provided, the End Time defaults to the current time, and the Start Time is set to 1 hour before the End Time.
//
// - If both times are provided but the time span exceeds 1 Day, an abnormal result is returned.
//
// @param request - ListHotlineRecordDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotlineRecordDetailResponse
func (client *Client) ListHotlineRecordDetailWithContext(ctx context.Context, request *ListHotlineRecordDetailRequest, runtime *dara.RuntimeOptions) (_result *ListHotlineRecordDetailResponse, _err error) {
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
		Action:      dara.String("ListHotlineRecordDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotlineRecordDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the outbound caller phone numbers for a specified agent based on the instance ID and agent account name.
//
// Description:
//
// ### Queries per second (QPS) limit
//
// - Per-user API call frequency: No rate limiting.
//
// - API frequency: 100 calls per second.
//
// > If the total calls from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListOutboundPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOutboundPhoneNumberResponse
func (client *Client) ListOutboundPhoneNumberWithContext(ctx context.Context, request *ListOutboundPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *ListOutboundPhoneNumberResponse, _err error) {
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
		Action:      dara.String("ListOutboundPhoneNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Intelligent Contact Robot call records.
//
// Description:
//
// - Before invoking this API, ensure that you have already initiated an outbound job using the Intelligent Contact Robot.
//
// - If you have not yet initiated an outbound job using the Intelligent Contact Robot, you can invoke the [RobotCall](https://help.aliyun.com/document_detail/2717996.html) API to start one.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListRobotCallDialogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRobotCallDialogResponse
func (client *Client) ListRobotCallDialogWithContext(ctx context.Context, request *ListRobotCallDialogRequest, runtime *dara.RuntimeOptions) (_result *ListRobotCallDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRobotCallDialog"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRobotCallDialogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of robot node information based on the robot ID, which is the script ID.
//
// Description:
//
// - Robot node information includes flow name, model name, node name, whether it is an output, and output ID.
//
// - Before invoking this API, you can confirm the robot ID (that is, the script ID) by following the instructions in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListRobotNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRobotNodeResponse
func (client *Client) ListRobotNodeWithContext(ctx context.Context, request *ListRobotNodeRequest, runtime *dara.RuntimeOptions) (_result *ListRobotNodeResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RobotId) {
		query["RobotId"] = request.RobotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRobotNode"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRobotNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of robot parameters by robot ID, which is also the script ID.
//
// Description:
//
// - Before invoking this API, ensure that you have already added input parameters for the robot.
//
// - If your robot does not have any input parameters, go to the [**Script Management**](https://aiccs.console.aliyun.com/patter/list) > **Configuration*	- > **Input and Output Parameters*	- interface and click to add an input parameter.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttle will be triggered.
//
// @param request - ListRobotParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRobotParamsResponse
func (client *Client) ListRobotParamsWithContext(ctx context.Context, request *ListRobotParamsRequest, runtime *dara.RuntimeOptions) (_result *ListRobotParamsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RobotId) {
		query["RobotId"] = request.RobotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRobotParams"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRobotParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke ListRoles to obtain the list of all roles under a tenant.
//
// @param request - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithContext(ctx context.Context, request *ListRolesRequest, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
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
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query skill groups in the Cloud Customer Service System by instance ID and skill group channel type.
//
// Description:
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupResponse
func (client *Client) ListSkillGroupWithContext(ctx context.Context, request *ListSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupResponse, _err error) {
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
		Action:      dara.String("ListSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Perform a paged query to retrieve the list of Intelligent Contact Robot call jobs. The response includes the total number of jobs and job details.
//
// Description:
//
// - You can use this API to obtain the list of Intelligent Contact Robot call jobs, or retrieve the job list from the [Task Management](https://aiccs.console.aliyun.com/job/list) interface.
//
// - If you have not created any Intelligent Contact Robot call jobs, you can click **Create Job*	- in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface or create a job by invoking the [CreateTask](https://help.aliyun.com/document_detail/2718003.html) API.
//
// - The optional parameters of this API serve as Filter Conditions for the Intelligent Contact Robot call job list. If these parameters are not specified, the API queries all jobs.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 500 calls per second.
//
// > If the combined invocations from multiple users exceed the API frequency limit, throttle will be triggered.
//
// @param request - ListTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskResponse
func (client *Client) ListTaskWithContext(ctx context.Context, request *ListTaskRequest, runtime *dara.RuntimeOptions) (_result *ListTaskResponse, _err error) {
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

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RobotName) {
		query["RobotName"] = request.RobotName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the detail information of Intelligent Contact Robot call jobs.
//
// Description:
//
// - You can invoke this API to obtain the detail information of Intelligent Contact Robot call jobs. Alternatively, you can view this information in the **Detail*	- interface under [**Task Management**](https://aiccs.console.aliyun.com/job/list).
//
// - Before invoking this API, ensure that you have created an Intelligent Contact Robot and successfully started a job.
//
// - If you do not have any successfully created Intelligent Contact Robot call jobs, you can create and start a job in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface, or use the [CreateTask](https://help.aliyun.com/document_detail/2718003.html) and [StartTask](https://help.aliyun.com/document_detail/2718005.html) APIs to create and start a job.
//
// - The optional parameters in this API serve as filter conditions for the detail information of Intelligent Contact Robot call jobs. If these parameters are not specified, all job details will be queried.
//
// ### queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - ListTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskDetailResponse
func (client *Client) ListTaskDetailWithContext(ctx context.Context, request *ListTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *ListTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Called) {
		query["Called"] = request.Called
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StatusCode) {
		query["StatusCode"] = request.StatusCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs full-duplex large language model operations.
//
// Description:
//
// - This operation can be called only during an intelligent outbound call. When you call this operation, set the **CallId*	- request parameter to the unique receipt ID of the ongoing call.
//
// - **CallId*	- is the **CallId*	- parameter returned by the [LlmSmartCallFullDuplex](https://help.aliyun.com/document_detail/2718012.html) operation.
//
// @param request - LlmFullDuplexCallOperateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LlmFullDuplexCallOperateResponse
func (client *Client) LlmFullDuplexCallOperateWithContext(ctx context.Context, request *LlmFullDuplexCallOperateRequest, runtime *dara.RuntimeOptions) (_result *LlmFullDuplexCallOperateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LlmFullDuplexCallOperate"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LlmFullDuplexCallOperateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Intelligent outbound call based on an LLM.
//
// Description:
//
// - Each API invocation supports adding only one called number. If you have multiple called numbers, invoke the API multiple times.
//
// - Before initiating an intelligent call based on an LLM, ensure that you have created an LLM application in the [Application Management](https://aiccs.console.aliyun.com/engine/llmApp) interface and have successfully requested a real number in the Voice Service [Number Management](https://dyvmsnext.console.aliyun.com/number/list/normal) interface.
//
// @param tmpReq - LlmSmartCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LlmSmartCallResponse
func (client *Client) LlmSmartCallWithContext(ctx context.Context, tmpReq *LlmSmartCallRequest, runtime *dara.RuntimeOptions) (_result *LlmSmartCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &LlmSmartCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizParam) {
		request.BizParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizParam, dara.String("BizParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PromptParam) {
		request.PromptParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromptParam, dara.String("PromptParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StartWordParam) {
		request.StartWordParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartWordParam, dara.String("StartWordParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.BizParamShrink) {
		query["BizParam"] = request.BizParamShrink
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallerNumber) {
		query["CallerNumber"] = request.CallerNumber
	}

	if !dara.IsNil(request.CustomerLineCode) {
		query["CustomerLineCode"] = request.CustomerLineCode
	}

	if !dara.IsNil(request.Extension) {
		query["Extension"] = request.Extension
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.PromptParamShrink) {
		query["PromptParam"] = request.PromptParamShrink
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.StartWordParamShrink) {
		query["StartWordParam"] = request.StartWordParamShrink
	}

	if !dara.IsNil(request.TtsSpeed) {
		query["TtsSpeed"] = request.TtsSpeed
	}

	if !dara.IsNil(request.TtsVoiceCode) {
		query["TtsVoiceCode"] = request.TtsVoiceCode
	}

	if !dara.IsNil(request.TtsVolume) {
		query["TtsVolume"] = request.TtsVolume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LlmSmartCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LlmSmartCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a smart call to an encrypted number using a large language model.
//
// @param tmpReq - LlmSmartCallEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LlmSmartCallEncryptResponse
func (client *Client) LlmSmartCallEncryptWithContext(ctx context.Context, tmpReq *LlmSmartCallEncryptRequest, runtime *dara.RuntimeOptions) (_result *LlmSmartCallEncryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &LlmSmartCallEncryptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PromptParam) {
		request.PromptParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromptParam, dara.String("PromptParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StartWordParam) {
		request.StartWordParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartWordParam, dara.String("StartWordParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.CallerNumber) {
		query["CallerNumber"] = request.CallerNumber
	}

	if !dara.IsNil(request.EncryptCalledNumber) {
		query["EncryptCalledNumber"] = request.EncryptCalledNumber
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PromptParamShrink) {
		query["PromptParam"] = request.PromptParamShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartWordParamShrink) {
		query["StartWordParam"] = request.StartWordParamShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LlmSmartCallEncrypt"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LlmSmartCallEncryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a full-duplex Artificial Intelligence Cloud Call Service call with support for personalized configurations.
//
// Description:
//
// ## Operation description
//
// - Before calling this operation, make sure that you have configured the ASR callback URL.
//
// - This operation is available only to users whose UIDs are added to the whitelist.
//
// - The request rate limit for a single user is 100 QPS.
//
// - `CalledNumber` and `CallerNumber` are required parameters that specify the called number and the calling number, respectively.
//
// - Optional parameters such as `StartWordParam`, `TtsVoiceCode`, `TtsSpeed`, and `TtsVolume` allow you to customize the call experience. If these parameters are not specified, the default settings of the application are used.
//
// - The `SessionTimeout` parameter specifies the maximum call duration in seconds. The call is automatically ended when the specified duration is exceeded. For the default value and valid range, refer to the relevant documentation.
//
// - The `OutId` field can be used to track a unique identifier in an external system. The value must be 1 to 32 bytes in length.
//
// @param tmpReq - LlmSmartCallFullDuplexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LlmSmartCallFullDuplexResponse
func (client *Client) LlmSmartCallFullDuplexWithContext(ctx context.Context, tmpReq *LlmSmartCallFullDuplexRequest, runtime *dara.RuntimeOptions) (_result *LlmSmartCallFullDuplexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &LlmSmartCallFullDuplexShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StartWordParam) {
		request.StartWordParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartWordParam, dara.String("StartWordParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallerNumber) {
		query["CallerNumber"] = request.CallerNumber
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.StartWordParamShrink) {
		query["StartWordParam"] = request.StartWordParamShrink
	}

	if !dara.IsNil(request.TtsSpeed) {
		query["TtsSpeed"] = request.TtsSpeed
	}

	if !dara.IsNil(request.TtsVoiceCode) {
		query["TtsVoiceCode"] = request.TtsVoiceCode
	}

	if !dara.IsNil(request.TtsVolume) {
		query["TtsVolume"] = request.TtsVolume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LlmSmartCallFullDuplex"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LlmSmartCallFullDuplexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke MakeCall to initiate a call.
//
// @param request - MakeCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MakeCallResponse
func (client *Client) MakeCallWithContext(ctx context.Context, request *MakeCallRequest, runtime *dara.RuntimeOptions) (_result *MakeCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.CommandCode) {
		query["CommandCode"] = request.CommandCode
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.OuterAccountId) {
		query["OuterAccountId"] = request.OuterAccountId
	}

	if !dara.IsNil(request.OuterAccountType) {
		query["OuterAccountType"] = request.OuterAccountType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MakeCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MakeCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can invoke MakeDoubleCall to initiate a call by using the server-side software development kit (SDK).
//
// @param request - MakeDoubleCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MakeDoubleCallResponse
func (client *Client) MakeDoubleCallWithContext(ctx context.Context, request *MakeDoubleCallRequest, runtime *dara.RuntimeOptions) (_result *MakeDoubleCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.BizData) {
		query["BizData"] = request.BizData
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberPhone) {
		query["MemberPhone"] = request.MemberPhone
	}

	if !dara.IsNil(request.OutboundCallNumber) {
		query["OutboundCallNumber"] = request.OutboundCallNumber
	}

	if !dara.IsNil(request.ServicerPhone) {
		query["ServicerPhone"] = request.ServicerPhone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MakeDoubleCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MakeDoubleCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a paginated list of agents.
//
// @param request - PageQueryAgentListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageQueryAgentListResponse
func (client *Client) PageQueryAgentListWithContext(ctx context.Context, request *PageQueryAgentListRequest, runtime *dara.RuntimeOptions) (_result *PageQueryAgentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageQueryAgentList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageQueryAgentListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a paginated list of agents (agency mode V2)
//
// @param request - PageQueryAgentListNewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageQueryAgentListNewResponse
func (client *Client) PageQueryAgentListNewWithContext(ctx context.Context, request *PageQueryAgentListNewRequest, runtime *dara.RuntimeOptions) (_result *PageQueryAgentListNewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.IsAvailable) {
		query["IsAvailable"] = request.IsAvailable
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageQueryAgentListNew"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageQueryAgentListNewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a paginated query for call task details.
//
// Description:
//
// - You must create a call task before you can query its details.
//
// - You can create a call task on the **call task management*	- page or by calling the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926796.html) API.
//
// @param tmpReq - QueryAiCallDetailPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAiCallDetailPageResponse
func (client *Client) QueryAiCallDetailPageWithContext(ctx context.Context, tmpReq *QueryAiCallDetailPageRequest, runtime *dara.RuntimeOptions) (_result *QueryAiCallDetailPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryAiCallDetailPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DetailIds) {
		request.DetailIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetailIds, dara.String("DetailIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		query["BatchId"] = request.BatchId
	}

	if !dara.IsNil(request.CallResult) {
		query["CallResult"] = request.CallResult
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.DetailIdsShrink) {
		query["DetailIds"] = request.DetailIdsShrink
	}

	if !dara.IsNil(request.EncryptionType) {
		query["EncryptionType"] = request.EncryptionType
	}

	if !dara.IsNil(request.EndCallingTime) {
		query["EndCallingTime"] = request.EndCallingTime
	}

	if !dara.IsNil(request.EndImportedTime) {
		query["EndImportedTime"] = request.EndImportedTime
	}

	if !dara.IsNil(request.MajorIntent) {
		query["MajorIntent"] = request.MajorIntent
	}

	if !dara.IsNil(request.MaxConversationDuration) {
		query["MaxConversationDuration"] = request.MaxConversationDuration
	}

	if !dara.IsNil(request.MinConversationDuration) {
		query["MinConversationDuration"] = request.MinConversationDuration
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartCallingTime) {
		query["StartCallingTime"] = request.StartCallingTime
	}

	if !dara.IsNil(request.StartImportedTime) {
		query["StartImportedTime"] = request.StartImportedTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAiCallDetailPage"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAiCallDetailPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets basic information for a call task.
//
// @param request - QueryAiCallTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAiCallTaskDetailResponse
func (client *Client) QueryAiCallTaskDetailWithContext(ctx context.Context, request *QueryAiCallTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryAiCallTaskDetailResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAiCallTaskDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAiCallTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of call tasks.
//
// @param request - QueryAiCallTaskPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAiCallTaskPageResponse
func (client *Client) QueryAiCallTaskPageWithContext(ctx context.Context, request *QueryAiCallTaskPageRequest, runtime *dara.RuntimeOptions) (_result *QueryAiCallTaskPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAiCallTaskPage"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAiCallTaskPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the details of a communication agent.
//
// @param request - QueryAiVoiceAgentDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAiVoiceAgentDetailResponse
func (client *Client) QueryAiVoiceAgentDetailWithContext(ctx context.Context, request *QueryAiVoiceAgentDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryAiVoiceAgentDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAiVoiceAgentDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAiVoiceAgentDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets details for a specified agent, including its basic, branch, and version information.
//
// Description:
//
// ## Usage notes
//
// - This API retrieves the details of a communication agent.
//
// - If you do not specify `BranchId` and `VersionId`, the API returns the configuration for the latest published version on the effective branch. If the effective branch only contains a draft version, no configuration is returned.
//
// - If you specify only `BranchId`, the API returns the configuration for the latest published version on the specified branch. If the specified branch only contains a draft version, no configuration is returned.
//
// - If you specify both `BranchId` and `VersionId`, the API returns the configuration for the specified version.
//
// - When `BranchDeployStatus` is `1` (branch deployed) and `VersionPublishStatus` is `1` (version published), imported outbound tasks use the configuration of this branch.
//
// @param request - QueryAiVoiceAgentDetailNewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAiVoiceAgentDetailNewResponse
func (client *Client) QueryAiVoiceAgentDetailNewWithContext(ctx context.Context, request *QueryAiVoiceAgentDetailNewRequest, runtime *dara.RuntimeOptions) (_result *QueryAiVoiceAgentDetailNewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.BranchId) {
		query["BranchId"] = request.BranchId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAiVoiceAgentDetailNew"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAiVoiceAgentDetailNewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets call details.
//
// Description:
//
// - This API retrieves call details. You can also view these details on the **call task management*	- > **details*	- > **execution history*	- > **completed*	- > **call details*	- page.
//
// - Before calling this API, make sure you have created a call task and imported called number data.
//
// - You can create a call task and import called number data either on the **call task management*	- page or by using the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926796.html) and [ImportTaskNumberDatas](https://help.aliyun.com/document_detail/2926815.html) APIs.
//
// @param request - QueryConversationDetailInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConversationDetailInfoResponse
func (client *Client) QueryConversationDetailInfoWithContext(ctx context.Context, request *QueryConversationDetailInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryConversationDetailInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		query["BatchId"] = request.BatchId
	}

	if !dara.IsNil(request.DetailId) {
		query["DetailId"] = request.DetailId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConversationDetailInfo"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConversationDetailInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a call task and call records.
//
// Description:
//
// - You can call this operation to query call details. You can also view call details in the **Call Task Management*	- > **Details*	- > **Execution Records*	- > **Completed*	- > **Call Details*	- console.
//
// - Before calling this operation, make sure that you have created a call task and imported callee data.
//
// - If you do not have a created call task, create a call task and import callee data in the **Call Task Management*	- console, or call the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926796.html) and [ImportTaskNumberDatas](https://help.aliyun.com/document_detail/2926815.html) operations to create a call task and import callee data.
//
// @param request - QueryConversationDetailInfoNewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConversationDetailInfoNewResponse
func (client *Client) QueryConversationDetailInfoNewWithContext(ctx context.Context, request *QueryConversationDetailInfoNewRequest, runtime *dara.RuntimeOptions) (_result *QueryConversationDetailInfoNewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.DetailId) {
		query["DetailId"] = request.DetailId
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConversationDetailInfoNew"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConversationDetailInfoNewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke QueryHotlineInQueue to obtain hotline agent data for a skill group by skill group ID.
//
// @param request - QueryHotlineInQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotlineInQueueResponse
func (client *Client) QueryHotlineInQueueWithContext(ctx context.Context, request *QueryHotlineInQueueRequest, runtime *dara.RuntimeOptions) (_result *QueryHotlineInQueueResponse, _err error) {
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
		Action:      dara.String("QueryHotlineInQueue"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryHotlineInQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the hotline number configuration list by instance ID. Fuzzy query by hotline number is supported.
//
// Description:
//
// The hotline number configuration includes information such as the hotline number, number location, carrier, and whether it is used for incoming calls. For details, see [Return Parameters](#api-detail-40).
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: 100 times per second.
//
// - API invocation frequency: 100 times per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param tmpReq - QueryHotlineNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotlineNumberResponse
func (client *Client) QueryHotlineNumberWithContext(ctx context.Context, tmpReq *QueryHotlineNumberRequest, runtime *dara.RuntimeOptions) (_result *QueryHotlineNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryHotlineNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryHotlineNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryHotlineNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询呼入CallId
//
// @param request - QueryInboundCallIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInboundCallIdResponse
func (client *Client) QueryInboundCallIdWithContext(ctx context.Context, request *QueryInboundCallIdRequest, runtime *dara.RuntimeOptions) (_result *QueryInboundCallIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallDate) {
		query["CallDate"] = request.CallDate
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInboundCallId"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInboundCallIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOutboundTaskResponse
func (client *Client) QueryOutboundTaskWithContext(ctx context.Context, request *QueryOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *QueryOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ani) {
		query["Ani"] = request.Ani
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroup) {
		query["SkillGroup"] = request.SkillGroup
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the skill group list of the Cloud Customer Service System by instance ID.
//
// Description:
//
// This API allows you to query information such as the display name, description, channel type, and ID of skill groups. For details, see [Response Parameters](#api-detail-40).
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 1000 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - QuerySkillGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySkillGroupsResponse
func (client *Client) QuerySkillGroupsWithContext(ctx context.Context, request *QuerySkillGroupsRequest, runtime *dara.RuntimeOptions) (_result *QuerySkillGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySkillGroups"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskDetailResponse
func (client *Client) QueryTaskDetailWithContext(ctx context.Context, request *QueryTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ani) {
		query["Ani"] = request.Ani
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DepartmentIdList) {
		query["DepartmentIdList"] = request.DepartmentIdList
	}

	if !dara.IsNil(request.Dnis) {
		query["Dnis"] = request.Dnis
	}

	if !dara.IsNil(request.EndReasonList) {
		query["EndReasonList"] = request.EndReasonList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundTaskId) {
		query["OutboundTaskId"] = request.OutboundTaskId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PriorityList) {
		query["PriorityList"] = request.PriorityList
	}

	if !dara.IsNil(request.ServicerId) {
		query["ServicerId"] = request.ServicerId
	}

	if !dara.IsNil(request.ServicerName) {
		query["ServicerName"] = request.ServicerName
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.SkillGroup) {
		query["SkillGroup"] = request.SkillGroup
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskDetail"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the QueryTickets API to query ticket information.
//
// @param tmpReq - QueryTicketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTicketsResponse
func (client *Client) QueryTicketsWithContext(ctx context.Context, tmpReq *QueryTicketsRequest, runtime *dara.RuntimeOptions) (_result *QueryTicketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CaseId) {
		body["CaseId"] = request.CaseId
	}

	if !dara.IsNil(request.CaseStatus) {
		body["CaseStatus"] = request.CaseStatus
	}

	if !dara.IsNil(request.CaseType) {
		body["CaseType"] = request.CaseType
	}

	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChannelType) {
		body["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DealId) {
		body["DealId"] = request.DealId
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrType) {
		body["SrType"] = request.SrType
	}

	if !dara.IsNil(request.TaskStatus) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TouchId) {
		body["TouchId"] = request.TouchId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTickets"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTicketsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query contact data by instance, list of contact IDs, list of session IDs, and so on.
//
// Description:
//
// This API allows you to query contact data such as contact status, session recipient, agent name, contact channel, satisfaction rating, and evaluation status. For details, see [Response parameters](#api-detail-40).
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No Rate Limiting.
//
// - API frequency: 50 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - QueryTouchListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTouchListResponse
func (client *Client) QueryTouchListWithContext(ctx context.Context, request *QueryTouchListRequest, runtime *dara.RuntimeOptions) (_result *QueryTouchListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChannelType) {
		body["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CloseTimeEnd) {
		body["CloseTimeEnd"] = request.CloseTimeEnd
	}

	if !dara.IsNil(request.CloseTimeStart) {
		body["CloseTimeStart"] = request.CloseTimeStart
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EvaluationLevel) {
		body["EvaluationLevel"] = request.EvaluationLevel
	}

	if !dara.IsNil(request.EvaluationScore) {
		body["EvaluationScore"] = request.EvaluationScore
	}

	if !dara.IsNil(request.EvaluationStatus) {
		body["EvaluationStatus"] = request.EvaluationStatus
	}

	if !dara.IsNil(request.FirstTimeEnd) {
		body["FirstTimeEnd"] = request.FirstTimeEnd
	}

	if !dara.IsNil(request.FirstTimeStart) {
		body["FirstTimeStart"] = request.FirstTimeStart
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberId) {
		body["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.MemberName) {
		body["MemberName"] = request.MemberName
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueId) {
		body["QueueId"] = request.QueueId
	}

	if !dara.IsNil(request.ServicerId) {
		body["ServicerId"] = request.ServicerId
	}

	if !dara.IsNil(request.ServicerName) {
		body["ServicerName"] = request.ServicerName
	}

	if !dara.IsNil(request.TouchId) {
		body["TouchId"] = request.TouchId
	}

	if !dara.IsNil(request.TouchType) {
		body["TouchType"] = request.TouchType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTouchList"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTouchListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove agent accounts from a skill group by specifying the skill group and agent IDs.
//
// Description:
//
// - If you need to provide agent or skill group information, refer to the instructions in the [Request Parameters](#api-detail-35) section.
//
// - You can invoke [ListAgentBySkillGroupId](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-listagentbyskillgroupid) to retrieve agent information under a skill group and verify whether the removal aligns with your expectations.
//
// - If you accidentally remove an agent, you can invoke [UpdateAgent](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-updateagent) to update the agent data and reassign the skill group to the agent.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 50 times/second.
//
// - API-wide invocation frequency: 100 times/second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param tmpReq - RemoveAgentFromSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAgentFromSkillGroupResponse
func (client *Client) RemoveAgentFromSkillGroupWithContext(ctx context.Context, tmpReq *RemoveAgentFromSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveAgentFromSkillGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveAgentFromSkillGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentIdsShrink) {
		query["AgentIds"] = request.AgentIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAgentFromSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAgentFromSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a skill group in the Cloud Customer Service System based on the AICCS instance ID and skill group ID.
//
// Description:
//
// - Deletion is a sensitive operation. Proceed with caution.
//
// - Before deletion, confirm the AICCS instance ID and the skill group ID to be deleted. Refer to the guidance in [Request Parameters](#api-detail-35) for details on how to obtain them.
//
// - After deletion, you can invoke the [QuerySkillGroups](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-queryskillgroups) API to verify the skill group status.
//
// - If you accidentally delete a skill group, you can recreate it by invoking the [CreateSkillGroup](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-createskillgroup) API.
//
// ### Queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 1000 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - RemoveSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSkillGroupResponse
func (client *Client) RemoveSkillGroupWithContext(ctx context.Context, request *RemoveSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveSkillGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		body["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the ResetHotlineNumber API to reset the inbound (IVR flow) and outbound (effective scope) configuration information of a hotline number.
//
// @param tmpReq - ResetHotlineNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetHotlineNumberResponse
func (client *Client) ResetHotlineNumberWithContext(ctx context.Context, tmpReq *ResetHotlineNumberRequest, runtime *dara.RuntimeOptions) (_result *ResetHotlineNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResetHotlineNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OutboundRangeList) {
		request.OutboundRangeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundRangeList, dara.String("OutboundRangeList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableInbound) {
		body["EnableInbound"] = request.EnableInbound
	}

	if !dara.IsNil(request.EnableInboundEvaluation) {
		body["EnableInboundEvaluation"] = request.EnableInboundEvaluation
	}

	if !dara.IsNil(request.EnableOutbound) {
		body["EnableOutbound"] = request.EnableOutbound
	}

	if !dara.IsNil(request.EnableOutboundEvaluation) {
		body["EnableOutboundEvaluation"] = request.EnableOutboundEvaluation
	}

	if !dara.IsNil(request.EvaluationLevel) {
		body["EvaluationLevel"] = request.EvaluationLevel
	}

	if !dara.IsNil(request.HotlineNumber) {
		body["HotlineNumber"] = request.HotlineNumber
	}

	if !dara.IsNil(request.InboundFlowId) {
		body["InboundFlowId"] = request.InboundFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundAllDepart) {
		body["OutboundAllDepart"] = request.OutboundAllDepart
	}

	if !dara.IsNil(request.OutboundRangeListShrink) {
		body["OutboundRangeList"] = request.OutboundRangeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetHotlineNumber"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetHotlineNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RestartOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartOutboundTaskResponse
func (client *Client) RestartOutboundTaskWithContext(ctx context.Context, request *RestartOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *RestartOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundTaskId) {
		query["OutboundTaskId"] = request.OutboundTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate an outbound call using the Intelligent Contact Robot.
//
// Description:
//
// - Before initiating an outbound call using the Intelligent Contact Robot, ensure that you already have a reviewed and approved script in the [Script Management](https://aiccs.console.aliyun.com/patter/list) interface and an approved real number in the Voice Service [Real Number Management](https://dyvmsnext.console.aliyun.com/number/list/normal) interface.
//
// - You can obtain the creation time of the call from the **date*	- parameter in the **Response Header*	- after invoking this API.
//
// > For example, if the **date*	- parameter is: `"date": "Mon, 24 Jun 2024 03:40:31 GMT"`, then the call creation time is: `"2024-06-24 03:40:31"`.
//
// ### queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 1000 calls per second.
//
// > If the total calls from multiple users exceed the API frequency limit, throttle will be triggered.
//
// @param request - RobotCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RobotCallResponse
func (client *Client) RobotCallWithContext(ctx context.Context, request *RobotCallRequest, runtime *dara.RuntimeOptions) (_result *RobotCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.EarlyMediaAsr) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.RecordFlag) {
		query["RecordFlag"] = request.RecordFlag
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RobotId) {
		query["RobotId"] = request.RobotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RobotCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RobotCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate an Intelligent Speech Interaction call based on the callee\\"s caller ID, callee number, and intelligent outbound call audio file.
//
// Description:
//
// - The following characters cannot appear in the Intelligent Speech Interaction SendCcoSmartCall callback: `@ = : "" $ { } ^ 	- ￥`.
//
// - After invoking this API, the **Data*	- field in the response contains the unique receipt ID for this call, which can be used when invoking the [SendCcoSmartCallOperate](https://help.aliyun.com/document_detail/2718013.html) API.
//
// ### Queries Per Second (QPS) Limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - SendCcoSmartCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCcoSmartCallResponse
func (client *Client) SendCcoSmartCallWithContext(ctx context.Context, request *SendCcoSmartCallRequest, runtime *dara.RuntimeOptions) (_result *SendCcoSmartCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionCodeBreak) {
		query["ActionCodeBreak"] = request.ActionCodeBreak
	}

	if !dara.IsNil(request.ActionCodeTimeBreak) {
		query["ActionCodeTimeBreak"] = request.ActionCodeTimeBreak
	}

	if !dara.IsNil(request.AsrAlsAmId) {
		query["AsrAlsAmId"] = request.AsrAlsAmId
	}

	if !dara.IsNil(request.AsrBaseId) {
		query["AsrBaseId"] = request.AsrBaseId
	}

	if !dara.IsNil(request.AsrModelId) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.AsrVocabularyId) {
		query["AsrVocabularyId"] = request.AsrVocabularyId
	}

	if !dara.IsNil(request.BackgroundFileCode) {
		query["BackgroundFileCode"] = request.BackgroundFileCode
	}

	if !dara.IsNil(request.BackgroundSpeed) {
		query["BackgroundSpeed"] = request.BackgroundSpeed
	}

	if !dara.IsNil(request.BackgroundVolume) {
		query["BackgroundVolume"] = request.BackgroundVolume
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.DynamicId) {
		query["DynamicId"] = request.DynamicId
	}

	if !dara.IsNil(request.EarlyMediaAsr) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !dara.IsNil(request.EnableITN) {
		query["EnableITN"] = request.EnableITN
	}

	if !dara.IsNil(request.MuteTime) {
		query["MuteTime"] = request.MuteTime
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PauseTime) {
		query["PauseTime"] = request.PauseTime
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RecordFlag) {
		query["RecordFlag"] = request.RecordFlag
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.TtsConf) {
		query["TtsConf"] = request.TtsConf
	}

	if !dara.IsNil(request.TtsSpeed) {
		query["TtsSpeed"] = request.TtsSpeed
	}

	if !dara.IsNil(request.TtsStyle) {
		query["TtsStyle"] = request.TtsStyle
	}

	if !dara.IsNil(request.TtsVolume) {
		query["TtsVolume"] = request.TtsVolume
	}

	if !dara.IsNil(request.VoiceCode) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.VoiceCodeParam) {
		query["VoiceCodeParam"] = request.VoiceCodeParam
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendCcoSmartCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendCcoSmartCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a specified action during an Intelligent Outbound Call, applicable only to scenarios such as parallel transfer to a human agent or allowing a human agent to listen in on the man-machine dialogue.
//
// Description:
//
// - This API can be successfully invoked only during an ongoing Intelligent Outbound Call. When invoking it, note that the **CallId*	- in the request parameters must be set to the unique receipt ID of the active call.
//
// - The **CallId*	- is the **Data*	- parameter returned when you invoke the [SendCcoSmartCall](https://help.aliyun.com/document_detail/2718012.html) API.
//
// @param request - SendCcoSmartCallOperateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCcoSmartCallOperateResponse
func (client *Client) SendCcoSmartCallOperateWithContext(ctx context.Context, request *SendCcoSmartCallOperateRequest, runtime *dara.RuntimeOptions) (_result *SendCcoSmartCallOperateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendCcoSmartCallOperate"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendCcoSmartCallOperateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Send a hotline heartbeat request based on the instance ID, agent account name, and heartbeat signature.
//
// Description:
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - SendHotlineHeartBeatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendHotlineHeartBeatResponse
func (client *Client) SendHotlineHeartBeatWithContext(ctx context.Context, request *SendHotlineHeartBeatRequest, runtime *dara.RuntimeOptions) (_result *SendHotlineHeartBeatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendHotlineHeartBeat"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a call task.
//
// Description:
//
// - Before you start a call task, ensure its status is Stopped.
//
// - If you do not have a call task, create one on the **Call Task Management*	- page or by calling the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926796.html) operation.
//
// @param request - StartAiCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAiCallTaskResponse
func (client *Client) StartAiCallTaskWithContext(ctx context.Context, request *StartAiCallTaskRequest, runtime *dara.RuntimeOptions) (_result *StartAiCallTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAiCallTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAiCallTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start an outbound calling job by instance ID and job ID.
//
// Description:
//
// - You can invoke this API to start an outbound calling job that is in the paused state.
//
// - An outbound calling job in the stopped state cannot be started again.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 20 times per second.
//
// - API invocation frequency: 20 times per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - StartAiOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAiOutboundTaskResponse
func (client *Client) StartAiOutboundTaskWithContext(ctx context.Context, request *StartAiOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *StartAiOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAiOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAiOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate an outbound call based on the instance ID, agent account name, hotline outbound caller number, and callee number.
//
// Description:
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - StartCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCallResponse
func (client *Client) StartCallWithContext(ctx context.Context, request *StartCallRequest, runtime *dara.RuntimeOptions) (_result *StartCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.Callee) {
		body["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		body["Caller"] = request.Caller
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCall"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke StartCallV2 to initiate an outbound call V2.
//
// @param request - StartCallV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCallV2Response
func (client *Client) StartCallV2WithContext(ctx context.Context, request *StartCallV2Request, runtime *dara.RuntimeOptions) (_result *StartCallV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.Callee) {
		body["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		body["Caller"] = request.Caller
	}

	if !dara.IsNil(request.CallerType) {
		body["CallerType"] = request.CallerType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCallV2"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCallV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke StartChatWork to switch an online agent to the working status.
//
// @param request - StartChatWorkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartChatWorkResponse
func (client *Client) StartChatWorkWithContext(ctx context.Context, request *StartChatWorkRequest, runtime *dara.RuntimeOptions) (_result *StartChatWorkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartChatWork"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartChatWorkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sign in a hotline agent by instance ID and agent account name to start the hotline service.
//
// Description:
//
// - The **Data*	- parameter returned by this API is the token required to initiate a heartbeat.
//
// - If the agent takes a break, you can pause the hotline service by invoking the [SuspendHotlineService](https://help.aliyun.com/document_detail/2718046.html) API.
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - StartHotlineServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartHotlineServiceResponse
func (client *Client) StartHotlineServiceWithContext(ctx context.Context, request *StartHotlineServiceRequest, runtime *dara.RuntimeOptions) (_result *StartHotlineServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartHotlineService"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the StartMicroOutbound API to initiate an outbound call request.
//
// @param request - StartMicroOutboundRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMicroOutboundResponse
func (client *Client) StartMicroOutboundWithContext(ctx context.Context, request *StartMicroOutboundRequest, runtime *dara.RuntimeOptions) (_result *StartMicroOutboundResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.CommandCode) {
		query["CommandCode"] = request.CommandCode
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMicroOutbound"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMicroOutboundResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start a successfully created Intelligent Contact Robot calling job.
//
// Description:
//
// - You can invoke this API to start a successfully created Intelligent Contact Robot calling job, or manually start the job on the [Task Management](https://aiccs.console.aliyun.com/job/list) page.
//
// - Before invoking this API, ensure that you already have a successfully created Intelligent Contact Robot calling job.
//
// - If you do not have a successfully created Intelligent Contact Robot outbound calling job, you can click to create a job on the [Task Management](https://aiccs.console.aliyun.com/job/list) page or create one by using the [CreateTask](https://help.aliyun.com/document_detail/2718003.html) API.
//
// ### queries per second (QPS) Limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - StartTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTaskResponse
func (client *Client) StartTaskWithContext(ctx context.Context, request *StartTaskRequest, runtime *dara.RuntimeOptions) (_result *StartTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartNow) {
		query["StartNow"] = request.StartNow
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an outbound call task.
//
// Description:
//
// - Before you stop an outbound call task, ensure that the task has been created and its status is `running`.
//
// - If you have not created an outbound call task, create one on the Outbound Task Management page or call the [CreateAiCallTask](https://help.aliyun.com/document_detail/2926796.html) operation.
//
// @param request - StopAiCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAiCallTaskResponse
func (client *Client) StopAiCallTaskWithContext(ctx context.Context, request *StopAiCallTaskRequest, runtime *dara.RuntimeOptions) (_result *StopAiCallTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAiCallTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAiCallTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pause an outbound calling job by instance ID and job ID.
//
// Description:
//
// After pausing an outbound calling job, you can invoke the [StartAiOutboundTask](https://help.aliyun.com/document_detail/2718027.html) API to restart it.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 20 times/second.
//
// - API invocation frequency: 20 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - StopAiOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAiOutboundTaskResponse
func (client *Client) StopAiOutboundTaskWithContext(ctx context.Context, request *StopAiOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *StopAiOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAiOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAiOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pause an Intelligent Contact Robot calling job that has been successfully started.
//
// Description:
//
// - You can use this API to pause an Intelligent Contact Robot calling job that has been successfully started, or manually pause the job on the [Task Management](https://aiccs.console.aliyun.com/job/list) interface.
//
// - Before invoking this API, ensure that you already have a successfully started Intelligent Contact Robot calling job.
//
// - If you do not have a successfully started Intelligent Contact Robot calling job, click **Start*	- on the [Task Management](https://aiccs.console.aliyun.com/job/list) interface or start a job by using the [StartTask](https://help.aliyun.com/document_detail/2718005.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 500 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - StopTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskResponse
func (client *Client) StopTaskWithContext(ctx context.Context, request *StopTaskRequest, runtime *dara.RuntimeOptions) (_result *StopTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a recording notes task.
//
// @param request - SubmitAudioNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAudioNoteResponse
func (client *Client) SubmitAudioNoteWithContext(ctx context.Context, request *SubmitAudioNoteRequest, runtime *dara.RuntimeOptions) (_result *SubmitAudioNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.LlmModelId) {
		query["LlmModelId"] = request.LlmModelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAudioNote"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAudioNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pause hotline service when an agent takes a short break.
//
// Description:
//
// If the break ends and you need to resume hotline service, you can invoke the [StartHotlineService](https://help.aliyun.com/document_detail/2718045.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 100 queries per second (QPS).
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - SuspendHotlineServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendHotlineServiceResponse
func (client *Client) SuspendHotlineServiceWithContext(ctx context.Context, request *SuspendHotlineServiceRequest, runtime *dara.RuntimeOptions) (_result *SuspendHotlineServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendHotlineService"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuspendOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendOutboundTaskResponse
func (client *Client) SuspendOutboundTaskWithContext(ctx context.Context, request *SuspendOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *SuspendOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundTaskId) {
		query["OutboundTaskId"] = request.OutboundTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminate an outbound call job by instance ID and job ID.
//
// Description:
//
// - This API supports terminating an outbound call job. Once terminated, the job cannot be restarted.
//
// - To temporarily stop an outbound call job, you can invoke the [StopAiOutboundTask](https://help.aliyun.com/document_detail/2718024.html) API.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 20 times/second.
//
// - API-wide invocation frequency: 20 times/second.
//
// > If the total invocations from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - TerminateAiOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateAiOutboundTaskResponse
func (client *Client) TerminateAiOutboundTaskWithContext(ctx context.Context, request *TerminateAiOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *TerminateAiOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateAiOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateAiOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TestLargeModel
//
// @param tmpReq - TestLargeModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestLargeModelResponse
func (client *Client) TestLargeModelWithContext(ctx context.Context, tmpReq *TestLargeModelRequest, runtime *dara.RuntimeOptions) (_result *TestLargeModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TestLargeModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BaseModel) {
		request.BaseModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseModel, dara.String("BaseModel"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseModelShrink) {
		query["BaseModel"] = request.BaseModelShrink
	}

	if !dara.IsNil(request.ModelCode) {
		query["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserDialogContent) {
		query["UserDialogContent"] = request.UserDialogContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestLargeModel"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestLargeModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the TransferCallToSkillGroup API to execute a single-step or two-step transfer to a skill group.
//
// @param request - TransferCallToSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferCallToSkillGroupResponse
func (client *Client) TransferCallToSkillGroupWithContext(ctx context.Context, request *TransferCallToSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *TransferCallToSkillGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.CallId) {
		body["CallId"] = request.CallId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.HoldConnectionId) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSingleTransfer) {
		body["IsSingleTransfer"] = request.IsSingleTransfer
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.SkillGroupId) {
		body["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferCallToSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify agent information in the Cloud Customer Service System based on the instance ID and agent account name. You can modify the agent\\"s display name and the skill groups to which the agent belongs.
//
// Description:
//
// - Before invoking this API, we recommend that you confirm your Artificial Intelligence Cloud Call Service (AICCS) instance ID. For guidance on how to obtain it, see the description of [Request Parameters](#api-detail-35).
//
// - After the update, you can invoke the [GetAgent](https://help.aliyun.com/document_detail/2717961.html) API to confirm whether the update meets your expectations.
//
// > Currently, only the display name and assigned skill groups can be modified.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: No rate limiting.
//
// - API-wide invocation frequency: 100 queries per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API-wide frequency limit.
//
// @param request - UpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithContext(ctx context.Context, request *UpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		body["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		body["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgent"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("PUT"),
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

// Summary:
//
// Updates a call task configuration.
//
// Description:
//
// Ensure the call task is stopped before you update its configuration.
//
// @param tmpReq - UpdateAiCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAiCallTaskResponse
func (client *Client) UpdateAiCallTaskWithContext(ctx context.Context, tmpReq *UpdateAiCallTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateAiCallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAiCallTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallDay) {
		request.CallDayShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallDay, dara.String("CallDay"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallRetryReason) {
		request.CallRetryReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallRetryReason, dara.String("CallRetryReason"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallTime) {
		request.CallTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTime, dara.String("CallTime"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallDayShrink) {
		query["CallDay"] = request.CallDayShrink
	}

	if !dara.IsNil(request.CallRetryInterval) {
		query["CallRetryInterval"] = request.CallRetryInterval
	}

	if !dara.IsNil(request.CallRetryReasonShrink) {
		query["CallRetryReason"] = request.CallRetryReasonShrink
	}

	if !dara.IsNil(request.CallRetryTimes) {
		query["CallRetryTimes"] = request.CallRetryTimes
	}

	if !dara.IsNil(request.CallTimeShrink) {
		query["CallTime"] = request.CallTimeShrink
	}

	if !dara.IsNil(request.LineEncoding) {
		query["LineEncoding"] = request.LineEncoding
	}

	if !dara.IsNil(request.LinePhoneNum) {
		query["LinePhoneNum"] = request.LinePhoneNum
	}

	if !dara.IsNil(request.MissCallRetry) {
		query["MissCallRetry"] = request.MissCallRetry
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneType) {
		query["PhoneType"] = request.PhoneType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.StartType) {
		query["StartType"] = request.StartType
	}

	if !dara.IsNil(request.TaskCps) {
		query["TaskCps"] = request.TaskCps
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStartTime) {
		query["TaskStartTime"] = request.TaskStartTime
	}

	if !dara.IsNil(request.VirtualNumber) {
		query["VirtualNumber"] = request.VirtualNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAiCallTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAiCallTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update an outbound call job by instance ID and job ID.
//
// Description:
//
// This API supports updating information such as the task name, task description, outbound caller number, and skill group ID. For details, see [Request Parameters](#api-detail-35).
//
// ### Queries per second (QPS) limit
//
// - Per-user invocation frequency: No rate limiting.
//
// - API frequency: 20 times per second.
//
// > Throttling is triggered if the total invocations from multiple users exceed the API frequency limit.
//
// @param tmpReq - UpdateAiOutboundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAiOutboundTaskResponse
func (client *Client) UpdateAiOutboundTaskWithContext(ctx context.Context, tmpReq *UpdateAiOutboundTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateAiOutboundTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAiOutboundTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OutboundNums) {
		request.OutboundNumsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundNums, dara.String("OutboundNums"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecallRule) {
		request.RecallRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecallRule, dara.String("RecallRule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrentRate) {
		query["ConcurrentRate"] = request.ConcurrentRate
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionTime) {
		query["ExecutionTime"] = request.ExecutionTime
	}

	if !dara.IsNil(request.ForecastCallRate) {
		query["ForecastCallRate"] = request.ForecastCallRate
	}

	if !dara.IsNil(request.HandlerId) {
		query["HandlerId"] = request.HandlerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NumRepeated) {
		query["NumRepeated"] = request.NumRepeated
	}

	if !dara.IsNil(request.OutboundNumsShrink) {
		query["OutboundNums"] = request.OutboundNumsShrink
	}

	if !dara.IsNil(request.RecallRuleShrink) {
		query["RecallRule"] = request.RecallRuleShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAiOutboundTask"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAiOutboundTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the department name based on the Artificial Intelligence Cloud Call Service (AICCS) instance ID and department ID.
//
// Description:
//
// - This API supports updating only the department name and does not support updating the department ID.
//
// - Before invoking this API, we recommend that you confirm the AICCS instance ID and department ID. For guidance on how to obtain them, see the instructions in [Request Parameters](#api-detail-35).
//
// - After the update is complete, you can invoke the [GetAllDepartment](https://help.aliyun.com/document_detail/2717975.html) API to confirm whether the department information matches your expectations.
//
// ### Queries per second (QPS) limits
//
// - Per-user invocation frequency: 100 times/second.
//
// - API-wide invocation frequency: 100 times/second.
//
// > If the total invocations from multiple users exceed the API-wide frequency limit, throttling will be triggered.
//
// @param request - UpdateDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDepartmentResponse
func (client *Client) UpdateDepartmentWithContext(ctx context.Context, request *UpdateDepartmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.DepartmentName) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDepartment"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a large language model.
//
// @param tmpReq - UpdateLargeModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLargeModelResponse
func (client *Client) UpdateLargeModelWithContext(ctx context.Context, tmpReq *UpdateLargeModelRequest, runtime *dara.RuntimeOptions) (_result *UpdateLargeModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLargeModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BaseModel) {
		request.BaseModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BaseModel, dara.String("BaseModel"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.BaseModelShrink) {
		query["BaseModel"] = request.BaseModelShrink
	}

	if !dara.IsNil(request.ModelCode) {
		query["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelUrl) {
		query["ModelUrl"] = request.ModelUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Temperature) {
		query["Temperature"] = request.Temperature
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.TopP) {
		query["TopP"] = request.TopP
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLargeModel"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLargeModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Model Application
//
// @param tmpReq - UpdateModelApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelApplicationResponse
func (client *Client) UpdateModelApplicationWithContext(ctx context.Context, tmpReq *UpdateModelApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdateModelApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateModelApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InterruptConfig) {
		request.InterruptConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterruptConfig, dara.String("InterruptConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCode) {
		query["ApplicationCode"] = request.ApplicationCode
	}

	if !dara.IsNil(request.ApplicationCps) {
		query["ApplicationCps"] = request.ApplicationCps
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.CallAssistantHangup) {
		query["CallAssistantHangup"] = request.CallAssistantHangup
	}

	if !dara.IsNil(request.CallAssistantRecognize) {
		query["CallAssistantRecognize"] = request.CallAssistantRecognize
	}

	if !dara.IsNil(request.CallConnectedTriggerModel) {
		query["CallConnectedTriggerModel"] = request.CallConnectedTriggerModel
	}

	if !dara.IsNil(request.DtmfAllowedDigits) {
		query["DtmfAllowedDigits"] = request.DtmfAllowedDigits
	}

	if !dara.IsNil(request.DtmfAutoValidateEnable) {
		query["DtmfAutoValidateEnable"] = request.DtmfAutoValidateEnable
	}

	if !dara.IsNil(request.DtmfDigitCount) {
		query["DtmfDigitCount"] = request.DtmfDigitCount
	}

	if !dara.IsNil(request.DtmfInputTimeout) {
		query["DtmfInputTimeout"] = request.DtmfInputTimeout
	}

	if !dara.IsNil(request.DtmfOutOfRangeAction) {
		query["DtmfOutOfRangeAction"] = request.DtmfOutOfRangeAction
	}

	if !dara.IsNil(request.DtmfRetryPlayTimes) {
		query["DtmfRetryPlayTimes"] = request.DtmfRetryPlayTimes
	}

	if !dara.IsNil(request.DtmfRetryPromptText) {
		query["DtmfRetryPromptText"] = request.DtmfRetryPromptText
	}

	if !dara.IsNil(request.DtmfSendMaxCount) {
		query["DtmfSendMaxCount"] = request.DtmfSendMaxCount
	}

	if !dara.IsNil(request.DtmfSendWaitTimeout) {
		query["DtmfSendWaitTimeout"] = request.DtmfSendWaitTimeout
	}

	if !dara.IsNil(request.DyvmsSceneName) {
		query["DyvmsSceneName"] = request.DyvmsSceneName
	}

	if !dara.IsNil(request.EnableDtmfReceive) {
		query["EnableDtmfReceive"] = request.EnableDtmfReceive
	}

	if !dara.IsNil(request.EnableDtmfSend) {
		query["EnableDtmfSend"] = request.EnableDtmfSend
	}

	if !dara.IsNil(request.EnableMorse) {
		query["EnableMorse"] = request.EnableMorse
	}

	if !dara.IsNil(request.InterruptConfigShrink) {
		query["InterruptConfig"] = request.InterruptConfigShrink
	}

	if !dara.IsNil(request.ModelCode) {
		query["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.MuteActive) {
		query["MuteActive"] = request.MuteActive
	}

	if !dara.IsNil(request.MuteDuration) {
		query["MuteDuration"] = request.MuteDuration
	}

	if !dara.IsNil(request.MuteHangupNum) {
		query["MuteHangupNum"] = request.MuteHangupNum
	}

	if !dara.IsNil(request.MutePushMode) {
		query["MutePushMode"] = request.MutePushMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.QualificationName) {
		query["QualificationName"] = request.QualificationName
	}

	if !dara.IsNil(request.RecordingFile) {
		query["RecordingFile"] = request.RecordingFile
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SpeechContent) {
		query["SpeechContent"] = request.SpeechContent
	}

	if !dara.IsNil(request.SpeechId) {
		query["SpeechId"] = request.SpeechId
	}

	if !dara.IsNil(request.StartWord) {
		query["StartWord"] = request.StartWord
	}

	if !dara.IsNil(request.StartWordType) {
		query["StartWordType"] = request.StartWordType
	}

	if !dara.IsNil(request.TtsConfigShrink) {
		query["TtsConfig"] = request.TtsConfigShrink
	}

	if !dara.IsNil(request.UsageDesc) {
		query["UsageDesc"] = request.UsageDesc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelApplication"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke UpdateOuterAccount to update an external account based on the external account ID.
//
// @param request - UpdateOuterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOuterAccountResponse
func (client *Client) UpdateOuterAccountWithContext(ctx context.Context, request *UpdateOuterAccountRequest, runtime *dara.RuntimeOptions) (_result *UpdateOuterAccountResponse, _err error) {
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
		Action:      dara.String("UpdateOuterAccount"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOuterAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the skill group information in the Cloud Customer Service System based on the AICCS instance ID and skill group ID. This API supports updating the skill group description and display name.
//
// Description:
//
// - This API does not support updating the skill group ID or skill group name.
//
// - After the update, you can invoke the [QuerySkillGroups](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-queryskillgroups) API to query the skill group information.
//
// - If you need to provide the AICCS instance ID and skill group ID, refer to the instructions in the [Request Parameters](#api-detail-35) section.
//
// ### Queries per second (QPS) limit
//
// - Per-user call frequency: No rate limiting.
//
// - API call frequency: 1000 calls per second.
//
// > If the total calls from multiple users exceed the API frequency limit, throttling will be triggered.
//
// @param request - UpdateSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillGroupResponse
func (client *Client) UpdateSkillGroupWithContext(ctx context.Context, request *UpdateSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.SkillGroupName) {
		query["SkillGroupName"] = request.SkillGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillGroup"),
		Version:     dara.String("2019-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
