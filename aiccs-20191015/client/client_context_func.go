// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 新增热线号码
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
// 追加任务明细
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
// 批量取消大模型解决方案Ai外呼明细任务
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
// 删除智能外呼任务
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
// 修改在线客服状态
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
// 创建任务
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
// 创建智能外呼任务（预测式外呼、自动外呼）
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
// 创建智能外呼任务批次
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
// 创建部门信息
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
// 创建外呼任务
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
// 创建坐席并开通sso登录能力
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
// 删除坐席账号
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
// 删除智能外呼任务
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
// 删除部门信息
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
// 删除热线号码
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
// 热线号码加密
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
// 坐席纬度基础状态量
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
// 热线坐席纬度详情汇总
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
// 坐席服务状态量
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
// 坐席纬度统计量
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
// 获取智能外呼任务业务自定义信息
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
// 智能外呼任务配置详情查询
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
// 智能外呼任务执行详情
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
// 智能外呼任务列表查询
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
// 智能外呼任务执行进度
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
// 查询对话内容
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
// 获取通话录音文件
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
// 获取热线配置号码列表
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
// 获取会员信息
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
// 获取技能组分组
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
// 部门纬度坐席状态量
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
// 查询话务动作结果数据
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
// 获取热线聊天记录
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
// 获取热线当前信息
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
// 服务统计量数据
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
// 热线检测获取mcu ip地址
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
// 在线坐席信息
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
// 在线服务总量
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
// 技能组纬度队列信息
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
// 获取录音链接
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
// 获取RtcToken
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
// 部门纬度坐席信息数据
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
// 技能组坐席状态详情
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
// 技能组坐席汇总状态量
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
// 技能组纬度状态量
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
// 技能组纬度服务能力
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
// 技能组服务状态量
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
// 技能组状态总量
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
// 移动端呼叫挂断
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
// 通信智能引擎中主动挂断通话
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
// 语音智能体外呼任务导入单条数据
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
// 导入任务号码数据
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
// 智能外呼任务导入号码
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
// 查询所有机器人列表
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
// 获取tts音色列表
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
// 根据时间段查询在线会话详情，包含会话内容，时间段范围最长不超过1天
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
// 查看对话记录
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
// 根据时间段查询热线详情列表，包含热线通话信息，时间段范围最长不超过1天
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
// 查看对话记录
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
// 查询机器人输出列表
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
// 查询参数列表
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
// 获取租户下的所有角色
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
// 查询任务列表
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
// 任务详情查看通话列表
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
// 高德全双工
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
// 基于大模型的智能外呼
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
// 大模型外呼加密号码接口
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
// 基于大模型的智能外呼
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
// 移动端发起呼叫
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
// 分页查询智能体列表
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
// 分页查询智能体列表（代运营模式V2）
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
// 查询明细记录
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
// 查询任务详情
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
// 查询任务列表
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
// 查询智能体明细
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
// 查询智能体详情（代运营模式V2）
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
// 查询通话详情信息
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
// 查询通话记录接口-新
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
// 查询热线号码配置
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
// 从技能组中移除坐席
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
// 重置热线号码
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
// 启动任务
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
// 启动智能外呼任务
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
// 修改在线客服为上班状态
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
// 立即或定时启动智能外呼任务
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
// 停止任务
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
// 手动暂停智能外呼任务
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
// 暂停外呼任务
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
// 终止智能外呼任务
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
// 更新AI外呼任务配置
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
// 更新智能外呼任务（预测式外呼、自动外呼）
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
// 更新部门信息
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
